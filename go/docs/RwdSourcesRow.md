# RwdSourcesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | Pointer to **NullableString** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] 
**DataSourceName** | Pointer to **NullableString** | The data source&#39;s name as the data holder registered it; null on a stub entry with no name. | [optional] 
**DataSourceAcronym** | Pointer to **NullableString** | The acronym the data holder registered for the source. | [optional] 
**DataHolder** | Pointer to **NullableString** | The organisation that holds the data, as typed in the register (free text, not resolved to a company). | [optional] 
**DataSourceType** | Pointer to **NullableString** | The data source type(s) the holder ticked, as one comma-separated string as the register exports it; rwd_source_types has one row per declared type. | [optional] 
**DataSourceTypeOther** | Pointer to **NullableString** | The holder&#39;s free-text description when it chose the register&#39;s other data source type. | [optional] 
**DataSourceDomain** | Pointer to **NullableString** | The register&#39;s data source domain (for example Claims, Primary care, Secondary care). | [optional] 
**CareSetting** | Pointer to **NullableString** | The care setting the holder declares for the source. | [optional] 
**MainFinancialSupport** | Pointer to **NullableString** | The main financial support of the data source as the holder declared it. | [optional] 
**CountryNames** | Pointer to **[]string** | The countries the source covers, one array element per register country, split by the register&#39;s own country list (so Iran, Islamic Republic of stays one country). Matches the group-by-able membership filter &#39;Spain&#39; &#x3D; ANY(country_names). | [optional] 
**UnmatchedCountryFragments** | Pointer to **[]string** | Country text items the country vocabulary of this build does not hold. | [optional] 
**CountryCount** | Pointer to **NullableInt32** | The number of countries in country_names. | [optional] 
**CatalogueUrl** | Pointer to **NullableString** | The address of this entry on the HMA-EMA catalogue site (the register&#39;s URL LINK). | [optional] 
**DataSourceWebsite** | Pointer to **NullableString** | The data source&#39;s own website as the holder registered it. | [optional] 
**IsDarwinEuDataPartner** | Pointer to **NullableBool** | Whether the holder declares the source a DARWIN EU data partner (the EMA-coordinated European network of real-world data sources); null when unanswered. | [optional] 
**IsQualified** | Pointer to **NullableBool** | The register&#39;s Data source qualification answer; true when the holder declares a qualification, null when unanswered. | [optional] 
**RefreshCadence** | Pointer to **NullableString** | How often the holder says the database takes new data, as typed (the register&#39;s Data source refresh). | [optional] 
**LastRefreshedOn** | Pointer to **NullableString** | The date the holder says the database last took new data (the register&#39;s Data source last refresh). | [optional] 
**RecordUpdatedAt** | Pointer to **NullableTime** | When the registrant last edited this catalogue entry. A different clock from last_refreshed_on: it dates the web form, not the database. | [optional] 
**RecordFirstPublishedAt** | Pointer to **NullableTime** | When the catalogue entry was first published. | [optional] 
**IsOutdated** | Pointer to **NullableBool** | EMA&#39;s own Outdated flag, carried through untouched; the register flags none of the data sources, so read false as EMA&#39;s answer and not as a currency finding. | [optional] 
**RecordUpdatedOver12MonthsAgo** | Pointer to **NullableBool** | True when the catalogue entry was last edited more than twelve months before the snapshot. About the web form only; it says nothing about whether the database has taken new data (see database_refresh_state). | [optional] 
**RefreshIntervalDays** | Pointer to **NullableInt32** | The interval in days the longest declared refresh cadence is read as; the threshold the database_refresh_state verdict used. Null when no declared cadence was recognised. | [optional] 
**RefreshIntervalBasis** | Pointer to **NullableString** | Which declared cadence refresh_interval_days came from. | [optional] 
**RefreshIntervalDaysShortest** | Pointer to **NullableInt32** | The interval in days of the strictest declared cadence, for a source that declared several. | [optional] 
**RefreshIntervalBasisShortest** | Pointer to **NullableString** | Which declared cadence refresh_interval_days_shortest came from. | [optional] 
**RefreshCadenceDeclared** | Pointer to **[]string** | Every cadence the holder declared, in declared order; the month names of one answer count as one declaration. Positionally paired with refresh_cadence_interval_days and refresh_cadence_is_month_names. Empty when none was declared. | [optional] 
**RefreshCadenceIntervalDays** | Pointer to **[]string** | The interval in days each declared cadence was read as, positionally paired with refresh_cadence_declared, rendered as text because the catalog has only text arrays. | [optional] 
**RefreshCadenceIsMonthNames** | Pointer to **[]string** | Whether each declared cadence is a list of month names (priced as 365 days over the number of months named); rendered as text (true or false per element) because the catalog has only text arrays. | [optional] 
**DaysSinceLastRefresh** | Pointer to **NullableInt32** | Days from last_refreshed_on to the catalogue snapshot date; null when no refresh date is declared. | [optional] 
**DatabaseRefreshState** | Pointer to **NullableString** | Adherence verdict about the holder&#39;s own promise: the declared cadence against the declared last refresh, measured to the snapshot date. Not a currency answer: use refresh_recency_band for how old the data is. | [optional] 
**RefreshStateStrict** | Pointer to **NullableString** | The same adherence verdict as database_refresh_state, read against the strictest declared cadence instead of the longest. | [optional] 
**RefreshRecencyBand** | Pointer to **NullableString** | Absolute age of the last database refresh at the snapshot date, in disjoint bands, with no reference to any declared cadence. Sources that declared no refresh date have their own band; they are unanswered, not stale. | [optional] 
**CollectionLastBand** | Pointer to **NullableString** | The band of the declared last collection date (the observable window), banded with the same ruler as refresh_recency_band; null when no collection date is declared. | [optional] 
**CollectionContradictsRefreshBand** | Pointer to **NullableBool** | True when the last collection date and the last refresh date fall in different currency bands. A fact to print beside the currency screen, not a control. | [optional] 
**PrescriptionsOfMedicines** | Pointer to **NullableBool** | Whether the data holder declares that prescriptions of medicines are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**PrescriptionsVocabulary** | Pointer to **NullableString** | The coding scheme(s) the holder declares for prescriptions, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] 
**PrescriptionsVocabularyOther** | Pointer to **NullableString** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] 
**DispensingOfMedicines** | Pointer to **NullableBool** | Whether the data holder declares that dispensing of medicines is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**DispensingVocabulary** | Pointer to **NullableString** | The coding scheme(s) the holder declares for dispensing, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] 
**DispensingVocabularyOther** | Pointer to **NullableString** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] 
**MedicinalProductVocabulary** | Pointer to **NullableString** | The coding scheme(s) the holder declares for medicinal product information, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] 
**MedicinalProductVocabularyOther** | Pointer to **NullableString** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] 
**IndicationForUse** | Pointer to **NullableBool** | Whether the data holder declares that the indication for use of a medicine is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**IndicationVocabulary** | Pointer to **NullableString** | The coding scheme(s) the holder declares for the indication for use, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] 
**IndicationVocabularyOther** | Pointer to **NullableString** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] 
**DiagnosticCodes** | Pointer to **NullableBool** | Whether the data holder declares that diagnostic codes are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**DiagnosisMedicalEventVocabulary** | Pointer to **NullableString** | The coding scheme(s) the holder declares for diagnoses and medical events, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] 
**DiagnosisMedicalEventVocabularyOther** | Pointer to **NullableString** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] 
**HospitalAdmissionAndOrDischarge** | Pointer to **NullableBool** | Whether the data holder declares that hospital admission and/or discharge is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**IcuAdmission** | Pointer to **NullableBool** | Whether the data holder declares that intensive care unit (ICU) admission is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**CauseOfDeath** | Pointer to **NullableBool** | Whether the data holder declares that cause of death is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**CauseOfDeathVocabulary** | Pointer to **NullableString** | The coding scheme(s) the holder declares for cause of death, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] 
**CauseOfDeathVocabularyOther** | Pointer to **NullableString** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] 
**DiseaseInformation** | Pointer to **NullableBool** | Whether the data holder declares that disease information is collected: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**RareDiseases** | Pointer to **NullableBool** | Whether the data holder declares that rare diseases are identifiable: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**PregnancyAndOrNeonates** | Pointer to **NullableBool** | Whether the data holder declares that pregnancy and/or neonates are covered: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**AdvanceTherapyMedicinalProductsAtmp** | Pointer to **NullableBool** | Whether the data holder declares that advanced therapy medicinal products (ATMP) are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**Contraception** | Pointer to **NullableBool** | Whether the data holder declares that contraception is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**MedicalDevices** | Pointer to **NullableBool** | Whether the data holder declares that medical devices are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**AdministrationOfVaccines** | Pointer to **NullableBool** | Whether the data holder declares that administration of vaccines is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**Procedures** | Pointer to **NullableBool** | Whether the data holder declares that procedures are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**ProceduresVocabulary** | Pointer to **NullableString** | The coding scheme(s) the holder declares for procedures, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] 
**ProceduresVocabularyOther** | Pointer to **NullableString** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] 
**ClinicalMeasurements** | Pointer to **NullableBool** | Whether the data holder declares that clinical measurements are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**HealthcareProvider** | Pointer to **NullableBool** | Whether the data holder declares that the healthcare provider is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**GeneticData** | Pointer to **NullableBool** | Whether the data holder declares that genetic data are held: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**GeneticDataVocabulary** | Pointer to **NullableString** | The coding scheme(s) the holder declares for genetic data, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] 
**GeneticDataVocabularyOther** | Pointer to **NullableString** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] 
**BiomarkerData** | Pointer to **NullableBool** | Whether the data holder declares that biomarker data are held: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**BiomarkerDataVocabulary** | Pointer to **NullableString** | The coding scheme(s) the holder declares for biomarker data, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] 
**BiomarkerVocabularyOther** | Pointer to **NullableString** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] 
**PatientReportedOutcomes** | Pointer to **NullableBool** | Whether the data holder declares that patient-reported outcomes are held: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**PatientGeneratedData** | Pointer to **NullableBool** | Whether the data holder declares that patient-generated data are held: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**UnitsOfHealthcareUtilisation** | Pointer to **NullableBool** | Whether the data holder declares that units of healthcare utilisation are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**UniqueIdentifierForPersons** | Pointer to **NullableBool** | Whether the data holder declares that a unique identifier for persons exists: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**QualityOfLifeMeasurements** | Pointer to **NullableBool** | Whether the data holder declares that quality of life measurements are held: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**QualityOfLifeMeasurementsVocabulary** | Pointer to **NullableString** | The coding scheme(s) the holder declares for quality of life measurements, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] 
**QualityOfLifeMeasurementsOther** | Pointer to **NullableString** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] 
**LifestyleFactors** | Pointer to **NullableBool** | Whether the data holder declares that lifestyle factors are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**SociodemographicInformation** | Pointer to **NullableBool** | Whether the data holder declares that sociodemographic information is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**PopulationSize** | Pointer to **NullableInt64** | The population size the holder declared for the source (all registered persons). | [optional] 
**ActivePopulationSize** | Pointer to **NullableInt64** | The active population size the holder declared for the source. | [optional] 
**MedianObservationTimeYears** | Pointer to **NullableFloat32** | The median observation time in years exactly as the holder declared it. Never corrected: see median_observation_time_plausibility. | [optional] 
**MedianObservationTimeActiveYears** | Pointer to **NullableFloat32** | The median observation time in years for the active population exactly as the holder declared it. Never corrected: see median_observation_time_active_plausibility. | [optional] 
**EstablishedOn** | Pointer to **NullableString** | The date the data source was established. | [optional] 
**CollectionFirstOn** | Pointer to **NullableString** | The date of the first data collection. | [optional] 
**CollectionLastOn** | Pointer to **NullableString** | The date of the last data collection; null on most sources and a future date on some, so it is a fact to read and not a field to sort on. | [optional] 
**DeclaredHistorySpanYears** | Pointer to **NullableFloat32** | The data window in years the holder&#39;s own dates give the database (earliest of established and first collection, to last collection or the snapshot). | [optional] 
**MedianObservationTimePlausibility** | Pointer to **NullableString** | Classifies the declared median observation time without correcting it: plausible, exceeds_ceiling (more than 100 years), exceeds_declared_span (more than a year longer than the holder&#39;s own data window) or not_declared. | [optional] 
**MedianObservationTimeActivePlausibility** | Pointer to **NullableString** | The same classification for the active-population median observation time. | [optional] 
**EstimatedPopulationCoverageText** | Pointer to **NullableString** | The estimated percentage of the population covered, as free text exactly as typed (values include Unsure and Not available); deliberately not parsed into a number. | [optional] 
**AgeGroups** | Pointer to **NullableString** | The age groups the holder declares the source covers, as one comma-separated string. | [optional] 
**AgeBandCount** | Pointer to **NullableInt32** | The number of declared age bands (rwd_source_age_bands has one row per band). | [optional] 
**HasPopulationBreakdown** | Pointer to **NullableBool** | Whether the holder published a population size breakdown by age band at all. | [optional] 
**HasActivePopulationBreakdown** | Pointer to **NullableBool** | Whether the holder published an active population breakdown by age band at all. | [optional] 
**AgeBandsAreNested** | Pointer to **NullableBool** | True when the declared age bands overlap (an open-ended elderly band declared beside narrower bands inside it). Age bands from such a source must not be summed; no sum is published. | [optional] 
**CdmMapping** | Pointer to **NullableBool** | Whether the holder declares the data are mapped to a common data model (CDM), such as OMOP. Read it with cdm_readiness, which says whether the transform is finished. | [optional] 
**EtlStatus** | Pointer to **NullableString** | The extract-transform-load (ETL) status of the CDM mapping as typed; one comma-joined answer per declared mapping (rwd_source_etl_mappings has one row per mapping). | [optional] 
**EtlCdmVersion** | Pointer to **NullableString** | The CDM version(s) declared for the ETL, comma-joined as typed. | [optional] 
**EtlFrequency** | Pointer to **NullableString** | The ETL frequency declared, comma-joined as typed. | [optional] 
**EtlSpecificationsFile** | Pointer to **NullableString** | The ETL specifications file the holder attached, if any (empty at source on the registered sources). | [optional] 
**DeclaredMappingCount** | Pointer to **NullableInt32** | How many CDM mappings the holder declared, after register join artefacts are removed; 0 when the question was never answered. | [optional] 
**CdmReadiness** | Pointer to **NullableString** | The CDM fact as one closed label that carries the word the holder wrote about the transform, so cdm_mapping alone is never read as ready. Exhaustive and never null: not_declared, no_mapping, mapping_without_etl_status, etl_completed, etl_in_progress, etl_planned, etl_not_done, etl_mixed (several mappings whose statuses differ or one has none) or etl_other (an unrecognised word). | [optional] 
**Linkage** | Pointer to **NullableBool** | Whether the holder declares that the data can be linked to other data sources. | [optional] 
**HasPreLinkedDeclaration** | Pointer to **NullableBool** | Whether the holder gave a description of a linkage already made (pre-linked). | [optional] 
**HasPossibleLinkageDeclaration** | Pointer to **NullableBool** | Whether the holder gave a description of a linkage that is possible but not done. | [optional] 
**LinkedDataSource** | Pointer to **NullableString** | The linked data source name(s) as typed, comma-joined; rwd_source_linkages has one row per declared linkage. | [optional] 
**LinkedDataSourceOther** | Pointer to **NullableString** | The linked data source named in free text when the holder chose Other data source. | [optional] 
**LinkedLinkageCompleteness** | Pointer to **NullableString** | The linkage completeness the holder declared, as typed. | [optional] 
**LinkedLinkageVariable** | Pointer to **NullableString** | The linkage variable(s) the holder declared, as typed. | [optional] 
**LinkedLinkageStrategy** | Pointer to **NullableString** | The linkage strategy the holder declared, as typed. | [optional] 
**DeclaredLinkageCount** | Pointer to **NullableInt32** | How many linkages the holder declared; 0 when none. | [optional] 
**LinkageState** | Pointer to **NullableString** | Five states, exhaustive and never null: pre_linked, possible_not_done, declared_without_detail, no_linkage, not_declared. Pre-linked and possible linkage are kept apart on purpose. | [optional] 
**BiospecimenAccess** | Pointer to **NullableBool** | Whether biospecimens can be accessed through the source; true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**AccessToSubjectDetails** | Pointer to **NullableBool** | Whether the holder can access subject details; true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**PossibilityOfDataValidation** | Pointer to **NullableBool** | Whether the data can be validated; true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**RecordsArePreservedIndefinitely** | Pointer to **NullableBool** | Whether records are preserved indefinitely; true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**DataSourcePreservationLengthYears** | Pointer to **NullableString** | How long the data are preserved, in years, as typed. | [optional] 
**ApprovalForPublication** | Pointer to **NullableBool** | Whether approval for publication is needed; true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] 
**InformedConsentForUseOfDataForResearch** | Pointer to **NullableString** | The informed consent regime for research use as the holder declared it, as typed. | [optional] 
**GovernanceDetailsLink** | Pointer to **NullableString** | A link to the source&#39;s governance details, as the holder registered it. | [optional] 
**DeclaredCapabilityCount** | Pointer to **NullableInt32** | How many of the 35 capability and governance questions the holder answered (a count of answered form fields, not a measure of the data held). | [optional] 
**DeclarableCapabilityCount** | Pointer to **NullableInt32** | How many such questions the form offers (35); the denominator of declared_capability_count. | [optional] 
**RegisterSourceCount** | Pointer to **NullableInt32** | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional] 
**StudyCount** | Pointer to **NullableInt32** | How many registered studies named this source. A registered study NAMED the source; it is never evidence that the source holds data on a condition. Read it with study_register_total and studies_naming_a_registered_source, because a zero is not a finding that the source is unused. | [optional] 
**StudyRegisterTotal** | Pointer to **NullableInt32** | The number of studies in the whole study register at the snapshot. | [optional] 
**StudiesNamingARegisteredSource** | Pointer to **NullableInt32** | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdSourcesRow

`func NewRwdSourcesRow() *RwdSourcesRow`

NewRwdSourcesRow instantiates a new RwdSourcesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdSourcesRowWithDefaults

`func NewRwdSourcesRowWithDefaults() *RwdSourcesRow`

NewRwdSourcesRowWithDefaults instantiates a new RwdSourcesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *RwdSourcesRow) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *RwdSourcesRow) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *RwdSourcesRow) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *RwdSourcesRow) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *RwdSourcesRow) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *RwdSourcesRow) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetDataSourceName

`func (o *RwdSourcesRow) GetDataSourceName() string`

GetDataSourceName returns the DataSourceName field if non-nil, zero value otherwise.

### GetDataSourceNameOk

`func (o *RwdSourcesRow) GetDataSourceNameOk() (*string, bool)`

GetDataSourceNameOk returns a tuple with the DataSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceName

`func (o *RwdSourcesRow) SetDataSourceName(v string)`

SetDataSourceName sets DataSourceName field to given value.

### HasDataSourceName

`func (o *RwdSourcesRow) HasDataSourceName() bool`

HasDataSourceName returns a boolean if a field has been set.

### SetDataSourceNameNil

`func (o *RwdSourcesRow) SetDataSourceNameNil(b bool)`

 SetDataSourceNameNil sets the value for DataSourceName to be an explicit nil

### UnsetDataSourceName
`func (o *RwdSourcesRow) UnsetDataSourceName()`

UnsetDataSourceName ensures that no value is present for DataSourceName, not even an explicit nil
### GetDataSourceAcronym

`func (o *RwdSourcesRow) GetDataSourceAcronym() string`

GetDataSourceAcronym returns the DataSourceAcronym field if non-nil, zero value otherwise.

### GetDataSourceAcronymOk

`func (o *RwdSourcesRow) GetDataSourceAcronymOk() (*string, bool)`

GetDataSourceAcronymOk returns a tuple with the DataSourceAcronym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceAcronym

`func (o *RwdSourcesRow) SetDataSourceAcronym(v string)`

SetDataSourceAcronym sets DataSourceAcronym field to given value.

### HasDataSourceAcronym

`func (o *RwdSourcesRow) HasDataSourceAcronym() bool`

HasDataSourceAcronym returns a boolean if a field has been set.

### SetDataSourceAcronymNil

`func (o *RwdSourcesRow) SetDataSourceAcronymNil(b bool)`

 SetDataSourceAcronymNil sets the value for DataSourceAcronym to be an explicit nil

### UnsetDataSourceAcronym
`func (o *RwdSourcesRow) UnsetDataSourceAcronym()`

UnsetDataSourceAcronym ensures that no value is present for DataSourceAcronym, not even an explicit nil
### GetDataHolder

`func (o *RwdSourcesRow) GetDataHolder() string`

GetDataHolder returns the DataHolder field if non-nil, zero value otherwise.

### GetDataHolderOk

`func (o *RwdSourcesRow) GetDataHolderOk() (*string, bool)`

GetDataHolderOk returns a tuple with the DataHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataHolder

`func (o *RwdSourcesRow) SetDataHolder(v string)`

SetDataHolder sets DataHolder field to given value.

### HasDataHolder

`func (o *RwdSourcesRow) HasDataHolder() bool`

HasDataHolder returns a boolean if a field has been set.

### SetDataHolderNil

`func (o *RwdSourcesRow) SetDataHolderNil(b bool)`

 SetDataHolderNil sets the value for DataHolder to be an explicit nil

### UnsetDataHolder
`func (o *RwdSourcesRow) UnsetDataHolder()`

UnsetDataHolder ensures that no value is present for DataHolder, not even an explicit nil
### GetDataSourceType

`func (o *RwdSourcesRow) GetDataSourceType() string`

GetDataSourceType returns the DataSourceType field if non-nil, zero value otherwise.

### GetDataSourceTypeOk

`func (o *RwdSourcesRow) GetDataSourceTypeOk() (*string, bool)`

GetDataSourceTypeOk returns a tuple with the DataSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceType

`func (o *RwdSourcesRow) SetDataSourceType(v string)`

SetDataSourceType sets DataSourceType field to given value.

### HasDataSourceType

`func (o *RwdSourcesRow) HasDataSourceType() bool`

HasDataSourceType returns a boolean if a field has been set.

### SetDataSourceTypeNil

