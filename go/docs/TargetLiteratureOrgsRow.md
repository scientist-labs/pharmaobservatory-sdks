# TargetLiteratureOrgsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. | [optional] 
**OrgId** | Pointer to **NullableString** | Legal-entity level: the resolved parent organization record the byline names. | [optional] 
**OrgName** | Pointer to **NullableString** | The organization&#39;s name. | [optional] 
**OrgType** | Pointer to **NullableString** | The organization&#39;s type. | [optional] 
**OrgCountry** | Pointer to **NullableString** | The organization&#39;s country code (null when not recorded). | [optional] 
**NPapers** | Pointer to **NullableInt32** | Distinct papers naming the organization on a byline. | [optional] 
**NTrialLinkedPapers** | Pointer to **NullableInt32** | The part of n_papers a registered clinical trial cites. | [optional] 
**LatestPubYear** | Pointer to **NullableInt32** | The newest publication year among those papers. | [optional] 
**OrgArm** | Pointer to **NullableString** | company when the organization is company-typed; institution for every other organization. | [optional] 

## Methods

### NewTargetLiteratureOrgsRow

`func NewTargetLiteratureOrgsRow() *TargetLiteratureOrgsRow`

NewTargetLiteratureOrgsRow instantiates a new TargetLiteratureOrgsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetLiteratureOrgsRowWithDefaults

`func NewTargetLiteratureOrgsRowWithDefaults() *TargetLiteratureOrgsRow`

NewTargetLiteratureOrgsRowWithDefaults instantiates a new TargetLiteratureOrgsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetLiteratureOrgsRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetLiteratureOrgsRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetLiteratureOrgsRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetLiteratureOrgsRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetLiteratureOrgsRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetLiteratureOrgsRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetOrgId

`func (o *TargetLiteratureOrgsRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *TargetLiteratureOrgsRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *TargetLiteratureOrgsRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *TargetLiteratureOrgsRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *TargetLiteratureOrgsRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *TargetLiteratureOrgsRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOrgName

`func (o *TargetLiteratureOrgsRow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *TargetLiteratureOrgsRow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *TargetLiteratureOrgsRow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *TargetLiteratureOrgsRow) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *TargetLiteratureOrgsRow) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *TargetLiteratureOrgsRow) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetOrgType

`func (o *TargetLiteratureOrgsRow) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *TargetLiteratureOrgsRow) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *TargetLiteratureOrgsRow) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *TargetLiteratureOrgsRow) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### SetOrgTypeNil

`func (o *TargetLiteratureOrgsRow) SetOrgTypeNil(b bool)`

 SetOrgTypeNil sets the value for OrgType to be an explicit nil

### UnsetOrgType
`func (o *TargetLiteratureOrgsRow) UnsetOrgType()`

UnsetOrgType ensures that no value is present for OrgType, not even an explicit nil
### GetOrgCountry

`func (o *TargetLiteratureOrgsRow) GetOrgCountry() string`

GetOrgCountry returns the OrgCountry field if non-nil, zero value otherwise.

### GetOrgCountryOk

`func (o *TargetLiteratureOrgsRow) GetOrgCountryOk() (*string, bool)`

GetOrgCountryOk returns a tuple with the OrgCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCountry

`func (o *TargetLiteratureOrgsRow) SetOrgCountry(v string)`

SetOrgCountry sets OrgCountry field to given value.

### HasOrgCountry

`func (o *TargetLiteratureOrgsRow) HasOrgCountry() bool`

HasOrgCountry returns a boolean if a field has been set.

### SetOrgCountryNil

`func (o *TargetLiteratureOrgsRow) SetOrgCountryNil(b bool)`

 SetOrgCountryNil sets the value for OrgCountry to be an explicit nil

### UnsetOrgCountry
`func (o *TargetLiteratureOrgsRow) UnsetOrgCountry()`

UnsetOrgCountry ensures that no value is present for OrgCountry, not even an explicit nil
### GetNPapers

`func (o *TargetLiteratureOrgsRow) GetNPapers() int32`

GetNPapers returns the NPapers field if non-nil, zero value otherwise.

### GetNPapersOk

`func (o *TargetLiteratureOrgsRow) GetNPapersOk() (*int32, bool)`

GetNPapersOk returns a tuple with the NPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPapers

`func (o *TargetLiteratureOrgsRow) SetNPapers(v int32)`

SetNPapers sets NPapers field to given value.

### HasNPapers

`func (o *TargetLiteratureOrgsRow) HasNPapers() bool`

HasNPapers returns a boolean if a field has been set.

### SetNPapersNil

`func (o *TargetLiteratureOrgsRow) SetNPapersNil(b bool)`

 SetNPapersNil sets the value for NPapers to be an explicit nil

### UnsetNPapers
`func (o *TargetLiteratureOrgsRow) UnsetNPapers()`

UnsetNPapers ensures that no value is present for NPapers, not even an explicit nil
### GetNTrialLinkedPapers

`func (o *TargetLiteratureOrgsRow) GetNTrialLinkedPapers() int32`

GetNTrialLinkedPapers returns the NTrialLinkedPapers field if non-nil, zero value otherwise.

### GetNTrialLinkedPapersOk

`func (o *TargetLiteratureOrgsRow) GetNTrialLinkedPapersOk() (*int32, bool)`

GetNTrialLinkedPapersOk returns a tuple with the NTrialLinkedPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialLinkedPapers

`func (o *TargetLiteratureOrgsRow) SetNTrialLinkedPapers(v int32)`

SetNTrialLinkedPapers sets NTrialLinkedPapers field to given value.

### HasNTrialLinkedPapers

`func (o *TargetLiteratureOrgsRow) HasNTrialLinkedPapers() bool`

HasNTrialLinkedPapers returns a boolean if a field has been set.

### SetNTrialLinkedPapersNil

`func (o *TargetLiteratureOrgsRow) SetNTrialLinkedPapersNil(b bool)`

 SetNTrialLinkedPapersNil sets the value for NTrialLinkedPapers to be an explicit nil

### UnsetNTrialLinkedPapers
`func (o *TargetLiteratureOrgsRow) UnsetNTrialLinkedPapers()`

UnsetNTrialLinkedPapers ensures that no value is present for NTrialLinkedPapers, not even an explicit nil
### GetLatestPubYear

`func (o *TargetLiteratureOrgsRow) GetLatestPubYear() int32`

GetLatestPubYear returns the LatestPubYear field if non-nil, zero value otherwise.

### GetLatestPubYearOk

`func (o *TargetLiteratureOrgsRow) GetLatestPubYearOk() (*int32, bool)`

GetLatestPubYearOk returns a tuple with the LatestPubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPubYear

`func (o *TargetLiteratureOrgsRow) SetLatestPubYear(v int32)`

SetLatestPubYear sets LatestPubYear field to given value.

### HasLatestPubYear

`func (o *TargetLiteratureOrgsRow) HasLatestPubYear() bool`

HasLatestPubYear returns a boolean if a field has been set.

### SetLatestPubYearNil

`func (o *TargetLiteratureOrgsRow) SetLatestPubYearNil(b bool)`

 SetLatestPubYearNil sets the value for LatestPubYear to be an explicit nil

### UnsetLatestPubYear
`func (o *TargetLiteratureOrgsRow) UnsetLatestPubYear()`

UnsetLatestPubYear ensures that no value is present for LatestPubYear, not even an explicit nil
### GetOrgArm

`func (o *TargetLiteratureOrgsRow) GetOrgArm() string`

GetOrgArm returns the OrgArm field if non-nil, zero value otherwise.

### GetOrgArmOk

`func (o *TargetLiteratureOrgsRow) GetOrgArmOk() (*string, bool)`

GetOrgArmOk returns a tuple with the OrgArm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgArm

`func (o *TargetLiteratureOrgsRow) SetOrgArm(v string)`

SetOrgArm sets OrgArm field to given value.

### HasOrgArm

`func (o *TargetLiteratureOrgsRow) HasOrgArm() bool`

HasOrgArm returns a boolean if a field has been set.

### SetOrgArmNil

`func (o *TargetLiteratureOrgsRow) SetOrgArmNil(b bool)`

 SetOrgArmNil sets the value for OrgArm to be an explicit nil

### UnsetOrgArm
`func (o *TargetLiteratureOrgsRow) UnsetOrgArm()`

UnsetOrgArm ensures that no value is present for OrgArm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


