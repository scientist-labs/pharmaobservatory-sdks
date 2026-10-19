# PharmaObservatory::CompaniesApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**companies_biomarker_bqp_assertions**](CompaniesApi.md#companies_biomarker_bqp_assertions) | **GET** /api/v1/companies/{id}/biomarker_bqp_assertions | FDA biomarker qualification projects |
| [**companies_biomarker_cdx_assertions**](CompaniesApi.md#companies_biomarker_cdx_assertions) | **GET** /api/v1/companies/{id}/biomarker_cdx_assertions | FDA companion diagnostic assertions |
| [**companies_biomarker_publishing_companies**](CompaniesApi.md#companies_biomarker_publishing_companies) | **GET** /api/v1/companies/{id}/biomarker_publishing_companies | Companies publishing on a biomarker |
| [**companies_company_members**](CompaniesApi.md#companies_company_members) | **GET** /api/v1/companies/{id}/company_members | Company members |
| [**companies_company_programs**](CompaniesApi.md#companies_company_programs) | **GET** /api/v1/companies/{id}/company_programs | Company programs |
| [**companies_company_targets**](CompaniesApi.md#companies_company_targets) | **GET** /api/v1/companies/{id}/company_targets | Targets a company pursues |
| [**companies_drug_patent_cliff**](CompaniesApi.md#companies_drug_patent_cliff) | **GET** /api/v1/companies/{id}/drug_patent_cliff | Patent-cliff exposure |
| [**companies_get**](CompaniesApi.md#companies_get) | **GET** /api/v1/companies/{id} | One company |
| [**companies_institutional_holdings**](CompaniesApi.md#companies_institutional_holdings) | **GET** /api/v1/companies/{id}/institutional_holdings | Institutional holdings (SEC Form 13F) |
| [**companies_legal_entities**](CompaniesApi.md#companies_legal_entities) | **GET** /api/v1/companies/{id}/legal_entities | Legal entities |
| [**companies_list**](CompaniesApi.md#companies_list) | **GET** /api/v1/companies | List companies |
| [**companies_pdufa_dates**](CompaniesApi.md#companies_pdufa_dates) | **GET** /api/v1/companies/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts |
| [**companies_sponsor_recruitment_by_phase**](CompaniesApi.md#companies_sponsor_recruitment_by_phase) | **GET** /api/v1/companies/{id}/sponsor_recruitment_by_phase | Trial recruitment duration by sponsor and phase |
| [**companies_sponsor_recruitment_duration**](CompaniesApi.md#companies_sponsor_recruitment_duration) | **GET** /api/v1/companies/{id}/sponsor_recruitment_duration | Trial recruitment duration by sponsor |
| [**companies_sponsor_trial_start_forecast**](CompaniesApi.md#companies_sponsor_trial_start_forecast) | **GET** /api/v1/companies/{id}/sponsor_trial_start_forecast | Trial-start forecast by sponsor |
| [**companies_sponsor_trials**](CompaniesApi.md#companies_sponsor_trials) | **GET** /api/v1/companies/{id}/sponsor_trials | Sponsor trials |


## companies_biomarker_bqp_assertions

> <BiomarkerBqpAssertionsResponse> companies_biomarker_bqp_assertions(id, opts)

FDA biomarker qualification projects

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'route_key' # String | 
}

begin
  # FDA biomarker qualification projects
  result = api_instance.companies_biomarker_bqp_assertions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_biomarker_bqp_assertions: #{e}"
end
```

#### Using the companies_biomarker_bqp_assertions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerBqpAssertionsResponse>, Integer, Hash)> companies_biomarker_bqp_assertions_with_http_info(id, opts)

```ruby
begin
  # FDA biomarker qualification projects
  data, status_code, headers = api_instance.companies_biomarker_bqp_assertions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerBqpAssertionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_biomarker_bqp_assertions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**BiomarkerBqpAssertionsResponse**](BiomarkerBqpAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_biomarker_cdx_assertions

> <BiomarkerCdxAssertionsResponse> companies_biomarker_cdx_assertions(id, opts)

FDA companion diagnostic assertions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'biomarker_id' # String | 
}

begin
  # FDA companion diagnostic assertions
  result = api_instance.companies_biomarker_cdx_assertions(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_biomarker_cdx_assertions: #{e}"
end
```

#### Using the companies_biomarker_cdx_assertions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerCdxAssertionsResponse>, Integer, Hash)> companies_biomarker_cdx_assertions_with_http_info(id, opts)

```ruby
begin
  # FDA companion diagnostic assertions
  data, status_code, headers = api_instance.companies_biomarker_cdx_assertions_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerCdxAssertionsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_biomarker_cdx_assertions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**BiomarkerCdxAssertionsResponse**](BiomarkerCdxAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_biomarker_publishing_companies

> <BiomarkerPublishingCompaniesResponse> companies_biomarker_publishing_companies(id, opts)

Companies publishing on a biomarker

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'n_publications' # String | 
}

begin
  # Companies publishing on a biomarker
  result = api_instance.companies_biomarker_publishing_companies(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_biomarker_publishing_companies: #{e}"
end
```

#### Using the companies_biomarker_publishing_companies_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkerPublishingCompaniesResponse>, Integer, Hash)> companies_biomarker_publishing_companies_with_http_info(id, opts)

```ruby
begin
  # Companies publishing on a biomarker
  data, status_code, headers = api_instance.companies_biomarker_publishing_companies_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkerPublishingCompaniesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_biomarker_publishing_companies_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**BiomarkerPublishingCompaniesResponse**](BiomarkerPublishingCompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_company_members

> <CompanyMembersResponse> companies_company_members(id, opts)

Company members

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'company_id' # String | 
}

begin
  # Company members
  result = api_instance.companies_company_members(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_company_members: #{e}"
end
```

#### Using the companies_company_members_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CompanyMembersResponse>, Integer, Hash)> companies_company_members_with_http_info(id, opts)

```ruby
begin
  # Company members
  data, status_code, headers = api_instance.companies_company_members_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CompanyMembersResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_company_members_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**CompanyMembersResponse**](CompanyMembersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_company_programs

> <CompanyProgramsResponse> companies_company_programs(id, opts)

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

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'max_phase' # String | 
}

begin
  # Company programs
  result = api_instance.companies_company_programs(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_company_programs: #{e}"
end
```

#### Using the companies_company_programs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CompanyProgramsResponse>, Integer, Hash)> companies_company_programs_with_http_info(id, opts)

```ruby
begin
  # Company programs
  data, status_code, headers = api_instance.companies_company_programs_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CompanyProgramsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_company_programs_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**CompanyProgramsResponse**](CompanyProgramsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_company_targets

> <CompanyTargetsResponse> companies_company_targets(id, opts)

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

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'max_phase' # String | 
}

begin
  # Targets a company pursues
  result = api_instance.companies_company_targets(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_company_targets: #{e}"
end
```

#### Using the companies_company_targets_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CompanyTargetsResponse>, Integer, Hash)> companies_company_targets_with_http_info(id, opts)

```ruby
begin
  # Targets a company pursues
  data, status_code, headers = api_instance.companies_company_targets_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CompanyTargetsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_company_targets_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**CompanyTargetsResponse**](CompanyTargetsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_drug_patent_cliff

> <DrugPatentCliffResponse> companies_drug_patent_cliff(id, opts)

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

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'cliff_date' # String | 
}

begin
  # Patent-cliff exposure
  result = api_instance.companies_drug_patent_cliff(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_drug_patent_cliff: #{e}"
end
```

#### Using the companies_drug_patent_cliff_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DrugPatentCliffResponse>, Integer, Hash)> companies_drug_patent_cliff_with_http_info(id, opts)

```ruby
begin
  # Patent-cliff exposure
  data, status_code, headers = api_instance.companies_drug_patent_cliff_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DrugPatentCliffResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_drug_patent_cliff_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**DrugPatentCliffResponse**](DrugPatentCliffResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_get

> <CompaniesRecord> companies_get(id)

One company

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 

begin
  # One company
  result = api_instance.companies_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_get: #{e}"
end
```

#### Using the companies_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CompaniesRecord>, Integer, Hash)> companies_get_with_http_info(id)

```ruby
begin
  # One company
  data, status_code, headers = api_instance.companies_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CompaniesRecord>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**CompaniesRecord**](CompaniesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_institutional_holdings

> <InstitutionalHoldingsResponse> companies_institutional_holdings(id, opts)

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

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'reported_value_amount' # String | 
}

begin
  # Institutional holdings (SEC Form 13F)
  result = api_instance.companies_institutional_holdings(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_institutional_holdings: #{e}"
end
```

#### Using the companies_institutional_holdings_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InstitutionalHoldingsResponse>, Integer, Hash)> companies_institutional_holdings_with_http_info(id, opts)

```ruby
begin
  # Institutional holdings (SEC Form 13F)
  data, status_code, headers = api_instance.companies_institutional_holdings_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InstitutionalHoldingsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_institutional_holdings_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**InstitutionalHoldingsResponse**](InstitutionalHoldingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_legal_entities

> <LegalEntitiesResponse> companies_legal_entities(id, opts)

Legal entities

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'activity_score' # String | 
}

begin
  # Legal entities
  result = api_instance.companies_legal_entities(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_legal_entities: #{e}"
end
```

#### Using the companies_legal_entities_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<LegalEntitiesResponse>, Integer, Hash)> companies_legal_entities_with_http_info(id, opts)

```ruby
begin
  # Legal entities
  data, status_code, headers = api_instance.companies_legal_entities_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <LegalEntitiesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_legal_entities_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**LegalEntitiesResponse**](LegalEntitiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_list

> <CompaniesResponse> companies_list(opts)

List companies

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::CompaniesApi.new
opts = {
  limit: 2, # Integer | 
  sort: 'activity_score' # String | 
}

begin
  # List companies
  result = api_instance.companies_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_list: #{e}"
end
```

#### Using the companies_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CompaniesResponse>, Integer, Hash)> companies_list_with_http_info(opts)

```ruby
begin
  # List companies
  data, status_code, headers = api_instance.companies_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CompaniesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_list_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**CompaniesResponse**](CompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_pdufa_dates

> <PdufaDatesResponse> companies_pdufa_dates(id, opts)

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

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'decision_date' # String | 
}

begin
  # FDA decision dates (PDUFA) and catalysts
  result = api_instance.companies_pdufa_dates(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_pdufa_dates: #{e}"
end
```

#### Using the companies_pdufa_dates_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PdufaDatesResponse>, Integer, Hash)> companies_pdufa_dates_with_http_info(id, opts)

```ruby
begin
  # FDA decision dates (PDUFA) and catalysts
  data, status_code, headers = api_instance.companies_pdufa_dates_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PdufaDatesResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_pdufa_dates_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**PdufaDatesResponse**](PdufaDatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_sponsor_recruitment_by_phase

> <SponsorRecruitmentByPhaseResponse> companies_sponsor_recruitment_by_phase(id, opts)

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

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'sponsor_org_id' # String | 
}

begin
  # Trial recruitment duration by sponsor and phase
  result = api_instance.companies_sponsor_recruitment_by_phase(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_sponsor_recruitment_by_phase: #{e}"
end
```

#### Using the companies_sponsor_recruitment_by_phase_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SponsorRecruitmentByPhaseResponse>, Integer, Hash)> companies_sponsor_recruitment_by_phase_with_http_info(id, opts)

```ruby
begin
  # Trial recruitment duration by sponsor and phase
  data, status_code, headers = api_instance.companies_sponsor_recruitment_by_phase_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SponsorRecruitmentByPhaseResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_sponsor_recruitment_by_phase_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**SponsorRecruitmentByPhaseResponse**](SponsorRecruitmentByPhaseResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_sponsor_recruitment_duration

> <SponsorRecruitmentDurationResponse> companies_sponsor_recruitment_duration(id, opts)

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

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'sponsor_org_id' # String | 
}

begin
  # Trial recruitment duration by sponsor
  result = api_instance.companies_sponsor_recruitment_duration(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_sponsor_recruitment_duration: #{e}"
end
```

#### Using the companies_sponsor_recruitment_duration_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SponsorRecruitmentDurationResponse>, Integer, Hash)> companies_sponsor_recruitment_duration_with_http_info(id, opts)

```ruby
begin
  # Trial recruitment duration by sponsor
  data, status_code, headers = api_instance.companies_sponsor_recruitment_duration_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SponsorRecruitmentDurationResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_sponsor_recruitment_duration_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**SponsorRecruitmentDurationResponse**](SponsorRecruitmentDurationResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_sponsor_trial_start_forecast

> <SponsorTrialStartForecastResponse> companies_sponsor_trial_start_forecast(id, opts)

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

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'sponsor_org_id' # String | 
}

begin
  # Trial-start forecast by sponsor
  result = api_instance.companies_sponsor_trial_start_forecast(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_sponsor_trial_start_forecast: #{e}"
end
```

#### Using the companies_sponsor_trial_start_forecast_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SponsorTrialStartForecastResponse>, Integer, Hash)> companies_sponsor_trial_start_forecast_with_http_info(id, opts)

```ruby
begin
  # Trial-start forecast by sponsor
  data, status_code, headers = api_instance.companies_sponsor_trial_start_forecast_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SponsorTrialStartForecastResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_sponsor_trial_start_forecast_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**SponsorTrialStartForecastResponse**](SponsorTrialStartForecastResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## companies_sponsor_trials

> <SponsorTrialsResponse> companies_sponsor_trials(id, opts)

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

api_instance = PharmaObservatory::CompaniesApi.new
id = 'company:c0185cf6c39a2536ad436cf6255494c0' # String | 
opts = {
  limit: 2, # Integer | 
  scope: 'company', # String | 
  sort: 'start_date' # String | 
}

begin
  # Sponsor trials
  result = api_instance.companies_sponsor_trials(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_sponsor_trials: #{e}"
end
```

#### Using the companies_sponsor_trials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SponsorTrialsResponse>, Integer, Hash)> companies_sponsor_trials_with_http_info(id, opts)

```ruby
begin
  # Sponsor trials
  data, status_code, headers = api_instance.companies_sponsor_trials_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SponsorTrialsResponse>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling CompaniesApi->companies_sponsor_trials_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **limit** | **Integer** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**SponsorTrialsResponse**](SponsorTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

