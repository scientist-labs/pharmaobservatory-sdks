# PharmaObservatory::TargetAssociationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] |
| **datasource** | **String** | The Open Targets datasource name (for example genetic_association or literature). | [optional] |
| **n_disease_associations** | **Integer** | Distinct disease associations the datasource holds for the target. | [optional] |
| **min_association_score** | **Float** | The lowest association score in the datasource for the target. | [optional] |
| **max_association_score** | **Float** | The highest association score in the datasource for the target. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetAssociationsRow.new(
  target_id: null,
  datasource: null,
  n_disease_associations: null,
  min_association_score: null,
  max_association_score: null
)
```

