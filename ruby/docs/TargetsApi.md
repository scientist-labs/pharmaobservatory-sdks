# PharmaObservatory::TargetsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**targets_biomarker_evidence**](TargetsApi.md#targets_biomarker_evidence) | **GET** /api/v1/targets/{id}/biomarker_evidence | Biomarker evidence |
| [**targets_company_targets**](TargetsApi.md#targets_company_targets) | **GET** /api/v1/targets/{id}/company_targets | Targets a company pursues |
| [**targets_get**](TargetsApi.md#targets_get) | **GET** /api/v1/targets/{id} | One target |
| [**targets_list**](TargetsApi.md#targets_list) | **GET** /api/v1/targets | List targets |
| [**targets_target_associations**](TargetsApi.md#targets_target_associations) | **GET** /api/v1/targets/{id}/target_associations | Target-disease associations by datasource |
| [**targets_target_crowding**](TargetsApi.md#targets_target_crowding) | **GET** /api/v1/targets/{id}/target_crowding | Target crowding by phase |
| [**targets_target_drugs**](TargetsApi.md#targets_target_drugs) | **GET** /api/v1/targets/{id}/target_drugs | Drugs acting on a target |
| [**targets_target_evidence**](TargetsApi.md#targets_target_evidence) | **GET** /api/v1/targets/{id}/target_evidence | Target evidence |
| [**targets_target_patent_mentions**](TargetsApi.md#targets_target_patent_mentions) | **GET** /api/v1/targets/{id}/target_patent_mentions | Target patent mentions |
| [**targets_target_programs**](TargetsApi.md#targets_target_programs) | **GET** /api/v1/targets/{id}/target_programs | Target development programs |
| [**targets_target_publications**](TargetsApi.md#targets_target_publications) | **GET** /api/v1/targets/{id}/target_publications | Target publications |
| [**targets_target_regulatory_precedent**](TargetsApi.md#targets_target_regulatory_precedent) | **GET** /api/v1/targets/{id}/target_regulatory_precedent | Target regulatory precedent |
| [**targets_target_trials**](TargetsApi.md#targets_target_trials) | **GET** /api/v1/targets/{id}/target_trials | Trials targeting a target |


## targets_biomarker_evidence

> <BiomarkerEvidenceResponse> targets_biomarker_evidence(id, opts)

Biomarker evidence

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Biomarker evidence
  result = api_instance.targets_biomarker_evidence(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_biomarker_evidence: #{e}"
end
```

#### Using the targets_biomarker_evidence_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerEvidenceResponse>, Integer, Hash)> targets_biomarker_evidence_with_http_info(id, opts)

```ruby
begin
  # Biomarker evidence
  data, status_code, headers = api_instance.targets_biomarker_evidence_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerEvidenceResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_biomarker_evidence_with_http_info: #{e}"
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

[**BiomarkerEvidenceResponse**](BiomarkerEvidenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_company_targets

> <CompanyTargetsResponse> targets_company_targets(id, opts)

Targets a company pursues

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Targets a company pursues
  result = api_instance.targets_company_targets(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_company_targets: #{e}"
end
```

#### Using the targets_company_targets_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CompanyTargetsResponse>, Integer, Hash)> targets_company_targets_with_http_info(id, opts)

```ruby
begin
  # Targets a company pursues
  data, status_code, headers = api_instance.targets_company_targets_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CompanyTargetsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_company_targets_with_http_info: #{e}"
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

[**CompanyTargetsResponse**](CompanyTargetsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_get

> <TargetsRecord> targets_get(id)

One target

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 

begin
  # One target
  result = api_instance.targets_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_get: #{e}"
end
```

#### Using the targets_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetsRecord>, Integer, Hash)> targets_get_with_http_info(id)

```ruby
begin
  # One target
  data, status_code, headers = api_instance.targets_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetsRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**TargetsRecord**](TargetsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_list

> <TargetsResponse> targets_list(opts)

List targets

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List targets
  result = api_instance.targets_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_list: #{e}"
end
```

#### Using the targets_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetsResponse>, Integer, Hash)> targets_list_with_http_info(opts)

```ruby
begin
  # List targets
  data, status_code, headers = api_instance.targets_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_list_with_http_info: #{e}"
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

[**TargetsResponse**](TargetsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_target_associations

> <TargetAssociationsResponse> targets_target_associations(id, opts)

Target-disease associations by datasource

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Target-disease associations by datasource
  result = api_instance.targets_target_associations(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_associations: #{e}"
end
```

#### Using the targets_target_associations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetAssociationsResponse>, Integer, Hash)> targets_target_associations_with_http_info(id, opts)

```ruby
begin
  # Target-disease associations by datasource
  data, status_code, headers = api_instance.targets_target_associations_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetAssociationsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_associations_with_http_info: #{e}"
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

[**TargetAssociationsResponse**](TargetAssociationsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_target_crowding

> <TargetCrowdingResponse> targets_target_crowding(id, opts)

Target crowding by phase

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Target crowding by phase
  result = api_instance.targets_target_crowding(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_crowding: #{e}"
end
```

#### Using the targets_target_crowding_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetCrowdingResponse>, Integer, Hash)> targets_target_crowding_with_http_info(id, opts)

```ruby
begin
  # Target crowding by phase
  data, status_code, headers = api_instance.targets_target_crowding_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetCrowdingResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_crowding_with_http_info: #{e}"
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

[**TargetCrowdingResponse**](TargetCrowdingResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_target_drugs

> <TargetDrugsResponse> targets_target_drugs(id, opts)

Drugs acting on a target

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Drugs acting on a target
  result = api_instance.targets_target_drugs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_drugs: #{e}"
end
```

#### Using the targets_target_drugs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetDrugsResponse>, Integer, Hash)> targets_target_drugs_with_http_info(id, opts)

```ruby
begin
  # Drugs acting on a target
  data, status_code, headers = api_instance.targets_target_drugs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetDrugsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_drugs_with_http_info: #{e}"
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

[**TargetDrugsResponse**](TargetDrugsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_target_evidence

> <TargetEvidenceResponse> targets_target_evidence(id, opts)

Target evidence

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Target evidence
  result = api_instance.targets_target_evidence(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_evidence: #{e}"
end
```

#### Using the targets_target_evidence_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetEvidenceResponse>, Integer, Hash)> targets_target_evidence_with_http_info(id, opts)

```ruby
begin
  # Target evidence
  data, status_code, headers = api_instance.targets_target_evidence_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetEvidenceResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_evidence_with_http_info: #{e}"
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

[**TargetEvidenceResponse**](TargetEvidenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_target_patent_mentions

> <TargetPatentMentionsResponse> targets_target_patent_mentions(id, opts)

Target patent mentions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Target patent mentions
  result = api_instance.targets_target_patent_mentions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_patent_mentions: #{e}"
end
```

#### Using the targets_target_patent_mentions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetPatentMentionsResponse>, Integer, Hash)> targets_target_patent_mentions_with_http_info(id, opts)

```ruby
begin
  # Target patent mentions
  data, status_code, headers = api_instance.targets_target_patent_mentions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetPatentMentionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_patent_mentions_with_http_info: #{e}"
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

[**TargetPatentMentionsResponse**](TargetPatentMentionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_target_programs

> <TargetProgramsResponse> targets_target_programs(id, opts)

Target development programs

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Target development programs
  result = api_instance.targets_target_programs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_programs: #{e}"
end
```

#### Using the targets_target_programs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetProgramsResponse>, Integer, Hash)> targets_target_programs_with_http_info(id, opts)

```ruby
begin
  # Target development programs
  data, status_code, headers = api_instance.targets_target_programs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetProgramsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_programs_with_http_info: #{e}"
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

[**TargetProgramsResponse**](TargetProgramsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_target_publications

> <TargetPublicationsResponse> targets_target_publications(id, opts)

Target publications

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Target publications
  result = api_instance.targets_target_publications(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_publications: #{e}"
end
```

#### Using the targets_target_publications_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetPublicationsResponse>, Integer, Hash)> targets_target_publications_with_http_info(id, opts)

```ruby
begin
  # Target publications
  data, status_code, headers = api_instance.targets_target_publications_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetPublicationsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_publications_with_http_info: #{e}"
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

[**TargetPublicationsResponse**](TargetPublicationsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_target_regulatory_precedent

> <TargetRegulatoryPrecedentResponse> targets_target_regulatory_precedent(id, opts)

Target regulatory precedent

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Target regulatory precedent
  result = api_instance.targets_target_regulatory_precedent(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_regulatory_precedent: #{e}"
end
```

#### Using the targets_target_regulatory_precedent_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetRegulatoryPrecedentResponse>, Integer, Hash)> targets_target_regulatory_precedent_with_http_info(id, opts)

```ruby
begin
  # Target regulatory precedent
  data, status_code, headers = api_instance.targets_target_regulatory_precedent_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetRegulatoryPrecedentResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_regulatory_precedent_with_http_info: #{e}"
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

[**TargetRegulatoryPrecedentResponse**](TargetRegulatoryPrecedentResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## targets_target_trials

> <TargetTrialsResponse> targets_target_trials(id, opts)

Trials targeting a target

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::TargetsApi.new
id = 'ENSG00000171759' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trials targeting a target
  result = api_instance.targets_target_trials(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_trials: #{e}"
end
```

#### Using the targets_target_trials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TargetTrialsResponse>, Integer, Hash)> targets_target_trials_with_http_info(id, opts)

```ruby
begin
  # Trials targeting a target
  data, status_code, headers = api_instance.targets_target_trials_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TargetTrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling TargetsApi->targets_target_trials_with_http_info: #{e}"
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

[**TargetTrialsResponse**](TargetTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

