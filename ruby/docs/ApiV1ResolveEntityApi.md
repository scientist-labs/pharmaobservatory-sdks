# PharmaObservatory::ApiV1ResolveEntityApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_resolve_entity_get**](ApiV1ResolveEntityApi.md#api_v1_resolve_entity_get) | **GET** /api/v1/resolve_entity | Resolve a free-text or coded entity mention to a canonical Pharma Observatory ID |


## api_v1_resolve_entity_get

> <ApiV1ResolveEntityGet200Response> api_v1_resolve_entity_get(opts)

Resolve a free-text or coded entity mention to a canonical Pharma Observatory ID

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1ResolveEntityApi.new
opts = {
  query: 'Anchorloss Biotherapeutics', # String | 
  raw_key: 'biomarin', # String | 
  source_type: 'aact_sponsor', # String | 
  spine: 'institution' # String | 
}

begin
  # Resolve a free-text or coded entity mention to a canonical Pharma Observatory ID
  result = api_instance.api_v1_resolve_entity_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ResolveEntityApi->api_v1_resolve_entity_get: #{e}"
end
```

#### Using the api_v1_resolve_entity_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1ResolveEntityGet200Response>, Integer, Hash)> api_v1_resolve_entity_get_with_http_info(opts)

```ruby
begin
  # Resolve a free-text or coded entity mention to a canonical Pharma Observatory ID
  data, status_code, headers = api_instance.api_v1_resolve_entity_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1ResolveEntityGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ResolveEntityApi->api_v1_resolve_entity_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **query** | **String** |  | [optional] |
| **raw_key** | **String** |  | [optional] |
| **source_type** | **String** |  | [optional] |
| **spine** | **String** |  | [optional] |

### Return type

[**ApiV1ResolveEntityGet200Response**](ApiV1ResolveEntityGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

