package model

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

var txn1 *Transaction
var txn2 *Transaction
var txn3 *Transaction
var sourceID string
var destID string

//warmUpWalletcreate multiple wallets
func warmUpTransactions() {
	sourceID = generateUUID()
	destID = generateUUID()

	txn1 = &Transaction{}
	txn1.ID = generateUUID()
	txn1.CreateTS = time.Now().Unix()
	txn1.Status = TxnStatusPending
	txn1.DestinationWalletID = destID
	txn1.DestinationUserID = destID
	txn1.SourceWalletID = sourceID
	txn1.SourceUserID = sourceID
	txn1.Message = "Thanks you"
	txn1.Amount = 80

	err := txn1.Create()
	catch(err, false)

	txn2 = &Transaction{}
	txn2.ID = generateUUID()
	txn2.CreateTS = time.Now().Unix()
	txn2.Status = TxnStatusPending
	txn2.DestinationWalletID = destID
	txn2.DestinationUserID = destID
	txn2.SourceWalletID = sourceID
	txn2.SourceUserID = sourceID
	txn2.Message = "Thanks you"
	txn2.Amount = 80

	err = txn2.Create()
	catch(err, false)

	txn3 = &Transaction{}
	txn3.ID = generateUUID()
	txn3.CreateTS = time.Now().Unix()
	txn3.Status = TxnStatusPending
	txn3.DestinationWalletID = destID
	txn3.DestinationUserID = destID
	txn3.SourceWalletID = sourceID
	txn3.SourceUserID = sourceID
	txn3.Message = "Thanks you"
	txn3.Amount = 80

	err = txn3.Create()
	catch(err, false)

}

//tearDownTransactions delete transactions created during testing
func tearDownTransactions() {
	txn1.Delete()
	txn2.Delete()
	txn3.Delete()
}

func TestGetTransactions(t *testing.T) {
	initTest()
	warmUpTransactions()
	Convey("Test Get transactions", t, func() {
		Convey("Get transactions by source_user_id", func() {
			transaction := Transaction{}
			transactions, _ := transaction.GetTransactionBySourceUserID(sourceID)

			So(len(transactions), ShouldNotBeZeroValue)
		})
		Convey("Get transactions by dest_user_id", func() {
			transaction := Transaction{}
			transactions, _ := transaction.GetTransactionByDestinationUserID(destID)

			So(len(transactions), ShouldNotBeZeroValue)
		})
		Convey("Get transactions by transactionID", func() {
			transaction := Transaction{}
			transaction.GetTransactionByID(txn1.ID)

			So(transaction.ID, ShouldEqual, txn1.ID)
		})
	})
	tearDownTransactions()
}