`func (o *RwdSourcesRow) SetDataSourceTypeNil(b bool)`

 SetDataSourceTypeNil sets the value for DataSourceType to be an explicit nil

### UnsetDataSourceType
`func (o *RwdSourcesRow) UnsetDataSourceType()`

UnsetDataSourceType ensures that no value is present for DataSourceType, not even an explicit nil
### GetDataSourceTypeOther

`func (o *RwdSourcesRow) GetDataSourceTypeOther() string`

GetDataSourceTypeOther returns the DataSourceTypeOther field if non-nil, zero value otherwise.

### GetDataSourceTypeOtherOk

`func (o *RwdSourcesRow) GetDataSourceTypeOtherOk() (*string, bool)`

GetDataSourceTypeOtherOk returns a tuple with the DataSourceTypeOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceTypeOther

`func (o *RwdSourcesRow) SetDataSourceTypeOther(v string)`

SetDataSourceTypeOther sets DataSourceTypeOther field to given value.

### HasDataSourceTypeOther

`func (o *RwdSourcesRow) HasDataSourceTypeOther() bool`

HasDataSourceTypeOther returns a boolean if a field has been set.

### SetDataSourceTypeOtherNil

`func (o *RwdSourcesRow) SetDataSourceTypeOtherNil(b bool)`

 SetDataSourceTypeOtherNil sets the value for DataSourceTypeOther to be an explicit nil

### UnsetDataSourceTypeOther
`func (o *RwdSourcesRow) UnsetDataSourceTypeOther()`

UnsetDataSourceTypeOther ensures that no value is present for DataSourceTypeOther, not even an explicit nil
### GetDataSourceDomain

`func (o *RwdSourcesRow) GetDataSourceDomain() string`

GetDataSourceDomain returns the DataSourceDomain field if non-nil, zero value otherwise.

### GetDataSourceDomainOk

`func (o *RwdSourcesRow) GetDataSourceDomainOk() (*string, bool)`

GetDataSourceDomainOk returns a tuple with the DataSourceDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceDomain

`func (o *RwdSourcesRow) SetDataSourceDomain(v string)`

SetDataSourceDomain sets DataSourceDomain field to given value.

### HasDataSourceDomain

`func (o *RwdSourcesRow) HasDataSourceDomain() bool`

HasDataSourceDomain returns a boolean if a field has been set.

### SetDataSourceDomainNil

`func (o *RwdSourcesRow) SetDataSourceDomainNil(b bool)`

 SetDataSourceDomainNil sets the value for DataSourceDomain to be an explicit nil

### UnsetDataSourceDomain
`func (o *RwdSourcesRow) UnsetDataSourceDomain()`

UnsetDataSourceDomain ensures that no value is present for DataSourceDomain, not even an explicit nil
### GetCareSetting

`func (o *RwdSourcesRow) GetCareSetting() string`

GetCareSetting returns the CareSetting field if non-nil, zero value otherwise.

### GetCareSettingOk

`func (o *RwdSourcesRow) GetCareSettingOk() (*string, bool)`

GetCareSettingOk returns a tuple with the CareSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareSetting

`func (o *RwdSourcesRow) SetCareSetting(v string)`

SetCareSetting sets CareSetting field to given value.

### HasCareSetting

`func (o *RwdSourcesRow) HasCareSetting() bool`

HasCareSetting returns a boolean if a field has been set.

### SetCareSettingNil

`func (o *RwdSourcesRow) SetCareSettingNil(b bool)`

 SetCareSettingNil sets the value for CareSetting to be an explicit nil

### UnsetCareSetting
`func (o *RwdSourcesRow) UnsetCareSetting()`

UnsetCareSetting ensures that no value is present for CareSetting, not even an explicit nil
### GetMainFinancialSupport

`func (o *RwdSourcesRow) GetMainFinancialSupport() string`

GetMainFinancialSupport returns the MainFinancialSupport field if non-nil, zero value otherwise.

### GetMainFinancialSupportOk

`func (o *RwdSourcesRow) GetMainFinancialSupportOk() (*string, bool)`

GetMainFinancialSupportOk returns a tuple with the MainFinancialSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainFinancialSupport

`func (o *RwdSourcesRow) SetMainFinancialSupport(v string)`

SetMainFinancialSupport sets MainFinancialSupport field to given value.

### HasMainFinancialSupport

`func (o *RwdSourcesRow) HasMainFinancialSupport() bool`

HasMainFinancialSupport returns a boolean if a field has been set.

### SetMainFinancialSupportNil

`func (o *RwdSourcesRow) SetMainFinancialSupportNil(b bool)`

 SetMainFinancialSupportNil sets the value for MainFinancialSupport to be an explicit nil

### UnsetMainFinancialSupport
`func (o *RwdSourcesRow) UnsetMainFinancialSupport()`

UnsetMainFinancialSupport ensures that no value is present for MainFinancialSupport, not even an explicit nil
### GetCountryNames

`func (o *RwdSourcesRow) GetCountryNames() []string`

GetCountryNames returns the CountryNames field if non-nil, zero value otherwise.

### GetCountryNamesOk

`func (o *RwdSourcesRow) GetCountryNamesOk() (*[]string, bool)`

GetCountryNamesOk returns a tuple with the CountryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryNames

`func (o *RwdSourcesRow) SetCountryNames(v []string)`

SetCountryNames sets CountryNames field to given value.

### HasCountryNames

`func (o *RwdSourcesRow) HasCountryNames() bool`

HasCountryNames returns a boolean if a field has been set.

### SetCountryNamesNil

`func (o *RwdSourcesRow) SetCountryNamesNil(b bool)`

 SetCountryNamesNil sets the value for CountryNames to be an explicit nil

### UnsetCountryNames
`func (o *RwdSourcesRow) UnsetCountryNames()`

UnsetCountryNames ensures that no value is present for CountryNames, not even an explicit nil
### GetUnmatchedCountryFragments

`func (o *RwdSourcesRow) GetUnmatchedCountryFragments() []string`

GetUnmatchedCountryFragments returns the UnmatchedCountryFragments field if non-nil, zero value otherwise.

### GetUnmatchedCountryFragmentsOk

`func (o *RwdSourcesRow) GetUnmatchedCountryFragmentsOk() (*[]string, bool)`

GetUnmatchedCountryFragmentsOk returns a tuple with the UnmatchedCountryFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedCountryFragments

`func (o *RwdSourcesRow) SetUnmatchedCountryFragments(v []string)`

SetUnmatchedCountryFragments sets UnmatchedCountryFragments field to given value.

### HasUnmatchedCountryFragments

`func (o *RwdSourcesRow) HasUnmatchedCountryFragments() bool`

HasUnmatchedCountryFragments returns a boolean if a field has been set.

### SetUnmatchedCountryFragmentsNil

`func (o *RwdSourcesRow) SetUnmatchedCountryFragmentsNil(b bool)`

 SetUnmatchedCountryFragmentsNil sets the value for UnmatchedCountryFragments to be an explicit nil

### UnsetUnmatchedCountryFragments
`func (o *RwdSourcesRow) UnsetUnmatchedCountryFragments()`

UnsetUnmatchedCountryFragments ensures that no value is present for UnmatchedCountryFragments, not even an explicit nil
### GetCountryCount

`func (o *RwdSourcesRow) GetCountryCount() int32`

GetCountryCount returns the CountryCount field if non-nil, zero value otherwise.

### GetCountryCountOk

`func (o *RwdSourcesRow) GetCountryCountOk() (*int32, bool)`

GetCountryCountOk returns a tuple with the CountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCount

`func (o *RwdSourcesRow) SetCountryCount(v int32)`

SetCountryCount sets CountryCount field to given value.

### HasCountryCount

`func (o *RwdSourcesRow) HasCountryCount() bool`

HasCountryCount returns a boolean if a field has been set.

### SetCountryCountNil

`func (o *RwdSourcesRow) SetCountryCountNil(b bool)`

 SetCountryCountNil sets the value for CountryCount to be an explicit nil

### UnsetCountryCount
`func (o *RwdSourcesRow) UnsetCountryCount()`

UnsetCountryCount ensures that no value is present for CountryCount, not even an explicit nil
### GetCatalogueUrl

`func (o *RwdSourcesRow) GetCatalogueUrl() string`

GetCatalogueUrl returns the CatalogueUrl field if non-nil, zero value otherwise.

### GetCatalogueUrlOk

`func (o *RwdSourcesRow) GetCatalogueUrlOk() (*string, bool)`

GetCatalogueUrlOk returns a tuple with the CatalogueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueUrl

`func (o *RwdSourcesRow) SetCatalogueUrl(v string)`

SetCatalogueUrl sets CatalogueUrl field to given value.

### HasCatalogueUrl

`func (o *RwdSourcesRow) HasCatalogueUrl() bool`

HasCatalogueUrl returns a boolean if a field has been set.

### SetCatalogueUrlNil

`func (o *RwdSourcesRow) SetCatalogueUrlNil(b bool)`

 SetCatalogueUrlNil sets the value for CatalogueUrl to be an explicit nil

### UnsetCatalogueUrl
`func (o *RwdSourcesRow) UnsetCatalogueUrl()`

UnsetCatalogueUrl ensures that no value is present for CatalogueUrl, not even an explicit nil
### GetDataSourceWebsite

`func (o *RwdSourcesRow) GetDataSourceWebsite() string`

GetDataSourceWebsite returns the DataSourceWebsite field if non-nil, zero value otherwise.

### GetDataSourceWebsiteOk

`func (o *RwdSourcesRow) GetDataSourceWebsiteOk() (*string, bool)`

GetDataSourceWebsiteOk returns a tuple with the DataSourceWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceWebsite

`func (o *RwdSourcesRow) SetDataSourceWebsite(v string)`

SetDataSourceWebsite sets DataSourceWebsite field to given value.

### HasDataSourceWebsite

`func (o *RwdSourcesRow) HasDataSourceWebsite() bool`

HasDataSourceWebsite returns a boolean if a field has been set.

### SetDataSourceWebsiteNil

`func (o *RwdSourcesRow) SetDataSourceWebsiteNil(b bool)`

 SetDataSourceWebsiteNil sets the value for DataSourceWebsite to be an explicit nil

### UnsetDataSourceWebsite
`func (o *RwdSourcesRow) UnsetDataSourceWebsite()`

UnsetDataSourceWebsite ensures that no value is present for DataSourceWebsite, not even an explicit nil
### GetIsDarwinEuDataPartner

`func (o *RwdSourcesRow) GetIsDarwinEuDataPartner() bool`

GetIsDarwinEuDataPartner returns the IsDarwinEuDataPartner field if non-nil, zero value otherwise.

### GetIsDarwinEuDataPartnerOk

`func (o *RwdSourcesRow) GetIsDarwinEuDataPartnerOk() (*bool, bool)`

GetIsDarwinEuDataPartnerOk returns a tuple with the IsDarwinEuDataPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDarwinEuDataPartner

`func (o *RwdSourcesRow) SetIsDarwinEuDataPartner(v bool)`

SetIsDarwinEuDataPartner sets IsDarwinEuDataPartner field to given value.

### HasIsDarwinEuDataPartner

`func (o *RwdSourcesRow) HasIsDarwinEuDataPartner() bool`

HasIsDarwinEuDataPartner returns a boolean if a field has been set.

### SetIsDarwinEuDataPartnerNil

`func (o *RwdSourcesRow) SetIsDarwinEuDataPartnerNil(b bool)`

 SetIsDarwinEuDataPartnerNil sets the value for IsDarwinEuDataPartner to be an explicit nil

### UnsetIsDarwinEuDataPartner
`func (o *RwdSourcesRow) UnsetIsDarwinEuDataPartner()`

UnsetIsDarwinEuDataPartner ensures that no value is present for IsDarwinEuDataPartner, not even an explicit nil
### GetIsQualified

`func (o *RwdSourcesRow) GetIsQualified() bool`

GetIsQualified returns the IsQualified field if non-nil, zero value otherwise.

### GetIsQualifiedOk

`func (o *RwdSourcesRow) GetIsQualifiedOk() (*bool, bool)`

GetIsQualifiedOk returns a tuple with the IsQualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQualified

`func (o *RwdSourcesRow) SetIsQualified(v bool)`

SetIsQualified sets IsQualified field to given value.

### HasIsQualified

`func (o *RwdSourcesRow) HasIsQualified() bool`

HasIsQualified returns a boolean if a field has been set.

### SetIsQualifiedNil

`func (o *RwdSourcesRow) SetIsQualifiedNil(b bool)`

 SetIsQualifiedNil sets the value for IsQualified to be an explicit nil

### UnsetIsQualified
`func (o *RwdSourcesRow) UnsetIsQualified()`

UnsetIsQualified ensures that no value is present for IsQualified, not even an explicit nil
### GetRefreshCadence

`func (o *RwdSourcesRow) GetRefreshCadence() string`

GetRefreshCadence returns the RefreshCadence field if non-nil, zero value otherwise.

### GetRefreshCadenceOk

`func (o *RwdSourcesRow) GetRefreshCadenceOk() (*string, bool)`

GetRefreshCadenceOk returns a tuple with the RefreshCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshCadence

`func (o *RwdSourcesRow) SetRefreshCadence(v string)`

SetRefreshCadence sets RefreshCadence field to given value.

### HasRefreshCadence

`func (o *RwdSourcesRow) HasRefreshCadence() bool`

HasRefreshCadence returns a boolean if a field has been set.

### SetRefreshCadenceNil

`func (o *RwdSourcesRow) SetRefreshCadenceNil(b bool)`

 SetRefreshCadenceNil sets the value for RefreshCadence to be an explicit nil

### UnsetRefreshCadence
`func (o *RwdSourcesRow) UnsetRefreshCadence()`

UnsetRefreshCadence ensures that no value is present for RefreshCadence, not even an explicit nil
### GetLastRefreshedOn

`func (o *RwdSourcesRow) GetLastRefreshedOn() string`

GetLastRefreshedOn returns the LastRefreshedOn field if non-nil, zero value otherwise.

### GetLastRefreshedOnOk

`func (o *RwdSourcesRow) GetLastRefreshedOnOk() (*string, bool)`

GetLastRefreshedOnOk returns a tuple with the LastRefreshedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshedOn

`func (o *RwdSourcesRow) SetLastRefreshedOn(v string)`

SetLastRefreshedOn sets LastRefreshedOn field to given value.

### HasLastRefreshedOn

`func (o *RwdSourcesRow) HasLastRefreshedOn() bool`

HasLastRefreshedOn returns a boolean if a field has been set.

### SetLastRefreshedOnNil

`func (o *RwdSourcesRow) SetLastRefreshedOnNil(b bool)`

 SetLastRefreshedOnNil sets the value for LastRefreshedOn to be an explicit nil

### UnsetLastRefreshedOn
`func (o *RwdSourcesRow) UnsetLastRefreshedOn()`

UnsetLastRefreshedOn ensures that no value is present for LastRefreshedOn, not even an explicit nil
### GetRecordUpdatedAt

`func (o *RwdSourcesRow) GetRecordUpdatedAt() time.Time`

GetRecordUpdatedAt returns the RecordUpdatedAt field if non-nil, zero value otherwise.

### GetRecordUpdatedAtOk

`func (o *RwdSourcesRow) GetRecordUpdatedAtOk() (*time.Time, bool)`

GetRecordUpdatedAtOk returns a tuple with the RecordUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedAt

`func (o *RwdSourcesRow) SetRecordUpdatedAt(v time.Time)`

SetRecordUpdatedAt sets RecordUpdatedAt field to given value.

### HasRecordUpdatedAt

`func (o *RwdSourcesRow) HasRecordUpdatedAt() bool`

HasRecordUpdatedAt returns a boolean if a field has been set.

### SetRecordUpdatedAtNil

`func (o *RwdSourcesRow) SetRecordUpdatedAtNil(b bool)`

 SetRecordUpdatedAtNil sets the value for RecordUpdatedAt to be an explicit nil

### UnsetRecordUpdatedAt
`func (o *RwdSourcesRow) UnsetRecordUpdatedAt()`

UnsetRecordUpdatedAt ensures that no value is present for RecordUpdatedAt, not even an explicit nil
### GetRecordFirstPublishedAt

`func (o *RwdSourcesRow) GetRecordFirstPublishedAt() time.Time`

GetRecordFirstPublishedAt returns the RecordFirstPublishedAt field if non-nil, zero value otherwise.

### GetRecordFirstPublishedAtOk

`func (o *RwdSourcesRow) GetRecordFirstPublishedAtOk() (*time.Time, bool)`

GetRecordFirstPublishedAtOk returns a tuple with the RecordFirstPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFirstPublishedAt

`func (o *RwdSourcesRow) SetRecordFirstPublishedAt(v time.Time)`

SetRecordFirstPublishedAt sets RecordFirstPublishedAt field to given value.

### HasRecordFirstPublishedAt

`func (o *RwdSourcesRow) HasRecordFirstPublishedAt() bool`

HasRecordFirstPublishedAt returns a boolean if a field has been set.

### SetRecordFirstPublishedAtNil

`func (o *RwdSourcesRow) SetRecordFirstPublishedAtNil(b bool)`

 SetRecordFirstPublishedAtNil sets the value for RecordFirstPublishedAt to be an explicit nil

### UnsetRecordFirstPublishedAt
`func (o *RwdSourcesRow) UnsetRecordFirstPublishedAt()`

UnsetRecordFirstPublishedAt ensures that no value is present for RecordFirstPublishedAt, not even an explicit nil
### GetIsOutdated

`func (o *RwdSourcesRow) GetIsOutdated() bool`

GetIsOutdated returns the IsOutdated field if non-nil, zero value otherwise.

### GetIsOutdatedOk

`func (o *RwdSourcesRow) GetIsOutdatedOk() (*bool, bool)`

GetIsOutdatedOk returns a tuple with the IsOutdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOutdated

`func (o *RwdSourcesRow) SetIsOutdated(v bool)`

SetIsOutdated sets IsOutdated field to given value.

### HasIsOutdated

`func (o *RwdSourcesRow) HasIsOutdated() bool`

HasIsOutdated returns a boolean if a field has been set.

### SetIsOutdatedNil

`func (o *RwdSourcesRow) SetIsOutdatedNil(b bool)`

 SetIsOutdatedNil sets the value for IsOutdated to be an explicit nil

### UnsetIsOutdated
`func (o *RwdSourcesRow) UnsetIsOutdated()`

UnsetIsOutdated ensures that no value is present for IsOutdated, not even an explicit nil
### GetRecordUpdatedOver12MonthsAgo

`func (o *RwdSourcesRow) GetRecordUpdatedOver12MonthsAgo() bool`

GetRecordUpdatedOver12MonthsAgo returns the RecordUpdatedOver12MonthsAgo field if non-nil, zero value otherwise.

### GetRecordUpdatedOver12MonthsAgoOk

`func (o *RwdSourcesRow) GetRecordUpdatedOver12MonthsAgoOk() (*bool, bool)`

GetRecordUpdatedOver12MonthsAgoOk returns a tuple with the RecordUpdatedOver12MonthsAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedOver12MonthsAgo

`func (o *RwdSourcesRow) SetRecordUpdatedOver12MonthsAgo(v bool)`

SetRecordUpdatedOver12MonthsAgo sets RecordUpdatedOver12MonthsAgo field to given value.

### HasRecordUpdatedOver12MonthsAgo

`func (o *RwdSourcesRow) HasRecordUpdatedOver12MonthsAgo() bool`

HasRecordUpdatedOver12MonthsAgo returns a boolean if a field has been set.

### SetRecordUpdatedOver12MonthsAgoNil

`func (o *RwdSourcesRow) SetRecordUpdatedOver12MonthsAgoNil(b bool)`

 SetRecordUpdatedOver12MonthsAgoNil sets the value for RecordUpdatedOver12MonthsAgo to be an explicit nil

### UnsetRecordUpdatedOver12MonthsAgo
`func (o *RwdSourcesRow) UnsetRecordUpdatedOver12MonthsAgo()`

UnsetRecordUpdatedOver12MonthsAgo ensures that no value is present for RecordUpdatedOver12MonthsAgo, not even an explicit nil
### GetRefreshIntervalDays

`func (o *RwdSourcesRow) GetRefreshIntervalDays() int32`

GetRefreshIntervalDays returns the RefreshIntervalDays field if non-nil, zero value otherwise.

### GetRefreshIntervalDaysOk

`func (o *RwdSourcesRow) GetRefreshIntervalDaysOk() (*int32, bool)`

GetRefreshIntervalDaysOk returns a tuple with the RefreshIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshIntervalDays

`func (o *RwdSourcesRow) SetRefreshIntervalDays(v int32)`

SetRefreshIntervalDays sets RefreshIntervalDays field to given value.

### HasRefreshIntervalDays

`func (o *RwdSourcesRow) HasRefreshIntervalDays() bool`

HasRefreshIntervalDays returns a boolean if a field has been set.

### SetRefreshIntervalDaysNil

`func (o *RwdSourcesRow) SetRefreshIntervalDaysNil(b bool)`

 SetRefreshIntervalDaysNil sets the value for RefreshIntervalDays to be an explicit nil

### UnsetRefreshIntervalDays
`func (o *RwdSourcesRow) UnsetRefreshIntervalDays()`

UnsetRefreshIntervalDays ensures that no value is present for RefreshIntervalDays, not even an explicit nil
### GetRefreshIntervalBasis

`func (o *RwdSourcesRow) GetRefreshIntervalBasis() string`

GetRefreshIntervalBasis returns the RefreshIntervalBasis field if non-nil, zero value otherwise.

### GetRefreshIntervalBasisOk

`func (o *RwdSourcesRow) GetRefreshIntervalBasisOk() (*string, bool)`

GetRefreshIntervalBasisOk returns a tuple with the RefreshIntervalBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshIntervalBasis

`func (o *RwdSourcesRow) SetRefreshIntervalBasis(v string)`

SetRefreshIntervalBasis sets RefreshIntervalBasis field to given value.

### HasRefreshIntervalBasis

`func (o *RwdSourcesRow) HasRefreshIntervalBasis() bool`

HasRefreshIntervalBasis returns a boolean if a field has been set.

### SetRefreshIntervalBasisNil

