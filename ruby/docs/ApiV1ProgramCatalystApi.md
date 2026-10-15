# PharmaObservatory::ApiV1ProgramCatalystApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_licensing_candidates_get**](ApiV1ProgramCatalystApi.md#api_v1_licensing_candidates_get) | **GET** /api/v1/licensing_candidates | licensing_candidates |
| [**api_v1_org_pipeline_stages_get**](ApiV1ProgramCatalystApi.md#api_v1_org_pipeline_stages_get) | **GET** /api/v1/org_pipeline_stages | pipeline_stages |
| [**api_v1_patent_cliff_exposure_get**](ApiV1ProgramCatalystApi.md#api_v1_patent_cliff_exposure_get) | **GET** /api/v1/patent_cliff_exposure | patent_cliff_exposure |
| [**api_v1_program_catalysts_get**](ApiV1ProgramCatalystApi.md#api_v1_program_catalysts_get) | **GET** /api/v1/program_catalysts | catalysts |


## api_v1_licensing_candidates_get

> <ApiV1LicensingCandidatesGet200Response> api_v1_licensing_candidates_get(opts)

licensing_candidates

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1ProgramCatalystApi.new
opts = {
  condition_id: 'D003920' # String | 
}

begin
  # licensing_candidates
  result = api_instance.api_v1_licensing_candidates_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ProgramCatalystApi->api_v1_licensing_candidates_get: #{e}"
end
```

#### Using the api_v1_licensing_candidates_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1LicensingCandidatesGet200Response>, Integer, Hash)> api_v1_licensing_candidates_get_with_http_info(opts)

```ruby
begin
  # licensing_candidates
  data, status_code, headers = api_instance.api_v1_licensing_candidates_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1LicensingCandidatesGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ProgramCatalystApi->api_v1_licensing_candidates_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  | [optional] |

### Return type

[**ApiV1LicensingCandidatesGet200Response**](ApiV1LicensingCandidatesGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_org_pipeline_stages_get

> <ApiV1OrgPipelineStagesGet200Response> api_v1_org_pipeline_stages_get(opts)

pipeline_stages

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1ProgramCatalystApi.new
opts = {
  company_id: 'edgar:0001048477' # String | 
}

begin
  # pipeline_stages
  result = api_instance.api_v1_org_pipeline_stages_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ProgramCatalystApi->api_v1_org_pipeline_stages_get: #{e}"
end
```

#### Using the api_v1_org_pipeline_stages_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1OrgPipelineStagesGet200Response>, Integer, Hash)> api_v1_org_pipeline_stages_get_with_http_info(opts)

```ruby
begin
  # pipeline_stages
  data, status_code, headers = api_instance.api_v1_org_pipeline_stages_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1OrgPipelineStagesGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ProgramCatalystApi->api_v1_org_pipeline_stages_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** |  | [optional] |

### Return type

[**ApiV1OrgPipelineStagesGet200Response**](ApiV1OrgPipelineStagesGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_patent_cliff_exposure_get

> <ApiV1PatentCliffExposureGet200Response> api_v1_patent_cliff_exposure_get(opts)

patent_cliff_exposure

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1ProgramCatalystApi.new
opts = {
  company_ids: ['inner_example'], # Array<String> | 
  drug_ids: ['inner_example'] # Array<String> | 
}

begin
  # patent_cliff_exposure
  result = api_instance.api_v1_patent_cliff_exposure_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ProgramCatalystApi->api_v1_patent_cliff_exposure_get: #{e}"
end
```

#### Using the api_v1_patent_cliff_exposure_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1PatentCliffExposureGet200Response>, Integer, Hash)> api_v1_patent_cliff_exposure_get_with_http_info(opts)

```ruby
begin
  # patent_cliff_exposure
  data, status_code, headers = api_instance.api_v1_patent_cliff_exposure_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1PatentCliffExposureGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ProgramCatalystApi->api_v1_patent_cliff_exposure_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_ids** | [**Array&lt;String&gt;**](String.md) |  | [optional] |
| **drug_ids** | [**Array&lt;String&gt;**](String.md) |  | [optional] |

### Return type

[**ApiV1PatentCliffExposureGet200Response**](ApiV1PatentCliffExposureGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_program_catalysts_get

> <ApiV1ProgramCatalystsGet200Response> api_v1_program_catalysts_get(opts)

catalysts

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1ProgramCatalystApi.new
opts = {
  company_ids: ['inner_example'], # Array<String> | 
  drug_ids: ['inner_example'] # Array<String> | 
}

begin
  # catalysts
  result = api_instance.api_v1_program_catalysts_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ProgramCatalystApi->api_v1_program_catalysts_get: #{e}"
end
```

#### Using the api_v1_program_catalysts_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1ProgramCatalystsGet200Response>, Integer, Hash)> api_v1_program_catalysts_get_with_http_info(opts)

```ruby
begin
  # catalysts
  data, status_code, headers = api_instance.api_v1_program_catalysts_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1ProgramCatalystsGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1ProgramCatalystApi->api_v1_program_catalysts_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_ids** | [**Array&lt;String&gt;**](String.md) |  | [optional] |
| **drug_ids** | [**Array&lt;String&gt;**](String.md) |  | [optional] |

### Return type

[**ApiV1ProgramCatalystsGet200Response**](ApiV1ProgramCatalystsGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

