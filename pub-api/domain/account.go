package domain

import "github.com/SaravananPitchaimuthu/pub-api/dto"

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountId: a.AccountId}
}

const DbTSLayout = "2006-01-02 15:04:05"

func NewAccount(customerId string, account_type string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		AccountType: account_type,
		Amount:      amount,
		Status:      "1",
		OpeningDate: DbTSLayout,
	}
}