`func (o *RwdSourcesRow) SetRefreshIntervalBasisNil(b bool)`

 SetRefreshIntervalBasisNil sets the value for RefreshIntervalBasis to be an explicit nil

### UnsetRefreshIntervalBasis
`func (o *RwdSourcesRow) UnsetRefreshIntervalBasis()`

UnsetRefreshIntervalBasis ensures that no value is present for RefreshIntervalBasis, not even an explicit nil
### GetRefreshIntervalDaysShortest

`func (o *RwdSourcesRow) GetRefreshIntervalDaysShortest() int32`

GetRefreshIntervalDaysShortest returns the RefreshIntervalDaysShortest field if non-nil, zero value otherwise.

### GetRefreshIntervalDaysShortestOk

`func (o *RwdSourcesRow) GetRefreshIntervalDaysShortestOk() (*int32, bool)`

GetRefreshIntervalDaysShortestOk returns a tuple with the RefreshIntervalDaysShortest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshIntervalDaysShortest

`func (o *RwdSourcesRow) SetRefreshIntervalDaysShortest(v int32)`

SetRefreshIntervalDaysShortest sets RefreshIntervalDaysShortest field to given value.

### HasRefreshIntervalDaysShortest

`func (o *RwdSourcesRow) HasRefreshIntervalDaysShortest() bool`

HasRefreshIntervalDaysShortest returns a boolean if a field has been set.

### SetRefreshIntervalDaysShortestNil

`func (o *RwdSourcesRow) SetRefreshIntervalDaysShortestNil(b bool)`

 SetRefreshIntervalDaysShortestNil sets the value for RefreshIntervalDaysShortest to be an explicit nil

### UnsetRefreshIntervalDaysShortest
`func (o *RwdSourcesRow) UnsetRefreshIntervalDaysShortest()`

UnsetRefreshIntervalDaysShortest ensures that no value is present for RefreshIntervalDaysShortest, not even an explicit nil
### GetRefreshIntervalBasisShortest

`func (o *RwdSourcesRow) GetRefreshIntervalBasisShortest() string`

GetRefreshIntervalBasisShortest returns the RefreshIntervalBasisShortest field if non-nil, zero value otherwise.

### GetRefreshIntervalBasisShortestOk

`func (o *RwdSourcesRow) GetRefreshIntervalBasisShortestOk() (*string, bool)`

GetRefreshIntervalBasisShortestOk returns a tuple with the RefreshIntervalBasisShortest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshIntervalBasisShortest

`func (o *RwdSourcesRow) SetRefreshIntervalBasisShortest(v string)`

SetRefreshIntervalBasisShortest sets RefreshIntervalBasisShortest field to given value.

### HasRefreshIntervalBasisShortest

`func (o *RwdSourcesRow) HasRefreshIntervalBasisShortest() bool`

HasRefreshIntervalBasisShortest returns a boolean if a field has been set.

### SetRefreshIntervalBasisShortestNil

`func (o *RwdSourcesRow) SetRefreshIntervalBasisShortestNil(b bool)`

 SetRefreshIntervalBasisShortestNil sets the value for RefreshIntervalBasisShortest to be an explicit nil

### UnsetRefreshIntervalBasisShortest
`func (o *RwdSourcesRow) UnsetRefreshIntervalBasisShortest()`

UnsetRefreshIntervalBasisShortest ensures that no value is present for RefreshIntervalBasisShortest, not even an explicit nil
### GetRefreshCadenceDeclared

`func (o *RwdSourcesRow) GetRefreshCadenceDeclared() []string`

GetRefreshCadenceDeclared returns the RefreshCadenceDeclared field if non-nil, zero value otherwise.

### GetRefreshCadenceDeclaredOk

`func (o *RwdSourcesRow) GetRefreshCadenceDeclaredOk() (*[]string, bool)`

GetRefreshCadenceDeclaredOk returns a tuple with the RefreshCadenceDeclared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshCadenceDeclared

`func (o *RwdSourcesRow) SetRefreshCadenceDeclared(v []string)`

SetRefreshCadenceDeclared sets RefreshCadenceDeclared field to given value.

### HasRefreshCadenceDeclared

`func (o *RwdSourcesRow) HasRefreshCadenceDeclared() bool`

HasRefreshCadenceDeclared returns a boolean if a field has been set.

### SetRefreshCadenceDeclaredNil

`func (o *RwdSourcesRow) SetRefreshCadenceDeclaredNil(b bool)`

 SetRefreshCadenceDeclaredNil sets the value for RefreshCadenceDeclared to be an explicit nil

### UnsetRefreshCadenceDeclared
`func (o *RwdSourcesRow) UnsetRefreshCadenceDeclared()`

UnsetRefreshCadenceDeclared ensures that no value is present for RefreshCadenceDeclared, not even an explicit nil
### GetRefreshCadenceIntervalDays

`func (o *RwdSourcesRow) GetRefreshCadenceIntervalDays() []string`

GetRefreshCadenceIntervalDays returns the RefreshCadenceIntervalDays field if non-nil, zero value otherwise.

### GetRefreshCadenceIntervalDaysOk

`func (o *RwdSourcesRow) GetRefreshCadenceIntervalDaysOk() (*[]string, bool)`

GetRefreshCadenceIntervalDaysOk returns a tuple with the RefreshCadenceIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshCadenceIntervalDays

`func (o *RwdSourcesRow) SetRefreshCadenceIntervalDays(v []string)`

SetRefreshCadenceIntervalDays sets RefreshCadenceIntervalDays field to given value.

### HasRefreshCadenceIntervalDays

`func (o *RwdSourcesRow) HasRefreshCadenceIntervalDays() bool`

HasRefreshCadenceIntervalDays returns a boolean if a field has been set.

### SetRefreshCadenceIntervalDaysNil

`func (o *RwdSourcesRow) SetRefreshCadenceIntervalDaysNil(b bool)`

 SetRefreshCadenceIntervalDaysNil sets the value for RefreshCadenceIntervalDays to be an explicit nil

### UnsetRefreshCadenceIntervalDays
`func (o *RwdSourcesRow) UnsetRefreshCadenceIntervalDays()`

UnsetRefreshCadenceIntervalDays ensures that no value is present for RefreshCadenceIntervalDays, not even an explicit nil
### GetRefreshCadenceIsMonthNames

`func (o *RwdSourcesRow) GetRefreshCadenceIsMonthNames() []string`

GetRefreshCadenceIsMonthNames returns the RefreshCadenceIsMonthNames field if non-nil, zero value otherwise.

### GetRefreshCadenceIsMonthNamesOk

`func (o *RwdSourcesRow) GetRefreshCadenceIsMonthNamesOk() (*[]string, bool)`

GetRefreshCadenceIsMonthNamesOk returns a tuple with the RefreshCadenceIsMonthNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshCadenceIsMonthNames

`func (o *RwdSourcesRow) SetRefreshCadenceIsMonthNames(v []string)`

SetRefreshCadenceIsMonthNames sets RefreshCadenceIsMonthNames field to given value.

### HasRefreshCadenceIsMonthNames

`func (o *RwdSourcesRow) HasRefreshCadenceIsMonthNames() bool`

HasRefreshCadenceIsMonthNames returns a boolean if a field has been set.

### SetRefreshCadenceIsMonthNamesNil

`func (o *RwdSourcesRow) SetRefreshCadenceIsMonthNamesNil(b bool)`

 SetRefreshCadenceIsMonthNamesNil sets the value for RefreshCadenceIsMonthNames to be an explicit nil

### UnsetRefreshCadenceIsMonthNames
`func (o *RwdSourcesRow) UnsetRefreshCadenceIsMonthNames()`

UnsetRefreshCadenceIsMonthNames ensures that no value is present for RefreshCadenceIsMonthNames, not even an explicit nil
### GetDaysSinceLastRefresh

`func (o *RwdSourcesRow) GetDaysSinceLastRefresh() int32`

GetDaysSinceLastRefresh returns the DaysSinceLastRefresh field if non-nil, zero value otherwise.

### GetDaysSinceLastRefreshOk

`func (o *RwdSourcesRow) GetDaysSinceLastRefreshOk() (*int32, bool)`

GetDaysSinceLastRefreshOk returns a tuple with the DaysSinceLastRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSinceLastRefresh

`func (o *RwdSourcesRow) SetDaysSinceLastRefresh(v int32)`

SetDaysSinceLastRefresh sets DaysSinceLastRefresh field to given value.

### HasDaysSinceLastRefresh

`func (o *RwdSourcesRow) HasDaysSinceLastRefresh() bool`

HasDaysSinceLastRefresh returns a boolean if a field has been set.

### SetDaysSinceLastRefreshNil

`func (o *RwdSourcesRow) SetDaysSinceLastRefreshNil(b bool)`

 SetDaysSinceLastRefreshNil sets the value for DaysSinceLastRefresh to be an explicit nil

### UnsetDaysSinceLastRefresh
`func (o *RwdSourcesRow) UnsetDaysSinceLastRefresh()`

UnsetDaysSinceLastRefresh ensures that no value is present for DaysSinceLastRefresh, not even an explicit nil
### GetDatabaseRefreshState

`func (o *RwdSourcesRow) GetDatabaseRefreshState() string`

GetDatabaseRefreshState returns the DatabaseRefreshState field if non-nil, zero value otherwise.

### GetDatabaseRefreshStateOk

`func (o *RwdSourcesRow) GetDatabaseRefreshStateOk() (*string, bool)`

GetDatabaseRefreshStateOk returns a tuple with the DatabaseRefreshState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseRefreshState

`func (o *RwdSourcesRow) SetDatabaseRefreshState(v string)`

SetDatabaseRefreshState sets DatabaseRefreshState field to given value.

### HasDatabaseRefreshState

`func (o *RwdSourcesRow) HasDatabaseRefreshState() bool`

HasDatabaseRefreshState returns a boolean if a field has been set.

### SetDatabaseRefreshStateNil

`func (o *RwdSourcesRow) SetDatabaseRefreshStateNil(b bool)`

 SetDatabaseRefreshStateNil sets the value for DatabaseRefreshState to be an explicit nil

### UnsetDatabaseRefreshState
`func (o *RwdSourcesRow) UnsetDatabaseRefreshState()`

UnsetDatabaseRefreshState ensures that no value is present for DatabaseRefreshState, not even an explicit nil
### GetRefreshStateStrict

`func (o *RwdSourcesRow) GetRefreshStateStrict() string`

GetRefreshStateStrict returns the RefreshStateStrict field if non-nil, zero value otherwise.

### GetRefreshStateStrictOk

`func (o *RwdSourcesRow) GetRefreshStateStrictOk() (*string, bool)`

GetRefreshStateStrictOk returns a tuple with the RefreshStateStrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshStateStrict

`func (o *RwdSourcesRow) SetRefreshStateStrict(v string)`

SetRefreshStateStrict sets RefreshStateStrict field to given value.

### HasRefreshStateStrict

`func (o *RwdSourcesRow) HasRefreshStateStrict() bool`

HasRefreshStateStrict returns a boolean if a field has been set.

### SetRefreshStateStrictNil

`func (o *RwdSourcesRow) SetRefreshStateStrictNil(b bool)`

 SetRefreshStateStrictNil sets the value for RefreshStateStrict to be an explicit nil

### UnsetRefreshStateStrict
`func (o *RwdSourcesRow) UnsetRefreshStateStrict()`

UnsetRefreshStateStrict ensures that no value is present for RefreshStateStrict, not even an explicit nil
### GetRefreshRecencyBand

`func (o *RwdSourcesRow) GetRefreshRecencyBand() string`

GetRefreshRecencyBand returns the RefreshRecencyBand field if non-nil, zero value otherwise.

### GetRefreshRecencyBandOk

`func (o *RwdSourcesRow) GetRefreshRecencyBandOk() (*string, bool)`

GetRefreshRecencyBandOk returns a tuple with the RefreshRecencyBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshRecencyBand

`func (o *RwdSourcesRow) SetRefreshRecencyBand(v string)`

SetRefreshRecencyBand sets RefreshRecencyBand field to given value.

### HasRefreshRecencyBand

`func (o *RwdSourcesRow) HasRefreshRecencyBand() bool`

HasRefreshRecencyBand returns a boolean if a field has been set.

### SetRefreshRecencyBandNil

`func (o *RwdSourcesRow) SetRefreshRecencyBandNil(b bool)`

 SetRefreshRecencyBandNil sets the value for RefreshRecencyBand to be an explicit nil

### UnsetRefreshRecencyBand
`func (o *RwdSourcesRow) UnsetRefreshRecencyBand()`

UnsetRefreshRecencyBand ensures that no value is present for RefreshRecencyBand, not even an explicit nil
### GetCollectionLastBand

`func (o *RwdSourcesRow) GetCollectionLastBand() string`

GetCollectionLastBand returns the CollectionLastBand field if non-nil, zero value otherwise.

### GetCollectionLastBandOk

`func (o *RwdSourcesRow) GetCollectionLastBandOk() (*string, bool)`

GetCollectionLastBandOk returns a tuple with the CollectionLastBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionLastBand

`func (o *RwdSourcesRow) SetCollectionLastBand(v string)`

SetCollectionLastBand sets CollectionLastBand field to given value.

### HasCollectionLastBand

`func (o *RwdSourcesRow) HasCollectionLastBand() bool`

HasCollectionLastBand returns a boolean if a field has been set.

### SetCollectionLastBandNil

`func (o *RwdSourcesRow) SetCollectionLastBandNil(b bool)`

 SetCollectionLastBandNil sets the value for CollectionLastBand to be an explicit nil

### UnsetCollectionLastBand
`func (o *RwdSourcesRow) UnsetCollectionLastBand()`

UnsetCollectionLastBand ensures that no value is present for CollectionLastBand, not even an explicit nil
### GetCollectionContradictsRefreshBand

`func (o *RwdSourcesRow) GetCollectionContradictsRefreshBand() bool`

GetCollectionContradictsRefreshBand returns the CollectionContradictsRefreshBand field if non-nil, zero value otherwise.

### GetCollectionContradictsRefreshBandOk

`func (o *RwdSourcesRow) GetCollectionContradictsRefreshBandOk() (*bool, bool)`

GetCollectionContradictsRefreshBandOk returns a tuple with the CollectionContradictsRefreshBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionContradictsRefreshBand

`func (o *RwdSourcesRow) SetCollectionContradictsRefreshBand(v bool)`

SetCollectionContradictsRefreshBand sets CollectionContradictsRefreshBand field to given value.

### HasCollectionContradictsRefreshBand

`func (o *RwdSourcesRow) HasCollectionContradictsRefreshBand() bool`

HasCollectionContradictsRefreshBand returns a boolean if a field has been set.

### SetCollectionContradictsRefreshBandNil

`func (o *RwdSourcesRow) SetCollectionContradictsRefreshBandNil(b bool)`

 SetCollectionContradictsRefreshBandNil sets the value for CollectionContradictsRefreshBand to be an explicit nil

### UnsetCollectionContradictsRefreshBand
`func (o *RwdSourcesRow) UnsetCollectionContradictsRefreshBand()`

UnsetCollectionContradictsRefreshBand ensures that no value is present for CollectionContradictsRefreshBand, not even an explicit nil
### GetPrescriptionsOfMedicines

`func (o *RwdSourcesRow) GetPrescriptionsOfMedicines() bool`

GetPrescriptionsOfMedicines returns the PrescriptionsOfMedicines field if non-nil, zero value otherwise.

### GetPrescriptionsOfMedicinesOk

`func (o *RwdSourcesRow) GetPrescriptionsOfMedicinesOk() (*bool, bool)`

GetPrescriptionsOfMedicinesOk returns a tuple with the PrescriptionsOfMedicines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescriptionsOfMedicines

`func (o *RwdSourcesRow) SetPrescriptionsOfMedicines(v bool)`

SetPrescriptionsOfMedicines sets PrescriptionsOfMedicines field to given value.

### HasPrescriptionsOfMedicines

`func (o *RwdSourcesRow) HasPrescriptionsOfMedicines() bool`

HasPrescriptionsOfMedicines returns a boolean if a field has been set.

### SetPrescriptionsOfMedicinesNil

`func (o *RwdSourcesRow) SetPrescriptionsOfMedicinesNil(b bool)`

 SetPrescriptionsOfMedicinesNil sets the value for PrescriptionsOfMedicines to be an explicit nil

### UnsetPrescriptionsOfMedicines
`func (o *RwdSourcesRow) UnsetPrescriptionsOfMedicines()`

UnsetPrescriptionsOfMedicines ensures that no value is present for PrescriptionsOfMedicines, not even an explicit nil
### GetPrescriptionsVocabulary

`func (o *RwdSourcesRow) GetPrescriptionsVocabulary() string`

GetPrescriptionsVocabulary returns the PrescriptionsVocabulary field if non-nil, zero value otherwise.

### GetPrescriptionsVocabularyOk

`func (o *RwdSourcesRow) GetPrescriptionsVocabularyOk() (*string, bool)`

GetPrescriptionsVocabularyOk returns a tuple with the PrescriptionsVocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescriptionsVocabulary

`func (o *RwdSourcesRow) SetPrescriptionsVocabulary(v string)`

SetPrescriptionsVocabulary sets PrescriptionsVocabulary field to given value.

### HasPrescriptionsVocabulary

`func (o *RwdSourcesRow) HasPrescriptionsVocabulary() bool`

HasPrescriptionsVocabulary returns a boolean if a field has been set.

### SetPrescriptionsVocabularyNil

`func (o *RwdSourcesRow) SetPrescriptionsVocabularyNil(b bool)`

 SetPrescriptionsVocabularyNil sets the value for PrescriptionsVocabulary to be an explicit nil

### UnsetPrescriptionsVocabulary
`func (o *RwdSourcesRow) UnsetPrescriptionsVocabulary()`

UnsetPrescriptionsVocabulary ensures that no value is present for PrescriptionsVocabulary, not even an explicit nil
### GetPrescriptionsVocabularyOther

`func (o *RwdSourcesRow) GetPrescriptionsVocabularyOther() string`

GetPrescriptionsVocabularyOther returns the PrescriptionsVocabularyOther field if non-nil, zero value otherwise.

### GetPrescriptionsVocabularyOtherOk

`func (o *RwdSourcesRow) GetPrescriptionsVocabularyOtherOk() (*string, bool)`

GetPrescriptionsVocabularyOtherOk returns a tuple with the PrescriptionsVocabularyOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescriptionsVocabularyOther

`func (o *RwdSourcesRow) SetPrescriptionsVocabularyOther(v string)`

SetPrescriptionsVocabularyOther sets PrescriptionsVocabularyOther field to given value.

### HasPrescriptionsVocabularyOther

`func (o *RwdSourcesRow) HasPrescriptionsVocabularyOther() bool`

HasPrescriptionsVocabularyOther returns a boolean if a field has been set.

### SetPrescriptionsVocabularyOtherNil

`func (o *RwdSourcesRow) SetPrescriptionsVocabularyOtherNil(b bool)`

 SetPrescriptionsVocabularyOtherNil sets the value for PrescriptionsVocabularyOther to be an explicit nil

### UnsetPrescriptionsVocabularyOther
`func (o *RwdSourcesRow) UnsetPrescriptionsVocabularyOther()`

UnsetPrescriptionsVocabularyOther ensures that no value is present for PrescriptionsVocabularyOther, not even an explicit nil
### GetDispensingOfMedicines

`func (o *RwdSourcesRow) GetDispensingOfMedicines() bool`

GetDispensingOfMedicines returns the DispensingOfMedicines field if non-nil, zero value otherwise.

### GetDispensingOfMedicinesOk

`func (o *RwdSourcesRow) GetDispensingOfMedicinesOk() (*bool, bool)`

GetDispensingOfMedicinesOk returns a tuple with the DispensingOfMedicines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispensingOfMedicines

`func (o *RwdSourcesRow) SetDispensingOfMedicines(v bool)`

SetDispensingOfMedicines sets DispensingOfMedicines field to given value.

### HasDispensingOfMedicines

`func (o *RwdSourcesRow) HasDispensingOfMedicines() bool`

HasDispensingOfMedicines returns a boolean if a field has been set.

### SetDispensingOfMedicinesNil

`func (o *RwdSourcesRow) SetDispensingOfMedicinesNil(b bool)`

 SetDispensingOfMedicinesNil sets the value for DispensingOfMedicines to be an explicit nil

### UnsetDispensingOfMedicines
`func (o *RwdSourcesRow) UnsetDispensingOfMedicines()`

UnsetDispensingOfMedicines ensures that no value is present for DispensingOfMedicines, not even an explicit nil
### GetDispensingVocabulary

`func (o *RwdSourcesRow) GetDispensingVocabulary() string`

GetDispensingVocabulary returns the DispensingVocabulary field if non-nil, zero value otherwise.

### GetDispensingVocabularyOk

`func (o *RwdSourcesRow) GetDispensingVocabularyOk() (*string, bool)`

GetDispensingVocabularyOk returns a tuple with the DispensingVocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispensingVocabulary

`func (o *RwdSourcesRow) SetDispensingVocabulary(v string)`

SetDispensingVocabulary sets DispensingVocabulary field to given value.

### HasDispensingVocabulary

`func (o *RwdSourcesRow) HasDispensingVocabulary() bool`

HasDispensingVocabulary returns a boolean if a field has been set.

### SetDispensingVocabularyNil

`func (o *RwdSourcesRow) SetDispensingVocabularyNil(b bool)`

 SetDispensingVocabularyNil sets the value for DispensingVocabulary to be an explicit nil

### UnsetDispensingVocabulary
`func (o *RwdSourcesRow) UnsetDispensingVocabulary()`

UnsetDispensingVocabulary ensures that no value is present for DispensingVocabulary, not even an explicit nil
### GetDispensingVocabularyOther

`func (o *RwdSourcesRow) GetDispensingVocabularyOther() string`

GetDispensingVocabularyOther returns the DispensingVocabularyOther field if non-nil, zero value otherwise.

### GetDispensingVocabularyOtherOk

`func (o *RwdSourcesRow) GetDispensingVocabularyOtherOk() (*string, bool)`

GetDispensingVocabularyOtherOk returns a tuple with the DispensingVocabularyOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispensingVocabularyOther

`func (o *RwdSourcesRow) SetDispensingVocabularyOther(v string)`

