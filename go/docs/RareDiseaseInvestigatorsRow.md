# RareDiseaseInvestigatorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**DiseaseName** | Pointer to **NullableString** | The disease&#39;s preferred English name. | [optional] 
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key. | [optional] 
**InvestigatorPersonName** | Pointer to **NullableString** | The investigator&#39;s name as the registry filed it. | [optional] 
**IdentityBasis** | Pointer to **NullableString** | What the person&#39;s identity rests on: curated, npi, orcid, reporter, affinity or self. | [optional] 
**MergedFragmentCount** | Pointer to **NullableInt64** | Name fragments merged into this person. | [optional] 
**NameIsAmbiguousInDisease** | Pointer to **NullableBool** | True when another investigator of the same disease shares the name. | [optional] 
**PiTrialCount** | Pointer to **NullableInt64** | Exact direct trials of the disease that name the person as principal investigator. | [optional] 
**FirstTrialStartDate** | Pointer to **NullableString** | The earliest start date among those trials. | [optional] 
**LastTrialStartDate** | Pointer to **NullableString** | The latest start date among those trials. | [optional] 
**LastTrialStartYear** | Pointer to **NullableInt32** | The year of the latest start. | [optional] 
**EvidenceTrialId** | Pointer to **NullableString** | One of the trials that names the person (the latest started). | [optional] 
**LastRecordUpdateDate** | Pointer to **NullableString** | The latest registry update among those trials. | [optional] 
**RegistryReferenceDate** | Pointer to **NullableString** | The registry&#39;s latest update date the snapshot reflects. | [optional] 
**RecencyState** | Pointer to **NullableString** | How recent the latest trial start is against the registry snapshot: within_5_years, 5_to_10_years, over_10_years or no_start_date. | [optional] 
**TrialSiteNames** | Pointer to **[]string** | The sites of the person&#39;s trials as the registry lists them. | [optional] 
**TrialSiteCount** | Pointer to **NullableInt64** | How many such sites. | [optional] 
**StatedAffiliations** | Pointer to **[]string** | The affiliations the registry states for the person. | [optional] 
**StatedAffiliationCount** | Pointer to **NullableInt64** | How many stated affiliations. | [optional] 
**AffiliationEvidence** | Pointer to **NullableString** | trial_site when a trial site is listed, stated_affiliation_only when only an affiliation is stated, none otherwise. | [optional] 

## Methods

### NewRareDiseaseInvestigatorsRow

`func NewRareDiseaseInvestigatorsRow() *RareDiseaseInvestigatorsRow`

NewRareDiseaseInvestigatorsRow instantiates a new RareDiseaseInvestigatorsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseInvestigatorsRowWithDefaults

`func NewRareDiseaseInvestigatorsRowWithDefaults() *RareDiseaseInvestigatorsRow`

NewRareDiseaseInvestigatorsRowWithDefaults instantiates a new RareDiseaseInvestigatorsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrphaCode

`func (o *RareDiseaseInvestigatorsRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseInvestigatorsRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseInvestigatorsRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseInvestigatorsRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseInvestigatorsRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseInvestigatorsRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetDiseaseName

`func (o *RareDiseaseInvestigatorsRow) GetDiseaseName() string`

GetDiseaseName returns the DiseaseName field if non-nil, zero value otherwise.

### GetDiseaseNameOk

`func (o *RareDiseaseInvestigatorsRow) GetDiseaseNameOk() (*string, bool)`

GetDiseaseNameOk returns a tuple with the DiseaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiseaseName

`func (o *RareDiseaseInvestigatorsRow) SetDiseaseName(v string)`

SetDiseaseName sets DiseaseName field to given value.

### HasDiseaseName

`func (o *RareDiseaseInvestigatorsRow) HasDiseaseName() bool`

HasDiseaseName returns a boolean if a field has been set.

### SetDiseaseNameNil

`func (o *RareDiseaseInvestigatorsRow) SetDiseaseNameNil(b bool)`

 SetDiseaseNameNil sets the value for DiseaseName to be an explicit nil

### UnsetDiseaseName
`func (o *RareDiseaseInvestigatorsRow) UnsetDiseaseName()`

UnsetDiseaseName ensures that no value is present for DiseaseName, not even an explicit nil
### GetInvestigatorPersonId

`func (o *RareDiseaseInvestigatorsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *RareDiseaseInvestigatorsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *RareDiseaseInvestigatorsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *RareDiseaseInvestigatorsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *RareDiseaseInvestigatorsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *RareDiseaseInvestigatorsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetInvestigatorPersonName

`func (o *RareDiseaseInvestigatorsRow) GetInvestigatorPersonName() string`

GetInvestigatorPersonName returns the InvestigatorPersonName field if non-nil, zero value otherwise.

### GetInvestigatorPersonNameOk

`func (o *RareDiseaseInvestigatorsRow) GetInvestigatorPersonNameOk() (*string, bool)`

GetInvestigatorPersonNameOk returns a tuple with the InvestigatorPersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonName

`func (o *RareDiseaseInvestigatorsRow) SetInvestigatorPersonName(v string)`

SetInvestigatorPersonName sets InvestigatorPersonName field to given value.

### HasInvestigatorPersonName

`func (o *RareDiseaseInvestigatorsRow) HasInvestigatorPersonName() bool`

HasInvestigatorPersonName returns a boolean if a field has been set.

### SetInvestigatorPersonNameNil

`func (o *RareDiseaseInvestigatorsRow) SetInvestigatorPersonNameNil(b bool)`

 SetInvestigatorPersonNameNil sets the value for InvestigatorPersonName to be an explicit nil

### UnsetInvestigatorPersonName
`func (o *RareDiseaseInvestigatorsRow) UnsetInvestigatorPersonName()`

UnsetInvestigatorPersonName ensures that no value is present for InvestigatorPersonName, not even an explicit nil
### GetIdentityBasis

`func (o *RareDiseaseInvestigatorsRow) GetIdentityBasis() string`

GetIdentityBasis returns the IdentityBasis field if non-nil, zero value otherwise.

### GetIdentityBasisOk

`func (o *RareDiseaseInvestigatorsRow) GetIdentityBasisOk() (*string, bool)`

GetIdentityBasisOk returns a tuple with the IdentityBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityBasis

`func (o *RareDiseaseInvestigatorsRow) SetIdentityBasis(v string)`

SetIdentityBasis sets IdentityBasis field to given value.

### HasIdentityBasis

`func (o *RareDiseaseInvestigatorsRow) HasIdentityBasis() bool`

HasIdentityBasis returns a boolean if a field has been set.

### SetIdentityBasisNil

`func (o *RareDiseaseInvestigatorsRow) SetIdentityBasisNil(b bool)`

 SetIdentityBasisNil sets the value for IdentityBasis to be an explicit nil

### UnsetIdentityBasis
`func (o *RareDiseaseInvestigatorsRow) UnsetIdentityBasis()`

UnsetIdentityBasis ensures that no value is present for IdentityBasis, not even an explicit nil
### GetMergedFragmentCount

`func (o *RareDiseaseInvestigatorsRow) GetMergedFragmentCount() int64`

GetMergedFragmentCount returns the MergedFragmentCount field if non-nil, zero value otherwise.

### GetMergedFragmentCountOk

`func (o *RareDiseaseInvestigatorsRow) GetMergedFragmentCountOk() (*int64, bool)`

GetMergedFragmentCountOk returns a tuple with the MergedFragmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedFragmentCount

`func (o *RareDiseaseInvestigatorsRow) SetMergedFragmentCount(v int64)`

SetMergedFragmentCount sets MergedFragmentCount field to given value.

### HasMergedFragmentCount

`func (o *RareDiseaseInvestigatorsRow) HasMergedFragmentCount() bool`

HasMergedFragmentCount returns a boolean if a field has been set.

### SetMergedFragmentCountNil

`func (o *RareDiseaseInvestigatorsRow) SetMergedFragmentCountNil(b bool)`

 SetMergedFragmentCountNil sets the value for MergedFragmentCount to be an explicit nil

### UnsetMergedFragmentCount
`func (o *RareDiseaseInvestigatorsRow) UnsetMergedFragmentCount()`

UnsetMergedFragmentCount ensures that no value is present for MergedFragmentCount, not even an explicit nil
### GetNameIsAmbiguousInDisease

`func (o *RareDiseaseInvestigatorsRow) GetNameIsAmbiguousInDisease() bool`

GetNameIsAmbiguousInDisease returns the NameIsAmbiguousInDisease field if non-nil, zero value otherwise.

### GetNameIsAmbiguousInDiseaseOk

`func (o *RareDiseaseInvestigatorsRow) GetNameIsAmbiguousInDiseaseOk() (*bool, bool)`

GetNameIsAmbiguousInDiseaseOk returns a tuple with the NameIsAmbiguousInDisease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIsAmbiguousInDisease

`func (o *RareDiseaseInvestigatorsRow) SetNameIsAmbiguousInDisease(v bool)`

SetNameIsAmbiguousInDisease sets NameIsAmbiguousInDisease field to given value.

### HasNameIsAmbiguousInDisease

`func (o *RareDiseaseInvestigatorsRow) HasNameIsAmbiguousInDisease() bool`

HasNameIsAmbiguousInDisease returns a boolean if a field has been set.

### SetNameIsAmbiguousInDiseaseNil

`func (o *RareDiseaseInvestigatorsRow) SetNameIsAmbiguousInDiseaseNil(b bool)`

 SetNameIsAmbiguousInDiseaseNil sets the value for NameIsAmbiguousInDisease to be an explicit nil

### UnsetNameIsAmbiguousInDisease
`func (o *RareDiseaseInvestigatorsRow) UnsetNameIsAmbiguousInDisease()`

UnsetNameIsAmbiguousInDisease ensures that no value is present for NameIsAmbiguousInDisease, not even an explicit nil
### GetPiTrialCount

`func (o *RareDiseaseInvestigatorsRow) GetPiTrialCount() int64`

GetPiTrialCount returns the PiTrialCount field if non-nil, zero value otherwise.

### GetPiTrialCountOk

`func (o *RareDiseaseInvestigatorsRow) GetPiTrialCountOk() (*int64, bool)`

GetPiTrialCountOk returns a tuple with the PiTrialCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiTrialCount

`func (o *RareDiseaseInvestigatorsRow) SetPiTrialCount(v int64)`

SetPiTrialCount sets PiTrialCount field to given value.

### HasPiTrialCount

`func (o *RareDiseaseInvestigatorsRow) HasPiTrialCount() bool`

HasPiTrialCount returns a boolean if a field has been set.

### SetPiTrialCountNil

`func (o *RareDiseaseInvestigatorsRow) SetPiTrialCountNil(b bool)`

 SetPiTrialCountNil sets the value for PiTrialCount to be an explicit nil

### UnsetPiTrialCount
`func (o *RareDiseaseInvestigatorsRow) UnsetPiTrialCount()`

UnsetPiTrialCount ensures that no value is present for PiTrialCount, not even an explicit nil
### GetFirstTrialStartDate

`func (o *RareDiseaseInvestigatorsRow) GetFirstTrialStartDate() string`

GetFirstTrialStartDate returns the FirstTrialStartDate field if non-nil, zero value otherwise.

### GetFirstTrialStartDateOk

`func (o *RareDiseaseInvestigatorsRow) GetFirstTrialStartDateOk() (*string, bool)`

GetFirstTrialStartDateOk returns a tuple with the FirstTrialStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTrialStartDate

`func (o *RareDiseaseInvestigatorsRow) SetFirstTrialStartDate(v string)`

SetFirstTrialStartDate sets FirstTrialStartDate field to given value.

### HasFirstTrialStartDate

`func (o *RareDiseaseInvestigatorsRow) HasFirstTrialStartDate() bool`

HasFirstTrialStartDate returns a boolean if a field has been set.

### SetFirstTrialStartDateNil

`func (o *RareDiseaseInvestigatorsRow) SetFirstTrialStartDateNil(b bool)`

 SetFirstTrialStartDateNil sets the value for FirstTrialStartDate to be an explicit nil

### UnsetFirstTrialStartDate
`func (o *RareDiseaseInvestigatorsRow) UnsetFirstTrialStartDate()`

UnsetFirstTrialStartDate ensures that no value is present for FirstTrialStartDate, not even an explicit nil
### GetLastTrialStartDate

`func (o *RareDiseaseInvestigatorsRow) GetLastTrialStartDate() string`

GetLastTrialStartDate returns the LastTrialStartDate field if non-nil, zero value otherwise.

### GetLastTrialStartDateOk

`func (o *RareDiseaseInvestigatorsRow) GetLastTrialStartDateOk() (*string, bool)`

GetLastTrialStartDateOk returns a tuple with the LastTrialStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrialStartDate

`func (o *RareDiseaseInvestigatorsRow) SetLastTrialStartDate(v string)`

SetLastTrialStartDate sets LastTrialStartDate field to given value.

### HasLastTrialStartDate

`func (o *RareDiseaseInvestigatorsRow) HasLastTrialStartDate() bool`

HasLastTrialStartDate returns a boolean if a field has been set.

### SetLastTrialStartDateNil

`func (o *RareDiseaseInvestigatorsRow) SetLastTrialStartDateNil(b bool)`

 SetLastTrialStartDateNil sets the value for LastTrialStartDate to be an explicit nil

### UnsetLastTrialStartDate
`func (o *RareDiseaseInvestigatorsRow) UnsetLastTrialStartDate()`

UnsetLastTrialStartDate ensures that no value is present for LastTrialStartDate, not even an explicit nil
### GetLastTrialStartYear

`func (o *RareDiseaseInvestigatorsRow) GetLastTrialStartYear() int32`

GetLastTrialStartYear returns the LastTrialStartYear field if non-nil, zero value otherwise.

### GetLastTrialStartYearOk

`func (o *RareDiseaseInvestigatorsRow) GetLastTrialStartYearOk() (*int32, bool)`

GetLastTrialStartYearOk returns a tuple with the LastTrialStartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrialStartYear

`func (o *RareDiseaseInvestigatorsRow) SetLastTrialStartYear(v int32)`

SetLastTrialStartYear sets LastTrialStartYear field to given value.

### HasLastTrialStartYear

`func (o *RareDiseaseInvestigatorsRow) HasLastTrialStartYear() bool`

HasLastTrialStartYear returns a boolean if a field has been set.

### SetLastTrialStartYearNil

`func (o *RareDiseaseInvestigatorsRow) SetLastTrialStartYearNil(b bool)`

 SetLastTrialStartYearNil sets the value for LastTrialStartYear to be an explicit nil

### UnsetLastTrialStartYear
`func (o *RareDiseaseInvestigatorsRow) UnsetLastTrialStartYear()`

UnsetLastTrialStartYear ensures that no value is present for LastTrialStartYear, not even an explicit nil
### GetEvidenceTrialId

`func (o *RareDiseaseInvestigatorsRow) GetEvidenceTrialId() string`

GetEvidenceTrialId returns the EvidenceTrialId field if non-nil, zero value otherwise.

### GetEvidenceTrialIdOk

`func (o *RareDiseaseInvestigatorsRow) GetEvidenceTrialIdOk() (*string, bool)`

GetEvidenceTrialIdOk returns a tuple with the EvidenceTrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceTrialId

`func (o *RareDiseaseInvestigatorsRow) SetEvidenceTrialId(v string)`

SetEvidenceTrialId sets EvidenceTrialId field to given value.

### HasEvidenceTrialId

`func (o *RareDiseaseInvestigatorsRow) HasEvidenceTrialId() bool`

HasEvidenceTrialId returns a boolean if a field has been set.

### SetEvidenceTrialIdNil

`func (o *RareDiseaseInvestigatorsRow) SetEvidenceTrialIdNil(b bool)`

 SetEvidenceTrialIdNil sets the value for EvidenceTrialId to be an explicit nil

### UnsetEvidenceTrialId
`func (o *RareDiseaseInvestigatorsRow) UnsetEvidenceTrialId()`

UnsetEvidenceTrialId ensures that no value is present for EvidenceTrialId, not even an explicit nil
### GetLastRecordUpdateDate

`func (o *RareDiseaseInvestigatorsRow) GetLastRecordUpdateDate() string`

GetLastRecordUpdateDate returns the LastRecordUpdateDate field if non-nil, zero value otherwise.

### GetLastRecordUpdateDateOk

`func (o *RareDiseaseInvestigatorsRow) GetLastRecordUpdateDateOk() (*string, bool)`

GetLastRecordUpdateDateOk returns a tuple with the LastRecordUpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecordUpdateDate

`func (o *RareDiseaseInvestigatorsRow) SetLastRecordUpdateDate(v string)`

SetLastRecordUpdateDate sets LastRecordUpdateDate field to given value.

### HasLastRecordUpdateDate

`func (o *RareDiseaseInvestigatorsRow) HasLastRecordUpdateDate() bool`

HasLastRecordUpdateDate returns a boolean if a field has been set.

### SetLastRecordUpdateDateNil

`func (o *RareDiseaseInvestigatorsRow) SetLastRecordUpdateDateNil(b bool)`

 SetLastRecordUpdateDateNil sets the value for LastRecordUpdateDate to be an explicit nil

### UnsetLastRecordUpdateDate
`func (o *RareDiseaseInvestigatorsRow) UnsetLastRecordUpdateDate()`

UnsetLastRecordUpdateDate ensures that no value is present for LastRecordUpdateDate, not even an explicit nil
### GetRegistryReferenceDate

`func (o *RareDiseaseInvestigatorsRow) GetRegistryReferenceDate() string`

GetRegistryReferenceDate returns the RegistryReferenceDate field if non-nil, zero value otherwise.

### GetRegistryReferenceDateOk

`func (o *RareDiseaseInvestigatorsRow) GetRegistryReferenceDateOk() (*string, bool)`

GetRegistryReferenceDateOk returns a tuple with the RegistryReferenceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryReferenceDate

`func (o *RareDiseaseInvestigatorsRow) SetRegistryReferenceDate(v string)`

SetRegistryReferenceDate sets RegistryReferenceDate field to given value.

### HasRegistryReferenceDate

`func (o *RareDiseaseInvestigatorsRow) HasRegistryReferenceDate() bool`

HasRegistryReferenceDate returns a boolean if a field has been set.

### SetRegistryReferenceDateNil

`func (o *RareDiseaseInvestigatorsRow) SetRegistryReferenceDateNil(b bool)`

 SetRegistryReferenceDateNil sets the value for RegistryReferenceDate to be an explicit nil

### UnsetRegistryReferenceDate
`func (o *RareDiseaseInvestigatorsRow) UnsetRegistryReferenceDate()`

UnsetRegistryReferenceDate ensures that no value is present for RegistryReferenceDate, not even an explicit nil
### GetRecencyState

`func (o *RareDiseaseInvestigatorsRow) GetRecencyState() string`

GetRecencyState returns the RecencyState field if non-nil, zero value otherwise.

### GetRecencyStateOk

`func (o *RareDiseaseInvestigatorsRow) GetRecencyStateOk() (*string, bool)`

GetRecencyStateOk returns a tuple with the RecencyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecencyState

`func (o *RareDiseaseInvestigatorsRow) SetRecencyState(v string)`

SetRecencyState sets RecencyState field to given value.

### HasRecencyState

`func (o *RareDiseaseInvestigatorsRow) HasRecencyState() bool`

HasRecencyState returns a boolean if a field has been set.

### SetRecencyStateNil

`func (o *RareDiseaseInvestigatorsRow) SetRecencyStateNil(b bool)`

 SetRecencyStateNil sets the value for RecencyState to be an explicit nil

### UnsetRecencyState
`func (o *RareDiseaseInvestigatorsRow) UnsetRecencyState()`

UnsetRecencyState ensures that no value is present for RecencyState, not even an explicit nil
### GetTrialSiteNames

`func (o *RareDiseaseInvestigatorsRow) GetTrialSiteNames() []string`

GetTrialSiteNames returns the TrialSiteNames field if non-nil, zero value otherwise.

### GetTrialSiteNamesOk

`func (o *RareDiseaseInvestigatorsRow) GetTrialSiteNamesOk() (*[]string, bool)`

GetTrialSiteNamesOk returns a tuple with the TrialSiteNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialSiteNames

`func (o *RareDiseaseInvestigatorsRow) SetTrialSiteNames(v []string)`

SetTrialSiteNames sets TrialSiteNames field to given value.

### HasTrialSiteNames

`func (o *RareDiseaseInvestigatorsRow) HasTrialSiteNames() bool`

HasTrialSiteNames returns a boolean if a field has been set.

### SetTrialSiteNamesNil

`func (o *RareDiseaseInvestigatorsRow) SetTrialSiteNamesNil(b bool)`

 SetTrialSiteNamesNil sets the value for TrialSiteNames to be an explicit nil

### UnsetTrialSiteNames
`func (o *RareDiseaseInvestigatorsRow) UnsetTrialSiteNames()`

UnsetTrialSiteNames ensures that no value is present for TrialSiteNames, not even an explicit nil
### GetTrialSiteCount

`func (o *RareDiseaseInvestigatorsRow) GetTrialSiteCount() int64`

GetTrialSiteCount returns the TrialSiteCount field if non-nil, zero value otherwise.

### GetTrialSiteCountOk

`func (o *RareDiseaseInvestigatorsRow) GetTrialSiteCountOk() (*int64, bool)`

GetTrialSiteCountOk returns a tuple with the TrialSiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialSiteCount

`func (o *RareDiseaseInvestigatorsRow) SetTrialSiteCount(v int64)`

SetTrialSiteCount sets TrialSiteCount field to given value.

### HasTrialSiteCount

`func (o *RareDiseaseInvestigatorsRow) HasTrialSiteCount() bool`

HasTrialSiteCount returns a boolean if a field has been set.

### SetTrialSiteCountNil

`func (o *RareDiseaseInvestigatorsRow) SetTrialSiteCountNil(b bool)`

 SetTrialSiteCountNil sets the value for TrialSiteCount to be an explicit nil

### UnsetTrialSiteCount
`func (o *RareDiseaseInvestigatorsRow) UnsetTrialSiteCount()`

UnsetTrialSiteCount ensures that no value is present for TrialSiteCount, not even an explicit nil
### GetStatedAffiliations

`func (o *RareDiseaseInvestigatorsRow) GetStatedAffiliations() []string`

GetStatedAffiliations returns the StatedAffiliations field if non-nil, zero value otherwise.

### GetStatedAffiliationsOk

`func (o *RareDiseaseInvestigatorsRow) GetStatedAffiliationsOk() (*[]string, bool)`

GetStatedAffiliationsOk returns a tuple with the StatedAffiliations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatedAffiliations

`func (o *RareDiseaseInvestigatorsRow) SetStatedAffiliations(v []string)`

SetStatedAffiliations sets StatedAffiliations field to given value.

### HasStatedAffiliations

`func (o *RareDiseaseInvestigatorsRow) HasStatedAffiliations() bool`

HasStatedAffiliations returns a boolean if a field has been set.

### SetStatedAffiliationsNil

`func (o *RareDiseaseInvestigatorsRow) SetStatedAffiliationsNil(b bool)`

 SetStatedAffiliationsNil sets the value for StatedAffiliations to be an explicit nil

### UnsetStatedAffiliations
`func (o *RareDiseaseInvestigatorsRow) UnsetStatedAffiliations()`

UnsetStatedAffiliations ensures that no value is present for StatedAffiliations, not even an explicit nil
### GetStatedAffiliationCount

`func (o *RareDiseaseInvestigatorsRow) GetStatedAffiliationCount() int64`

GetStatedAffiliationCount returns the StatedAffiliationCount field if non-nil, zero value otherwise.

### GetStatedAffiliationCountOk

`func (o *RareDiseaseInvestigatorsRow) GetStatedAffiliationCountOk() (*int64, bool)`

GetStatedAffiliationCountOk returns a tuple with the StatedAffiliationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatedAffiliationCount

`func (o *RareDiseaseInvestigatorsRow) SetStatedAffiliationCount(v int64)`

SetStatedAffiliationCount sets StatedAffiliationCount field to given value.

### HasStatedAffiliationCount

`func (o *RareDiseaseInvestigatorsRow) HasStatedAffiliationCount() bool`

HasStatedAffiliationCount returns a boolean if a field has been set.

### SetStatedAffiliationCountNil

`func (o *RareDiseaseInvestigatorsRow) SetStatedAffiliationCountNil(b bool)`

 SetStatedAffiliationCountNil sets the value for StatedAffiliationCount to be an explicit nil

### UnsetStatedAffiliationCount
`func (o *RareDiseaseInvestigatorsRow) UnsetStatedAffiliationCount()`

UnsetStatedAffiliationCount ensures that no value is present for StatedAffiliationCount, not even an explicit nil
### GetAffiliationEvidence

`func (o *RareDiseaseInvestigatorsRow) GetAffiliationEvidence() string`

GetAffiliationEvidence returns the AffiliationEvidence field if non-nil, zero value otherwise.

### GetAffiliationEvidenceOk

`func (o *RareDiseaseInvestigatorsRow) GetAffiliationEvidenceOk() (*string, bool)`

GetAffiliationEvidenceOk returns a tuple with the AffiliationEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliationEvidence

`func (o *RareDiseaseInvestigatorsRow) SetAffiliationEvidence(v string)`

SetAffiliationEvidence sets AffiliationEvidence field to given value.

### HasAffiliationEvidence

`func (o *RareDiseaseInvestigatorsRow) HasAffiliationEvidence() bool`

HasAffiliationEvidence returns a boolean if a field has been set.

### SetAffiliationEvidenceNil

`func (o *RareDiseaseInvestigatorsRow) SetAffiliationEvidenceNil(b bool)`

 SetAffiliationEvidenceNil sets the value for AffiliationEvidence to be an explicit nil

### UnsetAffiliationEvidence
`func (o *RareDiseaseInvestigatorsRow) UnsetAffiliationEvidence()`

UnsetAffiliationEvidence ensures that no value is present for AffiliationEvidence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


