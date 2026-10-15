# PharmaObservatory::ApiV1BaseApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_whoami_get**](ApiV1BaseApi.md#api_v1_whoami_get) | **GET** /api/v1/whoami | whoami |


## api_v1_whoami_get

> <ApiV1WhoamiGet200Response> api_v1_whoami_get

whoami

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1BaseApi.new

begin
  # whoami
  result = api_instance.api_v1_whoami_get
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1BaseApi->api_v1_whoami_get: #{e}"
end
```

#### Using the api_v1_whoami_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1WhoamiGet200Response>, Integer, Hash)> api_v1_whoami_get_with_http_info

```ruby
begin
  # whoami
  data, status_code, headers = api_instance.api_v1_whoami_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1WhoamiGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1BaseApi->api_v1_whoami_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**ApiV1WhoamiGet200Response**](ApiV1WhoamiGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

