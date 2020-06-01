package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gmbanaag/wallet2020/internal/app/model"
	. "github.com/smartystreets/goconvey/convey"
)

var wallet1 *model.Wallet
var wallet2 *model.Wallet
var wallet3 *model.Wallet
var wallet4 *model.Wallet
var wallet5 *model.Wallet
var wallet6 *model.Wallet
var wallet7 *model.Wallet
var wallet8 *model.Wallet

//warmUpWalletcreate multiple wallets
func warmUpWallet() {
	wallet1ID := generateUUID()
	wallet1 = &model.Wallet{}
	wallet1.ID = wallet1ID
	wallet1.Balance = 100
	wallet1.UserID = wallet1ID
	wallet1.CountryISO = "SG"
	wallet1.Currency = "SGD"
	wallet1.CreateTS = time.Now().Unix()
	err := wallet1.Create()
	catch(err, false)

	wallet2ID := generateUUID()
	wallet2 = &model.Wallet{}
	wallet2.ID = wallet2ID
	wallet2.Balance = 100
	wallet2.UserID = wallet2ID
	wallet2.CountryISO = "SG"
	wallet2.Currency = "SGD"
	wallet2.CreateTS = time.Now().Unix()
	err = wallet2.Create()
	catch(err, false)

	wallet3ID := generateUUID()
	wallet3 = &model.Wallet{}
	wallet3.ID = wallet3ID
	wallet3.Balance = 100
	wallet3.UserID = wallet3ID
	wallet3.CountryISO = "SG"
	wallet3.Currency = "USD"
	wallet3.CreateTS = time.Now().Unix()
	err = wallet3.Create()
	catch(err, false)

	wallet4 = &model.Wallet{}
	wallet4ID := generateUUID()
	wallet4.ID = wallet4ID
	wallet4.Balance = 100
	wallet4.UserID = wallet4ID
	wallet4.CountryISO = "SG"
	wallet4.Currency = "USD"
	wallet4.CreateTS = time.Now().Unix()
	err = wallet4.Create()
	catch(err, false)

	wallet5 := &model.Wallet{}
	wallet5ID := generateUUID()
	wallet5.ID = wallet5ID
	wallet5.Balance = 100
	wallet5.UserID = wallet5ID
	wallet5.CountryISO = "SG"
	wallet5.Currency = "SGD"
	wallet5.CreateTS = time.Now().Unix()
	err = wallet5.Create()
	catch(err, false)

	wallet6 := &model.Wallet{}
	wallet6ID := generateUUID()
	wallet6.ID = wallet6ID
	wallet6.Balance = 100
	wallet6.UserID = wallet6ID
	wallet6.CountryISO = "SG"
	wallet6.Currency = "SGD"
	wallet6.CreateTS = time.Now().Unix()
	err = wallet6.Create()
	catch(err, false)

	wallet7 := &model.Wallet{}
	wallet7ID := generateUUID()
	wallet7.ID = wallet7ID
	wallet7.Balance = 100
	wallet7.UserID = wallet7ID
	wallet7.CountryISO = "SG"
	wallet7.Currency = "USD"
	wallet7.CreateTS = time.Now().Unix()
	err = wallet7.Create()
	catch(err, false)

	wallet8 := &model.Wallet{}
	wallet8ID := generateUUID()
	wallet8.ID = wallet8ID
	wallet8.Balance = 100
	wallet8.UserID = wallet8ID
	wallet8.CountryISO = "SG"
	wallet8.Currency = "USD"
	wallet8.CreateTS = time.Now().Unix()
	err = wallet8.Create()
	catch(err, false)
}

//tearDownWallet delete wallets created during testing
func tearDownWallet() {
	wallet1.Delete()
	wallet2.Delete()
	wallet3.Delete()
	wallet4.Delete()
	wallet5.Delete()
	wallet6.Delete()
	wallet7.Delete()
	wallet8.Delete()
}

