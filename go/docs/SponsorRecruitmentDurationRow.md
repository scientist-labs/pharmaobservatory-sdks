# SponsorRecruitmentDurationRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the series is for (the lead sponsor). | [optional] 
**SponsorName** | Pointer to **NullableString** | Display name of the organization record. | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company the organization belongs to (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**NTrialsInSeries** | Pointer to **NullableInt32** | Trials the organization leads. | [optional] 
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
**NPredictedSuperseded** | Pointer to **NullableInt32** | Predicted trials whose CURRENT record shows the outcome (COMPLETED, TERMINATED, WITHDRAWN or an ACTUAL primary completion); kept out of every predicted figure, as the trial page withholds them. | [optional] 
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

### NewSponsorRecruitmentDurationRow

`func NewSponsorRecruitmentDurationRow() *SponsorRecruitmentDurationRow`

NewSponsorRecruitmentDurationRow instantiates a new SponsorRecruitmentDurationRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSponsorRecruitmentDurationRowWithDefaults

`func NewSponsorRecruitmentDurationRowWithDefaults() *SponsorRecruitmentDurationRow`

NewSponsorRecruitmentDurationRowWithDefaults instantiates a new SponsorRecruitmentDurationRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSponsorOrgId

`func (o *SponsorRecruitmentDurationRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *SponsorRecruitmentDurationRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *SponsorRecruitmentDurationRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *SponsorRecruitmentDurationRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *SponsorRecruitmentDurationRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *SponsorRecruitmentDurationRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorName

`func (o *SponsorRecruitmentDurationRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *SponsorRecruitmentDurationRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *SponsorRecruitmentDurationRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *SponsorRecruitmentDurationRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *SponsorRecruitmentDurationRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *SponsorRecruitmentDurationRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorCompanyId

`func (o *SponsorRecruitmentDurationRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *SponsorRecruitmentDurationRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *SponsorRecruitmentDurationRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *SponsorRecruitmentDurationRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *SponsorRecruitmentDurationRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *SponsorRecruitmentDurationRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *SponsorRecruitmentDurationRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *SponsorRecruitmentDurationRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *SponsorRecruitmentDurationRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *SponsorRecruitmentDurationRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *SponsorRecruitmentDurationRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *SponsorRecruitmentDurationRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *SponsorRecruitmentDurationRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *SponsorRecruitmentDurationRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *SponsorRecruitmentDurationRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *SponsorRecruitmentDurationRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *SponsorRecruitmentDurationRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *SponsorRecruitmentDurationRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *SponsorRecruitmentDurationRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *SponsorRecruitmentDurationRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *SponsorRecruitmentDurationRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *SponsorRecruitmentDurationRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *SponsorRecruitmentDurationRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *SponsorRecruitmentDurationRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetNTrialsInSeries

`func (o *SponsorRecruitmentDurationRow) GetNTrialsInSeries() int32`

GetNTrialsInSeries returns the NTrialsInSeries field if non-nil, zero value otherwise.

### GetNTrialsInSeriesOk

`func (o *SponsorRecruitmentDurationRow) GetNTrialsInSeriesOk() (*int32, bool)`

GetNTrialsInSeriesOk returns a tuple with the NTrialsInSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsInSeries

`func (o *SponsorRecruitmentDurationRow) SetNTrialsInSeries(v int32)`

SetNTrialsInSeries sets NTrialsInSeries field to given value.

### HasNTrialsInSeries

`func (o *SponsorRecruitmentDurationRow) HasNTrialsInSeries() bool`

HasNTrialsInSeries returns a boolean if a field has been set.

### SetNTrialsInSeriesNil

`func (o *SponsorRecruitmentDurationRow) SetNTrialsInSeriesNil(b bool)`

 SetNTrialsInSeriesNil sets the value for NTrialsInSeries to be an explicit nil

### UnsetNTrialsInSeries
`func (o *SponsorRecruitmentDurationRow) UnsetNTrialsInSeries()`

UnsetNTrialsInSeries ensures that no value is present for NTrialsInSeries, not even an explicit nil
### GetNEvaluated

`func (o *SponsorRecruitmentDurationRow) GetNEvaluated() int32`

GetNEvaluated returns the NEvaluated field if non-nil, zero value otherwise.

### GetNEvaluatedOk

`func (o *SponsorRecruitmentDurationRow) GetNEvaluatedOk() (*int32, bool)`

GetNEvaluatedOk returns a tuple with the NEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEvaluated

`func (o *SponsorRecruitmentDurationRow) SetNEvaluated(v int32)`

SetNEvaluated sets NEvaluated field to given value.

### HasNEvaluated

`func (o *SponsorRecruitmentDurationRow) HasNEvaluated() bool`

HasNEvaluated returns a boolean if a field has been set.

### SetNEvaluatedNil

`func (o *SponsorRecruitmentDurationRow) SetNEvaluatedNil(b bool)`

 SetNEvaluatedNil sets the value for NEvaluated to be an explicit nil

### UnsetNEvaluated
`func (o *SponsorRecruitmentDurationRow) UnsetNEvaluated()`

UnsetNEvaluated ensures that no value is present for NEvaluated, not even an explicit nil
### GetNNotEvaluated

`func (o *SponsorRecruitmentDurationRow) GetNNotEvaluated() int32`

GetNNotEvaluated returns the NNotEvaluated field if non-nil, zero value otherwise.

### GetNNotEvaluatedOk

`func (o *SponsorRecruitmentDurationRow) GetNNotEvaluatedOk() (*int32, bool)`

GetNNotEvaluatedOk returns a tuple with the NNotEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNNotEvaluated

`func (o *SponsorRecruitmentDurationRow) SetNNotEvaluated(v int32)`

SetNNotEvaluated sets NNotEvaluated field to given value.

### HasNNotEvaluated

`func (o *SponsorRecruitmentDurationRow) HasNNotEvaluated() bool`

HasNNotEvaluated returns a boolean if a field has been set.

### SetNNotEvaluatedNil

`func (o *SponsorRecruitmentDurationRow) SetNNotEvaluatedNil(b bool)`

 SetNNotEvaluatedNil sets the value for NNotEvaluated to be an explicit nil

### UnsetNNotEvaluated
`func (o *SponsorRecruitmentDurationRow) UnsetNNotEvaluated()`

UnsetNNotEvaluated ensures that no value is present for NNotEvaluated, not even an explicit nil
### GetNOutOfPopulation

`func (o *SponsorRecruitmentDurationRow) GetNOutOfPopulation() int32`

GetNOutOfPopulation returns the NOutOfPopulation field if non-nil, zero value otherwise.

### GetNOutOfPopulationOk

`func (o *SponsorRecruitmentDurationRow) GetNOutOfPopulationOk() (*int32, bool)`

GetNOutOfPopulationOk returns a tuple with the NOutOfPopulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOutOfPopulation

`func (o *SponsorRecruitmentDurationRow) SetNOutOfPopulation(v int32)`

SetNOutOfPopulation sets NOutOfPopulation field to given value.

### HasNOutOfPopulation

`func (o *SponsorRecruitmentDurationRow) HasNOutOfPopulation() bool`

HasNOutOfPopulation returns a boolean if a field has been set.

### SetNOutOfPopulationNil

`func (o *SponsorRecruitmentDurationRow) SetNOutOfPopulationNil(b bool)`

 SetNOutOfPopulationNil sets the value for NOutOfPopulation to be an explicit nil

### UnsetNOutOfPopulation
`func (o *SponsorRecruitmentDurationRow) UnsetNOutOfPopulation()`

UnsetNOutOfPopulation ensures that no value is present for NOutOfPopulation, not even an explicit nil
### GetNCompleted

`func (o *SponsorRecruitmentDurationRow) GetNCompleted() int32`

GetNCompleted returns the NCompleted field if non-nil, zero value otherwise.

### GetNCompletedOk

`func (o *SponsorRecruitmentDurationRow) GetNCompletedOk() (*int32, bool)`

GetNCompletedOk returns a tuple with the NCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompleted

`func (o *SponsorRecruitmentDurationRow) SetNCompleted(v int32)`

SetNCompleted sets NCompleted field to given value.

### HasNCompleted

`func (o *SponsorRecruitmentDurationRow) HasNCompleted() bool`

HasNCompleted returns a boolean if a field has been set.

### SetNCompletedNil

`func (o *SponsorRecruitmentDurationRow) SetNCompletedNil(b bool)`

 SetNCompletedNil sets the value for NCompleted to be an explicit nil

### UnsetNCompleted
`func (o *SponsorRecruitmentDurationRow) UnsetNCompleted()`

UnsetNCompleted ensures that no value is present for NCompleted, not even an explicit nil
### GetNStopped

`func (o *SponsorRecruitmentDurationRow) GetNStopped() int32`

GetNStopped returns the NStopped field if non-nil, zero value otherwise.

### GetNStoppedOk

`func (o *SponsorRecruitmentDurationRow) GetNStoppedOk() (*int32, bool)`

GetNStoppedOk returns a tuple with the NStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNStopped

`func (o *SponsorRecruitmentDurationRow) SetNStopped(v int32)`

SetNStopped sets NStopped field to given value.

### HasNStopped

`func (o *SponsorRecruitmentDurationRow) HasNStopped() bool`

HasNStopped returns a boolean if a field has been set.

### SetNStoppedNil

`func (o *SponsorRecruitmentDurationRow) SetNStoppedNil(b bool)`

 SetNStoppedNil sets the value for NStopped to be an explicit nil

### UnsetNStopped
`func (o *SponsorRecruitmentDurationRow) UnsetNStopped()`

UnsetNStopped ensures that no value is present for NStopped, not even an explicit nil
### GetNNoRegistrationRecord

`func (o *SponsorRecruitmentDurationRow) GetNNoRegistrationRecord() int32`

GetNNoRegistrationRecord returns the NNoRegistrationRecord field if non-nil, zero value otherwise.

### GetNNoRegistrationRecordOk

`func (o *SponsorRecruitmentDurationRow) GetNNoRegistrationRecordOk() (*int32, bool)`

GetNNoRegistrationRecordOk returns a tuple with the NNoRegistrationRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNNoRegistrationRecord

`func (o *SponsorRecruitmentDurationRow) SetNNoRegistrationRecord(v int32)`

SetNNoRegistrationRecord sets NNoRegistrationRecord field to given value.

### HasNNoRegistrationRecord

`func (o *SponsorRecruitmentDurationRow) HasNNoRegistrationRecord() bool`

HasNNoRegistrationRecord returns a boolean if a field has been set.

### SetNNoRegistrationRecordNil

`func (o *SponsorRecruitmentDurationRow) SetNNoRegistrationRecordNil(b bool)`

 SetNNoRegistrationRecordNil sets the value for NNoRegistrationRecord to be an explicit nil

### UnsetNNoRegistrationRecord
`func (o *SponsorRecruitmentDurationRow) UnsetNNoRegistrationRecord()`

UnsetNNoRegistrationRecord ensures that no value is present for NNoRegistrationRecord, not even an explicit nil
### GetNRetrospectiveRegistration

`func (o *SponsorRecruitmentDurationRow) GetNRetrospectiveRegistration() int32`

GetNRetrospectiveRegistration returns the NRetrospectiveRegistration field if non-nil, zero value otherwise.

### GetNRetrospectiveRegistrationOk

`func (o *SponsorRecruitmentDurationRow) GetNRetrospectiveRegistrationOk() (*int32, bool)`

GetNRetrospectiveRegistrationOk returns a tuple with the NRetrospectiveRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRetrospectiveRegistration

`func (o *SponsorRecruitmentDurationRow) SetNRetrospectiveRegistration(v int32)`

SetNRetrospectiveRegistration sets NRetrospectiveRegistration field to given value.

### HasNRetrospectiveRegistration

`func (o *SponsorRecruitmentDurationRow) HasNRetrospectiveRegistration() bool`

HasNRetrospectiveRegistration returns a boolean if a field has been set.

### SetNRetrospectiveRegistrationNil

`func (o *SponsorRecruitmentDurationRow) SetNRetrospectiveRegistrationNil(b bool)`

 SetNRetrospectiveRegistrationNil sets the value for NRetrospectiveRegistration to be an explicit nil

### UnsetNRetrospectiveRegistration
`func (o *SponsorRecruitmentDurationRow) UnsetNRetrospectiveRegistration()`

UnsetNRetrospectiveRegistration ensures that no value is present for NRetrospectiveRegistration, not even an explicit nil
### GetNStartNotActual

`func (o *SponsorRecruitmentDurationRow) GetNStartNotActual() int32`

GetNStartNotActual returns the NStartNotActual field if non-nil, zero value otherwise.

### GetNStartNotActualOk

`func (o *SponsorRecruitmentDurationRow) GetNStartNotActualOk() (*int32, bool)`

GetNStartNotActualOk returns a tuple with the NStartNotActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNStartNotActual

`func (o *SponsorRecruitmentDurationRow) SetNStartNotActual(v int32)`

SetNStartNotActual sets NStartNotActual field to given value.

### HasNStartNotActual

`func (o *SponsorRecruitmentDurationRow) HasNStartNotActual() bool`

HasNStartNotActual returns a boolean if a field has been set.

### SetNStartNotActualNil

`func (o *SponsorRecruitmentDurationRow) SetNStartNotActualNil(b bool)`

 SetNStartNotActualNil sets the value for NStartNotActual to be an explicit nil

### UnsetNStartNotActual
`func (o *SponsorRecruitmentDurationRow) UnsetNStartNotActual()`

UnsetNStartNotActual ensures that no value is present for NStartNotActual, not even an explicit nil
### GetNFitFailed

`func (o *SponsorRecruitmentDurationRow) GetNFitFailed() int32`

GetNFitFailed returns the NFitFailed field if non-nil, zero value otherwise.

### GetNFitFailedOk

`func (o *SponsorRecruitmentDurationRow) GetNFitFailedOk() (*int32, bool)`

GetNFitFailedOk returns a tuple with the NFitFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFitFailed

`func (o *SponsorRecruitmentDurationRow) SetNFitFailed(v int32)`

SetNFitFailed sets NFitFailed field to given value.

### HasNFitFailed

`func (o *SponsorRecruitmentDurationRow) HasNFitFailed() bool`

HasNFitFailed returns a boolean if a field has been set.

### SetNFitFailedNil

`func (o *SponsorRecruitmentDurationRow) SetNFitFailedNil(b bool)`

 SetNFitFailedNil sets the value for NFitFailed to be an explicit nil

### UnsetNFitFailed
`func (o *SponsorRecruitmentDurationRow) UnsetNFitFailed()`

UnsetNFitFailed ensures that no value is present for NFitFailed, not even an explicit nil
### GetNPredicted

`func (o *SponsorRecruitmentDurationRow) GetNPredicted() int32`

GetNPredicted returns the NPredicted field if non-nil, zero value otherwise.

### GetNPredictedOk

`func (o *SponsorRecruitmentDurationRow) GetNPredictedOk() (*int32, bool)`

GetNPredictedOk returns a tuple with the NPredicted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredicted

`func (o *SponsorRecruitmentDurationRow) SetNPredicted(v int32)`

SetNPredicted sets NPredicted field to given value.

### HasNPredicted

`func (o *SponsorRecruitmentDurationRow) HasNPredicted() bool`

HasNPredicted returns a boolean if a field has been set.

### SetNPredictedNil

`func (o *SponsorRecruitmentDurationRow) SetNPredictedNil(b bool)`

 SetNPredictedNil sets the value for NPredicted to be an explicit nil

### UnsetNPredicted
`func (o *SponsorRecruitmentDurationRow) UnsetNPredicted()`

UnsetNPredicted ensures that no value is present for NPredicted, not even an explicit nil
### GetNPredictedSuperseded

`func (o *SponsorRecruitmentDurationRow) GetNPredictedSuperseded() int32`

GetNPredictedSuperseded returns the NPredictedSuperseded field if non-nil, zero value otherwise.

### GetNPredictedSupersededOk

`func (o *SponsorRecruitmentDurationRow) GetNPredictedSupersededOk() (*int32, bool)`

GetNPredictedSupersededOk returns a tuple with the NPredictedSuperseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedSuperseded

`func (o *SponsorRecruitmentDurationRow) SetNPredictedSuperseded(v int32)`

SetNPredictedSuperseded sets NPredictedSuperseded field to given value.

### HasNPredictedSuperseded

`func (o *SponsorRecruitmentDurationRow) HasNPredictedSuperseded() bool`

HasNPredictedSuperseded returns a boolean if a field has been set.

### SetNPredictedSupersededNil

`func (o *SponsorRecruitmentDurationRow) SetNPredictedSupersededNil(b bool)`

 SetNPredictedSupersededNil sets the value for NPredictedSuperseded to be an explicit nil

### UnsetNPredictedSuperseded
`func (o *SponsorRecruitmentDurationRow) UnsetNPredictedSuperseded()`

UnsetNPredictedSuperseded ensures that no value is present for NPredictedSuperseded, not even an explicit nil
### GetNPredictedOpen

`func (o *SponsorRecruitmentDurationRow) GetNPredictedOpen() int32`

GetNPredictedOpen returns the NPredictedOpen field if non-nil, zero value otherwise.

### GetNPredictedOpenOk

`func (o *SponsorRecruitmentDurationRow) GetNPredictedOpenOk() (*int32, bool)`

GetNPredictedOpenOk returns a tuple with the NPredictedOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedOpen

`func (o *SponsorRecruitmentDurationRow) SetNPredictedOpen(v int32)`

SetNPredictedOpen sets NPredictedOpen field to given value.

### HasNPredictedOpen

`func (o *SponsorRecruitmentDurationRow) HasNPredictedOpen() bool`

HasNPredictedOpen returns a boolean if a field has been set.

### SetNPredictedOpenNil

`func (o *SponsorRecruitmentDurationRow) SetNPredictedOpenNil(b bool)`

 SetNPredictedOpenNil sets the value for NPredictedOpen to be an explicit nil

### UnsetNPredictedOpen
`func (o *SponsorRecruitmentDurationRow) UnsetNPredictedOpen()`

UnsetNPredictedOpen ensures that no value is present for NPredictedOpen, not even an explicit nil
### GetNPredictedAtCap

`func (o *SponsorRecruitmentDurationRow) GetNPredictedAtCap() int32`

GetNPredictedAtCap returns the NPredictedAtCap field if non-nil, zero value otherwise.

### GetNPredictedAtCapOk

`func (o *SponsorRecruitmentDurationRow) GetNPredictedAtCapOk() (*int32, bool)`

GetNPredictedAtCapOk returns a tuple with the NPredictedAtCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedAtCap

`func (o *SponsorRecruitmentDurationRow) SetNPredictedAtCap(v int32)`

SetNPredictedAtCap sets NPredictedAtCap field to given value.

### HasNPredictedAtCap

`func (o *SponsorRecruitmentDurationRow) HasNPredictedAtCap() bool`

HasNPredictedAtCap returns a boolean if a field has been set.

### SetNPredictedAtCapNil

`func (o *SponsorRecruitmentDurationRow) SetNPredictedAtCapNil(b bool)`

 SetNPredictedAtCapNil sets the value for NPredictedAtCap to be an explicit nil

### UnsetNPredictedAtCap
`func (o *SponsorRecruitmentDurationRow) UnsetNPredictedAtCap()`

UnsetNPredictedAtCap ensures that no value is present for NPredictedAtCap, not even an explicit nil
### GetNPredictedOverdue

`func (o *SponsorRecruitmentDurationRow) GetNPredictedOverdue() int32`

GetNPredictedOverdue returns the NPredictedOverdue field if non-nil, zero value otherwise.

### GetNPredictedOverdueOk

`func (o *SponsorRecruitmentDurationRow) GetNPredictedOverdueOk() (*int32, bool)`

GetNPredictedOverdueOk returns a tuple with the NPredictedOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedOverdue

`func (o *SponsorRecruitmentDurationRow) SetNPredictedOverdue(v int32)`

SetNPredictedOverdue sets NPredictedOverdue field to given value.

### HasNPredictedOverdue

`func (o *SponsorRecruitmentDurationRow) HasNPredictedOverdue() bool`

HasNPredictedOverdue returns a boolean if a field has been set.

### SetNPredictedOverdueNil

`func (o *SponsorRecruitmentDurationRow) SetNPredictedOverdueNil(b bool)`

 SetNPredictedOverdueNil sets the value for NPredictedOverdue to be an explicit nil

### UnsetNPredictedOverdue
`func (o *SponsorRecruitmentDurationRow) UnsetNPredictedOverdue()`

UnsetNPredictedOverdue ensures that no value is present for NPredictedOverdue, not even an explicit nil
### GetNPredictedErrorMeasured

`func (o *SponsorRecruitmentDurationRow) GetNPredictedErrorMeasured() int32`

GetNPredictedErrorMeasured returns the NPredictedErrorMeasured field if non-nil, zero value otherwise.

### GetNPredictedErrorMeasuredOk

`func (o *SponsorRecruitmentDurationRow) GetNPredictedErrorMeasuredOk() (*int32, bool)`

GetNPredictedErrorMeasuredOk returns a tuple with the NPredictedErrorMeasured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedErrorMeasured

`func (o *SponsorRecruitmentDurationRow) SetNPredictedErrorMeasured(v int32)`

SetNPredictedErrorMeasured sets NPredictedErrorMeasured field to given value.

### HasNPredictedErrorMeasured

`func (o *SponsorRecruitmentDurationRow) HasNPredictedErrorMeasured() bool`

HasNPredictedErrorMeasured returns a boolean if a field has been set.

### SetNPredictedErrorMeasuredNil

`func (o *SponsorRecruitmentDurationRow) SetNPredictedErrorMeasuredNil(b bool)`

 SetNPredictedErrorMeasuredNil sets the value for NPredictedErrorMeasured to be an explicit nil

### UnsetNPredictedErrorMeasured
`func (o *SponsorRecruitmentDurationRow) UnsetNPredictedErrorMeasured()`

UnsetNPredictedErrorMeasured ensures that no value is present for NPredictedErrorMeasured, not even an explicit nil
### GetPredictedMonthsMedian

`func (o *SponsorRecruitmentDurationRow) GetPredictedMonthsMedian() float32`

GetPredictedMonthsMedian returns the PredictedMonthsMedian field if non-nil, zero value otherwise.

### GetPredictedMonthsMedianOk

`func (o *SponsorRecruitmentDurationRow) GetPredictedMonthsMedianOk() (*float32, bool)`

GetPredictedMonthsMedianOk returns a tuple with the PredictedMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMonthsMedian

`func (o *SponsorRecruitmentDurationRow) SetPredictedMonthsMedian(v float32)`

SetPredictedMonthsMedian sets PredictedMonthsMedian field to given value.

### HasPredictedMonthsMedian

`func (o *SponsorRecruitmentDurationRow) HasPredictedMonthsMedian() bool`

HasPredictedMonthsMedian returns a boolean if a field has been set.

### SetPredictedMonthsMedianNil

`func (o *SponsorRecruitmentDurationRow) SetPredictedMonthsMedianNil(b bool)`

 SetPredictedMonthsMedianNil sets the value for PredictedMonthsMedian to be an explicit nil

### UnsetPredictedMonthsMedian
`func (o *SponsorRecruitmentDurationRow) UnsetPredictedMonthsMedian()`

UnsetPredictedMonthsMedian ensures that no value is present for PredictedMonthsMedian, not even an explicit nil
### GetPredictedMonthsP25

`func (o *SponsorRecruitmentDurationRow) GetPredictedMonthsP25() float32`

GetPredictedMonthsP25 returns the PredictedMonthsP25 field if non-nil, zero value otherwise.

### GetPredictedMonthsP25Ok

`func (o *SponsorRecruitmentDurationRow) GetPredictedMonthsP25Ok() (*float32, bool)`

GetPredictedMonthsP25Ok returns a tuple with the PredictedMonthsP25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMonthsP25

`func (o *SponsorRecruitmentDurationRow) SetPredictedMonthsP25(v float32)`

SetPredictedMonthsP25 sets PredictedMonthsP25 field to given value.

### HasPredictedMonthsP25

`func (o *SponsorRecruitmentDurationRow) HasPredictedMonthsP25() bool`

HasPredictedMonthsP25 returns a boolean if a field has been set.

### SetPredictedMonthsP25Nil

`func (o *SponsorRecruitmentDurationRow) SetPredictedMonthsP25Nil(b bool)`

 SetPredictedMonthsP25Nil sets the value for PredictedMonthsP25 to be an explicit nil

### UnsetPredictedMonthsP25
`func (o *SponsorRecruitmentDurationRow) UnsetPredictedMonthsP25()`

UnsetPredictedMonthsP25 ensures that no value is present for PredictedMonthsP25, not even an explicit nil
### GetPredictedMonthsP75

`func (o *SponsorRecruitmentDurationRow) GetPredictedMonthsP75() float32`

GetPredictedMonthsP75 returns the PredictedMonthsP75 field if non-nil, zero value otherwise.

### GetPredictedMonthsP75Ok

`func (o *SponsorRecruitmentDurationRow) GetPredictedMonthsP75Ok() (*float32, bool)`

GetPredictedMonthsP75Ok returns a tuple with the PredictedMonthsP75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMonthsP75

`func (o *SponsorRecruitmentDurationRow) SetPredictedMonthsP75(v float32)`

SetPredictedMonthsP75 sets PredictedMonthsP75 field to given value.

### HasPredictedMonthsP75

`func (o *SponsorRecruitmentDurationRow) HasPredictedMonthsP75() bool`

HasPredictedMonthsP75 returns a boolean if a field has been set.

### SetPredictedMonthsP75Nil

`func (o *SponsorRecruitmentDurationRow) SetPredictedMonthsP75Nil(b bool)`

 SetPredictedMonthsP75Nil sets the value for PredictedMonthsP75 to be an explicit nil

### UnsetPredictedMonthsP75
`func (o *SponsorRecruitmentDurationRow) UnsetPredictedMonthsP75()`

UnsetPredictedMonthsP75 ensures that no value is present for PredictedMonthsP75, not even an explicit nil
### GetBaselineMonthsMedian

`func (o *SponsorRecruitmentDurationRow) GetBaselineMonthsMedian() float32`

GetBaselineMonthsMedian returns the BaselineMonthsMedian field if non-nil, zero value otherwise.

### GetBaselineMonthsMedianOk

`func (o *SponsorRecruitmentDurationRow) GetBaselineMonthsMedianOk() (*float32, bool)`

GetBaselineMonthsMedianOk returns a tuple with the BaselineMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineMonthsMedian

`func (o *SponsorRecruitmentDurationRow) SetBaselineMonthsMedian(v float32)`

SetBaselineMonthsMedian sets BaselineMonthsMedian field to given value.

### HasBaselineMonthsMedian

`func (o *SponsorRecruitmentDurationRow) HasBaselineMonthsMedian() bool`

HasBaselineMonthsMedian returns a boolean if a field has been set.

### SetBaselineMonthsMedianNil

`func (o *SponsorRecruitmentDurationRow) SetBaselineMonthsMedianNil(b bool)`

 SetBaselineMonthsMedianNil sets the value for BaselineMonthsMedian to be an explicit nil

### UnsetBaselineMonthsMedian
`func (o *SponsorRecruitmentDurationRow) UnsetBaselineMonthsMedian()`

UnsetBaselineMonthsMedian ensures that no value is present for BaselineMonthsMedian, not even an explicit nil
### GetNRealized

`func (o *SponsorRecruitmentDurationRow) GetNRealized() int32`

GetNRealized returns the NRealized field if non-nil, zero value otherwise.

### GetNRealizedOk

`func (o *SponsorRecruitmentDurationRow) GetNRealizedOk() (*int32, bool)`

GetNRealizedOk returns a tuple with the NRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealized

`func (o *SponsorRecruitmentDurationRow) SetNRealized(v int32)`

SetNRealized sets NRealized field to given value.

### HasNRealized

`func (o *SponsorRecruitmentDurationRow) HasNRealized() bool`

HasNRealized returns a boolean if a field has been set.

### SetNRealizedNil

`func (o *SponsorRecruitmentDurationRow) SetNRealizedNil(b bool)`

 SetNRealizedNil sets the value for NRealized to be an explicit nil

### UnsetNRealized
`func (o *SponsorRecruitmentDurationRow) UnsetNRealized()`

UnsetNRealized ensures that no value is present for NRealized, not even an explicit nil
### GetRealizedMonthsMedian

`func (o *SponsorRecruitmentDurationRow) GetRealizedMonthsMedian() float32`

GetRealizedMonthsMedian returns the RealizedMonthsMedian field if non-nil, zero value otherwise.

### GetRealizedMonthsMedianOk

`func (o *SponsorRecruitmentDurationRow) GetRealizedMonthsMedianOk() (*float32, bool)`

GetRealizedMonthsMedianOk returns a tuple with the RealizedMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedMonthsMedian

`func (o *SponsorRecruitmentDurationRow) SetRealizedMonthsMedian(v float32)`

SetRealizedMonthsMedian sets RealizedMonthsMedian field to given value.

### HasRealizedMonthsMedian

`func (o *SponsorRecruitmentDurationRow) HasRealizedMonthsMedian() bool`

HasRealizedMonthsMedian returns a boolean if a field has been set.

### SetRealizedMonthsMedianNil

`func (o *SponsorRecruitmentDurationRow) SetRealizedMonthsMedianNil(b bool)`

 SetRealizedMonthsMedianNil sets the value for RealizedMonthsMedian to be an explicit nil

### UnsetRealizedMonthsMedian
`func (o *SponsorRecruitmentDurationRow) UnsetRealizedMonthsMedian()`

UnsetRealizedMonthsMedian ensures that no value is present for RealizedMonthsMedian, not even an explicit nil
### GetRealizedMonthsP25

`func (o *SponsorRecruitmentDurationRow) GetRealizedMonthsP25() float32`

GetRealizedMonthsP25 returns the RealizedMonthsP25 field if non-nil, zero value otherwise.

### GetRealizedMonthsP25Ok

`func (o *SponsorRecruitmentDurationRow) GetRealizedMonthsP25Ok() (*float32, bool)`

GetRealizedMonthsP25Ok returns a tuple with the RealizedMonthsP25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedMonthsP25

`func (o *SponsorRecruitmentDurationRow) SetRealizedMonthsP25(v float32)`

SetRealizedMonthsP25 sets RealizedMonthsP25 field to given value.

### HasRealizedMonthsP25

`func (o *SponsorRecruitmentDurationRow) HasRealizedMonthsP25() bool`

HasRealizedMonthsP25 returns a boolean if a field has been set.

### SetRealizedMonthsP25Nil

`func (o *SponsorRecruitmentDurationRow) SetRealizedMonthsP25Nil(b bool)`

 SetRealizedMonthsP25Nil sets the value for RealizedMonthsP25 to be an explicit nil

### UnsetRealizedMonthsP25
`func (o *SponsorRecruitmentDurationRow) UnsetRealizedMonthsP25()`

UnsetRealizedMonthsP25 ensures that no value is present for RealizedMonthsP25, not even an explicit nil
### GetRealizedMonthsP75

`func (o *SponsorRecruitmentDurationRow) GetRealizedMonthsP75() float32`

GetRealizedMonthsP75 returns the RealizedMonthsP75 field if non-nil, zero value otherwise.

### GetRealizedMonthsP75Ok

`func (o *SponsorRecruitmentDurationRow) GetRealizedMonthsP75Ok() (*float32, bool)`

GetRealizedMonthsP75Ok returns a tuple with the RealizedMonthsP75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedMonthsP75

`func (o *SponsorRecruitmentDurationRow) SetRealizedMonthsP75(v float32)`

SetRealizedMonthsP75 sets RealizedMonthsP75 field to given value.

### HasRealizedMonthsP75

`func (o *SponsorRecruitmentDurationRow) HasRealizedMonthsP75() bool`

HasRealizedMonthsP75 returns a boolean if a field has been set.

### SetRealizedMonthsP75Nil

`func (o *SponsorRecruitmentDurationRow) SetRealizedMonthsP75Nil(b bool)`

 SetRealizedMonthsP75Nil sets the value for RealizedMonthsP75 to be an explicit nil

### UnsetRealizedMonthsP75
`func (o *SponsorRecruitmentDurationRow) UnsetRealizedMonthsP75()`

UnsetRealizedMonthsP75 ensures that no value is present for RealizedMonthsP75, not even an explicit nil
### GetNRealizedOverCap

`func (o *SponsorRecruitmentDurationRow) GetNRealizedOverCap() int32`

GetNRealizedOverCap returns the NRealizedOverCap field if non-nil, zero value otherwise.

### GetNRealizedOverCapOk

`func (o *SponsorRecruitmentDurationRow) GetNRealizedOverCapOk() (*int32, bool)`

GetNRealizedOverCapOk returns a tuple with the NRealizedOverCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealizedOverCap

`func (o *SponsorRecruitmentDurationRow) SetNRealizedOverCap(v int32)`

SetNRealizedOverCap sets NRealizedOverCap field to given value.

### HasNRealizedOverCap

`func (o *SponsorRecruitmentDurationRow) HasNRealizedOverCap() bool`

HasNRealizedOverCap returns a boolean if a field has been set.

### SetNRealizedOverCapNil

`func (o *SponsorRecruitmentDurationRow) SetNRealizedOverCapNil(b bool)`

 SetNRealizedOverCapNil sets the value for NRealizedOverCap to be an explicit nil

### UnsetNRealizedOverCap
`func (o *SponsorRecruitmentDurationRow) UnsetNRealizedOverCap()`

UnsetNRealizedOverCap ensures that no value is present for NRealizedOverCap, not even an explicit nil
### GetMinN

`func (o *SponsorRecruitmentDurationRow) GetMinN() int32`

GetMinN returns the MinN field if non-nil, zero value otherwise.

### GetMinNOk

`func (o *SponsorRecruitmentDurationRow) GetMinNOk() (*int32, bool)`

GetMinNOk returns a tuple with the MinN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinN

`func (o *SponsorRecruitmentDurationRow) SetMinN(v int32)`

SetMinN sets MinN field to given value.

### HasMinN

`func (o *SponsorRecruitmentDurationRow) HasMinN() bool`

HasMinN returns a boolean if a field has been set.

### SetMinNNil

`func (o *SponsorRecruitmentDurationRow) SetMinNNil(b bool)`

 SetMinNNil sets the value for MinN to be an explicit nil

### UnsetMinN
`func (o *SponsorRecruitmentDurationRow) UnsetMinN()`

UnsetMinN ensures that no value is present for MinN, not even an explicit nil
### GetRollupStatus

`func (o *SponsorRecruitmentDurationRow) GetRollupStatus() string`

GetRollupStatus returns the RollupStatus field if non-nil, zero value otherwise.

### GetRollupStatusOk

`func (o *SponsorRecruitmentDurationRow) GetRollupStatusOk() (*string, bool)`

GetRollupStatusOk returns a tuple with the RollupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupStatus

`func (o *SponsorRecruitmentDurationRow) SetRollupStatus(v string)`

SetRollupStatus sets RollupStatus field to given value.

### HasRollupStatus

`func (o *SponsorRecruitmentDurationRow) HasRollupStatus() bool`

HasRollupStatus returns a boolean if a field has been set.

### SetRollupStatusNil

`func (o *SponsorRecruitmentDurationRow) SetRollupStatusNil(b bool)`

 SetRollupStatusNil sets the value for RollupStatus to be an explicit nil

### UnsetRollupStatus
`func (o *SponsorRecruitmentDurationRow) UnsetRollupStatus()`

UnsetRollupStatus ensures that no value is present for RollupStatus, not even an explicit nil
### GetRealizedStatus

`func (o *SponsorRecruitmentDurationRow) GetRealizedStatus() string`

GetRealizedStatus returns the RealizedStatus field if non-nil, zero value otherwise.

### GetRealizedStatusOk

`func (o *SponsorRecruitmentDurationRow) GetRealizedStatusOk() (*string, bool)`

GetRealizedStatusOk returns a tuple with the RealizedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedStatus

`func (o *SponsorRecruitmentDurationRow) SetRealizedStatus(v string)`

SetRealizedStatus sets RealizedStatus field to given value.

### HasRealizedStatus

`func (o *SponsorRecruitmentDurationRow) HasRealizedStatus() bool`

HasRealizedStatus returns a boolean if a field has been set.

### SetRealizedStatusNil

`func (o *SponsorRecruitmentDurationRow) SetRealizedStatusNil(b bool)`

 SetRealizedStatusNil sets the value for RealizedStatus to be an explicit nil

### UnsetRealizedStatus
`func (o *SponsorRecruitmentDurationRow) UnsetRealizedStatus()`

UnsetRealizedStatus ensures that no value is present for RealizedStatus, not even an explicit nil
### GetTrainingCutoff

`func (o *SponsorRecruitmentDurationRow) GetTrainingCutoff() string`

GetTrainingCutoff returns the TrainingCutoff field if non-nil, zero value otherwise.

### GetTrainingCutoffOk

`func (o *SponsorRecruitmentDurationRow) GetTrainingCutoffOk() (*string, bool)`

GetTrainingCutoffOk returns a tuple with the TrainingCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingCutoff

`func (o *SponsorRecruitmentDurationRow) SetTrainingCutoff(v string)`

SetTrainingCutoff sets TrainingCutoff field to given value.

### HasTrainingCutoff

`func (o *SponsorRecruitmentDurationRow) HasTrainingCutoff() bool`

HasTrainingCutoff returns a boolean if a field has been set.

### SetTrainingCutoffNil

`func (o *SponsorRecruitmentDurationRow) SetTrainingCutoffNil(b bool)`

 SetTrainingCutoffNil sets the value for TrainingCutoff to be an explicit nil

### UnsetTrainingCutoff
`func (o *SponsorRecruitmentDurationRow) UnsetTrainingCutoff()`

UnsetTrainingCutoff ensures that no value is present for TrainingCutoff, not even an explicit nil
### GetModelName

`func (o *SponsorRecruitmentDurationRow) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *SponsorRecruitmentDurationRow) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *SponsorRecruitmentDurationRow) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *SponsorRecruitmentDurationRow) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### SetModelNameNil

