# PharmaObservatory::CompanyMembersRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | The company the legal entity is a member of. | [optional] |
| **company_name** | **String** | The company&#39;s display name. | [optional] |
| **legal_entity_id** | **String** | A registry-scoped legal entity id (edgar:, ror:, lei: or rx_provider:). A registration, not an organization record id. | [optional] |
| **link_basis** | **String** | The source that placed the legal entity in the company (gleif, edgar, ror, name_key and so on). | [optional] |
| **parent_company_id** | **String** | The company&#39;s direct owner, null for a root. | [optional] |
| **parent_company_name** | **String** | Display name of the direct owner. | [optional] |
| **parent_link_basis** | **String** | The authority of the first ownership hop. | [optional] |
| **since** | **Date** | Date the ownership link began, when every hop is dated. | [optional] |
| **since_precision** | **String** | Precision of since. | [optional] |
| **since_source** | **String** | Where the date came from. | [optional] |
| **group_id** | **String** | Group level: the current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **hop_depth** | **Integer** | Hops from the company to its ultimate parent (0 for a root). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyMembersRow.new(
  company_id: null,
  company_name: null,
  legal_entity_id: null,
  link_basis: null,
  parent_company_id: null,
  parent_company_name: null,
  parent_link_basis: null,
  since: null,
  since_precision: null,
  since_source: null,
  group_id: null,
  group_name: null,
  hop_depth: null
)
```

