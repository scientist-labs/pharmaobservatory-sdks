# PharmaObservatory::HospitalsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ccn** | **String** | CMS Certification Number, the six-character Medicare provider id the hospital page address carries. | [optional] |
| **hospital_name** | **String** | The hospital&#39;s facility name in the CMS registry. | [optional] |
| **city** | **String** | The hospital&#39;s city. | [optional] |
| **state** | **String** | The hospital&#39;s two-letter state or territory code. | [optional] |
| **bed_count** | **Integer** | Certified bed count in the latest registry file. Null when the file has no bed record. | [optional] |
| **hospital_type** | **String** | The CMS hospital type. | [optional] |
| **ownership** | **String** | The CMS ownership category. | [optional] |
| **is_teaching** | **Boolean** | Whether the hospital is a teaching hospital. | [optional] |
| **emergency_services** | **String** | Whether the hospital provides emergency services, as CMS files it. Null where CMS left it blank. | [optional] |
| **latitude** | **Float** | Latitude of the hospital&#39;s own coordinate. Null exactly when geocode_method is unmatched. | [optional] |
| **longitude** | **Float** | Longitude of the hospital&#39;s own coordinate. Null exactly when geocode_method is unmatched. | [optional] |
| **geocode_method** | **String** | How the coordinate was placed: als (Amazon Location Service address match), zip_centroid (coarse ZIP-code centroid) or unmatched (no coordinate). | [optional] |
| **hrr_num** | **Integer** | Dartmouth Atlas Hospital Referral Region number. Null when hrr_match is not zip. | [optional] |
| **hrr_name** | **String** | The Hospital Referral Region, named for its hub city and state. The hub state can differ from the hospital&#39;s own state. | [optional] |
| **hsa_num** | **Integer** | Dartmouth Atlas Hospital Service Area number. Null when hrr_match is not zip. | [optional] |
| **hsa_name** | **String** | The Hospital Service Area, named for its hub city and state. | [optional] |
| **hrr_match** | **String** | zip (region found by ZIP), territory_uncovered (the Atlas does not cover US territories) or zip_unmatched (a mainland ZIP absent from the 2019 crosswalk). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::HospitalsRow.new(
  ccn: null,
  hospital_name: null,
  city: null,
  state: null,
  bed_count: null,
  hospital_type: null,
  ownership: null,
  is_teaching: null,
  emergency_services: null,
  latitude: null,
  longitude: null,
  geocode_method: null,
  hrr_num: null,
  hrr_name: null,
  hsa_num: null,
  hsa_name: null,
  hrr_match: null
)
```

