# InvestorPortfolioRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestorId** | Pointer to **NullableString** | The investor key. | [optional] 
**CompanyKey** | Pointer to **NullableString** | The company&#39;s key on this dataset: its organization id where it resolves, else the SEC CIK, else investor_site: followed by the normalized name as printed on the investor&#39;s page. | [optional] 
**CompanyKeyKind** | Pointer to **NullableString** | What company_key is: org, cik or investor_site. | [optional] 
**Sources** | Pointer to **[]string** | Where the row comes from: sec_edgar (S-1 and 424B holder tables), investor_site (the investor&#39;s own portfolio page), or both. | [optional] 
**IssuerCik** | Pointer to **NullableString** | The issuer&#39;s SEC Central Index Key (CIK), ten digits. Null for a row only the investor&#39;s website names. | [optional] 
**IssuerName** | Pointer to **NullableString** | The company&#39;s name: its organization record&#39;s name, else the SEC company name, else the name as printed on the investor&#39;s page. | [optional] 
**CompanyNameAsListed** | Pointer to **NullableString** | The company&#39;s name as printed on the investor&#39;s own portfolio page. Null for a row only SEC filings name. | [optional] 
**CompanyWebsiteUrl** | Pointer to **NullableString** | The company&#39;s website as linked on the investor&#39;s own portfolio page. Null when the page gave none. | [optional] 
**IssuerOrgId** | Pointer to **NullableString** | Legal-entity level: the company&#39;s organization record (null when unresolved). | [optional] 
**IssuerCompanyId** | Pointer to **NullableString** | Company level: the company&#39;s record (null when unresolved). | [optional] 
**IssuerCompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**IssuerGroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**IssuerGroupName** | Pointer to **NullableString** | Display name of the company&#39;s group. | [optional] 
**RelationshipTypes** | Pointer to **[]string** | How the investor is tied to the company: holder_5pct (a 5 percent holder) or selling_stockholder from filings, and portfolio_listed_current, portfolio_listed_realized, portfolio_listed_exited or portfolio_listed (no status printed) from the investor&#39;s own page. | [optional] 
**EdgeCount** | Pointer to **NullableInt64** | Filing mentions behind the row, across every filing and relationship type. 0 for a row only the investor&#39;s website names. | [optional] 
**FirstObservedDate** | Pointer to **NullableString** | The earliest filing date the investor was named against the issuer. Null for a row only the investor&#39;s website names. | [optional] 
**LastObservedDate** | Pointer to **NullableString** | The latest filing date the investor was named against the issuer. Null for a row only the investor&#39;s website names. | [optional] 
**ListedFrom** | Pointer to **NullableString** | The retrieval date of the first crawl of the investor&#39;s own portfolio page that listed the company (a crawl date, not an investment date). Null for a row only SEC filings name. | [optional] 
**ListedTo** | Pointer to **NullableString** | The retrieval date of the first later crawl that did not list the company, or the date the page went away. Null while the company is still listed, and for a row only SEC filings name. | [optional] 
**RemovalReason** | Pointer to **NullableString** | Why listed_to is set: absent_from_later_version (a later crawl of the page omitted the company) or page_gone (the page went away). Null exactly when listed_to is null. | [optional] 
**IsRemoved** | Pointer to **NullableBool** | True when the investor&#39;s own page no longer lists the company (listed_to is set). SEC filings do not reopen it. | [optional] 
**StatusOnPage** | Pointer to **NullableString** | The status the investor&#39;s page states for the company, verbatim: current, realized, exited or unknown (the page listed it with no status). Null for a row only SEC filings name. | [optional] 
**StatusDate** | Pointer to **NullableString** | The date the page itself prints for the status, as printed (for example 2024 or March 2024). Null unless the page printed one. | [optional] 
**StatusFirstObservedAt** | Pointer to **NullableString** | The date we first saw this status for the company on the investor&#39;s page: an observation date, never the date the event happened. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The address of the investor&#39;s portfolio page the listing was read from. Null for a row only SEC filings name. | [optional] 
**SiteLastRetrievedAt** | Pointer to **NullableTime** | When the investor&#39;s portfolio page was last retrieved with the company listed on it, in UTC. | [optional] 
**IsLifeSciencesIssuer** | Pointer to **NullableBool** | True when the issuer&#39;s SIC (Standard Industrial Classification) code is a life-sciences code, false when it is another code, null when the issuer has no SIC evidence. | [optional] 

## Methods

### NewInvestorPortfolioRow

`func NewInvestorPortfolioRow() *InvestorPortfolioRow`

NewInvestorPortfolioRow instantiates a new InvestorPortfolioRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestorPortfolioRowWithDefaults

`func NewInvestorPortfolioRowWithDefaults() *InvestorPortfolioRow`

NewInvestorPortfolioRowWithDefaults instantiates a new InvestorPortfolioRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestorId

`func (o *InvestorPortfolioRow) GetInvestorId() string`

GetInvestorId returns the InvestorId field if non-nil, zero value otherwise.

### GetInvestorIdOk

`func (o *InvestorPortfolioRow) GetInvestorIdOk() (*string, bool)`

GetInvestorIdOk returns a tuple with the InvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorId

`func (o *InvestorPortfolioRow) SetInvestorId(v string)`

SetInvestorId sets InvestorId field to given value.

### HasInvestorId

`func (o *InvestorPortfolioRow) HasInvestorId() bool`

HasInvestorId returns a boolean if a field has been set.

### SetInvestorIdNil

`func (o *InvestorPortfolioRow) SetInvestorIdNil(b bool)`

 SetInvestorIdNil sets the value for InvestorId to be an explicit nil

### UnsetInvestorId
`func (o *InvestorPortfolioRow) UnsetInvestorId()`

UnsetInvestorId ensures that no value is present for InvestorId, not even an explicit nil
### GetCompanyKey

`func (o *InvestorPortfolioRow) GetCompanyKey() string`

GetCompanyKey returns the CompanyKey field if non-nil, zero value otherwise.

### GetCompanyKeyOk

`func (o *InvestorPortfolioRow) GetCompanyKeyOk() (*string, bool)`

GetCompanyKeyOk returns a tuple with the CompanyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyKey

`func (o *InvestorPortfolioRow) SetCompanyKey(v string)`

SetCompanyKey sets CompanyKey field to given value.

### HasCompanyKey

`func (o *InvestorPortfolioRow) HasCompanyKey() bool`

HasCompanyKey returns a boolean if a field has been set.

### SetCompanyKeyNil

`func (o *InvestorPortfolioRow) SetCompanyKeyNil(b bool)`

 SetCompanyKeyNil sets the value for CompanyKey to be an explicit nil

### UnsetCompanyKey
`func (o *InvestorPortfolioRow) UnsetCompanyKey()`

UnsetCompanyKey ensures that no value is present for CompanyKey, not even an explicit nil
### GetCompanyKeyKind

`func (o *InvestorPortfolioRow) GetCompanyKeyKind() string`

GetCompanyKeyKind returns the CompanyKeyKind field if non-nil, zero value otherwise.

### GetCompanyKeyKindOk

`func (o *InvestorPortfolioRow) GetCompanyKeyKindOk() (*string, bool)`

GetCompanyKeyKindOk returns a tuple with the CompanyKeyKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyKeyKind

`func (o *InvestorPortfolioRow) SetCompanyKeyKind(v string)`

SetCompanyKeyKind sets CompanyKeyKind field to given value.

### HasCompanyKeyKind

`func (o *InvestorPortfolioRow) HasCompanyKeyKind() bool`

HasCompanyKeyKind returns a boolean if a field has been set.

### SetCompanyKeyKindNil

`func (o *InvestorPortfolioRow) SetCompanyKeyKindNil(b bool)`

 SetCompanyKeyKindNil sets the value for CompanyKeyKind to be an explicit nil

### UnsetCompanyKeyKind
`func (o *InvestorPortfolioRow) UnsetCompanyKeyKind()`

UnsetCompanyKeyKind ensures that no value is present for CompanyKeyKind, not even an explicit nil
### GetSources

`func (o *InvestorPortfolioRow) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *InvestorPortfolioRow) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *InvestorPortfolioRow) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *InvestorPortfolioRow) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *InvestorPortfolioRow) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *InvestorPortfolioRow) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetIssuerCik

`func (o *InvestorPortfolioRow) GetIssuerCik() string`

GetIssuerCik returns the IssuerCik field if non-nil, zero value otherwise.

### GetIssuerCikOk

`func (o *InvestorPortfolioRow) GetIssuerCikOk() (*string, bool)`

GetIssuerCikOk returns a tuple with the IssuerCik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCik

`func (o *InvestorPortfolioRow) SetIssuerCik(v string)`

SetIssuerCik sets IssuerCik field to given value.

### HasIssuerCik

`func (o *InvestorPortfolioRow) HasIssuerCik() bool`

HasIssuerCik returns a boolean if a field has been set.

### SetIssuerCikNil

`func (o *InvestorPortfolioRow) SetIssuerCikNil(b bool)`

 SetIssuerCikNil sets the value for IssuerCik to be an explicit nil

### UnsetIssuerCik
`func (o *InvestorPortfolioRow) UnsetIssuerCik()`

UnsetIssuerCik ensures that no value is present for IssuerCik, not even an explicit nil
### GetIssuerName

`func (o *InvestorPortfolioRow) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *InvestorPortfolioRow) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *InvestorPortfolioRow) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *InvestorPortfolioRow) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### SetIssuerNameNil

