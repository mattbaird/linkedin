package linkedin

import (
	"bytes"
	"encoding/json"
	"net/http"
)

/*
Limitations
AD ACCOUNTS
	Limited to 5,000 campaigns (regardless of campaign status.)
	Maximum of 1,000 concurrent campaigns in ACTIVE status at any given time.
	Maximum of 15 active creatives and 85 inactive creatives.
CAMPAIGNS
	Active until it reaches its end time or gets deleted.
	Paused campaigns are considered active until their designated end times.
CREATIVES
	Creatives must match the ad format selected during campaign creation.
	If no Ad format is set, it will be set by the first creative created under that campaign.
	Dynamic, Carousel, and Video Ad Campaigns must have their format set during creation.
*/

const REST_CREATE_CAMPAIGN = "https://api.linkedin.com/v2/adCampaignsV2" //POST

func (lic *LinkedInClient) CreateCampaign(create CreateCampaignRequest) (Campaign, error) {
	lic.logger.Debugf("LinkedInClient CreateCampaign called")
	retval := Campaign{}
	createCampaignJson, err := json.Marshal(create)
	if err != nil {
		return retval, err
	}
	req, err := http.NewRequest(POST, REST_CREATE_CAMPAIGN, bytes.NewBuffer(createCampaignJson))
	if err != nil {
		return retval, err
	}
	_, err = lic.callRestAPI(req, &retval)
	if err != nil {
		return retval, err
	}
	return retval, err
}
