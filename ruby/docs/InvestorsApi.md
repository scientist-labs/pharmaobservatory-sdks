# PharmaObservatory::InvestorsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**investors_get**](InvestorsApi.md#investors_get) | **GET** /api/v1/investors/{id} | One investor |
| [**investors_institutional_holdings**](InvestorsApi.md#investors_institutional_holdings) | **GET** /api/v1/investors/{id}/institutional_holdings | Institutional holdings (SEC Form 13F) |
| [**investors_investor_funds**](InvestorsApi.md#investors_investor_funds) | **GET** /api/v1/investors/{id}/investor_funds | Investor private funds by type |
| [**investors_investor_portfolio**](InvestorsApi.md#investors_investor_portfolio) | **GET** /api/v1/investors/{id}/investor_portfolio | Investor portfolio (offering filings) |
| [**investors_investor_private_funds**](InvestorsApi.md#investors_investor_private_funds) | **GET** /api/v1/investors/{id}/investor_private_funds | Investor private funds |
| [**investors_list**](InvestorsApi.md#investors_list) | **GET** /api/v1/investors | List investors |
| [**investors_managers**](InvestorsApi.md#investors_managers) | **GET** /api/v1/investors/{id}/managers | Institutional managers (SEC Form 13F) |
| [**investors_ownership_positions**](InvestorsApi.md#investors_ownership_positions) | **GET** /api/v1/investors/{id}/ownership_positions | Ownership and control positions |


## investors_get

> <InvestorsRecord> investors_get(id)

One investor

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestorsApi.new
id = 'adv:283882' # String | 

begin
  # One investor
  result = api_instance.investors_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_get: #{e}"
end
```

#### Using the investors_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestorsRecord>, Integer, Hash)> investors_get_with_http_info(id)

```ruby
begin
  # One investor
  data, status_code, headers = api_instance.investors_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestorsRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**InvestorsRecord**](InvestorsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investors_institutional_holdings

> <InstitutionalHoldingsResponse> investors_institutional_holdings(id, opts)

Institutional holdings (SEC Form 13F)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestorsApi.new
id = 'adv:283882' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Institutional holdings (SEC Form 13F)
  result = api_instance.investors_institutional_holdings(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_institutional_holdings: #{e}"
end
```

#### Using the investors_institutional_holdings_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InstitutionalHoldingsResponse>, Integer, Hash)> investors_institutional_holdings_with_http_info(id, opts)

```ruby
begin
  # Institutional holdings (SEC Form 13F)
  data, status_code, headers = api_instance.investors_institutional_holdings_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InstitutionalHoldingsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_institutional_holdings_with_http_info: #{e}"
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

[**InstitutionalHoldingsResponse**](InstitutionalHoldingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investors_investor_funds

> <InvestorFundsResponse> investors_investor_funds(id, opts)

Investor private funds by type

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestorsApi.new
id = 'adv:283882' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investor private funds by type
  result = api_instance.investors_investor_funds(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_investor_funds: #{e}"
end
```

#### Using the investors_investor_funds_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestorFundsResponse>, Integer, Hash)> investors_investor_funds_with_http_info(id, opts)

```ruby
begin
  # Investor private funds by type
  data, status_code, headers = api_instance.investors_investor_funds_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestorFundsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_investor_funds_with_http_info: #{e}"
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

[**InvestorFundsResponse**](InvestorFundsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investors_investor_portfolio

> <InvestorPortfolioResponse> investors_investor_portfolio(id, opts)

Investor portfolio (offering filings)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestorsApi.new
id = 'adv:283882' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investor portfolio (offering filings)
  result = api_instance.investors_investor_portfolio(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_investor_portfolio: #{e}"
end
```

#### Using the investors_investor_portfolio_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestorPortfolioResponse>, Integer, Hash)> investors_investor_portfolio_with_http_info(id, opts)

```ruby
begin
  # Investor portfolio (offering filings)
  data, status_code, headers = api_instance.investors_investor_portfolio_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestorPortfolioResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_investor_portfolio_with_http_info: #{e}"
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

[**InvestorPortfolioResponse**](InvestorPortfolioResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investors_investor_private_funds

> <InvestorPrivateFundsResponse> investors_investor_private_funds(id, opts)

Investor private funds

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestorsApi.new
id = 'adv:283882' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investor private funds
  result = api_instance.investors_investor_private_funds(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_investor_private_funds: #{e}"
end
```

#### Using the investors_investor_private_funds_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestorPrivateFundsResponse>, Integer, Hash)> investors_investor_private_funds_with_http_info(id, opts)

```ruby
begin
  # Investor private funds
  data, status_code, headers = api_instance.investors_investor_private_funds_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestorPrivateFundsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_investor_private_funds_with_http_info: #{e}"
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

[**InvestorPrivateFundsResponse**](InvestorPrivateFundsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investors_list

> <InvestorsResponse> investors_list(opts)

List investors

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestorsApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List investors
  result = api_instance.investors_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_list: #{e}"
end
```

#### Using the investors_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InvestorsResponse>, Integer, Hash)> investors_list_with_http_info(opts)

```ruby
begin
  # List investors
  data, status_code, headers = api_instance.investors_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InvestorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_list_with_http_info: #{e}"
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

[**InvestorsResponse**](InvestorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investors_managers

> <ManagersResponse> investors_managers(id, opts)

Institutional managers (SEC Form 13F)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestorsApi.new
id = 'adv:283882' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Institutional managers (SEC Form 13F)
  result = api_instance.investors_managers(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_managers: #{e}"
end
```

#### Using the investors_managers_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ManagersResponse>, Integer, Hash)> investors_managers_with_http_info(id, opts)

```ruby
begin
  # Institutional managers (SEC Form 13F)
  data, status_code, headers = api_instance.investors_managers_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ManagersResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_managers_with_http_info: #{e}"
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

[**ManagersResponse**](ManagersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investors_ownership_positions

> <OwnershipPositionsResponse> investors_ownership_positions(id, opts)

Ownership and control positions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::InvestorsApi.new
id = 'adv:283882' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Ownership and control positions
  result = api_instance.investors_ownership_positions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_ownership_positions: #{e}"
end
```

#### Using the investors_ownership_positions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OwnershipPositionsResponse>, Integer, Hash)> investors_ownership_positions_with_http_info(id, opts)

```ruby
begin
  # Ownership and control positions
  data, status_code, headers = api_instance.investors_ownership_positions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OwnershipPositionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling InvestorsApi->investors_ownership_positions_with_http_info: #{e}"
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

[**OwnershipPositionsResponse**](OwnershipPositionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

