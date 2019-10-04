package models

import "github.com/jinzhu/gorm"

type Wallet struct {
	//"id": 138,
	gorm.Model
	CustomerID     int64 `gorm:"size:255;not null;unique"`
	AvailableBal   float64
	LedgerBal      float64
	AccountType    string
	WalletNo       string
	Currency       string
	Status         string
	DateCreated    string
	DateBalUpdated string
}
