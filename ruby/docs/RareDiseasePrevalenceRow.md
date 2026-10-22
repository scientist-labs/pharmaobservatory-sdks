# PharmaObservatory::RareDiseasePrevalenceRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **prevalence_assertion_id** | **String** | The identifier of the prevalence observation. | [optional] |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **prevalence_type** | **String** | The kind of observation: point prevalence, prevalence at birth, annual incidence, lifetime prevalence or cases and families. | [optional] |
| **prevalence_class** | **String** | The Orphanet prevalence class of the observation, for example 1-9 / 100 000; null when the source gives none. | [optional] |
| **prevalence_value_source** | **String** | The prevalence value as the source writes it. | [optional] |
| **normalized_prevalence_ratio** | **Float** | The observation as a ratio (cases per person) where one can be derived. | [optional] |
| **prevalence_lower_bound** | **Float** | The lower bound of the prevalence class as a ratio. | [optional] |
| **prevalence_upper_bound** | **Float** | The upper bound of the prevalence class as a ratio. | [optional] |
| **population_denominator** | **Float** | The denominator of the figure (for example 100000). | [optional] |
| **population_unit** | **String** | The unit of the denominator. | [optional] |
| **population_scope** | **String** | The population the observation covers. | [optional] |
| **age_scope** | **String** | The age group the observation covers. | [optional] |
| **qualification** | **String** | How the figure is qualified at the source (value and class, class only, cases, families). | [optional] |
| **validation_status** | **String** | validated when Orphanet marks the observation validated, not_yet_validated otherwise. | [optional] |
| **geography** | **String** | The geography the observation covers (Worldwide, a country or a region). | [optional] |
| **observation_date** | **Date** | The date of the observation where the source gives one. | [optional] |
| **uncertainty_lower** | **Float** | The lower end of the stated uncertainty range. | [optional] |
| **uncertainty_upper** | **Float** | The upper end of the stated uncertainty range. | [optional] |
| **source** | **String** | The evidence source Orphanet cites for the observation. | [optional] |
| **source_snapshot_date** | **Date** | The date of the Orphadata Science snapshot the row was read from. | [optional] |
| **source_url** | **String** | The URL of the source file the row was read from. | [optional] |
| **license_name** | **String** | The licence of the Orphadata Science file. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseasePrevalenceRow.new(
  prevalence_assertion_id: null,
  orpha_code: null,
  prevalence_type: null,
  prevalence_class: null,
  prevalence_value_source: null,
  normalized_prevalence_ratio: null,
  prevalence_lower_bound: null,
  prevalence_upper_bound: null,
  population_denominator: null,
  population_unit: null,
  population_scope: null,
  age_scope: null,
  qualification: null,
  validation_status: null,
  geography: null,
  observation_date: null,
  uncertainty_lower: null,
  uncertainty_upper: null,
  source: null,
  source_snapshot_date: null,
  source_url: null,
  license_name: null,
  attribution: null
)
```

