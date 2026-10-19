# PharmaObservatory::SiteInvestigatorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **site_id** | **String** | The site key: one facility spelling from the registry. | [optional] |
| **investigator_person_name** | **String** | The investigator&#39;s name as filed. | [optional] |
| **investigator_person_id** | **String** | The investigator&#39;s person key. | [optional] |
| **investigator_affiliation** | **String** | The affiliation as filed. | [optional] |
| **n_trials** | **Integer** | Trials at the site that name the investigator. | [optional] |
| **is_principal_investigator** | **Boolean** | True when the investigator is named principal investigator on a trial at the site. | [optional] |
| **site_rank** | **Integer** | Rank within the site: principal investigators first, then trial count. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SiteInvestigatorsRow.new(
  site_id: null,
  investigator_person_name: null,
  investigator_person_id: null,
  investigator_affiliation: null,
  n_trials: null,
  is_principal_investigator: null,
  site_rank: null
)
```