SetDispensingVocabularyOther sets DispensingVocabularyOther field to given value.

### HasDispensingVocabularyOther

`func (o *RwdSourcesRow) HasDispensingVocabularyOther() bool`

HasDispensingVocabularyOther returns a boolean if a field has been set.

### SetDispensingVocabularyOtherNil

`func (o *RwdSourcesRow) SetDispensingVocabularyOtherNil(b bool)`

 SetDispensingVocabularyOtherNil sets the value for DispensingVocabularyOther to be an explicit nil

### UnsetDispensingVocabularyOther
`func (o *RwdSourcesRow) UnsetDispensingVocabularyOther()`

UnsetDispensingVocabularyOther ensures that no value is present for DispensingVocabularyOther, not even an explicit nil
### GetMedicinalProductVocabulary

`func (o *RwdSourcesRow) GetMedicinalProductVocabulary() string`

GetMedicinalProductVocabulary returns the MedicinalProductVocabulary field if non-nil, zero value otherwise.

### GetMedicinalProductVocabularyOk

`func (o *RwdSourcesRow) GetMedicinalProductVocabularyOk() (*string, bool)`

GetMedicinalProductVocabularyOk returns a tuple with the MedicinalProductVocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicinalProductVocabulary

`func (o *RwdSourcesRow) SetMedicinalProductVocabulary(v string)`

SetMedicinalProductVocabulary sets MedicinalProductVocabulary field to given value.

### HasMedicinalProductVocabulary

`func (o *RwdSourcesRow) HasMedicinalProductVocabulary() bool`

HasMedicinalProductVocabulary returns a boolean if a field has been set.

### SetMedicinalProductVocabularyNil

`func (o *RwdSourcesRow) SetMedicinalProductVocabularyNil(b bool)`

 SetMedicinalProductVocabularyNil sets the value for MedicinalProductVocabulary to be an explicit nil

### UnsetMedicinalProductVocabulary
`func (o *RwdSourcesRow) UnsetMedicinalProductVocabulary()`

UnsetMedicinalProductVocabulary ensures that no value is present for MedicinalProductVocabulary, not even an explicit nil
### GetMedicinalProductVocabularyOther

`func (o *RwdSourcesRow) GetMedicinalProductVocabularyOther() string`

GetMedicinalProductVocabularyOther returns the MedicinalProductVocabularyOther field if non-nil, zero value otherwise.

### GetMedicinalProductVocabularyOtherOk

`func (o *RwdSourcesRow) GetMedicinalProductVocabularyOtherOk() (*string, bool)`

GetMedicinalProductVocabularyOtherOk returns a tuple with the MedicinalProductVocabularyOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicinalProductVocabularyOther

`func (o *RwdSourcesRow) SetMedicinalProductVocabularyOther(v string)`

SetMedicinalProductVocabularyOther sets MedicinalProductVocabularyOther field to given value.

### HasMedicinalProductVocabularyOther

`func (o *RwdSourcesRow) HasMedicinalProductVocabularyOther() bool`

HasMedicinalProductVocabularyOther returns a boolean if a field has been set.

### SetMedicinalProductVocabularyOtherNil

`func (o *RwdSourcesRow) SetMedicinalProductVocabularyOtherNil(b bool)`

 SetMedicinalProductVocabularyOtherNil sets the value for MedicinalProductVocabularyOther to be an explicit nil

### UnsetMedicinalProductVocabularyOther
`func (o *RwdSourcesRow) UnsetMedicinalProductVocabularyOther()`

UnsetMedicinalProductVocabularyOther ensures that no value is present for MedicinalProductVocabularyOther, not even an explicit nil
### GetIndicationForUse

`func (o *RwdSourcesRow) GetIndicationForUse() bool`

GetIndicationForUse returns the IndicationForUse field if non-nil, zero value otherwise.

### GetIndicationForUseOk

`func (o *RwdSourcesRow) GetIndicationForUseOk() (*bool, bool)`

GetIndicationForUseOk returns a tuple with the IndicationForUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicationForUse

`func (o *RwdSourcesRow) SetIndicationForUse(v bool)`

SetIndicationForUse sets IndicationForUse field to given value.

### HasIndicationForUse

`func (o *RwdSourcesRow) HasIndicationForUse() bool`

HasIndicationForUse returns a boolean if a field has been set.

### SetIndicationForUseNil

`func (o *RwdSourcesRow) SetIndicationForUseNil(b bool)`

 SetIndicationForUseNil sets the value for IndicationForUse to be an explicit nil

### UnsetIndicationForUse
`func (o *RwdSourcesRow) UnsetIndicationForUse()`

UnsetIndicationForUse ensures that no value is present for IndicationForUse, not even an explicit nil
### GetIndicationVocabulary

`func (o *RwdSourcesRow) GetIndicationVocabulary() string`

GetIndicationVocabulary returns the IndicationVocabulary field if non-nil, zero value otherwise.

### GetIndicationVocabularyOk

`func (o *RwdSourcesRow) GetIndicationVocabularyOk() (*string, bool)`

GetIndicationVocabularyOk returns a tuple with the IndicationVocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicationVocabulary

`func (o *RwdSourcesRow) SetIndicationVocabulary(v string)`

SetIndicationVocabulary sets IndicationVocabulary field to given value.

### HasIndicationVocabulary

`func (o *RwdSourcesRow) HasIndicationVocabulary() bool`

HasIndicationVocabulary returns a boolean if a field has been set.

### SetIndicationVocabularyNil

`func (o *RwdSourcesRow) SetIndicationVocabularyNil(b bool)`

 SetIndicationVocabularyNil sets the value for IndicationVocabulary to be an explicit nil

### UnsetIndicationVocabulary
`func (o *RwdSourcesRow) UnsetIndicationVocabulary()`

UnsetIndicationVocabulary ensures that no value is present for IndicationVocabulary, not even an explicit nil
### GetIndicationVocabularyOther

`func (o *RwdSourcesRow) GetIndicationVocabularyOther() string`

GetIndicationVocabularyOther returns the IndicationVocabularyOther field if non-nil, zero value otherwise.

### GetIndicationVocabularyOtherOk

`func (o *RwdSourcesRow) GetIndicationVocabularyOtherOk() (*string, bool)`

GetIndicationVocabularyOtherOk returns a tuple with the IndicationVocabularyOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicationVocabularyOther

`func (o *RwdSourcesRow) SetIndicationVocabularyOther(v string)`

SetIndicationVocabularyOther sets IndicationVocabularyOther field to given value.

### HasIndicationVocabularyOther

`func (o *RwdSourcesRow) HasIndicationVocabularyOther() bool`

HasIndicationVocabularyOther returns a boolean if a field has been set.

### SetIndicationVocabularyOtherNil

`func (o *RwdSourcesRow) SetIndicationVocabularyOtherNil(b bool)`

 SetIndicationVocabularyOtherNil sets the value for IndicationVocabularyOther to be an explicit nil

### UnsetIndicationVocabularyOther
`func (o *RwdSourcesRow) UnsetIndicationVocabularyOther()`

UnsetIndicationVocabularyOther ensures that no value is present for IndicationVocabularyOther, not even an explicit nil
### GetDiagnosticCodes

`func (o *RwdSourcesRow) GetDiagnosticCodes() bool`

GetDiagnosticCodes returns the DiagnosticCodes field if non-nil, zero value otherwise.

### GetDiagnosticCodesOk

`func (o *RwdSourcesRow) GetDiagnosticCodesOk() (*bool, bool)`

GetDiagnosticCodesOk returns a tuple with the DiagnosticCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticCodes

`func (o *RwdSourcesRow) SetDiagnosticCodes(v bool)`

SetDiagnosticCodes sets DiagnosticCodes field to given value.

### HasDiagnosticCodes

`func (o *RwdSourcesRow) HasDiagnosticCodes() bool`

HasDiagnosticCodes returns a boolean if a field has been set.

### SetDiagnosticCodesNil

`func (o *RwdSourcesRow) SetDiagnosticCodesNil(b bool)`

 SetDiagnosticCodesNil sets the value for DiagnosticCodes to be an explicit nil

### UnsetDiagnosticCodes
`func (o *RwdSourcesRow) UnsetDiagnosticCodes()`

UnsetDiagnosticCodes ensures that no value is present for DiagnosticCodes, not even an explicit nil
### GetDiagnosisMedicalEventVocabulary

`func (o *RwdSourcesRow) GetDiagnosisMedicalEventVocabulary() string`

GetDiagnosisMedicalEventVocabulary returns the DiagnosisMedicalEventVocabulary field if non-nil, zero value otherwise.

### GetDiagnosisMedicalEventVocabularyOk

`func (o *RwdSourcesRow) GetDiagnosisMedicalEventVocabularyOk() (*string, bool)`

GetDiagnosisMedicalEventVocabularyOk returns a tuple with the DiagnosisMedicalEventVocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosisMedicalEventVocabulary

`func (o *RwdSourcesRow) SetDiagnosisMedicalEventVocabulary(v string)`

SetDiagnosisMedicalEventVocabulary sets DiagnosisMedicalEventVocabulary field to given value.

### HasDiagnosisMedicalEventVocabulary

`func (o *RwdSourcesRow) HasDiagnosisMedicalEventVocabulary() bool`

HasDiagnosisMedicalEventVocabulary returns a boolean if a field has been set.

### SetDiagnosisMedicalEventVocabularyNil

`func (o *RwdSourcesRow) SetDiagnosisMedicalEventVocabularyNil(b bool)`

 SetDiagnosisMedicalEventVocabularyNil sets the value for DiagnosisMedicalEventVocabulary to be an explicit nil

### UnsetDiagnosisMedicalEventVocabulary
`func (o *RwdSourcesRow) UnsetDiagnosisMedicalEventVocabulary()`

UnsetDiagnosisMedicalEventVocabulary ensures that no value is present for DiagnosisMedicalEventVocabulary, not even an explicit nil
### GetDiagnosisMedicalEventVocabularyOther

`func (o *RwdSourcesRow) GetDiagnosisMedicalEventVocabularyOther() string`

GetDiagnosisMedicalEventVocabularyOther returns the DiagnosisMedicalEventVocabularyOther field if non-nil, zero value otherwise.

### GetDiagnosisMedicalEventVocabularyOtherOk

`func (o *RwdSourcesRow) GetDiagnosisMedicalEventVocabularyOtherOk() (*string, bool)`

GetDiagnosisMedicalEventVocabularyOtherOk returns a tuple with the DiagnosisMedicalEventVocabularyOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosisMedicalEventVocabularyOther

`func (o *RwdSourcesRow) SetDiagnosisMedicalEventVocabularyOther(v string)`

SetDiagnosisMedicalEventVocabularyOther sets DiagnosisMedicalEventVocabularyOther field to given value.

### HasDiagnosisMedicalEventVocabularyOther

`func (o *RwdSourcesRow) HasDiagnosisMedicalEventVocabularyOther() bool`

HasDiagnosisMedicalEventVocabularyOther returns a boolean if a field has been set.

### SetDiagnosisMedicalEventVocabularyOtherNil

`func (o *RwdSourcesRow) SetDiagnosisMedicalEventVocabularyOtherNil(b bool)`

 SetDiagnosisMedicalEventVocabularyOtherNil sets the value for DiagnosisMedicalEventVocabularyOther to be an explicit nil

### UnsetDiagnosisMedicalEventVocabularyOther
`func (o *RwdSourcesRow) UnsetDiagnosisMedicalEventVocabularyOther()`

UnsetDiagnosisMedicalEventVocabularyOther ensures that no value is present for DiagnosisMedicalEventVocabularyOther, not even an explicit nil
### GetHospitalAdmissionAndOrDischarge

`func (o *RwdSourcesRow) GetHospitalAdmissionAndOrDischarge() bool`

GetHospitalAdmissionAndOrDischarge returns the HospitalAdmissionAndOrDischarge field if non-nil, zero value otherwise.

### GetHospitalAdmissionAndOrDischargeOk

`func (o *RwdSourcesRow) GetHospitalAdmissionAndOrDischargeOk() (*bool, bool)`

GetHospitalAdmissionAndOrDischargeOk returns a tuple with the HospitalAdmissionAndOrDischarge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHospitalAdmissionAndOrDischarge

`func (o *RwdSourcesRow) SetHospitalAdmissionAndOrDischarge(v bool)`

SetHospitalAdmissionAndOrDischarge sets HospitalAdmissionAndOrDischarge field to given value.

### HasHospitalAdmissionAndOrDischarge

`func (o *RwdSourcesRow) HasHospitalAdmissionAndOrDischarge() bool`

HasHospitalAdmissionAndOrDischarge returns a boolean if a field has been set.

### SetHospitalAdmissionAndOrDischargeNil

`func (o *RwdSourcesRow) SetHospitalAdmissionAndOrDischargeNil(b bool)`

 SetHospitalAdmissionAndOrDischargeNil sets the value for HospitalAdmissionAndOrDischarge to be an explicit nil

### UnsetHospitalAdmissionAndOrDischarge
`func (o *RwdSourcesRow) UnsetHospitalAdmissionAndOrDischarge()`

UnsetHospitalAdmissionAndOrDischarge ensures that no value is present for HospitalAdmissionAndOrDischarge, not even an explicit nil
### GetIcuAdmission

`func (o *RwdSourcesRow) GetIcuAdmission() bool`

GetIcuAdmission returns the IcuAdmission field if non-nil, zero value otherwise.

### GetIcuAdmissionOk

`func (o *RwdSourcesRow) GetIcuAdmissionOk() (*bool, bool)`

GetIcuAdmissionOk returns a tuple with the IcuAdmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcuAdmission

`func (o *RwdSourcesRow) SetIcuAdmission(v bool)`

SetIcuAdmission sets IcuAdmission field to given value.

### HasIcuAdmission

`func (o *RwdSourcesRow) HasIcuAdmission() bool`

HasIcuAdmission returns a boolean if a field has been set.

### SetIcuAdmissionNil

`func (o *RwdSourcesRow) SetIcuAdmissionNil(b bool)`

 SetIcuAdmissionNil sets the value for IcuAdmission to be an explicit nil

### UnsetIcuAdmission
`func (o *RwdSourcesRow) UnsetIcuAdmission()`

UnsetIcuAdmission ensures that no value is present for IcuAdmission, not even an explicit nil
### GetCauseOfDeath

`func (o *RwdSourcesRow) GetCauseOfDeath() bool`

GetCauseOfDeath returns the CauseOfDeath field if non-nil, zero value otherwise.

### GetCauseOfDeathOk

`func (o *RwdSourcesRow) GetCauseOfDeathOk() (*bool, bool)`

GetCauseOfDeathOk returns a tuple with the CauseOfDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseOfDeath

`func (o *RwdSourcesRow) SetCauseOfDeath(v bool)`

SetCauseOfDeath sets CauseOfDeath field to given value.

### HasCauseOfDeath

`func (o *RwdSourcesRow) HasCauseOfDeath() bool`

HasCauseOfDeath returns a boolean if a field has been set.

### SetCauseOfDeathNil

`func (o *RwdSourcesRow) SetCauseOfDeathNil(b bool)`

 SetCauseOfDeathNil sets the value for CauseOfDeath to be an explicit nil

### UnsetCauseOfDeath
`func (o *RwdSourcesRow) UnsetCauseOfDeath()`

UnsetCauseOfDeath ensures that no value is present for CauseOfDeath, not even an explicit nil
### GetCauseOfDeathVocabulary

`func (o *RwdSourcesRow) GetCauseOfDeathVocabulary() string`

GetCauseOfDeathVocabulary returns the CauseOfDeathVocabulary field if non-nil, zero value otherwise.

### GetCauseOfDeathVocabularyOk

`func (o *RwdSourcesRow) GetCauseOfDeathVocabularyOk() (*string, bool)`

GetCauseOfDeathVocabularyOk returns a tuple with the CauseOfDeathVocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseOfDeathVocabulary

`func (o *RwdSourcesRow) SetCauseOfDeathVocabulary(v string)`

SetCauseOfDeathVocabulary sets CauseOfDeathVocabulary field to given value.

### HasCauseOfDeathVocabulary

`func (o *RwdSourcesRow) HasCauseOfDeathVocabulary() bool`

HasCauseOfDeathVocabulary returns a boolean if a field has been set.

### SetCauseOfDeathVocabularyNil

`func (o *RwdSourcesRow) SetCauseOfDeathVocabularyNil(b bool)`

 SetCauseOfDeathVocabularyNil sets the value for CauseOfDeathVocabulary to be an explicit nil

### UnsetCauseOfDeathVocabulary
`func (o *RwdSourcesRow) UnsetCauseOfDeathVocabulary()`

UnsetCauseOfDeathVocabulary ensures that no value is present for CauseOfDeathVocabulary, not even an explicit nil
### GetCauseOfDeathVocabularyOther

`func (o *RwdSourcesRow) GetCauseOfDeathVocabularyOther() string`

GetCauseOfDeathVocabularyOther returns the CauseOfDeathVocabularyOther field if non-nil, zero value otherwise.

### GetCauseOfDeathVocabularyOtherOk

`func (o *RwdSourcesRow) GetCauseOfDeathVocabularyOtherOk() (*string, bool)`

GetCauseOfDeathVocabularyOtherOk returns a tuple with the CauseOfDeathVocabularyOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseOfDeathVocabularyOther

`func (o *RwdSourcesRow) SetCauseOfDeathVocabularyOther(v string)`

SetCauseOfDeathVocabularyOther sets CauseOfDeathVocabularyOther field to given value.

### HasCauseOfDeathVocabularyOther

`func (o *RwdSourcesRow) HasCauseOfDeathVocabularyOther() bool`

HasCauseOfDeathVocabularyOther returns a boolean if a field has been set.

### SetCauseOfDeathVocabularyOtherNil

`func (o *RwdSourcesRow) SetCauseOfDeathVocabularyOtherNil(b bool)`

 SetCauseOfDeathVocabularyOtherNil sets the value for CauseOfDeathVocabularyOther to be an explicit nil

### UnsetCauseOfDeathVocabularyOther
`func (o *RwdSourcesRow) UnsetCauseOfDeathVocabularyOther()`

UnsetCauseOfDeathVocabularyOther ensures that no value is present for CauseOfDeathVocabularyOther, not even an explicit nil
### GetDiseaseInformation

`func (o *RwdSourcesRow) GetDiseaseInformation() bool`

GetDiseaseInformation returns the DiseaseInformation field if non-nil, zero value otherwise.

### GetDiseaseInformationOk

`func (o *RwdSourcesRow) GetDiseaseInformationOk() (*bool, bool)`

GetDiseaseInformationOk returns a tuple with the DiseaseInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiseaseInformation

`func (o *RwdSourcesRow) SetDiseaseInformation(v bool)`

SetDiseaseInformation sets DiseaseInformation field to given value.

### HasDiseaseInformation

`func (o *RwdSourcesRow) HasDiseaseInformation() bool`

HasDiseaseInformation returns a boolean if a field has been set.

### SetDiseaseInformationNil

`func (o *RwdSourcesRow) SetDiseaseInformationNil(b bool)`

 SetDiseaseInformationNil sets the value for DiseaseInformation to be an explicit nil

### UnsetDiseaseInformation
`func (o *RwdSourcesRow) UnsetDiseaseInformation()`

UnsetDiseaseInformation ensures that no value is present for DiseaseInformation, not even an explicit nil
### GetRareDiseases

`func (o *RwdSourcesRow) GetRareDiseases() bool`

GetRareDiseases returns the RareDiseases field if non-nil, zero value otherwise.

### GetRareDiseasesOk

`func (o *RwdSourcesRow) GetRareDiseasesOk() (*bool, bool)`

GetRareDiseasesOk returns a tuple with the RareDiseases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRareDiseases

`func (o *RwdSourcesRow) SetRareDiseases(v bool)`

SetRareDiseases sets RareDiseases field to given value.

### HasRareDiseases

`func (o *RwdSourcesRow) HasRareDiseases() bool`

HasRareDiseases returns a boolean if a field has been set.

### SetRareDiseasesNil

`func (o *RwdSourcesRow) SetRareDiseasesNil(b bool)`

 SetRareDiseasesNil sets the value for RareDiseases to be an explicit nil

### UnsetRareDiseases
`func (o *RwdSourcesRow) UnsetRareDiseases()`

UnsetRareDiseases ensures that no value is present for RareDiseases, not even an explicit nil
### GetPregnancyAndOrNeonates

`func (o *RwdSourcesRow) GetPregnancyAndOrNeonates() bool`

GetPregnancyAndOrNeonates returns the PregnancyAndOrNeonates field if non-nil, zero value otherwise.

### GetPregnancyAndOrNeonatesOk

`func (o *RwdSourcesRow) GetPregnancyAndOrNeonatesOk() (*bool, bool)`

GetPregnancyAndOrNeonatesOk returns a tuple with the PregnancyAndOrNeonates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPregnancyAndOrNeonates

`func (o *RwdSourcesRow) SetPregnancyAndOrNeonates(v bool)`

SetPregnancyAndOrNeonates sets PregnancyAndOrNeonates field to given value.

### HasPregnancyAndOrNeonates

`func (o *RwdSourcesRow) HasPregnancyAndOrNeonates() bool`

HasPregnancyAndOrNeonates returns a boolean if a field has been set.

### SetPregnancyAndOrNeonatesNil

`func (o *RwdSourcesRow) SetPregnancyAndOrNeonatesNil(b bool)`

 SetPregnancyAndOrNeonatesNil sets the value for PregnancyAndOrNeonates to be an explicit nil

### UnsetPregnancyAndOrNeonates
`func (o *RwdSourcesRow) UnsetPregnancyAndOrNeonates()`

UnsetPregnancyAndOrNeonates ensures that no value is present for PregnancyAndOrNeonates, not even an explicit nil
### GetAdvanceTherapyMedicinalProductsAtmp

`func (o *RwdSourcesRow) GetAdvanceTherapyMedicinalProductsAtmp() bool`

GetAdvanceTherapyMedicinalProductsAtmp returns the AdvanceTherapyMedicinalProductsAtmp field if non-nil, zero value otherwise.

### GetAdvanceTherapyMedicinalProductsAtmpOk

`func (o *RwdSourcesRow) GetAdvanceTherapyMedicinalProductsAtmpOk() (*bool, bool)`

