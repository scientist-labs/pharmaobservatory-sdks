# PharmaObservatory::CellLinesApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**cell_lines_get**](CellLinesApi.md#cell_lines_get) | **GET** /api/v1/cell_lines/{id} | One cell line |
| [**cell_lines_list**](CellLinesApi.md#cell_lines_list) | **GET** /api/v1/cell_lines | List cell lines |


## cell_lines_get

> <CellLinesRecord> cell_lines_get(id)

One cell line

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
  # One cell line
  result = api_instance.cell_lines_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_get: #{e}"
end
```

#### Using the cell_lines_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CellLinesRecord>, Integer, Hash)> cell_lines_get_with_http_info(id)

```ruby
begin
  # One cell line
  data, status_code, headers = api_instance.cell_lines_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CellLinesRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**CellLinesRecord**](CellLinesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## cell_lines_list

> <CellLinesResponse> cell_lines_list(opts)

List cell lines

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
  limit: 2, # Integer | 
  sort: 'cell_line_id' # String | 
}

begin
  # List cell lines
  result = api_instance.cell_lines_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_list: #{e}"
end
```

#### Using the cell_lines_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CellLinesResponse>, Integer, Hash)> cell_lines_list_with_http_info(opts)

```ruby
begin
  # List cell lines
  data, status_code, headers = api_instance.cell_lines_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CellLinesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CellLinesApi->cell_lines_list_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**CellLinesResponse**](CellLinesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

