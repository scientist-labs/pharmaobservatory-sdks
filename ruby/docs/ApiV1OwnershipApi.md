# PharmaObservatory::ApiV1OwnershipApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**api_v1_investor_profile_get**](ApiV1OwnershipApi.md#api_v1_investor_profile_get) | **GET** /api/v1/investor_profile | investor_profile |
| [**api_v1_org_financing_and_ownership_get**](ApiV1OwnershipApi.md#api_v1_org_financing_and_ownership_get) | **GET** /api/v1/org_financing_and_ownership | org_financing_and_ownership |


## api_v1_investor_profile_get

> <ApiV1InvestorProfileGet200Response> api_v1_investor_profile_get(opts)

investor_profile

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1OwnershipApi.new
opts = {
  investor_id: 'adv:283882', # String | 
  investor_name: 'Bioventure Capital Advisers' # String | 
}

begin
  # investor_profile
  result = api_instance.api_v1_investor_profile_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1OwnershipApi->api_v1_investor_profile_get: #{e}"
end
```

#### Using the api_v1_investor_profile_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1InvestorProfileGet200Response>, Integer, Hash)> api_v1_investor_profile_get_with_http_info(opts)

```ruby
begin
  # investor_profile
  data, status_code, headers = api_instance.api_v1_investor_profile_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1InvestorProfileGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1OwnershipApi->api_v1_investor_profile_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investor_id** | **String** |  | [optional] |
| **investor_name** | **String** |  | [optional] |

### Return type

[**ApiV1InvestorProfileGet200Response**](ApiV1InvestorProfileGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## api_v1_org_financing_and_ownership_get

> <ApiV1OrgFinancingAndOwnershipGet200Response> api_v1_org_financing_and_ownership_get(opts)

org_financing_and_ownership

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::ApiV1OwnershipApi.new
opts = {
  edgar_cik: 1114448, # Integer | 
  org_id: 'edgar:0001048477', # String | 
  org_name: 'Novartis Ag', # String | 
  ror_id: '01hcyya48' # String | 
}

begin
  # org_financing_and_ownership
  result = api_instance.api_v1_org_financing_and_ownership_get(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1OwnershipApi->api_v1_org_financing_and_ownership_get: #{e}"
end
```

#### Using the api_v1_org_financing_and_ownership_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiV1OrgFinancingAndOwnershipGet200Response>, Integer, Hash)> api_v1_org_financing_and_ownership_get_with_http_info(opts)

```ruby
begin
  # org_financing_and_ownership
  data, status_code, headers = api_instance.api_v1_org_financing_and_ownership_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiV1OrgFinancingAndOwnershipGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling ApiV1OwnershipApi->api_v1_org_financing_and_ownership_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **edgar_cik** | **Integer** |  | [optional] |
| **org_id** | **String** |  | [optional] |
| **org_name** | **String** |  | [optional] |
| **ror_id** | **String** |  | [optional] |

### Return type

[**ApiV1OrgFinancingAndOwnershipGet200Response**](ApiV1OrgFinancingAndOwnershipGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

