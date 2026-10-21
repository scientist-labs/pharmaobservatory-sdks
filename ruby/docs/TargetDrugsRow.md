# PharmaObservatory::TargetDrugsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] |
| **target_symbol** | **String** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] |
| **drug_id** | **String** | The drug&#39;s canonical id (the same id the drug datasets use as drug_id). | [optional] |
| **molecule_id** | **String** | The molecule the drug resolves to, the id condition_drugs also calls molecule_id; equal to drug_id for a drug that resolves to no other. | [optional] |
| **drug_name** | **String** | The drug&#39;s name. | [optional] |
| **primary_action_type** | **String** | The action type of the best-evidence mechanism (INHIBITOR, AGONIST and so on); null where the source gives none. | [optional] |
| **evidence_source** | **String** | The best evidence tier that links the drug to the target: curated_moa (an Open Targets curated mechanism edge), chembl_moa (a ChEMBL mechanism) or intervention_seed (a curated investigational asset). | [optional] |
| **evidence_rank** | **Integer** | The rank of evidence_source, 1 for curated_moa, 2 for chembl_moa, 3 for intervention_seed. | [optional] |
| **evidence_sources** | **Array&lt;String&gt;** | Every evidence tier that links the drug to the target; more than one is a corroboration signal. | [optional] |
| **approved_anywhere** | **Boolean** | Whether the drug is approved by any authority we hold. | [optional] |
| **drug_n_trials** | **Float** | Distinct trials of the drug, on any target. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetDrugsRow.new(
  target_id: null,
  target_symbol: null,
  drug_id: null,
  molecule_id: null,
  drug_name: null,
  primary_action_type: null,
  evidence_source: null,
  evidence_rank: null,
  evidence_sources: null,
  approved_anywhere: null,
  drug_n_trials: null
)
```

