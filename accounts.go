package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
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

func Create(account Account) (*Account, error) {
	postBody, _ := json.Marshal(account)
	response, err := http.Post("http://localhost:8080/v1/organisation/accounts", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var returnedAccount Account
	err2 := json.Unmarshal([]byte(string(body)), &returnedAccount)

	if err2 != nil {
		return nil, err2
	}

	if response.StatusCode != 201 {
		err = errors.New(string(body))
		return nil, err
	}

	return &returnedAccount, err
}

func Fetch(accountId string) (*Account, error) {
	response, err := http.Get("http://localhost:8080/v1/organisation/accounts/" + accountId)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var account Account
	err2 := json.Unmarshal([]byte(string(body)), &account)

	if err2 != nil {
		return nil, err2
	}

	if response.StatusCode != 200 {
		err = errors.New(string(body))
		return nil, err
	}

	return &account, err
}

func Delete(accountId string, version int) error {
	request, err := http.NewRequest(
		"DELETE",
		"http://localhost:8080/v1/organisation/accounts/"+accountId+"?version="+strconv.Itoa(version),
		nil)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	if response.StatusCode == 404 {
		err = errors.New("Account " + accountId + " not found")
		return err
	}

	if response.StatusCode != 204 {
		err = errors.New(string(body))
		return err
	}

	log.Println("Account " + accountId + "deleted at " + string(body))

	return nil
}
