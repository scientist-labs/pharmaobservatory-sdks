# PharmaObservatory::DrugRecallsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **recall_number** | **String** | The FDA recall number. | [optional] |
| **classification** | **String** | The FDA recall class (Class I, Class II or Class III). | [optional] |
| **status** | **String** | The recall status (for example Ongoing, Completed, Terminated). | [optional] |
| **recalling_firm** | **String** | The firm that recalled the product, as filed. | [optional] |
| **product_description** | **String** | The recalled product as described in the report. | [optional] |
| **reason_for_recall** | **String** | The reason given for the recall. | [optional] |
| **recall_initiation_date** | **Date** | The date the recall began. | [optional] |
| **report_date** | **Date** | The date the FDA posted the report. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugRecallsRow.new(
  molecule_id: null,
  recall_number: null,
  classification: null,
  status: null,
  recalling_firm: null,
  product_description: null,
  reason_for_recall: null,
  recall_initiation_date: null,
  report_date: null
)
```

