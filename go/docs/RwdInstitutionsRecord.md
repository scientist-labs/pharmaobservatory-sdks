# RwdInstitutionsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** |  | 
**Id** | **string** |  | 
**Record** | [**RwdInstitutionsRow**](RwdInstitutionsRow.md) |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Datasets** | [**[]DatasetListing**](DatasetListing.md) |  | 
**Attribution** | Pointer to [**[]DatasetAttribution**](DatasetAttribution.md) |  | [optional] 
**SourceCredit** | Pointer to **string** |  | [optional] 
**Freshness** | Pointer to [**BiomarkerAssayEvidenceResponseFreshness**](BiomarkerAssayEvidenceResponseFreshness.md) |  | [optional] 
**Build** | Pointer to [**BiomarkerAssayEvidenceResponseBuild**](BiomarkerAssayEvidenceResponseBuild.md) |  | [optional] 
**BuildState** | Pointer to [**DatasetBuildState**](DatasetBuildState.md) |  | [optional] 
**Restriction** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRwdInstitutionsRecord

`func NewRwdInstitutionsRecord(objectType string, id string, record RwdInstitutionsRow, datasets []DatasetListing, ) *RwdInstitutionsRecord`

NewRwdInstitutionsRecord instantiates a new RwdInstitutionsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdInstitutionsRecordWithDefaults

`func NewRwdInstitutionsRecordWithDefaults() *RwdInstitutionsRecord`

NewRwdInstitutionsRecordWithDefaults instantiates a new RwdInstitutionsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *RwdInstitutionsRecord) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RwdInstitutionsRecord) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RwdInstitutionsRecord) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *RwdInstitutionsRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RwdInstitutionsRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RwdInstitutionsRecord) SetId(v string)`

SetId sets Id field to given value.


### GetRecord

`func (o *RwdInstitutionsRecord) GetRecord() RwdInstitutionsRow`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *RwdInstitutionsRecord) GetRecordOk() (*RwdInstitutionsRow, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *RwdInstitutionsRecord) SetRecord(v RwdInstitutionsRow)`

SetRecord sets Record field to given value.


### GetAttributes

`func (o *RwdInstitutionsRecord) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RwdInstitutionsRecord) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RwdInstitutionsRecord) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RwdInstitutionsRecord) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDatasets

`func (o *RwdInstitutionsRecord) GetDatasets() []DatasetListing`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *RwdInstitutionsRecord) GetDatasetsOk() (*[]DatasetListing, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *RwdInstitutionsRecord) SetDatasets(v []DatasetListing)`

SetDatasets sets Datasets field to given value.


### GetAttribution

`func (o *RwdInstitutionsRecord) GetAttribution() []DatasetAttribution`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdInstitutionsRecord) GetAttributionOk() (*[]DatasetAttribution, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdInstitutionsRecord) SetAttribution(v []DatasetAttribution)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdInstitutionsRecord) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### GetSourceCredit

`func (o *RwdInstitutionsRecord) GetSourceCredit() string`

GetSourceCredit returns the SourceCredit field if non-nil, zero value otherwise.

### GetSourceCreditOk

`func (o *RwdInstitutionsRecord) GetSourceCreditOk() (*string, bool)`

GetSourceCreditOk returns a tuple with the SourceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCredit

`func (o *RwdInstitutionsRecord) SetSourceCredit(v string)`

SetSourceCredit sets SourceCredit field to given value.

### HasSourceCredit

`func (o *RwdInstitutionsRecord) HasSourceCredit() bool`

HasSourceCredit returns a boolean if a field has been set.

### GetFreshness

`func (o *RwdInstitutionsRecord) GetFreshness() BiomarkerAssayEvidenceResponseFreshness`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *RwdInstitutionsRecord) GetFreshnessOk() (*BiomarkerAssayEvidenceResponseFreshness, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *RwdInstitutionsRecord) SetFreshness(v BiomarkerAssayEvidenceResponseFreshness)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *RwdInstitutionsRecord) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetBuild

`func (o *RwdInstitutionsRecord) GetBuild() BiomarkerAssayEvidenceResponseBuild`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *RwdInstitutionsRecord) GetBuildOk() (*BiomarkerAssayEvidenceResponseBuild, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *RwdInstitutionsRecord) SetBuild(v BiomarkerAssayEvidenceResponseBuild)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *RwdInstitutionsRecord) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetBuildState

`func (o *RwdInstitutionsRecord) GetBuildState() DatasetBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *RwdInstitutionsRecord) GetBuildStateOk() (*DatasetBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *RwdInstitutionsRecord) SetBuildState(v DatasetBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *RwdInstitutionsRecord) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### GetRestriction

`func (o *RwdInstitutionsRecord) GetRestriction() map[string]interface{}`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *RwdInstitutionsRecord) GetRestrictionOk() (*map[string]interface{}, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *RwdInstitutionsRecord) SetRestriction(v map[string]interface{})`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *RwdInstitutionsRecord) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### SetRestrictionNil

`func (o *RwdInstitutionsRecord) SetRestrictionNil(b bool)`

 SetRestrictionNil sets the value for Restriction to be an explicit nil

### UnsetRestriction
`func (o *RwdInstitutionsRecord) UnsetRestriction()`

UnsetRestriction ensures that no value is present for Restriction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


