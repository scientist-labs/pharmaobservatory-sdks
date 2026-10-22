# BrandCompaniesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**BrandKey** | Pointer to **NullableString** | The brand key: an upper-case alphanumeric trade-name token (the brand&#39;s leading name tokens), the same key the brands pages use. Not a display name; brand_name is the display form. | [optional] 
**CompanyRole** | Pointer to **NullableString** | The applicant&#39;s role for the brand: approval_holder or generic_applicant. | [optional] 
**ApplicantNameKey** | Pointer to **NullableString** | The normalised applicant name, the grain key that separates two spellings of one applicant. | [optional] 
**ApplicantName** | Pointer to **NullableString** | The applicant name FDA published, as filed. | [optional] 
**ApplicantOrgId** | Pointer to **NullableString** | Legal-entity level: the applicant&#39;s organization record (null when unresolved). | [optional] 
**ApplicantCompanyId** | Pointer to **NullableString** | Company level: the applicant&#39;s company (null when unresolved). | [optional] 
**ApplicantCompanyName** | Pointer to **NullableString** | Display name of the applicant&#39;s company. | [optional] 
**ApplicantGroupId** | Pointer to **NullableString** | Group level: the applicant&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**ApplicantGroupName** | Pointer to **NullableString** | Display name of the applicant&#39;s group. | [optional] 
**MatchScope** | Pointer to **NullableString** | How firmly the applicant name resolved to the organization (spine_unique, anchored and establishment_registry are folded onto an organization the spine already held; anchored_scoped, registry_scoped, alias_scoped and none are newly minted organizations). Null when the applicant did not resolve. | [optional] 
**NApplications** | Pointer to **NullableInt32** | Distinct FDA applications the applicant holds for the brand. | [optional] 
**FirstApprovalDate** | Pointer to **NullableString** | The earliest approval date among those applications. | [optional] 
**LatestApprovalDate** | Pointer to **NullableString** | The latest approval date among those applications. | [optional] 
**IsMarketedNow** | Pointer to **NullableBool** | True when at least one of the applicant&#39;s products is currently marketed. | [optional] 
**SourceSystems** | Pointer to **[]string** | The FDA records the applicant was read from (Drugs@FDA, Orange Book, Purple Book). | [optional] 
**AsOf** | Pointer to **NullableString** | The ingest snapshot date the row was read from. | [optional] 

## Methods

### NewBrandCompaniesRow

`func NewBrandCompaniesRow() *BrandCompaniesRow`

NewBrandCompaniesRow instantiates a new BrandCompaniesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandCompaniesRowWithDefaults

`func NewBrandCompaniesRowWithDefaults() *BrandCompaniesRow`

NewBrandCompaniesRowWithDefaults instantiates a new BrandCompaniesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *BrandCompaniesRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *BrandCompaniesRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *BrandCompaniesRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *BrandCompaniesRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *BrandCompaniesRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *BrandCompaniesRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetBrandKey

`func (o *BrandCompaniesRow) GetBrandKey() string`

GetBrandKey returns the BrandKey field if non-nil, zero value otherwise.

### GetBrandKeyOk

`func (o *BrandCompaniesRow) GetBrandKeyOk() (*string, bool)`

GetBrandKeyOk returns a tuple with the BrandKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandKey

`func (o *BrandCompaniesRow) SetBrandKey(v string)`

SetBrandKey sets BrandKey field to given value.

### HasBrandKey

`func (o *BrandCompaniesRow) HasBrandKey() bool`

HasBrandKey returns a boolean if a field has been set.

### SetBrandKeyNil

`func (o *BrandCompaniesRow) SetBrandKeyNil(b bool)`

 SetBrandKeyNil sets the value for BrandKey to be an explicit nil

### UnsetBrandKey
`func (o *BrandCompaniesRow) UnsetBrandKey()`

UnsetBrandKey ensures that no value is present for BrandKey, not even an explicit nil
### GetCompanyRole

`func (o *BrandCompaniesRow) GetCompanyRole() string`

