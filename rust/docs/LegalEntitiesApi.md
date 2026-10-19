# \LegalEntitiesApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**legal_entities_period_company_programs**](LegalEntitiesApi.md#legal_entities_period_company_programs) | **GET** /api/v1/legal_entities/{id}/company_programs | Company programs
[**legal_entities_period_drug_patent_cliff**](LegalEntitiesApi.md#legal_entities_period_drug_patent_cliff) | **GET** /api/v1/legal_entities/{id}/drug_patent_cliff | Patent-cliff exposure
[**legal_entities_period_get**](LegalEntitiesApi.md#legal_entities_period_get) | **GET** /api/v1/legal_entities/{id} | One legal entity
[**legal_entities_period_institutional_holdings**](LegalEntitiesApi.md#legal_entities_period_institutional_holdings) | **GET** /api/v1/legal_entities/{id}/institutional_holdings | Institutional holdings (SEC Form 13F)
[**legal_entities_period_list**](LegalEntitiesApi.md#legal_entities_period_list) | **GET** /api/v1/legal_entities | List legal entities
[**legal_entities_period_pdufa_dates**](LegalEntitiesApi.md#legal_entities_period_pdufa_dates) | **GET** /api/v1/legal_entities/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts
[**legal_entities_period_sponsor_recruitment_by_phase**](LegalEntitiesApi.md#legal_entities_period_sponsor_recruitment_by_phase) | **GET** /api/v1/legal_entities/{id}/sponsor_recruitment_by_phase | Trial recruitment duration by sponsor and phase
[**legal_entities_period_sponsor_recruitment_duration**](LegalEntitiesApi.md#legal_entities_period_sponsor_recruitment_duration) | **GET** /api/v1/legal_entities/{id}/sponsor_recruitment_duration | Trial recruitment duration by sponsor
[**legal_entities_period_sponsor_trial_start_forecast**](LegalEntitiesApi.md#legal_entities_period_sponsor_trial_start_forecast) | **GET** /api/v1/legal_entities/{id}/sponsor_trial_start_forecast | Trial-start forecast by sponsor
[**legal_entities_period_sponsor_trials**](LegalEntitiesApi.md#legal_entities_period_sponsor_trials) | **GET** /api/v1/legal_entities/{id}/sponsor_trials | Sponsor trials



## legal_entities_period_company_programs

> models::CompanyProgramsResponse legal_entities_period_company_programs(id, limit, sort)
Company programs

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::CompanyProgramsResponse**](CompanyProgramsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## legal_entities_period_drug_patent_cliff

> models::DrugPatentCliffResponse legal_entities_period_drug_patent_cliff(id, limit, sort)
Patent-cliff exposure

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::DrugPatentCliffResponse**](DrugPatentCliffResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## legal_entities_period_get

> models::LegalEntitiesRecord legal_entities_period_get(id)
One legal entity

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::LegalEntitiesRecord**](LegalEntitiesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## legal_entities_period_institutional_holdings

> models::InstitutionalHoldingsResponse legal_entities_period_institutional_holdings(id, limit, sort)
Institutional holdings (SEC Form 13F)

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::InstitutionalHoldingsResponse**](InstitutionalHoldingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## legal_entities_period_list

> models::LegalEntitiesResponse legal_entities_period_list(limit, sort)
List legal entities

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::LegalEntitiesResponse**](LegalEntitiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## legal_entities_period_pdufa_dates

> models::PdufaDatesResponse legal_entities_period_pdufa_dates(id, limit, sort)
FDA decision dates (PDUFA) and catalysts

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::PdufaDatesResponse**](PdufaDatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## legal_entities_period_sponsor_recruitment_by_phase

> models::SponsorRecruitmentByPhaseResponse legal_entities_period_sponsor_recruitment_by_phase(id, limit, sort)
Trial recruitment duration by sponsor and phase

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::SponsorRecruitmentByPhaseResponse**](SponsorRecruitmentByPhaseResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## legal_entities_period_sponsor_recruitment_duration

> models::SponsorRecruitmentDurationResponse legal_entities_period_sponsor_recruitment_duration(id, limit, sort)
Trial recruitment duration by sponsor

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::SponsorRecruitmentDurationResponse**](SponsorRecruitmentDurationResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## legal_entities_period_sponsor_trial_start_forecast

> models::SponsorTrialStartForecastResponse legal_entities_period_sponsor_trial_start_forecast(id, limit, sort)
Trial-start forecast by sponsor

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::SponsorTrialStartForecastResponse**](SponsorTrialStartForecastResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## legal_entities_period_sponsor_trials

> models::SponsorTrialsResponse legal_entities_period_sponsor_trials(id, limit, sort)
Sponsor trials

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::SponsorTrialsResponse**](SponsorTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

