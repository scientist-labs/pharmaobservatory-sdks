# PharmaObservatory::RareDiseaseInvestigatorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **disease_name** | **String** | The disease&#39;s preferred English name. | [optional] |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key. | [optional] |
| **investigator_person_name** | **String** | The investigator&#39;s name as the registry filed it. | [optional] |
| **identity_basis** | **String** | What the person&#39;s identity rests on: curated, npi, orcid, reporter, affinity or self. | [optional] |
| **merged_fragment_count** | **Integer** | Name fragments merged into this person. | [optional] |
| **name_is_ambiguous_in_disease** | **Boolean** | True when another investigator of the same disease shares the name. | [optional] |
| **pi_trial_count** | **Integer** | Exact direct trials of the disease that name the person as principal investigator. | [optional] |
| **first_trial_start_date** | **Date** | The earliest start date among those trials. | [optional] |
| **last_trial_start_date** | **Date** | The latest start date among those trials. | [optional] |
| **last_trial_start_year** | **Integer** | The year of the latest start. | [optional] |
| **evidence_trial_id** | **String** | One of the trials that names the person (the latest started). | [optional] |
| **last_record_update_date** | **Date** | The latest registry update among those trials. | [optional] |
| **registry_reference_date** | **Date** | The registry&#39;s latest update date the snapshot reflects. | [optional] |
| **recency_state** | **String** | How recent the latest trial start is against the registry snapshot: within_5_years, 5_to_10_years, over_10_years or no_start_date. | [optional] |
| **trial_site_names** | **Array&lt;String&gt;** | The sites of the person&#39;s trials as the registry lists them. | [optional] |
| **trial_site_count** | **Integer** | How many such sites. | [optional] |
| **stated_affiliations** | **Array&lt;String&gt;** | The affiliations the registry states for the person. | [optional] |
| **stated_affiliation_count** | **Integer** | How many stated affiliations. | [optional] |
| **affiliation_evidence** | **String** | trial_site when a trial site is listed, stated_affiliation_only when only an affiliation is stated, none otherwise. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseInvestigatorsRow.new(
  orpha_code: null,
  disease_name: null,
  investigator_person_id: null,
  investigator_person_name: null,
  identity_basis: null,
  merged_fragment_count: null,
  name_is_ambiguous_in_disease: null,
  pi_trial_count: null,
  first_trial_start_date: null,
  last_trial_start_date: null,
  last_trial_start_year: null,
  evidence_trial_id: null,
  last_record_update_date: null,
  registry_reference_date: null,
  recency_state: null,
  trial_site_names: null,
  trial_site_count: null,
  stated_affiliations: null,
  stated_affiliation_count: null,
  affiliation_evidence: null
)
```

