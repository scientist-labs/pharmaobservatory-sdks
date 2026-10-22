# PharmaObservatory::RareDiseaseHierarchyRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **hierarchy_assertion_id** | **String** | The identifier of the subclass assertion. | [optional] |
| **orpha_code** | **String** | The ORPHA code of the child disease the assertion is entered on. | [optional] |
| **preferred_name** | **String** | The child disease&#39;s preferred English name. | [optional] |
| **parent_orpha_code** | **String** | The ORPHA code of the direct parent disease. | [optional] |
| **parent_preferred_name** | **String** | The parent disease&#39;s preferred English name. | [optional] |
| **parent_status** | **String** | The parent concept&#39;s lifecycle state. | [optional] |
| **child_status** | **String** | The child concept&#39;s lifecycle state. | [optional] |
| **predicate** | **String** | The relation asserted from child to parent (subclass_of). | [optional] |
| **source_release** | **String** | The Orphadata Science release the row was read from. | [optional] |
| **source_url** | **String** | The URL of the source file the row was read from. | [optional] |
| **license_name** | **String** | The licence of the ontology release. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseHierarchyRow.new(
  hierarchy_assertion_id: null,
  orpha_code: null,
  preferred_name: null,
  parent_orpha_code: null,
  parent_preferred_name: null,
  parent_status: null,
  child_status: null,
  predicate: null,
  source_release: null,
  source_url: null,
  license_name: null,
  attribution: null
)
```

