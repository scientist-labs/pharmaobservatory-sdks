# ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | [**NullableApiV1OrgFinancingAndOwnershipGet200ResponseFinancingSummary**](ApiV1OrgFinancingAndOwnershipGet200ResponseFinancingSummary.md) |  | 
**Rounds** | [**[]ApiV1OrgFinancingAndOwnershipGet200ResponseFinancingRoundsInner**](ApiV1OrgFinancingAndOwnershipGet200ResponseFinancingRoundsInner.md) |  | 
**RoundsTotal** | **int32** |  | 
**RoundsTruncated** | **bool** |  | 

## Methods

### NewApiV1OrgFinancingAndOwnershipGet200ResponseFinancing

`func NewApiV1OrgFinancingAndOwnershipGet200ResponseFinancing(summary NullableApiV1OrgFinancingAndOwnershipGet200ResponseFinancingSummary, rounds []ApiV1OrgFinancingAndOwnershipGet200ResponseFinancingRoundsInner, roundsTotal int32, roundsTruncated bool, ) *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing`

NewApiV1OrgFinancingAndOwnershipGet200ResponseFinancing instantiates a new ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1OrgFinancingAndOwnershipGet200ResponseFinancingWithDefaults

`func NewApiV1OrgFinancingAndOwnershipGet200ResponseFinancingWithDefaults() *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing`

NewApiV1OrgFinancingAndOwnershipGet200ResponseFinancingWithDefaults instantiates a new ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) GetSummary() ApiV1OrgFinancingAndOwnershipGet200ResponseFinancingSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) GetSummaryOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseFinancingSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) SetSummary(v ApiV1OrgFinancingAndOwnershipGet200ResponseFinancingSummary)`

SetSummary sets Summary field to given value.


### SetSummaryNil

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetRounds

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) GetRounds() []ApiV1OrgFinancingAndOwnershipGet200ResponseFinancingRoundsInner`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) GetRoundsOk() (*[]ApiV1OrgFinancingAndOwnershipGet200ResponseFinancingRoundsInner, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) SetRounds(v []ApiV1OrgFinancingAndOwnershipGet200ResponseFinancingRoundsInner)`

SetRounds sets Rounds field to given value.


### GetRoundsTotal

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) GetRoundsTotal() int32`

GetRoundsTotal returns the RoundsTotal field if non-nil, zero value otherwise.

### GetRoundsTotalOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) GetRoundsTotalOk() (*int32, bool)`

GetRoundsTotalOk returns a tuple with the RoundsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundsTotal

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) SetRoundsTotal(v int32)`

SetRoundsTotal sets RoundsTotal field to given value.


### GetRoundsTruncated

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) GetRoundsTruncated() bool`

GetRoundsTruncated returns the RoundsTruncated field if non-nil, zero value otherwise.

### GetRoundsTruncatedOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) GetRoundsTruncatedOk() (*bool, bool)`

GetRoundsTruncatedOk returns a tuple with the RoundsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundsTruncated

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing) SetRoundsTruncated(v bool)`

SetRoundsTruncated sets RoundsTruncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


