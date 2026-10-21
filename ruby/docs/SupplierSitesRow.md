# PharmaObservatory::SupplierSitesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **site_id** | **String** | The supplier site id. | [optional] |
| **provider_id** | **String** | The supplier id: the Scientist.com provider id as text. | [optional] |
| **supplier_name** | **String** | The supplier&#39;s name as listed on the marketplace. | [optional] |
| **category** | **String** | The kind of site as the supplier recorded it (Headquarters, Laboratory, Warehouse, Other). Examples, not a closed set. | [optional] |
| **city** | **String** | City as recorded. | [optional] |
| **state** | **String** | State or province as recorded. | [optional] |
| **country** | **String** | Country as recorded. | [optional] |
| **latitude** | **Float** | Latitude of the site. | [optional] |
| **longitude** | **Float** | Longitude of the site. | [optional] |
| **company_type** | **String** | The supplier&#39;s company type, repeated on each of its sites. | [optional] |
| **service_areas** | **Array&lt;String&gt;** | The supplier&#39;s service areas, repeated on each of its sites. Filter with &#39;Biology&#39; &#x3D; ANY(service_areas). | [optional] |
| **hq_country** | **String** | ISO 3166-1 alpha-2 country code of the supplier&#39;s headquarters point, repeated on each of its sites. | [optional] |
| **hq_region** | **String** | The broad world region of the supplier&#39;s headquarters country, repeated on each of its sites; other when unknown. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SupplierSitesRow.new(
  site_id: null,
  provider_id: null,
  supplier_name: null,
  category: null,
  city: null,
  state: null,
  country: null,
  latitude: null,
  longitude: null,
  company_type: null,
  service_areas: null,
  hq_country: null,
  hq_region: null
)
```

