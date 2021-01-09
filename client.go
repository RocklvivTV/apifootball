package apifootball

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

const (
	apiurl     = "https://api-football-v1.p.rapidapi.com/v2/"
	hostHeader = "api-football-v1.p.rapidapi.com"
)

var (
	// Maximum request a day
	requestsDaily int = 100
	// Requests per minut limits
	rpmLimit int = 30
	// Counting requests per day
	countRequestsDaily int
	// Counting requests per minute
	countRequestsMinute int
	// Daily timeout in case if daily requests exceeded
	dailyTimeout time.Duration = 12
	// Timeout in case if requests per minute exceeded
	timeout time.Duration = 5
	err     error
)

// APIClient represents client to work with API-FOOTBALL
type APIClient struct {
	httpClient    http.Client
	apikey        string
	rpm           int
	reqDailyLimit int
	// dailyTimeout represents an hours
	dailyTimeout time.Duration
	// timeout represents an minutes
	timeout time.Duration
}

// NewAPIClient creates new client to communicate with API-FOOTBALL RapidApi
func NewAPIClient(apikey string) *APIClient {
	var h http.Client
	// if len(apikey) < 1 || apikey == "" {
	// 	log.Println(fmt.Errorf("the API Key is empty or not set"))
	// }
	return &APIClient{
		apikey:        apikey,
		httpClient:    h,
		rpm:           rpmLimit,
		reqDailyLimit: requestsDaily,
		dailyTimeout:  dailyTimeout,
		timeout:       timeout,
	}
}

// SetRequestsLimit provides a possibility to update requests rates
func (c *APIClient) SetRequestsLimit(rpm, requestsDaily int) *APIClient {
	c.reqDailyLimit = requestsDaily
	c.rpm = rpm
	return c
}

// GetRequestsLimit returns current values of cRPM, CRPD
func (c *APIClient) GetRequestsLimit() (int, int) {
	return countRequestsDaily, countRequestsMinute
}

// SetTimeoutes provides a possibility to configure timeoutes
func (c *APIClient) SetTimeoutes(dailyTimeout, timeout time.Duration) *APIClient {
	c.timeout = timeout
	c.dailyTimeout = dailyTimeout
	return c
}

// DoRequests prepare and send request to RapidAPI
func (c *APIClient) DoRequests(method, url string, values url.Values) (js *json.Decoder, err error) {
	req := &http.Request{
		Method: method,
		URL:    prepareURL(url, values),
		Header: http.Header{},
	}

	req.Header.Set("x-rapidapi-host", hostHeader)
	req.Header.Set("x-rapidapi-key", c.apikey)

	if countRequestsDaily >= c.reqDailyLimit {
		log.Println(fmt.Sprintf("Exceeded maximum requests per day. Limit: %d, Used: %d", c.reqDailyLimit, countRequestsDaily))
		log.Println(fmt.Sprintf("Timeout for next %d hours", c.dailyTimeout))
		time.Sleep(c.dailyTimeout * time.Hour)
	}

	if countRequestsMinute >= c.rpm {
		log.Println(fmt.Sprintf("Exceeded maximum requests per minute. Limit: %d, Used: %d", c.rpm, countRequestsMinute))
		log.Println(fmt.Sprintf("Timeout for next %d minutes", c.timeout))
		time.Sleep(c.timeout * time.Minute)
		countRequestsMinute = 0
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Println(fmt.Errorf("Failed to make a request. Error: %s", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()
	countRequestsMinute++
	countRequestsDaily++

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Failed request. StatusCode: %d, - Status: %s", resp.StatusCode, resp.Status)
		log.Println(resp.Request.URL)
		return nil, err
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)

	js = json.NewDecoder(buf)
	return js, nil
}

// UpdateRPMCounter provides a possibility to set a proper count of RequestsPerDay and RequestsPerMinute if required
func (c *APIClient) UpdateRPMCounter(cRPD, cRPM int) {
	countRequestsDaily = cRPD
	countRequestsMinute = cRPM
}

// Prepares URL with parameters.
func prepareURL(path string, values url.Values) *url.URL {
	if values == nil {
		values = url.Values{}
	}

	r, err := url.Parse(apiurl)
	if err != nil {
		log.Println(err)
	}

	for k, v := range values {
		for _, vv := range v {
			if reflect.TypeOf(vv).Kind() == reflect.String && vv == "" {
				values.Del(k)
			} else if reflect.TypeOf(vv).Kind() == reflect.Int && len(vv) >= 0 {
				values.Del(k)
			}
		}
	}

	updatedURL := &url.URL{Path: path, RawQuery: values.Encode()}
	return r.ResolveReference(updatedURL)
}
