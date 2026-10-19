# \CompaniesApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**companies_period_biomarker_bqp_assertions**](CompaniesApi.md#companies_period_biomarker_bqp_assertions) | **GET** /api/v1/companies/{id}/biomarker_bqp_assertions | FDA biomarker qualification projects
[**companies_period_biomarker_cdx_assertions**](CompaniesApi.md#companies_period_biomarker_cdx_assertions) | **GET** /api/v1/companies/{id}/biomarker_cdx_assertions | FDA companion diagnostic assertions
[**companies_period_biomarker_publishing_companies**](CompaniesApi.md#companies_period_biomarker_publishing_companies) | **GET** /api/v1/companies/{id}/biomarker_publishing_companies | Companies publishing on a biomarker
[**companies_period_biomarker_validation**](CompaniesApi.md#companies_period_biomarker_validation) | **GET** /api/v1/companies/{id}/biomarker_validation | Biomarker validation ladder
[**companies_period_company_assets**](CompaniesApi.md#companies_period_company_assets) | **GET** /api/v1/companies/{id}/company_assets | Company assets (approved and marketed drugs)
[**companies_period_company_funding_events**](CompaniesApi.md#companies_period_company_funding_events) | **GET** /api/v1/companies/{id}/company_funding_events | Company funding events
[**companies_period_company_funding_summary**](CompaniesApi.md#companies_period_company_funding_summary) | **GET** /api/v1/companies/{id}/company_funding_summary | Company funding summary
[**companies_period_company_members**](CompaniesApi.md#companies_period_company_members) | **GET** /api/v1/companies/{id}/company_members | Company members
[**companies_period_company_programs**](CompaniesApi.md#companies_period_company_programs) | **GET** /api/v1/companies/{id}/company_programs | Company programs
[**companies_period_company_targets**](CompaniesApi.md#companies_period_company_targets) | **GET** /api/v1/companies/{id}/company_targets | Targets a company pursues
[**companies_period_drug_patent_cliff**](CompaniesApi.md#companies_period_drug_patent_cliff) | **GET** /api/v1/companies/{id}/drug_patent_cliff | Patent-cliff exposure
[**companies_period_get**](CompaniesApi.md#companies_period_get) | **GET** /api/v1/companies/{id} | One company
[**companies_period_institutional_holdings**](CompaniesApi.md#companies_period_institutional_holdings) | **GET** /api/v1/companies/{id}/institutional_holdings | Institutional holdings (SEC Form 13F)
[**companies_period_legal_entities**](CompaniesApi.md#companies_period_legal_entities) | **GET** /api/v1/companies/{id}/legal_entities | Legal entities
[**companies_period_list**](CompaniesApi.md#companies_period_list) | **GET** /api/v1/companies | List companies
[**companies_period_ownership_positions**](CompaniesApi.md#companies_period_ownership_positions) | **GET** /api/v1/companies/{id}/ownership_positions | Ownership and control positions
[**companies_period_pdufa_dates**](CompaniesApi.md#companies_period_pdufa_dates) | **GET** /api/v1/companies/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts
[**companies_period_sponsor_recruitment_by_phase**](CompaniesApi.md#companies_period_sponsor_recruitment_by_phase) | **GET** /api/v1/companies/{id}/sponsor_recruitment_by_phase | Trial recruitment duration by sponsor and phase
[**companies_period_sponsor_recruitment_duration**](CompaniesApi.md#companies_period_sponsor_recruitment_duration) | **GET** /api/v1/companies/{id}/sponsor_recruitment_duration | Trial recruitment duration by sponsor
[**companies_period_sponsor_trial_start_forecast**](CompaniesApi.md#companies_period_sponsor_trial_start_forecast) | **GET** /api/v1/companies/{id}/sponsor_trial_start_forecast | Trial-start forecast by sponsor
[**companies_period_sponsor_trials**](CompaniesApi.md#companies_period_sponsor_trials) | **GET** /api/v1/companies/{id}/sponsor_trials | Sponsor trials



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

