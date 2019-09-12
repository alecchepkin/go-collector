package components

import (
	"collector/modeles"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Proxy struct {
	url string
	log *logrus.Entry
}

func (t *Proxy) SetLog(log *logrus.Entry) *Proxy {
	t.log = log.WithField("comp", "proxy")
	return t
}

//
//func (t *Proxy) SetUrl(url string) {
//	t.url = url
//}

func NewProxy(url string, log *logrus.Entry) *Proxy {
	p := &Proxy{url, log}
	return p
}
func (t *Proxy) GetClientToken(acc *modeles.Cabinet) (token string, err error) {

	timeout, err := time.ParseDuration("60s")
	if err != nil {
		t.log.Panic(err)
	}
	httpClient := &http.Client{Timeout: timeout}
	params := url.Values{}

	if acc.IsAgency {
		params.Set("grant_type", "client_credentials")
		params.Set("client_id", acc.ClientID)
		params.Set("client_secret", acc.ClientSecret)

	} else if acc.ParentID > 0 {
		if acc.Parent == nil {
			t.log.Panic("acc.Parent is nil")
		}
		params.Set("grant_type", "agency_client_credentials")
		params.Set("client_id", acc.Parent.ClientID)
		params.Set("client_secret", acc.Parent.ClientSecret)
		params.Set("agency_client_name", acc.Username)

	} else if acc.IsExternal {
		//todo: Temporarily we use access token received from banner_collector
		t.log.Debug("external acc, use db token")
		token = acc.AccessToken
		return
		/*
			params.Set("grant_type", "refresh_token")
			params.Set("client_id", extClientId)
			params.Set("client_secret", extClientSecret)
			params.Set("refresh_token", acc.RefreshToken)
		*/
	} else {
		err = errors.New("something wrong, undefined type acc")
		return
	}

	apiUrl := t.url + "/api/v2/oauth2/token.json"

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(params.Encode()))

	if err != nil {
		return
	}

	req.Header.Set("Content-Method", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)

	if err != nil {
		return
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.log.Panic("Error read body: " + err.Error())
	}
	if resp.StatusCode != 200 {
		err = errors.New(string(r))
		return
	}
	pResp := struct {
		AccessToken string `json:"access_token"`
	}{}
	err = json.Unmarshal(r, &pResp)
	if err != nil {
		t.log.Panic(err)
	}

	token = pResp.AccessToken
	return
}

func (t *Proxy) UpdateTokens(cabs []*modeles.Cabinet) {
	l := len(cabs)
	t.log.Info("tokens updating cabs:", l)
	res := make(chan int)
	for i, acc := range cabs {
		t.log.Debug("token routine:", i+1)
		go func(cab *modeles.Cabinet) {
			token, err := t.GetClientToken(cab)
			if err != nil {
				t.log.Errorln(err, "Cabinet Id:", cab.CabinetID)
			}
			cab.AccessToken = token
			res <- 1
		}(acc)

	}

	for i := 0; i < l; i++ {
		<-res
	}

}
