package accounts

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Account struct {
	Data Data `json:"data"`
}

type Data struct {
	Type           string     `json:"type"`
	Id             string     `json:"id"`
	OrganizationId string     `json:"organization_id"`
	Attributes     Attributes `json:"attributes"`
}

type Attributes struct {
	Name         []string `json:"name"`
	Country      string   `json:"country"`
	BaseCurrency string   `json:"base_currency"`
	BankId       string   `json:"bank_id"`
	BankIdCode   string   `json:"bank_id_code"`
	Bic          string   `json:"bic"`
}

func Create(account Account) (string, int) {
	postBody, _ := json.Marshal(account)
	response, err := http.Post("http://localhost:8080/v1/organisation/accounts", "application/json", bytes.NewBuffer(postBody))
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
