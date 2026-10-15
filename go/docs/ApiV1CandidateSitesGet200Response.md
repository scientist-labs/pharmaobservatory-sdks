# ApiV1CandidateSitesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**ApiV1CandidateSitesGet200ResponseFilters**](ApiV1CandidateSitesGet200ResponseFilters.md) |  | [optional] 
**MatchedCount** | Pointer to **int32** |  | [optional] 
**SaturatedCount** | Pointer to **int32** |  | [optional] 
**ReturnedCount** | Pointer to **int32** |  | [optional] 
**Sites** | Pointer to [**[]ApiV1CandidateSitesGet200ResponseSitesInner**](ApiV1CandidateSitesGet200ResponseSitesInner.md) |  | [optional] 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1CandidateSitesGet200Response

`func NewApiV1CandidateSitesGet200Response(buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, ) *ApiV1CandidateSitesGet200Response`

NewApiV1CandidateSitesGet200Response instantiates a new ApiV1CandidateSitesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1CandidateSitesGet200ResponseWithDefaults

`func NewApiV1CandidateSitesGet200ResponseWithDefaults() *ApiV1CandidateSitesGet200Response`

NewApiV1CandidateSitesGet200ResponseWithDefaults instantiates a new ApiV1CandidateSitesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ApiV1CandidateSitesGet200Response) GetFilters() ApiV1CandidateSitesGet200ResponseFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiV1CandidateSitesGet200Response) GetFiltersOk() (*ApiV1CandidateSitesGet200ResponseFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiV1CandidateSitesGet200Response) SetFilters(v ApiV1CandidateSitesGet200ResponseFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ApiV1CandidateSitesGet200Response) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMatchedCount

`func (o *ApiV1CandidateSitesGet200Response) GetMatchedCount() int32`

GetMatchedCount returns the MatchedCount field if non-nil, zero value otherwise.

### GetMatchedCountOk

`func (o *ApiV1CandidateSitesGet200Response) GetMatchedCountOk() (*int32, bool)`

GetMatchedCountOk returns a tuple with the MatchedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCount

`func (o *ApiV1CandidateSitesGet200Response) SetMatchedCount(v int32)`

SetMatchedCount sets MatchedCount field to given value.

### HasMatchedCount

`func (o *ApiV1CandidateSitesGet200Response) HasMatchedCount() bool`

HasMatchedCount returns a boolean if a field has been set.

### GetSaturatedCount

`func (o *ApiV1CandidateSitesGet200Response) GetSaturatedCount() int32`

GetSaturatedCount returns the SaturatedCount field if non-nil, zero value otherwise.

### GetSaturatedCountOk

`func (o *ApiV1CandidateSitesGet200Response) GetSaturatedCountOk() (*int32, bool)`

GetSaturatedCountOk returns a tuple with the SaturatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturatedCount

`func (o *ApiV1CandidateSitesGet200Response) SetSaturatedCount(v int32)`

SetSaturatedCount sets SaturatedCount field to given value.

### HasSaturatedCount

`func (o *ApiV1CandidateSitesGet200Response) HasSaturatedCount() bool`

HasSaturatedCount returns a boolean if a field has been set.

### GetReturnedCount

`func (o *ApiV1CandidateSitesGet200Response) GetReturnedCount() int32`

GetReturnedCount returns the ReturnedCount field if non-nil, zero value otherwise.

### GetReturnedCountOk

`func (o *ApiV1CandidateSitesGet200Response) GetReturnedCountOk() (*int32, bool)`

GetReturnedCountOk returns a tuple with the ReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedCount

`func (o *ApiV1CandidateSitesGet200Response) SetReturnedCount(v int32)`

SetReturnedCount sets ReturnedCount field to given value.

### HasReturnedCount

`func (o *ApiV1CandidateSitesGet200Response) HasReturnedCount() bool`

HasReturnedCount returns a boolean if a field has been set.

### GetSites

`func (o *ApiV1CandidateSitesGet200Response) GetSites() []ApiV1CandidateSitesGet200ResponseSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ApiV1CandidateSitesGet200Response) GetSitesOk() (*[]ApiV1CandidateSitesGet200ResponseSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ApiV1CandidateSitesGet200Response) SetSites(v []ApiV1CandidateSitesGet200ResponseSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ApiV1CandidateSitesGet200Response) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetBuildState

`func (o *ApiV1CandidateSitesGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1CandidateSitesGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1CandidateSitesGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetError

`func (o *ApiV1CandidateSitesGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1CandidateSitesGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1CandidateSitesGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1CandidateSitesGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1CandidateSitesGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1CandidateSitesGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1CandidateSitesGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1CandidateSitesGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


