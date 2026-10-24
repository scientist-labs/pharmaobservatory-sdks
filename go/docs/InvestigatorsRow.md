# InvestigatorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] 
**InvestigatorPersonName** | Pointer to **NullableString** | The investigator&#39;s name as the registry filed it, resolved across the person&#39;s merged name fragments. | [optional] 
**InvestigatorAffiliation** | Pointer to **NullableString** | The investigator&#39;s affiliation as filed (the most-trialled affiliation across the merged fragments). | [optional] 
**Roles** | Pointer to **[]string** | Distinct registry roles the person holds across their trials (PRINCIPAL_INVESTIGATOR, STUDY_DIRECTOR, STUDY_CHAIR, SUB_INVESTIGATOR); empty when the registry names none. | [optional] 
**IsPrincipalInvestigator** | Pointer to **NullableBool** | True when the person is named principal investigator on at least one trial. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Distinct registered trials across the person&#39;s merged fragments. | [optional] 
**NTrialPublications** | Pointer to **NullableInt64** | Distinct publications the person&#39;s trials register (trial-attributed, not authorship). | [optional] 
**NAuthoredPublications** | Pointer to **NullableInt64** | Publications linked to the person as an author. | [optional] 
**MergedFragmentCount** | Pointer to **NullableInt64** | Name fragments the resolution merged into this person. | [optional] 
**NpiAnchored** | Pointer to **NullableBool** | True when the person&#39;s fragments were merged on a US National Provider Identifier. | [optional] 
**OrcidAnchored** | Pointer to **NullableBool** | True when the person&#39;s fragments were merged on an ORCID. | [optional] 
**IdentityAnchor** | Pointer to **NullableString** | npi when the person is NPI-anchored, orcid when ORCID-anchored only, none otherwise. | [optional] 
**PrimarySpecialty** | Pointer to **NullableString** | The clinician&#39;s primary specialty in the CMS Doctors and Clinicians file, upper-case as filed (null without an NPI or clinician record). | [optional] 
**HasNihFunding** | Pointer to **NullableBool** | True when the person holds at least one corroborated NIH RePORTER project. | [optional] 
**NihAwardDollars** | Pointer to **NullableFloat32** | Corroborated NIH RePORTER award dollars (zero when only uncorroborated links exist). | [optional] 
**NihProjectCount** | Pointer to **NullableInt64** | Corroborated NIH RePORTER projects. | [optional] 
**NihAwardDollarsInclUncorroborated** | Pointer to **NullableFloat32** | NIH RePORTER award dollars including projects linked on organization affinity alone. | [optional] 
**NihProjectCountInclUncorroborated** | Pointer to **NullableInt64** | NIH RePORTER projects including those linked on organization affinity alone. | [optional] 
**NihCorroborated** | Pointer to **NullableBool** | True when at least one contributing RePORTER link is confirmed by a trial or a publication; null when the person has no funding record. | [optional] 
**NihAwardCount** | Pointer to **NullableInt64** | Distinct NIH awards (null when the person has no funding record). | [optional] 
**NihFirstFiscalYear** | Pointer to **NullableInt32** | First fiscal year of the person&#39;s NIH projects. | [optional] 
**NihLastFiscalYear** | Pointer to **NullableInt32** | Last fiscal year of the person&#39;s NIH projects. | [optional] 
**NihTopIc** | Pointer to **NullableString** | The NIH institute or center that administers the most of the person&#39;s projects. | [optional] 
**HasHospitalAffiliation** | Pointer to **NullableBool** | True when the person has a CMS facility affiliation (NPI-keyed, so false for a person with no NPI by lookup, not by fact). | [optional] 
**NAffiliatedHospitals** | Pointer to **NullableInt64** | CMS hospitals the person is affiliated with. | [optional] 
**PrimaryCcn** | Pointer to **NullableString** | The lowest affiliated CMS Certification Number, the key into the hospitals dataset. | [optional] 
**HasExpertiseTopics** | Pointer to **NullableBool** | True when the person has authored papers with MeSH topics (investigator_topics has rows for them). | [optional] 

## Methods

### NewInvestigatorsRow

`func NewInvestigatorsRow() *InvestigatorsRow`

NewInvestigatorsRow instantiates a new InvestigatorsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigatorsRowWithDefaults

`func NewInvestigatorsRowWithDefaults() *InvestigatorsRow`

