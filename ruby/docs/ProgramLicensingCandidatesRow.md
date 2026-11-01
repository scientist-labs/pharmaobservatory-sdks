# PharmaObservatory::ProgramLicensingCandidatesRow

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
| **condition_ancestor_ids** | **Array&lt;String&gt;** | Every MeSH descriptor id at or above the condition in the tree, the condition itself included. Filter a subtree with &#x3D; ANY. | [optional] |
| **max_phase** | **Integer** | Highest phase reached: 0 early phase 1 through 6 for phase 4 (marketed); -1 when unstated. | [optional] |
| **n_trials** | **Integer** | Lead-sponsored trials in the program. | [optional] |
| **n_active** | **Integer** | Of those, trials in an active status. | [optional] |
| **last_year** | **Integer** | Latest trial start year. | [optional] |
| **last_update_year** | **Integer** | Year the program&#39;s most recently updated trial record was posted. | [optional] |
| **withdrawn** | **Boolean** | True when the drug has been withdrawn. | [optional] |
| **owned** | **Boolean** | True when the sponsor is judged to develop the drug (an experimental arm, a regulatory holder, or an inferred lead); false when it is the backbone or comparator of a combination, one of many leads of a crowded drug, or not the holder of a generic. Every row is kept; this column labels it. | [optional] |
| **ownership_basis** | **String** | Why owned is what it is: experimental_arm, regulatory_holder or inferred_lead (owned), or backbone, comparator, crowded_not_lead or generic_not_holder (not owned). inferred_lead is an inference made where the trials give no arm evidence either way. | [optional] |
| **cause_name** | **String** | The Global Burden of Disease cause matched to the condition. | [optional] |
| **dalys_millions** | **Float** | Global disability-adjusted life years for the matched cause, in millions. | [optional] |
| **deaths_millions** | **Float** | Global deaths for the matched cause, in millions. | [optional] |
| **druggability** | **Integer** | Druggability class of the matched cause: 1 highly druggable, 2 druggable, 3 challenging, 4 very challenging. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ProgramLicensingCandidatesRow.new(
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
  condition_ancestor_ids: null,
  max_phase: null,
  n_trials: null,
  n_active: null,
  last_year: null,
  last_update_year: null,
  withdrawn: null,
  owned: null,
  ownership_basis: null,
  cause_name: null,
  dalys_millions: null,
  deaths_millions: null,
  druggability: null
)
```

