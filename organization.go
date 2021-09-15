package linkedin

import (
	"fmt"
	"net/http"
)

const REST_LOOKUP_ORGANIZATION_BY_ID = "https://api.linkedin.com/v2/organizations/%d"                                  // GET {organization ID}
const REST_LOOKUP_ORGANIZATION_BY_VANITY_NAME = "https://api.linkedin.com/v2/organizations?q=vanityName&vanityName=%s" //GET

func (lic *LinkedInClient) LookUpOrganizationByOrgId(orgId int) (Organization, error) {
	lic.logger.Debugf("LinkedInClient LookUpOrganizationByOrgId called")
	retval := Organization{}
	apiUrl := fmt.Sprintf(REST_LOOKUP_ORGANIZATION_BY_ID, orgId)
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

func (lic *LinkedInClient) LookUpOrganizationByVanityName(vanityName string) (Organization, error) {
	lic.logger.Debugf("LinkedInClient LookUpOrganizationByVanityName called")
	retval := Organization{}
	apiUrl := fmt.Sprintf(REST_LOOKUP_ORGANIZATION_BY_VANITY_NAME, vanityName)
	req, err := http.NewRequest(GET, apiUrl, nil)
	if err != nil {
		return retval, err
	}
	resp, err := lic.callRestAPI(req, &retval)
	if err != nil {
		return retval, err
	}
	if resp.StatusCode >= 300 {
		return retval, fmt.Errorf("error looking up org by vanity name [%d]:%v", resp.StatusCode, resp.Status)
	}
	return retval, err
}
