# SubscriptionsList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | [**[]SubscriptionsList200ResponseSubscriptionsInner**](SubscriptionsList200ResponseSubscriptionsInner.md) |  | 
**Count** | **int32** |  | 
**BuildState** | [**ChangesList200ResponseBuildState**](ChangesList200ResponseBuildState.md) |  | 

## Methods

### NewSubscriptionsList200Response

`func NewSubscriptionsList200Response(subscriptions []SubscriptionsList200ResponseSubscriptionsInner, count int32, buildState ChangesList200ResponseBuildState, ) *SubscriptionsList200Response`

NewSubscriptionsList200Response instantiates a new SubscriptionsList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsList200ResponseWithDefaults

`func NewSubscriptionsList200ResponseWithDefaults() *SubscriptionsList200Response`

NewSubscriptionsList200ResponseWithDefaults instantiates a new SubscriptionsList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *SubscriptionsList200Response) GetSubscriptions() []SubscriptionsList200ResponseSubscriptionsInner`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *SubscriptionsList200Response) GetSubscriptionsOk() (*[]SubscriptionsList200ResponseSubscriptionsInner, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *SubscriptionsList200Response) SetSubscriptions(v []SubscriptionsList200ResponseSubscriptionsInner)`

SetSubscriptions sets Subscriptions field to given value.


### GetCount

`func (o *SubscriptionsList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SubscriptionsList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SubscriptionsList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetBuildState

`func (o *SubscriptionsList200Response) GetBuildState() ChangesList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *SubscriptionsList200Response) GetBuildStateOk() (*ChangesList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *SubscriptionsList200Response) SetBuildState(v ChangesList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


