# PharmaObservatory::ApiV1CandidateSiteApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_candidate_sites_get**](ApiV1CandidateSiteApi.md#api_v1_candidate_sites_get) | **GET** /api/v1/candidate_sites | Rank candidate clinical trial sites for a condition by composite capacity, activity, and competition scores |


## api_v1_candidate_sites_get

> <ApiV1CandidateSitesGet200Response> api_v1_candidate_sites_get(opts)

Rank candidate clinical trial sites for a condition by composite capacity, activity, and competition scores

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1CandidateSiteApi.new
opts = {
  condition_id: 'D001523', # String | 
  country: 'United States', # String | 
  exclude_saturated: true, # Boolean | 
  limit: 500 # Integer | 
}

begin
  # Rank candidate clinical trial sites for a condition by composite capacity, activity, and competition scores
  result = api_instance.api_v1_candidate_sites_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1CandidateSiteApi->api_v1_candidate_sites_get: #{e}"
end
```

#### Using the api_v1_candidate_sites_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1CandidateSitesGet200Response>, Integer, Hash)> api_v1_candidate_sites_get_with_http_info(opts)

```ruby
begin
  # Rank candidate clinical trial sites for a condition by composite capacity, activity, and competition scores
  data, status_code, headers = api_instance.api_v1_candidate_sites_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1CandidateSitesGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1CandidateSiteApi->api_v1_candidate_sites_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  | [optional] |
| **country** | **String** |  | [optional] |
| **exclude_saturated** | **Boolean** |  | [optional] |
| **limit** | **Integer** |  | [optional] |

### Return type

[**ApiV1CandidateSitesGet200Response**](ApiV1CandidateSitesGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

