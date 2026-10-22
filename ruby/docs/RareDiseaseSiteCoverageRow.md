# PharmaObservatory::RareDiseaseSiteCoverageRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **concept_type** | **String** | Whether the concept is a group of disorders, a disorder or a subtype. | [optional] |
| **snapshot_date** | **Date** | The date of the registry snapshot. | [optional] |
| **exact_state** | **String** | How far the exact arm reaches: no_mapping, no_trials, no_site_data or computed. | [optional] |
| **exact_n_conditions** | **Integer** | MeSH conditions the disease maps to under the exact arm. | [optional] |
| **exact_n_trials** | **Integer** | Trials linked under the exact arm. | [optional] |
| **exact_n_trials_with_site** | **Integer** | Trials of the exact arm with at least one geocoded site. | [optional] |
| **exact_n_sites_total** | **Integer** | Distinct sites of the exact arm. | [optional] |
| **exact_n_sites_emitted** | **Integer** | Sites of the exact arm published in rare_disease_trial_sites (capped at 500 per disease and arm). | [optional] |
| **exact_n_sites_emitted_geocoded** | **Integer** | Published sites of the exact arm that carry a coordinate. | [optional] |
| **exact_n_facilities_ungeocoded** | **Integer** | Facility listings of the exact arm with no coordinate. | [optional] |
| **unspecified_state** | **String** | How far the unspecified arm reaches: no_mapping, no_trials, no_site_data or computed. | [optional] |
| **unspecified_n_conditions** | **Integer** | MeSH conditions the disease maps to under the unspecified arm. | [optional] |
| **unspecified_n_trials** | **Integer** | Trials linked under the unspecified arm. | [optional] |
| **unspecified_n_trials_with_site** | **Integer** | Trials of the unspecified arm with at least one geocoded site. | [optional] |
| **unspecified_n_sites_total** | **Integer** | Distinct sites of the unspecified arm. | [optional] |
| **unspecified_n_sites_emitted** | **Integer** | Sites of the unspecified arm published in rare_disease_trial_sites (capped at 500 per disease and arm). | [optional] |
| **unspecified_n_sites_emitted_geocoded** | **Integer** | Published sites of the unspecified arm that carry a coordinate. | [optional] |
| **unspecified_n_facilities_ungeocoded** | **Integer** | Facility listings of the unspecified arm with no coordinate. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseSiteCoverageRow.new(
  orpha_code: null,
  concept_type: null,
  snapshot_date: null,
  exact_state: null,
  exact_n_conditions: null,
  exact_n_trials: null,
  exact_n_trials_with_site: null,
  exact_n_sites_total: null,
  exact_n_sites_emitted: null,
  exact_n_sites_emitted_geocoded: null,
  exact_n_facilities_ungeocoded: null,
  unspecified_state: null,
  unspecified_n_conditions: null,
  unspecified_n_trials: null,
  unspecified_n_trials_with_site: null,
  unspecified_n_sites_total: null,
  unspecified_n_sites_emitted: null,
  unspecified_n_sites_emitted_geocoded: null,
  unspecified_n_facilities_ungeocoded: null
)
```

