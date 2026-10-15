# ApiV1RwdStudyScreenGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**ApiV1RwdStudyScreenGet200ResponseFilters**](ApiV1RwdStudyScreenGet200ResponseFilters.md) |  | 
**MatchedStudies** | **NullableInt32** |  | 
**RegisterTotal** | **NullableInt32** |  | 
**DeliveryDistribution** | [**NullableApiV1RwdStudyScreenGet200ResponseDeliveryDistribution**](ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution.md) |  | 
**CoverageCaveat** | **string** |  | 
**TierProvisionalNote** | **string** |  | 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Restriction** | [**ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction**](ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction.md) |  | 
**ConditionFilterBuilt** | Pointer to **bool** | false means a condition was named and the condition mart is not built in this warehouse, so the filter could not run and matched_studies, register_total and delivery_distribution are unknown (null), not zero. true means the filter ran, or no condition was named. | [optional] 

## Methods

### NewApiV1RwdStudyScreenGet200Response

`func NewApiV1RwdStudyScreenGet200Response(filters ApiV1RwdStudyScreenGet200ResponseFilters, matchedStudies NullableInt32, registerTotal NullableInt32, deliveryDistribution NullableApiV1RwdStudyScreenGet200ResponseDeliveryDistribution, coverageCaveat string, tierProvisionalNote string, buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, restriction ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction, ) *ApiV1RwdStudyScreenGet200Response`

NewApiV1RwdStudyScreenGet200Response instantiates a new ApiV1RwdStudyScreenGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RwdStudyScreenGet200ResponseWithDefaults

`func NewApiV1RwdStudyScreenGet200ResponseWithDefaults() *ApiV1RwdStudyScreenGet200Response`

