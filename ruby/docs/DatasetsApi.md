# PharmaObservatory::DatasetsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**datasets_get**](DatasetsApi.md#datasets_get) | **GET** /api/v1/datasets/{name} | Describe a dataset |
| [**datasets_list**](DatasetsApi.md#datasets_list) | **GET** /api/v1/datasets | List the datasets |


## datasets_get

> Hash&lt;String, Object&gt; datasets_get(name)

Describe a dataset

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DatasetsApi.new
name = 'sponsor_trials' # String | 

begin
  # Describe a dataset
  result = api_instance.datasets_get(name)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DatasetsApi->datasets_get: #{e}"
end
```

#### Using the datasets_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Hash&lt;String, Object&gt;, Integer, Hash)> datasets_get_with_http_info(name)

```ruby
begin
  # Describe a dataset
  data, status_code, headers = api_instance.datasets_get_with_http_info(name)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Hash&lt;String, Object&gt;
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DatasetsApi->datasets_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **name** | **String** |  |  |

### Return type

**Hash&lt;String, Object&gt;**

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## datasets_list

> <DatasetSummaryList> datasets_list

List the datasets

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DatasetsApi.new

begin
  # List the datasets
  result = api_instance.datasets_list
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DatasetsApi->datasets_list: #{e}"
end
```

#### Using the datasets_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DatasetSummaryList>, Integer, Hash)> datasets_list_with_http_info

```ruby
begin
  # List the datasets
  data, status_code, headers = api_instance.datasets_list_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DatasetSummaryList>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DatasetsApi->datasets_list_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**DatasetSummaryList**](DatasetSummaryList.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

