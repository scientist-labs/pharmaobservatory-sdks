# ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**State** | **string** |  | 
**Drawn** | **bool** |  | 
**Cutoff** | **string** |  | 
**Threshold** | **int32** |  | 
**StartsPerYear** | **string** |  | 
**Bucket** | **string** |  | 
**Years** | [**[]ApiV1OrgPipelineStagesGet200ResponseTrialStartForecastYearsInner**](ApiV1OrgPipelineStagesGet200ResponseTrialStartForecastYearsInner.md) |  | 

## Methods

### NewApiV1OrgPipelineStagesGet200ResponseTrialStartForecast

`func NewApiV1OrgPipelineStagesGet200ResponseTrialStartForecast(kind string, state string, drawn bool, cutoff string, threshold int32, startsPerYear string, bucket string, years []ApiV1OrgPipelineStagesGet200ResponseTrialStartForecastYearsInner, ) *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast`

NewApiV1OrgPipelineStagesGet200ResponseTrialStartForecast instantiates a new ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1OrgPipelineStagesGet200ResponseTrialStartForecastWithDefaults

`func NewApiV1OrgPipelineStagesGet200ResponseTrialStartForecastWithDefaults() *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast`

NewApiV1OrgPipelineStagesGet200ResponseTrialStartForecastWithDefaults instantiates a new ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) SetKind(v string)`

SetKind sets Kind field to given value.


### GetState

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) SetState(v string)`

SetState sets State field to given value.


### GetDrawn

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetDrawn() bool`

GetDrawn returns the Drawn field if non-nil, zero value otherwise.

### GetDrawnOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetDrawnOk() (*bool, bool)`

GetDrawnOk returns a tuple with the Drawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawn

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) SetDrawn(v bool)`

SetDrawn sets Drawn field to given value.


### GetCutoff

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetCutoff() string`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetCutoffOk() (*string, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoff

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) SetCutoff(v string)`

SetCutoff sets Cutoff field to given value.


### GetThreshold

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetStartsPerYear

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetStartsPerYear() string`

GetStartsPerYear returns the StartsPerYear field if non-nil, zero value otherwise.

### GetStartsPerYearOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetStartsPerYearOk() (*string, bool)`

GetStartsPerYearOk returns a tuple with the StartsPerYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsPerYear

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) SetStartsPerYear(v string)`

SetStartsPerYear sets StartsPerYear field to given value.


### GetBucket

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetYears

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetYears() []ApiV1OrgPipelineStagesGet200ResponseTrialStartForecastYearsInner`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) GetYearsOk() (*[]ApiV1OrgPipelineStagesGet200ResponseTrialStartForecastYearsInner, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast) SetYears(v []ApiV1OrgPipelineStagesGet200ResponseTrialStartForecastYearsInner)`

SetYears sets Years field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


