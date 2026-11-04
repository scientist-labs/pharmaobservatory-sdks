# SubscriptionsCreate201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**SubscriptionsCreate201ResponseSubscription**](SubscriptionsCreate201ResponseSubscription.md) |  | 
**Created** | **bool** |  | 
**BuildState** | [**ChangesList200ResponseBuildState**](ChangesList200ResponseBuildState.md) |  | 

## Methods

### NewSubscriptionsCreate201Response

`func NewSubscriptionsCreate201Response(subscription SubscriptionsCreate201ResponseSubscription, created bool, buildState ChangesList200ResponseBuildState, ) *SubscriptionsCreate201Response`

NewSubscriptionsCreate201Response instantiates a new SubscriptionsCreate201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsCreate201ResponseWithDefaults

`func NewSubscriptionsCreate201ResponseWithDefaults() *SubscriptionsCreate201Response`

NewSubscriptionsCreate201ResponseWithDefaults instantiates a new SubscriptionsCreate201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *SubscriptionsCreate201Response) GetSubscription() SubscriptionsCreate201ResponseSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SubscriptionsCreate201Response) GetSubscriptionOk() (*SubscriptionsCreate201ResponseSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SubscriptionsCreate201Response) SetSubscription(v SubscriptionsCreate201ResponseSubscription)`

SetSubscription sets Subscription field to given value.


### GetCreated

`func (o *SubscriptionsCreate201Response) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubscriptionsCreate201Response) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubscriptionsCreate201Response) SetCreated(v bool)`

SetCreated sets Created field to given value.


### GetBuildState

`func (o *SubscriptionsCreate201Response) GetBuildState() ChangesList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *SubscriptionsCreate201Response) GetBuildStateOk() (*ChangesList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *SubscriptionsCreate201Response) SetBuildState(v ChangesList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


