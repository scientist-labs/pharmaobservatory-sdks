# PharmaObservatory::SubscriptionsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**changes_list**](SubscriptionsApi.md#changes_list) | **GET** /api/v1/changes | Changes since a cursor |
| [**subscriptions_create**](SubscriptionsApi.md#subscriptions_create) | **POST** /api/v1/subscriptions | Follow an entity or a dataset query |
| [**subscriptions_delete**](SubscriptionsApi.md#subscriptions_delete) | **DELETE** /api/v1/subscriptions/{id} | Stop following |
| [**subscriptions_list**](SubscriptionsApi.md#subscriptions_list) | **GET** /api/v1/subscriptions | List your subscriptions |


## changes_list

> <ChangesList200Response> changes_list(opts)

Changes since a cursor

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SubscriptionsApi.new
opts = {
  cursor: 'eyJ2IjoxLCJvIjoiMjAyNi0xMS0wMlQwNjowMjowMC4wMDAwMDBaIiwiaSI6OTAwMn0', # String | 
  limit: 2 # Integer | 
}

begin
  # Changes since a cursor
  result = api_instance.changes_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SubscriptionsApi->changes_list: #{e}"
end
```

#### Using the changes_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ChangesList200Response>, Integer, Hash)> changes_list_with_http_info(opts)

```ruby
begin
  # Changes since a cursor
  data, status_code, headers = api_instance.changes_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ChangesList200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SubscriptionsApi->changes_list_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cursor** | **String** |  | [optional] |
| **limit** | **Integer** |  | [optional] |

### Return type

[**ChangesList200Response**](ChangesList200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## subscriptions_create

> <SubscriptionsCreate201Response> subscriptions_create(opts)

Follow an entity or a dataset query

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SubscriptionsApi.new
opts = {
  subscriptions_create_request: PharmaObservatory::SubscriptionsCreateRequest.new({object_type: 'object_type_example', entity_id: 'entity_id_example'}) # SubscriptionsCreateRequest | 
}

begin
  # Follow an entity or a dataset query
  result = api_instance.subscriptions_create(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SubscriptionsApi->subscriptions_create: #{e}"
end
```

#### Using the subscriptions_create_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SubscriptionsCreate201Response>, Integer, Hash)> subscriptions_create_with_http_info(opts)

```ruby
begin
  # Follow an entity or a dataset query
  data, status_code, headers = api_instance.subscriptions_create_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SubscriptionsCreate201Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SubscriptionsApi->subscriptions_create_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **subscriptions_create_request** | [**SubscriptionsCreateRequest**](SubscriptionsCreateRequest.md) |  | [optional] |

### Return type

[**SubscriptionsCreate201Response**](SubscriptionsCreate201Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## subscriptions_delete

> <SubscriptionsDelete200Response> subscriptions_delete(id)

Stop following

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SubscriptionsApi.new
id = 4242 # Integer | 

begin
  # Stop following
  result = api_instance.subscriptions_delete(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SubscriptionsApi->subscriptions_delete: #{e}"
end
```

#### Using the subscriptions_delete_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SubscriptionsDelete200Response>, Integer, Hash)> subscriptions_delete_with_http_info(id)

```ruby
begin
  # Stop following
  data, status_code, headers = api_instance.subscriptions_delete_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SubscriptionsDelete200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SubscriptionsApi->subscriptions_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |

### Return type

[**SubscriptionsDelete200Response**](SubscriptionsDelete200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## subscriptions_list

> <SubscriptionsList200Response> subscriptions_list

List your subscriptions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SubscriptionsApi.new

begin
  # List your subscriptions
  result = api_instance.subscriptions_list
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SubscriptionsApi->subscriptions_list: #{e}"
end
```

#### Using the subscriptions_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SubscriptionsList200Response>, Integer, Hash)> subscriptions_list_with_http_info

```ruby
begin
  # List your subscriptions
  data, status_code, headers = api_instance.subscriptions_list_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SubscriptionsList200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SubscriptionsApi->subscriptions_list_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**SubscriptionsList200Response**](SubscriptionsList200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

