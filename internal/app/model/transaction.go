package model

import "fmt"

//Transaction holds transaction information
type Transaction struct {
	ID                  string  `json:"id" gorm:"primary_key"`
	SourceUserID        string  `json:"source_user_id" gorm:"ForeignKey:id; not null"`
	DestinationUserID   string  `json:"destination_user_id" gorm:"ForeignKey:id: not null"`
	SourceWalletID      string  `json:"source_wallet_id" gorm:"ForeignKey:id; not null"`
	DestinationWalletID string  `json:"destination_wallet_id" gorm:"ForeignKey:id; not null"`
	DeviceID            string  `json:"device_id,omitempty" gorm:"ForeignKey:id; not null"`
	Amount              float64 `json:"amount"`
	Message             string  `json:"message" gorm:"not null"`
	Status              string  `json:"states" gorm:"not null"`
	CreateTS            int64   `json:"create_ts" gorm:"not null"`
	UpdateTS            int64   `json:"update_ts" gorm:"not null"`
}

const (
	//TxnStatusPending has a pending outcome
	TxnStatusPending = "pending"
	//TxnStatusSuccess has a success outcome
	TxnStatusSuccess = "success"
	//TxnStatusFailed has a failed outcome
	TxnStatusFailed = "failed"
)

//CacheKey returns the key mapped to this model
func (Transaction) CacheKey() string {
	return "_txn_"
}

//TableName in the database
func (Transaction) TableName() string {
	return "transactions"
}

//Create transaction
func (t *Transaction) Create() error {
	return DB.Table(t.TableName()).Create(&t).Error
}

//Update transaction
func (t *Transaction) Update() error {
	return DB.Table(t.TableName()).Save(&t).Error
}

//Delete transaction
func (t *Transaction) Delete() error {
	return DB.Table(t.TableName()).Delete(&t).Error
}

//GetTransactionByID get transaction by id
func (t *Transaction) GetTransactionByID(transactionID string) error {
	return DB.Table(t.TableName()).Where("id = ? ", transactionID).Find(&t).Error
}

//GetAllTransactions get transaction by source_user_id
func (t *Transaction) GetAllTransactions() ([]Transaction, error) {
	transactions := []Transaction{}
	err := DB.Table(t.TableName()).Find(&transactions).Error

	if err == nil && len(transactions) == 0 {
		err = fmt.Errorf("no transactions retrieved")
	}

	return transactions, err
}

//GetTransactionBySourceUserID get transaction by source_user_id
func (t *Transaction) GetTransactionBySourceUserID(userID string) ([]Transaction, error) {
	transactions := []Transaction{}
	err := DB.Table(t.TableName()).Where("source_user_id = ? ", userID).Find(&transactions).Error

	if err == nil && len(transactions) == 0 {
		err = fmt.Errorf("no transactions for source user %s", userID)
	}

	return transactions, err
}

//GetTransactionByDestinationUserID get transaction by source_user_id
func (t *Transaction) GetTransactionByDestinationUserID(userID string) ([]Transaction, error) {
	transactions := []Transaction{}
	err := DB.Table(t.TableName()).Where("destination_user_id = ? ", userID).Find(&transactions).Error

	if err == nil && len(transactions) == 0 {
		err = fmt.Errorf("no transactions for destination user %s", userID)
	}

	return transactions, err
}
