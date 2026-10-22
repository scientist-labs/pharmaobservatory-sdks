# PharmaObservatory::RwdStudiesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **study_id** | **String** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] |
| **eu_pas_number** | **String** | The EU PAS (post-authorisation study) register number of the study. | [optional] |
| **study_title** | **String** | The study&#39;s title as registered. | [optional] |
| **catalogue_url** | **String** | The address of this entry on the HMA-EMA catalogue site (the register&#39;s URL LINK). | [optional] |
| **study_status** | **String** | The study&#39;s status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional] |
| **is_ongoing** | **Boolean** | True when the study status is Ongoing. | [optional] |
| **is_finalised** | **Boolean** | True when the study status is Finalised. | [optional] |
| **is_darwin_eu_study** | **Boolean** | Whether the study is flagged a DARWIN EU study. | [optional] |
| **has_encepp_seal** | **Boolean** | Whether the study has been awarded the ENCePP (European Network of Centres for Pharmacoepidemiology and Pharmacovigilance) seal. | [optional] |
| **required_by_regulator** | **String** | The registrant&#39;s answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional] |
| **mandate_stratum** | **String** | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional] |
| **rmp_requirement** | **String** | The risk management plan (RMP) category the registrant says required the study, as typed. | [optional] |
| **has_rmp_category** | **Boolean** | Whether the registrant gave an RMP requirement category. | [optional] |
| **regulatory_procedure_number** | **String** | The regulatory procedure number the registrant gave. | [optional] |
| **funding_source_text** | **String** | The source of funding as the registrant typed it. | [optional] |
| **funding_sources** | **Array&lt;String&gt;** | The funding sources split into an array. | [optional] |
| **is_pharma_funded** | **Boolean** | Whether a funding source is a pharmaceutical company. | [optional] |
| **register_study_type** | **String** | The register&#39;s own Study type value, published as register_study_type because study_type is a shared enumeration of another source&#39;s vocabulary. | [optional] |
| **scope_of_the_study** | **String** | The scope of the study as registered (for example Safety study, Drug utilisation study). | [optional] |
| **study_topic** | **String** | The study topic as registered. | [optional] |
| **study_design** | **String** | The study design as registered. | [optional] |
| **non_interventional_study_design** | **String** | The non-interventional study design as registered. | [optional] |
| **main_study_objective** | **String** | The main study objective as the registrant typed it (free text). | [optional] |
| **data_collection_methods** | **String** | The data collection methods as registered. | [optional] |
| **medicinal_product_name** | **String** | The medicinal product name(s) as registered. | [optional] |
| **study_drug_inn** | **String** | The study drug international non-proprietary name(s) (INN) or common name as typed. | [optional] |
| **atc_code_text** | **String** | The Anatomical Therapeutic Chemical (ATC) code entry as typed, often with the substance name beside the code. | [optional] |
| **drug_mention_count** | **Integer** | How many drug mentions the study carries (rwd_study_drugs has one row per mention). | [optional] |
| **atc_level5_mention_count** | **Integer** | How many of the study&#39;s drug mentions carry an ATC level 5 code. | [optional] |
| **spine_matched_mention_count** | **Integer** | How many of those mentions resolved to a canonical drug. | [optional] |
| **atc_level5_codes** | **Array&lt;String&gt;** | The ATC level 5 (seven-character) codes the study names. | [optional] |
| **atc_level1_codes** | **Array&lt;String&gt;** | The level-1 ATC codes the study names. | [optional] |
| **medical_condition** | **String** | The medicinal condition to be studied, as typed. | [optional] |
| **additional_medical_conditions** | **String** | Additional medical conditions, as typed. | [optional] |
| **condition_count** | **Integer** | How many distinct condition terms the study names. | [optional] |
| **primary_condition_count** | **Integer** | How many of them are primary. | [optional] |
| **mesh_condition_count** | **Integer** | How many of them resolved to a MeSH descriptor. | [optional] |
| **condition_terms** | **Array&lt;String&gt;** | The study&#39;s condition terms, primary and additional, as an array. | [optional] |
| **population_age_groups** | **String** | The population age groups as registered. | [optional] |
| **special_population_of_interest** | **String** | The special population of interest as registered. | [optional] |
| **estimated_number_of_subjects** | **Integer** | The estimated number of subjects as registered; null when blank and also when the register wrote 0 (a zero is an unfilled box: see subjects_reported_zero). | [optional] |
| **subjects_reported_zero** | **Boolean** | True when the register held a 0 for the estimated number of subjects (published as null in estimated_number_of_subjects). | [optional] |
| **subjects_band** | **String** | A size band of estimated_number_of_subjects: under_30, under_100, under_1k, under_100k or over_100k; null when the figure is null. | [optional] |
| **conducting_institution_names** | **Array&lt;String&gt;** | The conducting institutions that match the institution register, as the register spells them; items that did not match are in unmatched_institution_fragments. | [optional] |
| **conducting_institution_count** | **Integer** | How many conducting institutions the study names. | [optional] |
| **additional_institution_names** | **Array&lt;String&gt;** | The additional institutions that match the institution register; unmatched items are in unmatched_institution_fragments. | [optional] |
| **additional_institution_count** | **Integer** | How many additional institutions the study names. | [optional] |
| **unmatched_institution_fragments** | **Array&lt;String&gt;** | Institution text items that did not match the institution register. | [optional] |
| **network_names** | **Array&lt;String&gt;** | The networks conducting the study, as registered. | [optional] |
| **network_count** | **Integer** | How many networks the study names. | [optional] |
| **data_source_names** | **Array&lt;String&gt;** | The data sources the study names, split into an array. | [optional] |
| **data_source_count** | **Integer** | How many data source items the study names. | [optional] |
| **unmatched_data_source_fragments** | **Array&lt;String&gt;** | Data source text items that did not match the data source register (the free-text Other linked data sources box is left whole). | [optional] |
| **conductor_attribution** | **String** | registered when the study names at least one conducting institution that is in the institution register, none otherwise. A study with none is not evidence of no experience: the register gives no further split. | [optional] |
| **study_country_names** | **Array&lt;String&gt;** | The countries of the study, one array element per register country. | [optional] |
| **study_country_count** | **Integer** | How many countries the study names. | [optional] |
| **unmatched_country_fragments** | **Array&lt;String&gt;** | Country text items the country vocabulary of this build does not hold. | [optional] |
| **funding_contract_signed_planned_on** | **Date** | Planned date the funding contract was signed. | [optional] |
| **funding_contract_signed_actual_on** | **Date** | Actual date the funding contract was signed. | [optional] |
| **study_start_planned_on** | **Date** | Planned study start date. | [optional] |
| **study_start_actual_on** | **Date** | Actual study start date. | [optional] |
| **data_analysis_start_planned_on** | **Date** | Planned data analysis start date. | [optional] |
| **data_analysis_start_actual_on** | **Date** | Actual data analysis start date. | [optional] |
| **interim_report_planned_on** | **Date** | Planned date of the interim report, if expected. | [optional] |
| **interim_report_actual_on** | **Date** | Actual date of the interim report. | [optional] |
| **final_report_planned_on** | **Date** | Planned date of the final study report. | [optional] |
| **final_report_actual_on** | **Date** | Actual date of the final study report. | [optional] |
| **has_final_report_pair** | **Boolean** | True when both the planned and the actual final report dates are present. | [optional] |
| **final_report_is_backfilled** | **Boolean** | True when the planned and actual final report dates are both present and equal to the day, which reads as a plan entered after the fact rather than perfect delivery. | [optional] |
| **final_report_slip_days** | **Integer** | Actual minus planned final report date in days (positive is late); null when either date is missing or the two dates are equal (a backfill). | [optional] |
| **study_start_year** | **Integer** | The calendar year of the actual study start. | [optional] |
| **started_in_last_5_years** | **Boolean** | Whether the study started within five years of the snapshot. | [optional] |
| **record_updated_at** | **Time** | When the registrant last edited this catalogue entry. A different clock from last_refreshed_on: it dates the web form, not the database. | [optional] |
| **record_first_published_at** | **Time** | When the catalogue entry was first published. | [optional] |
| **register_source_count** | **Integer** | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional] |
| **study_register_total** | **Integer** | The number of studies in the whole study register at the snapshot. | [optional] |
| **studies_naming_a_registered_source** | **Integer** | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional] |
| **studies_naming_a_registered_conductor** | **Integer** | How many registered studies name a conducting institution that is itself registered. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdStudiesRow.new(
  study_id: null,
  eu_pas_number: null,
  study_title: null,
  catalogue_url: null,
  study_status: null,
  is_ongoing: null,
  is_finalised: null,
  is_darwin_eu_study: null,
  has_encepp_seal: null,
  required_by_regulator: null,
  mandate_stratum: null,
  rmp_requirement: null,
  has_rmp_category: null,
  regulatory_procedure_number: null,
  funding_source_text: null,
  funding_sources: null,
  is_pharma_funded: null,
  register_study_type: null,
  scope_of_the_study: null,
  study_topic: null,
  study_design: null,
  non_interventional_study_design: null,
  main_study_objective: null,
  data_collection_methods: null,
  medicinal_product_name: null,
  study_drug_inn: null,
  atc_code_text: null,
  drug_mention_count: null,
  atc_level5_mention_count: null,
  spine_matched_mention_count: null,
  atc_level5_codes: null,
  atc_level1_codes: null,
  medical_condition: null,
  additional_medical_conditions: null,
  condition_count: null,
  primary_condition_count: null,
  mesh_condition_count: null,
  condition_terms: null,
  population_age_groups: null,
  special_population_of_interest: null,
  estimated_number_of_subjects: null,
  subjects_reported_zero: null,
  subjects_band: null,
  conducting_institution_names: null,
  conducting_institution_count: null,
  additional_institution_names: null,
  additional_institution_count: null,
  unmatched_institution_fragments: null,
  network_names: null,
  network_count: null,
  data_source_names: null,
  data_source_count: null,
  unmatched_data_source_fragments: null,
  conductor_attribution: null,
  study_country_names: null,
  study_country_count: null,
  unmatched_country_fragments: null,
  funding_contract_signed_planned_on: null,
  funding_contract_signed_actual_on: null,
  study_start_planned_on: null,
  study_start_actual_on: null,
  data_analysis_start_planned_on: null,
  data_analysis_start_actual_on: null,
  interim_report_planned_on: null,
  interim_report_actual_on: null,
  final_report_planned_on: null,
  final_report_actual_on: null,
  has_final_report_pair: null,
  final_report_is_backfilled: null,
  final_report_slip_days: null,
  study_start_year: null,
  started_in_last_5_years: null,
  record_updated_at: null,
  record_first_published_at: null,
  register_source_count: null,
  study_register_total: null,
  studies_naming_a_registered_source: null,
  studies_naming_a_registered_conductor: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

