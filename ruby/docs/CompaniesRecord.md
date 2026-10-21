# PharmaObservatory::CompaniesRecord

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **object_type** | **String** |  |  |
| **id** | **String** |  |  |
| **record** | [**CompaniesRow**](CompaniesRow.md) |  |  |
| **attributes** | [**CompaniesRecordAttributes**](CompaniesRecordAttributes.md) |  | [optional] |
| **datasets** | [**Array&lt;DatasetListing&gt;**](DatasetListing.md) |  |  |
| **attribution** | [**Array&lt;DatasetAttribution&gt;**](DatasetAttribution.md) |  | [optional] |
| **source_credit** | **String** |  | [optional] |
| **freshness** | [**BiomarkerAssayEvidenceResponseFreshness**](BiomarkerAssayEvidenceResponseFreshness.md) |  | [optional] |
| **build** | [**BiomarkerAssayEvidenceResponseBuild**](BiomarkerAssayEvidenceResponseBuild.md) |  | [optional] |
| **build_state** | [**DatasetBuildState**](DatasetBuildState.md) |  | [optional] |
| **restriction** | **Hash&lt;String, Object&gt;** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompaniesRecord.new(
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

