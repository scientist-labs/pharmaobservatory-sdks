# PharmaObservatory::BiomarkersRecord

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **object_type** | **String** |  |  |
| **id** | **String** |  |  |
| **record** | [**BiomarkersRow**](BiomarkersRow.md) |  |  |
| **attributes** | **Object** |  | [optional] |
| **datasets** | [**Array&lt;DatasetListing&gt;**](DatasetListing.md) |  |  |
| **attribution** | [**Array&lt;DatasetAttribution&gt;**](DatasetAttribution.md) |  | [optional] |
| **source_credit** | **String** |  | [optional] |
| **freshness** | [**AnimalWelfareAssurancesResponseFreshness**](AnimalWelfareAssurancesResponseFreshness.md) |  | [optional] |
| **build** | [**AnimalWelfareAssurancesResponseBuild**](AnimalWelfareAssurancesResponseBuild.md) |  | [optional] |
| **build_state** | [**DatasetBuildState**](DatasetBuildState.md) |  | [optional] |
| **restriction** | **Hash&lt;String, Object&gt;** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkersRecord.new(
  object_type: null,
  id: null,
  record: null,
  attributes: null,
  datasets: null,
  attribution: null,
  source_credit: null,
  freshness: null,
  build: null,
  build_state: null,
  restriction: null
)
```

