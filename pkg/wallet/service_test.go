package wallet

import (
	"github.com/bakhodur-nazriev/wallet/pkg/types"
	"reflect"
	"testing"
)

func TestService_RegisterAccount_alreadyRegistered(t *testing.T) {

}

func TestService_RegisterAccount_success(t *testing.T) {

}

func TestService_FindAccountById_exist(t *testing.T) {
	accounts := []types.Account{
		{ID: 1, Phone: "9924412789", Balance: 25_000},
		{ID: 2, Phone: "992918658735", Balance: 45_000_00},
		{ID: 3, Phone: "9922213245", Balance: 36_500_00},
		{ID: 4, Phone: "992000254987", Balance: 65_000_00},
	}

	expected := []types.Account{
		{ID: 1, Phone: "9924412789", Balance: 25_000},
	}

	result := 0

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestService_FindAccountById_notFound(t *testing.T) {

}
