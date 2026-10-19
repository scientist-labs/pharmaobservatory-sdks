# \QueryApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**query_period_result**](QueryApi.md#query_period_result) | **GET** /api/v1/results/{result_id} | Page a stored result
[**query_period_run**](QueryApi.md#query_period_run) | **POST** /api/v1/query | Run a SQL query



## query_period_result

> models::ResultPage query_period_result(result_id, cursor)
Page a stored result

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**result_id** | **String** |  | [required] |
**cursor** | Option<**String**> |  |  |

### Return type

[**models::ResultPage**](ResultPage.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## query_period_run

> models::QueryResponse query_period_run(query_run_request)
Run a SQL query

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**query_run_request** | Option<[**QueryRunRequest**](QueryRunRequest.md)> |  |  |

### Return type

[**models::QueryResponse**](QueryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

