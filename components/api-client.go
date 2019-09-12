package components

import (
	"collector/modeles"
	"collector/responses"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ApiClient struct {
	host string
	log  *logrus.Entry
}

func NewApiClient(host string, log *logrus.Entry) *ApiClient {
	t := &ApiClient{}
	t.host = host
	if log == nil {
		panic("log is nil")
	}
	t.log = log.WithField("comp", "api-client")
	return t
}

type params struct {
	method  string
	path    string
	query   url.Values
	body    url.Values
	token   string
	respRaw bool
}

func (t *ApiClient) SetHost(host string) *ApiClient {
	t.host = host
	return t
}

func (t *ApiClient) AgencyClients(cab *modeles.Cabinet) (resp responses.AgencyClientsResp, err error) {

	err = t.doRequest(&resp, params{
		method: "GET",
		path:   "/api/v2/agency/clients.json",
		query:  url.Values{},
		body:   url.Values{},
		token:  cab.AccessToken,
	})

	return
}

func (t *ApiClient) User(cab *modeles.Cabinet) (user responses.User, err error) {
	query := url.Values{}
	query.Set("fields", "id,username,language,firstname,lastname,email,types,status,additional_info,available_mailings,mailings,account,agency,agency_username,branch_username,max_active_banners,active_banners,partner,info_currency,currency,dmp,branch,email_settings,notifications,additional_emails")

	err = t.doRequest(&user, params{
		method:  "POST",
		path:    "/api/v2/user.json",
		query:   query,
		body:    url.Values{},
		token:   cab.AccessToken,
		respRaw: true,
	})
	return
}

func (t *ApiClient) Campaigns(cab *modeles.Cabinet) (resp responses.CampaignsResp, err error) {
	query := url.Values{}
	query.Set("fields", "id,name,status,issues,objective,price,targetings,package_id,created,updated")
	err = t.doRequest(&resp, params{
		method: "GET",
		path:   "/api/v2/campaigns.json",
		query:  query,
		body:   url.Values{},
		token:  cab.AccessToken,
	})
	return
}

func (t *ApiClient) Packages(cab *modeles.Cabinet) (resp responses.PackagesResp, err error) {
	query := url.Values{}
	err = t.doRequest(&resp, params{
		method: "GET",
		path:   "/api/v2/packages.json",
		query:  query,
		body:   url.Values{},
		token:  cab.AccessToken,
	})
	return
}
func (t *ApiClient) Banners(cab *modeles.Cabinet) (resp responses.BannersResp, err error) {
	query := url.Values{}
	err = t.doRequest(&resp, params{
		method: "GET",
		path:   "/api/v2/banners.json",
		query:  query,
		body:   url.Values{},
		token:  cab.AccessToken,
	})
	return
}

func (t *ApiClient) StatBanners(cab *modeles.Cabinet, bannerIds []int64, dateFrom time.Time, dateTo time.Time) (resp responses.StatBannersDayResp, err error) {

	query := url.Values{}

	ids := func() (s string) {
		l := len(bannerIds)
		for key, id := range bannerIds {
			s += fmt.Sprintf("%d", id)
			if key+1 != l {
				s += ","
			}
		}
		return
	}()
	query.Set("id", ids)
	query.Set("date_from", dateFrom.Format("2006-01-02"))
	query.Set("date_to", dateTo.Format("2006-01-02"))
	err = t.doRequest(&resp, params{
		method: "GET",
		path:   "/api/v2/statistics/banners/day.json",
		query:  query,
		body:   url.Values{},
		token:  cab.AccessToken,
	})

	return
}

func (t *ApiClient) Regions(cab *modeles.Cabinet) (resp responses.RegionsResp, err error) {
	query := url.Values{}
	err = t.doRequest(&resp, params{
		method: "GET",
		path:   "/api/v2/regions.json",
		query:  query,
		body:   url.Values{},
		token:  cab.AccessToken,
	})
	return
}

func (t *ApiClient) doRequest(resp interface{}, params params) (err error) {
	timeout, err := time.ParseDuration("60s")
	if err != nil {
		return
	}
	httpClient := &http.Client{Timeout: timeout}

	apiUrl := t.host + params.path
	query := params.query.Encode()
	if len(query) > 0 {
		apiUrl += "?" + query
	}
	req, err := http.NewRequest(params.method, apiUrl, strings.NewReader(params.body.Encode()))
	if err != nil {
		err = t.errRequest(err, req, []byte{})
		return
	}
	req.Header.Set("Content-Method", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+params.token)
	clientResp, err := httpClient.Do(req)
	if err != nil {
		err = t.errRequest(err, req, []byte{})
		return
	}
	respBytes, err := ioutil.ReadAll(clientResp.Body)
	if err != nil {
		t.log.Debug(FormatRequest.Format(FormatRequest{}, req))
		return
	}
	if clientResp.StatusCode != 200 {
		err = t.errRequest(errors.New("status code "+clientResp.Status), req, respBytes)
		return
	}
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		err = t.errRequest(err, req, respBytes)
		return
	}
	return
}

func (t ApiClient) errRequest(err error, req *http.Request, resp []byte) error {
	r := FormatRequest.Format(FormatRequest{}, req)
	e := errors.New(err.Error() + "\n" + "req:" + r + "\n" + "resp:" + string(resp))
	return e
}
