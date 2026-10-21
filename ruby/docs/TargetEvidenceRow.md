# PharmaObservatory::TargetEvidenceRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_evidence_id** | **String** | The row id: a hash of the target, drug, indication, trial and sponsor organization. | [optional] |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] |
| **target_symbol** | **String** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] |
| **drug_id** | **String** | The drug&#39;s canonical id (the same id the drug datasets use as drug_id). | [optional] |
| **drug_name** | **String** | The drug&#39;s name. | [optional] |
| **approved_anywhere** | **Boolean** | Whether the drug is approved by any authority we hold. | [optional] |
| **primary_action_type** | **String** | The action type of the drug&#39;s best-evidence mechanism on the target; null where the source gives none. | [optional] |
| **evidence_source** | **String** | The best evidence tier that links the drug to the target: curated_moa (an Open Targets curated mechanism edge), chembl_moa (a ChEMBL mechanism) or intervention_seed (a curated investigational asset). | [optional] |
| **condition_id** | **String** | The condition as a MeSH descriptor id (D followed by digits), the id the conditions datasets use. | [optional] |
| **condition_name** | **String** | The condition&#39;s name. | [optional] |
| **trial_id** | **String** | The ClinicalTrials.gov NCT id, the id the trials datasets use. | [optional] |
| **max_phase** | **Integer** | The furthest phase the program (drug and indication) reached, 0 to 4 (0 is phase not reported). | [optional] |
| **sponsor_org_id** | **String** | Legal-entity level: the trial&#39;s lead sponsor organization as filed; null when it does not resolve. | [optional] |
| **sponsor_name** | **String** | The lead sponsor&#39;s name as filed. | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the lead sponsor organization belongs to; null when none. | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetEvidenceRow.new(
  target_evidence_id: null,
  target_id: null,
  target_symbol: null,
  drug_id: null,
  drug_name: null,
  approved_anywhere: null,
  primary_action_type: null,
  evidence_source: null,
  condition_id: null,
  condition_name: null,
  trial_id: null,
  max_phase: null,
  sponsor_org_id: null,
  sponsor_name: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null
)
```

