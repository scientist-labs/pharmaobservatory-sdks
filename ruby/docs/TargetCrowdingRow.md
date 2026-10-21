# PharmaObservatory::TargetCrowdingRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] |
| **target_symbol** | **String** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] |
| **max_phase** | **Integer** | The band: the furthest phase a program reached, 0 to 4 (0 is phase not reported). | [optional] |
| **n_programs** | **Integer** | Distinct drug and indication programs whose furthest phase is the band. | [optional] |
| **n_active_programs** | **Integer** | Programs in the band with no terminated, withdrawn or suspended trial. | [optional] |
| **n_discontinued** | **Integer** | Programs in the band with at least one terminated, withdrawn or suspended trial. | [optional] |
| **n_drugs** | **Integer** | Distinct drugs among the band&#39;s programs. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetCrowdingRow.new(
  target_id: null,
  target_symbol: null,
  max_phase: null,
  n_programs: null,
  n_active_programs: null,
  n_discontinued: null,
  n_drugs: null
)
```

