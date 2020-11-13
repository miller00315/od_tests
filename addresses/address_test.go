// UNITY TEST

package addresses_test

import (
	"testing"
	. "tests-introduction/addresses" // . é  Alias que se refere ao apcote principal
)

type testScene struct {
	insertedAddress string
	expectedAddress string
}

func TestTypeAddress(t *testing.T) {
	t.Parallel()

	testScenes := []testScene{
		{"Rua abc", "Rua"},
		{"Avenida a", "Avenida"},
		{"Rodovia b", "Rodovia"},
		{"Praça a", "Invalid type"},
		{"", "Invalid type"},
	}

	for _, scene := range testScenes {
		receivedAddresTypes := TypeAddress(scene.insertedAddress)

		if receivedAddresTypes != scene.expectedAddress {
			t.Errorf("Tipo recebido diferente do esperado %s", receivedAddresTypes)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Deu muito ruim")
	}
}

// go test ./... -v --cover
// go test --coverprofile testResult.txt
// go tool cover --func=testResult.txt
// go tool cover --html=testResult.txt
