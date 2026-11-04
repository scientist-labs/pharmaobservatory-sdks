# ChangesList200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**SubscriptionId** | **int32** |  | 
**EventId** | **string** |  | 
**Dataset** | **string** |  | 
**Key** | **string** |  | 
**Title** | **string** |  | 
**EventType** | **string** |  | 
**EventLabel** | **string** |  | 
**Field** | **string** |  | 
**OldValue** | **string** |  | 
**NewValue** | **string** |  | 
**SourceDate** | **string** |  | 
**ObservedAt** | **string** |  | 
**MatchKind** | **string** |  | 
**EntryKeys** | [**ChangesList200ResponseItemsInnerEntryKeys**](ChangesList200ResponseItemsInnerEntryKeys.md) |  | 
**SourceCredit** | **string** |  | 

## Methods

### NewChangesList200ResponseItemsInner

`func NewChangesList200ResponseItemsInner(id int32, subscriptionId int32, eventId string, dataset string, key string, title string, eventType string, eventLabel string, field string, oldValue string, newValue string, sourceDate string, observedAt string, matchKind string, entryKeys ChangesList200ResponseItemsInnerEntryKeys, sourceCredit string, ) *ChangesList200ResponseItemsInner`

NewChangesList200ResponseItemsInner instantiates a new ChangesList200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangesList200ResponseItemsInnerWithDefaults

`func NewChangesList200ResponseItemsInnerWithDefaults() *ChangesList200ResponseItemsInner`

NewChangesList200ResponseItemsInnerWithDefaults instantiates a new ChangesList200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangesList200ResponseItemsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangesList200ResponseItemsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangesList200ResponseItemsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetSubscriptionId

`func (o *ChangesList200ResponseItemsInner) GetSubscriptionId() int32`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ChangesList200ResponseItemsInner) GetSubscriptionIdOk() (*int32, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ChangesList200ResponseItemsInner) SetSubscriptionId(v int32)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetEventId

`func (o *ChangesList200ResponseItemsInner) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ChangesList200ResponseItemsInner) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ChangesList200ResponseItemsInner) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetDataset

`func (o *ChangesList200ResponseItemsInner) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *ChangesList200ResponseItemsInner) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *ChangesList200ResponseItemsInner) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetKey

`func (o *ChangesList200ResponseItemsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ChangesList200ResponseItemsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ChangesList200ResponseItemsInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetTitle

`func (o *ChangesList200ResponseItemsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChangesList200ResponseItemsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChangesList200ResponseItemsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEventType

`func (o *ChangesList200ResponseItemsInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ChangesList200ResponseItemsInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ChangesList200ResponseItemsInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventLabel

`func (o *ChangesList200ResponseItemsInner) GetEventLabel() string`

GetEventLabel returns the EventLabel field if non-nil, zero value otherwise.

### GetEventLabelOk

`func (o *ChangesList200ResponseItemsInner) GetEventLabelOk() (*string, bool)`

GetEventLabelOk returns a tuple with the EventLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLabel

`func (o *ChangesList200ResponseItemsInner) SetEventLabel(v string)`

SetEventLabel sets EventLabel field to given value.


### GetField

`func (o *ChangesList200ResponseItemsInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ChangesList200ResponseItemsInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ChangesList200ResponseItemsInner) SetField(v string)`

SetField sets Field field to given value.


### GetOldValue

`func (o *ChangesList200ResponseItemsInner) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *ChangesList200ResponseItemsInner) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *ChangesList200ResponseItemsInner) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.


### GetNewValue

`func (o *ChangesList200ResponseItemsInner) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *ChangesList200ResponseItemsInner) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *ChangesList200ResponseItemsInner) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.


### GetSourceDate

`func (o *ChangesList200ResponseItemsInner) GetSourceDate() string`

GetSourceDate returns the SourceDate field if non-nil, zero value otherwise.

### GetSourceDateOk

`func (o *ChangesList200ResponseItemsInner) GetSourceDateOk() (*string, bool)`

GetSourceDateOk returns a tuple with the SourceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDate

`func (o *ChangesList200ResponseItemsInner) SetSourceDate(v string)`

SetSourceDate sets SourceDate field to given value.


### GetObservedAt

`func (o *ChangesList200ResponseItemsInner) GetObservedAt() string`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *ChangesList200ResponseItemsInner) GetObservedAtOk() (*string, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *ChangesList200ResponseItemsInner) SetObservedAt(v string)`

SetObservedAt sets ObservedAt field to given value.


### GetMatchKind

`func (o *ChangesList200ResponseItemsInner) GetMatchKind() string`

GetMatchKind returns the MatchKind field if non-nil, zero value otherwise.

### GetMatchKindOk

`func (o *ChangesList200ResponseItemsInner) GetMatchKindOk() (*string, bool)`

GetMatchKindOk returns a tuple with the MatchKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchKind

`func (o *ChangesList200ResponseItemsInner) SetMatchKind(v string)`

SetMatchKind sets MatchKind field to given value.


### GetEntryKeys

`func (o *ChangesList200ResponseItemsInner) GetEntryKeys() ChangesList200ResponseItemsInnerEntryKeys`

GetEntryKeys returns the EntryKeys field if non-nil, zero value otherwise.

### GetEntryKeysOk

`func (o *ChangesList200ResponseItemsInner) GetEntryKeysOk() (*ChangesList200ResponseItemsInnerEntryKeys, bool)`

GetEntryKeysOk returns a tuple with the EntryKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryKeys

`func (o *ChangesList200ResponseItemsInner) SetEntryKeys(v ChangesList200ResponseItemsInnerEntryKeys)`

SetEntryKeys sets EntryKeys field to given value.


### GetSourceCredit

`func (o *ChangesList200ResponseItemsInner) GetSourceCredit() string`

GetSourceCredit returns the SourceCredit field if non-nil, zero value otherwise.

### GetSourceCreditOk

`func (o *ChangesList200ResponseItemsInner) GetSourceCreditOk() (*string, bool)`

GetSourceCreditOk returns a tuple with the SourceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCredit

`func (o *ChangesList200ResponseItemsInner) SetSourceCredit(v string)`

SetSourceCredit sets SourceCredit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


