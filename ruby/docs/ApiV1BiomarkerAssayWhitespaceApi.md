# PharmaObservatory::ApiV1BiomarkerAssayWhitespaceApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_biomarker_assay_whitespace_get**](ApiV1BiomarkerAssayWhitespaceApi.md#api_v1_biomarker_assay_whitespace_get) | **GET** /api/v1/biomarker_assay_whitespace | index |


## api_v1_biomarker_assay_whitespace_get

> <ApiV1BiomarkerAssayWhitespaceGet200Response> api_v1_biomarker_assay_whitespace_get(opts)

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

api_instance = PharmaObservatory::ApiV1BiomarkerAssayWhitespaceApi.new
opts = {
  bucket: 'served', # String | 
  gene: 'A1BG', # String | 
  gene_evidence_tier: 'literature_only', # String | 
  gene_supply_status: 'declared_positive', # String | 
  limit: 700, # Integer | 
  q: 'EGFR' # String | 
}

begin
  # index
  result = api_instance.api_v1_biomarker_assay_whitespace_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1BiomarkerAssayWhitespaceApi->api_v1_biomarker_assay_whitespace_get: #{e}"
end
```

#### Using the api_v1_biomarker_assay_whitespace_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1BiomarkerAssayWhitespaceGet200Response>, Integer, Hash)> api_v1_biomarker_assay_whitespace_get_with_http_info(opts)

```ruby
begin
  # index
  data, status_code, headers = api_instance.api_v1_biomarker_assay_whitespace_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1BiomarkerAssayWhitespaceGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1BiomarkerAssayWhitespaceApi->api_v1_biomarker_assay_whitespace_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **bucket** | **String** |  | [optional] |
| **gene** | **String** |  | [optional] |
| **gene_evidence_tier** | **String** |  | [optional] |
| **gene_supply_status** | **String** |  | [optional] |
| **limit** | **Integer** |  | [optional] |
| **q** | **String** |  | [optional] |

### Return type

[**ApiV1BiomarkerAssayWhitespaceGet200Response**](ApiV1BiomarkerAssayWhitespaceGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

