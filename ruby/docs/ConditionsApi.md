# PharmaObservatory::ConditionsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**conditions_company_programs**](ConditionsApi.md#conditions_company_programs) | **GET** /api/v1/conditions/{id}/company_programs | Company programs |
| [**conditions_condition_adverse_events**](ConditionsApi.md#conditions_condition_adverse_events) | **GET** /api/v1/conditions/{id}/condition_adverse_events | Condition adverse events |
| [**conditions_condition_burden**](ConditionsApi.md#conditions_condition_burden) | **GET** /api/v1/conditions/{id}/condition_burden | Condition burden |
| [**conditions_condition_delivery**](ConditionsApi.md#conditions_condition_delivery) | **GET** /api/v1/conditions/{id}/condition_delivery | Condition delivery |
| [**conditions_condition_drugs**](ConditionsApi.md#conditions_condition_drugs) | **GET** /api/v1/conditions/{id}/condition_drugs | Condition drugs |
| [**conditions_condition_prevalence**](ConditionsApi.md#conditions_condition_prevalence) | **GET** /api/v1/conditions/{id}/condition_prevalence | Condition prevalence |
| [**conditions_condition_publications**](ConditionsApi.md#conditions_condition_publications) | **GET** /api/v1/conditions/{id}/condition_publications | Condition publications |
| [**conditions_condition_sites**](ConditionsApi.md#conditions_condition_sites) | **GET** /api/v1/conditions/{id}/condition_sites | Condition sites |
| [**conditions_condition_trials**](ConditionsApi.md#conditions_condition_trials) | **GET** /api/v1/conditions/{id}/condition_trials | Condition trials |
| [**conditions_get**](ConditionsApi.md#conditions_get) | **GET** /api/v1/conditions/{id} | One condition |
| [**conditions_list**](ConditionsApi.md#conditions_list) | **GET** /api/v1/conditions | List conditions |
| [**conditions_pdufa_dates**](ConditionsApi.md#conditions_pdufa_dates) | **GET** /api/v1/conditions/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts |


## conditions_company_programs

> <CompanyProgramsResponse> conditions_company_programs(id, opts)

Company programs

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Company programs
  result = api_instance.conditions_company_programs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_company_programs: #{e}"
end
```

#### Using the conditions_company_programs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CompanyProgramsResponse>, Integer, Hash)> conditions_company_programs_with_http_info(id, opts)

```ruby
begin
  # Company programs
  data, status_code, headers = api_instance.conditions_company_programs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CompanyProgramsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_company_programs_with_http_info: #{e}"
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

[**CompanyProgramsResponse**](CompanyProgramsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_condition_adverse_events

> <ConditionAdverseEventsResponse> conditions_condition_adverse_events(id, opts)

Condition adverse events

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Condition adverse events
  result = api_instance.conditions_condition_adverse_events(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_adverse_events: #{e}"
end
```

#### Using the conditions_condition_adverse_events_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ConditionAdverseEventsResponse>, Integer, Hash)> conditions_condition_adverse_events_with_http_info(id, opts)

```ruby
begin
  # Condition adverse events
  data, status_code, headers = api_instance.conditions_condition_adverse_events_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ConditionAdverseEventsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_adverse_events_with_http_info: #{e}"
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

[**ConditionAdverseEventsResponse**](ConditionAdverseEventsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_condition_burden

> <ConditionBurdenResponse> conditions_condition_burden(id, opts)

Condition burden

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Condition burden
  result = api_instance.conditions_condition_burden(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_burden: #{e}"
end
```

#### Using the conditions_condition_burden_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ConditionBurdenResponse>, Integer, Hash)> conditions_condition_burden_with_http_info(id, opts)

```ruby
begin
  # Condition burden
  data, status_code, headers = api_instance.conditions_condition_burden_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ConditionBurdenResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_burden_with_http_info: #{e}"
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

[**ConditionBurdenResponse**](ConditionBurdenResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_condition_delivery

> <ConditionDeliveryResponse> conditions_condition_delivery(id, opts)

Condition delivery

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Condition delivery
  result = api_instance.conditions_condition_delivery(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_delivery: #{e}"
end
```

#### Using the conditions_condition_delivery_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ConditionDeliveryResponse>, Integer, Hash)> conditions_condition_delivery_with_http_info(id, opts)

```ruby
begin
  # Condition delivery
  data, status_code, headers = api_instance.conditions_condition_delivery_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ConditionDeliveryResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_delivery_with_http_info: #{e}"
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

[**ConditionDeliveryResponse**](ConditionDeliveryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_condition_drugs

> <ConditionDrugsResponse> conditions_condition_drugs(id, opts)

Condition drugs

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Condition drugs
  result = api_instance.conditions_condition_drugs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_drugs: #{e}"
end
```

#### Using the conditions_condition_drugs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ConditionDrugsResponse>, Integer, Hash)> conditions_condition_drugs_with_http_info(id, opts)

```ruby
begin
  # Condition drugs
  data, status_code, headers = api_instance.conditions_condition_drugs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ConditionDrugsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_drugs_with_http_info: #{e}"
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

[**ConditionDrugsResponse**](ConditionDrugsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_condition_prevalence

> <ConditionPrevalenceResponse> conditions_condition_prevalence(id, opts)

Condition prevalence

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Condition prevalence
  result = api_instance.conditions_condition_prevalence(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_prevalence: #{e}"
end
```

#### Using the conditions_condition_prevalence_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ConditionPrevalenceResponse>, Integer, Hash)> conditions_condition_prevalence_with_http_info(id, opts)

```ruby
begin
  # Condition prevalence
  data, status_code, headers = api_instance.conditions_condition_prevalence_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ConditionPrevalenceResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_prevalence_with_http_info: #{e}"
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

[**ConditionPrevalenceResponse**](ConditionPrevalenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_condition_publications

> <ConditionPublicationsResponse> conditions_condition_publications(id, opts)

Condition publications

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Condition publications
  result = api_instance.conditions_condition_publications(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_publications: #{e}"
end
```

#### Using the conditions_condition_publications_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ConditionPublicationsResponse>, Integer, Hash)> conditions_condition_publications_with_http_info(id, opts)

```ruby
begin
  # Condition publications
  data, status_code, headers = api_instance.conditions_condition_publications_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ConditionPublicationsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_publications_with_http_info: #{e}"
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

[**ConditionPublicationsResponse**](ConditionPublicationsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_condition_sites

> <ConditionSitesResponse> conditions_condition_sites(id, opts)

Condition sites

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Condition sites
  result = api_instance.conditions_condition_sites(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_sites: #{e}"
end
```

#### Using the conditions_condition_sites_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ConditionSitesResponse>, Integer, Hash)> conditions_condition_sites_with_http_info(id, opts)

```ruby
begin
  # Condition sites
  data, status_code, headers = api_instance.conditions_condition_sites_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ConditionSitesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_sites_with_http_info: #{e}"
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

[**ConditionSitesResponse**](ConditionSitesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_condition_trials

> <ConditionTrialsResponse> conditions_condition_trials(id, opts)

Condition trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Condition trials
  result = api_instance.conditions_condition_trials(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_trials: #{e}"
end
```

#### Using the conditions_condition_trials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ConditionTrialsResponse>, Integer, Hash)> conditions_condition_trials_with_http_info(id, opts)

```ruby
begin
  # Condition trials
  data, status_code, headers = api_instance.conditions_condition_trials_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ConditionTrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_condition_trials_with_http_info: #{e}"
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

[**ConditionTrialsResponse**](ConditionTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_get

> <ConditionsRecord> conditions_get(id)

One condition

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 

begin
  # One condition
  result = api_instance.conditions_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_get: #{e}"
end
```

#### Using the conditions_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ConditionsRecord>, Integer, Hash)> conditions_get_with_http_info(id)

```ruby
begin
  # One condition
  data, status_code, headers = api_instance.conditions_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ConditionsRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**ConditionsRecord**](ConditionsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_list

> <ConditionsResponse> conditions_list(opts)

List conditions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List conditions
  result = api_instance.conditions_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_list: #{e}"
end
```

#### Using the conditions_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ConditionsResponse>, Integer, Hash)> conditions_list_with_http_info(opts)

```ruby
begin
  # List conditions
  data, status_code, headers = api_instance.conditions_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ConditionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_list_with_http_info: #{e}"
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

[**ConditionsResponse**](ConditionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## conditions_pdufa_dates

> <PdufaDatesResponse> conditions_pdufa_dates(id, opts)

FDA decision dates (PDUFA) and catalysts

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ConditionsApi.new
id = 'D001523' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA decision dates (PDUFA) and catalysts
  result = api_instance.conditions_pdufa_dates(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_pdufa_dates: #{e}"
end
```

#### Using the conditions_pdufa_dates_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PdufaDatesResponse>, Integer, Hash)> conditions_pdufa_dates_with_http_info(id, opts)

```ruby
begin
  # FDA decision dates (PDUFA) and catalysts
  data, status_code, headers = api_instance.conditions_pdufa_dates_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PdufaDatesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ConditionsApi->conditions_pdufa_dates_with_http_info: #{e}"
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

[**PdufaDatesResponse**](PdufaDatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