func TestTransferPost(t *testing.T) {
	warmUpWallet()
	Convey("Process POST /v1/transfer request", t, func() {
		rq := TransferRequest{}
		rq.SourceWalletID = wallet1.ID
		rq.DestinationWalletID = wallet2.ID
		rq.Amount = 20
		rq.Message = "Thank you"

		req, _ := json.Marshal(rq)
		Convey("failed transfer", func() {
			Convey("authorization is missing", func() {
				request, err := http.NewRequest(http.MethodPost, "/v1/transfer", bytes.NewBuffer(req))
				request.Header.Add("Content-Type", "application/json")

				if err != nil {
					t.Fatal(err)
				}

				recorder := httptest.NewRecorder()
				testrouter.ServeHTTP(recorder, request)
				So(recorder.Code, ShouldEqual, http.StatusUnauthorized)
			})

			Convey("authorization token is invalid", func() {
				request, err := http.NewRequest(http.MethodPost, "/v1/transfer", bytes.NewBuffer(req))
				request.Header.Add("Authorization", "Bearer xxx")
				request.Header.Add("Content-Type", "application/json")

				if err != nil {
					t.Fatal(err)
				}

				recorder := httptest.NewRecorder()
				testrouter.ServeHTTP(recorder, request)
				So(recorder.Code, ShouldEqual, http.StatusUnauthorized)
			})
		})
		Convey("successful transfer", func() {
			Convey("authorization token is valid", func() {
				request, err := http.NewRequest(http.MethodPost, "/v1/transfer", bytes.NewBuffer(req))
				request.Header.Add("Authorization", "Bearer validfortestingpurposes")
				request.Header.Add("X-User-ID", wallet1.UserID)
				request.Header.Add("Content-Type", "application/json")

				if err != nil {
					t.Fatal(err)
				}

				recorder := httptest.NewRecorder()
				testrouter.ServeHTTP(recorder, request)

				responseBody := TransferResponse{}

				err = json.NewDecoder(recorder.Body).Decode(&responseBody)
				catch(err, false)

				So(responseBody.Balance, ShouldEqual, 80)
				So(responseBody.TransactionID, ShouldNotBeEmpty)
				So(responseBody.Result, ShouldEqual, model.TxnStatusSuccess)

				So(recorder.Code, ShouldEqual, http.StatusOK)
			})
		})
	})
}

func TestProcessTransferTransaction(t *testing.T) {
	Convey("Process transfer", t, func() {
		Convey("successful transfer", func() {
			transferRequest = TransferRequest{}
			transferRequest.SourceWalletID = wallet3.ID
			transferRequest.DestinationWalletID = wallet4.ID
			transferRequest.Amount = 20
			transferRequest.Message = "Thank you"

			_ = prepareWallets()
			transaction, _ := createTransaction()
			_ = processTransferTransaction(transaction)

			srcWallet := model.Wallet{}
			srcWallet.GetWalletByID(transferRequest.SourceWalletID)

			destWallet := model.Wallet{}
			destWallet.GetWalletByID(transferRequest.DestinationWalletID)

			So(transaction.Status, ShouldEqual, model.TxnStatusSuccess)
			So(srcWallet.Balance, ShouldEqual, sourceWallet.Balance-transferRequest.Amount)
			So(destWallet.Balance, ShouldEqual, destinationWallet.Balance+transferRequest.Amount)
		})

		Convey("failed transfer", func() {
			Convey("insufficient balance", func() {
				transferRequest = TransferRequest{}
				transferRequest.SourceWalletID = wallet3.ID
				transferRequest.DestinationWalletID = wallet4.ID
				transferRequest.Amount = 140
				transferRequest.Message = "Thank you"

				err := prepareWallets()

				So(err, ShouldNotBeNil)

				transaction, _ := createTransaction()

				So(transaction.Status, ShouldEqual, model.TxnStatusPending)
			})
		})
	})
	tearDownWallet()
}
