# ApiV1TrialsDiscontinuationReasonsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Filters** | Pointer to [**ApiV1TrialsDiscontinuationReasonsGet200ResponseFilters**](ApiV1TrialsDiscontinuationReasonsGet200ResponseFilters.md) |  | [optional] 
**MatchedCount** | Pointer to **int32** |  | [optional] 
**TrialsWithWhyStopped** | Pointer to **int32** |  | [optional] 
**WhyStoppedCoverageRate** | Pointer to **float32** |  | [optional] 
**CoverageCaveat** | Pointer to **string** |  | [optional] 
**Trials** | Pointer to [**[]ApiV1TrialsDiscontinuationReasonsGet200ResponseTrialsInner**](ApiV1TrialsDiscontinuationReasonsGet200ResponseTrialsInner.md) |  | [optional] 

## Methods

### NewApiV1TrialsDiscontinuationReasonsGet200Response

`func NewApiV1TrialsDiscontinuationReasonsGet200Response(buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, ) *ApiV1TrialsDiscontinuationReasonsGet200Response`

NewApiV1TrialsDiscontinuationReasonsGet200Response instantiates a new ApiV1TrialsDiscontinuationReasonsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsDiscontinuationReasonsGet200ResponseWithDefaults

`func NewApiV1TrialsDiscontinuationReasonsGet200ResponseWithDefaults() *ApiV1TrialsDiscontinuationReasonsGet200Response`

NewApiV1TrialsDiscontinuationReasonsGet200ResponseWithDefaults instantiates a new ApiV1TrialsDiscontinuationReasonsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetBuildState

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetFilters

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetFilters() ApiV1TrialsDiscontinuationReasonsGet200ResponseFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetFiltersOk() (*ApiV1TrialsDiscontinuationReasonsGet200ResponseFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) SetFilters(v ApiV1TrialsDiscontinuationReasonsGet200ResponseFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMatchedCount

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetMatchedCount() int32`

GetMatchedCount returns the MatchedCount field if non-nil, zero value otherwise.

### GetMatchedCountOk

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetMatchedCountOk() (*int32, bool)`

GetMatchedCountOk returns a tuple with the MatchedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCount

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) SetMatchedCount(v int32)`

SetMatchedCount sets MatchedCount field to given value.

### HasMatchedCount

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) HasMatchedCount() bool`

HasMatchedCount returns a boolean if a field has been set.

### GetTrialsWithWhyStopped

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetTrialsWithWhyStopped() int32`

GetTrialsWithWhyStopped returns the TrialsWithWhyStopped field if non-nil, zero value otherwise.

### GetTrialsWithWhyStoppedOk

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetTrialsWithWhyStoppedOk() (*int32, bool)`

GetTrialsWithWhyStoppedOk returns a tuple with the TrialsWithWhyStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialsWithWhyStopped

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) SetTrialsWithWhyStopped(v int32)`

SetTrialsWithWhyStopped sets TrialsWithWhyStopped field to given value.

### HasTrialsWithWhyStopped

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) HasTrialsWithWhyStopped() bool`

HasTrialsWithWhyStopped returns a boolean if a field has been set.

### GetWhyStoppedCoverageRate

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetWhyStoppedCoverageRate() float32`

GetWhyStoppedCoverageRate returns the WhyStoppedCoverageRate field if non-nil, zero value otherwise.

### GetWhyStoppedCoverageRateOk

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetWhyStoppedCoverageRateOk() (*float32, bool)`

GetWhyStoppedCoverageRateOk returns a tuple with the WhyStoppedCoverageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhyStoppedCoverageRate

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) SetWhyStoppedCoverageRate(v float32)`

SetWhyStoppedCoverageRate sets WhyStoppedCoverageRate field to given value.

### HasWhyStoppedCoverageRate

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) HasWhyStoppedCoverageRate() bool`

HasWhyStoppedCoverageRate returns a boolean if a field has been set.

### GetCoverageCaveat

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetCoverageCaveat() string`

GetCoverageCaveat returns the CoverageCaveat field if non-nil, zero value otherwise.

### GetCoverageCaveatOk

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetCoverageCaveatOk() (*string, bool)`

GetCoverageCaveatOk returns a tuple with the CoverageCaveat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageCaveat

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) SetCoverageCaveat(v string)`

SetCoverageCaveat sets CoverageCaveat field to given value.

### HasCoverageCaveat

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) HasCoverageCaveat() bool`

HasCoverageCaveat returns a boolean if a field has been set.

### GetTrials

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetTrials() []ApiV1TrialsDiscontinuationReasonsGet200ResponseTrialsInner`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) GetTrialsOk() (*[]ApiV1TrialsDiscontinuationReasonsGet200ResponseTrialsInner, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) SetTrials(v []ApiV1TrialsDiscontinuationReasonsGet200ResponseTrialsInner)`

SetTrials sets Trials field to given value.

### HasTrials

`func (o *ApiV1TrialsDiscontinuationReasonsGet200Response) HasTrials() bool`

HasTrials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


