# PharmaObservatory::DrugApprovalsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **authority** | **String** | The approving authority. | [optional] |
| **jurisdiction** | **String** | The jurisdiction the authority covers. | [optional] |
| **status** | **String** | The approval status the source states. | [optional] |
| **first_approval_date** | **Date** | The earliest authorisation or first-marketed date the source gives. | [optional] |
| **approval_id** | **String** | The product number the authority uses (EMA product number or Health Canada drug identification number). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugApprovalsRow.new(
  molecule_id: null,
  authority: null,
  jurisdiction: null,
  status: null,
  first_approval_date: null,
  approval_id: null
)
```