`func (o *InvestorPortfolioRow) SetIssuerNameNil(b bool)`

 SetIssuerNameNil sets the value for IssuerName to be an explicit nil

### UnsetIssuerName
`func (o *InvestorPortfolioRow) UnsetIssuerName()`

UnsetIssuerName ensures that no value is present for IssuerName, not even an explicit nil
### GetCompanyNameAsListed

`func (o *InvestorPortfolioRow) GetCompanyNameAsListed() string`

GetCompanyNameAsListed returns the CompanyNameAsListed field if non-nil, zero value otherwise.

### GetCompanyNameAsListedOk

`func (o *InvestorPortfolioRow) GetCompanyNameAsListedOk() (*string, bool)`

GetCompanyNameAsListedOk returns a tuple with the CompanyNameAsListed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNameAsListed

`func (o *InvestorPortfolioRow) SetCompanyNameAsListed(v string)`

SetCompanyNameAsListed sets CompanyNameAsListed field to given value.

### HasCompanyNameAsListed

`func (o *InvestorPortfolioRow) HasCompanyNameAsListed() bool`

HasCompanyNameAsListed returns a boolean if a field has been set.

### SetCompanyNameAsListedNil

`func (o *InvestorPortfolioRow) SetCompanyNameAsListedNil(b bool)`

 SetCompanyNameAsListedNil sets the value for CompanyNameAsListed to be an explicit nil

