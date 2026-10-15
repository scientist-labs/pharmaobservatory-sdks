# ApiV1TrialsRecordsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requested** | **[]string** |  | 
**Trials** | [**[]ApiV1TrialsRecordsGet200ResponseTrialsInner**](ApiV1TrialsRecordsGet200ResponseTrialsInner.md) |  | 
**NotFound** | **[]string** |  | 
**ExtractionCoverage** | Pointer to [**ApiV1TrialsRecordsGet200ResponseExtractionCoverage**](ApiV1TrialsRecordsGet200ResponseExtractionCoverage.md) |  | [optional] 
**BuildState** | [**ApiV1TrialsRecordsGet200ResponseBuildState**](ApiV1TrialsRecordsGet200ResponseBuildState.md) |  | 
**Sections** | Pointer to [**ApiV1TrialsRecordsGet200ResponseSections**](ApiV1TrialsRecordsGet200ResponseSections.md) |  | [optional] 

## Methods

### NewApiV1TrialsRecordsGet200Response

`func NewApiV1TrialsRecordsGet200Response(requested []string, trials []ApiV1TrialsRecordsGet200ResponseTrialsInner, notFound []string, buildState ApiV1TrialsRecordsGet200ResponseBuildState, ) *ApiV1TrialsRecordsGet200Response`

NewApiV1TrialsRecordsGet200Response instantiates a new ApiV1TrialsRecordsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsRecordsGet200ResponseWithDefaults

`func NewApiV1TrialsRecordsGet200ResponseWithDefaults() *ApiV1TrialsRecordsGet200Response`

NewApiV1TrialsRecordsGet200ResponseWithDefaults instantiates a new ApiV1TrialsRecordsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequested

`func (o *ApiV1TrialsRecordsGet200Response) GetRequested() []string`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *ApiV1TrialsRecordsGet200Response) GetRequestedOk() (*[]string, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *ApiV1TrialsRecordsGet200Response) SetRequested(v []string)`

SetRequested sets Requested field to given value.


### GetTrials

`func (o *ApiV1TrialsRecordsGet200Response) GetTrials() []ApiV1TrialsRecordsGet200ResponseTrialsInner`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *ApiV1TrialsRecordsGet200Response) GetTrialsOk() (*[]ApiV1TrialsRecordsGet200ResponseTrialsInner, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *ApiV1TrialsRecordsGet200Response) SetTrials(v []ApiV1TrialsRecordsGet200ResponseTrialsInner)`

SetTrials sets Trials field to given value.


### GetNotFound

`func (o *ApiV1TrialsRecordsGet200Response) GetNotFound() []string`

GetNotFound returns the NotFound field if non-nil, zero value otherwise.

### GetNotFoundOk

`func (o *ApiV1TrialsRecordsGet200Response) GetNotFoundOk() (*[]string, bool)`

GetNotFoundOk returns a tuple with the NotFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFound

`func (o *ApiV1TrialsRecordsGet200Response) SetNotFound(v []string)`

SetNotFound sets NotFound field to given value.


### GetExtractionCoverage

`func (o *ApiV1TrialsRecordsGet200Response) GetExtractionCoverage() ApiV1TrialsRecordsGet200ResponseExtractionCoverage`

GetExtractionCoverage returns the ExtractionCoverage field if non-nil, zero value otherwise.

### GetExtractionCoverageOk

`func (o *ApiV1TrialsRecordsGet200Response) GetExtractionCoverageOk() (*ApiV1TrialsRecordsGet200ResponseExtractionCoverage, bool)`

GetExtractionCoverageOk returns a tuple with the ExtractionCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionCoverage

`func (o *ApiV1TrialsRecordsGet200Response) SetExtractionCoverage(v ApiV1TrialsRecordsGet200ResponseExtractionCoverage)`

SetExtractionCoverage sets ExtractionCoverage field to given value.

### HasExtractionCoverage

`func (o *ApiV1TrialsRecordsGet200Response) HasExtractionCoverage() bool`

HasExtractionCoverage returns a boolean if a field has been set.

### GetBuildState

`func (o *ApiV1TrialsRecordsGet200Response) GetBuildState() ApiV1TrialsRecordsGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1TrialsRecordsGet200Response) GetBuildStateOk() (*ApiV1TrialsRecordsGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1TrialsRecordsGet200Response) SetBuildState(v ApiV1TrialsRecordsGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetSections

`func (o *ApiV1TrialsRecordsGet200Response) GetSections() ApiV1TrialsRecordsGet200ResponseSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ApiV1TrialsRecordsGet200Response) GetSectionsOk() (*ApiV1TrialsRecordsGet200ResponseSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ApiV1TrialsRecordsGet200Response) SetSections(v ApiV1TrialsRecordsGet200ResponseSections)`

SetSections sets Sections field to given value.

### HasSections

`func (o *ApiV1TrialsRecordsGet200Response) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


