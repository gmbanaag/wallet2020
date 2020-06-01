package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gmbanaag/wallet2020/internal/app/logger"
	"github.com/gmbanaag/wallet2020/internal/app/model"
	//"github.com/gmbanaag/wallet2020/internal/app/notification"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

//TransferRequest request payload
type TransferRequest struct {
	SourceWalletID      string  `json:"source_wallet_id"`
	DestinationWalletID string  `json:"destination_wallet_id"`
	Amount              float64 `json:"amount"`
	Message             string  `json:"message"`
}

//TransferResponse response payload
type TransferResponse struct {
	TransactionID string  `json:"transaction_id,omitempty"`
	Result        string  `json:"result,omitempty"`
	Balance       float64 `json:"balance,omitempty"`
	ErrorCode     string  `json:"error_code,omitempty"`
	ErrorMsg      string  `json:"error_msg,omitempty"`
}

var transferRequest TransferRequest
var response Response
var responsePayload TransferResponse
var sourceWallet model.Wallet
var destinationWallet model.Wallet

//Transfer handler for transfer between wallets
func (h *Handler) Transfer(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&transferRequest)
	catch(err, false)

	err = transferRequest.validate(transferRequest)
	if err != nil {
		responsePayload = TransferResponse{ErrorCode: "invalid_request", ErrorMsg: err.Error()}
		response = Response{Code: http.StatusBadRequest, Payload: responsePayload}
	} else {
		reqID = r.Context().Value(middleware.RequestIDKey)
		user := r.Context().Value(h.cfg.UserCtxKey)
		retrieveUser(user)

		err = prepareWallets()
		if err != nil{
			responsePayload = TransferResponse{ErrorCode: "invalid_wallet", ErrorMsg: err.Error()}
			response = Response{Code: http.StatusBadRequest, Payload: responsePayload}
			render.Status(r, response.Code)
			render.DefaultResponder(w, r, response.Payload)
			return
		}

		err = validateTransfer()
		if err != nil {
			responsePayload = TransferResponse{ErrorCode: "invalid_request", ErrorMsg: err.Error()}
			response = Response{Code: http.StatusBadRequest, Payload: responsePayload}
		} else {
			transaction, err := createTransaction()
			if err != nil {
				logger.LogError(fmt.Sprintf("%s [%s]",err.Error(),reqID))

				responsePayload = TransferResponse{ErrorCode: "server_error", ErrorMsg: err.Error()}
				response = Response{Code: http.StatusBadRequest, Payload: responsePayload}
			} else {
				response = processTransferTransaction(transaction)
			}

		}
	}
	render.Status(r, response.Code)
	render.DefaultResponder(w, r, response.Payload)
	return
}

func prepareWallets() error {
	//placeholder for future locking mechanism
	sourceWallet = model.Wallet{}
	err := sourceWallet.GetWalletByID(transferRequest.SourceWalletID)
	if err != nil {
		logger.LogError(fmt.Sprintf("%s [%s]",err.Error(),reqID))
		return err
	}

	if sourceWallet.Balance < transferRequest.Amount {
		emsg := fmt.Errorf("wallet %s has insufficient balance", transferRequest.SourceWalletID)
		return emsg
	}

	destinationWallet = model.Wallet{}
	err = destinationWallet.GetWalletByID(transferRequest.DestinationWalletID)
	if err != nil {
		logger.LogError(fmt.Sprintf("%s [%s]",err.Error(),reqID))
		return err
	}

	return nil
}

func validateTransfer() error {
	//check if user owns the wallet
	if sourceWallet.UserID != userToken.UserID {
		emsg := fmt.Errorf("wallet %s does not belong to user %s", transferRequest.SourceWalletID, userToken.UserID)
		return emsg
	}

	if destinationWallet.UserID == "" {
		emsg := fmt.Errorf("wallet %s does not does not exist", transferRequest.DestinationWalletID)
		return emsg
	}

	if destinationWallet.Currency != sourceWallet.Currency {
		emsg := fmt.Errorf("source %s and destination %s wallets have different currency", transferRequest.SourceWalletID, transferRequest.DestinationWalletID)
		return emsg
	}

	return nil
}

//processTransferTransaction
func processTransferTransaction(transaction *model.Transaction) Response {
	wallet := model.Wallet{}
	err := wallet.ProcessTransfer(sourceWallet, destinationWallet, transferRequest.Amount, transaction)
	if err != nil {
		logger.LogError(fmt.Sprintf("%s [%s]",err.Error(),reqID))

		//make sure that the transaction is marked failed and not pending
		if transaction.Status == model.TxnStatusPending {
			transaction.Status = model.TxnStatusFailed
			err := transaction.Update()
			if err != nil {
				logger.LogError(fmt.Sprintf("%s [%s]",err.Error(),reqID))
			}
		}
		responsePayload = TransferResponse{ErrorCode: "server_error", ErrorMsg: err.Error()}
		response = Response{Code: http.StatusInternalServerError, Payload: responsePayload}
	} else {
		srcWallet := model.Wallet{}
		err := srcWallet.GetWalletByID(transferRequest.SourceWalletID)
		if err != nil {
			logger.LogError(fmt.Sprintf("%s [%s]",err.Error(),reqID))
			//omit balance if cannot be retrieved at this time
			responsePayload = TransferResponse{TransactionID: transaction.ID, Result: transaction.Status}
		} else {
			responsePayload = TransferResponse{TransactionID: transaction.ID, Balance: srcWallet.Balance, Result: transaction.Status}

		}

		/*Send to a notification service to notify the recipient user
		notification := notication.Client{Host: config., APIKey: a.Config.OAuthTokeninfo}
		notifyUser := notification.NotifyUser{}
		notification.SendNotication(notifyUser)
		*/

		response = Response{Code: http.StatusOK, Payload: responsePayload}
	}

	return response
}

func createTransaction() (*model.Transaction, error) {
	transaction := model.Transaction{}
	transaction.ID = generateUUID()
	transaction.CreateTS = time.Now().Unix()
	transaction.Status = model.TxnStatusPending
	transaction.DestinationWalletID = destinationWallet.ID
	transaction.DestinationUserID = destinationWallet.UserID
	transaction.SourceWalletID = sourceWallet.ID
	transaction.SourceUserID = sourceWallet.UserID
	transaction.Message = transferRequest.Message
	transaction.Amount = transferRequest.Amount

	err := transaction.Create()
	if err != nil {
		return &transaction, err
	}

	return &transaction, nil
}

//Validate request payload
func (r TransferRequest) validate(req TransferRequest) error {
	if req.SourceWalletID == "" || req.DestinationWalletID == "" || req.Amount <= 0 {
		return fmt.Errorf("%s", "missing parameters")
	}
	return nil
}