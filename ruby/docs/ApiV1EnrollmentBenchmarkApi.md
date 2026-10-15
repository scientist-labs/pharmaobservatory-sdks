# PharmaObservatory::ApiV1EnrollmentBenchmarkApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_enrollment_benchmark_get**](ApiV1EnrollmentBenchmarkApi.md#api_v1_enrollment_benchmark_get) | **GET** /api/v1/enrollment_benchmark | show |


## api_v1_enrollment_benchmark_get

> <ApiV1EnrollmentBenchmarkGet200Response> api_v1_enrollment_benchmark_get(opts)

show

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1EnrollmentBenchmarkApi.new
opts = {
  condition_id: 'D001523', # String | 
  phase: 'PHASE2' # String | 
}

begin
  # show
  result = api_instance.api_v1_enrollment_benchmark_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1EnrollmentBenchmarkApi->api_v1_enrollment_benchmark_get: #{e}"
end
```

#### Using the api_v1_enrollment_benchmark_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1EnrollmentBenchmarkGet200Response>, Integer, Hash)> api_v1_enrollment_benchmark_get_with_http_info(opts)

```ruby
begin
  # show
  data, status_code, headers = api_instance.api_v1_enrollment_benchmark_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1EnrollmentBenchmarkGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1EnrollmentBenchmarkApi->api_v1_enrollment_benchmark_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  | [optional] |
| **phase** | **String** |  | [optional] |

### Return type

[**ApiV1EnrollmentBenchmarkGet200Response**](ApiV1EnrollmentBenchmarkGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

