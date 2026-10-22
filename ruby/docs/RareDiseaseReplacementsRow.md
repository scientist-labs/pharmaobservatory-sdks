# PharmaObservatory::RareDiseaseReplacementsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **replacement_assertion_id** | **String** | The identifier of the replacement assertion. | [optional] |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **preferred_name** | **String** | The deprecated disease&#39;s preferred English name. | [optional] |
| **replacement_orpha_code** | **String** | The ORPHA code of the concept that replaces the deprecated one. | [optional] |
| **replacement_preferred_name** | **String** | The replacement concept&#39;s preferred English name. | [optional] |
| **replacement_status** | **String** | The replacement concept&#39;s lifecycle state. | [optional] |
| **predicate** | **String** | moved_to when Orphanet moved the concept to the replacement, referred_to when it referred the concept there. | [optional] |
| **assertion_source** | **String** | The Orphanet file the assertion was read from (ordo or orphadata_product1). | [optional] |
| **source_release** | **String** | The Orphadata Science release the row was read from. | [optional] |
| **source_url** | **String** | The URL of the source file the row was read from. | [optional] |
| **license_name** | **String** | The licence the assertion&#39;s file is published under. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseReplacementsRow.new(
  replacement_assertion_id: null,
  orpha_code: null,
  preferred_name: null,
  replacement_orpha_code: null,
  replacement_preferred_name: null,
  replacement_status: null,
  predicate: null,
  assertion_source: null,
  source_release: null,
  source_url: null,
  license_name: null,
  attribution: null
)
```