GetAdvanceTherapyMedicinalProductsAtmpOk returns a tuple with the AdvanceTherapyMedicinalProductsAtmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceTherapyMedicinalProductsAtmp

`func (o *RwdSourcesRow) SetAdvanceTherapyMedicinalProductsAtmp(v bool)`

SetAdvanceTherapyMedicinalProductsAtmp sets AdvanceTherapyMedicinalProductsAtmp field to given value.

### HasAdvanceTherapyMedicinalProductsAtmp

`func (o *RwdSourcesRow) HasAdvanceTherapyMedicinalProductsAtmp() bool`

HasAdvanceTherapyMedicinalProductsAtmp returns a boolean if a field has been set.

### SetAdvanceTherapyMedicinalProductsAtmpNil

`func (o *RwdSourcesRow) SetAdvanceTherapyMedicinalProductsAtmpNil(b bool)`

 SetAdvanceTherapyMedicinalProductsAtmpNil sets the value for AdvanceTherapyMedicinalProductsAtmp to be an explicit nil

### UnsetAdvanceTherapyMedicinalProductsAtmp
`func (o *RwdSourcesRow) UnsetAdvanceTherapyMedicinalProductsAtmp()`

UnsetAdvanceTherapyMedicinalProductsAtmp ensures that no value is present for AdvanceTherapyMedicinalProductsAtmp, not even an explicit nil
### GetContraception

`func (o *RwdSourcesRow) GetContraception() bool`

GetContraception returns the Contraception field if non-nil, zero value otherwise.

### GetContraceptionOk

`func (o *RwdSourcesRow) GetContraceptionOk() (*bool, bool)`

GetContraceptionOk returns a tuple with the Contraception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContraception

`func (o *RwdSourcesRow) SetContraception(v bool)`

SetContraception sets Contraception field to given value.

### HasContraception

`func (o *RwdSourcesRow) HasContraception() bool`

HasContraception returns a boolean if a field has been set.

### SetContraceptionNil

`func (o *RwdSourcesRow) SetContraceptionNil(b bool)`

 SetContraceptionNil sets the value for Contraception to be an explicit nil

### UnsetContraception
`func (o *RwdSourcesRow) UnsetContraception()`

UnsetContraception ensures that no value is present for Contraception, not even an explicit nil
### GetMedicalDevices

`func (o *RwdSourcesRow) GetMedicalDevices() bool`

GetMedicalDevices returns the MedicalDevices field if non-nil, zero value otherwise.

### GetMedicalDevicesOk

`func (o *RwdSourcesRow) GetMedicalDevicesOk() (*bool, bool)`

GetMedicalDevicesOk returns a tuple with the MedicalDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicalDevices

`func (o *RwdSourcesRow) SetMedicalDevices(v bool)`

SetMedicalDevices sets MedicalDevices field to given value.

### HasMedicalDevices

`func (o *RwdSourcesRow) HasMedicalDevices() bool`

HasMedicalDevices returns a boolean if a field has been set.

### SetMedicalDevicesNil

`func (o *RwdSourcesRow) SetMedicalDevicesNil(b bool)`

 SetMedicalDevicesNil sets the value for MedicalDevices to be an explicit nil

### UnsetMedicalDevices
`func (o *RwdSourcesRow) UnsetMedicalDevices()`

UnsetMedicalDevices ensures that no value is present for MedicalDevices, not even an explicit nil
### GetAdministrationOfVaccines

`func (o *RwdSourcesRow) GetAdministrationOfVaccines() bool`

GetAdministrationOfVaccines returns the AdministrationOfVaccines field if non-nil, zero value otherwise.

### GetAdministrationOfVaccinesOk

`func (o *RwdSourcesRow) GetAdministrationOfVaccinesOk() (*bool, bool)`

GetAdministrationOfVaccinesOk returns a tuple with the AdministrationOfVaccines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrationOfVaccines

`func (o *RwdSourcesRow) SetAdministrationOfVaccines(v bool)`

SetAdministrationOfVaccines sets AdministrationOfVaccines field to given value.

### HasAdministrationOfVaccines

`func (o *RwdSourcesRow) HasAdministrationOfVaccines() bool`

HasAdministrationOfVaccines returns a boolean if a field has been set.

### SetAdministrationOfVaccinesNil

`func (o *RwdSourcesRow) SetAdministrationOfVaccinesNil(b bool)`

 SetAdministrationOfVaccinesNil sets the value for AdministrationOfVaccines to be an explicit nil

### UnsetAdministrationOfVaccines
`func (o *RwdSourcesRow) UnsetAdministrationOfVaccines()`

UnsetAdministrationOfVaccines ensures that no value is present for AdministrationOfVaccines, not even an explicit nil
### GetProcedures

`func (o *RwdSourcesRow) GetProcedures() bool`

GetProcedures returns the Procedures field if non-nil, zero value otherwise.

### GetProceduresOk

`func (o *RwdSourcesRow) GetProceduresOk() (*bool, bool)`

GetProceduresOk returns a tuple with the Procedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedures

`func (o *RwdSourcesRow) SetProcedures(v bool)`

SetProcedures sets Procedures field to given value.

### HasProcedures

`func (o *RwdSourcesRow) HasProcedures() bool`

HasProcedures returns a boolean if a field has been set.

### SetProceduresNil

`func (o *RwdSourcesRow) SetProceduresNil(b bool)`

 SetProceduresNil sets the value for Procedures to be an explicit nil

### UnsetProcedures
`func (o *RwdSourcesRow) UnsetProcedures()`

UnsetProcedures ensures that no value is present for Procedures, not even an explicit nil
### GetProceduresVocabulary

`func (o *RwdSourcesRow) GetProceduresVocabulary() string`

GetProceduresVocabulary returns the ProceduresVocabulary field if non-nil, zero value otherwise.

### GetProceduresVocabularyOk

`func (o *RwdSourcesRow) GetProceduresVocabularyOk() (*string, bool)`

GetProceduresVocabularyOk returns a tuple with the ProceduresVocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProceduresVocabulary

`func (o *RwdSourcesRow) SetProceduresVocabulary(v string)`

SetProceduresVocabulary sets ProceduresVocabulary field to given value.

### HasProceduresVocabulary

`func (o *RwdSourcesRow) HasProceduresVocabulary() bool`

HasProceduresVocabulary returns a boolean if a field has been set.

### SetProceduresVocabularyNil

`func (o *RwdSourcesRow) SetProceduresVocabularyNil(b bool)`

 SetProceduresVocabularyNil sets the value for ProceduresVocabulary to be an explicit nil

### UnsetProceduresVocabulary
`func (o *RwdSourcesRow) UnsetProceduresVocabulary()`

UnsetProceduresVocabulary ensures that no value is present for ProceduresVocabulary, not even an explicit nil
### GetProceduresVocabularyOther

`func (o *RwdSourcesRow) GetProceduresVocabularyOther() string`

GetProceduresVocabularyOther returns the ProceduresVocabularyOther field if non-nil, zero value otherwise.

### GetProceduresVocabularyOtherOk

`func (o *RwdSourcesRow) GetProceduresVocabularyOtherOk() (*string, bool)`

GetProceduresVocabularyOtherOk returns a tuple with the ProceduresVocabularyOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProceduresVocabularyOther

`func (o *RwdSourcesRow) SetProceduresVocabularyOther(v string)`

SetProceduresVocabularyOther sets ProceduresVocabularyOther field to given value.

### HasProceduresVocabularyOther

`func (o *RwdSourcesRow) HasProceduresVocabularyOther() bool`

HasProceduresVocabularyOther returns a boolean if a field has been set.

### SetProceduresVocabularyOtherNil

`func (o *RwdSourcesRow) SetProceduresVocabularyOtherNil(b bool)`

 SetProceduresVocabularyOtherNil sets the value for ProceduresVocabularyOther to be an explicit nil

### UnsetProceduresVocabularyOther
`func (o *RwdSourcesRow) UnsetProceduresVocabularyOther()`

UnsetProceduresVocabularyOther ensures that no value is present for ProceduresVocabularyOther, not even an explicit nil
### GetClinicalMeasurements

`func (o *RwdSourcesRow) GetClinicalMeasurements() bool`

GetClinicalMeasurements returns the ClinicalMeasurements field if non-nil, zero value otherwise.

### GetClinicalMeasurementsOk

`func (o *RwdSourcesRow) GetClinicalMeasurementsOk() (*bool, bool)`

GetClinicalMeasurementsOk returns a tuple with the ClinicalMeasurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinicalMeasurements

`func (o *RwdSourcesRow) SetClinicalMeasurements(v bool)`

SetClinicalMeasurements sets ClinicalMeasurements field to given value.

### HasClinicalMeasurements

`func (o *RwdSourcesRow) HasClinicalMeasurements() bool`

HasClinicalMeasurements returns a boolean if a field has been set.

### SetClinicalMeasurementsNil

`func (o *RwdSourcesRow) SetClinicalMeasurementsNil(b bool)`

 SetClinicalMeasurementsNil sets the value for ClinicalMeasurements to be an explicit nil

### UnsetClinicalMeasurements
`func (o *RwdSourcesRow) UnsetClinicalMeasurements()`

UnsetClinicalMeasurements ensures that no value is present for ClinicalMeasurements, not even an explicit nil
### GetHealthcareProvider

`func (o *RwdSourcesRow) GetHealthcareProvider() bool`

GetHealthcareProvider returns the HealthcareProvider field if non-nil, zero value otherwise.

### GetHealthcareProviderOk

`func (o *RwdSourcesRow) GetHealthcareProviderOk() (*bool, bool)`

GetHealthcareProviderOk returns a tuple with the HealthcareProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcareProvider

`func (o *RwdSourcesRow) SetHealthcareProvider(v bool)`

SetHealthcareProvider sets HealthcareProvider field to given value.

### HasHealthcareProvider

`func (o *RwdSourcesRow) HasHealthcareProvider() bool`

HasHealthcareProvider returns a boolean if a field has been set.

### SetHealthcareProviderNil

`func (o *RwdSourcesRow) SetHealthcareProviderNil(b bool)`

 SetHealthcareProviderNil sets the value for HealthcareProvider to be an explicit nil

### UnsetHealthcareProvider
`func (o *RwdSourcesRow) UnsetHealthcareProvider()`

UnsetHealthcareProvider ensures that no value is present for HealthcareProvider, not even an explicit nil
### GetGeneticData

`func (o *RwdSourcesRow) GetGeneticData() bool`

GetGeneticData returns the GeneticData field if non-nil, zero value otherwise.

### GetGeneticDataOk

`func (o *RwdSourcesRow) GetGeneticDataOk() (*bool, bool)`

GetGeneticDataOk returns a tuple with the GeneticData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneticData

`func (o *RwdSourcesRow) SetGeneticData(v bool)`

SetGeneticData sets GeneticData field to given value.

### HasGeneticData

`func (o *RwdSourcesRow) HasGeneticData() bool`

HasGeneticData returns a boolean if a field has been set.

### SetGeneticDataNil

`func (o *RwdSourcesRow) SetGeneticDataNil(b bool)`

 SetGeneticDataNil sets the value for GeneticData to be an explicit nil

### UnsetGeneticData
`func (o *RwdSourcesRow) UnsetGeneticData()`

UnsetGeneticData ensures that no value is present for GeneticData, not even an explicit nil
### GetGeneticDataVocabulary

`func (o *RwdSourcesRow) GetGeneticDataVocabulary() string`

GetGeneticDataVocabulary returns the GeneticDataVocabulary field if non-nil, zero value otherwise.

### GetGeneticDataVocabularyOk

`func (o *RwdSourcesRow) GetGeneticDataVocabularyOk() (*string, bool)`

GetGeneticDataVocabularyOk returns a tuple with the GeneticDataVocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneticDataVocabulary

`func (o *RwdSourcesRow) SetGeneticDataVocabulary(v string)`

SetGeneticDataVocabulary sets GeneticDataVocabulary field to given value.

### HasGeneticDataVocabulary

`func (o *RwdSourcesRow) HasGeneticDataVocabulary() bool`

HasGeneticDataVocabulary returns a boolean if a field has been set.

### SetGeneticDataVocabularyNil

`func (o *RwdSourcesRow) SetGeneticDataVocabularyNil(b bool)`

 SetGeneticDataVocabularyNil sets the value for GeneticDataVocabulary to be an explicit nil

### UnsetGeneticDataVocabulary
`func (o *RwdSourcesRow) UnsetGeneticDataVocabulary()`

UnsetGeneticDataVocabulary ensures that no value is present for GeneticDataVocabulary, not even an explicit nil
### GetGeneticDataVocabularyOther

`func (o *RwdSourcesRow) GetGeneticDataVocabularyOther() string`

GetGeneticDataVocabularyOther returns the GeneticDataVocabularyOther field if non-nil, zero value otherwise.

### GetGeneticDataVocabularyOtherOk

`func (o *RwdSourcesRow) GetGeneticDataVocabularyOtherOk() (*string, bool)`

GetGeneticDataVocabularyOtherOk returns a tuple with the GeneticDataVocabularyOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneticDataVocabularyOther

`func (o *RwdSourcesRow) SetGeneticDataVocabularyOther(v string)`

SetGeneticDataVocabularyOther sets GeneticDataVocabularyOther field to given value.

### HasGeneticDataVocabularyOther

`func (o *RwdSourcesRow) HasGeneticDataVocabularyOther() bool`

HasGeneticDataVocabularyOther returns a boolean if a field has been set.

### SetGeneticDataVocabularyOtherNil

`func (o *RwdSourcesRow) SetGeneticDataVocabularyOtherNil(b bool)`

 SetGeneticDataVocabularyOtherNil sets the value for GeneticDataVocabularyOther to be an explicit nil

### UnsetGeneticDataVocabularyOther
`func (o *RwdSourcesRow) UnsetGeneticDataVocabularyOther()`

UnsetGeneticDataVocabularyOther ensures that no value is present for GeneticDataVocabularyOther, not even an explicit nil
### GetBiomarkerData

`func (o *RwdSourcesRow) GetBiomarkerData() bool`

GetBiomarkerData returns the BiomarkerData field if non-nil, zero value otherwise.

### GetBiomarkerDataOk

`func (o *RwdSourcesRow) GetBiomarkerDataOk() (*bool, bool)`

GetBiomarkerDataOk returns a tuple with the BiomarkerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerData

`func (o *RwdSourcesRow) SetBiomarkerData(v bool)`

SetBiomarkerData sets BiomarkerData field to given value.

### HasBiomarkerData

`func (o *RwdSourcesRow) HasBiomarkerData() bool`

HasBiomarkerData returns a boolean if a field has been set.

### SetBiomarkerDataNil

`func (o *RwdSourcesRow) SetBiomarkerDataNil(b bool)`

 SetBiomarkerDataNil sets the value for BiomarkerData to be an explicit nil

### UnsetBiomarkerData
`func (o *RwdSourcesRow) UnsetBiomarkerData()`

UnsetBiomarkerData ensures that no value is present for BiomarkerData, not even an explicit nil
### GetBiomarkerDataVocabulary

`func (o *RwdSourcesRow) GetBiomarkerDataVocabulary() string`

GetBiomarkerDataVocabulary returns the BiomarkerDataVocabulary field if non-nil, zero value otherwise.

### GetBiomarkerDataVocabularyOk

`func (o *RwdSourcesRow) GetBiomarkerDataVocabularyOk() (*string, bool)`

GetBiomarkerDataVocabularyOk returns a tuple with the BiomarkerDataVocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerDataVocabulary

`func (o *RwdSourcesRow) SetBiomarkerDataVocabulary(v string)`

SetBiomarkerDataVocabulary sets BiomarkerDataVocabulary field to given value.

### HasBiomarkerDataVocabulary

`func (o *RwdSourcesRow) HasBiomarkerDataVocabulary() bool`

HasBiomarkerDataVocabulary returns a boolean if a field has been set.

### SetBiomarkerDataVocabularyNil

`func (o *RwdSourcesRow) SetBiomarkerDataVocabularyNil(b bool)`

 SetBiomarkerDataVocabularyNil sets the value for BiomarkerDataVocabulary to be an explicit nil

### UnsetBiomarkerDataVocabulary
`func (o *RwdSourcesRow) UnsetBiomarkerDataVocabulary()`

UnsetBiomarkerDataVocabulary ensures that no value is present for BiomarkerDataVocabulary, not even an explicit nil
### GetBiomarkerVocabularyOther

`func (o *RwdSourcesRow) GetBiomarkerVocabularyOther() string`

GetBiomarkerVocabularyOther returns the BiomarkerVocabularyOther field if non-nil, zero value otherwise.

### GetBiomarkerVocabularyOtherOk

`func (o *RwdSourcesRow) GetBiomarkerVocabularyOtherOk() (*string, bool)`

GetBiomarkerVocabularyOtherOk returns a tuple with the BiomarkerVocabularyOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerVocabularyOther

`func (o *RwdSourcesRow) SetBiomarkerVocabularyOther(v string)`

SetBiomarkerVocabularyOther sets BiomarkerVocabularyOther field to given value.

### HasBiomarkerVocabularyOther

`func (o *RwdSourcesRow) HasBiomarkerVocabularyOther() bool`

HasBiomarkerVocabularyOther returns a boolean if a field has been set.

### SetBiomarkerVocabularyOtherNil

`func (o *RwdSourcesRow) SetBiomarkerVocabularyOtherNil(b bool)`

 SetBiomarkerVocabularyOtherNil sets the value for BiomarkerVocabularyOther to be an explicit nil

### UnsetBiomarkerVocabularyOther
`func (o *RwdSourcesRow) UnsetBiomarkerVocabularyOther()`

UnsetBiomarkerVocabularyOther ensures that no value is present for BiomarkerVocabularyOther, not even an explicit nil
### GetPatientReportedOutcomes

`func (o *RwdSourcesRow) GetPatientReportedOutcomes() bool`

GetPatientReportedOutcomes returns the PatientReportedOutcomes field if non-nil, zero value otherwise.

### GetPatientReportedOutcomesOk

`func (o *RwdSourcesRow) GetPatientReportedOutcomesOk() (*bool, bool)`

GetPatientReportedOutcomesOk returns a tuple with the PatientReportedOutcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatientReportedOutcomes

`func (o *RwdSourcesRow) SetPatientReportedOutcomes(v bool)`

SetPatientReportedOutcomes sets PatientReportedOutcomes field to given value.

### HasPatientReportedOutcomes

`func (o *RwdSourcesRow) HasPatientReportedOutcomes() bool`

HasPatientReportedOutcomes returns a boolean if a field has been set.

### SetPatientReportedOutcomesNil

`func (o *RwdSourcesRow) SetPatientReportedOutcomesNil(b bool)`

 SetPatientReportedOutcomesNil sets the value for PatientReportedOutcomes to be an explicit nil

### UnsetPatientReportedOutcomes
`func (o *RwdSourcesRow) UnsetPatientReportedOutcomes()`

UnsetPatientReportedOutcomes ensures that no value is present for PatientReportedOutcomes, not even an explicit nil
### GetPatientGeneratedData

`func (o *RwdSourcesRow) GetPatientGeneratedData() bool`

GetPatientGeneratedData returns the PatientGeneratedData field if non-nil, zero value otherwise.

### GetPatientGeneratedDataOk

`func (o *RwdSourcesRow) GetPatientGeneratedDataOk() (*bool, bool)`

GetPatientGeneratedDataOk returns a tuple with the PatientGeneratedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatientGeneratedData

`func (o *RwdSourcesRow) SetPatientGeneratedData(v bool)`

SetPatientGeneratedData sets PatientGeneratedData field to given value.

### HasPatientGeneratedData

`func (o *RwdSourcesRow) HasPatientGeneratedData() bool`

HasPatientGeneratedData returns a boolean if a field has been set.

### SetPatientGeneratedDataNil

`func (o *RwdSourcesRow) SetPatientGeneratedDataNil(b bool)`

 SetPatientGeneratedDataNil sets the value for PatientGeneratedData to be an explicit nil

### UnsetPatientGeneratedData
`func (o *RwdSourcesRow) UnsetPatientGeneratedData()`

UnsetPatientGeneratedData ensures that no value is present for PatientGeneratedData, not even an explicit nil
### GetUnitsOfHealthcareUtilisation

`func (o *RwdSourcesRow) GetUnitsOfHealthcareUtilisation() bool`

GetUnitsOfHealthcareUtilisation returns the UnitsOfHealthcareUtilisation field if non-nil, zero value otherwise.

### GetUnitsOfHealthcareUtilisationOk

`func (o *RwdSourcesRow) GetUnitsOfHealthcareUtilisationOk() (*bool, bool)`

GetUnitsOfHealthcareUtilisationOk returns a tuple with the UnitsOfHealthcareUtilisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsOfHealthcareUtilisation

`func (o *RwdSourcesRow) SetUnitsOfHealthcareUtilisation(v bool)`

SetUnitsOfHealthcareUtilisation sets UnitsOfHealthcareUtilisation field to given value.

### HasUnitsOfHealthcareUtilisation

`func (o *RwdSourcesRow) HasUnitsOfHealthcareUtilisation() bool`

HasUnitsOfHealthcareUtilisation returns a boolean if a field has been set.

### SetUnitsOfHealthcareUtilisationNil

`func (o *RwdSourcesRow) SetUnitsOfHealthcareUtilisationNil(b bool)`

 SetUnitsOfHealthcareUtilisationNil sets the value for UnitsOfHealthcareUtilisation to be an explicit nil

### UnsetUnitsOfHealthcareUtilisation
`func (o *RwdSourcesRow) UnsetUnitsOfHealthcareUtilisation()`

UnsetUnitsOfHealthcareUtilisation ensures that no value is present for UnitsOfHealthcareUtilisation, not even an explicit nil
### GetUniqueIdentifierForPersons

`func (o *RwdSourcesRow) GetUniqueIdentifierForPersons() bool`

GetUniqueIdentifierForPersons returns the UniqueIdentifierForPersons field if non-nil, zero value otherwise.

### GetUniqueIdentifierForPersonsOk

`func (o *RwdSourcesRow) GetUniqueIdentifierForPersonsOk() (*bool, bool)`

