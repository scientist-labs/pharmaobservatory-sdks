# ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**NOpen** | **int32** |  | 
**NAtCap** | **int32** |  | 
**NOverdue** | **int32** |  | 
**PredictedMedian** | **interface{}** |  | 
**BaselineMedian** | **interface{}** |  | 
**NRealized** | **int32** |  | 
**RealizedMedian** | **NullableFloat32** |  | 
**Mae** | **NullableFloat32** |  | 
**Relmae** | **NullableFloat32** |  | 
**ErrorNTrials** | **NullableInt32** |  | 
**ErrorNYears** | **NullableInt32** |  | 
**PredictedCell** | **string** |  | 
**RegisteredCell** | **string** |  | 
**Error** | **string** |  | 
**PredictedMedianAtHorizonCap** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner

`func NewApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner(bucket string, nOpen int32, nAtCap int32, nOverdue int32, predictedMedian interface{}, baselineMedian interface{}, nRealized int32, realizedMedian NullableFloat32, mae NullableFloat32, relmae NullableFloat32, errorNTrials NullableInt32, errorNYears NullableInt32, predictedCell string, registeredCell string, error_ string, ) *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner`

NewApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner instantiates a new ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInnerWithDefaults

`func NewApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInnerWithDefaults() *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner`

NewApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInnerWithDefaults instantiates a new ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetNOpen

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetNOpen() int32`

GetNOpen returns the NOpen field if non-nil, zero value otherwise.

### GetNOpenOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetNOpenOk() (*int32, bool)`

GetNOpenOk returns a tuple with the NOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOpen

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetNOpen(v int32)`

SetNOpen sets NOpen field to given value.


### GetNAtCap

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetNAtCap() int32`

GetNAtCap returns the NAtCap field if non-nil, zero value otherwise.

### GetNAtCapOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetNAtCapOk() (*int32, bool)`

GetNAtCapOk returns a tuple with the NAtCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAtCap

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetNAtCap(v int32)`

SetNAtCap sets NAtCap field to given value.


### GetNOverdue

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetNOverdue() int32`

GetNOverdue returns the NOverdue field if non-nil, zero value otherwise.

### GetNOverdueOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetNOverdueOk() (*int32, bool)`

GetNOverdueOk returns a tuple with the NOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOverdue

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetNOverdue(v int32)`

SetNOverdue sets NOverdue field to given value.


### GetPredictedMedian

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetPredictedMedian() interface{}`

GetPredictedMedian returns the PredictedMedian field if non-nil, zero value otherwise.

### GetPredictedMedianOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetPredictedMedianOk() (*interface{}, bool)`

GetPredictedMedianOk returns a tuple with the PredictedMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMedian

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetPredictedMedian(v interface{})`

SetPredictedMedian sets PredictedMedian field to given value.


### SetPredictedMedianNil

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetPredictedMedianNil(b bool)`

 SetPredictedMedianNil sets the value for PredictedMedian to be an explicit nil

### UnsetPredictedMedian
`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) UnsetPredictedMedian()`

UnsetPredictedMedian ensures that no value is present for PredictedMedian, not even an explicit nil
### GetBaselineMedian

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetBaselineMedian() interface{}`

GetBaselineMedian returns the BaselineMedian field if non-nil, zero value otherwise.

### GetBaselineMedianOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetBaselineMedianOk() (*interface{}, bool)`

GetBaselineMedianOk returns a tuple with the BaselineMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineMedian

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetBaselineMedian(v interface{})`

SetBaselineMedian sets BaselineMedian field to given value.


### SetBaselineMedianNil

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetBaselineMedianNil(b bool)`

 SetBaselineMedianNil sets the value for BaselineMedian to be an explicit nil

### UnsetBaselineMedian
`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) UnsetBaselineMedian()`

UnsetBaselineMedian ensures that no value is present for BaselineMedian, not even an explicit nil
### GetNRealized

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetNRealized() int32`

GetNRealized returns the NRealized field if non-nil, zero value otherwise.

### GetNRealizedOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetNRealizedOk() (*int32, bool)`

GetNRealizedOk returns a tuple with the NRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealized

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetNRealized(v int32)`

SetNRealized sets NRealized field to given value.


### GetRealizedMedian

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetRealizedMedian() float32`

GetRealizedMedian returns the RealizedMedian field if non-nil, zero value otherwise.

### GetRealizedMedianOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetRealizedMedianOk() (*float32, bool)`

GetRealizedMedianOk returns a tuple with the RealizedMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedMedian

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetRealizedMedian(v float32)`

SetRealizedMedian sets RealizedMedian field to given value.


### SetRealizedMedianNil

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetRealizedMedianNil(b bool)`

 SetRealizedMedianNil sets the value for RealizedMedian to be an explicit nil

