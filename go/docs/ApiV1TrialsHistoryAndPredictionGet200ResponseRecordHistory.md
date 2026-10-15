# ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** |  | 
**SitesAvailable** | **bool** |  | 
**CoverageAvailable** | **bool** |  | 
**SnapshotCount** | **int32** |  | 
**FirstAsOf** | **NullableString** |  | 
**LastAsOf** | **NullableString** |  | 
**Window** | **NullableString** |  | 
**FullyCaptured** | **bool** |  | 
**Observations** | [**[]ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner.md) |  | 
**Changes** | [**[]ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryChangesInner**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryChangesInner.md) |  | 

## Methods

### NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory

`func NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory(available bool, sitesAvailable bool, coverageAvailable bool, snapshotCount int32, firstAsOf NullableString, lastAsOf NullableString, window NullableString, fullyCaptured bool, observations []ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner, changes []ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryChangesInner, ) *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory`

NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory instantiates a new ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryWithDefaults

`func NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryWithDefaults() *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory`

NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryWithDefaults instantiates a new ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetSitesAvailable

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetSitesAvailable() bool`

GetSitesAvailable returns the SitesAvailable field if non-nil, zero value otherwise.

### GetSitesAvailableOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetSitesAvailableOk() (*bool, bool)`

GetSitesAvailableOk returns a tuple with the SitesAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesAvailable

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetSitesAvailable(v bool)`

SetSitesAvailable sets SitesAvailable field to given value.


### GetCoverageAvailable

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetCoverageAvailable() bool`

GetCoverageAvailable returns the CoverageAvailable field if non-nil, zero value otherwise.

### GetCoverageAvailableOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetCoverageAvailableOk() (*bool, bool)`

GetCoverageAvailableOk returns a tuple with the CoverageAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageAvailable

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetCoverageAvailable(v bool)`

SetCoverageAvailable sets CoverageAvailable field to given value.


### GetSnapshotCount

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetSnapshotCount() int32`

GetSnapshotCount returns the SnapshotCount field if non-nil, zero value otherwise.

### GetSnapshotCountOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetSnapshotCountOk() (*int32, bool)`

GetSnapshotCountOk returns a tuple with the SnapshotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCount

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetSnapshotCount(v int32)`

SetSnapshotCount sets SnapshotCount field to given value.


### GetFirstAsOf

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetFirstAsOf() string`

GetFirstAsOf returns the FirstAsOf field if non-nil, zero value otherwise.

### GetFirstAsOfOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetFirstAsOfOk() (*string, bool)`

GetFirstAsOfOk returns a tuple with the FirstAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAsOf

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetFirstAsOf(v string)`

SetFirstAsOf sets FirstAsOf field to given value.


### SetFirstAsOfNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetFirstAsOfNil(b bool)`

 SetFirstAsOfNil sets the value for FirstAsOf to be an explicit nil

### UnsetFirstAsOf
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) UnsetFirstAsOf()`

UnsetFirstAsOf ensures that no value is present for FirstAsOf, not even an explicit nil
### GetLastAsOf

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetLastAsOf() string`

GetLastAsOf returns the LastAsOf field if non-nil, zero value otherwise.

### GetLastAsOfOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetLastAsOfOk() (*string, bool)`

GetLastAsOfOk returns a tuple with the LastAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAsOf

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetLastAsOf(v string)`

SetLastAsOf sets LastAsOf field to given value.


### SetLastAsOfNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetLastAsOfNil(b bool)`

 SetLastAsOfNil sets the value for LastAsOf to be an explicit nil

### UnsetLastAsOf
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) UnsetLastAsOf()`

UnsetLastAsOf ensures that no value is present for LastAsOf, not even an explicit nil
### GetWindow

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetWindow(v string)`

SetWindow sets Window field to given value.


### SetWindowNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetWindowNil(b bool)`

 SetWindowNil sets the value for Window to be an explicit nil

### UnsetWindow
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) UnsetWindow()`

UnsetWindow ensures that no value is present for Window, not even an explicit nil
### GetFullyCaptured

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetFullyCaptured() bool`

GetFullyCaptured returns the FullyCaptured field if non-nil, zero value otherwise.

### GetFullyCapturedOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetFullyCapturedOk() (*bool, bool)`

GetFullyCapturedOk returns a tuple with the FullyCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyCaptured

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetFullyCaptured(v bool)`

SetFullyCaptured sets FullyCaptured field to given value.


### GetObservations

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetObservations() []ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetObservationsOk() (*[]ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetObservations(v []ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner)`

SetObservations sets Observations field to given value.


### GetChanges

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetChanges() []ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryChangesInner`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) GetChangesOk() (*[]ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryChangesInner, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory) SetChanges(v []ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryChangesInner)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


