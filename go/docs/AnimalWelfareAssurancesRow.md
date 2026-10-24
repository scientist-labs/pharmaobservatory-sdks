# AnimalWelfareAssurancesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstitutionOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the OLAW institution resolved to; the entry key for the legal_entities object type. | [optional] 
**InstitutionCompanyId** | Pointer to **NullableString** | Company level: the organization&#39;s company (null when the organization is not in a company); the entry key for the companies object type. | [optional] 
**InstitutionCompanyName** | Pointer to **NullableString** | Display name of the organization&#39;s company. | [optional] 
**InstitutionGroupId** | Pointer to **NullableString** | Group level: the organization&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**InstitutionGroupName** | Pointer to **NullableString** | Display name of the organization&#39;s group. | [optional] 
**Source** | Pointer to **NullableString** | The animal-welfare source: nih_olaw today. A second source (the Canadian Council on Animal Care list) can be added as another value without changing the grain. | [optional] 
**AssuranceId** | Pointer to **NullableString** | The OLAW assurance id of the organization&#39;s best evidence (a primary assurance beats a covered component). | [optional] 
**AssuranceStatusCode** | Pointer to **NullableString** | OLAW&#39;s raw two-digit assurance status code as published (for example 02 or 03). Never decoded, because no OLAW page documents it. | [optional] 
**AssuranceKind** | Pointer to **NullableString** | OLAW&#39;s PRIMARY_OR_COVERED_COMP field as published: PRIMARY or COVERED_COMP (a covered component of a primary assurance). | [optional] 
**IsDomestic** | Pointer to **NullableBool** | True for a domestic (US) assurance, false for a foreign one, from OLAW&#39;s assurance type (D or F). | [optional] 
**MatchTier** | Pointer to **NullableString** | How well the OLAW institution name reached the organization: high (name, country and state agree), medium (name and country agree) or low (name only). The web badge shows high and medium. | [optional] 
**AssuranceCount** | Pointer to **NullableInt64** | How many OLAW assurance rows resolved to this organization (a university and its covered components). | [optional] 
**AssuranceIds** | Pointer to **[]string** | Every OLAW assurance id that resolved to this organization, sorted. | [optional] 
**InstitutionNames** | Pointer to **[]string** | The institution names as OLAW lists them, for every assurance that resolved to this organization. | [optional] 
**SourceAsof** | Pointer to **NullableTime** | When OLAW&#39;s lookup tool says its list was last updated; the source does not state a timezone. | [optional] 
**SourceCredit** | Pointer to **NullableString** | The plain source credit line to show beside any figure from this dataset. | [optional] 

## Methods

### NewAnimalWelfareAssurancesRow

`func NewAnimalWelfareAssurancesRow() *AnimalWelfareAssurancesRow`

NewAnimalWelfareAssurancesRow instantiates a new AnimalWelfareAssurancesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnimalWelfareAssurancesRowWithDefaults

`func NewAnimalWelfareAssurancesRowWithDefaults() *AnimalWelfareAssurancesRow`

NewAnimalWelfareAssurancesRowWithDefaults instantiates a new AnimalWelfareAssurancesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutionOrgId

`func (o *AnimalWelfareAssurancesRow) GetInstitutionOrgId() string`

GetInstitutionOrgId returns the InstitutionOrgId field if non-nil, zero value otherwise.

### GetInstitutionOrgIdOk

`func (o *AnimalWelfareAssurancesRow) GetInstitutionOrgIdOk() (*string, bool)`

GetInstitutionOrgIdOk returns a tuple with the InstitutionOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionOrgId

`func (o *AnimalWelfareAssurancesRow) SetInstitutionOrgId(v string)`

SetInstitutionOrgId sets InstitutionOrgId field to given value.

### HasInstitutionOrgId

`func (o *AnimalWelfareAssurancesRow) HasInstitutionOrgId() bool`

HasInstitutionOrgId returns a boolean if a field has been set.

### SetInstitutionOrgIdNil

`func (o *AnimalWelfareAssurancesRow) SetInstitutionOrgIdNil(b bool)`

 SetInstitutionOrgIdNil sets the value for InstitutionOrgId to be an explicit nil

