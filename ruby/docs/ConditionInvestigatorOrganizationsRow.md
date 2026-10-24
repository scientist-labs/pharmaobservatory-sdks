# PharmaObservatory::ConditionInvestigatorOrganizationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition&#39;s MeSH descriptor id, the id in the conditions page address (D001943). The entry key of the conditions object type. | [optional] |
| **condition_name** | **String** | The condition&#39;s MeSH name. | [optional] |
| **organization_name** | **String** | The organization&#39;s name in the registry that resolved it (ROR, else SEC EDGAR), never the investigator&#39;s affiliation text as filed. | [optional] |
| **organization_org_id** | **String** | Legal-entity level: the organization record the investigators&#39; affiliation resolved to. | [optional] |
| **organization_company_id** | **String** | Company level: the company the organization belongs to (null when unresolved). | [optional] |
| **organization_company_name** | **String** | Display name of the organization&#39;s company. | [optional] |
| **organization_group_id** | **String** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **organization_group_name** | **String** | Display name of the organization&#39;s group. | [optional] |
| **n_investigators** | **Integer** | Distinct investigators at the organization named on a trial of the condition, in any role. Always at least 2. | [optional] |
| **n_principal_investigators** | **Integer** | Of n_investigators, those who hold the principal investigator role on one of the condition&#39;s trials. | [optional] |
| **n_investigators_total** | **Integer** | Distinct investigators across the condition&#39;s trials, at any organization or none. A constant of the condition repeated on every row: never sum it. | [optional] |
| **n_at_resolved_organizations** | **Integer** | Of n_investigators_total, those whose affiliation resolved to an organization at a high-confidence tier, listed or not. A constant of the condition. | [optional] |
| **n_at_unresolved_affiliations** | **Integer** | Of n_investigators_total, those with no resolved organization, including those with no affiliation on file. A constant of the condition. | [optional] |
| **n_principal_investigators_total** | **Integer** | Of n_investigators_total, those who hold the principal investigator role on one of the condition&#39;s trials. A constant of the condition. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionInvestigatorOrganizationsRow.new(
  condition_id: null,
  condition_name: null,
  organization_name: null,
  organization_org_id: null,
  organization_company_id: null,
  organization_company_name: null,
  organization_group_id: null,
  organization_group_name: null,
  n_investigators: null,
  n_principal_investigators: null,
  n_investigators_total: null,
  n_at_resolved_organizations: null,
  n_at_unresolved_affiliations: null,
  n_principal_investigators_total: null
)
```