GetCompanyRole returns the CompanyRole field if non-nil, zero value otherwise.

### GetCompanyRoleOk

`func (o *BrandCompaniesRow) GetCompanyRoleOk() (*string, bool)`

GetCompanyRoleOk returns a tuple with the CompanyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyRole

`func (o *BrandCompaniesRow) SetCompanyRole(v string)`

SetCompanyRole sets CompanyRole field to given value.

### HasCompanyRole

`func (o *BrandCompaniesRow) HasCompanyRole() bool`

HasCompanyRole returns a boolean if a field has been set.

### SetCompanyRoleNil

`func (o *BrandCompaniesRow) SetCompanyRoleNil(b bool)`

 SetCompanyRoleNil sets the value for CompanyRole to be an explicit nil

### UnsetCompanyRole
`func (o *BrandCompaniesRow) UnsetCompanyRole()`

UnsetCompanyRole ensures that no value is present for CompanyRole, not even an explicit nil
### GetApplicantNameKey

`func (o *BrandCompaniesRow) GetApplicantNameKey() string`

GetApplicantNameKey returns the ApplicantNameKey field if non-nil, zero value otherwise.

### GetApplicantNameKeyOk

`func (o *BrandCompaniesRow) GetApplicantNameKeyOk() (*string, bool)`

GetApplicantNameKeyOk returns a tuple with the ApplicantNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantNameKey

`func (o *BrandCompaniesRow) SetApplicantNameKey(v string)`

SetApplicantNameKey sets ApplicantNameKey field to given value.

### HasApplicantNameKey

`func (o *BrandCompaniesRow) HasApplicantNameKey() bool`

HasApplicantNameKey returns a boolean if a field has been set.

### SetApplicantNameKeyNil

`func (o *BrandCompaniesRow) SetApplicantNameKeyNil(b bool)`

 SetApplicantNameKeyNil sets the value for ApplicantNameKey to be an explicit nil

### UnsetApplicantNameKey
`func (o *BrandCompaniesRow) UnsetApplicantNameKey()`

UnsetApplicantNameKey ensures that no value is present for ApplicantNameKey, not even an explicit nil
### GetApplicantName

`func (o *BrandCompaniesRow) GetApplicantName() string`

GetApplicantName returns the ApplicantName field if non-nil, zero value otherwise.

### GetApplicantNameOk

`func (o *BrandCompaniesRow) GetApplicantNameOk() (*string, bool)`

GetApplicantNameOk returns a tuple with the ApplicantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantName

`func (o *BrandCompaniesRow) SetApplicantName(v string)`

SetApplicantName sets ApplicantName field to given value.

### HasApplicantName

`func (o *BrandCompaniesRow) HasApplicantName() bool`

HasApplicantName returns a boolean if a field has been set.

### SetApplicantNameNil

`func (o *BrandCompaniesRow) SetApplicantNameNil(b bool)`

 SetApplicantNameNil sets the value for ApplicantName to be an explicit nil

### UnsetApplicantName
`func (o *BrandCompaniesRow) UnsetApplicantName()`

UnsetApplicantName ensures that no value is present for ApplicantName, not even an explicit nil
### GetApplicantOrgId

`func (o *BrandCompaniesRow) GetApplicantOrgId() string`

GetApplicantOrgId returns the ApplicantOrgId field if non-nil, zero value otherwise.

### GetApplicantOrgIdOk

`func (o *BrandCompaniesRow) GetApplicantOrgIdOk() (*string, bool)`

GetApplicantOrgIdOk returns a tuple with the ApplicantOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantOrgId

`func (o *BrandCompaniesRow) SetApplicantOrgId(v string)`

SetApplicantOrgId sets ApplicantOrgId field to given value.

### HasApplicantOrgId

`func (o *BrandCompaniesRow) HasApplicantOrgId() bool`

HasApplicantOrgId returns a boolean if a field has been set.

### SetApplicantOrgIdNil