GetUniqueIdentifierForPersonsOk returns a tuple with the UniqueIdentifierForPersons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifierForPersons

`func (o *RwdSourcesRow) SetUniqueIdentifierForPersons(v bool)`

SetUniqueIdentifierForPersons sets UniqueIdentifierForPersons field to given value.

### HasUniqueIdentifierForPersons

`func (o *RwdSourcesRow) HasUniqueIdentifierForPersons() bool`

HasUniqueIdentifierForPersons returns a boolean if a field has been set.

### SetUniqueIdentifierForPersonsNil

`func (o *RwdSourcesRow) SetUniqueIdentifierForPersonsNil(b bool)`

 SetUniqueIdentifierForPersonsNil sets the value for UniqueIdentifierForPersons to be an explicit nil

### UnsetUniqueIdentifierForPersons
`func (o *RwdSourcesRow) UnsetUniqueIdentifierForPersons()`

UnsetUniqueIdentifierForPersons ensures that no value is present for UniqueIdentifierForPersons, not even an explicit nil
### GetQualityOfLifeMeasurements

`func (o *RwdSourcesRow) GetQualityOfLifeMeasurements() bool`

GetQualityOfLifeMeasurements returns the QualityOfLifeMeasurements field if non-nil, zero value otherwise.

### GetQualityOfLifeMeasurementsOk

`func (o *RwdSourcesRow) GetQualityOfLifeMeasurementsOk() (*bool, bool)`

GetQualityOfLifeMeasurementsOk returns a tuple with the QualityOfLifeMeasurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityOfLifeMeasurements

`func (o *RwdSourcesRow) SetQualityOfLifeMeasurements(v bool)`

SetQualityOfLifeMeasurements sets QualityOfLifeMeasurements field to given value.

### HasQualityOfLifeMeasurements

`func (o *RwdSourcesRow) HasQualityOfLifeMeasurements() bool`

HasQualityOfLifeMeasurements returns a boolean if a field has been set.

### SetQualityOfLifeMeasurementsNil

`func (o *RwdSourcesRow) SetQualityOfLifeMeasurementsNil(b bool)`

 SetQualityOfLifeMeasurementsNil sets the value for QualityOfLifeMeasurements to be an explicit nil

### UnsetQualityOfLifeMeasurements
`func (o *RwdSourcesRow) UnsetQualityOfLifeMeasurements()`

UnsetQualityOfLifeMeasurements ensures that no value is present for QualityOfLifeMeasurements, not even an explicit nil
### GetQualityOfLifeMeasurementsVocabulary

`func (o *RwdSourcesRow) GetQualityOfLifeMeasurementsVocabulary() string`

GetQualityOfLifeMeasurementsVocabulary returns the QualityOfLifeMeasurementsVocabulary field if non-nil, zero value otherwise.

### GetQualityOfLifeMeasurementsVocabularyOk

`func (o *RwdSourcesRow) GetQualityOfLifeMeasurementsVocabularyOk() (*string, bool)`

GetQualityOfLifeMeasurementsVocabularyOk returns a tuple with the QualityOfLifeMeasurementsVocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityOfLifeMeasurementsVocabulary

`func (o *RwdSourcesRow) SetQualityOfLifeMeasurementsVocabulary(v string)`

SetQualityOfLifeMeasurementsVocabulary sets QualityOfLifeMeasurementsVocabulary field to given value.

### HasQualityOfLifeMeasurementsVocabulary

`func (o *RwdSourcesRow) HasQualityOfLifeMeasurementsVocabulary() bool`

HasQualityOfLifeMeasurementsVocabulary returns a boolean if a field has been set.

### SetQualityOfLifeMeasurementsVocabularyNil

`func (o *RwdSourcesRow) SetQualityOfLifeMeasurementsVocabularyNil(b bool)`

 SetQualityOfLifeMeasurementsVocabularyNil sets the value for QualityOfLifeMeasurementsVocabulary to be an explicit nil

### UnsetQualityOfLifeMeasurementsVocabulary
`func (o *RwdSourcesRow) UnsetQualityOfLifeMeasurementsVocabulary()`

UnsetQualityOfLifeMeasurementsVocabulary ensures that no value is present for QualityOfLifeMeasurementsVocabulary, not even an explicit nil
### GetQualityOfLifeMeasurementsOther

`func (o *RwdSourcesRow) GetQualityOfLifeMeasurementsOther() string`

GetQualityOfLifeMeasurementsOther returns the QualityOfLifeMeasurementsOther field if non-nil, zero value otherwise.

### GetQualityOfLifeMeasurementsOtherOk

`func (o *RwdSourcesRow) GetQualityOfLifeMeasurementsOtherOk() (*string, bool)`

GetQualityOfLifeMeasurementsOtherOk returns a tuple with the QualityOfLifeMeasurementsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityOfLifeMeasurementsOther

`func (o *RwdSourcesRow) SetQualityOfLifeMeasurementsOther(v string)`

SetQualityOfLifeMeasurementsOther sets QualityOfLifeMeasurementsOther field to given value.

### HasQualityOfLifeMeasurementsOther

`func (o *RwdSourcesRow) HasQualityOfLifeMeasurementsOther() bool`

HasQualityOfLifeMeasurementsOther returns a boolean if a field has been set.

### SetQualityOfLifeMeasurementsOtherNil

`func (o *RwdSourcesRow) SetQualityOfLifeMeasurementsOtherNil(b bool)`

 SetQualityOfLifeMeasurementsOtherNil sets the value for QualityOfLifeMeasurementsOther to be an explicit nil

### UnsetQualityOfLifeMeasurementsOther
`func (o *RwdSourcesRow) UnsetQualityOfLifeMeasurementsOther()`

UnsetQualityOfLifeMeasurementsOther ensures that no value is present for QualityOfLifeMeasurementsOther, not even an explicit nil
### GetLifestyleFactors

`func (o *RwdSourcesRow) GetLifestyleFactors() bool`

GetLifestyleFactors returns the LifestyleFactors field if non-nil, zero value otherwise.

### GetLifestyleFactorsOk

`func (o *RwdSourcesRow) GetLifestyleFactorsOk() (*bool, bool)`

GetLifestyleFactorsOk returns a tuple with the LifestyleFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifestyleFactors

`func (o *RwdSourcesRow) SetLifestyleFactors(v bool)`

SetLifestyleFactors sets LifestyleFactors field to given value.

### HasLifestyleFactors

`func (o *RwdSourcesRow) HasLifestyleFactors() bool`

HasLifestyleFactors returns a boolean if a field has been set.

### SetLifestyleFactorsNil

`func (o *RwdSourcesRow) SetLifestyleFactorsNil(b bool)`

 SetLifestyleFactorsNil sets the value for LifestyleFactors to be an explicit nil

### UnsetLifestyleFactors
`func (o *RwdSourcesRow) UnsetLifestyleFactors()`

UnsetLifestyleFactors ensures that no value is present for LifestyleFactors, not even an explicit nil
### GetSociodemographicInformation

`func (o *RwdSourcesRow) GetSociodemographicInformation() bool`

GetSociodemographicInformation returns the SociodemographicInformation field if non-nil, zero value otherwise.

### GetSociodemographicInformationOk

`func (o *RwdSourcesRow) GetSociodemographicInformationOk() (*bool, bool)`

GetSociodemographicInformationOk returns a tuple with the SociodemographicInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSociodemographicInformation

`func (o *RwdSourcesRow) SetSociodemographicInformation(v bool)`

SetSociodemographicInformation sets SociodemographicInformation field to given value.

### HasSociodemographicInformation

`func (o *RwdSourcesRow) HasSociodemographicInformation() bool`

HasSociodemographicInformation returns a boolean if a field has been set.

### SetSociodemographicInformationNil

`func (o *RwdSourcesRow) SetSociodemographicInformationNil(b bool)`

 SetSociodemographicInformationNil sets the value for SociodemographicInformation to be an explicit nil

### UnsetSociodemographicInformation
`func (o *RwdSourcesRow) UnsetSociodemographicInformation()`

UnsetSociodemographicInformation ensures that no value is present for SociodemographicInformation, not even an explicit nil
### GetPopulationSize

`func (o *RwdSourcesRow) GetPopulationSize() int64`

GetPopulationSize returns the PopulationSize field if non-nil, zero value otherwise.

### GetPopulationSizeOk

`func (o *RwdSourcesRow) GetPopulationSizeOk() (*int64, bool)`

GetPopulationSizeOk returns a tuple with the PopulationSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationSize

`func (o *RwdSourcesRow) SetPopulationSize(v int64)`

SetPopulationSize sets PopulationSize field to given value.

### HasPopulationSize

`func (o *RwdSourcesRow) HasPopulationSize() bool`

HasPopulationSize returns a boolean if a field has been set.

### SetPopulationSizeNil

`func (o *RwdSourcesRow) SetPopulationSizeNil(b bool)`

 SetPopulationSizeNil sets the value for PopulationSize to be an explicit nil

### UnsetPopulationSize
`func (o *RwdSourcesRow) UnsetPopulationSize()`

UnsetPopulationSize ensures that no value is present for PopulationSize, not even an explicit nil
### GetActivePopulationSize

`func (o *RwdSourcesRow) GetActivePopulationSize() int64`

GetActivePopulationSize returns the ActivePopulationSize field if non-nil, zero value otherwise.

### GetActivePopulationSizeOk

`func (o *RwdSourcesRow) GetActivePopulationSizeOk() (*int64, bool)`

GetActivePopulationSizeOk returns a tuple with the ActivePopulationSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePopulationSize

`func (o *RwdSourcesRow) SetActivePopulationSize(v int64)`

SetActivePopulationSize sets ActivePopulationSize field to given value.

### HasActivePopulationSize

`func (o *RwdSourcesRow) HasActivePopulationSize() bool`

HasActivePopulationSize returns a boolean if a field has been set.

### SetActivePopulationSizeNil

`func (o *RwdSourcesRow) SetActivePopulationSizeNil(b bool)`

 SetActivePopulationSizeNil sets the value for ActivePopulationSize to be an explicit nil

### UnsetActivePopulationSize
`func (o *RwdSourcesRow) UnsetActivePopulationSize()`

UnsetActivePopulationSize ensures that no value is present for ActivePopulationSize, not even an explicit nil
### GetMedianObservationTimeYears

`func (o *RwdSourcesRow) GetMedianObservationTimeYears() float32`

GetMedianObservationTimeYears returns the MedianObservationTimeYears field if non-nil, zero value otherwise.

### GetMedianObservationTimeYearsOk

`func (o *RwdSourcesRow) GetMedianObservationTimeYearsOk() (*float32, bool)`

GetMedianObservationTimeYearsOk returns a tuple with the MedianObservationTimeYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianObservationTimeYears

`func (o *RwdSourcesRow) SetMedianObservationTimeYears(v float32)`

SetMedianObservationTimeYears sets MedianObservationTimeYears field to given value.

### HasMedianObservationTimeYears

`func (o *RwdSourcesRow) HasMedianObservationTimeYears() bool`

HasMedianObservationTimeYears returns a boolean if a field has been set.

### SetMedianObservationTimeYearsNil

`func (o *RwdSourcesRow) SetMedianObservationTimeYearsNil(b bool)`

 SetMedianObservationTimeYearsNil sets the value for MedianObservationTimeYears to be an explicit nil

### UnsetMedianObservationTimeYears
`func (o *RwdSourcesRow) UnsetMedianObservationTimeYears()`

UnsetMedianObservationTimeYears ensures that no value is present for MedianObservationTimeYears, not even an explicit nil
### GetMedianObservationTimeActiveYears

`func (o *RwdSourcesRow) GetMedianObservationTimeActiveYears() float32`

GetMedianObservationTimeActiveYears returns the MedianObservationTimeActiveYears field if non-nil, zero value otherwise.

### GetMedianObservationTimeActiveYearsOk

`func (o *RwdSourcesRow) GetMedianObservationTimeActiveYearsOk() (*float32, bool)`

GetMedianObservationTimeActiveYearsOk returns a tuple with the MedianObservationTimeActiveYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianObservationTimeActiveYears

`func (o *RwdSourcesRow) SetMedianObservationTimeActiveYears(v float32)`

SetMedianObservationTimeActiveYears sets MedianObservationTimeActiveYears field to given value.

### HasMedianObservationTimeActiveYears

`func (o *RwdSourcesRow) HasMedianObservationTimeActiveYears() bool`

HasMedianObservationTimeActiveYears returns a boolean if a field has been set.

### SetMedianObservationTimeActiveYearsNil

`func (o *RwdSourcesRow) SetMedianObservationTimeActiveYearsNil(b bool)`

 SetMedianObservationTimeActiveYearsNil sets the value for MedianObservationTimeActiveYears to be an explicit nil

### UnsetMedianObservationTimeActiveYears
`func (o *RwdSourcesRow) UnsetMedianObservationTimeActiveYears()`

UnsetMedianObservationTimeActiveYears ensures that no value is present for MedianObservationTimeActiveYears, not even an explicit nil
### GetEstablishedOn

`func (o *RwdSourcesRow) GetEstablishedOn() string`

GetEstablishedOn returns the EstablishedOn field if non-nil, zero value otherwise.

### GetEstablishedOnOk

`func (o *RwdSourcesRow) GetEstablishedOnOk() (*string, bool)`

GetEstablishedOnOk returns a tuple with the EstablishedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablishedOn

`func (o *RwdSourcesRow) SetEstablishedOn(v string)`

SetEstablishedOn sets EstablishedOn field to given value.

### HasEstablishedOn

`func (o *RwdSourcesRow) HasEstablishedOn() bool`

HasEstablishedOn returns a boolean if a field has been set.

### SetEstablishedOnNil

`func (o *RwdSourcesRow) SetEstablishedOnNil(b bool)`

 SetEstablishedOnNil sets the value for EstablishedOn to be an explicit nil

### UnsetEstablishedOn
`func (o *RwdSourcesRow) UnsetEstablishedOn()`

UnsetEstablishedOn ensures that no value is present for EstablishedOn, not even an explicit nil
### GetCollectionFirstOn

`func (o *RwdSourcesRow) GetCollectionFirstOn() string`

GetCollectionFirstOn returns the CollectionFirstOn field if non-nil, zero value otherwise.

### GetCollectionFirstOnOk

`func (o *RwdSourcesRow) GetCollectionFirstOnOk() (*string, bool)`

GetCollectionFirstOnOk returns a tuple with the CollectionFirstOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionFirstOn

`func (o *RwdSourcesRow) SetCollectionFirstOn(v string)`

SetCollectionFirstOn sets CollectionFirstOn field to given value.

### HasCollectionFirstOn

`func (o *RwdSourcesRow) HasCollectionFirstOn() bool`

HasCollectionFirstOn returns a boolean if a field has been set.

### SetCollectionFirstOnNil

`func (o *RwdSourcesRow) SetCollectionFirstOnNil(b bool)`

 SetCollectionFirstOnNil sets the value for CollectionFirstOn to be an explicit nil

### UnsetCollectionFirstOn
`func (o *RwdSourcesRow) UnsetCollectionFirstOn()`

UnsetCollectionFirstOn ensures that no value is present for CollectionFirstOn, not even an explicit nil
### GetCollectionLastOn

`func (o *RwdSourcesRow) GetCollectionLastOn() string`

GetCollectionLastOn returns the CollectionLastOn field if non-nil, zero value otherwise.

### GetCollectionLastOnOk

`func (o *RwdSourcesRow) GetCollectionLastOnOk() (*string, bool)`

GetCollectionLastOnOk returns a tuple with the CollectionLastOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionLastOn

`func (o *RwdSourcesRow) SetCollectionLastOn(v string)`

SetCollectionLastOn sets CollectionLastOn field to given value.

### HasCollectionLastOn

`func (o *RwdSourcesRow) HasCollectionLastOn() bool`

HasCollectionLastOn returns a boolean if a field has been set.

### SetCollectionLastOnNil

`func (o *RwdSourcesRow) SetCollectionLastOnNil(b bool)`

 SetCollectionLastOnNil sets the value for CollectionLastOn to be an explicit nil

### UnsetCollectionLastOn
`func (o *RwdSourcesRow) UnsetCollectionLastOn()`

UnsetCollectionLastOn ensures that no value is present for CollectionLastOn, not even an explicit nil
### GetDeclaredHistorySpanYears

`func (o *RwdSourcesRow) GetDeclaredHistorySpanYears() float32`

GetDeclaredHistorySpanYears returns the DeclaredHistorySpanYears field if non-nil, zero value otherwise.

### GetDeclaredHistorySpanYearsOk

`func (o *RwdSourcesRow) GetDeclaredHistorySpanYearsOk() (*float32, bool)`

GetDeclaredHistorySpanYearsOk returns a tuple with the DeclaredHistorySpanYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredHistorySpanYears

`func (o *RwdSourcesRow) SetDeclaredHistorySpanYears(v float32)`

SetDeclaredHistorySpanYears sets DeclaredHistorySpanYears field to given value.

### HasDeclaredHistorySpanYears

`func (o *RwdSourcesRow) HasDeclaredHistorySpanYears() bool`

HasDeclaredHistorySpanYears returns a boolean if a field has been set.

### SetDeclaredHistorySpanYearsNil

`func (o *RwdSourcesRow) SetDeclaredHistorySpanYearsNil(b bool)`

 SetDeclaredHistorySpanYearsNil sets the value for DeclaredHistorySpanYears to be an explicit nil

### UnsetDeclaredHistorySpanYears
`func (o *RwdSourcesRow) UnsetDeclaredHistorySpanYears()`

UnsetDeclaredHistorySpanYears ensures that no value is present for DeclaredHistorySpanYears, not even an explicit nil
### GetMedianObservationTimePlausibility

`func (o *RwdSourcesRow) GetMedianObservationTimePlausibility() string`

GetMedianObservationTimePlausibility returns the MedianObservationTimePlausibility field if non-nil, zero value otherwise.

### GetMedianObservationTimePlausibilityOk

`func (o *RwdSourcesRow) GetMedianObservationTimePlausibilityOk() (*string, bool)`

GetMedianObservationTimePlausibilityOk returns a tuple with the MedianObservationTimePlausibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianObservationTimePlausibility

`func (o *RwdSourcesRow) SetMedianObservationTimePlausibility(v string)`

SetMedianObservationTimePlausibility sets MedianObservationTimePlausibility field to given value.

### HasMedianObservationTimePlausibility

`func (o *RwdSourcesRow) HasMedianObservationTimePlausibility() bool`

HasMedianObservationTimePlausibility returns a boolean if a field has been set.

### SetMedianObservationTimePlausibilityNil

`func (o *RwdSourcesRow) SetMedianObservationTimePlausibilityNil(b bool)`

 SetMedianObservationTimePlausibilityNil sets the value for MedianObservationTimePlausibility to be an explicit nil

### UnsetMedianObservationTimePlausibility
`func (o *RwdSourcesRow) UnsetMedianObservationTimePlausibility()`

UnsetMedianObservationTimePlausibility ensures that no value is present for MedianObservationTimePlausibility, not even an explicit nil
### GetMedianObservationTimeActivePlausibility

`func (o *RwdSourcesRow) GetMedianObservationTimeActivePlausibility() string`

GetMedianObservationTimeActivePlausibility returns the MedianObservationTimeActivePlausibility field if non-nil, zero value otherwise.

### GetMedianObservationTimeActivePlausibilityOk

`func (o *RwdSourcesRow) GetMedianObservationTimeActivePlausibilityOk() (*string, bool)`

GetMedianObservationTimeActivePlausibilityOk returns a tuple with the MedianObservationTimeActivePlausibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianObservationTimeActivePlausibility

`func (o *RwdSourcesRow) SetMedianObservationTimeActivePlausibility(v string)`

SetMedianObservationTimeActivePlausibility sets MedianObservationTimeActivePlausibility field to given value.

### HasMedianObservationTimeActivePlausibility

`func (o *RwdSourcesRow) HasMedianObservationTimeActivePlausibility() bool`

HasMedianObservationTimeActivePlausibility returns a boolean if a field has been set.

### SetMedianObservationTimeActivePlausibilityNil

`func (o *RwdSourcesRow) SetMedianObservationTimeActivePlausibilityNil(b bool)`

 SetMedianObservationTimeActivePlausibilityNil sets the value for MedianObservationTimeActivePlausibility to be an explicit nil

### UnsetMedianObservationTimeActivePlausibility
`func (o *RwdSourcesRow) UnsetMedianObservationTimeActivePlausibility()`

UnsetMedianObservationTimeActivePlausibility ensures that no value is present for MedianObservationTimeActivePlausibility, not even an explicit nil
### GetEstimatedPopulationCoverageText

`func (o *RwdSourcesRow) GetEstimatedPopulationCoverageText() string`

GetEstimatedPopulationCoverageText returns the EstimatedPopulationCoverageText field if non-nil, zero value otherwise.

### GetEstimatedPopulationCoverageTextOk

`func (o *RwdSourcesRow) GetEstimatedPopulationCoverageTextOk() (*string, bool)`

GetEstimatedPopulationCoverageTextOk returns a tuple with the EstimatedPopulationCoverageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedPopulationCoverageText

`func (o *RwdSourcesRow) SetEstimatedPopulationCoverageText(v string)`

SetEstimatedPopulationCoverageText sets EstimatedPopulationCoverageText field to given value.

### HasEstimatedPopulationCoverageText

`func (o *RwdSourcesRow) HasEstimatedPopulationCoverageText() bool`

HasEstimatedPopulationCoverageText returns a boolean if a field has been set.

### SetEstimatedPopulationCoverageTextNil

`func (o *RwdSourcesRow) SetEstimatedPopulationCoverageTextNil(b bool)`

 SetEstimatedPopulationCoverageTextNil sets the value for EstimatedPopulationCoverageText to be an explicit nil

### UnsetEstimatedPopulationCoverageText
`func (o *RwdSourcesRow) UnsetEstimatedPopulationCoverageText()`

UnsetEstimatedPopulationCoverageText ensures that no value is present for EstimatedPopulationCoverageText, not even an explicit nil
### GetAgeGroups

`func (o *RwdSourcesRow) GetAgeGroups() string`

GetAgeGroups returns the AgeGroups field if non-nil, zero value otherwise.

### GetAgeGroupsOk

