# PharmaObservatory::RwdCoverageRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **chain_key** | **String** | Which denominator chain the row belongs to. | [optional] |
| **chain_label** | **String** | The chain&#39;s plain-words label. | [optional] |
| **chain_is_a_partition** | **Boolean** | True when the chain&#39;s steps partition the denominator (they sum to it); false when steps overlap. | [optional] |
| **step_ordinal** | **Integer** | The step&#39;s position in the chain; 0 is the denominator itself, published as a step so the chain reads top to bottom. | [optional] |
| **step_key** | **String** | The step&#39;s key. | [optional] |
| **step_label** | **String** | The step&#39;s plain-words label. | [optional] |
| **register_key** | **String** | Which of the five exports the row concerns: study, data_source_human, data_source_veterinary, institution or network. The entry key of the rwd_registers object type. | [optional] |
| **measured_value** | **Integer** | The measured count at this step. | [optional] |
| **denominator_key** | **String** | Which denominator the step is read against. | [optional] |
| **denominator_label** | **String** | The denominator&#39;s plain-words label. | [optional] |
| **denominator_value** | **Integer** | The denominator&#39;s value. | [optional] |
| **share_of_denominator** | **Float** | measured_value over denominator_value, as a fraction. | [optional] |
| **derivation_note** | **String** | How the step is derived. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **snapshot_date** | **Date** | The date of the catalogue snapshot. | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdCoverageRow.new(
  chain_key: null,
  chain_label: null,
  chain_is_a_partition: null,
  step_ordinal: null,
  step_key: null,
  step_label: null,
  register_key: null,
  measured_value: null,
  denominator_key: null,
  denominator_label: null,
  denominator_value: null,
  share_of_denominator: null,
  derivation_note: null,
  snapshot_label: null,
  snapshot_date: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

