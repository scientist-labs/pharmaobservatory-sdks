# \ApiV1TrialApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_trials_discontinuation_reasons_get**](ApiV1TrialApi.md#api_v1_trials_discontinuation_reasons_get) | **GET** /api/v1/trials/discontinuation-reasons | discontinuation_reasons
[**api_v1_trials_eligibility_generalizability_scan_get**](ApiV1TrialApi.md#api_v1_trials_eligibility_generalizability_scan_get) | **GET** /api/v1/trials/eligibility-generalizability-scan | eligibility_generalizability_scan
[**api_v1_trials_history_and_prediction_get**](ApiV1TrialApi.md#api_v1_trials_history_and_prediction_get) | **GET** /api/v1/trials/history-and-prediction | history_and_prediction
[**api_v1_trials_records_get**](ApiV1TrialApi.md#api_v1_trials_records_get) | **GET** /api/v1/trials/records | records
[**api_v1_trials_reporting_completeness_get**](ApiV1TrialApi.md#api_v1_trials_reporting_completeness_get) | **GET** /api/v1/trials/reporting-completeness | reporting_completeness



## api_v1_trials_discontinuation_reasons_get

> models::ApiV1TrialsDiscontinuationReasonsGet200Response api_v1_trials_discontinuation_reasons_get(condition_id)
discontinuation_reasons

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**condition_id** | Option<**String**> |  |  |

### Return type

[**models::ApiV1TrialsDiscontinuationReasonsGet200Response**](_api_v1_trials_discontinuation_reasons_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_trials_eligibility_generalizability_scan_get

> models::ApiV1TrialsEligibilityGeneralizabilityScanGet200Response api_v1_trials_eligibility_generalizability_scan_get(condition_id, limit, nct_ids)
eligibility_generalizability_scan

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**condition_id** | Option<**String**> |  |  |
**limit** | Option<**i32**> |  |  |
**nct_ids** | Option<[**Vec<String>**](String.md)> |  |  |

### Return type

[**models::ApiV1TrialsEligibilityGeneralizabilityScanGet200Response**](_api_v1_trials_eligibility_generalizability_scan_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_trials_history_and_prediction_get

> models::ApiV1TrialsHistoryAndPredictionGet200Response api_v1_trials_history_and_prediction_get(nct_id)
history_and_prediction

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**nct_id** | Option<**String**> |  |  |

### Return type

[**models::ApiV1TrialsHistoryAndPredictionGet200Response**](_api_v1_trials_history_and_prediction_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_trials_records_get

> models::ApiV1TrialsRecordsGet200Response api_v1_trials_records_get(include_extracted, nct_ids)
records

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**include_extracted** | Option<**bool**> |  |  |
**nct_ids** | Option<[**Vec<String>**](String.md)> |  |  |

### Return type

[**models::ApiV1TrialsRecordsGet200Response**](_api_v1_trials_records_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_trials_reporting_completeness_get

> models::ApiV1TrialsReportingCompletenessGet200Response api_v1_trials_reporting_completeness_get(condition_id)
reporting_completeness

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**condition_id** | Option<**String**> |  |  |

### Return type

[**models::ApiV1TrialsReportingCompletenessGet200Response**](_api_v1_trials_reporting_completeness_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

