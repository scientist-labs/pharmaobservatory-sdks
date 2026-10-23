# PharmaObservatory::BiomarkerTraversalStatusRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The biomarker&#39;s route key, the entry key for the biomarkers object type. | [optional] |
| **biomarker_id** | **String** | The biomarker&#39;s own key. | [optional] |
| **biomarker_kind** | **String** | What kind of biomarker the row describes. | [optional] |
| **traversal_key** | **String** | The traversal. | [optional] |
| **display_order** | **Integer** | The order the web page lists the traversals in, one to nine. | [optional] |
| **traversal_name** | **String** | The traversal&#39;s display name. | [optional] |
| **applicability_state** | **String** | applicable: measured for this kind of biomarker. inapplicable: does not apply to this kind. not_measured: applies in principle but the corpus is not mined for this kind yet. | [optional] |
| **context_mode** | **String** | How the traversal reads for this biomarker. direct: its own rows. parent_gene: the variant&#39;s gene rows. parent_and_exact: the gene&#39;s rows plus the exact variant&#39;s. child_variants: the gene&#39;s variant children. source_linked: shown when a source assertion resolves to it. bridge_deferred: depends on a bridge that has not landed. none: no traversal. | [optional] |
| **reason_code** | **String** | The code behind the applicability state and context mode; state_detail is its sentence. | [optional] |
| **state_detail** | **String** | The sentence the web page prints for the reason code. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerTraversalStatusRow.new(
  route_key: null,
  biomarker_id: null,
  biomarker_kind: null,
  traversal_key: null,
  display_order: null,
  traversal_name: null,
  applicability_state: null,
  context_mode: null,
  reason_code: null,
  state_detail: null
)
```

