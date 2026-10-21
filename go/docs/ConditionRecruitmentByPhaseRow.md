# ConditionRecruitmentByPhaseRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s MeSH heading. | [optional] 
**PhaseBucket** | Pointer to **NullableString** | The phase bucket of the trial (one vocabulary for predicted and realized rows). | [optional] 
**NPredictedOpen** | Pointer to **NullableInt32** | Per phase bucket, the same count as the series column of that name. | [optional] 
**NPredictedAtCap** | Pointer to **NullableInt32** | Per phase bucket, the same count as the series column of that name. | [optional] 
**NPredictedOverdue** | Pointer to **NullableInt32** | Per phase bucket, the same count as the series column of that name. | [optional] 
**PredictedMonthsMedian** | Pointer to **NullableFloat32** | Median over the bucket&#39;s rows; NULL under min_n. | [optional] 
**BaselineMonthsMedian** | Pointer to **NullableFloat32** | Median over the bucket&#39;s rows; NULL under min_n. | [optional] 
**NRealized** | Pointer to **NullableInt32** | Per phase bucket, the same count as the series column of that name. | [optional] 
**RealizedMonthsMedian** | Pointer to **NullableFloat32** | Median over the bucket&#39;s rows; NULL under min_n. | [optional] 
**ErrorPhaseMaeMonths** | Pointer to **NullableFloat32** | Backtest mean absolute error in months for this phase bucket; NULL on an unmeasured bucket or one with no open predicted trial. | [optional] 
**ErrorPhaseRelmaeVsBaseline** | Pointer to **NullableFloat32** | The model&#39;s mean absolute error (MAE) over the baseline&#39;s MAE in this phase bucket. | [optional] 
**ErrorPhaseBaselineMaeMonths** | Pointer to **NullableFloat32** | The baseline&#39;s mean absolute error (MAE) in months in this phase bucket. | [optional] 
**ErrorPhaseNTrials** | Pointer to **NullableInt32** | Backtest trials behind the phase figure. | [optional] 
**ErrorPhaseNYears** | Pointer to **NullableInt32** | Backtest years behind the phase figure. | [optional] 
**MinN** | Pointer to **NullableInt32** | The floor applied: a median is withheld below this many trials, carried so a reader sees the floor used. | [optional] 
**TrainingCutoff** | Pointer to **NullableString** | The date the model&#39;s training data ends. | [optional] 

## Methods

### NewConditionRecruitmentByPhaseRow

`func NewConditionRecruitmentByPhaseRow() *ConditionRecruitmentByPhaseRow`

NewConditionRecruitmentByPhaseRow instantiates a new ConditionRecruitmentByPhaseRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionRecruitmentByPhaseRowWithDefaults

`func NewConditionRecruitmentByPhaseRowWithDefaults() *ConditionRecruitmentByPhaseRow`

NewConditionRecruitmentByPhaseRowWithDefaults instantiates a new ConditionRecruitmentByPhaseRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionRecruitmentByPhaseRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionRecruitmentByPhaseRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionRecruitmentByPhaseRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionRecruitmentByPhaseRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionRecruitmentByPhaseRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionRecruitmentByPhaseRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *ConditionRecruitmentByPhaseRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ConditionRecruitmentByPhaseRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ConditionRecruitmentByPhaseRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *ConditionRecruitmentByPhaseRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *ConditionRecruitmentByPhaseRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *ConditionRecruitmentByPhaseRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetPhaseBucket

`func (o *ConditionRecruitmentByPhaseRow) GetPhaseBucket() string`

GetPhaseBucket returns the PhaseBucket field if non-nil, zero value otherwise.

### GetPhaseBucketOk

`func (o *ConditionRecruitmentByPhaseRow) GetPhaseBucketOk() (*string, bool)`

GetPhaseBucketOk returns a tuple with the PhaseBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseBucket

`func (o *ConditionRecruitmentByPhaseRow) SetPhaseBucket(v string)`

SetPhaseBucket sets PhaseBucket field to given value.

### HasPhaseBucket

