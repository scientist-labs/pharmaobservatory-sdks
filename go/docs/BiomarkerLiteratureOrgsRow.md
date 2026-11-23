# BiomarkerLiteratureOrgsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The route key of the gene, the entry key for the biomarkers object type. | [optional] 
**OrgId** | Pointer to **NullableString** | Legal-entity level: the resolved parent organization record the byline names. | [optional] 
**OrgName** | Pointer to **NullableString** | The organization&#39;s name. | [optional] 
**OrgType** | Pointer to **NullableString** | The organization&#39;s type. | [optional] 
**OrgCountry** | Pointer to **NullableString** | The organization&#39;s country code (null when not recorded). | [optional] 
**NPapers** | Pointer to **NullableInt32** | Distinct papers naming the organization on a byline. | [optional] 
**NTrialLinkedPapers** | Pointer to **NullableInt32** | The part of n_papers a registered clinical trial cites. | [optional] 
**LatestPubYear** | Pointer to **NullableInt32** | The newest publication year among those papers. | [optional] 
**OrgArm** | Pointer to **NullableString** | company when the organization is company-typed; institution for every other organization. | [optional] 

## Methods

### NewBiomarkerLiteratureOrgsRow

`func NewBiomarkerLiteratureOrgsRow() *BiomarkerLiteratureOrgsRow`

NewBiomarkerLiteratureOrgsRow instantiates a new BiomarkerLiteratureOrgsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerLiteratureOrgsRowWithDefaults

`func NewBiomarkerLiteratureOrgsRowWithDefaults() *BiomarkerLiteratureOrgsRow`

NewBiomarkerLiteratureOrgsRowWithDefaults instantiates a new BiomarkerLiteratureOrgsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerLiteratureOrgsRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerLiteratureOrgsRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerLiteratureOrgsRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerLiteratureOrgsRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerLiteratureOrgsRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerLiteratureOrgsRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetOrgId

`func (o *BiomarkerLiteratureOrgsRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *BiomarkerLiteratureOrgsRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *BiomarkerLiteratureOrgsRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *BiomarkerLiteratureOrgsRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *BiomarkerLiteratureOrgsRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *BiomarkerLiteratureOrgsRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOrgName

`func (o *BiomarkerLiteratureOrgsRow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *BiomarkerLiteratureOrgsRow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *BiomarkerLiteratureOrgsRow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *BiomarkerLiteratureOrgsRow) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *BiomarkerLiteratureOrgsRow) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *BiomarkerLiteratureOrgsRow) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetOrgType

`func (o *BiomarkerLiteratureOrgsRow) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *BiomarkerLiteratureOrgsRow) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *BiomarkerLiteratureOrgsRow) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *BiomarkerLiteratureOrgsRow) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### SetOrgTypeNil

`func (o *BiomarkerLiteratureOrgsRow) SetOrgTypeNil(b bool)`

 SetOrgTypeNil sets the value for OrgType to be an explicit nil

### UnsetOrgType
`func (o *BiomarkerLiteratureOrgsRow) UnsetOrgType()`

UnsetOrgType ensures that no value is present for OrgType, not even an explicit nil
### GetOrgCountry

`func (o *BiomarkerLiteratureOrgsRow) GetOrgCountry() string`

GetOrgCountry returns the OrgCountry field if non-nil, zero value otherwise.

### GetOrgCountryOk

`func (o *BiomarkerLiteratureOrgsRow) GetOrgCountryOk() (*string, bool)`

GetOrgCountryOk returns a tuple with the OrgCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCountry

`func (o *BiomarkerLiteratureOrgsRow) SetOrgCountry(v string)`

SetOrgCountry sets OrgCountry field to given value.

### HasOrgCountry

`func (o *BiomarkerLiteratureOrgsRow) HasOrgCountry() bool`

HasOrgCountry returns a boolean if a field has been set.

### SetOrgCountryNil

`func (o *BiomarkerLiteratureOrgsRow) SetOrgCountryNil(b bool)`

 SetOrgCountryNil sets the value for OrgCountry to be an explicit nil

### UnsetOrgCountry
`func (o *BiomarkerLiteratureOrgsRow) UnsetOrgCountry()`

UnsetOrgCountry ensures that no value is present for OrgCountry, not even an explicit nil
### GetNPapers

`func (o *BiomarkerLiteratureOrgsRow) GetNPapers() int32`

GetNPapers returns the NPapers field if non-nil, zero value otherwise.

### GetNPapersOk

`func (o *BiomarkerLiteratureOrgsRow) GetNPapersOk() (*int32, bool)`

GetNPapersOk returns a tuple with the NPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPapers

`func (o *BiomarkerLiteratureOrgsRow) SetNPapers(v int32)`

SetNPapers sets NPapers field to given value.

### HasNPapers

`func (o *BiomarkerLiteratureOrgsRow) HasNPapers() bool`

HasNPapers returns a boolean if a field has been set.

### SetNPapersNil

`func (o *BiomarkerLiteratureOrgsRow) SetNPapersNil(b bool)`

 SetNPapersNil sets the value for NPapers to be an explicit nil

### UnsetNPapers
`func (o *BiomarkerLiteratureOrgsRow) UnsetNPapers()`

UnsetNPapers ensures that no value is present for NPapers, not even an explicit nil
### GetNTrialLinkedPapers

`func (o *BiomarkerLiteratureOrgsRow) GetNTrialLinkedPapers() int32`

GetNTrialLinkedPapers returns the NTrialLinkedPapers field if non-nil, zero value otherwise.

### GetNTrialLinkedPapersOk

`func (o *BiomarkerLiteratureOrgsRow) GetNTrialLinkedPapersOk() (*int32, bool)`

GetNTrialLinkedPapersOk returns a tuple with the NTrialLinkedPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialLinkedPapers

`func (o *BiomarkerLiteratureOrgsRow) SetNTrialLinkedPapers(v int32)`

SetNTrialLinkedPapers sets NTrialLinkedPapers field to given value.

### HasNTrialLinkedPapers

`func (o *BiomarkerLiteratureOrgsRow) HasNTrialLinkedPapers() bool`

HasNTrialLinkedPapers returns a boolean if a field has been set.

### SetNTrialLinkedPapersNil

`func (o *BiomarkerLiteratureOrgsRow) SetNTrialLinkedPapersNil(b bool)`

 SetNTrialLinkedPapersNil sets the value for NTrialLinkedPapers to be an explicit nil

### UnsetNTrialLinkedPapers
`func (o *BiomarkerLiteratureOrgsRow) UnsetNTrialLinkedPapers()`

UnsetNTrialLinkedPapers ensures that no value is present for NTrialLinkedPapers, not even an explicit nil
### GetLatestPubYear

`func (o *BiomarkerLiteratureOrgsRow) GetLatestPubYear() int32`

GetLatestPubYear returns the LatestPubYear field if non-nil, zero value otherwise.

### GetLatestPubYearOk

`func (o *BiomarkerLiteratureOrgsRow) GetLatestPubYearOk() (*int32, bool)`

GetLatestPubYearOk returns a tuple with the LatestPubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPubYear

`func (o *BiomarkerLiteratureOrgsRow) SetLatestPubYear(v int32)`

SetLatestPubYear sets LatestPubYear field to given value.

### HasLatestPubYear

`func (o *BiomarkerLiteratureOrgsRow) HasLatestPubYear() bool`

HasLatestPubYear returns a boolean if a field has been set.

### SetLatestPubYearNil

`func (o *BiomarkerLiteratureOrgsRow) SetLatestPubYearNil(b bool)`

 SetLatestPubYearNil sets the value for LatestPubYear to be an explicit nil

### UnsetLatestPubYear
`func (o *BiomarkerLiteratureOrgsRow) UnsetLatestPubYear()`

UnsetLatestPubYear ensures that no value is present for LatestPubYear, not even an explicit nil
### GetOrgArm

`func (o *BiomarkerLiteratureOrgsRow) GetOrgArm() string`

GetOrgArm returns the OrgArm field if non-nil, zero value otherwise.

### GetOrgArmOk

`func (o *BiomarkerLiteratureOrgsRow) GetOrgArmOk() (*string, bool)`

GetOrgArmOk returns a tuple with the OrgArm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgArm

`func (o *BiomarkerLiteratureOrgsRow) SetOrgArm(v string)`

SetOrgArm sets OrgArm field to given value.

### HasOrgArm

`func (o *BiomarkerLiteratureOrgsRow) HasOrgArm() bool`

HasOrgArm returns a boolean if a field has been set.

### SetOrgArmNil

`func (o *BiomarkerLiteratureOrgsRow) SetOrgArmNil(b bool)`

 SetOrgArmNil sets the value for OrgArm to be an explicit nil

### UnsetOrgArm
`func (o *BiomarkerLiteratureOrgsRow) UnsetOrgArm()`

UnsetOrgArm ensures that no value is present for OrgArm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