### UnsetCompanyNameAsListed
`func (o *InvestorPortfolioRow) UnsetCompanyNameAsListed()`

UnsetCompanyNameAsListed ensures that no value is present for CompanyNameAsListed, not even an explicit nil
### GetCompanyWebsiteUrl

`func (o *InvestorPortfolioRow) GetCompanyWebsiteUrl() string`

GetCompanyWebsiteUrl returns the CompanyWebsiteUrl field if non-nil, zero value otherwise.

### GetCompanyWebsiteUrlOk

`func (o *InvestorPortfolioRow) GetCompanyWebsiteUrlOk() (*string, bool)`

GetCompanyWebsiteUrlOk returns a tuple with the CompanyWebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyWebsiteUrl

`func (o *InvestorPortfolioRow) SetCompanyWebsiteUrl(v string)`

SetCompanyWebsiteUrl sets CompanyWebsiteUrl field to given value.

### HasCompanyWebsiteUrl

`func (o *InvestorPortfolioRow) HasCompanyWebsiteUrl() bool`

HasCompanyWebsiteUrl returns a boolean if a field has been set.

### SetCompanyWebsiteUrlNil

`func (o *InvestorPortfolioRow) SetCompanyWebsiteUrlNil(b bool)`

 SetCompanyWebsiteUrlNil sets the value for CompanyWebsiteUrl to be an explicit nil