`func (o *ConditionRecruitmentByPhaseRow) HasPhaseBucket() bool`

HasPhaseBucket returns a boolean if a field has been set.

### SetPhaseBucketNil

`func (o *ConditionRecruitmentByPhaseRow) SetPhaseBucketNil(b bool)`

 SetPhaseBucketNil sets the value for PhaseBucket to be an explicit nil

### UnsetPhaseBucket
`func (o *ConditionRecruitmentByPhaseRow) UnsetPhaseBucket()`

UnsetPhaseBucket ensures that no value is present for PhaseBucket, not even an explicit nil
### GetNPredictedOpen

`func (o *ConditionRecruitmentByPhaseRow) GetNPredictedOpen() int32`

GetNPredictedOpen returns the NPredictedOpen field if non-nil, zero value otherwise.

### GetNPredictedOpenOk

`func (o *ConditionRecruitmentByPhaseRow) GetNPredictedOpenOk() (*int32, bool)`

GetNPredictedOpenOk returns a tuple with the NPredictedOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedOpen

`func (o *ConditionRecruitmentByPhaseRow) SetNPredictedOpen(v int32)`

SetNPredictedOpen sets NPredictedOpen field to given value.

### HasNPredictedOpen

`func (o *ConditionRecruitmentByPhaseRow) HasNPredictedOpen() bool`

HasNPredictedOpen returns a boolean if a field has been set.

### SetNPredictedOpenNil

`func (o *ConditionRecruitmentByPhaseRow) SetNPredictedOpenNil(b bool)`

 SetNPredictedOpenNil sets the value for NPredictedOpen to be an explicit nil

### UnsetNPredictedOpen
`func (o *ConditionRecruitmentByPhaseRow) UnsetNPredictedOpen()`

UnsetNPredictedOpen ensures that no value is present for NPredictedOpen, not even an explicit nil
### GetNPredictedAtCap

`func (o *ConditionRecruitmentByPhaseRow) GetNPredictedAtCap() int32`

GetNPredictedAtCap returns the NPredictedAtCap field if non-nil, zero value otherwise.

### GetNPredictedAtCapOk

`func (o *ConditionRecruitmentByPhaseRow) GetNPredictedAtCapOk() (*int32, bool)`

GetNPredictedAtCapOk returns a tuple with the NPredictedAtCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedAtCap

`func (o *ConditionRecruitmentByPhaseRow) SetNPredictedAtCap(v int32)`

SetNPredictedAtCap sets NPredictedAtCap field to given value.

### HasNPredictedAtCap

`func (o *ConditionRecruitmentByPhaseRow) HasNPredictedAtCap() bool`

HasNPredictedAtCap returns a boolean if a field has been set.

### SetNPredictedAtCapNil

`func (o *ConditionRecruitmentByPhaseRow) SetNPredictedAtCapNil(b bool)`

 SetNPredictedAtCapNil sets the value for NPredictedAtCap to be an explicit nil

### UnsetNPredictedAtCap
`func (o *ConditionRecruitmentByPhaseRow) UnsetNPredictedAtCap()`

UnsetNPredictedAtCap ensures that no value is present for NPredictedAtCap, not even an explicit nil
### GetNPredictedOverdue

`func (o *ConditionRecruitmentByPhaseRow) GetNPredictedOverdue() int32`

GetNPredictedOverdue returns the NPredictedOverdue field if non-nil, zero value otherwise.

### GetNPredictedOverdueOk

`func (o *ConditionRecruitmentByPhaseRow) GetNPredictedOverdueOk() (*int32, bool)`

GetNPredictedOverdueOk returns a tuple with the NPredictedOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedOverdue

`func (o *ConditionRecruitmentByPhaseRow) SetNPredictedOverdue(v int32)`

SetNPredictedOverdue sets NPredictedOverdue field to given value.

### HasNPredictedOverdue

`func (o *ConditionRecruitmentByPhaseRow) HasNPredictedOverdue() bool`

HasNPredictedOverdue returns a boolean if a field has been set.

### SetNPredictedOverdueNil

