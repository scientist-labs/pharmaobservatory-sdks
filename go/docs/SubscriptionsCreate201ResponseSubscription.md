# SubscriptionsCreate201ResponseSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Label** | **string** |  | 
**Target** | [**SubscriptionsCreate201ResponseSubscriptionTarget**](SubscriptionsCreate201ResponseSubscriptionTarget.md) |  | 
**EventTypes** | **[]string** |  | 
**Paused** | **bool** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewSubscriptionsCreate201ResponseSubscription

`func NewSubscriptionsCreate201ResponseSubscription(id int32, label string, target SubscriptionsCreate201ResponseSubscriptionTarget, eventTypes []string, paused bool, createdAt string, ) *SubscriptionsCreate201ResponseSubscription`

NewSubscriptionsCreate201ResponseSubscription instantiates a new SubscriptionsCreate201ResponseSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsCreate201ResponseSubscriptionWithDefaults

`func NewSubscriptionsCreate201ResponseSubscriptionWithDefaults() *SubscriptionsCreate201ResponseSubscription`

NewSubscriptionsCreate201ResponseSubscriptionWithDefaults instantiates a new SubscriptionsCreate201ResponseSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionsCreate201ResponseSubscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionsCreate201ResponseSubscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionsCreate201ResponseSubscription) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *SubscriptionsCreate201ResponseSubscription) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SubscriptionsCreate201ResponseSubscription) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SubscriptionsCreate201ResponseSubscription) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTarget

`func (o *SubscriptionsCreate201ResponseSubscription) GetTarget() SubscriptionsCreate201ResponseSubscriptionTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SubscriptionsCreate201ResponseSubscription) GetTargetOk() (*SubscriptionsCreate201ResponseSubscriptionTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SubscriptionsCreate201ResponseSubscription) SetTarget(v SubscriptionsCreate201ResponseSubscriptionTarget)`

SetTarget sets Target field to given value.


### GetEventTypes

`func (o *SubscriptionsCreate201ResponseSubscription) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *SubscriptionsCreate201ResponseSubscription) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *SubscriptionsCreate201ResponseSubscription) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetPaused

`func (o *SubscriptionsCreate201ResponseSubscription) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *SubscriptionsCreate201ResponseSubscription) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *SubscriptionsCreate201ResponseSubscription) SetPaused(v bool)`

SetPaused sets Paused field to given value.


### GetCreatedAt

`func (o *SubscriptionsCreate201ResponseSubscription) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionsCreate201ResponseSubscription) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionsCreate201ResponseSubscription) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


