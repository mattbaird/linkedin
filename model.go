package linkedin

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

type CreateCampaignResponse struct {
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
