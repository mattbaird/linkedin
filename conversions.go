package linkedin

const REST_GET_CONVERSIONS = "https://api.linkedin.com/v2/conversions/%s?account=%s" // GET
//{conversionId}
//{sponsoredAccountUrn}

const REST_BATCH_GET_CONVERSIONS = "https://api.linkedin.com/v2/conversions?account=urn%3Ali%3AsponsoredAccount%3A519072844&ids=List(104012,104004)" //GET
//account=urn%3Ali%3AsponsoredAccount%3A519072844&
//ids=List(104012,104004)

const REST_GET_CONVERSIONS_BY_AD_ACCOUNT = "https://api.linkedin.com/v2/conversions?q=account&account=urn%3Ali%3AsponsoredAccount%3A519072844" //GET
//q=account&
//account=urn%3Ali%3AsponsoredAccount%3A519072844"
