# PharmaObservatory::LegalEntitiesApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**legal_entities_company_programs**](LegalEntitiesApi.md#legal_entities_company_programs) | **GET** /api/v1/legal_entities/{id}/company_programs | Company programs |
| [**legal_entities_drug_patent_cliff**](LegalEntitiesApi.md#legal_entities_drug_patent_cliff) | **GET** /api/v1/legal_entities/{id}/drug_patent_cliff | Patent-cliff exposure |
| [**legal_entities_get**](LegalEntitiesApi.md#legal_entities_get) | **GET** /api/v1/legal_entities/{id} | One legal entity |
| [**legal_entities_institutional_holdings**](LegalEntitiesApi.md#legal_entities_institutional_holdings) | **GET** /api/v1/legal_entities/{id}/institutional_holdings | Institutional holdings (SEC Form 13F) |
| [**legal_entities_list**](LegalEntitiesApi.md#legal_entities_list) | **GET** /api/v1/legal_entities | List legal entities |
| [**legal_entities_pdufa_dates**](LegalEntitiesApi.md#legal_entities_pdufa_dates) | **GET** /api/v1/legal_entities/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts |
| [**legal_entities_sponsor_recruitment_by_phase**](LegalEntitiesApi.md#legal_entities_sponsor_recruitment_by_phase) | **GET** /api/v1/legal_entities/{id}/sponsor_recruitment_by_phase | Trial recruitment duration by sponsor and phase |
| [**legal_entities_sponsor_recruitment_duration**](LegalEntitiesApi.md#legal_entities_sponsor_recruitment_duration) | **GET** /api/v1/legal_entities/{id}/sponsor_recruitment_duration | Trial recruitment duration by sponsor |
| [**legal_entities_sponsor_trial_start_forecast**](LegalEntitiesApi.md#legal_entities_sponsor_trial_start_forecast) | **GET** /api/v1/legal_entities/{id}/sponsor_trial_start_forecast | Trial-start forecast by sponsor |
| [**legal_entities_sponsor_trials**](LegalEntitiesApi.md#legal_entities_sponsor_trials) | **GET** /api/v1/legal_entities/{id}/sponsor_trials | Sponsor trials |


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
  limit: 2, # Integer | 
  sort: 'max_phase' # String | 
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
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**CompanyProgramsResponse**](CompanyProgramsResponse.md)

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
  limit: 2, # Integer | 
  sort: 'cliff_date' # String | 
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
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**DrugPatentCliffResponse**](DrugPatentCliffResponse.md)

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
  limit: 2, # Integer | 
  sort: 'reported_value_amount' # String | 
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
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**InstitutionalHoldingsResponse**](InstitutionalHoldingsResponse.md)

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
  limit: 2, # Integer | 
  sort: 'activity_score' # String | 
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
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**LegalEntitiesResponse**](LegalEntitiesResponse.md)

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
  limit: 2, # Integer | 
  sort: 'decision_date' # String | 
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
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**PdufaDatesResponse**](PdufaDatesResponse.md)

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
  limit: 2, # Integer | 
  sort: 'sponsor_org_id' # String | 
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
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

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
  limit: 2, # Integer | 
  sort: 'sponsor_org_id' # String | 
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
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

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
  limit: 2, # Integer | 
  sort: 'sponsor_org_id' # String | 
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
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

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
  limit: 2, # Integer | 
  sort: 'start_date' # String | 
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
| **limit** | **Integer** |  | [optional] |
| **sort** | **String** |  | [optional] |

### Return type

[**SponsorTrialsResponse**](SponsorTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

