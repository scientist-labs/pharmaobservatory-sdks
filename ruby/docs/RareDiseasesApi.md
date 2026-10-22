# PharmaObservatory::RareDiseasesApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**rare_diseases_get**](RareDiseasesApi.md#rare_diseases_get) | **GET** /api/v1/rare_diseases/{id} | One rare disease |
| [**rare_diseases_list**](RareDiseasesApi.md#rare_diseases_list) | **GET** /api/v1/rare_diseases | List rare diseases |
| [**rare_diseases_rare_disease_genes**](RareDiseasesApi.md#rare_diseases_rare_disease_genes) | **GET** /api/v1/rare_diseases/{id}/rare_disease_genes | Rare disease genes |
| [**rare_diseases_rare_disease_hierarchy**](RareDiseasesApi.md#rare_diseases_rare_disease_hierarchy) | **GET** /api/v1/rare_diseases/{id}/rare_disease_hierarchy | Rare disease hierarchy |
| [**rare_diseases_rare_disease_investigator_coverage**](RareDiseasesApi.md#rare_diseases_rare_disease_investigator_coverage) | **GET** /api/v1/rare_diseases/{id}/rare_disease_investigator_coverage | Rare disease investigator coverage |
| [**rare_diseases_rare_disease_investigators**](RareDiseasesApi.md#rare_diseases_rare_disease_investigators) | **GET** /api/v1/rare_diseases/{id}/rare_disease_investigators | Rare disease investigators |
| [**rare_diseases_rare_disease_mappings**](RareDiseasesApi.md#rare_diseases_rare_disease_mappings) | **GET** /api/v1/rare_diseases/{id}/rare_disease_mappings | Rare disease mappings |
| [**rare_diseases_rare_disease_organizations**](RareDiseasesApi.md#rare_diseases_rare_disease_organizations) | **GET** /api/v1/rare_diseases/{id}/rare_disease_organizations | Rare disease organizations |
| [**rare_diseases_rare_disease_prevalence**](RareDiseasesApi.md#rare_diseases_rare_disease_prevalence) | **GET** /api/v1/rare_diseases/{id}/rare_disease_prevalence | Rare disease prevalence |
| [**rare_diseases_rare_disease_rarity**](RareDiseasesApi.md#rare_diseases_rare_disease_rarity) | **GET** /api/v1/rare_diseases/{id}/rare_disease_rarity | Rare disease rarity assessments |
| [**rare_diseases_rare_disease_rarity_evidence**](RareDiseasesApi.md#rare_diseases_rare_disease_rarity_evidence) | **GET** /api/v1/rare_diseases/{id}/rare_disease_rarity_evidence | Rare disease rarity evidence |
| [**rare_diseases_rare_disease_replacements**](RareDiseasesApi.md#rare_diseases_rare_disease_replacements) | **GET** /api/v1/rare_diseases/{id}/rare_disease_replacements | Rare disease replacements |
| [**rare_diseases_rare_disease_site_coverage**](RareDiseasesApi.md#rare_diseases_rare_disease_site_coverage) | **GET** /api/v1/rare_diseases/{id}/rare_disease_site_coverage | Rare disease site coverage |
| [**rare_diseases_rare_disease_sponsors**](RareDiseasesApi.md#rare_diseases_rare_disease_sponsors) | **GET** /api/v1/rare_diseases/{id}/rare_disease_sponsors | Rare disease sponsors |
| [**rare_diseases_rare_disease_trial_coverage**](RareDiseasesApi.md#rare_diseases_rare_disease_trial_coverage) | **GET** /api/v1/rare_diseases/{id}/rare_disease_trial_coverage | Rare disease trial coverage |
| [**rare_diseases_rare_disease_trial_sites**](RareDiseasesApi.md#rare_diseases_rare_disease_trial_sites) | **GET** /api/v1/rare_diseases/{id}/rare_disease_trial_sites | Rare disease trial sites |
| [**rare_diseases_rare_disease_trials**](RareDiseasesApi.md#rare_diseases_rare_disease_trials) | **GET** /api/v1/rare_diseases/{id}/rare_disease_trials | Rare disease trials |


## rare_diseases_get

> <RareDiseasesRecord> rare_diseases_get(id)

One rare disease

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 

begin
  # One rare disease
  result = api_instance.rare_diseases_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_get: #{e}"
end
```

#### Using the rare_diseases_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseasesRecord>, Integer, Hash)> rare_diseases_get_with_http_info(id)

```ruby
begin
  # One rare disease
  data, status_code, headers = api_instance.rare_diseases_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseasesRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |

### Return type

[**RareDiseasesRecord**](RareDiseasesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_list

> <RareDiseasesResponse> rare_diseases_list(opts)

List rare diseases

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List rare diseases
  result = api_instance.rare_diseases_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_list: #{e}"
end
```

#### Using the rare_diseases_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseasesResponse>, Integer, Hash)> rare_diseases_list_with_http_info(opts)

```ruby
begin
  # List rare diseases
  data, status_code, headers = api_instance.rare_diseases_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseasesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_list_with_http_info: #{e}"
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

[**RareDiseasesResponse**](RareDiseasesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_genes

> <RareDiseaseGenesResponse> rare_diseases_rare_disease_genes(id, opts)

Rare disease genes

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease genes
  result = api_instance.rare_diseases_rare_disease_genes(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_genes: #{e}"
end
```

#### Using the rare_diseases_rare_disease_genes_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseGenesResponse>, Integer, Hash)> rare_diseases_rare_disease_genes_with_http_info(id, opts)

```ruby
begin
  # Rare disease genes
  data, status_code, headers = api_instance.rare_diseases_rare_disease_genes_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseGenesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_genes_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseGenesResponse**](RareDiseaseGenesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_hierarchy

> <RareDiseaseHierarchyResponse> rare_diseases_rare_disease_hierarchy(id, opts)

Rare disease hierarchy

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease hierarchy
  result = api_instance.rare_diseases_rare_disease_hierarchy(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_hierarchy: #{e}"
end
```

#### Using the rare_diseases_rare_disease_hierarchy_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseHierarchyResponse>, Integer, Hash)> rare_diseases_rare_disease_hierarchy_with_http_info(id, opts)

```ruby
begin
  # Rare disease hierarchy
  data, status_code, headers = api_instance.rare_diseases_rare_disease_hierarchy_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseHierarchyResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_hierarchy_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseHierarchyResponse**](RareDiseaseHierarchyResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_investigator_coverage

> <RareDiseaseInvestigatorCoverageResponse> rare_diseases_rare_disease_investigator_coverage(id, opts)

Rare disease investigator coverage

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease investigator coverage
  result = api_instance.rare_diseases_rare_disease_investigator_coverage(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_investigator_coverage: #{e}"
end
```

#### Using the rare_diseases_rare_disease_investigator_coverage_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseInvestigatorCoverageResponse>, Integer, Hash)> rare_diseases_rare_disease_investigator_coverage_with_http_info(id, opts)

```ruby
begin
  # Rare disease investigator coverage
  data, status_code, headers = api_instance.rare_diseases_rare_disease_investigator_coverage_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseInvestigatorCoverageResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_investigator_coverage_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseInvestigatorCoverageResponse**](RareDiseaseInvestigatorCoverageResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_investigators

> <RareDiseaseInvestigatorsResponse> rare_diseases_rare_disease_investigators(id, opts)

Rare disease investigators

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease investigators
  result = api_instance.rare_diseases_rare_disease_investigators(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_investigators: #{e}"
end
```

#### Using the rare_diseases_rare_disease_investigators_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseInvestigatorsResponse>, Integer, Hash)> rare_diseases_rare_disease_investigators_with_http_info(id, opts)

```ruby
begin
  # Rare disease investigators
  data, status_code, headers = api_instance.rare_diseases_rare_disease_investigators_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseInvestigatorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_investigators_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseInvestigatorsResponse**](RareDiseaseInvestigatorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_mappings

> <RareDiseaseMappingsResponse> rare_diseases_rare_disease_mappings(id, opts)

Rare disease mappings

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease mappings
  result = api_instance.rare_diseases_rare_disease_mappings(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_mappings: #{e}"
end
```

#### Using the rare_diseases_rare_disease_mappings_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseMappingsResponse>, Integer, Hash)> rare_diseases_rare_disease_mappings_with_http_info(id, opts)

```ruby
begin
  # Rare disease mappings
  data, status_code, headers = api_instance.rare_diseases_rare_disease_mappings_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseMappingsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_mappings_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseMappingsResponse**](RareDiseaseMappingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_organizations

> <RareDiseaseOrganizationsResponse> rare_diseases_rare_disease_organizations(id, opts)

Rare disease organizations

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease organizations
  result = api_instance.rare_diseases_rare_disease_organizations(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_organizations: #{e}"
end
```

#### Using the rare_diseases_rare_disease_organizations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseOrganizationsResponse>, Integer, Hash)> rare_diseases_rare_disease_organizations_with_http_info(id, opts)

```ruby
begin
  # Rare disease organizations
  data, status_code, headers = api_instance.rare_diseases_rare_disease_organizations_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseOrganizationsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_organizations_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseOrganizationsResponse**](RareDiseaseOrganizationsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_prevalence

> <RareDiseasePrevalenceResponse> rare_diseases_rare_disease_prevalence(id, opts)

Rare disease prevalence

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease prevalence
  result = api_instance.rare_diseases_rare_disease_prevalence(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_prevalence: #{e}"
end
```

#### Using the rare_diseases_rare_disease_prevalence_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseasePrevalenceResponse>, Integer, Hash)> rare_diseases_rare_disease_prevalence_with_http_info(id, opts)

```ruby
begin
  # Rare disease prevalence
  data, status_code, headers = api_instance.rare_diseases_rare_disease_prevalence_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseasePrevalenceResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_prevalence_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseasePrevalenceResponse**](RareDiseasePrevalenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_rarity

> <RareDiseaseRarityResponse> rare_diseases_rare_disease_rarity(id, opts)

Rare disease rarity assessments

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease rarity assessments
  result = api_instance.rare_diseases_rare_disease_rarity(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_rarity: #{e}"
end
```

#### Using the rare_diseases_rare_disease_rarity_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseRarityResponse>, Integer, Hash)> rare_diseases_rare_disease_rarity_with_http_info(id, opts)

```ruby
begin
  # Rare disease rarity assessments
  data, status_code, headers = api_instance.rare_diseases_rare_disease_rarity_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseRarityResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_rarity_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseRarityResponse**](RareDiseaseRarityResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_rarity_evidence

> <RareDiseaseRarityEvidenceResponse> rare_diseases_rare_disease_rarity_evidence(id, opts)

Rare disease rarity evidence

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease rarity evidence
  result = api_instance.rare_diseases_rare_disease_rarity_evidence(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_rarity_evidence: #{e}"
end
```

#### Using the rare_diseases_rare_disease_rarity_evidence_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseRarityEvidenceResponse>, Integer, Hash)> rare_diseases_rare_disease_rarity_evidence_with_http_info(id, opts)

```ruby
begin
  # Rare disease rarity evidence
  data, status_code, headers = api_instance.rare_diseases_rare_disease_rarity_evidence_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseRarityEvidenceResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_rarity_evidence_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseRarityEvidenceResponse**](RareDiseaseRarityEvidenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_replacements

> <RareDiseaseReplacementsResponse> rare_diseases_rare_disease_replacements(id, opts)

Rare disease replacements

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease replacements
  result = api_instance.rare_diseases_rare_disease_replacements(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_replacements: #{e}"
end
```

#### Using the rare_diseases_rare_disease_replacements_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseReplacementsResponse>, Integer, Hash)> rare_diseases_rare_disease_replacements_with_http_info(id, opts)

```ruby
begin
  # Rare disease replacements
  data, status_code, headers = api_instance.rare_diseases_rare_disease_replacements_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseReplacementsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_replacements_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseReplacementsResponse**](RareDiseaseReplacementsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_site_coverage

> <RareDiseaseSiteCoverageResponse> rare_diseases_rare_disease_site_coverage(id, opts)

Rare disease site coverage

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease site coverage
  result = api_instance.rare_diseases_rare_disease_site_coverage(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_site_coverage: #{e}"
end
```

#### Using the rare_diseases_rare_disease_site_coverage_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseSiteCoverageResponse>, Integer, Hash)> rare_diseases_rare_disease_site_coverage_with_http_info(id, opts)

```ruby
begin
  # Rare disease site coverage
  data, status_code, headers = api_instance.rare_diseases_rare_disease_site_coverage_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseSiteCoverageResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_site_coverage_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseSiteCoverageResponse**](RareDiseaseSiteCoverageResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_sponsors

> <RareDiseaseSponsorsResponse> rare_diseases_rare_disease_sponsors(id, opts)

Rare disease sponsors

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease sponsors
  result = api_instance.rare_diseases_rare_disease_sponsors(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_sponsors: #{e}"
end
```

#### Using the rare_diseases_rare_disease_sponsors_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseSponsorsResponse>, Integer, Hash)> rare_diseases_rare_disease_sponsors_with_http_info(id, opts)

```ruby
begin
  # Rare disease sponsors
  data, status_code, headers = api_instance.rare_diseases_rare_disease_sponsors_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseSponsorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_sponsors_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseSponsorsResponse**](RareDiseaseSponsorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_trial_coverage

> <RareDiseaseTrialCoverageResponse> rare_diseases_rare_disease_trial_coverage(id, opts)

Rare disease trial coverage

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease trial coverage
  result = api_instance.rare_diseases_rare_disease_trial_coverage(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_trial_coverage: #{e}"
end
```

#### Using the rare_diseases_rare_disease_trial_coverage_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseTrialCoverageResponse>, Integer, Hash)> rare_diseases_rare_disease_trial_coverage_with_http_info(id, opts)

```ruby
begin
  # Rare disease trial coverage
  data, status_code, headers = api_instance.rare_diseases_rare_disease_trial_coverage_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseTrialCoverageResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_trial_coverage_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseTrialCoverageResponse**](RareDiseaseTrialCoverageResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_trial_sites

> <RareDiseaseTrialSitesResponse> rare_diseases_rare_disease_trial_sites(id, opts)

Rare disease trial sites

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease trial sites
  result = api_instance.rare_diseases_rare_disease_trial_sites(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_trial_sites: #{e}"
end
```

#### Using the rare_diseases_rare_disease_trial_sites_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseTrialSitesResponse>, Integer, Hash)> rare_diseases_rare_disease_trial_sites_with_http_info(id, opts)

```ruby
begin
  # Rare disease trial sites
  data, status_code, headers = api_instance.rare_diseases_rare_disease_trial_sites_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseTrialSitesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_trial_sites_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseTrialSitesResponse**](RareDiseaseTrialSitesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rare_diseases_rare_disease_trials

> <RareDiseaseTrialsResponse> rare_diseases_rare_disease_trials(id, opts)

Rare disease trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RareDiseasesApi.new
id = 95708 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease trials
  result = api_instance.rare_diseases_rare_disease_trials(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_trials: #{e}"
end
```

#### Using the rare_diseases_rare_disease_trials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseTrialsResponse>, Integer, Hash)> rare_diseases_rare_disease_trials_with_http_info(id, opts)

```ruby
begin
  # Rare disease trials
  data, status_code, headers = api_instance.rare_diseases_rare_disease_trials_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseTrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RareDiseasesApi->rare_diseases_rare_disease_trials_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **cursor** | **String** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | [optional] |
| **limit** | **Integer** | Rows per page, up to 500 (default 50). | [optional] |
| **sort** | **String** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | [optional] |
| **group_by** | **String** | A dimension to group by; a date takes :year, :quarter or :month. | [optional] |
| **measure** | **String** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | [optional] |

### Return type

[**RareDiseaseTrialsResponse**](RareDiseaseTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

