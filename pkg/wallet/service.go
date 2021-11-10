package wallet

import "github.com/bakhodur-nazriev/wallet/pkg/types"

type Service struct {
	nextAccountID int64
	accounts      []types.Account
	payments      []types.Payment
}

// Payment introduce info about payment
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

type Phone string

// Account introduce info about bill of users
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

func (s *Service) RegisterAccount(phone types.Phone) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return
		}
	}

	s.nextAccountID++
	s.accounts = append(s.accounts, types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	})
}

func (s *Service) Deposit(accountID int64, amount types.Money) {
	if amount <= 0 {
		return
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if account.ID == accountID {
			account = &acc
			break
		}
	}

	if account == nil {
		return
	}

	account.Balance += amount
}
