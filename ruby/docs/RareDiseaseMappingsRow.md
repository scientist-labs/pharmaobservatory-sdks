# PharmaObservatory::RareDiseaseMappingsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **xref_assertion_id** | **String** | The identifier of the cross-reference assertion. | [optional] |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **target_namespace** | **String** | The vocabulary the target code belongs to (ICD-10, ICD-11, MESH, MONDO, UMLS, GARD, ORPHA and others). | [optional] |
| **target_code** | **String** | The code in the target vocabulary. | [optional] |
| **target_curie** | **String** | The target as a prefixed identifier (namespace and code). | [optional] |
| **mapping_relation** | **String** | The relation the source states from the disease to the target: exact, narrower, broader, undetermined or unspecified. | [optional] |
| **assertion_source** | **String** | The file the assertion was read from: ordo, orphadata_product1 or mondo. | [optional] |
| **confidence** | **Float** | The source&#39;s confidence in the assertion, where it states one. | [optional] |
| **condition_id** | **String** | The MeSH descriptor id of the matching condition, on MESH rows; the entry key of the conditions dataset. | [optional] |
| **condition_name** | **String** | The matching condition&#39;s name, on MESH rows. | [optional] |
| **source_release** | **String** | The Orphadata Science release the row was read from. | [optional] |
| **source_url** | **String** | The URL of the source file the row was read from. | [optional] |
| **license_name** | **String** | The licence the assertion&#39;s file is published under. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseMappingsRow.new(
  xref_assertion_id: null,
  orpha_code: null,
  target_namespace: null,
  target_code: null,
  target_curie: null,
  mapping_relation: null,
  assertion_source: null,
  confidence: null,
  condition_id: null,
  condition_name: null,
  source_release: null,
  source_url: null,
  license_name: null,
  attribution: null
)
```

