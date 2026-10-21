# ConditionTrialStartForecastRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s MeSH heading. | [optional] 
**TargetYear** | Pointer to **NullableInt32** | The calendar year forecast. | [optional] 
**HorizonOffset** | Pointer to **NullableInt32** | Years ahead of the build year (0 is the current year). | [optional] 
**MonthsObserved** | Pointer to **NullableInt32** | Months of the target year already observed. | [optional] 
**MonthsForecast** | Pointer to **NullableInt32** | Months of the target year forecast. | [optional] 
**InputObservedThroughCutoff** | Pointer to **NullableInt32** | Trials started in the target year up to the training cutoff. | [optional] 
**InputPlannedRemaining** | Pointer to **NullableInt32** | Trials registered as planned to start in the rest of the target year. | [optional] 
**ForecastPoint** | Pointer to **NullableFloat32** | The forecast trial starts for the whole target year. A forecast of the eventual count, never an actual count, and it carries no interval. | [optional] 
**ForecastStatus** | Pointer to **NullableString** | The series&#39; forecast status. | [optional] 
**InputTrialsWindow60m** | Pointer to **NullableInt32** | Trials of the condition (as in condition_trials) with a start date in the 60 calendar months ending with the cutoff month. | [optional] 
**InputStartsPerYear60m** | Pointer to **NullableFloat32** | Average trial starts per year over the last 60 months. | [optional] 
**SizeBucket** | Pointer to **NullableString** | The series size bucket the backtest error is measured in. | [optional] 
**BacktestHorizon** | Pointer to **NullableString** | The horizon the backtest error was measured at. | [optional] 
**BacktestSmapePct** | Pointer to **NullableFloat32** | Backtest symmetric mean absolute percentage error. | [optional] 
**BacktestRelBias** | Pointer to **NullableFloat32** | Backtest relative bias. | [optional] 
**BacktestRelmaeVsSnaive** | Pointer to **NullableFloat32** | Backtest mean absolute error relative to the seasonal naive baseline. | [optional] 
**BacktestNSeries** | Pointer to **NullableInt32** | Series the backtest covered. | [optional] 
**BacktestSeriesKind** | Pointer to **NullableString** | The kind of series the backtest error was measured on. | [optional] 
**BacktestRef** | Pointer to **NullableString** | The backtest run reference. | [optional] 
**PlannedStartCoefficient** | Pointer to **NullableFloat32** | The multiple of input_planned_remaining the planned-start method adds to the months after the training cutoff; null on a row built with the flat mean alone. | [optional] 
**PlannedStartStricterCheckAgrees** | Pointer to **NullableBool** | On a planned-start row, whether the stricter re-check (trained only on target years at least two before the cutoff year) adopted the same cell; null on a flat-mean row. | [optional] 
**PlannedStartStricterCheckRelBias** | Pointer to **NullableFloat32** | On a planned-start row, the relative bias (forecast minus actual over actual) of the stricter re-check; null on a flat-mean row. | [optional] 
**ModelName** | Pointer to **NullableString** | The forecasting model. | [optional] 
**ModelVersion** | Pointer to **NullableString** | The model version. | [optional] 
**TrainingCutoff** | Pointer to **NullableString** | The date the model&#39;s training data ends. | [optional] 
**Method** | Pointer to **NullableString** | How the forecast was produced. | [optional] 

## Methods

### NewConditionTrialStartForecastRow

`func NewConditionTrialStartForecastRow() *ConditionTrialStartForecastRow`

NewConditionTrialStartForecastRow instantiates a new ConditionTrialStartForecastRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionTrialStartForecastRowWithDefaults

`func NewConditionTrialStartForecastRowWithDefaults() *ConditionTrialStartForecastRow`

NewConditionTrialStartForecastRowWithDefaults instantiates a new ConditionTrialStartForecastRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionTrialStartForecastRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionTrialStartForecastRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionTrialStartForecastRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionTrialStartForecastRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionTrialStartForecastRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionTrialStartForecastRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *ConditionTrialStartForecastRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ConditionTrialStartForecastRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ConditionTrialStartForecastRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *ConditionTrialStartForecastRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *ConditionTrialStartForecastRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *ConditionTrialStartForecastRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetTargetYear

