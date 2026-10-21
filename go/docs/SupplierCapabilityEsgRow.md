# SupplierCapabilityEsgRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**VendorOrgId** | Pointer to **NullableString** | Legal-entity level: the contract research organization&#39;s organization record. | [optional] 
**VendorName** | Pointer to **NullableString** | The organization&#39;s name as filed. | [optional] 
**VendorCompanyId** | Pointer to **NullableString** | Company level: the contract research organization&#39;s company (null when unresolved). | [optional] 
**VendorCompanyName** | Pointer to **NullableString** | Display name of the contract research organization&#39;s company. | [optional] 
**VendorGroupId** | Pointer to **NullableString** | Group level: the contract research organization&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**VendorGroupName** | Pointer to **NullableString** | Display name of the contract research organization&#39;s group. | [optional] 
**CroAttributionMethod** | Pointer to **NullableString** | How the vendor was inferred for the trial: sponsor_match or collaborator_match (the organization is named on the registry record and is on the curated list) or llm_extraction. | [optional] 
**CroConfidence** | Pointer to **NullableFloat32** | Confidence of the inference, 0 to 1. | [optional] 
**Phase** | Pointer to **NullableString** | The registry&#39;s phase code of the trial (PHASE1, PHASE2, PHASE3 and so on). Null when the registry gives none. | [optional] 
**StartYear** | Pointer to **NullableInt32** | Start year of the trial. | [optional] 
**LeadSponsorName** | Pointer to **NullableString** | The trial&#39;s lead sponsor as filed. | [optional] 
**FacilityCountries** | Pointer to **[]string** | Countries of the trial&#39;s facilities. Filter on one country (for example facility_countries&#x3D;United States) to match trials that list it. | [optional] 
**ConditionIds** | Pointer to **[]string** | The MeSH (Medical Subject Headings) condition ids of the trial, direct and ancestor (D followed by six to nine digits). Filter on one id (for example condition_ids&#x3D;D003316) to scope the rows to a condition. | [optional] 
**EnrollmentDeltaKnown** | Pointer to **NullableBool** | True when the trial has both a planned and an actual enrollment, so its attainment is graded. | [optional] 
**EnrollmentMet** | Pointer to **NullableBool** | True when actual enrollment reached the plan. Null when not graded. | [optional] 
**EnrollmentAttainmentPct** | Pointer to **NullableFloat32** | Actual enrollment as a percentage of planned. Null when not graded. | [optional] 
**ClimateTargetState** | Pointer to **NullableString** | The vendor&#39;s climate target state in the open lane (extracted from its SEC filings): disclosed when a net-zero or near-term target is claimed, true_negative when only a no is filed, not_disclosed otherwise, including a vendor the lane does not cover. The same value on every row of the vendor. | [optional] 
**DiversityTargetState** | Pointer to **NullableString** | The vendor&#39;s diversity target state in the open lane, over the has_dei_targets claim, by the same rule. | [optional] 
**EsgClaimCount** | Pointer to **NullableInt32** | The vendor&#39;s extracted claims (any type) that carry a value. | [optional] 
**EsgLatestDisclosureDate** | Pointer to **NullableString** | The latest disclosure date among the vendor&#39;s extracted claims. | [optional] 
**SbtiClimateTargetState** | Pointer to **NullableString** | The vendor&#39;s climate target state in the Science Based Targets initiative lane, by the same rule. Null for a caller without the licence for this column; see withheld_columns. | [optional] 
**SbtiClaimCount** | Pointer to **NullableInt32** | The vendor&#39;s SBTi claims that carry a value. Null for a caller without the licence for this column; see withheld_columns. | [optional] 
**SbtiLatestDisclosureDate** | Pointer to **NullableString** | The latest disclosure date among the vendor&#39;s SBTi claims. Null for a caller without the licence for this column; see withheld_columns. | [optional] 

## Methods

### NewSupplierCapabilityEsgRow

`func NewSupplierCapabilityEsgRow() *SupplierCapabilityEsgRow`

NewSupplierCapabilityEsgRow instantiates a new SupplierCapabilityEsgRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierCapabilityEsgRowWithDefaults

`func NewSupplierCapabilityEsgRowWithDefaults() *SupplierCapabilityEsgRow`

NewSupplierCapabilityEsgRowWithDefaults instantiates a new SupplierCapabilityEsgRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *SupplierCapabilityEsgRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *SupplierCapabilityEsgRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *SupplierCapabilityEsgRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *SupplierCapabilityEsgRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *SupplierCapabilityEsgRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *SupplierCapabilityEsgRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetVendorOrgId

`func (o *SupplierCapabilityEsgRow) GetVendorOrgId() string`

GetVendorOrgId returns the VendorOrgId field if non-nil, zero value otherwise.

### GetVendorOrgIdOk

`func (o *SupplierCapabilityEsgRow) GetVendorOrgIdOk() (*string, bool)`

GetVendorOrgIdOk returns a tuple with the VendorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorOrgId

`func (o *SupplierCapabilityEsgRow) SetVendorOrgId(v string)`

SetVendorOrgId sets VendorOrgId field to given value.

### HasVendorOrgId

`func (o *SupplierCapabilityEsgRow) HasVendorOrgId() bool`

HasVendorOrgId returns a boolean if a field has been set.

### SetVendorOrgIdNil

`func (o *SupplierCapabilityEsgRow) SetVendorOrgIdNil(b bool)`

 SetVendorOrgIdNil sets the value for VendorOrgId to be an explicit nil

### UnsetVendorOrgId
`func (o *SupplierCapabilityEsgRow) UnsetVendorOrgId()`

UnsetVendorOrgId ensures that no value is present for VendorOrgId, not even an explicit nil
### GetVendorName

`func (o *SupplierCapabilityEsgRow) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *SupplierCapabilityEsgRow) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *SupplierCapabilityEsgRow) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *SupplierCapabilityEsgRow) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### SetVendorNameNil

`func (o *SupplierCapabilityEsgRow) SetVendorNameNil(b bool)`

 SetVendorNameNil sets the value for VendorName to be an explicit nil

### UnsetVendorName
`func (o *SupplierCapabilityEsgRow) UnsetVendorName()`

UnsetVendorName ensures that no value is present for VendorName, not even an explicit nil
### GetVendorCompanyId

`func (o *SupplierCapabilityEsgRow) GetVendorCompanyId() string`

GetVendorCompanyId returns the VendorCompanyId field if non-nil, zero value otherwise.

### GetVendorCompanyIdOk

`func (o *SupplierCapabilityEsgRow) GetVendorCompanyIdOk() (*string, bool)`

GetVendorCompanyIdOk returns a tuple with the VendorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCompanyId

`func (o *SupplierCapabilityEsgRow) SetVendorCompanyId(v string)`

SetVendorCompanyId sets VendorCompanyId field to given value.

### HasVendorCompanyId

`func (o *SupplierCapabilityEsgRow) HasVendorCompanyId() bool`

HasVendorCompanyId returns a boolean if a field has been set.

### SetVendorCompanyIdNil

`func (o *SupplierCapabilityEsgRow) SetVendorCompanyIdNil(b bool)`

 SetVendorCompanyIdNil sets the value for VendorCompanyId to be an explicit nil

### UnsetVendorCompanyId
`func (o *SupplierCapabilityEsgRow) UnsetVendorCompanyId()`

UnsetVendorCompanyId ensures that no value is present for VendorCompanyId, not even an explicit nil
### GetVendorCompanyName

`func (o *SupplierCapabilityEsgRow) GetVendorCompanyName() string`

GetVendorCompanyName returns the VendorCompanyName field if non-nil, zero value otherwise.

### GetVendorCompanyNameOk

`func (o *SupplierCapabilityEsgRow) GetVendorCompanyNameOk() (*string, bool)`

GetVendorCompanyNameOk returns a tuple with the VendorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCompanyName

`func (o *SupplierCapabilityEsgRow) SetVendorCompanyName(v string)`

SetVendorCompanyName sets VendorCompanyName field to given value.

### HasVendorCompanyName

`func (o *SupplierCapabilityEsgRow) HasVendorCompanyName() bool`

HasVendorCompanyName returns a boolean if a field has been set.

### SetVendorCompanyNameNil

`func (o *SupplierCapabilityEsgRow) SetVendorCompanyNameNil(b bool)`

 SetVendorCompanyNameNil sets the value for VendorCompanyName to be an explicit nil

### UnsetVendorCompanyName
`func (o *SupplierCapabilityEsgRow) UnsetVendorCompanyName()`

UnsetVendorCompanyName ensures that no value is present for VendorCompanyName, not even an explicit nil
### GetVendorGroupId

`func (o *SupplierCapabilityEsgRow) GetVendorGroupId() string`

GetVendorGroupId returns the VendorGroupId field if non-nil, zero value otherwise.

### GetVendorGroupIdOk

`func (o *SupplierCapabilityEsgRow) GetVendorGroupIdOk() (*string, bool)`

GetVendorGroupIdOk returns a tuple with the VendorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorGroupId

`func (o *SupplierCapabilityEsgRow) SetVendorGroupId(v string)`

SetVendorGroupId sets VendorGroupId field to given value.

### HasVendorGroupId

`func (o *SupplierCapabilityEsgRow) HasVendorGroupId() bool`

HasVendorGroupId returns a boolean if a field has been set.

### SetVendorGroupIdNil

`func (o *SupplierCapabilityEsgRow) SetVendorGroupIdNil(b bool)`

 SetVendorGroupIdNil sets the value for VendorGroupId to be an explicit nil

### UnsetVendorGroupId
`func (o *SupplierCapabilityEsgRow) UnsetVendorGroupId()`

UnsetVendorGroupId ensures that no value is present for VendorGroupId, not even an explicit nil
### GetVendorGroupName

`func (o *SupplierCapabilityEsgRow) GetVendorGroupName() string`

GetVendorGroupName returns the VendorGroupName field if non-nil, zero value otherwise.

### GetVendorGroupNameOk

`func (o *SupplierCapabilityEsgRow) GetVendorGroupNameOk() (*string, bool)`

GetVendorGroupNameOk returns a tuple with the VendorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorGroupName

`func (o *SupplierCapabilityEsgRow) SetVendorGroupName(v string)`

SetVendorGroupName sets VendorGroupName field to given value.

### HasVendorGroupName

`func (o *SupplierCapabilityEsgRow) HasVendorGroupName() bool`

HasVendorGroupName returns a boolean if a field has been set.

### SetVendorGroupNameNil

`func (o *SupplierCapabilityEsgRow) SetVendorGroupNameNil(b bool)`

 SetVendorGroupNameNil sets the value for VendorGroupName to be an explicit nil

### UnsetVendorGroupName
`func (o *SupplierCapabilityEsgRow) UnsetVendorGroupName()`

UnsetVendorGroupName ensures that no value is present for VendorGroupName, not even an explicit nil
### GetCroAttributionMethod

`func (o *SupplierCapabilityEsgRow) GetCroAttributionMethod() string`

GetCroAttributionMethod returns the CroAttributionMethod field if non-nil, zero value otherwise.

### GetCroAttributionMethodOk

`func (o *SupplierCapabilityEsgRow) GetCroAttributionMethodOk() (*string, bool)`

GetCroAttributionMethodOk returns a tuple with the CroAttributionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCroAttributionMethod

`func (o *SupplierCapabilityEsgRow) SetCroAttributionMethod(v string)`

SetCroAttributionMethod sets CroAttributionMethod field to given value.

### HasCroAttributionMethod

`func (o *SupplierCapabilityEsgRow) HasCroAttributionMethod() bool`

HasCroAttributionMethod returns a boolean if a field has been set.

### SetCroAttributionMethodNil

`func (o *SupplierCapabilityEsgRow) SetCroAttributionMethodNil(b bool)`

 SetCroAttributionMethodNil sets the value for CroAttributionMethod to be an explicit nil

### UnsetCroAttributionMethod
`func (o *SupplierCapabilityEsgRow) UnsetCroAttributionMethod()`

UnsetCroAttributionMethod ensures that no value is present for CroAttributionMethod, not even an explicit nil
### GetCroConfidence

`func (o *SupplierCapabilityEsgRow) GetCroConfidence() float32`

GetCroConfidence returns the CroConfidence field if non-nil, zero value otherwise.

### GetCroConfidenceOk

`func (o *SupplierCapabilityEsgRow) GetCroConfidenceOk() (*float32, bool)`

GetCroConfidenceOk returns a tuple with the CroConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCroConfidence

`func (o *SupplierCapabilityEsgRow) SetCroConfidence(v float32)`

SetCroConfidence sets CroConfidence field to given value.

### HasCroConfidence

`func (o *SupplierCapabilityEsgRow) HasCroConfidence() bool`

HasCroConfidence returns a boolean if a field has been set.

### SetCroConfidenceNil

`func (o *SupplierCapabilityEsgRow) SetCroConfidenceNil(b bool)`

 SetCroConfidenceNil sets the value for CroConfidence to be an explicit nil

### UnsetCroConfidence
`func (o *SupplierCapabilityEsgRow) UnsetCroConfidence()`

UnsetCroConfidence ensures that no value is present for CroConfidence, not even an explicit nil
### GetPhase

`func (o *SupplierCapabilityEsgRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SupplierCapabilityEsgRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SupplierCapabilityEsgRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *SupplierCapabilityEsgRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *SupplierCapabilityEsgRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *SupplierCapabilityEsgRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetStartYear

`func (o *SupplierCapabilityEsgRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *SupplierCapabilityEsgRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *SupplierCapabilityEsgRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *SupplierCapabilityEsgRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *SupplierCapabilityEsgRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *SupplierCapabilityEsgRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetLeadSponsorName

`func (o *SupplierCapabilityEsgRow) GetLeadSponsorName() string`

GetLeadSponsorName returns the LeadSponsorName field if non-nil, zero value otherwise.

### GetLeadSponsorNameOk

`func (o *SupplierCapabilityEsgRow) GetLeadSponsorNameOk() (*string, bool)`

GetLeadSponsorNameOk returns a tuple with the LeadSponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorName

`func (o *SupplierCapabilityEsgRow) SetLeadSponsorName(v string)`

SetLeadSponsorName sets LeadSponsorName field to given value.

### HasLeadSponsorName

`func (o *SupplierCapabilityEsgRow) HasLeadSponsorName() bool`

HasLeadSponsorName returns a boolean if a field has been set.

### SetLeadSponsorNameNil

`func (o *SupplierCapabilityEsgRow) SetLeadSponsorNameNil(b bool)`

 SetLeadSponsorNameNil sets the value for LeadSponsorName to be an explicit nil

### UnsetLeadSponsorName
`func (o *SupplierCapabilityEsgRow) UnsetLeadSponsorName()`

UnsetLeadSponsorName ensures that no value is present for LeadSponsorName, not even an explicit nil
### GetFacilityCountries

`func (o *SupplierCapabilityEsgRow) GetFacilityCountries() []string`

GetFacilityCountries returns the FacilityCountries field if non-nil, zero value otherwise.

### GetFacilityCountriesOk

`func (o *SupplierCapabilityEsgRow) GetFacilityCountriesOk() (*[]string, bool)`

GetFacilityCountriesOk returns a tuple with the FacilityCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountries

`func (o *SupplierCapabilityEsgRow) SetFacilityCountries(v []string)`

SetFacilityCountries sets FacilityCountries field to given value.

### HasFacilityCountries

`func (o *SupplierCapabilityEsgRow) HasFacilityCountries() bool`

HasFacilityCountries returns a boolean if a field has been set.

### SetFacilityCountriesNil

`func (o *SupplierCapabilityEsgRow) SetFacilityCountriesNil(b bool)`

 SetFacilityCountriesNil sets the value for FacilityCountries to be an explicit nil

### UnsetFacilityCountries
`func (o *SupplierCapabilityEsgRow) UnsetFacilityCountries()`

UnsetFacilityCountries ensures that no value is present for FacilityCountries, not even an explicit nil
### GetConditionIds

`func (o *SupplierCapabilityEsgRow) GetConditionIds() []string`

GetConditionIds returns the ConditionIds field if non-nil, zero value otherwise.

### GetConditionIdsOk

`func (o *SupplierCapabilityEsgRow) GetConditionIdsOk() (*[]string, bool)`

GetConditionIdsOk returns a tuple with the ConditionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIds

`func (o *SupplierCapabilityEsgRow) SetConditionIds(v []string)`

SetConditionIds sets ConditionIds field to given value.

### HasConditionIds

`func (o *SupplierCapabilityEsgRow) HasConditionIds() bool`

HasConditionIds returns a boolean if a field has been set.

### SetConditionIdsNil

`func (o *SupplierCapabilityEsgRow) SetConditionIdsNil(b bool)`

 SetConditionIdsNil sets the value for ConditionIds to be an explicit nil

### UnsetConditionIds
`func (o *SupplierCapabilityEsgRow) UnsetConditionIds()`

UnsetConditionIds ensures that no value is present for ConditionIds, not even an explicit nil
### GetEnrollmentDeltaKnown

`func (o *SupplierCapabilityEsgRow) GetEnrollmentDeltaKnown() bool`

GetEnrollmentDeltaKnown returns the EnrollmentDeltaKnown field if non-nil, zero value otherwise.

### GetEnrollmentDeltaKnownOk

`func (o *SupplierCapabilityEsgRow) GetEnrollmentDeltaKnownOk() (*bool, bool)`

GetEnrollmentDeltaKnownOk returns a tuple with the EnrollmentDeltaKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentDeltaKnown

`func (o *SupplierCapabilityEsgRow) SetEnrollmentDeltaKnown(v bool)`

SetEnrollmentDeltaKnown sets EnrollmentDeltaKnown field to given value.

### HasEnrollmentDeltaKnown

`func (o *SupplierCapabilityEsgRow) HasEnrollmentDeltaKnown() bool`

HasEnrollmentDeltaKnown returns a boolean if a field has been set.

### SetEnrollmentDeltaKnownNil

`func (o *SupplierCapabilityEsgRow) SetEnrollmentDeltaKnownNil(b bool)`

 SetEnrollmentDeltaKnownNil sets the value for EnrollmentDeltaKnown to be an explicit nil

### UnsetEnrollmentDeltaKnown
`func (o *SupplierCapabilityEsgRow) UnsetEnrollmentDeltaKnown()`

UnsetEnrollmentDeltaKnown ensures that no value is present for EnrollmentDeltaKnown, not even an explicit nil
### GetEnrollmentMet

`func (o *SupplierCapabilityEsgRow) GetEnrollmentMet() bool`

GetEnrollmentMet returns the EnrollmentMet field if non-nil, zero value otherwise.

### GetEnrollmentMetOk

`func (o *SupplierCapabilityEsgRow) GetEnrollmentMetOk() (*bool, bool)`

GetEnrollmentMetOk returns a tuple with the EnrollmentMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMet

`func (o *SupplierCapabilityEsgRow) SetEnrollmentMet(v bool)`

SetEnrollmentMet sets EnrollmentMet field to given value.

### HasEnrollmentMet

`func (o *SupplierCapabilityEsgRow) HasEnrollmentMet() bool`

HasEnrollmentMet returns a boolean if a field has been set.

### SetEnrollmentMetNil

`func (o *SupplierCapabilityEsgRow) SetEnrollmentMetNil(b bool)`

 SetEnrollmentMetNil sets the value for EnrollmentMet to be an explicit nil

### UnsetEnrollmentMet
`func (o *SupplierCapabilityEsgRow) UnsetEnrollmentMet()`

UnsetEnrollmentMet ensures that no value is present for EnrollmentMet, not even an explicit nil
### GetEnrollmentAttainmentPct

`func (o *SupplierCapabilityEsgRow) GetEnrollmentAttainmentPct() float32`

GetEnrollmentAttainmentPct returns the EnrollmentAttainmentPct field if non-nil, zero value otherwise.

### GetEnrollmentAttainmentPctOk

`func (o *SupplierCapabilityEsgRow) GetEnrollmentAttainmentPctOk() (*float32, bool)`

GetEnrollmentAttainmentPctOk returns a tuple with the EnrollmentAttainmentPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentAttainmentPct

`func (o *SupplierCapabilityEsgRow) SetEnrollmentAttainmentPct(v float32)`

SetEnrollmentAttainmentPct sets EnrollmentAttainmentPct field to given value.

### HasEnrollmentAttainmentPct

`func (o *SupplierCapabilityEsgRow) HasEnrollmentAttainmentPct() bool`

HasEnrollmentAttainmentPct returns a boolean if a field has been set.

### SetEnrollmentAttainmentPctNil

`func (o *SupplierCapabilityEsgRow) SetEnrollmentAttainmentPctNil(b bool)`

 SetEnrollmentAttainmentPctNil sets the value for EnrollmentAttainmentPct to be an explicit nil

### UnsetEnrollmentAttainmentPct
`func (o *SupplierCapabilityEsgRow) UnsetEnrollmentAttainmentPct()`

UnsetEnrollmentAttainmentPct ensures that no value is present for EnrollmentAttainmentPct, not even an explicit nil
### GetClimateTargetState

`func (o *SupplierCapabilityEsgRow) GetClimateTargetState() string`

GetClimateTargetState returns the ClimateTargetState field if non-nil, zero value otherwise.

### GetClimateTargetStateOk

`func (o *SupplierCapabilityEsgRow) GetClimateTargetStateOk() (*string, bool)`

GetClimateTargetStateOk returns a tuple with the ClimateTargetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClimateTargetState

`func (o *SupplierCapabilityEsgRow) SetClimateTargetState(v string)`

SetClimateTargetState sets ClimateTargetState field to given value.

### HasClimateTargetState

`func (o *SupplierCapabilityEsgRow) HasClimateTargetState() bool`

HasClimateTargetState returns a boolean if a field has been set.

### SetClimateTargetStateNil

`func (o *SupplierCapabilityEsgRow) SetClimateTargetStateNil(b bool)`

 SetClimateTargetStateNil sets the value for ClimateTargetState to be an explicit nil

### UnsetClimateTargetState
`func (o *SupplierCapabilityEsgRow) UnsetClimateTargetState()`

UnsetClimateTargetState ensures that no value is present for ClimateTargetState, not even an explicit nil
### GetDiversityTargetState

`func (o *SupplierCapabilityEsgRow) GetDiversityTargetState() string`

GetDiversityTargetState returns the DiversityTargetState field if non-nil, zero value otherwise.

### GetDiversityTargetStateOk

`func (o *SupplierCapabilityEsgRow) GetDiversityTargetStateOk() (*string, bool)`

GetDiversityTargetStateOk returns a tuple with the DiversityTargetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversityTargetState

`func (o *SupplierCapabilityEsgRow) SetDiversityTargetState(v string)`

SetDiversityTargetState sets DiversityTargetState field to given value.

### HasDiversityTargetState

`func (o *SupplierCapabilityEsgRow) HasDiversityTargetState() bool`

HasDiversityTargetState returns a boolean if a field has been set.

### SetDiversityTargetStateNil

`func (o *SupplierCapabilityEsgRow) SetDiversityTargetStateNil(b bool)`

 SetDiversityTargetStateNil sets the value for DiversityTargetState to be an explicit nil

### UnsetDiversityTargetState
`func (o *SupplierCapabilityEsgRow) UnsetDiversityTargetState()`

UnsetDiversityTargetState ensures that no value is present for DiversityTargetState, not even an explicit nil
### GetEsgClaimCount

`func (o *SupplierCapabilityEsgRow) GetEsgClaimCount() int32`

GetEsgClaimCount returns the EsgClaimCount field if non-nil, zero value otherwise.

### GetEsgClaimCountOk

`func (o *SupplierCapabilityEsgRow) GetEsgClaimCountOk() (*int32, bool)`

GetEsgClaimCountOk returns a tuple with the EsgClaimCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgClaimCount

`func (o *SupplierCapabilityEsgRow) SetEsgClaimCount(v int32)`

SetEsgClaimCount sets EsgClaimCount field to given value.

### HasEsgClaimCount

`func (o *SupplierCapabilityEsgRow) HasEsgClaimCount() bool`

HasEsgClaimCount returns a boolean if a field has been set.

### SetEsgClaimCountNil

`func (o *SupplierCapabilityEsgRow) SetEsgClaimCountNil(b bool)`

 SetEsgClaimCountNil sets the value for EsgClaimCount to be an explicit nil

### UnsetEsgClaimCount
`func (o *SupplierCapabilityEsgRow) UnsetEsgClaimCount()`

UnsetEsgClaimCount ensures that no value is present for EsgClaimCount, not even an explicit nil
### GetEsgLatestDisclosureDate

`func (o *SupplierCapabilityEsgRow) GetEsgLatestDisclosureDate() string`

GetEsgLatestDisclosureDate returns the EsgLatestDisclosureDate field if non-nil, zero value otherwise.

### GetEsgLatestDisclosureDateOk

`func (o *SupplierCapabilityEsgRow) GetEsgLatestDisclosureDateOk() (*string, bool)`

GetEsgLatestDisclosureDateOk returns a tuple with the EsgLatestDisclosureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgLatestDisclosureDate

`func (o *SupplierCapabilityEsgRow) SetEsgLatestDisclosureDate(v string)`

SetEsgLatestDisclosureDate sets EsgLatestDisclosureDate field to given value.

### HasEsgLatestDisclosureDate

`func (o *SupplierCapabilityEsgRow) HasEsgLatestDisclosureDate() bool`

HasEsgLatestDisclosureDate returns a boolean if a field has been set.

### SetEsgLatestDisclosureDateNil

`func (o *SupplierCapabilityEsgRow) SetEsgLatestDisclosureDateNil(b bool)`

 SetEsgLatestDisclosureDateNil sets the value for EsgLatestDisclosureDate to be an explicit nil

### UnsetEsgLatestDisclosureDate
`func (o *SupplierCapabilityEsgRow) UnsetEsgLatestDisclosureDate()`

UnsetEsgLatestDisclosureDate ensures that no value is present for EsgLatestDisclosureDate, not even an explicit nil
### GetSbtiClimateTargetState

`func (o *SupplierCapabilityEsgRow) GetSbtiClimateTargetState() string`

GetSbtiClimateTargetState returns the SbtiClimateTargetState field if non-nil, zero value otherwise.

### GetSbtiClimateTargetStateOk

`func (o *SupplierCapabilityEsgRow) GetSbtiClimateTargetStateOk() (*string, bool)`

GetSbtiClimateTargetStateOk returns a tuple with the SbtiClimateTargetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbtiClimateTargetState

`func (o *SupplierCapabilityEsgRow) SetSbtiClimateTargetState(v string)`

SetSbtiClimateTargetState sets SbtiClimateTargetState field to given value.

### HasSbtiClimateTargetState

`func (o *SupplierCapabilityEsgRow) HasSbtiClimateTargetState() bool`

HasSbtiClimateTargetState returns a boolean if a field has been set.

### SetSbtiClimateTargetStateNil

`func (o *SupplierCapabilityEsgRow) SetSbtiClimateTargetStateNil(b bool)`

 SetSbtiClimateTargetStateNil sets the value for SbtiClimateTargetState to be an explicit nil

### UnsetSbtiClimateTargetState
`func (o *SupplierCapabilityEsgRow) UnsetSbtiClimateTargetState()`

UnsetSbtiClimateTargetState ensures that no value is present for SbtiClimateTargetState, not even an explicit nil
### GetSbtiClaimCount

`func (o *SupplierCapabilityEsgRow) GetSbtiClaimCount() int32`

GetSbtiClaimCount returns the SbtiClaimCount field if non-nil, zero value otherwise.

### GetSbtiClaimCountOk

`func (o *SupplierCapabilityEsgRow) GetSbtiClaimCountOk() (*int32, bool)`

GetSbtiClaimCountOk returns a tuple with the SbtiClaimCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbtiClaimCount

`func (o *SupplierCapabilityEsgRow) SetSbtiClaimCount(v int32)`

SetSbtiClaimCount sets SbtiClaimCount field to given value.

### HasSbtiClaimCount

`func (o *SupplierCapabilityEsgRow) HasSbtiClaimCount() bool`

HasSbtiClaimCount returns a boolean if a field has been set.

### SetSbtiClaimCountNil

`func (o *SupplierCapabilityEsgRow) SetSbtiClaimCountNil(b bool)`

 SetSbtiClaimCountNil sets the value for SbtiClaimCount to be an explicit nil

### UnsetSbtiClaimCount
`func (o *SupplierCapabilityEsgRow) UnsetSbtiClaimCount()`

UnsetSbtiClaimCount ensures that no value is present for SbtiClaimCount, not even an explicit nil
### GetSbtiLatestDisclosureDate

`func (o *SupplierCapabilityEsgRow) GetSbtiLatestDisclosureDate() string`

GetSbtiLatestDisclosureDate returns the SbtiLatestDisclosureDate field if non-nil, zero value otherwise.

### GetSbtiLatestDisclosureDateOk

`func (o *SupplierCapabilityEsgRow) GetSbtiLatestDisclosureDateOk() (*string, bool)`

GetSbtiLatestDisclosureDateOk returns a tuple with the SbtiLatestDisclosureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbtiLatestDisclosureDate

`func (o *SupplierCapabilityEsgRow) SetSbtiLatestDisclosureDate(v string)`

SetSbtiLatestDisclosureDate sets SbtiLatestDisclosureDate field to given value.

### HasSbtiLatestDisclosureDate

`func (o *SupplierCapabilityEsgRow) HasSbtiLatestDisclosureDate() bool`

HasSbtiLatestDisclosureDate returns a boolean if a field has been set.

### SetSbtiLatestDisclosureDateNil

`func (o *SupplierCapabilityEsgRow) SetSbtiLatestDisclosureDateNil(b bool)`

 SetSbtiLatestDisclosureDateNil sets the value for SbtiLatestDisclosureDate to be an explicit nil

### UnsetSbtiLatestDisclosureDate
`func (o *SupplierCapabilityEsgRow) UnsetSbtiLatestDisclosureDate()`

UnsetSbtiLatestDisclosureDate ensures that no value is present for SbtiLatestDisclosureDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


