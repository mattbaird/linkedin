package linkedin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const REST_CREATE_AD_ACCOUNT = "https://api.linkedin.com/v2/adAccountsV2" // POST
const REST_FETCH_AD_ACCOUNT = "https://api.linkedin.com/v2/adAccountsV2/%s"

func (lic *LinkedInClient) CreateAdAccount(currency string, name string, notifiedOnCampaignOptimization bool,
	notifiedOnCreativeApproval bool, notifiedOnCreativeRejection bool, notifiedOnEndOfCampaign bool,
	orgId int, accountType string) (string, error) {
	lic.logger.Debugf("LinkedInClient CreateAdAccount called")
	adAccountId := ""
	err := lic.checkAndRefresh()
	if err != nil {
		return adAccountId, err
	}
	account, err := newAdAccountRequest(currency, name, notifiedOnCampaignOptimization,
		notifiedOnCreativeApproval, notifiedOnCreativeRejection, notifiedOnEndOfCampaign,
		orgId, accountType)
	if err != nil {
		return adAccountId, err
	}
	accountJson, err := json.Marshal(account)
	if err != nil {
		return adAccountId, err
	}
	req, err := http.NewRequest(POST, REST_CREATE_AD_ACCOUNT, bytes.NewBuffer(accountJson))
	if err != nil {
		return adAccountId, err
	}
	resp, err := lic.callRestAPI(req, nil)
	if err != nil {
		return adAccountId, err
	}
	// The Ad Account's ID is returned back in the X-LinkedIn-Id response header if creation is successful.
	adAccountId = resp.Header.Get("X-LinkedIn-Id")
	if len(strings.TrimSpace(adAccountId)) == 0 {
		return adAccountId, errors.New("no X-LinkedIn-Id header value returned")
	}
	return adAccountId, err
}

func (lic *LinkedInClient) FetchAdAccount(accountId string) (FetchAdAccountResponse, error) {
	lic.logger.Debugf("LinkedInClient CreateAdAccount called")
	retval := FetchAdAccountResponse{}
	err := lic.checkAndRefresh()
	if err != nil {
		return retval, err
	}
	apiUrl := fmt.Sprintf(REST_FETCH_AD_ACCOUNT, accountId)
	req, err := http.NewRequest(POST, apiUrl, nil)
	if err != nil {
		return retval, err
	}
	_, err = lic.callRestAPI(req, &retval)
	return retval, err
}
