# PharmaObservatory::RwdWithheldFieldsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **register_key** | **String** | Which of the five exports the row concerns: study, data_source_human, data_source_veterinary, institution or network. The entry key of the rwd_registers object type. | [optional] |
| **register_label** | **String** | The register&#39;s plain-words label. | [optional] |
| **source_ordinal** | **Integer** | The withheld column&#39;s position in the export. | [optional] |
| **source_header** | **String** | The export&#39;s header for the withheld column. | [optional] |
| **column_name** | **String** | The name of the withheld column in the lane&#39;s landed schema. | [optional] |
| **withheld_reason** | **String** | Why the column is not published. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdWithheldFieldsRow.new(
  register_key: null,
  register_label: null,
  source_ordinal: null,
  source_header: null,
  column_name: null,
  withheld_reason: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

