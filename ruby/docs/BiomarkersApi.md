# PharmaObservatory::BiomarkersApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**biomarkers_biomarker_bqp_assertions**](BiomarkersApi.md#biomarkers_biomarker_bqp_assertions) | **GET** /api/v1/biomarkers/{id}/biomarker_bqp_assertions | FDA biomarker qualification projects |
| [**biomarkers_biomarker_cdx_assertions**](BiomarkersApi.md#biomarkers_biomarker_cdx_assertions) | **GET** /api/v1/biomarkers/{id}/biomarker_cdx_assertions | FDA companion diagnostic assertions |
| [**biomarkers_biomarker_evidence**](BiomarkersApi.md#biomarkers_biomarker_evidence) | **GET** /api/v1/biomarkers/{id}/biomarker_evidence | Biomarker evidence |
| [**biomarkers_biomarker_literature**](BiomarkersApi.md#biomarkers_biomarker_literature) | **GET** /api/v1/biomarkers/{id}/biomarker_literature | Trial-cited literature mentions |
| [**biomarkers_biomarker_publishing_companies**](BiomarkersApi.md#biomarkers_biomarker_publishing_companies) | **GET** /api/v1/biomarkers/{id}/biomarker_publishing_companies | Companies publishing on a biomarker |
| [**biomarkers_biomarker_variant_mentions**](BiomarkersApi.md#biomarkers_biomarker_variant_mentions) | **GET** /api/v1/biomarkers/{id}/biomarker_variant_mentions | Protein variant mentions |
| [**biomarkers_get**](BiomarkersApi.md#biomarkers_get) | **GET** /api/v1/biomarkers/{id} | One biomarker |
| [**biomarkers_list**](BiomarkersApi.md#biomarkers_list) | **GET** /api/v1/biomarkers | List biomarkers |


## biomarkers_biomarker_bqp_assertions

> <BiomarkerBqpAssertionsResponse> biomarkers_biomarker_bqp_assertions(id, opts)

FDA biomarker qualification projects

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 
opts = {
  limit: 2, # Integer | 
  sort: 'route_key' # String | 
}

begin
  # FDA biomarker qualification projects
  result = api_instance.biomarkers_biomarker_bqp_assertions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_bqp_assertions: #{e}"
end
```

#### Using the biomarkers_biomarker_bqp_assertions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerBqpAssertionsResponse>, Integer, Hash)> biomarkers_biomarker_bqp_assertions_with_http_info(id, opts)

```ruby
begin
  # FDA biomarker qualification projects
  data, status_code, headers = api_instance.biomarkers_biomarker_bqp_assertions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerBqpAssertionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_bqp_assertions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**BiomarkerBqpAssertionsResponse**](BiomarkerBqpAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_cdx_assertions

> <BiomarkerCdxAssertionsResponse> biomarkers_biomarker_cdx_assertions(id, opts)

FDA companion diagnostic assertions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 
opts = {
  limit: 2, # Integer | 
  sort: 'biomarker_id' # String | 
}

begin
  # FDA companion diagnostic assertions
  result = api_instance.biomarkers_biomarker_cdx_assertions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_cdx_assertions: #{e}"
end
```

#### Using the biomarkers_biomarker_cdx_assertions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerCdxAssertionsResponse>, Integer, Hash)> biomarkers_biomarker_cdx_assertions_with_http_info(id, opts)

```ruby
begin
  # FDA companion diagnostic assertions
  data, status_code, headers = api_instance.biomarkers_biomarker_cdx_assertions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerCdxAssertionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_cdx_assertions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**BiomarkerCdxAssertionsResponse**](BiomarkerCdxAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_evidence

> <BiomarkerEvidenceResponse> biomarkers_biomarker_evidence(id, opts)

Biomarker evidence

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 
opts = {
  limit: 2, # Integer | 
  sort: 'gene_symbol' # String | 
}

begin
  # Biomarker evidence
  result = api_instance.biomarkers_biomarker_evidence(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_evidence: #{e}"
end
```

#### Using the biomarkers_biomarker_evidence_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerEvidenceResponse>, Integer, Hash)> biomarkers_biomarker_evidence_with_http_info(id, opts)

```ruby
begin
  # Biomarker evidence
  data, status_code, headers = api_instance.biomarkers_biomarker_evidence_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerEvidenceResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_evidence_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**BiomarkerEvidenceResponse**](BiomarkerEvidenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_literature

> <BiomarkerLiteratureResponse> biomarkers_biomarker_literature(id, opts)

Trial-cited literature mentions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 
opts = {
  limit: 2, # Integer | 
  sort: 'pub_year' # String | 
}

begin
  # Trial-cited literature mentions
  result = api_instance.biomarkers_biomarker_literature(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_literature: #{e}"
end
```

#### Using the biomarkers_biomarker_literature_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerLiteratureResponse>, Integer, Hash)> biomarkers_biomarker_literature_with_http_info(id, opts)

```ruby
begin
  # Trial-cited literature mentions
  data, status_code, headers = api_instance.biomarkers_biomarker_literature_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerLiteratureResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_literature_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**BiomarkerLiteratureResponse**](BiomarkerLiteratureResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_publishing_companies

> <BiomarkerPublishingCompaniesResponse> biomarkers_biomarker_publishing_companies(id, opts)

Companies publishing on a biomarker

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 
opts = {
  limit: 2, # Integer | 
  sort: 'n_publications' # String | 
}

begin
  # Companies publishing on a biomarker
  result = api_instance.biomarkers_biomarker_publishing_companies(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_publishing_companies: #{e}"
end
```

#### Using the biomarkers_biomarker_publishing_companies_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerPublishingCompaniesResponse>, Integer, Hash)> biomarkers_biomarker_publishing_companies_with_http_info(id, opts)

```ruby
begin
  # Companies publishing on a biomarker
  data, status_code, headers = api_instance.biomarkers_biomarker_publishing_companies_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerPublishingCompaniesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_publishing_companies_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**BiomarkerPublishingCompaniesResponse**](BiomarkerPublishingCompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_variant_mentions

> <BiomarkerVariantMentionsResponse> biomarkers_biomarker_variant_mentions(id, opts)

Protein variant mentions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 
opts = {
  limit: 2, # Integer | 
  sort: 'n_papers' # String | 
}

begin
  # Protein variant mentions
  result = api_instance.biomarkers_biomarker_variant_mentions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_variant_mentions: #{e}"
end
```

#### Using the biomarkers_biomarker_variant_mentions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerVariantMentionsResponse>, Integer, Hash)> biomarkers_biomarker_variant_mentions_with_http_info(id, opts)

```ruby
begin
  # Protein variant mentions
  data, status_code, headers = api_instance.biomarkers_biomarker_variant_mentions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerVariantMentionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_variant_mentions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**BiomarkerVariantMentionsResponse**](BiomarkerVariantMentionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_get

> <BiomarkersRecord> biomarkers_get(id)

One biomarker

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # One biomarker
  result = api_instance.biomarkers_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_get: #{e}"
end
```

#### Using the biomarkers_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersRecord>, Integer, Hash)> biomarkers_get_with_http_info(id)

```ruby
begin
  # One biomarker
  data, status_code, headers = api_instance.biomarkers_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersRecord**](BiomarkersRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_list

> <BiomarkersResponse> biomarkers_list(opts)

List biomarkers

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
opts = {
  limit: 2, # Integer | 
  sort: 'evidence_tier' # String | 
}

begin
  # List biomarkers
  result = api_instance.biomarkers_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_list: #{e}"
end
```

#### Using the biomarkers_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersResponse>, Integer, Hash)> biomarkers_list_with_http_info(opts)

```ruby
begin
  # List biomarkers
  data, status_code, headers = api_instance.biomarkers_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_list_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**BiomarkersResponse**](BiomarkersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

