# ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**Measured** | **bool** |  | 
**NTrials** | **NullableInt32** |  | 
**NYears** | **NullableInt32** |  | 
**MaeMonths** | **NullableFloat32** |  | 
**RelmaeVsBaseline** | **NullableFloat32** |  | 
**BaselineMaeMonths** | **NullableFloat32** |  | 
**Comparator** | **string** |  | 
**Statement** | **string** |  | 

## Methods

### NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize

`func NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize(bucket string, measured bool, nTrials NullableInt32, nYears NullableInt32, maeMonths NullableFloat32, relmaeVsBaseline NullableFloat32, baselineMaeMonths NullableFloat32, comparator string, statement string, ) *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize`

NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize instantiates a new ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSizeWithDefaults

`func NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSizeWithDefaults() *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize`

NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSizeWithDefaults instantiates a new ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetMeasured

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetMeasured() bool`

GetMeasured returns the Measured field if non-nil, zero value otherwise.

### GetMeasuredOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetMeasuredOk() (*bool, bool)`

GetMeasuredOk returns a tuple with the Measured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasured

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetMeasured(v bool)`

SetMeasured sets Measured field to given value.


### GetNTrials

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.


### SetNTrialsNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNYears

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetNYears() int32`

GetNYears returns the NYears field if non-nil, zero value otherwise.

### GetNYearsOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetNYearsOk() (*int32, bool)`

GetNYearsOk returns a tuple with the NYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNYears

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetNYears(v int32)`

SetNYears sets NYears field to given value.


### SetNYearsNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetNYearsNil(b bool)`

 SetNYearsNil sets the value for NYears to be an explicit nil

### UnsetNYears
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) UnsetNYears()`

UnsetNYears ensures that no value is present for NYears, not even an explicit nil
### GetMaeMonths

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetMaeMonths() float32`

GetMaeMonths returns the MaeMonths field if non-nil, zero value otherwise.

### GetMaeMonthsOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetMaeMonthsOk() (*float32, bool)`

GetMaeMonthsOk returns a tuple with the MaeMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaeMonths

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetMaeMonths(v float32)`

SetMaeMonths sets MaeMonths field to given value.


### SetMaeMonthsNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetMaeMonthsNil(b bool)`

 SetMaeMonthsNil sets the value for MaeMonths to be an explicit nil

### UnsetMaeMonths
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) UnsetMaeMonths()`

UnsetMaeMonths ensures that no value is present for MaeMonths, not even an explicit nil
### GetRelmaeVsBaseline

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetRelmaeVsBaseline() float32`

GetRelmaeVsBaseline returns the RelmaeVsBaseline field if non-nil, zero value otherwise.

### GetRelmaeVsBaselineOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetRelmaeVsBaselineOk() (*float32, bool)`

GetRelmaeVsBaselineOk returns a tuple with the RelmaeVsBaseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelmaeVsBaseline

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetRelmaeVsBaseline(v float32)`

SetRelmaeVsBaseline sets RelmaeVsBaseline field to given value.


### SetRelmaeVsBaselineNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetRelmaeVsBaselineNil(b bool)`

 SetRelmaeVsBaselineNil sets the value for RelmaeVsBaseline to be an explicit nil

### UnsetRelmaeVsBaseline
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) UnsetRelmaeVsBaseline()`

UnsetRelmaeVsBaseline ensures that no value is present for RelmaeVsBaseline, not even an explicit nil
### GetBaselineMaeMonths

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetBaselineMaeMonths() float32`

GetBaselineMaeMonths returns the BaselineMaeMonths field if non-nil, zero value otherwise.

### GetBaselineMaeMonthsOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetBaselineMaeMonthsOk() (*float32, bool)`

GetBaselineMaeMonthsOk returns a tuple with the BaselineMaeMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineMaeMonths

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetBaselineMaeMonths(v float32)`

SetBaselineMaeMonths sets BaselineMaeMonths field to given value.


### SetBaselineMaeMonthsNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetBaselineMaeMonthsNil(b bool)`

 SetBaselineMaeMonthsNil sets the value for BaselineMaeMonths to be an explicit nil

### UnsetBaselineMaeMonths
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) UnsetBaselineMaeMonths()`

UnsetBaselineMaeMonths ensures that no value is present for BaselineMaeMonths, not even an explicit nil
### GetComparator

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetComparator() string`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetComparatorOk() (*string, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetComparator(v string)`

SetComparator sets Comparator field to given value.


### GetStatement

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredErrorSize) SetStatement(v string)`

SetStatement sets Statement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


