# SponsorRecruitmentByPhaseRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the series is for (the lead sponsor). | [optional] 
**SponsorName** | Pointer to **NullableString** | Display name of the organization record. | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company the organization belongs to (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
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

### NewSponsorRecruitmentByPhaseRow

`func NewSponsorRecruitmentByPhaseRow() *SponsorRecruitmentByPhaseRow`

NewSponsorRecruitmentByPhaseRow instantiates a new SponsorRecruitmentByPhaseRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSponsorRecruitmentByPhaseRowWithDefaults

`func NewSponsorRecruitmentByPhaseRowWithDefaults() *SponsorRecruitmentByPhaseRow`

NewSponsorRecruitmentByPhaseRowWithDefaults instantiates a new SponsorRecruitmentByPhaseRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSponsorOrgId

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *SponsorRecruitmentByPhaseRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *SponsorRecruitmentByPhaseRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorName

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *SponsorRecruitmentByPhaseRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *SponsorRecruitmentByPhaseRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorCompanyId

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *SponsorRecruitmentByPhaseRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *SponsorRecruitmentByPhaseRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *SponsorRecruitmentByPhaseRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *SponsorRecruitmentByPhaseRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *SponsorRecruitmentByPhaseRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *SponsorRecruitmentByPhaseRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *SponsorRecruitmentByPhaseRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *SponsorRecruitmentByPhaseRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *SponsorRecruitmentByPhaseRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *SponsorRecruitmentByPhaseRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetPhaseBucket

`func (o *SponsorRecruitmentByPhaseRow) GetPhaseBucket() string`

GetPhaseBucket returns the PhaseBucket field if non-nil, zero value otherwise.

### GetPhaseBucketOk

`func (o *SponsorRecruitmentByPhaseRow) GetPhaseBucketOk() (*string, bool)`

GetPhaseBucketOk returns a tuple with the PhaseBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseBucket

`func (o *SponsorRecruitmentByPhaseRow) SetPhaseBucket(v string)`

SetPhaseBucket sets PhaseBucket field to given value.

### HasPhaseBucket

`func (o *SponsorRecruitmentByPhaseRow) HasPhaseBucket() bool`

HasPhaseBucket returns a boolean if a field has been set.

### SetPhaseBucketNil

`func (o *SponsorRecruitmentByPhaseRow) SetPhaseBucketNil(b bool)`

 SetPhaseBucketNil sets the value for PhaseBucket to be an explicit nil

### UnsetPhaseBucket
`func (o *SponsorRecruitmentByPhaseRow) UnsetPhaseBucket()`

UnsetPhaseBucket ensures that no value is present for PhaseBucket, not even an explicit nil
### GetNPredictedOpen

`func (o *SponsorRecruitmentByPhaseRow) GetNPredictedOpen() int32`

GetNPredictedOpen returns the NPredictedOpen field if non-nil, zero value otherwise.

### GetNPredictedOpenOk

`func (o *SponsorRecruitmentByPhaseRow) GetNPredictedOpenOk() (*int32, bool)`

GetNPredictedOpenOk returns a tuple with the NPredictedOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedOpen

`func (o *SponsorRecruitmentByPhaseRow) SetNPredictedOpen(v int32)`

SetNPredictedOpen sets NPredictedOpen field to given value.

### HasNPredictedOpen

`func (o *SponsorRecruitmentByPhaseRow) HasNPredictedOpen() bool`

HasNPredictedOpen returns a boolean if a field has been set.

### SetNPredictedOpenNil

`func (o *SponsorRecruitmentByPhaseRow) SetNPredictedOpenNil(b bool)`

 SetNPredictedOpenNil sets the value for NPredictedOpen to be an explicit nil

### UnsetNPredictedOpen
`func (o *SponsorRecruitmentByPhaseRow) UnsetNPredictedOpen()`

UnsetNPredictedOpen ensures that no value is present for NPredictedOpen, not even an explicit nil
### GetNPredictedAtCap

`func (o *SponsorRecruitmentByPhaseRow) GetNPredictedAtCap() int32`

GetNPredictedAtCap returns the NPredictedAtCap field if non-nil, zero value otherwise.

### GetNPredictedAtCapOk

`func (o *SponsorRecruitmentByPhaseRow) GetNPredictedAtCapOk() (*int32, bool)`

GetNPredictedAtCapOk returns a tuple with the NPredictedAtCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedAtCap

`func (o *SponsorRecruitmentByPhaseRow) SetNPredictedAtCap(v int32)`

SetNPredictedAtCap sets NPredictedAtCap field to given value.

### HasNPredictedAtCap

`func (o *SponsorRecruitmentByPhaseRow) HasNPredictedAtCap() bool`

HasNPredictedAtCap returns a boolean if a field has been set.

### SetNPredictedAtCapNil

`func (o *SponsorRecruitmentByPhaseRow) SetNPredictedAtCapNil(b bool)`

 SetNPredictedAtCapNil sets the value for NPredictedAtCap to be an explicit nil

### UnsetNPredictedAtCap
`func (o *SponsorRecruitmentByPhaseRow) UnsetNPredictedAtCap()`

UnsetNPredictedAtCap ensures that no value is present for NPredictedAtCap, not even an explicit nil
### GetNPredictedOverdue

`func (o *SponsorRecruitmentByPhaseRow) GetNPredictedOverdue() int32`

GetNPredictedOverdue returns the NPredictedOverdue field if non-nil, zero value otherwise.

### GetNPredictedOverdueOk

`func (o *SponsorRecruitmentByPhaseRow) GetNPredictedOverdueOk() (*int32, bool)`

GetNPredictedOverdueOk returns a tuple with the NPredictedOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPredictedOverdue

`func (o *SponsorRecruitmentByPhaseRow) SetNPredictedOverdue(v int32)`

SetNPredictedOverdue sets NPredictedOverdue field to given value.

### HasNPredictedOverdue

`func (o *SponsorRecruitmentByPhaseRow) HasNPredictedOverdue() bool`

HasNPredictedOverdue returns a boolean if a field has been set.

### SetNPredictedOverdueNil

`func (o *SponsorRecruitmentByPhaseRow) SetNPredictedOverdueNil(b bool)`

 SetNPredictedOverdueNil sets the value for NPredictedOverdue to be an explicit nil

### UnsetNPredictedOverdue
`func (o *SponsorRecruitmentByPhaseRow) UnsetNPredictedOverdue()`

UnsetNPredictedOverdue ensures that no value is present for NPredictedOverdue, not even an explicit nil
### GetPredictedMonthsMedian

`func (o *SponsorRecruitmentByPhaseRow) GetPredictedMonthsMedian() float32`

GetPredictedMonthsMedian returns the PredictedMonthsMedian field if non-nil, zero value otherwise.

### GetPredictedMonthsMedianOk

`func (o *SponsorRecruitmentByPhaseRow) GetPredictedMonthsMedianOk() (*float32, bool)`

GetPredictedMonthsMedianOk returns a tuple with the PredictedMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMonthsMedian

`func (o *SponsorRecruitmentByPhaseRow) SetPredictedMonthsMedian(v float32)`

SetPredictedMonthsMedian sets PredictedMonthsMedian field to given value.

### HasPredictedMonthsMedian

`func (o *SponsorRecruitmentByPhaseRow) HasPredictedMonthsMedian() bool`

HasPredictedMonthsMedian returns a boolean if a field has been set.

### SetPredictedMonthsMedianNil

`func (o *SponsorRecruitmentByPhaseRow) SetPredictedMonthsMedianNil(b bool)`

 SetPredictedMonthsMedianNil sets the value for PredictedMonthsMedian to be an explicit nil

### UnsetPredictedMonthsMedian
`func (o *SponsorRecruitmentByPhaseRow) UnsetPredictedMonthsMedian()`

UnsetPredictedMonthsMedian ensures that no value is present for PredictedMonthsMedian, not even an explicit nil
### GetBaselineMonthsMedian

`func (o *SponsorRecruitmentByPhaseRow) GetBaselineMonthsMedian() float32`

GetBaselineMonthsMedian returns the BaselineMonthsMedian field if non-nil, zero value otherwise.

### GetBaselineMonthsMedianOk

`func (o *SponsorRecruitmentByPhaseRow) GetBaselineMonthsMedianOk() (*float32, bool)`

GetBaselineMonthsMedianOk returns a tuple with the BaselineMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineMonthsMedian

`func (o *SponsorRecruitmentByPhaseRow) SetBaselineMonthsMedian(v float32)`

SetBaselineMonthsMedian sets BaselineMonthsMedian field to given value.

### HasBaselineMonthsMedian

`func (o *SponsorRecruitmentByPhaseRow) HasBaselineMonthsMedian() bool`

HasBaselineMonthsMedian returns a boolean if a field has been set.

### SetBaselineMonthsMedianNil

`func (o *SponsorRecruitmentByPhaseRow) SetBaselineMonthsMedianNil(b bool)`

 SetBaselineMonthsMedianNil sets the value for BaselineMonthsMedian to be an explicit nil

### UnsetBaselineMonthsMedian
`func (o *SponsorRecruitmentByPhaseRow) UnsetBaselineMonthsMedian()`

UnsetBaselineMonthsMedian ensures that no value is present for BaselineMonthsMedian, not even an explicit nil
### GetNRealized

`func (o *SponsorRecruitmentByPhaseRow) GetNRealized() int32`

GetNRealized returns the NRealized field if non-nil, zero value otherwise.

### GetNRealizedOk

`func (o *SponsorRecruitmentByPhaseRow) GetNRealizedOk() (*int32, bool)`

GetNRealizedOk returns a tuple with the NRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealized

`func (o *SponsorRecruitmentByPhaseRow) SetNRealized(v int32)`

SetNRealized sets NRealized field to given value.

### HasNRealized

`func (o *SponsorRecruitmentByPhaseRow) HasNRealized() bool`

HasNRealized returns a boolean if a field has been set.

### SetNRealizedNil

`func (o *SponsorRecruitmentByPhaseRow) SetNRealizedNil(b bool)`

 SetNRealizedNil sets the value for NRealized to be an explicit nil

### UnsetNRealized
`func (o *SponsorRecruitmentByPhaseRow) UnsetNRealized()`

UnsetNRealized ensures that no value is present for NRealized, not even an explicit nil
### GetRealizedMonthsMedian

`func (o *SponsorRecruitmentByPhaseRow) GetRealizedMonthsMedian() float32`

GetRealizedMonthsMedian returns the RealizedMonthsMedian field if non-nil, zero value otherwise.

### GetRealizedMonthsMedianOk

`func (o *SponsorRecruitmentByPhaseRow) GetRealizedMonthsMedianOk() (*float32, bool)`

GetRealizedMonthsMedianOk returns a tuple with the RealizedMonthsMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedMonthsMedian

`func (o *SponsorRecruitmentByPhaseRow) SetRealizedMonthsMedian(v float32)`

SetRealizedMonthsMedian sets RealizedMonthsMedian field to given value.

### HasRealizedMonthsMedian

`func (o *SponsorRecruitmentByPhaseRow) HasRealizedMonthsMedian() bool`

HasRealizedMonthsMedian returns a boolean if a field has been set.

### SetRealizedMonthsMedianNil

`func (o *SponsorRecruitmentByPhaseRow) SetRealizedMonthsMedianNil(b bool)`

 SetRealizedMonthsMedianNil sets the value for RealizedMonthsMedian to be an explicit nil

### UnsetRealizedMonthsMedian
`func (o *SponsorRecruitmentByPhaseRow) UnsetRealizedMonthsMedian()`

UnsetRealizedMonthsMedian ensures that no value is present for RealizedMonthsMedian, not even an explicit nil
### GetErrorPhaseMaeMonths

`func (o *SponsorRecruitmentByPhaseRow) GetErrorPhaseMaeMonths() float32`

GetErrorPhaseMaeMonths returns the ErrorPhaseMaeMonths field if non-nil, zero value otherwise.

### GetErrorPhaseMaeMonthsOk

`func (o *SponsorRecruitmentByPhaseRow) GetErrorPhaseMaeMonthsOk() (*float32, bool)`

GetErrorPhaseMaeMonthsOk returns a tuple with the ErrorPhaseMaeMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPhaseMaeMonths

`func (o *SponsorRecruitmentByPhaseRow) SetErrorPhaseMaeMonths(v float32)`

SetErrorPhaseMaeMonths sets ErrorPhaseMaeMonths field to given value.

### HasErrorPhaseMaeMonths

`func (o *SponsorRecruitmentByPhaseRow) HasErrorPhaseMaeMonths() bool`

HasErrorPhaseMaeMonths returns a boolean if a field has been set.

### SetErrorPhaseMaeMonthsNil

`func (o *SponsorRecruitmentByPhaseRow) SetErrorPhaseMaeMonthsNil(b bool)`

 SetErrorPhaseMaeMonthsNil sets the value for ErrorPhaseMaeMonths to be an explicit nil

### UnsetErrorPhaseMaeMonths
`func (o *SponsorRecruitmentByPhaseRow) UnsetErrorPhaseMaeMonths()`

UnsetErrorPhaseMaeMonths ensures that no value is present for ErrorPhaseMaeMonths, not even an explicit nil
### GetErrorPhaseRelmaeVsBaseline

`func (o *SponsorRecruitmentByPhaseRow) GetErrorPhaseRelmaeVsBaseline() float32`

GetErrorPhaseRelmaeVsBaseline returns the ErrorPhaseRelmaeVsBaseline field if non-nil, zero value otherwise.

### GetErrorPhaseRelmaeVsBaselineOk

`func (o *SponsorRecruitmentByPhaseRow) GetErrorPhaseRelmaeVsBaselineOk() (*float32, bool)`

GetErrorPhaseRelmaeVsBaselineOk returns a tuple with the ErrorPhaseRelmaeVsBaseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPhaseRelmaeVsBaseline

`func (o *SponsorRecruitmentByPhaseRow) SetErrorPhaseRelmaeVsBaseline(v float32)`

SetErrorPhaseRelmaeVsBaseline sets ErrorPhaseRelmaeVsBaseline field to given value.

### HasErrorPhaseRelmaeVsBaseline

`func (o *SponsorRecruitmentByPhaseRow) HasErrorPhaseRelmaeVsBaseline() bool`

HasErrorPhaseRelmaeVsBaseline returns a boolean if a field has been set.

### SetErrorPhaseRelmaeVsBaselineNil

`func (o *SponsorRecruitmentByPhaseRow) SetErrorPhaseRelmaeVsBaselineNil(b bool)`

 SetErrorPhaseRelmaeVsBaselineNil sets the value for ErrorPhaseRelmaeVsBaseline to be an explicit nil

### UnsetErrorPhaseRelmaeVsBaseline
`func (o *SponsorRecruitmentByPhaseRow) UnsetErrorPhaseRelmaeVsBaseline()`

UnsetErrorPhaseRelmaeVsBaseline ensures that no value is present for ErrorPhaseRelmaeVsBaseline, not even an explicit nil
### GetErrorPhaseBaselineMaeMonths

`func (o *SponsorRecruitmentByPhaseRow) GetErrorPhaseBaselineMaeMonths() float32`

GetErrorPhaseBaselineMaeMonths returns the ErrorPhaseBaselineMaeMonths field if non-nil, zero value otherwise.

### GetErrorPhaseBaselineMaeMonthsOk

`func (o *SponsorRecruitmentByPhaseRow) GetErrorPhaseBaselineMaeMonthsOk() (*float32, bool)`

GetErrorPhaseBaselineMaeMonthsOk returns a tuple with the ErrorPhaseBaselineMaeMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPhaseBaselineMaeMonths

`func (o *SponsorRecruitmentByPhaseRow) SetErrorPhaseBaselineMaeMonths(v float32)`

SetErrorPhaseBaselineMaeMonths sets ErrorPhaseBaselineMaeMonths field to given value.

### HasErrorPhaseBaselineMaeMonths

`func (o *SponsorRecruitmentByPhaseRow) HasErrorPhaseBaselineMaeMonths() bool`

HasErrorPhaseBaselineMaeMonths returns a boolean if a field has been set.

### SetErrorPhaseBaselineMaeMonthsNil

`func (o *SponsorRecruitmentByPhaseRow) SetErrorPhaseBaselineMaeMonthsNil(b bool)`

 SetErrorPhaseBaselineMaeMonthsNil sets the value for ErrorPhaseBaselineMaeMonths to be an explicit nil

### UnsetErrorPhaseBaselineMaeMonths
`func (o *SponsorRecruitmentByPhaseRow) UnsetErrorPhaseBaselineMaeMonths()`

UnsetErrorPhaseBaselineMaeMonths ensures that no value is present for ErrorPhaseBaselineMaeMonths, not even an explicit nil
### GetErrorPhaseNTrials

`func (o *SponsorRecruitmentByPhaseRow) GetErrorPhaseNTrials() int32`

GetErrorPhaseNTrials returns the ErrorPhaseNTrials field if non-nil, zero value otherwise.

### GetErrorPhaseNTrialsOk

`func (o *SponsorRecruitmentByPhaseRow) GetErrorPhaseNTrialsOk() (*int32, bool)`

GetErrorPhaseNTrialsOk returns a tuple with the ErrorPhaseNTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPhaseNTrials

`func (o *SponsorRecruitmentByPhaseRow) SetErrorPhaseNTrials(v int32)`

SetErrorPhaseNTrials sets ErrorPhaseNTrials field to given value.

### HasErrorPhaseNTrials

`func (o *SponsorRecruitmentByPhaseRow) HasErrorPhaseNTrials() bool`

HasErrorPhaseNTrials returns a boolean if a field has been set.

### SetErrorPhaseNTrialsNil

`func (o *SponsorRecruitmentByPhaseRow) SetErrorPhaseNTrialsNil(b bool)`

 SetErrorPhaseNTrialsNil sets the value for ErrorPhaseNTrials to be an explicit nil

### UnsetErrorPhaseNTrials
`func (o *SponsorRecruitmentByPhaseRow) UnsetErrorPhaseNTrials()`

UnsetErrorPhaseNTrials ensures that no value is present for ErrorPhaseNTrials, not even an explicit nil
### GetErrorPhaseNYears

`func (o *SponsorRecruitmentByPhaseRow) GetErrorPhaseNYears() int32`

GetErrorPhaseNYears returns the ErrorPhaseNYears field if non-nil, zero value otherwise.

### GetErrorPhaseNYearsOk

`func (o *SponsorRecruitmentByPhaseRow) GetErrorPhaseNYearsOk() (*int32, bool)`

GetErrorPhaseNYearsOk returns a tuple with the ErrorPhaseNYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPhaseNYears

`func (o *SponsorRecruitmentByPhaseRow) SetErrorPhaseNYears(v int32)`

SetErrorPhaseNYears sets ErrorPhaseNYears field to given value.

### HasErrorPhaseNYears

`func (o *SponsorRecruitmentByPhaseRow) HasErrorPhaseNYears() bool`

HasErrorPhaseNYears returns a boolean if a field has been set.

### SetErrorPhaseNYearsNil

`func (o *SponsorRecruitmentByPhaseRow) SetErrorPhaseNYearsNil(b bool)`

 SetErrorPhaseNYearsNil sets the value for ErrorPhaseNYears to be an explicit nil

### UnsetErrorPhaseNYears
`func (o *SponsorRecruitmentByPhaseRow) UnsetErrorPhaseNYears()`

UnsetErrorPhaseNYears ensures that no value is present for ErrorPhaseNYears, not even an explicit nil
### GetMinN

`func (o *SponsorRecruitmentByPhaseRow) GetMinN() int32`

GetMinN returns the MinN field if non-nil, zero value otherwise.

### GetMinNOk

`func (o *SponsorRecruitmentByPhaseRow) GetMinNOk() (*int32, bool)`

GetMinNOk returns a tuple with the MinN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinN

`func (o *SponsorRecruitmentByPhaseRow) SetMinN(v int32)`

SetMinN sets MinN field to given value.

### HasMinN

`func (o *SponsorRecruitmentByPhaseRow) HasMinN() bool`

HasMinN returns a boolean if a field has been set.

### SetMinNNil

`func (o *SponsorRecruitmentByPhaseRow) SetMinNNil(b bool)`

 SetMinNNil sets the value for MinN to be an explicit nil

### UnsetMinN
`func (o *SponsorRecruitmentByPhaseRow) UnsetMinN()`

UnsetMinN ensures that no value is present for MinN, not even an explicit nil
### GetTrainingCutoff

`func (o *SponsorRecruitmentByPhaseRow) GetTrainingCutoff() string`

GetTrainingCutoff returns the TrainingCutoff field if non-nil, zero value otherwise.

### GetTrainingCutoffOk

`func (o *SponsorRecruitmentByPhaseRow) GetTrainingCutoffOk() (*string, bool)`

GetTrainingCutoffOk returns a tuple with the TrainingCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingCutoff

`func (o *SponsorRecruitmentByPhaseRow) SetTrainingCutoff(v string)`

SetTrainingCutoff sets TrainingCutoff field to given value.

### HasTrainingCutoff

`func (o *SponsorRecruitmentByPhaseRow) HasTrainingCutoff() bool`

HasTrainingCutoff returns a boolean if a field has been set.

### SetTrainingCutoffNil

`func (o *SponsorRecruitmentByPhaseRow) SetTrainingCutoffNil(b bool)`

 SetTrainingCutoffNil sets the value for TrainingCutoff to be an explicit nil

### UnsetTrainingCutoff
`func (o *SponsorRecruitmentByPhaseRow) UnsetTrainingCutoff()`

UnsetTrainingCutoff ensures that no value is present for TrainingCutoff, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


