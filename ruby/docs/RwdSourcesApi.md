# PharmaObservatory::RwdSourcesApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**rwd_sources_get**](RwdSourcesApi.md#rwd_sources_get) | **GET** /api/v1/rwd_sources/{id} | One real-world data source |
| [**rwd_sources_list**](RwdSourcesApi.md#rwd_sources_list) | **GET** /api/v1/rwd_sources | List real-world data sources |
| [**rwd_sources_rwd_source_age_bands**](RwdSourcesApi.md#rwd_sources_rwd_source_age_bands) | **GET** /api/v1/rwd_sources/{id}/rwd_source_age_bands | Real-world data source age bands |
| [**rwd_sources_rwd_source_conditions**](RwdSourcesApi.md#rwd_sources_rwd_source_conditions) | **GET** /api/v1/rwd_sources/{id}/rwd_source_conditions | Real-world data sources named by registered studies of a condition |
| [**rwd_sources_rwd_source_etl_mappings**](RwdSourcesApi.md#rwd_sources_rwd_source_etl_mappings) | **GET** /api/v1/rwd_sources/{id}/rwd_source_etl_mappings | Real-world data source common data model mappings |
| [**rwd_sources_rwd_source_linkages**](RwdSourcesApi.md#rwd_sources_rwd_source_linkages) | **GET** /api/v1/rwd_sources/{id}/rwd_source_linkages | Real-world data source linkages |
| [**rwd_sources_rwd_source_types**](RwdSourcesApi.md#rwd_sources_rwd_source_types) | **GET** /api/v1/rwd_sources/{id}/rwd_source_types | Real-world data source types |
| [**rwd_sources_rwd_source_vocabularies**](RwdSourcesApi.md#rwd_sources_rwd_source_vocabularies) | **GET** /api/v1/rwd_sources/{id}/rwd_source_vocabularies | Real-world data source coding schemes |
| [**rwd_sources_rwd_study_data_sources**](RwdSourcesApi.md#rwd_sources_rwd_study_data_sources) | **GET** /api/v1/rwd_sources/{id}/rwd_study_data_sources | Registered real-world data study data sources |


## rwd_sources_get

> <RwdSourcesRecord> rwd_sources_get(id)

One real-world data source

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdSourcesApi.new
id = 900000201 # Integer | 

begin
  # One real-world data source
  result = api_instance.rwd_sources_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_get: #{e}"
end
```

#### Using the rwd_sources_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdSourcesRecord>, Integer, Hash)> rwd_sources_get_with_http_info(id)

```ruby
begin
  # One real-world data source
  data, status_code, headers = api_instance.rwd_sources_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdSourcesRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |

### Return type

[**RwdSourcesRecord**](RwdSourcesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_sources_list

> <RwdSourcesResponse> rwd_sources_list(opts)

List real-world data sources

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdSourcesApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List real-world data sources
  result = api_instance.rwd_sources_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_list: #{e}"
end
```

#### Using the rwd_sources_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdSourcesResponse>, Integer, Hash)> rwd_sources_list_with_http_info(opts)

```ruby
begin
  # List real-world data sources
  data, status_code, headers = api_instance.rwd_sources_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdSourcesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_list_with_http_info: #{e}"
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

[**RwdSourcesResponse**](RwdSourcesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_sources_rwd_source_age_bands

> <RwdSourceAgeBandsResponse> rwd_sources_rwd_source_age_bands(id, opts)

Real-world data source age bands

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdSourcesApi.new
id = 900000201 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Real-world data source age bands
  result = api_instance.rwd_sources_rwd_source_age_bands(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_age_bands: #{e}"
end
```

#### Using the rwd_sources_rwd_source_age_bands_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdSourceAgeBandsResponse>, Integer, Hash)> rwd_sources_rwd_source_age_bands_with_http_info(id, opts)

```ruby
begin
  # Real-world data source age bands
  data, status_code, headers = api_instance.rwd_sources_rwd_source_age_bands_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdSourceAgeBandsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_age_bands_with_http_info: #{e}"
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

[**RwdSourceAgeBandsResponse**](RwdSourceAgeBandsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_sources_rwd_source_conditions

> <RwdSourceConditionsResponse> rwd_sources_rwd_source_conditions(id, opts)

Real-world data sources named by registered studies of a condition

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdSourcesApi.new
id = 900000201 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Real-world data sources named by registered studies of a condition
  result = api_instance.rwd_sources_rwd_source_conditions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_conditions: #{e}"
end
```

#### Using the rwd_sources_rwd_source_conditions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdSourceConditionsResponse>, Integer, Hash)> rwd_sources_rwd_source_conditions_with_http_info(id, opts)

```ruby
begin
  # Real-world data sources named by registered studies of a condition
  data, status_code, headers = api_instance.rwd_sources_rwd_source_conditions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdSourceConditionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_conditions_with_http_info: #{e}"
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

[**RwdSourceConditionsResponse**](RwdSourceConditionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_sources_rwd_source_etl_mappings

> <RwdSourceEtlMappingsResponse> rwd_sources_rwd_source_etl_mappings(id, opts)

Real-world data source common data model mappings

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdSourcesApi.new
id = 900000201 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Real-world data source common data model mappings
  result = api_instance.rwd_sources_rwd_source_etl_mappings(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_etl_mappings: #{e}"
end
```

#### Using the rwd_sources_rwd_source_etl_mappings_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdSourceEtlMappingsResponse>, Integer, Hash)> rwd_sources_rwd_source_etl_mappings_with_http_info(id, opts)

```ruby
begin
  # Real-world data source common data model mappings
  data, status_code, headers = api_instance.rwd_sources_rwd_source_etl_mappings_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdSourceEtlMappingsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_etl_mappings_with_http_info: #{e}"
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

[**RwdSourceEtlMappingsResponse**](RwdSourceEtlMappingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_sources_rwd_source_linkages

> <RwdSourceLinkagesResponse> rwd_sources_rwd_source_linkages(id, opts)

Real-world data source linkages

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdSourcesApi.new
id = 900000201 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Real-world data source linkages
  result = api_instance.rwd_sources_rwd_source_linkages(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_linkages: #{e}"
end
```

#### Using the rwd_sources_rwd_source_linkages_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdSourceLinkagesResponse>, Integer, Hash)> rwd_sources_rwd_source_linkages_with_http_info(id, opts)

```ruby
begin
  # Real-world data source linkages
  data, status_code, headers = api_instance.rwd_sources_rwd_source_linkages_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdSourceLinkagesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_linkages_with_http_info: #{e}"
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

[**RwdSourceLinkagesResponse**](RwdSourceLinkagesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_sources_rwd_source_types

> <RwdSourceTypesResponse> rwd_sources_rwd_source_types(id, opts)

Real-world data source types

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdSourcesApi.new
id = 900000201 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Real-world data source types
  result = api_instance.rwd_sources_rwd_source_types(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_types: #{e}"
end
```

#### Using the rwd_sources_rwd_source_types_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdSourceTypesResponse>, Integer, Hash)> rwd_sources_rwd_source_types_with_http_info(id, opts)

```ruby
begin
  # Real-world data source types
  data, status_code, headers = api_instance.rwd_sources_rwd_source_types_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdSourceTypesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_types_with_http_info: #{e}"
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

[**RwdSourceTypesResponse**](RwdSourceTypesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_sources_rwd_source_vocabularies

> <RwdSourceVocabulariesResponse> rwd_sources_rwd_source_vocabularies(id, opts)

Real-world data source coding schemes

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdSourcesApi.new
id = 900000201 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Real-world data source coding schemes
  result = api_instance.rwd_sources_rwd_source_vocabularies(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_vocabularies: #{e}"
end
```

#### Using the rwd_sources_rwd_source_vocabularies_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdSourceVocabulariesResponse>, Integer, Hash)> rwd_sources_rwd_source_vocabularies_with_http_info(id, opts)

```ruby
begin
  # Real-world data source coding schemes
  data, status_code, headers = api_instance.rwd_sources_rwd_source_vocabularies_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdSourceVocabulariesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_source_vocabularies_with_http_info: #{e}"
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

[**RwdSourceVocabulariesResponse**](RwdSourceVocabulariesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_sources_rwd_study_data_sources

> <RwdStudyDataSourcesResponse> rwd_sources_rwd_study_data_sources(id, opts)

Registered real-world data study data sources

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdSourcesApi.new
id = 900000201 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Registered real-world data study data sources
  result = api_instance.rwd_sources_rwd_study_data_sources(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_study_data_sources: #{e}"
end
```

#### Using the rwd_sources_rwd_study_data_sources_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdStudyDataSourcesResponse>, Integer, Hash)> rwd_sources_rwd_study_data_sources_with_http_info(id, opts)

```ruby
begin
  # Registered real-world data study data sources
  data, status_code, headers = api_instance.rwd_sources_rwd_study_data_sources_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdStudyDataSourcesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdSourcesApi->rwd_sources_rwd_study_data_sources_with_http_info: #{e}"
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

[**RwdStudyDataSourcesResponse**](RwdStudyDataSourcesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

