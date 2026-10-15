# PharmaObservatory::ApiV1ConditionApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_conditions_condition_id_burden_vs_evidence_get**](ApiV1ConditionApi.md#api_v1_conditions_condition_id_burden_vs_evidence_get) | **GET** /api/v1/conditions/{condition_id}/burden_vs_evidence | burden_vs_evidence |
| [**api_v1_conditions_condition_id_evidence_freshness_get**](ApiV1ConditionApi.md#api_v1_conditions_condition_id_evidence_freshness_get) | **GET** /api/v1/conditions/{condition_id}/evidence_freshness | evidence_freshness |
| [**api_v1_conditions_condition_id_sponsor_landscape_get**](ApiV1ConditionApi.md#api_v1_conditions_condition_id_sponsor_landscape_get) | **GET** /api/v1/conditions/{condition_id}/sponsor_landscape | sponsor_landscape |


## api_v1_conditions_condition_id_burden_vs_evidence_get

> <ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response> api_v1_conditions_condition_id_burden_vs_evidence_get(condition_id)

burden_vs_evidence

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1ConditionApi.new
condition_id = 'D020521' # String | 

begin
  # burden_vs_evidence
  result = api_instance.api_v1_conditions_condition_id_burden_vs_evidence_get(condition_id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ConditionApi->api_v1_conditions_condition_id_burden_vs_evidence_get: #{e}"
end
```

#### Using the api_v1_conditions_condition_id_burden_vs_evidence_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response>, Integer, Hash)> api_v1_conditions_condition_id_burden_vs_evidence_get_with_http_info(condition_id)

```ruby
begin
  # burden_vs_evidence
  data, status_code, headers = api_instance.api_v1_conditions_condition_id_burden_vs_evidence_get_with_http_info(condition_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ConditionApi->api_v1_conditions_condition_id_burden_vs_evidence_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  |  |

### Return type

[**ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response**](ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_conditions_condition_id_evidence_freshness_get

> <ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response> api_v1_conditions_condition_id_evidence_freshness_get(condition_id)

evidence_freshness

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1ConditionApi.new
condition_id = 'D064419' # String | 

begin
  # evidence_freshness
  result = api_instance.api_v1_conditions_condition_id_evidence_freshness_get(condition_id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ConditionApi->api_v1_conditions_condition_id_evidence_freshness_get: #{e}"
end
```

#### Using the api_v1_conditions_condition_id_evidence_freshness_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response>, Integer, Hash)> api_v1_conditions_condition_id_evidence_freshness_get_with_http_info(condition_id)

```ruby
begin
  # evidence_freshness
  data, status_code, headers = api_instance.api_v1_conditions_condition_id_evidence_freshness_get_with_http_info(condition_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ConditionApi->api_v1_conditions_condition_id_evidence_freshness_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  |  |

### Return type

[**ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_conditions_condition_id_sponsor_landscape_get

> <ApiV1ConditionsConditionIdSponsorLandscapeGet200Response> api_v1_conditions_condition_id_sponsor_landscape_get(condition_id)

sponsor_landscape

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1ConditionApi.new
condition_id = 'D064419' # String | 

begin
  # sponsor_landscape
  result = api_instance.api_v1_conditions_condition_id_sponsor_landscape_get(condition_id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ConditionApi->api_v1_conditions_condition_id_sponsor_landscape_get: #{e}"
end
```

#### Using the api_v1_conditions_condition_id_sponsor_landscape_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1ConditionsConditionIdSponsorLandscapeGet200Response>, Integer, Hash)> api_v1_conditions_condition_id_sponsor_landscape_get_with_http_info(condition_id)

```ruby
begin
  # sponsor_landscape
  data, status_code, headers = api_instance.api_v1_conditions_condition_id_sponsor_landscape_get_with_http_info(condition_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1ConditionsConditionIdSponsorLandscapeGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ConditionApi->api_v1_conditions_condition_id_sponsor_landscape_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  |  |

### Return type

[**ApiV1ConditionsConditionIdSponsorLandscapeGet200Response**](ApiV1ConditionsConditionIdSponsorLandscapeGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

