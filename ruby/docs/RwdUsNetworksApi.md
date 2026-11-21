# PharmaObservatory::RwdUsNetworksApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**rwd_us_networks_get**](RwdUsNetworksApi.md#rwd_us_networks_get) | **GET** /api/v1/rwd_us_networks/{id} | One us real-world data network |
| [**rwd_us_networks_list**](RwdUsNetworksApi.md#rwd_us_networks_list) | **GET** /api/v1/rwd_us_networks | List us real-world data networks |
| [**rwd_us_networks_rwd_us_network_memberships**](RwdUsNetworksApi.md#rwd_us_networks_rwd_us_network_memberships) | **GET** /api/v1/rwd_us_networks/{id}/rwd_us_network_memberships | US real-world data network membership (N3C) |
| [**rwd_us_networks_rwd_us_network_projects**](RwdUsNetworksApi.md#rwd_us_networks_rwd_us_network_projects) | **GET** /api/v1/rwd_us_networks/{id}/rwd_us_network_projects | US real-world data network projects (N3C) |


## rwd_us_networks_get

> <RwdUsNetworksRecord> rwd_us_networks_get(id)

One us real-world data network

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdUsNetworksApi.new
id = 'n3c' # String | 

begin
  # One us real-world data network
  result = api_instance.rwd_us_networks_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdUsNetworksApi->rwd_us_networks_get: #{e}"
end
```

#### Using the rwd_us_networks_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdUsNetworksRecord>, Integer, Hash)> rwd_us_networks_get_with_http_info(id)

```ruby
begin
  # One us real-world data network
  data, status_code, headers = api_instance.rwd_us_networks_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdUsNetworksRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdUsNetworksApi->rwd_us_networks_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**RwdUsNetworksRecord**](RwdUsNetworksRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_us_networks_list

> <RwdUsNetworksResponse> rwd_us_networks_list(opts)

List us real-world data networks

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdUsNetworksApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List us real-world data networks
  result = api_instance.rwd_us_networks_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdUsNetworksApi->rwd_us_networks_list: #{e}"
end
```

#### Using the rwd_us_networks_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdUsNetworksResponse>, Integer, Hash)> rwd_us_networks_list_with_http_info(opts)

```ruby
begin
  # List us real-world data networks
  data, status_code, headers = api_instance.rwd_us_networks_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdUsNetworksResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdUsNetworksApi->rwd_us_networks_list_with_http_info: #{e}"
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

[**RwdUsNetworksResponse**](RwdUsNetworksResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_us_networks_rwd_us_network_memberships

> <RwdUsNetworkMembershipsResponse> rwd_us_networks_rwd_us_network_memberships(id, opts)

US real-world data network membership (N3C)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdUsNetworksApi.new
id = 'n3c' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # US real-world data network membership (N3C)
  result = api_instance.rwd_us_networks_rwd_us_network_memberships(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdUsNetworksApi->rwd_us_networks_rwd_us_network_memberships: #{e}"
end
```

#### Using the rwd_us_networks_rwd_us_network_memberships_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdUsNetworkMembershipsResponse>, Integer, Hash)> rwd_us_networks_rwd_us_network_memberships_with_http_info(id, opts)

```ruby
begin
  # US real-world data network membership (N3C)
  data, status_code, headers = api_instance.rwd_us_networks_rwd_us_network_memberships_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdUsNetworkMembershipsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdUsNetworksApi->rwd_us_networks_rwd_us_network_memberships_with_http_info: #{e}"
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

[**RwdUsNetworkMembershipsResponse**](RwdUsNetworkMembershipsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## rwd_us_networks_rwd_us_network_projects

> <RwdUsNetworkProjectsResponse> rwd_us_networks_rwd_us_network_projects(id, opts)

US real-world data network projects (N3C)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::RwdUsNetworksApi.new
id = 'n3c' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # US real-world data network projects (N3C)
  result = api_instance.rwd_us_networks_rwd_us_network_projects(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdUsNetworksApi->rwd_us_networks_rwd_us_network_projects: #{e}"
end
```

#### Using the rwd_us_networks_rwd_us_network_projects_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RwdUsNetworkProjectsResponse>, Integer, Hash)> rwd_us_networks_rwd_us_network_projects_with_http_info(id, opts)

```ruby
begin
  # US real-world data network projects (N3C)
  data, status_code, headers = api_instance.rwd_us_networks_rwd_us_network_projects_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RwdUsNetworkProjectsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling RwdUsNetworksApi->rwd_us_networks_rwd_us_network_projects_with_http_info: #{e}"
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

[**RwdUsNetworkProjectsResponse**](RwdUsNetworkProjectsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

