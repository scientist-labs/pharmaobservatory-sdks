# ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**AsOf** | **string** |  | 
**Enrollment** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerEnrollment**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerEnrollment.md) |  | 
**StartDate** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate.md) |  | 
**PrimaryCompletionDate** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate.md) |  | 
**CompletionDate** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate.md) |  | 
**OverallStatus** | [**ApiV1TrialsHistoryRecordHistoryObservationsInnerOverallStatus**](ApiV1TrialsHistoryRecordHistoryObservationsInnerOverallStatus.md) |  | 
**Sites** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerSites**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerSites.md) |  | 

## Methods

### NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner

`func NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner(kind string, asOf string, enrollment ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerEnrollment, startDate ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate, primaryCompletionDate ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate, completionDate ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate, overallStatus ApiV1TrialsHistoryRecordHistoryObservationsInnerOverallStatus, sites ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerSites, ) *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner`

NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner instantiates a new ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerWithDefaults

`func NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerWithDefaults() *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner`

NewApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerWithDefaults instantiates a new ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAsOf

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.


### GetEnrollment

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetEnrollment() ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetEnrollmentOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) SetEnrollment(v ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerEnrollment)`

SetEnrollment sets Enrollment field to given value.


### GetStartDate

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetStartDate() ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetStartDateOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) SetStartDate(v ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate)`

SetStartDate sets StartDate field to given value.


### GetPrimaryCompletionDate

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetPrimaryCompletionDate() ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate`

GetPrimaryCompletionDate returns the PrimaryCompletionDate field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetPrimaryCompletionDateOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate, bool)`

GetPrimaryCompletionDateOk returns a tuple with the PrimaryCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDate

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) SetPrimaryCompletionDate(v ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate)`

SetPrimaryCompletionDate sets PrimaryCompletionDate field to given value.


### GetCompletionDate

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetCompletionDate() ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetCompletionDateOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) SetCompletionDate(v ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate)`

SetCompletionDate sets CompletionDate field to given value.


### GetOverallStatus

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetOverallStatus() ApiV1TrialsHistoryRecordHistoryObservationsInnerOverallStatus`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetOverallStatusOk() (*ApiV1TrialsHistoryRecordHistoryObservationsInnerOverallStatus, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) SetOverallStatus(v ApiV1TrialsHistoryRecordHistoryObservationsInnerOverallStatus)`

SetOverallStatus sets OverallStatus field to given value.


### GetSites

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetSites() ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerSites`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) GetSitesOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerSites, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner) SetSites(v ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerSites)`

SetSites sets Sites field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


