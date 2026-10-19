# TrialOutcomeMeasuresRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**OutcomeId** | Pointer to **NullableInt32** | The outcome measure&#39;s registry id. | [optional] 
**OutcomeType** | Pointer to **NullableString** | PRIMARY, SECONDARY or OTHER. | [optional] 
**Title** | Pointer to **NullableString** | The outcome measure&#39;s title. | [optional] 
**TimeFrame** | Pointer to **NullableString** | The time frame the measure covers. | [optional] 
**Population** | Pointer to **NullableString** | The population the measure is reported for. | [optional] 

## Methods

### NewTrialOutcomeMeasuresRow

`func NewTrialOutcomeMeasuresRow() *TrialOutcomeMeasuresRow`

NewTrialOutcomeMeasuresRow instantiates a new TrialOutcomeMeasuresRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialOutcomeMeasuresRowWithDefaults

`func NewTrialOutcomeMeasuresRowWithDefaults() *TrialOutcomeMeasuresRow`

NewTrialOutcomeMeasuresRowWithDefaults instantiates a new TrialOutcomeMeasuresRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialOutcomeMeasuresRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialOutcomeMeasuresRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialOutcomeMeasuresRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialOutcomeMeasuresRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialOutcomeMeasuresRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialOutcomeMeasuresRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetOutcomeId

`func (o *TrialOutcomeMeasuresRow) GetOutcomeId() int32`

GetOutcomeId returns the OutcomeId field if non-nil, zero value otherwise.

### GetOutcomeIdOk

`func (o *TrialOutcomeMeasuresRow) GetOutcomeIdOk() (*int32, bool)`

GetOutcomeIdOk returns a tuple with the OutcomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeId

`func (o *TrialOutcomeMeasuresRow) SetOutcomeId(v int32)`

SetOutcomeId sets OutcomeId field to given value.

### HasOutcomeId

`func (o *TrialOutcomeMeasuresRow) HasOutcomeId() bool`

HasOutcomeId returns a boolean if a field has been set.

### SetOutcomeIdNil

`func (o *TrialOutcomeMeasuresRow) SetOutcomeIdNil(b bool)`

 SetOutcomeIdNil sets the value for OutcomeId to be an explicit nil

### UnsetOutcomeId
`func (o *TrialOutcomeMeasuresRow) UnsetOutcomeId()`

UnsetOutcomeId ensures that no value is present for OutcomeId, not even an explicit nil
### GetOutcomeType

`func (o *TrialOutcomeMeasuresRow) GetOutcomeType() string`

GetOutcomeType returns the OutcomeType field if non-nil, zero value otherwise.

### GetOutcomeTypeOk

`func (o *TrialOutcomeMeasuresRow) GetOutcomeTypeOk() (*string, bool)`

GetOutcomeTypeOk returns a tuple with the OutcomeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeType

`func (o *TrialOutcomeMeasuresRow) SetOutcomeType(v string)`

SetOutcomeType sets OutcomeType field to given value.

### HasOutcomeType

`func (o *TrialOutcomeMeasuresRow) HasOutcomeType() bool`

HasOutcomeType returns a boolean if a field has been set.

### SetOutcomeTypeNil

`func (o *TrialOutcomeMeasuresRow) SetOutcomeTypeNil(b bool)`

 SetOutcomeTypeNil sets the value for OutcomeType to be an explicit nil

### UnsetOutcomeType
`func (o *TrialOutcomeMeasuresRow) UnsetOutcomeType()`

UnsetOutcomeType ensures that no value is present for OutcomeType, not even an explicit nil
### GetTitle

`func (o *TrialOutcomeMeasuresRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrialOutcomeMeasuresRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrialOutcomeMeasuresRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrialOutcomeMeasuresRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TrialOutcomeMeasuresRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TrialOutcomeMeasuresRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTimeFrame

`func (o *TrialOutcomeMeasuresRow) GetTimeFrame() string`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *TrialOutcomeMeasuresRow) GetTimeFrameOk() (*string, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *TrialOutcomeMeasuresRow) SetTimeFrame(v string)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *TrialOutcomeMeasuresRow) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### SetTimeFrameNil

`func (o *TrialOutcomeMeasuresRow) SetTimeFrameNil(b bool)`

 SetTimeFrameNil sets the value for TimeFrame to be an explicit nil

### UnsetTimeFrame
`func (o *TrialOutcomeMeasuresRow) UnsetTimeFrame()`

UnsetTimeFrame ensures that no value is present for TimeFrame, not even an explicit nil
### GetPopulation

`func (o *TrialOutcomeMeasuresRow) GetPopulation() string`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *TrialOutcomeMeasuresRow) GetPopulationOk() (*string, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *TrialOutcomeMeasuresRow) SetPopulation(v string)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *TrialOutcomeMeasuresRow) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.

### SetPopulationNil

`func (o *TrialOutcomeMeasuresRow) SetPopulationNil(b bool)`

 SetPopulationNil sets the value for Population to be an explicit nil

### UnsetPopulation
`func (o *TrialOutcomeMeasuresRow) UnsetPopulation()`

UnsetPopulation ensures that no value is present for Population, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


