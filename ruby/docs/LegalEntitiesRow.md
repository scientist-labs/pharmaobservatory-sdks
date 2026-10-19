# PharmaObservatory::LegalEntitiesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **org_id** | **String** | The organization record id: a registered legal entity (edgar:, ror:) or a name-resolved sponsor (sponsor:). | [optional] |
| **org_name** | **String** | The organization&#39;s name. | [optional] |
| **is_name_resolved_sponsor** | **Boolean** | True for a sponsor resolved from a filed name alone rather than from a registration. | [optional] |
| **org_origin** | **String** | How the organization record was minted. | [optional] |
| **country** | **String** | Country as the source states it. | [optional] |
| **org_type** | **String** | Organization type. | [optional] |
| **edgar_cik** | **String** | SEC EDGAR central index key, when it files. | [optional] |
| **edgar_ticker** | **String** | Exchange ticker, when listed. | [optional] |
| **ror_id** | **String** | Research Organization Registry id, when it has one. | [optional] |
| **lei** | **String** | Legal Entity Identifier, when it has one. | [optional] |
| **lifecycle_status** | **String** | Lifecycle status from the registry sources. | [optional] |
| **num_lead_trials** | **Integer** | Trials the organization leads. | [optional] |
| **num_trials** | **Integer** | Trials the organization leads or collaborates on. | [optional] |
| **total_lead_enrollment** | **Float** | Planned enrollment across led trials. | [optional] |
| **num_recent_lead_trials** | **Integer** | Led trials started in the recent window. | [optional] |
| **activity_score** | **Float** | Trial activity score. | [optional] |
| **company_id** | **String** | Company level: the company the organization belongs to (null when unresolved). | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::LegalEntitiesRow.new(
  org_id: null,
  org_name: null,
  is_name_resolved_sponsor: null,
  org_origin: null,
  country: null,
  org_type: null,
  edgar_cik: null,
  edgar_ticker: null,
  ror_id: null,
  lei: null,
  lifecycle_status: null,
  num_lead_trials: null,
  num_trials: null,
  total_lead_enrollment: null,
  num_recent_lead_trials: null,
  activity_score: null,
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null
)
```

