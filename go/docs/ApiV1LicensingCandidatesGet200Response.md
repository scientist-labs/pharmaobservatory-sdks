# ApiV1LicensingCandidatesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | **string** |  | 
**ConditionName** | **string** |  | 
**Candidates** | [**[]ApiV1LicensingCandidatesGet200ResponseCandidatesInner**](ApiV1LicensingCandidatesGet200ResponseCandidatesInner.md) |  | 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Restriction** | [**ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction**](ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction.md) |  | 

## Methods

### NewApiV1LicensingCandidatesGet200Response

`func NewApiV1LicensingCandidatesGet200Response(conditionId string, conditionName string, candidates []ApiV1LicensingCandidatesGet200ResponseCandidatesInner, buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, restriction ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction, ) *ApiV1LicensingCandidatesGet200Response`

NewApiV1LicensingCandidatesGet200Response instantiates a new ApiV1LicensingCandidatesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1LicensingCandidatesGet200ResponseWithDefaults

`func NewApiV1LicensingCandidatesGet200ResponseWithDefaults() *ApiV1LicensingCandidatesGet200Response`

NewApiV1LicensingCandidatesGet200ResponseWithDefaults instantiates a new ApiV1LicensingCandidatesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ApiV1LicensingCandidatesGet200Response) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ApiV1LicensingCandidatesGet200Response) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ApiV1LicensingCandidatesGet200Response) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.


### GetConditionName

`func (o *ApiV1LicensingCandidatesGet200Response) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ApiV1LicensingCandidatesGet200Response) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ApiV1LicensingCandidatesGet200Response) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.


### GetCandidates

`func (o *ApiV1LicensingCandidatesGet200Response) GetCandidates() []ApiV1LicensingCandidatesGet200ResponseCandidatesInner`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *ApiV1LicensingCandidatesGet200Response) GetCandidatesOk() (*[]ApiV1LicensingCandidatesGet200ResponseCandidatesInner, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *ApiV1LicensingCandidatesGet200Response) SetCandidates(v []ApiV1LicensingCandidatesGet200ResponseCandidatesInner)`

SetCandidates sets Candidates field to given value.


### GetBuildState

`func (o *ApiV1LicensingCandidatesGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1LicensingCandidatesGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1LicensingCandidatesGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetRestriction

`func (o *ApiV1LicensingCandidatesGet200Response) GetRestriction() ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ApiV1LicensingCandidatesGet200Response) GetRestrictionOk() (*ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ApiV1LicensingCandidatesGet200Response) SetRestriction(v ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction)`

SetRestriction sets Restriction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


