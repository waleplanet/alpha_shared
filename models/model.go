package models

import "github.com/jinzhu/gorm"

type Wallet struct {
	//"id": 138,
	gorm.Model
	CustomerID     int64   `json:"customer_id" gorm:"size:255;not null;unique"`
	AvailableBal   float64 `json:"available_balance"`
	LedgerBal      float64 `json:"ledger_balance"`
	AccountType    string  `gorm:"type:varchar(1);not null" json:"account_type"`
	WalletID       int64   `gorm:"size:255;not null" json:"id"`
	WalletNo       int64   `gorm:"size:255;not null" json:"wallet_number"`
	Currency       string  `gorm:"type:varchar(20);not null" json:"currency"`
	Status         string  `json:"status"`
	DateCreated    string  `json:"date_created"`
	DateBalUpdated string  `json:"date_balance_updated"`
}