NewInvestigatorsRowWithDefaults instantiates a new InvestigatorsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestigatorPersonId

`func (o *InvestigatorsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *InvestigatorsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *InvestigatorsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *InvestigatorsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *InvestigatorsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *InvestigatorsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetInvestigatorPersonName

`func (o *InvestigatorsRow) GetInvestigatorPersonName() string`

GetInvestigatorPersonName returns the InvestigatorPersonName field if non-nil, zero value otherwise.

### GetInvestigatorPersonNameOk

`func (o *InvestigatorsRow) GetInvestigatorPersonNameOk() (*string, bool)`

GetInvestigatorPersonNameOk returns a tuple with the InvestigatorPersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonName

`func (o *InvestigatorsRow) SetInvestigatorPersonName(v string)`

SetInvestigatorPersonName sets InvestigatorPersonName field to given value.

### HasInvestigatorPersonName

`func (o *InvestigatorsRow) HasInvestigatorPersonName() bool`

HasInvestigatorPersonName returns a boolean if a field has been set.

### SetInvestigatorPersonNameNil

`func (o *InvestigatorsRow) SetInvestigatorPersonNameNil(b bool)`

 SetInvestigatorPersonNameNil sets the value for InvestigatorPersonName to be an explicit nil

### UnsetInvestigatorPersonName
`func (o *InvestigatorsRow) UnsetInvestigatorPersonName()`

UnsetInvestigatorPersonName ensures that no value is present for InvestigatorPersonName, not even an explicit nil
### GetInvestigatorAffiliation

`func (o *InvestigatorsRow) GetInvestigatorAffiliation() string`

GetInvestigatorAffiliation returns the InvestigatorAffiliation field if non-nil, zero value otherwise.

### GetInvestigatorAffiliationOk

`func (o *InvestigatorsRow) GetInvestigatorAffiliationOk() (*string, bool)`

GetInvestigatorAffiliationOk returns a tuple with the InvestigatorAffiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorAffiliation

`func (o *InvestigatorsRow) SetInvestigatorAffiliation(v string)`

SetInvestigatorAffiliation sets InvestigatorAffiliation field to given value.

### HasInvestigatorAffiliation

`func (o *InvestigatorsRow) HasInvestigatorAffiliation() bool`

HasInvestigatorAffiliation returns a boolean if a field has been set.

### SetInvestigatorAffiliationNil

`func (o *InvestigatorsRow) SetInvestigatorAffiliationNil(b bool)`

 SetInvestigatorAffiliationNil sets the value for InvestigatorAffiliation to be an explicit nil

### UnsetInvestigatorAffiliation
`func (o *InvestigatorsRow) UnsetInvestigatorAffiliation()`

UnsetInvestigatorAffiliation ensures that no value is present for InvestigatorAffiliation, not even an explicit nil
### GetRoles

`func (o *InvestigatorsRow) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InvestigatorsRow) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *InvestigatorsRow) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *InvestigatorsRow) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *InvestigatorsRow) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *InvestigatorsRow) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetIsPrincipalInvestigator

`func (o *InvestigatorsRow) GetIsPrincipalInvestigator() bool`

GetIsPrincipalInvestigator returns the IsPrincipalInvestigator field if non-nil, zero value otherwise.

### GetIsPrincipalInvestigatorOk

`func (o *InvestigatorsRow) GetIsPrincipalInvestigatorOk() (*bool, bool)`

GetIsPrincipalInvestigatorOk returns a tuple with the IsPrincipalInvestigator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrincipalInvestigator

`func (o *InvestigatorsRow) SetIsPrincipalInvestigator(v bool)`

SetIsPrincipalInvestigator sets IsPrincipalInvestigator field to given value.

### HasIsPrincipalInvestigator

`func (o *InvestigatorsRow) HasIsPrincipalInvestigator() bool`

HasIsPrincipalInvestigator returns a boolean if a field has been set.

### SetIsPrincipalInvestigatorNil

`func (o *InvestigatorsRow) SetIsPrincipalInvestigatorNil(b bool)`

 SetIsPrincipalInvestigatorNil sets the value for IsPrincipalInvestigator to be an explicit nil

### UnsetIsPrincipalInvestigator
`func (o *InvestigatorsRow) UnsetIsPrincipalInvestigator()`

