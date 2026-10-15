# ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**State** | **string** |  | 
**RealizedState** | **string** |  | 
**Cutoff** | **string** |  | 
**MinN** | **int32** |  | 
**CapMonths** | **int32** |  | 
**ModelVersion** | **string** |  | 
**Method** | **string** |  | 
**NTrials** | **int32** |  | 
**NEvaluated** | **int32** |  | 
**NNotEvaluated** | **int32** |  | 
**Statuses** | [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupStatuses**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupStatuses.md) |  | 
**NSuperseded** | **int32** |  | 
**NOpen** | **int32** |  | 
**NAtCap** | **int32** |  | 
**NOverdue** | **int32** |  | 
**NErrorMeasured** | **int32** |  | 
**Predicted** | [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPredicted**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPredicted.md) |  | 
**NRealized** | **int32** |  | 
**NRealizedOverCap** | **int32** |  | 
**Realized** | [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupRealized**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupRealized.md) |  | 
**Phases** | [**[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPhasesInner**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPhasesInner.md) |  | 
**Scope** | **string** |  | 
**Title** | **string** |  | 
**Body** | **string** |  | 
**Figures** | [**[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupFiguresInner**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupFiguresInner.md) |  | 
**StatusCounts** | [**[]ApiV1SponsorLandscapeRecruitmentRollupStatusCountsInner**](ApiV1SponsorLandscapeRecruitmentRollupStatusCountsInner.md) |  | 
**Source** | **string** |  | 
**PredictedAtHorizonCap** | Pointer to [**ApiV1SponsorLandscapeRecruitmentRollupPredictedAtHorizonCap**](ApiV1SponsorLandscapeRecruitmentRollupPredictedAtHorizonCap.md) |  | [optional] 

## Methods

### NewApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup

`func NewApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup(kind string, state string, realizedState string, cutoff string, minN int32, capMonths int32, modelVersion string, method string, nTrials int32, nEvaluated int32, nNotEvaluated int32, statuses ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupStatuses, nSuperseded int32, nOpen int32, nAtCap int32, nOverdue int32, nErrorMeasured int32, predicted ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPredicted, nRealized int32, nRealizedOverCap int32, realized ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupRealized, phases []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPhasesInner, scope string, title string, body string, figures []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupFiguresInner, statusCounts []ApiV1SponsorLandscapeRecruitmentRollupStatusCountsInner, source string, ) *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup`

NewApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup instantiates a new ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupWithDefaults

`func NewApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupWithDefaults() *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup`

NewApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupWithDefaults instantiates a new ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetKind(v string)`

SetKind sets Kind field to given value.


### GetState

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetState(v string)`

SetState sets State field to given value.


### GetRealizedState

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetRealizedState() string`

GetRealizedState returns the RealizedState field if non-nil, zero value otherwise.

### GetRealizedStateOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetRealizedStateOk() (*string, bool)`

GetRealizedStateOk returns a tuple with the RealizedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedState

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetRealizedState(v string)`

SetRealizedState sets RealizedState field to given value.


### GetCutoff

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetCutoff() string`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetCutoffOk() (*string, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoff

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetCutoff(v string)`

SetCutoff sets Cutoff field to given value.


### GetMinN

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetMinN() int32`

GetMinN returns the MinN field if non-nil, zero value otherwise.

### GetMinNOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetMinNOk() (*int32, bool)`

GetMinNOk returns a tuple with the MinN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinN

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetMinN(v int32)`

SetMinN sets MinN field to given value.


### GetCapMonths

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetCapMonths() int32`

GetCapMonths returns the CapMonths field if non-nil, zero value otherwise.

### GetCapMonthsOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetCapMonthsOk() (*int32, bool)`

GetCapMonthsOk returns a tuple with the CapMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapMonths

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetCapMonths(v int32)`

SetCapMonths sets CapMonths field to given value.


### GetModelVersion

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.


### GetMethod

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetNTrials

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.


### GetNEvaluated

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNEvaluated() int32`

GetNEvaluated returns the NEvaluated field if non-nil, zero value otherwise.

### GetNEvaluatedOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNEvaluatedOk() (*int32, bool)`

GetNEvaluatedOk returns a tuple with the NEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEvaluated

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetNEvaluated(v int32)`

SetNEvaluated sets NEvaluated field to given value.


### GetNNotEvaluated

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNNotEvaluated() int32`

GetNNotEvaluated returns the NNotEvaluated field if non-nil, zero value otherwise.

### GetNNotEvaluatedOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNNotEvaluatedOk() (*int32, bool)`

GetNNotEvaluatedOk returns a tuple with the NNotEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNNotEvaluated

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetNNotEvaluated(v int32)`

SetNNotEvaluated sets NNotEvaluated field to given value.


### GetStatuses

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetStatuses() ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupStatuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetStatusesOk() (*ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupStatuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetStatuses(v ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupStatuses)`

SetStatuses sets Statuses field to given value.


### GetNSuperseded

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNSuperseded() int32`

GetNSuperseded returns the NSuperseded field if non-nil, zero value otherwise.

### GetNSupersededOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNSupersededOk() (*int32, bool)`

GetNSupersededOk returns a tuple with the NSuperseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSuperseded

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetNSuperseded(v int32)`

SetNSuperseded sets NSuperseded field to given value.


### GetNOpen

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNOpen() int32`

GetNOpen returns the NOpen field if non-nil, zero value otherwise.

### GetNOpenOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNOpenOk() (*int32, bool)`

GetNOpenOk returns a tuple with the NOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOpen

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetNOpen(v int32)`

SetNOpen sets NOpen field to given value.


### GetNAtCap

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNAtCap() int32`

GetNAtCap returns the NAtCap field if non-nil, zero value otherwise.

### GetNAtCapOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNAtCapOk() (*int32, bool)`

GetNAtCapOk returns a tuple with the NAtCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAtCap

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetNAtCap(v int32)`

SetNAtCap sets NAtCap field to given value.


### GetNOverdue

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNOverdue() int32`

GetNOverdue returns the NOverdue field if non-nil, zero value otherwise.

### GetNOverdueOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNOverdueOk() (*int32, bool)`

GetNOverdueOk returns a tuple with the NOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOverdue

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetNOverdue(v int32)`

SetNOverdue sets NOverdue field to given value.


### GetNErrorMeasured

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNErrorMeasured() int32`

GetNErrorMeasured returns the NErrorMeasured field if non-nil, zero value otherwise.

### GetNErrorMeasuredOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNErrorMeasuredOk() (*int32, bool)`

GetNErrorMeasuredOk returns a tuple with the NErrorMeasured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNErrorMeasured

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetNErrorMeasured(v int32)`

SetNErrorMeasured sets NErrorMeasured field to given value.


### GetPredicted

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetPredicted() ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPredicted`

GetPredicted returns the Predicted field if non-nil, zero value otherwise.

### GetPredictedOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetPredictedOk() (*ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPredicted, bool)`

GetPredictedOk returns a tuple with the Predicted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicted

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetPredicted(v ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPredicted)`

SetPredicted sets Predicted field to given value.


### GetNRealized

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNRealized() int32`

GetNRealized returns the NRealized field if non-nil, zero value otherwise.

### GetNRealizedOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNRealizedOk() (*int32, bool)`

GetNRealizedOk returns a tuple with the NRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealized

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetNRealized(v int32)`

SetNRealized sets NRealized field to given value.


### GetNRealizedOverCap

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNRealizedOverCap() int32`

GetNRealizedOverCap returns the NRealizedOverCap field if non-nil, zero value otherwise.

### GetNRealizedOverCapOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetNRealizedOverCapOk() (*int32, bool)`

GetNRealizedOverCapOk returns a tuple with the NRealizedOverCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealizedOverCap

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetNRealizedOverCap(v int32)`

SetNRealizedOverCap sets NRealizedOverCap field to given value.


### GetRealized

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetRealized() ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupRealized`

GetRealized returns the Realized field if non-nil, zero value otherwise.

### GetRealizedOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetRealizedOk() (*ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupRealized, bool)`

GetRealizedOk returns a tuple with the Realized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealized

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetRealized(v ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupRealized)`

SetRealized sets Realized field to given value.


### GetPhases

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetPhases() []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPhasesInner`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetPhasesOk() (*[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPhasesInner, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetPhases(v []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupPhasesInner)`

SetPhases sets Phases field to given value.


### GetScope

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTitle

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetBody(v string)`

SetBody sets Body field to given value.


### GetFigures

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetFigures() []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupFiguresInner`

GetFigures returns the Figures field if non-nil, zero value otherwise.

### GetFiguresOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetFiguresOk() (*[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupFiguresInner, bool)`

GetFiguresOk returns a tuple with the Figures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigures

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetFigures(v []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupFiguresInner)`

SetFigures sets Figures field to given value.


### GetStatusCounts

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetStatusCounts() []ApiV1SponsorLandscapeRecruitmentRollupStatusCountsInner`

GetStatusCounts returns the StatusCounts field if non-nil, zero value otherwise.

### GetStatusCountsOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetStatusCountsOk() (*[]ApiV1SponsorLandscapeRecruitmentRollupStatusCountsInner, bool)`

GetStatusCountsOk returns a tuple with the StatusCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCounts

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetStatusCounts(v []ApiV1SponsorLandscapeRecruitmentRollupStatusCountsInner)`

SetStatusCounts sets StatusCounts field to given value.


### GetSource

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetSource(v string)`

SetSource sets Source field to given value.


### GetPredictedAtHorizonCap

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetPredictedAtHorizonCap() ApiV1SponsorLandscapeRecruitmentRollupPredictedAtHorizonCap`

GetPredictedAtHorizonCap returns the PredictedAtHorizonCap field if non-nil, zero value otherwise.

### GetPredictedAtHorizonCapOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) GetPredictedAtHorizonCapOk() (*ApiV1SponsorLandscapeRecruitmentRollupPredictedAtHorizonCap, bool)`

GetPredictedAtHorizonCapOk returns a tuple with the PredictedAtHorizonCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedAtHorizonCap

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) SetPredictedAtHorizonCap(v ApiV1SponsorLandscapeRecruitmentRollupPredictedAtHorizonCap)`

SetPredictedAtHorizonCap sets PredictedAtHorizonCap field to given value.

### HasPredictedAtHorizonCap

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup) HasPredictedAtHorizonCap() bool`

HasPredictedAtHorizonCap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


