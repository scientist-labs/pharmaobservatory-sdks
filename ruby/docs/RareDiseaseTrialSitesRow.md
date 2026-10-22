# PharmaObservatory::RareDiseaseTrialSitesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **mapping_arm** | **String** | exact or unspecified: the mapping arm the site&#39;s trials belong to. | [optional] |
| **site_id** | **String** | The site key: one facility spelling from the registry; joins the sites dataset. | [optional] |
| **site_rank** | **Integer** | The site&#39;s rank by trial count within its disease and arm (1 is the most trials); a position in a list, not a quality score. | [optional] |
| **site_name** | **String** | The facility name as filed in the registry. | [optional] |
| **city** | **String** | The facility&#39;s city as filed. | [optional] |
| **state** | **String** | The facility&#39;s state or region as filed. | [optional] |
| **country** | **String** | The facility&#39;s country as filed. | [optional] |
| **latitude** | **Float** | The site&#39;s latitude in degrees; null when no coordinate is known. | [optional] |
| **longitude** | **Float** | The site&#39;s longitude in degrees; null when no coordinate is known. | [optional] |
| **coord_source** | **String** | The source of the coordinate: aact (the registry&#39;s own), geonames (GeoNames) or als (Amazon Location Service). | [optional] |
| **condition_ids** | **Array&lt;String&gt;** | The MeSH condition ids through which the disease reaches the site&#39;s trials. | [optional] |
| **n_trials** | **Integer** | Registered trials of this disease and arm that list the site. | [optional] |
| **n_trials_direct** | **Integer** | Those trials whose own condition list names the mapped condition. | [optional] |
| **n_active_trials** | **Integer** | Those trials that are recruiting, not yet recruiting, active or enrolling by invitation. | [optional] |
| **first_trial_year** | **Integer** | The earliest start year among those trials. | [optional] |
| **last_trial_year** | **Integer** | The latest start year among those trials. | [optional] |
| **last_update_posted_date** | **Date** | The latest registry update among those trials. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseTrialSitesRow.new(
  orpha_code: null,
  mapping_arm: null,
  site_id: null,
  site_rank: null,
  site_name: null,
  city: null,
  state: null,
  country: null,
  latitude: null,
  longitude: null,
  coord_source: null,
  condition_ids: null,
  n_trials: null,
  n_trials_direct: null,
  n_active_trials: null,
  first_trial_year: null,
  last_trial_year: null,
  last_update_posted_date: null
)
```

