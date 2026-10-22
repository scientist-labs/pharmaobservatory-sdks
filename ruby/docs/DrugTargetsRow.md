# PharmaObservatory::DrugTargetsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **ensembl_id** | **String** | The target&#39;s Ensembl gene id: the entry key of the targets object type. | [optional] |
| **symbol** | **String** | The gene symbol: the route key of the biomarker page for the same gene. | [optional] |
| **target_name** | **String** | The target&#39;s name. | [optional] |
| **uniprot** | **String** | The UniProt accession of the target protein. | [optional] |
| **has_pdb** | **Boolean** | True when the Protein Data Bank holds a structure for the protein. | [optional] |
| **primary_action_type** | **String** | The mechanism of action (for example INHIBITOR, AGONIST) from the best-ranked evidence that states one. | [optional] |
| **evidence_source** | **String** | The best-ranked evidence source for the link. | [optional] |
| **evidence_rank** | **Integer** | The rank of that source; lower is stronger. | [optional] |
| **evidence_sources** | **Array&lt;String&gt;** | Every evidence source for the link. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugTargetsRow.new(
  molecule_id: null,
  ensembl_id: null,
  symbol: null,
  target_name: null,
  uniprot: null,
  has_pdb: null,
  primary_action_type: null,
  evidence_source: null,
  evidence_rank: null,
  evidence_sources: null
)
```

