package redmine

import (
	"errors"
	"strings"
	"time"

	"github.com/goccy/go-json"
)

type Account struct {
	Id          int       `json:"id"`
	Login       string    `json:"login"`
	Admin       bool      `json:"admin"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Mail        string    `json:"mail"`
	CreatedOn   time.Time `json:"created_on"`
	LastLoginOn time.Time `json:"last_login_on"`
	ApiKey      string    `json:"api_key"`
}

type accountRequest struct {
	Account Account `json:"user"`
}

//goland:noinspection GoUnhandledErrorResult
func (c *Client) MyAccount() (*Account, error) {
	url := c.endpoint + "/my/account.json?key=" + c.apikey
	res, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, errors.New("not found")
	}

	decoder := json.NewDecoder(res.Body)
	var r accountRequest
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}

	return &r.Account, nil
}
