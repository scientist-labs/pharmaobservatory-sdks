# PharmaObservatory::ConditionPrevalenceRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **disease_prevalence_id** | **String** | The figure&#39;s key. | [optional] |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The condition&#39;s MeSH heading. | [optional] |
| **condition_mapping_relation** | **String** | How the figure&#39;s disease maps to the condition: exact, narrower, broader, undetermined or unspecified. | [optional] |
| **disease_id** | **String** | The rare disease key of the figure&#39;s disease (orpha:&lt;code&gt;). | [optional] |
| **disease_name** | **String** | The disease&#39;s name in the source. | [optional] |
| **orpha_code** | **String** | The Orphanet disease number. | [optional] |
| **prevalence_type** | **String** | What the figure measures: point_prevalence, lifetime_prevalence, annual_incidence, prevalence_at_birth or cases_families. | [optional] |
| **prevalence_class** | **String** | The source&#39;s prevalence band, as filed. | [optional] |
| **prevalence_lower_bound** | **Float** | Lower bound of the prevalence band as a ratio of the population (0.0001 is 1 in 10,000). | [optional] |
| **prevalence_upper_bound** | **Float** | Upper bound of the prevalence band as a ratio of the population (0.0005 is 5 in 10,000). | [optional] |
| **population_denominator** | **Float** | The population size the source states the band against (for example 10000 for a band per 10,000 persons). | [optional] |
| **population_unit** | **String** | The unit of the population denominator. | [optional] |
| **qualification** | **String** | The source&#39;s qualification of the figure (class_only, value_and_class and so on). | [optional] |
| **validation_status** | **String** | The source&#39;s validation status for the figure (validated or not_yet_validated). | [optional] |
| **geography** | **String** | The geography the figure describes, as the source states it. | [optional] |
| **jurisdiction_code** | **String** | US, EU (European Union) or GB (Great Britain) where the geography is that jurisdiction; null otherwise. | [optional] |
| **observation_date** | **Date** | The date the source assigns to the observation. | [optional] |
| **source_snapshot_date** | **Date** | The date of the source release the figure was read from. | [optional] |
| **source** | **String** | The publication reference the source cites for the figure. | [optional] |
| **source_url** | **String** | The source release URL. | [optional] |
| **license_name** | **String** | The figure&#39;s licence. | [optional] |
| **attribution** | **String** | The credit line the licence requires. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionPrevalenceRow.new(
  disease_prevalence_id: null,
  condition_id: null,
  condition_name: null,
  condition_mapping_relation: null,
  disease_id: null,
  disease_name: null,
  orpha_code: null,
  prevalence_type: null,
  prevalence_class: null,
  prevalence_lower_bound: null,
  prevalence_upper_bound: null,
  population_denominator: null,
  population_unit: null,
  qualification: null,
  validation_status: null,
  geography: null,
  jurisdiction_code: null,
  observation_date: null,
  source_snapshot_date: null,
  source: null,
  source_url: null,
  license_name: null,
  attribution: null
)
```

