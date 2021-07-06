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

type API struct {
	Client  *http.Client
	baseURL string
}

func (api *API) Create(account Account) (*Account, error) {
	postBody, _ := json.Marshal(account)
	response, err := api.Client.Post(api.baseURL+"/v1/organisation/accounts", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var returnedAccount Account
	err2 := json.Unmarshal([]byte(string(body)), &returnedAccount)

	if err2 != nil {
		return nil, err2
	}

	if response.StatusCode != http.StatusCreated {
		err = errors.New(string(body))
		return nil, err
	}

	return &returnedAccount, err
}

func (api *API) Fetch(accountId string) (*Account, error) {
	response, err := api.Client.Get(api.baseURL + "/v1/organisation/accounts/" + accountId)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var account Account
	err2 := json.Unmarshal([]byte(string(body)), &account)

	if err2 != nil {
		return nil, err2
	}

	if response.StatusCode != http.StatusOK {
		err = errors.New(string(body))
		return nil, err
	}

	return &account, err
}

func (api *API) Delete(accountId string, version int) error {
	request, err := http.NewRequest(
		"DELETE",
		api.baseURL+"/v1/organisation/accounts/"+accountId+"?version="+strconv.Itoa(version),
		nil)

	if err != nil {
		log.Println(err)
		return err
	}

	response, err := api.Client.Do(request)

	if err != nil {
		log.Println(err)
		return err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
		return err
	}

	if response.StatusCode == http.StatusNotFound {
		err = errors.New("Account " + accountId + " not found")
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		err = errors.New(string(body))
		return err
	}

	log.Println("Account " + accountId + "deleted at " + string(body))

	return nil
}
