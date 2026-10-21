# PharmaObservatory::TargetPatentMentionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] |
| **target_symbol** | **String** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] |
| **drug_id** | **String** | The drug&#39;s canonical id (the same id the drug datasets use as drug_id). | [optional] |
| **drug_name** | **String** | The drug&#39;s name. | [optional] |
| **patent_number** | **String** | The patent number. | [optional] |
| **country** | **String** | The patent office&#39;s country code. | [optional] |
| **publication_date** | **Date** | The patent&#39;s publication date. | [optional] |
| **title** | **String** | The patent&#39;s title. | [optional] |
| **assignee** | **String** | The patent&#39;s assignee as filed; null where none. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetPatentMentionsRow.new(
  target_id: null,
  target_symbol: null,
  drug_id: null,
  drug_name: null,
  patent_number: null,
  country: null,
  publication_date: null,
  title: null,
  assignee: null
)
```

