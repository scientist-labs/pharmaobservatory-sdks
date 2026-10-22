# PharmaObservatory::HospitalsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**hospitals_get**](HospitalsApi.md#hospitals_get) | **GET** /api/v1/hospitals/{id} | One hospital |
| [**hospitals_hospital_beds_yearly**](HospitalsApi.md#hospitals_hospital_beds_yearly) | **GET** /api/v1/hospitals/{id}/hospital_beds_yearly | Hospital beds by year |
| [**hospitals_hospital_facility_changes**](HospitalsApi.md#hospitals_hospital_facility_changes) | **GET** /api/v1/hospitals/{id}/hospital_facility_changes | Hospital facility changes |
| [**hospitals_list**](HospitalsApi.md#hospitals_list) | **GET** /api/v1/hospitals | List hospitals |
| [**hospitals_site_hospitals**](HospitalsApi.md#hospitals_site_hospitals) | **GET** /api/v1/hospitals/{id}/site_hospitals | Site hospitals |


## hospitals_get

> <HospitalsRecord> hospitals_get(id)

One hospital

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::HospitalsApi.new
id = 12355 # Integer | 

begin
  # One hospital
  result = api_instance.hospitals_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling HospitalsApi->hospitals_get: #{e}"
end
```

#### Using the hospitals_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<HospitalsRecord>, Integer, Hash)> hospitals_get_with_http_info(id)

```ruby
begin
  # One hospital
  data, status_code, headers = api_instance.hospitals_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <HospitalsRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling HospitalsApi->hospitals_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |

### Return type

[**HospitalsRecord**](HospitalsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## hospitals_hospital_beds_yearly

> <HospitalBedsYearlyResponse> hospitals_hospital_beds_yearly(id, opts)

Hospital beds by year

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::HospitalsApi.new
id = 12355 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Hospital beds by year
  result = api_instance.hospitals_hospital_beds_yearly(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling HospitalsApi->hospitals_hospital_beds_yearly: #{e}"
end
```

#### Using the hospitals_hospital_beds_yearly_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<HospitalBedsYearlyResponse>, Integer, Hash)> hospitals_hospital_beds_yearly_with_http_info(id, opts)

```ruby
begin
  # Hospital beds by year
  data, status_code, headers = api_instance.hospitals_hospital_beds_yearly_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <HospitalBedsYearlyResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling HospitalsApi->hospitals_hospital_beds_yearly_with_http_info: #{e}"
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

[**HospitalBedsYearlyResponse**](HospitalBedsYearlyResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## hospitals_hospital_facility_changes

> <HospitalFacilityChangesResponse> hospitals_hospital_facility_changes(id, opts)

Hospital facility changes

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::HospitalsApi.new
id = 12355 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Hospital facility changes
  result = api_instance.hospitals_hospital_facility_changes(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling HospitalsApi->hospitals_hospital_facility_changes: #{e}"
end
```

#### Using the hospitals_hospital_facility_changes_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<HospitalFacilityChangesResponse>, Integer, Hash)> hospitals_hospital_facility_changes_with_http_info(id, opts)

```ruby
begin
  # Hospital facility changes
  data, status_code, headers = api_instance.hospitals_hospital_facility_changes_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <HospitalFacilityChangesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling HospitalsApi->hospitals_hospital_facility_changes_with_http_info: #{e}"
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

[**HospitalFacilityChangesResponse**](HospitalFacilityChangesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## hospitals_list

> <HospitalsResponse> hospitals_list(opts)

List hospitals

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::HospitalsApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List hospitals
  result = api_instance.hospitals_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling HospitalsApi->hospitals_list: #{e}"
end
```

#### Using the hospitals_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<HospitalsResponse>, Integer, Hash)> hospitals_list_with_http_info(opts)

```ruby
begin
  # List hospitals
  data, status_code, headers = api_instance.hospitals_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <HospitalsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling HospitalsApi->hospitals_list_with_http_info: #{e}"
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

[**HospitalsResponse**](HospitalsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## hospitals_site_hospitals

> <SiteHospitalsResponse> hospitals_site_hospitals(id, opts)

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

api_instance = PharmaObservatory::HospitalsApi.new
id = 12355 # Integer | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Site hospitals
  result = api_instance.hospitals_site_hospitals(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling HospitalsApi->hospitals_site_hospitals: #{e}"
end
```

#### Using the hospitals_site_hospitals_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SiteHospitalsResponse>, Integer, Hash)> hospitals_site_hospitals_with_http_info(id, opts)

```ruby
begin
  # Site hospitals
  data, status_code, headers = api_instance.hospitals_site_hospitals_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SiteHospitalsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling HospitalsApi->hospitals_site_hospitals_with_http_info: #{e}"
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

[**SiteHospitalsResponse**](SiteHospitalsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

