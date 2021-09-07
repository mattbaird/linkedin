package linkedin

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const REST_GET_AD_ANALYTICS = "GET https://api.linkedin.com/v2/adAnalyticsV2?q=%s&dateRange.start.month=%d&dateRange.start.day=%d&dateRange.start.year=%d&timeGranularity=%s&pivot=%s&campaigns=%s" //GET

//GET https://api.linkedin.com/v2/adAnalyticsV2?
//q=analytics&dateRange.start.month=1&dateRange.start.day=1&
//dateRange.start.year=2016&timeGranularity=MONTHLY&
//pivot=CREATIVE&campaigns=urn:li:sponsoredCampaign:112466001

func makeCampaignUrn(id int) string {
	return fmt.Sprintf("urn:li:sponsoredCampaign:%d", id)
}

func makeArrayOfCampaignUrns(campaigns []int) string {
	valuesText := []string{}
	for i := range campaigns {
		id := campaigns[i]
		text := makeCampaignUrn(id)
		valuesText = append(valuesText, text)
	}
	return strings.Join(valuesText, "+")
}

func (lic *LinkedInClient) GetAdAnalytics(start time.Time, granularity string,
	pivot string, campaigns []int) (AdAnalyticsResponse, error) {
	lic.logger.Debugf("LinkedInClient CreateCampaign called")
	retval := AdAnalyticsResponse{}
	err := lic.checkAndRefresh()
	if err != nil {
		return retval, err
	}
	restUrl := fmt.Sprintf(REST_GET_AD_ANALYTICS, "analytics",
		start.Month(), start.Day(), start.Year(), granularity, pivot, makeArrayOfCampaignUrns(campaigns))
	req, err := http.NewRequest(GET, restUrl, nil)
	if err != nil {
		return retval, err
	}
	_, err = lic.callRestAPI(req, &retval)
	if err != nil {
		return retval, err
	}
	return retval, err
}
