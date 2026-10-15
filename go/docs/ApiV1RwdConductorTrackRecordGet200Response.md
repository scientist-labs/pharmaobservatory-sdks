# ApiV1RwdConductorTrackRecordGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstitutionName** | **NullableString** |  | 
**HmaEmaId** | **NullableString** |  | 
**CanonicalOrgId** | **NullableString** |  | 
**Found** | **bool** |  | 
**Candidates** | **[]interface{}** |  | 
**Conductor** | [**NullableApiV1RwdConductorTrackRecordGet200ResponseConductor**](ApiV1RwdConductorTrackRecordGet200ResponseConductor.md) |  | 
**DeliveryByMandate** | [**NullableApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandate**](ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandate.md) |  | 
**Footprint** | [**NullableApiV1RwdConductorTrackRecordGet200ResponseFootprint**](ApiV1RwdConductorTrackRecordGet200ResponseFootprint.md) |  | 
**CoverageCaveat** | **string** |  | 
**TierProvisionalNote** | **string** |  | 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Restriction** | [**ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction**](ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction.md) |  | 
**ResolutionConfidence** | [**ApiV1ResolveEntityGet200ResponseResolutionConfidence**](ApiV1ResolveEntityGet200ResponseResolutionConfidence.md) |  | 

## Methods

### NewApiV1RwdConductorTrackRecordGet200Response

`func NewApiV1RwdConductorTrackRecordGet200Response(institutionName NullableString, hmaEmaId NullableString, canonicalOrgId NullableString, found bool, candidates []interface{}, conductor NullableApiV1RwdConductorTrackRecordGet200ResponseConductor, deliveryByMandate NullableApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandate, footprint NullableApiV1RwdConductorTrackRecordGet200ResponseFootprint, coverageCaveat string, tierProvisionalNote string, buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, restriction ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction, resolutionConfidence ApiV1ResolveEntityGet200ResponseResolutionConfidence, ) *ApiV1RwdConductorTrackRecordGet200Response`

NewApiV1RwdConductorTrackRecordGet200Response instantiates a new ApiV1RwdConductorTrackRecordGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RwdConductorTrackRecordGet200ResponseWithDefaults

`func NewApiV1RwdConductorTrackRecordGet200ResponseWithDefaults() *ApiV1RwdConductorTrackRecordGet200Response`

NewApiV1RwdConductorTrackRecordGet200ResponseWithDefaults instantiates a new ApiV1RwdConductorTrackRecordGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutionName

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.


### SetInstitutionNameNil

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetInstitutionNameNil(b bool)`

 SetInstitutionNameNil sets the value for InstitutionName to be an explicit nil

### UnsetInstitutionName
`func (o *ApiV1RwdConductorTrackRecordGet200Response) UnsetInstitutionName()`

UnsetInstitutionName ensures that no value is present for InstitutionName, not even an explicit nil
### GetHmaEmaId

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetHmaEmaId() string`

GetHmaEmaId returns the HmaEmaId field if non-nil, zero value otherwise.

### GetHmaEmaIdOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetHmaEmaIdOk() (*string, bool)`

GetHmaEmaIdOk returns a tuple with the HmaEmaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmaEmaId

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetHmaEmaId(v string)`

SetHmaEmaId sets HmaEmaId field to given value.


### SetHmaEmaIdNil

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetHmaEmaIdNil(b bool)`

 SetHmaEmaIdNil sets the value for HmaEmaId to be an explicit nil

### UnsetHmaEmaId
`func (o *ApiV1RwdConductorTrackRecordGet200Response) UnsetHmaEmaId()`

UnsetHmaEmaId ensures that no value is present for HmaEmaId, not even an explicit nil
### GetCanonicalOrgId

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetCanonicalOrgId() string`

GetCanonicalOrgId returns the CanonicalOrgId field if non-nil, zero value otherwise.

### GetCanonicalOrgIdOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetCanonicalOrgIdOk() (*string, bool)`

GetCanonicalOrgIdOk returns a tuple with the CanonicalOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalOrgId

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetCanonicalOrgId(v string)`

SetCanonicalOrgId sets CanonicalOrgId field to given value.


### SetCanonicalOrgIdNil

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetCanonicalOrgIdNil(b bool)`

 SetCanonicalOrgIdNil sets the value for CanonicalOrgId to be an explicit nil

### UnsetCanonicalOrgId
`func (o *ApiV1RwdConductorTrackRecordGet200Response) UnsetCanonicalOrgId()`

