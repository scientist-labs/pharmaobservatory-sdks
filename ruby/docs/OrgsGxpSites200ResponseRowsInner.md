# PharmaObservatory::OrgsGxpSites200ResponseRowsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **facility_name** | **String** |  |  |
| **city** | **String** |  |  |
| **country** | **String** |  |  |
| **regulators** | **Array&lt;String&gt;** |  |  |
| **canonical_org_id** | **String** |  |  |
| **org_name** | **String** |  |  |
| **facility_key** | **String** |  | [optional] |
| **state_region** | **Object** |  | [optional] |
| **severity_level** | **String** |  | [optional] |
| **severity_raw_verdict** | **String** |  | [optional] |
| **severity_signal_date** | **String** |  | [optional] |
| **match_tier** | **String** |  | [optional] |
| **restriction_tier** | **String** |  | [optional] |
| **register_domain** | **String** |  | [optional] |
| **latitude** | **Object** |  | [optional] |
| **longitude** | **Object** |  | [optional] |
| **geocode_precision** | **Object** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsGxpSites200ResponseRowsInner.new(
  facility_name: null,
  city: null,
  country: null,
  regulators: null,
  canonical_org_id: null,
  org_name: null,
  facility_key: null,
  state_region: null,
  severity_level: null,
  severity_raw_verdict: null,
  severity_signal_date: null,
  match_tier: null,
  restriction_tier: null,
  register_domain: null,
  latitude: null,
  longitude: null,
  geocode_precision: null
)
```

