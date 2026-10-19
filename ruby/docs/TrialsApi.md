# PharmaObservatory::TrialsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**trials_get**](TrialsApi.md#trials_get) | **GET** /api/v1/trials/{id} | One trial |
| [**trials_list**](TrialsApi.md#trials_list) | **GET** /api/v1/trials | List trials |
| [**trials_sponsor_trials**](TrialsApi.md#trials_sponsor_trials) | **GET** /api/v1/trials/{id}/sponsor_trials | Sponsor trials |
| [**trials_trial_adverse_events**](TrialsApi.md#trials_trial_adverse_events) | **GET** /api/v1/trials/{id}/trial_adverse_events | Trial adverse events |
| [**trials_trial_changes**](TrialsApi.md#trials_trial_changes) | **GET** /api/v1/trials/{id}/trial_changes | Trial changes |
| [**trials_trial_delivery**](TrialsApi.md#trials_trial_delivery) | **GET** /api/v1/trials/{id}/trial_delivery | Trial delivery |
| [**trials_trial_demographics**](TrialsApi.md#trials_trial_demographics) | **GET** /api/v1/trials/{id}/trial_demographics | Trial demographics |
| [**trials_trial_drugs**](TrialsApi.md#trials_trial_drugs) | **GET** /api/v1/trials/{id}/trial_drugs | Trial drugs |
| [**trials_trial_enrollment_history**](TrialsApi.md#trials_trial_enrollment_history) | **GET** /api/v1/trials/{id}/trial_enrollment_history | Trial enrollment history |
| [**trials_trial_enrollment_rate**](TrialsApi.md#trials_trial_enrollment_rate) | **GET** /api/v1/trials/{id}/trial_enrollment_rate | Trial enrollment rate |
| [**trials_trial_extracted_endpoints**](TrialsApi.md#trials_trial_extracted_endpoints) | **GET** /api/v1/trials/{id}/trial_extracted_endpoints | Trial extracted endpoints |
| [**trials_trial_outcome_measures**](TrialsApi.md#trials_trial_outcome_measures) | **GET** /api/v1/trials/{id}/trial_outcome_measures | Trial outcome measures |
| [**trials_trial_sites**](TrialsApi.md#trials_trial_sites) | **GET** /api/v1/trials/{id}/trial_sites | Trial sites |


## trials_get

> <TrialsRecord> trials_get(id)

One trial

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 

begin
  # One trial
  result = api_instance.trials_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_get: #{e}"
end
```

#### Using the trials_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialsRecord>, Integer, Hash)> trials_get_with_http_info(id)

```ruby
begin
  # One trial
  data, status_code, headers = api_instance.trials_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialsRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**TrialsRecord**](TrialsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_list

> <TrialsResponse> trials_list(opts)

List trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List trials
  result = api_instance.trials_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_list: #{e}"
end
```

#### Using the trials_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialsResponse>, Integer, Hash)> trials_list_with_http_info(opts)

```ruby
begin
  # List trials
  data, status_code, headers = api_instance.trials_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_list_with_http_info: #{e}"
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

[**TrialsResponse**](TrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_sponsor_trials

> <SponsorTrialsResponse> trials_sponsor_trials(id, opts)

Sponsor trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Sponsor trials
  result = api_instance.trials_sponsor_trials(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_sponsor_trials: #{e}"
end
```

#### Using the trials_sponsor_trials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SponsorTrialsResponse>, Integer, Hash)> trials_sponsor_trials_with_http_info(id, opts)

```ruby
begin
  # Sponsor trials
  data, status_code, headers = api_instance.trials_sponsor_trials_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SponsorTrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_sponsor_trials_with_http_info: #{e}"
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

[**SponsorTrialsResponse**](SponsorTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_trial_adverse_events

> <TrialAdverseEventsResponse> trials_trial_adverse_events(id, opts)

Trial adverse events

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial adverse events
  result = api_instance.trials_trial_adverse_events(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_adverse_events: #{e}"
end
```

#### Using the trials_trial_adverse_events_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialAdverseEventsResponse>, Integer, Hash)> trials_trial_adverse_events_with_http_info(id, opts)

```ruby
begin
  # Trial adverse events
  data, status_code, headers = api_instance.trials_trial_adverse_events_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialAdverseEventsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_adverse_events_with_http_info: #{e}"
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

[**TrialAdverseEventsResponse**](TrialAdverseEventsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_trial_changes

> <TrialChangesResponse> trials_trial_changes(id, opts)

Trial changes

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial changes
  result = api_instance.trials_trial_changes(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_changes: #{e}"
end
```

#### Using the trials_trial_changes_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialChangesResponse>, Integer, Hash)> trials_trial_changes_with_http_info(id, opts)

```ruby
begin
  # Trial changes
  data, status_code, headers = api_instance.trials_trial_changes_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialChangesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_changes_with_http_info: #{e}"
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

[**TrialChangesResponse**](TrialChangesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_trial_delivery

> <TrialDeliveryResponse> trials_trial_delivery(id, opts)

Trial delivery

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial delivery
  result = api_instance.trials_trial_delivery(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_delivery: #{e}"
end
```

#### Using the trials_trial_delivery_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialDeliveryResponse>, Integer, Hash)> trials_trial_delivery_with_http_info(id, opts)

```ruby
begin
  # Trial delivery
  data, status_code, headers = api_instance.trials_trial_delivery_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialDeliveryResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_delivery_with_http_info: #{e}"
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

[**TrialDeliveryResponse**](TrialDeliveryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_trial_demographics

> <TrialDemographicsResponse> trials_trial_demographics(id, opts)

Trial demographics

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial demographics
  result = api_instance.trials_trial_demographics(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_demographics: #{e}"
end
```

#### Using the trials_trial_demographics_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialDemographicsResponse>, Integer, Hash)> trials_trial_demographics_with_http_info(id, opts)

```ruby
begin
  # Trial demographics
  data, status_code, headers = api_instance.trials_trial_demographics_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialDemographicsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_demographics_with_http_info: #{e}"
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

[**TrialDemographicsResponse**](TrialDemographicsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_trial_drugs

> <TrialDrugsResponse> trials_trial_drugs(id, opts)

Trial drugs

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial drugs
  result = api_instance.trials_trial_drugs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_drugs: #{e}"
end
```

#### Using the trials_trial_drugs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialDrugsResponse>, Integer, Hash)> trials_trial_drugs_with_http_info(id, opts)

```ruby
begin
  # Trial drugs
  data, status_code, headers = api_instance.trials_trial_drugs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialDrugsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_drugs_with_http_info: #{e}"
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

[**TrialDrugsResponse**](TrialDrugsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_trial_enrollment_history

> <TrialEnrollmentHistoryResponse> trials_trial_enrollment_history(id, opts)

Trial enrollment history

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial enrollment history
  result = api_instance.trials_trial_enrollment_history(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_enrollment_history: #{e}"
end
```

#### Using the trials_trial_enrollment_history_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialEnrollmentHistoryResponse>, Integer, Hash)> trials_trial_enrollment_history_with_http_info(id, opts)

```ruby
begin
  # Trial enrollment history
  data, status_code, headers = api_instance.trials_trial_enrollment_history_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialEnrollmentHistoryResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_enrollment_history_with_http_info: #{e}"
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

[**TrialEnrollmentHistoryResponse**](TrialEnrollmentHistoryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_trial_enrollment_rate

> <TrialEnrollmentRateResponse> trials_trial_enrollment_rate(id, opts)

Trial enrollment rate

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial enrollment rate
  result = api_instance.trials_trial_enrollment_rate(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_enrollment_rate: #{e}"
end
```

#### Using the trials_trial_enrollment_rate_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialEnrollmentRateResponse>, Integer, Hash)> trials_trial_enrollment_rate_with_http_info(id, opts)

```ruby
begin
  # Trial enrollment rate
  data, status_code, headers = api_instance.trials_trial_enrollment_rate_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialEnrollmentRateResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_enrollment_rate_with_http_info: #{e}"
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

[**TrialEnrollmentRateResponse**](TrialEnrollmentRateResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_trial_extracted_endpoints

> <TrialExtractedEndpointsResponse> trials_trial_extracted_endpoints(id, opts)

Trial extracted endpoints

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial extracted endpoints
  result = api_instance.trials_trial_extracted_endpoints(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_extracted_endpoints: #{e}"
end
```

#### Using the trials_trial_extracted_endpoints_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialExtractedEndpointsResponse>, Integer, Hash)> trials_trial_extracted_endpoints_with_http_info(id, opts)

```ruby
begin
  # Trial extracted endpoints
  data, status_code, headers = api_instance.trials_trial_extracted_endpoints_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialExtractedEndpointsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_extracted_endpoints_with_http_info: #{e}"
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

[**TrialExtractedEndpointsResponse**](TrialExtractedEndpointsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_trial_outcome_measures

> <TrialOutcomeMeasuresResponse> trials_trial_outcome_measures(id, opts)

Trial outcome measures

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial outcome measures
  result = api_instance.trials_trial_outcome_measures(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_outcome_measures: #{e}"
end
```

#### Using the trials_trial_outcome_measures_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialOutcomeMeasuresResponse>, Integer, Hash)> trials_trial_outcome_measures_with_http_info(id, opts)

```ruby
begin
  # Trial outcome measures
  data, status_code, headers = api_instance.trials_trial_outcome_measures_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialOutcomeMeasuresResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_outcome_measures_with_http_info: #{e}"
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

[**TrialOutcomeMeasuresResponse**](TrialOutcomeMeasuresResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## trials_trial_sites

> <TrialSitesResponse> trials_trial_sites(id, opts)

Trial sites

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TrialsApi.new
id = 'NCT00000125' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial sites
  result = api_instance.trials_trial_sites(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_sites: #{e}"
end
```

#### Using the trials_trial_sites_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialSitesResponse>, Integer, Hash)> trials_trial_sites_with_http_info(id, opts)

```ruby
begin
  # Trial sites
  data, status_code, headers = api_instance.trials_trial_sites_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialSitesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TrialsApi->trials_trial_sites_with_http_info: #{e}"
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

[**TrialSitesResponse**](TrialSitesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