### UnsetRealizedMedian
`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) UnsetRealizedMedian()`

UnsetRealizedMedian ensures that no value is present for RealizedMedian, not even an explicit nil
### GetMae

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetMae() float32`

GetMae returns the Mae field if non-nil, zero value otherwise.

### GetMaeOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetMaeOk() (*float32, bool)`

GetMaeOk returns a tuple with the Mae field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMae

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetMae(v float32)`

SetMae sets Mae field to given value.


### SetMaeNil

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetMaeNil(b bool)`

 SetMaeNil sets the value for Mae to be an explicit nil

### UnsetMae
`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) UnsetMae()`

UnsetMae ensures that no value is present for Mae, not even an explicit nil
### GetRelmae

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetRelmae() float32`

GetRelmae returns the Relmae field if non-nil, zero value otherwise.

### GetRelmaeOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetRelmaeOk() (*float32, bool)`

GetRelmaeOk returns a tuple with the Relmae field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelmae

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetRelmae(v float32)`

SetRelmae sets Relmae field to given value.


### SetRelmaeNil

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetRelmaeNil(b bool)`

 SetRelmaeNil sets the value for Relmae to be an explicit nil

### UnsetRelmae
`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) UnsetRelmae()`

UnsetRelmae ensures that no value is present for Relmae, not even an explicit nil
### GetErrorNTrials

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetErrorNTrials() int32`

GetErrorNTrials returns the ErrorNTrials field if non-nil, zero value otherwise.

### GetErrorNTrialsOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetErrorNTrialsOk() (*int32, bool)`

GetErrorNTrialsOk returns a tuple with the ErrorNTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNTrials

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetErrorNTrials(v int32)`

SetErrorNTrials sets ErrorNTrials field to given value.


### SetErrorNTrialsNil

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetErrorNTrialsNil(b bool)`

 SetErrorNTrialsNil sets the value for ErrorNTrials to be an explicit nil

### UnsetErrorNTrials
`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) UnsetErrorNTrials()`

UnsetErrorNTrials ensures that no value is present for ErrorNTrials, not even an explicit nil
### GetErrorNYears

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetErrorNYears() int32`

GetErrorNYears returns the ErrorNYears field if non-nil, zero value otherwise.

### GetErrorNYearsOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetErrorNYearsOk() (*int32, bool)`

GetErrorNYearsOk returns a tuple with the ErrorNYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNYears

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetErrorNYears(v int32)`

SetErrorNYears sets ErrorNYears field to given value.


### SetErrorNYearsNil

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetErrorNYearsNil(b bool)`

 SetErrorNYearsNil sets the value for ErrorNYears to be an explicit nil

### UnsetErrorNYears
`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) UnsetErrorNYears()`

UnsetErrorNYears ensures that no value is present for ErrorNYears, not even an explicit nil
### GetPredictedCell

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetPredictedCell() string`

GetPredictedCell returns the PredictedCell field if non-nil, zero value otherwise.

### GetPredictedCellOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetPredictedCellOk() (*string, bool)`

GetPredictedCellOk returns a tuple with the PredictedCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedCell

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetPredictedCell(v string)`

SetPredictedCell sets PredictedCell field to given value.


### GetRegisteredCell

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetRegisteredCell() string`

GetRegisteredCell returns the RegisteredCell field if non-nil, zero value otherwise.

### GetRegisteredCellOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetRegisteredCellOk() (*string, bool)`

GetRegisteredCellOk returns a tuple with the RegisteredCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredCell

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetRegisteredCell(v string)`

SetRegisteredCell sets RegisteredCell field to given value.


### GetError

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetError(v string)`

SetError sets Error field to given value.


### GetPredictedMedianAtHorizonCap

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetPredictedMedianAtHorizonCap() bool`

GetPredictedMedianAtHorizonCap returns the PredictedMedianAtHorizonCap field if non-nil, zero value otherwise.

### GetPredictedMedianAtHorizonCapOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) GetPredictedMedianAtHorizonCapOk() (*bool, bool)`

GetPredictedMedianAtHorizonCapOk returns a tuple with the PredictedMedianAtHorizonCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMedianAtHorizonCap

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) SetPredictedMedianAtHorizonCap(v bool)`

SetPredictedMedianAtHorizonCap sets PredictedMedianAtHorizonCap field to given value.

### HasPredictedMedianAtHorizonCap

`func (o *ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner) HasPredictedMedianAtHorizonCap() bool`

HasPredictedMedianAtHorizonCap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


