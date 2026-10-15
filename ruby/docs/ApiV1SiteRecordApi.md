# PharmaObservatory::ApiV1SiteRecordApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_site_record_get**](ApiV1SiteRecordApi.md#api_v1_site_record_get) | **GET** /api/v1/site_record | Look up a clinical trial site&#39;s full record, including hosted trials and named investigators |


## api_v1_site_record_get

> <ApiV1SiteRecordGet200Response> api_v1_site_record_get(opts)

Look up a clinical trial site's full record, including hosted trials and named investigators

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1SiteRecordApi.new
opts = {
  investigator_id: 'aact:e7c2d751838e731e6898645b06a70eb1', # String | 
  limit: 1, # Integer | 
  org_id: 'edgar:0001060955', # String | 
  site_id: '87a040a71a6f82289add8872820314eb' # String | 
}

begin
  # Look up a clinical trial site's full record, including hosted trials and named investigators
  result = api_instance.api_v1_site_record_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1SiteRecordApi->api_v1_site_record_get: #{e}"
end
```

#### Using the api_v1_site_record_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1SiteRecordGet200Response>, Integer, Hash)> api_v1_site_record_get_with_http_info(opts)

```ruby
begin
  # Look up a clinical trial site's full record, including hosted trials and named investigators
  data, status_code, headers = api_instance.api_v1_site_record_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1SiteRecordGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1SiteRecordApi->api_v1_site_record_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investigator_id** | **String** |  | [optional] |
| **limit** | **Integer** |  | [optional] |
| **org_id** | **String** |  | [optional] |
| **site_id** | **String** |  | [optional] |

### Return type

[**ApiV1SiteRecordGet200Response**](ApiV1SiteRecordGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

