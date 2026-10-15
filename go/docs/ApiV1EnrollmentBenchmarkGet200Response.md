# ApiV1EnrollmentBenchmarkGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Filters** | Pointer to [**ApiV1EnrollmentBenchmarkGet200ResponseFilters**](ApiV1EnrollmentBenchmarkGet200ResponseFilters.md) |  | [optional] 
**Matched** | Pointer to **bool** |  | [optional] 
**Benchmark** | Pointer to [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmark**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmark.md) |  | [optional] 

## Methods

### NewApiV1EnrollmentBenchmarkGet200Response

`func NewApiV1EnrollmentBenchmarkGet200Response(buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, ) *ApiV1EnrollmentBenchmarkGet200Response`

NewApiV1EnrollmentBenchmarkGet200Response instantiates a new ApiV1EnrollmentBenchmarkGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1EnrollmentBenchmarkGet200ResponseWithDefaults

`func NewApiV1EnrollmentBenchmarkGet200ResponseWithDefaults() *ApiV1EnrollmentBenchmarkGet200Response`

NewApiV1EnrollmentBenchmarkGet200ResponseWithDefaults instantiates a new ApiV1EnrollmentBenchmarkGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1EnrollmentBenchmarkGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1EnrollmentBenchmarkGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1EnrollmentBenchmarkGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1EnrollmentBenchmarkGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetBuildState

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1EnrollmentBenchmarkGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetFilters

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetFilters() ApiV1EnrollmentBenchmarkGet200ResponseFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetFiltersOk() (*ApiV1EnrollmentBenchmarkGet200ResponseFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiV1EnrollmentBenchmarkGet200Response) SetFilters(v ApiV1EnrollmentBenchmarkGet200ResponseFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ApiV1EnrollmentBenchmarkGet200Response) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMatched

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetMatched() bool`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetMatchedOk() (*bool, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *ApiV1EnrollmentBenchmarkGet200Response) SetMatched(v bool)`

SetMatched sets Matched field to given value.

### HasMatched

`func (o *ApiV1EnrollmentBenchmarkGet200Response) HasMatched() bool`

HasMatched returns a boolean if a field has been set.

### GetBenchmark

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetBenchmark() ApiV1EnrollmentBenchmarkGet200ResponseBenchmark`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *ApiV1EnrollmentBenchmarkGet200Response) GetBenchmarkOk() (*ApiV1EnrollmentBenchmarkGet200ResponseBenchmark, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *ApiV1EnrollmentBenchmarkGet200Response) SetBenchmark(v ApiV1EnrollmentBenchmarkGet200ResponseBenchmark)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *ApiV1EnrollmentBenchmarkGet200Response) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


