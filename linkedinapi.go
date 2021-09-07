package linkedin

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const REST_TARGETING_FACETS = "https://api.linkedin.com/v2/adTargetingFacets"                                                                                                                                                                                                                                                                                                                                                              // GET
const REST_TARGETING_FACET_ENTITIES = "https://api.linkedin.com/v2/adTargetingEntities?q=adTargetingFacet&queryVersion=QUERY_USES_URNS&facet=urn:li:adTargetingFacet:seniorities"                                                                                                                                                                                                                                                          // GET
const REST_AUDIENCE_COUNTS = "https://api.linkedin.com/v2/audienceCountsV2?q=targetingCriteria&target.includedTargetingFacets.locations[0]=urn:li:geo:101174742&target.includedTargetingFacets.locations[1]=urn:li:geo:103644278&target.excludingTargetingFacets.seniorities[0]=urn:li:seniority:3"                                                                                                                                        //GET
const REST_AD_BUDGET_PRICING = "https://api.linkedin.com/v2/adBudgetPricing?account=urn:li:sponsoredAccount:502616245&bidType=CPM&campaignType=TEXT_AD&matchType=EXACT&q=criteria&target.includedTargetingFacets.locations[0]=urn:li:geo:101174742&target.includedTargetingFacets.locations[1]=urn:li:geo:103644278&target.excludingTargetingFacets.seniorities[0]=urn:li:seniority:3&dailyBudget.amount=100&dailyBudget.currencyCode=USD" //GET
const REST_CREATE_AD_CREATIVE = "https://api.linkedin.com/v2/adCreativesV2"                                                                                                                                                                                                                                                                                                                                                                //POST

const REST_GENERATE_ACCESS_TOKEN = "https://www.linkedin.com/oauth/v2/accessToken" //POST

const HEADER_CONTENT_TYPE = "Content-Type"
const APPLICATION_JSON = "application/json"

const HEADER_AUTHORIZATION = "Authorization"

const POST = "POST"
const GET = "GET"

type LinkedInClient struct {
	clientCredentials string
	clientSecret      string
	client            *http.Client
	token             Token
	tokenAcquiredAt   int64
	logger            *logrus.Logger
}

// NewClient creates a new Linkedin API Client
func NewClient(clientCredentials, clientSecret string) (LinkedInClient, error) {
	var (
		conn *tls.Conn
		err  error
	)
	if len(strings.TrimSpace(clientSecret)) == 0 {
		return LinkedInClient{}, errors.New("missing client credentials")
	}

	if len(strings.TrimSpace(clientCredentials)) == 0 {
		return LinkedInClient{}, errors.New("missing client credentials")
	}
	retval := LinkedInClient{clientCredentials: clientCredentials, clientSecret: clientSecret}
	// set default logger
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	retval.SetLogger(logger)
	tlsConfig := http.DefaultTransport.(*http.Transport).TLSClientConfig

	retval.client = &http.Client{
		Transport: &http.Transport{
			DialTLS: func(network, addr string) (net.Conn, error) {
				conn, err = tls.Dial(network, addr, tlsConfig)
				return conn, err
			},
		},
	}

	return retval, nil
}

// SetLogger associates a logrus instance to the linkedIn api client
func (lic *LinkedInClient) SetLogger(l *logrus.Logger) {
	lic.logger = l
}

// Authenticate generates an access token, issue a HTTP POST against accessToken
// with your Client ID and Client Secret values
func (lic *LinkedInClient) Authenticate() error {
	lic.logger.Debugf("LinkedInClient Authenticate called")
	formValues := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", lic.clientCredentials, lic.clientSecret)
	req, err := http.NewRequest(POST, REST_GENERATE_ACCESS_TOKEN, strings.NewReader(formValues))
	if err != nil {
		return err
	}
	req.Header.Set(HEADER_CONTENT_TYPE, "application/x-www-form-urlencoded")
	resp, err := lic.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("error authenticating LinkedIn:%v", resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &lic.token); err != nil {
		return err
	}
	lic.logger.Printf("token:%v", lic.token)
	lic.tokenAcquiredAt = time.Now().UnixMilli()
	return nil
}

// GetTargetingFacets calls the LinkedIn API
func (lic *LinkedInClient) GetTargetingFacets() (TargetingFacetsResponse, error) {
	lic.logger.Debugf("LinkedInClient GetTargetingFacets called")
	retval := TargetingFacetsResponse{}
	err := lic.checkAndRefresh()
	if err != nil {
		return retval, err
	}
	req, err := http.NewRequest(POST, REST_TARGETING_FACETS, nil)
	if err != nil {
		return retval, err
	}
	_, err = lic.callRestAPI(req, &retval)
	return retval, err
}

func (lic *LinkedInClient) callRestAPI(req *http.Request, target interface{}) (*http.Response, error) {
	req.Header.Set(HEADER_AUTHORIZATION, fmt.Sprintf("Bearer %s", lic.token.AccessToken))
	if len(req.Header.Get(HEADER_CONTENT_TYPE)) == 0 {
		req.Header.Set(HEADER_CONTENT_TYPE, APPLICATION_JSON)
	}
	resp, err := lic.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("error during API Call [%v]:%v", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &target); err != nil {
		return nil, err
	}
	return resp, nil
}

func (lic *LinkedInClient) shouldRefresh() bool {
	expiresInInt, err := strconv.ParseInt(lic.token.ExpiresIn, 10, 64)
	if err != nil {
		lic.logger.Printf("error converting lic.token.ExpiresIn to number:%v", err)
		return false
	}
	// check with a buffer
	// expiresInInt is in seconds, so convert to millis
	if lic.tokenAcquiredAt+(expiresInInt*1000) > time.Now().UnixMilli()+2000 {
		return true
	}
	return false
}

func (lic *LinkedInClient) checkAndRefresh() error {
	if lic.shouldRefresh() {
		return lic.Authenticate()
	}
	return nil
}
