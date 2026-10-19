# PharmaObservatory::TrialSitesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **site_key** | **String** | A site key built from the normalized facility name, city, state and country. | [optional] |
| **listing_ordinal** | **Integer** | 1 for the first listing of the site on the trial; higher for a repeat listing of the same site on the same trial. | [optional] |
| **facility_name** | **String** | The facility name as filed. | [optional] |
| **city** | **String** | City. | [optional] |
| **state** | **String** | State or region. | [optional] |
| **country** | **String** | Country. | [optional] |
| **postal** | **String** | Postal code. | [optional] |
| **latitude** | **Float** | Latitude of the site, in degrees. | [optional] |
| **longitude** | **Float** | Longitude of the site, in degrees. | [optional] |
| **geo_precision** | **String** | How exact the coordinate is (postal, city and so on). | [optional] |
| **coord_source** | **String** | Where the coordinate came from: aact, geonames or als. | [optional] |
| **trial_status** | **String** | The trial&#39;s overall status code. | [optional] |
| **trial_start_year** | **Integer** | Year the trial started. | [optional] |
| **trial_title** | **String** | The trial&#39;s brief title. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialSitesRow.new(
  trial_id: null,
  site_key: null,
  listing_ordinal: null,
  facility_name: null,
  city: null,
  state: null,
  country: null,
  postal: null,
  latitude: null,
  longitude: null,
  geo_precision: null,
  coord_source: null,
  trial_status: null,
  trial_start_year: null,
  trial_title: null
)
```

