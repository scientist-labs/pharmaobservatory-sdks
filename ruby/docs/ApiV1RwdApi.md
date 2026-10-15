# PharmaObservatory::ApiV1RwdApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_rwd_conductor_track_record_get**](ApiV1RwdApi.md#api_v1_rwd_conductor_track_record_get) | **GET** /api/v1/rwd_conductor_track_record | conductor_track_record |
| [**api_v1_rwd_study_screen_get**](ApiV1RwdApi.md#api_v1_rwd_study_screen_get) | **GET** /api/v1/rwd_study_screen | study_screen |


## api_v1_rwd_conductor_track_record_get

> <ApiV1RwdConductorTrackRecordGet200Response> api_v1_rwd_conductor_track_record_get(opts)

conductor_track_record

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1RwdApi.new
opts = {
  canonical_org_id: 'ror:rareres01', # String | 
  hma_ema_id: 900000005, # Integer | 
  institution_name: 'Seedmark Consulting' # String | 
}

begin
  # conductor_track_record
  result = api_instance.api_v1_rwd_conductor_track_record_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1RwdApi->api_v1_rwd_conductor_track_record_get: #{e}"
end
```

#### Using the api_v1_rwd_conductor_track_record_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1RwdConductorTrackRecordGet200Response>, Integer, Hash)> api_v1_rwd_conductor_track_record_get_with_http_info(opts)

```ruby
begin
  # conductor_track_record
  data, status_code, headers = api_instance.api_v1_rwd_conductor_track_record_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1RwdConductorTrackRecordGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1RwdApi->api_v1_rwd_conductor_track_record_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **canonical_org_id** | **String** |  | [optional] |
| **hma_ema_id** | **Integer** |  | [optional] |
| **institution_name** | **String** |  | [optional] |

### Return type

[**ApiV1RwdConductorTrackRecordGet200Response**](ApiV1RwdConductorTrackRecordGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_rwd_study_screen_get

> <ApiV1RwdStudyScreenGet200Response> api_v1_rwd_study_screen_get(opts)

study_screen

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1RwdApi.new
opts = {
  condition: 'Cataract', # String | 
  mandate: 'mandated' # String | 
}

begin
  # study_screen
  result = api_instance.api_v1_rwd_study_screen_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1RwdApi->api_v1_rwd_study_screen_get: #{e}"
end
```

#### Using the api_v1_rwd_study_screen_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1RwdStudyScreenGet200Response>, Integer, Hash)> api_v1_rwd_study_screen_get_with_http_info(opts)

```ruby
begin
  # study_screen
  data, status_code, headers = api_instance.api_v1_rwd_study_screen_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1RwdStudyScreenGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1RwdApi->api_v1_rwd_study_screen_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition** | **String** |  | [optional] |
| **mandate** | **String** |  | [optional] |

### Return type

[**ApiV1RwdStudyScreenGet200Response**](ApiV1RwdStudyScreenGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