`func (o *BrandCompaniesRow) SetApplicantOrgIdNil(b bool)`

 SetApplicantOrgIdNil sets the value for ApplicantOrgId to be an explicit nil

### UnsetApplicantOrgId
`func (o *BrandCompaniesRow) UnsetApplicantOrgId()`

UnsetApplicantOrgId ensures that no value is present for ApplicantOrgId, not even an explicit nil
### GetApplicantCompanyId

`func (o *BrandCompaniesRow) GetApplicantCompanyId() string`

GetApplicantCompanyId returns the ApplicantCompanyId field if non-nil, zero value otherwise.

### GetApplicantCompanyIdOk

`func (o *BrandCompaniesRow) GetApplicantCompanyIdOk() (*string, bool)`

GetApplicantCompanyIdOk returns a tuple with the ApplicantCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantCompanyId

`func (o *BrandCompaniesRow) SetApplicantCompanyId(v string)`

SetApplicantCompanyId sets ApplicantCompanyId field to given value.

### HasApplicantCompanyId

`func (o *BrandCompaniesRow) HasApplicantCompanyId() bool`

HasApplicantCompanyId returns a boolean if a field has been set.

### SetApplicantCompanyIdNil

`func (o *BrandCompaniesRow) SetApplicantCompanyIdNil(b bool)`

 SetApplicantCompanyIdNil sets the value for ApplicantCompanyId to be an explicit nil

### UnsetApplicantCompanyId
`func (o *BrandCompaniesRow) UnsetApplicantCompanyId()`

UnsetApplicantCompanyId ensures that no value is present for ApplicantCompanyId, not even an explicit nil
### GetApplicantCompanyName

`func (o *BrandCompaniesRow) GetApplicantCompanyName() string`

GetApplicantCompanyName returns the ApplicantCompanyName field if non-nil, zero value otherwise.

### GetApplicantCompanyNameOk

`func (o *BrandCompaniesRow) GetApplicantCompanyNameOk() (*string, bool)`

GetApplicantCompanyNameOk returns a tuple with the ApplicantCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantCompanyName

`func (o *BrandCompaniesRow) SetApplicantCompanyName(v string)`

SetApplicantCompanyName sets ApplicantCompanyName field to given value.

### HasApplicantCompanyName

`func (o *BrandCompaniesRow) HasApplicantCompanyName() bool`

HasApplicantCompanyName returns a boolean if a field has been set.

### SetApplicantCompanyNameNil

`func (o *BrandCompaniesRow) SetApplicantCompanyNameNil(b bool)`

 SetApplicantCompanyNameNil sets the value for ApplicantCompanyName to be an explicit nil

### UnsetApplicantCompanyName
`func (o *BrandCompaniesRow) UnsetApplicantCompanyName()`

UnsetApplicantCompanyName ensures that no value is present for ApplicantCompanyName, not even an explicit nil
### GetApplicantGroupId

`func (o *BrandCompaniesRow) GetApplicantGroupId() string`

GetApplicantGroupId returns the ApplicantGroupId field if non-nil, zero value otherwise.

### GetApplicantGroupIdOk

`func (o *BrandCompaniesRow) GetApplicantGroupIdOk() (*string, bool)`

GetApplicantGroupIdOk returns a tuple with the ApplicantGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantGroupId

`func (o *BrandCompaniesRow) SetApplicantGroupId(v string)`

SetApplicantGroupId sets ApplicantGroupId field to given value.

### HasApplicantGroupId

`func (o *BrandCompaniesRow) HasApplicantGroupId() bool`

HasApplicantGroupId returns a boolean if a field has been set.

### SetApplicantGroupIdNil

`func (o *BrandCompaniesRow) SetApplicantGroupIdNil(b bool)`

 SetApplicantGroupIdNil sets the value for ApplicantGroupId to be an explicit nil

### UnsetApplicantGroupId
`func (o *BrandCompaniesRow) UnsetApplicantGroupId()`

UnsetApplicantGroupId ensures that no value is present for ApplicantGroupId, not even an explicit nil
### GetApplicantGroupName