### UnsetCompanyWebsiteUrl
`func (o *InvestorPortfolioRow) UnsetCompanyWebsiteUrl()`

UnsetCompanyWebsiteUrl ensures that no value is present for CompanyWebsiteUrl, not even an explicit nil
### GetIssuerOrgId

`func (o *InvestorPortfolioRow) GetIssuerOrgId() string`

GetIssuerOrgId returns the IssuerOrgId field if non-nil, zero value otherwise.

### GetIssuerOrgIdOk

`func (o *InvestorPortfolioRow) GetIssuerOrgIdOk() (*string, bool)`

GetIssuerOrgIdOk returns a tuple with the IssuerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerOrgId

`func (o *InvestorPortfolioRow) SetIssuerOrgId(v string)`

SetIssuerOrgId sets IssuerOrgId field to given value.

### HasIssuerOrgId

`func (o *InvestorPortfolioRow) HasIssuerOrgId() bool`

HasIssuerOrgId returns a boolean if a field has been set.

### SetIssuerOrgIdNil

`func (o *InvestorPortfolioRow) SetIssuerOrgIdNil(b bool)`

 SetIssuerOrgIdNil sets the value for IssuerOrgId to be an explicit nil

### UnsetIssuerOrgId
`func (o *InvestorPortfolioRow) UnsetIssuerOrgId()`

UnsetIssuerOrgId ensures that no value is present for IssuerOrgId, not even an explicit nil
### GetIssuerCompanyId

`func (o *InvestorPortfolioRow) GetIssuerCompanyId() string`

GetIssuerCompanyId returns the IssuerCompanyId field if non-nil, zero value otherwise.

### GetIssuerCompanyIdOk

`func (o *InvestorPortfolioRow) GetIssuerCompanyIdOk() (*string, bool)`

GetIssuerCompanyIdOk returns a tuple with the IssuerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCompanyId

`func (o *InvestorPortfolioRow) SetIssuerCompanyId(v string)`

SetIssuerCompanyId sets IssuerCompanyId field to given value.

### HasIssuerCompanyId

`func (o *InvestorPortfolioRow) HasIssuerCompanyId() bool`

HasIssuerCompanyId returns a boolean if a field has been set.

### SetIssuerCompanyIdNil

`func (o *InvestorPortfolioRow) SetIssuerCompanyIdNil(b bool)`

 SetIssuerCompanyIdNil sets the value for IssuerCompanyId to be an explicit nil

### UnsetIssuerCompanyId
`func (o *InvestorPortfolioRow) UnsetIssuerCompanyId()`

UnsetIssuerCompanyId ensures that no value is present for IssuerCompanyId, not even an explicit nil
### GetIssuerCompanyName

`func (o *InvestorPortfolioRow) GetIssuerCompanyName() string`

GetIssuerCompanyName returns the IssuerCompanyName field if non-nil, zero value otherwise.

### GetIssuerCompanyNameOk

`func (o *InvestorPortfolioRow) GetIssuerCompanyNameOk() (*string, bool)`

GetIssuerCompanyNameOk returns a tuple with the IssuerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCompanyName

`func (o *InvestorPortfolioRow) SetIssuerCompanyName(v string)`

SetIssuerCompanyName sets IssuerCompanyName field to given value.

### HasIssuerCompanyName

`func (o *InvestorPortfolioRow) HasIssuerCompanyName() bool`

HasIssuerCompanyName returns a boolean if a field has been set.

### SetIssuerCompanyNameNil

`func (o *InvestorPortfolioRow) SetIssuerCompanyNameNil(b bool)`

 SetIssuerCompanyNameNil sets the value for IssuerCompanyName to be an explicit nil

### UnsetIssuerCompanyName
`func (o *InvestorPortfolioRow) UnsetIssuerCompanyName()`

UnsetIssuerCompanyName ensures that no value is present for IssuerCompanyName, not even an explicit nil
### GetIssuerGroupId

`func (o *InvestorPortfolioRow) GetIssuerGroupId() string`

GetIssuerGroupId returns the IssuerGroupId field if non-nil, zero value otherwise.

### GetIssuerGroupIdOk

