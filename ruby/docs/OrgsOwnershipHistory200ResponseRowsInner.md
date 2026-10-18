# PharmaObservatory::OrgsOwnershipHistory200ResponseRowsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ownership_observation_id** | **String** |  |  |
| **direction** | **String** |  |  |
| **source_class** | **String** |  |  |
| **source_system** | **String** |  |  |
| **ownership_dimension** | **String** |  |  |
| **relationship_type** | **String** |  |  |
| **owner_name_raw** | **String** |  |  |
| **owner_source_ref** | **String** |  |  |
| **subject_name_raw** | **String** |  |  |
| **subject_source_ref** | **String** |  |  |
| **review_state** | **String** |  |  |
| **stake_value_type** | **String** |  |  |
| **stake_percent** | **Float** |  |  |
| **percentage_basis** | **String** |  |  |
| **observation_date** | **String** |  |  |
| **is_current_edge** | **Boolean** |  |  |
| **current_edge_state** | **String** |  |  |
| **restriction_tier** | **String** |  |  |
| **source_license** | **String** |  |  |
| **attribution** | **String** |  |  |
| **contains_personal_data** | **Boolean** |  | [optional] |
| **privacy_status** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsOwnershipHistory200ResponseRowsInner.new(
  ownership_observation_id: null,
  direction: null,
  source_class: null,
  source_system: null,
  ownership_dimension: null,
  relationship_type: null,
  owner_name_raw: null,
  owner_source_ref: null,
  subject_name_raw: null,
  subject_source_ref: null,
  review_state: null,
  stake_value_type: null,
  stake_percent: null,
  percentage_basis: null,
  observation_date: null,
  is_current_edge: null,
  current_edge_state: null,
  restriction_tier: null,
  source_license: null,
  attribution: null,
  contains_personal_data: null,
  privacy_status: null
)
```

