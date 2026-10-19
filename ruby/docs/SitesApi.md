# PharmaObservatory::SitesApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**sites_get**](SitesApi.md#sites_get) | **GET** /api/v1/sites/{id} | One site |
| [**sites_list**](SitesApi.md#sites_list) | **GET** /api/v1/sites | List sites |
| [**sites_site_beds_yearly**](SitesApi.md#sites_site_beds_yearly) | **GET** /api/v1/sites/{id}/site_beds_yearly | Site beds by year |
| [**sites_site_cro_trials**](SitesApi.md#sites_site_cro_trials) | **GET** /api/v1/sites/{id}/site_cro_trials | Site CRO trials |
| [**sites_site_hospitals**](SitesApi.md#sites_site_hospitals) | **GET** /api/v1/sites/{id}/site_hospitals | Site hospitals |
| [**sites_site_investigators**](SitesApi.md#sites_site_investigators) | **GET** /api/v1/sites/{id}/site_investigators | Site investigators |
| [**sites_site_trial_conditions**](SitesApi.md#sites_site_trial_conditions) | **GET** /api/v1/sites/{id}/site_trial_conditions | Site trial conditions |
| [**sites_site_trial_drugs**](SitesApi.md#sites_site_trial_drugs) | **GET** /api/v1/sites/{id}/site_trial_drugs | Site trial drugs |
| [**sites_site_trials**](SitesApi.md#sites_site_trials) | **GET** /api/v1/sites/{id}/site_trials | Site trials |


## sites_get

> <SitesRecord> sites_get(id)

One site

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SitesApi.new
id = '23c3485b8d2a74c67877558364ce64c9' # String | 

begin
  # One site
  result = api_instance.sites_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_get: #{e}"
end
```

#### Using the sites_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SitesRecord>, Integer, Hash)> sites_get_with_http_info(id)

```ruby
begin
  # One site
  data, status_code, headers = api_instance.sites_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SitesRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**SitesRecord**](SitesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## sites_list

> <SitesResponse> sites_list(opts)

List sites

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SitesApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List sites
  result = api_instance.sites_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_list: #{e}"
end
```

#### Using the sites_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SitesResponse>, Integer, Hash)> sites_list_with_http_info(opts)

```ruby
begin
  # List sites
  data, status_code, headers = api_instance.sites_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SitesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_list_with_http_info: #{e}"
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

[**SitesResponse**](SitesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## sites_site_beds_yearly

> <SiteBedsYearlyResponse> sites_site_beds_yearly(id, opts)

Site beds by year

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SitesApi.new
id = '23c3485b8d2a74c67877558364ce64c9' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Site beds by year
  result = api_instance.sites_site_beds_yearly(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_beds_yearly: #{e}"
end
```

#### Using the sites_site_beds_yearly_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SiteBedsYearlyResponse>, Integer, Hash)> sites_site_beds_yearly_with_http_info(id, opts)

```ruby
begin
  # Site beds by year
  data, status_code, headers = api_instance.sites_site_beds_yearly_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SiteBedsYearlyResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_beds_yearly_with_http_info: #{e}"
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

[**SiteBedsYearlyResponse**](SiteBedsYearlyResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## sites_site_cro_trials

> <SiteCroTrialsResponse> sites_site_cro_trials(id, opts)

Site CRO trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SitesApi.new
id = '23c3485b8d2a74c67877558364ce64c9' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Site CRO trials
  result = api_instance.sites_site_cro_trials(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_cro_trials: #{e}"
end
```

#### Using the sites_site_cro_trials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SiteCroTrialsResponse>, Integer, Hash)> sites_site_cro_trials_with_http_info(id, opts)

```ruby
begin
  # Site CRO trials
  data, status_code, headers = api_instance.sites_site_cro_trials_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SiteCroTrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_cro_trials_with_http_info: #{e}"
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

[**SiteCroTrialsResponse**](SiteCroTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## sites_site_hospitals

> <SiteHospitalsResponse> sites_site_hospitals(id, opts)

Site hospitals

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SitesApi.new
id = '23c3485b8d2a74c67877558364ce64c9' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Site hospitals
  result = api_instance.sites_site_hospitals(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_hospitals: #{e}"
end
```

#### Using the sites_site_hospitals_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SiteHospitalsResponse>, Integer, Hash)> sites_site_hospitals_with_http_info(id, opts)

```ruby
begin
  # Site hospitals
  data, status_code, headers = api_instance.sites_site_hospitals_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SiteHospitalsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_hospitals_with_http_info: #{e}"
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

[**SiteHospitalsResponse**](SiteHospitalsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## sites_site_investigators

> <SiteInvestigatorsResponse> sites_site_investigators(id, opts)

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

api_instance = PharmaObservatory::SitesApi.new
id = '23c3485b8d2a74c67877558364ce64c9' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Site investigators
  result = api_instance.sites_site_investigators(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_investigators: #{e}"
end
```

#### Using the sites_site_investigators_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SiteInvestigatorsResponse>, Integer, Hash)> sites_site_investigators_with_http_info(id, opts)

```ruby
begin
  # Site investigators
  data, status_code, headers = api_instance.sites_site_investigators_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SiteInvestigatorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_investigators_with_http_info: #{e}"
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


## sites_site_trial_conditions

> <SiteTrialConditionsResponse> sites_site_trial_conditions(id, opts)

Site trial conditions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SitesApi.new
id = '23c3485b8d2a74c67877558364ce64c9' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Site trial conditions
  result = api_instance.sites_site_trial_conditions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_trial_conditions: #{e}"
end
```

#### Using the sites_site_trial_conditions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SiteTrialConditionsResponse>, Integer, Hash)> sites_site_trial_conditions_with_http_info(id, opts)

```ruby
begin
  # Site trial conditions
  data, status_code, headers = api_instance.sites_site_trial_conditions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SiteTrialConditionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_trial_conditions_with_http_info: #{e}"
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

[**SiteTrialConditionsResponse**](SiteTrialConditionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## sites_site_trial_drugs

> <SiteTrialDrugsResponse> sites_site_trial_drugs(id, opts)

Site trial drugs

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SitesApi.new
id = '23c3485b8d2a74c67877558364ce64c9' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Site trial drugs
  result = api_instance.sites_site_trial_drugs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_trial_drugs: #{e}"
end
```

#### Using the sites_site_trial_drugs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SiteTrialDrugsResponse>, Integer, Hash)> sites_site_trial_drugs_with_http_info(id, opts)

```ruby
begin
  # Site trial drugs
  data, status_code, headers = api_instance.sites_site_trial_drugs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SiteTrialDrugsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_trial_drugs_with_http_info: #{e}"
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

[**SiteTrialDrugsResponse**](SiteTrialDrugsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## sites_site_trials

> <SiteTrialsResponse> sites_site_trials(id, opts)

Site trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::SitesApi.new
id = '23c3485b8d2a74c67877558364ce64c9' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Site trials
  result = api_instance.sites_site_trials(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_trials: #{e}"
end
```

#### Using the sites_site_trials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SiteTrialsResponse>, Integer, Hash)> sites_site_trials_with_http_info(id, opts)

```ruby
begin
  # Site trials
  data, status_code, headers = api_instance.sites_site_trials_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SiteTrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling SitesApi->sites_site_trials_with_http_info: #{e}"
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

[**SiteTrialsResponse**](SiteTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

