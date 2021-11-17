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
