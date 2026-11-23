# ConditionLiteratureOrgsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**OrgId** | Pointer to **NullableString** | Legal-entity level: the resolved parent organization record the byline names. | [optional] 
**OrgName** | Pointer to **NullableString** | The organization&#39;s name. | [optional] 
**OrgType** | Pointer to **NullableString** | The organization&#39;s type. | [optional] 
**OrgCountry** | Pointer to **NullableString** | The organization&#39;s country code (null when not recorded). | [optional] 
**NPapers** | Pointer to **NullableInt32** | Distinct papers naming the organization on a byline. | [optional] 
**NTrialLinkedPapers** | Pointer to **NullableInt32** | The part of n_papers a registered clinical trial cites. | [optional] 
**LatestPubYear** | Pointer to **NullableInt32** | The newest publication year among those papers. | [optional] 
**OrgArm** | Pointer to **NullableString** | company when the organization is company-typed; institution for every other organization. | [optional] 

## Methods

### NewConditionLiteratureOrgsRow

`func NewConditionLiteratureOrgsRow() *ConditionLiteratureOrgsRow`

NewConditionLiteratureOrgsRow instantiates a new ConditionLiteratureOrgsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionLiteratureOrgsRowWithDefaults

`func NewConditionLiteratureOrgsRowWithDefaults() *ConditionLiteratureOrgsRow`

NewConditionLiteratureOrgsRowWithDefaults instantiates a new ConditionLiteratureOrgsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionLiteratureOrgsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionLiteratureOrgsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionLiteratureOrgsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionLiteratureOrgsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionLiteratureOrgsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionLiteratureOrgsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetOrgId

`func (o *ConditionLiteratureOrgsRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ConditionLiteratureOrgsRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ConditionLiteratureOrgsRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ConditionLiteratureOrgsRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *ConditionLiteratureOrgsRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *ConditionLiteratureOrgsRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOrgName

`func (o *ConditionLiteratureOrgsRow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ConditionLiteratureOrgsRow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ConditionLiteratureOrgsRow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ConditionLiteratureOrgsRow) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *ConditionLiteratureOrgsRow) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *ConditionLiteratureOrgsRow) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetOrgType

