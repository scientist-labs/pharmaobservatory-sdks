# PharmaObservatory::ConditionSiteYearsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **site_id** | **String** | The site key (a normalized facility, city and country). | [optional] |
| **start_year** | **Integer** | Year the trials started. Null for a trial with no start date. | [optional] |
| **site_name** | **String** | The facility name as the registry lists it. | [optional] |
| **city** | **String** | City as listed. | [optional] |
| **state** | **String** | State or province as listed. | [optional] |
| **country** | **String** | Country as listed. | [optional] |
| **latitude** | **Float** | Latitude of the finest coordinate the site record has. | [optional] |
| **longitude** | **Float** | Longitude of the finest coordinate the site record has. | [optional] |
| **n_trials** | **Integer** | Trials of the condition that started that year and list the site. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionSiteYearsRow.new(
  condition_id: null,
  site_id: null,
  start_year: null,
  site_name: null,
  city: null,
  state: null,
  country: null,
  latitude: null,
  longitude: null,
  n_trials: null
)
```

