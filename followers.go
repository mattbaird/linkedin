package linkedin

import (
	"fmt"
	"net/http"
)

const REST_GET_FOLLOWER_STATISTICS = "https://api.linkedin.com/v2/organizationalEntityFollowerStatistics?q=organizationalEntity&organizationalEntity=%s"

func makeUrn(id int) string {
	return fmt.Sprintf("urn:li:organization:%d", id)
}

func (lic *LinkedInClient) GetFollowerStatistics(orgId int) (FollowerStatistics, error) {
	lic.logger.Debugf("LinkedInClient GetFollowerStatistics called")
	retval := FollowerStatistics{}
	apiUrl := fmt.Sprintf(REST_GET_FOLLOWER_STATISTICS, makeUrn(orgId))
	req, err := http.NewRequest(GET, apiUrl, nil)
	if err != nil {
		return retval, err
	}
	_, err = lic.callRestAPI(req, &retval)
	if err != nil {
		return retval, err
	}
	return retval, err
}
