# PharmaObservatory::RwdSourceAgeBandsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data_source_id** | **String** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] |
| **band_ordinal** | **Integer** | The position of the age band in the source&#39;s declared list, from 1. | [optional] |
| **age_band** | **String** | The age band label as the holder declared it. Bands can overlap (see is_nested_scheme); never sum them. | [optional] |
| **population_size** | **Integer** | The population size the holder declared for the source (all registered persons). | [optional] |
| **active_population_size** | **Integer** | The active population size the holder declared for the source. | [optional] |
| **has_population_breakdown** | **Boolean** | Whether the holder published a population size breakdown by age band at all. | [optional] |
| **has_active_population_breakdown** | **Boolean** | Whether the holder published an active population breakdown by age band at all. | [optional] |
| **is_nested_scheme** | **Boolean** | True when the source&#39;s bands overlap, so a sum over bands double counts. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdSourceAgeBandsRow.new(
  data_source_id: null,
  band_ordinal: null,
  age_band: null,
  population_size: null,
  active_population_size: null,
  has_population_breakdown: null,
  has_active_population_breakdown: null,
  is_nested_scheme: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

