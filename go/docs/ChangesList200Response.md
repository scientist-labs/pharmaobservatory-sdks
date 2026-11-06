# ChangesList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ChangesList200ResponseItemsInner**](ChangesList200ResponseItemsInner.md) |  | 
**Count** | **int32** |  | 
**NextCursor** | **string** |  | 
**HasMore** | **bool** |  | 
**BuildState** | [**ChangesList200ResponseBuildState**](ChangesList200ResponseBuildState.md) |  | 
**OldestAvailable** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChangesList200Response

`func NewChangesList200Response(items []ChangesList200ResponseItemsInner, count int32, nextCursor string, hasMore bool, buildState ChangesList200ResponseBuildState, ) *ChangesList200Response`

NewChangesList200Response instantiates a new ChangesList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangesList200ResponseWithDefaults

`func NewChangesList200ResponseWithDefaults() *ChangesList200Response`

NewChangesList200ResponseWithDefaults instantiates a new ChangesList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ChangesList200Response) GetItems() []ChangesList200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ChangesList200Response) GetItemsOk() (*[]ChangesList200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ChangesList200Response) SetItems(v []ChangesList200ResponseItemsInner)`

SetItems sets Items field to given value.


### GetCount

`func (o *ChangesList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ChangesList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ChangesList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNextCursor

`func (o *ChangesList200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ChangesList200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ChangesList200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### GetHasMore

`func (o *ChangesList200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ChangesList200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ChangesList200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetBuildState

`func (o *ChangesList200Response) GetBuildState() ChangesList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ChangesList200Response) GetBuildStateOk() (*ChangesList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ChangesList200Response) SetBuildState(v ChangesList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetOldestAvailable

`func (o *ChangesList200Response) GetOldestAvailable() string`

GetOldestAvailable returns the OldestAvailable field if non-nil, zero value otherwise.

### GetOldestAvailableOk

`func (o *ChangesList200Response) GetOldestAvailableOk() (*string, bool)`

GetOldestAvailableOk returns a tuple with the OldestAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestAvailable

`func (o *ChangesList200Response) SetOldestAvailable(v string)`

SetOldestAvailable sets OldestAvailable field to given value.

### HasOldestAvailable

`func (o *ChangesList200Response) HasOldestAvailable() bool`

HasOldestAvailable returns a boolean if a field has been set.

### SetOldestAvailableNil

`func (o *ChangesList200Response) SetOldestAvailableNil(b bool)`

 SetOldestAvailableNil sets the value for OldestAvailable to be an explicit nil

### UnsetOldestAvailable
`func (o *ChangesList200Response) UnsetOldestAvailable()`

UnsetOldestAvailable ensures that no value is present for OldestAvailable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


