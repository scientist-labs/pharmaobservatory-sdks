# PharmaObservatory::CompanyProgramsRow

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
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The condition&#39;s name. | [optional] |
| **max_phase** | **Integer** | Highest phase reached: 0 early phase 1 through 6 for phase 4 (marketed); -1 when unstated. | [optional] |
| **n_trials** | **Integer** | Lead-sponsored trials in the program. | [optional] |
| **n_active** | **Integer** | Of those, trials in an active status. | [optional] |
| **last_year** | **Integer** | Latest trial start year. | [optional] |
| **withdrawn** | **Boolean** | True when the drug has been withdrawn. | [optional] |
| **lifecycle** | **String** | approved, active or abandoned. | [optional] |
| **owned** | **Boolean** | True when the sponsor is judged to develop the drug (an experimental arm, a regulatory holder, or an inferred lead); false when it is the backbone or comparator of a combination, one of many leads of a crowded drug, or not the holder of a generic. Every row is kept; this column labels it. | [optional] |
| **ownership_basis** | **String** | Why owned is what it is: experimental_arm, regulatory_holder or inferred_lead (owned), or backbone, comparator, crowded_not_lead or generic_not_holder (not owned). inferred_lead is an inference made where the trials give no arm evidence either way. | [optional] |
| **dalys_millions** | **Float** | Disability-adjusted life years (DALYs, millions) of the condition&#39;s global burden, from the Global Burden of Disease (GBD) 2021 study of the Institute for Health Metrics and Evaluation (IHME). Null for a caller without the licence for this column; see withheld_columns. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyProgramsRow.new(
  sponsor_org_id: null,
  sponsor_name: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  drug_id: null,
  drug_name: null,
  condition_id: null,
  condition_name: null,
  max_phase: null,
  n_trials: null,
  n_active: null,
  last_year: null,
  withdrawn: null,
  lifecycle: null,
  owned: null,
  ownership_basis: null,
  dalys_millions: null
)
```

