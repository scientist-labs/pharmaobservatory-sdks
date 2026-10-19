# SponsorTrialStartForecastRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the forecast series is for. | [optional] 
**SponsorName** | Pointer to **NullableString** | Display name of the organization record. | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company the organization belongs to (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**TargetYear** | Pointer to **NullableInt32** | The calendar year forecast. | [optional] 
**HorizonOffset** | Pointer to **NullableInt32** | Years ahead of the build year (0 is the current year). | [optional] 
**MonthsObserved** | Pointer to **NullableInt32** | Months of the target year already observed. | [optional] 
**MonthsForecast** | Pointer to **NullableInt32** | Months of the target year forecast. | [optional] 
**InputObservedThroughCutoff** | Pointer to **NullableInt32** | Trials started in the target year up to the training cutoff. | [optional] 
**InputPlannedRemaining** | Pointer to **NullableInt32** | Trials registered as planned to start in the rest of the target year. | [optional] 
**ForecastPoint** | Pointer to **NullableFloat32** | The forecast trial starts for the whole target year. A forecast of the eventual count, never an actual count, and it carries no interval. | [optional] 
**ForecastStatus** | Pointer to **NullableString** | The series&#39; forecast status. | [optional] 
**InputTrialsWindow60m** | Pointer to **NullableInt32** | Trials the organization started in the last 60 months. | [optional] 
**InputStartsPerYear60m** | Pointer to **NullableFloat32** | Average trial starts per year over the last 60 months. | [optional] 
**SizeBucket** | Pointer to **NullableString** | The series size bucket the backtest error is measured in. | [optional] 
**BacktestHorizon** | Pointer to **NullableString** | The horizon the backtest error was measured at. | [optional] 
**BacktestSmapePct** | Pointer to **NullableFloat32** | Backtest symmetric mean absolute percentage error. | [optional] 
**BacktestRelBias** | Pointer to **NullableFloat32** | Backtest relative bias. | [optional] 
**BacktestRelmaeVsSnaive** | Pointer to **NullableFloat32** | Backtest mean absolute error relative to the seasonal naive baseline. | [optional] 
**BacktestNSeries** | Pointer to **NullableInt32** | Series the backtest covered. | [optional] 
**BacktestRef** | Pointer to **NullableString** | The backtest run reference. | [optional] 
**ModelName** | Pointer to **NullableString** | The forecasting model. | [optional] 
**ModelVersion** | Pointer to **NullableString** | The model version. | [optional] 
**TrainingCutoff** | Pointer to **NullableString** | The date the model&#39;s training data ends. | [optional] 
**Method** | Pointer to **NullableString** | How the forecast was produced. | [optional] 

## Methods

### NewSponsorTrialStartForecastRow

`func NewSponsorTrialStartForecastRow() *SponsorTrialStartForecastRow`

NewSponsorTrialStartForecastRow instantiates a new SponsorTrialStartForecastRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSponsorTrialStartForecastRowWithDefaults

`func NewSponsorTrialStartForecastRowWithDefaults() *SponsorTrialStartForecastRow`

NewSponsorTrialStartForecastRowWithDefaults instantiates a new SponsorTrialStartForecastRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSponsorOrgId

`func (o *SponsorTrialStartForecastRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *SponsorTrialStartForecastRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *SponsorTrialStartForecastRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *SponsorTrialStartForecastRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *SponsorTrialStartForecastRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *SponsorTrialStartForecastRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorName

`func (o *SponsorTrialStartForecastRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *SponsorTrialStartForecastRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *SponsorTrialStartForecastRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *SponsorTrialStartForecastRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *SponsorTrialStartForecastRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *SponsorTrialStartForecastRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorCompanyId

`func (o *SponsorTrialStartForecastRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *SponsorTrialStartForecastRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *SponsorTrialStartForecastRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *SponsorTrialStartForecastRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *SponsorTrialStartForecastRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *SponsorTrialStartForecastRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *SponsorTrialStartForecastRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *SponsorTrialStartForecastRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *SponsorTrialStartForecastRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *SponsorTrialStartForecastRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *SponsorTrialStartForecastRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *SponsorTrialStartForecastRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *SponsorTrialStartForecastRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *SponsorTrialStartForecastRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *SponsorTrialStartForecastRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *SponsorTrialStartForecastRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *SponsorTrialStartForecastRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *SponsorTrialStartForecastRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *SponsorTrialStartForecastRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *SponsorTrialStartForecastRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *SponsorTrialStartForecastRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *SponsorTrialStartForecastRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *SponsorTrialStartForecastRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *SponsorTrialStartForecastRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetTargetYear

`func (o *SponsorTrialStartForecastRow) GetTargetYear() int32`

GetTargetYear returns the TargetYear field if non-nil, zero value otherwise.

### GetTargetYearOk

`func (o *SponsorTrialStartForecastRow) GetTargetYearOk() (*int32, bool)`

GetTargetYearOk returns a tuple with the TargetYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetYear

`func (o *SponsorTrialStartForecastRow) SetTargetYear(v int32)`

SetTargetYear sets TargetYear field to given value.

### HasTargetYear

`func (o *SponsorTrialStartForecastRow) HasTargetYear() bool`

HasTargetYear returns a boolean if a field has been set.

### SetTargetYearNil

`func (o *SponsorTrialStartForecastRow) SetTargetYearNil(b bool)`

 SetTargetYearNil sets the value for TargetYear to be an explicit nil

### UnsetTargetYear
`func (o *SponsorTrialStartForecastRow) UnsetTargetYear()`

UnsetTargetYear ensures that no value is present for TargetYear, not even an explicit nil
### GetHorizonOffset

`func (o *SponsorTrialStartForecastRow) GetHorizonOffset() int32`

GetHorizonOffset returns the HorizonOffset field if non-nil, zero value otherwise.

### GetHorizonOffsetOk

`func (o *SponsorTrialStartForecastRow) GetHorizonOffsetOk() (*int32, bool)`

GetHorizonOffsetOk returns a tuple with the HorizonOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizonOffset

`func (o *SponsorTrialStartForecastRow) SetHorizonOffset(v int32)`

SetHorizonOffset sets HorizonOffset field to given value.

### HasHorizonOffset

`func (o *SponsorTrialStartForecastRow) HasHorizonOffset() bool`

HasHorizonOffset returns a boolean if a field has been set.

### SetHorizonOffsetNil

`func (o *SponsorTrialStartForecastRow) SetHorizonOffsetNil(b bool)`

 SetHorizonOffsetNil sets the value for HorizonOffset to be an explicit nil

### UnsetHorizonOffset
`func (o *SponsorTrialStartForecastRow) UnsetHorizonOffset()`

UnsetHorizonOffset ensures that no value is present for HorizonOffset, not even an explicit nil
### GetMonthsObserved

`func (o *SponsorTrialStartForecastRow) GetMonthsObserved() int32`

GetMonthsObserved returns the MonthsObserved field if non-nil, zero value otherwise.

### GetMonthsObservedOk

`func (o *SponsorTrialStartForecastRow) GetMonthsObservedOk() (*int32, bool)`

GetMonthsObservedOk returns a tuple with the MonthsObserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthsObserved

`func (o *SponsorTrialStartForecastRow) SetMonthsObserved(v int32)`

SetMonthsObserved sets MonthsObserved field to given value.

### HasMonthsObserved

`func (o *SponsorTrialStartForecastRow) HasMonthsObserved() bool`

HasMonthsObserved returns a boolean if a field has been set.

### SetMonthsObservedNil

`func (o *SponsorTrialStartForecastRow) SetMonthsObservedNil(b bool)`

 SetMonthsObservedNil sets the value for MonthsObserved to be an explicit nil

### UnsetMonthsObserved
`func (o *SponsorTrialStartForecastRow) UnsetMonthsObserved()`

UnsetMonthsObserved ensures that no value is present for MonthsObserved, not even an explicit nil
### GetMonthsForecast

`func (o *SponsorTrialStartForecastRow) GetMonthsForecast() int32`

GetMonthsForecast returns the MonthsForecast field if non-nil, zero value otherwise.

### GetMonthsForecastOk

`func (o *SponsorTrialStartForecastRow) GetMonthsForecastOk() (*int32, bool)`

GetMonthsForecastOk returns a tuple with the MonthsForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthsForecast

`func (o *SponsorTrialStartForecastRow) SetMonthsForecast(v int32)`

SetMonthsForecast sets MonthsForecast field to given value.

### HasMonthsForecast

`func (o *SponsorTrialStartForecastRow) HasMonthsForecast() bool`

HasMonthsForecast returns a boolean if a field has been set.

### SetMonthsForecastNil

`func (o *SponsorTrialStartForecastRow) SetMonthsForecastNil(b bool)`

 SetMonthsForecastNil sets the value for MonthsForecast to be an explicit nil

### UnsetMonthsForecast
`func (o *SponsorTrialStartForecastRow) UnsetMonthsForecast()`

UnsetMonthsForecast ensures that no value is present for MonthsForecast, not even an explicit nil
### GetInputObservedThroughCutoff

`func (o *SponsorTrialStartForecastRow) GetInputObservedThroughCutoff() int32`

GetInputObservedThroughCutoff returns the InputObservedThroughCutoff field if non-nil, zero value otherwise.

### GetInputObservedThroughCutoffOk

`func (o *SponsorTrialStartForecastRow) GetInputObservedThroughCutoffOk() (*int32, bool)`

GetInputObservedThroughCutoffOk returns a tuple with the InputObservedThroughCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputObservedThroughCutoff

`func (o *SponsorTrialStartForecastRow) SetInputObservedThroughCutoff(v int32)`

SetInputObservedThroughCutoff sets InputObservedThroughCutoff field to given value.

### HasInputObservedThroughCutoff

`func (o *SponsorTrialStartForecastRow) HasInputObservedThroughCutoff() bool`

HasInputObservedThroughCutoff returns a boolean if a field has been set.

### SetInputObservedThroughCutoffNil

`func (o *SponsorTrialStartForecastRow) SetInputObservedThroughCutoffNil(b bool)`

 SetInputObservedThroughCutoffNil sets the value for InputObservedThroughCutoff to be an explicit nil

### UnsetInputObservedThroughCutoff
`func (o *SponsorTrialStartForecastRow) UnsetInputObservedThroughCutoff()`

UnsetInputObservedThroughCutoff ensures that no value is present for InputObservedThroughCutoff, not even an explicit nil
### GetInputPlannedRemaining

`func (o *SponsorTrialStartForecastRow) GetInputPlannedRemaining() int32`

GetInputPlannedRemaining returns the InputPlannedRemaining field if non-nil, zero value otherwise.

### GetInputPlannedRemainingOk

`func (o *SponsorTrialStartForecastRow) GetInputPlannedRemainingOk() (*int32, bool)`

GetInputPlannedRemainingOk returns a tuple with the InputPlannedRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPlannedRemaining

`func (o *SponsorTrialStartForecastRow) SetInputPlannedRemaining(v int32)`

SetInputPlannedRemaining sets InputPlannedRemaining field to given value.

### HasInputPlannedRemaining

`func (o *SponsorTrialStartForecastRow) HasInputPlannedRemaining() bool`

HasInputPlannedRemaining returns a boolean if a field has been set.

### SetInputPlannedRemainingNil

`func (o *SponsorTrialStartForecastRow) SetInputPlannedRemainingNil(b bool)`

 SetInputPlannedRemainingNil sets the value for InputPlannedRemaining to be an explicit nil

### UnsetInputPlannedRemaining
`func (o *SponsorTrialStartForecastRow) UnsetInputPlannedRemaining()`

UnsetInputPlannedRemaining ensures that no value is present for InputPlannedRemaining, not even an explicit nil
### GetForecastPoint

`func (o *SponsorTrialStartForecastRow) GetForecastPoint() float32`

GetForecastPoint returns the ForecastPoint field if non-nil, zero value otherwise.

### GetForecastPointOk

`func (o *SponsorTrialStartForecastRow) GetForecastPointOk() (*float32, bool)`

GetForecastPointOk returns a tuple with the ForecastPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastPoint

`func (o *SponsorTrialStartForecastRow) SetForecastPoint(v float32)`

SetForecastPoint sets ForecastPoint field to given value.

### HasForecastPoint

`func (o *SponsorTrialStartForecastRow) HasForecastPoint() bool`

HasForecastPoint returns a boolean if a field has been set.

### SetForecastPointNil

`func (o *SponsorTrialStartForecastRow) SetForecastPointNil(b bool)`

 SetForecastPointNil sets the value for ForecastPoint to be an explicit nil

### UnsetForecastPoint
`func (o *SponsorTrialStartForecastRow) UnsetForecastPoint()`

UnsetForecastPoint ensures that no value is present for ForecastPoint, not even an explicit nil
### GetForecastStatus

`func (o *SponsorTrialStartForecastRow) GetForecastStatus() string`

GetForecastStatus returns the ForecastStatus field if non-nil, zero value otherwise.

### GetForecastStatusOk

`func (o *SponsorTrialStartForecastRow) GetForecastStatusOk() (*string, bool)`

GetForecastStatusOk returns a tuple with the ForecastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastStatus

`func (o *SponsorTrialStartForecastRow) SetForecastStatus(v string)`

SetForecastStatus sets ForecastStatus field to given value.

### HasForecastStatus

`func (o *SponsorTrialStartForecastRow) HasForecastStatus() bool`

HasForecastStatus returns a boolean if a field has been set.

### SetForecastStatusNil

`func (o *SponsorTrialStartForecastRow) SetForecastStatusNil(b bool)`

 SetForecastStatusNil sets the value for ForecastStatus to be an explicit nil

### UnsetForecastStatus
`func (o *SponsorTrialStartForecastRow) UnsetForecastStatus()`

UnsetForecastStatus ensures that no value is present for ForecastStatus, not even an explicit nil
### GetInputTrialsWindow60m

`func (o *SponsorTrialStartForecastRow) GetInputTrialsWindow60m() int32`

GetInputTrialsWindow60m returns the InputTrialsWindow60m field if non-nil, zero value otherwise.

### GetInputTrialsWindow60mOk

`func (o *SponsorTrialStartForecastRow) GetInputTrialsWindow60mOk() (*int32, bool)`

GetInputTrialsWindow60mOk returns a tuple with the InputTrialsWindow60m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTrialsWindow60m

`func (o *SponsorTrialStartForecastRow) SetInputTrialsWindow60m(v int32)`

SetInputTrialsWindow60m sets InputTrialsWindow60m field to given value.

### HasInputTrialsWindow60m

`func (o *SponsorTrialStartForecastRow) HasInputTrialsWindow60m() bool`

HasInputTrialsWindow60m returns a boolean if a field has been set.

### SetInputTrialsWindow60mNil

`func (o *SponsorTrialStartForecastRow) SetInputTrialsWindow60mNil(b bool)`

 SetInputTrialsWindow60mNil sets the value for InputTrialsWindow60m to be an explicit nil

### UnsetInputTrialsWindow60m
`func (o *SponsorTrialStartForecastRow) UnsetInputTrialsWindow60m()`

UnsetInputTrialsWindow60m ensures that no value is present for InputTrialsWindow60m, not even an explicit nil
### GetInputStartsPerYear60m

`func (o *SponsorTrialStartForecastRow) GetInputStartsPerYear60m() float32`

GetInputStartsPerYear60m returns the InputStartsPerYear60m field if non-nil, zero value otherwise.

### GetInputStartsPerYear60mOk

`func (o *SponsorTrialStartForecastRow) GetInputStartsPerYear60mOk() (*float32, bool)`

GetInputStartsPerYear60mOk returns a tuple with the InputStartsPerYear60m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputStartsPerYear60m

`func (o *SponsorTrialStartForecastRow) SetInputStartsPerYear60m(v float32)`

SetInputStartsPerYear60m sets InputStartsPerYear60m field to given value.

### HasInputStartsPerYear60m

`func (o *SponsorTrialStartForecastRow) HasInputStartsPerYear60m() bool`

HasInputStartsPerYear60m returns a boolean if a field has been set.

### SetInputStartsPerYear60mNil

`func (o *SponsorTrialStartForecastRow) SetInputStartsPerYear60mNil(b bool)`

 SetInputStartsPerYear60mNil sets the value for InputStartsPerYear60m to be an explicit nil

### UnsetInputStartsPerYear60m
`func (o *SponsorTrialStartForecastRow) UnsetInputStartsPerYear60m()`

UnsetInputStartsPerYear60m ensures that no value is present for InputStartsPerYear60m, not even an explicit nil
### GetSizeBucket

`func (o *SponsorTrialStartForecastRow) GetSizeBucket() string`

GetSizeBucket returns the SizeBucket field if non-nil, zero value otherwise.

### GetSizeBucketOk

`func (o *SponsorTrialStartForecastRow) GetSizeBucketOk() (*string, bool)`

GetSizeBucketOk returns a tuple with the SizeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBucket

`func (o *SponsorTrialStartForecastRow) SetSizeBucket(v string)`

SetSizeBucket sets SizeBucket field to given value.

### HasSizeBucket

`func (o *SponsorTrialStartForecastRow) HasSizeBucket() bool`

HasSizeBucket returns a boolean if a field has been set.

### SetSizeBucketNil

`func (o *SponsorTrialStartForecastRow) SetSizeBucketNil(b bool)`

 SetSizeBucketNil sets the value for SizeBucket to be an explicit nil

### UnsetSizeBucket
`func (o *SponsorTrialStartForecastRow) UnsetSizeBucket()`

UnsetSizeBucket ensures that no value is present for SizeBucket, not even an explicit nil
### GetBacktestHorizon

`func (o *SponsorTrialStartForecastRow) GetBacktestHorizon() string`

GetBacktestHorizon returns the BacktestHorizon field if non-nil, zero value otherwise.

### GetBacktestHorizonOk

`func (o *SponsorTrialStartForecastRow) GetBacktestHorizonOk() (*string, bool)`

GetBacktestHorizonOk returns a tuple with the BacktestHorizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestHorizon

`func (o *SponsorTrialStartForecastRow) SetBacktestHorizon(v string)`

SetBacktestHorizon sets BacktestHorizon field to given value.

### HasBacktestHorizon

`func (o *SponsorTrialStartForecastRow) HasBacktestHorizon() bool`

HasBacktestHorizon returns a boolean if a field has been set.

### SetBacktestHorizonNil

`func (o *SponsorTrialStartForecastRow) SetBacktestHorizonNil(b bool)`

 SetBacktestHorizonNil sets the value for BacktestHorizon to be an explicit nil

### UnsetBacktestHorizon
`func (o *SponsorTrialStartForecastRow) UnsetBacktestHorizon()`

UnsetBacktestHorizon ensures that no value is present for BacktestHorizon, not even an explicit nil
### GetBacktestSmapePct

`func (o *SponsorTrialStartForecastRow) GetBacktestSmapePct() float32`

GetBacktestSmapePct returns the BacktestSmapePct field if non-nil, zero value otherwise.

### GetBacktestSmapePctOk

`func (o *SponsorTrialStartForecastRow) GetBacktestSmapePctOk() (*float32, bool)`

GetBacktestSmapePctOk returns a tuple with the BacktestSmapePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestSmapePct

`func (o *SponsorTrialStartForecastRow) SetBacktestSmapePct(v float32)`

SetBacktestSmapePct sets BacktestSmapePct field to given value.

### HasBacktestSmapePct

`func (o *SponsorTrialStartForecastRow) HasBacktestSmapePct() bool`

HasBacktestSmapePct returns a boolean if a field has been set.

### SetBacktestSmapePctNil

`func (o *SponsorTrialStartForecastRow) SetBacktestSmapePctNil(b bool)`

 SetBacktestSmapePctNil sets the value for BacktestSmapePct to be an explicit nil

### UnsetBacktestSmapePct
`func (o *SponsorTrialStartForecastRow) UnsetBacktestSmapePct()`

UnsetBacktestSmapePct ensures that no value is present for BacktestSmapePct, not even an explicit nil
### GetBacktestRelBias

`func (o *SponsorTrialStartForecastRow) GetBacktestRelBias() float32`

GetBacktestRelBias returns the BacktestRelBias field if non-nil, zero value otherwise.

### GetBacktestRelBiasOk

`func (o *SponsorTrialStartForecastRow) GetBacktestRelBiasOk() (*float32, bool)`

GetBacktestRelBiasOk returns a tuple with the BacktestRelBias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestRelBias

`func (o *SponsorTrialStartForecastRow) SetBacktestRelBias(v float32)`

SetBacktestRelBias sets BacktestRelBias field to given value.

### HasBacktestRelBias

`func (o *SponsorTrialStartForecastRow) HasBacktestRelBias() bool`

HasBacktestRelBias returns a boolean if a field has been set.

### SetBacktestRelBiasNil

`func (o *SponsorTrialStartForecastRow) SetBacktestRelBiasNil(b bool)`

 SetBacktestRelBiasNil sets the value for BacktestRelBias to be an explicit nil

### UnsetBacktestRelBias
`func (o *SponsorTrialStartForecastRow) UnsetBacktestRelBias()`

UnsetBacktestRelBias ensures that no value is present for BacktestRelBias, not even an explicit nil
### GetBacktestRelmaeVsSnaive

`func (o *SponsorTrialStartForecastRow) GetBacktestRelmaeVsSnaive() float32`

GetBacktestRelmaeVsSnaive returns the BacktestRelmaeVsSnaive field if non-nil, zero value otherwise.

### GetBacktestRelmaeVsSnaiveOk

`func (o *SponsorTrialStartForecastRow) GetBacktestRelmaeVsSnaiveOk() (*float32, bool)`

GetBacktestRelmaeVsSnaiveOk returns a tuple with the BacktestRelmaeVsSnaive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestRelmaeVsSnaive

`func (o *SponsorTrialStartForecastRow) SetBacktestRelmaeVsSnaive(v float32)`

SetBacktestRelmaeVsSnaive sets BacktestRelmaeVsSnaive field to given value.

### HasBacktestRelmaeVsSnaive

`func (o *SponsorTrialStartForecastRow) HasBacktestRelmaeVsSnaive() bool`

HasBacktestRelmaeVsSnaive returns a boolean if a field has been set.

### SetBacktestRelmaeVsSnaiveNil

`func (o *SponsorTrialStartForecastRow) SetBacktestRelmaeVsSnaiveNil(b bool)`

 SetBacktestRelmaeVsSnaiveNil sets the value for BacktestRelmaeVsSnaive to be an explicit nil

### UnsetBacktestRelmaeVsSnaive
`func (o *SponsorTrialStartForecastRow) UnsetBacktestRelmaeVsSnaive()`

UnsetBacktestRelmaeVsSnaive ensures that no value is present for BacktestRelmaeVsSnaive, not even an explicit nil
### GetBacktestNSeries

`func (o *SponsorTrialStartForecastRow) GetBacktestNSeries() int32`

GetBacktestNSeries returns the BacktestNSeries field if non-nil, zero value otherwise.

### GetBacktestNSeriesOk

`func (o *SponsorTrialStartForecastRow) GetBacktestNSeriesOk() (*int32, bool)`

GetBacktestNSeriesOk returns a tuple with the BacktestNSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestNSeries

`func (o *SponsorTrialStartForecastRow) SetBacktestNSeries(v int32)`

SetBacktestNSeries sets BacktestNSeries field to given value.

### HasBacktestNSeries

`func (o *SponsorTrialStartForecastRow) HasBacktestNSeries() bool`

HasBacktestNSeries returns a boolean if a field has been set.

### SetBacktestNSeriesNil

`func (o *SponsorTrialStartForecastRow) SetBacktestNSeriesNil(b bool)`

 SetBacktestNSeriesNil sets the value for BacktestNSeries to be an explicit nil

### UnsetBacktestNSeries
`func (o *SponsorTrialStartForecastRow) UnsetBacktestNSeries()`

UnsetBacktestNSeries ensures that no value is present for BacktestNSeries, not even an explicit nil
### GetBacktestRef

`func (o *SponsorTrialStartForecastRow) GetBacktestRef() string`

GetBacktestRef returns the BacktestRef field if non-nil, zero value otherwise.

### GetBacktestRefOk

`func (o *SponsorTrialStartForecastRow) GetBacktestRefOk() (*string, bool)`

GetBacktestRefOk returns a tuple with the BacktestRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktestRef

`func (o *SponsorTrialStartForecastRow) SetBacktestRef(v string)`

SetBacktestRef sets BacktestRef field to given value.

### HasBacktestRef

`func (o *SponsorTrialStartForecastRow) HasBacktestRef() bool`

HasBacktestRef returns a boolean if a field has been set.

### SetBacktestRefNil

`func (o *SponsorTrialStartForecastRow) SetBacktestRefNil(b bool)`

 SetBacktestRefNil sets the value for BacktestRef to be an explicit nil

### UnsetBacktestRef
`func (o *SponsorTrialStartForecastRow) UnsetBacktestRef()`

UnsetBacktestRef ensures that no value is present for BacktestRef, not even an explicit nil
### GetModelName

`func (o *SponsorTrialStartForecastRow) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *SponsorTrialStartForecastRow) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *SponsorTrialStartForecastRow) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *SponsorTrialStartForecastRow) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### SetModelNameNil

