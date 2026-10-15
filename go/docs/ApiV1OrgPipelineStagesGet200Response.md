# ApiV1OrgPipelineStagesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ApprovedMarketed** | [**[]ApiV1OrgPipelineStagesGet200ResponseApprovedMarketedInner**](ApiV1OrgPipelineStagesGet200ResponseApprovedMarketedInner.md) |  | 
**InDevelopment** | [**[]ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner**](ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner.md) |  | 
**AbandonedDormant** | [**[]ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner**](ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner.md) |  | 
**BuildState** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  | 
**Sections** | [**ApiV1OrgPipelineStagesGet200ResponseSections**](ApiV1OrgPipelineStagesGet200ResponseSections.md) |  | 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**NameWithheldNote** | Pointer to **NullableString** |  | [optional] 
**NameBuilt** | Pointer to **bool** |  | [optional] 
**TrialsByYear** | Pointer to [**[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner.md) |  | [optional] 
**TrialStartForecast** | Pointer to [**ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast**](ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast.md) |  | [optional] 
**Statement** | Pointer to **string** |  | [optional] 
**RecruitmentRollup** | Pointer to [**ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup**](ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup.md) |  | [optional] 
**RecruitmentStatement** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1OrgPipelineStagesGet200Response

`func NewApiV1OrgPipelineStagesGet200Response(approvedMarketed []ApiV1OrgPipelineStagesGet200ResponseApprovedMarketedInner, inDevelopment []ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner, abandonedDormant []ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner, buildState ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, sections ApiV1OrgPipelineStagesGet200ResponseSections, ) *ApiV1OrgPipelineStagesGet200Response`

NewApiV1OrgPipelineStagesGet200Response instantiates a new ApiV1OrgPipelineStagesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1OrgPipelineStagesGet200ResponseWithDefaults

`func NewApiV1OrgPipelineStagesGet200ResponseWithDefaults() *ApiV1OrgPipelineStagesGet200Response`

NewApiV1OrgPipelineStagesGet200ResponseWithDefaults instantiates a new ApiV1OrgPipelineStagesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ApiV1OrgPipelineStagesGet200Response) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApiV1OrgPipelineStagesGet200Response) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ApiV1OrgPipelineStagesGet200Response) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetName

`func (o *ApiV1OrgPipelineStagesGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1OrgPipelineStagesGet200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV1OrgPipelineStagesGet200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApprovedMarketed

`func (o *ApiV1OrgPipelineStagesGet200Response) GetApprovedMarketed() []ApiV1OrgPipelineStagesGet200ResponseApprovedMarketedInner`

GetApprovedMarketed returns the ApprovedMarketed field if non-nil, zero value otherwise.

### GetApprovedMarketedOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetApprovedMarketedOk() (*[]ApiV1OrgPipelineStagesGet200ResponseApprovedMarketedInner, bool)`

GetApprovedMarketedOk returns a tuple with the ApprovedMarketed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedMarketed

`func (o *ApiV1OrgPipelineStagesGet200Response) SetApprovedMarketed(v []ApiV1OrgPipelineStagesGet200ResponseApprovedMarketedInner)`

SetApprovedMarketed sets ApprovedMarketed field to given value.


### GetInDevelopment

`func (o *ApiV1OrgPipelineStagesGet200Response) GetInDevelopment() []ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner`

GetInDevelopment returns the InDevelopment field if non-nil, zero value otherwise.

### GetInDevelopmentOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetInDevelopmentOk() (*[]ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner, bool)`

GetInDevelopmentOk returns a tuple with the InDevelopment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInDevelopment

`func (o *ApiV1OrgPipelineStagesGet200Response) SetInDevelopment(v []ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner)`

SetInDevelopment sets InDevelopment field to given value.


### GetAbandonedDormant

`func (o *ApiV1OrgPipelineStagesGet200Response) GetAbandonedDormant() []ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner`

GetAbandonedDormant returns the AbandonedDormant field if non-nil, zero value otherwise.

### GetAbandonedDormantOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetAbandonedDormantOk() (*[]ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner, bool)`

GetAbandonedDormantOk returns a tuple with the AbandonedDormant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonedDormant

`func (o *ApiV1OrgPipelineStagesGet200Response) SetAbandonedDormant(v []ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner)`

SetAbandonedDormant sets AbandonedDormant field to given value.


### GetBuildState