`func (o *InvestorPortfolioRow) GetIssuerGroupIdOk() (*string, bool)`

GetIssuerGroupIdOk returns a tuple with the IssuerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerGroupId

`func (o *InvestorPortfolioRow) SetIssuerGroupId(v string)`

SetIssuerGroupId sets IssuerGroupId field to given value.

### HasIssuerGroupId

`func (o *InvestorPortfolioRow) HasIssuerGroupId() bool`

HasIssuerGroupId returns a boolean if a field has been set.

### SetIssuerGroupIdNil

`func (o *InvestorPortfolioRow) SetIssuerGroupIdNil(b bool)`

 SetIssuerGroupIdNil sets the value for IssuerGroupId to be an explicit nil

### UnsetIssuerGroupId
`func (o *InvestorPortfolioRow) UnsetIssuerGroupId()`

UnsetIssuerGroupId ensures that no value is present for IssuerGroupId, not even an explicit nil
### GetIssuerGroupName

`func (o *InvestorPortfolioRow) GetIssuerGroupName() string`

GetIssuerGroupName returns the IssuerGroupName field if non-nil, zero value otherwise.

### GetIssuerGroupNameOk

`func (o *InvestorPortfolioRow) GetIssuerGroupNameOk() (*string, bool)`

GetIssuerGroupNameOk returns a tuple with the IssuerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerGroupName

`func (o *InvestorPortfolioRow) SetIssuerGroupName(v string)`

SetIssuerGroupName sets IssuerGroupName field to given value.

### HasIssuerGroupName

`func (o *InvestorPortfolioRow) HasIssuerGroupName() bool`

HasIssuerGroupName returns a boolean if a field has been set.

### SetIssuerGroupNameNil

`func (o *InvestorPortfolioRow) SetIssuerGroupNameNil(b bool)`

 SetIssuerGroupNameNil sets the value for IssuerGroupName to be an explicit nil

### UnsetIssuerGroupName
`func (o *InvestorPortfolioRow) UnsetIssuerGroupName()`

UnsetIssuerGroupName ensures that no value is present for IssuerGroupName, not even an explicit nil
### GetRelationshipTypes

`func (o *InvestorPortfolioRow) GetRelationshipTypes() []string`

GetRelationshipTypes returns the RelationshipTypes field if non-nil, zero value otherwise.

### GetRelationshipTypesOk

`func (o *InvestorPortfolioRow) GetRelationshipTypesOk() (*[]string, bool)`

GetRelationshipTypesOk returns a tuple with the RelationshipTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipTypes

`func (o *InvestorPortfolioRow) SetRelationshipTypes(v []string)`

SetRelationshipTypes sets RelationshipTypes field to given value.

### HasRelationshipTypes

`func (o *InvestorPortfolioRow) HasRelationshipTypes() bool`

HasRelationshipTypes returns a boolean if a field has been set.

### SetRelationshipTypesNil

`func (o *InvestorPortfolioRow) SetRelationshipTypesNil(b bool)`

 SetRelationshipTypesNil sets the value for RelationshipTypes to be an explicit nil

### UnsetRelationshipTypes
`func (o *InvestorPortfolioRow) UnsetRelationshipTypes()`

UnsetRelationshipTypes ensures that no value is present for RelationshipTypes, not even an explicit nil
### GetEdgeCount

`func (o *InvestorPortfolioRow) GetEdgeCount() int64`

GetEdgeCount returns the EdgeCount field if non-nil, zero value otherwise.

### GetEdgeCountOk

`func (o *InvestorPortfolioRow) GetEdgeCountOk() (*int64, bool)`

GetEdgeCountOk returns a tuple with the EdgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCount

`func (o *InvestorPortfolioRow) SetEdgeCount(v int64)`

SetEdgeCount sets EdgeCount field to given value.

### HasEdgeCount

`func (o *InvestorPortfolioRow) HasEdgeCount() bool`

HasEdgeCount returns a boolean if a field has been set.

### SetEdgeCountNil

`func (o *InvestorPortfolioRow) SetEdgeCountNil(b bool)`

 SetEdgeCountNil sets the value for EdgeCount to be an explicit nil

