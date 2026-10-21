# OrgDisclosuresRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **NullableString** | Legal-entity level: the organization record. | [optional] 
**OrgName** | Pointer to **NullableString** | The organization&#39;s name. | [optional] 
**CompanyId** | Pointer to **NullableString** | Company level: the company (null when unresolved). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**ClaimCategory** | Pointer to **NullableString** | The claim&#39;s domain: climate or diversity. | [optional] 
**ClaimType** | Pointer to **NullableString** | Which claim the row is about. Climate: net_zero_target, near_term_target, sbti_status_claimed, target_temperature_claimed. Diversity: board_pct_women, workforce_pct_women, leadership_pct_women, workforce_pct_underrepresented, gender_pay_gap_pct, has_dei_targets, external_recognition_claimed. | [optional] 
**DisclosureSource** | Pointer to **NullableString** | Where the claim comes from: sec_edgar_extraction (figures extracted from the company&#39;s own US Securities and Exchange Commission filings, open) or sbti (the Science Based Targets initiative dashboard, licensable). | [optional] 
**DisclosureState** | Pointer to **NullableString** | disclosed: the source states a value or a positive claim. true_negative: the source states a real negative (false or none). not_disclosed: the source covers this organization and states no value for the claim. A claim withheld from a caller who lacks the licence is not a state: the row is absent for that caller and the response names the withheld tier. | [optional] 
**ClaimValue** | Pointer to **NullableString** | The claimed value as text: true or false for a yes-or-no claim, a number for a percentage, the source&#39;s own word for a status. Null when disclosure_state is not_disclosed. | [optional] 
**ClaimDetail** | Pointer to **NullableString** | A qualifier of the claim, such as the target year. Null when the source gives none. | [optional] 
**DisclosureDate** | Pointer to **NullableString** | The date of the disclosure: the filing date for an extracted claim, the dashboard&#39;s last-updated date for an SBTi claim. Null when the source gives none. | [optional] 
**ReportingYear** | Pointer to **NullableString** | The reporting year the company states for an extracted claim. Null for an SBTi claim. | [optional] 
**EvidenceQuote** | Pointer to **NullableString** | The sentence of the filing the figure was extracted from. Null for an SBTi claim. | [optional] 
**Confidence** | Pointer to **NullableFloat32** | The extraction confidence, 0 to 1. Null for an SBTi claim, which carries none. | [optional] 
**SourceRecordId** | Pointer to **NullableString** | The source record: the filing document id for an extracted claim, the SBTi company id for an SBTi claim. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the claim&#39;s source. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit line for the claim&#39;s source. Show it wherever the claim is shown. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | none for an extracted claim, licensable for an SBTi claim. The view withholds a licensable row from a caller who does not hold the licence. | [optional] 

## Methods

### NewOrgDisclosuresRow

`func NewOrgDisclosuresRow() *OrgDisclosuresRow`

NewOrgDisclosuresRow instantiates a new OrgDisclosuresRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgDisclosuresRowWithDefaults

`func NewOrgDisclosuresRowWithDefaults() *OrgDisclosuresRow`

NewOrgDisclosuresRowWithDefaults instantiates a new OrgDisclosuresRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrgDisclosuresRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgDisclosuresRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgDisclosuresRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrgDisclosuresRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *OrgDisclosuresRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *OrgDisclosuresRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOrgName

`func (o *OrgDisclosuresRow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *OrgDisclosuresRow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *OrgDisclosuresRow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *OrgDisclosuresRow) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *OrgDisclosuresRow) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *OrgDisclosuresRow) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetCompanyId

`func (o *OrgDisclosuresRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *OrgDisclosuresRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *OrgDisclosuresRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *OrgDisclosuresRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *OrgDisclosuresRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *OrgDisclosuresRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *OrgDisclosuresRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrgDisclosuresRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrgDisclosuresRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrgDisclosuresRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *OrgDisclosuresRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *OrgDisclosuresRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *OrgDisclosuresRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *OrgDisclosuresRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *OrgDisclosuresRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *OrgDisclosuresRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *OrgDisclosuresRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *OrgDisclosuresRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *OrgDisclosuresRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *OrgDisclosuresRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *OrgDisclosuresRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *OrgDisclosuresRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *OrgDisclosuresRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *OrgDisclosuresRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetClaimCategory

`func (o *OrgDisclosuresRow) GetClaimCategory() string`

GetClaimCategory returns the ClaimCategory field if non-nil, zero value otherwise.

### GetClaimCategoryOk

`func (o *OrgDisclosuresRow) GetClaimCategoryOk() (*string, bool)`

GetClaimCategoryOk returns a tuple with the ClaimCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimCategory

`func (o *OrgDisclosuresRow) SetClaimCategory(v string)`

SetClaimCategory sets ClaimCategory field to given value.

### HasClaimCategory

`func (o *OrgDisclosuresRow) HasClaimCategory() bool`

HasClaimCategory returns a boolean if a field has been set.

### SetClaimCategoryNil

`func (o *OrgDisclosuresRow) SetClaimCategoryNil(b bool)`

 SetClaimCategoryNil sets the value for ClaimCategory to be an explicit nil

### UnsetClaimCategory
`func (o *OrgDisclosuresRow) UnsetClaimCategory()`

UnsetClaimCategory ensures that no value is present for ClaimCategory, not even an explicit nil
### GetClaimType

`func (o *OrgDisclosuresRow) GetClaimType() string`

GetClaimType returns the ClaimType field if non-nil, zero value otherwise.

### GetClaimTypeOk

`func (o *OrgDisclosuresRow) GetClaimTypeOk() (*string, bool)`

GetClaimTypeOk returns a tuple with the ClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimType

`func (o *OrgDisclosuresRow) SetClaimType(v string)`

SetClaimType sets ClaimType field to given value.

### HasClaimType

`func (o *OrgDisclosuresRow) HasClaimType() bool`

HasClaimType returns a boolean if a field has been set.

### SetClaimTypeNil

`func (o *OrgDisclosuresRow) SetClaimTypeNil(b bool)`

 SetClaimTypeNil sets the value for ClaimType to be an explicit nil

### UnsetClaimType
`func (o *OrgDisclosuresRow) UnsetClaimType()`

UnsetClaimType ensures that no value is present for ClaimType, not even an explicit nil
### GetDisclosureSource

`func (o *OrgDisclosuresRow) GetDisclosureSource() string`

GetDisclosureSource returns the DisclosureSource field if non-nil, zero value otherwise.

### GetDisclosureSourceOk

`func (o *OrgDisclosuresRow) GetDisclosureSourceOk() (*string, bool)`

GetDisclosureSourceOk returns a tuple with the DisclosureSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureSource

`func (o *OrgDisclosuresRow) SetDisclosureSource(v string)`

SetDisclosureSource sets DisclosureSource field to given value.

### HasDisclosureSource

`func (o *OrgDisclosuresRow) HasDisclosureSource() bool`

HasDisclosureSource returns a boolean if a field has been set.

### SetDisclosureSourceNil

`func (o *OrgDisclosuresRow) SetDisclosureSourceNil(b bool)`

 SetDisclosureSourceNil sets the value for DisclosureSource to be an explicit nil

### UnsetDisclosureSource
`func (o *OrgDisclosuresRow) UnsetDisclosureSource()`

UnsetDisclosureSource ensures that no value is present for DisclosureSource, not even an explicit nil
### GetDisclosureState

`func (o *OrgDisclosuresRow) GetDisclosureState() string`

GetDisclosureState returns the DisclosureState field if non-nil, zero value otherwise.

### GetDisclosureStateOk

`func (o *OrgDisclosuresRow) GetDisclosureStateOk() (*string, bool)`

GetDisclosureStateOk returns a tuple with the DisclosureState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureState

`func (o *OrgDisclosuresRow) SetDisclosureState(v string)`

SetDisclosureState sets DisclosureState field to given value.

### HasDisclosureState

`func (o *OrgDisclosuresRow) HasDisclosureState() bool`

HasDisclosureState returns a boolean if a field has been set.

### SetDisclosureStateNil

`func (o *OrgDisclosuresRow) SetDisclosureStateNil(b bool)`

 SetDisclosureStateNil sets the value for DisclosureState to be an explicit nil

### UnsetDisclosureState
`func (o *OrgDisclosuresRow) UnsetDisclosureState()`

UnsetDisclosureState ensures that no value is present for DisclosureState, not even an explicit nil
### GetClaimValue

`func (o *OrgDisclosuresRow) GetClaimValue() string`

GetClaimValue returns the ClaimValue field if non-nil, zero value otherwise.

### GetClaimValueOk

`func (o *OrgDisclosuresRow) GetClaimValueOk() (*string, bool)`

GetClaimValueOk returns a tuple with the ClaimValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimValue

`func (o *OrgDisclosuresRow) SetClaimValue(v string)`

SetClaimValue sets ClaimValue field to given value.

### HasClaimValue

`func (o *OrgDisclosuresRow) HasClaimValue() bool`

HasClaimValue returns a boolean if a field has been set.

### SetClaimValueNil

`func (o *OrgDisclosuresRow) SetClaimValueNil(b bool)`

 SetClaimValueNil sets the value for ClaimValue to be an explicit nil

### UnsetClaimValue
`func (o *OrgDisclosuresRow) UnsetClaimValue()`

UnsetClaimValue ensures that no value is present for ClaimValue, not even an explicit nil
### GetClaimDetail

`func (o *OrgDisclosuresRow) GetClaimDetail() string`

GetClaimDetail returns the ClaimDetail field if non-nil, zero value otherwise.

### GetClaimDetailOk

`func (o *OrgDisclosuresRow) GetClaimDetailOk() (*string, bool)`

GetClaimDetailOk returns a tuple with the ClaimDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimDetail

`func (o *OrgDisclosuresRow) SetClaimDetail(v string)`

SetClaimDetail sets ClaimDetail field to given value.

### HasClaimDetail

`func (o *OrgDisclosuresRow) HasClaimDetail() bool`

HasClaimDetail returns a boolean if a field has been set.

### SetClaimDetailNil

`func (o *OrgDisclosuresRow) SetClaimDetailNil(b bool)`

 SetClaimDetailNil sets the value for ClaimDetail to be an explicit nil

### UnsetClaimDetail
`func (o *OrgDisclosuresRow) UnsetClaimDetail()`

UnsetClaimDetail ensures that no value is present for ClaimDetail, not even an explicit nil
### GetDisclosureDate

`func (o *OrgDisclosuresRow) GetDisclosureDate() string`

GetDisclosureDate returns the DisclosureDate field if non-nil, zero value otherwise.

### GetDisclosureDateOk

`func (o *OrgDisclosuresRow) GetDisclosureDateOk() (*string, bool)`

GetDisclosureDateOk returns a tuple with the DisclosureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureDate

`func (o *OrgDisclosuresRow) SetDisclosureDate(v string)`

SetDisclosureDate sets DisclosureDate field to given value.

### HasDisclosureDate

`func (o *OrgDisclosuresRow) HasDisclosureDate() bool`

HasDisclosureDate returns a boolean if a field has been set.

### SetDisclosureDateNil

`func (o *OrgDisclosuresRow) SetDisclosureDateNil(b bool)`

 SetDisclosureDateNil sets the value for DisclosureDate to be an explicit nil

### UnsetDisclosureDate
`func (o *OrgDisclosuresRow) UnsetDisclosureDate()`

UnsetDisclosureDate ensures that no value is present for DisclosureDate, not even an explicit nil
### GetReportingYear

`func (o *OrgDisclosuresRow) GetReportingYear() string`

GetReportingYear returns the ReportingYear field if non-nil, zero value otherwise.

### GetReportingYearOk

`func (o *OrgDisclosuresRow) GetReportingYearOk() (*string, bool)`

GetReportingYearOk returns a tuple with the ReportingYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingYear

`func (o *OrgDisclosuresRow) SetReportingYear(v string)`

SetReportingYear sets ReportingYear field to given value.

### HasReportingYear

`func (o *OrgDisclosuresRow) HasReportingYear() bool`

HasReportingYear returns a boolean if a field has been set.

### SetReportingYearNil

`func (o *OrgDisclosuresRow) SetReportingYearNil(b bool)`

 SetReportingYearNil sets the value for ReportingYear to be an explicit nil

### UnsetReportingYear
`func (o *OrgDisclosuresRow) UnsetReportingYear()`

UnsetReportingYear ensures that no value is present for ReportingYear, not even an explicit nil
### GetEvidenceQuote

`func (o *OrgDisclosuresRow) GetEvidenceQuote() string`

GetEvidenceQuote returns the EvidenceQuote field if non-nil, zero value otherwise.

### GetEvidenceQuoteOk

`func (o *OrgDisclosuresRow) GetEvidenceQuoteOk() (*string, bool)`

GetEvidenceQuoteOk returns a tuple with the EvidenceQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceQuote

`func (o *OrgDisclosuresRow) SetEvidenceQuote(v string)`

SetEvidenceQuote sets EvidenceQuote field to given value.

### HasEvidenceQuote

`func (o *OrgDisclosuresRow) HasEvidenceQuote() bool`

HasEvidenceQuote returns a boolean if a field has been set.

### SetEvidenceQuoteNil

`func (o *OrgDisclosuresRow) SetEvidenceQuoteNil(b bool)`

 SetEvidenceQuoteNil sets the value for EvidenceQuote to be an explicit nil

### UnsetEvidenceQuote
`func (o *OrgDisclosuresRow) UnsetEvidenceQuote()`

UnsetEvidenceQuote ensures that no value is present for EvidenceQuote, not even an explicit nil
### GetConfidence

`func (o *OrgDisclosuresRow) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *OrgDisclosuresRow) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *OrgDisclosuresRow) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *OrgDisclosuresRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *OrgDisclosuresRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *OrgDisclosuresRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetSourceRecordId

`func (o *OrgDisclosuresRow) GetSourceRecordId() string`

GetSourceRecordId returns the SourceRecordId field if non-nil, zero value otherwise.

### GetSourceRecordIdOk

`func (o *OrgDisclosuresRow) GetSourceRecordIdOk() (*string, bool)`

GetSourceRecordIdOk returns a tuple with the SourceRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRecordId

`func (o *OrgDisclosuresRow) SetSourceRecordId(v string)`

SetSourceRecordId sets SourceRecordId field to given value.

### HasSourceRecordId

`func (o *OrgDisclosuresRow) HasSourceRecordId() bool`

HasSourceRecordId returns a boolean if a field has been set.

### SetSourceRecordIdNil

`func (o *OrgDisclosuresRow) SetSourceRecordIdNil(b bool)`

 SetSourceRecordIdNil sets the value for SourceRecordId to be an explicit nil

### UnsetSourceRecordId
`func (o *OrgDisclosuresRow) UnsetSourceRecordId()`

UnsetSourceRecordId ensures that no value is present for SourceRecordId, not even an explicit nil
### GetLicenseName

`func (o *OrgDisclosuresRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *OrgDisclosuresRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *OrgDisclosuresRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *OrgDisclosuresRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *OrgDisclosuresRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *OrgDisclosuresRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *OrgDisclosuresRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgDisclosuresRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgDisclosuresRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *OrgDisclosuresRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *OrgDisclosuresRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *OrgDisclosuresRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil
### GetRestrictionTier

`func (o *OrgDisclosuresRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *OrgDisclosuresRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *OrgDisclosuresRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *OrgDisclosuresRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *OrgDisclosuresRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *OrgDisclosuresRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