`func (o *ApiV1OrgPipelineStagesGet200Response) GetBuildState() ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetBuildStateOk() (*ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1OrgPipelineStagesGet200Response) SetBuildState(v ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetSections

`func (o *ApiV1OrgPipelineStagesGet200Response) GetSections() ApiV1OrgPipelineStagesGet200ResponseSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetSectionsOk() (*ApiV1OrgPipelineStagesGet200ResponseSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ApiV1OrgPipelineStagesGet200Response) SetSections(v ApiV1OrgPipelineStagesGet200ResponseSections)`

SetSections sets Sections field to given value.


### GetError

`func (o *ApiV1OrgPipelineStagesGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1OrgPipelineStagesGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1OrgPipelineStagesGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1OrgPipelineStagesGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1OrgPipelineStagesGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1OrgPipelineStagesGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNameWithheldNote

`func (o *ApiV1OrgPipelineStagesGet200Response) GetNameWithheldNote() string`

GetNameWithheldNote returns the NameWithheldNote field if non-nil, zero value otherwise.

### GetNameWithheldNoteOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetNameWithheldNoteOk() (*string, bool)`

GetNameWithheldNoteOk returns a tuple with the NameWithheldNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameWithheldNote

`func (o *ApiV1OrgPipelineStagesGet200Response) SetNameWithheldNote(v string)`

SetNameWithheldNote sets NameWithheldNote field to given value.

### HasNameWithheldNote

`func (o *ApiV1OrgPipelineStagesGet200Response) HasNameWithheldNote() bool`

HasNameWithheldNote returns a boolean if a field has been set.

### SetNameWithheldNoteNil

`func (o *ApiV1OrgPipelineStagesGet200Response) SetNameWithheldNoteNil(b bool)`

 SetNameWithheldNoteNil sets the value for NameWithheldNote to be an explicit nil

### UnsetNameWithheldNote
`func (o *ApiV1OrgPipelineStagesGet200Response) UnsetNameWithheldNote()`

UnsetNameWithheldNote ensures that no value is present for NameWithheldNote, not even an explicit nil
### GetNameBuilt

`func (o *ApiV1OrgPipelineStagesGet200Response) GetNameBuilt() bool`

GetNameBuilt returns the NameBuilt field if non-nil, zero value otherwise.

### GetNameBuiltOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetNameBuiltOk() (*bool, bool)`

GetNameBuiltOk returns a tuple with the NameBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameBuilt

`func (o *ApiV1OrgPipelineStagesGet200Response) SetNameBuilt(v bool)`

SetNameBuilt sets NameBuilt field to given value.

### HasNameBuilt

`func (o *ApiV1OrgPipelineStagesGet200Response) HasNameBuilt() bool`

HasNameBuilt returns a boolean if a field has been set.

### GetTrialsByYear

`func (o *ApiV1OrgPipelineStagesGet200Response) GetTrialsByYear() []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner`

GetTrialsByYear returns the TrialsByYear field if non-nil, zero value otherwise.

### GetTrialsByYearOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetTrialsByYearOk() (*[]ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner, bool)`

GetTrialsByYearOk returns a tuple with the TrialsByYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialsByYear

`func (o *ApiV1OrgPipelineStagesGet200Response) SetTrialsByYear(v []ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner)`

SetTrialsByYear sets TrialsByYear field to given value.

### HasTrialsByYear

`func (o *ApiV1OrgPipelineStagesGet200Response) HasTrialsByYear() bool`

HasTrialsByYear returns a boolean if a field has been set.

### GetTrialStartForecast

`func (o *ApiV1OrgPipelineStagesGet200Response) GetTrialStartForecast() ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast`

GetTrialStartForecast returns the TrialStartForecast field if non-nil, zero value otherwise.

### GetTrialStartForecastOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetTrialStartForecastOk() (*ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast, bool)`

GetTrialStartForecastOk returns a tuple with the TrialStartForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStartForecast

`func (o *ApiV1OrgPipelineStagesGet200Response) SetTrialStartForecast(v ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast)`

SetTrialStartForecast sets TrialStartForecast field to given value.

### HasTrialStartForecast

`func (o *ApiV1OrgPipelineStagesGet200Response) HasTrialStartForecast() bool`

HasTrialStartForecast returns a boolean if a field has been set.

### GetStatement

`func (o *ApiV1OrgPipelineStagesGet200Response) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ApiV1OrgPipelineStagesGet200Response) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *ApiV1OrgPipelineStagesGet200Response) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetRecruitmentRollup

`func (o *ApiV1OrgPipelineStagesGet200Response) GetRecruitmentRollup() ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup`

GetRecruitmentRollup returns the RecruitmentRollup field if non-nil, zero value otherwise.

### GetRecruitmentRollupOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetRecruitmentRollupOk() (*ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup, bool)`

GetRecruitmentRollupOk returns a tuple with the RecruitmentRollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecruitmentRollup

`func (o *ApiV1OrgPipelineStagesGet200Response) SetRecruitmentRollup(v ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup)`

SetRecruitmentRollup sets RecruitmentRollup field to given value.

### HasRecruitmentRollup

`func (o *ApiV1OrgPipelineStagesGet200Response) HasRecruitmentRollup() bool`

HasRecruitmentRollup returns a boolean if a field has been set.

### GetRecruitmentStatement

`func (o *ApiV1OrgPipelineStagesGet200Response) GetRecruitmentStatement() string`

GetRecruitmentStatement returns the RecruitmentStatement field if non-nil, zero value otherwise.

### GetRecruitmentStatementOk

`func (o *ApiV1OrgPipelineStagesGet200Response) GetRecruitmentStatementOk() (*string, bool)`

GetRecruitmentStatementOk returns a tuple with the RecruitmentStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecruitmentStatement

`func (o *ApiV1OrgPipelineStagesGet200Response) SetRecruitmentStatement(v string)`

SetRecruitmentStatement sets RecruitmentStatement field to given value.

### HasRecruitmentStatement

`func (o *ApiV1OrgPipelineStagesGet200Response) HasRecruitmentStatement() bool`

HasRecruitmentStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