### UnsetEdgeCount
`func (o *InvestorPortfolioRow) UnsetEdgeCount()`

UnsetEdgeCount ensures that no value is present for EdgeCount, not even an explicit nil
### GetFirstObservedDate

`func (o *InvestorPortfolioRow) GetFirstObservedDate() string`

GetFirstObservedDate returns the FirstObservedDate field if non-nil, zero value otherwise.

### GetFirstObservedDateOk

`func (o *InvestorPortfolioRow) GetFirstObservedDateOk() (*string, bool)`

GetFirstObservedDateOk returns a tuple with the FirstObservedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstObservedDate

`func (o *InvestorPortfolioRow) SetFirstObservedDate(v string)`

SetFirstObservedDate sets FirstObservedDate field to given value.

### HasFirstObservedDate

`func (o *InvestorPortfolioRow) HasFirstObservedDate() bool`

HasFirstObservedDate returns a boolean if a field has been set.

### SetFirstObservedDateNil

`func (o *InvestorPortfolioRow) SetFirstObservedDateNil(b bool)`

 SetFirstObservedDateNil sets the value for FirstObservedDate to be an explicit nil

### UnsetFirstObservedDate
`func (o *InvestorPortfolioRow) UnsetFirstObservedDate()`

UnsetFirstObservedDate ensures that no value is present for FirstObservedDate, not even an explicit nil
### GetLastObservedDate

`func (o *InvestorPortfolioRow) GetLastObservedDate() string`

GetLastObservedDate returns the LastObservedDate field if non-nil, zero value otherwise.

### GetLastObservedDateOk

`func (o *InvestorPortfolioRow) GetLastObservedDateOk() (*string, bool)`

GetLastObservedDateOk returns a tuple with the LastObservedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastObservedDate

`func (o *InvestorPortfolioRow) SetLastObservedDate(v string)`

SetLastObservedDate sets LastObservedDate field to given value.

### HasLastObservedDate

`func (o *InvestorPortfolioRow) HasLastObservedDate() bool`

HasLastObservedDate returns a boolean if a field has been set.

### SetLastObservedDateNil

`func (o *InvestorPortfolioRow) SetLastObservedDateNil(b bool)`

 SetLastObservedDateNil sets the value for LastObservedDate to be an explicit nil

### UnsetLastObservedDate
`func (o *InvestorPortfolioRow) UnsetLastObservedDate()`

UnsetLastObservedDate ensures that no value is present for LastObservedDate, not even an explicit nil
### GetListedFrom

`func (o *InvestorPortfolioRow) GetListedFrom() string`

GetListedFrom returns the ListedFrom field if non-nil, zero value otherwise.

### GetListedFromOk

`func (o *InvestorPortfolioRow) GetListedFromOk() (*string, bool)`

GetListedFromOk returns a tuple with the ListedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedFrom

`func (o *InvestorPortfolioRow) SetListedFrom(v string)`

SetListedFrom sets ListedFrom field to given value.

### HasListedFrom

`func (o *InvestorPortfolioRow) HasListedFrom() bool`

HasListedFrom returns a boolean if a field has been set.

### SetListedFromNil

`func (o *InvestorPortfolioRow) SetListedFromNil(b bool)`

 SetListedFromNil sets the value for ListedFrom to be an explicit nil

### UnsetListedFrom
`func (o *InvestorPortfolioRow) UnsetListedFrom()`

UnsetListedFrom ensures that no value is present for ListedFrom, not even an explicit nil
### GetListedTo

`func (o *InvestorPortfolioRow) GetListedTo() string`

GetListedTo returns the ListedTo field if non-nil, zero value otherwise.

### GetListedToOk

`func (o *InvestorPortfolioRow) GetListedToOk() (*string, bool)`

GetListedToOk returns a tuple with the ListedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedTo

`func (o *InvestorPortfolioRow) SetListedTo(v string)`

SetListedTo sets ListedTo field to given value.

### HasListedTo

`func (o *InvestorPortfolioRow) HasListedTo() bool`

HasListedTo returns a boolean if a field has been set.

### SetListedToNil

