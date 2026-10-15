# ApiV1TrialsEligibilityGeneralizabilityScanGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters**](ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters.md) |  | [optional] 
**MatchedCount** | Pointer to **int32** |  | [optional] 
**ReturnedCount** | Pointer to **int32** |  | [optional] 
**Trials** | Pointer to [**[]ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner**](ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner.md) |  | [optional] 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1TrialsEligibilityGeneralizabilityScanGet200Response

`func NewApiV1TrialsEligibilityGeneralizabilityScanGet200Response(buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, ) *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response`

NewApiV1TrialsEligibilityGeneralizabilityScanGet200Response instantiates a new ApiV1TrialsEligibilityGeneralizabilityScanGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseWithDefaults

`func NewApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseWithDefaults() *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response`

NewApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseWithDefaults instantiates a new ApiV1TrialsEligibilityGeneralizabilityScanGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetFilters

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetFilters() ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetFiltersOk() (*ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) SetFilters(v ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMatchedCount

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetMatchedCount() int32`

GetMatchedCount returns the MatchedCount field if non-nil, zero value otherwise.

### GetMatchedCountOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetMatchedCountOk() (*int32, bool)`

GetMatchedCountOk returns a tuple with the MatchedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCount

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) SetMatchedCount(v int32)`

SetMatchedCount sets MatchedCount field to given value.

### HasMatchedCount

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) HasMatchedCount() bool`

HasMatchedCount returns a boolean if a field has been set.

### GetReturnedCount

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetReturnedCount() int32`

GetReturnedCount returns the ReturnedCount field if non-nil, zero value otherwise.

### GetReturnedCountOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetReturnedCountOk() (*int32, bool)`

GetReturnedCountOk returns a tuple with the ReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedCount

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) SetReturnedCount(v int32)`

SetReturnedCount sets ReturnedCount field to given value.

### HasReturnedCount

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) HasReturnedCount() bool`

HasReturnedCount returns a boolean if a field has been set.

### GetTrials

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetTrials() []ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetTrialsOk() (*[]ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) SetTrials(v []ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner)`

SetTrials sets Trials field to given value.

### HasTrials

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) HasTrials() bool`

HasTrials returns a boolean if a field has been set.

### GetBuildState

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetError

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


