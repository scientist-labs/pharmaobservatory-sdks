# PharmaObservatory::CompanyPeersRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **peer_company_id** | **String** | Company level: the peer company. | [optional] |
| **peer_company_name** | **String** | Display name of the peer company. | [optional] |
| **peer_group_id** | **String** | Group level: the peer&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **peer_group_name** | **String** | Display name of the peer&#39;s group. | [optional] |
| **peer_rank** | **Integer** | The peer&#39;s rank for the company, 1 is the closest match: match_level first, then the closeness of lead-trial counts, then of biologic share. Dense 1 to N. | [optional] |
| **match_level** | **Integer** | What the peer matched the company on. 1 stage, modality, area and size; 2 stage, modality and area; 3 stage, modality and size; 4 stage and modality. Stage and modality always match. | [optional] |
| **matched_area** | **Boolean** | True when the peer shares the company&#39;s primary therapeutic area. | [optional] |
| **matched_size** | **Boolean** | True when the peer shares the company&#39;s size band. | [optional] |
| **lead_stage** | **String** | The peer&#39;s lead stage, the furthest phase of its live owned programs. Equal to the company&#39;s lead stage on every row. | [optional] |
| **lead_drug_name** | **String** | The peer&#39;s lead drug, the highest-phase live owned drug (ties: most trials). | [optional] |
| **biologic_share** | **Float** | Share, 0 to 1, of the peer&#39;s live owned drugs that are biologic (GSRS protein or nucleic acid substance type). | [optional] |
| **modality_band** | **String** | biologic when biologic_share is at least 0.5, else non_biologic. Equal to the company&#39;s band on every row. | [optional] |
| **primary_area** | **String** | The peer&#39;s primary therapeutic area: the MeSH top-level category with the most live owned program trials. | [optional] |
| **num_lead_trials** | **Integer** | Distinct lead-sponsored trials of the peer&#39;s subtree. | [optional] |
| **size_band** | **String** | Band of num_lead_trials: 1-4, 5-19, 20-99 or 100+. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyPeersRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  peer_company_id: null,
  peer_company_name: null,
  peer_group_id: null,
  peer_group_name: null,
  peer_rank: null,
  match_level: null,
  matched_area: null,
  matched_size: null,
  lead_stage: null,
  lead_drug_name: null,
  biologic_share: null,
  modality_band: null,
  primary_area: null,
  num_lead_trials: null,
  size_band: null
)
```