`func (o *InvestorPortfolioRow) SetListedToNil(b bool)`

 SetListedToNil sets the value for ListedTo to be an explicit nil

### UnsetListedTo
`func (o *InvestorPortfolioRow) UnsetListedTo()`

UnsetListedTo ensures that no value is present for ListedTo, not even an explicit nil
### GetRemovalReason

`func (o *InvestorPortfolioRow) GetRemovalReason() string`

GetRemovalReason returns the RemovalReason field if non-nil, zero value otherwise.

### GetRemovalReasonOk

`func (o *InvestorPortfolioRow) GetRemovalReasonOk() (*string, bool)`

GetRemovalReasonOk returns a tuple with the RemovalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalReason

`func (o *InvestorPortfolioRow) SetRemovalReason(v string)`

SetRemovalReason sets RemovalReason field to given value.

### HasRemovalReason

`func (o *InvestorPortfolioRow) HasRemovalReason() bool`

HasRemovalReason returns a boolean if a field has been set.

### SetRemovalReasonNil

`func (o *InvestorPortfolioRow) SetRemovalReasonNil(b bool)`

 SetRemovalReasonNil sets the value for RemovalReason to be an explicit nil

### UnsetRemovalReason
`func (o *InvestorPortfolioRow) UnsetRemovalReason()`

UnsetRemovalReason ensures that no value is present for RemovalReason, not even an explicit nil
### GetIsRemoved

`func (o *InvestorPortfolioRow) GetIsRemoved() bool`

GetIsRemoved returns the IsRemoved field if non-nil, zero value otherwise.

### GetIsRemovedOk

`func (o *InvestorPortfolioRow) GetIsRemovedOk() (*bool, bool)`

GetIsRemovedOk returns a tuple with the IsRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoved

`func (o *InvestorPortfolioRow) SetIsRemoved(v bool)`

SetIsRemoved sets IsRemoved field to given value.

### HasIsRemoved

`func (o *InvestorPortfolioRow) HasIsRemoved() bool`

HasIsRemoved returns a boolean if a field has been set.

### SetIsRemovedNil

`func (o *InvestorPortfolioRow) SetIsRemovedNil(b bool)`

 SetIsRemovedNil sets the value for IsRemoved to be an explicit nil

### UnsetIsRemoved
`func (o *InvestorPortfolioRow) UnsetIsRemoved()`

UnsetIsRemoved ensures that no value is present for IsRemoved, not even an explicit nil
### GetStatusOnPage

`func (o *InvestorPortfolioRow) GetStatusOnPage() string`

GetStatusOnPage returns the StatusOnPage field if non-nil, zero value otherwise.

### GetStatusOnPageOk

`func (o *InvestorPortfolioRow) GetStatusOnPageOk() (*string, bool)`

GetStatusOnPageOk returns a tuple with the StatusOnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOnPage

`func (o *InvestorPortfolioRow) SetStatusOnPage(v string)`

SetStatusOnPage sets StatusOnPage field to given value.

### HasStatusOnPage

`func (o *InvestorPortfolioRow) HasStatusOnPage() bool`

HasStatusOnPage returns a boolean if a field has been set.

### SetStatusOnPageNil

`func (o *InvestorPortfolioRow) SetStatusOnPageNil(b bool)`

 SetStatusOnPageNil sets the value for StatusOnPage to be an explicit nil

### UnsetStatusOnPage
`func (o *InvestorPortfolioRow) UnsetStatusOnPage()`

UnsetStatusOnPage ensures that no value is present for StatusOnPage, not even an explicit nil
### GetStatusDate

`func (o *InvestorPortfolioRow) GetStatusDate() string`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *InvestorPortfolioRow) GetStatusDateOk() (*string, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *InvestorPortfolioRow) SetStatusDate(v string)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *InvestorPortfolioRow) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *InvestorPortfolioRow) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *InvestorPortfolioRow) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusFirstObservedAt

`func (o *InvestorPortfolioRow) GetStatusFirstObservedAt() string`

GetStatusFirstObservedAt returns the StatusFirstObservedAt field if non-nil, zero value otherwise.

### GetStatusFirstObservedAtOk

