# PharmaObservatory::CellLinesList200ResponseRowsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **scientist_model_id** | **String** |  |  |
| **model_lot_id** | **String** |  |  |
| **cell_type** | **String** |  |  |
| **supplier** | **String** |  |  |
| **sample_species** | **String** |  |  |
| **safety_status** | **String** |  |  |
| **safety_label** | **String** |  |  |
| **confirmed_problematic** | **Boolean** |  |  |
| **cellosaurus_accession** | **String** |  |  |
| **cellosaurus_url** | **String** |  |  |
| **problematic_comment** | **String** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CellLinesList200ResponseRowsInner.new(
  scientist_model_id: null,
  model_lot_id: null,
  cell_type: null,
  supplier: null,
  sample_species: null,
  safety_status: null,
  safety_label: null,
  confirmed_problematic: null,
  cellosaurus_accession: null,
  cellosaurus_url: null,
  problematic_comment: null
)
```

