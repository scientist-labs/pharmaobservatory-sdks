# PharmaObservatory::ApiV1TrialApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_trials_discontinuation_reasons_get**](ApiV1TrialApi.md#api_v1_trials_discontinuation_reasons_get) | **GET** /api/v1/trials/discontinuation-reasons | discontinuation_reasons |
| [**api_v1_trials_eligibility_generalizability_scan_get**](ApiV1TrialApi.md#api_v1_trials_eligibility_generalizability_scan_get) | **GET** /api/v1/trials/eligibility-generalizability-scan | eligibility_generalizability_scan |
| [**api_v1_trials_history_and_prediction_get**](ApiV1TrialApi.md#api_v1_trials_history_and_prediction_get) | **GET** /api/v1/trials/history-and-prediction | history_and_prediction |
| [**api_v1_trials_records_get**](ApiV1TrialApi.md#api_v1_trials_records_get) | **GET** /api/v1/trials/records | records |
| [**api_v1_trials_reporting_completeness_get**](ApiV1TrialApi.md#api_v1_trials_reporting_completeness_get) | **GET** /api/v1/trials/reporting-completeness | reporting_completeness |


## api_v1_trials_discontinuation_reasons_get

> <ApiV1TrialsDiscontinuationReasonsGet200Response> api_v1_trials_discontinuation_reasons_get(opts)

discontinuation_reasons

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1TrialApi.new
opts = {
  condition_id: 'D064419' # String | 
}

begin
  # discontinuation_reasons
  result = api_instance.api_v1_trials_discontinuation_reasons_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1TrialApi->api_v1_trials_discontinuation_reasons_get: #{e}"
end
```

#### Using the api_v1_trials_discontinuation_reasons_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1TrialsDiscontinuationReasonsGet200Response>, Integer, Hash)> api_v1_trials_discontinuation_reasons_get_with_http_info(opts)

```ruby
begin
  # discontinuation_reasons
  data, status_code, headers = api_instance.api_v1_trials_discontinuation_reasons_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1TrialsDiscontinuationReasonsGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1TrialApi->api_v1_trials_discontinuation_reasons_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  | [optional] |

### Return type

[**ApiV1TrialsDiscontinuationReasonsGet200Response**](ApiV1TrialsDiscontinuationReasonsGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_trials_eligibility_generalizability_scan_get

> <ApiV1TrialsEligibilityGeneralizabilityScanGet200Response> api_v1_trials_eligibility_generalizability_scan_get(opts)

eligibility_generalizability_scan

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1TrialApi.new
opts = {
  condition_id: 'D064419', # String | 
  limit: 5, # Integer | 
  nct_ids: ['inner_example'] # Array<String> | 
}

begin
  # eligibility_generalizability_scan
  result = api_instance.api_v1_trials_eligibility_generalizability_scan_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1TrialApi->api_v1_trials_eligibility_generalizability_scan_get: #{e}"
end
```

#### Using the api_v1_trials_eligibility_generalizability_scan_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1TrialsEligibilityGeneralizabilityScanGet200Response>, Integer, Hash)> api_v1_trials_eligibility_generalizability_scan_get_with_http_info(opts)

```ruby
begin
  # eligibility_generalizability_scan
  data, status_code, headers = api_instance.api_v1_trials_eligibility_generalizability_scan_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1TrialsEligibilityGeneralizabilityScanGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1TrialApi->api_v1_trials_eligibility_generalizability_scan_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  | [optional] |
| **limit** | **Integer** |  | [optional] |
| **nct_ids** | [**Array&lt;String&gt;**](String.md) |  | [optional] |

### Return type

[**ApiV1TrialsEligibilityGeneralizabilityScanGet200Response**](ApiV1TrialsEligibilityGeneralizabilityScanGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_trials_history_and_prediction_get

> <ApiV1TrialsHistoryAndPredictionGet200Response> api_v1_trials_history_and_prediction_get(opts)

history_and_prediction

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1TrialApi.new
opts = {
  nct_id: 'NCT05734196' # String | 
}

begin
  # history_and_prediction
  result = api_instance.api_v1_trials_history_and_prediction_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1TrialApi->api_v1_trials_history_and_prediction_get: #{e}"
end
```

#### Using the api_v1_trials_history_and_prediction_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1TrialsHistoryAndPredictionGet200Response>, Integer, Hash)> api_v1_trials_history_and_prediction_get_with_http_info(opts)

```ruby
begin
  # history_and_prediction
  data, status_code, headers = api_instance.api_v1_trials_history_and_prediction_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1TrialsHistoryAndPredictionGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1TrialApi->api_v1_trials_history_and_prediction_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **nct_id** | **String** |  | [optional] |

### Return type

[**ApiV1TrialsHistoryAndPredictionGet200Response**](ApiV1TrialsHistoryAndPredictionGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_trials_records_get

> <ApiV1TrialsRecordsGet200Response> api_v1_trials_records_get(opts)

records

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1TrialApi.new
opts = {
  include_extracted: true, # Boolean | 
  nct_ids: ['inner_example'] # Array<String> | 
}

begin
  # records
  result = api_instance.api_v1_trials_records_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1TrialApi->api_v1_trials_records_get: #{e}"
end
```

#### Using the api_v1_trials_records_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1TrialsRecordsGet200Response>, Integer, Hash)> api_v1_trials_records_get_with_http_info(opts)

```ruby
begin
  # records
  data, status_code, headers = api_instance.api_v1_trials_records_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1TrialsRecordsGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1TrialApi->api_v1_trials_records_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **include_extracted** | **Boolean** |  | [optional] |
| **nct_ids** | [**Array&lt;String&gt;**](String.md) |  | [optional] |

### Return type

[**ApiV1TrialsRecordsGet200Response**](ApiV1TrialsRecordsGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_trials_reporting_completeness_get

> <ApiV1TrialsReportingCompletenessGet200Response> api_v1_trials_reporting_completeness_get(opts)

reporting_completeness

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1TrialApi.new
opts = {
  condition_id: 'D010661' # String | 
}

begin
  # reporting_completeness
  result = api_instance.api_v1_trials_reporting_completeness_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1TrialApi->api_v1_trials_reporting_completeness_get: #{e}"
end
```

#### Using the api_v1_trials_reporting_completeness_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1TrialsReportingCompletenessGet200Response>, Integer, Hash)> api_v1_trials_reporting_completeness_get_with_http_info(opts)

```ruby
begin
  # reporting_completeness
  data, status_code, headers = api_instance.api_v1_trials_reporting_completeness_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1TrialsReportingCompletenessGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1TrialApi->api_v1_trials_reporting_completeness_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  | [optional] |

### Return type

[**ApiV1TrialsReportingCompletenessGet200Response**](ApiV1TrialsReportingCompletenessGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

