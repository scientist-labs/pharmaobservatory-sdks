# PharmaObservatory::BiomarkerLiteratureOrgsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The route key of the gene, the entry key for the biomarkers object type. | [optional] |
| **org_id** | **String** | Legal-entity level: the resolved parent organization record the byline names. | [optional] |
| **org_name** | **String** | The organization&#39;s name. | [optional] |
| **org_type** | **String** | The organization&#39;s type. | [optional] |
| **org_country** | **String** | The organization&#39;s country code (null when not recorded). | [optional] |
| **n_papers** | **Integer** | Distinct papers naming the organization on a byline. | [optional] |
| **n_trial_linked_papers** | **Integer** | The part of n_papers a registered clinical trial cites. | [optional] |
| **latest_pub_year** | **Integer** | The newest publication year among those papers. | [optional] |
| **org_arm** | **String** | company when the organization is company-typed; institution for every other organization. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerLiteratureOrgsRow.new(
  route_key: null,
  org_id: null,
  org_name: null,
  org_type: null,
  org_country: null,
  n_papers: null,
  n_trial_linked_papers: null,
  latest_pub_year: null,
  org_arm: null
)
```

