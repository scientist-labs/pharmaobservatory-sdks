# PharmaObservatory::DrugAspCodesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **hcpcs** | **String** | The HCPCS billing code. | [optional] |
| **hcpcs_desc** | **String** | The code&#39;s short description. | [optional] |
| **hcpcs_dosage** | **String** | The billing unit the code is priced per. | [optional] |
| **asp_payment_limit** | **Float** | The Medicare payment limit per billing unit. | [optional] |
| **priced_ndcs** | **Integer** | Packages priced under the code. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugAspCodesRow.new(
  molecule_id: null,
  hcpcs: null,
  hcpcs_desc: null,
  hcpcs_dosage: null,
  asp_payment_limit: null,
  priced_ndcs: null
)
```

