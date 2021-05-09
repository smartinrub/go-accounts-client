package accounts

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Account struct {
	Data Data `json:"data"`
}

type Data struct {
	Type           string     `json:"type"`
	Id             string     `json:"id"`
	OrganizationId string     `json:"organisation_id"`
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

func Create(account Account) (*Account, int) {
	postBody, _ := json.Marshal(account)
	response, err := http.Post("http://localhost:8080/v1/organisation/accounts", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalln(err)
		return nil, http.StatusInternalServerError
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var returnedAccount Account
	json.Unmarshal([]byte(string(body)), &returnedAccount)

	return &returnedAccount, http.StatusCreated
}

func Fetch(accountId string) (*Account, int) {
	response, err := http.Get("http://localhost:8080/v1/organisation/accounts/" + accountId)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
		return nil, http.StatusInternalServerError
	}

	var account Account
	json.Unmarshal([]byte(string(body)), &account)

	return &account, http.StatusOK
}

func Delete(accountId string, version int) int {
	request, err := http.NewRequest(
		"DELETE",
		"http://localhost:8080/v1/organisation/accounts/"+accountId+"?version="+strconv.Itoa(version),
		nil)

	if err != nil {
		log.Fatalln(err)
		return http.StatusInternalServerError
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
		return http.StatusInternalServerError
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
		return http.StatusInternalServerError
	}

	log.Print(string(body))

	return http.StatusNoContent
}
