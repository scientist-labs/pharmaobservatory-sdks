# PharmaObservatory::RwdSourceEtlMappingsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data_source_id** | **String** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] |
| **mapping_ordinal** | **Integer** | The position of the mapping among the source&#39;s declared CDM mappings, from 1. | [optional] |
| **declared_mapping_count** | **Integer** | How many CDM mappings the holder declared, after register join artefacts are removed; 0 when the question was never answered. | [optional] |
| **etl_status** | **String** | The extract-transform-load (ETL) status of the CDM mapping as typed; one comma-joined answer per declared mapping (rwd_source_etl_mappings has one row per mapping). | [optional] |
| **etl_status_aligned** | **Boolean** | False when the register list for this column splits into a different number of pieces than the number of mappings declared; the row value is then null and the whole answer is in etl_status_source. | [optional] |
| **etl_status_source** | **String** | The whole answer the ETL status was split from. | [optional] |
| **etl_cdm_version** | **String** | The CDM version(s) declared for the ETL, comma-joined as typed. | [optional] |
| **etl_cdm_version_aligned** | **Boolean** | False when the register list for this column splits into a different number of pieces than the number of mappings declared; the row value is then null and the whole answer is in etl_cdm_version_source. | [optional] |
| **etl_cdm_version_source** | **String** | The whole answer the CDM version was split from. | [optional] |
| **etl_frequency** | **String** | The ETL frequency declared, comma-joined as typed. | [optional] |
| **etl_frequency_aligned** | **Boolean** | False when the register list for this column splits into a different number of pieces than the number of mappings declared; the row value is then null and the whole answer is in etl_frequency_source. | [optional] |
| **etl_frequency_source** | **String** | The whole answer the ETL frequency was split from. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdSourceEtlMappingsRow.new(
  data_source_id: null,
  mapping_ordinal: null,
  declared_mapping_count: null,
  etl_status: null,
  etl_status_aligned: null,
  etl_status_source: null,
  etl_cdm_version: null,
  etl_cdm_version_aligned: null,
  etl_cdm_version_source: null,
  etl_frequency: null,
  etl_frequency_aligned: null,
  etl_frequency_source: null,
  restriction_tier: null,
  attribution: null
)
```

