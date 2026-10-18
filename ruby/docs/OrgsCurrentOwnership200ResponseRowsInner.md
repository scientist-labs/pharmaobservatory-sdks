# PharmaObservatory::OrgsCurrentOwnership200ResponseRowsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **owner_rollup_id** | **String** |  |  |
| **direction** | **String** |  |  |
| **source_class** | **String** |  |  |
| **source_system** | **String** |  |  |
| **ownership_dimension** | **String** |  |  |
| **relationship_type** | **String** |  |  |
| **owner_display_name** | **String** |  |  |
| **owner_kind** | **String** |  |  |
| **owner_org_id** | **Object** |  |  |
| **owner_investor_id** | **Object** |  |  |
| **subject_name_raw** | **String** |  |  |
| **subject_org_id** | **String** |  |  |
| **reported_percent** | **Float** |  |  |
| **percentage_basis** | **String** |  |  |
| **percent_basis** | **String** |  |  |
| **equity_shares_standing** | **Float** |  |  |
| **shares_basis** | **String** |  |  |
| **edge_state** | **String** |  |  |
| **position_count** | **Integer** |  |  |
| **standing_position_count** | **Integer** |  |  |
| **ceased_position_count** | **Integer** |  |  |
| **no_longer_reported_position_count** | **Integer** |  |  |
| **last_activity_date** | **String** |  |  |
| **restriction_tier** | **String** |  |  |
| **source_license** | **String** |  |  |
| **attribution** | **String** |  |  |
| **contains_personal_data** | **Boolean** |  | [optional] |
| **privacy_status** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsCurrentOwnership200ResponseRowsInner.new(
  owner_rollup_id: null,
  direction: null,
  source_class: null,
  source_system: null,
  ownership_dimension: null,
  relationship_type: null,
  owner_display_name: null,
  owner_kind: null,
  owner_org_id: null,
  owner_investor_id: null,
  subject_name_raw: null,
  subject_org_id: null,
  reported_percent: null,
  percentage_basis: null,
  percent_basis: null,
  equity_shares_standing: null,
  shares_basis: null,
  edge_state: null,
  position_count: null,
  standing_position_count: null,
  ceased_position_count: null,
  no_longer_reported_position_count: null,
  last_activity_date: null,
  restriction_tier: null,
  source_license: null,
  attribution: null,
  contains_personal_data: null,
  privacy_status: null
)
```

