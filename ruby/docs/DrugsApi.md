# PharmaObservatory::DrugsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**drugs_brand_companies**](DrugsApi.md#drugs_brand_companies) | **GET** /api/v1/drugs/{id}/brand_companies | Brand companies |
| [**drugs_brand_revenue**](DrugsApi.md#drugs_brand_revenue) | **GET** /api/v1/drugs/{id}/brand_revenue | Brand revenue |
| [**drugs_catalysts**](DrugsApi.md#drugs_catalysts) | **GET** /api/v1/drugs/{id}/catalysts | Catalysts: trial readouts, patent cliffs, FDA decision dates and press-release events |
| [**drugs_drug_approvals**](DrugsApi.md#drugs_drug_approvals) | **GET** /api/v1/drugs/{id}/drug_approvals | Drug approvals |
| [**drugs_drug_asp_codes**](DrugsApi.md#drugs_drug_asp_codes) | **GET** /api/v1/drugs/{id}/drug_asp_codes | Drug ASP payment limits |
| [**drugs_drug_brands**](DrugsApi.md#drugs_drug_brands) | **GET** /api/v1/drugs/{id}/drug_brands | Drug brands |
| [**drugs_drug_combinations**](DrugsApi.md#drugs_drug_combinations) | **GET** /api/v1/drugs/{id}/drug_combinations | Drug combinations |
| [**drugs_drug_declared_patents**](DrugsApi.md#drugs_drug_declared_patents) | **GET** /api/v1/drugs/{id}/drug_declared_patents | Drug declared patents |
| [**drugs_drug_investigators**](DrugsApi.md#drugs_drug_investigators) | **GET** /api/v1/drugs/{id}/drug_investigators | Investigators on a drug&#39;s trials |
| [**drugs_drug_label_safety**](DrugsApi.md#drugs_drug_label_safety) | **GET** /api/v1/drugs/{id}/drug_label_safety | Drug label safety |
| [**drugs_drug_medicare_spend**](DrugsApi.md#drugs_drug_medicare_spend) | **GET** /api/v1/drugs/{id}/drug_medicare_spend | Drug Medicare spend |
| [**drugs_drug_patent_cliff**](DrugsApi.md#drugs_drug_patent_cliff) | **GET** /api/v1/drugs/{id}/drug_patent_cliff | Patent-cliff exposure |
| [**drugs_drug_patent_mentions**](DrugsApi.md#drugs_drug_patent_mentions) | **GET** /api/v1/drugs/{id}/drug_patent_mentions | Drug patent mentions |
| [**drugs_drug_patent_owners**](DrugsApi.md#drugs_drug_patent_owners) | **GET** /api/v1/drugs/{id}/drug_patent_owners | Drug patent owners |
| [**drugs_drug_pricing**](DrugsApi.md#drugs_drug_pricing) | **GET** /api/v1/drugs/{id}/drug_pricing | Drug pricing (NADAC) |
| [**drugs_drug_recalls**](DrugsApi.md#drugs_drug_recalls) | **GET** /api/v1/drugs/{id}/drug_recalls | Drug recalls |
| [**drugs_drug_regulatory**](DrugsApi.md#drugs_drug_regulatory) | **GET** /api/v1/drugs/{id}/drug_regulatory | Drug regulatory status |
| [**drugs_drug_sites**](DrugsApi.md#drugs_drug_sites) | **GET** /api/v1/drugs/{id}/drug_sites | Drug sites |
| [**drugs_drug_targets**](DrugsApi.md#drugs_drug_targets) | **GET** /api/v1/drugs/{id}/drug_targets | Drug targets |
| [**drugs_drug_trial_conditions**](DrugsApi.md#drugs_drug_trial_conditions) | **GET** /api/v1/drugs/{id}/drug_trial_conditions | Drug trial conditions |
| [**drugs_drug_trials**](DrugsApi.md#drugs_drug_trials) | **GET** /api/v1/drugs/{id}/drug_trials | Drug trials |
| [**drugs_drug_variants**](DrugsApi.md#drugs_drug_variants) | **GET** /api/v1/drugs/{id}/drug_variants | Drug variants |
| [**drugs_get**](DrugsApi.md#drugs_get) | **GET** /api/v1/drugs/{id} | One drug |
| [**drugs_list**](DrugsApi.md#drugs_list) | **GET** /api/v1/drugs | List drugs |
| [**drugs_pdufa_dates**](DrugsApi.md#drugs_pdufa_dates) | **GET** /api/v1/drugs/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts |
| [**drugs_programs**](DrugsApi.md#drugs_programs) | **GET** /api/v1/drugs/{id}/programs | Programs |
| [**drugs_sec_pdufa_dates**](DrugsApi.md#drugs_sec_pdufa_dates) | **GET** /api/v1/drugs/{id}/sec_pdufa_dates | FDA dates disclosed in SEC filings |
| [**drugs_trial_drugs**](DrugsApi.md#drugs_trial_drugs) | **GET** /api/v1/drugs/{id}/trial_drugs | Trial drugs |


## drugs_brand_companies

> <BrandCompaniesResponse> drugs_brand_companies(id, opts)

Brand companies

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Brand companies
  result = api_instance.drugs_brand_companies(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_brand_companies: #{e}"
end
```

#### Using the drugs_brand_companies_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BrandCompaniesResponse>, Integer, Hash)> drugs_brand_companies_with_http_info(id, opts)

```ruby
begin
  # Brand companies
  data, status_code, headers = api_instance.drugs_brand_companies_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BrandCompaniesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_brand_companies_with_http_info: #{e}"
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

[**BrandCompaniesResponse**](BrandCompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_brand_revenue

> <BrandRevenueResponse> drugs_brand_revenue(id, opts)

Brand revenue

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Brand revenue
  result = api_instance.drugs_brand_revenue(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_brand_revenue: #{e}"
end
```

#### Using the drugs_brand_revenue_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BrandRevenueResponse>, Integer, Hash)> drugs_brand_revenue_with_http_info(id, opts)

```ruby
begin
  # Brand revenue
  data, status_code, headers = api_instance.drugs_brand_revenue_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BrandRevenueResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_brand_revenue_with_http_info: #{e}"
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

[**BrandRevenueResponse**](BrandRevenueResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_catalysts

> <CatalystsResponse> drugs_catalysts(id, opts)

Catalysts: trial readouts, patent cliffs, FDA decision dates and press-release events

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Catalysts: trial readouts, patent cliffs, FDA decision dates and press-release events
  result = api_instance.drugs_catalysts(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_catalysts: #{e}"
end
```

#### Using the drugs_catalysts_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CatalystsResponse>, Integer, Hash)> drugs_catalysts_with_http_info(id, opts)

```ruby
begin
  # Catalysts: trial readouts, patent cliffs, FDA decision dates and press-release events
  data, status_code, headers = api_instance.drugs_catalysts_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CatalystsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_catalysts_with_http_info: #{e}"
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

[**CatalystsResponse**](CatalystsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_approvals

> <DrugApprovalsResponse> drugs_drug_approvals(id, opts)

Drug approvals

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug approvals
  result = api_instance.drugs_drug_approvals(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_approvals: #{e}"
end
```

#### Using the drugs_drug_approvals_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugApprovalsResponse>, Integer, Hash)> drugs_drug_approvals_with_http_info(id, opts)

```ruby
begin
  # Drug approvals
  data, status_code, headers = api_instance.drugs_drug_approvals_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugApprovalsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_approvals_with_http_info: #{e}"
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

[**DrugApprovalsResponse**](DrugApprovalsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_asp_codes

> <DrugAspCodesResponse> drugs_drug_asp_codes(id, opts)

Drug ASP payment limits

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug ASP payment limits
  result = api_instance.drugs_drug_asp_codes(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_asp_codes: #{e}"
end
```

#### Using the drugs_drug_asp_codes_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugAspCodesResponse>, Integer, Hash)> drugs_drug_asp_codes_with_http_info(id, opts)

```ruby
begin
  # Drug ASP payment limits
  data, status_code, headers = api_instance.drugs_drug_asp_codes_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugAspCodesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_asp_codes_with_http_info: #{e}"
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

[**DrugAspCodesResponse**](DrugAspCodesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_brands

> <DrugBrandsResponse> drugs_drug_brands(id, opts)

Drug brands

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug brands
  result = api_instance.drugs_drug_brands(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_brands: #{e}"
end
```

#### Using the drugs_drug_brands_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugBrandsResponse>, Integer, Hash)> drugs_drug_brands_with_http_info(id, opts)

```ruby
begin
  # Drug brands
  data, status_code, headers = api_instance.drugs_drug_brands_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugBrandsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_brands_with_http_info: #{e}"
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

[**DrugBrandsResponse**](DrugBrandsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_combinations

> <DrugCombinationsResponse> drugs_drug_combinations(id, opts)

Drug combinations

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug combinations
  result = api_instance.drugs_drug_combinations(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_combinations: #{e}"
end
```

#### Using the drugs_drug_combinations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugCombinationsResponse>, Integer, Hash)> drugs_drug_combinations_with_http_info(id, opts)

```ruby
begin
  # Drug combinations
  data, status_code, headers = api_instance.drugs_drug_combinations_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugCombinationsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_combinations_with_http_info: #{e}"
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

[**DrugCombinationsResponse**](DrugCombinationsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_declared_patents

> <DrugDeclaredPatentsResponse> drugs_drug_declared_patents(id, opts)

Drug declared patents

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug declared patents
  result = api_instance.drugs_drug_declared_patents(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_declared_patents: #{e}"
end
```

#### Using the drugs_drug_declared_patents_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugDeclaredPatentsResponse>, Integer, Hash)> drugs_drug_declared_patents_with_http_info(id, opts)

```ruby
begin
  # Drug declared patents
  data, status_code, headers = api_instance.drugs_drug_declared_patents_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugDeclaredPatentsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_declared_patents_with_http_info: #{e}"
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

[**DrugDeclaredPatentsResponse**](DrugDeclaredPatentsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_investigators

> <DrugInvestigatorsResponse> drugs_drug_investigators(id, opts)

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

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Investigators on a drug's trials
  result = api_instance.drugs_drug_investigators(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_investigators: #{e}"
end
```

#### Using the drugs_drug_investigators_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugInvestigatorsResponse>, Integer, Hash)> drugs_drug_investigators_with_http_info(id, opts)

```ruby
begin
  # Investigators on a drug's trials
  data, status_code, headers = api_instance.drugs_drug_investigators_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugInvestigatorsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_investigators_with_http_info: #{e}"
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


## drugs_drug_label_safety

> <DrugLabelSafetyResponse> drugs_drug_label_safety(id, opts)

Drug label safety

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug label safety
  result = api_instance.drugs_drug_label_safety(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_label_safety: #{e}"
end
```

#### Using the drugs_drug_label_safety_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugLabelSafetyResponse>, Integer, Hash)> drugs_drug_label_safety_with_http_info(id, opts)

```ruby
begin
  # Drug label safety
  data, status_code, headers = api_instance.drugs_drug_label_safety_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugLabelSafetyResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_label_safety_with_http_info: #{e}"
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

[**DrugLabelSafetyResponse**](DrugLabelSafetyResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_medicare_spend

> <DrugMedicareSpendResponse> drugs_drug_medicare_spend(id, opts)

Drug Medicare spend

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug Medicare spend
  result = api_instance.drugs_drug_medicare_spend(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_medicare_spend: #{e}"
end
```

#### Using the drugs_drug_medicare_spend_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugMedicareSpendResponse>, Integer, Hash)> drugs_drug_medicare_spend_with_http_info(id, opts)

```ruby
begin
  # Drug Medicare spend
  data, status_code, headers = api_instance.drugs_drug_medicare_spend_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugMedicareSpendResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_medicare_spend_with_http_info: #{e}"
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

[**DrugMedicareSpendResponse**](DrugMedicareSpendResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_patent_cliff

> <DrugPatentCliffResponse> drugs_drug_patent_cliff(id, opts)

Patent-cliff exposure

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Patent-cliff exposure
  result = api_instance.drugs_drug_patent_cliff(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_patent_cliff: #{e}"
end
```

#### Using the drugs_drug_patent_cliff_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugPatentCliffResponse>, Integer, Hash)> drugs_drug_patent_cliff_with_http_info(id, opts)

```ruby
begin
  # Patent-cliff exposure
  data, status_code, headers = api_instance.drugs_drug_patent_cliff_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugPatentCliffResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_patent_cliff_with_http_info: #{e}"
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

[**DrugPatentCliffResponse**](DrugPatentCliffResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_patent_mentions

> <DrugPatentMentionsResponse> drugs_drug_patent_mentions(id, opts)

Drug patent mentions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug patent mentions
  result = api_instance.drugs_drug_patent_mentions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_patent_mentions: #{e}"
end
```

#### Using the drugs_drug_patent_mentions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugPatentMentionsResponse>, Integer, Hash)> drugs_drug_patent_mentions_with_http_info(id, opts)

```ruby
begin
  # Drug patent mentions
  data, status_code, headers = api_instance.drugs_drug_patent_mentions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugPatentMentionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_patent_mentions_with_http_info: #{e}"
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

[**DrugPatentMentionsResponse**](DrugPatentMentionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_patent_owners

> <DrugPatentOwnersResponse> drugs_drug_patent_owners(id, opts)

Drug patent owners

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug patent owners
  result = api_instance.drugs_drug_patent_owners(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_patent_owners: #{e}"
end
```

#### Using the drugs_drug_patent_owners_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugPatentOwnersResponse>, Integer, Hash)> drugs_drug_patent_owners_with_http_info(id, opts)

```ruby
begin
  # Drug patent owners
  data, status_code, headers = api_instance.drugs_drug_patent_owners_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugPatentOwnersResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_patent_owners_with_http_info: #{e}"
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

[**DrugPatentOwnersResponse**](DrugPatentOwnersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_pricing

> <DrugPricingResponse> drugs_drug_pricing(id, opts)

Drug pricing (NADAC)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug pricing (NADAC)
  result = api_instance.drugs_drug_pricing(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_pricing: #{e}"
end
```

#### Using the drugs_drug_pricing_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugPricingResponse>, Integer, Hash)> drugs_drug_pricing_with_http_info(id, opts)

```ruby
begin
  # Drug pricing (NADAC)
  data, status_code, headers = api_instance.drugs_drug_pricing_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugPricingResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_pricing_with_http_info: #{e}"
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

[**DrugPricingResponse**](DrugPricingResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_recalls

> <DrugRecallsResponse> drugs_drug_recalls(id, opts)

Drug recalls

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug recalls
  result = api_instance.drugs_drug_recalls(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_recalls: #{e}"
end
```

#### Using the drugs_drug_recalls_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugRecallsResponse>, Integer, Hash)> drugs_drug_recalls_with_http_info(id, opts)

```ruby
begin
  # Drug recalls
  data, status_code, headers = api_instance.drugs_drug_recalls_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugRecallsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_recalls_with_http_info: #{e}"
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

[**DrugRecallsResponse**](DrugRecallsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_regulatory

> <DrugRegulatoryResponse> drugs_drug_regulatory(id, opts)

Drug regulatory status

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug regulatory status
  result = api_instance.drugs_drug_regulatory(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_regulatory: #{e}"
end
```

#### Using the drugs_drug_regulatory_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugRegulatoryResponse>, Integer, Hash)> drugs_drug_regulatory_with_http_info(id, opts)

```ruby
begin
  # Drug regulatory status
  data, status_code, headers = api_instance.drugs_drug_regulatory_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugRegulatoryResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_regulatory_with_http_info: #{e}"
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

[**DrugRegulatoryResponse**](DrugRegulatoryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_sites

> <DrugSitesResponse> drugs_drug_sites(id, opts)

Drug sites

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug sites
  result = api_instance.drugs_drug_sites(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_sites: #{e}"
end
```

#### Using the drugs_drug_sites_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugSitesResponse>, Integer, Hash)> drugs_drug_sites_with_http_info(id, opts)

```ruby
begin
  # Drug sites
  data, status_code, headers = api_instance.drugs_drug_sites_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugSitesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_sites_with_http_info: #{e}"
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

[**DrugSitesResponse**](DrugSitesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_targets

> <DrugTargetsResponse> drugs_drug_targets(id, opts)

Drug targets

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug targets
  result = api_instance.drugs_drug_targets(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_targets: #{e}"
end
```

#### Using the drugs_drug_targets_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugTargetsResponse>, Integer, Hash)> drugs_drug_targets_with_http_info(id, opts)

```ruby
begin
  # Drug targets
  data, status_code, headers = api_instance.drugs_drug_targets_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugTargetsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_targets_with_http_info: #{e}"
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

[**DrugTargetsResponse**](DrugTargetsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_trial_conditions

> <DrugTrialConditionsResponse> drugs_drug_trial_conditions(id, opts)

Drug trial conditions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug trial conditions
  result = api_instance.drugs_drug_trial_conditions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_trial_conditions: #{e}"
end
```

#### Using the drugs_drug_trial_conditions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugTrialConditionsResponse>, Integer, Hash)> drugs_drug_trial_conditions_with_http_info(id, opts)

```ruby
begin
  # Drug trial conditions
  data, status_code, headers = api_instance.drugs_drug_trial_conditions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugTrialConditionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_trial_conditions_with_http_info: #{e}"
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

[**DrugTrialConditionsResponse**](DrugTrialConditionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_trials

> <DrugTrialsResponse> drugs_drug_trials(id, opts)

Drug trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug trials
  result = api_instance.drugs_drug_trials(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_trials: #{e}"
end
```

#### Using the drugs_drug_trials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugTrialsResponse>, Integer, Hash)> drugs_drug_trials_with_http_info(id, opts)

```ruby
begin
  # Drug trials
  data, status_code, headers = api_instance.drugs_drug_trials_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugTrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_trials_with_http_info: #{e}"
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

[**DrugTrialsResponse**](DrugTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_drug_variants

> <DrugVariantsResponse> drugs_drug_variants(id, opts)

Drug variants

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drug variants
  result = api_instance.drugs_drug_variants(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_variants: #{e}"
end
```

#### Using the drugs_drug_variants_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugVariantsResponse>, Integer, Hash)> drugs_drug_variants_with_http_info(id, opts)

```ruby
begin
  # Drug variants
  data, status_code, headers = api_instance.drugs_drug_variants_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugVariantsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_drug_variants_with_http_info: #{e}"
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

[**DrugVariantsResponse**](DrugVariantsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_get

> <DrugsRecord> drugs_get(id)

One drug

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 

begin
  # One drug
  result = api_instance.drugs_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_get: #{e}"
end
```

#### Using the drugs_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugsRecord>, Integer, Hash)> drugs_get_with_http_info(id)

```ruby
begin
  # One drug
  data, status_code, headers = api_instance.drugs_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugsRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**DrugsRecord**](DrugsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_list

> <DrugsResponse> drugs_list(opts)

List drugs

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List drugs
  result = api_instance.drugs_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_list: #{e}"
end
```

#### Using the drugs_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugsResponse>, Integer, Hash)> drugs_list_with_http_info(opts)

```ruby
begin
  # List drugs
  data, status_code, headers = api_instance.drugs_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_list_with_http_info: #{e}"
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

[**DrugsResponse**](DrugsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_pdufa_dates

> <PdufaDatesResponse> drugs_pdufa_dates(id, opts)

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

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA decision dates (PDUFA) and catalysts
  result = api_instance.drugs_pdufa_dates(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_pdufa_dates: #{e}"
end
```

#### Using the drugs_pdufa_dates_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PdufaDatesResponse>, Integer, Hash)> drugs_pdufa_dates_with_http_info(id, opts)

```ruby
begin
  # FDA decision dates (PDUFA) and catalysts
  data, status_code, headers = api_instance.drugs_pdufa_dates_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PdufaDatesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_pdufa_dates_with_http_info: #{e}"
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


## drugs_programs

> <ProgramsResponse> drugs_programs(id, opts)

Programs

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Programs
  result = api_instance.drugs_programs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_programs: #{e}"
end
```

#### Using the drugs_programs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ProgramsResponse>, Integer, Hash)> drugs_programs_with_http_info(id, opts)

```ruby
begin
  # Programs
  data, status_code, headers = api_instance.drugs_programs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ProgramsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_programs_with_http_info: #{e}"
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

[**ProgramsResponse**](ProgramsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_sec_pdufa_dates

> <SecPdufaDatesResponse> drugs_sec_pdufa_dates(id, opts)

FDA dates disclosed in SEC filings

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA dates disclosed in SEC filings
  result = api_instance.drugs_sec_pdufa_dates(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_sec_pdufa_dates: #{e}"
end
```

#### Using the drugs_sec_pdufa_dates_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SecPdufaDatesResponse>, Integer, Hash)> drugs_sec_pdufa_dates_with_http_info(id, opts)

```ruby
begin
  # FDA dates disclosed in SEC filings
  data, status_code, headers = api_instance.drugs_sec_pdufa_dates_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SecPdufaDatesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_sec_pdufa_dates_with_http_info: #{e}"
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

[**SecPdufaDatesResponse**](SecPdufaDatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## drugs_trial_drugs

> <TrialDrugsResponse> drugs_trial_drugs(id, opts)

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

api_instance = PharmaObservatory::DrugsApi.new
id = 'unii:53AXN4NNHX' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial drugs
  result = api_instance.drugs_trial_drugs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_trial_drugs: #{e}"
end
```

#### Using the drugs_trial_drugs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TrialDrugsResponse>, Integer, Hash)> drugs_trial_drugs_with_http_info(id, opts)

```ruby
begin
  # Trial drugs
  data, status_code, headers = api_instance.drugs_trial_drugs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TrialDrugsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling DrugsApi->drugs_trial_drugs_with_http_info: #{e}"
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

