# \CompaniesApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**companies_period_biomarker_bqp_assertions**](CompaniesApi.md#companies_period_biomarker_bqp_assertions) | **GET** /api/v1/companies/{id}/biomarker_bqp_assertions | FDA biomarker qualification projects
[**companies_period_biomarker_cdx_assertions**](CompaniesApi.md#companies_period_biomarker_cdx_assertions) | **GET** /api/v1/companies/{id}/biomarker_cdx_assertions | FDA companion diagnostic assertions
[**companies_period_biomarker_publishing_companies**](CompaniesApi.md#companies_period_biomarker_publishing_companies) | **GET** /api/v1/companies/{id}/biomarker_publishing_companies | Companies publishing on a biomarker
[**companies_period_company_members**](CompaniesApi.md#companies_period_company_members) | **GET** /api/v1/companies/{id}/company_members | Company members
[**companies_period_company_programs**](CompaniesApi.md#companies_period_company_programs) | **GET** /api/v1/companies/{id}/company_programs | Company programs
[**companies_period_company_targets**](CompaniesApi.md#companies_period_company_targets) | **GET** /api/v1/companies/{id}/company_targets | Targets a company pursues
[**companies_period_drug_patent_cliff**](CompaniesApi.md#companies_period_drug_patent_cliff) | **GET** /api/v1/companies/{id}/drug_patent_cliff | Patent-cliff exposure
[**companies_period_get**](CompaniesApi.md#companies_period_get) | **GET** /api/v1/companies/{id} | One company
[**companies_period_institutional_holdings**](CompaniesApi.md#companies_period_institutional_holdings) | **GET** /api/v1/companies/{id}/institutional_holdings | Institutional holdings (SEC Form 13F)
[**companies_period_legal_entities**](CompaniesApi.md#companies_period_legal_entities) | **GET** /api/v1/companies/{id}/legal_entities | Legal entities
[**companies_period_list**](CompaniesApi.md#companies_period_list) | **GET** /api/v1/companies | List companies
[**companies_period_pdufa_dates**](CompaniesApi.md#companies_period_pdufa_dates) | **GET** /api/v1/companies/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts
[**companies_period_sponsor_recruitment_by_phase**](CompaniesApi.md#companies_period_sponsor_recruitment_by_phase) | **GET** /api/v1/companies/{id}/sponsor_recruitment_by_phase | Trial recruitment duration by sponsor and phase
[**companies_period_sponsor_recruitment_duration**](CompaniesApi.md#companies_period_sponsor_recruitment_duration) | **GET** /api/v1/companies/{id}/sponsor_recruitment_duration | Trial recruitment duration by sponsor
[**companies_period_sponsor_trial_start_forecast**](CompaniesApi.md#companies_period_sponsor_trial_start_forecast) | **GET** /api/v1/companies/{id}/sponsor_trial_start_forecast | Trial-start forecast by sponsor
[**companies_period_sponsor_trials**](CompaniesApi.md#companies_period_sponsor_trials) | **GET** /api/v1/companies/{id}/sponsor_trials | Sponsor trials



## companies_period_biomarker_bqp_assertions

> models::BiomarkerBqpAssertionsResponse companies_period_biomarker_bqp_assertions(id, limit, scope, sort)
FDA biomarker qualification projects

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::BiomarkerBqpAssertionsResponse**](BiomarkerBqpAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_biomarker_cdx_assertions

> models::BiomarkerCdxAssertionsResponse companies_period_biomarker_cdx_assertions(id, limit, scope, sort)
FDA companion diagnostic assertions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::BiomarkerCdxAssertionsResponse**](BiomarkerCdxAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_biomarker_publishing_companies

> models::BiomarkerPublishingCompaniesResponse companies_period_biomarker_publishing_companies(id, limit, scope, sort)
Companies publishing on a biomarker

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::BiomarkerPublishingCompaniesResponse**](BiomarkerPublishingCompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_company_members

> models::CompanyMembersResponse companies_period_company_members(id, limit, scope, sort)
Company members

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::CompanyMembersResponse**](CompanyMembersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_company_programs

> models::CompanyProgramsResponse companies_period_company_programs(id, limit, scope, sort)
Company programs

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::CompanyProgramsResponse**](CompanyProgramsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_company_targets

> models::CompanyTargetsResponse companies_period_company_targets(id, limit, scope, sort)
Targets a company pursues

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::CompanyTargetsResponse**](CompanyTargetsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_drug_patent_cliff

> models::DrugPatentCliffResponse companies_period_drug_patent_cliff(id, limit, scope, sort)
Patent-cliff exposure

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

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

> models::InstitutionalHoldingsResponse companies_period_institutional_holdings(id, limit, scope, sort)
Institutional holdings (SEC Form 13F)

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::InstitutionalHoldingsResponse**](InstitutionalHoldingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_legal_entities

> models::LegalEntitiesResponse companies_period_legal_entities(id, limit, scope, sort)
Legal entities

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::LegalEntitiesResponse**](LegalEntitiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_list

> models::CompaniesResponse companies_period_list(limit, sort)
List companies

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::CompaniesResponse**](CompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_pdufa_dates

> models::PdufaDatesResponse companies_period_pdufa_dates(id, limit, scope, sort)
FDA decision dates (PDUFA) and catalysts

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::PdufaDatesResponse**](PdufaDatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_sponsor_recruitment_by_phase

> models::SponsorRecruitmentByPhaseResponse companies_period_sponsor_recruitment_by_phase(id, limit, scope, sort)
Trial recruitment duration by sponsor and phase

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::SponsorRecruitmentByPhaseResponse**](SponsorRecruitmentByPhaseResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_sponsor_recruitment_duration

> models::SponsorRecruitmentDurationResponse companies_period_sponsor_recruitment_duration(id, limit, scope, sort)
Trial recruitment duration by sponsor

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::SponsorRecruitmentDurationResponse**](SponsorRecruitmentDurationResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_sponsor_trial_start_forecast

> models::SponsorTrialStartForecastResponse companies_period_sponsor_trial_start_forecast(id, limit, scope, sort)
Trial-start forecast by sponsor

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::SponsorTrialStartForecastResponse**](SponsorTrialStartForecastResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## companies_period_sponsor_trials

> models::SponsorTrialsResponse companies_period_sponsor_trials(id, limit, scope, sort)
Sponsor trials

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**scope** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::SponsorTrialsResponse**](SponsorTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

