# PharmaObservatory::RareDiseaseSponsorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **sponsor_name** | **String** | The sponsor&#39;s name as filed in the registry. | [optional] |
| **sponsor_org_id** | **String** | Legal-entity level: the organization record the sponsor resolved to. | [optional] |
| **sponsor_company_id** | **String** | Company level: the sponsor&#39;s company (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **registry_sponsor_name** | **String** | The sponsor&#39;s most common name as filed across the disease&#39;s trials. | [optional] |
| **org_type** | **String** | The organization&#39;s type. | [optional] |
| **sponsor_agency_class** | **String** | The funder class the registry declared for the sponsor. | [optional] |
| **sponsor_resolution_confidence** | **String** | How firmly the sponsor resolved to an organization record (high, medium, low or unresolved). | [optional] |
| **n_trials** | **Integer** | Distinct exact direct trials of the disease the sponsor leads or collaborates on. | [optional] |
| **n_trials_lead** | **Integer** | Those trials the sponsor leads. | [optional] |
| **n_trials_collaborator** | **Integer** | Those trials the sponsor collaborates on. | [optional] |
| **n_completed** | **Integer** | Those trials the registry marks completed. | [optional] |
| **n_terminated** | **Integer** | Those trials marked terminated. | [optional] |
| **n_withdrawn** | **Integer** | Those trials marked withdrawn. | [optional] |
| **n_suspended** | **Integer** | Those trials marked suspended. | [optional] |
| **n_open** | **Integer** | Those trials recruiting or not yet recruiting. | [optional] |
| **n_ongoing** | **Integer** | Those trials active but not recruiting, or enrolling by invitation. | [optional] |
| **n_other** | **Integer** | Those trials in any other status. | [optional] |
| **n_unverified_3y** | **Integer** | Those trials whose record was last updated three or more years before the snapshot. | [optional] |
| **first_trial_start** | **Date** | The earliest start date among those trials. | [optional] |
| **last_trial_start** | **Date** | The latest start date among those trials. | [optional] |
| **first_trial_start_year** | **Integer** | The year of the earliest start. | [optional] |
| **last_trial_start_year** | **Integer** | The year of the latest start. | [optional] |
| **last_update_posted_date** | **Date** | The latest registry update among those trials. | [optional] |
| **is_dormant** | **Boolean** | True when the sponsor&#39;s latest trial start is more than five years before the registry snapshot. | [optional] |
| **mapping_arm** | **String** | The mapping arm the dataset is built at (always exact). | [optional] |
| **edge_scope** | **String** | The edge scope the dataset is built at (always direct). | [optional] |
| **freshness_as_of** | **Date** | The registry&#39;s latest update date the snapshot reflects. | [optional] |
| **license_name** | **String** | The licence of the trial registry data. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseSponsorsRow.new(
  orpha_code: null,
  sponsor_name: null,
  sponsor_org_id: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  registry_sponsor_name: null,
  org_type: null,
  sponsor_agency_class: null,
  sponsor_resolution_confidence: null,
  n_trials: null,
  n_trials_lead: null,
  n_trials_collaborator: null,
  n_completed: null,
  n_terminated: null,
  n_withdrawn: null,
  n_suspended: null,
  n_open: null,
  n_ongoing: null,
  n_other: null,
  n_unverified_3y: null,
  first_trial_start: null,
  last_trial_start: null,
  first_trial_start_year: null,
  last_trial_start_year: null,
  last_update_posted_date: null,
  is_dormant: null,
  mapping_arm: null,
  edge_scope: null,
  freshness_as_of: null,
  license_name: null,
  attribution: null
)
```

