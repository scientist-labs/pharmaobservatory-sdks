# PharmaObservatory::RwdInstitutionsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**rwd_institutions_get**](RwdInstitutionsApi.md#rwd_institutions_get) | **GET** /api/v1/rwd_institutions/{id} | One real-world data register institution |
| [**rwd_institutions_list**](RwdInstitutionsApi.md#rwd_institutions_list) | **GET** /api/v1/rwd_institutions | List real-world data register institutions |


## rwd_institutions_get

> <RwdInstitutionsRecord> rwd_institutions_get(id)

One real-world data register institution

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdInstitutionsApi.new
id = 900000005 # Integer | 

begin
  # One real-world data register institution
  result = api_instance.rwd_institutions_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdInstitutionsApi->rwd_institutions_get: #{e}"
end
```

#### Using the rwd_institutions_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdInstitutionsRecord>, Integer, Hash)> rwd_institutions_get_with_http_info(id)

```ruby
begin
  # One real-world data register institution
  data, status_code, headers = api_instance.rwd_institutions_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdInstitutionsRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdInstitutionsApi->rwd_institutions_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |

### Return type

[**RwdInstitutionsRecord**](RwdInstitutionsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_institutions_list

> <RwdInstitutionsResponse> rwd_institutions_list(opts)

List real-world data register institutions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdInstitutionsApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List real-world data register institutions
  result = api_instance.rwd_institutions_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdInstitutionsApi->rwd_institutions_list: #{e}"
end
```

#### Using the rwd_institutions_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdInstitutionsResponse>, Integer, Hash)> rwd_institutions_list_with_http_info(opts)

```ruby
begin
  # List real-world data register institutions
  data, status_code, headers = api_instance.rwd_institutions_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdInstitutionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdInstitutionsApi->rwd_institutions_list_with_http_info: #{e}"
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

[**RwdInstitutionsResponse**](RwdInstitutionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

