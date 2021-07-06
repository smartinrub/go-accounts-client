package accounts

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	account := Account{
		Data: Data{
			Type:           "accounts",
			Id:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			OrganizationId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: Attributes{
				Name:         []string{"sergiobank"},
				Country:      "GB",
				BaseCurrency: "GBP",
				BankId:       "400300",
				BankIdCode:   "GBDSC",
				Bic:          "NWBKGB22",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/v1/organisation/accounts")
		json, err := json.Marshal(account)
		if err != nil {
			log.Fatalln(err)
		}
		rw.WriteHeader(http.StatusCreated)
		rw.Write(json)
	}))
	defer server.Close()

	api := API{server.Client(), server.URL}
	acc, err := api.Create(account)

	assert.NoError(t, err)
	assert.Equal(t, acc, &account)
}

func TestFetchAccount(t *testing.T) {
	account := Account{
		Data: Data{
			Type:           "accounts",
			Id:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			OrganizationId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: Attributes{
				Name:         []string{"sergiobank"},
				Country:      "GB",
				BaseCurrency: "GBP",
				BankId:       "400300",
				BankIdCode:   "GBDSC",
				Bic:          "NWBKGB22",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
		json, err := json.Marshal(account)
		if err != nil {
			log.Fatalln(err)
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(json)
	}))
	defer server.Close()

	api := API{server.Client(), server.URL}
	acc, err := api.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	assert.NoError(t, err)
	assert.Equal(t, acc, &account)
}

func TestDeleteAccount(t *testing.T) {
	account := Account{
		Data: Data{
			Type:           "accounts",
			Id:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			OrganizationId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: Attributes{
				Name:         []string{"sergiobank"},
				Country:      "GB",
				BaseCurrency: "GBP",
				BankId:       "400300",
				BankIdCode:   "GBDSC",
				Bic:          "NWBKGB22",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc?version=0")
		json, err := json.Marshal(account)
		if err != nil {
			log.Fatalln(err)
		}
		rw.WriteHeader(http.StatusNoContent)
		rw.Write(json)
	}))
	defer server.Close()

	api := API{server.Client(), server.URL}
	err := api.Delete("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)

	assert.NoError(t, err)
}
