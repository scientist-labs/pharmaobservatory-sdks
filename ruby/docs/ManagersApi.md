# PharmaObservatory::ManagersApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**managers_get**](ManagersApi.md#managers_get) | **GET** /api/v1/managers/{id} | One form 13f manager |
| [**managers_institutional_holdings**](ManagersApi.md#managers_institutional_holdings) | **GET** /api/v1/managers/{id}/institutional_holdings | Institutional holdings (SEC Form 13F) |
| [**managers_list**](ManagersApi.md#managers_list) | **GET** /api/v1/managers | List form 13f managers |


## managers_get

> <ManagersRecord> managers_get(id)

One form 13f manager

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ManagersApi.new
id = 'edgar_cik:0002900005' # String | 

begin
  # One form 13f manager
  result = api_instance.managers_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ManagersApi->managers_get: #{e}"
end
```

#### Using the managers_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ManagersRecord>, Integer, Hash)> managers_get_with_http_info(id)

```ruby
begin
  # One form 13f manager
  data, status_code, headers = api_instance.managers_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ManagersRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ManagersApi->managers_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**ManagersRecord**](ManagersRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## managers_institutional_holdings

> <InstitutionalHoldingsResponse> managers_institutional_holdings(id, opts)

Institutional holdings (SEC Form 13F)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ManagersApi.new
id = 'edgar_cik:0002900005' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Institutional holdings (SEC Form 13F)
  result = api_instance.managers_institutional_holdings(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ManagersApi->managers_institutional_holdings: #{e}"
end
```

#### Using the managers_institutional_holdings_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InstitutionalHoldingsResponse>, Integer, Hash)> managers_institutional_holdings_with_http_info(id, opts)

```ruby
begin
  # Institutional holdings (SEC Form 13F)
  data, status_code, headers = api_instance.managers_institutional_holdings_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InstitutionalHoldingsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ManagersApi->managers_institutional_holdings_with_http_info: #{e}"
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

[**InstitutionalHoldingsResponse**](InstitutionalHoldingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## managers_list

> <ManagersResponse> managers_list(opts)

List form 13f managers

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ManagersApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List form 13f managers
  result = api_instance.managers_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ManagersApi->managers_list: #{e}"
end
```

#### Using the managers_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ManagersResponse>, Integer, Hash)> managers_list_with_http_info(opts)

```ruby
begin
  # List form 13f managers
  data, status_code, headers = api_instance.managers_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ManagersResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ManagersApi->managers_list_with_http_info: #{e}"
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

[**ManagersResponse**](ManagersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

