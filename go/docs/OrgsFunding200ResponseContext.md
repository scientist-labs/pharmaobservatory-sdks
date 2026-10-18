# OrgsFunding200ResponseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderNote** | **string** |  | 
**RemainderNoun** | **string** |  | 
**UndisclosedEvents** | Pointer to **int32** |  | [optional] 
**Summary** | Pointer to [**OrgsFunding200ResponseContextSummary**](OrgsFunding200ResponseContextSummary.md) |  | [optional] 
**DownwardCorrections** | Pointer to **int32** |  | [optional] 
**DownwardCorrectionsUsd** | Pointer to **int32** |  | [optional] 
**UnresolvedRecognitionRecords** | Pointer to **int32** |  | [optional] 
**UnresolvedUsd** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrgsFunding200ResponseContext

`func NewOrgsFunding200ResponseContext(orderNote string, remainderNoun string, ) *OrgsFunding200ResponseContext`

NewOrgsFunding200ResponseContext instantiates a new OrgsFunding200ResponseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsFunding200ResponseContextWithDefaults

`func NewOrgsFunding200ResponseContextWithDefaults() *OrgsFunding200ResponseContext`

NewOrgsFunding200ResponseContextWithDefaults instantiates a new OrgsFunding200ResponseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNote

`func (o *OrgsFunding200ResponseContext) GetOrderNote() string`

GetOrderNote returns the OrderNote field if non-nil, zero value otherwise.

### GetOrderNoteOk

`func (o *OrgsFunding200ResponseContext) GetOrderNoteOk() (*string, bool)`

GetOrderNoteOk returns a tuple with the OrderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNote

`func (o *OrgsFunding200ResponseContext) SetOrderNote(v string)`

SetOrderNote sets OrderNote field to given value.


### GetRemainderNoun

`func (o *OrgsFunding200ResponseContext) GetRemainderNoun() string`

GetRemainderNoun returns the RemainderNoun field if non-nil, zero value otherwise.

### GetRemainderNounOk

`func (o *OrgsFunding200ResponseContext) GetRemainderNounOk() (*string, bool)`

GetRemainderNounOk returns a tuple with the RemainderNoun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainderNoun

`func (o *OrgsFunding200ResponseContext) SetRemainderNoun(v string)`

SetRemainderNoun sets RemainderNoun field to given value.


### GetUndisclosedEvents

`func (o *OrgsFunding200ResponseContext) GetUndisclosedEvents() int32`

GetUndisclosedEvents returns the UndisclosedEvents field if non-nil, zero value otherwise.

### GetUndisclosedEventsOk

`func (o *OrgsFunding200ResponseContext) GetUndisclosedEventsOk() (*int32, bool)`

GetUndisclosedEventsOk returns a tuple with the UndisclosedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndisclosedEvents

`func (o *OrgsFunding200ResponseContext) SetUndisclosedEvents(v int32)`

SetUndisclosedEvents sets UndisclosedEvents field to given value.

### HasUndisclosedEvents

`func (o *OrgsFunding200ResponseContext) HasUndisclosedEvents() bool`

HasUndisclosedEvents returns a boolean if a field has been set.

### GetSummary

`func (o *OrgsFunding200ResponseContext) GetSummary() OrgsFunding200ResponseContextSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *OrgsFunding200ResponseContext) GetSummaryOk() (*OrgsFunding200ResponseContextSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *OrgsFunding200ResponseContext) SetSummary(v OrgsFunding200ResponseContextSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *OrgsFunding200ResponseContext) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDownwardCorrections

`func (o *OrgsFunding200ResponseContext) GetDownwardCorrections() int32`

GetDownwardCorrections returns the DownwardCorrections field if non-nil, zero value otherwise.

### GetDownwardCorrectionsOk

`func (o *OrgsFunding200ResponseContext) GetDownwardCorrectionsOk() (*int32, bool)`

GetDownwardCorrectionsOk returns a tuple with the DownwardCorrections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownwardCorrections

`func (o *OrgsFunding200ResponseContext) SetDownwardCorrections(v int32)`

SetDownwardCorrections sets DownwardCorrections field to given value.

### HasDownwardCorrections

`func (o *OrgsFunding200ResponseContext) HasDownwardCorrections() bool`

HasDownwardCorrections returns a boolean if a field has been set.

### GetDownwardCorrectionsUsd

`func (o *OrgsFunding200ResponseContext) GetDownwardCorrectionsUsd() int32`

GetDownwardCorrectionsUsd returns the DownwardCorrectionsUsd field if non-nil, zero value otherwise.

### GetDownwardCorrectionsUsdOk

`func (o *OrgsFunding200ResponseContext) GetDownwardCorrectionsUsdOk() (*int32, bool)`

GetDownwardCorrectionsUsdOk returns a tuple with the DownwardCorrectionsUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownwardCorrectionsUsd

`func (o *OrgsFunding200ResponseContext) SetDownwardCorrectionsUsd(v int32)`

SetDownwardCorrectionsUsd sets DownwardCorrectionsUsd field to given value.

### HasDownwardCorrectionsUsd

`func (o *OrgsFunding200ResponseContext) HasDownwardCorrectionsUsd() bool`

HasDownwardCorrectionsUsd returns a boolean if a field has been set.

### GetUnresolvedRecognitionRecords

`func (o *OrgsFunding200ResponseContext) GetUnresolvedRecognitionRecords() int32`

GetUnresolvedRecognitionRecords returns the UnresolvedRecognitionRecords field if non-nil, zero value otherwise.

### GetUnresolvedRecognitionRecordsOk

`func (o *OrgsFunding200ResponseContext) GetUnresolvedRecognitionRecordsOk() (*int32, bool)`

GetUnresolvedRecognitionRecordsOk returns a tuple with the UnresolvedRecognitionRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedRecognitionRecords

`func (o *OrgsFunding200ResponseContext) SetUnresolvedRecognitionRecords(v int32)`

SetUnresolvedRecognitionRecords sets UnresolvedRecognitionRecords field to given value.

### HasUnresolvedRecognitionRecords

`func (o *OrgsFunding200ResponseContext) HasUnresolvedRecognitionRecords() bool`

HasUnresolvedRecognitionRecords returns a boolean if a field has been set.

### GetUnresolvedUsd

`func (o *OrgsFunding200ResponseContext) GetUnresolvedUsd() int32`

GetUnresolvedUsd returns the UnresolvedUsd field if non-nil, zero value otherwise.

### GetUnresolvedUsdOk

`func (o *OrgsFunding200ResponseContext) GetUnresolvedUsdOk() (*int32, bool)`

GetUnresolvedUsdOk returns a tuple with the UnresolvedUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedUsd

`func (o *OrgsFunding200ResponseContext) SetUnresolvedUsd(v int32)`

SetUnresolvedUsd sets UnresolvedUsd field to given value.

### HasUnresolvedUsd

`func (o *OrgsFunding200ResponseContext) HasUnresolvedUsd() bool`

HasUnresolvedUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


