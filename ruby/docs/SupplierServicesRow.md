# PharmaObservatory::SupplierServicesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **provider_id** | **String** | The supplier id: the Scientist.com provider id as text. | [optional] |
| **ware_id** | **String** | The marketplace service listing id. | [optional] |
| **service_name** | **String** | The service name as the supplier listed it, as plain text. | [optional] |
| **ware_type** | **String** | The kind of marketplace listing. | [optional] |
| **discipline** | **String** | The top of the marketplace category tree. | [optional] |
| **category** | **String** | The marketplace category. | [optional] |
| **subcategory** | **String** | The marketplace subcategory. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SupplierServicesRow.new(
  provider_id: null,
  ware_id: null,
  service_name: null,
  ware_type: null,
  discipline: null,
  category: null,
  subcategory: null
)
```

