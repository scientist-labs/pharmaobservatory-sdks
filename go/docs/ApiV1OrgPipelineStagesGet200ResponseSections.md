# ApiV1OrgPipelineStagesGet200ResponseSections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InDevelopment** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**AbandonedDormant** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**ApprovedMarketed** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**TrialStartForecastSeries** | Pointer to [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] 
**TrialStartForecastYears** | Pointer to [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] 
**RecruitmentRollupSeries** | Pointer to [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] 
**RecruitmentRollupPhase** | Pointer to [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] 

## Methods

### NewApiV1OrgPipelineStagesGet200ResponseSections

`func NewApiV1OrgPipelineStagesGet200ResponseSections(inDevelopment ApiV1BiomarkerAssayScreenGet200ResponseBuildState, abandonedDormant ApiV1BiomarkerAssayScreenGet200ResponseBuildState, approvedMarketed ApiV1BiomarkerAssayScreenGet200ResponseBuildState, ) *ApiV1OrgPipelineStagesGet200ResponseSections`

NewApiV1OrgPipelineStagesGet200ResponseSections instantiates a new ApiV1OrgPipelineStagesGet200ResponseSections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1OrgPipelineStagesGet200ResponseSectionsWithDefaults

`func NewApiV1OrgPipelineStagesGet200ResponseSectionsWithDefaults() *ApiV1OrgPipelineStagesGet200ResponseSections`

NewApiV1OrgPipelineStagesGet200ResponseSectionsWithDefaults instantiates a new ApiV1OrgPipelineStagesGet200ResponseSections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInDevelopment

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetInDevelopment() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetInDevelopment returns the InDevelopment field if non-nil, zero value otherwise.

### GetInDevelopmentOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetInDevelopmentOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetInDevelopmentOk returns a tuple with the InDevelopment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInDevelopment

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) SetInDevelopment(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetInDevelopment sets InDevelopment field to given value.


### GetAbandonedDormant

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetAbandonedDormant() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetAbandonedDormant returns the AbandonedDormant field if non-nil, zero value otherwise.

### GetAbandonedDormantOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetAbandonedDormantOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetAbandonedDormantOk returns a tuple with the AbandonedDormant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonedDormant

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) SetAbandonedDormant(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetAbandonedDormant sets AbandonedDormant field to given value.


### GetApprovedMarketed

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetApprovedMarketed() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetApprovedMarketed returns the ApprovedMarketed field if non-nil, zero value otherwise.

### GetApprovedMarketedOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetApprovedMarketedOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetApprovedMarketedOk returns a tuple with the ApprovedMarketed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedMarketed

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) SetApprovedMarketed(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetApprovedMarketed sets ApprovedMarketed field to given value.


### GetTrialStartForecastSeries

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetTrialStartForecastSeries() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetTrialStartForecastSeries returns the TrialStartForecastSeries field if non-nil, zero value otherwise.

### GetTrialStartForecastSeriesOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetTrialStartForecastSeriesOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetTrialStartForecastSeriesOk returns a tuple with the TrialStartForecastSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStartForecastSeries

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) SetTrialStartForecastSeries(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetTrialStartForecastSeries sets TrialStartForecastSeries field to given value.

### HasTrialStartForecastSeries

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) HasTrialStartForecastSeries() bool`

HasTrialStartForecastSeries returns a boolean if a field has been set.

### GetTrialStartForecastYears

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetTrialStartForecastYears() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetTrialStartForecastYears returns the TrialStartForecastYears field if non-nil, zero value otherwise.

### GetTrialStartForecastYearsOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetTrialStartForecastYearsOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetTrialStartForecastYearsOk returns a tuple with the TrialStartForecastYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStartForecastYears

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) SetTrialStartForecastYears(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetTrialStartForecastYears sets TrialStartForecastYears field to given value.

### HasTrialStartForecastYears

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) HasTrialStartForecastYears() bool`

HasTrialStartForecastYears returns a boolean if a field has been set.

### GetRecruitmentRollupSeries

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetRecruitmentRollupSeries() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetRecruitmentRollupSeries returns the RecruitmentRollupSeries field if non-nil, zero value otherwise.

### GetRecruitmentRollupSeriesOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetRecruitmentRollupSeriesOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetRecruitmentRollupSeriesOk returns a tuple with the RecruitmentRollupSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecruitmentRollupSeries

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) SetRecruitmentRollupSeries(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetRecruitmentRollupSeries sets RecruitmentRollupSeries field to given value.

### HasRecruitmentRollupSeries

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) HasRecruitmentRollupSeries() bool`

HasRecruitmentRollupSeries returns a boolean if a field has been set.

### GetRecruitmentRollupPhase

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetRecruitmentRollupPhase() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetRecruitmentRollupPhase returns the RecruitmentRollupPhase field if non-nil, zero value otherwise.

### GetRecruitmentRollupPhaseOk

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) GetRecruitmentRollupPhaseOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetRecruitmentRollupPhaseOk returns a tuple with the RecruitmentRollupPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecruitmentRollupPhase

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) SetRecruitmentRollupPhase(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetRecruitmentRollupPhase sets RecruitmentRollupPhase field to given value.

### HasRecruitmentRollupPhase

`func (o *ApiV1OrgPipelineStagesGet200ResponseSections) HasRecruitmentRollupPhase() bool`

HasRecruitmentRollupPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


