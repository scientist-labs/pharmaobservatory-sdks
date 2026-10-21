# \CompaniesApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**companies_period_biomarker_bqp_assertions**](CompaniesApi.md#companies_period_biomarker_bqp_assertions) | **GET** /api/v1/companies/{id}/biomarker_bqp_assertions | FDA biomarker qualification projects
[**companies_period_biomarker_cdx_assertions**](CompaniesApi.md#companies_period_biomarker_cdx_assertions) | **GET** /api/v1/companies/{id}/biomarker_cdx_assertions | FDA companion diagnostic assertions
[**companies_period_biomarker_publishing_companies**](CompaniesApi.md#companies_period_biomarker_publishing_companies) | **GET** /api/v1/companies/{id}/biomarker_publishing_companies | Companies publishing on a biomarker
[**companies_period_biomarker_validation**](CompaniesApi.md#companies_period_biomarker_validation) | **GET** /api/v1/companies/{id}/biomarker_validation | Biomarker validation ladder
[**companies_period_catalysts**](CompaniesApi.md#companies_period_catalysts) | **GET** /api/v1/companies/{id}/catalysts | Catalysts: trial readouts, patent cliffs, FDA decision dates and press-release events
[**companies_period_clia_certificates**](CompaniesApi.md#companies_period_clia_certificates) | **GET** /api/v1/companies/{id}/clia_certificates | CLIA laboratory certificates
[**companies_period_company_assets**](CompaniesApi.md#companies_period_company_assets) | **GET** /api/v1/companies/{id}/company_assets | Company assets (approved and marketed drugs)
[**companies_period_company_funding_events**](CompaniesApi.md#companies_period_company_funding_events) | **GET** /api/v1/companies/{id}/company_funding_events | Company funding events
[**companies_period_company_funding_summary**](CompaniesApi.md#companies_period_company_funding_summary) | **GET** /api/v1/companies/{id}/company_funding_summary | Company funding summary
[**companies_period_company_members**](CompaniesApi.md#companies_period_company_members) | **GET** /api/v1/companies/{id}/company_members | Company members
[**companies_period_company_programs**](CompaniesApi.md#companies_period_company_programs) | **GET** /api/v1/companies/{id}/company_programs | Company programs
[**companies_period_company_targets**](CompaniesApi.md#companies_period_company_targets) | **GET** /api/v1/companies/{id}/company_targets | Targets a company pursues
[**companies_period_cro_footprint**](CompaniesApi.md#companies_period_cro_footprint) | **GET** /api/v1/companies/{id}/cro_footprint | CRO footprint
[**companies_period_cro_trials**](CompaniesApi.md#companies_period_cro_trials) | **GET** /api/v1/companies/{id}/cro_trials | CRO trials
[**companies_period_debarments**](CompaniesApi.md#companies_period_debarments) | **GET** /api/v1/companies/{id}/debarments | FDA firm debarments
[**companies_period_drug_patent_cliff**](CompaniesApi.md#companies_period_drug_patent_cliff) | **GET** /api/v1/companies/{id}/drug_patent_cliff | Patent-cliff exposure
[**companies_period_enforcement_events**](CompaniesApi.md#companies_period_enforcement_events) | **GET** /api/v1/companies/{id}/enforcement_events | FDA enforcement events
[**companies_period_get**](CompaniesApi.md#companies_period_get) | **GET** /api/v1/companies/{id} | One company
[**companies_period_gxp_facilities**](CompaniesApi.md#companies_period_gxp_facilities) | **GET** /api/v1/companies/{id}/gxp_facilities | GxP facilities and quality signals
[**companies_period_import_refusals**](CompaniesApi.md#companies_period_import_refusals) | **GET** /api/v1/companies/{id}/import_refusals | FDA import refusals
[**companies_period_inspection_findings**](CompaniesApi.md#companies_period_inspection_findings) | **GET** /api/v1/companies/{id}/inspection_findings | FDA inspection findings
[**companies_period_inspections**](CompaniesApi.md#companies_period_inspections) | **GET** /api/v1/companies/{id}/inspections | FDA inspections
[**companies_period_institutional_holdings**](CompaniesApi.md#companies_period_institutional_holdings) | **GET** /api/v1/companies/{id}/institutional_holdings | Institutional holdings (SEC Form 13F)
[**companies_period_legal_entities**](CompaniesApi.md#companies_period_legal_entities) | **GET** /api/v1/companies/{id}/legal_entities | Legal entities
[**companies_period_list**](CompaniesApi.md#companies_period_list) | **GET** /api/v1/companies | List companies
[**companies_period_org_disclosures**](CompaniesApi.md#companies_period_org_disclosures) | **GET** /api/v1/companies/{id}/org_disclosures | Organization disclosures (climate and diversity)
[**companies_period_org_relationships**](CompaniesApi.md#companies_period_org_relationships) | **GET** /api/v1/companies/{id}/org_relationships | Organization relationships
[**companies_period_ownership_positions**](CompaniesApi.md#companies_period_ownership_positions) | **GET** /api/v1/companies/{id}/ownership_positions | Ownership and control positions
[**companies_period_pdufa_dates**](CompaniesApi.md#companies_period_pdufa_dates) | **GET** /api/v1/companies/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts
[**companies_period_recalls**](CompaniesApi.md#companies_period_recalls) | **GET** /api/v1/companies/{id}/recalls | FDA recalls
[**companies_period_sponsor_recruitment_by_phase**](CompaniesApi.md#companies_period_sponsor_recruitment_by_phase) | **GET** /api/v1/companies/{id}/sponsor_recruitment_by_phase | Trial recruitment duration by sponsor and phase
[**companies_period_sponsor_recruitment_duration**](CompaniesApi.md#companies_period_sponsor_recruitment_duration) | **GET** /api/v1/companies/{id}/sponsor_recruitment_duration | Trial recruitment duration by sponsor
[**companies_period_sponsor_trial_start_forecast**](CompaniesApi.md#companies_period_sponsor_trial_start_forecast) | **GET** /api/v1/companies/{id}/sponsor_trial_start_forecast | Trial-start forecast by sponsor
[**companies_period_sponsor_trials**](CompaniesApi.md#companies_period_sponsor_trials) | **GET** /api/v1/companies/{id}/sponsor_trials | Sponsor trials
[**companies_period_supplier_capability_esg**](CompaniesApi.md#companies_period_supplier_capability_esg) | **GET** /api/v1/companies/{id}/supplier_capability_esg | Supplier capability and ESG disclosure
[**companies_period_untitled_letters**](CompaniesApi.md#companies_period_untitled_letters) | **GET** /api/v1/companies/{id}/untitled_letters | FDA untitled letters
[**companies_period_warning_letters**](CompaniesApi.md#companies_period_warning_letters) | **GET** /api/v1/companies/{id}/warning_letters | FDA warning letters



## companies_period_biomarker_bqp_assertions

> models::BiomarkerBqpAssertionsResponse companies_period_biomarker_bqp_assertions(id, cursor, limit, sort, group_by, measure, scope)
FDA biomarker qualification projects

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::BiomarkerBqpAssertionsResponse**](BiomarkerBqpAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_biomarker_cdx_assertions

> models::BiomarkerCdxAssertionsResponse companies_period_biomarker_cdx_assertions(id, cursor, limit, sort, group_by, measure, scope)
FDA companion diagnostic assertions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::BiomarkerCdxAssertionsResponse**](BiomarkerCdxAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_biomarker_publishing_companies

> models::BiomarkerPublishingCompaniesResponse companies_period_biomarker_publishing_companies(id, cursor, limit, sort, group_by, measure, scope)
Companies publishing on a biomarker

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::BiomarkerPublishingCompaniesResponse**](BiomarkerPublishingCompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_biomarker_validation

> models::BiomarkerValidationResponse companies_period_biomarker_validation(id, cursor, limit, sort, group_by, measure, scope)
Biomarker validation ladder

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::BiomarkerValidationResponse**](BiomarkerValidationResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_catalysts

> models::CatalystsResponse companies_period_catalysts(id, cursor, limit, sort, group_by, measure, scope)
Catalysts: trial readouts, patent cliffs, FDA decision dates and press-release events

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::CatalystsResponse**](CatalystsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_clia_certificates

> models::CliaCertificatesResponse companies_period_clia_certificates(id, cursor, limit, sort, group_by, measure, scope)
CLIA laboratory certificates

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::CliaCertificatesResponse**](CliaCertificatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_company_assets

> models::CompanyAssetsResponse companies_period_company_assets(id, cursor, limit, sort, group_by, measure, scope)
Company assets (approved and marketed drugs)

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::CompanyAssetsResponse**](CompanyAssetsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_company_funding_events

> models::CompanyFundingEventsResponse companies_period_company_funding_events(id, cursor, limit, sort, group_by, measure, scope)
Company funding events

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::CompanyFundingEventsResponse**](CompanyFundingEventsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_company_funding_summary

> models::CompanyFundingSummaryResponse companies_period_company_funding_summary(id, cursor, limit, sort, group_by, measure, scope)
Company funding summary

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::CompanyFundingSummaryResponse**](CompanyFundingSummaryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_company_members

> models::CompanyMembersResponse companies_period_company_members(id, cursor, limit, sort, group_by, measure, scope)
Company members

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::CompanyMembersResponse**](CompanyMembersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_company_programs

> models::CompanyProgramsResponse companies_period_company_programs(id, cursor, limit, sort, group_by, measure, scope)
Company programs

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::CompanyProgramsResponse**](CompanyProgramsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_company_targets

> models::CompanyTargetsResponse companies_period_company_targets(id, cursor, limit, sort, group_by, measure, scope)
Targets a company pursues

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::CompanyTargetsResponse**](CompanyTargetsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_cro_footprint

> models::CroFootprintResponse companies_period_cro_footprint(id, cursor, limit, sort, group_by, measure, scope)
CRO footprint

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::CroFootprintResponse**](CroFootprintResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_cro_trials

> models::CroTrialsResponse companies_period_cro_trials(id, cursor, limit, sort, group_by, measure, scope)
CRO trials

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::CroTrialsResponse**](CroTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_debarments

> models::DebarmentsResponse companies_period_debarments(id, cursor, limit, sort, group_by, measure, scope)
FDA firm debarments

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::DebarmentsResponse**](DebarmentsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_drug_patent_cliff

> models::DrugPatentCliffResponse companies_period_drug_patent_cliff(id, cursor, limit, sort, group_by, measure, scope)
Patent-cliff exposure

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::DrugPatentCliffResponse**](DrugPatentCliffResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_enforcement_events

> models::EnforcementEventsResponse companies_period_enforcement_events(id, cursor, limit, sort, group_by, measure, scope)
FDA enforcement events

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::EnforcementEventsResponse**](EnforcementEventsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_get

> models::CompaniesRecord companies_period_get(id)
One company

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::CompaniesRecord**](CompaniesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_gxp_facilities

> models::GxpFacilitiesResponse companies_period_gxp_facilities(id, cursor, limit, sort, group_by, measure, scope)
GxP facilities and quality signals

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::GxpFacilitiesResponse**](GxpFacilitiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_import_refusals

> models::ImportRefusalsResponse companies_period_import_refusals(id, cursor, limit, sort, group_by, measure, scope)
FDA import refusals

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::ImportRefusalsResponse**](ImportRefusalsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_inspection_findings

> models::InspectionFindingsResponse companies_period_inspection_findings(id, cursor, limit, sort, group_by, measure, scope)
FDA inspection findings

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::InspectionFindingsResponse**](InspectionFindingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_inspections

> models::InspectionsResponse companies_period_inspections(id, cursor, limit, sort, group_by, measure, scope)
FDA inspections

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::InspectionsResponse**](InspectionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_institutional_holdings

> models::InstitutionalHoldingsResponse companies_period_institutional_holdings(id, cursor, limit, sort, group_by, measure, scope)
Institutional holdings (SEC Form 13F)

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::InstitutionalHoldingsResponse**](InstitutionalHoldingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_legal_entities

> models::LegalEntitiesResponse companies_period_legal_entities(id, cursor, limit, sort, group_by, measure, scope)
Legal entities

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::LegalEntitiesResponse**](LegalEntitiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_list

> models::CompaniesResponse companies_period_list(cursor, limit, sort, group_by, measure)
List companies

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::CompaniesResponse**](CompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_org_disclosures

> models::OrgDisclosuresResponse companies_period_org_disclosures(id, cursor, limit, sort, group_by, measure, scope)
Organization disclosures (climate and diversity)

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::OrgDisclosuresResponse**](OrgDisclosuresResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_org_relationships

> models::OrgRelationshipsResponse companies_period_org_relationships(id, cursor, limit, sort, group_by, measure, scope)
Organization relationships

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::OrgRelationshipsResponse**](OrgRelationshipsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_ownership_positions

> models::OwnershipPositionsResponse companies_period_ownership_positions(id, cursor, limit, sort, group_by, measure, scope)
Ownership and control positions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::OwnershipPositionsResponse**](OwnershipPositionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_pdufa_dates

> models::PdufaDatesResponse companies_period_pdufa_dates(id, cursor, limit, sort, group_by, measure, scope)
FDA decision dates (PDUFA) and catalysts

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::PdufaDatesResponse**](PdufaDatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_recalls

> models::RecallsResponse companies_period_recalls(id, cursor, limit, sort, group_by, measure, scope)
FDA recalls

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::RecallsResponse**](RecallsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_sponsor_recruitment_by_phase

> models::SponsorRecruitmentByPhaseResponse companies_period_sponsor_recruitment_by_phase(id, cursor, limit, sort, group_by, measure, scope)
Trial recruitment duration by sponsor and phase

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::SponsorRecruitmentByPhaseResponse**](SponsorRecruitmentByPhaseResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_sponsor_recruitment_duration

> models::SponsorRecruitmentDurationResponse companies_period_sponsor_recruitment_duration(id, cursor, limit, sort, group_by, measure, scope)
Trial recruitment duration by sponsor

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::SponsorRecruitmentDurationResponse**](SponsorRecruitmentDurationResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_sponsor_trial_start_forecast

> models::SponsorTrialStartForecastResponse companies_period_sponsor_trial_start_forecast(id, cursor, limit, sort, group_by, measure, scope)
Trial-start forecast by sponsor

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::SponsorTrialStartForecastResponse**](SponsorTrialStartForecastResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_sponsor_trials

> models::SponsorTrialsResponse companies_period_sponsor_trials(id, cursor, limit, sort, group_by, measure, scope)
Sponsor trials

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::SponsorTrialsResponse**](SponsorTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_supplier_capability_esg

> models::SupplierCapabilityEsgResponse companies_period_supplier_capability_esg(id, cursor, limit, sort, group_by, measure, scope)
Supplier capability and ESG disclosure

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::SupplierCapabilityEsgResponse**](SupplierCapabilityEsgResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_untitled_letters

> models::UntitledLettersResponse companies_period_untitled_letters(id, cursor, limit, sort, group_by, measure, scope)
FDA untitled letters

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::UntitledLettersResponse**](UntitledLettersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_warning_letters

> models::WarningLettersResponse companies_period_warning_letters(id, cursor, limit, sort, group_by, measure, scope)
FDA warning letters

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |
**scope** | Option<**String**> | company (this company only) or group (its corporate group). |  |

### Return type

[**models::WarningLettersResponse**](WarningLettersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

