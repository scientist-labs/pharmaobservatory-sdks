# OrgsOrgRecruitment200ResponseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**RealizedState** | **string** |  | 
**Kind** | **string** |  | 
**Cutoff** | **string** |  | 
**MinN** | **int32** |  | 
**Cap** | **int32** |  | 
**ModelVersion** | **string** |  | 
**Method** | **string** |  | 
**NTrials** | **int32** |  | 
**NEvaluated** | **int32** |  | 
**NNotEvaluated** | **int32** |  | 
**Statuses** | [**OrgsOrgRecruitment200ResponseContextStatuses**](OrgsOrgRecruitment200ResponseContextStatuses.md) |  | 
**NSuperseded** | **int32** |  | 
**NOpen** | **int32** |  | 
**NAtCap** | **int32** |  | 
**NOverdue** | **int32** |  | 
**NErrorMeasured** | **int32** |  | 
**Predicted** | [**OrgsOrgRecruitment200ResponseContextPredicted**](OrgsOrgRecruitment200ResponseContextPredicted.md) |  | 
**NRealized** | **int32** |  | 
**NRealizedOverCap** | **int32** |  | 
**Realized** | [**OrgsOrgRecruitment200ResponseContextRealized**](OrgsOrgRecruitment200ResponseContextRealized.md) |  | 

## Methods

### NewOrgsOrgRecruitment200ResponseContext

`func NewOrgsOrgRecruitment200ResponseContext(state string, realizedState string, kind string, cutoff string, minN int32, cap int32, modelVersion string, method string, nTrials int32, nEvaluated int32, nNotEvaluated int32, statuses OrgsOrgRecruitment200ResponseContextStatuses, nSuperseded int32, nOpen int32, nAtCap int32, nOverdue int32, nErrorMeasured int32, predicted OrgsOrgRecruitment200ResponseContextPredicted, nRealized int32, nRealizedOverCap int32, realized OrgsOrgRecruitment200ResponseContextRealized, ) *OrgsOrgRecruitment200ResponseContext`

NewOrgsOrgRecruitment200ResponseContext instantiates a new OrgsOrgRecruitment200ResponseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsOrgRecruitment200ResponseContextWithDefaults

`func NewOrgsOrgRecruitment200ResponseContextWithDefaults() *OrgsOrgRecruitment200ResponseContext`

NewOrgsOrgRecruitment200ResponseContextWithDefaults instantiates a new OrgsOrgRecruitment200ResponseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *OrgsOrgRecruitment200ResponseContext) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrgsOrgRecruitment200ResponseContext) SetState(v string)`

SetState sets State field to given value.


### GetRealizedState

`func (o *OrgsOrgRecruitment200ResponseContext) GetRealizedState() string`

GetRealizedState returns the RealizedState field if non-nil, zero value otherwise.

### GetRealizedStateOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetRealizedStateOk() (*string, bool)`

GetRealizedStateOk returns a tuple with the RealizedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedState

`func (o *OrgsOrgRecruitment200ResponseContext) SetRealizedState(v string)`

SetRealizedState sets RealizedState field to given value.


### GetKind

`func (o *OrgsOrgRecruitment200ResponseContext) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrgsOrgRecruitment200ResponseContext) SetKind(v string)`

SetKind sets Kind field to given value.


### GetCutoff

`func (o *OrgsOrgRecruitment200ResponseContext) GetCutoff() string`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetCutoffOk() (*string, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoff

`func (o *OrgsOrgRecruitment200ResponseContext) SetCutoff(v string)`

SetCutoff sets Cutoff field to given value.


### GetMinN

`func (o *OrgsOrgRecruitment200ResponseContext) GetMinN() int32`

GetMinN returns the MinN field if non-nil, zero value otherwise.

### GetMinNOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetMinNOk() (*int32, bool)`

GetMinNOk returns a tuple with the MinN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinN

`func (o *OrgsOrgRecruitment200ResponseContext) SetMinN(v int32)`

SetMinN sets MinN field to given value.


### GetCap

`func (o *OrgsOrgRecruitment200ResponseContext) GetCap() int32`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetCapOk() (*int32, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *OrgsOrgRecruitment200ResponseContext) SetCap(v int32)`

SetCap sets Cap field to given value.


### GetModelVersion

`func (o *OrgsOrgRecruitment200ResponseContext) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OrgsOrgRecruitment200ResponseContext) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.


### GetMethod

`func (o *OrgsOrgRecruitment200ResponseContext) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OrgsOrgRecruitment200ResponseContext) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetNTrials

`func (o *OrgsOrgRecruitment200ResponseContext) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *OrgsOrgRecruitment200ResponseContext) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.


### GetNEvaluated

`func (o *OrgsOrgRecruitment200ResponseContext) GetNEvaluated() int32`

GetNEvaluated returns the NEvaluated field if non-nil, zero value otherwise.

### GetNEvaluatedOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetNEvaluatedOk() (*int32, bool)`

GetNEvaluatedOk returns a tuple with the NEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEvaluated

`func (o *OrgsOrgRecruitment200ResponseContext) SetNEvaluated(v int32)`

SetNEvaluated sets NEvaluated field to given value.


### GetNNotEvaluated

`func (o *OrgsOrgRecruitment200ResponseContext) GetNNotEvaluated() int32`

GetNNotEvaluated returns the NNotEvaluated field if non-nil, zero value otherwise.

### GetNNotEvaluatedOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetNNotEvaluatedOk() (*int32, bool)`

