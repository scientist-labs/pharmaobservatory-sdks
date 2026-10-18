# PharmaObservatory::BiomarkersApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**biomarkers_assay_provider_ledger**](BiomarkersApi.md#biomarkers_assay_provider_ledger) | **GET** /api/v1/biomarkers/{id}/assay_provider_ledger | Assay-provider ledger |
| [**biomarkers_cgi**](BiomarkersApi.md#biomarkers_cgi) | **GET** /api/v1/biomarkers/{id}/cgi | Actionable biomarkers |
| [**biomarkers_coverage_by_traversal**](BiomarkersApi.md#biomarkers_coverage_by_traversal) | **GET** /api/v1/biomarkers/{id}/coverage_by_traversal | Coverage by traversal |
| [**biomarkers_evidence_over_time**](BiomarkersApi.md#biomarkers_evidence_over_time) | **GET** /api/v1/biomarkers/{id}/evidence_over_time | Evidence over time |
| [**biomarkers_fda_bqp**](BiomarkersApi.md#biomarkers_fda_bqp) | **GET** /api/v1/biomarkers/{id}/fda_bqp | FDA biomarker qualification projects |
| [**biomarkers_fda_cdx**](BiomarkersApi.md#biomarkers_fda_cdx) | **GET** /api/v1/biomarkers/{id}/fda_cdx | FDA companion diagnostic assertions |
| [**biomarkers_get**](BiomarkersApi.md#biomarkers_get) | **GET** /api/v1/biomarkers/{id} | One biomarkers record |
| [**biomarkers_hub_tier_tables**](BiomarkersApi.md#biomarkers_hub_tier_tables) | **GET** /api/v1/biomarkers/hub_tier_tables | Biomarker evidence, labeled by type |
| [**biomarkers_identity**](BiomarkersApi.md#biomarkers_identity) | **GET** /api/v1/biomarkers/{id}/identity | Biomarker identity |
| [**biomarkers_list**](BiomarkersApi.md#biomarkers_list) | **GET** /api/v1/biomarkers | Biomarkers |
| [**biomarkers_literature**](BiomarkersApi.md#biomarkers_literature) | **GET** /api/v1/biomarkers/{id}/literature | Trial-cited literature mentions |
| [**biomarkers_ot_association**](BiomarkersApi.md#biomarkers_ot_association) | **GET** /api/v1/biomarkers/{id}/ot_association | Target-disease associations |
| [**biomarkers_population_allele_frequency**](BiomarkersApi.md#biomarkers_population_allele_frequency) | **GET** /api/v1/biomarkers/{id}/population_allele_frequency | Population allele frequency |
| [**biomarkers_publishing_org**](BiomarkersApi.md#biomarkers_publishing_org) | **GET** /api/v1/biomarkers/{id}/publishing_org | Organizations publishing |
| [**biomarkers_related_biomarkers**](BiomarkersApi.md#biomarkers_related_biomarkers) | **GET** /api/v1/biomarkers/{id}/related_biomarkers | Related biomarkers |
| [**biomarkers_related_conditions**](BiomarkersApi.md#biomarkers_related_conditions) | **GET** /api/v1/biomarkers/{id}/related_conditions | Related conditions |
| [**biomarkers_related_organizations**](BiomarkersApi.md#biomarkers_related_organizations) | **GET** /api/v1/biomarkers/{id}/related_organizations | Related organisations |
| [**biomarkers_relations**](BiomarkersApi.md#biomarkers_relations) | **GET** /api/v1/biomarkers/{id}/relations | Literature-extracted relations |
| [**biomarkers_sample_screening**](BiomarkersApi.md#biomarkers_sample_screening) | **GET** /api/v1/biomarkers/{id}/sample_screening | Human sample screening reference |
| [**biomarkers_trial_mention**](BiomarkersApi.md#biomarkers_trial_mention) | **GET** /api/v1/biomarkers/{id}/trial_mention | Trial mentions |
| [**biomarkers_variant_mentions**](BiomarkersApi.md#biomarkers_variant_mentions) | **GET** /api/v1/biomarkers/{id}/variant_mentions | Protein variant mentions |
| [**biomarkers_vendor_sourcing**](BiomarkersApi.md#biomarkers_vendor_sourcing) | **GET** /api/v1/biomarkers/{id}/vendor_sourcing | Vendor sourcing |
| [**biomarkers_where_it_appears**](BiomarkersApi.md#biomarkers_where_it_appears) | **GET** /api/v1/biomarkers/{id}/where_it_appears | Where it appears |


## biomarkers_assay_provider_ledger

> <BiomarkersAssayProviderLedger200Response> biomarkers_assay_provider_ledger(id)

Assay-provider ledger

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ALK' # String | 

begin
  # Assay-provider ledger
  result = api_instance.biomarkers_assay_provider_ledger(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_assay_provider_ledger: #{e}"
end
```

#### Using the biomarkers_assay_provider_ledger_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersAssayProviderLedger200Response>, Integer, Hash)> biomarkers_assay_provider_ledger_with_http_info(id)

```ruby
begin
  # Assay-provider ledger
  data, status_code, headers = api_instance.biomarkers_assay_provider_ledger_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersAssayProviderLedger200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_assay_provider_ledger_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersAssayProviderLedger200Response**](BiomarkersAssayProviderLedger200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_cgi

> <BiomarkersCgi200Response> biomarkers_cgi(id)

Actionable biomarkers

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # Actionable biomarkers
  result = api_instance.biomarkers_cgi(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_cgi: #{e}"
end
```

#### Using the biomarkers_cgi_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersCgi200Response>, Integer, Hash)> biomarkers_cgi_with_http_info(id)

```ruby
begin
  # Actionable biomarkers
  data, status_code, headers = api_instance.biomarkers_cgi_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersCgi200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_cgi_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersCgi200Response**](BiomarkersCgi200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_coverage_by_traversal

> <BiomarkersCoverageByTraversal200Response> biomarkers_coverage_by_traversal(id)

Coverage by traversal

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # Coverage by traversal
  result = api_instance.biomarkers_coverage_by_traversal(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_coverage_by_traversal: #{e}"
end
```

#### Using the biomarkers_coverage_by_traversal_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersCoverageByTraversal200Response>, Integer, Hash)> biomarkers_coverage_by_traversal_with_http_info(id)

```ruby
begin
  # Coverage by traversal
  data, status_code, headers = api_instance.biomarkers_coverage_by_traversal_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersCoverageByTraversal200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_coverage_by_traversal_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersCoverageByTraversal200Response**](BiomarkersCoverageByTraversal200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_evidence_over_time

> <BiomarkersEvidenceOverTime200Response> biomarkers_evidence_over_time(id)

Evidence over time

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # Evidence over time
  result = api_instance.biomarkers_evidence_over_time(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_evidence_over_time: #{e}"
end
```

#### Using the biomarkers_evidence_over_time_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersEvidenceOverTime200Response>, Integer, Hash)> biomarkers_evidence_over_time_with_http_info(id)

```ruby
begin
  # Evidence over time
  data, status_code, headers = api_instance.biomarkers_evidence_over_time_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersEvidenceOverTime200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_evidence_over_time_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersEvidenceOverTime200Response**](BiomarkersEvidenceOverTime200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_fda_bqp

> <BiomarkersFdaBqp200Response> biomarkers_fda_bqp(id)

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

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'bqp.safety-biomarker-to-be-used-with-traditional-ind.2516104f1f' # String | 

begin
  # FDA biomarker qualification projects
  result = api_instance.biomarkers_fda_bqp(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_fda_bqp: #{e}"
end
```

#### Using the biomarkers_fda_bqp_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersFdaBqp200Response>, Integer, Hash)> biomarkers_fda_bqp_with_http_info(id)

```ruby
begin
  # FDA biomarker qualification projects
  data, status_code, headers = api_instance.biomarkers_fda_bqp_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersFdaBqp200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_fda_bqp_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersFdaBqp200Response**](BiomarkersFdaBqp200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_fda_cdx

> <BiomarkersFdaCdx200Response> biomarkers_fda_cdx(id)

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

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'f.pd-l1' # String | 

begin
  # FDA companion diagnostic assertions
  result = api_instance.biomarkers_fda_cdx(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_fda_cdx: #{e}"
end
```

#### Using the biomarkers_fda_cdx_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersFdaCdx200Response>, Integer, Hash)> biomarkers_fda_cdx_with_http_info(id)

```ruby
begin
  # FDA companion diagnostic assertions
  data, status_code, headers = api_instance.biomarkers_fda_cdx_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersFdaCdx200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_fda_cdx_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersFdaCdx200Response**](BiomarkersFdaCdx200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_get

> <BiomarkersGet200Response> biomarkers_get(id)

One biomarkers record

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # One biomarkers record
  result = api_instance.biomarkers_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_get: #{e}"
end
```

#### Using the biomarkers_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersGet200Response>, Integer, Hash)> biomarkers_get_with_http_info(id)

```ruby
begin
  # One biomarkers record
  data, status_code, headers = api_instance.biomarkers_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersGet200Response**](BiomarkersGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_hub_tier_tables

> <BiomarkersHubTierTables200Response> biomarkers_hub_tier_tables

Biomarker evidence, labeled by type

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new

begin
  # Biomarker evidence, labeled by type
  result = api_instance.biomarkers_hub_tier_tables
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_hub_tier_tables: #{e}"
end
```

#### Using the biomarkers_hub_tier_tables_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersHubTierTables200Response>, Integer, Hash)> biomarkers_hub_tier_tables_with_http_info

```ruby
begin
  # Biomarker evidence, labeled by type
  data, status_code, headers = api_instance.biomarkers_hub_tier_tables_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersHubTierTables200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_hub_tier_tables_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**BiomarkersHubTierTables200Response**](BiomarkersHubTierTables200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_identity

> <BiomarkersIdentity200Response> biomarkers_identity(id)

Biomarker identity

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # Biomarker identity
  result = api_instance.biomarkers_identity(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_identity: #{e}"
end
```

#### Using the biomarkers_identity_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersIdentity200Response>, Integer, Hash)> biomarkers_identity_with_http_info(id)

```ruby
begin
  # Biomarker identity
  data, status_code, headers = api_instance.biomarkers_identity_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersIdentity200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_identity_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersIdentity200Response**](BiomarkersIdentity200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_list

> <BiomarkersList200Response> biomarkers_list(opts)

Biomarkers

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
opts = {
  cursor: 'eyJwIjoyLCJsIjozLCJmIjoic3BlYy1maW5nZXJwcmludCJ9', # String | 
  dir: 'asc', # String | 
  evidence: 0, # Integer | 
  facets: true, # Boolean | 
  identity: 'resolved', # String | 
  kind: 'gene', # String | 
  limit: 1, # Integer | 
  q: 'BRAF', # String | 
  sort: 'name', # String | 
  validated: 1 # Integer | 
}

begin
  # Biomarkers
  result = api_instance.biomarkers_list(opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_list: #{e}"
end
```

#### Using the biomarkers_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersList200Response>, Integer, Hash)> biomarkers_list_with_http_info(opts)

```ruby
begin
  # Biomarkers
  data, status_code, headers = api_instance.biomarkers_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersList200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_list_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cursor** | **String** |  | [optional] |
| **dir** | **String** |  | [optional] |
| **evidence** | **Integer** |  | [optional] |
| **facets** | **Boolean** |  | [optional] |
| **identity** | **String** |  | [optional] |
| **kind** | **String** |  | [optional] |
| **limit** | **Integer** |  | [optional] |
| **q** | **String** |  | [optional] |
| **sort** | **String** |  | [optional] |
| **validated** | **Integer** |  | [optional] |

### Return type

[**BiomarkersList200Response**](BiomarkersList200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_literature

> <BiomarkersLiterature200Response> biomarkers_literature(id)

Trial-cited literature mentions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ERBB2' # String | 

begin
  # Trial-cited literature mentions
  result = api_instance.biomarkers_literature(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_literature: #{e}"
end
```

#### Using the biomarkers_literature_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersLiterature200Response>, Integer, Hash)> biomarkers_literature_with_http_info(id)

```ruby
begin
  # Trial-cited literature mentions
  data, status_code, headers = api_instance.biomarkers_literature_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersLiterature200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_literature_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersLiterature200Response**](BiomarkersLiterature200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_ot_association

> <BiomarkersOtAssociation200Response> biomarkers_ot_association(id)

Target-disease associations

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'MMP9' # String | 

begin
  # Target-disease associations
  result = api_instance.biomarkers_ot_association(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_ot_association: #{e}"
end
```

#### Using the biomarkers_ot_association_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersOtAssociation200Response>, Integer, Hash)> biomarkers_ot_association_with_http_info(id)

```ruby
begin
  # Target-disease associations
  data, status_code, headers = api_instance.biomarkers_ot_association_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersOtAssociation200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_ot_association_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersOtAssociation200Response**](BiomarkersOtAssociation200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_population_allele_frequency

> <BiomarkersPopulationAlleleFrequency200Response> biomarkers_population_allele_frequency(id)

Population allele frequency

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'v.a1bg.p-a170t.a6cb074848' # String | 

begin
  # Population allele frequency
  result = api_instance.biomarkers_population_allele_frequency(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_population_allele_frequency: #{e}"
end
```

#### Using the biomarkers_population_allele_frequency_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersPopulationAlleleFrequency200Response>, Integer, Hash)> biomarkers_population_allele_frequency_with_http_info(id)

```ruby
begin
  # Population allele frequency
  data, status_code, headers = api_instance.biomarkers_population_allele_frequency_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersPopulationAlleleFrequency200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_population_allele_frequency_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersPopulationAlleleFrequency200Response**](BiomarkersPopulationAlleleFrequency200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_publishing_org

> <BiomarkersPublishingOrg200Response> biomarkers_publishing_org(id)

Organizations publishing

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # Organizations publishing
  result = api_instance.biomarkers_publishing_org(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_publishing_org: #{e}"
end
```

#### Using the biomarkers_publishing_org_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersPublishingOrg200Response>, Integer, Hash)> biomarkers_publishing_org_with_http_info(id)

```ruby
begin
  # Organizations publishing
  data, status_code, headers = api_instance.biomarkers_publishing_org_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersPublishingOrg200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_publishing_org_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersPublishingOrg200Response**](BiomarkersPublishingOrg200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_related_biomarkers

> <BiomarkersRelatedBiomarkers200Response> biomarkers_related_biomarkers(id)

Related biomarkers

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # Related biomarkers
  result = api_instance.biomarkers_related_biomarkers(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_related_biomarkers: #{e}"
end
```

#### Using the biomarkers_related_biomarkers_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersRelatedBiomarkers200Response>, Integer, Hash)> biomarkers_related_biomarkers_with_http_info(id)

```ruby
begin
  # Related biomarkers
  data, status_code, headers = api_instance.biomarkers_related_biomarkers_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersRelatedBiomarkers200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_related_biomarkers_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersRelatedBiomarkers200Response**](BiomarkersRelatedBiomarkers200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_related_conditions

> <BiomarkersRelatedConditions200Response> biomarkers_related_conditions(id)

Related conditions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # Related conditions
  result = api_instance.biomarkers_related_conditions(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_related_conditions: #{e}"
end
```

#### Using the biomarkers_related_conditions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersRelatedConditions200Response>, Integer, Hash)> biomarkers_related_conditions_with_http_info(id)

```ruby
begin
  # Related conditions
  data, status_code, headers = api_instance.biomarkers_related_conditions_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersRelatedConditions200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_related_conditions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersRelatedConditions200Response**](BiomarkersRelatedConditions200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_related_organizations

> <BiomarkersRelatedOrganizations200Response> biomarkers_related_organizations(id)

Related organisations

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # Related organisations
  result = api_instance.biomarkers_related_organizations(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_related_organizations: #{e}"
end
```

#### Using the biomarkers_related_organizations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersRelatedOrganizations200Response>, Integer, Hash)> biomarkers_related_organizations_with_http_info(id)

```ruby
begin
  # Related organisations
  data, status_code, headers = api_instance.biomarkers_related_organizations_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersRelatedOrganizations200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_related_organizations_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersRelatedOrganizations200Response**](BiomarkersRelatedOrganizations200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_relations

> <BiomarkersRelations200Response> biomarkers_relations(id, opts)

Literature-extracted relations

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'BRAF' # String | 
opts = {
  chemicals: 'x' # String | 
}

begin
  # Literature-extracted relations
  result = api_instance.biomarkers_relations(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_relations: #{e}"
end
```

#### Using the biomarkers_relations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersRelations200Response>, Integer, Hash)> biomarkers_relations_with_http_info(id, opts)

```ruby
begin
  # Literature-extracted relations
  data, status_code, headers = api_instance.biomarkers_relations_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersRelations200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_relations_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **chemicals** | **String** |  | [optional] |

### Return type

[**BiomarkersRelations200Response**](BiomarkersRelations200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_sample_screening

> <BiomarkersSampleScreening200Response> biomarkers_sample_screening(id, opts)

Human sample screening reference

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'BRAF' # String | 
opts = {
  positive_samples: 20 # Integer | 
}

begin
  # Human sample screening reference
  result = api_instance.biomarkers_sample_screening(id, opts)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_sample_screening: #{e}"
end
```

#### Using the biomarkers_sample_screening_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersSampleScreening200Response>, Integer, Hash)> biomarkers_sample_screening_with_http_info(id, opts)

```ruby
begin
  # Human sample screening reference
  data, status_code, headers = api_instance.biomarkers_sample_screening_with_http_info(id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersSampleScreening200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_sample_screening_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **positive_samples** | **Integer** |  | [optional] |

### Return type

[**BiomarkersSampleScreening200Response**](BiomarkersSampleScreening200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_trial_mention

> <BiomarkersTrialMention200Response> biomarkers_trial_mention(id)

Trial mentions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ALB' # String | 

begin
  # Trial mentions
  result = api_instance.biomarkers_trial_mention(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_trial_mention: #{e}"
end
```

#### Using the biomarkers_trial_mention_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersTrialMention200Response>, Integer, Hash)> biomarkers_trial_mention_with_http_info(id)

```ruby
begin
  # Trial mentions
  data, status_code, headers = api_instance.biomarkers_trial_mention_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersTrialMention200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_trial_mention_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersTrialMention200Response**](BiomarkersTrialMention200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_variant_mentions

> <BiomarkersVariantMentions200Response> biomarkers_variant_mentions(id)

Protein variant mentions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'BRAF' # String | 

begin
  # Protein variant mentions
  result = api_instance.biomarkers_variant_mentions(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_variant_mentions: #{e}"
end
```

#### Using the biomarkers_variant_mentions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersVariantMentions200Response>, Integer, Hash)> biomarkers_variant_mentions_with_http_info(id)

```ruby
begin
  # Protein variant mentions
  data, status_code, headers = api_instance.biomarkers_variant_mentions_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersVariantMentions200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_variant_mentions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersVariantMentions200Response**](BiomarkersVariantMentions200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_vendor_sourcing

> <BiomarkersVendorSourcing200Response> biomarkers_vendor_sourcing(id)

Vendor sourcing

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ALB' # String | 

begin
  # Vendor sourcing
  result = api_instance.biomarkers_vendor_sourcing(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_vendor_sourcing: #{e}"
end
```

#### Using the biomarkers_vendor_sourcing_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersVendorSourcing200Response>, Integer, Hash)> biomarkers_vendor_sourcing_with_http_info(id)

```ruby
begin
  # Vendor sourcing
  data, status_code, headers = api_instance.biomarkers_vendor_sourcing_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersVendorSourcing200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_vendor_sourcing_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersVendorSourcing200Response**](BiomarkersVendorSourcing200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## biomarkers_where_it_appears

> <BiomarkersWhereItAppears200Response> biomarkers_where_it_appears(id)

Where it appears

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::BiomarkersApi.new
id = 'ESR1' # String | 

begin
  # Where it appears
  result = api_instance.biomarkers_where_it_appears(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_where_it_appears: #{e}"
end
```

#### Using the biomarkers_where_it_appears_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<BiomarkersWhereItAppears200Response>, Integer, Hash)> biomarkers_where_it_appears_with_http_info(id)

```ruby
begin
  # Where it appears
  data, status_code, headers = api_instance.biomarkers_where_it_appears_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <BiomarkersWhereItAppears200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling BiomarkersApi->biomarkers_where_it_appears_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**BiomarkersWhereItAppears200Response**](BiomarkersWhereItAppears200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

