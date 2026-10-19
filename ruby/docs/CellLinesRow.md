# PharmaObservatory::CellLinesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cell_line_id** | **String** | The catalog listing&#39;s model id (scientist_model_id). | [optional] |
| **alternate_cell_line_ids** | **String** | Other catalog model ids for the same line. | [optional] |
| **model_lot_id** | **String** | The lot id. | [optional] |
| **cell_type** | **String** | The cell line&#39;s name. | [optional] |
| **provider_id** | **String** | The supplier&#39;s catalog id. | [optional] |
| **supplier** | **String** | The supplier as listed. | [optional] |
| **sample_species** | **String** | The species. | [optional] |
| **storefront_enable** | **Boolean** | Whether the listing is enabled in the storefront. | [optional] |
| **mutation_data** | **String** | Mutation data as listed. | [optional] |
| **status** | **String** | The listing status. | [optional] |
| **safety_status** | **String** | confirmed_problematic, review_candidate, no_problematic_match or missing_cell_type. | [optional] |
| **resolution_status** | **String** | resolved, unresolved_review, no_match or missing_input. | [optional] |
| **confirmed_problematic** | **Boolean** | Whether the line is confirmed problematic in Cellosaurus. | [optional] |
| **cellosaurus_accession** | **String** | The displayed Cellosaurus accession. | [optional] |
| **cellosaurus_url** | **String** | Link to the Cellosaurus entry. | [optional] |
| **problem_category** | **String** | The problem category, for a problematic line. | [optional] |
| **problematic_comment** | **String** | The Cellosaurus comment, for a problematic line. | [optional] |
| **iclac_registration_ids** | **Array&lt;String&gt;** | International Cell Line Authentication Committee (ICLAC) registration ids. | [optional] |
| **release_version** | **String** | The Cellosaurus release the check ran against. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CellLinesRow.new(
  cell_line_id: null,
  alternate_cell_line_ids: null,
  model_lot_id: null,
  cell_type: null,
  provider_id: null,
  supplier: null,
  sample_species: null,
  storefront_enable: null,
  mutation_data: null,
  status: null,
  safety_status: null,
  resolution_status: null,
  confirmed_problematic: null,
  cellosaurus_accession: null,
  cellosaurus_url: null,
  problem_category: null,
  problematic_comment: null,
  iclac_registration_ids: null,
  release_version: null
)
```

