package accounts

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Account struct {
	Country      string `json:"country"`
	BaseCurrency string `json:"base_currency"`
	BankId       string `json:"bank_id"`
}

func Create(account Account) (string, int) {
	postBody, _ := json.Marshal(account)
	response, err := http.Post("http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(body), http.StatusCreated
}

func Fetch(accountId string) (string, int) {
	response, err := http.Get("http://localhost:8080/v1/organisation/accounts/" + accountId)
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(body), http.StatusOK
}

func Delete() {

}
