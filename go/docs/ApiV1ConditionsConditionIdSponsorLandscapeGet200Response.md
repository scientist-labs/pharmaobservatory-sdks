# ApiV1ConditionsConditionIdSponsorLandscapeGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | **string** |  | 
**Name** | **NullableString** |  | 
**Sponsors** | [**[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSponsorsInner**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSponsorsInner.md) |  | 
**PhaseBreakdown** | [**[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponsePhaseBreakdownInner**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponsePhaseBreakdownInner.md) |  | 
**ActiveYearSpan** | [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseActiveYearSpan**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseActiveYearSpan.md) |  | 
**BuildState** | Pointer to [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  | [optional] 
**Sections** | Pointer to [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSections**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSections.md) |  | [optional] 
**TrialsByYear** | Pointer to [**[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner.md) |  | [optional] 
**TrialStartForecast** | Pointer to [**NullableApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialStartForecast**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialStartForecast.md) |  | [optional] 
**Statement** | Pointer to **NullableString** |  | [optional] 
**RecruitmentRollup** | Pointer to [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup.md) |  | [optional] 
**RecruitmentStatement** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1ConditionsConditionIdSponsorLandscapeGet200Response

`func NewApiV1ConditionsConditionIdSponsorLandscapeGet200Response(conditionId string, name NullableString, sponsors []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSponsorsInner, phaseBreakdown []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponsePhaseBreakdownInner, activeYearSpan ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseActiveYearSpan, ) *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response`

NewApiV1ConditionsConditionIdSponsorLandscapeGet200Response instantiates a new ApiV1ConditionsConditionIdSponsorLandscapeGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseWithDefaults

`func NewApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseWithDefaults() *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response`

NewApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseWithDefaults instantiates a new ApiV1ConditionsConditionIdSponsorLandscapeGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.


### GetName

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSponsors

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetSponsors() []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSponsorsInner`

GetSponsors returns the Sponsors field if non-nil, zero value otherwise.

### GetSponsorsOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetSponsorsOk() (*[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSponsorsInner, bool)`

GetSponsorsOk returns a tuple with the Sponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsors

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetSponsors(v []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSponsorsInner)`

SetSponsors sets Sponsors field to given value.


### GetPhaseBreakdown

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetPhaseBreakdown() []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponsePhaseBreakdownInner`

GetPhaseBreakdown returns the PhaseBreakdown field if non-nil, zero value otherwise.

### GetPhaseBreakdownOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetPhaseBreakdownOk() (*[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponsePhaseBreakdownInner, bool)`

GetPhaseBreakdownOk returns a tuple with the PhaseBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseBreakdown

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetPhaseBreakdown(v []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponsePhaseBreakdownInner)`

SetPhaseBreakdown sets PhaseBreakdown field to given value.


### GetActiveYearSpan

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetActiveYearSpan() ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseActiveYearSpan`

GetActiveYearSpan returns the ActiveYearSpan field if non-nil, zero value otherwise.

### GetActiveYearSpanOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetActiveYearSpanOk() (*ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseActiveYearSpan, bool)`

GetActiveYearSpanOk returns a tuple with the ActiveYearSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveYearSpan

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetActiveYearSpan(v ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseActiveYearSpan)`

SetActiveYearSpan sets ActiveYearSpan field to given value.


### GetBuildState

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetBuildState() ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetBuildStateOk() (*ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetBuildState(v ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### GetSections

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetSections() ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetSectionsOk() (*ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetSections(v ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSections)`

SetSections sets Sections field to given value.

### HasSections

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetTrialsByYear

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetTrialsByYear() []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner`

GetTrialsByYear returns the TrialsByYear field if non-nil, zero value otherwise.

### GetTrialsByYearOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetTrialsByYearOk() (*[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner, bool)`

GetTrialsByYearOk returns a tuple with the TrialsByYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialsByYear

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetTrialsByYear(v []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner)`

SetTrialsByYear sets TrialsByYear field to given value.

### HasTrialsByYear

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) HasTrialsByYear() bool`

HasTrialsByYear returns a boolean if a field has been set.

### GetTrialStartForecast

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetTrialStartForecast() ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialStartForecast`

GetTrialStartForecast returns the TrialStartForecast field if non-nil, zero value otherwise.

### GetTrialStartForecastOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetTrialStartForecastOk() (*ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialStartForecast, bool)`

GetTrialStartForecastOk returns a tuple with the TrialStartForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStartForecast

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetTrialStartForecast(v ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialStartForecast)`

SetTrialStartForecast sets TrialStartForecast field to given value.

### HasTrialStartForecast

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) HasTrialStartForecast() bool`

HasTrialStartForecast returns a boolean if a field has been set.

### SetTrialStartForecastNil

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetTrialStartForecastNil(b bool)`

 SetTrialStartForecastNil sets the value for TrialStartForecast to be an explicit nil

### UnsetTrialStartForecast
`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) UnsetTrialStartForecast()`

UnsetTrialStartForecast ensures that no value is present for TrialStartForecast, not even an explicit nil
### GetStatement

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### SetStatementNil

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetStatementNil(b bool)`

 SetStatementNil sets the value for Statement to be an explicit nil

### UnsetStatement
`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) UnsetStatement()`

UnsetStatement ensures that no value is present for Statement, not even an explicit nil
### GetRecruitmentRollup

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetRecruitmentRollup() ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup`

GetRecruitmentRollup returns the RecruitmentRollup field if non-nil, zero value otherwise.

### GetRecruitmentRollupOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetRecruitmentRollupOk() (*ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup, bool)`

GetRecruitmentRollupOk returns a tuple with the RecruitmentRollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecruitmentRollup

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetRecruitmentRollup(v ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup)`

SetRecruitmentRollup sets RecruitmentRollup field to given value.

### HasRecruitmentRollup

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) HasRecruitmentRollup() bool`

HasRecruitmentRollup returns a boolean if a field has been set.

### GetRecruitmentStatement

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetRecruitmentStatement() string`

GetRecruitmentStatement returns the RecruitmentStatement field if non-nil, zero value otherwise.

### GetRecruitmentStatementOk

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) GetRecruitmentStatementOk() (*string, bool)`

GetRecruitmentStatementOk returns a tuple with the RecruitmentStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecruitmentStatement

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) SetRecruitmentStatement(v string)`

SetRecruitmentStatement sets RecruitmentStatement field to given value.

### HasRecruitmentStatement

`func (o *ApiV1ConditionsConditionIdSponsorLandscapeGet200Response) HasRecruitmentStatement() bool`

HasRecruitmentStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