### UnsetInstitutionOrgId
`func (o *AnimalWelfareAssurancesRow) UnsetInstitutionOrgId()`

UnsetInstitutionOrgId ensures that no value is present for InstitutionOrgId, not even an explicit nil
### GetInstitutionCompanyId

`func (o *AnimalWelfareAssurancesRow) GetInstitutionCompanyId() string`

GetInstitutionCompanyId returns the InstitutionCompanyId field if non-nil, zero value otherwise.

### GetInstitutionCompanyIdOk

`func (o *AnimalWelfareAssurancesRow) GetInstitutionCompanyIdOk() (*string, bool)`

GetInstitutionCompanyIdOk returns a tuple with the InstitutionCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCompanyId

`func (o *AnimalWelfareAssurancesRow) SetInstitutionCompanyId(v string)`

SetInstitutionCompanyId sets InstitutionCompanyId field to given value.

### HasInstitutionCompanyId

`func (o *AnimalWelfareAssurancesRow) HasInstitutionCompanyId() bool`

HasInstitutionCompanyId returns a boolean if a field has been set.

### SetInstitutionCompanyIdNil

`func (o *AnimalWelfareAssurancesRow) SetInstitutionCompanyIdNil(b bool)`

 SetInstitutionCompanyIdNil sets the value for InstitutionCompanyId to be an explicit nil

### UnsetInstitutionCompanyId
`func (o *AnimalWelfareAssurancesRow) UnsetInstitutionCompanyId()`

UnsetInstitutionCompanyId ensures that no value is present for InstitutionCompanyId, not even an explicit nil
### GetInstitutionCompanyName

`func (o *AnimalWelfareAssurancesRow) GetInstitutionCompanyName() string`

GetInstitutionCompanyName returns the InstitutionCompanyName field if non-nil, zero value otherwise.

### GetInstitutionCompanyNameOk

`func (o *AnimalWelfareAssurancesRow) GetInstitutionCompanyNameOk() (*string, bool)`

GetInstitutionCompanyNameOk returns a tuple with the InstitutionCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCompanyName

`func (o *AnimalWelfareAssurancesRow) SetInstitutionCompanyName(v string)`

SetInstitutionCompanyName sets InstitutionCompanyName field to given value.

### HasInstitutionCompanyName

`func (o *AnimalWelfareAssurancesRow) HasInstitutionCompanyName() bool`

HasInstitutionCompanyName returns a boolean if a field has been set.

### SetInstitutionCompanyNameNil

`func (o *AnimalWelfareAssurancesRow) SetInstitutionCompanyNameNil(b bool)`

 SetInstitutionCompanyNameNil sets the value for InstitutionCompanyName to be an explicit nil

### UnsetInstitutionCompanyName
`func (o *AnimalWelfareAssurancesRow) UnsetInstitutionCompanyName()`

UnsetInstitutionCompanyName ensures that no value is present for InstitutionCompanyName, not even an explicit nil
### GetInstitutionGroupId

`func (o *AnimalWelfareAssurancesRow) GetInstitutionGroupId() string`

GetInstitutionGroupId returns the InstitutionGroupId field if non-nil, zero value otherwise.

### GetInstitutionGroupIdOk

`func (o *AnimalWelfareAssurancesRow) GetInstitutionGroupIdOk() (*string, bool)`

GetInstitutionGroupIdOk returns a tuple with the InstitutionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionGroupId

`func (o *AnimalWelfareAssurancesRow) SetInstitutionGroupId(v string)`

SetInstitutionGroupId sets InstitutionGroupId field to given value.

### HasInstitutionGroupId

`func (o *AnimalWelfareAssurancesRow) HasInstitutionGroupId() bool`

HasInstitutionGroupId returns a boolean if a field has been set.

### SetInstitutionGroupIdNil

`func (o *AnimalWelfareAssurancesRow) SetInstitutionGroupIdNil(b bool)`

 SetInstitutionGroupIdNil sets the value for InstitutionGroupId to be an explicit nil

### UnsetInstitutionGroupId
`func (o *AnimalWelfareAssurancesRow) UnsetInstitutionGroupId()`

