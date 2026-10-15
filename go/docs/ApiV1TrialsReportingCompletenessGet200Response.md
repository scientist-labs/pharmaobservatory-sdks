# ApiV1TrialsReportingCompletenessGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters**](ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters.md) |  | [optional] 
**MatchedCount** | Pointer to **int32** |  | [optional] 
**ReportedCount** | Pointer to **int32** |  | [optional] 
**ReportedShare** | Pointer to **float32** |  | [optional] 
**DemographicFieldPopulationRate** | Pointer to [**ApiV1TrialsReportingCompletenessGet200ResponseDemographicFieldPopulationRate**](ApiV1TrialsReportingCompletenessGet200ResponseDemographicFieldPopulationRate.md) |  | [optional] 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1TrialsReportingCompletenessGet200Response

`func NewApiV1TrialsReportingCompletenessGet200Response(buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, ) *ApiV1TrialsReportingCompletenessGet200Response`

NewApiV1TrialsReportingCompletenessGet200Response instantiates a new ApiV1TrialsReportingCompletenessGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsReportingCompletenessGet200ResponseWithDefaults

`func NewApiV1TrialsReportingCompletenessGet200ResponseWithDefaults() *ApiV1TrialsReportingCompletenessGet200Response`

NewApiV1TrialsReportingCompletenessGet200ResponseWithDefaults instantiates a new ApiV1TrialsReportingCompletenessGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetFilters() ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetFiltersOk() (*ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiV1TrialsReportingCompletenessGet200Response) SetFilters(v ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ApiV1TrialsReportingCompletenessGet200Response) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMatchedCount

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetMatchedCount() int32`

GetMatchedCount returns the MatchedCount field if non-nil, zero value otherwise.

### GetMatchedCountOk

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetMatchedCountOk() (*int32, bool)`

GetMatchedCountOk returns a tuple with the MatchedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCount

`func (o *ApiV1TrialsReportingCompletenessGet200Response) SetMatchedCount(v int32)`

SetMatchedCount sets MatchedCount field to given value.

### HasMatchedCount

`func (o *ApiV1TrialsReportingCompletenessGet200Response) HasMatchedCount() bool`

HasMatchedCount returns a boolean if a field has been set.

### GetReportedCount

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetReportedCount() int32`

GetReportedCount returns the ReportedCount field if non-nil, zero value otherwise.

### GetReportedCountOk

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetReportedCountOk() (*int32, bool)`

GetReportedCountOk returns a tuple with the ReportedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedCount

`func (o *ApiV1TrialsReportingCompletenessGet200Response) SetReportedCount(v int32)`

SetReportedCount sets ReportedCount field to given value.

### HasReportedCount

`func (o *ApiV1TrialsReportingCompletenessGet200Response) HasReportedCount() bool`

HasReportedCount returns a boolean if a field has been set.

### GetReportedShare

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetReportedShare() float32`

GetReportedShare returns the ReportedShare field if non-nil, zero value otherwise.

### GetReportedShareOk

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetReportedShareOk() (*float32, bool)`

GetReportedShareOk returns a tuple with the ReportedShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedShare

`func (o *ApiV1TrialsReportingCompletenessGet200Response) SetReportedShare(v float32)`

SetReportedShare sets ReportedShare field to given value.

### HasReportedShare

`func (o *ApiV1TrialsReportingCompletenessGet200Response) HasReportedShare() bool`

HasReportedShare returns a boolean if a field has been set.

### GetDemographicFieldPopulationRate

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetDemographicFieldPopulationRate() ApiV1TrialsReportingCompletenessGet200ResponseDemographicFieldPopulationRate`

GetDemographicFieldPopulationRate returns the DemographicFieldPopulationRate field if non-nil, zero value otherwise.

### GetDemographicFieldPopulationRateOk

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetDemographicFieldPopulationRateOk() (*ApiV1TrialsReportingCompletenessGet200ResponseDemographicFieldPopulationRate, bool)`

GetDemographicFieldPopulationRateOk returns a tuple with the DemographicFieldPopulationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographicFieldPopulationRate

`func (o *ApiV1TrialsReportingCompletenessGet200Response) SetDemographicFieldPopulationRate(v ApiV1TrialsReportingCompletenessGet200ResponseDemographicFieldPopulationRate)`

SetDemographicFieldPopulationRate sets DemographicFieldPopulationRate field to given value.

### HasDemographicFieldPopulationRate

`func (o *ApiV1TrialsReportingCompletenessGet200Response) HasDemographicFieldPopulationRate() bool`

HasDemographicFieldPopulationRate returns a boolean if a field has been set.

### GetBuildState

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1TrialsReportingCompletenessGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetError

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1TrialsReportingCompletenessGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1TrialsReportingCompletenessGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1TrialsReportingCompletenessGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1TrialsReportingCompletenessGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1TrialsReportingCompletenessGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