`func (o *BrandCompaniesRow) GetApplicantGroupName() string`

GetApplicantGroupName returns the ApplicantGroupName field if non-nil, zero value otherwise.

### GetApplicantGroupNameOk

`func (o *BrandCompaniesRow) GetApplicantGroupNameOk() (*string, bool)`

GetApplicantGroupNameOk returns a tuple with the ApplicantGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantGroupName

`func (o *BrandCompaniesRow) SetApplicantGroupName(v string)`

SetApplicantGroupName sets ApplicantGroupName field to given value.

### HasApplicantGroupName

`func (o *BrandCompaniesRow) HasApplicantGroupName() bool`

HasApplicantGroupName returns a boolean if a field has been set.

### SetApplicantGroupNameNil

`func (o *BrandCompaniesRow) SetApplicantGroupNameNil(b bool)`

 SetApplicantGroupNameNil sets the value for ApplicantGroupName to be an explicit nil

### UnsetApplicantGroupName
`func (o *BrandCompaniesRow) UnsetApplicantGroupName()`

UnsetApplicantGroupName ensures that no value is present for ApplicantGroupName, not even an explicit nil
### GetMatchScope

`func (o *BrandCompaniesRow) GetMatchScope() string`

GetMatchScope returns the MatchScope field if non-nil, zero value otherwise.

### GetMatchScopeOk

`func (o *BrandCompaniesRow) GetMatchScopeOk() (*string, bool)`

GetMatchScopeOk returns a tuple with the MatchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchScope

`func (o *BrandCompaniesRow) SetMatchScope(v string)`

SetMatchScope sets MatchScope field to given value.

### HasMatchScope

`func (o *BrandCompaniesRow) HasMatchScope() bool`

HasMatchScope returns a boolean if a field has been set.

### SetMatchScopeNil

`func (o *BrandCompaniesRow) SetMatchScopeNil(b bool)`

 SetMatchScopeNil sets the value for MatchScope to be an explicit nil

### UnsetMatchScope
`func (o *BrandCompaniesRow) UnsetMatchScope()`

UnsetMatchScope ensures that no value is present for MatchScope, not even an explicit nil
### GetNApplications

`func (o *BrandCompaniesRow) GetNApplications() int32`

GetNApplications returns the NApplications field if non-nil, zero value otherwise.

### GetNApplicationsOk

`func (o *BrandCompaniesRow) GetNApplicationsOk() (*int32, bool)`

GetNApplicationsOk returns a tuple with the NApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNApplications

`func (o *BrandCompaniesRow) SetNApplications(v int32)`

SetNApplications sets NApplications field to given value.

### HasNApplications

`func (o *BrandCompaniesRow) HasNApplications() bool`

HasNApplications returns a boolean if a field has been set.

### SetNApplicationsNil

`func (o *BrandCompaniesRow) SetNApplicationsNil(b bool)`

 SetNApplicationsNil sets the value for NApplications to be an explicit nil

### UnsetNApplications
`func (o *BrandCompaniesRow) UnsetNApplications()`

UnsetNApplications ensures that no value is present for NApplications, not even an explicit nil
### GetFirstApprovalDate

`func (o *BrandCompaniesRow) GetFirstApprovalDate() string`

GetFirstApprovalDate returns the FirstApprovalDate field if non-nil, zero value otherwise.

### GetFirstApprovalDateOk

`func (o *BrandCompaniesRow) GetFirstApprovalDateOk() (*string, bool)`

GetFirstApprovalDateOk returns a tuple with the FirstApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstApprovalDate

`func (o *BrandCompaniesRow) SetFirstApprovalDate(v string)`

SetFirstApprovalDate sets FirstApprovalDate field to given value.

### HasFirstApprovalDate

`func (o *BrandCompaniesRow) HasFirstApprovalDate() bool`

HasFirstApprovalDate returns a boolean if a field has been set.

### SetFirstApprovalDateNil

