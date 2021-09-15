package linkedin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	client *http.Client
	logger *logrus.Logger
}

// NewClient creates a new Linkedin API Client
func NewClient(client *http.Client) LinkedInClient {
	retval := LinkedInClient{}
	// set default logger
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	retval.SetLogger(logger)
	retval.client = client
	return retval
}

// SetLogger associates a logrus instance to the linkedIn api client
func (lic *LinkedInClient) SetLogger(l *logrus.Logger) {
	lic.logger = l
}

// GetTargetingFacets calls the LinkedIn API
func (lic *LinkedInClient) GetTargetingFacets() (TargetingFacetsResponse, error) {
	lic.logger.Debugf("LinkedInClient GetTargetingFacets called")
	retval := TargetingFacetsResponse{}
	req, err := http.NewRequest(POST, REST_TARGETING_FACETS, nil)
	if err != nil {
		return retval, err
	}
	_, err = lic.callRestAPI(req, &retval)
	return retval, err
}

func (lic *LinkedInClient) callRestAPI(req *http.Request, target interface{}) (*http.Response, error) {
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