`func (o *ConditionLiteratureOrgsRow) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *ConditionLiteratureOrgsRow) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *ConditionLiteratureOrgsRow) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *ConditionLiteratureOrgsRow) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### SetOrgTypeNil

`func (o *ConditionLiteratureOrgsRow) SetOrgTypeNil(b bool)`

 SetOrgTypeNil sets the value for OrgType to be an explicit nil

### UnsetOrgType
`func (o *ConditionLiteratureOrgsRow) UnsetOrgType()`

UnsetOrgType ensures that no value is present for OrgType, not even an explicit nil
### GetOrgCountry

`func (o *ConditionLiteratureOrgsRow) GetOrgCountry() string`

GetOrgCountry returns the OrgCountry field if non-nil, zero value otherwise.

### GetOrgCountryOk

`func (o *ConditionLiteratureOrgsRow) GetOrgCountryOk() (*string, bool)`

GetOrgCountryOk returns a tuple with the OrgCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCountry

`func (o *ConditionLiteratureOrgsRow) SetOrgCountry(v string)`

SetOrgCountry sets OrgCountry field to given value.

### HasOrgCountry

`func (o *ConditionLiteratureOrgsRow) HasOrgCountry() bool`

HasOrgCountry returns a boolean if a field has been set.

### SetOrgCountryNil

`func (o *ConditionLiteratureOrgsRow) SetOrgCountryNil(b bool)`

 SetOrgCountryNil sets the value for OrgCountry to be an explicit nil

### UnsetOrgCountry
`func (o *ConditionLiteratureOrgsRow) UnsetOrgCountry()`

UnsetOrgCountry ensures that no value is present for OrgCountry, not even an explicit nil
### GetNPapers

`func (o *ConditionLiteratureOrgsRow) GetNPapers() int32`

GetNPapers returns the NPapers field if non-nil, zero value otherwise.

### GetNPapersOk

`func (o *ConditionLiteratureOrgsRow) GetNPapersOk() (*int32, bool)`

GetNPapersOk returns a tuple with the NPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPapers

`func (o *ConditionLiteratureOrgsRow) SetNPapers(v int32)`

SetNPapers sets NPapers field to given value.

### HasNPapers

`func (o *ConditionLiteratureOrgsRow) HasNPapers() bool`

HasNPapers returns a boolean if a field has been set.

### SetNPapersNil

`func (o *ConditionLiteratureOrgsRow) SetNPapersNil(b bool)`

 SetNPapersNil sets the value for NPapers to be an explicit nil

### UnsetNPapers
`func (o *ConditionLiteratureOrgsRow) UnsetNPapers()`

UnsetNPapers ensures that no value is present for NPapers, not even an explicit nil
### GetNTrialLinkedPapers

`func (o *ConditionLiteratureOrgsRow) GetNTrialLinkedPapers() int32`

GetNTrialLinkedPapers returns the NTrialLinkedPapers field if non-nil, zero value otherwise.

### GetNTrialLinkedPapersOk

`func (o *ConditionLiteratureOrgsRow) GetNTrialLinkedPapersOk() (*int32, bool)`

GetNTrialLinkedPapersOk returns a tuple with the NTrialLinkedPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialLinkedPapers

`func (o *ConditionLiteratureOrgsRow) SetNTrialLinkedPapers(v int32)`

SetNTrialLinkedPapers sets NTrialLinkedPapers field to given value.

### HasNTrialLinkedPapers

`func (o *ConditionLiteratureOrgsRow) HasNTrialLinkedPapers() bool`

HasNTrialLinkedPapers returns a boolean if a field has been set.

### SetNTrialLinkedPapersNil

`func (o *ConditionLiteratureOrgsRow) SetNTrialLinkedPapersNil(b bool)`

 SetNTrialLinkedPapersNil sets the value for NTrialLinkedPapers to be an explicit nil

### UnsetNTrialLinkedPapers
`func (o *ConditionLiteratureOrgsRow) UnsetNTrialLinkedPapers()`

UnsetNTrialLinkedPapers ensures that no value is present for NTrialLinkedPapers, not even an explicit nil
### GetLatestPubYear

`func (o *ConditionLiteratureOrgsRow) GetLatestPubYear() int32`

GetLatestPubYear returns the LatestPubYear field if non-nil, zero value otherwise.

### GetLatestPubYearOk

`func (o *ConditionLiteratureOrgsRow) GetLatestPubYearOk() (*int32, bool)`

GetLatestPubYearOk returns a tuple with the LatestPubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPubYear

`func (o *ConditionLiteratureOrgsRow) SetLatestPubYear(v int32)`

SetLatestPubYear sets LatestPubYear field to given value.

### HasLatestPubYear

`func (o *ConditionLiteratureOrgsRow) HasLatestPubYear() bool`

HasLatestPubYear returns a boolean if a field has been set.

### SetLatestPubYearNil

`func (o *ConditionLiteratureOrgsRow) SetLatestPubYearNil(b bool)`

 SetLatestPubYearNil sets the value for LatestPubYear to be an explicit nil

### UnsetLatestPubYear
`func (o *ConditionLiteratureOrgsRow) UnsetLatestPubYear()`

UnsetLatestPubYear ensures that no value is present for LatestPubYear, not even an explicit nil
### GetOrgArm

`func (o *ConditionLiteratureOrgsRow) GetOrgArm() string`

GetOrgArm returns the OrgArm field if non-nil, zero value otherwise.

### GetOrgArmOk

`func (o *ConditionLiteratureOrgsRow) GetOrgArmOk() (*string, bool)`

GetOrgArmOk returns a tuple with the OrgArm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgArm

`func (o *ConditionLiteratureOrgsRow) SetOrgArm(v string)`

SetOrgArm sets OrgArm field to given value.

### HasOrgArm

`func (o *ConditionLiteratureOrgsRow) HasOrgArm() bool`

HasOrgArm returns a boolean if a field has been set.

### SetOrgArmNil

`func (o *ConditionLiteratureOrgsRow) SetOrgArmNil(b bool)`

 SetOrgArmNil sets the value for OrgArm to be an explicit nil

### UnsetOrgArm
`func (o *ConditionLiteratureOrgsRow) UnsetOrgArm()`

UnsetOrgArm ensures that no value is present for OrgArm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


