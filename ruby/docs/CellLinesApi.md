# PharmaObservatory::CellLinesApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**cell_lines_get**](CellLinesApi.md#cell_lines_get) | **GET** /api/v1/cell-lines/{id} | One cell-lines record |
| [**cell_lines_identity**](CellLinesApi.md#cell_lines_identity) | **GET** /api/v1/cell-lines/{id}/identity | Identity and availability |
| [**cell_lines_list**](CellLinesApi.md#cell_lines_list) | **GET** /api/v1/cell-lines | Cell line identity safety |
| [**cell_lines_safety_resolution**](CellLinesApi.md#cell_lines_safety_resolution) | **GET** /api/v1/cell-lines/{id}/safety_resolution | Safety resolution |
| [**cell_lines_status_counts**](CellLinesApi.md#cell_lines_status_counts) | **GET** /api/v1/cell-lines/status_counts | Cell line safety summary |


## cell_lines_get

> <CellLinesGet200Response> cell_lines_get(id)

One cell-lines record

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CellLinesApi.new
id = 'SCI-CONF' # String | 

begin
  # One cell-lines record
  result = api_instance.cell_lines_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_get: #{e}"
end
```

#### Using the cell_lines_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CellLinesGet200Response>, Integer, Hash)> cell_lines_get_with_http_info(id)

```ruby
begin
  # One cell-lines record
  data, status_code, headers = api_instance.cell_lines_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CellLinesGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**CellLinesGet200Response**](CellLinesGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## cell_lines_identity

> <CellLinesIdentity200Response> cell_lines_identity(id)

Identity and availability

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CellLinesApi.new
id = 'SCI-CONF' # String | 

begin
  # Identity and availability
  result = api_instance.cell_lines_identity(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_identity: #{e}"
end
```

#### Using the cell_lines_identity_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CellLinesIdentity200Response>, Integer, Hash)> cell_lines_identity_with_http_info(id)

```ruby
begin
  # Identity and availability
  data, status_code, headers = api_instance.cell_lines_identity_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CellLinesIdentity200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_identity_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**CellLinesIdentity200Response**](CellLinesIdentity200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## cell_lines_list

> <CellLinesList200Response> cell_lines_list(opts)

Cell line identity safety

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CellLinesApi.new
opts = {
  cursor: 'eyJwIjoyLCJsIjoyLCJmIjoic3BlYy1maW5nZXJwcmludCJ9', # String | 
  facets: true, # Boolean | 
  limit: 1, # Integer | 
  safety_status: 'confirmed_problematic' # String | 
}

begin
  # Cell line identity safety
  result = api_instance.cell_lines_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_list: #{e}"
end
```

#### Using the cell_lines_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CellLinesList200Response>, Integer, Hash)> cell_lines_list_with_http_info(opts)

```ruby
begin
  # Cell line identity safety
  data, status_code, headers = api_instance.cell_lines_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CellLinesList200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_list_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cursor** | **String** |  | [optional] |
| **facets** | **Boolean** |  | [optional] |
| **limit** | **Integer** |  | [optional] |
| **safety_status** | **String** |  | [optional] |

### Return type

[**CellLinesList200Response**](CellLinesList200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## cell_lines_safety_resolution

> <CellLinesSafetyResolution200Response> cell_lines_safety_resolution(id)

Safety resolution

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CellLinesApi.new
id = 'SCI-CONF' # String | 

begin
  # Safety resolution
  result = api_instance.cell_lines_safety_resolution(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_safety_resolution: #{e}"
end
```

#### Using the cell_lines_safety_resolution_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CellLinesSafetyResolution200Response>, Integer, Hash)> cell_lines_safety_resolution_with_http_info(id)

```ruby
begin
  # Safety resolution
  data, status_code, headers = api_instance.cell_lines_safety_resolution_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CellLinesSafetyResolution200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_safety_resolution_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**CellLinesSafetyResolution200Response**](CellLinesSafetyResolution200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## cell_lines_status_counts

> <CellLinesStatusCounts200Response> cell_lines_status_counts(opts)

Cell line safety summary

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CellLinesApi.new
opts = {
  species: 'Mouse' # String | 
}

begin
  # Cell line safety summary
  result = api_instance.cell_lines_status_counts(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_status_counts: #{e}"
end
```

#### Using the cell_lines_status_counts_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CellLinesStatusCounts200Response>, Integer, Hash)> cell_lines_status_counts_with_http_info(opts)

```ruby
begin
  # Cell line safety summary
  data, status_code, headers = api_instance.cell_lines_status_counts_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CellLinesStatusCounts200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_status_counts_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **species** | **String** |  | [optional] |

### Return type

[**CellLinesStatusCounts200Response**](CellLinesStatusCounts200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

