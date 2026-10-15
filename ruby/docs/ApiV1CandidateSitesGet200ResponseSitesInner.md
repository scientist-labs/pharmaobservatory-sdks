# PharmaObservatory::ApiV1CandidateSitesGet200ResponseSitesInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **rank** | **Integer** |  |  |
| **site_key** | **String** |  |  |
| **facility_name** | **String** |  |  |
| **city** | **String** |  |  |
| **state** | **String** |  |  |
| **country** | **String** |  |  |
| **n_trials_in_indication** | **Integer** |  |  |
| **estimated_patients_in_indication** | **String** |  |  |
| **n_active_in_indication** | **Integer** |  |  |
| **n_active_same_phase** | **Object** |  |  |
| **composite_score** | **Float** |  |  |
| **capacity_z** | **Float** |  |  |
| **activity_z** | **Float** |  |  |
| **competition_z** | **Float** |  |  |
| **saturated** | **Boolean** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1CandidateSitesGet200ResponseSitesInner.new(
  rank: null,
  site_key: null,
  facility_name: null,
  city: null,
  state: null,
  country: null,
  n_trials_in_indication: null,
  estimated_patients_in_indication: null,
  n_active_in_indication: null,
  n_active_same_phase: null,
  composite_score: null,
  capacity_z: null,
  activity_z: null,
  competition_z: null,
  saturated: null
)
```

