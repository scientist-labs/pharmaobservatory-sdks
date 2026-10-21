# TargetAssociationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] 
**Datasource** | Pointer to **NullableString** | The Open Targets datasource name (for example genetic_association or literature). | [optional] 
**NDiseaseAssociations** | Pointer to **NullableInt64** | Distinct disease associations the datasource holds for the target. | [optional] 
**MinAssociationScore** | Pointer to **NullableFloat64** | The lowest association score in the datasource for the target. | [optional] 
**MaxAssociationScore** | Pointer to **NullableFloat64** | The highest association score in the datasource for the target. | [optional] 

## Methods

### NewTargetAssociationsRow

`func NewTargetAssociationsRow() *TargetAssociationsRow`

NewTargetAssociationsRow instantiates a new TargetAssociationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetAssociationsRowWithDefaults

`func NewTargetAssociationsRowWithDefaults() *TargetAssociationsRow`

NewTargetAssociationsRowWithDefaults instantiates a new TargetAssociationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetAssociationsRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetAssociationsRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetAssociationsRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetAssociationsRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetAssociationsRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetAssociationsRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetDatasource

`func (o *TargetAssociationsRow) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *TargetAssociationsRow) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *TargetAssociationsRow) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *TargetAssociationsRow) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### SetDatasourceNil

`func (o *TargetAssociationsRow) SetDatasourceNil(b bool)`

 SetDatasourceNil sets the value for Datasource to be an explicit nil

### UnsetDatasource
`func (o *TargetAssociationsRow) UnsetDatasource()`

UnsetDatasource ensures that no value is present for Datasource, not even an explicit nil
### GetNDiseaseAssociations

`func (o *TargetAssociationsRow) GetNDiseaseAssociations() int64`

GetNDiseaseAssociations returns the NDiseaseAssociations field if non-nil, zero value otherwise.

### GetNDiseaseAssociationsOk

`func (o *TargetAssociationsRow) GetNDiseaseAssociationsOk() (*int64, bool)`

GetNDiseaseAssociationsOk returns a tuple with the NDiseaseAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDiseaseAssociations

`func (o *TargetAssociationsRow) SetNDiseaseAssociations(v int64)`

SetNDiseaseAssociations sets NDiseaseAssociations field to given value.

### HasNDiseaseAssociations

`func (o *TargetAssociationsRow) HasNDiseaseAssociations() bool`

HasNDiseaseAssociations returns a boolean if a field has been set.

### SetNDiseaseAssociationsNil

`func (o *TargetAssociationsRow) SetNDiseaseAssociationsNil(b bool)`

 SetNDiseaseAssociationsNil sets the value for NDiseaseAssociations to be an explicit nil

### UnsetNDiseaseAssociations
`func (o *TargetAssociationsRow) UnsetNDiseaseAssociations()`

UnsetNDiseaseAssociations ensures that no value is present for NDiseaseAssociations, not even an explicit nil
### GetMinAssociationScore

`func (o *TargetAssociationsRow) GetMinAssociationScore() float64`

GetMinAssociationScore returns the MinAssociationScore field if non-nil, zero value otherwise.

### GetMinAssociationScoreOk

`func (o *TargetAssociationsRow) GetMinAssociationScoreOk() (*float64, bool)`

GetMinAssociationScoreOk returns a tuple with the MinAssociationScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAssociationScore

`func (o *TargetAssociationsRow) SetMinAssociationScore(v float64)`

SetMinAssociationScore sets MinAssociationScore field to given value.

### HasMinAssociationScore

`func (o *TargetAssociationsRow) HasMinAssociationScore() bool`

HasMinAssociationScore returns a boolean if a field has been set.

### SetMinAssociationScoreNil

`func (o *TargetAssociationsRow) SetMinAssociationScoreNil(b bool)`

 SetMinAssociationScoreNil sets the value for MinAssociationScore to be an explicit nil

### UnsetMinAssociationScore
`func (o *TargetAssociationsRow) UnsetMinAssociationScore()`

UnsetMinAssociationScore ensures that no value is present for MinAssociationScore, not even an explicit nil
### GetMaxAssociationScore

`func (o *TargetAssociationsRow) GetMaxAssociationScore() float64`

GetMaxAssociationScore returns the MaxAssociationScore field if non-nil, zero value otherwise.

### GetMaxAssociationScoreOk

`func (o *TargetAssociationsRow) GetMaxAssociationScoreOk() (*float64, bool)`

GetMaxAssociationScoreOk returns a tuple with the MaxAssociationScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAssociationScore

`func (o *TargetAssociationsRow) SetMaxAssociationScore(v float64)`

SetMaxAssociationScore sets MaxAssociationScore field to given value.

### HasMaxAssociationScore

`func (o *TargetAssociationsRow) HasMaxAssociationScore() bool`

HasMaxAssociationScore returns a boolean if a field has been set.

### SetMaxAssociationScoreNil

`func (o *TargetAssociationsRow) SetMaxAssociationScoreNil(b bool)`

 SetMaxAssociationScoreNil sets the value for MaxAssociationScore to be an explicit nil

### UnsetMaxAssociationScore
`func (o *TargetAssociationsRow) UnsetMaxAssociationScore()`

UnsetMaxAssociationScore ensures that no value is present for MaxAssociationScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


