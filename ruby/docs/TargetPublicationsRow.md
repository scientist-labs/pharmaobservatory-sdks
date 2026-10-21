# PharmaObservatory::TargetPublicationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] |
| **target_symbol** | **String** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] |
| **pmid** | **String** | The PubMed id. | [optional] |
| **citation** | **String** | The citation string registrants gave. | [optional] |
| **title** | **String** | The paper&#39;s title from PubMed; null when PubMed has no record. | [optional] |
| **pub_year** | **Integer** | The publication year. | [optional] |
| **n_linked_trials** | **Integer** | Distinct trials in the registry that cite the paper, for any target. | [optional] |
| **n_target_trials** | **Integer** | Distinct trials of this target that cite the paper. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetPublicationsRow.new(
  target_id: null,
  target_symbol: null,
  pmid: null,
  citation: null,
  title: null,
  pub_year: null,
  n_linked_trials: null,
  n_target_trials: null
)
```

