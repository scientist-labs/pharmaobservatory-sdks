# PharmaObservatory::ConditionSitesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **site_id** | **String** | The site key (a normalized facility, city and country). | [optional] |
| **site_name** | **String** | The facility name as the registry lists it. | [optional] |
| **city** | **String** | City as listed. | [optional] |
| **country** | **String** | Country as listed. | [optional] |
| **n_trials** | **Integer** | Trials for this condition listed at the site. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionSitesRow.new(
  condition_id: null,
  site_id: null,
  site_name: null,
  city: null,
  country: null,
  n_trials: null
)
```

