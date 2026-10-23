# PharmaObservatory::RwdConductorsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**rwd_conductors_get**](RwdConductorsApi.md#rwd_conductors_get) | **GET** /api/v1/rwd_conductors/{id} | One real-world data study conductor |
| [**rwd_conductors_list**](RwdConductorsApi.md#rwd_conductors_list) | **GET** /api/v1/rwd_conductors | List real-world data study conductors |
| [**rwd_conductors_rwd_conductor_delivery**](RwdConductorsApi.md#rwd_conductors_rwd_conductor_delivery) | **GET** /api/v1/rwd_conductors/{id}/rwd_conductor_delivery | Real-world data study conductor delivery by mandate (HMA-EMA catalogue) |
| [**rwd_conductors_rwd_conductor_footprint**](RwdConductorsApi.md#rwd_conductors_rwd_conductor_footprint) | **GET** /api/v1/rwd_conductors/{id}/rwd_conductor_footprint | Real-world data study conductor footprint (HMA-EMA catalogue) |
| [**rwd_conductors_rwd_conductor_studies**](RwdConductorsApi.md#rwd_conductors_rwd_conductor_studies) | **GET** /api/v1/rwd_conductors/{id}/rwd_conductor_studies | Real-world data study conductor studies (HMA-EMA catalogue) |


## rwd_conductors_get

> <RwdConductorsRecord> rwd_conductors_get(id)

One real-world data study conductor

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdConductorsApi.new
id = 'seedmark-consulting' # String | 

begin
  # One real-world data study conductor
  result = api_instance.rwd_conductors_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdConductorsApi->rwd_conductors_get: #{e}"
end
```

#### Using the rwd_conductors_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdConductorsRecord>, Integer, Hash)> rwd_conductors_get_with_http_info(id)

```ruby
begin
  # One real-world data study conductor
  data, status_code, headers = api_instance.rwd_conductors_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdConductorsRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdConductorsApi->rwd_conductors_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**RwdConductorsRecord**](RwdConductorsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_conductors_list

> <RwdConductorsResponse> rwd_conductors_list(opts)

List real-world data study conductors

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdConductorsApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List real-world data study conductors
  result = api_instance.rwd_conductors_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdConductorsApi->rwd_conductors_list: #{e}"
end
```

#### Using the rwd_conductors_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdConductorsResponse>, Integer, Hash)> rwd_conductors_list_with_http_info(opts)

```ruby
begin
  # List real-world data study conductors
  data, status_code, headers = api_instance.rwd_conductors_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdConductorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdConductorsApi->rwd_conductors_list_with_http_info: #{e}"
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

[**RwdConductorsResponse**](RwdConductorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_conductors_rwd_conductor_delivery

> <RwdConductorDeliveryResponse> rwd_conductors_rwd_conductor_delivery(id, opts)

Real-world data study conductor delivery by mandate (HMA-EMA catalogue)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdConductorsApi.new
id = 'seedmark-consulting' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Real-world data study conductor delivery by mandate (HMA-EMA catalogue)
  result = api_instance.rwd_conductors_rwd_conductor_delivery(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdConductorsApi->rwd_conductors_rwd_conductor_delivery: #{e}"
end
```

#### Using the rwd_conductors_rwd_conductor_delivery_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdConductorDeliveryResponse>, Integer, Hash)> rwd_conductors_rwd_conductor_delivery_with_http_info(id, opts)

```ruby
begin
  # Real-world data study conductor delivery by mandate (HMA-EMA catalogue)
  data, status_code, headers = api_instance.rwd_conductors_rwd_conductor_delivery_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdConductorDeliveryResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdConductorsApi->rwd_conductors_rwd_conductor_delivery_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RwdConductorDeliveryResponse**](RwdConductorDeliveryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_conductors_rwd_conductor_footprint

> <RwdConductorFootprintResponse> rwd_conductors_rwd_conductor_footprint(id, opts)

Real-world data study conductor footprint (HMA-EMA catalogue)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdConductorsApi.new
id = 'seedmark-consulting' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Real-world data study conductor footprint (HMA-EMA catalogue)
  result = api_instance.rwd_conductors_rwd_conductor_footprint(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdConductorsApi->rwd_conductors_rwd_conductor_footprint: #{e}"
end
```

#### Using the rwd_conductors_rwd_conductor_footprint_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdConductorFootprintResponse>, Integer, Hash)> rwd_conductors_rwd_conductor_footprint_with_http_info(id, opts)

```ruby
begin
  # Real-world data study conductor footprint (HMA-EMA catalogue)
  data, status_code, headers = api_instance.rwd_conductors_rwd_conductor_footprint_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdConductorFootprintResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdConductorsApi->rwd_conductors_rwd_conductor_footprint_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RwdConductorFootprintResponse**](RwdConductorFootprintResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_conductors_rwd_conductor_studies

> <RwdConductorStudiesResponse> rwd_conductors_rwd_conductor_studies(id, opts)

Real-world data study conductor studies (HMA-EMA catalogue)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdConductorsApi.new
id = 'seedmark-consulting' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Real-world data study conductor studies (HMA-EMA catalogue)
  result = api_instance.rwd_conductors_rwd_conductor_studies(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdConductorsApi->rwd_conductors_rwd_conductor_studies: #{e}"
end
```

#### Using the rwd_conductors_rwd_conductor_studies_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdConductorStudiesResponse>, Integer, Hash)> rwd_conductors_rwd_conductor_studies_with_http_info(id, opts)

```ruby
begin
  # Real-world data study conductor studies (HMA-EMA catalogue)
  data, status_code, headers = api_instance.rwd_conductors_rwd_conductor_studies_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdConductorStudiesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdConductorsApi->rwd_conductors_rwd_conductor_studies_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RwdConductorStudiesResponse**](RwdConductorStudiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

