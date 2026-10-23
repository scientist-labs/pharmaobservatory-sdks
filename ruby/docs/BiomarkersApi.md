# PharmaObservatory::BiomarkersApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**biomarkers_biomarker_allele_frequencies**](BiomarkersApi.md#biomarkers_biomarker_allele_frequencies) | **GET** /api/v1/biomarkers/{id}/biomarker_allele_frequencies | Population allele frequency of protein variants (gnomAD) |
| [**biomarkers_biomarker_assay_evidence**](BiomarkersApi.md#biomarkers_biomarker_assay_evidence) | **GET** /api/v1/biomarkers/{id}/biomarker_assay_evidence | Biomarker assay evidence from the literature |
| [**biomarkers_biomarker_assay_providers**](BiomarkersApi.md#biomarkers_biomarker_assay_providers) | **GET** /api/v1/biomarkers/{id}/biomarker_assay_providers | Biomarker assay providers (from validating papers) |
| [**biomarkers_biomarker_assay_whitespace**](BiomarkersApi.md#biomarkers_biomarker_assay_whitespace) | **GET** /api/v1/biomarkers/{id}/biomarker_assay_whitespace | Biomarker assay whitespace (evidence against supplier supply) |
| [**biomarkers_biomarker_bqp_assertions**](BiomarkersApi.md#biomarkers_biomarker_bqp_assertions) | **GET** /api/v1/biomarkers/{id}/biomarker_bqp_assertions | FDA biomarker qualification projects |
| [**biomarkers_biomarker_cdx_assertions**](BiomarkersApi.md#biomarkers_biomarker_cdx_assertions) | **GET** /api/v1/biomarkers/{id}/biomarker_cdx_assertions | FDA companion diagnostic assertions |
| [**biomarkers_biomarker_evidence**](BiomarkersApi.md#biomarkers_biomarker_evidence) | **GET** /api/v1/biomarkers/{id}/biomarker_evidence | Biomarker evidence |
| [**biomarkers_biomarker_evidence_by_year**](BiomarkersApi.md#biomarkers_biomarker_evidence_by_year) | **GET** /api/v1/biomarkers/{id}/biomarker_evidence_by_year | Biomarker evidence by year (trial and publication mentions) |
| [**biomarkers_biomarker_literature**](BiomarkersApi.md#biomarkers_biomarker_literature) | **GET** /api/v1/biomarkers/{id}/biomarker_literature | Trial-cited literature mentions |
| [**biomarkers_biomarker_publishing_companies**](BiomarkersApi.md#biomarkers_biomarker_publishing_companies) | **GET** /api/v1/biomarkers/{id}/biomarker_publishing_companies | Companies publishing on a biomarker |
| [**biomarkers_biomarker_relations**](BiomarkersApi.md#biomarkers_biomarker_relations) | **GET** /api/v1/biomarkers/{id}/biomarker_relations | Literature-extracted relations |
| [**biomarkers_biomarker_sample_coverage**](BiomarkersApi.md#biomarkers_biomarker_sample_coverage) | **GET** /api/v1/biomarkers/{id}/biomarker_sample_coverage | Human sample screening coverage (GDC) |
| [**biomarkers_biomarker_sample_prevalence**](BiomarkersApi.md#biomarkers_biomarker_sample_prevalence) | **GET** /api/v1/biomarkers/{id}/biomarker_sample_prevalence | Human sample screening reference (GDC cohorts) |
| [**biomarkers_biomarker_traversal_status**](BiomarkersApi.md#biomarkers_biomarker_traversal_status) | **GET** /api/v1/biomarkers/{id}/biomarker_traversal_status | Biomarker coverage by traversal |
| [**biomarkers_biomarker_validation**](BiomarkersApi.md#biomarkers_biomarker_validation) | **GET** /api/v1/biomarkers/{id}/biomarker_validation | Biomarker validation ladder |
| [**biomarkers_biomarker_variant_mentions**](BiomarkersApi.md#biomarkers_biomarker_variant_mentions) | **GET** /api/v1/biomarkers/{id}/biomarker_variant_mentions | Protein variant mentions |
| [**biomarkers_get**](BiomarkersApi.md#biomarkers_get) | **GET** /api/v1/biomarkers/{id} | One biomarker |
| [**biomarkers_list**](BiomarkersApi.md#biomarkers_list) | **GET** /api/v1/biomarkers | List biomarkers |


## biomarkers_biomarker_allele_frequencies

> <BiomarkerAlleleFrequenciesResponse> biomarkers_biomarker_allele_frequencies(id, opts)

Population allele frequency of protein variants (gnomAD)

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Population allele frequency of protein variants (gnomAD)
  result = api_instance.biomarkers_biomarker_allele_frequencies(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_allele_frequencies: #{e}"
end
```

#### Using the biomarkers_biomarker_allele_frequencies_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerAlleleFrequenciesResponse>, Integer, Hash)> biomarkers_biomarker_allele_frequencies_with_http_info(id, opts)

```ruby
begin
  # Population allele frequency of protein variants (gnomAD)
  data, status_code, headers = api_instance.biomarkers_biomarker_allele_frequencies_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerAlleleFrequenciesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_allele_frequencies_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerAlleleFrequenciesResponse**](BiomarkerAlleleFrequenciesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_assay_evidence

> <BiomarkerAssayEvidenceResponse> biomarkers_biomarker_assay_evidence(id, opts)

Biomarker assay evidence from the literature

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Biomarker assay evidence from the literature
  result = api_instance.biomarkers_biomarker_assay_evidence(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_assay_evidence: #{e}"
end
```

#### Using the biomarkers_biomarker_assay_evidence_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerAssayEvidenceResponse>, Integer, Hash)> biomarkers_biomarker_assay_evidence_with_http_info(id, opts)

```ruby
begin
  # Biomarker assay evidence from the literature
  data, status_code, headers = api_instance.biomarkers_biomarker_assay_evidence_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerAssayEvidenceResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_assay_evidence_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerAssayEvidenceResponse**](BiomarkerAssayEvidenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_assay_providers

> <BiomarkerAssayProvidersResponse> biomarkers_biomarker_assay_providers(id, opts)

Biomarker assay providers (from validating papers)

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Biomarker assay providers (from validating papers)
  result = api_instance.biomarkers_biomarker_assay_providers(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_assay_providers: #{e}"
end
```

#### Using the biomarkers_biomarker_assay_providers_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerAssayProvidersResponse>, Integer, Hash)> biomarkers_biomarker_assay_providers_with_http_info(id, opts)

```ruby
begin
  # Biomarker assay providers (from validating papers)
  data, status_code, headers = api_instance.biomarkers_biomarker_assay_providers_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerAssayProvidersResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_assay_providers_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerAssayProvidersResponse**](BiomarkerAssayProvidersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_assay_whitespace

> <BiomarkerAssayWhitespaceResponse> biomarkers_biomarker_assay_whitespace(id, opts)

Biomarker assay whitespace (evidence against supplier supply)

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Biomarker assay whitespace (evidence against supplier supply)
  result = api_instance.biomarkers_biomarker_assay_whitespace(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_assay_whitespace: #{e}"
end
```

#### Using the biomarkers_biomarker_assay_whitespace_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerAssayWhitespaceResponse>, Integer, Hash)> biomarkers_biomarker_assay_whitespace_with_http_info(id, opts)

```ruby
begin
  # Biomarker assay whitespace (evidence against supplier supply)
  data, status_code, headers = api_instance.biomarkers_biomarker_assay_whitespace_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerAssayWhitespaceResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_assay_whitespace_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerAssayWhitespaceResponse**](BiomarkerAssayWhitespaceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
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
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
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
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
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
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerEvidenceResponse**](BiomarkerEvidenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_evidence_by_year

> <BiomarkerEvidenceByYearResponse> biomarkers_biomarker_evidence_by_year(id, opts)

Biomarker evidence by year (trial and publication mentions)

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Biomarker evidence by year (trial and publication mentions)
  result = api_instance.biomarkers_biomarker_evidence_by_year(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_evidence_by_year: #{e}"
end
```

#### Using the biomarkers_biomarker_evidence_by_year_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerEvidenceByYearResponse>, Integer, Hash)> biomarkers_biomarker_evidence_by_year_with_http_info(id, opts)

```ruby
begin
  # Biomarker evidence by year (trial and publication mentions)
  data, status_code, headers = api_instance.biomarkers_biomarker_evidence_by_year_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerEvidenceByYearResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_evidence_by_year_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerEvidenceByYearResponse**](BiomarkerEvidenceByYearResponse.md)

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
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
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
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
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerPublishingCompaniesResponse**](BiomarkerPublishingCompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_relations

> <BiomarkerRelationsResponse> biomarkers_biomarker_relations(id, opts)

Literature-extracted relations

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Literature-extracted relations
  result = api_instance.biomarkers_biomarker_relations(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_relations: #{e}"
end
```

#### Using the biomarkers_biomarker_relations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerRelationsResponse>, Integer, Hash)> biomarkers_biomarker_relations_with_http_info(id, opts)

```ruby
begin
  # Literature-extracted relations
  data, status_code, headers = api_instance.biomarkers_biomarker_relations_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerRelationsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_relations_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerRelationsResponse**](BiomarkerRelationsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_sample_coverage

> <BiomarkerSampleCoverageResponse> biomarkers_biomarker_sample_coverage(id, opts)

Human sample screening coverage (GDC)

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Human sample screening coverage (GDC)
  result = api_instance.biomarkers_biomarker_sample_coverage(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_sample_coverage: #{e}"
end
```

#### Using the biomarkers_biomarker_sample_coverage_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerSampleCoverageResponse>, Integer, Hash)> biomarkers_biomarker_sample_coverage_with_http_info(id, opts)

```ruby
begin
  # Human sample screening coverage (GDC)
  data, status_code, headers = api_instance.biomarkers_biomarker_sample_coverage_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerSampleCoverageResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_sample_coverage_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerSampleCoverageResponse**](BiomarkerSampleCoverageResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_sample_prevalence

> <BiomarkerSamplePrevalenceResponse> biomarkers_biomarker_sample_prevalence(id, opts)

Human sample screening reference (GDC cohorts)

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Human sample screening reference (GDC cohorts)
  result = api_instance.biomarkers_biomarker_sample_prevalence(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_sample_prevalence: #{e}"
end
```

#### Using the biomarkers_biomarker_sample_prevalence_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerSamplePrevalenceResponse>, Integer, Hash)> biomarkers_biomarker_sample_prevalence_with_http_info(id, opts)

```ruby
begin
  # Human sample screening reference (GDC cohorts)
  data, status_code, headers = api_instance.biomarkers_biomarker_sample_prevalence_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerSamplePrevalenceResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_sample_prevalence_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerSamplePrevalenceResponse**](BiomarkerSamplePrevalenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_traversal_status

> <BiomarkerTraversalStatusResponse> biomarkers_biomarker_traversal_status(id, opts)

Biomarker coverage by traversal

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Biomarker coverage by traversal
  result = api_instance.biomarkers_biomarker_traversal_status(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_traversal_status: #{e}"
end
```

#### Using the biomarkers_biomarker_traversal_status_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerTraversalStatusResponse>, Integer, Hash)> biomarkers_biomarker_traversal_status_with_http_info(id, opts)

```ruby
begin
  # Biomarker coverage by traversal
  data, status_code, headers = api_instance.biomarkers_biomarker_traversal_status_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerTraversalStatusResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_traversal_status_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerTraversalStatusResponse**](BiomarkerTraversalStatusResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_biomarker_validation

> <BiomarkerValidationResponse> biomarkers_biomarker_validation(id, opts)

Biomarker validation ladder

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Biomarker validation ladder
  result = api_instance.biomarkers_biomarker_validation(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_validation: #{e}"
end
```

#### Using the biomarkers_biomarker_validation_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerValidationResponse>, Integer, Hash)> biomarkers_biomarker_validation_with_http_info(id, opts)

```ruby
begin
  # Biomarker validation ladder
  data, status_code, headers = api_instance.biomarkers_biomarker_validation_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerValidationResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_biomarker_validation_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkerValidationResponse**](BiomarkerValidationResponse.md)

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
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
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

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
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
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
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**BiomarkersResponse**](BiomarkersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

