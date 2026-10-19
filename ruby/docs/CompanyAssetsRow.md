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
| **max_phase** | **Integer** | Highest phase the sponsor reached on the drug: 0 early phase 1 through 6 for phase 4 (marketed); -1 when unstated. | [optional] |
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
  n_trials: null,
  n_active: null,
  last_year: null,
  withdrawn: null
)
```

