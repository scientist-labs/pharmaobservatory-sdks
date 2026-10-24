# PharmaObservatory::InvestigatorsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**investigators_drug_investigators**](InvestigatorsApi.md#investigators_drug_investigators) | **GET** /api/v1/investigators/{id}/drug_investigators | Investigators on a drug&#39;s trials |
| [**investigators_get**](InvestigatorsApi.md#investigators_get) | **GET** /api/v1/investigators/{id} | One investigator |
| [**investigators_investigator_conditions**](InvestigatorsApi.md#investigators_investigator_conditions) | **GET** /api/v1/investigators/{id}/investigator_conditions | Investigator conditions |
| [**investigators_investigator_debarments**](InvestigatorsApi.md#investigators_investigator_debarments) | **GET** /api/v1/investigators/{id}/investigator_debarments | Investigator FDA debarments |
| [**investigators_investigator_hospitals**](InvestigatorsApi.md#investigators_investigator_hospitals) | **GET** /api/v1/investigators/{id}/investigator_hospitals | Investigator hospital affiliations |
| [**investigators_investigator_inspections**](InvestigatorsApi.md#investigators_investigator_inspections) | **GET** /api/v1/investigators/{id}/investigator_inspections | Investigator FDA BIMO inspections |
| [**investigators_investigator_publications**](InvestigatorsApi.md#investigators_investigator_publications) | **GET** /api/v1/investigators/{id}/investigator_publications | Investigator publications |
| [**investigators_investigator_topics**](InvestigatorsApi.md#investigators_investigator_topics) | **GET** /api/v1/investigators/{id}/investigator_topics | Investigator expertise topics |
| [**investigators_investigator_trials**](InvestigatorsApi.md#investigators_investigator_trials) | **GET** /api/v1/investigators/{id}/investigator_trials | Investigator trials |
| [**investigators_list**](InvestigatorsApi.md#investigators_list) | **GET** /api/v1/investigators | List investigators |
| [**investigators_rare_disease_investigators**](InvestigatorsApi.md#investigators_rare_disease_investigators) | **GET** /api/v1/investigators/{id}/rare_disease_investigators | Rare disease investigators |
| [**investigators_site_investigators**](InvestigatorsApi.md#investigators_site_investigators) | **GET** /api/v1/investigators/{id}/site_investigators | Site investigators |
| [**investigators_target_investigators**](InvestigatorsApi.md#investigators_target_investigators) | **GET** /api/v1/investigators/{id}/target_investigators | Investigators on a target&#39;s trials |
| [**investigators_trial_kols**](InvestigatorsApi.md#investigators_trial_kols) | **GET** /api/v1/investigators/{id}/trial_kols | Trial investigators |


## investigators_drug_investigators

> <DrugInvestigatorsResponse> investigators_drug_investigators(id, opts)

Investigators on a drug's trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investigators on a drug's trials
  result = api_instance.investigators_drug_investigators(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_drug_investigators: #{e}"
end
```

#### Using the investigators_drug_investigators_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugInvestigatorsResponse>, Integer, Hash)> investigators_drug_investigators_with_http_info(id, opts)

```ruby
begin
  # Investigators on a drug's trials
  data, status_code, headers = api_instance.investigators_drug_investigators_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugInvestigatorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_drug_investigators_with_http_info: #{e}"
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

[**DrugInvestigatorsResponse**](DrugInvestigatorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_get

> <InvestigatorsRecord> investigators_get(id)

One investigator

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 

begin
  # One investigator
  result = api_instance.investigators_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_get: #{e}"
end
```

#### Using the investigators_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestigatorsRecord>, Integer, Hash)> investigators_get_with_http_info(id)

```ruby
begin
  # One investigator
  data, status_code, headers = api_instance.investigators_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestigatorsRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**InvestigatorsRecord**](InvestigatorsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_investigator_conditions

> <InvestigatorConditionsResponse> investigators_investigator_conditions(id, opts)

Investigator conditions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investigator conditions
  result = api_instance.investigators_investigator_conditions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_conditions: #{e}"
end
```

#### Using the investigators_investigator_conditions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestigatorConditionsResponse>, Integer, Hash)> investigators_investigator_conditions_with_http_info(id, opts)

```ruby
begin
  # Investigator conditions
  data, status_code, headers = api_instance.investigators_investigator_conditions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestigatorConditionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_conditions_with_http_info: #{e}"
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

[**InvestigatorConditionsResponse**](InvestigatorConditionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_investigator_debarments

> <InvestigatorDebarmentsResponse> investigators_investigator_debarments(id, opts)

Investigator FDA debarments

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investigator FDA debarments
  result = api_instance.investigators_investigator_debarments(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_debarments: #{e}"
end
```

#### Using the investigators_investigator_debarments_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestigatorDebarmentsResponse>, Integer, Hash)> investigators_investigator_debarments_with_http_info(id, opts)

```ruby
begin
  # Investigator FDA debarments
  data, status_code, headers = api_instance.investigators_investigator_debarments_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestigatorDebarmentsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_debarments_with_http_info: #{e}"
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

[**InvestigatorDebarmentsResponse**](InvestigatorDebarmentsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_investigator_hospitals

> <InvestigatorHospitalsResponse> investigators_investigator_hospitals(id, opts)

Investigator hospital affiliations

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investigator hospital affiliations
  result = api_instance.investigators_investigator_hospitals(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_hospitals: #{e}"
end
```

#### Using the investigators_investigator_hospitals_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestigatorHospitalsResponse>, Integer, Hash)> investigators_investigator_hospitals_with_http_info(id, opts)

```ruby
begin
  # Investigator hospital affiliations
  data, status_code, headers = api_instance.investigators_investigator_hospitals_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestigatorHospitalsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_hospitals_with_http_info: #{e}"
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

[**InvestigatorHospitalsResponse**](InvestigatorHospitalsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_investigator_inspections

> <InvestigatorInspectionsResponse> investigators_investigator_inspections(id, opts)

Investigator FDA BIMO inspections

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investigator FDA BIMO inspections
  result = api_instance.investigators_investigator_inspections(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_inspections: #{e}"
end
```

#### Using the investigators_investigator_inspections_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestigatorInspectionsResponse>, Integer, Hash)> investigators_investigator_inspections_with_http_info(id, opts)

```ruby
begin
  # Investigator FDA BIMO inspections
  data, status_code, headers = api_instance.investigators_investigator_inspections_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestigatorInspectionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_inspections_with_http_info: #{e}"
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

[**InvestigatorInspectionsResponse**](InvestigatorInspectionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_investigator_publications

> <InvestigatorPublicationsResponse> investigators_investigator_publications(id, opts)

Investigator publications

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investigator publications
  result = api_instance.investigators_investigator_publications(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_publications: #{e}"
end
```

#### Using the investigators_investigator_publications_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestigatorPublicationsResponse>, Integer, Hash)> investigators_investigator_publications_with_http_info(id, opts)

```ruby
begin
  # Investigator publications
  data, status_code, headers = api_instance.investigators_investigator_publications_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestigatorPublicationsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_publications_with_http_info: #{e}"
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

[**InvestigatorPublicationsResponse**](InvestigatorPublicationsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_investigator_topics

> <InvestigatorTopicsResponse> investigators_investigator_topics(id, opts)

Investigator expertise topics

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investigator expertise topics
  result = api_instance.investigators_investigator_topics(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_topics: #{e}"
end
```

#### Using the investigators_investigator_topics_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestigatorTopicsResponse>, Integer, Hash)> investigators_investigator_topics_with_http_info(id, opts)

```ruby
begin
  # Investigator expertise topics
  data, status_code, headers = api_instance.investigators_investigator_topics_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestigatorTopicsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_topics_with_http_info: #{e}"
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

[**InvestigatorTopicsResponse**](InvestigatorTopicsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_investigator_trials

> <InvestigatorTrialsResponse> investigators_investigator_trials(id, opts)

Investigator trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investigator trials
  result = api_instance.investigators_investigator_trials(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_trials: #{e}"
end
```

#### Using the investigators_investigator_trials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestigatorTrialsResponse>, Integer, Hash)> investigators_investigator_trials_with_http_info(id, opts)

```ruby
begin
  # Investigator trials
  data, status_code, headers = api_instance.investigators_investigator_trials_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestigatorTrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_investigator_trials_with_http_info: #{e}"
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

[**InvestigatorTrialsResponse**](InvestigatorTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_list

> <InvestigatorsResponse> investigators_list(opts)

List investigators

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List investigators
  result = api_instance.investigators_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_list: #{e}"
end
```

#### Using the investigators_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestigatorsResponse>, Integer, Hash)> investigators_list_with_http_info(opts)

```ruby
begin
  # List investigators
  data, status_code, headers = api_instance.investigators_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestigatorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_list_with_http_info: #{e}"
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

[**InvestigatorsResponse**](InvestigatorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_rare_disease_investigators

> <RareDiseaseInvestigatorsResponse> investigators_rare_disease_investigators(id, opts)

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

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Rare disease investigators
  result = api_instance.investigators_rare_disease_investigators(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_rare_disease_investigators: #{e}"
end
```

#### Using the investigators_rare_disease_investigators_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RareDiseaseInvestigatorsResponse>, Integer, Hash)> investigators_rare_disease_investigators_with_http_info(id, opts)

```ruby
begin
  # Rare disease investigators
  data, status_code, headers = api_instance.investigators_rare_disease_investigators_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RareDiseaseInvestigatorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_rare_disease_investigators_with_http_info: #{e}"
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

[**RareDiseaseInvestigatorsResponse**](RareDiseaseInvestigatorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_site_investigators

> <SiteInvestigatorsResponse> investigators_site_investigators(id, opts)

Site investigators

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Site investigators
  result = api_instance.investigators_site_investigators(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_site_investigators: #{e}"
end
```

#### Using the investigators_site_investigators_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SiteInvestigatorsResponse>, Integer, Hash)> investigators_site_investigators_with_http_info(id, opts)

```ruby
begin
  # Site investigators
  data, status_code, headers = api_instance.investigators_site_investigators_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SiteInvestigatorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_site_investigators_with_http_info: #{e}"
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

[**SiteInvestigatorsResponse**](SiteInvestigatorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_target_investigators

> <TargetInvestigatorsResponse> investigators_target_investigators(id, opts)

Investigators on a target's trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investigators on a target's trials
  result = api_instance.investigators_target_investigators(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_target_investigators: #{e}"
end
```

#### Using the investigators_target_investigators_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetInvestigatorsResponse>, Integer, Hash)> investigators_target_investigators_with_http_info(id, opts)

```ruby
begin
  # Investigators on a target's trials
  data, status_code, headers = api_instance.investigators_target_investigators_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetInvestigatorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_target_investigators_with_http_info: #{e}"
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

[**TargetInvestigatorsResponse**](TargetInvestigatorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigators_trial_kols

> <TrialKolsResponse> investigators_trial_kols(id, opts)

Trial investigators

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestigatorsApi.new
id = 'npi:9900000381' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial investigators
  result = api_instance.investigators_trial_kols(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_trial_kols: #{e}"
end
```

#### Using the investigators_trial_kols_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialKolsResponse>, Integer, Hash)> investigators_trial_kols_with_http_info(id, opts)

```ruby
begin
  # Trial investigators
  data, status_code, headers = api_instance.investigators_trial_kols_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialKolsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestigatorsApi->investigators_trial_kols_with_http_info: #{e}"
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

[**TrialKolsResponse**](TrialKolsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

