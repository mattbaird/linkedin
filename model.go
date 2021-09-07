package linkedin

import (
	"fmt"
	"strings"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

type CreateAdAccountRequest struct {
	Currency                       string `json:"currency"`
	Name                           string `json:"name"`
	NotifiedOnCampaignOptimization bool   `json:"notifiedOnCampaignOptimization"`
	NotifiedOnCreativeApproval     bool   `json:"notifiedOnCreativeApproval"`
	NotifiedOnCreativeRejection    bool   `json:"notifiedOnCreativeRejection"`
	NotifiedOnEndOfCampaign        bool   `json:"notifiedOnEndOfCampaign"`
	Reference                      string `json:"reference"`
	Type                           string `json:"type"`
}

type FetchAdAccountResponse struct {
	Test              bool `json:"test"`
	ChangeAuditStamps struct {
		Created struct {
			Actor string `json:"actor"`
			Time  int64  `json:"time"`
		} `json:"created"`
		LastModified struct {
			Actor string `json:"actor"`
			Time  int64  `json:"time"`
		} `json:"lastModified"`
	} `json:"changeAuditStamps"`
	Currency                       string   `json:"currency"`
	ID                             int      `json:"id"`
	Name                           string   `json:"name"`
	NotifiedOnCampaignOptimization bool     `json:"notifiedOnCampaignOptimization"`
	NotifiedOnCreativeApproval     bool     `json:"notifiedOnCreativeApproval"`
	NotifiedOnCreativeRejection    bool     `json:"notifiedOnCreativeRejection"`
	NotifiedOnEndOfCampaign        bool     `json:"notifiedOnEndOfCampaign"`
	Reference                      string   `json:"reference"`
	ServingStatuses                []string `json:"servingStatuses"`
	Status                         string   `json:"status"`
	Type                           string   `json:"type"`
	Version                        struct {
		VersionTag string `json:"versionTag"`
	} `json:"version"`
}

// we only support creating organization ad accounts (reference urn)
func newAdAccountRequest(currency string, name string, notifiedOnCampaignOptimization bool,
	notifiedOnCreativeApproval bool, notifiedOnCreativeRejection bool, notifiedOnEndOfCampaign bool,
	orgId int, accountType string) (CreateAdAccountRequest, error) {
	if len(strings.TrimSpace(name)) == 0 {
		return CreateAdAccountRequest{}, fmt.Errorf("name was empty")
	}
	if !(accountType == "BUSINESS" || accountType == "ENTERPRISE") {
		return CreateAdAccountRequest{}, fmt.Errorf("accountType must be one of BUSINESS or ENTERPRISE")
	}
	actualCurrency := "USD"
	if len(strings.TrimSpace(currency)) != 0 {
		actualCurrency = currency
	}
	return CreateAdAccountRequest{
		Currency:                       actualCurrency,
		Name:                           name,
		NotifiedOnCampaignOptimization: notifiedOnCampaignOptimization,
		NotifiedOnCreativeApproval:     notifiedOnCreativeApproval,
		NotifiedOnCreativeRejection:    notifiedOnCreativeRejection,
		NotifiedOnEndOfCampaign:        notifiedOnEndOfCampaign,
		Reference:                      fmt.Sprintf("urn:li:organization:%v", orgId),
		Type:                           accountType,
	}, nil
}

type CreateCampaignGroupResponse struct {
	Account     string `json:"account"`
	Name        string `json:"name"`
	RunSchedule struct {
		End   int `json:"end"`
		Start int `json:"start"`
	} `json:"runSchedule"`
	Status      string `json:"status"`
	TotalBudget struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currencyCode"`
	} `json:"totalBudget"`
}

