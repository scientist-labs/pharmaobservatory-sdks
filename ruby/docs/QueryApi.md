# PharmaObservatory::QueryApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**query_result**](QueryApi.md#query_result) | **GET** /api/v1/results/{result_id} | Page a stored result |
| [**query_run**](QueryApi.md#query_run) | **POST** /api/v1/query | Run a SQL query |


## query_result

> <ResultPage> query_result(result_id, opts)

Page a stored result

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::QueryApi.new
result_id = '5d3f6a1e-0c1b-4f7e-9a52-2b8c1d4e6f70' # String | 
opts = {
  cursor: 'eyJvIjoyMDB9' # String | 
}

begin
  # Page a stored result
  result = api_instance.query_result(result_id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling QueryApi->query_result: #{e}"
end
```

#### Using the query_result_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ResultPage>, Integer, Hash)> query_result_with_http_info(result_id, opts)

```ruby
begin
  # Page a stored result
  data, status_code, headers = api_instance.query_result_with_http_info(result_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ResultPage>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling QueryApi->query_result_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **result_id** | **String** |  |  |
| **cursor** | **String** |  | [optional] |

### Return type

[**ResultPage**](ResultPage.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## query_run

> <QueryResponse> query_run(opts)

Run a SQL query

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::QueryApi.new
opts = {
  query_run_request: PharmaObservatory::QueryRunRequest.new({sql: 'sql_example'}) # QueryRunRequest | 
}

begin
  # Run a SQL query
  result = api_instance.query_run(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling QueryApi->query_run: #{e}"
end
```

#### Using the query_run_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<QueryResponse>, Integer, Hash)> query_run_with_http_info(opts)

```ruby
begin
  # Run a SQL query
  data, status_code, headers = api_instance.query_run_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <QueryResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling QueryApi->query_run_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **query_run_request** | [**QueryRunRequest**](QueryRunRequest.md) |  | [optional] |

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

