# PharmaObservatory::CompanyPeerProfileRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **peer_state** | **String** | measured when the company has at least one peer; no_live_owned_program when it has no live owned program; phase_not_reported when its live owned programs carry no phase; no_candidates when stage and modality are known and no other company shares them. | [optional] |
| **n_live_owned_programs** | **Integer** | Distinct live owned (drug, condition) programs of the company&#39;s subtree. | [optional] |
| **n_live_owned_drugs** | **Integer** | Distinct live owned drugs of the company&#39;s subtree. | [optional] |
| **lead_stage** | **String** | Furthest phase of the live owned programs: approved, phase3, phase2 (Phase 2/3 included), phase1 (Phase 1/2 included) or early_phase1. Null when no program carries a phase. | [optional] |
| **lead_max_phase** | **Integer** | The phase code behind lead_stage: 6 Phase 4 or approved, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1. | [optional] |
| **lead_drug_id** | **String** | The canonical drug id of the lead drug, the highest-phase live owned drug (ties: most trials, then drug id). | [optional] |
| **lead_drug_name** | **String** | Name of the lead drug. | [optional] |
| **n_biologic_drugs** | **Integer** | Live owned drugs whose GSRS substance type is protein or nucleic acid. | [optional] |
| **biologic_share** | **Float** | n_biologic_drugs over n_live_owned_drugs, 0 to 1. Null when the company has no live owned drug. | [optional] |
| **modality_band** | **String** | biologic when biologic_share is at least 0.5, else non_biologic. | [optional] |
| **primary_area** | **String** | MeSH top-level category with the most live owned program trials. | [optional] |
| **primary_area_trials** | **Integer** | Live owned program trials in primary_area. | [optional] |
| **num_lead_trials** | **Integer** | Distinct lead-sponsored trials of the company&#39;s subtree. | [optional] |
| **size_band** | **String** | Band of num_lead_trials: 1-4, 5-19, 20-99 or 100+. | [optional] |
| **n_cell_candidates** | **Integer** | Other measured companies outside the company&#39;s group that share stage, modality, area and size (match level 1), before the top 10 are kept. | [optional] |
| **n_peers** | **Integer** | Peer rows kept for the company, at most 10. | [optional] |
| **coverage_measured** | **Integer** | Companies in this dataset with peer_state measured. The same value on every row. | [optional] |
| **coverage_total** | **Integer** | Companies in this dataset. The same value on every row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyPeerProfileRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  peer_state: null,
  n_live_owned_programs: null,
  n_live_owned_drugs: null,
  lead_stage: null,
  lead_max_phase: null,
  lead_drug_id: null,
  lead_drug_name: null,
  n_biologic_drugs: null,
  biologic_share: null,
  modality_band: null,
  primary_area: null,
  primary_area_trials: null,
  num_lead_trials: null,
  size_band: null,
  n_cell_candidates: null,
  n_peers: null,
  coverage_measured: null,
  coverage_total: null
)
```

