# PharmaObservatory::LegalEntitiesApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**legal_entities_catalysts**](LegalEntitiesApi.md#legal_entities_catalysts) | **GET** /api/v1/legal_entities/{id}/catalysts | Catalysts: trial readouts, patent cliffs, FDA decision dates and press-release events |
| [**legal_entities_clia_certificates**](LegalEntitiesApi.md#legal_entities_clia_certificates) | **GET** /api/v1/legal_entities/{id}/clia_certificates | CLIA laboratory certificates |
| [**legal_entities_company_assets**](LegalEntitiesApi.md#legal_entities_company_assets) | **GET** /api/v1/legal_entities/{id}/company_assets | Company assets (approved and marketed drugs) |
| [**legal_entities_company_programs**](LegalEntitiesApi.md#legal_entities_company_programs) | **GET** /api/v1/legal_entities/{id}/company_programs | Company programs |
| [**legal_entities_debarments**](LegalEntitiesApi.md#legal_entities_debarments) | **GET** /api/v1/legal_entities/{id}/debarments | FDA firm debarments |
| [**legal_entities_drug_patent_cliff**](LegalEntitiesApi.md#legal_entities_drug_patent_cliff) | **GET** /api/v1/legal_entities/{id}/drug_patent_cliff | Patent-cliff exposure |
| [**legal_entities_enforcement_events**](LegalEntitiesApi.md#legal_entities_enforcement_events) | **GET** /api/v1/legal_entities/{id}/enforcement_events | FDA enforcement events |
| [**legal_entities_get**](LegalEntitiesApi.md#legal_entities_get) | **GET** /api/v1/legal_entities/{id} | One legal entity |
| [**legal_entities_gxp_facilities**](LegalEntitiesApi.md#legal_entities_gxp_facilities) | **GET** /api/v1/legal_entities/{id}/gxp_facilities | GxP facilities and quality signals |
| [**legal_entities_import_refusals**](LegalEntitiesApi.md#legal_entities_import_refusals) | **GET** /api/v1/legal_entities/{id}/import_refusals | FDA import refusals |
| [**legal_entities_inspection_findings**](LegalEntitiesApi.md#legal_entities_inspection_findings) | **GET** /api/v1/legal_entities/{id}/inspection_findings | FDA inspection findings |
| [**legal_entities_inspections**](LegalEntitiesApi.md#legal_entities_inspections) | **GET** /api/v1/legal_entities/{id}/inspections | FDA inspections |
| [**legal_entities_institutional_holdings**](LegalEntitiesApi.md#legal_entities_institutional_holdings) | **GET** /api/v1/legal_entities/{id}/institutional_holdings | Institutional holdings (SEC Form 13F) |
| [**legal_entities_institutional_quarter_totals**](LegalEntitiesApi.md#legal_entities_institutional_quarter_totals) | **GET** /api/v1/legal_entities/{id}/institutional_quarter_totals | Institutional quarter totals (SEC Form 13F) |
| [**legal_entities_list**](LegalEntitiesApi.md#legal_entities_list) | **GET** /api/v1/legal_entities | List legal entities |
| [**legal_entities_org_disclosures**](LegalEntitiesApi.md#legal_entities_org_disclosures) | **GET** /api/v1/legal_entities/{id}/org_disclosures | Organization disclosures (climate and diversity) |
| [**legal_entities_org_relationships**](LegalEntitiesApi.md#legal_entities_org_relationships) | **GET** /api/v1/legal_entities/{id}/org_relationships | Organization relationships |
| [**legal_entities_ownership_percentages**](LegalEntitiesApi.md#legal_entities_ownership_percentages) | **GET** /api/v1/legal_entities/{id}/ownership_percentages | Ownership percentages |
| [**legal_entities_ownership_positions**](LegalEntitiesApi.md#legal_entities_ownership_positions) | **GET** /api/v1/legal_entities/{id}/ownership_positions | Ownership and control positions |
| [**legal_entities_pdufa_dates**](LegalEntitiesApi.md#legal_entities_pdufa_dates) | **GET** /api/v1/legal_entities/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts |
| [**legal_entities_program_licensing_candidates**](LegalEntitiesApi.md#legal_entities_program_licensing_candidates) | **GET** /api/v1/legal_entities/{id}/program_licensing_candidates | Program licensing candidates |
| [**legal_entities_recalls**](LegalEntitiesApi.md#legal_entities_recalls) | **GET** /api/v1/legal_entities/{id}/recalls | FDA recalls |
| [**legal_entities_sponsor_recruitment_by_phase**](LegalEntitiesApi.md#legal_entities_sponsor_recruitment_by_phase) | **GET** /api/v1/legal_entities/{id}/sponsor_recruitment_by_phase | Trial recruitment duration by sponsor and phase |
| [**legal_entities_sponsor_recruitment_duration**](LegalEntitiesApi.md#legal_entities_sponsor_recruitment_duration) | **GET** /api/v1/legal_entities/{id}/sponsor_recruitment_duration | Trial recruitment duration by sponsor |
| [**legal_entities_sponsor_trial_start_forecast**](LegalEntitiesApi.md#legal_entities_sponsor_trial_start_forecast) | **GET** /api/v1/legal_entities/{id}/sponsor_trial_start_forecast | Trial-start forecast by sponsor |
| [**legal_entities_sponsor_trials**](LegalEntitiesApi.md#legal_entities_sponsor_trials) | **GET** /api/v1/legal_entities/{id}/sponsor_trials | Sponsor trials |
| [**legal_entities_supplier_capability_esg**](LegalEntitiesApi.md#legal_entities_supplier_capability_esg) | **GET** /api/v1/legal_entities/{id}/supplier_capability_esg | Supplier capability and ESG disclosure |
| [**legal_entities_unmatched_product_revenue**](LegalEntitiesApi.md#legal_entities_unmatched_product_revenue) | **GET** /api/v1/legal_entities/{id}/unmatched_product_revenue | Unmatched product revenue |
| [**legal_entities_untitled_letters**](LegalEntitiesApi.md#legal_entities_untitled_letters) | **GET** /api/v1/legal_entities/{id}/untitled_letters | FDA untitled letters |
| [**legal_entities_warning_letters**](LegalEntitiesApi.md#legal_entities_warning_letters) | **GET** /api/v1/legal_entities/{id}/warning_letters | FDA warning letters |


## legal_entities_catalysts

> <CatalystsResponse> legal_entities_catalysts(id, opts)

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

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Catalysts: trial readouts, patent cliffs, FDA decision dates and press-release events
  result = api_instance.legal_entities_catalysts(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_catalysts: #{e}"
end
```

#### Using the legal_entities_catalysts_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CatalystsResponse>, Integer, Hash)> legal_entities_catalysts_with_http_info(id, opts)

```ruby
begin
  # Catalysts: trial readouts, patent cliffs, FDA decision dates and press-release events
  data, status_code, headers = api_instance.legal_entities_catalysts_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CatalystsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_catalysts_with_http_info: #{e}"
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


## legal_entities_clia_certificates

> <CliaCertificatesResponse> legal_entities_clia_certificates(id, opts)

CLIA laboratory certificates

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # CLIA laboratory certificates
  result = api_instance.legal_entities_clia_certificates(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_clia_certificates: #{e}"
end
```

#### Using the legal_entities_clia_certificates_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CliaCertificatesResponse>, Integer, Hash)> legal_entities_clia_certificates_with_http_info(id, opts)

```ruby
begin
  # CLIA laboratory certificates
  data, status_code, headers = api_instance.legal_entities_clia_certificates_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CliaCertificatesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_clia_certificates_with_http_info: #{e}"
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

[**CliaCertificatesResponse**](CliaCertificatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_company_assets

> <CompanyAssetsResponse> legal_entities_company_assets(id, opts)

Company assets (approved and marketed drugs)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Company assets (approved and marketed drugs)
  result = api_instance.legal_entities_company_assets(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_company_assets: #{e}"
end
```

#### Using the legal_entities_company_assets_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CompanyAssetsResponse>, Integer, Hash)> legal_entities_company_assets_with_http_info(id, opts)

```ruby
begin
  # Company assets (approved and marketed drugs)
  data, status_code, headers = api_instance.legal_entities_company_assets_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CompanyAssetsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_company_assets_with_http_info: #{e}"
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

[**CompanyAssetsResponse**](CompanyAssetsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_company_programs

> <CompanyProgramsResponse> legal_entities_company_programs(id, opts)

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

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Company programs
  result = api_instance.legal_entities_company_programs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_company_programs: #{e}"
end
```

#### Using the legal_entities_company_programs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CompanyProgramsResponse>, Integer, Hash)> legal_entities_company_programs_with_http_info(id, opts)

```ruby
begin
  # Company programs
  data, status_code, headers = api_instance.legal_entities_company_programs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CompanyProgramsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_company_programs_with_http_info: #{e}"
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


## legal_entities_debarments

> <DebarmentsResponse> legal_entities_debarments(id, opts)

FDA firm debarments

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA firm debarments
  result = api_instance.legal_entities_debarments(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_debarments: #{e}"
end
```

#### Using the legal_entities_debarments_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DebarmentsResponse>, Integer, Hash)> legal_entities_debarments_with_http_info(id, opts)

```ruby
begin
  # FDA firm debarments
  data, status_code, headers = api_instance.legal_entities_debarments_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DebarmentsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_debarments_with_http_info: #{e}"
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

[**DebarmentsResponse**](DebarmentsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_drug_patent_cliff

> <DrugPatentCliffResponse> legal_entities_drug_patent_cliff(id, opts)

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

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Patent-cliff exposure
  result = api_instance.legal_entities_drug_patent_cliff(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_drug_patent_cliff: #{e}"
end
```

#### Using the legal_entities_drug_patent_cliff_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugPatentCliffResponse>, Integer, Hash)> legal_entities_drug_patent_cliff_with_http_info(id, opts)

```ruby
begin
  # Patent-cliff exposure
  data, status_code, headers = api_instance.legal_entities_drug_patent_cliff_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugPatentCliffResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_drug_patent_cliff_with_http_info: #{e}"
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


## legal_entities_enforcement_events

> <EnforcementEventsResponse> legal_entities_enforcement_events(id, opts)

FDA enforcement events

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA enforcement events
  result = api_instance.legal_entities_enforcement_events(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_enforcement_events: #{e}"
end
```

#### Using the legal_entities_enforcement_events_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<EnforcementEventsResponse>, Integer, Hash)> legal_entities_enforcement_events_with_http_info(id, opts)

```ruby
begin
  # FDA enforcement events
  data, status_code, headers = api_instance.legal_entities_enforcement_events_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <EnforcementEventsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_enforcement_events_with_http_info: #{e}"
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

[**EnforcementEventsResponse**](EnforcementEventsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_get

> <LegalEntitiesRecord> legal_entities_get(id)

One legal entity

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 

begin
  # One legal entity
  result = api_instance.legal_entities_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_get: #{e}"
end
```

#### Using the legal_entities_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<LegalEntitiesRecord>, Integer, Hash)> legal_entities_get_with_http_info(id)

```ruby
begin
  # One legal entity
  data, status_code, headers = api_instance.legal_entities_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <LegalEntitiesRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**LegalEntitiesRecord**](LegalEntitiesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_gxp_facilities

> <GxpFacilitiesResponse> legal_entities_gxp_facilities(id, opts)

GxP facilities and quality signals

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # GxP facilities and quality signals
  result = api_instance.legal_entities_gxp_facilities(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_gxp_facilities: #{e}"
end
```

#### Using the legal_entities_gxp_facilities_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<GxpFacilitiesResponse>, Integer, Hash)> legal_entities_gxp_facilities_with_http_info(id, opts)

```ruby
begin
  # GxP facilities and quality signals
  data, status_code, headers = api_instance.legal_entities_gxp_facilities_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <GxpFacilitiesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_gxp_facilities_with_http_info: #{e}"
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

[**GxpFacilitiesResponse**](GxpFacilitiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_import_refusals

> <ImportRefusalsResponse> legal_entities_import_refusals(id, opts)

FDA import refusals

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA import refusals
  result = api_instance.legal_entities_import_refusals(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_import_refusals: #{e}"
end
```

#### Using the legal_entities_import_refusals_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ImportRefusalsResponse>, Integer, Hash)> legal_entities_import_refusals_with_http_info(id, opts)

```ruby
begin
  # FDA import refusals
  data, status_code, headers = api_instance.legal_entities_import_refusals_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ImportRefusalsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_import_refusals_with_http_info: #{e}"
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

[**ImportRefusalsResponse**](ImportRefusalsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_inspection_findings

> <InspectionFindingsResponse> legal_entities_inspection_findings(id, opts)

FDA inspection findings

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA inspection findings
  result = api_instance.legal_entities_inspection_findings(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_inspection_findings: #{e}"
end
```

#### Using the legal_entities_inspection_findings_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InspectionFindingsResponse>, Integer, Hash)> legal_entities_inspection_findings_with_http_info(id, opts)

```ruby
begin
  # FDA inspection findings
  data, status_code, headers = api_instance.legal_entities_inspection_findings_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InspectionFindingsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_inspection_findings_with_http_info: #{e}"
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

[**InspectionFindingsResponse**](InspectionFindingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_inspections

> <InspectionsResponse> legal_entities_inspections(id, opts)

FDA inspections

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA inspections
  result = api_instance.legal_entities_inspections(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_inspections: #{e}"
end
```

#### Using the legal_entities_inspections_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InspectionsResponse>, Integer, Hash)> legal_entities_inspections_with_http_info(id, opts)

```ruby
begin
  # FDA inspections
  data, status_code, headers = api_instance.legal_entities_inspections_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InspectionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_inspections_with_http_info: #{e}"
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

[**InspectionsResponse**](InspectionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_institutional_holdings

> <InstitutionalHoldingsResponse> legal_entities_institutional_holdings(id, opts)

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

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Institutional holdings (SEC Form 13F)
  result = api_instance.legal_entities_institutional_holdings(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_institutional_holdings: #{e}"
end
```

#### Using the legal_entities_institutional_holdings_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InstitutionalHoldingsResponse>, Integer, Hash)> legal_entities_institutional_holdings_with_http_info(id, opts)

```ruby
begin
  # Institutional holdings (SEC Form 13F)
  data, status_code, headers = api_instance.legal_entities_institutional_holdings_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InstitutionalHoldingsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_institutional_holdings_with_http_info: #{e}"
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


## legal_entities_institutional_quarter_totals

> <InstitutionalQuarterTotalsResponse> legal_entities_institutional_quarter_totals(id, opts)

Institutional quarter totals (SEC Form 13F)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Institutional quarter totals (SEC Form 13F)
  result = api_instance.legal_entities_institutional_quarter_totals(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_institutional_quarter_totals: #{e}"
end
```

#### Using the legal_entities_institutional_quarter_totals_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InstitutionalQuarterTotalsResponse>, Integer, Hash)> legal_entities_institutional_quarter_totals_with_http_info(id, opts)

```ruby
begin
  # Institutional quarter totals (SEC Form 13F)
  data, status_code, headers = api_instance.legal_entities_institutional_quarter_totals_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InstitutionalQuarterTotalsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_institutional_quarter_totals_with_http_info: #{e}"
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

[**InstitutionalQuarterTotalsResponse**](InstitutionalQuarterTotalsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_list

> <LegalEntitiesResponse> legal_entities_list(opts)

List legal entities

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # List legal entities
  result = api_instance.legal_entities_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_list: #{e}"
end
```

#### Using the legal_entities_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<LegalEntitiesResponse>, Integer, Hash)> legal_entities_list_with_http_info(opts)

```ruby
begin
  # List legal entities
  data, status_code, headers = api_instance.legal_entities_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <LegalEntitiesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_list_with_http_info: #{e}"
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

[**LegalEntitiesResponse**](LegalEntitiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_org_disclosures

> <OrgDisclosuresResponse> legal_entities_org_disclosures(id, opts)

Organization disclosures (climate and diversity)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Organization disclosures (climate and diversity)
  result = api_instance.legal_entities_org_disclosures(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_org_disclosures: #{e}"
end
```

#### Using the legal_entities_org_disclosures_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgDisclosuresResponse>, Integer, Hash)> legal_entities_org_disclosures_with_http_info(id, opts)

```ruby
begin
  # Organization disclosures (climate and diversity)
  data, status_code, headers = api_instance.legal_entities_org_disclosures_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgDisclosuresResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_org_disclosures_with_http_info: #{e}"
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

[**OrgDisclosuresResponse**](OrgDisclosuresResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_org_relationships

> <OrgRelationshipsResponse> legal_entities_org_relationships(id, opts)

Organization relationships

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Organization relationships
  result = api_instance.legal_entities_org_relationships(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_org_relationships: #{e}"
end
```

#### Using the legal_entities_org_relationships_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgRelationshipsResponse>, Integer, Hash)> legal_entities_org_relationships_with_http_info(id, opts)

```ruby
begin
  # Organization relationships
  data, status_code, headers = api_instance.legal_entities_org_relationships_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgRelationshipsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_org_relationships_with_http_info: #{e}"
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

[**OrgRelationshipsResponse**](OrgRelationshipsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_ownership_percentages

> <OwnershipPercentagesResponse> legal_entities_ownership_percentages(id, opts)

Ownership percentages

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Ownership percentages
  result = api_instance.legal_entities_ownership_percentages(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_ownership_percentages: #{e}"
end
```

#### Using the legal_entities_ownership_percentages_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OwnershipPercentagesResponse>, Integer, Hash)> legal_entities_ownership_percentages_with_http_info(id, opts)

```ruby
begin
  # Ownership percentages
  data, status_code, headers = api_instance.legal_entities_ownership_percentages_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OwnershipPercentagesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_ownership_percentages_with_http_info: #{e}"
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

[**OwnershipPercentagesResponse**](OwnershipPercentagesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_ownership_positions

> <OwnershipPositionsResponse> legal_entities_ownership_positions(id, opts)

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

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Ownership and control positions
  result = api_instance.legal_entities_ownership_positions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_ownership_positions: #{e}"
end
```

#### Using the legal_entities_ownership_positions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OwnershipPositionsResponse>, Integer, Hash)> legal_entities_ownership_positions_with_http_info(id, opts)

```ruby
begin
  # Ownership and control positions
  data, status_code, headers = api_instance.legal_entities_ownership_positions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OwnershipPositionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_ownership_positions_with_http_info: #{e}"
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


## legal_entities_pdufa_dates

> <PdufaDatesResponse> legal_entities_pdufa_dates(id, opts)

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

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA decision dates (PDUFA) and catalysts
  result = api_instance.legal_entities_pdufa_dates(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_pdufa_dates: #{e}"
end
```

#### Using the legal_entities_pdufa_dates_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PdufaDatesResponse>, Integer, Hash)> legal_entities_pdufa_dates_with_http_info(id, opts)

```ruby
begin
  # FDA decision dates (PDUFA) and catalysts
  data, status_code, headers = api_instance.legal_entities_pdufa_dates_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PdufaDatesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_pdufa_dates_with_http_info: #{e}"
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


## legal_entities_program_licensing_candidates

> <ProgramLicensingCandidatesResponse> legal_entities_program_licensing_candidates(id, opts)

Program licensing candidates

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Program licensing candidates
  result = api_instance.legal_entities_program_licensing_candidates(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_program_licensing_candidates: #{e}"
end
```

#### Using the legal_entities_program_licensing_candidates_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ProgramLicensingCandidatesResponse>, Integer, Hash)> legal_entities_program_licensing_candidates_with_http_info(id, opts)

```ruby
begin
  # Program licensing candidates
  data, status_code, headers = api_instance.legal_entities_program_licensing_candidates_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ProgramLicensingCandidatesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_program_licensing_candidates_with_http_info: #{e}"
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

[**ProgramLicensingCandidatesResponse**](ProgramLicensingCandidatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_recalls

> <RecallsResponse> legal_entities_recalls(id, opts)

FDA recalls

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA recalls
  result = api_instance.legal_entities_recalls(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_recalls: #{e}"
end
```

#### Using the legal_entities_recalls_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<RecallsResponse>, Integer, Hash)> legal_entities_recalls_with_http_info(id, opts)

```ruby
begin
  # FDA recalls
  data, status_code, headers = api_instance.legal_entities_recalls_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <RecallsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_recalls_with_http_info: #{e}"
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

[**RecallsResponse**](RecallsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_sponsor_recruitment_by_phase

> <SponsorRecruitmentByPhaseResponse> legal_entities_sponsor_recruitment_by_phase(id, opts)

Trial recruitment duration by sponsor and phase

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial recruitment duration by sponsor and phase
  result = api_instance.legal_entities_sponsor_recruitment_by_phase(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_sponsor_recruitment_by_phase: #{e}"
end
```

#### Using the legal_entities_sponsor_recruitment_by_phase_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SponsorRecruitmentByPhaseResponse>, Integer, Hash)> legal_entities_sponsor_recruitment_by_phase_with_http_info(id, opts)

```ruby
begin
  # Trial recruitment duration by sponsor and phase
  data, status_code, headers = api_instance.legal_entities_sponsor_recruitment_by_phase_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SponsorRecruitmentByPhaseResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_sponsor_recruitment_by_phase_with_http_info: #{e}"
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

[**SponsorRecruitmentByPhaseResponse**](SponsorRecruitmentByPhaseResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_sponsor_recruitment_duration

> <SponsorRecruitmentDurationResponse> legal_entities_sponsor_recruitment_duration(id, opts)

Trial recruitment duration by sponsor

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial recruitment duration by sponsor
  result = api_instance.legal_entities_sponsor_recruitment_duration(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_sponsor_recruitment_duration: #{e}"
end
```

#### Using the legal_entities_sponsor_recruitment_duration_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SponsorRecruitmentDurationResponse>, Integer, Hash)> legal_entities_sponsor_recruitment_duration_with_http_info(id, opts)

```ruby
begin
  # Trial recruitment duration by sponsor
  data, status_code, headers = api_instance.legal_entities_sponsor_recruitment_duration_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SponsorRecruitmentDurationResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_sponsor_recruitment_duration_with_http_info: #{e}"
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

[**SponsorRecruitmentDurationResponse**](SponsorRecruitmentDurationResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_sponsor_trial_start_forecast

> <SponsorTrialStartForecastResponse> legal_entities_sponsor_trial_start_forecast(id, opts)

Trial-start forecast by sponsor

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Trial-start forecast by sponsor
  result = api_instance.legal_entities_sponsor_trial_start_forecast(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_sponsor_trial_start_forecast: #{e}"
end
```

#### Using the legal_entities_sponsor_trial_start_forecast_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SponsorTrialStartForecastResponse>, Integer, Hash)> legal_entities_sponsor_trial_start_forecast_with_http_info(id, opts)

```ruby
begin
  # Trial-start forecast by sponsor
  data, status_code, headers = api_instance.legal_entities_sponsor_trial_start_forecast_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SponsorTrialStartForecastResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_sponsor_trial_start_forecast_with_http_info: #{e}"
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

[**SponsorTrialStartForecastResponse**](SponsorTrialStartForecastResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_sponsor_trials

> <SponsorTrialsResponse> legal_entities_sponsor_trials(id, opts)

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

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Sponsor trials
  result = api_instance.legal_entities_sponsor_trials(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_sponsor_trials: #{e}"
end
```

#### Using the legal_entities_sponsor_trials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SponsorTrialsResponse>, Integer, Hash)> legal_entities_sponsor_trials_with_http_info(id, opts)

```ruby
begin
  # Sponsor trials
  data, status_code, headers = api_instance.legal_entities_sponsor_trials_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SponsorTrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_sponsor_trials_with_http_info: #{e}"
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


## legal_entities_supplier_capability_esg

> <SupplierCapabilityEsgResponse> legal_entities_supplier_capability_esg(id, opts)

Supplier capability and ESG disclosure

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Supplier capability and ESG disclosure
  result = api_instance.legal_entities_supplier_capability_esg(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_supplier_capability_esg: #{e}"
end
```

#### Using the legal_entities_supplier_capability_esg_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SupplierCapabilityEsgResponse>, Integer, Hash)> legal_entities_supplier_capability_esg_with_http_info(id, opts)

```ruby
begin
  # Supplier capability and ESG disclosure
  data, status_code, headers = api_instance.legal_entities_supplier_capability_esg_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SupplierCapabilityEsgResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_supplier_capability_esg_with_http_info: #{e}"
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

[**SupplierCapabilityEsgResponse**](SupplierCapabilityEsgResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_unmatched_product_revenue

> <UnmatchedProductRevenueResponse> legal_entities_unmatched_product_revenue(id, opts)

Unmatched product revenue

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # Unmatched product revenue
  result = api_instance.legal_entities_unmatched_product_revenue(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_unmatched_product_revenue: #{e}"
end
```

#### Using the legal_entities_unmatched_product_revenue_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<UnmatchedProductRevenueResponse>, Integer, Hash)> legal_entities_unmatched_product_revenue_with_http_info(id, opts)

```ruby
begin
  # Unmatched product revenue
  data, status_code, headers = api_instance.legal_entities_unmatched_product_revenue_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <UnmatchedProductRevenueResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_unmatched_product_revenue_with_http_info: #{e}"
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

[**UnmatchedProductRevenueResponse**](UnmatchedProductRevenueResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_untitled_letters

> <UntitledLettersResponse> legal_entities_untitled_letters(id, opts)

FDA untitled letters

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA untitled letters
  result = api_instance.legal_entities_untitled_letters(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_untitled_letters: #{e}"
end
```

#### Using the legal_entities_untitled_letters_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<UntitledLettersResponse>, Integer, Hash)> legal_entities_untitled_letters_with_http_info(id, opts)

```ruby
begin
  # FDA untitled letters
  data, status_code, headers = api_instance.legal_entities_untitled_letters_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <UntitledLettersResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_untitled_letters_with_http_info: #{e}"
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

[**UntitledLettersResponse**](UntitledLettersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## legal_entities_warning_letters

> <WarningLettersResponse> legal_entities_warning_letters(id, opts)

FDA warning letters

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::LegalEntitiesApi.new
id = 'edgar:0001114448' # String | 
opts = {
  cursor: 'cursor_example', # String | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it.
  limit: 56, # Integer | Rows per page, up to 500 (default 50).
  sort: 'sort_example', # String | Comma separated columns, a leading minus for descending: sort=a,-b.
  group_by: 'group_by_example', # String | A dimension to group by; a date takes :year, :quarter or :month.
  measure: 'measure_example' # String | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col).
}

begin
  # FDA warning letters
  result = api_instance.legal_entities_warning_letters(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_warning_letters: #{e}"
end
```

#### Using the legal_entities_warning_letters_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<WarningLettersResponse>, Integer, Hash)> legal_entities_warning_letters_with_http_info(id, opts)

```ruby
begin
  # FDA warning letters
  data, status_code, headers = api_instance.legal_entities_warning_letters_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <WarningLettersResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling LegalEntitiesApi->legal_entities_warning_letters_with_http_info: #{e}"
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

[**WarningLettersResponse**](WarningLettersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

