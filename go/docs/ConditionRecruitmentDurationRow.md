# ConditionRecruitmentDurationRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s MeSH heading. | [optional] 
**NTrialsInSeries** | Pointer to **NullableInt32** | Trials in the condition series (the rows of condition_trials). | [optional] 
**NEvaluated** | Pointer to **NullableInt32** | Of those, trials with a row in the per-trial predictions at the cutoff. | [optional] 
**NNotEvaluated** | Pointer to **NullableInt32** | n_trials_in_series - n_evaluated: registered after the cutoff or a coverage gap; the data cannot tell which. | [optional] 
**NOutOfPopulation** | Pointer to **NullableInt32** | Evaluated trials with prediction_status out_of_population. | [optional] 
**NCompleted** | Pointer to **NullableInt32** | Evaluated trials with prediction_status completed. | [optional] 
**NStopped** | Pointer to **NullableInt32** | Evaluated trials with prediction_status stopped. | [optional] 
**NNoRegistrationRecord** | Pointer to **NullableInt32** | Evaluated trials with prediction_status no_registration_record. | [optional] 
**NRetrospectiveRegistration** | Pointer to **NullableInt32** | Evaluated trials with prediction_status retrospective_registration. | [optional] 
**NStartNotActual** | Pointer to **NullableInt32** | Evaluated trials with prediction_status start_not_actual. | [optional] 
**NFitFailed** | Pointer to **NullableInt32** | Evaluated trials with prediction_status fit_failed. | [optional] 
**NPredicted** | Pointer to **NullableInt32** | Evaluated trials with prediction_status predicted, superseded or not. | [optional] 
**NPredictedSuperseded** | Pointer to **NullableInt32** | Predicted trials whose CURRENT record shows the outcome (COMPLETED, TERMINATED, WITHDRAWN or an ACTUAL primary completion); kept out of every predicted figure. | [optional] 
**NPredictedOpen** | Pointer to **NullableInt32** | n_predicted - n_predicted_superseded; the denominator of every predicted figure. | [optional] 
**NPredictedAtCap** | Pointer to **NullableInt32** | Open predicted trials whose figure sits at the horizon cap (an at-cap row contributes the cap to the median). | [optional] 
**NPredictedOverdue** | Pointer to **NullableInt32** | Open predicted trials running past their predicted end at the cutoff. | [optional] 
**NPredictedErrorMeasured** | Pointer to **NullableInt32** | Open predicted trials in a phase bucket the backtest measured (error_phase_mae_months not null). | [optional] 
**PredictedMonthsMedian** | Pointer to **NullableFloat32** | Median of the carried predicted months over open predicted trials (capped at horizon_cap_months). NULL under min_n. | [optional] 
**PredictedMonthsP25** | Pointer to **NullableFloat32** | 25th percentile, same rows and floor. | [optional] 
**PredictedMonthsP75** | Pointer to **NullableFloat32** | 75th percentile, same rows and floor. | [optional] 
**BaselineMonthsMedian** | Pointer to **NullableFloat32** | Median of the carried baseline months over the same open predicted trials. NULL under min_n. | [optional] 
**NRealized** | Pointer to **NullableInt32** | Trials with a registered span (primary completion typed ACTUAL today, plausible start and completion), from the realized-duration input. | [optional] 
**RealizedMonthsMedian** | Pointer to **NullableFloat32** | Median registered start to primary completion in months over n_realized trials. NULL under min_n. | [optional] 
**RealizedMonthsP25** | Pointer to **NullableFloat32** | 25th percentile, same rows and floor. | [optional] 
**RealizedMonthsP75** | Pointer to **NullableFloat32** | 75th percentile, same rows and floor. | [optional] 
**NRealizedOverCap** | Pointer to **NullableInt32** | Realized spans above horizon_cap_months, so a reader sees how many would have hit the cap. | [optional] 
**MinN** | Pointer to **NullableInt32** | The floor applied: a median is withheld below this many trials, carried so a reader sees the floor used. | [optional] 
**RollupStatus** | Pointer to **NullableString** | rolled_up (open predicted &gt;&#x3D; min_n), below_floor (1 &lt;&#x3D; open &lt; min_n, counts only), none_predicted (evaluated, no open prediction), none_evaluated. | [optional] 
**RealizedStatus** | Pointer to **NullableString** | rolled_up (n_realized &gt;&#x3D; min_n), below_floor, or none. | [optional] 
**TrainingCutoff** | Pointer to **NullableString** | The date the model&#39;s training data ends. | [optional] 
**ModelName** | Pointer to **NullableString** | The recruitment-duration model. | [optional] 
**ModelVersion** | Pointer to **NullableString** | The model version. | [optional] 
**HorizonCapMonths** | Pointer to **NullableInt32** | The cap on predicted months (36 today), from the prediction rows. | [optional] 
**Method** | Pointer to **NullableString** | The run&#39;s method sentence from the run provenance, never re-authored. | [optional] 

## Methods

### NewConditionRecruitmentDurationRow

`func NewConditionRecruitmentDurationRow() *ConditionRecruitmentDurationRow`

NewConditionRecruitmentDurationRow instantiates a new ConditionRecruitmentDurationRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionRecruitmentDurationRowWithDefaults

`func NewConditionRecruitmentDurationRowWithDefaults() *ConditionRecruitmentDurationRow`

NewConditionRecruitmentDurationRowWithDefaults instantiates a new ConditionRecruitmentDurationRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionRecruitmentDurationRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionRecruitmentDurationRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionRecruitmentDurationRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionRecruitmentDurationRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionRecruitmentDurationRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionRecruitmentDurationRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *ConditionRecruitmentDurationRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ConditionRecruitmentDurationRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ConditionRecruitmentDurationRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *ConditionRecruitmentDurationRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *ConditionRecruitmentDurationRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *ConditionRecruitmentDurationRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetNTrialsInSeries

`func (o *ConditionRecruitmentDurationRow) GetNTrialsInSeries() int32`

GetNTrialsInSeries returns the NTrialsInSeries field if non-nil, zero value otherwise.

### GetNTrialsInSeriesOk

`func (o *ConditionRecruitmentDurationRow) GetNTrialsInSeriesOk() (*int32, bool)`

GetNTrialsInSeriesOk returns a tuple with the NTrialsInSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsInSeries

`func (o *ConditionRecruitmentDurationRow) SetNTrialsInSeries(v int32)`

SetNTrialsInSeries sets NTrialsInSeries field to given value.

### HasNTrialsInSeries

`func (o *ConditionRecruitmentDurationRow) HasNTrialsInSeries() bool`

HasNTrialsInSeries returns a boolean if a field has been set.

### SetNTrialsInSeriesNil

`func (o *ConditionRecruitmentDurationRow) SetNTrialsInSeriesNil(b bool)`

 SetNTrialsInSeriesNil sets the value for NTrialsInSeries to be an explicit nil

### UnsetNTrialsInSeries
`func (o *ConditionRecruitmentDurationRow) UnsetNTrialsInSeries()`

UnsetNTrialsInSeries ensures that no value is present for NTrialsInSeries, not even an explicit nil
### GetNEvaluated

`func (o *ConditionRecruitmentDurationRow) GetNEvaluated() int32`

GetNEvaluated returns the NEvaluated field if non-nil, zero value otherwise.

### GetNEvaluatedOk

`func (o *ConditionRecruitmentDurationRow) GetNEvaluatedOk() (*int32, bool)`

GetNEvaluatedOk returns a tuple with the NEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEvaluated

`func (o *ConditionRecruitmentDurationRow) SetNEvaluated(v int32)`

SetNEvaluated sets NEvaluated field to given value.

### HasNEvaluated

`func (o *ConditionRecruitmentDurationRow) HasNEvaluated() bool`

HasNEvaluated returns a boolean if a field has been set.

### SetNEvaluatedNil

`func (o *ConditionRecruitmentDurationRow) SetNEvaluatedNil(b bool)`

 SetNEvaluatedNil sets the value for NEvaluated to be an explicit nil

### UnsetNEvaluated
`func (o *ConditionRecruitmentDurationRow) UnsetNEvaluated()`

UnsetNEvaluated ensures that no value is present for NEvaluated, not even an explicit nil
### GetNNotEvaluated

`func (o *ConditionRecruitmentDurationRow) GetNNotEvaluated() int32`

GetNNotEvaluated returns the NNotEvaluated field if non-nil, zero value otherwise.

### GetNNotEvaluatedOk

`func (o *ConditionRecruitmentDurationRow) GetNNotEvaluatedOk() (*int32, bool)`

GetNNotEvaluatedOk returns a tuple with the NNotEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNNotEvaluated

`func (o *ConditionRecruitmentDurationRow) SetNNotEvaluated(v int32)`

SetNNotEvaluated sets NNotEvaluated field to given value.

### HasNNotEvaluated

`func (o *ConditionRecruitmentDurationRow) HasNNotEvaluated() bool`

HasNNotEvaluated returns a boolean if a field has been set.

### SetNNotEvaluatedNil

`func (o *ConditionRecruitmentDurationRow) SetNNotEvaluatedNil(b bool)`

 SetNNotEvaluatedNil sets the value for NNotEvaluated to be an explicit nil

### UnsetNNotEvaluated
`func (o *ConditionRecruitmentDurationRow) UnsetNNotEvaluated()`

UnsetNNotEvaluated ensures that no value is present for NNotEvaluated, not even an explicit nil
### GetNOutOfPopulation

`func (o *ConditionRecruitmentDurationRow) GetNOutOfPopulation() int32`

GetNOutOfPopulation returns the NOutOfPopulation field if non-nil, zero value otherwise.

### GetNOutOfPopulationOk

`func (o *ConditionRecruitmentDurationRow) GetNOutOfPopulationOk() (*int32, bool)`

GetNOutOfPopulationOk returns a tuple with the NOutOfPopulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOutOfPopulation

`func (o *ConditionRecruitmentDurationRow) SetNOutOfPopulation(v int32)`

SetNOutOfPopulation sets NOutOfPopulation field to given value.

### HasNOutOfPopulation

`func (o *ConditionRecruitmentDurationRow) HasNOutOfPopulation() bool`

HasNOutOfPopulation returns a boolean if a field has been set.

### SetNOutOfPopulationNil

`func (o *ConditionRecruitmentDurationRow) SetNOutOfPopulationNil(b bool)`

 SetNOutOfPopulationNil sets the value for NOutOfPopulation to be an explicit nil

### UnsetNOutOfPopulation
`func (o *ConditionRecruitmentDurationRow) UnsetNOutOfPopulation()`

UnsetNOutOfPopulation ensures that no value is present for NOutOfPopulation, not even an explicit nil
### GetNCompleted

`func (o *ConditionRecruitmentDurationRow) GetNCompleted() int32`

GetNCompleted returns the NCompleted field if non-nil, zero value otherwise.

### GetNCompletedOk

`func (o *ConditionRecruitmentDurationRow) GetNCompletedOk() (*int32, bool)`

GetNCompletedOk returns a tuple with the NCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompleted

`func (o *ConditionRecruitmentDurationRow) SetNCompleted(v int32)`

SetNCompleted sets NCompleted field to given value.

### HasNCompleted

`func (o *ConditionRecruitmentDurationRow) HasNCompleted() bool`

HasNCompleted returns a boolean if a field has been set.

### SetNCompletedNil

`func (o *ConditionRecruitmentDurationRow) SetNCompletedNil(b bool)`

 SetNCompletedNil sets the value for NCompleted to be an explicit nil

### UnsetNCompleted
`func (o *ConditionRecruitmentDurationRow) UnsetNCompleted()`

UnsetNCompleted ensures that no value is present for NCompleted, not even an explicit nil
### GetNStopped

`func (o *ConditionRecruitmentDurationRow) GetNStopped() int32`

GetNStopped returns the NStopped field if non-nil, zero value otherwise.

### GetNStoppedOk

`func (o *ConditionRecruitmentDurationRow) GetNStoppedOk() (*int32, bool)`

GetNStoppedOk returns a tuple with the NStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNStopped

`func (o *ConditionRecruitmentDurationRow) SetNStopped(v int32)`

SetNStopped sets NStopped field to given value.

### HasNStopped

`func (o *ConditionRecruitmentDurationRow) HasNStopped() bool`

HasNStopped returns a boolean if a field has been set.

### SetNStoppedNil

`func (o *ConditionRecruitmentDurationRow) SetNStoppedNil(b bool)`

 SetNStoppedNil sets the value for NStopped to be an explicit nil

### UnsetNStopped
`func (o *ConditionRecruitmentDurationRow) UnsetNStopped()`

UnsetNStopped ensures that no value is present for NStopped, not even an explicit nil
### GetNNoRegistrationRecord

`func (o *ConditionRecruitmentDurationRow) GetNNoRegistrationRecord() int32`

GetNNoRegistrationRecord returns the NNoRegistrationRecord field if non-nil, zero value otherwise.

### GetNNoRegistrationRecordOk

`func (o *ConditionRecruitmentDurationRow) GetNNoRegistrationRecordOk() (*int32, bool)`

GetNNoRegistrationRecordOk returns a tuple with the NNoRegistrationRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNNoRegistrationRecord

`func (o *ConditionRecruitmentDurationRow) SetNNoRegistrationRecord(v int32)`

SetNNoRegistrationRecord sets NNoRegistrationRecord field to given value.

### HasNNoRegistrationRecord

`func (o *ConditionRecruitmentDurationRow) HasNNoRegistrationRecord() bool`

HasNNoRegistrationRecord returns a boolean if a field has been set.

### SetNNoRegistrationRecordNil

`func (o *ConditionRecruitmentDurationRow) SetNNoRegistrationRecordNil(b bool)`

 SetNNoRegistrationRecordNil sets the value for NNoRegistrationRecord to be an explicit nil

### UnsetNNoRegistrationRecord
`func (o *ConditionRecruitmentDurationRow) UnsetNNoRegistrationRecord()`

UnsetNNoRegistrationRecord ensures that no value is present for NNoRegistrationRecord, not even an explicit nil
### GetNRetrospectiveRegistration

`func (o *ConditionRecruitmentDurationRow) GetNRetrospectiveRegistration() int32`

GetNRetrospectiveRegistration returns the NRetrospectiveRegistration field if non-nil, zero value otherwise.

### GetNRetrospectiveRegistrationOk

`func (o *ConditionRecruitmentDurationRow) GetNRetrospectiveRegistrationOk() (*int32, bool)`

GetNRetrospectiveRegistrationOk returns a tuple with the NRetrospectiveRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRetrospectiveRegistration

`func (o *ConditionRecruitmentDurationRow) SetNRetrospectiveRegistration(v int32)`

SetNRetrospectiveRegistration sets NRetrospectiveRegistration field to given value.

### HasNRetrospectiveRegistration

`func (o *ConditionRecruitmentDurationRow) HasNRetrospectiveRegistration() bool`

HasNRetrospectiveRegistration returns a boolean if a field has been set.

### SetNRetrospectiveRegistrationNil

`func (o *ConditionRecruitmentDurationRow) SetNRetrospectiveRegistrationNil(b bool)`

 SetNRetrospectiveRegistrationNil sets the value for NRetrospectiveRegistration to be an explicit nil

### UnsetNRetrospectiveRegistration
`func (o *ConditionRecruitmentDurationRow) UnsetNRetrospectiveRegistration()`

UnsetNRetrospectiveRegistration ensures that no value is present for NRetrospectiveRegistration, not even an explicit nil
### GetNStartNotActual

`func (o *ConditionRecruitmentDurationRow) GetNStartNotActual() int32`

GetNStartNotActual returns the NStartNotActual field if non-nil, zero value otherwise.

### GetNStartNotActualOk

`func (o *ConditionRecruitmentDurationRow) GetNStartNotActualOk() (*int32, bool)`

GetNStartNotActualOk returns a tuple with the NStartNotActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNStartNotActual

`func (o *ConditionRecruitmentDurationRow) SetNStartNotActual(v int32)`

SetNStartNotActual sets NStartNotActual field to given value.

### HasNStartNotActual

`func (o *ConditionRecruitmentDurationRow) HasNStartNotActual() bool`

HasNStartNotActual returns a boolean if a field has been set.

### SetNStartNotActualNil

`func (o *ConditionRecruitmentDurationRow) SetNStartNotActualNil(b bool)`

 SetNStartNotActualNil sets the value for NStartNotActual to be an explicit nil

### UnsetNStartNotActual
`func (o *ConditionRecruitmentDurationRow) UnsetNStartNotActual()`

UnsetNStartNotActual ensures that no value is present for NStartNotActual, not even an explicit nil
### GetNFitFailed

`func (o *ConditionRecruitmentDurationRow) GetNFitFailed() int32`

GetNFitFailed returns the NFitFailed field if non-nil, zero value otherwise.

### GetNFitFailedOk

`func (o *ConditionRecruitmentDurationRow) GetNFitFailedOk() (*int32, bool)`

GetNFitFailedOk returns a tuple with the NFitFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFitFailed

`func (o *ConditionRecruitmentDurationRow) SetNFitFailed(v int32)`

SetNFitFailed sets NFitFailed field to given value.

### HasNFitFailed

`func (o *ConditionRecruitmentDurationRow) HasNFitFailed() bool`

HasNFitFailed returns a boolean if a field has been set.

### SetNFitFailedNil

`func (o *ConditionRecruitmentDurationRow) SetNFitFailedNil(b bool)`

 SetNFitFailedNil sets the value for NFitFailed to be an explicit nil

### UnsetNFitFailed
`func (o *ConditionRecruitmentDurationRow) UnsetNFitFailed()`

UnsetNFitFailed ensures that no value is present for NFitFailed, not even an explicit nil
### GetNPredicted

`func (o *ConditionRecruitmentDurationRow) GetNPredicted() int32`

GetNPredicted returns the NPredicted field if non-nil, zero value otherwise.

### GetNPredictedOk

`func (o *ConditionRecruitmentDurationRow) GetNPredictedOk() (*int32, bool)`

GetNPredictedOk returns a tuple with the NPredicted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredicted

`func (o *ConditionRecruitmentDurationRow) SetNPredicted(v int32)`

SetNPredicted sets NPredicted field to given value.

### HasNPredicted

`func (o *ConditionRecruitmentDurationRow) HasNPredicted() bool`

HasNPredicted returns a boolean if a field has been set.

### SetNPredictedNil

`func (o *ConditionRecruitmentDurationRow) SetNPredictedNil(b bool)`

 SetNPredictedNil sets the value for NPredicted to be an explicit nil

### UnsetNPredicted
`func (o *ConditionRecruitmentDurationRow) UnsetNPredicted()`

UnsetNPredicted ensures that no value is present for NPredicted, not even an explicit nil
### GetNPredictedSuperseded

`func (o *ConditionRecruitmentDurationRow) GetNPredictedSuperseded() int32`

GetNPredictedSuperseded returns the NPredictedSuperseded field if non-nil, zero value otherwise.

### GetNPredictedSupersededOk

`func (o *ConditionRecruitmentDurationRow) GetNPredictedSupersededOk() (*int32, bool)`

GetNPredictedSupersededOk returns a tuple with the NPredictedSuperseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedSuperseded

`func (o *ConditionRecruitmentDurationRow) SetNPredictedSuperseded(v int32)`

SetNPredictedSuperseded sets NPredictedSuperseded field to given value.

### HasNPredictedSuperseded

`func (o *ConditionRecruitmentDurationRow) HasNPredictedSuperseded() bool`

HasNPredictedSuperseded returns a boolean if a field has been set.

### SetNPredictedSupersededNil

`func (o *ConditionRecruitmentDurationRow) SetNPredictedSupersededNil(b bool)`

 SetNPredictedSupersededNil sets the value for NPredictedSuperseded to be an explicit nil

### UnsetNPredictedSuperseded
`func (o *ConditionRecruitmentDurationRow) UnsetNPredictedSuperseded()`

UnsetNPredictedSuperseded ensures that no value is present for NPredictedSuperseded, not even an explicit nil
### GetNPredictedOpen

`func (o *ConditionRecruitmentDurationRow) GetNPredictedOpen() int32`

GetNPredictedOpen returns the NPredictedOpen field if non-nil, zero value otherwise.

### GetNPredictedOpenOk

`func (o *ConditionRecruitmentDurationRow) GetNPredictedOpenOk() (*int32, bool)`

GetNPredictedOpenOk returns a tuple with the NPredictedOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedOpen

`func (o *ConditionRecruitmentDurationRow) SetNPredictedOpen(v int32)`

SetNPredictedOpen sets NPredictedOpen field to given value.

### HasNPredictedOpen

`func (o *ConditionRecruitmentDurationRow) HasNPredictedOpen() bool`

HasNPredictedOpen returns a boolean if a field has been set.

### SetNPredictedOpenNil

`func (o *ConditionRecruitmentDurationRow) SetNPredictedOpenNil(b bool)`

 SetNPredictedOpenNil sets the value for NPredictedOpen to be an explicit nil

### UnsetNPredictedOpen
`func (o *ConditionRecruitmentDurationRow) UnsetNPredictedOpen()`

UnsetNPredictedOpen ensures that no value is present for NPredictedOpen, not even an explicit nil
### GetNPredictedAtCap

`func (o *ConditionRecruitmentDurationRow) GetNPredictedAtCap() int32`

GetNPredictedAtCap returns the NPredictedAtCap field if non-nil, zero value otherwise.

### GetNPredictedAtCapOk

`func (o *ConditionRecruitmentDurationRow) GetNPredictedAtCapOk() (*int32, bool)`

GetNPredictedAtCapOk returns a tuple with the NPredictedAtCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedAtCap

`func (o *ConditionRecruitmentDurationRow) SetNPredictedAtCap(v int32)`

SetNPredictedAtCap sets NPredictedAtCap field to given value.

### HasNPredictedAtCap

`func (o *ConditionRecruitmentDurationRow) HasNPredictedAtCap() bool`

HasNPredictedAtCap returns a boolean if a field has been set.

### SetNPredictedAtCapNil

`func (o *ConditionRecruitmentDurationRow) SetNPredictedAtCapNil(b bool)`

 SetNPredictedAtCapNil sets the value for NPredictedAtCap to be an explicit nil

### UnsetNPredictedAtCap
`func (o *ConditionRecruitmentDurationRow) UnsetNPredictedAtCap()`

UnsetNPredictedAtCap ensures that no value is present for NPredictedAtCap, not even an explicit nil
### GetNPredictedOverdue

`func (o *ConditionRecruitmentDurationRow) GetNPredictedOverdue() int32`

GetNPredictedOverdue returns the NPredictedOverdue field if non-nil, zero value otherwise.

### GetNPredictedOverdueOk

`func (o *ConditionRecruitmentDurationRow) GetNPredictedOverdueOk() (*int32, bool)`

GetNPredictedOverdueOk returns a tuple with the NPredictedOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedOverdue

`func (o *ConditionRecruitmentDurationRow) SetNPredictedOverdue(v int32)`

SetNPredictedOverdue sets NPredictedOverdue field to given value.

### HasNPredictedOverdue

`func (o *ConditionRecruitmentDurationRow) HasNPredictedOverdue() bool`

HasNPredictedOverdue returns a boolean if a field has been set.

### SetNPredictedOverdueNil

`func (o *ConditionRecruitmentDurationRow) SetNPredictedOverdueNil(b bool)`

 SetNPredictedOverdueNil sets the value for NPredictedOverdue to be an explicit nil

### UnsetNPredictedOverdue
`func (o *ConditionRecruitmentDurationRow) UnsetNPredictedOverdue()`

UnsetNPredictedOverdue ensures that no value is present for NPredictedOverdue, not even an explicit nil
### GetNPredictedErrorMeasured

`func (o *ConditionRecruitmentDurationRow) GetNPredictedErrorMeasured() int32`

GetNPredictedErrorMeasured returns the NPredictedErrorMeasured field if non-nil, zero value otherwise.

### GetNPredictedErrorMeasuredOk

`func (o *ConditionRecruitmentDurationRow) GetNPredictedErrorMeasuredOk() (*int32, bool)`

GetNPredictedErrorMeasuredOk returns a tuple with the NPredictedErrorMeasured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedErrorMeasured

`func (o *ConditionRecruitmentDurationRow) SetNPredictedErrorMeasured(v int32)`

SetNPredictedErrorMeasured sets NPredictedErrorMeasured field to given value.

### HasNPredictedErrorMeasured

`func (o *ConditionRecruitmentDurationRow) HasNPredictedErrorMeasured() bool`

HasNPredictedErrorMeasured returns a boolean if a field has been set.

### SetNPredictedErrorMeasuredNil

`func (o *ConditionRecruitmentDurationRow) SetNPredictedErrorMeasuredNil(b bool)`

 SetNPredictedErrorMeasuredNil sets the value for NPredictedErrorMeasured to be an explicit nil

### UnsetNPredictedErrorMeasured
`func (o *ConditionRecruitmentDurationRow) UnsetNPredictedErrorMeasured()`

UnsetNPredictedErrorMeasured ensures that no value is present for NPredictedErrorMeasured, not even an explicit nil
### GetPredictedMonthsMedian

`func (o *ConditionRecruitmentDurationRow) GetPredictedMonthsMedian() float32`

GetPredictedMonthsMedian returns the PredictedMonthsMedian field if non-nil, zero value otherwise.

### GetPredictedMonthsMedianOk

`func (o *ConditionRecruitmentDurationRow) GetPredictedMonthsMedianOk() (*float32, bool)`

GetPredictedMonthsMedianOk returns a tuple with the PredictedMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMonthsMedian

`func (o *ConditionRecruitmentDurationRow) SetPredictedMonthsMedian(v float32)`

SetPredictedMonthsMedian sets PredictedMonthsMedian field to given value.

### HasPredictedMonthsMedian

`func (o *ConditionRecruitmentDurationRow) HasPredictedMonthsMedian() bool`

HasPredictedMonthsMedian returns a boolean if a field has been set.

### SetPredictedMonthsMedianNil

`func (o *ConditionRecruitmentDurationRow) SetPredictedMonthsMedianNil(b bool)`

 SetPredictedMonthsMedianNil sets the value for PredictedMonthsMedian to be an explicit nil

### UnsetPredictedMonthsMedian
`func (o *ConditionRecruitmentDurationRow) UnsetPredictedMonthsMedian()`

UnsetPredictedMonthsMedian ensures that no value is present for PredictedMonthsMedian, not even an explicit nil
### GetPredictedMonthsP25

`func (o *ConditionRecruitmentDurationRow) GetPredictedMonthsP25() float32`

GetPredictedMonthsP25 returns the PredictedMonthsP25 field if non-nil, zero value otherwise.

### GetPredictedMonthsP25Ok

`func (o *ConditionRecruitmentDurationRow) GetPredictedMonthsP25Ok() (*float32, bool)`

GetPredictedMonthsP25Ok returns a tuple with the PredictedMonthsP25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMonthsP25

`func (o *ConditionRecruitmentDurationRow) SetPredictedMonthsP25(v float32)`

SetPredictedMonthsP25 sets PredictedMonthsP25 field to given value.

### HasPredictedMonthsP25

`func (o *ConditionRecruitmentDurationRow) HasPredictedMonthsP25() bool`

HasPredictedMonthsP25 returns a boolean if a field has been set.

### SetPredictedMonthsP25Nil

`func (o *ConditionRecruitmentDurationRow) SetPredictedMonthsP25Nil(b bool)`

 SetPredictedMonthsP25Nil sets the value for PredictedMonthsP25 to be an explicit nil

### UnsetPredictedMonthsP25
`func (o *ConditionRecruitmentDurationRow) UnsetPredictedMonthsP25()`

UnsetPredictedMonthsP25 ensures that no value is present for PredictedMonthsP25, not even an explicit nil
### GetPredictedMonthsP75

`func (o *ConditionRecruitmentDurationRow) GetPredictedMonthsP75() float32`

GetPredictedMonthsP75 returns the PredictedMonthsP75 field if non-nil, zero value otherwise.

### GetPredictedMonthsP75Ok

`func (o *ConditionRecruitmentDurationRow) GetPredictedMonthsP75Ok() (*float32, bool)`

GetPredictedMonthsP75Ok returns a tuple with the PredictedMonthsP75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMonthsP75

`func (o *ConditionRecruitmentDurationRow) SetPredictedMonthsP75(v float32)`

SetPredictedMonthsP75 sets PredictedMonthsP75 field to given value.

### HasPredictedMonthsP75

`func (o *ConditionRecruitmentDurationRow) HasPredictedMonthsP75() bool`

HasPredictedMonthsP75 returns a boolean if a field has been set.

### SetPredictedMonthsP75Nil

`func (o *ConditionRecruitmentDurationRow) SetPredictedMonthsP75Nil(b bool)`

 SetPredictedMonthsP75Nil sets the value for PredictedMonthsP75 to be an explicit nil

### UnsetPredictedMonthsP75
`func (o *ConditionRecruitmentDurationRow) UnsetPredictedMonthsP75()`

UnsetPredictedMonthsP75 ensures that no value is present for PredictedMonthsP75, not even an explicit nil
### GetBaselineMonthsMedian

`func (o *ConditionRecruitmentDurationRow) GetBaselineMonthsMedian() float32`

GetBaselineMonthsMedian returns the BaselineMonthsMedian field if non-nil, zero value otherwise.

### GetBaselineMonthsMedianOk

`func (o *ConditionRecruitmentDurationRow) GetBaselineMonthsMedianOk() (*float32, bool)`

GetBaselineMonthsMedianOk returns a tuple with the BaselineMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineMonthsMedian

`func (o *ConditionRecruitmentDurationRow) SetBaselineMonthsMedian(v float32)`

SetBaselineMonthsMedian sets BaselineMonthsMedian field to given value.

### HasBaselineMonthsMedian

`func (o *ConditionRecruitmentDurationRow) HasBaselineMonthsMedian() bool`

HasBaselineMonthsMedian returns a boolean if a field has been set.

### SetBaselineMonthsMedianNil

`func (o *ConditionRecruitmentDurationRow) SetBaselineMonthsMedianNil(b bool)`

 SetBaselineMonthsMedianNil sets the value for BaselineMonthsMedian to be an explicit nil

### UnsetBaselineMonthsMedian
`func (o *ConditionRecruitmentDurationRow) UnsetBaselineMonthsMedian()`

UnsetBaselineMonthsMedian ensures that no value is present for BaselineMonthsMedian, not even an explicit nil
### GetNRealized

`func (o *ConditionRecruitmentDurationRow) GetNRealized() int32`

GetNRealized returns the NRealized field if non-nil, zero value otherwise.

### GetNRealizedOk

`func (o *ConditionRecruitmentDurationRow) GetNRealizedOk() (*int32, bool)`

GetNRealizedOk returns a tuple with the NRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealized

`func (o *ConditionRecruitmentDurationRow) SetNRealized(v int32)`

SetNRealized sets NRealized field to given value.

### HasNRealized

`func (o *ConditionRecruitmentDurationRow) HasNRealized() bool`

HasNRealized returns a boolean if a field has been set.

### SetNRealizedNil

`func (o *ConditionRecruitmentDurationRow) SetNRealizedNil(b bool)`

 SetNRealizedNil sets the value for NRealized to be an explicit nil

### UnsetNRealized
`func (o *ConditionRecruitmentDurationRow) UnsetNRealized()`

UnsetNRealized ensures that no value is present for NRealized, not even an explicit nil
### GetRealizedMonthsMedian

`func (o *ConditionRecruitmentDurationRow) GetRealizedMonthsMedian() float32`

GetRealizedMonthsMedian returns the RealizedMonthsMedian field if non-nil, zero value otherwise.

### GetRealizedMonthsMedianOk

`func (o *ConditionRecruitmentDurationRow) GetRealizedMonthsMedianOk() (*float32, bool)`

GetRealizedMonthsMedianOk returns a tuple with the RealizedMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedMonthsMedian

`func (o *ConditionRecruitmentDurationRow) SetRealizedMonthsMedian(v float32)`

SetRealizedMonthsMedian sets RealizedMonthsMedian field to given value.

### HasRealizedMonthsMedian

`func (o *ConditionRecruitmentDurationRow) HasRealizedMonthsMedian() bool`

HasRealizedMonthsMedian returns a boolean if a field has been set.

### SetRealizedMonthsMedianNil

`func (o *ConditionRecruitmentDurationRow) SetRealizedMonthsMedianNil(b bool)`

 SetRealizedMonthsMedianNil sets the value for RealizedMonthsMedian to be an explicit nil

### UnsetRealizedMonthsMedian
`func (o *ConditionRecruitmentDurationRow) UnsetRealizedMonthsMedian()`

UnsetRealizedMonthsMedian ensures that no value is present for RealizedMonthsMedian, not even an explicit nil
### GetRealizedMonthsP25

`func (o *ConditionRecruitmentDurationRow) GetRealizedMonthsP25() float32`

GetRealizedMonthsP25 returns the RealizedMonthsP25 field if non-nil, zero value otherwise.

### GetRealizedMonthsP25Ok

`func (o *ConditionRecruitmentDurationRow) GetRealizedMonthsP25Ok() (*float32, bool)`

GetRealizedMonthsP25Ok returns a tuple with the RealizedMonthsP25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedMonthsP25

`func (o *ConditionRecruitmentDurationRow) SetRealizedMonthsP25(v float32)`

SetRealizedMonthsP25 sets RealizedMonthsP25 field to given value.

### HasRealizedMonthsP25

`func (o *ConditionRecruitmentDurationRow) HasRealizedMonthsP25() bool`

HasRealizedMonthsP25 returns a boolean if a field has been set.

### SetRealizedMonthsP25Nil

`func (o *ConditionRecruitmentDurationRow) SetRealizedMonthsP25Nil(b bool)`

 SetRealizedMonthsP25Nil sets the value for RealizedMonthsP25 to be an explicit nil

### UnsetRealizedMonthsP25
`func (o *ConditionRecruitmentDurationRow) UnsetRealizedMonthsP25()`

UnsetRealizedMonthsP25 ensures that no value is present for RealizedMonthsP25, not even an explicit nil
### GetRealizedMonthsP75

`func (o *ConditionRecruitmentDurationRow) GetRealizedMonthsP75() float32`

GetRealizedMonthsP75 returns the RealizedMonthsP75 field if non-nil, zero value otherwise.

### GetRealizedMonthsP75Ok

`func (o *ConditionRecruitmentDurationRow) GetRealizedMonthsP75Ok() (*float32, bool)`

GetRealizedMonthsP75Ok returns a tuple with the RealizedMonthsP75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedMonthsP75

`func (o *ConditionRecruitmentDurationRow) SetRealizedMonthsP75(v float32)`

SetRealizedMonthsP75 sets RealizedMonthsP75 field to given value.

### HasRealizedMonthsP75

`func (o *ConditionRecruitmentDurationRow) HasRealizedMonthsP75() bool`

HasRealizedMonthsP75 returns a boolean if a field has been set.

### SetRealizedMonthsP75Nil

`func (o *ConditionRecruitmentDurationRow) SetRealizedMonthsP75Nil(b bool)`

 SetRealizedMonthsP75Nil sets the value for RealizedMonthsP75 to be an explicit nil

### UnsetRealizedMonthsP75
`func (o *ConditionRecruitmentDurationRow) UnsetRealizedMonthsP75()`

UnsetRealizedMonthsP75 ensures that no value is present for RealizedMonthsP75, not even an explicit nil
### GetNRealizedOverCap

`func (o *ConditionRecruitmentDurationRow) GetNRealizedOverCap() int32`

GetNRealizedOverCap returns the NRealizedOverCap field if non-nil, zero value otherwise.

### GetNRealizedOverCapOk

`func (o *ConditionRecruitmentDurationRow) GetNRealizedOverCapOk() (*int32, bool)`

GetNRealizedOverCapOk returns a tuple with the NRealizedOverCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealizedOverCap

`func (o *ConditionRecruitmentDurationRow) SetNRealizedOverCap(v int32)`

SetNRealizedOverCap sets NRealizedOverCap field to given value.

### HasNRealizedOverCap

`func (o *ConditionRecruitmentDurationRow) HasNRealizedOverCap() bool`

HasNRealizedOverCap returns a boolean if a field has been set.

### SetNRealizedOverCapNil

`func (o *ConditionRecruitmentDurationRow) SetNRealizedOverCapNil(b bool)`

 SetNRealizedOverCapNil sets the value for NRealizedOverCap to be an explicit nil

### UnsetNRealizedOverCap
`func (o *ConditionRecruitmentDurationRow) UnsetNRealizedOverCap()`

UnsetNRealizedOverCap ensures that no value is present for NRealizedOverCap, not even an explicit nil
### GetMinN

`func (o *ConditionRecruitmentDurationRow) GetMinN() int32`

GetMinN returns the MinN field if non-nil, zero value otherwise.

### GetMinNOk

`func (o *ConditionRecruitmentDurationRow) GetMinNOk() (*int32, bool)`

GetMinNOk returns a tuple with the MinN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinN

`func (o *ConditionRecruitmentDurationRow) SetMinN(v int32)`

SetMinN sets MinN field to given value.

### HasMinN

`func (o *ConditionRecruitmentDurationRow) HasMinN() bool`

HasMinN returns a boolean if a field has been set.

### SetMinNNil

`func (o *ConditionRecruitmentDurationRow) SetMinNNil(b bool)`

 SetMinNNil sets the value for MinN to be an explicit nil

### UnsetMinN
`func (o *ConditionRecruitmentDurationRow) UnsetMinN()`

UnsetMinN ensures that no value is present for MinN, not even an explicit nil
### GetRollupStatus

`func (o *ConditionRecruitmentDurationRow) GetRollupStatus() string`

GetRollupStatus returns the RollupStatus field if non-nil, zero value otherwise.

### GetRollupStatusOk

`func (o *ConditionRecruitmentDurationRow) GetRollupStatusOk() (*string, bool)`

GetRollupStatusOk returns a tuple with the RollupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupStatus

`func (o *ConditionRecruitmentDurationRow) SetRollupStatus(v string)`

SetRollupStatus sets RollupStatus field to given value.

### HasRollupStatus

`func (o *ConditionRecruitmentDurationRow) HasRollupStatus() bool`

HasRollupStatus returns a boolean if a field has been set.

### SetRollupStatusNil

`func (o *ConditionRecruitmentDurationRow) SetRollupStatusNil(b bool)`

 SetRollupStatusNil sets the value for RollupStatus to be an explicit nil

### UnsetRollupStatus
`func (o *ConditionRecruitmentDurationRow) UnsetRollupStatus()`

UnsetRollupStatus ensures that no value is present for RollupStatus, not even an explicit nil
### GetRealizedStatus

`func (o *ConditionRecruitmentDurationRow) GetRealizedStatus() string`

GetRealizedStatus returns the RealizedStatus field if non-nil, zero value otherwise.

### GetRealizedStatusOk

`func (o *ConditionRecruitmentDurationRow) GetRealizedStatusOk() (*string, bool)`

GetRealizedStatusOk returns a tuple with the RealizedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedStatus

`func (o *ConditionRecruitmentDurationRow) SetRealizedStatus(v string)`

SetRealizedStatus sets RealizedStatus field to given value.

### HasRealizedStatus

`func (o *ConditionRecruitmentDurationRow) HasRealizedStatus() bool`

HasRealizedStatus returns a boolean if a field has been set.

### SetRealizedStatusNil

`func (o *ConditionRecruitmentDurationRow) SetRealizedStatusNil(b bool)`

 SetRealizedStatusNil sets the value for RealizedStatus to be an explicit nil

### UnsetRealizedStatus
`func (o *ConditionRecruitmentDurationRow) UnsetRealizedStatus()`

UnsetRealizedStatus ensures that no value is present for RealizedStatus, not even an explicit nil
### GetTrainingCutoff

`func (o *ConditionRecruitmentDurationRow) GetTrainingCutoff() string`

GetTrainingCutoff returns the TrainingCutoff field if non-nil, zero value otherwise.

### GetTrainingCutoffOk

`func (o *ConditionRecruitmentDurationRow) GetTrainingCutoffOk() (*string, bool)`

GetTrainingCutoffOk returns a tuple with the TrainingCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingCutoff

`func (o *ConditionRecruitmentDurationRow) SetTrainingCutoff(v string)`

SetTrainingCutoff sets TrainingCutoff field to given value.

### HasTrainingCutoff

`func (o *ConditionRecruitmentDurationRow) HasTrainingCutoff() bool`

HasTrainingCutoff returns a boolean if a field has been set.

### SetTrainingCutoffNil

`func (o *ConditionRecruitmentDurationRow) SetTrainingCutoffNil(b bool)`

 SetTrainingCutoffNil sets the value for TrainingCutoff to be an explicit nil

### UnsetTrainingCutoff
`func (o *ConditionRecruitmentDurationRow) UnsetTrainingCutoff()`

UnsetTrainingCutoff ensures that no value is present for TrainingCutoff, not even an explicit nil
### GetModelName

`func (o *ConditionRecruitmentDurationRow) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ConditionRecruitmentDurationRow) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ConditionRecruitmentDurationRow) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ConditionRecruitmentDurationRow) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### SetModelNameNil

`func (o *ConditionRecruitmentDurationRow) SetModelNameNil(b bool)`

 SetModelNameNil sets the value for ModelName to be an explicit nil

### UnsetModelName
`func (o *ConditionRecruitmentDurationRow) UnsetModelName()`

UnsetModelName ensures that no value is present for ModelName, not even an explicit nil
### GetModelVersion

`func (o *ConditionRecruitmentDurationRow) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ConditionRecruitmentDurationRow) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ConditionRecruitmentDurationRow) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ConditionRecruitmentDurationRow) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### SetModelVersionNil

