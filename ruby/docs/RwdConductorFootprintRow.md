# PharmaObservatory::RwdConductorFootprintRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **conductor_name_key** | **String** | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. | [optional] |
| **conductor_id** | **String** | The conductor slug: the register&#39;s normalised institution name key with spaces as hyphens, the id in the web page address (/rwd-conductors/&lt;id&gt;). The entry key of the rwd_conductors object type; injective because no name key carries a hyphen. | [optional] |
| **conductor_name** | **String** | The conductor&#39;s display name: one register spelling of the key. | [optional] |
| **dimension** | **String** | What the footprint value is: delivered_country (a country of a conducting study) or delivered_data_source (a registered data source a conducting study names). | [optional] |
| **value** | **String** | The country or data source name the row counts. | [optional] |
| **study_count** | **Integer** | Distinct conducting studies of the conductor for this country or data source. | [optional] |
| **conductor_study_count** | **Integer** | Distinct conducting studies of the conductor, the denominator beside study_count. | [optional] |
| **data_source_ids** | **Array&lt;String&gt;** | The register Data source IDs of the registered sources with this name (delivered_data_source rows only; null for a country). | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdConductorFootprintRow.new(
  conductor_name_key: null,
  conductor_id: null,
  conductor_name: null,
  dimension: null,
  value: null,
  study_count: null,
  conductor_study_count: null,
  data_source_ids: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