`func (o *InvestorPortfolioRow) GetStatusFirstObservedAtOk() (*string, bool)`

GetStatusFirstObservedAtOk returns a tuple with the StatusFirstObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusFirstObservedAt

`func (o *InvestorPortfolioRow) SetStatusFirstObservedAt(v string)`

SetStatusFirstObservedAt sets StatusFirstObservedAt field to given value.

### HasStatusFirstObservedAt

`func (o *InvestorPortfolioRow) HasStatusFirstObservedAt() bool`

HasStatusFirstObservedAt returns a boolean if a field has been set.

### SetStatusFirstObservedAtNil

`func (o *InvestorPortfolioRow) SetStatusFirstObservedAtNil(b bool)`

 SetStatusFirstObservedAtNil sets the value for StatusFirstObservedAt to be an explicit nil

### UnsetStatusFirstObservedAt
`func (o *InvestorPortfolioRow) UnsetStatusFirstObservedAt()`

UnsetStatusFirstObservedAt ensures that no value is present for StatusFirstObservedAt, not even an explicit nil
### GetSourceUrl

`func (o *InvestorPortfolioRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *InvestorPortfolioRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *InvestorPortfolioRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *InvestorPortfolioRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *InvestorPortfolioRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *InvestorPortfolioRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSiteLastRetrievedAt

`func (o *InvestorPortfolioRow) GetSiteLastRetrievedAt() time.Time`

GetSiteLastRetrievedAt returns the SiteLastRetrievedAt field if non-nil, zero value otherwise.

### GetSiteLastRetrievedAtOk

`func (o *InvestorPortfolioRow) GetSiteLastRetrievedAtOk() (*time.Time, bool)`

GetSiteLastRetrievedAtOk returns a tuple with the SiteLastRetrievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteLastRetrievedAt

`func (o *InvestorPortfolioRow) SetSiteLastRetrievedAt(v time.Time)`

SetSiteLastRetrievedAt sets SiteLastRetrievedAt field to given value.

### HasSiteLastRetrievedAt

`func (o *InvestorPortfolioRow) HasSiteLastRetrievedAt() bool`

HasSiteLastRetrievedAt returns a boolean if a field has been set.

### SetSiteLastRetrievedAtNil

`func (o *InvestorPortfolioRow) SetSiteLastRetrievedAtNil(b bool)`

 SetSiteLastRetrievedAtNil sets the value for SiteLastRetrievedAt to be an explicit nil

### UnsetSiteLastRetrievedAt
`func (o *InvestorPortfolioRow) UnsetSiteLastRetrievedAt()`

UnsetSiteLastRetrievedAt ensures that no value is present for SiteLastRetrievedAt, not even an explicit nil
### GetIsLifeSciencesIssuer

`func (o *InvestorPortfolioRow) GetIsLifeSciencesIssuer() bool`

GetIsLifeSciencesIssuer returns the IsLifeSciencesIssuer field if non-nil, zero value otherwise.

### GetIsLifeSciencesIssuerOk

`func (o *InvestorPortfolioRow) GetIsLifeSciencesIssuerOk() (*bool, bool)`

GetIsLifeSciencesIssuerOk returns a tuple with the IsLifeSciencesIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLifeSciencesIssuer

`func (o *InvestorPortfolioRow) SetIsLifeSciencesIssuer(v bool)`

SetIsLifeSciencesIssuer sets IsLifeSciencesIssuer field to given value.

### HasIsLifeSciencesIssuer

`func (o *InvestorPortfolioRow) HasIsLifeSciencesIssuer() bool`

HasIsLifeSciencesIssuer returns a boolean if a field has been set.

### SetIsLifeSciencesIssuerNil

`func (o *InvestorPortfolioRow) SetIsLifeSciencesIssuerNil(b bool)`

 SetIsLifeSciencesIssuerNil sets the value for IsLifeSciencesIssuer to be an explicit nil

### UnsetIsLifeSciencesIssuer
`func (o *InvestorPortfolioRow) UnsetIsLifeSciencesIssuer()`

UnsetIsLifeSciencesIssuer ensures that no value is present for IsLifeSciencesIssuer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


