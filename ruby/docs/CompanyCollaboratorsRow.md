# PharmaObservatory::CompanyCollaboratorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The MeSH descriptor name of the condition. | [optional] |
| **disease_rank** | **Integer** | The condition&#39;s rank among the company&#39;s disease areas by trial count (1 is the busiest). Only the top 10 are kept. | [optional] |
| **collaborator_company_id** | **String** | Company level: the company that co-sponsors the trials. | [optional] |
| **collaborator_company_name** | **String** | Display name of the collaborator&#39;s company. | [optional] |
| **collaborator_group_id** | **String** | Group level: the collaborator&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **collaborator_group_name** | **String** | Display name of the collaborator&#39;s group. | [optional] |
| **collaborator_rank** | **Integer** | 0 for the company itself, 1 to 6 for its collaborators ranked by shared trials across its top disease areas. | [optional] |
| **shared_trials** | **Integer** | Distinct trials the company shares with the collaborator in this disease area (for rank 0, the company&#39;s own trials in it). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyCollaboratorsRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  condition_id: null,
  condition_name: null,
  disease_rank: null,
  collaborator_company_id: null,
  collaborator_company_name: null,
  collaborator_group_id: null,
  collaborator_group_name: null,
  collaborator_rank: null,
  shared_trials: null
)
```

