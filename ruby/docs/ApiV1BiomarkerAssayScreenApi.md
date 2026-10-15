# PharmaObservatory::ApiV1BiomarkerAssayScreenApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_biomarker_assay_screen_get**](ApiV1BiomarkerAssayScreenApi.md#api_v1_biomarker_assay_screen_get) | **GET** /api/v1/biomarker_assay_screen | index |


## api_v1_biomarker_assay_screen_get

> <ApiV1BiomarkerAssayScreenGet200Response> api_v1_biomarker_assay_screen_get(opts)

index

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1BiomarkerAssayScreenApi.new
opts = {
  assay_provider_company: 'Roche', # String | 
  document_id: 'doc-pubmed-biomarker-0002', # String | 
  gene: 'BRCA2', # String | 
  gene_resolution_status: 'unresolved', # String | 
  pmid: 20349101, # Integer | 
  validation_type_mentioned: 'clinical' # String | 
}

begin
  # index
  result = api_instance.api_v1_biomarker_assay_screen_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1BiomarkerAssayScreenApi->api_v1_biomarker_assay_screen_get: #{e}"
end
```

#### Using the api_v1_biomarker_assay_screen_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1BiomarkerAssayScreenGet200Response>, Integer, Hash)> api_v1_biomarker_assay_screen_get_with_http_info(opts)

```ruby
begin
  # index
  data, status_code, headers = api_instance.api_v1_biomarker_assay_screen_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1BiomarkerAssayScreenGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1BiomarkerAssayScreenApi->api_v1_biomarker_assay_screen_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **assay_provider_company** | **String** |  | [optional] |
| **document_id** | **String** |  | [optional] |
| **gene** | **String** |  | [optional] |
| **gene_resolution_status** | **String** |  | [optional] |
| **pmid** | **Integer** |  | [optional] |
| **validation_type_mentioned** | **String** |  | [optional] |

### Return type

[**ApiV1BiomarkerAssayScreenGet200Response**](ApiV1BiomarkerAssayScreenGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