`func (o *ConditionRecruitmentByPhaseRow) SetNPredictedOverdueNil(b bool)`

 SetNPredictedOverdueNil sets the value for NPredictedOverdue to be an explicit nil

### UnsetNPredictedOverdue
`func (o *ConditionRecruitmentByPhaseRow) UnsetNPredictedOverdue()`

UnsetNPredictedOverdue ensures that no value is present for NPredictedOverdue, not even an explicit nil
### GetPredictedMonthsMedian

`func (o *ConditionRecruitmentByPhaseRow) GetPredictedMonthsMedian() float32`

GetPredictedMonthsMedian returns the PredictedMonthsMedian field if non-nil, zero value otherwise.

### GetPredictedMonthsMedianOk

`func (o *ConditionRecruitmentByPhaseRow) GetPredictedMonthsMedianOk() (*float32, bool)`

GetPredictedMonthsMedianOk returns a tuple with the PredictedMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMonthsMedian

`func (o *ConditionRecruitmentByPhaseRow) SetPredictedMonthsMedian(v float32)`

SetPredictedMonthsMedian sets PredictedMonthsMedian field to given value.

### HasPredictedMonthsMedian

`func (o *ConditionRecruitmentByPhaseRow) HasPredictedMonthsMedian() bool`

HasPredictedMonthsMedian returns a boolean if a field has been set.

### SetPredictedMonthsMedianNil

`func (o *ConditionRecruitmentByPhaseRow) SetPredictedMonthsMedianNil(b bool)`

 SetPredictedMonthsMedianNil sets the value for PredictedMonthsMedian to be an explicit nil

### UnsetPredictedMonthsMedian
`func (o *ConditionRecruitmentByPhaseRow) UnsetPredictedMonthsMedian()`

UnsetPredictedMonthsMedian ensures that no value is present for PredictedMonthsMedian, not even an explicit nil
### GetBaselineMonthsMedian

`func (o *ConditionRecruitmentByPhaseRow) GetBaselineMonthsMedian() float32`

GetBaselineMonthsMedian returns the BaselineMonthsMedian field if non-nil, zero value otherwise.

### GetBaselineMonthsMedianOk

`func (o *ConditionRecruitmentByPhaseRow) GetBaselineMonthsMedianOk() (*float32, bool)`

GetBaselineMonthsMedianOk returns a tuple with the BaselineMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineMonthsMedian

`func (o *ConditionRecruitmentByPhaseRow) SetBaselineMonthsMedian(v float32)`

SetBaselineMonthsMedian sets BaselineMonthsMedian field to given value.

### HasBaselineMonthsMedian

`func (o *ConditionRecruitmentByPhaseRow) HasBaselineMonthsMedian() bool`

HasBaselineMonthsMedian returns a boolean if a field has been set.

### SetBaselineMonthsMedianNil

`func (o *ConditionRecruitmentByPhaseRow) SetBaselineMonthsMedianNil(b bool)`

 SetBaselineMonthsMedianNil sets the value for BaselineMonthsMedian to be an explicit nil

### UnsetBaselineMonthsMedian
`func (o *ConditionRecruitmentByPhaseRow) UnsetBaselineMonthsMedian()`

UnsetBaselineMonthsMedian ensures that no value is present for BaselineMonthsMedian, not even an explicit nil
### GetNRealized

`func (o *ConditionRecruitmentByPhaseRow) GetNRealized() int32`

GetNRealized returns the NRealized field if non-nil, zero value otherwise.

### GetNRealizedOk

`func (o *ConditionRecruitmentByPhaseRow) GetNRealizedOk() (*int32, bool)`

GetNRealizedOk returns a tuple with the NRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealized

`func (o *ConditionRecruitmentByPhaseRow) SetNRealized(v int32)`

SetNRealized sets NRealized field to given value.

### HasNRealized

`func (o *ConditionRecruitmentByPhaseRow) HasNRealized() bool`

HasNRealized returns a boolean if a field has been set.

### SetNRealizedNil

