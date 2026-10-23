# PharmaObservatory::RwdStudiesApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**rwd_studies_get**](RwdStudiesApi.md#rwd_studies_get) | **GET** /api/v1/rwd_studies/{id} | One real-world data study |
| [**rwd_studies_list**](RwdStudiesApi.md#rwd_studies_list) | **GET** /api/v1/rwd_studies | List real-world data studies |
| [**rwd_studies_rwd_conductor_studies**](RwdStudiesApi.md#rwd_studies_rwd_conductor_studies) | **GET** /api/v1/rwd_studies/{id}/rwd_conductor_studies | Real-world data study conductor studies (HMA-EMA catalogue) |
| [**rwd_studies_rwd_study_conditions**](RwdStudiesApi.md#rwd_studies_rwd_study_conditions) | **GET** /api/v1/rwd_studies/{id}/rwd_study_conditions | Registered real-world data study conditions |
| [**rwd_studies_rwd_study_data_sources**](RwdStudiesApi.md#rwd_studies_rwd_study_data_sources) | **GET** /api/v1/rwd_studies/{id}/rwd_study_data_sources | Registered real-world data study data sources |
| [**rwd_studies_rwd_study_drugs**](RwdStudiesApi.md#rwd_studies_rwd_study_drugs) | **GET** /api/v1/rwd_studies/{id}/rwd_study_drugs | Registered real-world data study drugs |
| [**rwd_studies_rwd_study_milestones**](RwdStudiesApi.md#rwd_studies_rwd_study_milestones) | **GET** /api/v1/rwd_studies/{id}/rwd_study_milestones | Registered real-world data study milestones |


## rwd_studies_get

> <RwdStudiesRecord> rwd_studies_get(id)

One real-world data study

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdStudiesApi.new
id = 900000301 # Integer | 

begin
  # One real-world data study
  result = api_instance.rwd_studies_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_get: #{e}"
end
```

#### Using the rwd_studies_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdStudiesRecord>, Integer, Hash)> rwd_studies_get_with_http_info(id)

```ruby
begin
  # One real-world data study
  data, status_code, headers = api_instance.rwd_studies_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdStudiesRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |

### Return type

[**RwdStudiesRecord**](RwdStudiesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_studies_list

> <RwdStudiesResponse> rwd_studies_list(opts)

List real-world data studies

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdStudiesApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List real-world data studies
  result = api_instance.rwd_studies_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_list: #{e}"
end
```

#### Using the rwd_studies_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdStudiesResponse>, Integer, Hash)> rwd_studies_list_with_http_info(opts)

```ruby
begin
  # List real-world data studies
  data, status_code, headers = api_instance.rwd_studies_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdStudiesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_list_with_http_info: #{e}"
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

[**RwdStudiesResponse**](RwdStudiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_studies_rwd_conductor_studies

> <RwdConductorStudiesResponse> rwd_studies_rwd_conductor_studies(id, opts)

Real-world data study conductor studies (HMA-EMA catalogue)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdStudiesApi.new
id = 900000301 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Real-world data study conductor studies (HMA-EMA catalogue)
  result = api_instance.rwd_studies_rwd_conductor_studies(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_rwd_conductor_studies: #{e}"
end
```

#### Using the rwd_studies_rwd_conductor_studies_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdConductorStudiesResponse>, Integer, Hash)> rwd_studies_rwd_conductor_studies_with_http_info(id, opts)

```ruby
begin
  # Real-world data study conductor studies (HMA-EMA catalogue)
  data, status_code, headers = api_instance.rwd_studies_rwd_conductor_studies_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdConductorStudiesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_rwd_conductor_studies_with_http_info: #{e}"
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

[**RwdConductorStudiesResponse**](RwdConductorStudiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_studies_rwd_study_conditions

> <RwdStudyConditionsResponse> rwd_studies_rwd_study_conditions(id, opts)

Registered real-world data study conditions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdStudiesApi.new
id = 900000301 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Registered real-world data study conditions
  result = api_instance.rwd_studies_rwd_study_conditions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_rwd_study_conditions: #{e}"
end
```

#### Using the rwd_studies_rwd_study_conditions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdStudyConditionsResponse>, Integer, Hash)> rwd_studies_rwd_study_conditions_with_http_info(id, opts)

```ruby
begin
  # Registered real-world data study conditions
  data, status_code, headers = api_instance.rwd_studies_rwd_study_conditions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdStudyConditionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_rwd_study_conditions_with_http_info: #{e}"
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

[**RwdStudyConditionsResponse**](RwdStudyConditionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_studies_rwd_study_data_sources

> <RwdStudyDataSourcesResponse> rwd_studies_rwd_study_data_sources(id, opts)

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

api_instance = PharmaObservatory::RwdStudiesApi.new
id = 900000301 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Registered real-world data study data sources
  result = api_instance.rwd_studies_rwd_study_data_sources(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_rwd_study_data_sources: #{e}"
end
```

#### Using the rwd_studies_rwd_study_data_sources_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdStudyDataSourcesResponse>, Integer, Hash)> rwd_studies_rwd_study_data_sources_with_http_info(id, opts)

```ruby
begin
  # Registered real-world data study data sources
  data, status_code, headers = api_instance.rwd_studies_rwd_study_data_sources_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdStudyDataSourcesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_rwd_study_data_sources_with_http_info: #{e}"
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


## rwd_studies_rwd_study_drugs

> <RwdStudyDrugsResponse> rwd_studies_rwd_study_drugs(id, opts)

Registered real-world data study drugs

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdStudiesApi.new
id = 900000301 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Registered real-world data study drugs
  result = api_instance.rwd_studies_rwd_study_drugs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_rwd_study_drugs: #{e}"
end
```

#### Using the rwd_studies_rwd_study_drugs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdStudyDrugsResponse>, Integer, Hash)> rwd_studies_rwd_study_drugs_with_http_info(id, opts)

```ruby
begin
  # Registered real-world data study drugs
  data, status_code, headers = api_instance.rwd_studies_rwd_study_drugs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdStudyDrugsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_rwd_study_drugs_with_http_info: #{e}"
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

[**RwdStudyDrugsResponse**](RwdStudyDrugsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_studies_rwd_study_milestones

> <RwdStudyMilestonesResponse> rwd_studies_rwd_study_milestones(id, opts)

Registered real-world data study milestones

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdStudiesApi.new
id = 900000301 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Registered real-world data study milestones
  result = api_instance.rwd_studies_rwd_study_milestones(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_rwd_study_milestones: #{e}"
end
```

#### Using the rwd_studies_rwd_study_milestones_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdStudyMilestonesResponse>, Integer, Hash)> rwd_studies_rwd_study_milestones_with_http_info(id, opts)

```ruby
begin
  # Registered real-world data study milestones
  data, status_code, headers = api_instance.rwd_studies_rwd_study_milestones_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdStudyMilestonesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdStudiesApi->rwd_studies_rwd_study_milestones_with_http_info: #{e}"
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

[**RwdStudyMilestonesResponse**](RwdStudyMilestonesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

