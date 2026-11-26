# PharmaObservatory::RareDiseasesRecord

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sections** | **Hash&lt;String, Object&gt;** | Sections of the object&#39;s web page, by key: patient_organisations, the list the page shows of the organisations that name the disease or condition on their own websites (state, empty_state, coverage, groups of entries carrying the patient_organisation dataset&#39;s columns, the per-entry credit and the lane attribution). | [optional] |
| **object_type** | **String** |  |  |
| **id** | **String** |  |  |
| **record** | [**RareDiseasesRow**](RareDiseasesRow.md) |  |  |
| **attributes** | [**RareDiseasesRecordAttributes**](RareDiseasesRecordAttributes.md) |  | [optional] |
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

instance = PharmaObservatory::RareDiseasesRecord.new(
  sections: null,
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

