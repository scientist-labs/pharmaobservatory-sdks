# PharmaObservatory::SuppliersApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**suppliers_get**](SuppliersApi.md#suppliers_get) | **GET** /api/v1/suppliers/{id} | One supplier |
| [**suppliers_list**](SuppliersApi.md#suppliers_list) | **GET** /api/v1/suppliers | List suppliers |
| [**suppliers_supplier_services**](SuppliersApi.md#suppliers_supplier_services) | **GET** /api/v1/suppliers/{id}/supplier_services | Supplier services |
| [**suppliers_supplier_sites**](SuppliersApi.md#suppliers_supplier_sites) | **GET** /api/v1/suppliers/{id}/supplier_sites | Supplier sites |


## suppliers_get

> <SuppliersRecord> suppliers_get(id)

One supplier

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SuppliersApi.new
id = 990502 # Integer | 

begin
  # One supplier
  result = api_instance.suppliers_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SuppliersApi->suppliers_get: #{e}"
end
```

#### Using the suppliers_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SuppliersRecord>, Integer, Hash)> suppliers_get_with_http_info(id)

```ruby
begin
  # One supplier
  data, status_code, headers = api_instance.suppliers_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SuppliersRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SuppliersApi->suppliers_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |

### Return type

[**SuppliersRecord**](SuppliersRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## suppliers_list

> <SuppliersResponse> suppliers_list(opts)

List suppliers

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SuppliersApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List suppliers
  result = api_instance.suppliers_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SuppliersApi->suppliers_list: #{e}"
end
```

#### Using the suppliers_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SuppliersResponse>, Integer, Hash)> suppliers_list_with_http_info(opts)

```ruby
begin
  # List suppliers
  data, status_code, headers = api_instance.suppliers_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SuppliersResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SuppliersApi->suppliers_list_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**SuppliersResponse**](SuppliersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## suppliers_supplier_services

> <SupplierServicesResponse> suppliers_supplier_services(id, opts)

Supplier services

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SuppliersApi.new
id = 990502 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Supplier services
  result = api_instance.suppliers_supplier_services(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SuppliersApi->suppliers_supplier_services: #{e}"
end
```

#### Using the suppliers_supplier_services_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SupplierServicesResponse>, Integer, Hash)> suppliers_supplier_services_with_http_info(id, opts)

```ruby
begin
  # Supplier services
  data, status_code, headers = api_instance.suppliers_supplier_services_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SupplierServicesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SuppliersApi->suppliers_supplier_services_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**SupplierServicesResponse**](SupplierServicesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## suppliers_supplier_sites

> <SupplierSitesResponse> suppliers_supplier_sites(id, opts)

Supplier sites

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SuppliersApi.new
id = 990502 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Supplier sites
  result = api_instance.suppliers_supplier_sites(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SuppliersApi->suppliers_supplier_sites: #{e}"
end
```

#### Using the suppliers_supplier_sites_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SupplierSitesResponse>, Integer, Hash)> suppliers_supplier_sites_with_http_info(id, opts)

```ruby
begin
  # Supplier sites
  data, status_code, headers = api_instance.suppliers_supplier_sites_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SupplierSitesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SuppliersApi->suppliers_supplier_sites_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**SupplierSitesResponse**](SupplierSitesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

