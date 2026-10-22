# PharmaObservatory::RareDiseaseTrialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **mapping_arm** | **String** | exact when the link rests on an asserted equivalence, unspecified when it rests on a MONDO cross-reference that states no relation. | [optional] |
| **edge_scope** | **String** | direct when the trial&#39;s own condition list names the mapped condition, broader_family when only a broader condition does. | [optional] |
| **via_exact** | **Boolean** | True when an exact mapping links the trial to the disease. | [optional] |
| **via_unspecified** | **Boolean** | True when an unspecified-relation mapping links the trial to the disease. | [optional] |
| **via_direct** | **Boolean** | True when a direct edge links the trial to the disease. | [optional] |
| **via_broader** | **Boolean** | True when a broader-family edge links the trial to the disease. | [optional] |
| **mapping_target_curie** | **String** | The MeSH condition (as a prefixed identifier) the mapping reaches the trial through. | [optional] |
| **mapping_sources** | **Array&lt;String&gt;** | The files that assert the mapping (ordo, orphadata_product1 or mondo_legacy_bridge). | [optional] |
| **trial_title** | **String** | The trial&#39;s brief title. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall recruitment status (COMPLETED, RECRUITING, TERMINATED and so on). | [optional] |
| **study_type** | **String** | The registry&#39;s study type (INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS). | [optional] |
| **phase_label** | **String** | The trial phase as a display label (Phase 2, Phase 3, Observational and so on). | [optional] |
| **enrollment** | **Integer** | The trial&#39;s total enrollment as registered. | [optional] |
| **enrollment_type** | **String** | ACTUAL when the registered enrollment is the number enrolled, ESTIMATED when it is the planned number. | [optional] |
| **start_date** | **Date** | The trial&#39;s start date as registered. | [optional] |
| **start_year** | **Integer** | The year of start_date. | [optional] |
| **completion_date** | **Date** | The trial&#39;s completion date as registered. | [optional] |
| **why_stopped** | **String** | The reason the sponsor gave for stopping the trial, where it stopped early. | [optional] |
| **were_results_reported** | **Boolean** | True when the sponsor has posted results to the registry. | [optional] |
| **last_update_posted_date** | **Date** | The date the registry record was last updated. | [optional] |
| **facility_count** | **Integer** | Facilities the registry lists for the trial. | [optional] |
| **facility_country_count** | **Integer** | Countries the listed facilities are in. | [optional] |
| **facility_countries** | **Array&lt;String&gt;** | The countries the listed facilities are in. | [optional] |
| **days_since_update** | **Integer** | Days from the last registry update to the registry snapshot date. | [optional] |
| **is_unverified** | **Boolean** | True when the registry record was last updated two or more years before the snapshot date. | [optional] |
| **sponsor_name** | **String** | The sponsor&#39;s name as filed in the registry. | [optional] |
| **sponsor_org_id** | **String** | Legal-entity level: the organization record the sponsor resolved to (null when unresolved). | [optional] |
| **sponsor_company_id** | **String** | Company level: the sponsor&#39;s company (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **sponsor_agency_class** | **String** | The funder class the registry declared for the lead sponsor; null when the trial has no lead sponsor row. | [optional] |
| **sponsor_resolution_confidence** | **String** | How firmly the lead sponsor resolved to an organization record (high, medium or unresolved). | [optional] |
| **license_name** | **String** | The licence of the trial registry data. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseTrialsRow.new(
  orpha_code: null,
  trial_id: null,
  mapping_arm: null,
  edge_scope: null,
  via_exact: null,
  via_unspecified: null,
  via_direct: null,
  via_broader: null,
  mapping_target_curie: null,
  mapping_sources: null,
  trial_title: null,
  overall_status: null,
  study_type: null,
  phase_label: null,
  enrollment: null,
  enrollment_type: null,
  start_date: null,
  start_year: null,
  completion_date: null,
  why_stopped: null,
  were_results_reported: null,
  last_update_posted_date: null,
  facility_count: null,
  facility_country_count: null,
  facility_countries: null,
  days_since_update: null,
  is_unverified: null,
  sponsor_name: null,
  sponsor_org_id: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  sponsor_agency_class: null,
  sponsor_resolution_confidence: null,
  license_name: null,
  attribution: null
)
```

