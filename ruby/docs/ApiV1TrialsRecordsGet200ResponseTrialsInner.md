# PharmaObservatory::ApiV1TrialsRecordsGet200ResponseTrialsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **canonical_trial_id** | **String** |  |  |
| **nct_id** | **String** |  |  |
| **source_registry** | **String** |  |  |
| **registry** | **String** |  |  |
| **brief_title** | **String** |  |  |
| **official_title** | **String** |  |  |
| **study_type** | **String** |  |  |
| **overall_status** | **String** |  |  |
| **overall_status_summary** | **String** |  |  |
| **phase** | **String** |  |  |
| **phase_by_status** | **String** |  |  |
| **enrollment** | **Integer** |  |  |
| **enrollment_type** | **String** |  |  |
| **enrollment_flag_reason** | **Object** |  |  |
| **start_date** | **String** |  |  |
| **start_year** | **Integer** |  |  |
| **primary_completion_date** | **String** |  |  |
| **primary_completion_date_type** | **String** |  |  |
| **completion_date** | **String** |  |  |
| **completion_date_type** | **String** |  |  |
| **study_first_posted_date** | **String** |  |  |
| **last_update_posted_date** | **String** |  |  |
| **why_stopped** | **String** |  |  |
| **number_of_arms** | **Object** |  |  |
| **number_of_facilities** | **Integer** |  |  |
| **were_results_reported** | **Boolean** |  |  |
| **has_us_facility** | **Boolean** |  |  |
| **gender** | **String** |  |  |
| **minimum_age** | **String** |  |  |
| **maximum_age** | **String** |  |  |
| **healthy_volunteers** | **Boolean** |  |  |
| **criteria** | **String** |  |  |
| **lead_sponsor_name** | **String** |  |  |
| **lead_sponsor_id** | **String** |  |  |
| **collaborators** | **Array&lt;String&gt;** |  |  |
| **participating_sponsors** | **Array&lt;String&gt;** |  |  |
| **conditions** | **Array&lt;String&gt;** |  |  |
| **condition_mesh** | **Array&lt;String&gt;** |  |  |
| **interventions** | **Array&lt;String&gt;** |  |  |
| **intervention_types** | **Array&lt;String&gt;** |  |  |
| **intervention_mesh** | **Array&lt;String&gt;** |  |  |
| **facility_count** | **Integer** |  |  |
| **facility_countries** | **Array&lt;String&gt;** |  |  |
| **facility_country_count** | **Integer** |  |  |
| **enrollment_is_plausible** | **Boolean** |  |  |
| **phase_label** | **String** |  |  |
| **extracted** | [**ApiV1TrialsRecordsGet200ResponseTrialsInnerExtracted**](ApiV1TrialsRecordsGet200ResponseTrialsInnerExtracted.md) |  |  |
| **linked_drugs** | [**Array&lt;ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner&gt;**](ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner.md) |  | [optional] |
| **linked_drugs_built** | **Boolean** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsRecordsGet200ResponseTrialsInner.new(
  canonical_trial_id: null,
  nct_id: null,
  source_registry: null,
  registry: null,
  brief_title: null,
  official_title: null,
  study_type: null,
  overall_status: null,
  overall_status_summary: null,
  phase: null,
  phase_by_status: null,
  enrollment: null,
  enrollment_type: null,
  enrollment_flag_reason: null,
  start_date: null,
  start_year: null,
  primary_completion_date: null,
  primary_completion_date_type: null,
  completion_date: null,
  completion_date_type: null,
  study_first_posted_date: null,
  last_update_posted_date: null,
  why_stopped: null,
  number_of_arms: null,
  number_of_facilities: null,
  were_results_reported: null,
  has_us_facility: null,
  gender: null,
  minimum_age: null,
  maximum_age: null,
  healthy_volunteers: null,
  criteria: null,
  lead_sponsor_name: null,
  lead_sponsor_id: null,
  collaborators: null,
  participating_sponsors: null,
  conditions: null,
  condition_mesh: null,
  interventions: null,
  intervention_types: null,
  intervention_mesh: null,
  facility_count: null,
  facility_countries: null,
  facility_country_count: null,
  enrollment_is_plausible: null,
  phase_label: null,
  extracted: null,
  linked_drugs: null,
  linked_drugs_built: null
)
```