UnsetCanonicalOrgId ensures that no value is present for CanonicalOrgId, not even an explicit nil
### GetFound

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetFound(v bool)`

SetFound sets Found field to given value.


### GetCandidates

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetCandidates() []interface{}`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetCandidatesOk() (*[]interface{}, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetCandidates(v []interface{})`

SetCandidates sets Candidates field to given value.


### GetConductor

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetConductor() ApiV1RwdConductorTrackRecordGet200ResponseConductor`

GetConductor returns the Conductor field if non-nil, zero value otherwise.

### GetConductorOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetConductorOk() (*ApiV1RwdConductorTrackRecordGet200ResponseConductor, bool)`

GetConductorOk returns a tuple with the Conductor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductor

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetConductor(v ApiV1RwdConductorTrackRecordGet200ResponseConductor)`

SetConductor sets Conductor field to given value.


### SetConductorNil

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetConductorNil(b bool)`

 SetConductorNil sets the value for Conductor to be an explicit nil

### UnsetConductor
`func (o *ApiV1RwdConductorTrackRecordGet200Response) UnsetConductor()`

UnsetConductor ensures that no value is present for Conductor, not even an explicit nil
### GetDeliveryByMandate

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetDeliveryByMandate() ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandate`

GetDeliveryByMandate returns the DeliveryByMandate field if non-nil, zero value otherwise.

### GetDeliveryByMandateOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetDeliveryByMandateOk() (*ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandate, bool)`

GetDeliveryByMandateOk returns a tuple with the DeliveryByMandate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryByMandate

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetDeliveryByMandate(v ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandate)`

SetDeliveryByMandate sets DeliveryByMandate field to given value.


### SetDeliveryByMandateNil

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetDeliveryByMandateNil(b bool)`

 SetDeliveryByMandateNil sets the value for DeliveryByMandate to be an explicit nil

### UnsetDeliveryByMandate
`func (o *ApiV1RwdConductorTrackRecordGet200Response) UnsetDeliveryByMandate()`

UnsetDeliveryByMandate ensures that no value is present for DeliveryByMandate, not even an explicit nil
### GetFootprint

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetFootprint() ApiV1RwdConductorTrackRecordGet200ResponseFootprint`

GetFootprint returns the Footprint field if non-nil, zero value otherwise.

### GetFootprintOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetFootprintOk() (*ApiV1RwdConductorTrackRecordGet200ResponseFootprint, bool)`

GetFootprintOk returns a tuple with the Footprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFootprint

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetFootprint(v ApiV1RwdConductorTrackRecordGet200ResponseFootprint)`

SetFootprint sets Footprint field to given value.


### SetFootprintNil

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetFootprintNil(b bool)`

 SetFootprintNil sets the value for Footprint to be an explicit nil

### UnsetFootprint
`func (o *ApiV1RwdConductorTrackRecordGet200Response) UnsetFootprint()`

UnsetFootprint ensures that no value is present for Footprint, not even an explicit nil
### GetCoverageCaveat

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetCoverageCaveat() string`

GetCoverageCaveat returns the CoverageCaveat field if non-nil, zero value otherwise.

### GetCoverageCaveatOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetCoverageCaveatOk() (*string, bool)`

GetCoverageCaveatOk returns a tuple with the CoverageCaveat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageCaveat

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetCoverageCaveat(v string)`

SetCoverageCaveat sets CoverageCaveat field to given value.


### GetTierProvisionalNote

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetTierProvisionalNote() string`

GetTierProvisionalNote returns the TierProvisionalNote field if non-nil, zero value otherwise.

### GetTierProvisionalNoteOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetTierProvisionalNoteOk() (*string, bool)`

GetTierProvisionalNoteOk returns a tuple with the TierProvisionalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierProvisionalNote

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetTierProvisionalNote(v string)`

SetTierProvisionalNote sets TierProvisionalNote field to given value.


### GetBuildState

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetRestriction

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetRestriction() ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetRestrictionOk() (*ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetRestriction(v ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction)`

SetRestriction sets Restriction field to given value.


### GetResolutionConfidence

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetResolutionConfidence() ApiV1ResolveEntityGet200ResponseResolutionConfidence`

GetResolutionConfidence returns the ResolutionConfidence field if non-nil, zero value otherwise.

### GetResolutionConfidenceOk

`func (o *ApiV1RwdConductorTrackRecordGet200Response) GetResolutionConfidenceOk() (*ApiV1ResolveEntityGet200ResponseResolutionConfidence, bool)`

GetResolutionConfidenceOk returns a tuple with the ResolutionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionConfidence

`func (o *ApiV1RwdConductorTrackRecordGet200Response) SetResolutionConfidence(v ApiV1ResolveEntityGet200ResponseResolutionConfidence)`

SetResolutionConfidence sets ResolutionConfidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