`func (o *SponsorTrialStartForecastRow) SetModelNameNil(b bool)`

 SetModelNameNil sets the value for ModelName to be an explicit nil

### UnsetModelName
`func (o *SponsorTrialStartForecastRow) UnsetModelName()`

UnsetModelName ensures that no value is present for ModelName, not even an explicit nil
### GetModelVersion

`func (o *SponsorTrialStartForecastRow) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SponsorTrialStartForecastRow) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SponsorTrialStartForecastRow) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SponsorTrialStartForecastRow) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### SetModelVersionNil

`func (o *SponsorTrialStartForecastRow) SetModelVersionNil(b bool)`

 SetModelVersionNil sets the value for ModelVersion to be an explicit nil

### UnsetModelVersion
`func (o *SponsorTrialStartForecastRow) UnsetModelVersion()`

UnsetModelVersion ensures that no value is present for ModelVersion, not even an explicit nil
### GetTrainingCutoff

`func (o *SponsorTrialStartForecastRow) GetTrainingCutoff() string`

GetTrainingCutoff returns the TrainingCutoff field if non-nil, zero value otherwise.

### GetTrainingCutoffOk

`func (o *SponsorTrialStartForecastRow) GetTrainingCutoffOk() (*string, bool)`

GetTrainingCutoffOk returns a tuple with the TrainingCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingCutoff

`func (o *SponsorTrialStartForecastRow) SetTrainingCutoff(v string)`

SetTrainingCutoff sets TrainingCutoff field to given value.

### HasTrainingCutoff

`func (o *SponsorTrialStartForecastRow) HasTrainingCutoff() bool`

HasTrainingCutoff returns a boolean if a field has been set.

### SetTrainingCutoffNil

`func (o *SponsorTrialStartForecastRow) SetTrainingCutoffNil(b bool)`

 SetTrainingCutoffNil sets the value for TrainingCutoff to be an explicit nil

### UnsetTrainingCutoff
`func (o *SponsorTrialStartForecastRow) UnsetTrainingCutoff()`

UnsetTrainingCutoff ensures that no value is present for TrainingCutoff, not even an explicit nil
### GetMethod

`func (o *SponsorTrialStartForecastRow) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SponsorTrialStartForecastRow) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SponsorTrialStartForecastRow) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SponsorTrialStartForecastRow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *SponsorTrialStartForecastRow) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *SponsorTrialStartForecastRow) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