`func (o *BrandCompaniesRow) SetFirstApprovalDateNil(b bool)`

 SetFirstApprovalDateNil sets the value for FirstApprovalDate to be an explicit nil

### UnsetFirstApprovalDate
`func (o *BrandCompaniesRow) UnsetFirstApprovalDate()`

UnsetFirstApprovalDate ensures that no value is present for FirstApprovalDate, not even an explicit nil
### GetLatestApprovalDate

`func (o *BrandCompaniesRow) GetLatestApprovalDate() string`

GetLatestApprovalDate returns the LatestApprovalDate field if non-nil, zero value otherwise.

### GetLatestApprovalDateOk

`func (o *BrandCompaniesRow) GetLatestApprovalDateOk() (*string, bool)`

GetLatestApprovalDateOk returns a tuple with the LatestApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestApprovalDate

`func (o *BrandCompaniesRow) SetLatestApprovalDate(v string)`

SetLatestApprovalDate sets LatestApprovalDate field to given value.

### HasLatestApprovalDate

`func (o *BrandCompaniesRow) HasLatestApprovalDate() bool`

HasLatestApprovalDate returns a boolean if a field has been set.

### SetLatestApprovalDateNil

`func (o *BrandCompaniesRow) SetLatestApprovalDateNil(b bool)`

 SetLatestApprovalDateNil sets the value for LatestApprovalDate to be an explicit nil

### UnsetLatestApprovalDate
`func (o *BrandCompaniesRow) UnsetLatestApprovalDate()`

UnsetLatestApprovalDate ensures that no value is present for LatestApprovalDate, not even an explicit nil
### GetIsMarketedNow

`func (o *BrandCompaniesRow) GetIsMarketedNow() bool`

GetIsMarketedNow returns the IsMarketedNow field if non-nil, zero value otherwise.

### GetIsMarketedNowOk

`func (o *BrandCompaniesRow) GetIsMarketedNowOk() (*bool, bool)`

GetIsMarketedNowOk returns a tuple with the IsMarketedNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketedNow

`func (o *BrandCompaniesRow) SetIsMarketedNow(v bool)`

SetIsMarketedNow sets IsMarketedNow field to given value.

### HasIsMarketedNow

`func (o *BrandCompaniesRow) HasIsMarketedNow() bool`

HasIsMarketedNow returns a boolean if a field has been set.

### SetIsMarketedNowNil

`func (o *BrandCompaniesRow) SetIsMarketedNowNil(b bool)`

 SetIsMarketedNowNil sets the value for IsMarketedNow to be an explicit nil

### UnsetIsMarketedNow
`func (o *BrandCompaniesRow) UnsetIsMarketedNow()`

UnsetIsMarketedNow ensures that no value is present for IsMarketedNow, not even an explicit nil
### GetSourceSystems

`func (o *BrandCompaniesRow) GetSourceSystems() []string`

GetSourceSystems returns the SourceSystems field if non-nil, zero value otherwise.

### GetSourceSystemsOk

`func (o *BrandCompaniesRow) GetSourceSystemsOk() (*[]string, bool)`

GetSourceSystemsOk returns a tuple with the SourceSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystems

`func (o *BrandCompaniesRow) SetSourceSystems(v []string)`

SetSourceSystems sets SourceSystems field to given value.

### HasSourceSystems

`func (o *BrandCompaniesRow) HasSourceSystems() bool`

HasSourceSystems returns a boolean if a field has been set.

### SetSourceSystemsNil

`func (o *BrandCompaniesRow) SetSourceSystemsNil(b bool)`

 SetSourceSystemsNil sets the value for SourceSystems to be an explicit nil

### UnsetSourceSystems
`func (o *BrandCompaniesRow) UnsetSourceSystems()`

UnsetSourceSystems ensures that no value is present for SourceSystems, not even an explicit nil
### GetAsOf

`func (o *BrandCompaniesRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *BrandCompaniesRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *BrandCompaniesRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *BrandCompaniesRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *BrandCompaniesRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *BrandCompaniesRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