`func (o *ConditionRecruitmentByPhaseRow) SetNRealizedNil(b bool)`

 SetNRealizedNil sets the value for NRealized to be an explicit nil

### UnsetNRealized
`func (o *ConditionRecruitmentByPhaseRow) UnsetNRealized()`

UnsetNRealized ensures that no value is present for NRealized, not even an explicit nil
### GetRealizedMonthsMedian

`func (o *ConditionRecruitmentByPhaseRow) GetRealizedMonthsMedian() float32`

GetRealizedMonthsMedian returns the RealizedMonthsMedian field if non-nil, zero value otherwise.

### GetRealizedMonthsMedianOk

`func (o *ConditionRecruitmentByPhaseRow) GetRealizedMonthsMedianOk() (*float32, bool)`

GetRealizedMonthsMedianOk returns a tuple with the RealizedMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedMonthsMedian

`func (o *ConditionRecruitmentByPhaseRow) SetRealizedMonthsMedian(v float32)`

SetRealizedMonthsMedian sets RealizedMonthsMedian field to given value.

### HasRealizedMonthsMedian

`func (o *ConditionRecruitmentByPhaseRow) HasRealizedMonthsMedian() bool`

HasRealizedMonthsMedian returns a boolean if a field has been set.

### SetRealizedMonthsMedianNil

`func (o *ConditionRecruitmentByPhaseRow) SetRealizedMonthsMedianNil(b bool)`

 SetRealizedMonthsMedianNil sets the value for RealizedMonthsMedian to be an explicit nil

### UnsetRealizedMonthsMedian
`func (o *ConditionRecruitmentByPhaseRow) UnsetRealizedMonthsMedian()`

UnsetRealizedMonthsMedian ensures that no value is present for RealizedMonthsMedian, not even an explicit nil
### GetErrorPhaseMaeMonths

`func (o *ConditionRecruitmentByPhaseRow) GetErrorPhaseMaeMonths() float32`

GetErrorPhaseMaeMonths returns the ErrorPhaseMaeMonths field if non-nil, zero value otherwise.

### GetErrorPhaseMaeMonthsOk

`func (o *ConditionRecruitmentByPhaseRow) GetErrorPhaseMaeMonthsOk() (*float32, bool)`

GetErrorPhaseMaeMonthsOk returns a tuple with the ErrorPhaseMaeMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPhaseMaeMonths

`func (o *ConditionRecruitmentByPhaseRow) SetErrorPhaseMaeMonths(v float32)`

SetErrorPhaseMaeMonths sets ErrorPhaseMaeMonths field to given value.

### HasErrorPhaseMaeMonths

`func (o *ConditionRecruitmentByPhaseRow) HasErrorPhaseMaeMonths() bool`

HasErrorPhaseMaeMonths returns a boolean if a field has been set.

### SetErrorPhaseMaeMonthsNil

`func (o *ConditionRecruitmentByPhaseRow) SetErrorPhaseMaeMonthsNil(b bool)`

 SetErrorPhaseMaeMonthsNil sets the value for ErrorPhaseMaeMonths to be an explicit nil

### UnsetErrorPhaseMaeMonths
`func (o *ConditionRecruitmentByPhaseRow) UnsetErrorPhaseMaeMonths()`

UnsetErrorPhaseMaeMonths ensures that no value is present for ErrorPhaseMaeMonths, not even an explicit nil
### GetErrorPhaseRelmaeVsBaseline

`func (o *ConditionRecruitmentByPhaseRow) GetErrorPhaseRelmaeVsBaseline() float32`

GetErrorPhaseRelmaeVsBaseline returns the ErrorPhaseRelmaeVsBaseline field if non-nil, zero value otherwise.

### GetErrorPhaseRelmaeVsBaselineOk

`func (o *ConditionRecruitmentByPhaseRow) GetErrorPhaseRelmaeVsBaselineOk() (*float32, bool)`

GetErrorPhaseRelmaeVsBaselineOk returns a tuple with the ErrorPhaseRelmaeVsBaseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPhaseRelmaeVsBaseline

