# PharmaObservatory::TargetRegulatoryPrecedentRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] |
| **target_symbol** | **String** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] |
| **drug_id** | **String** | The drug&#39;s canonical id (the same id the drug datasets use as drug_id). | [optional] |
| **drug_name** | **String** | The drug&#39;s name. | [optional] |
| **authority** | **String** | The approving authority. | [optional] |
| **jurisdiction** | **String** | The authority&#39;s jurisdiction. | [optional] |
| **status** | **String** | The approval status the authority records. | [optional] |
| **first_approval_date** | **Date** | The authority&#39;s first approval date of the drug. | [optional] |
| **approval_id** | **String** | The authority&#39;s own approval identifier. | [optional] |
| **primary_action_type** | **String** | The action type of the drug&#39;s curated mechanism on the target. | [optional] |
| **moa_evidence_source** | **String** | The evidence tier of the mechanism; curated_moa on every row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetRegulatoryPrecedentRow.new(
  target_id: null,
  target_symbol: null,
  drug_id: null,
  drug_name: null,
  authority: null,
  jurisdiction: null,
  status: null,
  first_approval_date: null,
  approval_id: null,
  primary_action_type: null,
  moa_evidence_source: null
)
```