`func (o *ConditionTrialStartForecastRow) GetTargetYear() int32`

GetTargetYear returns the TargetYear field if non-nil, zero value otherwise.

### GetTargetYearOk

`func (o *ConditionTrialStartForecastRow) GetTargetYearOk() (*int32, bool)`

GetTargetYearOk returns a tuple with the TargetYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetYear

`func (o *ConditionTrialStartForecastRow) SetTargetYear(v int32)`

SetTargetYear sets TargetYear field to given value.

### HasTargetYear

`func (o *ConditionTrialStartForecastRow) HasTargetYear() bool`

HasTargetYear returns a boolean if a field has been set.

### SetTargetYearNil

`func (o *ConditionTrialStartForecastRow) SetTargetYearNil(b bool)`

 SetTargetYearNil sets the value for TargetYear to be an explicit nil

### UnsetTargetYear
`func (o *ConditionTrialStartForecastRow) UnsetTargetYear()`

UnsetTargetYear ensures that no value is present for TargetYear, not even an explicit nil
### GetHorizonOffset

`func (o *ConditionTrialStartForecastRow) GetHorizonOffset() int32`

GetHorizonOffset returns the HorizonOffset field if non-nil, zero value otherwise.

### GetHorizonOffsetOk

`func (o *ConditionTrialStartForecastRow) GetHorizonOffsetOk() (*int32, bool)`

GetHorizonOffsetOk returns a tuple with the HorizonOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizonOffset

`func (o *ConditionTrialStartForecastRow) SetHorizonOffset(v int32)`

SetHorizonOffset sets HorizonOffset field to given value.

### HasHorizonOffset

`func (o *ConditionTrialStartForecastRow) HasHorizonOffset() bool`

HasHorizonOffset returns a boolean if a field has been set.

### SetHorizonOffsetNil

`func (o *ConditionTrialStartForecastRow) SetHorizonOffsetNil(b bool)`

 SetHorizonOffsetNil sets the value for HorizonOffset to be an explicit nil

### UnsetHorizonOffset
`func (o *ConditionTrialStartForecastRow) UnsetHorizonOffset()`

UnsetHorizonOffset ensures that no value is present for HorizonOffset, not even an explicit nil
### GetMonthsObserved

`func (o *ConditionTrialStartForecastRow) GetMonthsObserved() int32`

GetMonthsObserved returns the MonthsObserved field if non-nil, zero value otherwise.

### GetMonthsObservedOk

`func (o *ConditionTrialStartForecastRow) GetMonthsObservedOk() (*int32, bool)`

GetMonthsObservedOk returns a tuple with the MonthsObserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthsObserved

`func (o *ConditionTrialStartForecastRow) SetMonthsObserved(v int32)`

SetMonthsObserved sets MonthsObserved field to given value.

### HasMonthsObserved

`func (o *ConditionTrialStartForecastRow) HasMonthsObserved() bool`

HasMonthsObserved returns a boolean if a field has been set.

### SetMonthsObservedNil

`func (o *ConditionTrialStartForecastRow) SetMonthsObservedNil(b bool)`

 SetMonthsObservedNil sets the value for MonthsObserved to be an explicit nil

### UnsetMonthsObserved
`func (o *ConditionTrialStartForecastRow) UnsetMonthsObserved()`

UnsetMonthsObserved ensures that no value is present for MonthsObserved, not even an explicit nil
### GetMonthsForecast

`func (o *ConditionTrialStartForecastRow) GetMonthsForecast() int32`

GetMonthsForecast returns the MonthsForecast field if non-nil, zero value otherwise.

### GetMonthsForecastOk

`func (o *ConditionTrialStartForecastRow) GetMonthsForecastOk() (*int32, bool)`

GetMonthsForecastOk returns a tuple with the MonthsForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthsForecast

`func (o *ConditionTrialStartForecastRow) SetMonthsForecast(v int32)`

SetMonthsForecast sets MonthsForecast field to given value.

### HasMonthsForecast

`func (o *ConditionTrialStartForecastRow) HasMonthsForecast() bool`

HasMonthsForecast returns a boolean if a field has been set.

### SetMonthsForecastNil

`func (o *ConditionTrialStartForecastRow) SetMonthsForecastNil(b bool)`

 SetMonthsForecastNil sets the value for MonthsForecast to be an explicit nil

