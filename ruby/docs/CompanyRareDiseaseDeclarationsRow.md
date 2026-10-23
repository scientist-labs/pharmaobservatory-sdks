# PharmaObservatory::CompanyRareDiseaseDeclarationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **assertion_id** | **String** | The declaration assertion id. | [optional] |
| **organization_name** | **String** | The organization&#39;s name as the source filed it. | [optional] |
| **organization_org_id** | **String** | Legal-entity level: the organization record of the organization. | [optional] |
| **organization_company_id** | **String** | Company level: the company the organization belongs to (null when unresolved). | [optional] |
| **organization_company_name** | **String** | Display name of the organization&#39;s company. | [optional] |
| **organization_group_id** | **String** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **organization_group_name** | **String** | Display name of the organization&#39;s group. | [optional] |
| **declared_function** | **String** | The function the source states for the organization. | [optional] |
| **evidence_type** | **String** | The kind of evidence the source gives. | [optional] |
| **authority_class** | **String** | How authoritative the source is for the claim. | [optional] |
| **source_key** | **String** | The source the assertion comes from. | [optional] |
| **source_record_id** | **String** | The source&#39;s own record id. | [optional] |
| **source_url** | **String** | A link to the source record. | [optional] |
| **observed_at** | **Time** | When the assertion was observed (UTC). | [optional] |
| **last_verified_at** | **Time** | When the assertion was last verified (UTC). | [optional] |
| **confidence** | **Float** | The source&#39;s or the extraction&#39;s own confidence; null when none is stated. | [optional] |
| **license_name** | **String** | The licence of the source. | [optional] |
| **attribution** | **String** | The credit line for the source. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyRareDiseaseDeclarationsRow.new(
  assertion_id: null,
  organization_name: null,
  organization_org_id: null,
  organization_company_id: null,
  organization_company_name: null,
  organization_group_id: null,
  organization_group_name: null,
  declared_function: null,
  evidence_type: null,
  authority_class: null,
  source_key: null,
  source_record_id: null,
  source_url: null,
  observed_at: null,
  last_verified_at: null,
  confidence: null,
  license_name: null,
  attribution: null
)
```

