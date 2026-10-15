# ApiV1BiomarkerAssayScreenGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**ApiV1BiomarkerAssayScreenGet200ResponseFilters**](ApiV1BiomarkerAssayScreenGet200ResponseFilters.md) |  | 
**MatchedCount** | **int32** |  | 
**ConfidenceCaveat** | **string** |  | 
**Evidence** | [**[]ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner**](ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner.md) |  | 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 

## Methods

### NewApiV1BiomarkerAssayScreenGet200Response

`func NewApiV1BiomarkerAssayScreenGet200Response(filters ApiV1BiomarkerAssayScreenGet200ResponseFilters, matchedCount int32, confidenceCaveat string, evidence []ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner, buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, ) *ApiV1BiomarkerAssayScreenGet200Response`

NewApiV1BiomarkerAssayScreenGet200Response instantiates a new ApiV1BiomarkerAssayScreenGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1BiomarkerAssayScreenGet200ResponseWithDefaults

`func NewApiV1BiomarkerAssayScreenGet200ResponseWithDefaults() *ApiV1BiomarkerAssayScreenGet200Response`

NewApiV1BiomarkerAssayScreenGet200ResponseWithDefaults instantiates a new ApiV1BiomarkerAssayScreenGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ApiV1BiomarkerAssayScreenGet200Response) GetFilters() ApiV1BiomarkerAssayScreenGet200ResponseFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiV1BiomarkerAssayScreenGet200Response) GetFiltersOk() (*ApiV1BiomarkerAssayScreenGet200ResponseFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiV1BiomarkerAssayScreenGet200Response) SetFilters(v ApiV1BiomarkerAssayScreenGet200ResponseFilters)`

SetFilters sets Filters field to given value.


### GetMatchedCount

`func (o *ApiV1BiomarkerAssayScreenGet200Response) GetMatchedCount() int32`

GetMatchedCount returns the MatchedCount field if non-nil, zero value otherwise.

### GetMatchedCountOk

`func (o *ApiV1BiomarkerAssayScreenGet200Response) GetMatchedCountOk() (*int32, bool)`

GetMatchedCountOk returns a tuple with the MatchedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCount

`func (o *ApiV1BiomarkerAssayScreenGet200Response) SetMatchedCount(v int32)`

SetMatchedCount sets MatchedCount field to given value.


### GetConfidenceCaveat

`func (o *ApiV1BiomarkerAssayScreenGet200Response) GetConfidenceCaveat() string`

GetConfidenceCaveat returns the ConfidenceCaveat field if non-nil, zero value otherwise.

### GetConfidenceCaveatOk

`func (o *ApiV1BiomarkerAssayScreenGet200Response) GetConfidenceCaveatOk() (*string, bool)`

GetConfidenceCaveatOk returns a tuple with the ConfidenceCaveat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceCaveat

`func (o *ApiV1BiomarkerAssayScreenGet200Response) SetConfidenceCaveat(v string)`

SetConfidenceCaveat sets ConfidenceCaveat field to given value.


### GetEvidence

`func (o *ApiV1BiomarkerAssayScreenGet200Response) GetEvidence() []ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *ApiV1BiomarkerAssayScreenGet200Response) GetEvidenceOk() (*[]ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *ApiV1BiomarkerAssayScreenGet200Response) SetEvidence(v []ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner)`

SetEvidence sets Evidence field to given value.


### GetBuildState

`func (o *ApiV1BiomarkerAssayScreenGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1BiomarkerAssayScreenGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1BiomarkerAssayScreenGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