UnsetIsPrincipalInvestigator ensures that no value is present for IsPrincipalInvestigator, not even an explicit nil
### GetNTrials

`func (o *InvestigatorsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *InvestigatorsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *InvestigatorsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *InvestigatorsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *InvestigatorsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *InvestigatorsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNTrialPublications

`func (o *InvestigatorsRow) GetNTrialPublications() int64`

GetNTrialPublications returns the NTrialPublications field if non-nil, zero value otherwise.

### GetNTrialPublicationsOk

`func (o *InvestigatorsRow) GetNTrialPublicationsOk() (*int64, bool)`

GetNTrialPublicationsOk returns a tuple with the NTrialPublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialPublications

`func (o *InvestigatorsRow) SetNTrialPublications(v int64)`

SetNTrialPublications sets NTrialPublications field to given value.

### HasNTrialPublications

`func (o *InvestigatorsRow) HasNTrialPublications() bool`

HasNTrialPublications returns a boolean if a field has been set.

### SetNTrialPublicationsNil

`func (o *InvestigatorsRow) SetNTrialPublicationsNil(b bool)`

 SetNTrialPublicationsNil sets the value for NTrialPublications to be an explicit nil

### UnsetNTrialPublications
`func (o *InvestigatorsRow) UnsetNTrialPublications()`

UnsetNTrialPublications ensures that no value is present for NTrialPublications, not even an explicit nil
### GetNAuthoredPublications

`func (o *InvestigatorsRow) GetNAuthoredPublications() int64`

GetNAuthoredPublications returns the NAuthoredPublications field if non-nil, zero value otherwise.

### GetNAuthoredPublicationsOk

`func (o *InvestigatorsRow) GetNAuthoredPublicationsOk() (*int64, bool)`

GetNAuthoredPublicationsOk returns a tuple with the NAuthoredPublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAuthoredPublications

`func (o *InvestigatorsRow) SetNAuthoredPublications(v int64)`

SetNAuthoredPublications sets NAuthoredPublications field to given value.

### HasNAuthoredPublications

`func (o *InvestigatorsRow) HasNAuthoredPublications() bool`

HasNAuthoredPublications returns a boolean if a field has been set.

### SetNAuthoredPublicationsNil

`func (o *InvestigatorsRow) SetNAuthoredPublicationsNil(b bool)`

 SetNAuthoredPublicationsNil sets the value for NAuthoredPublications to be an explicit nil

### UnsetNAuthoredPublications
`func (o *InvestigatorsRow) UnsetNAuthoredPublications()`

UnsetNAuthoredPublications ensures that no value is present for NAuthoredPublications, not even an explicit nil
### GetMergedFragmentCount

`func (o *InvestigatorsRow) GetMergedFragmentCount() int64`

GetMergedFragmentCount returns the MergedFragmentCount field if non-nil, zero value otherwise.

### GetMergedFragmentCountOk

`func (o *InvestigatorsRow) GetMergedFragmentCountOk() (*int64, bool)`

GetMergedFragmentCountOk returns a tuple with the MergedFragmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedFragmentCount

`func (o *InvestigatorsRow) SetMergedFragmentCount(v int64)`

SetMergedFragmentCount sets MergedFragmentCount field to given value.

### HasMergedFragmentCount

`func (o *InvestigatorsRow) HasMergedFragmentCount() bool`

HasMergedFragmentCount returns a boolean if a field has been set.

### SetMergedFragmentCountNil

`func (o *InvestigatorsRow) SetMergedFragmentCountNil(b bool)`

 SetMergedFragmentCountNil sets the value for MergedFragmentCount to be an explicit nil

### UnsetMergedFragmentCount
`func (o *InvestigatorsRow) UnsetMergedFragmentCount()`

UnsetMergedFragmentCount ensures that no value is present for MergedFragmentCount, not even an explicit nil
### GetNpiAnchored

`func (o *InvestigatorsRow) GetNpiAnchored() bool`

GetNpiAnchored returns the NpiAnchored field if non-nil, zero value otherwise.

### GetNpiAnchoredOk

`func (o *InvestigatorsRow) GetNpiAnchoredOk() (*bool, bool)`

GetNpiAnchoredOk returns a tuple with the NpiAnchored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpiAnchored

`func (o *InvestigatorsRow) SetNpiAnchored(v bool)`

SetNpiAnchored sets NpiAnchored field to given value.

### HasNpiAnchored

`func (o *InvestigatorsRow) HasNpiAnchored() bool`

HasNpiAnchored returns a boolean if a field has been set.

### SetNpiAnchoredNil

`func (o *InvestigatorsRow) SetNpiAnchoredNil(b bool)`

 SetNpiAnchoredNil sets the value for NpiAnchored to be an explicit nil

### UnsetNpiAnchored
`func (o *InvestigatorsRow) UnsetNpiAnchored()`

UnsetNpiAnchored ensures that no value is present for NpiAnchored, not even an explicit nil
### GetOrcidAnchored

`func (o *InvestigatorsRow) GetOrcidAnchored() bool`

GetOrcidAnchored returns the OrcidAnchored field if non-nil, zero value otherwise.

### GetOrcidAnchoredOk

`func (o *InvestigatorsRow) GetOrcidAnchoredOk() (*bool, bool)`

GetOrcidAnchoredOk returns a tuple with the OrcidAnchored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrcidAnchored

`func (o *InvestigatorsRow) SetOrcidAnchored(v bool)`

SetOrcidAnchored sets OrcidAnchored field to given value.

### HasOrcidAnchored

`func (o *InvestigatorsRow) HasOrcidAnchored() bool`

HasOrcidAnchored returns a boolean if a field has been set.

### SetOrcidAnchoredNil

`func (o *InvestigatorsRow) SetOrcidAnchoredNil(b bool)`

 SetOrcidAnchoredNil sets the value for OrcidAnchored to be an explicit nil

### UnsetOrcidAnchored
`func (o *InvestigatorsRow) UnsetOrcidAnchored()`

UnsetOrcidAnchored ensures that no value is present for OrcidAnchored, not even an explicit nil
### GetIdentityAnchor

`func (o *InvestigatorsRow) GetIdentityAnchor() string`

GetIdentityAnchor returns the IdentityAnchor field if non-nil, zero value otherwise.

### GetIdentityAnchorOk

`func (o *InvestigatorsRow) GetIdentityAnchorOk() (*string, bool)`

GetIdentityAnchorOk returns a tuple with the IdentityAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAnchor

`func (o *InvestigatorsRow) SetIdentityAnchor(v string)`

SetIdentityAnchor sets IdentityAnchor field to given value.

### HasIdentityAnchor

`func (o *InvestigatorsRow) HasIdentityAnchor() bool`

HasIdentityAnchor returns a boolean if a field has been set.

### SetIdentityAnchorNil

`func (o *InvestigatorsRow) SetIdentityAnchorNil(b bool)`

 SetIdentityAnchorNil sets the value for IdentityAnchor to be an explicit nil

### UnsetIdentityAnchor
`func (o *InvestigatorsRow) UnsetIdentityAnchor()`

UnsetIdentityAnchor ensures that no value is present for IdentityAnchor, not even an explicit nil
### GetPrimarySpecialty

`func (o *InvestigatorsRow) GetPrimarySpecialty() string`

GetPrimarySpecialty returns the PrimarySpecialty field if non-nil, zero value otherwise.

### GetPrimarySpecialtyOk

`func (o *InvestigatorsRow) GetPrimarySpecialtyOk() (*string, bool)`

GetPrimarySpecialtyOk returns a tuple with the PrimarySpecialty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySpecialty

`func (o *InvestigatorsRow) SetPrimarySpecialty(v string)`

SetPrimarySpecialty sets PrimarySpecialty field to given value.

### HasPrimarySpecialty

`func (o *InvestigatorsRow) HasPrimarySpecialty() bool`

HasPrimarySpecialty returns a boolean if a field has been set.

### SetPrimarySpecialtyNil

`func (o *InvestigatorsRow) SetPrimarySpecialtyNil(b bool)`

 SetPrimarySpecialtyNil sets the value for PrimarySpecialty to be an explicit nil

### UnsetPrimarySpecialty
`func (o *InvestigatorsRow) UnsetPrimarySpecialty()`

UnsetPrimarySpecialty ensures that no value is present for PrimarySpecialty, not even an explicit nil
### GetHasNihFunding

`func (o *InvestigatorsRow) GetHasNihFunding() bool`

GetHasNihFunding returns the HasNihFunding field if non-nil, zero value otherwise.

### GetHasNihFundingOk

`func (o *InvestigatorsRow) GetHasNihFundingOk() (*bool, bool)`

GetHasNihFundingOk returns a tuple with the HasNihFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNihFunding

`func (o *InvestigatorsRow) SetHasNihFunding(v bool)`

SetHasNihFunding sets HasNihFunding field to given value.

### HasHasNihFunding

`func (o *InvestigatorsRow) HasHasNihFunding() bool`

HasHasNihFunding returns a boolean if a field has been set.

### SetHasNihFundingNil

`func (o *InvestigatorsRow) SetHasNihFundingNil(b bool)`

 SetHasNihFundingNil sets the value for HasNihFunding to be an explicit nil

### UnsetHasNihFunding
`func (o *InvestigatorsRow) UnsetHasNihFunding()`

UnsetHasNihFunding ensures that no value is present for HasNihFunding, not even an explicit nil
### GetNihAwardDollars

`func (o *InvestigatorsRow) GetNihAwardDollars() float32`

GetNihAwardDollars returns the NihAwardDollars field if non-nil, zero value otherwise.

### GetNihAwardDollarsOk

`func (o *InvestigatorsRow) GetNihAwardDollarsOk() (*float32, bool)`

GetNihAwardDollarsOk returns a tuple with the NihAwardDollars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNihAwardDollars

`func (o *InvestigatorsRow) SetNihAwardDollars(v float32)`

SetNihAwardDollars sets NihAwardDollars field to given value.

### HasNihAwardDollars

`func (o *InvestigatorsRow) HasNihAwardDollars() bool`

HasNihAwardDollars returns a boolean if a field has been set.

### SetNihAwardDollarsNil

`func (o *InvestigatorsRow) SetNihAwardDollarsNil(b bool)`

 SetNihAwardDollarsNil sets the value for NihAwardDollars to be an explicit nil

### UnsetNihAwardDollars
`func (o *InvestigatorsRow) UnsetNihAwardDollars()`

UnsetNihAwardDollars ensures that no value is present for NihAwardDollars, not even an explicit nil
### GetNihProjectCount

`func (o *InvestigatorsRow) GetNihProjectCount() int64`

GetNihProjectCount returns the NihProjectCount field if non-nil, zero value otherwise.

### GetNihProjectCountOk

`func (o *InvestigatorsRow) GetNihProjectCountOk() (*int64, bool)`

GetNihProjectCountOk returns a tuple with the NihProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNihProjectCount

`func (o *InvestigatorsRow) SetNihProjectCount(v int64)`

SetNihProjectCount sets NihProjectCount field to given value.

### HasNihProjectCount

`func (o *InvestigatorsRow) HasNihProjectCount() bool`

HasNihProjectCount returns a boolean if a field has been set.

### SetNihProjectCountNil

`func (o *InvestigatorsRow) SetNihProjectCountNil(b bool)`

 SetNihProjectCountNil sets the value for NihProjectCount to be an explicit nil

### UnsetNihProjectCount
`func (o *InvestigatorsRow) UnsetNihProjectCount()`

UnsetNihProjectCount ensures that no value is present for NihProjectCount, not even an explicit nil
### GetNihAwardDollarsInclUncorroborated

`func (o *InvestigatorsRow) GetNihAwardDollarsInclUncorroborated() float32`

GetNihAwardDollarsInclUncorroborated returns the NihAwardDollarsInclUncorroborated field if non-nil, zero value otherwise.

### GetNihAwardDollarsInclUncorroboratedOk

`func (o *InvestigatorsRow) GetNihAwardDollarsInclUncorroboratedOk() (*float32, bool)`

GetNihAwardDollarsInclUncorroboratedOk returns a tuple with the NihAwardDollarsInclUncorroborated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNihAwardDollarsInclUncorroborated

`func (o *InvestigatorsRow) SetNihAwardDollarsInclUncorroborated(v float32)`

SetNihAwardDollarsInclUncorroborated sets NihAwardDollarsInclUncorroborated field to given value.

### HasNihAwardDollarsInclUncorroborated

`func (o *InvestigatorsRow) HasNihAwardDollarsInclUncorroborated() bool`

HasNihAwardDollarsInclUncorroborated returns a boolean if a field has been set.

### SetNihAwardDollarsInclUncorroboratedNil

`func (o *InvestigatorsRow) SetNihAwardDollarsInclUncorroboratedNil(b bool)`

 SetNihAwardDollarsInclUncorroboratedNil sets the value for NihAwardDollarsInclUncorroborated to be an explicit nil

### UnsetNihAwardDollarsInclUncorroborated
`func (o *InvestigatorsRow) UnsetNihAwardDollarsInclUncorroborated()`

UnsetNihAwardDollarsInclUncorroborated ensures that no value is present for NihAwardDollarsInclUncorroborated, not even an explicit nil
### GetNihProjectCountInclUncorroborated

`func (o *InvestigatorsRow) GetNihProjectCountInclUncorroborated() int64`

GetNihProjectCountInclUncorroborated returns the NihProjectCountInclUncorroborated field if non-nil, zero value otherwise.

### GetNihProjectCountInclUncorroboratedOk

`func (o *InvestigatorsRow) GetNihProjectCountInclUncorroboratedOk() (*int64, bool)`

GetNihProjectCountInclUncorroboratedOk returns a tuple with the NihProjectCountInclUncorroborated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNihProjectCountInclUncorroborated

`func (o *InvestigatorsRow) SetNihProjectCountInclUncorroborated(v int64)`

SetNihProjectCountInclUncorroborated sets NihProjectCountInclUncorroborated field to given value.

### HasNihProjectCountInclUncorroborated

`func (o *InvestigatorsRow) HasNihProjectCountInclUncorroborated() bool`

HasNihProjectCountInclUncorroborated returns a boolean if a field has been set.

### SetNihProjectCountInclUncorroboratedNil

`func (o *InvestigatorsRow) SetNihProjectCountInclUncorroboratedNil(b bool)`

 SetNihProjectCountInclUncorroboratedNil sets the value for NihProjectCountInclUncorroborated to be an explicit nil

### UnsetNihProjectCountInclUncorroborated
`func (o *InvestigatorsRow) UnsetNihProjectCountInclUncorroborated()`

UnsetNihProjectCountInclUncorroborated ensures that no value is present for NihProjectCountInclUncorroborated, not even an explicit nil
### GetNihCorroborated

`func (o *InvestigatorsRow) GetNihCorroborated() bool`

GetNihCorroborated returns the NihCorroborated field if non-nil, zero value otherwise.

### GetNihCorroboratedOk

`func (o *InvestigatorsRow) GetNihCorroboratedOk() (*bool, bool)`

GetNihCorroboratedOk returns a tuple with the NihCorroborated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNihCorroborated

`func (o *InvestigatorsRow) SetNihCorroborated(v bool)`

SetNihCorroborated sets NihCorroborated field to given value.

### HasNihCorroborated

`func (o *InvestigatorsRow) HasNihCorroborated() bool`

HasNihCorroborated returns a boolean if a field has been set.

### SetNihCorroboratedNil

`func (o *InvestigatorsRow) SetNihCorroboratedNil(b bool)`

 SetNihCorroboratedNil sets the value for NihCorroborated to be an explicit nil

### UnsetNihCorroborated
`func (o *InvestigatorsRow) UnsetNihCorroborated()`

UnsetNihCorroborated ensures that no value is present for NihCorroborated, not even an explicit nil
### GetNihAwardCount

`func (o *InvestigatorsRow) GetNihAwardCount() int64`

GetNihAwardCount returns the NihAwardCount field if non-nil, zero value otherwise.

### GetNihAwardCountOk

`func (o *InvestigatorsRow) GetNihAwardCountOk() (*int64, bool)`

GetNihAwardCountOk returns a tuple with the NihAwardCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNihAwardCount

`func (o *InvestigatorsRow) SetNihAwardCount(v int64)`

SetNihAwardCount sets NihAwardCount field to given value.

### HasNihAwardCount

`func (o *InvestigatorsRow) HasNihAwardCount() bool`

HasNihAwardCount returns a boolean if a field has been set.

### SetNihAwardCountNil

`func (o *InvestigatorsRow) SetNihAwardCountNil(b bool)`

 SetNihAwardCountNil sets the value for NihAwardCount to be an explicit nil

### UnsetNihAwardCount
`func (o *InvestigatorsRow) UnsetNihAwardCount()`

UnsetNihAwardCount ensures that no value is present for NihAwardCount, not even an explicit nil
### GetNihFirstFiscalYear

`func (o *InvestigatorsRow) GetNihFirstFiscalYear() int32`

GetNihFirstFiscalYear returns the NihFirstFiscalYear field if non-nil, zero value otherwise.

### GetNihFirstFiscalYearOk

`func (o *InvestigatorsRow) GetNihFirstFiscalYearOk() (*int32, bool)`

GetNihFirstFiscalYearOk returns a tuple with the NihFirstFiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNihFirstFiscalYear

`func (o *InvestigatorsRow) SetNihFirstFiscalYear(v int32)`

SetNihFirstFiscalYear sets NihFirstFiscalYear field to given value.

### HasNihFirstFiscalYear

`func (o *InvestigatorsRow) HasNihFirstFiscalYear() bool`

HasNihFirstFiscalYear returns a boolean if a field has been set.

### SetNihFirstFiscalYearNil

`func (o *InvestigatorsRow) SetNihFirstFiscalYearNil(b bool)`

 SetNihFirstFiscalYearNil sets the value for NihFirstFiscalYear to be an explicit nil

### UnsetNihFirstFiscalYear
`func (o *InvestigatorsRow) UnsetNihFirstFiscalYear()`

UnsetNihFirstFiscalYear ensures that no value is present for NihFirstFiscalYear, not even an explicit nil
### GetNihLastFiscalYear

`func (o *InvestigatorsRow) GetNihLastFiscalYear() int32`

GetNihLastFiscalYear returns the NihLastFiscalYear field if non-nil, zero value otherwise.

### GetNihLastFiscalYearOk

`func (o *InvestigatorsRow) GetNihLastFiscalYearOk() (*int32, bool)`

GetNihLastFiscalYearOk returns a tuple with the NihLastFiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNihLastFiscalYear

`func (o *InvestigatorsRow) SetNihLastFiscalYear(v int32)`

SetNihLastFiscalYear sets NihLastFiscalYear field to given value.

### HasNihLastFiscalYear

`func (o *InvestigatorsRow) HasNihLastFiscalYear() bool`

HasNihLastFiscalYear returns a boolean if a field has been set.

### SetNihLastFiscalYearNil

`func (o *InvestigatorsRow) SetNihLastFiscalYearNil(b bool)`

 SetNihLastFiscalYearNil sets the value for NihLastFiscalYear to be an explicit nil

### UnsetNihLastFiscalYear
`func (o *InvestigatorsRow) UnsetNihLastFiscalYear()`

UnsetNihLastFiscalYear ensures that no value is present for NihLastFiscalYear, not even an explicit nil
### GetNihTopIc

`func (o *InvestigatorsRow) GetNihTopIc() string`

GetNihTopIc returns the NihTopIc field if non-nil, zero value otherwise.

### GetNihTopIcOk

`func (o *InvestigatorsRow) GetNihTopIcOk() (*string, bool)`

GetNihTopIcOk returns a tuple with the NihTopIc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNihTopIc

`func (o *InvestigatorsRow) SetNihTopIc(v string)`

SetNihTopIc sets NihTopIc field to given value.

### HasNihTopIc

`func (o *InvestigatorsRow) HasNihTopIc() bool`

HasNihTopIc returns a boolean if a field has been set.

### SetNihTopIcNil

`func (o *InvestigatorsRow) SetNihTopIcNil(b bool)`

 SetNihTopIcNil sets the value for NihTopIc to be an explicit nil

### UnsetNihTopIc
`func (o *InvestigatorsRow) UnsetNihTopIc()`

UnsetNihTopIc ensures that no value is present for NihTopIc, not even an explicit nil
### GetHasHospitalAffiliation

`func (o *InvestigatorsRow) GetHasHospitalAffiliation() bool`

GetHasHospitalAffiliation returns the HasHospitalAffiliation field if non-nil, zero value otherwise.

### GetHasHospitalAffiliationOk

`func (o *InvestigatorsRow) GetHasHospitalAffiliationOk() (*bool, bool)`

GetHasHospitalAffiliationOk returns a tuple with the HasHospitalAffiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHospitalAffiliation

`func (o *InvestigatorsRow) SetHasHospitalAffiliation(v bool)`

SetHasHospitalAffiliation sets HasHospitalAffiliation field to given value.

### HasHasHospitalAffiliation

`func (o *InvestigatorsRow) HasHasHospitalAffiliation() bool`

HasHasHospitalAffiliation returns a boolean if a field has been set.

### SetHasHospitalAffiliationNil

`func (o *InvestigatorsRow) SetHasHospitalAffiliationNil(b bool)`

 SetHasHospitalAffiliationNil sets the value for HasHospitalAffiliation to be an explicit nil

### UnsetHasHospitalAffiliation
`func (o *InvestigatorsRow) UnsetHasHospitalAffiliation()`

UnsetHasHospitalAffiliation ensures that no value is present for HasHospitalAffiliation, not even an explicit nil
### GetNAffiliatedHospitals

`func (o *InvestigatorsRow) GetNAffiliatedHospitals() int64`

GetNAffiliatedHospitals returns the NAffiliatedHospitals field if non-nil, zero value otherwise.

### GetNAffiliatedHospitalsOk

`func (o *InvestigatorsRow) GetNAffiliatedHospitalsOk() (*int64, bool)`

GetNAffiliatedHospitalsOk returns a tuple with the NAffiliatedHospitals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAffiliatedHospitals

`func (o *InvestigatorsRow) SetNAffiliatedHospitals(v int64)`

SetNAffiliatedHospitals sets NAffiliatedHospitals field to given value.

### HasNAffiliatedHospitals

`func (o *InvestigatorsRow) HasNAffiliatedHospitals() bool`

HasNAffiliatedHospitals returns a boolean if a field has been set.

### SetNAffiliatedHospitalsNil

`func (o *InvestigatorsRow) SetNAffiliatedHospitalsNil(b bool)`

 SetNAffiliatedHospitalsNil sets the value for NAffiliatedHospitals to be an explicit nil

### UnsetNAffiliatedHospitals
`func (o *InvestigatorsRow) UnsetNAffiliatedHospitals()`

UnsetNAffiliatedHospitals ensures that no value is present for NAffiliatedHospitals, not even an explicit nil
### GetPrimaryCcn

`func (o *InvestigatorsRow) GetPrimaryCcn() string`

GetPrimaryCcn returns the PrimaryCcn field if non-nil, zero value otherwise.

### GetPrimaryCcnOk

`func (o *InvestigatorsRow) GetPrimaryCcnOk() (*string, bool)`

GetPrimaryCcnOk returns a tuple with the PrimaryCcn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCcn

`func (o *InvestigatorsRow) SetPrimaryCcn(v string)`

SetPrimaryCcn sets PrimaryCcn field to given value.

### HasPrimaryCcn

`func (o *InvestigatorsRow) HasPrimaryCcn() bool`

HasPrimaryCcn returns a boolean if a field has been set.

### SetPrimaryCcnNil

`func (o *InvestigatorsRow) SetPrimaryCcnNil(b bool)`

 SetPrimaryCcnNil sets the value for PrimaryCcn to be an explicit nil

### UnsetPrimaryCcn
`func (o *InvestigatorsRow) UnsetPrimaryCcn()`

UnsetPrimaryCcn ensures that no value is present for PrimaryCcn, not even an explicit nil
### GetHasExpertiseTopics

`func (o *InvestigatorsRow) GetHasExpertiseTopics() bool`

GetHasExpertiseTopics returns the HasExpertiseTopics field if non-nil, zero value otherwise.

### GetHasExpertiseTopicsOk

`func (o *InvestigatorsRow) GetHasExpertiseTopicsOk() (*bool, bool)`

GetHasExpertiseTopicsOk returns a tuple with the HasExpertiseTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExpertiseTopics

`func (o *InvestigatorsRow) SetHasExpertiseTopics(v bool)`

SetHasExpertiseTopics sets HasExpertiseTopics field to given value.

### HasHasExpertiseTopics

`func (o *InvestigatorsRow) HasHasExpertiseTopics() bool`

HasHasExpertiseTopics returns a boolean if a field has been set.

### SetHasExpertiseTopicsNil

`func (o *InvestigatorsRow) SetHasExpertiseTopicsNil(b bool)`

 SetHasExpertiseTopicsNil sets the value for HasExpertiseTopics to be an explicit nil

### UnsetHasExpertiseTopics
`func (o *InvestigatorsRow) UnsetHasExpertiseTopics()`

UnsetHasExpertiseTopics ensures that no value is present for HasExpertiseTopics, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