### UnsetMonthsForecast
`func (o *ConditionTrialStartForecastRow) UnsetMonthsForecast()`

UnsetMonthsForecast ensures that no value is present for MonthsForecast, not even an explicit nil
### GetInputObservedThroughCutoff

`func (o *ConditionTrialStartForecastRow) GetInputObservedThroughCutoff() int32`

GetInputObservedThroughCutoff returns the InputObservedThroughCutoff field if non-nil, zero value otherwise.

### GetInputObservedThroughCutoffOk

`func (o *ConditionTrialStartForecastRow) GetInputObservedThroughCutoffOk() (*int32, bool)`

GetInputObservedThroughCutoffOk returns a tuple with the InputObservedThroughCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputObservedThroughCutoff

`func (o *ConditionTrialStartForecastRow) SetInputObservedThroughCutoff(v int32)`

SetInputObservedThroughCutoff sets InputObservedThroughCutoff field to given value.

### HasInputObservedThroughCutoff

`func (o *ConditionTrialStartForecastRow) HasInputObservedThroughCutoff() bool`

HasInputObservedThroughCutoff returns a boolean if a field has been set.

### SetInputObservedThroughCutoffNil

`func (o *ConditionTrialStartForecastRow) SetInputObservedThroughCutoffNil(b bool)`

 SetInputObservedThroughCutoffNil sets the value for InputObservedThroughCutoff to be an explicit nil

### UnsetInputObservedThroughCutoff
`func (o *ConditionTrialStartForecastRow) UnsetInputObservedThroughCutoff()`

UnsetInputObservedThroughCutoff ensures that no value is present for InputObservedThroughCutoff, not even an explicit nil
### GetInputPlannedRemaining

`func (o *ConditionTrialStartForecastRow) GetInputPlannedRemaining() int32`

GetInputPlannedRemaining returns the InputPlannedRemaining field if non-nil, zero value otherwise.

### GetInputPlannedRemainingOk

`func (o *ConditionTrialStartForecastRow) GetInputPlannedRemainingOk() (*int32, bool)`

GetInputPlannedRemainingOk returns a tuple with the InputPlannedRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPlannedRemaining

`func (o *ConditionTrialStartForecastRow) SetInputPlannedRemaining(v int32)`

SetInputPlannedRemaining sets InputPlannedRemaining field to given value.

### HasInputPlannedRemaining

`func (o *ConditionTrialStartForecastRow) HasInputPlannedRemaining() bool`

HasInputPlannedRemaining returns a boolean if a field has been set.

### SetInputPlannedRemainingNil

`func (o *ConditionTrialStartForecastRow) SetInputPlannedRemainingNil(b bool)`

 SetInputPlannedRemainingNil sets the value for InputPlannedRemaining to be an explicit nil

### UnsetInputPlannedRemaining
`func (o *ConditionTrialStartForecastRow) UnsetInputPlannedRemaining()`

UnsetInputPlannedRemaining ensures that no value is present for InputPlannedRemaining, not even an explicit nil
### GetForecastPoint

`func (o *ConditionTrialStartForecastRow) GetForecastPoint() float32`

GetForecastPoint returns the ForecastPoint field if non-nil, zero value otherwise.

### GetForecastPointOk

`func (o *ConditionTrialStartForecastRow) GetForecastPointOk() (*float32, bool)`

GetForecastPointOk returns a tuple with the ForecastPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastPoint

`func (o *ConditionTrialStartForecastRow) SetForecastPoint(v float32)`

SetForecastPoint sets ForecastPoint field to given value.

### HasForecastPoint

`func (o *ConditionTrialStartForecastRow) HasForecastPoint() bool`

HasForecastPoint returns a boolean if a field has been set.

### SetForecastPointNil

`func (o *ConditionTrialStartForecastRow) SetForecastPointNil(b bool)`

 SetForecastPointNil sets the value for ForecastPoint to be an explicit nil

### UnsetForecastPoint
`func (o *ConditionTrialStartForecastRow) UnsetForecastPoint()`

UnsetForecastPoint ensures that no value is present for ForecastPoint, not even an explicit nil
### GetForecastStatus

`func (o *ConditionTrialStartForecastRow) GetForecastStatus() string`

