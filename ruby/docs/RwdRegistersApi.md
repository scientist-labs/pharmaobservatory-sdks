# PharmaObservatory::RwdRegistersApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**rwd_registers_get**](RwdRegistersApi.md#rwd_registers_get) | **GET** /api/v1/rwd_registers/{id} | One real-world data register |
| [**rwd_registers_list**](RwdRegistersApi.md#rwd_registers_list) | **GET** /api/v1/rwd_registers | List real-world data registers |
| [**rwd_registers_rwd_coverage**](RwdRegistersApi.md#rwd_registers_rwd_coverage) | **GET** /api/v1/rwd_registers/{id}/rwd_coverage | HMA-EMA catalogue denominator chains |
| [**rwd_registers_rwd_withheld_fields**](RwdRegistersApi.md#rwd_registers_rwd_withheld_fields) | **GET** /api/v1/rwd_registers/{id}/rwd_withheld_fields | HMA-EMA catalogue columns withheld from publication |


## rwd_registers_get

> <RwdRegistersRecord> rwd_registers_get(id)

One real-world data register

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdRegistersApi.new
id = 'study' # String | 

begin
  # One real-world data register
  result = api_instance.rwd_registers_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdRegistersApi->rwd_registers_get: #{e}"
end
```

#### Using the rwd_registers_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdRegistersRecord>, Integer, Hash)> rwd_registers_get_with_http_info(id)

```ruby
begin
  # One real-world data register
  data, status_code, headers = api_instance.rwd_registers_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdRegistersRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdRegistersApi->rwd_registers_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**RwdRegistersRecord**](RwdRegistersRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_registers_list

> <RwdRegisterProvenanceResponse> rwd_registers_list(opts)

List real-world data registers

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdRegistersApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List real-world data registers
  result = api_instance.rwd_registers_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdRegistersApi->rwd_registers_list: #{e}"
end
```

#### Using the rwd_registers_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdRegisterProvenanceResponse>, Integer, Hash)> rwd_registers_list_with_http_info(opts)

```ruby
begin
  # List real-world data registers
  data, status_code, headers = api_instance.rwd_registers_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdRegisterProvenanceResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdRegistersApi->rwd_registers_list_with_http_info: #{e}"
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

[**RwdRegisterProvenanceResponse**](RwdRegisterProvenanceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_registers_rwd_coverage

> <RwdCoverageResponse> rwd_registers_rwd_coverage(id, opts)

HMA-EMA catalogue denominator chains

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdRegistersApi.new
id = 'study' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # HMA-EMA catalogue denominator chains
  result = api_instance.rwd_registers_rwd_coverage(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdRegistersApi->rwd_registers_rwd_coverage: #{e}"
end
```

#### Using the rwd_registers_rwd_coverage_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdCoverageResponse>, Integer, Hash)> rwd_registers_rwd_coverage_with_http_info(id, opts)

```ruby
begin
  # HMA-EMA catalogue denominator chains
  data, status_code, headers = api_instance.rwd_registers_rwd_coverage_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdCoverageResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdRegistersApi->rwd_registers_rwd_coverage_with_http_info: #{e}"
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

[**RwdCoverageResponse**](RwdCoverageResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_registers_rwd_withheld_fields

> <RwdWithheldFieldsResponse> rwd_registers_rwd_withheld_fields(id, opts)

HMA-EMA catalogue columns withheld from publication

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdRegistersApi.new
id = 'study' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # HMA-EMA catalogue columns withheld from publication
  result = api_instance.rwd_registers_rwd_withheld_fields(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdRegistersApi->rwd_registers_rwd_withheld_fields: #{e}"
end
```

#### Using the rwd_registers_rwd_withheld_fields_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdWithheldFieldsResponse>, Integer, Hash)> rwd_registers_rwd_withheld_fields_with_http_info(id, opts)

```ruby
begin
  # HMA-EMA catalogue columns withheld from publication
  data, status_code, headers = api_instance.rwd_registers_rwd_withheld_fields_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdWithheldFieldsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdRegistersApi->rwd_registers_rwd_withheld_fields_with_http_info: #{e}"
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

[**RwdWithheldFieldsResponse**](RwdWithheldFieldsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