`func (o *RwdSourcesRow) GetAgeGroupsOk() (*string, bool)`

GetAgeGroupsOk returns a tuple with the AgeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeGroups

`func (o *RwdSourcesRow) SetAgeGroups(v string)`

SetAgeGroups sets AgeGroups field to given value.

### HasAgeGroups

`func (o *RwdSourcesRow) HasAgeGroups() bool`

HasAgeGroups returns a boolean if a field has been set.

### SetAgeGroupsNil

`func (o *RwdSourcesRow) SetAgeGroupsNil(b bool)`

 SetAgeGroupsNil sets the value for AgeGroups to be an explicit nil

### UnsetAgeGroups
`func (o *RwdSourcesRow) UnsetAgeGroups()`

UnsetAgeGroups ensures that no value is present for AgeGroups, not even an explicit nil
### GetAgeBandCount

`func (o *RwdSourcesRow) GetAgeBandCount() int32`

GetAgeBandCount returns the AgeBandCount field if non-nil, zero value otherwise.

### GetAgeBandCountOk

`func (o *RwdSourcesRow) GetAgeBandCountOk() (*int32, bool)`

GetAgeBandCountOk returns a tuple with the AgeBandCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeBandCount

`func (o *RwdSourcesRow) SetAgeBandCount(v int32)`

SetAgeBandCount sets AgeBandCount field to given value.

### HasAgeBandCount

`func (o *RwdSourcesRow) HasAgeBandCount() bool`

HasAgeBandCount returns a boolean if a field has been set.

### SetAgeBandCountNil

`func (o *RwdSourcesRow) SetAgeBandCountNil(b bool)`

 SetAgeBandCountNil sets the value for AgeBandCount to be an explicit nil

### UnsetAgeBandCount
`func (o *RwdSourcesRow) UnsetAgeBandCount()`

UnsetAgeBandCount ensures that no value is present for AgeBandCount, not even an explicit nil
### GetHasPopulationBreakdown

`func (o *RwdSourcesRow) GetHasPopulationBreakdown() bool`

GetHasPopulationBreakdown returns the HasPopulationBreakdown field if non-nil, zero value otherwise.

### GetHasPopulationBreakdownOk

`func (o *RwdSourcesRow) GetHasPopulationBreakdownOk() (*bool, bool)`

GetHasPopulationBreakdownOk returns a tuple with the HasPopulationBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPopulationBreakdown

`func (o *RwdSourcesRow) SetHasPopulationBreakdown(v bool)`

SetHasPopulationBreakdown sets HasPopulationBreakdown field to given value.

### HasHasPopulationBreakdown

`func (o *RwdSourcesRow) HasHasPopulationBreakdown() bool`

HasHasPopulationBreakdown returns a boolean if a field has been set.

### SetHasPopulationBreakdownNil

`func (o *RwdSourcesRow) SetHasPopulationBreakdownNil(b bool)`

 SetHasPopulationBreakdownNil sets the value for HasPopulationBreakdown to be an explicit nil

### UnsetHasPopulationBreakdown
`func (o *RwdSourcesRow) UnsetHasPopulationBreakdown()`

UnsetHasPopulationBreakdown ensures that no value is present for HasPopulationBreakdown, not even an explicit nil
### GetHasActivePopulationBreakdown

`func (o *RwdSourcesRow) GetHasActivePopulationBreakdown() bool`

GetHasActivePopulationBreakdown returns the HasActivePopulationBreakdown field if non-nil, zero value otherwise.

### GetHasActivePopulationBreakdownOk

`func (o *RwdSourcesRow) GetHasActivePopulationBreakdownOk() (*bool, bool)`

GetHasActivePopulationBreakdownOk returns a tuple with the HasActivePopulationBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActivePopulationBreakdown

`func (o *RwdSourcesRow) SetHasActivePopulationBreakdown(v bool)`

SetHasActivePopulationBreakdown sets HasActivePopulationBreakdown field to given value.

### HasHasActivePopulationBreakdown

`func (o *RwdSourcesRow) HasHasActivePopulationBreakdown() bool`

HasHasActivePopulationBreakdown returns a boolean if a field has been set.

### SetHasActivePopulationBreakdownNil

`func (o *RwdSourcesRow) SetHasActivePopulationBreakdownNil(b bool)`

 SetHasActivePopulationBreakdownNil sets the value for HasActivePopulationBreakdown to be an explicit nil

### UnsetHasActivePopulationBreakdown
`func (o *RwdSourcesRow) UnsetHasActivePopulationBreakdown()`

UnsetHasActivePopulationBreakdown ensures that no value is present for HasActivePopulationBreakdown, not even an explicit nil
### GetAgeBandsAreNested

`func (o *RwdSourcesRow) GetAgeBandsAreNested() bool`

GetAgeBandsAreNested returns the AgeBandsAreNested field if non-nil, zero value otherwise.

### GetAgeBandsAreNestedOk

`func (o *RwdSourcesRow) GetAgeBandsAreNestedOk() (*bool, bool)`

GetAgeBandsAreNestedOk returns a tuple with the AgeBandsAreNested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeBandsAreNested

`func (o *RwdSourcesRow) SetAgeBandsAreNested(v bool)`

SetAgeBandsAreNested sets AgeBandsAreNested field to given value.

### HasAgeBandsAreNested

`func (o *RwdSourcesRow) HasAgeBandsAreNested() bool`

HasAgeBandsAreNested returns a boolean if a field has been set.

### SetAgeBandsAreNestedNil

`func (o *RwdSourcesRow) SetAgeBandsAreNestedNil(b bool)`

 SetAgeBandsAreNestedNil sets the value for AgeBandsAreNested to be an explicit nil

### UnsetAgeBandsAreNested
`func (o *RwdSourcesRow) UnsetAgeBandsAreNested()`

UnsetAgeBandsAreNested ensures that no value is present for AgeBandsAreNested, not even an explicit nil
### GetCdmMapping

`func (o *RwdSourcesRow) GetCdmMapping() bool`

GetCdmMapping returns the CdmMapping field if non-nil, zero value otherwise.

### GetCdmMappingOk

`func (o *RwdSourcesRow) GetCdmMappingOk() (*bool, bool)`

GetCdmMappingOk returns a tuple with the CdmMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdmMapping

`func (o *RwdSourcesRow) SetCdmMapping(v bool)`

SetCdmMapping sets CdmMapping field to given value.

### HasCdmMapping

`func (o *RwdSourcesRow) HasCdmMapping() bool`

HasCdmMapping returns a boolean if a field has been set.

### SetCdmMappingNil

`func (o *RwdSourcesRow) SetCdmMappingNil(b bool)`

 SetCdmMappingNil sets the value for CdmMapping to be an explicit nil

### UnsetCdmMapping
`func (o *RwdSourcesRow) UnsetCdmMapping()`

UnsetCdmMapping ensures that no value is present for CdmMapping, not even an explicit nil
### GetEtlStatus

`func (o *RwdSourcesRow) GetEtlStatus() string`

GetEtlStatus returns the EtlStatus field if non-nil, zero value otherwise.

### GetEtlStatusOk

`func (o *RwdSourcesRow) GetEtlStatusOk() (*string, bool)`

GetEtlStatusOk returns a tuple with the EtlStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlStatus

`func (o *RwdSourcesRow) SetEtlStatus(v string)`

SetEtlStatus sets EtlStatus field to given value.

### HasEtlStatus

`func (o *RwdSourcesRow) HasEtlStatus() bool`

HasEtlStatus returns a boolean if a field has been set.

### SetEtlStatusNil

`func (o *RwdSourcesRow) SetEtlStatusNil(b bool)`

 SetEtlStatusNil sets the value for EtlStatus to be an explicit nil

### UnsetEtlStatus
`func (o *RwdSourcesRow) UnsetEtlStatus()`

UnsetEtlStatus ensures that no value is present for EtlStatus, not even an explicit nil
### GetEtlCdmVersion

`func (o *RwdSourcesRow) GetEtlCdmVersion() string`

GetEtlCdmVersion returns the EtlCdmVersion field if non-nil, zero value otherwise.

### GetEtlCdmVersionOk

`func (o *RwdSourcesRow) GetEtlCdmVersionOk() (*string, bool)`

GetEtlCdmVersionOk returns a tuple with the EtlCdmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlCdmVersion

`func (o *RwdSourcesRow) SetEtlCdmVersion(v string)`

SetEtlCdmVersion sets EtlCdmVersion field to given value.

### HasEtlCdmVersion

`func (o *RwdSourcesRow) HasEtlCdmVersion() bool`

HasEtlCdmVersion returns a boolean if a field has been set.

### SetEtlCdmVersionNil

`func (o *RwdSourcesRow) SetEtlCdmVersionNil(b bool)`

 SetEtlCdmVersionNil sets the value for EtlCdmVersion to be an explicit nil

### UnsetEtlCdmVersion
`func (o *RwdSourcesRow) UnsetEtlCdmVersion()`

UnsetEtlCdmVersion ensures that no value is present for EtlCdmVersion, not even an explicit nil
### GetEtlFrequency

`func (o *RwdSourcesRow) GetEtlFrequency() string`

GetEtlFrequency returns the EtlFrequency field if non-nil, zero value otherwise.

### GetEtlFrequencyOk

`func (o *RwdSourcesRow) GetEtlFrequencyOk() (*string, bool)`

GetEtlFrequencyOk returns a tuple with the EtlFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlFrequency

`func (o *RwdSourcesRow) SetEtlFrequency(v string)`

SetEtlFrequency sets EtlFrequency field to given value.

### HasEtlFrequency

`func (o *RwdSourcesRow) HasEtlFrequency() bool`

HasEtlFrequency returns a boolean if a field has been set.

### SetEtlFrequencyNil

`func (o *RwdSourcesRow) SetEtlFrequencyNil(b bool)`

 SetEtlFrequencyNil sets the value for EtlFrequency to be an explicit nil

### UnsetEtlFrequency
`func (o *RwdSourcesRow) UnsetEtlFrequency()`

UnsetEtlFrequency ensures that no value is present for EtlFrequency, not even an explicit nil
### GetEtlSpecificationsFile

`func (o *RwdSourcesRow) GetEtlSpecificationsFile() string`

GetEtlSpecificationsFile returns the EtlSpecificationsFile field if non-nil, zero value otherwise.

### GetEtlSpecificationsFileOk

`func (o *RwdSourcesRow) GetEtlSpecificationsFileOk() (*string, bool)`

GetEtlSpecificationsFileOk returns a tuple with the EtlSpecificationsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlSpecificationsFile

`func (o *RwdSourcesRow) SetEtlSpecificationsFile(v string)`

SetEtlSpecificationsFile sets EtlSpecificationsFile field to given value.

### HasEtlSpecificationsFile

`func (o *RwdSourcesRow) HasEtlSpecificationsFile() bool`

HasEtlSpecificationsFile returns a boolean if a field has been set.

### SetEtlSpecificationsFileNil

`func (o *RwdSourcesRow) SetEtlSpecificationsFileNil(b bool)`

 SetEtlSpecificationsFileNil sets the value for EtlSpecificationsFile to be an explicit nil

### UnsetEtlSpecificationsFile
`func (o *RwdSourcesRow) UnsetEtlSpecificationsFile()`

UnsetEtlSpecificationsFile ensures that no value is present for EtlSpecificationsFile, not even an explicit nil
### GetDeclaredMappingCount

`func (o *RwdSourcesRow) GetDeclaredMappingCount() int32`

GetDeclaredMappingCount returns the DeclaredMappingCount field if non-nil, zero value otherwise.

### GetDeclaredMappingCountOk

`func (o *RwdSourcesRow) GetDeclaredMappingCountOk() (*int32, bool)`

GetDeclaredMappingCountOk returns a tuple with the DeclaredMappingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredMappingCount

`func (o *RwdSourcesRow) SetDeclaredMappingCount(v int32)`

SetDeclaredMappingCount sets DeclaredMappingCount field to given value.

### HasDeclaredMappingCount

`func (o *RwdSourcesRow) HasDeclaredMappingCount() bool`

HasDeclaredMappingCount returns a boolean if a field has been set.

### SetDeclaredMappingCountNil

`func (o *RwdSourcesRow) SetDeclaredMappingCountNil(b bool)`

 SetDeclaredMappingCountNil sets the value for DeclaredMappingCount to be an explicit nil

### UnsetDeclaredMappingCount
`func (o *RwdSourcesRow) UnsetDeclaredMappingCount()`

UnsetDeclaredMappingCount ensures that no value is present for DeclaredMappingCount, not even an explicit nil
### GetCdmReadiness

`func (o *RwdSourcesRow) GetCdmReadiness() string`

GetCdmReadiness returns the CdmReadiness field if non-nil, zero value otherwise.

### GetCdmReadinessOk

`func (o *RwdSourcesRow) GetCdmReadinessOk() (*string, bool)`

GetCdmReadinessOk returns a tuple with the CdmReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdmReadiness

`func (o *RwdSourcesRow) SetCdmReadiness(v string)`

SetCdmReadiness sets CdmReadiness field to given value.

### HasCdmReadiness

`func (o *RwdSourcesRow) HasCdmReadiness() bool`

HasCdmReadiness returns a boolean if a field has been set.

### SetCdmReadinessNil

`func (o *RwdSourcesRow) SetCdmReadinessNil(b bool)`

 SetCdmReadinessNil sets the value for CdmReadiness to be an explicit nil

### UnsetCdmReadiness
`func (o *RwdSourcesRow) UnsetCdmReadiness()`

UnsetCdmReadiness ensures that no value is present for CdmReadiness, not even an explicit nil
### GetLinkage

`func (o *RwdSourcesRow) GetLinkage() bool`

GetLinkage returns the Linkage field if non-nil, zero value otherwise.

### GetLinkageOk

`func (o *RwdSourcesRow) GetLinkageOk() (*bool, bool)`

GetLinkageOk returns a tuple with the Linkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkage

`func (o *RwdSourcesRow) SetLinkage(v bool)`

SetLinkage sets Linkage field to given value.

### HasLinkage

`func (o *RwdSourcesRow) HasLinkage() bool`

HasLinkage returns a boolean if a field has been set.

### SetLinkageNil

`func (o *RwdSourcesRow) SetLinkageNil(b bool)`

 SetLinkageNil sets the value for Linkage to be an explicit nil

### UnsetLinkage
`func (o *RwdSourcesRow) UnsetLinkage()`

UnsetLinkage ensures that no value is present for Linkage, not even an explicit nil
### GetHasPreLinkedDeclaration

`func (o *RwdSourcesRow) GetHasPreLinkedDeclaration() bool`

GetHasPreLinkedDeclaration returns the HasPreLinkedDeclaration field if non-nil, zero value otherwise.

### GetHasPreLinkedDeclarationOk

`func (o *RwdSourcesRow) GetHasPreLinkedDeclarationOk() (*bool, bool)`

GetHasPreLinkedDeclarationOk returns a tuple with the HasPreLinkedDeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreLinkedDeclaration

`func (o *RwdSourcesRow) SetHasPreLinkedDeclaration(v bool)`

SetHasPreLinkedDeclaration sets HasPreLinkedDeclaration field to given value.

### HasHasPreLinkedDeclaration

`func (o *RwdSourcesRow) HasHasPreLinkedDeclaration() bool`

HasHasPreLinkedDeclaration returns a boolean if a field has been set.

### SetHasPreLinkedDeclarationNil

`func (o *RwdSourcesRow) SetHasPreLinkedDeclarationNil(b bool)`

 SetHasPreLinkedDeclarationNil sets the value for HasPreLinkedDeclaration to be an explicit nil

### UnsetHasPreLinkedDeclaration
`func (o *RwdSourcesRow) UnsetHasPreLinkedDeclaration()`

UnsetHasPreLinkedDeclaration ensures that no value is present for HasPreLinkedDeclaration, not even an explicit nil
### GetHasPossibleLinkageDeclaration

`func (o *RwdSourcesRow) GetHasPossibleLinkageDeclaration() bool`

GetHasPossibleLinkageDeclaration returns the HasPossibleLinkageDeclaration field if non-nil, zero value otherwise.

### GetHasPossibleLinkageDeclarationOk

`func (o *RwdSourcesRow) GetHasPossibleLinkageDeclarationOk() (*bool, bool)`

GetHasPossibleLinkageDeclarationOk returns a tuple with the HasPossibleLinkageDeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPossibleLinkageDeclaration

`func (o *RwdSourcesRow) SetHasPossibleLinkageDeclaration(v bool)`

SetHasPossibleLinkageDeclaration sets HasPossibleLinkageDeclaration field to given value.

### HasHasPossibleLinkageDeclaration

`func (o *RwdSourcesRow) HasHasPossibleLinkageDeclaration() bool`

HasHasPossibleLinkageDeclaration returns a boolean if a field has been set.

### SetHasPossibleLinkageDeclarationNil

`func (o *RwdSourcesRow) SetHasPossibleLinkageDeclarationNil(b bool)`

 SetHasPossibleLinkageDeclarationNil sets the value for HasPossibleLinkageDeclaration to be an explicit nil

### UnsetHasPossibleLinkageDeclaration
`func (o *RwdSourcesRow) UnsetHasPossibleLinkageDeclaration()`

UnsetHasPossibleLinkageDeclaration ensures that no value is present for HasPossibleLinkageDeclaration, not even an explicit nil
### GetLinkedDataSource

`func (o *RwdSourcesRow) GetLinkedDataSource() string`

GetLinkedDataSource returns the LinkedDataSource field if non-nil, zero value otherwise.

### GetLinkedDataSourceOk

`func (o *RwdSourcesRow) GetLinkedDataSourceOk() (*string, bool)`

GetLinkedDataSourceOk returns a tuple with the LinkedDataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDataSource

`func (o *RwdSourcesRow) SetLinkedDataSource(v string)`

SetLinkedDataSource sets LinkedDataSource field to given value.

### HasLinkedDataSource

`func (o *RwdSourcesRow) HasLinkedDataSource() bool`

HasLinkedDataSource returns a boolean if a field has been set.

### SetLinkedDataSourceNil

`func (o *RwdSourcesRow) SetLinkedDataSourceNil(b bool)`

 SetLinkedDataSourceNil sets the value for LinkedDataSource to be an explicit nil

### UnsetLinkedDataSource
`func (o *RwdSourcesRow) UnsetLinkedDataSource()`

UnsetLinkedDataSource ensures that no value is present for LinkedDataSource, not even an explicit nil
### GetLinkedDataSourceOther

`func (o *RwdSourcesRow) GetLinkedDataSourceOther() string`

GetLinkedDataSourceOther returns the LinkedDataSourceOther field if non-nil, zero value otherwise.

### GetLinkedDataSourceOtherOk

`func (o *RwdSourcesRow) GetLinkedDataSourceOtherOk() (*string, bool)`

GetLinkedDataSourceOtherOk returns a tuple with the LinkedDataSourceOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDataSourceOther

`func (o *RwdSourcesRow) SetLinkedDataSourceOther(v string)`

SetLinkedDataSourceOther sets LinkedDataSourceOther field to given value.

### HasLinkedDataSourceOther

`func (o *RwdSourcesRow) HasLinkedDataSourceOther() bool`

HasLinkedDataSourceOther returns a boolean if a field has been set.

### SetLinkedDataSourceOtherNil

`func (o *RwdSourcesRow) SetLinkedDataSourceOtherNil(b bool)`

 SetLinkedDataSourceOtherNil sets the value for LinkedDataSourceOther to be an explicit nil

### UnsetLinkedDataSourceOther
`func (o *RwdSourcesRow) UnsetLinkedDataSourceOther()`

UnsetLinkedDataSourceOther ensures that no value is present for LinkedDataSourceOther, not even an explicit nil
### GetLinkedLinkageCompleteness

`func (o *RwdSourcesRow) GetLinkedLinkageCompleteness() string`

GetLinkedLinkageCompleteness returns the LinkedLinkageCompleteness field if non-nil, zero value otherwise.

### GetLinkedLinkageCompletenessOk

`func (o *RwdSourcesRow) GetLinkedLinkageCompletenessOk() (*string, bool)`

GetLinkedLinkageCompletenessOk returns a tuple with the LinkedLinkageCompleteness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedLinkageCompleteness

`func (o *RwdSourcesRow) SetLinkedLinkageCompleteness(v string)`

SetLinkedLinkageCompleteness sets LinkedLinkageCompleteness field to given value.

### HasLinkedLinkageCompleteness

`func (o *RwdSourcesRow) HasLinkedLinkageCompleteness() bool`

HasLinkedLinkageCompleteness returns a boolean if a field has been set.

### SetLinkedLinkageCompletenessNil

`func (o *RwdSourcesRow) SetLinkedLinkageCompletenessNil(b bool)`

 SetLinkedLinkageCompletenessNil sets the value for LinkedLinkageCompleteness to be an explicit nil

### UnsetLinkedLinkageCompleteness
`func (o *RwdSourcesRow) UnsetLinkedLinkageCompleteness()`

UnsetLinkedLinkageCompleteness ensures that no value is present for LinkedLinkageCompleteness, not even an explicit nil
### GetLinkedLinkageVariable

`func (o *RwdSourcesRow) GetLinkedLinkageVariable() string`

GetLinkedLinkageVariable returns the LinkedLinkageVariable field if non-nil, zero value otherwise.

### GetLinkedLinkageVariableOk

`func (o *RwdSourcesRow) GetLinkedLinkageVariableOk() (*string, bool)`

GetLinkedLinkageVariableOk returns a tuple with the LinkedLinkageVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedLinkageVariable

`func (o *RwdSourcesRow) SetLinkedLinkageVariable(v string)`

SetLinkedLinkageVariable sets LinkedLinkageVariable field to given value.

### HasLinkedLinkageVariable

`func (o *RwdSourcesRow) HasLinkedLinkageVariable() bool`

HasLinkedLinkageVariable returns a boolean if a field has been set.

### SetLinkedLinkageVariableNil

`func (o *RwdSourcesRow) SetLinkedLinkageVariableNil(b bool)`

 SetLinkedLinkageVariableNil sets the value for LinkedLinkageVariable to be an explicit nil

### UnsetLinkedLinkageVariable
`func (o *RwdSourcesRow) UnsetLinkedLinkageVariable()`