type GetCampaignGroupResponse struct {
	Account           string `json:"account"`
	Backfilled        bool   `json:"backfilled"`
	ChangeAuditStamps struct {
		Created struct {
			Actor string `json:"actor"`
			Time  int64  `json:"time"`
		} `json:"created"`
		LastModified struct {
			Actor string `json:"actor"`
			Time  int64  `json:"time"`
		} `json:"lastModified"`
	} `json:"changeAuditStamps"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	RunSchedule struct {
		End   int `json:"end"`
		Start int `json:"start"`
	} `json:"runSchedule"`
	Test                 bool     `json:"test"`
	ServingStatuses      []string `json:"servingStatuses"`
	AllowedCampaignTypes []string `json:"allowedCampaignTypes"`
	Status               string   `json:"status"`
}

type TargetingFacetsResponse struct {
	Elements []struct {
		FacetName              string   `json:"facetName"`
		AvailableEntityFinders []string `json:"availableEntityFinders"`
		EntityTypes            []string `json:"entityTypes"`
		URN                    string   `json:"$URN"`
	} `json:"elements"`
}

type TargetingFacetEntitiesResponse struct {
	Elements []struct {
		Urn      string `json:"urn"`
		FacetUrn string `json:"facetUrn"`
		Name     string `json:"name"`
	} `json:"elements"`
}

type AudienceCountsResponse struct {
	Elements []struct {
		Total  int `json:"total"`
		Active int `json:"active"`
	} `json:"elements"`
	Paging struct {
		Count int           `json:"count"`
		Links []interface{} `json:"links"`
		Start int           `json:"start"`
	} `json:"paging"`
}

type AdBudgetPricingResponse struct {
	Elements []struct {
		SuggestedBid struct {
			Default struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"default"`
			Min struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"min"`
			Max struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"max"`
		} `json:"suggestedBid"`
		DailyBudgetLimits struct {
			Default struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"default"`
			Min struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"min"`
			Max struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"max"`
		} `json:"dailyBudgetLimits"`
		BidLimits struct {
			Min struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"min"`
			Max struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"max"`
		} `json:"bidLimits"`
	} `json:"elements"`
	Paging struct {
		Count int           `json:"count"`
		Start int           `json:"start"`
		Links []interface{} `json:"links"`
	} `json:"paging"`
}

type Campaign struct {
}
type CreateCampaignRequest struct {
	Account                  string `json:"account"`
	AudienceExpansionEnabled bool   `json:"audienceExpansionEnabled"`
	CostType                 string `json:"costType"`
	CreativeSelection        string `json:"creativeSelection"`
	DailyBudget              struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currencyCode"`
	} `json:"dailyBudget"`
	Locale struct {
		Country  string `json:"country"`
		Language string `json:"language"`
	} `json:"locale"`
	Name                   string `json:"name"`
	OffsiteDeliveryEnabled bool   `json:"offsiteDeliveryEnabled"`
	RunSchedule            struct {
		Start int64 `json:"start"`
		End   int64 `json:"end"`
	} `json:"runSchedule"`
	TargetingCriteria struct {
		Include struct {
			And []struct {
				Or struct {
					UrnLiAdTargetingFacetLocations []string `json:"urn:li:adTargetingFacet:locations"`
				} `json:"or"`
			} `json:"and"`
		} `json:"include"`
		Exclude struct {
			Or struct {
				UrnLiAdTargetingFacetSeniorities []string `json:"urn:li:adTargetingFacet:seniorities"`
			} `json:"or"`
		} `json:"exclude"`
	} `json:"targetingCriteria"`
	Type     string `json:"type"`
	UnitCost struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currencyCode"`
	} `json:"unitCost"`
}

