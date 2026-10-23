# PharmaObservatory::CompanyAssetsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sponsor_org_id** | **String** | Legal-entity level: the lead sponsor&#39;s organization record. | [optional] |
| **sponsor_name** | **String** | The sponsor&#39;s organization name. | [optional] |
| **sponsor_company_id** | **String** | Company level: the sponsor&#39;s company (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **drug_id** | **String** | The drug key. | [optional] |
| **drug_name** | **String** | The drug&#39;s name. | [optional] |
| **max_phase** | **Integer** | Highest phase the sponsor reached on the drug on the 0 to 4 scale: 0 not reported, 1 through 4 for phase 1 through phase 4 (4 is marketed); a split phase folds up to its higher phase. | [optional] |
| **max_phase_label** | **String** | The furthest phase reached with the split phases kept: Early 1, Phase 1, Phase 1/2, Phase 2, Phase 2/3, Phase 3, Phase 4, or Not reported. max_phase is the same value folded onto 0 to 4 (a split phase counts as its higher phase). | [optional] |
| **max_phase_fine_rank** | **Integer** | The finest phase rank behind max_phase_label: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early 1, -1 not reported. Sort on it to order programs by furthest phase with the split phases kept; max_phase is the same value folded onto 0 to 4. | [optional] |
| **n_trials** | **Integer** | Distinct trials this sponsor leads on the drug. Counted once per trial within the sponsor organization, so it is not additive across a company&#39;s organizations. | [optional] |
| **n_active** | **Integer** | Of those trials, the number recruiting, not yet recruiting, active or enrolling by invitation. | [optional] |
| **last_year** | **Integer** | Latest start year among those trials. | [optional] |
| **withdrawn** | **Boolean** | True when the drug has been withdrawn from the market. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyAssetsRow.new(
  sponsor_org_id: null,
  sponsor_name: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  drug_id: null,
  drug_name: null,
  max_phase: null,
  max_phase_label: null,
  max_phase_fine_rank: null,
  n_trials: null,
  n_active: null,
  last_year: null,
  withdrawn: null
)
```

