package linkedin

import (
	"fmt"
	"net/http"
)

const REST_URN_SENIORITY = "https://api.linkedin.com/v2/seniorities/%d?locale.language=en&locale.country=US" //GET
const REST_URN_INDUSTRY = "https://api.linkedin.com/v2/industries/%d?locale.language=en&locale.country=US"   //GET
const REST_URN_FUNCTION = "https://api.linkedin.com/v2/functions/%d?locale=en_US"                            //GET

func (lic *LinkedInClient) ResolveURN(id int, typeOfUrn string) (ResolvedURN, error) {
	lic.logger.Debugf("LinkedInClient ResolveURN called")
	switch typeOfUrn {
	case "seniority":
		return lic.resolveURN(id, REST_URN_SENIORITY)
	case "industry":
		return lic.resolveURN(id, REST_URN_INDUSTRY)
	case "function":
		return lic.resolveURN(id, REST_URN_FUNCTION)
	default:
		return ResolvedURN{}, fmt.Errorf("cannot resolve type %v, must be one of seniority, industry or function", typeOfUrn)
	}
}

func (lic *LinkedInClient) resolveURN(id int, url string) (ResolvedURN, error) {
	retval := ResolvedURN{}
	apiUrl := fmt.Sprintf(url, id)
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
