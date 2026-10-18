# PharmaObservatory::OrgsApi

All URIs are relative to *https://pharmaobservatory.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**orgs_active_programs**](OrgsApi.md#orgs_active_programs) | **GET** /api/v1/orgs/{id}/active_programs | Active development programs |
| [**orgs_activity_summary**](OrgsApi.md#orgs_activity_summary) | **GET** /api/v1/orgs/{id}/activity_summary | Activity summary |
| [**orgs_approved_assets**](OrgsApi.md#orgs_approved_assets) | **GET** /api/v1/orgs/{id}/approved_assets | Approved assets |
| [**orgs_benzinga_catalysts**](OrgsApi.md#orgs_benzinga_catalysts) | **GET** /api/v1/orgs/{id}/benzinga_catalysts | Press-release catalysts |
| [**orgs_biomarker_capability**](OrgsApi.md#orgs_biomarker_capability) | **GET** /api/v1/orgs/{id}/biomarker_capability | Biomarker capability evidence |
| [**orgs_clia**](OrgsApi.md#orgs_clia) | **GET** /api/v1/orgs/{id}/clia | CLIA clinical laboratory certificates |
| [**orgs_climate**](OrgsApi.md#orgs_climate) | **GET** /api/v1/orgs/{id}/climate | Climate targets, as filed |
| [**orgs_collaborators**](OrgsApi.md#orgs_collaborators) | **GET** /api/v1/orgs/{id}/collaborators | Study collaborators |
| [**orgs_company**](OrgsApi.md#orgs_company) | **GET** /api/v1/orgs/{id}/company | Company |
| [**orgs_competitor_activity**](OrgsApi.md#orgs_competitor_activity) | **GET** /api/v1/orgs/{id}/competitor_activity | Competitor activity |
| [**orgs_competitors**](OrgsApi.md#orgs_competitors) | **GET** /api/v1/orgs/{id}/competitors | Competitors |
| [**orgs_coverage**](OrgsApi.md#orgs_coverage) | **GET** /api/v1/orgs/{id}/coverage | Coverage of this page |
| [**orgs_current_ownership**](OrgsApi.md#orgs_current_ownership) | **GET** /api/v1/orgs/{id}/current_ownership | Current ownership and control |
| [**orgs_declarations**](OrgsApi.md#orgs_declarations) | **GET** /api/v1/orgs/{id}/declarations | Organization-wide declarations |
| [**orgs_delivery**](OrgsApi.md#orgs_delivery) | **GET** /api/v1/orgs/{id}/delivery | Delivery reliability |
| [**orgs_description**](OrgsApi.md#orgs_description) | **GET** /api/v1/orgs/{id}/description | Description |
| [**orgs_disease_areas**](OrgsApi.md#orgs_disease_areas) | **GET** /api/v1/orgs/{id}/disease_areas | Disease areas |
| [**orgs_diversity**](OrgsApi.md#orgs_diversity) | **GET** /api/v1/orgs/{id}/diversity | Workforce and board diversity, as filed |
| [**orgs_esg_trend**](OrgsApi.md#orgs_esg_trend) | **GET** /api/v1/orgs/{id}/esg_trend | Climate and diversity by reporting year |
| [**orgs_financial_line_items**](OrgsApi.md#orgs_financial_line_items) | **GET** /api/v1/orgs/{id}/financial_line_items | Reported line items |
| [**orgs_financial_overview**](OrgsApi.md#orgs_financial_overview) | **GET** /api/v1/orgs/{id}/financial_overview | Financial overview |
| [**orgs_financials**](OrgsApi.md#orgs_financials) | **GET** /api/v1/orgs/{id}/financials | SEC-reported annual financials |
| [**orgs_funding**](OrgsApi.md#orgs_funding) | **GET** /api/v1/orgs/{id}/funding | Financing and grant events |
| [**orgs_get**](OrgsApi.md#orgs_get) | **GET** /api/v1/orgs/{id} | One orgs record |
| [**orgs_group_sections**](OrgsApi.md#orgs_group_sections) | **GET** /api/v1/orgs/{id}/group_sections | Group sections |
| [**orgs_group_structure**](OrgsApi.md#orgs_group_structure) | **GET** /api/v1/orgs/{id}/group_structure | Group structure |
| [**orgs_gxp_sites**](OrgsApi.md#orgs_gxp_sites) | **GET** /api/v1/orgs/{id}/gxp_sites | Manufacturing and distribution sites |
| [**orgs_gxp_sites_map**](OrgsApi.md#orgs_gxp_sites_map) | **GET** /api/v1/orgs/{id}/gxp_sites_map | Site map pins |
| [**orgs_identity**](OrgsApi.md#orgs_identity) | **GET** /api/v1/orgs/{id}/identity | Organization identity |
| [**orgs_import_refusals**](OrgsApi.md#orgs_import_refusals) | **GET** /api/v1/orgs/{id}/import_refusals | FDA import refusals |
| [**orgs_inspections**](OrgsApi.md#orgs_inspections) | **GET** /api/v1/orgs/{id}/inspections | FDA inspection record |
| [**orgs_institutional13f**](OrgsApi.md#orgs_institutional13f) | **GET** /api/v1/orgs/{id}/institutional_13f | Institutional holders (SEC Form 13F) |
| [**orgs_journals**](OrgsApi.md#orgs_journals) | **GET** /api/v1/orgs/{id}/journals | Journals published in |
| [**orgs_list**](OrgsApi.md#orgs_list) | **GET** /api/v1/orgs | Organizations |
| [**orgs_member_children**](OrgsApi.md#orgs_member_children) | **GET** /api/v1/orgs/{id}/member_children | Direct subsidiaries |
| [**orgs_member_identity**](OrgsApi.md#orgs_member_identity) | **GET** /api/v1/orgs/{id}/member_identity | Identity |
| [**orgs_member_records**](OrgsApi.md#orgs_member_records) | **GET** /api/v1/orgs/{id}/member_records | Records for this entity |
| [**orgs_members**](OrgsApi.md#orgs_members) | **GET** /api/v1/orgs/{id}/members | Companies grouped under this page |
| [**orgs_org_assets**](OrgsApi.md#orgs_org_assets) | **GET** /api/v1/orgs/{id}/org_assets | Approved and marketed assets |
| [**orgs_org_customer**](OrgsApi.md#orgs_org_customer) | **GET** /api/v1/orgs/{id}/org_customer | Scientist.com customer |
| [**orgs_org_gleif**](OrgsApi.md#orgs_org_gleif) | **GET** /api/v1/orgs/{id}/org_gleif | GLEIF corporate group |
| [**orgs_org_hosted**](OrgsApi.md#orgs_org_hosted) | **GET** /api/v1/orgs/{id}/org_hosted | Also a trial site |
| [**orgs_org_owned_drugs**](OrgsApi.md#orgs_org_owned_drugs) | **GET** /api/v1/orgs/{id}/org_owned_drugs | Drugs whose patents this organization owns |
| [**orgs_org_ownership**](OrgsApi.md#orgs_org_ownership) | **GET** /api/v1/orgs/{id}/org_ownership | Ownership and control (capped panel) |
| [**orgs_org_patent_mentions**](OrgsApi.md#orgs_org_patent_mentions) | **GET** /api/v1/orgs/{id}/org_patent_mentions | Patent mentions |
| [**orgs_org_people**](OrgsApi.md#orgs_org_people) | **GET** /api/v1/orgs/{id}/org_people | People |
| [**orgs_org_pipeline**](OrgsApi.md#orgs_org_pipeline) | **GET** /api/v1/orgs/{id}/org_pipeline | Pipeline: development lifecycle |
| [**orgs_org_programs**](OrgsApi.md#orgs_org_programs) | **GET** /api/v1/orgs/{id}/org_programs | Pipeline programs |
| [**orgs_org_publications**](OrgsApi.md#orgs_org_publications) | **GET** /api/v1/orgs/{id}/org_publications | Publications |
| [**orgs_org_recruitment**](OrgsApi.md#orgs_org_recruitment) | **GET** /api/v1/orgs/{id}/org_recruitment | Recruitment duration |
| [**orgs_org_relationships**](OrgsApi.md#orgs_org_relationships) | **GET** /api/v1/orgs/{id}/org_relationships | Organization relationships |
| [**orgs_org_supplier**](OrgsApi.md#orgs_org_supplier) | **GET** /api/v1/orgs/{id}/org_supplier | Scientist.com supplier |
| [**orgs_org_targets**](OrgsApi.md#orgs_org_targets) | **GET** /api/v1/orgs/{id}/org_targets | Targets being pursued |
| [**orgs_org_type_tally**](OrgsApi.md#orgs_org_type_tally) | **GET** /api/v1/orgs/org_type_tally | Organization pages by type |
| [**orgs_org_us_patents**](OrgsApi.md#orgs_org_us_patents) | **GET** /api/v1/orgs/{id}/org_us_patents | US granted patents |
| [**orgs_overview_digest**](OrgsApi.md#orgs_overview_digest) | **GET** /api/v1/orgs/{id}/overview_digest | Across the record |
| [**orgs_ownership_coverage**](OrgsApi.md#orgs_ownership_coverage) | **GET** /api/v1/orgs/{id}/ownership_coverage | Source coverage |
| [**orgs_ownership_history**](OrgsApi.md#orgs_ownership_history) | **GET** /api/v1/orgs/{id}/ownership_history | Reported history |
| [**orgs_ownership_percentages**](OrgsApi.md#orgs_ownership_percentages) | **GET** /api/v1/orgs/{id}/ownership_percentages | Share of shares outstanding |
| [**orgs_patent_cliff**](OrgsApi.md#orgs_patent_cliff) | **GET** /api/v1/orgs/{id}/patent_cliff | Patent-cliff exposure |
| [**orgs_patent_vs_trial**](OrgsApi.md#orgs_patent_vs_trial) | **GET** /api/v1/orgs/{id}/patent_vs_trial | Patents vs trials |
| [**orgs_patents**](OrgsApi.md#orgs_patents) | **GET** /api/v1/orgs/{id}/patents | US granted patents |
| [**orgs_pdufa_reconciliation**](OrgsApi.md#orgs_pdufa_reconciliation) | **GET** /api/v1/orgs/{id}/pdufa_reconciliation | SEC vs Benzinga reconciliation |
| [**orgs_peers**](OrgsApi.md#orgs_peers) | **GET** /api/v1/orgs/{id}/peers | Peer comparison |
| [**orgs_phase_status**](OrgsApi.md#orgs_phase_status) | **GET** /api/v1/orgs/{id}/phase_status | Phase status |
| [**orgs_publishing_biomarkers**](OrgsApi.md#orgs_publishing_biomarkers) | **GET** /api/v1/orgs/{id}/publishing_biomarkers | Biomarkers this organization publishes on |
| [**orgs_rare_activity**](OrgsApi.md#orgs_rare_activity) | **GET** /api/v1/orgs/{id}/rare_activity | Rare disease activity |
| [**orgs_recalls**](OrgsApi.md#orgs_recalls) | **GET** /api/v1/orgs/{id}/recalls | FDA recalls |
| [**orgs_registry_shell**](OrgsApi.md#orgs_registry_shell) | **GET** /api/v1/orgs/{id}/registry_shell | Operating-firm registry record |
| [**orgs_related_biomarkers**](OrgsApi.md#orgs_related_biomarkers) | **GET** /api/v1/orgs/{id}/related_biomarkers | Related biomarkers |
| [**orgs_related_conditions**](OrgsApi.md#orgs_related_conditions) | **GET** /api/v1/orgs/{id}/related_conditions | Related conditions |
| [**orgs_related_unfolded**](OrgsApi.md#orgs_related_unfolded) | **GET** /api/v1/orgs/{id}/related_unfolded | Related companies not grouped here |
| [**orgs_revenue_by_drug**](OrgsApi.md#orgs_revenue_by_drug) | **GET** /api/v1/orgs/{id}/revenue_by_drug | Revenue by drug |
| [**orgs_revenue_geography**](OrgsApi.md#orgs_revenue_geography) | **GET** /api/v1/orgs/{id}/revenue_geography | Revenue by geography |
| [**orgs_same_name_sec_filer**](OrgsApi.md#orgs_same_name_sec_filer) | **GET** /api/v1/orgs/{id}/same_name_sec_filer | Same-name SEC filer |
| [**orgs_sbti_crosscheck**](OrgsApi.md#orgs_sbti_crosscheck) | **GET** /api/v1/orgs/{id}/sbti_crosscheck | SBTi cross-check |
| [**orgs_scientist_com**](OrgsApi.md#orgs_scientist_com) | **GET** /api/v1/orgs/{id}/scientist_com | On Scientist.com |
| [**orgs_sec_filer_shell**](OrgsApi.md#orgs_sec_filer_shell) | **GET** /api/v1/orgs/{id}/sec_filer_shell | SEC filer record |
| [**orgs_sec_pdufa**](OrgsApi.md#orgs_sec_pdufa) | **GET** /api/v1/orgs/{id}/sec_pdufa | SEC PDUFA decisions |
| [**orgs_sponsor_delivery**](OrgsApi.md#orgs_sponsor_delivery) | **GET** /api/v1/orgs/{id}/sponsor_delivery | Sponsor delivery record |
| [**orgs_supplier_links**](OrgsApi.md#orgs_supplier_links) | **GET** /api/v1/orgs/supplier_links | Suppliers matching your search |
| [**orgs_timeline**](OrgsApi.md#orgs_timeline) | **GET** /api/v1/orgs/{id}/timeline | Latest dated events |
| [**orgs_top_conditions**](OrgsApi.md#orgs_top_conditions) | **GET** /api/v1/orgs/{id}/top_conditions | Top disease areas |
| [**orgs_top_drugs**](OrgsApi.md#orgs_top_drugs) | **GET** /api/v1/orgs/{id}/top_drugs | Top drugs (all indications) |
| [**orgs_trials_per_year**](OrgsApi.md#orgs_trials_per_year) | **GET** /api/v1/orgs/{id}/trials_per_year | Trials by start year as sponsor or collaborator |
| [**orgs_unmatched_revenue**](OrgsApi.md#orgs_unmatched_revenue) | **GET** /api/v1/orgs/{id}/unmatched_revenue | Product lines not matched to a brand |
| [**orgs_untitled_letters**](OrgsApi.md#orgs_untitled_letters) | **GET** /api/v1/orgs/{id}/untitled_letters | FDA untitled letters |
| [**orgs_warning_letters**](OrgsApi.md#orgs_warning_letters) | **GET** /api/v1/orgs/{id}/warning_letters | FDA warning letters |


## orgs_active_programs

> <OrgsActivePrograms200Response> orgs_active_programs(id)

Active development programs

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Active development programs
  result = api_instance.orgs_active_programs(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_active_programs: #{e}"
end
```

#### Using the orgs_active_programs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsActivePrograms200Response>, Integer, Hash)> orgs_active_programs_with_http_info(id)

```ruby
begin
  # Active development programs
  data, status_code, headers = api_instance.orgs_active_programs_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsActivePrograms200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_active_programs_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsActivePrograms200Response**](OrgsActivePrograms200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_activity_summary

> <OrgsActivitySummary200Response> orgs_activity_summary(id)

Activity summary

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Activity summary
  result = api_instance.orgs_activity_summary(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_activity_summary: #{e}"
end
```

#### Using the orgs_activity_summary_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsActivitySummary200Response>, Integer, Hash)> orgs_activity_summary_with_http_info(id)

```ruby
begin
  # Activity summary
  data, status_code, headers = api_instance.orgs_activity_summary_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsActivitySummary200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_activity_summary_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsActivitySummary200Response**](OrgsActivitySummary200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_approved_assets

> <OrgsApprovedAssets200Response> orgs_approved_assets(id)

Approved assets

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Approved assets
  result = api_instance.orgs_approved_assets(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_approved_assets: #{e}"
end
```

#### Using the orgs_approved_assets_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsApprovedAssets200Response>, Integer, Hash)> orgs_approved_assets_with_http_info(id)

```ruby
begin
  # Approved assets
  data, status_code, headers = api_instance.orgs_approved_assets_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsApprovedAssets200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_approved_assets_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsApprovedAssets200Response**](OrgsApprovedAssets200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_benzinga_catalysts

> <OrgsBenzingaCatalysts200Response> orgs_benzinga_catalysts(id)

Press-release catalysts

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Press-release catalysts
  result = api_instance.orgs_benzinga_catalysts(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_benzinga_catalysts: #{e}"
end
```

#### Using the orgs_benzinga_catalysts_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsBenzingaCatalysts200Response>, Integer, Hash)> orgs_benzinga_catalysts_with_http_info(id)

```ruby
begin
  # Press-release catalysts
  data, status_code, headers = api_instance.orgs_benzinga_catalysts_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsBenzingaCatalysts200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_benzinga_catalysts_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsBenzingaCatalysts200Response**](OrgsBenzingaCatalysts200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_biomarker_capability

> <OrgsBiomarkerCapability200Response> orgs_biomarker_capability(id)

Biomarker capability evidence

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Biomarker capability evidence
  result = api_instance.orgs_biomarker_capability(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_biomarker_capability: #{e}"
end
```

#### Using the orgs_biomarker_capability_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsBiomarkerCapability200Response>, Integer, Hash)> orgs_biomarker_capability_with_http_info(id)

```ruby
begin
  # Biomarker capability evidence
  data, status_code, headers = api_instance.orgs_biomarker_capability_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsBiomarkerCapability200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_biomarker_capability_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsBiomarkerCapability200Response**](OrgsBiomarkerCapability200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_clia

> <OrgsClia200Response> orgs_clia(id)

CLIA clinical laboratory certificates

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # CLIA clinical laboratory certificates
  result = api_instance.orgs_clia(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_clia: #{e}"
end
```

#### Using the orgs_clia_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsClia200Response>, Integer, Hash)> orgs_clia_with_http_info(id)

```ruby
begin
  # CLIA clinical laboratory certificates
  data, status_code, headers = api_instance.orgs_clia_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsClia200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_clia_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsClia200Response**](OrgsClia200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_climate

> <OrgsClimate200Response> orgs_climate(id)

Climate targets, as filed

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Climate targets, as filed
  result = api_instance.orgs_climate(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_climate: #{e}"
end
```

#### Using the orgs_climate_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsClimate200Response>, Integer, Hash)> orgs_climate_with_http_info(id)

```ruby
begin
  # Climate targets, as filed
  data, status_code, headers = api_instance.orgs_climate_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsClimate200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_climate_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsClimate200Response**](OrgsClimate200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_collaborators

> <OrgsCollaborators200Response> orgs_collaborators(id)

Study collaborators

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'sponsor:bae3f16c4a392772a747f9e7b149e9ad' # String | 

begin
  # Study collaborators
  result = api_instance.orgs_collaborators(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_collaborators: #{e}"
end
```

#### Using the orgs_collaborators_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsCollaborators200Response>, Integer, Hash)> orgs_collaborators_with_http_info(id)

```ruby
begin
  # Study collaborators
  data, status_code, headers = api_instance.orgs_collaborators_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsCollaborators200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_collaborators_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsCollaborators200Response**](OrgsCollaborators200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_company

> <OrgsCompany200Response> orgs_company(id)

Company

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Company
  result = api_instance.orgs_company(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_company: #{e}"
end
```

#### Using the orgs_company_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsCompany200Response>, Integer, Hash)> orgs_company_with_http_info(id)

```ruby
begin
  # Company
  data, status_code, headers = api_instance.orgs_company_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsCompany200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_company_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsCompany200Response**](OrgsCompany200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_competitor_activity

> <OrgsCollaborators200Response> orgs_competitor_activity(id)

Competitor activity

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'sponsor:29e5c2bafe3cbc660c0ceaac902442fc' # String | 

begin
  # Competitor activity
  result = api_instance.orgs_competitor_activity(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_competitor_activity: #{e}"
end
```

#### Using the orgs_competitor_activity_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsCollaborators200Response>, Integer, Hash)> orgs_competitor_activity_with_http_info(id)

```ruby
begin
  # Competitor activity
  data, status_code, headers = api_instance.orgs_competitor_activity_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsCollaborators200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_competitor_activity_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsCollaborators200Response**](OrgsCollaborators200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_competitors

> <OrgsCompetitors200Response> orgs_competitors(id)

Competitors

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'sponsor:b80857179c0f39712baa71f8183c3919' # String | 

begin
  # Competitors
  result = api_instance.orgs_competitors(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_competitors: #{e}"
end
```

#### Using the orgs_competitors_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsCompetitors200Response>, Integer, Hash)> orgs_competitors_with_http_info(id)

```ruby
begin
  # Competitors
  data, status_code, headers = api_instance.orgs_competitors_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsCompetitors200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_competitors_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsCompetitors200Response**](OrgsCompetitors200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_coverage

> <OrgsCoverage200Response> orgs_coverage(id)

Coverage of this page

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Coverage of this page
  result = api_instance.orgs_coverage(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_coverage: #{e}"
end
```

#### Using the orgs_coverage_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsCoverage200Response>, Integer, Hash)> orgs_coverage_with_http_info(id)

```ruby
begin
  # Coverage of this page
  data, status_code, headers = api_instance.orgs_coverage_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsCoverage200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_coverage_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsCoverage200Response**](OrgsCoverage200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_current_ownership

> <OrgsCurrentOwnership200Response> orgs_current_ownership(id)

Current ownership and control

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Current ownership and control
  result = api_instance.orgs_current_ownership(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_current_ownership: #{e}"
end
```

#### Using the orgs_current_ownership_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsCurrentOwnership200Response>, Integer, Hash)> orgs_current_ownership_with_http_info(id)

```ruby
begin
  # Current ownership and control
  data, status_code, headers = api_instance.orgs_current_ownership_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsCurrentOwnership200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_current_ownership_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsCurrentOwnership200Response**](OrgsCurrentOwnership200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_declarations

> <OrgsDeclarations200Response> orgs_declarations(id)

Organization-wide declarations

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'charity-commission:9000002' # String | 

begin
  # Organization-wide declarations
  result = api_instance.orgs_declarations(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_declarations: #{e}"
end
```

#### Using the orgs_declarations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsDeclarations200Response>, Integer, Hash)> orgs_declarations_with_http_info(id)

```ruby
begin
  # Organization-wide declarations
  data, status_code, headers = api_instance.orgs_declarations_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsDeclarations200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_declarations_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsDeclarations200Response**](OrgsDeclarations200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_delivery

> <OrgsDelivery200Response> orgs_delivery(id)

Delivery reliability

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Delivery reliability
  result = api_instance.orgs_delivery(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_delivery: #{e}"
end
```

#### Using the orgs_delivery_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsDelivery200Response>, Integer, Hash)> orgs_delivery_with_http_info(id)

```ruby
begin
  # Delivery reliability
  data, status_code, headers = api_instance.orgs_delivery_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsDelivery200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_delivery_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsDelivery200Response**](OrgsDelivery200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_description

> <OrgsDescription200Response> orgs_description(id)

Description

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Description
  result = api_instance.orgs_description(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_description: #{e}"
end
```

#### Using the orgs_description_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsDescription200Response>, Integer, Hash)> orgs_description_with_http_info(id)

```ruby
begin
  # Description
  data, status_code, headers = api_instance.orgs_description_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsDescription200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_description_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsDescription200Response**](OrgsDescription200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_disease_areas

> <OrgsDiseaseAreas200Response> orgs_disease_areas(id)

Disease areas

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Disease areas
  result = api_instance.orgs_disease_areas(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_disease_areas: #{e}"
end
```

#### Using the orgs_disease_areas_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsDiseaseAreas200Response>, Integer, Hash)> orgs_disease_areas_with_http_info(id)

```ruby
begin
  # Disease areas
  data, status_code, headers = api_instance.orgs_disease_areas_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsDiseaseAreas200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_disease_areas_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsDiseaseAreas200Response**](OrgsDiseaseAreas200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_diversity

> <OrgsDiversity200Response> orgs_diversity(id)

Workforce and board diversity, as filed

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Workforce and board diversity, as filed
  result = api_instance.orgs_diversity(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_diversity: #{e}"
end
```

#### Using the orgs_diversity_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsDiversity200Response>, Integer, Hash)> orgs_diversity_with_http_info(id)

```ruby
begin
  # Workforce and board diversity, as filed
  data, status_code, headers = api_instance.orgs_diversity_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsDiversity200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_diversity_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsDiversity200Response**](OrgsDiversity200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_esg_trend

> <OrgsEsgTrend200Response> orgs_esg_trend(id)

Climate and diversity by reporting year

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Climate and diversity by reporting year
  result = api_instance.orgs_esg_trend(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_esg_trend: #{e}"
end
```

#### Using the orgs_esg_trend_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsEsgTrend200Response>, Integer, Hash)> orgs_esg_trend_with_http_info(id)

```ruby
begin
  # Climate and diversity by reporting year
  data, status_code, headers = api_instance.orgs_esg_trend_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsEsgTrend200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_esg_trend_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsEsgTrend200Response**](OrgsEsgTrend200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_financial_line_items

> <OrgsFinancialLineItems200Response> orgs_financial_line_items(id)

Reported line items

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Reported line items
  result = api_instance.orgs_financial_line_items(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_financial_line_items: #{e}"
end
```

#### Using the orgs_financial_line_items_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsFinancialLineItems200Response>, Integer, Hash)> orgs_financial_line_items_with_http_info(id)

```ruby
begin
  # Reported line items
  data, status_code, headers = api_instance.orgs_financial_line_items_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsFinancialLineItems200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_financial_line_items_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsFinancialLineItems200Response**](OrgsFinancialLineItems200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_financial_overview

> <OrgsFinancialOverview200Response> orgs_financial_overview(id)

Financial overview

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Financial overview
  result = api_instance.orgs_financial_overview(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_financial_overview: #{e}"
end
```

#### Using the orgs_financial_overview_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsFinancialOverview200Response>, Integer, Hash)> orgs_financial_overview_with_http_info(id)

```ruby
begin
  # Financial overview
  data, status_code, headers = api_instance.orgs_financial_overview_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsFinancialOverview200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_financial_overview_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsFinancialOverview200Response**](OrgsFinancialOverview200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_financials

> <OrgsFinancials200Response> orgs_financials(id)

SEC-reported annual financials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # SEC-reported annual financials
  result = api_instance.orgs_financials(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_financials: #{e}"
end
```

#### Using the orgs_financials_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsFinancials200Response>, Integer, Hash)> orgs_financials_with_http_info(id)

```ruby
begin
  # SEC-reported annual financials
  data, status_code, headers = api_instance.orgs_financials_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsFinancials200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_financials_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsFinancials200Response**](OrgsFinancials200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_funding

> <OrgsFunding200Response> orgs_funding(id)

Financing and grant events

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Financing and grant events
  result = api_instance.orgs_funding(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_funding: #{e}"
end
```

#### Using the orgs_funding_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsFunding200Response>, Integer, Hash)> orgs_funding_with_http_info(id)

```ruby
begin
  # Financing and grant events
  data, status_code, headers = api_instance.orgs_funding_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsFunding200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_funding_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsFunding200Response**](OrgsFunding200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_get

> <OrgsGet200Response> orgs_get(id)

One orgs record

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # One orgs record
  result = api_instance.orgs_get(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_get: #{e}"
end
```

#### Using the orgs_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsGet200Response>, Integer, Hash)> orgs_get_with_http_info(id)

```ruby
begin
  # One orgs record
  data, status_code, headers = api_instance.orgs_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsGet200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsGet200Response**](OrgsGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_group_sections

> <OrgsGroupSections200Response> orgs_group_sections(id)

Group sections

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Group sections
  result = api_instance.orgs_group_sections(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_group_sections: #{e}"
end
```

#### Using the orgs_group_sections_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsGroupSections200Response>, Integer, Hash)> orgs_group_sections_with_http_info(id)

```ruby
begin
  # Group sections
  data, status_code, headers = api_instance.orgs_group_sections_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsGroupSections200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_group_sections_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsGroupSections200Response**](OrgsGroupSections200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_group_structure

> <OrgsGroupStructure200Response> orgs_group_structure(id)

Group structure

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Group structure
  result = api_instance.orgs_group_structure(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_group_structure: #{e}"
end
```

#### Using the orgs_group_structure_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsGroupStructure200Response>, Integer, Hash)> orgs_group_structure_with_http_info(id)

```ruby
begin
  # Group structure
  data, status_code, headers = api_instance.orgs_group_structure_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsGroupStructure200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_group_structure_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsGroupStructure200Response**](OrgsGroupStructure200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_gxp_sites

> <OrgsGxpSites200Response> orgs_gxp_sites(id)

Manufacturing and distribution sites

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Manufacturing and distribution sites
  result = api_instance.orgs_gxp_sites(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_gxp_sites: #{e}"
end
```

#### Using the orgs_gxp_sites_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsGxpSites200Response>, Integer, Hash)> orgs_gxp_sites_with_http_info(id)

```ruby
begin
  # Manufacturing and distribution sites
  data, status_code, headers = api_instance.orgs_gxp_sites_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsGxpSites200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_gxp_sites_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsGxpSites200Response**](OrgsGxpSites200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_gxp_sites_map

> <OrgsGxpSitesMap200Response> orgs_gxp_sites_map(id)

Site map pins

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Site map pins
  result = api_instance.orgs_gxp_sites_map(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_gxp_sites_map: #{e}"
end
```

#### Using the orgs_gxp_sites_map_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsGxpSitesMap200Response>, Integer, Hash)> orgs_gxp_sites_map_with_http_info(id)

```ruby
begin
  # Site map pins
  data, status_code, headers = api_instance.orgs_gxp_sites_map_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsGxpSitesMap200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_gxp_sites_map_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsGxpSitesMap200Response**](OrgsGxpSitesMap200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_identity

> <OrgsIdentity200Response> orgs_identity(id)

Organization identity

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Organization identity
  result = api_instance.orgs_identity(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_identity: #{e}"
end
```

#### Using the orgs_identity_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsIdentity200Response>, Integer, Hash)> orgs_identity_with_http_info(id)

```ruby
begin
  # Organization identity
  data, status_code, headers = api_instance.orgs_identity_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsIdentity200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_identity_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsIdentity200Response**](OrgsIdentity200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_import_refusals

> <OrgsImportRefusals200Response> orgs_import_refusals(id)

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

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # FDA import refusals
  result = api_instance.orgs_import_refusals(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_import_refusals: #{e}"
end
```

#### Using the orgs_import_refusals_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsImportRefusals200Response>, Integer, Hash)> orgs_import_refusals_with_http_info(id)

```ruby
begin
  # FDA import refusals
  data, status_code, headers = api_instance.orgs_import_refusals_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsImportRefusals200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_import_refusals_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsImportRefusals200Response**](OrgsImportRefusals200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_inspections

> <OrgsImportRefusals200Response> orgs_inspections(id)

FDA inspection record

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # FDA inspection record
  result = api_instance.orgs_inspections(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_inspections: #{e}"
end
```

#### Using the orgs_inspections_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsImportRefusals200Response>, Integer, Hash)> orgs_inspections_with_http_info(id)

```ruby
begin
  # FDA inspection record
  data, status_code, headers = api_instance.orgs_inspections_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsImportRefusals200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_inspections_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsImportRefusals200Response**](OrgsImportRefusals200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_institutional13f

> <OrgsInstitutional13f200Response> orgs_institutional13f(id)

Institutional holders (SEC Form 13F)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Institutional holders (SEC Form 13F)
  result = api_instance.orgs_institutional13f(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_institutional13f: #{e}"
end
```

#### Using the orgs_institutional13f_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsInstitutional13f200Response>, Integer, Hash)> orgs_institutional13f_with_http_info(id)

```ruby
begin
  # Institutional holders (SEC Form 13F)
  data, status_code, headers = api_instance.orgs_institutional13f_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsInstitutional13f200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_institutional13f_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsInstitutional13f200Response**](OrgsInstitutional13f200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_journals

> <OrgsJournals200Response> orgs_journals(id)

Journals published in

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Journals published in
  result = api_instance.orgs_journals(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_journals: #{e}"
end
```

#### Using the orgs_journals_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsJournals200Response>, Integer, Hash)> orgs_journals_with_http_info(id)

```ruby
begin
  # Journals published in
  data, status_code, headers = api_instance.orgs_journals_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsJournals200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_journals_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsJournals200Response**](OrgsJournals200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_list

> <OrgsList200Response> orgs_list

Organizations

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new

begin
  # Organizations
  result = api_instance.orgs_list
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_list: #{e}"
end
```

#### Using the orgs_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsList200Response>, Integer, Hash)> orgs_list_with_http_info

```ruby
begin
  # Organizations
  data, status_code, headers = api_instance.orgs_list_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsList200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_list_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**OrgsList200Response**](OrgsList200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_member_children

> <OrgsMemberChildren200Response> orgs_member_children(id)

Direct subsidiaries

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'company:0009c1abcb13474d358773a3b56ab4d4' # String | 

begin
  # Direct subsidiaries
  result = api_instance.orgs_member_children(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_member_children: #{e}"
end
```

#### Using the orgs_member_children_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsMemberChildren200Response>, Integer, Hash)> orgs_member_children_with_http_info(id)

```ruby
begin
  # Direct subsidiaries
  data, status_code, headers = api_instance.orgs_member_children_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsMemberChildren200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_member_children_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsMemberChildren200Response**](OrgsMemberChildren200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_member_identity

> <OrgsMemberIdentity200Response> orgs_member_identity(id)

Identity

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'company:0009c1abcb13474d358773a3b56ab4d4' # String | 

begin
  # Identity
  result = api_instance.orgs_member_identity(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_member_identity: #{e}"
end
```

#### Using the orgs_member_identity_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsMemberIdentity200Response>, Integer, Hash)> orgs_member_identity_with_http_info(id)

```ruby
begin
  # Identity
  data, status_code, headers = api_instance.orgs_member_identity_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsMemberIdentity200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_member_identity_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsMemberIdentity200Response**](OrgsMemberIdentity200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_member_records

> <OrgsMemberRecords200Response> orgs_member_records(id)

Records for this entity

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'company:0009c1abcb13474d358773a3b56ab4d4' # String | 

begin
  # Records for this entity
  result = api_instance.orgs_member_records(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_member_records: #{e}"
end
```

#### Using the orgs_member_records_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsMemberRecords200Response>, Integer, Hash)> orgs_member_records_with_http_info(id)

```ruby
begin
  # Records for this entity
  data, status_code, headers = api_instance.orgs_member_records_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsMemberRecords200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_member_records_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsMemberRecords200Response**](OrgsMemberRecords200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_members

> <OrgsMembers200Response> orgs_members(id)

Companies grouped under this page

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Companies grouped under this page
  result = api_instance.orgs_members(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_members: #{e}"
end
```

#### Using the orgs_members_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsMembers200Response>, Integer, Hash)> orgs_members_with_http_info(id)

```ruby
begin
  # Companies grouped under this page
  data, status_code, headers = api_instance.orgs_members_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsMembers200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_members_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsMembers200Response**](OrgsMembers200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_assets

> <OrgsOrgAssets200Response> orgs_org_assets(id)

Approved and marketed assets

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Approved and marketed assets
  result = api_instance.orgs_org_assets(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_assets: #{e}"
end
```

#### Using the orgs_org_assets_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgAssets200Response>, Integer, Hash)> orgs_org_assets_with_http_info(id)

```ruby
begin
  # Approved and marketed assets
  data, status_code, headers = api_instance.orgs_org_assets_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgAssets200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_assets_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgAssets200Response**](OrgsOrgAssets200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_customer

> <OrgsOrgCustomer200Response> orgs_org_customer(id)

Scientist.com customer

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Scientist.com customer
  result = api_instance.orgs_org_customer(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_customer: #{e}"
end
```

#### Using the orgs_org_customer_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgCustomer200Response>, Integer, Hash)> orgs_org_customer_with_http_info(id)

```ruby
begin
  # Scientist.com customer
  data, status_code, headers = api_instance.orgs_org_customer_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgCustomer200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_customer_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgCustomer200Response**](OrgsOrgCustomer200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_gleif

> <OrgsOrgGleif200Response> orgs_org_gleif(id)

GLEIF corporate group

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # GLEIF corporate group
  result = api_instance.orgs_org_gleif(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_gleif: #{e}"
end
```

#### Using the orgs_org_gleif_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgGleif200Response>, Integer, Hash)> orgs_org_gleif_with_http_info(id)

```ruby
begin
  # GLEIF corporate group
  data, status_code, headers = api_instance.orgs_org_gleif_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgGleif200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_gleif_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgGleif200Response**](OrgsOrgGleif200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_hosted

> <OrgsOrgHosted200Response> orgs_org_hosted(id)

Also a trial site

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'fdafirm:3010000003' # String | 

begin
  # Also a trial site
  result = api_instance.orgs_org_hosted(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_hosted: #{e}"
end
```

#### Using the orgs_org_hosted_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgHosted200Response>, Integer, Hash)> orgs_org_hosted_with_http_info(id)

```ruby
begin
  # Also a trial site
  data, status_code, headers = api_instance.orgs_org_hosted_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgHosted200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_hosted_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgHosted200Response**](OrgsOrgHosted200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_owned_drugs

> <OrgsOrgOwnedDrugs200Response> orgs_org_owned_drugs(id)

Drugs whose patents this organization owns

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'ror:04gndp242' # String | 

begin
  # Drugs whose patents this organization owns
  result = api_instance.orgs_org_owned_drugs(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_owned_drugs: #{e}"
end
```

#### Using the orgs_org_owned_drugs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgOwnedDrugs200Response>, Integer, Hash)> orgs_org_owned_drugs_with_http_info(id)

```ruby
begin
  # Drugs whose patents this organization owns
  data, status_code, headers = api_instance.orgs_org_owned_drugs_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgOwnedDrugs200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_owned_drugs_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgOwnedDrugs200Response**](OrgsOrgOwnedDrugs200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_ownership

> <OrgsOrgOwnership200Response> orgs_org_ownership(id)

Ownership and control (capped panel)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Ownership and control (capped panel)
  result = api_instance.orgs_org_ownership(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_ownership: #{e}"
end
```

#### Using the orgs_org_ownership_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgOwnership200Response>, Integer, Hash)> orgs_org_ownership_with_http_info(id)

```ruby
begin
  # Ownership and control (capped panel)
  data, status_code, headers = api_instance.orgs_org_ownership_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgOwnership200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_ownership_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgOwnership200Response**](OrgsOrgOwnership200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_patent_mentions

> <OrgsOrgPatentMentions200Response> orgs_org_patent_mentions(id)

Patent mentions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Patent mentions
  result = api_instance.orgs_org_patent_mentions(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_patent_mentions: #{e}"
end
```

#### Using the orgs_org_patent_mentions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgPatentMentions200Response>, Integer, Hash)> orgs_org_patent_mentions_with_http_info(id)

```ruby
begin
  # Patent mentions
  data, status_code, headers = api_instance.orgs_org_patent_mentions_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgPatentMentions200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_patent_mentions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgPatentMentions200Response**](OrgsOrgPatentMentions200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_people

> <OrgsOrgPeople200Response> orgs_org_people(id)

People

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # People
  result = api_instance.orgs_org_people(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_people: #{e}"
end
```

#### Using the orgs_org_people_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgPeople200Response>, Integer, Hash)> orgs_org_people_with_http_info(id)

```ruby
begin
  # People
  data, status_code, headers = api_instance.orgs_org_people_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgPeople200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_people_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgPeople200Response**](OrgsOrgPeople200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_pipeline

> <OrgsOrgPipeline200Response> orgs_org_pipeline(id)

Pipeline: development lifecycle

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Pipeline: development lifecycle
  result = api_instance.orgs_org_pipeline(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_pipeline: #{e}"
end
```

#### Using the orgs_org_pipeline_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgPipeline200Response>, Integer, Hash)> orgs_org_pipeline_with_http_info(id)

```ruby
begin
  # Pipeline: development lifecycle
  data, status_code, headers = api_instance.orgs_org_pipeline_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgPipeline200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_pipeline_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgPipeline200Response**](OrgsOrgPipeline200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_programs

> <OrgsOrgPrograms200Response> orgs_org_programs(id)

Pipeline programs

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Pipeline programs
  result = api_instance.orgs_org_programs(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_programs: #{e}"
end
```

#### Using the orgs_org_programs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgPrograms200Response>, Integer, Hash)> orgs_org_programs_with_http_info(id)

```ruby
begin
  # Pipeline programs
  data, status_code, headers = api_instance.orgs_org_programs_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgPrograms200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_programs_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgPrograms200Response**](OrgsOrgPrograms200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_publications

> <OrgsOrgPublications200Response> orgs_org_publications(id)

Publications

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Publications
  result = api_instance.orgs_org_publications(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_publications: #{e}"
end
```

#### Using the orgs_org_publications_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgPublications200Response>, Integer, Hash)> orgs_org_publications_with_http_info(id)

```ruby
begin
  # Publications
  data, status_code, headers = api_instance.orgs_org_publications_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgPublications200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_publications_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgPublications200Response**](OrgsOrgPublications200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_recruitment

> <OrgsOrgRecruitment200Response> orgs_org_recruitment(id)

Recruitment duration

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Recruitment duration
  result = api_instance.orgs_org_recruitment(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_recruitment: #{e}"
end
```

#### Using the orgs_org_recruitment_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgRecruitment200Response>, Integer, Hash)> orgs_org_recruitment_with_http_info(id)

```ruby
begin
  # Recruitment duration
  data, status_code, headers = api_instance.orgs_org_recruitment_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgRecruitment200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_recruitment_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgRecruitment200Response**](OrgsOrgRecruitment200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_relationships

> <OrgsOrgRelationships200Response> orgs_org_relationships(id)

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

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Organization relationships
  result = api_instance.orgs_org_relationships(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_relationships: #{e}"
end
```

#### Using the orgs_org_relationships_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgRelationships200Response>, Integer, Hash)> orgs_org_relationships_with_http_info(id)

```ruby
begin
  # Organization relationships
  data, status_code, headers = api_instance.orgs_org_relationships_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgRelationships200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_relationships_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgRelationships200Response**](OrgsOrgRelationships200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_supplier

> <OrgsOrgSupplier200Response> orgs_org_supplier(id)

Scientist.com supplier

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Scientist.com supplier
  result = api_instance.orgs_org_supplier(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_supplier: #{e}"
end
```

#### Using the orgs_org_supplier_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgSupplier200Response>, Integer, Hash)> orgs_org_supplier_with_http_info(id)

```ruby
begin
  # Scientist.com supplier
  data, status_code, headers = api_instance.orgs_org_supplier_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgSupplier200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_supplier_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgSupplier200Response**](OrgsOrgSupplier200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_targets

> <OrgsOrgTargets200Response> orgs_org_targets(id)

Targets being pursued

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Targets being pursued
  result = api_instance.orgs_org_targets(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_targets: #{e}"
end
```

#### Using the orgs_org_targets_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgTargets200Response>, Integer, Hash)> orgs_org_targets_with_http_info(id)

```ruby
begin
  # Targets being pursued
  data, status_code, headers = api_instance.orgs_org_targets_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgTargets200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_targets_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgTargets200Response**](OrgsOrgTargets200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_type_tally

> <OrgsOrgTypeTally200Response> orgs_org_type_tally

Organization pages by type

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new

begin
  # Organization pages by type
  result = api_instance.orgs_org_type_tally
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_type_tally: #{e}"
end
```

#### Using the orgs_org_type_tally_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgTypeTally200Response>, Integer, Hash)> orgs_org_type_tally_with_http_info

```ruby
begin
  # Organization pages by type
  data, status_code, headers = api_instance.orgs_org_type_tally_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgTypeTally200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_type_tally_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**OrgsOrgTypeTally200Response**](OrgsOrgTypeTally200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_org_us_patents

> <OrgsOrgUsPatents200Response> orgs_org_us_patents(id)

US granted patents

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # US granted patents
  result = api_instance.orgs_org_us_patents(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_us_patents: #{e}"
end
```

#### Using the orgs_org_us_patents_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOrgUsPatents200Response>, Integer, Hash)> orgs_org_us_patents_with_http_info(id)

```ruby
begin
  # US granted patents
  data, status_code, headers = api_instance.orgs_org_us_patents_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOrgUsPatents200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_org_us_patents_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOrgUsPatents200Response**](OrgsOrgUsPatents200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_overview_digest

> <OrgsOverviewDigest200Response> orgs_overview_digest(id)

Across the record

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Across the record
  result = api_instance.orgs_overview_digest(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_overview_digest: #{e}"
end
```

#### Using the orgs_overview_digest_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOverviewDigest200Response>, Integer, Hash)> orgs_overview_digest_with_http_info(id)

```ruby
begin
  # Across the record
  data, status_code, headers = api_instance.orgs_overview_digest_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOverviewDigest200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_overview_digest_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOverviewDigest200Response**](OrgsOverviewDigest200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_ownership_coverage

> <OrgsOwnershipCoverage200Response> orgs_ownership_coverage(id)

Source coverage

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Source coverage
  result = api_instance.orgs_ownership_coverage(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_ownership_coverage: #{e}"
end
```

#### Using the orgs_ownership_coverage_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOwnershipCoverage200Response>, Integer, Hash)> orgs_ownership_coverage_with_http_info(id)

```ruby
begin
  # Source coverage
  data, status_code, headers = api_instance.orgs_ownership_coverage_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOwnershipCoverage200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_ownership_coverage_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOwnershipCoverage200Response**](OrgsOwnershipCoverage200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_ownership_history

> <OrgsOwnershipHistory200Response> orgs_ownership_history(id)

Reported history

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Reported history
  result = api_instance.orgs_ownership_history(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_ownership_history: #{e}"
end
```

#### Using the orgs_ownership_history_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOwnershipHistory200Response>, Integer, Hash)> orgs_ownership_history_with_http_info(id)

```ruby
begin
  # Reported history
  data, status_code, headers = api_instance.orgs_ownership_history_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOwnershipHistory200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_ownership_history_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOwnershipHistory200Response**](OrgsOwnershipHistory200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_ownership_percentages

> <OrgsOwnershipPercentages200Response> orgs_ownership_percentages(id)

Share of shares outstanding

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Share of shares outstanding
  result = api_instance.orgs_ownership_percentages(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_ownership_percentages: #{e}"
end
```

#### Using the orgs_ownership_percentages_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsOwnershipPercentages200Response>, Integer, Hash)> orgs_ownership_percentages_with_http_info(id)

```ruby
begin
  # Share of shares outstanding
  data, status_code, headers = api_instance.orgs_ownership_percentages_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsOwnershipPercentages200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_ownership_percentages_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsOwnershipPercentages200Response**](OrgsOwnershipPercentages200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_patent_cliff

> <OrgsPatentCliff200Response> orgs_patent_cliff(id)

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

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Patent-cliff exposure
  result = api_instance.orgs_patent_cliff(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_patent_cliff: #{e}"
end
```

#### Using the orgs_patent_cliff_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsPatentCliff200Response>, Integer, Hash)> orgs_patent_cliff_with_http_info(id)

```ruby
begin
  # Patent-cliff exposure
  data, status_code, headers = api_instance.orgs_patent_cliff_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsPatentCliff200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_patent_cliff_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsPatentCliff200Response**](OrgsPatentCliff200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_patent_vs_trial

> <OrgsPatentVsTrial200Response> orgs_patent_vs_trial(id)

Patents vs trials

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Patents vs trials
  result = api_instance.orgs_patent_vs_trial(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_patent_vs_trial: #{e}"
end
```

#### Using the orgs_patent_vs_trial_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsPatentVsTrial200Response>, Integer, Hash)> orgs_patent_vs_trial_with_http_info(id)

```ruby
begin
  # Patents vs trials
  data, status_code, headers = api_instance.orgs_patent_vs_trial_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsPatentVsTrial200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_patent_vs_trial_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsPatentVsTrial200Response**](OrgsPatentVsTrial200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_patents

> <OrgsPatents200Response> orgs_patents(id)

US granted patents

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # US granted patents
  result = api_instance.orgs_patents(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_patents: #{e}"
end
```

#### Using the orgs_patents_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsPatents200Response>, Integer, Hash)> orgs_patents_with_http_info(id)

```ruby
begin
  # US granted patents
  data, status_code, headers = api_instance.orgs_patents_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsPatents200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_patents_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsPatents200Response**](OrgsPatents200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_pdufa_reconciliation

> <OrgsPdufaReconciliation200Response> orgs_pdufa_reconciliation(id)

SEC vs Benzinga reconciliation

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # SEC vs Benzinga reconciliation
  result = api_instance.orgs_pdufa_reconciliation(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_pdufa_reconciliation: #{e}"
end
```

#### Using the orgs_pdufa_reconciliation_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsPdufaReconciliation200Response>, Integer, Hash)> orgs_pdufa_reconciliation_with_http_info(id)

```ruby
begin
  # SEC vs Benzinga reconciliation
  data, status_code, headers = api_instance.orgs_pdufa_reconciliation_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsPdufaReconciliation200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_pdufa_reconciliation_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsPdufaReconciliation200Response**](OrgsPdufaReconciliation200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_peers

> <OrgsPeers200Response> orgs_peers(id)

Peer comparison

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'sponsor:b80857179c0f39712baa71f8183c3919' # String | 

begin
  # Peer comparison
  result = api_instance.orgs_peers(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_peers: #{e}"
end
```

#### Using the orgs_peers_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsPeers200Response>, Integer, Hash)> orgs_peers_with_http_info(id)

```ruby
begin
  # Peer comparison
  data, status_code, headers = api_instance.orgs_peers_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsPeers200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_peers_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsPeers200Response**](OrgsPeers200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_phase_status

> <OrgsPhaseStatus200Response> orgs_phase_status(id)

Phase status

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Phase status
  result = api_instance.orgs_phase_status(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_phase_status: #{e}"
end
```

#### Using the orgs_phase_status_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsPhaseStatus200Response>, Integer, Hash)> orgs_phase_status_with_http_info(id)

```ruby
begin
  # Phase status
  data, status_code, headers = api_instance.orgs_phase_status_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsPhaseStatus200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_phase_status_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsPhaseStatus200Response**](OrgsPhaseStatus200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_publishing_biomarkers

> <OrgsPublishingBiomarkers200Response> orgs_publishing_biomarkers(id)

Biomarkers this organization publishes on

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Biomarkers this organization publishes on
  result = api_instance.orgs_publishing_biomarkers(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_publishing_biomarkers: #{e}"
end
```

#### Using the orgs_publishing_biomarkers_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsPublishingBiomarkers200Response>, Integer, Hash)> orgs_publishing_biomarkers_with_http_info(id)

```ruby
begin
  # Biomarkers this organization publishes on
  data, status_code, headers = api_instance.orgs_publishing_biomarkers_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsPublishingBiomarkers200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_publishing_biomarkers_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsPublishingBiomarkers200Response**](OrgsPublishingBiomarkers200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_rare_activity

> <OrgsRareActivity200Response> orgs_rare_activity(id)

Rare disease activity

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0000200406' # String | 

begin
  # Rare disease activity
  result = api_instance.orgs_rare_activity(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_rare_activity: #{e}"
end
```

#### Using the orgs_rare_activity_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsRareActivity200Response>, Integer, Hash)> orgs_rare_activity_with_http_info(id)

```ruby
begin
  # Rare disease activity
  data, status_code, headers = api_instance.orgs_rare_activity_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsRareActivity200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_rare_activity_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsRareActivity200Response**](OrgsRareActivity200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_recalls

> <OrgsClia200Response> orgs_recalls(id)

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

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # FDA recalls
  result = api_instance.orgs_recalls(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_recalls: #{e}"
end
```

#### Using the orgs_recalls_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsClia200Response>, Integer, Hash)> orgs_recalls_with_http_info(id)

```ruby
begin
  # FDA recalls
  data, status_code, headers = api_instance.orgs_recalls_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsClia200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_recalls_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsClia200Response**](OrgsClia200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_registry_shell

> <OrgsRegistryShell200Response> orgs_registry_shell(id)

Operating-firm registry record

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'fdafirm:3010000003' # String | 

begin
  # Operating-firm registry record
  result = api_instance.orgs_registry_shell(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_registry_shell: #{e}"
end
```

#### Using the orgs_registry_shell_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsRegistryShell200Response>, Integer, Hash)> orgs_registry_shell_with_http_info(id)

```ruby
begin
  # Operating-firm registry record
  data, status_code, headers = api_instance.orgs_registry_shell_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsRegistryShell200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_registry_shell_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsRegistryShell200Response**](OrgsRegistryShell200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_related_biomarkers

> <OrgsMemberChildren200Response> orgs_related_biomarkers(id)

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

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Related biomarkers
  result = api_instance.orgs_related_biomarkers(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_related_biomarkers: #{e}"
end
```

#### Using the orgs_related_biomarkers_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsMemberChildren200Response>, Integer, Hash)> orgs_related_biomarkers_with_http_info(id)

```ruby
begin
  # Related biomarkers
  data, status_code, headers = api_instance.orgs_related_biomarkers_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsMemberChildren200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_related_biomarkers_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsMemberChildren200Response**](OrgsMemberChildren200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_related_conditions

> <OrgsRelatedConditions200Response> orgs_related_conditions(id)

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

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Related conditions
  result = api_instance.orgs_related_conditions(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_related_conditions: #{e}"
end
```

#### Using the orgs_related_conditions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsRelatedConditions200Response>, Integer, Hash)> orgs_related_conditions_with_http_info(id)

```ruby
begin
  # Related conditions
  data, status_code, headers = api_instance.orgs_related_conditions_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsRelatedConditions200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_related_conditions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsRelatedConditions200Response**](OrgsRelatedConditions200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_related_unfolded

> <OrgsRelatedUnfolded200Response> orgs_related_unfolded(id)

Related companies not grouped here

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'company:synth3499-bridge' # String | 

begin
  # Related companies not grouped here
  result = api_instance.orgs_related_unfolded(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_related_unfolded: #{e}"
end
```

#### Using the orgs_related_unfolded_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsRelatedUnfolded200Response>, Integer, Hash)> orgs_related_unfolded_with_http_info(id)

```ruby
begin
  # Related companies not grouped here
  data, status_code, headers = api_instance.orgs_related_unfolded_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsRelatedUnfolded200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_related_unfolded_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsRelatedUnfolded200Response**](OrgsRelatedUnfolded200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_revenue_by_drug

> <OrgsRevenueByDrug200Response> orgs_revenue_by_drug(id)

Revenue by drug

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Revenue by drug
  result = api_instance.orgs_revenue_by_drug(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_revenue_by_drug: #{e}"
end
```

#### Using the orgs_revenue_by_drug_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsRevenueByDrug200Response>, Integer, Hash)> orgs_revenue_by_drug_with_http_info(id)

```ruby
begin
  # Revenue by drug
  data, status_code, headers = api_instance.orgs_revenue_by_drug_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsRevenueByDrug200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_revenue_by_drug_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsRevenueByDrug200Response**](OrgsRevenueByDrug200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_revenue_geography

> <OrgsRevenueGeography200Response> orgs_revenue_geography(id)

Revenue by geography

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001048477' # String | 

begin
  # Revenue by geography
  result = api_instance.orgs_revenue_geography(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_revenue_geography: #{e}"
end
```

#### Using the orgs_revenue_geography_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsRevenueGeography200Response>, Integer, Hash)> orgs_revenue_geography_with_http_info(id)

```ruby
begin
  # Revenue by geography
  data, status_code, headers = api_instance.orgs_revenue_geography_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsRevenueGeography200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_revenue_geography_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsRevenueGeography200Response**](OrgsRevenueGeography200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_same_name_sec_filer

> <OrgsSameNameSecFiler200Response> orgs_same_name_sec_filer(id)

Same-name SEC filer

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0000059478' # String | 

begin
  # Same-name SEC filer
  result = api_instance.orgs_same_name_sec_filer(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_same_name_sec_filer: #{e}"
end
```

#### Using the orgs_same_name_sec_filer_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsSameNameSecFiler200Response>, Integer, Hash)> orgs_same_name_sec_filer_with_http_info(id)

```ruby
begin
  # Same-name SEC filer
  data, status_code, headers = api_instance.orgs_same_name_sec_filer_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsSameNameSecFiler200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_same_name_sec_filer_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsSameNameSecFiler200Response**](OrgsSameNameSecFiler200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_sbti_crosscheck

> <OrgsSbtiCrosscheck200Response> orgs_sbti_crosscheck(id)

SBTi cross-check

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0000353278' # String | 

begin
  # SBTi cross-check
  result = api_instance.orgs_sbti_crosscheck(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_sbti_crosscheck: #{e}"
end
```

#### Using the orgs_sbti_crosscheck_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsSbtiCrosscheck200Response>, Integer, Hash)> orgs_sbti_crosscheck_with_http_info(id)

```ruby
begin
  # SBTi cross-check
  data, status_code, headers = api_instance.orgs_sbti_crosscheck_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsSbtiCrosscheck200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_sbti_crosscheck_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsSbtiCrosscheck200Response**](OrgsSbtiCrosscheck200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_scientist_com

> <OrgsScientistCom200Response> orgs_scientist_com(id)

On Scientist.com

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # On Scientist.com
  result = api_instance.orgs_scientist_com(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_scientist_com: #{e}"
end
```

#### Using the orgs_scientist_com_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsScientistCom200Response>, Integer, Hash)> orgs_scientist_com_with_http_info(id)

```ruby
begin
  # On Scientist.com
  data, status_code, headers = api_instance.orgs_scientist_com_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsScientistCom200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_scientist_com_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsScientistCom200Response**](OrgsScientistCom200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_sec_filer_shell

> <OrgsSecFilerShell200Response> orgs_sec_filer_shell(id)

SEC filer record

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'secfiler:0001880000' # String | 

begin
  # SEC filer record
  result = api_instance.orgs_sec_filer_shell(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_sec_filer_shell: #{e}"
end
```

#### Using the orgs_sec_filer_shell_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsSecFilerShell200Response>, Integer, Hash)> orgs_sec_filer_shell_with_http_info(id)

```ruby
begin
  # SEC filer record
  data, status_code, headers = api_instance.orgs_sec_filer_shell_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsSecFilerShell200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_sec_filer_shell_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsSecFilerShell200Response**](OrgsSecFilerShell200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_sec_pdufa

> <OrgsSecPdufa200Response> orgs_sec_pdufa(id)

SEC PDUFA decisions

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # SEC PDUFA decisions
  result = api_instance.orgs_sec_pdufa(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_sec_pdufa: #{e}"
end
```

#### Using the orgs_sec_pdufa_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsSecPdufa200Response>, Integer, Hash)> orgs_sec_pdufa_with_http_info(id)

```ruby
begin
  # SEC PDUFA decisions
  data, status_code, headers = api_instance.orgs_sec_pdufa_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsSecPdufa200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_sec_pdufa_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsSecPdufa200Response**](OrgsSecPdufa200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_sponsor_delivery

> <OrgsMemberChildren200Response> orgs_sponsor_delivery(id)

Sponsor delivery record

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Sponsor delivery record
  result = api_instance.orgs_sponsor_delivery(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_sponsor_delivery: #{e}"
end
```

#### Using the orgs_sponsor_delivery_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsMemberChildren200Response>, Integer, Hash)> orgs_sponsor_delivery_with_http_info(id)

```ruby
begin
  # Sponsor delivery record
  data, status_code, headers = api_instance.orgs_sponsor_delivery_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsMemberChildren200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_sponsor_delivery_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsMemberChildren200Response**](OrgsMemberChildren200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_supplier_links

> <OrgsSupplierLinks200Response> orgs_supplier_links

Suppliers matching your search

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new

begin
  # Suppliers matching your search
  result = api_instance.orgs_supplier_links
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_supplier_links: #{e}"
end
```

#### Using the orgs_supplier_links_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsSupplierLinks200Response>, Integer, Hash)> orgs_supplier_links_with_http_info

```ruby
begin
  # Suppliers matching your search
  data, status_code, headers = api_instance.orgs_supplier_links_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsSupplierLinks200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_supplier_links_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**OrgsSupplierLinks200Response**](OrgsSupplierLinks200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_timeline

> <OrgsTimeline200Response> orgs_timeline(id)

Latest dated events

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Latest dated events
  result = api_instance.orgs_timeline(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_timeline: #{e}"
end
```

#### Using the orgs_timeline_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsTimeline200Response>, Integer, Hash)> orgs_timeline_with_http_info(id)

```ruby
begin
  # Latest dated events
  data, status_code, headers = api_instance.orgs_timeline_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsTimeline200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_timeline_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsTimeline200Response**](OrgsTimeline200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_top_conditions

> <OrgsTopConditions200Response> orgs_top_conditions(id)

Top disease areas

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Top disease areas
  result = api_instance.orgs_top_conditions(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_top_conditions: #{e}"
end
```

#### Using the orgs_top_conditions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsTopConditions200Response>, Integer, Hash)> orgs_top_conditions_with_http_info(id)

```ruby
begin
  # Top disease areas
  data, status_code, headers = api_instance.orgs_top_conditions_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsTopConditions200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_top_conditions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsTopConditions200Response**](OrgsTopConditions200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_top_drugs

> <OrgsTopDrugs200Response> orgs_top_drugs(id)

Top drugs (all indications)

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Top drugs (all indications)
  result = api_instance.orgs_top_drugs(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_top_drugs: #{e}"
end
```

#### Using the orgs_top_drugs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsTopDrugs200Response>, Integer, Hash)> orgs_top_drugs_with_http_info(id)

```ruby
begin
  # Top drugs (all indications)
  data, status_code, headers = api_instance.orgs_top_drugs_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsTopDrugs200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_top_drugs_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsTopDrugs200Response**](OrgsTopDrugs200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_trials_per_year

> <OrgsTrialsPerYear200Response> orgs_trials_per_year(id)

Trials by start year as sponsor or collaborator

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Trials by start year as sponsor or collaborator
  result = api_instance.orgs_trials_per_year(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_trials_per_year: #{e}"
end
```

#### Using the orgs_trials_per_year_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsTrialsPerYear200Response>, Integer, Hash)> orgs_trials_per_year_with_http_info(id)

```ruby
begin
  # Trials by start year as sponsor or collaborator
  data, status_code, headers = api_instance.orgs_trials_per_year_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsTrialsPerYear200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_trials_per_year_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsTrialsPerYear200Response**](OrgsTrialsPerYear200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_unmatched_revenue

> <OrgsUnmatchedRevenue200Response> orgs_unmatched_revenue(id)

Product lines not matched to a brand

### Examples

```ruby
require 'time'
require 'pharmaobservatory'
# setup authorization
PharmaObservatory.configure do |config|
  # Configure Bearer authorization (JWT): ApiBearerJwt
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # Product lines not matched to a brand
  result = api_instance.orgs_unmatched_revenue(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_unmatched_revenue: #{e}"
end
```

#### Using the orgs_unmatched_revenue_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsUnmatchedRevenue200Response>, Integer, Hash)> orgs_unmatched_revenue_with_http_info(id)

```ruby
begin
  # Product lines not matched to a brand
  data, status_code, headers = api_instance.orgs_unmatched_revenue_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsUnmatchedRevenue200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_unmatched_revenue_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsUnmatchedRevenue200Response**](OrgsUnmatchedRevenue200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_untitled_letters

> <OrgsClia200Response> orgs_untitled_letters(id)

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

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # FDA untitled letters
  result = api_instance.orgs_untitled_letters(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_untitled_letters: #{e}"
end
```

#### Using the orgs_untitled_letters_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsClia200Response>, Integer, Hash)> orgs_untitled_letters_with_http_info(id)

```ruby
begin
  # FDA untitled letters
  data, status_code, headers = api_instance.orgs_untitled_letters_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsClia200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_untitled_letters_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsClia200Response**](OrgsClia200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgs_warning_letters

> <OrgsClia200Response> orgs_warning_letters(id)

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

api_instance = PharmaObservatory::OrgsApi.new
id = 'edgar:0001114448' # String | 

begin
  # FDA warning letters
  result = api_instance.orgs_warning_letters(id)
  p result
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_warning_letters: #{e}"
end
```

#### Using the orgs_warning_letters_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OrgsClia200Response>, Integer, Hash)> orgs_warning_letters_with_http_info(id)

```ruby
begin
  # FDA warning letters
  data, status_code, headers = api_instance.orgs_warning_letters_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OrgsClia200Response>
rescue PharmaObservatory::ApiError => e
  puts "Error when calling OrgsApi->orgs_warning_letters_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |

### Return type

[**OrgsClia200Response**](OrgsClia200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

