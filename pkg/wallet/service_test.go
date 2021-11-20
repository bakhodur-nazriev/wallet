package wallet

import (
	"github.com/bakhodur-nazriev/wallet/pkg/types"
	"reflect"
	"testing"
)

func TestService_FindAccountByID_success(t *testing.T) {
	s := Service{
		accounts: []*types.Account{
			{ID: 1, Phone: "+992918123456", Balance: 150_000},
			{ID: 2, Phone: "+992901165432", Balance: 12_000_00},
			{ID: 3, Phone: "+992000224123", Balance: 240_00_00},
		},
	}

	expected := &types.Account{
		ID: 2, Phone: "+992901165432", Balance: 12_000_00,
	}

	result, _ := s.FindAccountByID(2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestService_FindAccountByID_fail(t *testing.T) {
	s := Service{
		accounts: []*types.Account{
			{ID: 1, Phone: "+992931442321", Balance: 15_000},
			{ID: 2, Phone: "+992918114587", Balance: 25_000},
			{ID: 3, Phone: "+992501334455", Balance: 35_000},
		},
	}

	expected := ErrAccountNotFound

	result, _ := s.FindAccountByID(4)

	if reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestService_Reject_success(t *testing.T) {
	// Creating service
	s := &Service{}

	// Register there user
	phone := types.Phone("+992985326598")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("Reject(): can't register account, error = %v", err)
		return
	}

	// Replenish his account
	err = s.Deposit(account.ID, 10_000_00)
	if err != nil {
		t.Errorf("Reject(): can't deposit account, error = %v", err)
		return
	}

	// Make payment to his account
	payment, err := s.Pay(account.ID, 1000_00, "auto")
	if err != nil {
		t.Errorf("Reject(): can't create payment, error = %v", err)
		return
	}

	// Try to cancel payment
	err = s.Reject(payment.ID)
	if err != nil {
		t.Errorf("Reject(): error = %v", err)
		return
	}
}

func TestService_FindPaymentByID_success(t *testing.T) {
	// Creating service
	s := &Service{}

	// Register there user
	phone := types.Phone("+9929876532")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("FindAccountByID(): can't register account, error = %v", err)
		return
	}

	// his account
	err = s.Deposit(account.ID, 10_000_00)
	if err != nil {
		t.Errorf("FindPaymentByID(): can't deposit account, error = %v", err)
		return
	}

	// Make payment to his account
	payment, err := s.Pay(account.ID, 1000_00, "auto")
	if err != nil {
		t.Errorf("FindPaymntByID(): can't create payment, error = %v", err)
		return
	}

	// Try to find payment
	got, err := s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("FindPaymentByID(): error = %v", err)
		return
	}

	// Compare paymentx
	if !reflect.DeepEqual(payment, got) {
		t.Errorf("FindPaymentByID(): wrong payment return = %v", err)
	}
}