`func (o *SponsorRecruitmentDurationRow) SetModelNameNil(b bool)`

 SetModelNameNil sets the value for ModelName to be an explicit nil

### UnsetModelName
`func (o *SponsorRecruitmentDurationRow) UnsetModelName()`

UnsetModelName ensures that no value is present for ModelName, not even an explicit nil
### GetModelVersion

`func (o *SponsorRecruitmentDurationRow) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SponsorRecruitmentDurationRow) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SponsorRecruitmentDurationRow) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SponsorRecruitmentDurationRow) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### SetModelVersionNil

`func (o *SponsorRecruitmentDurationRow) SetModelVersionNil(b bool)`

 SetModelVersionNil sets the value for ModelVersion to be an explicit nil

### UnsetModelVersion
`func (o *SponsorRecruitmentDurationRow) UnsetModelVersion()`

UnsetModelVersion ensures that no value is present for ModelVersion, not even an explicit nil
### GetHorizonCapMonths

`func (o *SponsorRecruitmentDurationRow) GetHorizonCapMonths() int32`

GetHorizonCapMonths returns the HorizonCapMonths field if non-nil, zero value otherwise.

### GetHorizonCapMonthsOk

`func (o *SponsorRecruitmentDurationRow) GetHorizonCapMonthsOk() (*int32, bool)`

GetHorizonCapMonthsOk returns a tuple with the HorizonCapMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizonCapMonths

`func (o *SponsorRecruitmentDurationRow) SetHorizonCapMonths(v int32)`

SetHorizonCapMonths sets HorizonCapMonths field to given value.

### HasHorizonCapMonths

`func (o *SponsorRecruitmentDurationRow) HasHorizonCapMonths() bool`

HasHorizonCapMonths returns a boolean if a field has been set.

### SetHorizonCapMonthsNil

`func (o *SponsorRecruitmentDurationRow) SetHorizonCapMonthsNil(b bool)`

 SetHorizonCapMonthsNil sets the value for HorizonCapMonths to be an explicit nil

### UnsetHorizonCapMonths
`func (o *SponsorRecruitmentDurationRow) UnsetHorizonCapMonths()`

UnsetHorizonCapMonths ensures that no value is present for HorizonCapMonths, not even an explicit nil
### GetMethod

`func (o *SponsorRecruitmentDurationRow) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SponsorRecruitmentDurationRow) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SponsorRecruitmentDurationRow) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SponsorRecruitmentDurationRow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *SponsorRecruitmentDurationRow) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *SponsorRecruitmentDurationRow) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