GetForecastStatus returns the ForecastStatus field if non-nil, zero value otherwise.

### GetForecastStatusOk

`func (o *ConditionTrialStartForecastRow) GetForecastStatusOk() (*string, bool)`

GetForecastStatusOk returns a tuple with the ForecastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastStatus

`func (o *ConditionTrialStartForecastRow) SetForecastStatus(v string)`

SetForecastStatus sets ForecastStatus field to given value.

### HasForecastStatus

`func (o *ConditionTrialStartForecastRow) HasForecastStatus() bool`

HasForecastStatus returns a boolean if a field has been set.

### SetForecastStatusNil

`func (o *ConditionTrialStartForecastRow) SetForecastStatusNil(b bool)`

 SetForecastStatusNil sets the value for ForecastStatus to be an explicit nil

### UnsetForecastStatus
`func (o *ConditionTrialStartForecastRow) UnsetForecastStatus()`

UnsetForecastStatus ensures that no value is present for ForecastStatus, not even an explicit nil
### GetInputTrialsWindow60m

`func (o *ConditionTrialStartForecastRow) GetInputTrialsWindow60m() int32`

GetInputTrialsWindow60m returns the InputTrialsWindow60m field if non-nil, zero value otherwise.

### GetInputTrialsWindow60mOk

`func (o *ConditionTrialStartForecastRow) GetInputTrialsWindow60mOk() (*int32, bool)`

GetInputTrialsWindow60mOk returns a tuple with the InputTrialsWindow60m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTrialsWindow60m

`func (o *ConditionTrialStartForecastRow) SetInputTrialsWindow60m(v int32)`

SetInputTrialsWindow60m sets InputTrialsWindow60m field to given value.

### HasInputTrialsWindow60m

`func (o *ConditionTrialStartForecastRow) HasInputTrialsWindow60m() bool`

HasInputTrialsWindow60m returns a boolean if a field has been set.

### SetInputTrialsWindow60mNil

`func (o *ConditionTrialStartForecastRow) SetInputTrialsWindow60mNil(b bool)`

 SetInputTrialsWindow60mNil sets the value for InputTrialsWindow60m to be an explicit nil

### UnsetInputTrialsWindow60m
`func (o *ConditionTrialStartForecastRow) UnsetInputTrialsWindow60m()`

UnsetInputTrialsWindow60m ensures that no value is present for InputTrialsWindow60m, not even an explicit nil
### GetInputStartsPerYear60m

`func (o *ConditionTrialStartForecastRow) GetInputStartsPerYear60m() float32`

GetInputStartsPerYear60m returns the InputStartsPerYear60m field if non-nil, zero value otherwise.

### GetInputStartsPerYear60mOk

`func (o *ConditionTrialStartForecastRow) GetInputStartsPerYear60mOk() (*float32, bool)`

GetInputStartsPerYear60mOk returns a tuple with the InputStartsPerYear60m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputStartsPerYear60m

`func (o *ConditionTrialStartForecastRow) SetInputStartsPerYear60m(v float32)`

SetInputStartsPerYear60m sets InputStartsPerYear60m field to given value.

### HasInputStartsPerYear60m

`func (o *ConditionTrialStartForecastRow) HasInputStartsPerYear60m() bool`

HasInputStartsPerYear60m returns a boolean if a field has been set.

### SetInputStartsPerYear60mNil

`func (o *ConditionTrialStartForecastRow) SetInputStartsPerYear60mNil(b bool)`

 SetInputStartsPerYear60mNil sets the value for InputStartsPerYear60m to be an explicit nil

### UnsetInputStartsPerYear60m
`func (o *ConditionTrialStartForecastRow) UnsetInputStartsPerYear60m()`

UnsetInputStartsPerYear60m ensures that no value is present for InputStartsPerYear60m, not even an explicit nil
### GetSizeBucket

`func (o *ConditionTrialStartForecastRow) GetSizeBucket() string`

GetSizeBucket returns the SizeBucket field if non-nil, zero value otherwise.

### GetSizeBucketOk

`func (o *ConditionTrialStartForecastRow) GetSizeBucketOk() (*string, bool)`

GetSizeBucketOk returns a tuple with the SizeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBucket

`func (o *ConditionTrialStartForecastRow) SetSizeBucket(v string)`

