# ApiV1InvestorProfileGet200ResponseFunds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rollup** | [**[]ApiV1InvestorProfileGet200ResponseFundsRollupInner**](ApiV1InvestorProfileGet200ResponseFundsRollupInner.md) |  | 
**RollupTotal** | **int32** |  | 

## Methods

### NewApiV1InvestorProfileGet200ResponseFunds

`func NewApiV1InvestorProfileGet200ResponseFunds(rollup []ApiV1InvestorProfileGet200ResponseFundsRollupInner, rollupTotal int32, ) *ApiV1InvestorProfileGet200ResponseFunds`

NewApiV1InvestorProfileGet200ResponseFunds instantiates a new ApiV1InvestorProfileGet200ResponseFunds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1InvestorProfileGet200ResponseFundsWithDefaults

`func NewApiV1InvestorProfileGet200ResponseFundsWithDefaults() *ApiV1InvestorProfileGet200ResponseFunds`

NewApiV1InvestorProfileGet200ResponseFundsWithDefaults instantiates a new ApiV1InvestorProfileGet200ResponseFunds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollup

`func (o *ApiV1InvestorProfileGet200ResponseFunds) GetRollup() []ApiV1InvestorProfileGet200ResponseFundsRollupInner`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *ApiV1InvestorProfileGet200ResponseFunds) GetRollupOk() (*[]ApiV1InvestorProfileGet200ResponseFundsRollupInner, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *ApiV1InvestorProfileGet200ResponseFunds) SetRollup(v []ApiV1InvestorProfileGet200ResponseFundsRollupInner)`

SetRollup sets Rollup field to given value.


### GetRollupTotal

`func (o *ApiV1InvestorProfileGet200ResponseFunds) GetRollupTotal() int32`

GetRollupTotal returns the RollupTotal field if non-nil, zero value otherwise.

### GetRollupTotalOk

`func (o *ApiV1InvestorProfileGet200ResponseFunds) GetRollupTotalOk() (*int32, bool)`

GetRollupTotalOk returns a tuple with the RollupTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupTotal

`func (o *ApiV1InvestorProfileGet200ResponseFunds) SetRollupTotal(v int32)`

SetRollupTotal sets RollupTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