`func (o *ConditionRecruitmentDurationRow) SetModelVersionNil(b bool)`

 SetModelVersionNil sets the value for ModelVersion to be an explicit nil

### UnsetModelVersion
`func (o *ConditionRecruitmentDurationRow) UnsetModelVersion()`

UnsetModelVersion ensures that no value is present for ModelVersion, not even an explicit nil
### GetHorizonCapMonths

`func (o *ConditionRecruitmentDurationRow) GetHorizonCapMonths() int32`

GetHorizonCapMonths returns the HorizonCapMonths field if non-nil, zero value otherwise.

### GetHorizonCapMonthsOk

`func (o *ConditionRecruitmentDurationRow) GetHorizonCapMonthsOk() (*int32, bool)`

GetHorizonCapMonthsOk returns a tuple with the HorizonCapMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizonCapMonths

`func (o *ConditionRecruitmentDurationRow) SetHorizonCapMonths(v int32)`

SetHorizonCapMonths sets HorizonCapMonths field to given value.

### HasHorizonCapMonths

`func (o *ConditionRecruitmentDurationRow) HasHorizonCapMonths() bool`

HasHorizonCapMonths returns a boolean if a field has been set.

### SetHorizonCapMonthsNil

`func (o *ConditionRecruitmentDurationRow) SetHorizonCapMonthsNil(b bool)`

 SetHorizonCapMonthsNil sets the value for HorizonCapMonths to be an explicit nil

### UnsetHorizonCapMonths
`func (o *ConditionRecruitmentDurationRow) UnsetHorizonCapMonths()`

UnsetHorizonCapMonths ensures that no value is present for HorizonCapMonths, not even an explicit nil
### GetMethod

`func (o *ConditionRecruitmentDurationRow) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ConditionRecruitmentDurationRow) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ConditionRecruitmentDurationRow) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ConditionRecruitmentDurationRow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *ConditionRecruitmentDurationRow) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *ConditionRecruitmentDurationRow) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


