package linkedin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

const REST_CREATE_AD_ACCOUNT = "https://api.linkedin.com/v2/adAccountsV2"                                                                                                                                                                                                                                                                                                                                                                  // POST
const REST_CREATE_CAMPAIGN_GROUP = "https://api.linkedin.com/v2/adCampaignGroupsV2"                                                                                                                                                                                                                                                                                                                                                        //POST
const REST_TARGETING_FACETS = "https://api.linkedin.com/v2/adTargetingFacets"                                                                                                                                                                                                                                                                                                                                                              // GET
const REST_TARGETING_FACET_ENTITIES = "https://api.linkedin.com/v2/adTargetingEntities?q=adTargetingFacet&queryVersion=QUERY_USES_URNS&facet=urn:li:adTargetingFacet:seniorities"                                                                                                                                                                                                                                                          // GET
const REST_AUDIENCE_COUNTS = "https://api.linkedin.com/v2/audienceCountsV2?q=targetingCriteria&target.includedTargetingFacets.locations[0]=urn:li:geo:101174742&target.includedTargetingFacets.locations[1]=urn:li:geo:103644278&target.excludingTargetingFacets.seniorities[0]=urn:li:seniority:3"                                                                                                                                        //GET
const REST_AD_BUDGET_PRICING = "https://api.linkedin.com/v2/adBudgetPricing?account=urn:li:sponsoredAccount:502616245&bidType=CPM&campaignType=TEXT_AD&matchType=EXACT&q=criteria&target.includedTargetingFacets.locations[0]=urn:li:geo:101174742&target.includedTargetingFacets.locations[1]=urn:li:geo:103644278&target.excludingTargetingFacets.seniorities[0]=urn:li:seniority:3&dailyBudget.amount=100&dailyBudget.currencyCode=USD" //GET
const REST_CREATE_CAMPAIGN = "https://api.linkedin.com/v2/adCampaignsV2"                                                                                                                                                                                                                                                                                                                                                                   //POST
const REST_CREATE_AD_CREATIVE = "https://api.linkedin.com/v2/adCreativesV2"                                                                                                                                                                                                                                                                                                                                                                //POST

const REST_GENERATE_ACCESS_TOKEN = "https://www.linkedin.com/oauth/v2/accessToken" //POST

const HEADER_CONTENT_TYPE = "Content-Type"

const POST = "POST"

type LinkedInClient struct {
	clientCredentials string
	clientSecret      string
	client            http.Client
	token             Token
	logger            *logrus.Logger
}

func NewClient(clientCredentials, clientSecret string) (LinkedInClient, error) {
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
		return errors.New("error authenticating")
	}
	lic.logger.Printf("no error no bad status code")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &lic.token); err != nil {
		return err
	}
	lic.logger.Printf("token:%v", lic.token)
	return nil
}

func (lic *LinkedInClient) GetTargetingFacets() TargetingFacetsResponse {
	lic.logger.Debugf("LinkedInClient GetTargetingFacets called")
	return TargetingFacetsResponse{}
}