NewApiV1RwdStudyScreenGet200ResponseWithDefaults instantiates a new ApiV1RwdStudyScreenGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ApiV1RwdStudyScreenGet200Response) GetFilters() ApiV1RwdStudyScreenGet200ResponseFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiV1RwdStudyScreenGet200Response) GetFiltersOk() (*ApiV1RwdStudyScreenGet200ResponseFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiV1RwdStudyScreenGet200Response) SetFilters(v ApiV1RwdStudyScreenGet200ResponseFilters)`

SetFilters sets Filters field to given value.


### GetMatchedStudies

`func (o *ApiV1RwdStudyScreenGet200Response) GetMatchedStudies() int32`

GetMatchedStudies returns the MatchedStudies field if non-nil, zero value otherwise.

### GetMatchedStudiesOk

`func (o *ApiV1RwdStudyScreenGet200Response) GetMatchedStudiesOk() (*int32, bool)`

GetMatchedStudiesOk returns a tuple with the MatchedStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedStudies

`func (o *ApiV1RwdStudyScreenGet200Response) SetMatchedStudies(v int32)`

SetMatchedStudies sets MatchedStudies field to given value.


### SetMatchedStudiesNil

`func (o *ApiV1RwdStudyScreenGet200Response) SetMatchedStudiesNil(b bool)`

 SetMatchedStudiesNil sets the value for MatchedStudies to be an explicit nil

### UnsetMatchedStudies
`func (o *ApiV1RwdStudyScreenGet200Response) UnsetMatchedStudies()`

UnsetMatchedStudies ensures that no value is present for MatchedStudies, not even an explicit nil
### GetRegisterTotal

`func (o *ApiV1RwdStudyScreenGet200Response) GetRegisterTotal() int32`

GetRegisterTotal returns the RegisterTotal field if non-nil, zero value otherwise.

### GetRegisterTotalOk

`func (o *ApiV1RwdStudyScreenGet200Response) GetRegisterTotalOk() (*int32, bool)`

GetRegisterTotalOk returns a tuple with the RegisterTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterTotal

`func (o *ApiV1RwdStudyScreenGet200Response) SetRegisterTotal(v int32)`

SetRegisterTotal sets RegisterTotal field to given value.


### SetRegisterTotalNil

`func (o *ApiV1RwdStudyScreenGet200Response) SetRegisterTotalNil(b bool)`

 SetRegisterTotalNil sets the value for RegisterTotal to be an explicit nil

### UnsetRegisterTotal
`func (o *ApiV1RwdStudyScreenGet200Response) UnsetRegisterTotal()`

UnsetRegisterTotal ensures that no value is present for RegisterTotal, not even an explicit nil
### GetDeliveryDistribution

`func (o *ApiV1RwdStudyScreenGet200Response) GetDeliveryDistribution() ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution`

GetDeliveryDistribution returns the DeliveryDistribution field if non-nil, zero value otherwise.

### GetDeliveryDistributionOk

`func (o *ApiV1RwdStudyScreenGet200Response) GetDeliveryDistributionOk() (*ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution, bool)`

GetDeliveryDistributionOk returns a tuple with the DeliveryDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDistribution

`func (o *ApiV1RwdStudyScreenGet200Response) SetDeliveryDistribution(v ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution)`

SetDeliveryDistribution sets DeliveryDistribution field to given value.


### SetDeliveryDistributionNil

`func (o *ApiV1RwdStudyScreenGet200Response) SetDeliveryDistributionNil(b bool)`

 SetDeliveryDistributionNil sets the value for DeliveryDistribution to be an explicit nil

### UnsetDeliveryDistribution
`func (o *ApiV1RwdStudyScreenGet200Response) UnsetDeliveryDistribution()`

UnsetDeliveryDistribution ensures that no value is present for DeliveryDistribution, not even an explicit nil
### GetCoverageCaveat

`func (o *ApiV1RwdStudyScreenGet200Response) GetCoverageCaveat() string`

GetCoverageCaveat returns the CoverageCaveat field if non-nil, zero value otherwise.

### GetCoverageCaveatOk

`func (o *ApiV1RwdStudyScreenGet200Response) GetCoverageCaveatOk() (*string, bool)`

GetCoverageCaveatOk returns a tuple with the CoverageCaveat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageCaveat

`func (o *ApiV1RwdStudyScreenGet200Response) SetCoverageCaveat(v string)`

SetCoverageCaveat sets CoverageCaveat field to given value.


### GetTierProvisionalNote

`func (o *ApiV1RwdStudyScreenGet200Response) GetTierProvisionalNote() string`

GetTierProvisionalNote returns the TierProvisionalNote field if non-nil, zero value otherwise.

### GetTierProvisionalNoteOk

`func (o *ApiV1RwdStudyScreenGet200Response) GetTierProvisionalNoteOk() (*string, bool)`

GetTierProvisionalNoteOk returns a tuple with the TierProvisionalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierProvisionalNote

`func (o *ApiV1RwdStudyScreenGet200Response) SetTierProvisionalNote(v string)`

SetTierProvisionalNote sets TierProvisionalNote field to given value.


### GetBuildState

`func (o *ApiV1RwdStudyScreenGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1RwdStudyScreenGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1RwdStudyScreenGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetRestriction

`func (o *ApiV1RwdStudyScreenGet200Response) GetRestriction() ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ApiV1RwdStudyScreenGet200Response) GetRestrictionOk() (*ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ApiV1RwdStudyScreenGet200Response) SetRestriction(v ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction)`

SetRestriction sets Restriction field to given value.


### GetConditionFilterBuilt

`func (o *ApiV1RwdStudyScreenGet200Response) GetConditionFilterBuilt() bool`

GetConditionFilterBuilt returns the ConditionFilterBuilt field if non-nil, zero value otherwise.

### GetConditionFilterBuiltOk

`func (o *ApiV1RwdStudyScreenGet200Response) GetConditionFilterBuiltOk() (*bool, bool)`

GetConditionFilterBuiltOk returns a tuple with the ConditionFilterBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionFilterBuilt

`func (o *ApiV1RwdStudyScreenGet200Response) SetConditionFilterBuilt(v bool)`

SetConditionFilterBuilt sets ConditionFilterBuilt field to given value.

### HasConditionFilterBuilt

`func (o *ApiV1RwdStudyScreenGet200Response) HasConditionFilterBuilt() bool`

HasConditionFilterBuilt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


