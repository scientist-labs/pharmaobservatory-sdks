# \ApiV1ProgramCatalystApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_licensing_candidates_get**](ApiV1ProgramCatalystApi.md#api_v1_licensing_candidates_get) | **GET** /api/v1/licensing_candidates | licensing_candidates
[**api_v1_org_pipeline_stages_get**](ApiV1ProgramCatalystApi.md#api_v1_org_pipeline_stages_get) | **GET** /api/v1/org_pipeline_stages | pipeline_stages
[**api_v1_patent_cliff_exposure_get**](ApiV1ProgramCatalystApi.md#api_v1_patent_cliff_exposure_get) | **GET** /api/v1/patent_cliff_exposure | patent_cliff_exposure
[**api_v1_program_catalysts_get**](ApiV1ProgramCatalystApi.md#api_v1_program_catalysts_get) | **GET** /api/v1/program_catalysts | catalysts



## api_v1_licensing_candidates_get

> models::ApiV1LicensingCandidatesGet200Response api_v1_licensing_candidates_get(condition_id)
licensing_candidates

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**condition_id** | Option<**String**> |  |  |

### Return type

[**models::ApiV1LicensingCandidatesGet200Response**](_api_v1_licensing_candidates_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_org_pipeline_stages_get

> models::ApiV1OrgPipelineStagesGet200Response api_v1_org_pipeline_stages_get(company_id)
pipeline_stages

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**company_id** | Option<**String**> |  |  |

### Return type

[**models::ApiV1OrgPipelineStagesGet200Response**](_api_v1_org_pipeline_stages_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_patent_cliff_exposure_get

> models::ApiV1PatentCliffExposureGet200Response api_v1_patent_cliff_exposure_get(company_ids, drug_ids)
patent_cliff_exposure

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**company_ids** | Option<[**Vec<String>**](String.md)> |  |  |
**drug_ids** | Option<[**Vec<String>**](String.md)> |  |  |

### Return type

[**models::ApiV1PatentCliffExposureGet200Response**](_api_v1_patent_cliff_exposure_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_program_catalysts_get

> models::ApiV1ProgramCatalystsGet200Response api_v1_program_catalysts_get(company_ids, drug_ids)
catalysts

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**company_ids** | Option<[**Vec<String>**](String.md)> |  |  |
**drug_ids** | Option<[**Vec<String>**](String.md)> |  |  |

### Return type

[**models::ApiV1ProgramCatalystsGet200Response**](_api_v1_program_catalysts_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