SetSizeBucket sets SizeBucket field to given value.

### HasSizeBucket

`func (o *ConditionTrialStartForecastRow) HasSizeBucket() bool`

HasSizeBucket returns a boolean if a field has been set.

### SetSizeBucketNil

`func (o *ConditionTrialStartForecastRow) SetSizeBucketNil(b bool)`

 SetSizeBucketNil sets the value for SizeBucket to be an explicit nil

### UnsetSizeBucket
`func (o *ConditionTrialStartForecastRow) UnsetSizeBucket()`

UnsetSizeBucket ensures that no value is present for SizeBucket, not even an explicit nil
### GetBacktestHorizon

`func (o *ConditionTrialStartForecastRow) GetBacktestHorizon() string`

GetBacktestHorizon returns the BacktestHorizon field if non-nil, zero value otherwise.

### GetBacktestHorizonOk

`func (o *ConditionTrialStartForecastRow) GetBacktestHorizonOk() (*string, bool)`

GetBacktestHorizonOk returns a tuple with the BacktestHorizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestHorizon

`func (o *ConditionTrialStartForecastRow) SetBacktestHorizon(v string)`

SetBacktestHorizon sets BacktestHorizon field to given value.

### HasBacktestHorizon

`func (o *ConditionTrialStartForecastRow) HasBacktestHorizon() bool`

HasBacktestHorizon returns a boolean if a field has been set.

### SetBacktestHorizonNil

`func (o *ConditionTrialStartForecastRow) SetBacktestHorizonNil(b bool)`

 SetBacktestHorizonNil sets the value for BacktestHorizon to be an explicit nil

### UnsetBacktestHorizon
`func (o *ConditionTrialStartForecastRow) UnsetBacktestHorizon()`

UnsetBacktestHorizon ensures that no value is present for BacktestHorizon, not even an explicit nil
### GetBacktestSmapePct

`func (o *ConditionTrialStartForecastRow) GetBacktestSmapePct() float32`

GetBacktestSmapePct returns the BacktestSmapePct field if non-nil, zero value otherwise.

### GetBacktestSmapePctOk

`func (o *ConditionTrialStartForecastRow) GetBacktestSmapePctOk() (*float32, bool)`

GetBacktestSmapePctOk returns a tuple with the BacktestSmapePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestSmapePct

`func (o *ConditionTrialStartForecastRow) SetBacktestSmapePct(v float32)`

SetBacktestSmapePct sets BacktestSmapePct field to given value.

### HasBacktestSmapePct

`func (o *ConditionTrialStartForecastRow) HasBacktestSmapePct() bool`

HasBacktestSmapePct returns a boolean if a field has been set.

### SetBacktestSmapePctNil

`func (o *ConditionTrialStartForecastRow) SetBacktestSmapePctNil(b bool)`

 SetBacktestSmapePctNil sets the value for BacktestSmapePct to be an explicit nil

### UnsetBacktestSmapePct
`func (o *ConditionTrialStartForecastRow) UnsetBacktestSmapePct()`

UnsetBacktestSmapePct ensures that no value is present for BacktestSmapePct, not even an explicit nil
### GetBacktestRelBias

`func (o *ConditionTrialStartForecastRow) GetBacktestRelBias() float32`

GetBacktestRelBias returns the BacktestRelBias field if non-nil, zero value otherwise.

### GetBacktestRelBiasOk

`func (o *ConditionTrialStartForecastRow) GetBacktestRelBiasOk() (*float32, bool)`

GetBacktestRelBiasOk returns a tuple with the BacktestRelBias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestRelBias

`func (o *ConditionTrialStartForecastRow) SetBacktestRelBias(v float32)`

SetBacktestRelBias sets BacktestRelBias field to given value.

### HasBacktestRelBias

`func (o *ConditionTrialStartForecastRow) HasBacktestRelBias() bool`

HasBacktestRelBias returns a boolean if a field has been set.

### SetBacktestRelBiasNil

`func (o *ConditionTrialStartForecastRow) SetBacktestRelBiasNil(b bool)`

 SetBacktestRelBiasNil sets the value for BacktestRelBias to be an explicit nil

### UnsetBacktestRelBias
`func (o *ConditionTrialStartForecastRow) UnsetBacktestRelBias()`