type CreateAdCreativeResponse struct {
	Campaign  string `json:"campaign"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Variables struct {
		ClickURI string `json:"clickUri"`
		Data     struct {
			ComLinkedinAdsTextAdCreativeVariables struct {
				Text  string `json:"text"`
				Title string `json:"title"`
			} `json:"com.linkedin.ads.TextAdCreativeVariables"`
		} `json:"data"`
	} `json:"variables"`
}

type AdAnalyticsResponse struct {
	Elements []struct {
		ChargeableClicks int `json:"chargeableClicks"`
		DateRange        struct {
			End struct {
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			} `json:"end"`
			Start struct {
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			} `json:"start"`
		} `json:"dateRange"`
		Impressions         int    `json:"impressions"`
		NonchargeableClicks int    `json:"nonchargeableClicks"`
		Pivot               string `json:"pivot"`
		Spend               struct {
			Amount       string `json:"amount"`
			CurrencyCode string `json:"currencyCode"`
		} `json:"spend"`
	} `json:"elements"`
	Paging struct {
		Count int `json:"count"`
		Start int `json:"start"`
	} `json:"paging"`
}

type GetConversionsResponse struct {
	PostClickAttributionWindowSize   int    `json:"postClickAttributionWindowSize"`
	ViewThroughAttributionWindowSize int    `json:"viewThroughAttributionWindowSize"`
	Created                          int64  `json:"created"`
	ImagePixelTag                    string `json:"imagePixelTag"`
	Type                             string `json:"type"`
	Enabled                          bool   `json:"enabled"`
	AssociatedCampaigns              []struct {
		Campaign     string `json:"campaign"`
		AssociatedAt int64  `json:"associatedAt"`
		Conversion   string `json:"conversion"`
	} `json:"associatedCampaigns"`
	Campaigns              []string `json:"campaigns"`
	URLMatchRuleExpression [][]struct {
		MatchType  string `json:"matchType"`
		MatchValue string `json:"matchValue"`
	} `json:"urlMatchRuleExpression"`
	Name            string        `json:"name"`
	LastModified    int64         `json:"lastModified"`
	ID              int           `json:"id"`
	AttributionType string        `json:"attributionType"`
	URLRules        []interface{} `json:"urlRules"`
	Account         string        `json:"account"`
}

type BatchGetConversionsResponse struct {
	Statuses struct {
	} `json:"statuses"`
	Results struct {
		Num104004 struct {
			PostClickAttributionWindowSize   int    `json:"postClickAttributionWindowSize"`
			Created                          int64  `json:"created"`
			ViewThroughAttributionWindowSize int    `json:"viewThroughAttributionWindowSize"`
			ImagePixelTag                    string `json:"imagePixelTag"`
			Type                             string `json:"type"`
			Enabled                          bool   `json:"enabled"`
			AssociatedCampaigns              []struct {
				Campaign     string `json:"campaign"`
				AssociatedAt int64  `json:"associatedAt"`
				Conversion   string `json:"conversion"`
			} `json:"associatedCampaigns"`
			Campaigns       []string      `json:"campaigns"`
			Name            string        `json:"name"`
			LastModified    int64         `json:"lastModified"`
			ID              int           `json:"id"`
			AttributionType string        `json:"attributionType"`
			URLRules        []interface{} `json:"urlRules"`
			Value           struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"value"`
			Account string `json:"account"`
		} `json:"104004"`
		Num104012 struct {
			PostClickAttributionWindowSize   int    `json:"postClickAttributionWindowSize"`
			Created                          int64  `json:"created"`
			ViewThroughAttributionWindowSize int    `json:"viewThroughAttributionWindowSize"`
			ImagePixelTag                    string `json:"imagePixelTag"`
			Type                             string `json:"type"`
			Enabled                          bool   `json:"enabled"`
			AssociatedCampaigns              []struct {
				Campaign     string `json:"campaign"`
				AssociatedAt int64  `json:"associatedAt"`
				Conversion   string `json:"conversion"`
			} `json:"associatedCampaigns"`
			Campaigns              []string `json:"campaigns"`
			Name                   string   `json:"name"`
			URLMatchRuleExpression [][]struct {
				MatchType  string `json:"matchType"`
				MatchValue string `json:"matchValue"`
			} `json:"urlMatchRuleExpression"`
			LastModified    int64  `json:"lastModified"`
			ID              int    `json:"id"`
			AttributionType string `json:"attributionType"`
			URLRules        []struct {
				MatchValue string `json:"matchValue"`
				Type       string `json:"type"`
			} `json:"urlRules"`
			Value struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"value"`
			Account string `json:"account"`
		} `json:"104012"`
	} `json:"results"`
	Errors struct {
	} `json:"errors"`
}

type FollowerStatistics struct {
	Paging struct {
		Start int           `json:"start"`
		Count int           `json:"count"`
		Links []interface{} `json:"links"`
	} `json:"paging"`
	Elements []struct {
		Followercountsbyassociationtype []struct {
			Followercounts struct {
				Organicfollowercount int `json:"organicFollowerCount"`
				Paidfollowercount    int `json:"paidFollowerCount"`
			} `json:"followerCounts"`
			Associationtype string `json:"associationType,omitempty"`
		} `json:"followerCountsByAssociationType"`
		Followercountsbyregion []struct {
			Region         string `json:"region"`
			Followercounts struct {
				Organicfollowercount int `json:"organicFollowerCount"`
				Paidfollowercount    int `json:"paidFollowerCount"`
			} `json:"followerCounts"`
		} `json:"followerCountsByRegion"`
		Followercountsbyseniority []struct {
			Followercounts struct {
				Organicfollowercount int `json:"organicFollowerCount"`
				Paidfollowercount    int `json:"paidFollowerCount"`
			} `json:"followerCounts"`
			Seniority string `json:"seniority"`
		} `json:"followerCountsBySeniority"`
		Followercountsbyindustry []struct {
			Followercounts struct {
				Organicfollowercount int `json:"organicFollowerCount"`
				Paidfollowercount    int `json:"paidFollowerCount"`
			} `json:"followerCounts"`
			Industry string `json:"industry"`
		} `json:"followerCountsByIndustry"`
		Followercountsbyfunction []struct {
			Followercounts struct {
				Organicfollowercount int `json:"organicFollowerCount"`
				Paidfollowercount    int `json:"paidFollowerCount"`
			} `json:"followerCounts"`
			Function string `json:"function"`
		} `json:"followerCountsByFunction"`
		Followercountsbystaffcountrange []struct {
			Followercounts struct {
				Organicfollowercount int `json:"organicFollowerCount"`
				Paidfollowercount    int `json:"paidFollowerCount"`
			} `json:"followerCounts"`
			Staffcountrange string `json:"staffCountRange"`
		} `json:"followerCountsByStaffCountRange"`
		Followercountsbycountry []struct {
			Followercounts struct {
				Organicfollowercount int `json:"organicFollowerCount"`
				Paidfollowercount    int `json:"paidFollowerCount"`
			} `json:"followerCounts"`
			Country string `json:"country"`
		} `json:"followerCountsByCountry"`
		Organizationalentity string `json:"organizationalEntity"`
	} `json:"elements"`
}

type Organization struct {
	Elements []struct {
		Urn              string        `json:"$URN"`
		Alternativenames []interface{} `json:"alternativeNames"`
		Autocreated      bool          `json:"autoCreated"`
		Defaultlocale    struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"defaultLocale"`
		Description struct {
			Localized struct {
				EnUs string `json:"en_US"`
			} `json:"localized"`
			Preferredlocale struct {
				Country  string `json:"country"`
				Language string `json:"language"`
			} `json:"preferredLocale"`
		} `json:"description"`
		ID                   int           `json:"id"`
		Industries           []string      `json:"industries"`
		Localizeddescription string        `json:"localizedDescription"`
		Localizedname        string        `json:"localizedName"`
		Localizedspecialties []interface{} `json:"localizedSpecialties"`
		Name                 struct {
			Localized struct {
				EnUs string `json:"en_US"`
			} `json:"localized"`
			Preferredlocale struct {
				Country  string `json:"country"`
				Language string `json:"language"`
			} `json:"preferredLocale"`
		} `json:"name"`
		Specialties []interface{} `json:"specialties"`
		Vanityname  string        `json:"vanityName"`
		Versiontag  string        `json:"versionTag"`
		Website     struct {
			Localized struct {
				EnUs string `json:"en_US"`
			} `json:"localized"`
			Preferredlocale struct {
				Country  string `json:"country"`
				Language string `json:"language"`
			} `json:"preferredLocale"`
		} `json:"website"`
	} `json:"elements"`
	Paging struct {
		Count int           `json:"count"`
		Links []interface{} `json:"links"`
		Start int           `json:"start"`
		Total int           `json:"total"`
	} `json:"paging"`
}

type ResolvedURN struct {
	Urn  string `json:"$URN"`
	Name struct {
		Localized struct {
			EnUs string `json:"en_US"`
		} `json:"localized"`
	} `json:"name"`
	ID int `json:"id"`
}
