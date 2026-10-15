# ApiV1InvestorProfileGet200ResponseOwnership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**View** | **string** |  | 
**Owners** | [**ApiV1InvestorProfileGet200ResponseOwnershipOwners**](ApiV1InvestorProfileGet200ResponseOwnershipOwners.md) |  | 
**Holdings** | [**ApiV1InvestorProfileGet200ResponseOwnershipHoldings**](ApiV1InvestorProfileGet200ResponseOwnershipHoldings.md) |  | 

## Methods

### NewApiV1InvestorProfileGet200ResponseOwnership

`func NewApiV1InvestorProfileGet200ResponseOwnership(view string, owners ApiV1InvestorProfileGet200ResponseOwnershipOwners, holdings ApiV1InvestorProfileGet200ResponseOwnershipHoldings, ) *ApiV1InvestorProfileGet200ResponseOwnership`

NewApiV1InvestorProfileGet200ResponseOwnership instantiates a new ApiV1InvestorProfileGet200ResponseOwnership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1InvestorProfileGet200ResponseOwnershipWithDefaults

`func NewApiV1InvestorProfileGet200ResponseOwnershipWithDefaults() *ApiV1InvestorProfileGet200ResponseOwnership`

NewApiV1InvestorProfileGet200ResponseOwnershipWithDefaults instantiates a new ApiV1InvestorProfileGet200ResponseOwnership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetView

`func (o *ApiV1InvestorProfileGet200ResponseOwnership) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ApiV1InvestorProfileGet200ResponseOwnership) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ApiV1InvestorProfileGet200ResponseOwnership) SetView(v string)`

SetView sets View field to given value.


### GetOwners

`func (o *ApiV1InvestorProfileGet200ResponseOwnership) GetOwners() ApiV1InvestorProfileGet200ResponseOwnershipOwners`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ApiV1InvestorProfileGet200ResponseOwnership) GetOwnersOk() (*ApiV1InvestorProfileGet200ResponseOwnershipOwners, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ApiV1InvestorProfileGet200ResponseOwnership) SetOwners(v ApiV1InvestorProfileGet200ResponseOwnershipOwners)`

SetOwners sets Owners field to given value.


### GetHoldings

`func (o *ApiV1InvestorProfileGet200ResponseOwnership) GetHoldings() ApiV1InvestorProfileGet200ResponseOwnershipHoldings`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *ApiV1InvestorProfileGet200ResponseOwnership) GetHoldingsOk() (*ApiV1InvestorProfileGet200ResponseOwnershipHoldings, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *ApiV1InvestorProfileGet200ResponseOwnership) SetHoldings(v ApiV1InvestorProfileGet200ResponseOwnershipHoldings)`

SetHoldings sets Holdings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