GetNNotEvaluatedOk returns a tuple with the NNotEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNNotEvaluated

`func (o *OrgsOrgRecruitment200ResponseContext) SetNNotEvaluated(v int32)`

SetNNotEvaluated sets NNotEvaluated field to given value.


### GetStatuses

`func (o *OrgsOrgRecruitment200ResponseContext) GetStatuses() OrgsOrgRecruitment200ResponseContextStatuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetStatusesOk() (*OrgsOrgRecruitment200ResponseContextStatuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *OrgsOrgRecruitment200ResponseContext) SetStatuses(v OrgsOrgRecruitment200ResponseContextStatuses)`

SetStatuses sets Statuses field to given value.


### GetNSuperseded

`func (o *OrgsOrgRecruitment200ResponseContext) GetNSuperseded() int32`

GetNSuperseded returns the NSuperseded field if non-nil, zero value otherwise.

### GetNSupersededOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetNSupersededOk() (*int32, bool)`

GetNSupersededOk returns a tuple with the NSuperseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSuperseded

`func (o *OrgsOrgRecruitment200ResponseContext) SetNSuperseded(v int32)`

SetNSuperseded sets NSuperseded field to given value.


### GetNOpen

`func (o *OrgsOrgRecruitment200ResponseContext) GetNOpen() int32`

GetNOpen returns the NOpen field if non-nil, zero value otherwise.

### GetNOpenOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetNOpenOk() (*int32, bool)`

GetNOpenOk returns a tuple with the NOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOpen

`func (o *OrgsOrgRecruitment200ResponseContext) SetNOpen(v int32)`

SetNOpen sets NOpen field to given value.


### GetNAtCap

`func (o *OrgsOrgRecruitment200ResponseContext) GetNAtCap() int32`

GetNAtCap returns the NAtCap field if non-nil, zero value otherwise.

### GetNAtCapOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetNAtCapOk() (*int32, bool)`

GetNAtCapOk returns a tuple with the NAtCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAtCap

`func (o *OrgsOrgRecruitment200ResponseContext) SetNAtCap(v int32)`

SetNAtCap sets NAtCap field to given value.


### GetNOverdue

`func (o *OrgsOrgRecruitment200ResponseContext) GetNOverdue() int32`

GetNOverdue returns the NOverdue field if non-nil, zero value otherwise.

### GetNOverdueOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetNOverdueOk() (*int32, bool)`

GetNOverdueOk returns a tuple with the NOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOverdue

`func (o *OrgsOrgRecruitment200ResponseContext) SetNOverdue(v int32)`

SetNOverdue sets NOverdue field to given value.


### GetNErrorMeasured

`func (o *OrgsOrgRecruitment200ResponseContext) GetNErrorMeasured() int32`

GetNErrorMeasured returns the NErrorMeasured field if non-nil, zero value otherwise.

### GetNErrorMeasuredOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetNErrorMeasuredOk() (*int32, bool)`

GetNErrorMeasuredOk returns a tuple with the NErrorMeasured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNErrorMeasured

`func (o *OrgsOrgRecruitment200ResponseContext) SetNErrorMeasured(v int32)`

SetNErrorMeasured sets NErrorMeasured field to given value.


### GetPredicted

`func (o *OrgsOrgRecruitment200ResponseContext) GetPredicted() OrgsOrgRecruitment200ResponseContextPredicted`

GetPredicted returns the Predicted field if non-nil, zero value otherwise.

### GetPredictedOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetPredictedOk() (*OrgsOrgRecruitment200ResponseContextPredicted, bool)`

GetPredictedOk returns a tuple with the Predicted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicted

`func (o *OrgsOrgRecruitment200ResponseContext) SetPredicted(v OrgsOrgRecruitment200ResponseContextPredicted)`

SetPredicted sets Predicted field to given value.


### GetNRealized

`func (o *OrgsOrgRecruitment200ResponseContext) GetNRealized() int32`

GetNRealized returns the NRealized field if non-nil, zero value otherwise.

### GetNRealizedOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetNRealizedOk() (*int32, bool)`

GetNRealizedOk returns a tuple with the NRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealized

`func (o *OrgsOrgRecruitment200ResponseContext) SetNRealized(v int32)`

SetNRealized sets NRealized field to given value.


### GetNRealizedOverCap

`func (o *OrgsOrgRecruitment200ResponseContext) GetNRealizedOverCap() int32`

GetNRealizedOverCap returns the NRealizedOverCap field if non-nil, zero value otherwise.

### GetNRealizedOverCapOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetNRealizedOverCapOk() (*int32, bool)`

GetNRealizedOverCapOk returns a tuple with the NRealizedOverCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRealizedOverCap

`func (o *OrgsOrgRecruitment200ResponseContext) SetNRealizedOverCap(v int32)`

SetNRealizedOverCap sets NRealizedOverCap field to given value.


### GetRealized

`func (o *OrgsOrgRecruitment200ResponseContext) GetRealized() OrgsOrgRecruitment200ResponseContextRealized`

GetRealized returns the Realized field if non-nil, zero value otherwise.

### GetRealizedOk

`func (o *OrgsOrgRecruitment200ResponseContext) GetRealizedOk() (*OrgsOrgRecruitment200ResponseContextRealized, bool)`

GetRealizedOk returns a tuple with the Realized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealized

`func (o *OrgsOrgRecruitment200ResponseContext) SetRealized(v OrgsOrgRecruitment200ResponseContextRealized)`

SetRealized sets Realized field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


