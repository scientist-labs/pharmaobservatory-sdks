# PharmaObservatory::ApiV1SiteRecordGet200ResponseSitesInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **site_key** | **String** |  |  |
| **facility_name** | **String** |  |  |
| **city** | **String** |  |  |
| **state** | **String** |  |  |
| **country** | **String** |  |  |
| **n_trials** | **Integer** |  |  |
| **n_active** | **Integer** |  |  |
| **active_years** | [**ApiV1SiteRecordGet200ResponseSitesInnerActiveYears**](ApiV1SiteRecordGet200ResponseSitesInnerActiveYears.md) |  |  |
| **latitude** | **String** |  | [optional] |
| **longitude** | **String** |  | [optional] |
| **trials** | [**Array&lt;ApiV1SiteRecordGet200ResponseSitesInnerTrialsInner&gt;**](ApiV1SiteRecordGet200ResponseSitesInnerTrialsInner.md) |  | [optional] |
| **indications** | [**Array&lt;ApiV1SiteRecordGet200ResponseSitesInnerIndicationsInner&gt;**](ApiV1SiteRecordGet200ResponseSitesInnerIndicationsInner.md) |  | [optional] |
| **drugs** | [**Array&lt;ApiV1SiteRecordGet200ResponseSitesInnerDrugsInner&gt;**](ApiV1SiteRecordGet200ResponseSitesInnerDrugsInner.md) |  | [optional] |
| **estimated_enrollment** | [**ApiV1SiteRecordGet200ResponseSitesInnerEstimatedEnrollment**](ApiV1SiteRecordGet200ResponseSitesInnerEstimatedEnrollment.md) |  | [optional] |
| **investigators** | [**Array&lt;ApiV1SiteRecordGet200ResponseSitesInnerInvestigatorsInner&gt;**](ApiV1SiteRecordGet200ResponseSitesInnerInvestigatorsInner.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1SiteRecordGet200ResponseSitesInner.new(
  site_key: null,
  facility_name: null,
  city: null,
  state: null,
  country: null,
  n_trials: null,
  n_active: null,
  active_years: null,
  latitude: null,
  longitude: null,
  trials: null,
  indications: null,
  drugs: null,
  estimated_enrollment: null,
  investigators: null
)
```

