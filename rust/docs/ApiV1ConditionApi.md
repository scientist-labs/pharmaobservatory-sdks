# \ApiV1ConditionApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_conditions_condition_id_burden_vs_evidence_get**](ApiV1ConditionApi.md#api_v1_conditions_condition_id_burden_vs_evidence_get) | **GET** /api/v1/conditions/{condition_id}/burden_vs_evidence | burden_vs_evidence
[**api_v1_conditions_condition_id_evidence_freshness_get**](ApiV1ConditionApi.md#api_v1_conditions_condition_id_evidence_freshness_get) | **GET** /api/v1/conditions/{condition_id}/evidence_freshness | evidence_freshness
[**api_v1_conditions_condition_id_sponsor_landscape_get**](ApiV1ConditionApi.md#api_v1_conditions_condition_id_sponsor_landscape_get) | **GET** /api/v1/conditions/{condition_id}/sponsor_landscape | sponsor_landscape



## api_v1_conditions_condition_id_burden_vs_evidence_get

> models::ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response api_v1_conditions_condition_id_burden_vs_evidence_get(condition_id)
burden_vs_evidence

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**condition_id** | **String** |  | [required] |

### Return type

[**models::ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response**](_api_v1_conditions__condition_id__burden_vs_evidence_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_conditions_condition_id_evidence_freshness_get

> models::ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response api_v1_conditions_condition_id_evidence_freshness_get(condition_id)
evidence_freshness

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**condition_id** | **String** |  | [required] |

### Return type

[**models::ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response**](_api_v1_conditions__condition_id__evidence_freshness_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_conditions_condition_id_sponsor_landscape_get

> models::ApiV1ConditionsConditionIdSponsorLandscapeGet200Response api_v1_conditions_condition_id_sponsor_landscape_get(condition_id)
sponsor_landscape

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**condition_id** | **String** |  | [required] |

### Return type

[**models::ApiV1ConditionsConditionIdSponsorLandscapeGet200Response**](_api_v1_conditions__condition_id__sponsor_landscape_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

