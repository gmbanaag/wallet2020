package model

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

var wallet1 *Wallet
var wallet2 *Wallet
var wallet3 *Wallet
var wallet4 *Wallet
var wallet5 *Wallet
var wallet6 *Wallet
var wallet7 *Wallet
var wallet8 *Wallet

//warmUpWalletcreate multiple wallets
func warmUpWallet() {
	wallet1ID := generateUUID()
	wallet1 = &Wallet{}
	wallet1.ID = wallet1ID
	wallet1.Balance = 100
	wallet1.UserID = wallet1ID
	wallet1.CountryISO = "SG"
	wallet1.Currency = "SGD"
	wallet1.CreateTS = time.Now().Unix()
	err := wallet1.Create()
	catch(err, false)

	wallet2ID := generateUUID()
	wallet2 = &Wallet{}
	wallet2.ID = wallet2ID
	wallet2.Balance = 100
	wallet2.UserID = wallet2ID
	wallet2.CountryISO = "SG"
	wallet2.Currency = "SGD"
	wallet2.CreateTS = time.Now().Unix()
	err = wallet2.Create()
	catch(err, false)

	wallet3ID := generateUUID()
	wallet3 = &Wallet{}
	wallet3.ID = wallet3ID
	wallet3.Balance = 100
	wallet3.UserID = wallet3ID
	wallet3.CountryISO = "SG"
	wallet3.Currency = "USD"
	wallet3.CreateTS = time.Now().Unix()
	err = wallet3.Create()
	catch(err, false)

	wallet4 = &Wallet{}
	wallet4ID := generateUUID()
	wallet4.ID = wallet4ID
	wallet4.Balance = 100
	wallet4.UserID = wallet4ID
	wallet4.CountryISO = "SG"
	wallet4.Currency = "USD"
	wallet4.CreateTS = time.Now().Unix()
	err = wallet4.Create()
	catch(err, false)

	wallet5 = &Wallet{}
	wallet5ID := generateUUID()
	wallet5.ID = wallet5ID
	wallet5.Balance = 100
	wallet5.UserID = wallet5ID
	wallet5.CountryISO = "SG"
	wallet5.Currency = "SGD"
	wallet5.CreateTS = time.Now().Unix()
	err = wallet5.Create()
	catch(err, false)

	wallet6 = &Wallet{}
	wallet6ID := generateUUID()
	wallet6.ID = wallet6ID
	wallet6.Balance = 100
	wallet6.UserID = wallet6ID
	wallet6.CountryISO = "SG"
	wallet6.Currency = "SGD"
	wallet6.CreateTS = time.Now().Unix()
	err = wallet6.Create()
	catch(err, false)

	wallet7 = &Wallet{}
	wallet7ID := generateUUID()
	wallet7.ID = wallet7ID
	wallet7.Balance = 100
	wallet7.UserID = wallet7ID
	wallet7.CountryISO = "SG"
	wallet7.Currency = "USD"
	wallet7.CreateTS = time.Now().Unix()
	err = wallet7.Create()
	catch(err, false)

	wallet8 = &Wallet{}
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

func TestGetWallet(t *testing.T) {
	warmUpWallet()
	Convey("Test get wallets", t, func() {
		Convey("Get wallets of walletID", func() {
			wallet := Wallet{}
			_ = wallet.GetWalletByID(wallet1.ID)

			So(wallet1.ID, ShouldEqual, wallet1.ID)
		})
		Convey("Get wallets of user", func() {
			wallet := Wallet{}
			wallets := []Wallet{}
			wallets, _ = wallet.GetWalletsByUserID(wallet1.UserID)

			So(len(wallets), ShouldNotBeZeroValue)
		})
	})
	tearDownWallet()
}