UnsetBacktestRelBias ensures that no value is present for BacktestRelBias, not even an explicit nil
### GetBacktestRelmaeVsSnaive

`func (o *ConditionTrialStartForecastRow) GetBacktestRelmaeVsSnaive() float32`

GetBacktestRelmaeVsSnaive returns the BacktestRelmaeVsSnaive field if non-nil, zero value otherwise.

### GetBacktestRelmaeVsSnaiveOk

`func (o *ConditionTrialStartForecastRow) GetBacktestRelmaeVsSnaiveOk() (*float32, bool)`

GetBacktestRelmaeVsSnaiveOk returns a tuple with the BacktestRelmaeVsSnaive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestRelmaeVsSnaive

`func (o *ConditionTrialStartForecastRow) SetBacktestRelmaeVsSnaive(v float32)`

SetBacktestRelmaeVsSnaive sets BacktestRelmaeVsSnaive field to given value.

### HasBacktestRelmaeVsSnaive

`func (o *ConditionTrialStartForecastRow) HasBacktestRelmaeVsSnaive() bool`

HasBacktestRelmaeVsSnaive returns a boolean if a field has been set.

### SetBacktestRelmaeVsSnaiveNil

`func (o *ConditionTrialStartForecastRow) SetBacktestRelmaeVsSnaiveNil(b bool)`

 SetBacktestRelmaeVsSnaiveNil sets the value for BacktestRelmaeVsSnaive to be an explicit nil

### UnsetBacktestRelmaeVsSnaive
`func (o *ConditionTrialStartForecastRow) UnsetBacktestRelmaeVsSnaive()`

UnsetBacktestRelmaeVsSnaive ensures that no value is present for BacktestRelmaeVsSnaive, not even an explicit nil
### GetBacktestNSeries

`func (o *ConditionTrialStartForecastRow) GetBacktestNSeries() int32`

GetBacktestNSeries returns the BacktestNSeries field if non-nil, zero value otherwise.

### GetBacktestNSeriesOk

`func (o *ConditionTrialStartForecastRow) GetBacktestNSeriesOk() (*int32, bool)`

GetBacktestNSeriesOk returns a tuple with the BacktestNSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestNSeries

`func (o *ConditionTrialStartForecastRow) SetBacktestNSeries(v int32)`

SetBacktestNSeries sets BacktestNSeries field to given value.

### HasBacktestNSeries

`func (o *ConditionTrialStartForecastRow) HasBacktestNSeries() bool`

HasBacktestNSeries returns a boolean if a field has been set.

### SetBacktestNSeriesNil

`func (o *ConditionTrialStartForecastRow) SetBacktestNSeriesNil(b bool)`

 SetBacktestNSeriesNil sets the value for BacktestNSeries to be an explicit nil

### UnsetBacktestNSeries
`func (o *ConditionTrialStartForecastRow) UnsetBacktestNSeries()`

UnsetBacktestNSeries ensures that no value is present for BacktestNSeries, not even an explicit nil
### GetBacktestSeriesKind

`func (o *ConditionTrialStartForecastRow) GetBacktestSeriesKind() string`

GetBacktestSeriesKind returns the BacktestSeriesKind field if non-nil, zero value otherwise.

### GetBacktestSeriesKindOk

`func (o *ConditionTrialStartForecastRow) GetBacktestSeriesKindOk() (*string, bool)`

GetBacktestSeriesKindOk returns a tuple with the BacktestSeriesKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestSeriesKind

`func (o *ConditionTrialStartForecastRow) SetBacktestSeriesKind(v string)`

SetBacktestSeriesKind sets BacktestSeriesKind field to given value.

### HasBacktestSeriesKind

`func (o *ConditionTrialStartForecastRow) HasBacktestSeriesKind() bool`

HasBacktestSeriesKind returns a boolean if a field has been set.

### SetBacktestSeriesKindNil

`func (o *ConditionTrialStartForecastRow) SetBacktestSeriesKindNil(b bool)`

 SetBacktestSeriesKindNil sets the value for BacktestSeriesKind to be an explicit nil

### UnsetBacktestSeriesKind
`func (o *ConditionTrialStartForecastRow) UnsetBacktestSeriesKind()`

UnsetBacktestSeriesKind ensures that no value is present for BacktestSeriesKind, not even an explicit nil
### GetBacktestRef

