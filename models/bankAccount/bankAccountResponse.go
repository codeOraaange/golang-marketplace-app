package bankaccount

import "time"

type BankAccountResponse struct {
	AccountID     int       `json:"account_id"`
	UserID        int       `json:"user_id"`
	BankName      string    `json:"bank_name"`
	AccountName   string    `json:"account_name"`
	AccountNumber string    `json:"account_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}