`func (o *ConditionRecruitmentByPhaseRow) SetErrorPhaseRelmaeVsBaseline(v float32)`

SetErrorPhaseRelmaeVsBaseline sets ErrorPhaseRelmaeVsBaseline field to given value.

### HasErrorPhaseRelmaeVsBaseline

`func (o *ConditionRecruitmentByPhaseRow) HasErrorPhaseRelmaeVsBaseline() bool`

HasErrorPhaseRelmaeVsBaseline returns a boolean if a field has been set.

### SetErrorPhaseRelmaeVsBaselineNil

`func (o *ConditionRecruitmentByPhaseRow) SetErrorPhaseRelmaeVsBaselineNil(b bool)`

 SetErrorPhaseRelmaeVsBaselineNil sets the value for ErrorPhaseRelmaeVsBaseline to be an explicit nil

### UnsetErrorPhaseRelmaeVsBaseline
`func (o *ConditionRecruitmentByPhaseRow) UnsetErrorPhaseRelmaeVsBaseline()`

UnsetErrorPhaseRelmaeVsBaseline ensures that no value is present for ErrorPhaseRelmaeVsBaseline, not even an explicit nil
### GetErrorPhaseBaselineMaeMonths

`func (o *ConditionRecruitmentByPhaseRow) GetErrorPhaseBaselineMaeMonths() float32`

GetErrorPhaseBaselineMaeMonths returns the ErrorPhaseBaselineMaeMonths field if non-nil, zero value otherwise.

### GetErrorPhaseBaselineMaeMonthsOk

`func (o *ConditionRecruitmentByPhaseRow) GetErrorPhaseBaselineMaeMonthsOk() (*float32, bool)`

GetErrorPhaseBaselineMaeMonthsOk returns a tuple with the ErrorPhaseBaselineMaeMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPhaseBaselineMaeMonths

`func (o *ConditionRecruitmentByPhaseRow) SetErrorPhaseBaselineMaeMonths(v float32)`

SetErrorPhaseBaselineMaeMonths sets ErrorPhaseBaselineMaeMonths field to given value.

### HasErrorPhaseBaselineMaeMonths

`func (o *ConditionRecruitmentByPhaseRow) HasErrorPhaseBaselineMaeMonths() bool`

HasErrorPhaseBaselineMaeMonths returns a boolean if a field has been set.

### SetErrorPhaseBaselineMaeMonthsNil

`func (o *ConditionRecruitmentByPhaseRow) SetErrorPhaseBaselineMaeMonthsNil(b bool)`

 SetErrorPhaseBaselineMaeMonthsNil sets the value for ErrorPhaseBaselineMaeMonths to be an explicit nil

### UnsetErrorPhaseBaselineMaeMonths
`func (o *ConditionRecruitmentByPhaseRow) UnsetErrorPhaseBaselineMaeMonths()`

UnsetErrorPhaseBaselineMaeMonths ensures that no value is present for ErrorPhaseBaselineMaeMonths, not even an explicit nil
### GetErrorPhaseNTrials

`func (o *ConditionRecruitmentByPhaseRow) GetErrorPhaseNTrials() int32`

GetErrorPhaseNTrials returns the ErrorPhaseNTrials field if non-nil, zero value otherwise.

### GetErrorPhaseNTrialsOk

`func (o *ConditionRecruitmentByPhaseRow) GetErrorPhaseNTrialsOk() (*int32, bool)`

GetErrorPhaseNTrialsOk returns a tuple with the ErrorPhaseNTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPhaseNTrials

`func (o *ConditionRecruitmentByPhaseRow) SetErrorPhaseNTrials(v int32)`

SetErrorPhaseNTrials sets ErrorPhaseNTrials field to given value.

### HasErrorPhaseNTrials

`func (o *ConditionRecruitmentByPhaseRow) HasErrorPhaseNTrials() bool`

HasErrorPhaseNTrials returns a boolean if a field has been set.

### SetErrorPhaseNTrialsNil

