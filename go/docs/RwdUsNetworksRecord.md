# RwdUsNetworksRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** |  | 
**Id** | **string** |  | 
**Record** | [**RwdUsNetworksRow**](RwdUsNetworksRow.md) |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Datasets** | [**[]DatasetListing**](DatasetListing.md) |  | 
**Attribution** | Pointer to [**[]DatasetAttribution**](DatasetAttribution.md) |  | [optional] 
**SourceCredit** | Pointer to **string** |  | [optional] 
**Freshness** | Pointer to [**AnimalWelfareAssurancesResponseFreshness**](AnimalWelfareAssurancesResponseFreshness.md) |  | [optional] 
**Build** | Pointer to [**AnimalWelfareAssurancesResponseBuild**](AnimalWelfareAssurancesResponseBuild.md) |  | [optional] 
**BuildState** | Pointer to [**DatasetBuildState**](DatasetBuildState.md) |  | [optional] 
**Restriction** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRwdUsNetworksRecord

`func NewRwdUsNetworksRecord(objectType string, id string, record RwdUsNetworksRow, datasets []DatasetListing, ) *RwdUsNetworksRecord`

NewRwdUsNetworksRecord instantiates a new RwdUsNetworksRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdUsNetworksRecordWithDefaults

`func NewRwdUsNetworksRecordWithDefaults() *RwdUsNetworksRecord`

NewRwdUsNetworksRecordWithDefaults instantiates a new RwdUsNetworksRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *RwdUsNetworksRecord) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RwdUsNetworksRecord) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RwdUsNetworksRecord) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *RwdUsNetworksRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RwdUsNetworksRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RwdUsNetworksRecord) SetId(v string)`

SetId sets Id field to given value.


### GetRecord

`func (o *RwdUsNetworksRecord) GetRecord() RwdUsNetworksRow`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *RwdUsNetworksRecord) GetRecordOk() (*RwdUsNetworksRow, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *RwdUsNetworksRecord) SetRecord(v RwdUsNetworksRow)`

SetRecord sets Record field to given value.


### GetAttributes

`func (o *RwdUsNetworksRecord) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RwdUsNetworksRecord) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RwdUsNetworksRecord) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RwdUsNetworksRecord) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDatasets

`func (o *RwdUsNetworksRecord) GetDatasets() []DatasetListing`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *RwdUsNetworksRecord) GetDatasetsOk() (*[]DatasetListing, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *RwdUsNetworksRecord) SetDatasets(v []DatasetListing)`

SetDatasets sets Datasets field to given value.


### GetAttribution

`func (o *RwdUsNetworksRecord) GetAttribution() []DatasetAttribution`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdUsNetworksRecord) GetAttributionOk() (*[]DatasetAttribution, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdUsNetworksRecord) SetAttribution(v []DatasetAttribution)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdUsNetworksRecord) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### GetSourceCredit

`func (o *RwdUsNetworksRecord) GetSourceCredit() string`

GetSourceCredit returns the SourceCredit field if non-nil, zero value otherwise.

### GetSourceCreditOk

`func (o *RwdUsNetworksRecord) GetSourceCreditOk() (*string, bool)`

GetSourceCreditOk returns a tuple with the SourceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCredit

`func (o *RwdUsNetworksRecord) SetSourceCredit(v string)`

SetSourceCredit sets SourceCredit field to given value.

### HasSourceCredit

`func (o *RwdUsNetworksRecord) HasSourceCredit() bool`

HasSourceCredit returns a boolean if a field has been set.

### GetFreshness

`func (o *RwdUsNetworksRecord) GetFreshness() AnimalWelfareAssurancesResponseFreshness`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *RwdUsNetworksRecord) GetFreshnessOk() (*AnimalWelfareAssurancesResponseFreshness, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *RwdUsNetworksRecord) SetFreshness(v AnimalWelfareAssurancesResponseFreshness)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *RwdUsNetworksRecord) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetBuild

`func (o *RwdUsNetworksRecord) GetBuild() AnimalWelfareAssurancesResponseBuild`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *RwdUsNetworksRecord) GetBuildOk() (*AnimalWelfareAssurancesResponseBuild, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *RwdUsNetworksRecord) SetBuild(v AnimalWelfareAssurancesResponseBuild)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *RwdUsNetworksRecord) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetBuildState

`func (o *RwdUsNetworksRecord) GetBuildState() DatasetBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *RwdUsNetworksRecord) GetBuildStateOk() (*DatasetBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *RwdUsNetworksRecord) SetBuildState(v DatasetBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *RwdUsNetworksRecord) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### GetRestriction

`func (o *RwdUsNetworksRecord) GetRestriction() map[string]interface{}`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *RwdUsNetworksRecord) GetRestrictionOk() (*map[string]interface{}, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *RwdUsNetworksRecord) SetRestriction(v map[string]interface{})`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *RwdUsNetworksRecord) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### SetRestrictionNil

`func (o *RwdUsNetworksRecord) SetRestrictionNil(b bool)`

 SetRestrictionNil sets the value for Restriction to be an explicit nil

### UnsetRestriction
`func (o *RwdUsNetworksRecord) UnsetRestriction()`

UnsetRestriction ensures that no value is present for Restriction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