UnsetLinkedLinkageVariable ensures that no value is present for LinkedLinkageVariable, not even an explicit nil
### GetLinkedLinkageStrategy

`func (o *RwdSourcesRow) GetLinkedLinkageStrategy() string`

GetLinkedLinkageStrategy returns the LinkedLinkageStrategy field if non-nil, zero value otherwise.

### GetLinkedLinkageStrategyOk

`func (o *RwdSourcesRow) GetLinkedLinkageStrategyOk() (*string, bool)`

GetLinkedLinkageStrategyOk returns a tuple with the LinkedLinkageStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedLinkageStrategy

`func (o *RwdSourcesRow) SetLinkedLinkageStrategy(v string)`

SetLinkedLinkageStrategy sets LinkedLinkageStrategy field to given value.

### HasLinkedLinkageStrategy

`func (o *RwdSourcesRow) HasLinkedLinkageStrategy() bool`

HasLinkedLinkageStrategy returns a boolean if a field has been set.

### SetLinkedLinkageStrategyNil

`func (o *RwdSourcesRow) SetLinkedLinkageStrategyNil(b bool)`

 SetLinkedLinkageStrategyNil sets the value for LinkedLinkageStrategy to be an explicit nil

### UnsetLinkedLinkageStrategy
`func (o *RwdSourcesRow) UnsetLinkedLinkageStrategy()`

UnsetLinkedLinkageStrategy ensures that no value is present for LinkedLinkageStrategy, not even an explicit nil
### GetDeclaredLinkageCount

`func (o *RwdSourcesRow) GetDeclaredLinkageCount() int32`

GetDeclaredLinkageCount returns the DeclaredLinkageCount field if non-nil, zero value otherwise.

### GetDeclaredLinkageCountOk

`func (o *RwdSourcesRow) GetDeclaredLinkageCountOk() (*int32, bool)`

GetDeclaredLinkageCountOk returns a tuple with the DeclaredLinkageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredLinkageCount

`func (o *RwdSourcesRow) SetDeclaredLinkageCount(v int32)`

SetDeclaredLinkageCount sets DeclaredLinkageCount field to given value.

### HasDeclaredLinkageCount

`func (o *RwdSourcesRow) HasDeclaredLinkageCount() bool`

HasDeclaredLinkageCount returns a boolean if a field has been set.

### SetDeclaredLinkageCountNil

`func (o *RwdSourcesRow) SetDeclaredLinkageCountNil(b bool)`

 SetDeclaredLinkageCountNil sets the value for DeclaredLinkageCount to be an explicit nil

### UnsetDeclaredLinkageCount
`func (o *RwdSourcesRow) UnsetDeclaredLinkageCount()`

UnsetDeclaredLinkageCount ensures that no value is present for DeclaredLinkageCount, not even an explicit nil
### GetLinkageState

`func (o *RwdSourcesRow) GetLinkageState() string`

GetLinkageState returns the LinkageState field if non-nil, zero value otherwise.

### GetLinkageStateOk

`func (o *RwdSourcesRow) GetLinkageStateOk() (*string, bool)`

GetLinkageStateOk returns a tuple with the LinkageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageState

`func (o *RwdSourcesRow) SetLinkageState(v string)`

SetLinkageState sets LinkageState field to given value.

### HasLinkageState

`func (o *RwdSourcesRow) HasLinkageState() bool`

HasLinkageState returns a boolean if a field has been set.

### SetLinkageStateNil

`func (o *RwdSourcesRow) SetLinkageStateNil(b bool)`

 SetLinkageStateNil sets the value for LinkageState to be an explicit nil

### UnsetLinkageState
`func (o *RwdSourcesRow) UnsetLinkageState()`

UnsetLinkageState ensures that no value is present for LinkageState, not even an explicit nil
### GetBiospecimenAccess

`func (o *RwdSourcesRow) GetBiospecimenAccess() bool`

GetBiospecimenAccess returns the BiospecimenAccess field if non-nil, zero value otherwise.

### GetBiospecimenAccessOk

`func (o *RwdSourcesRow) GetBiospecimenAccessOk() (*bool, bool)`

GetBiospecimenAccessOk returns a tuple with the BiospecimenAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiospecimenAccess

`func (o *RwdSourcesRow) SetBiospecimenAccess(v bool)`

SetBiospecimenAccess sets BiospecimenAccess field to given value.

### HasBiospecimenAccess

`func (o *RwdSourcesRow) HasBiospecimenAccess() bool`

HasBiospecimenAccess returns a boolean if a field has been set.

### SetBiospecimenAccessNil

`func (o *RwdSourcesRow) SetBiospecimenAccessNil(b bool)`

 SetBiospecimenAccessNil sets the value for BiospecimenAccess to be an explicit nil

### UnsetBiospecimenAccess
`func (o *RwdSourcesRow) UnsetBiospecimenAccess()`

UnsetBiospecimenAccess ensures that no value is present for BiospecimenAccess, not even an explicit nil
### GetAccessToSubjectDetails

`func (o *RwdSourcesRow) GetAccessToSubjectDetails() bool`

GetAccessToSubjectDetails returns the AccessToSubjectDetails field if non-nil, zero value otherwise.

### GetAccessToSubjectDetailsOk

`func (o *RwdSourcesRow) GetAccessToSubjectDetailsOk() (*bool, bool)`

GetAccessToSubjectDetailsOk returns a tuple with the AccessToSubjectDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToSubjectDetails

`func (o *RwdSourcesRow) SetAccessToSubjectDetails(v bool)`

SetAccessToSubjectDetails sets AccessToSubjectDetails field to given value.

### HasAccessToSubjectDetails

`func (o *RwdSourcesRow) HasAccessToSubjectDetails() bool`

HasAccessToSubjectDetails returns a boolean if a field has been set.

### SetAccessToSubjectDetailsNil

`func (o *RwdSourcesRow) SetAccessToSubjectDetailsNil(b bool)`

 SetAccessToSubjectDetailsNil sets the value for AccessToSubjectDetails to be an explicit nil

### UnsetAccessToSubjectDetails
`func (o *RwdSourcesRow) UnsetAccessToSubjectDetails()`

UnsetAccessToSubjectDetails ensures that no value is present for AccessToSubjectDetails, not even an explicit nil
### GetPossibilityOfDataValidation

`func (o *RwdSourcesRow) GetPossibilityOfDataValidation() bool`

GetPossibilityOfDataValidation returns the PossibilityOfDataValidation field if non-nil, zero value otherwise.

### GetPossibilityOfDataValidationOk

`func (o *RwdSourcesRow) GetPossibilityOfDataValidationOk() (*bool, bool)`

GetPossibilityOfDataValidationOk returns a tuple with the PossibilityOfDataValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibilityOfDataValidation

`func (o *RwdSourcesRow) SetPossibilityOfDataValidation(v bool)`

SetPossibilityOfDataValidation sets PossibilityOfDataValidation field to given value.

### HasPossibilityOfDataValidation

`func (o *RwdSourcesRow) HasPossibilityOfDataValidation() bool`

HasPossibilityOfDataValidation returns a boolean if a field has been set.

### SetPossibilityOfDataValidationNil

`func (o *RwdSourcesRow) SetPossibilityOfDataValidationNil(b bool)`

 SetPossibilityOfDataValidationNil sets the value for PossibilityOfDataValidation to be an explicit nil

### UnsetPossibilityOfDataValidation
`func (o *RwdSourcesRow) UnsetPossibilityOfDataValidation()`

UnsetPossibilityOfDataValidation ensures that no value is present for PossibilityOfDataValidation, not even an explicit nil
### GetRecordsArePreservedIndefinitely

`func (o *RwdSourcesRow) GetRecordsArePreservedIndefinitely() bool`

GetRecordsArePreservedIndefinitely returns the RecordsArePreservedIndefinitely field if non-nil, zero value otherwise.

### GetRecordsArePreservedIndefinitelyOk

`func (o *RwdSourcesRow) GetRecordsArePreservedIndefinitelyOk() (*bool, bool)`

GetRecordsArePreservedIndefinitelyOk returns a tuple with the RecordsArePreservedIndefinitely field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsArePreservedIndefinitely

`func (o *RwdSourcesRow) SetRecordsArePreservedIndefinitely(v bool)`

SetRecordsArePreservedIndefinitely sets RecordsArePreservedIndefinitely field to given value.

### HasRecordsArePreservedIndefinitely

`func (o *RwdSourcesRow) HasRecordsArePreservedIndefinitely() bool`

HasRecordsArePreservedIndefinitely returns a boolean if a field has been set.

### SetRecordsArePreservedIndefinitelyNil

`func (o *RwdSourcesRow) SetRecordsArePreservedIndefinitelyNil(b bool)`

 SetRecordsArePreservedIndefinitelyNil sets the value for RecordsArePreservedIndefinitely to be an explicit nil

### UnsetRecordsArePreservedIndefinitely
`func (o *RwdSourcesRow) UnsetRecordsArePreservedIndefinitely()`

UnsetRecordsArePreservedIndefinitely ensures that no value is present for RecordsArePreservedIndefinitely, not even an explicit nil
### GetDataSourcePreservationLengthYears

`func (o *RwdSourcesRow) GetDataSourcePreservationLengthYears() string`

GetDataSourcePreservationLengthYears returns the DataSourcePreservationLengthYears field if non-nil, zero value otherwise.

### GetDataSourcePreservationLengthYearsOk

`func (o *RwdSourcesRow) GetDataSourcePreservationLengthYearsOk() (*string, bool)`

GetDataSourcePreservationLengthYearsOk returns a tuple with the DataSourcePreservationLengthYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourcePreservationLengthYears

`func (o *RwdSourcesRow) SetDataSourcePreservationLengthYears(v string)`

SetDataSourcePreservationLengthYears sets DataSourcePreservationLengthYears field to given value.

### HasDataSourcePreservationLengthYears

`func (o *RwdSourcesRow) HasDataSourcePreservationLengthYears() bool`

HasDataSourcePreservationLengthYears returns a boolean if a field has been set.

### SetDataSourcePreservationLengthYearsNil

`func (o *RwdSourcesRow) SetDataSourcePreservationLengthYearsNil(b bool)`

 SetDataSourcePreservationLengthYearsNil sets the value for DataSourcePreservationLengthYears to be an explicit nil

### UnsetDataSourcePreservationLengthYears
`func (o *RwdSourcesRow) UnsetDataSourcePreservationLengthYears()`

UnsetDataSourcePreservationLengthYears ensures that no value is present for DataSourcePreservationLengthYears, not even an explicit nil
### GetApprovalForPublication

`func (o *RwdSourcesRow) GetApprovalForPublication() bool`

GetApprovalForPublication returns the ApprovalForPublication field if non-nil, zero value otherwise.

### GetApprovalForPublicationOk

`func (o *RwdSourcesRow) GetApprovalForPublicationOk() (*bool, bool)`

GetApprovalForPublicationOk returns a tuple with the ApprovalForPublication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalForPublication

`func (o *RwdSourcesRow) SetApprovalForPublication(v bool)`

SetApprovalForPublication sets ApprovalForPublication field to given value.

### HasApprovalForPublication

`func (o *RwdSourcesRow) HasApprovalForPublication() bool`

HasApprovalForPublication returns a boolean if a field has been set.

### SetApprovalForPublicationNil

`func (o *RwdSourcesRow) SetApprovalForPublicationNil(b bool)`

 SetApprovalForPublicationNil sets the value for ApprovalForPublication to be an explicit nil

### UnsetApprovalForPublication
`func (o *RwdSourcesRow) UnsetApprovalForPublication()`

UnsetApprovalForPublication ensures that no value is present for ApprovalForPublication, not even an explicit nil
### GetInformedConsentForUseOfDataForResearch

`func (o *RwdSourcesRow) GetInformedConsentForUseOfDataForResearch() string`

GetInformedConsentForUseOfDataForResearch returns the InformedConsentForUseOfDataForResearch field if non-nil, zero value otherwise.

### GetInformedConsentForUseOfDataForResearchOk

`func (o *RwdSourcesRow) GetInformedConsentForUseOfDataForResearchOk() (*string, bool)`

GetInformedConsentForUseOfDataForResearchOk returns a tuple with the InformedConsentForUseOfDataForResearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformedConsentForUseOfDataForResearch

`func (o *RwdSourcesRow) SetInformedConsentForUseOfDataForResearch(v string)`

SetInformedConsentForUseOfDataForResearch sets InformedConsentForUseOfDataForResearch field to given value.

### HasInformedConsentForUseOfDataForResearch

`func (o *RwdSourcesRow) HasInformedConsentForUseOfDataForResearch() bool`

HasInformedConsentForUseOfDataForResearch returns a boolean if a field has been set.

### SetInformedConsentForUseOfDataForResearchNil

`func (o *RwdSourcesRow) SetInformedConsentForUseOfDataForResearchNil(b bool)`

 SetInformedConsentForUseOfDataForResearchNil sets the value for InformedConsentForUseOfDataForResearch to be an explicit nil

### UnsetInformedConsentForUseOfDataForResearch
`func (o *RwdSourcesRow) UnsetInformedConsentForUseOfDataForResearch()`

UnsetInformedConsentForUseOfDataForResearch ensures that no value is present for InformedConsentForUseOfDataForResearch, not even an explicit nil
### GetGovernanceDetailsLink

`func (o *RwdSourcesRow) GetGovernanceDetailsLink() string`

GetGovernanceDetailsLink returns the GovernanceDetailsLink field if non-nil, zero value otherwise.

### GetGovernanceDetailsLinkOk

`func (o *RwdSourcesRow) GetGovernanceDetailsLinkOk() (*string, bool)`

GetGovernanceDetailsLinkOk returns a tuple with the GovernanceDetailsLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernanceDetailsLink

`func (o *RwdSourcesRow) SetGovernanceDetailsLink(v string)`

SetGovernanceDetailsLink sets GovernanceDetailsLink field to given value.

### HasGovernanceDetailsLink

`func (o *RwdSourcesRow) HasGovernanceDetailsLink() bool`

HasGovernanceDetailsLink returns a boolean if a field has been set.

### SetGovernanceDetailsLinkNil

`func (o *RwdSourcesRow) SetGovernanceDetailsLinkNil(b bool)`

 SetGovernanceDetailsLinkNil sets the value for GovernanceDetailsLink to be an explicit nil

### UnsetGovernanceDetailsLink
`func (o *RwdSourcesRow) UnsetGovernanceDetailsLink()`

UnsetGovernanceDetailsLink ensures that no value is present for GovernanceDetailsLink, not even an explicit nil
### GetDeclaredCapabilityCount

`func (o *RwdSourcesRow) GetDeclaredCapabilityCount() int32`

GetDeclaredCapabilityCount returns the DeclaredCapabilityCount field if non-nil, zero value otherwise.

### GetDeclaredCapabilityCountOk

`func (o *RwdSourcesRow) GetDeclaredCapabilityCountOk() (*int32, bool)`

GetDeclaredCapabilityCountOk returns a tuple with the DeclaredCapabilityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredCapabilityCount

`func (o *RwdSourcesRow) SetDeclaredCapabilityCount(v int32)`

SetDeclaredCapabilityCount sets DeclaredCapabilityCount field to given value.

### HasDeclaredCapabilityCount

`func (o *RwdSourcesRow) HasDeclaredCapabilityCount() bool`

HasDeclaredCapabilityCount returns a boolean if a field has been set.

### SetDeclaredCapabilityCountNil

`func (o *RwdSourcesRow) SetDeclaredCapabilityCountNil(b bool)`

 SetDeclaredCapabilityCountNil sets the value for DeclaredCapabilityCount to be an explicit nil

### UnsetDeclaredCapabilityCount
`func (o *RwdSourcesRow) UnsetDeclaredCapabilityCount()`

UnsetDeclaredCapabilityCount ensures that no value is present for DeclaredCapabilityCount, not even an explicit nil
### GetDeclarableCapabilityCount

`func (o *RwdSourcesRow) GetDeclarableCapabilityCount() int32`

GetDeclarableCapabilityCount returns the DeclarableCapabilityCount field if non-nil, zero value otherwise.

### GetDeclarableCapabilityCountOk

`func (o *RwdSourcesRow) GetDeclarableCapabilityCountOk() (*int32, bool)`

GetDeclarableCapabilityCountOk returns a tuple with the DeclarableCapabilityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarableCapabilityCount

`func (o *RwdSourcesRow) SetDeclarableCapabilityCount(v int32)`

SetDeclarableCapabilityCount sets DeclarableCapabilityCount field to given value.

### HasDeclarableCapabilityCount

`func (o *RwdSourcesRow) HasDeclarableCapabilityCount() bool`

HasDeclarableCapabilityCount returns a boolean if a field has been set.

### SetDeclarableCapabilityCountNil

`func (o *RwdSourcesRow) SetDeclarableCapabilityCountNil(b bool)`

 SetDeclarableCapabilityCountNil sets the value for DeclarableCapabilityCount to be an explicit nil

### UnsetDeclarableCapabilityCount
`func (o *RwdSourcesRow) UnsetDeclarableCapabilityCount()`

UnsetDeclarableCapabilityCount ensures that no value is present for DeclarableCapabilityCount, not even an explicit nil
### GetRegisterSourceCount

`func (o *RwdSourcesRow) GetRegisterSourceCount() int32`

GetRegisterSourceCount returns the RegisterSourceCount field if non-nil, zero value otherwise.

### GetRegisterSourceCountOk

`func (o *RwdSourcesRow) GetRegisterSourceCountOk() (*int32, bool)`

GetRegisterSourceCountOk returns a tuple with the RegisterSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterSourceCount

`func (o *RwdSourcesRow) SetRegisterSourceCount(v int32)`

SetRegisterSourceCount sets RegisterSourceCount field to given value.

### HasRegisterSourceCount

`func (o *RwdSourcesRow) HasRegisterSourceCount() bool`

HasRegisterSourceCount returns a boolean if a field has been set.

### SetRegisterSourceCountNil

`func (o *RwdSourcesRow) SetRegisterSourceCountNil(b bool)`

 SetRegisterSourceCountNil sets the value for RegisterSourceCount to be an explicit nil

### UnsetRegisterSourceCount
`func (o *RwdSourcesRow) UnsetRegisterSourceCount()`

UnsetRegisterSourceCount ensures that no value is present for RegisterSourceCount, not even an explicit nil
### GetStudyCount

`func (o *RwdSourcesRow) GetStudyCount() int32`

GetStudyCount returns the StudyCount field if non-nil, zero value otherwise.

### GetStudyCountOk

`func (o *RwdSourcesRow) GetStudyCountOk() (*int32, bool)`

GetStudyCountOk returns a tuple with the StudyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyCount

`func (o *RwdSourcesRow) SetStudyCount(v int32)`

SetStudyCount sets StudyCount field to given value.

### HasStudyCount

`func (o *RwdSourcesRow) HasStudyCount() bool`

HasStudyCount returns a boolean if a field has been set.

### SetStudyCountNil

`func (o *RwdSourcesRow) SetStudyCountNil(b bool)`

 SetStudyCountNil sets the value for StudyCount to be an explicit nil

### UnsetStudyCount
`func (o *RwdSourcesRow) UnsetStudyCount()`

UnsetStudyCount ensures that no value is present for StudyCount, not even an explicit nil
### GetStudyRegisterTotal

`func (o *RwdSourcesRow) GetStudyRegisterTotal() int32`

GetStudyRegisterTotal returns the StudyRegisterTotal field if non-nil, zero value otherwise.

### GetStudyRegisterTotalOk

`func (o *RwdSourcesRow) GetStudyRegisterTotalOk() (*int32, bool)`

GetStudyRegisterTotalOk returns a tuple with the StudyRegisterTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyRegisterTotal

`func (o *RwdSourcesRow) SetStudyRegisterTotal(v int32)`

SetStudyRegisterTotal sets StudyRegisterTotal field to given value.

### HasStudyRegisterTotal

`func (o *RwdSourcesRow) HasStudyRegisterTotal() bool`

HasStudyRegisterTotal returns a boolean if a field has been set.

### SetStudyRegisterTotalNil

`func (o *RwdSourcesRow) SetStudyRegisterTotalNil(b bool)`

 SetStudyRegisterTotalNil sets the value for StudyRegisterTotal to be an explicit nil

### UnsetStudyRegisterTotal
`func (o *RwdSourcesRow) UnsetStudyRegisterTotal()`

UnsetStudyRegisterTotal ensures that no value is present for StudyRegisterTotal, not even an explicit nil
### GetStudiesNamingARegisteredSource

`func (o *RwdSourcesRow) GetStudiesNamingARegisteredSource() int32`

GetStudiesNamingARegisteredSource returns the StudiesNamingARegisteredSource field if non-nil, zero value otherwise.

### GetStudiesNamingARegisteredSourceOk

`func (o *RwdSourcesRow) GetStudiesNamingARegisteredSourceOk() (*int32, bool)`

GetStudiesNamingARegisteredSourceOk returns a tuple with the StudiesNamingARegisteredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudiesNamingARegisteredSource

`func (o *RwdSourcesRow) SetStudiesNamingARegisteredSource(v int32)`

SetStudiesNamingARegisteredSource sets StudiesNamingARegisteredSource field to given value.

### HasStudiesNamingARegisteredSource

`func (o *RwdSourcesRow) HasStudiesNamingARegisteredSource() bool`

HasStudiesNamingARegisteredSource returns a boolean if a field has been set.

### SetStudiesNamingARegisteredSourceNil

`func (o *RwdSourcesRow) SetStudiesNamingARegisteredSourceNil(b bool)`

 SetStudiesNamingARegisteredSourceNil sets the value for StudiesNamingARegisteredSource to be an explicit nil

### UnsetStudiesNamingARegisteredSource
`func (o *RwdSourcesRow) UnsetStudiesNamingARegisteredSource()`

UnsetStudiesNamingARegisteredSource ensures that no value is present for StudiesNamingARegisteredSource, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdSourcesRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdSourcesRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdSourcesRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdSourcesRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdSourcesRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdSourcesRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdSourcesRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdSourcesRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdSourcesRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdSourcesRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdSourcesRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdSourcesRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdSourcesRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdSourcesRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdSourcesRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdSourcesRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdSourcesRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdSourcesRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdSourcesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdSourcesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdSourcesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdSourcesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdSourcesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdSourcesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


