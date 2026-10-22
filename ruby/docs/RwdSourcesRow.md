# PharmaObservatory::RwdSourcesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data_source_id** | **String** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] |
| **data_source_name** | **String** | The data source&#39;s name as the data holder registered it; null on a stub entry with no name. | [optional] |
| **data_source_acronym** | **String** | The acronym the data holder registered for the source. | [optional] |
| **data_holder** | **String** | The organisation that holds the data, as typed in the register (free text, not resolved to a company). | [optional] |
| **data_source_type** | **String** | The data source type(s) the holder ticked, as one comma-separated string as the register exports it; rwd_source_types has one row per declared type. | [optional] |
| **data_source_type_other** | **String** | The holder&#39;s free-text description when it chose the register&#39;s other data source type. | [optional] |
| **data_source_domain** | **String** | The register&#39;s data source domain (for example Claims, Primary care, Secondary care). | [optional] |
| **care_setting** | **String** | The care setting the holder declares for the source. | [optional] |
| **main_financial_support** | **String** | The main financial support of the data source as the holder declared it. | [optional] |
| **country_names** | **Array&lt;String&gt;** | The countries the source covers, one array element per register country, split by the register&#39;s own country list (so Iran, Islamic Republic of stays one country). Matches the group-by-able membership filter &#39;Spain&#39; &#x3D; ANY(country_names). | [optional] |
| **unmatched_country_fragments** | **Array&lt;String&gt;** | Country text items the country vocabulary of this build does not hold. | [optional] |
| **country_count** | **Integer** | The number of countries in country_names. | [optional] |
| **catalogue_url** | **String** | The address of this entry on the HMA-EMA catalogue site (the register&#39;s URL LINK). | [optional] |
| **data_source_website** | **String** | The data source&#39;s own website as the holder registered it. | [optional] |
| **is_darwin_eu_data_partner** | **Boolean** | Whether the holder declares the source a DARWIN EU data partner (the EMA-coordinated European network of real-world data sources); null when unanswered. | [optional] |
| **is_qualified** | **Boolean** | The register&#39;s Data source qualification answer; true when the holder declares a qualification, null when unanswered. | [optional] |
| **refresh_cadence** | **String** | How often the holder says the database takes new data, as typed (the register&#39;s Data source refresh). | [optional] |
| **last_refreshed_on** | **Date** | The date the holder says the database last took new data (the register&#39;s Data source last refresh). | [optional] |
| **record_updated_at** | **Time** | When the registrant last edited this catalogue entry. A different clock from last_refreshed_on: it dates the web form, not the database. | [optional] |
| **record_first_published_at** | **Time** | When the catalogue entry was first published. | [optional] |
| **is_outdated** | **Boolean** | EMA&#39;s own Outdated flag, carried through untouched; the register flags none of the data sources, so read false as EMA&#39;s answer and not as a currency finding. | [optional] |
| **record_updated_over_12_months_ago** | **Boolean** | True when the catalogue entry was last edited more than twelve months before the snapshot. About the web form only; it says nothing about whether the database has taken new data (see database_refresh_state). | [optional] |
| **refresh_interval_days** | **Integer** | The interval in days the longest declared refresh cadence is read as; the threshold the database_refresh_state verdict used. Null when no declared cadence was recognised. | [optional] |
| **refresh_interval_basis** | **String** | Which declared cadence refresh_interval_days came from. | [optional] |
| **refresh_interval_days_shortest** | **Integer** | The interval in days of the strictest declared cadence, for a source that declared several. | [optional] |
| **refresh_interval_basis_shortest** | **String** | Which declared cadence refresh_interval_days_shortest came from. | [optional] |
| **refresh_cadence_declared** | **Array&lt;String&gt;** | Every cadence the holder declared, in declared order; the month names of one answer count as one declaration. Positionally paired with refresh_cadence_interval_days and refresh_cadence_is_month_names. Empty when none was declared. | [optional] |
| **refresh_cadence_interval_days** | **Array&lt;String&gt;** | The interval in days each declared cadence was read as, positionally paired with refresh_cadence_declared, rendered as text because the catalog has only text arrays. | [optional] |
| **refresh_cadence_is_month_names** | **Array&lt;String&gt;** | Whether each declared cadence is a list of month names (priced as 365 days over the number of months named); rendered as text (true or false per element) because the catalog has only text arrays. | [optional] |
| **days_since_last_refresh** | **Integer** | Days from last_refreshed_on to the catalogue snapshot date; null when no refresh date is declared. | [optional] |
| **database_refresh_state** | **String** | Adherence verdict about the holder&#39;s own promise: the declared cadence against the declared last refresh, measured to the snapshot date. Not a currency answer: use refresh_recency_band for how old the data is. | [optional] |
| **refresh_state_strict** | **String** | The same adherence verdict as database_refresh_state, read against the strictest declared cadence instead of the longest. | [optional] |
| **refresh_recency_band** | **String** | Absolute age of the last database refresh at the snapshot date, in disjoint bands, with no reference to any declared cadence. Sources that declared no refresh date have their own band; they are unanswered, not stale. | [optional] |
| **collection_last_band** | **String** | The band of the declared last collection date (the observable window), banded with the same ruler as refresh_recency_band; null when no collection date is declared. | [optional] |
| **collection_contradicts_refresh_band** | **Boolean** | True when the last collection date and the last refresh date fall in different currency bands. A fact to print beside the currency screen, not a control. | [optional] |
| **prescriptions_of_medicines** | **Boolean** | Whether the data holder declares that prescriptions of medicines are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **prescriptions_vocabulary** | **String** | The coding scheme(s) the holder declares for prescriptions, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] |
| **prescriptions_vocabulary_other** | **String** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] |
| **dispensing_of_medicines** | **Boolean** | Whether the data holder declares that dispensing of medicines is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **dispensing_vocabulary** | **String** | The coding scheme(s) the holder declares for dispensing, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] |
| **dispensing_vocabulary_other** | **String** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] |
| **medicinal_product_vocabulary** | **String** | The coding scheme(s) the holder declares for medicinal product information, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] |
| **medicinal_product_vocabulary_other** | **String** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] |
| **indication_for_use** | **Boolean** | Whether the data holder declares that the indication for use of a medicine is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **indication_vocabulary** | **String** | The coding scheme(s) the holder declares for the indication for use, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] |
| **indication_vocabulary_other** | **String** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] |
| **diagnostic_codes** | **Boolean** | Whether the data holder declares that diagnostic codes are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **diagnosis_medical_event_vocabulary** | **String** | The coding scheme(s) the holder declares for diagnoses and medical events, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] |
| **diagnosis_medical_event_vocabulary_other** | **String** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] |
| **hospital_admission_and_or_discharge** | **Boolean** | Whether the data holder declares that hospital admission and/or discharge is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **icu_admission** | **Boolean** | Whether the data holder declares that intensive care unit (ICU) admission is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **cause_of_death** | **Boolean** | Whether the data holder declares that cause of death is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **cause_of_death_vocabulary** | **String** | The coding scheme(s) the holder declares for cause of death, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] |
| **cause_of_death_vocabulary_other** | **String** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] |
| **disease_information** | **Boolean** | Whether the data holder declares that disease information is collected: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **rare_diseases** | **Boolean** | Whether the data holder declares that rare diseases are identifiable: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **pregnancy_and_or_neonates** | **Boolean** | Whether the data holder declares that pregnancy and/or neonates are covered: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **advance_therapy_medicinal_products_atmp** | **Boolean** | Whether the data holder declares that advanced therapy medicinal products (ATMP) are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **contraception** | **Boolean** | Whether the data holder declares that contraception is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **medical_devices** | **Boolean** | Whether the data holder declares that medical devices are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **administration_of_vaccines** | **Boolean** | Whether the data holder declares that administration of vaccines is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **procedures** | **Boolean** | Whether the data holder declares that procedures are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **procedures_vocabulary** | **String** | The coding scheme(s) the holder declares for procedures, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] |
| **procedures_vocabulary_other** | **String** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] |
| **clinical_measurements** | **Boolean** | Whether the data holder declares that clinical measurements are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **healthcare_provider** | **Boolean** | Whether the data holder declares that the healthcare provider is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **genetic_data** | **Boolean** | Whether the data holder declares that genetic data are held: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **genetic_data_vocabulary** | **String** | The coding scheme(s) the holder declares for genetic data, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] |
| **genetic_data_vocabulary_other** | **String** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] |
| **biomarker_data** | **Boolean** | Whether the data holder declares that biomarker data are held: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **biomarker_data_vocabulary** | **String** | The coding scheme(s) the holder declares for biomarker data, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] |
| **biomarker_vocabulary_other** | **String** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] |
| **patient_reported_outcomes** | **Boolean** | Whether the data holder declares that patient-reported outcomes are held: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **patient_generated_data** | **Boolean** | Whether the data holder declares that patient-generated data are held: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **units_of_healthcare_utilisation** | **Boolean** | Whether the data holder declares that units of healthcare utilisation are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **unique_identifier_for_persons** | **Boolean** | Whether the data holder declares that a unique identifier for persons exists: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **quality_of_life_measurements** | **Boolean** | Whether the data holder declares that quality of life measurements are held: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **quality_of_life_measurements_vocabulary** | **String** | The coding scheme(s) the holder declares for quality of life measurements, comma-joined as typed; rwd_source_vocabularies has one row per scheme. | [optional] |
| **quality_of_life_measurements_other** | **String** | The holder&#39;s free-text description of the other coding scheme it declares for this capability. | [optional] |
| **lifestyle_factors** | **Boolean** | Whether the data holder declares that lifestyle factors are recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **sociodemographic_information** | **Boolean** | Whether the data holder declares that sociodemographic information is recorded: true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **population_size** | **Integer** | The population size the holder declared for the source (all registered persons). | [optional] |
| **active_population_size** | **Integer** | The active population size the holder declared for the source. | [optional] |
| **median_observation_time_years** | **Float** | The median observation time in years exactly as the holder declared it. Never corrected: see median_observation_time_plausibility. | [optional] |
| **median_observation_time_active_years** | **Float** | The median observation time in years for the active population exactly as the holder declared it. Never corrected: see median_observation_time_active_plausibility. | [optional] |
| **established_on** | **Date** | The date the data source was established. | [optional] |
| **collection_first_on** | **Date** | The date of the first data collection. | [optional] |
| **collection_last_on** | **Date** | The date of the last data collection; null on most sources and a future date on some, so it is a fact to read and not a field to sort on. | [optional] |
| **declared_history_span_years** | **Float** | The data window in years the holder&#39;s own dates give the database (earliest of established and first collection, to last collection or the snapshot). | [optional] |
| **median_observation_time_plausibility** | **String** | Classifies the declared median observation time without correcting it: plausible, exceeds_ceiling (more than 100 years), exceeds_declared_span (more than a year longer than the holder&#39;s own data window) or not_declared. | [optional] |
| **median_observation_time_active_plausibility** | **String** | The same classification for the active-population median observation time. | [optional] |
| **estimated_population_coverage_text** | **String** | The estimated percentage of the population covered, as free text exactly as typed (values include Unsure and Not available); deliberately not parsed into a number. | [optional] |
| **age_groups** | **String** | The age groups the holder declares the source covers, as one comma-separated string. | [optional] |
| **age_band_count** | **Integer** | The number of declared age bands (rwd_source_age_bands has one row per band). | [optional] |
| **has_population_breakdown** | **Boolean** | Whether the holder published a population size breakdown by age band at all. | [optional] |
| **has_active_population_breakdown** | **Boolean** | Whether the holder published an active population breakdown by age band at all. | [optional] |
| **age_bands_are_nested** | **Boolean** | True when the declared age bands overlap (an open-ended elderly band declared beside narrower bands inside it). Age bands from such a source must not be summed; no sum is published. | [optional] |
| **cdm_mapping** | **Boolean** | Whether the holder declares the data are mapped to a common data model (CDM), such as OMOP. Read it with cdm_readiness, which says whether the transform is finished. | [optional] |
| **etl_status** | **String** | The extract-transform-load (ETL) status of the CDM mapping as typed; one comma-joined answer per declared mapping (rwd_source_etl_mappings has one row per mapping). | [optional] |
| **etl_cdm_version** | **String** | The CDM version(s) declared for the ETL, comma-joined as typed. | [optional] |
| **etl_frequency** | **String** | The ETL frequency declared, comma-joined as typed. | [optional] |
| **etl_specifications_file** | **String** | The ETL specifications file the holder attached, if any (empty at source on the registered sources). | [optional] |
| **declared_mapping_count** | **Integer** | How many CDM mappings the holder declared, after register join artefacts are removed; 0 when the question was never answered. | [optional] |
| **cdm_readiness** | **String** | The CDM fact as one closed label that carries the word the holder wrote about the transform, so cdm_mapping alone is never read as ready. Exhaustive and never null: not_declared, no_mapping, mapping_without_etl_status, etl_completed, etl_in_progress, etl_planned, etl_not_done, etl_mixed (several mappings whose statuses differ or one has none) or etl_other (an unrecognised word). | [optional] |
| **linkage** | **Boolean** | Whether the holder declares that the data can be linked to other data sources. | [optional] |
| **has_pre_linked_declaration** | **Boolean** | Whether the holder gave a description of a linkage already made (pre-linked). | [optional] |
| **has_possible_linkage_declaration** | **Boolean** | Whether the holder gave a description of a linkage that is possible but not done. | [optional] |
| **linked_data_source** | **String** | The linked data source name(s) as typed, comma-joined; rwd_source_linkages has one row per declared linkage. | [optional] |
| **linked_data_source_other** | **String** | The linked data source named in free text when the holder chose Other data source. | [optional] |
| **linked_linkage_completeness** | **String** | The linkage completeness the holder declared, as typed. | [optional] |
| **linked_linkage_variable** | **String** | The linkage variable(s) the holder declared, as typed. | [optional] |
| **linked_linkage_strategy** | **String** | The linkage strategy the holder declared, as typed. | [optional] |
| **declared_linkage_count** | **Integer** | How many linkages the holder declared; 0 when none. | [optional] |
| **linkage_state** | **String** | Five states, exhaustive and never null: pre_linked, possible_not_done, declared_without_detail, no_linkage, not_declared. Pre-linked and possible linkage are kept apart on purpose. | [optional] |
| **biospecimen_access** | **Boolean** | Whether biospecimens can be accessed through the source; true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **access_to_subject_details** | **Boolean** | Whether the holder can access subject details; true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **possibility_of_data_validation** | **Boolean** | Whether the data can be validated; true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **records_are_preserved_indefinitely** | **Boolean** | Whether records are preserved indefinitely; true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **data_source_preservation_length_years** | **String** | How long the data are preserved, in years, as typed. | [optional] |
| **approval_for_publication** | **Boolean** | Whether approval for publication is needed; true when the data holder declares it, false when the holder declares it absent, null when the holder did not answer (never read null as no). | [optional] |
| **informed_consent_for_use_of_data_for_research** | **String** | The informed consent regime for research use as the holder declared it, as typed. | [optional] |
| **governance_details_link** | **String** | A link to the source&#39;s governance details, as the holder registered it. | [optional] |
| **declared_capability_count** | **Integer** | How many of the 35 capability and governance questions the holder answered (a count of answered form fields, not a measure of the data held). | [optional] |
| **declarable_capability_count** | **Integer** | How many such questions the form offers (35); the denominator of declared_capability_count. | [optional] |
| **register_source_count** | **Integer** | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional] |
| **study_count** | **Integer** | How many registered studies named this source. A registered study NAMED the source; it is never evidence that the source holds data on a condition. Read it with study_register_total and studies_naming_a_registered_source, because a zero is not a finding that the source is unused. | [optional] |
| **study_register_total** | **Integer** | The number of studies in the whole study register at the snapshot. | [optional] |
| **studies_naming_a_registered_source** | **Integer** | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdSourcesRow.new(
  data_source_id: null,
  data_source_name: null,
  data_source_acronym: null,
  data_holder: null,
  data_source_type: null,
  data_source_type_other: null,
  data_source_domain: null,
  care_setting: null,
  main_financial_support: null,
  country_names: null,
  unmatched_country_fragments: null,
  country_count: null,
  catalogue_url: null,
  data_source_website: null,
  is_darwin_eu_data_partner: null,
  is_qualified: null,
  refresh_cadence: null,
  last_refreshed_on: null,
  record_updated_at: null,
  record_first_published_at: null,
  is_outdated: null,
  record_updated_over_12_months_ago: null,
  refresh_interval_days: null,
  refresh_interval_basis: null,
  refresh_interval_days_shortest: null,
  refresh_interval_basis_shortest: null,
  refresh_cadence_declared: null,
  refresh_cadence_interval_days: null,
  refresh_cadence_is_month_names: null,
  days_since_last_refresh: null,
  database_refresh_state: null,
  refresh_state_strict: null,
  refresh_recency_band: null,
  collection_last_band: null,
  collection_contradicts_refresh_band: null,
  prescriptions_of_medicines: null,
  prescriptions_vocabulary: null,
  prescriptions_vocabulary_other: null,
  dispensing_of_medicines: null,
  dispensing_vocabulary: null,
  dispensing_vocabulary_other: null,
  medicinal_product_vocabulary: null,
  medicinal_product_vocabulary_other: null,
  indication_for_use: null,
  indication_vocabulary: null,
  indication_vocabulary_other: null,
  diagnostic_codes: null,
  diagnosis_medical_event_vocabulary: null,
  diagnosis_medical_event_vocabulary_other: null,
  hospital_admission_and_or_discharge: null,
  icu_admission: null,
  cause_of_death: null,
  cause_of_death_vocabulary: null,
  cause_of_death_vocabulary_other: null,
  disease_information: null,
  rare_diseases: null,
  pregnancy_and_or_neonates: null,
  advance_therapy_medicinal_products_atmp: null,
  contraception: null,
  medical_devices: null,
  administration_of_vaccines: null,
  procedures: null,
  procedures_vocabulary: null,
  procedures_vocabulary_other: null,
  clinical_measurements: null,
  healthcare_provider: null,
  genetic_data: null,
  genetic_data_vocabulary: null,
  genetic_data_vocabulary_other: null,
  biomarker_data: null,
  biomarker_data_vocabulary: null,
  biomarker_vocabulary_other: null,
  patient_reported_outcomes: null,
  patient_generated_data: null,
  units_of_healthcare_utilisation: null,
  unique_identifier_for_persons: null,
  quality_of_life_measurements: null,
  quality_of_life_measurements_vocabulary: null,
  quality_of_life_measurements_other: null,
  lifestyle_factors: null,
  sociodemographic_information: null,
  population_size: null,
  active_population_size: null,
  median_observation_time_years: null,
  median_observation_time_active_years: null,
  established_on: null,
  collection_first_on: null,
  collection_last_on: null,
  declared_history_span_years: null,
  median_observation_time_plausibility: null,
  median_observation_time_active_plausibility: null,
  estimated_population_coverage_text: null,
  age_groups: null,
  age_band_count: null,
  has_population_breakdown: null,
  has_active_population_breakdown: null,
  age_bands_are_nested: null,
  cdm_mapping: null,
  etl_status: null,
  etl_cdm_version: null,
  etl_frequency: null,
  etl_specifications_file: null,
  declared_mapping_count: null,
  cdm_readiness: null,
  linkage: null,
  has_pre_linked_declaration: null,
  has_possible_linkage_declaration: null,
  linked_data_source: null,
  linked_data_source_other: null,
  linked_linkage_completeness: null,
  linked_linkage_variable: null,
  linked_linkage_strategy: null,
  declared_linkage_count: null,
  linkage_state: null,
  biospecimen_access: null,
  access_to_subject_details: null,
  possibility_of_data_validation: null,
  records_are_preserved_indefinitely: null,
  data_source_preservation_length_years: null,
  approval_for_publication: null,
  informed_consent_for_use_of_data_for_research: null,
  governance_details_link: null,
  declared_capability_count: null,
  declarable_capability_count: null,
  register_source_count: null,
  study_count: null,
  study_register_total: null,
  studies_naming_a_registered_source: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

