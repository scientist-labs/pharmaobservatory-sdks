# ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | **string** |  | 
**Name** | **string** |  | 
**Found** | **bool** |  | 
**DalysM** | **float32** |  | 
**Trials** | **int32** |  | 
**TrialsPerMDaly** | **string** |  | 
**Rank** | **int32** |  | 
**TotalRanked** | **int32** |  | 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Restriction** | [**ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction**](ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction.md) |  | 

## Methods

### NewApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response

`func NewApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response(conditionId string, name string, found bool, dalysM float32, trials int32, trialsPerMDaly string, rank int32, totalRanked int32, buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, restriction ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction, ) *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response`

NewApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response instantiates a new ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseWithDefaults

`func NewApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseWithDefaults() *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response`

NewApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseWithDefaults instantiates a new ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.


### GetName

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) SetName(v string)`

SetName sets Name field to given value.


### GetFound

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) SetFound(v bool)`

SetFound sets Found field to given value.


### GetDalysM

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetDalysM() float32`

GetDalysM returns the DalysM field if non-nil, zero value otherwise.

### GetDalysMOk

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetDalysMOk() (*float32, bool)`

GetDalysMOk returns a tuple with the DalysM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDalysM

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) SetDalysM(v float32)`

SetDalysM sets DalysM field to given value.


### GetTrials

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetTrials() int32`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetTrialsOk() (*int32, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) SetTrials(v int32)`

SetTrials sets Trials field to given value.


### GetTrialsPerMDaly

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetTrialsPerMDaly() string`

GetTrialsPerMDaly returns the TrialsPerMDaly field if non-nil, zero value otherwise.

### GetTrialsPerMDalyOk

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetTrialsPerMDalyOk() (*string, bool)`

GetTrialsPerMDalyOk returns a tuple with the TrialsPerMDaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialsPerMDaly

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) SetTrialsPerMDaly(v string)`

SetTrialsPerMDaly sets TrialsPerMDaly field to given value.


### GetRank

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetTotalRanked

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetTotalRanked() int32`

GetTotalRanked returns the TotalRanked field if non-nil, zero value otherwise.

### GetTotalRankedOk

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetTotalRankedOk() (*int32, bool)`

GetTotalRankedOk returns a tuple with the TotalRanked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRanked

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) SetTotalRanked(v int32)`

SetTotalRanked sets TotalRanked field to given value.


### GetBuildState

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetRestriction

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetRestriction() ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) GetRestrictionOk() (*ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response) SetRestriction(v ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction)`

SetRestriction sets Restriction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