`func (o *ConditionTrialStartForecastRow) GetBacktestRef() string`

GetBacktestRef returns the BacktestRef field if non-nil, zero value otherwise.

### GetBacktestRefOk

`func (o *ConditionTrialStartForecastRow) GetBacktestRefOk() (*string, bool)`

GetBacktestRefOk returns a tuple with the BacktestRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestRef

`func (o *ConditionTrialStartForecastRow) SetBacktestRef(v string)`

SetBacktestRef sets BacktestRef field to given value.

### HasBacktestRef

`func (o *ConditionTrialStartForecastRow) HasBacktestRef() bool`

HasBacktestRef returns a boolean if a field has been set.

### SetBacktestRefNil

`func (o *ConditionTrialStartForecastRow) SetBacktestRefNil(b bool)`

 SetBacktestRefNil sets the value for BacktestRef to be an explicit nil

### UnsetBacktestRef
`func (o *ConditionTrialStartForecastRow) UnsetBacktestRef()`

UnsetBacktestRef ensures that no value is present for BacktestRef, not even an explicit nil
### GetPlannedStartCoefficient

`func (o *ConditionTrialStartForecastRow) GetPlannedStartCoefficient() float32`

GetPlannedStartCoefficient returns the PlannedStartCoefficient field if non-nil, zero value otherwise.

### GetPlannedStartCoefficientOk

`func (o *ConditionTrialStartForecastRow) GetPlannedStartCoefficientOk() (*float32, bool)`

GetPlannedStartCoefficientOk returns a tuple with the PlannedStartCoefficient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartCoefficient

`func (o *ConditionTrialStartForecastRow) SetPlannedStartCoefficient(v float32)`

SetPlannedStartCoefficient sets PlannedStartCoefficient field to given value.

### HasPlannedStartCoefficient

`func (o *ConditionTrialStartForecastRow) HasPlannedStartCoefficient() bool`

HasPlannedStartCoefficient returns a boolean if a field has been set.

### SetPlannedStartCoefficientNil

`func (o *ConditionTrialStartForecastRow) SetPlannedStartCoefficientNil(b bool)`

 SetPlannedStartCoefficientNil sets the value for PlannedStartCoefficient to be an explicit nil

### UnsetPlannedStartCoefficient
`func (o *ConditionTrialStartForecastRow) UnsetPlannedStartCoefficient()`

UnsetPlannedStartCoefficient ensures that no value is present for PlannedStartCoefficient, not even an explicit nil
### GetPlannedStartStricterCheckAgrees

`func (o *ConditionTrialStartForecastRow) GetPlannedStartStricterCheckAgrees() bool`

GetPlannedStartStricterCheckAgrees returns the PlannedStartStricterCheckAgrees field if non-nil, zero value otherwise.

### GetPlannedStartStricterCheckAgreesOk

`func (o *ConditionTrialStartForecastRow) GetPlannedStartStricterCheckAgreesOk() (*bool, bool)`

GetPlannedStartStricterCheckAgreesOk returns a tuple with the PlannedStartStricterCheckAgrees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartStricterCheckAgrees

`func (o *ConditionTrialStartForecastRow) SetPlannedStartStricterCheckAgrees(v bool)`

SetPlannedStartStricterCheckAgrees sets PlannedStartStricterCheckAgrees field to given value.

### HasPlannedStartStricterCheckAgrees

`func (o *ConditionTrialStartForecastRow) HasPlannedStartStricterCheckAgrees() bool`

HasPlannedStartStricterCheckAgrees returns a boolean if a field has been set.

### SetPlannedStartStricterCheckAgreesNil

`func (o *ConditionTrialStartForecastRow) SetPlannedStartStricterCheckAgreesNil(b bool)`

 SetPlannedStartStricterCheckAgreesNil sets the value for PlannedStartStricterCheckAgrees to be an explicit nil

### UnsetPlannedStartStricterCheckAgrees
`func (o *ConditionTrialStartForecastRow) UnsetPlannedStartStricterCheckAgrees()`

UnsetPlannedStartStricterCheckAgrees ensures that no value is present for PlannedStartStricterCheckAgrees, not even an explicit nil
### GetPlannedStartStricterCheckRelBias

