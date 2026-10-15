# ApiV1TrialsHistoryAndPredictionGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NctId** | Pointer to **string** |  | [optional] 
**Found** | Pointer to **bool** |  | [optional] 
**RecordHistory** | Pointer to [**NullableApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory.md) |  | [optional] 
**Predictions** | Pointer to [**[]ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner.md) |  | [optional] 
**PredictionsStatus** | Pointer to [**ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsStatus**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsStatus.md) |  | [optional] 
**BuildState** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  | 
**Sections** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseSections**](ApiV1TrialsHistoryAndPredictionGet200ResponseSections.md) |  | 
**Error** | Pointer to **string** |  | [optional] 
**UnbuiltRelations** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1TrialsHistoryAndPredictionGet200Response

`func NewApiV1TrialsHistoryAndPredictionGet200Response(buildState ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, sections ApiV1TrialsHistoryAndPredictionGet200ResponseSections, ) *ApiV1TrialsHistoryAndPredictionGet200Response`

NewApiV1TrialsHistoryAndPredictionGet200Response instantiates a new ApiV1TrialsHistoryAndPredictionGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsHistoryAndPredictionGet200ResponseWithDefaults

`func NewApiV1TrialsHistoryAndPredictionGet200ResponseWithDefaults() *ApiV1TrialsHistoryAndPredictionGet200Response`

NewApiV1TrialsHistoryAndPredictionGet200ResponseWithDefaults instantiates a new ApiV1TrialsHistoryAndPredictionGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNctId

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetNctId() string`

GetNctId returns the NctId field if non-nil, zero value otherwise.

### GetNctIdOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetNctIdOk() (*string, bool)`

GetNctIdOk returns a tuple with the NctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNctId

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetNctId(v string)`

SetNctId sets NctId field to given value.

### HasNctId

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) HasNctId() bool`

HasNctId returns a boolean if a field has been set.

### GetFound

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetFound(v bool)`

SetFound sets Found field to given value.

### HasFound

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) HasFound() bool`

HasFound returns a boolean if a field has been set.

### GetRecordHistory

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetRecordHistory() ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory`

GetRecordHistory returns the RecordHistory field if non-nil, zero value otherwise.

### GetRecordHistoryOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetRecordHistoryOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory, bool)`

GetRecordHistoryOk returns a tuple with the RecordHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordHistory

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetRecordHistory(v ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory)`

SetRecordHistory sets RecordHistory field to given value.

### HasRecordHistory

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) HasRecordHistory() bool`

HasRecordHistory returns a boolean if a field has been set.

### SetRecordHistoryNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetRecordHistoryNil(b bool)`

 SetRecordHistoryNil sets the value for RecordHistory to be an explicit nil

### UnsetRecordHistory
`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) UnsetRecordHistory()`

UnsetRecordHistory ensures that no value is present for RecordHistory, not even an explicit nil
### GetPredictions

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetPredictions() []ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner`

GetPredictions returns the Predictions field if non-nil, zero value otherwise.

### GetPredictionsOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetPredictionsOk() (*[]ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner, bool)`

GetPredictionsOk returns a tuple with the Predictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictions

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetPredictions(v []ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner)`

SetPredictions sets Predictions field to given value.

### HasPredictions

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) HasPredictions() bool`

HasPredictions returns a boolean if a field has been set.

### GetPredictionsStatus

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetPredictionsStatus() ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsStatus`

GetPredictionsStatus returns the PredictionsStatus field if non-nil, zero value otherwise.

### GetPredictionsStatusOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetPredictionsStatusOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsStatus, bool)`

GetPredictionsStatusOk returns a tuple with the PredictionsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionsStatus

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetPredictionsStatus(v ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsStatus)`

SetPredictionsStatus sets PredictionsStatus field to given value.

### HasPredictionsStatus

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) HasPredictionsStatus() bool`

HasPredictionsStatus returns a boolean if a field has been set.

### GetBuildState

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetBuildState() ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetBuildStateOk() (*ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetBuildState(v ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetSections

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetSections() ApiV1TrialsHistoryAndPredictionGet200ResponseSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetSectionsOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponseSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetSections(v ApiV1TrialsHistoryAndPredictionGet200ResponseSections)`

SetSections sets Sections field to given value.


### GetError

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetUnbuiltRelations

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetUnbuiltRelations() []string`

GetUnbuiltRelations returns the UnbuiltRelations field if non-nil, zero value otherwise.

### GetUnbuiltRelationsOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetUnbuiltRelationsOk() (*[]string, bool)`

GetUnbuiltRelationsOk returns a tuple with the UnbuiltRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbuiltRelations

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetUnbuiltRelations(v []string)`

SetUnbuiltRelations sets UnbuiltRelations field to given value.

### HasUnbuiltRelations

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) HasUnbuiltRelations() bool`

HasUnbuiltRelations returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1TrialsHistoryAndPredictionGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


