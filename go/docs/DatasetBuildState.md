# DatasetBuildState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Built** | Pointer to **bool** |  | [optional] 
**Generation** | Pointer to **NullableString** |  | [optional] 
**LastLoaded** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDatasetBuildState

`func NewDatasetBuildState() *DatasetBuildState`

NewDatasetBuildState instantiates a new DatasetBuildState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetBuildStateWithDefaults

`func NewDatasetBuildStateWithDefaults() *DatasetBuildState`

NewDatasetBuildStateWithDefaults instantiates a new DatasetBuildState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilt

`func (o *DatasetBuildState) GetBuilt() bool`

GetBuilt returns the Built field if non-nil, zero value otherwise.

### GetBuiltOk

`func (o *DatasetBuildState) GetBuiltOk() (*bool, bool)`

GetBuiltOk returns a tuple with the Built field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilt

`func (o *DatasetBuildState) SetBuilt(v bool)`

SetBuilt sets Built field to given value.

### HasBuilt

`func (o *DatasetBuildState) HasBuilt() bool`

HasBuilt returns a boolean if a field has been set.

### GetGeneration

`func (o *DatasetBuildState) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *DatasetBuildState) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *DatasetBuildState) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *DatasetBuildState) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### SetGenerationNil

`func (o *DatasetBuildState) SetGenerationNil(b bool)`

 SetGenerationNil sets the value for Generation to be an explicit nil

### UnsetGeneration
`func (o *DatasetBuildState) UnsetGeneration()`

UnsetGeneration ensures that no value is present for Generation, not even an explicit nil
### GetLastLoaded

`func (o *DatasetBuildState) GetLastLoaded() string`

GetLastLoaded returns the LastLoaded field if non-nil, zero value otherwise.

### GetLastLoadedOk

`func (o *DatasetBuildState) GetLastLoadedOk() (*string, bool)`

GetLastLoadedOk returns a tuple with the LastLoaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoaded

`func (o *DatasetBuildState) SetLastLoaded(v string)`

SetLastLoaded sets LastLoaded field to given value.

### HasLastLoaded

`func (o *DatasetBuildState) HasLastLoaded() bool`

HasLastLoaded returns a boolean if a field has been set.

### SetLastLoadedNil

`func (o *DatasetBuildState) SetLastLoadedNil(b bool)`

 SetLastLoadedNil sets the value for LastLoaded to be an explicit nil

### UnsetLastLoaded
`func (o *DatasetBuildState) UnsetLastLoaded()`

UnsetLastLoaded ensures that no value is present for LastLoaded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


