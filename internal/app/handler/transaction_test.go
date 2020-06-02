package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gmbanaag/wallet2020/internal/app/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTransactionGet(t *testing.T) {
	warmUpWallet()
	Convey("Process GET /v1/transactions request", t, func() {
		Convey("successful request", func() {
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

			Convey("authorization token is valid sent transactions", func() {
				request, err := http.NewRequest(http.MethodGet, "/v1/transactions/sent", bytes.NewBuffer([]byte(``)))
				request.Header.Add("Authorization", "Bearer validfortestingpurposes")
				request.Header.Add("X-User-ID", wallet3.UserID)
				request.Header.Add("Content-Type", "application/json")

				if err != nil {
					t.Fatal(err)
				}

				recorder := httptest.NewRecorder()
				testrouter.ServeHTTP(recorder, request)
				So(recorder.Code, ShouldEqual, http.StatusOK)
			})
		})

		Convey("admin requests", func() {
			Convey("authorization token is valid", func() {
				request, err := http.NewRequest(http.MethodGet, "/v1/admin/transactions", bytes.NewBuffer([]byte(``)))
				request.Header.Add("Authorization", "Bearer validfortestingpurposesadmin")
				request.Header.Add("X-User-ID", wallet2.UserID)
				request.Header.Add("Content-Type", "application/json")

				if err != nil {
					t.Fatal(err)
				}

				recorder := httptest.NewRecorder()
				testrouter.ServeHTTP(recorder, request)

				responseBody := []model.Transaction{}

				err = json.NewDecoder(recorder.Body).Decode(&responseBody)
				catch(err, false)

				So(recorder.Code, ShouldEqual, http.StatusOK)
			})
			Convey("authorization token is valid but no admin scope", func() {
				request, err := http.NewRequest(http.MethodGet, "/v1/admin/transactions", bytes.NewBuffer([]byte(``)))
				request.Header.Add("Authorization", "Bearer validfortestingpurposes")
				request.Header.Add("X-User-ID", wallet2.UserID)
				request.Header.Add("Content-Type", "application/json")

				if err != nil {
					t.Fatal(err)
				}

				recorder := httptest.NewRecorder()
				testrouter.ServeHTTP(recorder, request)

				responseBody := []model.Transaction{}

				err = json.NewDecoder(recorder.Body).Decode(&responseBody)
				catch(err, false)

				So(recorder.Code, ShouldEqual, http.StatusUnauthorized)
			})
		})
	})
}