`func (o *ConditionRecruitmentByPhaseRow) SetErrorPhaseNTrialsNil(b bool)`

 SetErrorPhaseNTrialsNil sets the value for ErrorPhaseNTrials to be an explicit nil

### UnsetErrorPhaseNTrials
`func (o *ConditionRecruitmentByPhaseRow) UnsetErrorPhaseNTrials()`

UnsetErrorPhaseNTrials ensures that no value is present for ErrorPhaseNTrials, not even an explicit nil
### GetErrorPhaseNYears

`func (o *ConditionRecruitmentByPhaseRow) GetErrorPhaseNYears() int32`

GetErrorPhaseNYears returns the ErrorPhaseNYears field if non-nil, zero value otherwise.

### GetErrorPhaseNYearsOk

`func (o *ConditionRecruitmentByPhaseRow) GetErrorPhaseNYearsOk() (*int32, bool)`

GetErrorPhaseNYearsOk returns a tuple with the ErrorPhaseNYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPhaseNYears

`func (o *ConditionRecruitmentByPhaseRow) SetErrorPhaseNYears(v int32)`

SetErrorPhaseNYears sets ErrorPhaseNYears field to given value.

### HasErrorPhaseNYears

`func (o *ConditionRecruitmentByPhaseRow) HasErrorPhaseNYears() bool`

HasErrorPhaseNYears returns a boolean if a field has been set.

### SetErrorPhaseNYearsNil

`func (o *ConditionRecruitmentByPhaseRow) SetErrorPhaseNYearsNil(b bool)`

 SetErrorPhaseNYearsNil sets the value for ErrorPhaseNYears to be an explicit nil

### UnsetErrorPhaseNYears
`func (o *ConditionRecruitmentByPhaseRow) UnsetErrorPhaseNYears()`

UnsetErrorPhaseNYears ensures that no value is present for ErrorPhaseNYears, not even an explicit nil
### GetMinN

`func (o *ConditionRecruitmentByPhaseRow) GetMinN() int32`

GetMinN returns the MinN field if non-nil, zero value otherwise.

### GetMinNOk

`func (o *ConditionRecruitmentByPhaseRow) GetMinNOk() (*int32, bool)`

GetMinNOk returns a tuple with the MinN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinN

`func (o *ConditionRecruitmentByPhaseRow) SetMinN(v int32)`

SetMinN sets MinN field to given value.

### HasMinN

`func (o *ConditionRecruitmentByPhaseRow) HasMinN() bool`

HasMinN returns a boolean if a field has been set.

### SetMinNNil

`func (o *ConditionRecruitmentByPhaseRow) SetMinNNil(b bool)`

 SetMinNNil sets the value for MinN to be an explicit nil

### UnsetMinN
`func (o *ConditionRecruitmentByPhaseRow) UnsetMinN()`

UnsetMinN ensures that no value is present for MinN, not even an explicit nil
### GetTrainingCutoff

`func (o *ConditionRecruitmentByPhaseRow) GetTrainingCutoff() string`

GetTrainingCutoff returns the TrainingCutoff field if non-nil, zero value otherwise.

### GetTrainingCutoffOk

`func (o *ConditionRecruitmentByPhaseRow) GetTrainingCutoffOk() (*string, bool)`

GetTrainingCutoffOk returns a tuple with the TrainingCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingCutoff

`func (o *ConditionRecruitmentByPhaseRow) SetTrainingCutoff(v string)`

SetTrainingCutoff sets TrainingCutoff field to given value.

### HasTrainingCutoff

`func (o *ConditionRecruitmentByPhaseRow) HasTrainingCutoff() bool`

HasTrainingCutoff returns a boolean if a field has been set.

### SetTrainingCutoffNil

`func (o *ConditionRecruitmentByPhaseRow) SetTrainingCutoffNil(b bool)`

 SetTrainingCutoffNil sets the value for TrainingCutoff to be an explicit nil

### UnsetTrainingCutoff
`func (o *ConditionRecruitmentByPhaseRow) UnsetTrainingCutoff()`

UnsetTrainingCutoff ensures that no value is present for TrainingCutoff, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


