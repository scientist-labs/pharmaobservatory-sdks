# \DatasetsApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**datasets_period_get**](DatasetsApi.md#datasets_period_get) | **GET** /api/v1/datasets/{name} | Describe a dataset
[**datasets_period_list**](DatasetsApi.md#datasets_period_list) | **GET** /api/v1/datasets | List the datasets



## datasets_period_get

> std::collections::HashMap<String, serde_json::Value> datasets_period_get(name)
Describe a dataset

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**name** | **String** |  | [required] |

### Return type

[**std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## datasets_period_list

> models::DatasetSummaryList datasets_period_list()
List the datasets

### Parameters

This endpoint does not need any parameter.

### Return type

[**models::DatasetSummaryList**](DatasetSummaryList.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

