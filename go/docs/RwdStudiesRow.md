# RwdStudiesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StudyId** | Pointer to **NullableString** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] 
**EuPasNumber** | Pointer to **NullableString** | The EU PAS (post-authorisation study) register number of the study. | [optional] 
**StudyTitle** | Pointer to **NullableString** | The study&#39;s title as registered. | [optional] 
**CatalogueUrl** | Pointer to **NullableString** | The address of this entry on the HMA-EMA catalogue site (the register&#39;s URL LINK). | [optional] 
**StudyStatus** | Pointer to **NullableString** | The study&#39;s status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional] 
**IsOngoing** | Pointer to **NullableBool** | True when the study status is Ongoing. | [optional] 
**IsFinalised** | Pointer to **NullableBool** | True when the study status is Finalised. | [optional] 
**IsDarwinEuStudy** | Pointer to **NullableBool** | Whether the study is flagged a DARWIN EU study. | [optional] 
**HasEnceppSeal** | Pointer to **NullableBool** | Whether the study has been awarded the ENCePP (European Network of Centres for Pharmacoepidemiology and Pharmacovigilance) seal. | [optional] 
**RequiredByRegulator** | Pointer to **NullableString** | The registrant&#39;s answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional] 
**MandateStratum** | Pointer to **NullableString** | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional] 
**RmpRequirement** | Pointer to **NullableString** | The risk management plan (RMP) category the registrant says required the study, as typed. | [optional] 
**HasRmpCategory** | Pointer to **NullableBool** | Whether the registrant gave an RMP requirement category. | [optional] 
**RegulatoryProcedureNumber** | Pointer to **NullableString** | The regulatory procedure number the registrant gave. | [optional] 
**FundingSourceText** | Pointer to **NullableString** | The source of funding as the registrant typed it. | [optional] 
**FundingSources** | Pointer to **[]string** | The funding sources split into an array. | [optional] 
**IsPharmaFunded** | Pointer to **NullableBool** | Whether a funding source is a pharmaceutical company. | [optional] 
**RegisterStudyType** | Pointer to **NullableString** | The register&#39;s own Study type value, published as register_study_type because study_type is a shared enumeration of another source&#39;s vocabulary. | [optional] 
**ScopeOfTheStudy** | Pointer to **NullableString** | The scope of the study as registered (for example Safety study, Drug utilisation study). | [optional] 
**StudyTopic** | Pointer to **NullableString** | The study topic as registered. | [optional] 
**StudyDesign** | Pointer to **NullableString** | The study design as registered. | [optional] 
**NonInterventionalStudyDesign** | Pointer to **NullableString** | The non-interventional study design as registered. | [optional] 
**MainStudyObjective** | Pointer to **NullableString** | The main study objective as the registrant typed it (free text). | [optional] 
**DataCollectionMethods** | Pointer to **NullableString** | The data collection methods as registered. | [optional] 
**MedicinalProductName** | Pointer to **NullableString** | The medicinal product name(s) as registered. | [optional] 
**StudyDrugInn** | Pointer to **NullableString** | The study drug international non-proprietary name(s) (INN) or common name as typed. | [optional] 
**AtcCodeText** | Pointer to **NullableString** | The Anatomical Therapeutic Chemical (ATC) code entry as typed, often with the substance name beside the code. | [optional] 
**DrugMentionCount** | Pointer to **NullableInt32** | How many drug mentions the study carries (rwd_study_drugs has one row per mention). | [optional] 
**AtcLevel5MentionCount** | Pointer to **NullableInt32** | How many of the study&#39;s drug mentions carry an ATC level 5 code. | [optional] 
**SpineMatchedMentionCount** | Pointer to **NullableInt32** | How many of those mentions resolved to a canonical drug. | [optional] 
**AtcLevel5Codes** | Pointer to **[]string** | The ATC level 5 (seven-character) codes the study names. | [optional] 
**AtcLevel1Codes** | Pointer to **[]string** | The level-1 ATC codes the study names. | [optional] 
**MedicalCondition** | Pointer to **NullableString** | The medicinal condition to be studied, as typed. | [optional] 
**AdditionalMedicalConditions** | Pointer to **NullableString** | Additional medical conditions, as typed. | [optional] 
**ConditionCount** | Pointer to **NullableInt32** | How many distinct condition terms the study names. | [optional] 
**PrimaryConditionCount** | Pointer to **NullableInt32** | How many of them are primary. | [optional] 
**MeshConditionCount** | Pointer to **NullableInt32** | How many of them resolved to a MeSH descriptor. | [optional] 
**ConditionTerms** | Pointer to **[]string** | The study&#39;s condition terms, primary and additional, as an array. | [optional] 
**PopulationAgeGroups** | Pointer to **NullableString** | The population age groups as registered. | [optional] 
**SpecialPopulationOfInterest** | Pointer to **NullableString** | The special population of interest as registered. | [optional] 
**EstimatedNumberOfSubjects** | Pointer to **NullableInt64** | The estimated number of subjects as registered; null when blank and also when the register wrote 0 (a zero is an unfilled box: see subjects_reported_zero). | [optional] 
**SubjectsReportedZero** | Pointer to **NullableBool** | True when the register held a 0 for the estimated number of subjects (published as null in estimated_number_of_subjects). | [optional] 
**SubjectsBand** | Pointer to **NullableString** | A size band of estimated_number_of_subjects: under_30, under_100, under_1k, under_100k or over_100k; null when the figure is null. | [optional] 
**ConductingInstitutionNames** | Pointer to **[]string** | The conducting institutions that match the institution register, as the register spells them; items that did not match are in unmatched_institution_fragments. | [optional] 
**ConductingInstitutionCount** | Pointer to **NullableInt32** | How many conducting institutions the study names. | [optional] 
**AdditionalInstitutionNames** | Pointer to **[]string** | The additional institutions that match the institution register; unmatched items are in unmatched_institution_fragments. | [optional] 
**AdditionalInstitutionCount** | Pointer to **NullableInt32** | How many additional institutions the study names. | [optional] 
**UnmatchedInstitutionFragments** | Pointer to **[]string** | Institution text items that did not match the institution register. | [optional] 
**NetworkNames** | Pointer to **[]string** | The networks conducting the study, as registered. | [optional] 
**NetworkCount** | Pointer to **NullableInt32** | How many networks the study names. | [optional] 
**DataSourceNames** | Pointer to **[]string** | The data sources the study names, split into an array. | [optional] 
**DataSourceCount** | Pointer to **NullableInt32** | How many data source items the study names. | [optional] 
**UnmatchedDataSourceFragments** | Pointer to **[]string** | Data source text items that did not match the data source register (the free-text Other linked data sources box is left whole). | [optional] 
**ConductorAttribution** | Pointer to **NullableString** | registered when the study names at least one conducting institution that is in the institution register, none otherwise. A study with none is not evidence of no experience: the register gives no further split. | [optional] 
**StudyCountryNames** | Pointer to **[]string** | The countries of the study, one array element per register country. | [optional] 
**StudyCountryCount** | Pointer to **NullableInt32** | How many countries the study names. | [optional] 
**UnmatchedCountryFragments** | Pointer to **[]string** | Country text items the country vocabulary of this build does not hold. | [optional] 
**FundingContractSignedPlannedOn** | Pointer to **NullableString** | Planned date the funding contract was signed. | [optional] 
**FundingContractSignedActualOn** | Pointer to **NullableString** | Actual date the funding contract was signed. | [optional] 
**StudyStartPlannedOn** | Pointer to **NullableString** | Planned study start date. | [optional] 
**StudyStartActualOn** | Pointer to **NullableString** | Actual study start date. | [optional] 
**DataAnalysisStartPlannedOn** | Pointer to **NullableString** | Planned data analysis start date. | [optional] 
**DataAnalysisStartActualOn** | Pointer to **NullableString** | Actual data analysis start date. | [optional] 
**InterimReportPlannedOn** | Pointer to **NullableString** | Planned date of the interim report, if expected. | [optional] 
**InterimReportActualOn** | Pointer to **NullableString** | Actual date of the interim report. | [optional] 
**FinalReportPlannedOn** | Pointer to **NullableString** | Planned date of the final study report. | [optional] 
**FinalReportActualOn** | Pointer to **NullableString** | Actual date of the final study report. | [optional] 
**HasFinalReportPair** | Pointer to **NullableBool** | True when both the planned and the actual final report dates are present. | [optional] 
**FinalReportIsBackfilled** | Pointer to **NullableBool** | True when the planned and actual final report dates are both present and equal to the day, which reads as a plan entered after the fact rather than perfect delivery. | [optional] 
**FinalReportSlipDays** | Pointer to **NullableInt32** | Actual minus planned final report date in days (positive is late); null when either date is missing or the two dates are equal (a backfill). | [optional] 
**StudyStartYear** | Pointer to **NullableInt32** | The calendar year of the actual study start. | [optional] 
**StartedInLast5Years** | Pointer to **NullableBool** | Whether the study started within five years of the snapshot. | [optional] 
**RecordUpdatedAt** | Pointer to **NullableTime** | When the registrant last edited this catalogue entry. A different clock from last_refreshed_on: it dates the web form, not the database. | [optional] 
**RecordFirstPublishedAt** | Pointer to **NullableTime** | When the catalogue entry was first published. | [optional] 
**RegisterSourceCount** | Pointer to **NullableInt32** | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional] 
**StudyRegisterTotal** | Pointer to **NullableInt32** | The number of studies in the whole study register at the snapshot. | [optional] 
**StudiesNamingARegisteredSource** | Pointer to **NullableInt32** | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional] 
**StudiesNamingARegisteredConductor** | Pointer to **NullableInt32** | How many registered studies name a conducting institution that is itself registered. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdStudiesRow

`func NewRwdStudiesRow() *RwdStudiesRow`

NewRwdStudiesRow instantiates a new RwdStudiesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdStudiesRowWithDefaults

`func NewRwdStudiesRowWithDefaults() *RwdStudiesRow`

NewRwdStudiesRowWithDefaults instantiates a new RwdStudiesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStudyId

`func (o *RwdStudiesRow) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *RwdStudiesRow) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *RwdStudiesRow) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.

### HasStudyId

`func (o *RwdStudiesRow) HasStudyId() bool`

HasStudyId returns a boolean if a field has been set.

### SetStudyIdNil

`func (o *RwdStudiesRow) SetStudyIdNil(b bool)`

 SetStudyIdNil sets the value for StudyId to be an explicit nil

### UnsetStudyId
`func (o *RwdStudiesRow) UnsetStudyId()`

UnsetStudyId ensures that no value is present for StudyId, not even an explicit nil
### GetEuPasNumber

`func (o *RwdStudiesRow) GetEuPasNumber() string`

GetEuPasNumber returns the EuPasNumber field if non-nil, zero value otherwise.

### GetEuPasNumberOk

`func (o *RwdStudiesRow) GetEuPasNumberOk() (*string, bool)`

GetEuPasNumberOk returns a tuple with the EuPasNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEuPasNumber

`func (o *RwdStudiesRow) SetEuPasNumber(v string)`

SetEuPasNumber sets EuPasNumber field to given value.

### HasEuPasNumber

`func (o *RwdStudiesRow) HasEuPasNumber() bool`

HasEuPasNumber returns a boolean if a field has been set.

### SetEuPasNumberNil

`func (o *RwdStudiesRow) SetEuPasNumberNil(b bool)`

 SetEuPasNumberNil sets the value for EuPasNumber to be an explicit nil

### UnsetEuPasNumber
`func (o *RwdStudiesRow) UnsetEuPasNumber()`

UnsetEuPasNumber ensures that no value is present for EuPasNumber, not even an explicit nil
### GetStudyTitle

`func (o *RwdStudiesRow) GetStudyTitle() string`

GetStudyTitle returns the StudyTitle field if non-nil, zero value otherwise.

### GetStudyTitleOk

`func (o *RwdStudiesRow) GetStudyTitleOk() (*string, bool)`

GetStudyTitleOk returns a tuple with the StudyTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyTitle

`func (o *RwdStudiesRow) SetStudyTitle(v string)`

SetStudyTitle sets StudyTitle field to given value.

### HasStudyTitle

`func (o *RwdStudiesRow) HasStudyTitle() bool`

HasStudyTitle returns a boolean if a field has been set.

### SetStudyTitleNil

`func (o *RwdStudiesRow) SetStudyTitleNil(b bool)`

 SetStudyTitleNil sets the value for StudyTitle to be an explicit nil

### UnsetStudyTitle
`func (o *RwdStudiesRow) UnsetStudyTitle()`

UnsetStudyTitle ensures that no value is present for StudyTitle, not even an explicit nil
### GetCatalogueUrl

`func (o *RwdStudiesRow) GetCatalogueUrl() string`

GetCatalogueUrl returns the CatalogueUrl field if non-nil, zero value otherwise.

### GetCatalogueUrlOk

`func (o *RwdStudiesRow) GetCatalogueUrlOk() (*string, bool)`

GetCatalogueUrlOk returns a tuple with the CatalogueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueUrl

`func (o *RwdStudiesRow) SetCatalogueUrl(v string)`

SetCatalogueUrl sets CatalogueUrl field to given value.

### HasCatalogueUrl

`func (o *RwdStudiesRow) HasCatalogueUrl() bool`

HasCatalogueUrl returns a boolean if a field has been set.

### SetCatalogueUrlNil

`func (o *RwdStudiesRow) SetCatalogueUrlNil(b bool)`

 SetCatalogueUrlNil sets the value for CatalogueUrl to be an explicit nil

### UnsetCatalogueUrl
`func (o *RwdStudiesRow) UnsetCatalogueUrl()`

UnsetCatalogueUrl ensures that no value is present for CatalogueUrl, not even an explicit nil
### GetStudyStatus

`func (o *RwdStudiesRow) GetStudyStatus() string`

GetStudyStatus returns the StudyStatus field if non-nil, zero value otherwise.

### GetStudyStatusOk

`func (o *RwdStudiesRow) GetStudyStatusOk() (*string, bool)`

GetStudyStatusOk returns a tuple with the StudyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStatus

`func (o *RwdStudiesRow) SetStudyStatus(v string)`

SetStudyStatus sets StudyStatus field to given value.

### HasStudyStatus

`func (o *RwdStudiesRow) HasStudyStatus() bool`

HasStudyStatus returns a boolean if a field has been set.

### SetStudyStatusNil

`func (o *RwdStudiesRow) SetStudyStatusNil(b bool)`

 SetStudyStatusNil sets the value for StudyStatus to be an explicit nil

### UnsetStudyStatus
`func (o *RwdStudiesRow) UnsetStudyStatus()`

UnsetStudyStatus ensures that no value is present for StudyStatus, not even an explicit nil
### GetIsOngoing

`func (o *RwdStudiesRow) GetIsOngoing() bool`

GetIsOngoing returns the IsOngoing field if non-nil, zero value otherwise.

### GetIsOngoingOk

`func (o *RwdStudiesRow) GetIsOngoingOk() (*bool, bool)`

GetIsOngoingOk returns a tuple with the IsOngoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOngoing

`func (o *RwdStudiesRow) SetIsOngoing(v bool)`

SetIsOngoing sets IsOngoing field to given value.

### HasIsOngoing

`func (o *RwdStudiesRow) HasIsOngoing() bool`

HasIsOngoing returns a boolean if a field has been set.

### SetIsOngoingNil

`func (o *RwdStudiesRow) SetIsOngoingNil(b bool)`

 SetIsOngoingNil sets the value for IsOngoing to be an explicit nil

### UnsetIsOngoing
`func (o *RwdStudiesRow) UnsetIsOngoing()`

UnsetIsOngoing ensures that no value is present for IsOngoing, not even an explicit nil
### GetIsFinalised

`func (o *RwdStudiesRow) GetIsFinalised() bool`

GetIsFinalised returns the IsFinalised field if non-nil, zero value otherwise.

### GetIsFinalisedOk

`func (o *RwdStudiesRow) GetIsFinalisedOk() (*bool, bool)`

GetIsFinalisedOk returns a tuple with the IsFinalised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinalised

`func (o *RwdStudiesRow) SetIsFinalised(v bool)`

SetIsFinalised sets IsFinalised field to given value.

### HasIsFinalised

`func (o *RwdStudiesRow) HasIsFinalised() bool`

HasIsFinalised returns a boolean if a field has been set.

### SetIsFinalisedNil

`func (o *RwdStudiesRow) SetIsFinalisedNil(b bool)`

 SetIsFinalisedNil sets the value for IsFinalised to be an explicit nil

### UnsetIsFinalised
`func (o *RwdStudiesRow) UnsetIsFinalised()`

UnsetIsFinalised ensures that no value is present for IsFinalised, not even an explicit nil
### GetIsDarwinEuStudy

`func (o *RwdStudiesRow) GetIsDarwinEuStudy() bool`

GetIsDarwinEuStudy returns the IsDarwinEuStudy field if non-nil, zero value otherwise.

### GetIsDarwinEuStudyOk

`func (o *RwdStudiesRow) GetIsDarwinEuStudyOk() (*bool, bool)`

GetIsDarwinEuStudyOk returns a tuple with the IsDarwinEuStudy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDarwinEuStudy

`func (o *RwdStudiesRow) SetIsDarwinEuStudy(v bool)`

SetIsDarwinEuStudy sets IsDarwinEuStudy field to given value.

### HasIsDarwinEuStudy

`func (o *RwdStudiesRow) HasIsDarwinEuStudy() bool`

HasIsDarwinEuStudy returns a boolean if a field has been set.

### SetIsDarwinEuStudyNil

`func (o *RwdStudiesRow) SetIsDarwinEuStudyNil(b bool)`

 SetIsDarwinEuStudyNil sets the value for IsDarwinEuStudy to be an explicit nil

### UnsetIsDarwinEuStudy
`func (o *RwdStudiesRow) UnsetIsDarwinEuStudy()`

UnsetIsDarwinEuStudy ensures that no value is present for IsDarwinEuStudy, not even an explicit nil
### GetHasEnceppSeal

`func (o *RwdStudiesRow) GetHasEnceppSeal() bool`

GetHasEnceppSeal returns the HasEnceppSeal field if non-nil, zero value otherwise.

### GetHasEnceppSealOk

`func (o *RwdStudiesRow) GetHasEnceppSealOk() (*bool, bool)`

GetHasEnceppSealOk returns a tuple with the HasEnceppSeal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEnceppSeal

`func (o *RwdStudiesRow) SetHasEnceppSeal(v bool)`

SetHasEnceppSeal sets HasEnceppSeal field to given value.

### HasHasEnceppSeal

`func (o *RwdStudiesRow) HasHasEnceppSeal() bool`

HasHasEnceppSeal returns a boolean if a field has been set.

### SetHasEnceppSealNil

`func (o *RwdStudiesRow) SetHasEnceppSealNil(b bool)`

 SetHasEnceppSealNil sets the value for HasEnceppSeal to be an explicit nil

### UnsetHasEnceppSeal
`func (o *RwdStudiesRow) UnsetHasEnceppSeal()`

UnsetHasEnceppSeal ensures that no value is present for HasEnceppSeal, not even an explicit nil
### GetRequiredByRegulator

`func (o *RwdStudiesRow) GetRequiredByRegulator() string`

GetRequiredByRegulator returns the RequiredByRegulator field if non-nil, zero value otherwise.

### GetRequiredByRegulatorOk

`func (o *RwdStudiesRow) GetRequiredByRegulatorOk() (*string, bool)`

GetRequiredByRegulatorOk returns a tuple with the RequiredByRegulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredByRegulator

`func (o *RwdStudiesRow) SetRequiredByRegulator(v string)`

SetRequiredByRegulator sets RequiredByRegulator field to given value.

### HasRequiredByRegulator

`func (o *RwdStudiesRow) HasRequiredByRegulator() bool`

HasRequiredByRegulator returns a boolean if a field has been set.

### SetRequiredByRegulatorNil

`func (o *RwdStudiesRow) SetRequiredByRegulatorNil(b bool)`

 SetRequiredByRegulatorNil sets the value for RequiredByRegulator to be an explicit nil

### UnsetRequiredByRegulator
`func (o *RwdStudiesRow) UnsetRequiredByRegulator()`

UnsetRequiredByRegulator ensures that no value is present for RequiredByRegulator, not even an explicit nil
### GetMandateStratum

`func (o *RwdStudiesRow) GetMandateStratum() string`

GetMandateStratum returns the MandateStratum field if non-nil, zero value otherwise.

### GetMandateStratumOk

`func (o *RwdStudiesRow) GetMandateStratumOk() (*string, bool)`

GetMandateStratumOk returns a tuple with the MandateStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStratum

`func (o *RwdStudiesRow) SetMandateStratum(v string)`

SetMandateStratum sets MandateStratum field to given value.

### HasMandateStratum

`func (o *RwdStudiesRow) HasMandateStratum() bool`

HasMandateStratum returns a boolean if a field has been set.

### SetMandateStratumNil

`func (o *RwdStudiesRow) SetMandateStratumNil(b bool)`

 SetMandateStratumNil sets the value for MandateStratum to be an explicit nil

### UnsetMandateStratum
`func (o *RwdStudiesRow) UnsetMandateStratum()`

UnsetMandateStratum ensures that no value is present for MandateStratum, not even an explicit nil
### GetRmpRequirement

`func (o *RwdStudiesRow) GetRmpRequirement() string`

GetRmpRequirement returns the RmpRequirement field if non-nil, zero value otherwise.

### GetRmpRequirementOk

`func (o *RwdStudiesRow) GetRmpRequirementOk() (*string, bool)`

GetRmpRequirementOk returns a tuple with the RmpRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmpRequirement

`func (o *RwdStudiesRow) SetRmpRequirement(v string)`

SetRmpRequirement sets RmpRequirement field to given value.

### HasRmpRequirement

`func (o *RwdStudiesRow) HasRmpRequirement() bool`

HasRmpRequirement returns a boolean if a field has been set.

### SetRmpRequirementNil

`func (o *RwdStudiesRow) SetRmpRequirementNil(b bool)`

 SetRmpRequirementNil sets the value for RmpRequirement to be an explicit nil

### UnsetRmpRequirement
`func (o *RwdStudiesRow) UnsetRmpRequirement()`

UnsetRmpRequirement ensures that no value is present for RmpRequirement, not even an explicit nil
### GetHasRmpCategory

`func (o *RwdStudiesRow) GetHasRmpCategory() bool`

GetHasRmpCategory returns the HasRmpCategory field if non-nil, zero value otherwise.

### GetHasRmpCategoryOk

`func (o *RwdStudiesRow) GetHasRmpCategoryOk() (*bool, bool)`

GetHasRmpCategoryOk returns a tuple with the HasRmpCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRmpCategory

`func (o *RwdStudiesRow) SetHasRmpCategory(v bool)`

SetHasRmpCategory sets HasRmpCategory field to given value.

### HasHasRmpCategory

`func (o *RwdStudiesRow) HasHasRmpCategory() bool`

HasHasRmpCategory returns a boolean if a field has been set.

### SetHasRmpCategoryNil

`func (o *RwdStudiesRow) SetHasRmpCategoryNil(b bool)`

 SetHasRmpCategoryNil sets the value for HasRmpCategory to be an explicit nil

### UnsetHasRmpCategory
`func (o *RwdStudiesRow) UnsetHasRmpCategory()`

UnsetHasRmpCategory ensures that no value is present for HasRmpCategory, not even an explicit nil
### GetRegulatoryProcedureNumber

`func (o *RwdStudiesRow) GetRegulatoryProcedureNumber() string`

GetRegulatoryProcedureNumber returns the RegulatoryProcedureNumber field if non-nil, zero value otherwise.

### GetRegulatoryProcedureNumberOk

`func (o *RwdStudiesRow) GetRegulatoryProcedureNumberOk() (*string, bool)`

GetRegulatoryProcedureNumberOk returns a tuple with the RegulatoryProcedureNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryProcedureNumber

`func (o *RwdStudiesRow) SetRegulatoryProcedureNumber(v string)`

SetRegulatoryProcedureNumber sets RegulatoryProcedureNumber field to given value.

### HasRegulatoryProcedureNumber

`func (o *RwdStudiesRow) HasRegulatoryProcedureNumber() bool`

HasRegulatoryProcedureNumber returns a boolean if a field has been set.

### SetRegulatoryProcedureNumberNil

`func (o *RwdStudiesRow) SetRegulatoryProcedureNumberNil(b bool)`

 SetRegulatoryProcedureNumberNil sets the value for RegulatoryProcedureNumber to be an explicit nil

### UnsetRegulatoryProcedureNumber
`func (o *RwdStudiesRow) UnsetRegulatoryProcedureNumber()`

UnsetRegulatoryProcedureNumber ensures that no value is present for RegulatoryProcedureNumber, not even an explicit nil
### GetFundingSourceText

`func (o *RwdStudiesRow) GetFundingSourceText() string`

GetFundingSourceText returns the FundingSourceText field if non-nil, zero value otherwise.

### GetFundingSourceTextOk

`func (o *RwdStudiesRow) GetFundingSourceTextOk() (*string, bool)`

GetFundingSourceTextOk returns a tuple with the FundingSourceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceText

`func (o *RwdStudiesRow) SetFundingSourceText(v string)`

SetFundingSourceText sets FundingSourceText field to given value.

### HasFundingSourceText

`func (o *RwdStudiesRow) HasFundingSourceText() bool`

HasFundingSourceText returns a boolean if a field has been set.

### SetFundingSourceTextNil

`func (o *RwdStudiesRow) SetFundingSourceTextNil(b bool)`

 SetFundingSourceTextNil sets the value for FundingSourceText to be an explicit nil

### UnsetFundingSourceText
`func (o *RwdStudiesRow) UnsetFundingSourceText()`

UnsetFundingSourceText ensures that no value is present for FundingSourceText, not even an explicit nil
### GetFundingSources

`func (o *RwdStudiesRow) GetFundingSources() []string`

GetFundingSources returns the FundingSources field if non-nil, zero value otherwise.

### GetFundingSourcesOk

`func (o *RwdStudiesRow) GetFundingSourcesOk() (*[]string, bool)`

GetFundingSourcesOk returns a tuple with the FundingSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSources

`func (o *RwdStudiesRow) SetFundingSources(v []string)`

SetFundingSources sets FundingSources field to given value.

### HasFundingSources

`func (o *RwdStudiesRow) HasFundingSources() bool`

HasFundingSources returns a boolean if a field has been set.

### SetFundingSourcesNil

`func (o *RwdStudiesRow) SetFundingSourcesNil(b bool)`

 SetFundingSourcesNil sets the value for FundingSources to be an explicit nil

### UnsetFundingSources
`func (o *RwdStudiesRow) UnsetFundingSources()`

UnsetFundingSources ensures that no value is present for FundingSources, not even an explicit nil
### GetIsPharmaFunded

`func (o *RwdStudiesRow) GetIsPharmaFunded() bool`

GetIsPharmaFunded returns the IsPharmaFunded field if non-nil, zero value otherwise.

### GetIsPharmaFundedOk

`func (o *RwdStudiesRow) GetIsPharmaFundedOk() (*bool, bool)`

GetIsPharmaFundedOk returns a tuple with the IsPharmaFunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPharmaFunded

`func (o *RwdStudiesRow) SetIsPharmaFunded(v bool)`

SetIsPharmaFunded sets IsPharmaFunded field to given value.

### HasIsPharmaFunded

`func (o *RwdStudiesRow) HasIsPharmaFunded() bool`

HasIsPharmaFunded returns a boolean if a field has been set.

### SetIsPharmaFundedNil

`func (o *RwdStudiesRow) SetIsPharmaFundedNil(b bool)`

 SetIsPharmaFundedNil sets the value for IsPharmaFunded to be an explicit nil

### UnsetIsPharmaFunded
`func (o *RwdStudiesRow) UnsetIsPharmaFunded()`

UnsetIsPharmaFunded ensures that no value is present for IsPharmaFunded, not even an explicit nil
### GetRegisterStudyType

`func (o *RwdStudiesRow) GetRegisterStudyType() string`

GetRegisterStudyType returns the RegisterStudyType field if non-nil, zero value otherwise.

### GetRegisterStudyTypeOk

`func (o *RwdStudiesRow) GetRegisterStudyTypeOk() (*string, bool)`

GetRegisterStudyTypeOk returns a tuple with the RegisterStudyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterStudyType

`func (o *RwdStudiesRow) SetRegisterStudyType(v string)`

SetRegisterStudyType sets RegisterStudyType field to given value.

### HasRegisterStudyType

`func (o *RwdStudiesRow) HasRegisterStudyType() bool`

HasRegisterStudyType returns a boolean if a field has been set.

### SetRegisterStudyTypeNil

`func (o *RwdStudiesRow) SetRegisterStudyTypeNil(b bool)`

 SetRegisterStudyTypeNil sets the value for RegisterStudyType to be an explicit nil

### UnsetRegisterStudyType
`func (o *RwdStudiesRow) UnsetRegisterStudyType()`

UnsetRegisterStudyType ensures that no value is present for RegisterStudyType, not even an explicit nil
### GetScopeOfTheStudy

`func (o *RwdStudiesRow) GetScopeOfTheStudy() string`

GetScopeOfTheStudy returns the ScopeOfTheStudy field if non-nil, zero value otherwise.

### GetScopeOfTheStudyOk

`func (o *RwdStudiesRow) GetScopeOfTheStudyOk() (*string, bool)`

GetScopeOfTheStudyOk returns a tuple with the ScopeOfTheStudy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOfTheStudy

`func (o *RwdStudiesRow) SetScopeOfTheStudy(v string)`

SetScopeOfTheStudy sets ScopeOfTheStudy field to given value.

### HasScopeOfTheStudy

`func (o *RwdStudiesRow) HasScopeOfTheStudy() bool`

HasScopeOfTheStudy returns a boolean if a field has been set.

### SetScopeOfTheStudyNil

`func (o *RwdStudiesRow) SetScopeOfTheStudyNil(b bool)`

 SetScopeOfTheStudyNil sets the value for ScopeOfTheStudy to be an explicit nil

### UnsetScopeOfTheStudy
`func (o *RwdStudiesRow) UnsetScopeOfTheStudy()`

UnsetScopeOfTheStudy ensures that no value is present for ScopeOfTheStudy, not even an explicit nil
### GetStudyTopic

`func (o *RwdStudiesRow) GetStudyTopic() string`

GetStudyTopic returns the StudyTopic field if non-nil, zero value otherwise.

### GetStudyTopicOk

`func (o *RwdStudiesRow) GetStudyTopicOk() (*string, bool)`

GetStudyTopicOk returns a tuple with the StudyTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyTopic

`func (o *RwdStudiesRow) SetStudyTopic(v string)`

SetStudyTopic sets StudyTopic field to given value.

### HasStudyTopic

`func (o *RwdStudiesRow) HasStudyTopic() bool`

HasStudyTopic returns a boolean if a field has been set.

### SetStudyTopicNil

`func (o *RwdStudiesRow) SetStudyTopicNil(b bool)`

 SetStudyTopicNil sets the value for StudyTopic to be an explicit nil

### UnsetStudyTopic
`func (o *RwdStudiesRow) UnsetStudyTopic()`

UnsetStudyTopic ensures that no value is present for StudyTopic, not even an explicit nil
### GetStudyDesign

`func (o *RwdStudiesRow) GetStudyDesign() string`

GetStudyDesign returns the StudyDesign field if non-nil, zero value otherwise.

### GetStudyDesignOk

`func (o *RwdStudiesRow) GetStudyDesignOk() (*string, bool)`

GetStudyDesignOk returns a tuple with the StudyDesign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyDesign

`func (o *RwdStudiesRow) SetStudyDesign(v string)`

SetStudyDesign sets StudyDesign field to given value.

### HasStudyDesign

`func (o *RwdStudiesRow) HasStudyDesign() bool`

HasStudyDesign returns a boolean if a field has been set.

### SetStudyDesignNil

`func (o *RwdStudiesRow) SetStudyDesignNil(b bool)`

 SetStudyDesignNil sets the value for StudyDesign to be an explicit nil

### UnsetStudyDesign
`func (o *RwdStudiesRow) UnsetStudyDesign()`

UnsetStudyDesign ensures that no value is present for StudyDesign, not even an explicit nil
### GetNonInterventionalStudyDesign

`func (o *RwdStudiesRow) GetNonInterventionalStudyDesign() string`

GetNonInterventionalStudyDesign returns the NonInterventionalStudyDesign field if non-nil, zero value otherwise.

### GetNonInterventionalStudyDesignOk

`func (o *RwdStudiesRow) GetNonInterventionalStudyDesignOk() (*string, bool)`

GetNonInterventionalStudyDesignOk returns a tuple with the NonInterventionalStudyDesign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonInterventionalStudyDesign

`func (o *RwdStudiesRow) SetNonInterventionalStudyDesign(v string)`

SetNonInterventionalStudyDesign sets NonInterventionalStudyDesign field to given value.

### HasNonInterventionalStudyDesign

`func (o *RwdStudiesRow) HasNonInterventionalStudyDesign() bool`

HasNonInterventionalStudyDesign returns a boolean if a field has been set.

### SetNonInterventionalStudyDesignNil

`func (o *RwdStudiesRow) SetNonInterventionalStudyDesignNil(b bool)`

 SetNonInterventionalStudyDesignNil sets the value for NonInterventionalStudyDesign to be an explicit nil

### UnsetNonInterventionalStudyDesign
`func (o *RwdStudiesRow) UnsetNonInterventionalStudyDesign()`

UnsetNonInterventionalStudyDesign ensures that no value is present for NonInterventionalStudyDesign, not even an explicit nil
### GetMainStudyObjective

`func (o *RwdStudiesRow) GetMainStudyObjective() string`

GetMainStudyObjective returns the MainStudyObjective field if non-nil, zero value otherwise.

### GetMainStudyObjectiveOk

`func (o *RwdStudiesRow) GetMainStudyObjectiveOk() (*string, bool)`

GetMainStudyObjectiveOk returns a tuple with the MainStudyObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainStudyObjective

`func (o *RwdStudiesRow) SetMainStudyObjective(v string)`

SetMainStudyObjective sets MainStudyObjective field to given value.

### HasMainStudyObjective

`func (o *RwdStudiesRow) HasMainStudyObjective() bool`

HasMainStudyObjective returns a boolean if a field has been set.

### SetMainStudyObjectiveNil

`func (o *RwdStudiesRow) SetMainStudyObjectiveNil(b bool)`

 SetMainStudyObjectiveNil sets the value for MainStudyObjective to be an explicit nil

### UnsetMainStudyObjective
`func (o *RwdStudiesRow) UnsetMainStudyObjective()`

UnsetMainStudyObjective ensures that no value is present for MainStudyObjective, not even an explicit nil
### GetDataCollectionMethods

`func (o *RwdStudiesRow) GetDataCollectionMethods() string`

GetDataCollectionMethods returns the DataCollectionMethods field if non-nil, zero value otherwise.

### GetDataCollectionMethodsOk

`func (o *RwdStudiesRow) GetDataCollectionMethodsOk() (*string, bool)`

GetDataCollectionMethodsOk returns a tuple with the DataCollectionMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectionMethods

`func (o *RwdStudiesRow) SetDataCollectionMethods(v string)`

SetDataCollectionMethods sets DataCollectionMethods field to given value.

### HasDataCollectionMethods

`func (o *RwdStudiesRow) HasDataCollectionMethods() bool`

HasDataCollectionMethods returns a boolean if a field has been set.

### SetDataCollectionMethodsNil

`func (o *RwdStudiesRow) SetDataCollectionMethodsNil(b bool)`

 SetDataCollectionMethodsNil sets the value for DataCollectionMethods to be an explicit nil

### UnsetDataCollectionMethods
`func (o *RwdStudiesRow) UnsetDataCollectionMethods()`

UnsetDataCollectionMethods ensures that no value is present for DataCollectionMethods, not even an explicit nil
### GetMedicinalProductName

`func (o *RwdStudiesRow) GetMedicinalProductName() string`

GetMedicinalProductName returns the MedicinalProductName field if non-nil, zero value otherwise.

### GetMedicinalProductNameOk

`func (o *RwdStudiesRow) GetMedicinalProductNameOk() (*string, bool)`

GetMedicinalProductNameOk returns a tuple with the MedicinalProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicinalProductName

`func (o *RwdStudiesRow) SetMedicinalProductName(v string)`

SetMedicinalProductName sets MedicinalProductName field to given value.

### HasMedicinalProductName

`func (o *RwdStudiesRow) HasMedicinalProductName() bool`

HasMedicinalProductName returns a boolean if a field has been set.

### SetMedicinalProductNameNil

`func (o *RwdStudiesRow) SetMedicinalProductNameNil(b bool)`

 SetMedicinalProductNameNil sets the value for MedicinalProductName to be an explicit nil

### UnsetMedicinalProductName
`func (o *RwdStudiesRow) UnsetMedicinalProductName()`

UnsetMedicinalProductName ensures that no value is present for MedicinalProductName, not even an explicit nil
### GetStudyDrugInn

`func (o *RwdStudiesRow) GetStudyDrugInn() string`

GetStudyDrugInn returns the StudyDrugInn field if non-nil, zero value otherwise.

### GetStudyDrugInnOk

`func (o *RwdStudiesRow) GetStudyDrugInnOk() (*string, bool)`

GetStudyDrugInnOk returns a tuple with the StudyDrugInn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyDrugInn

`func (o *RwdStudiesRow) SetStudyDrugInn(v string)`

SetStudyDrugInn sets StudyDrugInn field to given value.

### HasStudyDrugInn

`func (o *RwdStudiesRow) HasStudyDrugInn() bool`

HasStudyDrugInn returns a boolean if a field has been set.

### SetStudyDrugInnNil

`func (o *RwdStudiesRow) SetStudyDrugInnNil(b bool)`

 SetStudyDrugInnNil sets the value for StudyDrugInn to be an explicit nil

### UnsetStudyDrugInn
`func (o *RwdStudiesRow) UnsetStudyDrugInn()`

UnsetStudyDrugInn ensures that no value is present for StudyDrugInn, not even an explicit nil
### GetAtcCodeText

`func (o *RwdStudiesRow) GetAtcCodeText() string`

GetAtcCodeText returns the AtcCodeText field if non-nil, zero value otherwise.

### GetAtcCodeTextOk

`func (o *RwdStudiesRow) GetAtcCodeTextOk() (*string, bool)`

GetAtcCodeTextOk returns a tuple with the AtcCodeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcCodeText

`func (o *RwdStudiesRow) SetAtcCodeText(v string)`

SetAtcCodeText sets AtcCodeText field to given value.

### HasAtcCodeText

`func (o *RwdStudiesRow) HasAtcCodeText() bool`

HasAtcCodeText returns a boolean if a field has been set.

### SetAtcCodeTextNil

`func (o *RwdStudiesRow) SetAtcCodeTextNil(b bool)`

 SetAtcCodeTextNil sets the value for AtcCodeText to be an explicit nil

### UnsetAtcCodeText
`func (o *RwdStudiesRow) UnsetAtcCodeText()`

UnsetAtcCodeText ensures that no value is present for AtcCodeText, not even an explicit nil
### GetDrugMentionCount

`func (o *RwdStudiesRow) GetDrugMentionCount() int32`

GetDrugMentionCount returns the DrugMentionCount field if non-nil, zero value otherwise.

### GetDrugMentionCountOk

`func (o *RwdStudiesRow) GetDrugMentionCountOk() (*int32, bool)`

GetDrugMentionCountOk returns a tuple with the DrugMentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugMentionCount

`func (o *RwdStudiesRow) SetDrugMentionCount(v int32)`

SetDrugMentionCount sets DrugMentionCount field to given value.

### HasDrugMentionCount

`func (o *RwdStudiesRow) HasDrugMentionCount() bool`

HasDrugMentionCount returns a boolean if a field has been set.

### SetDrugMentionCountNil

`func (o *RwdStudiesRow) SetDrugMentionCountNil(b bool)`

 SetDrugMentionCountNil sets the value for DrugMentionCount to be an explicit nil

### UnsetDrugMentionCount
`func (o *RwdStudiesRow) UnsetDrugMentionCount()`

UnsetDrugMentionCount ensures that no value is present for DrugMentionCount, not even an explicit nil
### GetAtcLevel5MentionCount

`func (o *RwdStudiesRow) GetAtcLevel5MentionCount() int32`

GetAtcLevel5MentionCount returns the AtcLevel5MentionCount field if non-nil, zero value otherwise.

### GetAtcLevel5MentionCountOk

`func (o *RwdStudiesRow) GetAtcLevel5MentionCountOk() (*int32, bool)`

GetAtcLevel5MentionCountOk returns a tuple with the AtcLevel5MentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcLevel5MentionCount

`func (o *RwdStudiesRow) SetAtcLevel5MentionCount(v int32)`

SetAtcLevel5MentionCount sets AtcLevel5MentionCount field to given value.

### HasAtcLevel5MentionCount

`func (o *RwdStudiesRow) HasAtcLevel5MentionCount() bool`

HasAtcLevel5MentionCount returns a boolean if a field has been set.

### SetAtcLevel5MentionCountNil

`func (o *RwdStudiesRow) SetAtcLevel5MentionCountNil(b bool)`

 SetAtcLevel5MentionCountNil sets the value for AtcLevel5MentionCount to be an explicit nil

### UnsetAtcLevel5MentionCount
`func (o *RwdStudiesRow) UnsetAtcLevel5MentionCount()`

UnsetAtcLevel5MentionCount ensures that no value is present for AtcLevel5MentionCount, not even an explicit nil
### GetSpineMatchedMentionCount

`func (o *RwdStudiesRow) GetSpineMatchedMentionCount() int32`

GetSpineMatchedMentionCount returns the SpineMatchedMentionCount field if non-nil, zero value otherwise.

### GetSpineMatchedMentionCountOk

`func (o *RwdStudiesRow) GetSpineMatchedMentionCountOk() (*int32, bool)`

GetSpineMatchedMentionCountOk returns a tuple with the SpineMatchedMentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpineMatchedMentionCount

`func (o *RwdStudiesRow) SetSpineMatchedMentionCount(v int32)`

SetSpineMatchedMentionCount sets SpineMatchedMentionCount field to given value.

### HasSpineMatchedMentionCount

`func (o *RwdStudiesRow) HasSpineMatchedMentionCount() bool`

HasSpineMatchedMentionCount returns a boolean if a field has been set.

### SetSpineMatchedMentionCountNil

`func (o *RwdStudiesRow) SetSpineMatchedMentionCountNil(b bool)`

 SetSpineMatchedMentionCountNil sets the value for SpineMatchedMentionCount to be an explicit nil

### UnsetSpineMatchedMentionCount
`func (o *RwdStudiesRow) UnsetSpineMatchedMentionCount()`

UnsetSpineMatchedMentionCount ensures that no value is present for SpineMatchedMentionCount, not even an explicit nil
### GetAtcLevel5Codes

`func (o *RwdStudiesRow) GetAtcLevel5Codes() []string`

GetAtcLevel5Codes returns the AtcLevel5Codes field if non-nil, zero value otherwise.

### GetAtcLevel5CodesOk

`func (o *RwdStudiesRow) GetAtcLevel5CodesOk() (*[]string, bool)`

GetAtcLevel5CodesOk returns a tuple with the AtcLevel5Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcLevel5Codes

`func (o *RwdStudiesRow) SetAtcLevel5Codes(v []string)`

SetAtcLevel5Codes sets AtcLevel5Codes field to given value.

### HasAtcLevel5Codes

`func (o *RwdStudiesRow) HasAtcLevel5Codes() bool`

HasAtcLevel5Codes returns a boolean if a field has been set.

### SetAtcLevel5CodesNil

`func (o *RwdStudiesRow) SetAtcLevel5CodesNil(b bool)`

 SetAtcLevel5CodesNil sets the value for AtcLevel5Codes to be an explicit nil

### UnsetAtcLevel5Codes
`func (o *RwdStudiesRow) UnsetAtcLevel5Codes()`

UnsetAtcLevel5Codes ensures that no value is present for AtcLevel5Codes, not even an explicit nil
### GetAtcLevel1Codes

`func (o *RwdStudiesRow) GetAtcLevel1Codes() []string`

GetAtcLevel1Codes returns the AtcLevel1Codes field if non-nil, zero value otherwise.

### GetAtcLevel1CodesOk

`func (o *RwdStudiesRow) GetAtcLevel1CodesOk() (*[]string, bool)`

GetAtcLevel1CodesOk returns a tuple with the AtcLevel1Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcLevel1Codes

`func (o *RwdStudiesRow) SetAtcLevel1Codes(v []string)`

SetAtcLevel1Codes sets AtcLevel1Codes field to given value.

### HasAtcLevel1Codes

`func (o *RwdStudiesRow) HasAtcLevel1Codes() bool`

HasAtcLevel1Codes returns a boolean if a field has been set.

### SetAtcLevel1CodesNil

`func (o *RwdStudiesRow) SetAtcLevel1CodesNil(b bool)`

 SetAtcLevel1CodesNil sets the value for AtcLevel1Codes to be an explicit nil

### UnsetAtcLevel1Codes
`func (o *RwdStudiesRow) UnsetAtcLevel1Codes()`

UnsetAtcLevel1Codes ensures that no value is present for AtcLevel1Codes, not even an explicit nil
### GetMedicalCondition

`func (o *RwdStudiesRow) GetMedicalCondition() string`

GetMedicalCondition returns the MedicalCondition field if non-nil, zero value otherwise.

### GetMedicalConditionOk

`func (o *RwdStudiesRow) GetMedicalConditionOk() (*string, bool)`

GetMedicalConditionOk returns a tuple with the MedicalCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicalCondition

`func (o *RwdStudiesRow) SetMedicalCondition(v string)`

SetMedicalCondition sets MedicalCondition field to given value.

### HasMedicalCondition

`func (o *RwdStudiesRow) HasMedicalCondition() bool`

HasMedicalCondition returns a boolean if a field has been set.

### SetMedicalConditionNil

`func (o *RwdStudiesRow) SetMedicalConditionNil(b bool)`

 SetMedicalConditionNil sets the value for MedicalCondition to be an explicit nil

### UnsetMedicalCondition
`func (o *RwdStudiesRow) UnsetMedicalCondition()`

UnsetMedicalCondition ensures that no value is present for MedicalCondition, not even an explicit nil
### GetAdditionalMedicalConditions

`func (o *RwdStudiesRow) GetAdditionalMedicalConditions() string`

GetAdditionalMedicalConditions returns the AdditionalMedicalConditions field if non-nil, zero value otherwise.

### GetAdditionalMedicalConditionsOk

`func (o *RwdStudiesRow) GetAdditionalMedicalConditionsOk() (*string, bool)`

GetAdditionalMedicalConditionsOk returns a tuple with the AdditionalMedicalConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMedicalConditions

`func (o *RwdStudiesRow) SetAdditionalMedicalConditions(v string)`

SetAdditionalMedicalConditions sets AdditionalMedicalConditions field to given value.

### HasAdditionalMedicalConditions

`func (o *RwdStudiesRow) HasAdditionalMedicalConditions() bool`

HasAdditionalMedicalConditions returns a boolean if a field has been set.

### SetAdditionalMedicalConditionsNil

`func (o *RwdStudiesRow) SetAdditionalMedicalConditionsNil(b bool)`

 SetAdditionalMedicalConditionsNil sets the value for AdditionalMedicalConditions to be an explicit nil

### UnsetAdditionalMedicalConditions
`func (o *RwdStudiesRow) UnsetAdditionalMedicalConditions()`

UnsetAdditionalMedicalConditions ensures that no value is present for AdditionalMedicalConditions, not even an explicit nil
### GetConditionCount

`func (o *RwdStudiesRow) GetConditionCount() int32`

GetConditionCount returns the ConditionCount field if non-nil, zero value otherwise.

### GetConditionCountOk

`func (o *RwdStudiesRow) GetConditionCountOk() (*int32, bool)`

GetConditionCountOk returns a tuple with the ConditionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionCount

`func (o *RwdStudiesRow) SetConditionCount(v int32)`

SetConditionCount sets ConditionCount field to given value.

### HasConditionCount

`func (o *RwdStudiesRow) HasConditionCount() bool`

HasConditionCount returns a boolean if a field has been set.

### SetConditionCountNil

`func (o *RwdStudiesRow) SetConditionCountNil(b bool)`

 SetConditionCountNil sets the value for ConditionCount to be an explicit nil

### UnsetConditionCount
`func (o *RwdStudiesRow) UnsetConditionCount()`

UnsetConditionCount ensures that no value is present for ConditionCount, not even an explicit nil
### GetPrimaryConditionCount

`func (o *RwdStudiesRow) GetPrimaryConditionCount() int32`

GetPrimaryConditionCount returns the PrimaryConditionCount field if non-nil, zero value otherwise.

### GetPrimaryConditionCountOk

`func (o *RwdStudiesRow) GetPrimaryConditionCountOk() (*int32, bool)`

GetPrimaryConditionCountOk returns a tuple with the PrimaryConditionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryConditionCount

`func (o *RwdStudiesRow) SetPrimaryConditionCount(v int32)`

SetPrimaryConditionCount sets PrimaryConditionCount field to given value.

### HasPrimaryConditionCount

`func (o *RwdStudiesRow) HasPrimaryConditionCount() bool`

HasPrimaryConditionCount returns a boolean if a field has been set.

### SetPrimaryConditionCountNil

`func (o *RwdStudiesRow) SetPrimaryConditionCountNil(b bool)`

 SetPrimaryConditionCountNil sets the value for PrimaryConditionCount to be an explicit nil

### UnsetPrimaryConditionCount
`func (o *RwdStudiesRow) UnsetPrimaryConditionCount()`

UnsetPrimaryConditionCount ensures that no value is present for PrimaryConditionCount, not even an explicit nil
### GetMeshConditionCount

`func (o *RwdStudiesRow) GetMeshConditionCount() int32`

GetMeshConditionCount returns the MeshConditionCount field if non-nil, zero value otherwise.

### GetMeshConditionCountOk

`func (o *RwdStudiesRow) GetMeshConditionCountOk() (*int32, bool)`

GetMeshConditionCountOk returns a tuple with the MeshConditionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshConditionCount

`func (o *RwdStudiesRow) SetMeshConditionCount(v int32)`

SetMeshConditionCount sets MeshConditionCount field to given value.

### HasMeshConditionCount

`func (o *RwdStudiesRow) HasMeshConditionCount() bool`

HasMeshConditionCount returns a boolean if a field has been set.

### SetMeshConditionCountNil

`func (o *RwdStudiesRow) SetMeshConditionCountNil(b bool)`

 SetMeshConditionCountNil sets the value for MeshConditionCount to be an explicit nil

### UnsetMeshConditionCount
`func (o *RwdStudiesRow) UnsetMeshConditionCount()`

UnsetMeshConditionCount ensures that no value is present for MeshConditionCount, not even an explicit nil
### GetConditionTerms

`func (o *RwdStudiesRow) GetConditionTerms() []string`

GetConditionTerms returns the ConditionTerms field if non-nil, zero value otherwise.

### GetConditionTermsOk

`func (o *RwdStudiesRow) GetConditionTermsOk() (*[]string, bool)`

GetConditionTermsOk returns a tuple with the ConditionTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTerms

`func (o *RwdStudiesRow) SetConditionTerms(v []string)`

SetConditionTerms sets ConditionTerms field to given value.

### HasConditionTerms

`func (o *RwdStudiesRow) HasConditionTerms() bool`

HasConditionTerms returns a boolean if a field has been set.

### SetConditionTermsNil

`func (o *RwdStudiesRow) SetConditionTermsNil(b bool)`

 SetConditionTermsNil sets the value for ConditionTerms to be an explicit nil

### UnsetConditionTerms
`func (o *RwdStudiesRow) UnsetConditionTerms()`

UnsetConditionTerms ensures that no value is present for ConditionTerms, not even an explicit nil
### GetPopulationAgeGroups

`func (o *RwdStudiesRow) GetPopulationAgeGroups() string`

GetPopulationAgeGroups returns the PopulationAgeGroups field if non-nil, zero value otherwise.

### GetPopulationAgeGroupsOk

`func (o *RwdStudiesRow) GetPopulationAgeGroupsOk() (*string, bool)`

GetPopulationAgeGroupsOk returns a tuple with the PopulationAgeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationAgeGroups

`func (o *RwdStudiesRow) SetPopulationAgeGroups(v string)`

SetPopulationAgeGroups sets PopulationAgeGroups field to given value.

### HasPopulationAgeGroups

`func (o *RwdStudiesRow) HasPopulationAgeGroups() bool`

HasPopulationAgeGroups returns a boolean if a field has been set.

### SetPopulationAgeGroupsNil

`func (o *RwdStudiesRow) SetPopulationAgeGroupsNil(b bool)`

 SetPopulationAgeGroupsNil sets the value for PopulationAgeGroups to be an explicit nil

### UnsetPopulationAgeGroups
`func (o *RwdStudiesRow) UnsetPopulationAgeGroups()`

UnsetPopulationAgeGroups ensures that no value is present for PopulationAgeGroups, not even an explicit nil
### GetSpecialPopulationOfInterest

`func (o *RwdStudiesRow) GetSpecialPopulationOfInterest() string`

GetSpecialPopulationOfInterest returns the SpecialPopulationOfInterest field if non-nil, zero value otherwise.

### GetSpecialPopulationOfInterestOk

`func (o *RwdStudiesRow) GetSpecialPopulationOfInterestOk() (*string, bool)`

GetSpecialPopulationOfInterestOk returns a tuple with the SpecialPopulationOfInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPopulationOfInterest

`func (o *RwdStudiesRow) SetSpecialPopulationOfInterest(v string)`

SetSpecialPopulationOfInterest sets SpecialPopulationOfInterest field to given value.

### HasSpecialPopulationOfInterest

`func (o *RwdStudiesRow) HasSpecialPopulationOfInterest() bool`

HasSpecialPopulationOfInterest returns a boolean if a field has been set.

### SetSpecialPopulationOfInterestNil

`func (o *RwdStudiesRow) SetSpecialPopulationOfInterestNil(b bool)`

 SetSpecialPopulationOfInterestNil sets the value for SpecialPopulationOfInterest to be an explicit nil

### UnsetSpecialPopulationOfInterest
`func (o *RwdStudiesRow) UnsetSpecialPopulationOfInterest()`

UnsetSpecialPopulationOfInterest ensures that no value is present for SpecialPopulationOfInterest, not even an explicit nil
### GetEstimatedNumberOfSubjects

`func (o *RwdStudiesRow) GetEstimatedNumberOfSubjects() int64`

GetEstimatedNumberOfSubjects returns the EstimatedNumberOfSubjects field if non-nil, zero value otherwise.

### GetEstimatedNumberOfSubjectsOk

`func (o *RwdStudiesRow) GetEstimatedNumberOfSubjectsOk() (*int64, bool)`

GetEstimatedNumberOfSubjectsOk returns a tuple with the EstimatedNumberOfSubjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedNumberOfSubjects

`func (o *RwdStudiesRow) SetEstimatedNumberOfSubjects(v int64)`

SetEstimatedNumberOfSubjects sets EstimatedNumberOfSubjects field to given value.

### HasEstimatedNumberOfSubjects

`func (o *RwdStudiesRow) HasEstimatedNumberOfSubjects() bool`

HasEstimatedNumberOfSubjects returns a boolean if a field has been set.

### SetEstimatedNumberOfSubjectsNil

`func (o *RwdStudiesRow) SetEstimatedNumberOfSubjectsNil(b bool)`

 SetEstimatedNumberOfSubjectsNil sets the value for EstimatedNumberOfSubjects to be an explicit nil

### UnsetEstimatedNumberOfSubjects
`func (o *RwdStudiesRow) UnsetEstimatedNumberOfSubjects()`

UnsetEstimatedNumberOfSubjects ensures that no value is present for EstimatedNumberOfSubjects, not even an explicit nil
### GetSubjectsReportedZero

`func (o *RwdStudiesRow) GetSubjectsReportedZero() bool`

GetSubjectsReportedZero returns the SubjectsReportedZero field if non-nil, zero value otherwise.

### GetSubjectsReportedZeroOk

`func (o *RwdStudiesRow) GetSubjectsReportedZeroOk() (*bool, bool)`

GetSubjectsReportedZeroOk returns a tuple with the SubjectsReportedZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectsReportedZero

`func (o *RwdStudiesRow) SetSubjectsReportedZero(v bool)`

SetSubjectsReportedZero sets SubjectsReportedZero field to given value.

### HasSubjectsReportedZero

`func (o *RwdStudiesRow) HasSubjectsReportedZero() bool`

HasSubjectsReportedZero returns a boolean if a field has been set.

### SetSubjectsReportedZeroNil

`func (o *RwdStudiesRow) SetSubjectsReportedZeroNil(b bool)`

 SetSubjectsReportedZeroNil sets the value for SubjectsReportedZero to be an explicit nil

### UnsetSubjectsReportedZero
`func (o *RwdStudiesRow) UnsetSubjectsReportedZero()`

UnsetSubjectsReportedZero ensures that no value is present for SubjectsReportedZero, not even an explicit nil
### GetSubjectsBand

`func (o *RwdStudiesRow) GetSubjectsBand() string`

GetSubjectsBand returns the SubjectsBand field if non-nil, zero value otherwise.

### GetSubjectsBandOk

`func (o *RwdStudiesRow) GetSubjectsBandOk() (*string, bool)`

GetSubjectsBandOk returns a tuple with the SubjectsBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectsBand

`func (o *RwdStudiesRow) SetSubjectsBand(v string)`

SetSubjectsBand sets SubjectsBand field to given value.

### HasSubjectsBand

`func (o *RwdStudiesRow) HasSubjectsBand() bool`

HasSubjectsBand returns a boolean if a field has been set.

### SetSubjectsBandNil

`func (o *RwdStudiesRow) SetSubjectsBandNil(b bool)`

 SetSubjectsBandNil sets the value for SubjectsBand to be an explicit nil

### UnsetSubjectsBand
`func (o *RwdStudiesRow) UnsetSubjectsBand()`

UnsetSubjectsBand ensures that no value is present for SubjectsBand, not even an explicit nil
### GetConductingInstitutionNames

`func (o *RwdStudiesRow) GetConductingInstitutionNames() []string`

GetConductingInstitutionNames returns the ConductingInstitutionNames field if non-nil, zero value otherwise.

### GetConductingInstitutionNamesOk

`func (o *RwdStudiesRow) GetConductingInstitutionNamesOk() (*[]string, bool)`

GetConductingInstitutionNamesOk returns a tuple with the ConductingInstitutionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductingInstitutionNames

`func (o *RwdStudiesRow) SetConductingInstitutionNames(v []string)`

SetConductingInstitutionNames sets ConductingInstitutionNames field to given value.

### HasConductingInstitutionNames

`func (o *RwdStudiesRow) HasConductingInstitutionNames() bool`

HasConductingInstitutionNames returns a boolean if a field has been set.

### SetConductingInstitutionNamesNil

`func (o *RwdStudiesRow) SetConductingInstitutionNamesNil(b bool)`

 SetConductingInstitutionNamesNil sets the value for ConductingInstitutionNames to be an explicit nil

### UnsetConductingInstitutionNames
`func (o *RwdStudiesRow) UnsetConductingInstitutionNames()`

UnsetConductingInstitutionNames ensures that no value is present for ConductingInstitutionNames, not even an explicit nil
### GetConductingInstitutionCount

`func (o *RwdStudiesRow) GetConductingInstitutionCount() int32`

GetConductingInstitutionCount returns the ConductingInstitutionCount field if non-nil, zero value otherwise.

### GetConductingInstitutionCountOk

`func (o *RwdStudiesRow) GetConductingInstitutionCountOk() (*int32, bool)`

GetConductingInstitutionCountOk returns a tuple with the ConductingInstitutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductingInstitutionCount

`func (o *RwdStudiesRow) SetConductingInstitutionCount(v int32)`

SetConductingInstitutionCount sets ConductingInstitutionCount field to given value.

### HasConductingInstitutionCount

`func (o *RwdStudiesRow) HasConductingInstitutionCount() bool`

HasConductingInstitutionCount returns a boolean if a field has been set.

### SetConductingInstitutionCountNil

`func (o *RwdStudiesRow) SetConductingInstitutionCountNil(b bool)`

 SetConductingInstitutionCountNil sets the value for ConductingInstitutionCount to be an explicit nil

### UnsetConductingInstitutionCount
`func (o *RwdStudiesRow) UnsetConductingInstitutionCount()`

UnsetConductingInstitutionCount ensures that no value is present for ConductingInstitutionCount, not even an explicit nil
### GetAdditionalInstitutionNames

`func (o *RwdStudiesRow) GetAdditionalInstitutionNames() []string`

GetAdditionalInstitutionNames returns the AdditionalInstitutionNames field if non-nil, zero value otherwise.

### GetAdditionalInstitutionNamesOk

`func (o *RwdStudiesRow) GetAdditionalInstitutionNamesOk() (*[]string, bool)`

GetAdditionalInstitutionNamesOk returns a tuple with the AdditionalInstitutionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInstitutionNames

`func (o *RwdStudiesRow) SetAdditionalInstitutionNames(v []string)`

SetAdditionalInstitutionNames sets AdditionalInstitutionNames field to given value.

### HasAdditionalInstitutionNames

`func (o *RwdStudiesRow) HasAdditionalInstitutionNames() bool`

HasAdditionalInstitutionNames returns a boolean if a field has been set.

### SetAdditionalInstitutionNamesNil

`func (o *RwdStudiesRow) SetAdditionalInstitutionNamesNil(b bool)`

 SetAdditionalInstitutionNamesNil sets the value for AdditionalInstitutionNames to be an explicit nil

### UnsetAdditionalInstitutionNames
`func (o *RwdStudiesRow) UnsetAdditionalInstitutionNames()`

UnsetAdditionalInstitutionNames ensures that no value is present for AdditionalInstitutionNames, not even an explicit nil
### GetAdditionalInstitutionCount

`func (o *RwdStudiesRow) GetAdditionalInstitutionCount() int32`

GetAdditionalInstitutionCount returns the AdditionalInstitutionCount field if non-nil, zero value otherwise.

### GetAdditionalInstitutionCountOk

`func (o *RwdStudiesRow) GetAdditionalInstitutionCountOk() (*int32, bool)`

GetAdditionalInstitutionCountOk returns a tuple with the AdditionalInstitutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInstitutionCount

`func (o *RwdStudiesRow) SetAdditionalInstitutionCount(v int32)`

SetAdditionalInstitutionCount sets AdditionalInstitutionCount field to given value.

### HasAdditionalInstitutionCount

`func (o *RwdStudiesRow) HasAdditionalInstitutionCount() bool`

HasAdditionalInstitutionCount returns a boolean if a field has been set.

### SetAdditionalInstitutionCountNil

`func (o *RwdStudiesRow) SetAdditionalInstitutionCountNil(b bool)`

 SetAdditionalInstitutionCountNil sets the value for AdditionalInstitutionCount to be an explicit nil

### UnsetAdditionalInstitutionCount
`func (o *RwdStudiesRow) UnsetAdditionalInstitutionCount()`

UnsetAdditionalInstitutionCount ensures that no value is present for AdditionalInstitutionCount, not even an explicit nil
### GetUnmatchedInstitutionFragments

`func (o *RwdStudiesRow) GetUnmatchedInstitutionFragments() []string`

GetUnmatchedInstitutionFragments returns the UnmatchedInstitutionFragments field if non-nil, zero value otherwise.

### GetUnmatchedInstitutionFragmentsOk

`func (o *RwdStudiesRow) GetUnmatchedInstitutionFragmentsOk() (*[]string, bool)`

GetUnmatchedInstitutionFragmentsOk returns a tuple with the UnmatchedInstitutionFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedInstitutionFragments

`func (o *RwdStudiesRow) SetUnmatchedInstitutionFragments(v []string)`

SetUnmatchedInstitutionFragments sets UnmatchedInstitutionFragments field to given value.

### HasUnmatchedInstitutionFragments

`func (o *RwdStudiesRow) HasUnmatchedInstitutionFragments() bool`

HasUnmatchedInstitutionFragments returns a boolean if a field has been set.

### SetUnmatchedInstitutionFragmentsNil

`func (o *RwdStudiesRow) SetUnmatchedInstitutionFragmentsNil(b bool)`

 SetUnmatchedInstitutionFragmentsNil sets the value for UnmatchedInstitutionFragments to be an explicit nil

### UnsetUnmatchedInstitutionFragments
`func (o *RwdStudiesRow) UnsetUnmatchedInstitutionFragments()`

UnsetUnmatchedInstitutionFragments ensures that no value is present for UnmatchedInstitutionFragments, not even an explicit nil
### GetNetworkNames

`func (o *RwdStudiesRow) GetNetworkNames() []string`

GetNetworkNames returns the NetworkNames field if non-nil, zero value otherwise.

### GetNetworkNamesOk

`func (o *RwdStudiesRow) GetNetworkNamesOk() (*[]string, bool)`

GetNetworkNamesOk returns a tuple with the NetworkNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkNames

`func (o *RwdStudiesRow) SetNetworkNames(v []string)`

SetNetworkNames sets NetworkNames field to given value.

### HasNetworkNames

`func (o *RwdStudiesRow) HasNetworkNames() bool`

HasNetworkNames returns a boolean if a field has been set.

### SetNetworkNamesNil

`func (o *RwdStudiesRow) SetNetworkNamesNil(b bool)`

 SetNetworkNamesNil sets the value for NetworkNames to be an explicit nil

### UnsetNetworkNames
`func (o *RwdStudiesRow) UnsetNetworkNames()`

UnsetNetworkNames ensures that no value is present for NetworkNames, not even an explicit nil
### GetNetworkCount

`func (o *RwdStudiesRow) GetNetworkCount() int32`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *RwdStudiesRow) GetNetworkCountOk() (*int32, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *RwdStudiesRow) SetNetworkCount(v int32)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *RwdStudiesRow) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### SetNetworkCountNil

`func (o *RwdStudiesRow) SetNetworkCountNil(b bool)`

 SetNetworkCountNil sets the value for NetworkCount to be an explicit nil

### UnsetNetworkCount
`func (o *RwdStudiesRow) UnsetNetworkCount()`

UnsetNetworkCount ensures that no value is present for NetworkCount, not even an explicit nil
### GetDataSourceNames

`func (o *RwdStudiesRow) GetDataSourceNames() []string`

GetDataSourceNames returns the DataSourceNames field if non-nil, zero value otherwise.

### GetDataSourceNamesOk

`func (o *RwdStudiesRow) GetDataSourceNamesOk() (*[]string, bool)`

GetDataSourceNamesOk returns a tuple with the DataSourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceNames

`func (o *RwdStudiesRow) SetDataSourceNames(v []string)`

SetDataSourceNames sets DataSourceNames field to given value.

### HasDataSourceNames

`func (o *RwdStudiesRow) HasDataSourceNames() bool`

HasDataSourceNames returns a boolean if a field has been set.

### SetDataSourceNamesNil

`func (o *RwdStudiesRow) SetDataSourceNamesNil(b bool)`

 SetDataSourceNamesNil sets the value for DataSourceNames to be an explicit nil

### UnsetDataSourceNames
`func (o *RwdStudiesRow) UnsetDataSourceNames()`

UnsetDataSourceNames ensures that no value is present for DataSourceNames, not even an explicit nil
### GetDataSourceCount

`func (o *RwdStudiesRow) GetDataSourceCount() int32`

GetDataSourceCount returns the DataSourceCount field if non-nil, zero value otherwise.

### GetDataSourceCountOk

`func (o *RwdStudiesRow) GetDataSourceCountOk() (*int32, bool)`

GetDataSourceCountOk returns a tuple with the DataSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceCount

`func (o *RwdStudiesRow) SetDataSourceCount(v int32)`

SetDataSourceCount sets DataSourceCount field to given value.

### HasDataSourceCount

`func (o *RwdStudiesRow) HasDataSourceCount() bool`

HasDataSourceCount returns a boolean if a field has been set.

### SetDataSourceCountNil

`func (o *RwdStudiesRow) SetDataSourceCountNil(b bool)`

 SetDataSourceCountNil sets the value for DataSourceCount to be an explicit nil

### UnsetDataSourceCount
`func (o *RwdStudiesRow) UnsetDataSourceCount()`

UnsetDataSourceCount ensures that no value is present for DataSourceCount, not even an explicit nil
### GetUnmatchedDataSourceFragments

`func (o *RwdStudiesRow) GetUnmatchedDataSourceFragments() []string`

GetUnmatchedDataSourceFragments returns the UnmatchedDataSourceFragments field if non-nil, zero value otherwise.

### GetUnmatchedDataSourceFragmentsOk

`func (o *RwdStudiesRow) GetUnmatchedDataSourceFragmentsOk() (*[]string, bool)`

GetUnmatchedDataSourceFragmentsOk returns a tuple with the UnmatchedDataSourceFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedDataSourceFragments

`func (o *RwdStudiesRow) SetUnmatchedDataSourceFragments(v []string)`

SetUnmatchedDataSourceFragments sets UnmatchedDataSourceFragments field to given value.

### HasUnmatchedDataSourceFragments

`func (o *RwdStudiesRow) HasUnmatchedDataSourceFragments() bool`

HasUnmatchedDataSourceFragments returns a boolean if a field has been set.

### SetUnmatchedDataSourceFragmentsNil

`func (o *RwdStudiesRow) SetUnmatchedDataSourceFragmentsNil(b bool)`

 SetUnmatchedDataSourceFragmentsNil sets the value for UnmatchedDataSourceFragments to be an explicit nil

### UnsetUnmatchedDataSourceFragments
`func (o *RwdStudiesRow) UnsetUnmatchedDataSourceFragments()`

UnsetUnmatchedDataSourceFragments ensures that no value is present for UnmatchedDataSourceFragments, not even an explicit nil
### GetConductorAttribution

`func (o *RwdStudiesRow) GetConductorAttribution() string`

GetConductorAttribution returns the ConductorAttribution field if non-nil, zero value otherwise.

### GetConductorAttributionOk

`func (o *RwdStudiesRow) GetConductorAttributionOk() (*string, bool)`

GetConductorAttributionOk returns a tuple with the ConductorAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorAttribution

`func (o *RwdStudiesRow) SetConductorAttribution(v string)`

SetConductorAttribution sets ConductorAttribution field to given value.

### HasConductorAttribution

`func (o *RwdStudiesRow) HasConductorAttribution() bool`

HasConductorAttribution returns a boolean if a field has been set.

### SetConductorAttributionNil

`func (o *RwdStudiesRow) SetConductorAttributionNil(b bool)`

 SetConductorAttributionNil sets the value for ConductorAttribution to be an explicit nil

### UnsetConductorAttribution
`func (o *RwdStudiesRow) UnsetConductorAttribution()`

UnsetConductorAttribution ensures that no value is present for ConductorAttribution, not even an explicit nil
### GetStudyCountryNames

`func (o *RwdStudiesRow) GetStudyCountryNames() []string`

GetStudyCountryNames returns the StudyCountryNames field if non-nil, zero value otherwise.

### GetStudyCountryNamesOk

`func (o *RwdStudiesRow) GetStudyCountryNamesOk() (*[]string, bool)`

GetStudyCountryNamesOk returns a tuple with the StudyCountryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyCountryNames

`func (o *RwdStudiesRow) SetStudyCountryNames(v []string)`

SetStudyCountryNames sets StudyCountryNames field to given value.

### HasStudyCountryNames

`func (o *RwdStudiesRow) HasStudyCountryNames() bool`

HasStudyCountryNames returns a boolean if a field has been set.

### SetStudyCountryNamesNil

`func (o *RwdStudiesRow) SetStudyCountryNamesNil(b bool)`

 SetStudyCountryNamesNil sets the value for StudyCountryNames to be an explicit nil

### UnsetStudyCountryNames
`func (o *RwdStudiesRow) UnsetStudyCountryNames()`

UnsetStudyCountryNames ensures that no value is present for StudyCountryNames, not even an explicit nil
### GetStudyCountryCount

`func (o *RwdStudiesRow) GetStudyCountryCount() int32`

GetStudyCountryCount returns the StudyCountryCount field if non-nil, zero value otherwise.

### GetStudyCountryCountOk

`func (o *RwdStudiesRow) GetStudyCountryCountOk() (*int32, bool)`

GetStudyCountryCountOk returns a tuple with the StudyCountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyCountryCount

`func (o *RwdStudiesRow) SetStudyCountryCount(v int32)`

SetStudyCountryCount sets StudyCountryCount field to given value.

### HasStudyCountryCount

`func (o *RwdStudiesRow) HasStudyCountryCount() bool`

HasStudyCountryCount returns a boolean if a field has been set.

### SetStudyCountryCountNil

`func (o *RwdStudiesRow) SetStudyCountryCountNil(b bool)`

 SetStudyCountryCountNil sets the value for StudyCountryCount to be an explicit nil

### UnsetStudyCountryCount
`func (o *RwdStudiesRow) UnsetStudyCountryCount()`

UnsetStudyCountryCount ensures that no value is present for StudyCountryCount, not even an explicit nil
### GetUnmatchedCountryFragments

`func (o *RwdStudiesRow) GetUnmatchedCountryFragments() []string`

GetUnmatchedCountryFragments returns the UnmatchedCountryFragments field if non-nil, zero value otherwise.

### GetUnmatchedCountryFragmentsOk

`func (o *RwdStudiesRow) GetUnmatchedCountryFragmentsOk() (*[]string, bool)`

GetUnmatchedCountryFragmentsOk returns a tuple with the UnmatchedCountryFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedCountryFragments

`func (o *RwdStudiesRow) SetUnmatchedCountryFragments(v []string)`

SetUnmatchedCountryFragments sets UnmatchedCountryFragments field to given value.

### HasUnmatchedCountryFragments

`func (o *RwdStudiesRow) HasUnmatchedCountryFragments() bool`

HasUnmatchedCountryFragments returns a boolean if a field has been set.

### SetUnmatchedCountryFragmentsNil

`func (o *RwdStudiesRow) SetUnmatchedCountryFragmentsNil(b bool)`

 SetUnmatchedCountryFragmentsNil sets the value for UnmatchedCountryFragments to be an explicit nil

### UnsetUnmatchedCountryFragments
`func (o *RwdStudiesRow) UnsetUnmatchedCountryFragments()`

UnsetUnmatchedCountryFragments ensures that no value is present for UnmatchedCountryFragments, not even an explicit nil
### GetFundingContractSignedPlannedOn

`func (o *RwdStudiesRow) GetFundingContractSignedPlannedOn() string`

GetFundingContractSignedPlannedOn returns the FundingContractSignedPlannedOn field if non-nil, zero value otherwise.

### GetFundingContractSignedPlannedOnOk

`func (o *RwdStudiesRow) GetFundingContractSignedPlannedOnOk() (*string, bool)`

GetFundingContractSignedPlannedOnOk returns a tuple with the FundingContractSignedPlannedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingContractSignedPlannedOn

`func (o *RwdStudiesRow) SetFundingContractSignedPlannedOn(v string)`

SetFundingContractSignedPlannedOn sets FundingContractSignedPlannedOn field to given value.

### HasFundingContractSignedPlannedOn

`func (o *RwdStudiesRow) HasFundingContractSignedPlannedOn() bool`

HasFundingContractSignedPlannedOn returns a boolean if a field has been set.

### SetFundingContractSignedPlannedOnNil

`func (o *RwdStudiesRow) SetFundingContractSignedPlannedOnNil(b bool)`

 SetFundingContractSignedPlannedOnNil sets the value for FundingContractSignedPlannedOn to be an explicit nil

### UnsetFundingContractSignedPlannedOn
`func (o *RwdStudiesRow) UnsetFundingContractSignedPlannedOn()`

UnsetFundingContractSignedPlannedOn ensures that no value is present for FundingContractSignedPlannedOn, not even an explicit nil
### GetFundingContractSignedActualOn

`func (o *RwdStudiesRow) GetFundingContractSignedActualOn() string`

GetFundingContractSignedActualOn returns the FundingContractSignedActualOn field if non-nil, zero value otherwise.

### GetFundingContractSignedActualOnOk

`func (o *RwdStudiesRow) GetFundingContractSignedActualOnOk() (*string, bool)`

GetFundingContractSignedActualOnOk returns a tuple with the FundingContractSignedActualOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingContractSignedActualOn

`func (o *RwdStudiesRow) SetFundingContractSignedActualOn(v string)`

SetFundingContractSignedActualOn sets FundingContractSignedActualOn field to given value.

### HasFundingContractSignedActualOn

`func (o *RwdStudiesRow) HasFundingContractSignedActualOn() bool`

HasFundingContractSignedActualOn returns a boolean if a field has been set.

### SetFundingContractSignedActualOnNil

`func (o *RwdStudiesRow) SetFundingContractSignedActualOnNil(b bool)`

 SetFundingContractSignedActualOnNil sets the value for FundingContractSignedActualOn to be an explicit nil

### UnsetFundingContractSignedActualOn
`func (o *RwdStudiesRow) UnsetFundingContractSignedActualOn()`

UnsetFundingContractSignedActualOn ensures that no value is present for FundingContractSignedActualOn, not even an explicit nil
### GetStudyStartPlannedOn

`func (o *RwdStudiesRow) GetStudyStartPlannedOn() string`

GetStudyStartPlannedOn returns the StudyStartPlannedOn field if non-nil, zero value otherwise.

### GetStudyStartPlannedOnOk

`func (o *RwdStudiesRow) GetStudyStartPlannedOnOk() (*string, bool)`

GetStudyStartPlannedOnOk returns a tuple with the StudyStartPlannedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStartPlannedOn

`func (o *RwdStudiesRow) SetStudyStartPlannedOn(v string)`

SetStudyStartPlannedOn sets StudyStartPlannedOn field to given value.

### HasStudyStartPlannedOn

`func (o *RwdStudiesRow) HasStudyStartPlannedOn() bool`

HasStudyStartPlannedOn returns a boolean if a field has been set.

### SetStudyStartPlannedOnNil

`func (o *RwdStudiesRow) SetStudyStartPlannedOnNil(b bool)`

 SetStudyStartPlannedOnNil sets the value for StudyStartPlannedOn to be an explicit nil

### UnsetStudyStartPlannedOn
`func (o *RwdStudiesRow) UnsetStudyStartPlannedOn()`

UnsetStudyStartPlannedOn ensures that no value is present for StudyStartPlannedOn, not even an explicit nil
### GetStudyStartActualOn

`func (o *RwdStudiesRow) GetStudyStartActualOn() string`

GetStudyStartActualOn returns the StudyStartActualOn field if non-nil, zero value otherwise.

### GetStudyStartActualOnOk

`func (o *RwdStudiesRow) GetStudyStartActualOnOk() (*string, bool)`

GetStudyStartActualOnOk returns a tuple with the StudyStartActualOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStartActualOn

`func (o *RwdStudiesRow) SetStudyStartActualOn(v string)`

SetStudyStartActualOn sets StudyStartActualOn field to given value.

### HasStudyStartActualOn

`func (o *RwdStudiesRow) HasStudyStartActualOn() bool`

HasStudyStartActualOn returns a boolean if a field has been set.

### SetStudyStartActualOnNil

`func (o *RwdStudiesRow) SetStudyStartActualOnNil(b bool)`

 SetStudyStartActualOnNil sets the value for StudyStartActualOn to be an explicit nil

### UnsetStudyStartActualOn
`func (o *RwdStudiesRow) UnsetStudyStartActualOn()`

UnsetStudyStartActualOn ensures that no value is present for StudyStartActualOn, not even an explicit nil
### GetDataAnalysisStartPlannedOn

`func (o *RwdStudiesRow) GetDataAnalysisStartPlannedOn() string`

GetDataAnalysisStartPlannedOn returns the DataAnalysisStartPlannedOn field if non-nil, zero value otherwise.

### GetDataAnalysisStartPlannedOnOk

`func (o *RwdStudiesRow) GetDataAnalysisStartPlannedOnOk() (*string, bool)`

GetDataAnalysisStartPlannedOnOk returns a tuple with the DataAnalysisStartPlannedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAnalysisStartPlannedOn

`func (o *RwdStudiesRow) SetDataAnalysisStartPlannedOn(v string)`

SetDataAnalysisStartPlannedOn sets DataAnalysisStartPlannedOn field to given value.

### HasDataAnalysisStartPlannedOn

`func (o *RwdStudiesRow) HasDataAnalysisStartPlannedOn() bool`

HasDataAnalysisStartPlannedOn returns a boolean if a field has been set.

### SetDataAnalysisStartPlannedOnNil

`func (o *RwdStudiesRow) SetDataAnalysisStartPlannedOnNil(b bool)`

 SetDataAnalysisStartPlannedOnNil sets the value for DataAnalysisStartPlannedOn to be an explicit nil

### UnsetDataAnalysisStartPlannedOn
`func (o *RwdStudiesRow) UnsetDataAnalysisStartPlannedOn()`

UnsetDataAnalysisStartPlannedOn ensures that no value is present for DataAnalysisStartPlannedOn, not even an explicit nil
### GetDataAnalysisStartActualOn

`func (o *RwdStudiesRow) GetDataAnalysisStartActualOn() string`

GetDataAnalysisStartActualOn returns the DataAnalysisStartActualOn field if non-nil, zero value otherwise.

### GetDataAnalysisStartActualOnOk

`func (o *RwdStudiesRow) GetDataAnalysisStartActualOnOk() (*string, bool)`

GetDataAnalysisStartActualOnOk returns a tuple with the DataAnalysisStartActualOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAnalysisStartActualOn

`func (o *RwdStudiesRow) SetDataAnalysisStartActualOn(v string)`

SetDataAnalysisStartActualOn sets DataAnalysisStartActualOn field to given value.

### HasDataAnalysisStartActualOn

`func (o *RwdStudiesRow) HasDataAnalysisStartActualOn() bool`

HasDataAnalysisStartActualOn returns a boolean if a field has been set.

### SetDataAnalysisStartActualOnNil

`func (o *RwdStudiesRow) SetDataAnalysisStartActualOnNil(b bool)`

 SetDataAnalysisStartActualOnNil sets the value for DataAnalysisStartActualOn to be an explicit nil

### UnsetDataAnalysisStartActualOn
`func (o *RwdStudiesRow) UnsetDataAnalysisStartActualOn()`

UnsetDataAnalysisStartActualOn ensures that no value is present for DataAnalysisStartActualOn, not even an explicit nil
### GetInterimReportPlannedOn

`func (o *RwdStudiesRow) GetInterimReportPlannedOn() string`

GetInterimReportPlannedOn returns the InterimReportPlannedOn field if non-nil, zero value otherwise.

### GetInterimReportPlannedOnOk

`func (o *RwdStudiesRow) GetInterimReportPlannedOnOk() (*string, bool)`

GetInterimReportPlannedOnOk returns a tuple with the InterimReportPlannedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimReportPlannedOn

`func (o *RwdStudiesRow) SetInterimReportPlannedOn(v string)`

SetInterimReportPlannedOn sets InterimReportPlannedOn field to given value.

### HasInterimReportPlannedOn

`func (o *RwdStudiesRow) HasInterimReportPlannedOn() bool`

HasInterimReportPlannedOn returns a boolean if a field has been set.

### SetInterimReportPlannedOnNil

`func (o *RwdStudiesRow) SetInterimReportPlannedOnNil(b bool)`

 SetInterimReportPlannedOnNil sets the value for InterimReportPlannedOn to be an explicit nil

### UnsetInterimReportPlannedOn
`func (o *RwdStudiesRow) UnsetInterimReportPlannedOn()`

UnsetInterimReportPlannedOn ensures that no value is present for InterimReportPlannedOn, not even an explicit nil
### GetInterimReportActualOn

`func (o *RwdStudiesRow) GetInterimReportActualOn() string`

GetInterimReportActualOn returns the InterimReportActualOn field if non-nil, zero value otherwise.

### GetInterimReportActualOnOk

`func (o *RwdStudiesRow) GetInterimReportActualOnOk() (*string, bool)`

GetInterimReportActualOnOk returns a tuple with the InterimReportActualOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimReportActualOn

`func (o *RwdStudiesRow) SetInterimReportActualOn(v string)`

SetInterimReportActualOn sets InterimReportActualOn field to given value.

### HasInterimReportActualOn

`func (o *RwdStudiesRow) HasInterimReportActualOn() bool`

HasInterimReportActualOn returns a boolean if a field has been set.

### SetInterimReportActualOnNil

`func (o *RwdStudiesRow) SetInterimReportActualOnNil(b bool)`

 SetInterimReportActualOnNil sets the value for InterimReportActualOn to be an explicit nil

### UnsetInterimReportActualOn
`func (o *RwdStudiesRow) UnsetInterimReportActualOn()`

UnsetInterimReportActualOn ensures that no value is present for InterimReportActualOn, not even an explicit nil
### GetFinalReportPlannedOn

`func (o *RwdStudiesRow) GetFinalReportPlannedOn() string`

GetFinalReportPlannedOn returns the FinalReportPlannedOn field if non-nil, zero value otherwise.

### GetFinalReportPlannedOnOk

`func (o *RwdStudiesRow) GetFinalReportPlannedOnOk() (*string, bool)`

GetFinalReportPlannedOnOk returns a tuple with the FinalReportPlannedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportPlannedOn

`func (o *RwdStudiesRow) SetFinalReportPlannedOn(v string)`

SetFinalReportPlannedOn sets FinalReportPlannedOn field to given value.

### HasFinalReportPlannedOn

`func (o *RwdStudiesRow) HasFinalReportPlannedOn() bool`

HasFinalReportPlannedOn returns a boolean if a field has been set.

### SetFinalReportPlannedOnNil

`func (o *RwdStudiesRow) SetFinalReportPlannedOnNil(b bool)`

 SetFinalReportPlannedOnNil sets the value for FinalReportPlannedOn to be an explicit nil

### UnsetFinalReportPlannedOn
`func (o *RwdStudiesRow) UnsetFinalReportPlannedOn()`

UnsetFinalReportPlannedOn ensures that no value is present for FinalReportPlannedOn, not even an explicit nil
### GetFinalReportActualOn

`func (o *RwdStudiesRow) GetFinalReportActualOn() string`

GetFinalReportActualOn returns the FinalReportActualOn field if non-nil, zero value otherwise.

### GetFinalReportActualOnOk

`func (o *RwdStudiesRow) GetFinalReportActualOnOk() (*string, bool)`

GetFinalReportActualOnOk returns a tuple with the FinalReportActualOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportActualOn

`func (o *RwdStudiesRow) SetFinalReportActualOn(v string)`

SetFinalReportActualOn sets FinalReportActualOn field to given value.

### HasFinalReportActualOn

`func (o *RwdStudiesRow) HasFinalReportActualOn() bool`

HasFinalReportActualOn returns a boolean if a field has been set.

### SetFinalReportActualOnNil

`func (o *RwdStudiesRow) SetFinalReportActualOnNil(b bool)`

 SetFinalReportActualOnNil sets the value for FinalReportActualOn to be an explicit nil

### UnsetFinalReportActualOn
`func (o *RwdStudiesRow) UnsetFinalReportActualOn()`

UnsetFinalReportActualOn ensures that no value is present for FinalReportActualOn, not even an explicit nil
### GetHasFinalReportPair

`func (o *RwdStudiesRow) GetHasFinalReportPair() bool`

GetHasFinalReportPair returns the HasFinalReportPair field if non-nil, zero value otherwise.

### GetHasFinalReportPairOk

`func (o *RwdStudiesRow) GetHasFinalReportPairOk() (*bool, bool)`

GetHasFinalReportPairOk returns a tuple with the HasFinalReportPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFinalReportPair

`func (o *RwdStudiesRow) SetHasFinalReportPair(v bool)`

SetHasFinalReportPair sets HasFinalReportPair field to given value.

### HasHasFinalReportPair

`func (o *RwdStudiesRow) HasHasFinalReportPair() bool`

HasHasFinalReportPair returns a boolean if a field has been set.

### SetHasFinalReportPairNil

`func (o *RwdStudiesRow) SetHasFinalReportPairNil(b bool)`

 SetHasFinalReportPairNil sets the value for HasFinalReportPair to be an explicit nil

### UnsetHasFinalReportPair
`func (o *RwdStudiesRow) UnsetHasFinalReportPair()`

UnsetHasFinalReportPair ensures that no value is present for HasFinalReportPair, not even an explicit nil
### GetFinalReportIsBackfilled

`func (o *RwdStudiesRow) GetFinalReportIsBackfilled() bool`

GetFinalReportIsBackfilled returns the FinalReportIsBackfilled field if non-nil, zero value otherwise.

### GetFinalReportIsBackfilledOk

`func (o *RwdStudiesRow) GetFinalReportIsBackfilledOk() (*bool, bool)`

GetFinalReportIsBackfilledOk returns a tuple with the FinalReportIsBackfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportIsBackfilled

`func (o *RwdStudiesRow) SetFinalReportIsBackfilled(v bool)`

SetFinalReportIsBackfilled sets FinalReportIsBackfilled field to given value.

### HasFinalReportIsBackfilled

`func (o *RwdStudiesRow) HasFinalReportIsBackfilled() bool`

HasFinalReportIsBackfilled returns a boolean if a field has been set.

### SetFinalReportIsBackfilledNil

`func (o *RwdStudiesRow) SetFinalReportIsBackfilledNil(b bool)`

 SetFinalReportIsBackfilledNil sets the value for FinalReportIsBackfilled to be an explicit nil

### UnsetFinalReportIsBackfilled
`func (o *RwdStudiesRow) UnsetFinalReportIsBackfilled()`

UnsetFinalReportIsBackfilled ensures that no value is present for FinalReportIsBackfilled, not even an explicit nil
### GetFinalReportSlipDays

`func (o *RwdStudiesRow) GetFinalReportSlipDays() int32`

GetFinalReportSlipDays returns the FinalReportSlipDays field if non-nil, zero value otherwise.

### GetFinalReportSlipDaysOk

`func (o *RwdStudiesRow) GetFinalReportSlipDaysOk() (*int32, bool)`

GetFinalReportSlipDaysOk returns a tuple with the FinalReportSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportSlipDays

`func (o *RwdStudiesRow) SetFinalReportSlipDays(v int32)`

SetFinalReportSlipDays sets FinalReportSlipDays field to given value.

### HasFinalReportSlipDays

`func (o *RwdStudiesRow) HasFinalReportSlipDays() bool`

HasFinalReportSlipDays returns a boolean if a field has been set.

### SetFinalReportSlipDaysNil

`func (o *RwdStudiesRow) SetFinalReportSlipDaysNil(b bool)`

 SetFinalReportSlipDaysNil sets the value for FinalReportSlipDays to be an explicit nil

### UnsetFinalReportSlipDays
`func (o *RwdStudiesRow) UnsetFinalReportSlipDays()`

UnsetFinalReportSlipDays ensures that no value is present for FinalReportSlipDays, not even an explicit nil
### GetStudyStartYear

`func (o *RwdStudiesRow) GetStudyStartYear() int32`

GetStudyStartYear returns the StudyStartYear field if non-nil, zero value otherwise.

### GetStudyStartYearOk

`func (o *RwdStudiesRow) GetStudyStartYearOk() (*int32, bool)`

GetStudyStartYearOk returns a tuple with the StudyStartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStartYear

`func (o *RwdStudiesRow) SetStudyStartYear(v int32)`

SetStudyStartYear sets StudyStartYear field to given value.

### HasStudyStartYear

`func (o *RwdStudiesRow) HasStudyStartYear() bool`

HasStudyStartYear returns a boolean if a field has been set.

### SetStudyStartYearNil

`func (o *RwdStudiesRow) SetStudyStartYearNil(b bool)`

 SetStudyStartYearNil sets the value for StudyStartYear to be an explicit nil

### UnsetStudyStartYear
`func (o *RwdStudiesRow) UnsetStudyStartYear()`

UnsetStudyStartYear ensures that no value is present for StudyStartYear, not even an explicit nil
### GetStartedInLast5Years

`func (o *RwdStudiesRow) GetStartedInLast5Years() bool`

GetStartedInLast5Years returns the StartedInLast5Years field if non-nil, zero value otherwise.

### GetStartedInLast5YearsOk

`func (o *RwdStudiesRow) GetStartedInLast5YearsOk() (*bool, bool)`

GetStartedInLast5YearsOk returns a tuple with the StartedInLast5Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedInLast5Years

`func (o *RwdStudiesRow) SetStartedInLast5Years(v bool)`

SetStartedInLast5Years sets StartedInLast5Years field to given value.

### HasStartedInLast5Years

`func (o *RwdStudiesRow) HasStartedInLast5Years() bool`

HasStartedInLast5Years returns a boolean if a field has been set.

### SetStartedInLast5YearsNil

`func (o *RwdStudiesRow) SetStartedInLast5YearsNil(b bool)`

 SetStartedInLast5YearsNil sets the value for StartedInLast5Years to be an explicit nil

### UnsetStartedInLast5Years
`func (o *RwdStudiesRow) UnsetStartedInLast5Years()`

UnsetStartedInLast5Years ensures that no value is present for StartedInLast5Years, not even an explicit nil
### GetRecordUpdatedAt

`func (o *RwdStudiesRow) GetRecordUpdatedAt() time.Time`

GetRecordUpdatedAt returns the RecordUpdatedAt field if non-nil, zero value otherwise.

### GetRecordUpdatedAtOk

`func (o *RwdStudiesRow) GetRecordUpdatedAtOk() (*time.Time, bool)`

GetRecordUpdatedAtOk returns a tuple with the RecordUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedAt

`func (o *RwdStudiesRow) SetRecordUpdatedAt(v time.Time)`

SetRecordUpdatedAt sets RecordUpdatedAt field to given value.

### HasRecordUpdatedAt

`func (o *RwdStudiesRow) HasRecordUpdatedAt() bool`

HasRecordUpdatedAt returns a boolean if a field has been set.

### SetRecordUpdatedAtNil

`func (o *RwdStudiesRow) SetRecordUpdatedAtNil(b bool)`

 SetRecordUpdatedAtNil sets the value for RecordUpdatedAt to be an explicit nil

### UnsetRecordUpdatedAt
`func (o *RwdStudiesRow) UnsetRecordUpdatedAt()`

UnsetRecordUpdatedAt ensures that no value is present for RecordUpdatedAt, not even an explicit nil
### GetRecordFirstPublishedAt

`func (o *RwdStudiesRow) GetRecordFirstPublishedAt() time.Time`

GetRecordFirstPublishedAt returns the RecordFirstPublishedAt field if non-nil, zero value otherwise.

### GetRecordFirstPublishedAtOk

`func (o *RwdStudiesRow) GetRecordFirstPublishedAtOk() (*time.Time, bool)`

GetRecordFirstPublishedAtOk returns a tuple with the RecordFirstPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFirstPublishedAt

`func (o *RwdStudiesRow) SetRecordFirstPublishedAt(v time.Time)`

SetRecordFirstPublishedAt sets RecordFirstPublishedAt field to given value.

### HasRecordFirstPublishedAt

`func (o *RwdStudiesRow) HasRecordFirstPublishedAt() bool`

HasRecordFirstPublishedAt returns a boolean if a field has been set.

### SetRecordFirstPublishedAtNil

`func (o *RwdStudiesRow) SetRecordFirstPublishedAtNil(b bool)`

 SetRecordFirstPublishedAtNil sets the value for RecordFirstPublishedAt to be an explicit nil

### UnsetRecordFirstPublishedAt
`func (o *RwdStudiesRow) UnsetRecordFirstPublishedAt()`

UnsetRecordFirstPublishedAt ensures that no value is present for RecordFirstPublishedAt, not even an explicit nil
### GetRegisterSourceCount

`func (o *RwdStudiesRow) GetRegisterSourceCount() int32`

GetRegisterSourceCount returns the RegisterSourceCount field if non-nil, zero value otherwise.

### GetRegisterSourceCountOk

`func (o *RwdStudiesRow) GetRegisterSourceCountOk() (*int32, bool)`

GetRegisterSourceCountOk returns a tuple with the RegisterSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterSourceCount

`func (o *RwdStudiesRow) SetRegisterSourceCount(v int32)`

SetRegisterSourceCount sets RegisterSourceCount field to given value.

### HasRegisterSourceCount

`func (o *RwdStudiesRow) HasRegisterSourceCount() bool`

HasRegisterSourceCount returns a boolean if a field has been set.

### SetRegisterSourceCountNil

`func (o *RwdStudiesRow) SetRegisterSourceCountNil(b bool)`

 SetRegisterSourceCountNil sets the value for RegisterSourceCount to be an explicit nil

### UnsetRegisterSourceCount
`func (o *RwdStudiesRow) UnsetRegisterSourceCount()`

UnsetRegisterSourceCount ensures that no value is present for RegisterSourceCount, not even an explicit nil
### GetStudyRegisterTotal

`func (o *RwdStudiesRow) GetStudyRegisterTotal() int32`

GetStudyRegisterTotal returns the StudyRegisterTotal field if non-nil, zero value otherwise.

### GetStudyRegisterTotalOk

`func (o *RwdStudiesRow) GetStudyRegisterTotalOk() (*int32, bool)`

GetStudyRegisterTotalOk returns a tuple with the StudyRegisterTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyRegisterTotal

`func (o *RwdStudiesRow) SetStudyRegisterTotal(v int32)`

SetStudyRegisterTotal sets StudyRegisterTotal field to given value.

### HasStudyRegisterTotal

`func (o *RwdStudiesRow) HasStudyRegisterTotal() bool`

HasStudyRegisterTotal returns a boolean if a field has been set.

### SetStudyRegisterTotalNil

`func (o *RwdStudiesRow) SetStudyRegisterTotalNil(b bool)`

 SetStudyRegisterTotalNil sets the value for StudyRegisterTotal to be an explicit nil

### UnsetStudyRegisterTotal
`func (o *RwdStudiesRow) UnsetStudyRegisterTotal()`

UnsetStudyRegisterTotal ensures that no value is present for StudyRegisterTotal, not even an explicit nil
### GetStudiesNamingARegisteredSource

`func (o *RwdStudiesRow) GetStudiesNamingARegisteredSource() int32`

GetStudiesNamingARegisteredSource returns the StudiesNamingARegisteredSource field if non-nil, zero value otherwise.

### GetStudiesNamingARegisteredSourceOk

`func (o *RwdStudiesRow) GetStudiesNamingARegisteredSourceOk() (*int32, bool)`

GetStudiesNamingARegisteredSourceOk returns a tuple with the StudiesNamingARegisteredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudiesNamingARegisteredSource

`func (o *RwdStudiesRow) SetStudiesNamingARegisteredSource(v int32)`

SetStudiesNamingARegisteredSource sets StudiesNamingARegisteredSource field to given value.

### HasStudiesNamingARegisteredSource

`func (o *RwdStudiesRow) HasStudiesNamingARegisteredSource() bool`

HasStudiesNamingARegisteredSource returns a boolean if a field has been set.

### SetStudiesNamingARegisteredSourceNil

`func (o *RwdStudiesRow) SetStudiesNamingARegisteredSourceNil(b bool)`

 SetStudiesNamingARegisteredSourceNil sets the value for StudiesNamingARegisteredSource to be an explicit nil

### UnsetStudiesNamingARegisteredSource
`func (o *RwdStudiesRow) UnsetStudiesNamingARegisteredSource()`

UnsetStudiesNamingARegisteredSource ensures that no value is present for StudiesNamingARegisteredSource, not even an explicit nil
### GetStudiesNamingARegisteredConductor

`func (o *RwdStudiesRow) GetStudiesNamingARegisteredConductor() int32`

GetStudiesNamingARegisteredConductor returns the StudiesNamingARegisteredConductor field if non-nil, zero value otherwise.

### GetStudiesNamingARegisteredConductorOk

`func (o *RwdStudiesRow) GetStudiesNamingARegisteredConductorOk() (*int32, bool)`

GetStudiesNamingARegisteredConductorOk returns a tuple with the StudiesNamingARegisteredConductor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudiesNamingARegisteredConductor

`func (o *RwdStudiesRow) SetStudiesNamingARegisteredConductor(v int32)`

SetStudiesNamingARegisteredConductor sets StudiesNamingARegisteredConductor field to given value.

### HasStudiesNamingARegisteredConductor

`func (o *RwdStudiesRow) HasStudiesNamingARegisteredConductor() bool`

HasStudiesNamingARegisteredConductor returns a boolean if a field has been set.

### SetStudiesNamingARegisteredConductorNil

`func (o *RwdStudiesRow) SetStudiesNamingARegisteredConductorNil(b bool)`

 SetStudiesNamingARegisteredConductorNil sets the value for StudiesNamingARegisteredConductor to be an explicit nil

### UnsetStudiesNamingARegisteredConductor
`func (o *RwdStudiesRow) UnsetStudiesNamingARegisteredConductor()`

UnsetStudiesNamingARegisteredConductor ensures that no value is present for StudiesNamingARegisteredConductor, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdStudiesRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdStudiesRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdStudiesRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdStudiesRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdStudiesRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdStudiesRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdStudiesRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdStudiesRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdStudiesRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdStudiesRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdStudiesRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdStudiesRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdStudiesRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdStudiesRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdStudiesRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdStudiesRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdStudiesRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdStudiesRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdStudiesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdStudiesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdStudiesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdStudiesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdStudiesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdStudiesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


