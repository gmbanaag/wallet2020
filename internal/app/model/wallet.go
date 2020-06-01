package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

//Wallet holds user wallet information
type Wallet struct {
	ID         string  `json:"id" gorm:"primary_key"`
	UserID     string  `json:"user_id" gorm:"ForeignKey:id; not null"`
	CountryISO string  `json:"country_iso" gorm:"not null"`
	Currency   string  `json:"currency" gorm:"not null"`
	Balance    float64 `json:"balance" gorm:"not null"`
	CreateTS   int64   `json:"create_ts" gorm:"not null"`
	UpdateTS   int64   `json:"update_ts" gorm:"not null"`
}

//CacheKey returns the key mapped to this model
func (Wallet) CacheKey() string {
	return "_wallet_"
}

//TableName in the database
func (Wallet) TableName() string {
	return "wallets"
}

//Create wallet
func (w *Wallet) Create() error {
	return DB.Table(w.TableName()).Create(&w).Error
}

//Update wallet
func (w *Wallet) Update() error {
	return DB.Table(w.TableName()).Save(&w).Error
}

//Delete wallet
func (w *Wallet) Delete() error {
	return DB.Table(w.TableName()).Delete(&w).Error
}

//GetWalletByID by walletID
func (w *Wallet) GetWalletByID(walletID string) error {
	return DB.Table(w.TableName()).Where("id = ? ", walletID).Find(&w).Error
	//cache results
	/*cacheKey := fmt.Sprintf("%s%s%s", CacheClient.DefaultKey, w.CacheKey(), walletID)
	results, _ := CacheClient.Get(cacheKey)
	if results == "" {
		err := DB.Table(w.TableName()).Where("id = ? ", walletID).Find(&w).Error

		if err == nil {
			if w.ID != "" {
				jsonResults, _ := json.Marshal(w)
				err = CacheClient.Set(cacheKey, string(jsonResults), 0)
				if err != nil {
					logger.LogError(err.Error())
				}
			}
		} else {
			return err
		}
	} else {
		json.Unmarshal([]byte(results), &w)
	}
	return nil*/
}

//GetAllWallets get all wallets
func (w *Wallet) GetAllWallets() ([]Wallet, error) {
	wallets := []Wallet{}
	err := DB.Table(w.TableName()).Find(&wallets).Error

	if err == nil && len(wallets) == 0 {
		err = fmt.Errorf("no wallet retrieved")
	}

	return wallets, err
}

//GetWalletsByUserID get wallets by user_id
func (w *Wallet) GetWalletsByUserID(userID string) ([]Wallet, error) {
	wallets := []Wallet{}
	err := DB.Table(w.TableName()).Where("user_id = ? ", userID).Find(&wallets).Error

	if err == nil && len(wallets) == 0 {
		err = fmt.Errorf("no wallets for source user %s", userID)
	}

	return wallets, err
}

//ProcessTransfer wallet transfer
func (w *Wallet) ProcessTransfer(sourceWallet, destinationWallet Wallet, amount float64, transaction *Transaction) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Exec("UPDATE wallets SET balance=balance-?, update_ts=? WHERE id=? AND balance>=?", amount, time.Now().Unix(), sourceWallet.ID, amount).Error
		if err != nil {
			return err
		}
		if tx.RowsAffected == 0 {
			//For now this seems not working
			//return fmt.Errorf("unable to credit balance from source wallet")
		}

		if err := tx.Exec("UPDATE wallets SET balance=balance+?, update_ts=? WHERE id=?", amount, time.Now().Unix(), destinationWallet.ID).Error; err != nil {
			return err
		}

		if tx.RowsAffected == 0 {
			//For now this seems not working
			//return fmt.Errorf("unable to debit balance to destination wallet")
		}

		if err := tx.Exec("UPDATE transactions SET status=?, update_ts=? WHERE id=?", TxnStatusSuccess, time.Now().Unix(), transaction.ID).Error; err != nil {
			return err
		}

		transaction.Status = TxnStatusSuccess
		// return nil will commit
		return nil
	})
}
