# PharmaObservatory::InspectionFindingsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **citation_id** | **String** | The citation id. | [optional] |
| **inspection_id** | **String** | The inspection the citation belongs to; join to the inspections dataset. | [optional] |
| **fei** | **String** | The FDA establishment identifier (FEI) of the inspected site. | [optional] |
| **firm_name** | **String** | The firm&#39;s name as filed with the regulator. | [optional] |
| **firm_org_id** | **String** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] |
| **firm_company_id** | **String** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] |
| **firm_company_name** | **String** | Display name of the firm&#39;s company. | [optional] |
| **firm_group_id** | **String** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **firm_group_name** | **String** | Display name of the firm&#39;s group. | [optional] |
| **is_org_resolved** | **Boolean** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] |
| **match_tier** | **String** | How well the firm was attached to an organization: high, medium, low, registry, ambiguous or unresolved. | [optional] |
| **inspection_date** | **Date** | The date of the inspection. | [optional] |
| **classification** | **String** | The outcome of the inspection: NAI, VAI or OAI. | [optional] |
| **act_cfr_number** | **String** | The cited regulation (a Code of Federal Regulations section or Act citation). | [optional] |
| **short_description** | **String** | The FDA&#39;s short description of the observation. | [optional] |
| **program_area** | **String** | The FDA program area of the citation. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InspectionFindingsRow.new(
  citation_id: null,
  inspection_id: null,
  fei: null,
  firm_name: null,
  firm_org_id: null,
  firm_company_id: null,
  firm_company_name: null,
  firm_group_id: null,
  firm_group_name: null,
  is_org_resolved: null,
  match_tier: null,
  inspection_date: null,
  classification: null,
  act_cfr_number: null,
  short_description: null,
  program_area: null
)
```

