# SubscriptionsList200ResponseSubscriptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Label** | **string** |  | 
**Target** | [**SubscriptionsList200ResponseSubscriptionsInnerTarget**](SubscriptionsList200ResponseSubscriptionsInnerTarget.md) |  | 
**EventTypes** | **[]string** |  | 
**Paused** | **bool** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewSubscriptionsList200ResponseSubscriptionsInner

`func NewSubscriptionsList200ResponseSubscriptionsInner(id int32, label string, target SubscriptionsList200ResponseSubscriptionsInnerTarget, eventTypes []string, paused bool, createdAt string, ) *SubscriptionsList200ResponseSubscriptionsInner`

NewSubscriptionsList200ResponseSubscriptionsInner instantiates a new SubscriptionsList200ResponseSubscriptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsList200ResponseSubscriptionsInnerWithDefaults

`func NewSubscriptionsList200ResponseSubscriptionsInnerWithDefaults() *SubscriptionsList200ResponseSubscriptionsInner`

NewSubscriptionsList200ResponseSubscriptionsInnerWithDefaults instantiates a new SubscriptionsList200ResponseSubscriptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionsList200ResponseSubscriptionsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SubscriptionsList200ResponseSubscriptionsInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTarget

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetTarget() SubscriptionsList200ResponseSubscriptionsInnerTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetTargetOk() (*SubscriptionsList200ResponseSubscriptionsInnerTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SubscriptionsList200ResponseSubscriptionsInner) SetTarget(v SubscriptionsList200ResponseSubscriptionsInnerTarget)`

SetTarget sets Target field to given value.


### GetEventTypes

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *SubscriptionsList200ResponseSubscriptionsInner) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetPaused

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *SubscriptionsList200ResponseSubscriptionsInner) SetPaused(v bool)`

SetPaused sets Paused field to given value.


### GetCreatedAt

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionsList200ResponseSubscriptionsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionsList200ResponseSubscriptionsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


