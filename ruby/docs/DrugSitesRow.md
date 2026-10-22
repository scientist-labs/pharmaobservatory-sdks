# PharmaObservatory::DrugSitesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **site_id** | **String** | The site key: one facility spelling from the registry. | [optional] |
| **site_name** | **String** | The facility name as registered. | [optional] |
| **city** | **String** | The facility&#39;s city. | [optional] |
| **country** | **String** | The facility&#39;s country. | [optional] |
| **n_trials** | **Integer** | The drug&#39;s trials registered at the site. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugSitesRow.new(
  molecule_id: null,
  site_id: null,
  site_name: null,
  city: null,
  country: null,
  n_trials: null
)
```

