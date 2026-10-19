# PharmaObservatory::SiteHospitalsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **site_id** | **String** | The site key: one facility spelling from the registry. | [optional] |
| **ccn** | **String** | CMS Certification Number of the matched hospital. | [optional] |
| **hospital_name** | **String** | The hospital&#39;s name in the CMS registry. | [optional] |
| **hospital_type** | **String** | The CMS hospital type. | [optional] |
| **ownership** | **String** | The CMS ownership category. | [optional] |
| **bed_count** | **Integer** | Certified bed count in the latest registry file. | [optional] |
| **is_teaching** | **Boolean** | Whether the hospital is a teaching hospital. | [optional] |
| **hospital_latitude** | **Float** | Latitude of the hospital. | [optional] |
| **hospital_longitude** | **Float** | Longitude of the hospital. | [optional] |
| **geocode_method** | **String** | How the hospital&#39;s coordinate was placed. | [optional] |
| **ccn_n_trials** | **Integer** | Registered trials across the hospital&#39;s matched site records. | [optional] |
| **ccn_n_site_variants** | **Integer** | Site records matched to the same hospital. | [optional] |
| **is_primary_site** | **Boolean** | True for the matched site record with the most trials for this hospital. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SiteHospitalsRow.new(
  site_id: null,
  ccn: null,
  hospital_name: null,
  hospital_type: null,
  ownership: null,
  bed_count: null,
  is_teaching: null,
  hospital_latitude: null,
  hospital_longitude: null,
  geocode_method: null,
  ccn_n_trials: null,
  ccn_n_site_variants: null,
  is_primary_site: null
)
```

