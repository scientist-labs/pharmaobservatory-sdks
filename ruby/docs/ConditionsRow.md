# PharmaObservatory::ConditionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The condition&#39;s MeSH heading. | [optional] |
| **primary_category** | **String** | The condition&#39;s primary MeSH disease category. | [optional] |
| **categories** | **Array&lt;String&gt;** | The MeSH tree branches the condition sits under (C and F03 codes). | [optional] |
| **is_neoplasm** | **Boolean** | True when the condition sits under the MeSH neoplasms branch (C04). | [optional] |
| **tree_numbers** | **Array&lt;String&gt;** | The condition&#39;s MeSH tree numbers. | [optional] |
| **n_trials** | **Integer** | Trials tagged with the condition or with a condition beneath it in the MeSH tree. | [optional] |
| **n_direct_trials** | **Integer** | Trials tagged with the condition itself. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionsRow.new(
  condition_id: null,
  condition_name: null,
  primary_category: null,
  categories: null,
  is_neoplasm: null,
  tree_numbers: null,
  n_trials: null,
  n_direct_trials: null
)
```