UnsetInstitutionGroupId ensures that no value is present for InstitutionGroupId, not even an explicit nil
### GetInstitutionGroupName

`func (o *AnimalWelfareAssurancesRow) GetInstitutionGroupName() string`

GetInstitutionGroupName returns the InstitutionGroupName field if non-nil, zero value otherwise.

### GetInstitutionGroupNameOk

`func (o *AnimalWelfareAssurancesRow) GetInstitutionGroupNameOk() (*string, bool)`

GetInstitutionGroupNameOk returns a tuple with the InstitutionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionGroupName

`func (o *AnimalWelfareAssurancesRow) SetInstitutionGroupName(v string)`

SetInstitutionGroupName sets InstitutionGroupName field to given value.

### HasInstitutionGroupName

`func (o *AnimalWelfareAssurancesRow) HasInstitutionGroupName() bool`

HasInstitutionGroupName returns a boolean if a field has been set.

### SetInstitutionGroupNameNil

`func (o *AnimalWelfareAssurancesRow) SetInstitutionGroupNameNil(b bool)`

 SetInstitutionGroupNameNil sets the value for InstitutionGroupName to be an explicit nil

### UnsetInstitutionGroupName
`func (o *AnimalWelfareAssurancesRow) UnsetInstitutionGroupName()`

UnsetInstitutionGroupName ensures that no value is present for InstitutionGroupName, not even an explicit nil
### GetSource

`func (o *AnimalWelfareAssurancesRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AnimalWelfareAssurancesRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AnimalWelfareAssurancesRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AnimalWelfareAssurancesRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AnimalWelfareAssurancesRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AnimalWelfareAssurancesRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetAssuranceId

`func (o *AnimalWelfareAssurancesRow) GetAssuranceId() string`

GetAssuranceId returns the AssuranceId field if non-nil, zero value otherwise.

### GetAssuranceIdOk

`func (o *AnimalWelfareAssurancesRow) GetAssuranceIdOk() (*string, bool)`

GetAssuranceIdOk returns a tuple with the AssuranceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceId

`func (o *AnimalWelfareAssurancesRow) SetAssuranceId(v string)`

SetAssuranceId sets AssuranceId field to given value.

### HasAssuranceId

`func (o *AnimalWelfareAssurancesRow) HasAssuranceId() bool`

HasAssuranceId returns a boolean if a field has been set.

### SetAssuranceIdNil

`func (o *AnimalWelfareAssurancesRow) SetAssuranceIdNil(b bool)`

 SetAssuranceIdNil sets the value for AssuranceId to be an explicit nil

### UnsetAssuranceId
`func (o *AnimalWelfareAssurancesRow) UnsetAssuranceId()`

UnsetAssuranceId ensures that no value is present for AssuranceId, not even an explicit nil
### GetAssuranceStatusCode

`func (o *AnimalWelfareAssurancesRow) GetAssuranceStatusCode() string`

GetAssuranceStatusCode returns the AssuranceStatusCode field if non-nil, zero value otherwise.

### GetAssuranceStatusCodeOk

`func (o *AnimalWelfareAssurancesRow) GetAssuranceStatusCodeOk() (*string, bool)`

GetAssuranceStatusCodeOk returns a tuple with the AssuranceStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceStatusCode

`func (o *AnimalWelfareAssurancesRow) SetAssuranceStatusCode(v string)`

SetAssuranceStatusCode sets AssuranceStatusCode field to given value.

### HasAssuranceStatusCode

`func (o *AnimalWelfareAssurancesRow) HasAssuranceStatusCode() bool`

HasAssuranceStatusCode returns a boolean if a field has been set.

### SetAssuranceStatusCodeNil

`func (o *AnimalWelfareAssurancesRow) SetAssuranceStatusCodeNil(b bool)`

 SetAssuranceStatusCodeNil sets the value for AssuranceStatusCode to be an explicit nil

### UnsetAssuranceStatusCode
`func (o *AnimalWelfareAssurancesRow) UnsetAssuranceStatusCode()`

UnsetAssuranceStatusCode ensures that no value is present for AssuranceStatusCode, not even an explicit nil
### GetAssuranceKind

`func (o *AnimalWelfareAssurancesRow) GetAssuranceKind() string`

GetAssuranceKind returns the AssuranceKind field if non-nil, zero value otherwise.

### GetAssuranceKindOk

`func (o *AnimalWelfareAssurancesRow) GetAssuranceKindOk() (*string, bool)`

GetAssuranceKindOk returns a tuple with the AssuranceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceKind

`func (o *AnimalWelfareAssurancesRow) SetAssuranceKind(v string)`

SetAssuranceKind sets AssuranceKind field to given value.

### HasAssuranceKind

`func (o *AnimalWelfareAssurancesRow) HasAssuranceKind() bool`

HasAssuranceKind returns a boolean if a field has been set.

### SetAssuranceKindNil

`func (o *AnimalWelfareAssurancesRow) SetAssuranceKindNil(b bool)`

 SetAssuranceKindNil sets the value for AssuranceKind to be an explicit nil

### UnsetAssuranceKind
`func (o *AnimalWelfareAssurancesRow) UnsetAssuranceKind()`

UnsetAssuranceKind ensures that no value is present for AssuranceKind, not even an explicit nil
### GetIsDomestic

`func (o *AnimalWelfareAssurancesRow) GetIsDomestic() bool`

GetIsDomestic returns the IsDomestic field if non-nil, zero value otherwise.

### GetIsDomesticOk

`func (o *AnimalWelfareAssurancesRow) GetIsDomesticOk() (*bool, bool)`

GetIsDomesticOk returns a tuple with the IsDomestic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomestic

`func (o *AnimalWelfareAssurancesRow) SetIsDomestic(v bool)`

SetIsDomestic sets IsDomestic field to given value.

### HasIsDomestic

`func (o *AnimalWelfareAssurancesRow) HasIsDomestic() bool`

HasIsDomestic returns a boolean if a field has been set.

### SetIsDomesticNil

`func (o *AnimalWelfareAssurancesRow) SetIsDomesticNil(b bool)`

 SetIsDomesticNil sets the value for IsDomestic to be an explicit nil

### UnsetIsDomestic
`func (o *AnimalWelfareAssurancesRow) UnsetIsDomestic()`

UnsetIsDomestic ensures that no value is present for IsDomestic, not even an explicit nil
### GetMatchTier

`func (o *AnimalWelfareAssurancesRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *AnimalWelfareAssurancesRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *AnimalWelfareAssurancesRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *AnimalWelfareAssurancesRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *AnimalWelfareAssurancesRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *AnimalWelfareAssurancesRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetAssuranceCount

`func (o *AnimalWelfareAssurancesRow) GetAssuranceCount() int64`

GetAssuranceCount returns the AssuranceCount field if non-nil, zero value otherwise.

### GetAssuranceCountOk

`func (o *AnimalWelfareAssurancesRow) GetAssuranceCountOk() (*int64, bool)`

GetAssuranceCountOk returns a tuple with the AssuranceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceCount

`func (o *AnimalWelfareAssurancesRow) SetAssuranceCount(v int64)`

SetAssuranceCount sets AssuranceCount field to given value.

### HasAssuranceCount

`func (o *AnimalWelfareAssurancesRow) HasAssuranceCount() bool`

HasAssuranceCount returns a boolean if a field has been set.

### SetAssuranceCountNil

`func (o *AnimalWelfareAssurancesRow) SetAssuranceCountNil(b bool)`

 SetAssuranceCountNil sets the value for AssuranceCount to be an explicit nil

### UnsetAssuranceCount
`func (o *AnimalWelfareAssurancesRow) UnsetAssuranceCount()`

UnsetAssuranceCount ensures that no value is present for AssuranceCount, not even an explicit nil
### GetAssuranceIds

`func (o *AnimalWelfareAssurancesRow) GetAssuranceIds() []string`

GetAssuranceIds returns the AssuranceIds field if non-nil, zero value otherwise.

### GetAssuranceIdsOk

`func (o *AnimalWelfareAssurancesRow) GetAssuranceIdsOk() (*[]string, bool)`

GetAssuranceIdsOk returns a tuple with the AssuranceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceIds

`func (o *AnimalWelfareAssurancesRow) SetAssuranceIds(v []string)`

SetAssuranceIds sets AssuranceIds field to given value.

### HasAssuranceIds

`func (o *AnimalWelfareAssurancesRow) HasAssuranceIds() bool`

HasAssuranceIds returns a boolean if a field has been set.

### SetAssuranceIdsNil

`func (o *AnimalWelfareAssurancesRow) SetAssuranceIdsNil(b bool)`

 SetAssuranceIdsNil sets the value for AssuranceIds to be an explicit nil

### UnsetAssuranceIds
`func (o *AnimalWelfareAssurancesRow) UnsetAssuranceIds()`

UnsetAssuranceIds ensures that no value is present for AssuranceIds, not even an explicit nil
### GetInstitutionNames

`func (o *AnimalWelfareAssurancesRow) GetInstitutionNames() []string`

GetInstitutionNames returns the InstitutionNames field if non-nil, zero value otherwise.

### GetInstitutionNamesOk

`func (o *AnimalWelfareAssurancesRow) GetInstitutionNamesOk() (*[]string, bool)`

GetInstitutionNamesOk returns a tuple with the InstitutionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionNames

`func (o *AnimalWelfareAssurancesRow) SetInstitutionNames(v []string)`

SetInstitutionNames sets InstitutionNames field to given value.

### HasInstitutionNames

`func (o *AnimalWelfareAssurancesRow) HasInstitutionNames() bool`

HasInstitutionNames returns a boolean if a field has been set.

### SetInstitutionNamesNil

`func (o *AnimalWelfareAssurancesRow) SetInstitutionNamesNil(b bool)`

 SetInstitutionNamesNil sets the value for InstitutionNames to be an explicit nil

### UnsetInstitutionNames
`func (o *AnimalWelfareAssurancesRow) UnsetInstitutionNames()`

UnsetInstitutionNames ensures that no value is present for InstitutionNames, not even an explicit nil
### GetSourceAsof

`func (o *AnimalWelfareAssurancesRow) GetSourceAsof() time.Time`

GetSourceAsof returns the SourceAsof field if non-nil, zero value otherwise.

### GetSourceAsofOk

`func (o *AnimalWelfareAssurancesRow) GetSourceAsofOk() (*time.Time, bool)`

GetSourceAsofOk returns a tuple with the SourceAsof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAsof

`func (o *AnimalWelfareAssurancesRow) SetSourceAsof(v time.Time)`

SetSourceAsof sets SourceAsof field to given value.

### HasSourceAsof

`func (o *AnimalWelfareAssurancesRow) HasSourceAsof() bool`

HasSourceAsof returns a boolean if a field has been set.

### SetSourceAsofNil

`func (o *AnimalWelfareAssurancesRow) SetSourceAsofNil(b bool)`

 SetSourceAsofNil sets the value for SourceAsof to be an explicit nil

### UnsetSourceAsof
`func (o *AnimalWelfareAssurancesRow) UnsetSourceAsof()`

UnsetSourceAsof ensures that no value is present for SourceAsof, not even an explicit nil
### GetSourceCredit

`func (o *AnimalWelfareAssurancesRow) GetSourceCredit() string`

GetSourceCredit returns the SourceCredit field if non-nil, zero value otherwise.

### GetSourceCreditOk

`func (o *AnimalWelfareAssurancesRow) GetSourceCreditOk() (*string, bool)`

GetSourceCreditOk returns a tuple with the SourceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCredit

`func (o *AnimalWelfareAssurancesRow) SetSourceCredit(v string)`

SetSourceCredit sets SourceCredit field to given value.

### HasSourceCredit

`func (o *AnimalWelfareAssurancesRow) HasSourceCredit() bool`

HasSourceCredit returns a boolean if a field has been set.

### SetSourceCreditNil

`func (o *AnimalWelfareAssurancesRow) SetSourceCreditNil(b bool)`

 SetSourceCreditNil sets the value for SourceCredit to be an explicit nil

### UnsetSourceCredit
`func (o *AnimalWelfareAssurancesRow) UnsetSourceCredit()`

UnsetSourceCredit ensures that no value is present for SourceCredit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


