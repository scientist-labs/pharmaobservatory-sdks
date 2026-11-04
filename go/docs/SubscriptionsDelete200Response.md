# SubscriptionsDelete200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | **bool** |  | 
**Id** | **int32** |  | 
**Label** | **string** |  | 
**BuildState** | [**ChangesList200ResponseBuildState**](ChangesList200ResponseBuildState.md) |  | 

## Methods

### NewSubscriptionsDelete200Response

`func NewSubscriptionsDelete200Response(deleted bool, id int32, label string, buildState ChangesList200ResponseBuildState, ) *SubscriptionsDelete200Response`

NewSubscriptionsDelete200Response instantiates a new SubscriptionsDelete200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsDelete200ResponseWithDefaults

`func NewSubscriptionsDelete200ResponseWithDefaults() *SubscriptionsDelete200Response`

NewSubscriptionsDelete200ResponseWithDefaults instantiates a new SubscriptionsDelete200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *SubscriptionsDelete200Response) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SubscriptionsDelete200Response) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SubscriptionsDelete200Response) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetId

`func (o *SubscriptionsDelete200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionsDelete200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionsDelete200Response) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *SubscriptionsDelete200Response) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SubscriptionsDelete200Response) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SubscriptionsDelete200Response) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetBuildState

`func (o *SubscriptionsDelete200Response) GetBuildState() ChangesList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *SubscriptionsDelete200Response) GetBuildStateOk() (*ChangesList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *SubscriptionsDelete200Response) SetBuildState(v ChangesList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