`func (o *ConditionTrialStartForecastRow) GetPlannedStartStricterCheckRelBias() float32`

GetPlannedStartStricterCheckRelBias returns the PlannedStartStricterCheckRelBias field if non-nil, zero value otherwise.

### GetPlannedStartStricterCheckRelBiasOk

`func (o *ConditionTrialStartForecastRow) GetPlannedStartStricterCheckRelBiasOk() (*float32, bool)`

GetPlannedStartStricterCheckRelBiasOk returns a tuple with the PlannedStartStricterCheckRelBias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartStricterCheckRelBias

`func (o *ConditionTrialStartForecastRow) SetPlannedStartStricterCheckRelBias(v float32)`

SetPlannedStartStricterCheckRelBias sets PlannedStartStricterCheckRelBias field to given value.

### HasPlannedStartStricterCheckRelBias

`func (o *ConditionTrialStartForecastRow) HasPlannedStartStricterCheckRelBias() bool`

HasPlannedStartStricterCheckRelBias returns a boolean if a field has been set.

### SetPlannedStartStricterCheckRelBiasNil

`func (o *ConditionTrialStartForecastRow) SetPlannedStartStricterCheckRelBiasNil(b bool)`

 SetPlannedStartStricterCheckRelBiasNil sets the value for PlannedStartStricterCheckRelBias to be an explicit nil

### UnsetPlannedStartStricterCheckRelBias
`func (o *ConditionTrialStartForecastRow) UnsetPlannedStartStricterCheckRelBias()`

UnsetPlannedStartStricterCheckRelBias ensures that no value is present for PlannedStartStricterCheckRelBias, not even an explicit nil
### GetModelName

`func (o *ConditionTrialStartForecastRow) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ConditionTrialStartForecastRow) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ConditionTrialStartForecastRow) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ConditionTrialStartForecastRow) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### SetModelNameNil

`func (o *ConditionTrialStartForecastRow) SetModelNameNil(b bool)`

 SetModelNameNil sets the value for ModelName to be an explicit nil

### UnsetModelName
`func (o *ConditionTrialStartForecastRow) UnsetModelName()`

UnsetModelName ensures that no value is present for ModelName, not even an explicit nil
### GetModelVersion

`func (o *ConditionTrialStartForecastRow) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ConditionTrialStartForecastRow) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ConditionTrialStartForecastRow) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ConditionTrialStartForecastRow) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### SetModelVersionNil

`func (o *ConditionTrialStartForecastRow) SetModelVersionNil(b bool)`

 SetModelVersionNil sets the value for ModelVersion to be an explicit nil

### UnsetModelVersion
`func (o *ConditionTrialStartForecastRow) UnsetModelVersion()`

UnsetModelVersion ensures that no value is present for ModelVersion, not even an explicit nil
### GetTrainingCutoff

`func (o *ConditionTrialStartForecastRow) GetTrainingCutoff() string`

GetTrainingCutoff returns the TrainingCutoff field if non-nil, zero value otherwise.

### GetTrainingCutoffOk

`func (o *ConditionTrialStartForecastRow) GetTrainingCutoffOk() (*string, bool)`

GetTrainingCutoffOk returns a tuple with the TrainingCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingCutoff

`func (o *ConditionTrialStartForecastRow) SetTrainingCutoff(v string)`

SetTrainingCutoff sets TrainingCutoff field to given value.

### HasTrainingCutoff

`func (o *ConditionTrialStartForecastRow) HasTrainingCutoff() bool`

HasTrainingCutoff returns a boolean if a field has been set.

### SetTrainingCutoffNil

`func (o *ConditionTrialStartForecastRow) SetTrainingCutoffNil(b bool)`

 SetTrainingCutoffNil sets the value for TrainingCutoff to be an explicit nil

### UnsetTrainingCutoff
`func (o *ConditionTrialStartForecastRow) UnsetTrainingCutoff()`

UnsetTrainingCutoff ensures that no value is present for TrainingCutoff, not even an explicit nil
### GetMethod

`func (o *ConditionTrialStartForecastRow) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ConditionTrialStartForecastRow) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ConditionTrialStartForecastRow) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ConditionTrialStartForecastRow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *ConditionTrialStartForecastRow) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *ConditionTrialStartForecastRow) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


