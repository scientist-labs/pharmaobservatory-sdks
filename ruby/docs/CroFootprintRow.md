# PharmaObservatory::CroFootprintRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **vendor_company_id** | **String** | Company level: the contract research organization&#39;s company. | [optional] |
| **vendor_company_name** | **String** | Display name of the contract research organization&#39;s company. | [optional] |
| **vendor_group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **vendor_group_name** | **String** | Display name of the company&#39;s group. | [optional] |
| **country** | **String** | The country name as the registry&#39;s facility roster spells it. | [optional] |
| **n_trials** | **Integer** | Distinct trials the company delivered with a facility in the country. | [optional] |
| **n_sites** | **Integer** | Distinct sites (facility keys) in the country across those trials. | [optional] |
| **latitude** | **Float** | Latitude of the country&#39;s approximate centroid. Null for a country with no centroid. | [optional] |
| **longitude** | **Float** | Longitude of the country&#39;s approximate centroid. Null for a country with no centroid. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CroFootprintRow.new(
  vendor_company_id: null,
  vendor_company_name: null,
  vendor_group_id: null,
  vendor_group_name: null,
  country: null,
  n_trials: null,
  n_sites: null,
  latitude: null,
  longitude: null
)
```

