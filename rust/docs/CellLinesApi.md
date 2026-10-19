# \CellLinesApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cell_lines_period_get**](CellLinesApi.md#cell_lines_period_get) | **GET** /api/v1/cell_lines/{id} | One cell line
[**cell_lines_period_list**](CellLinesApi.md#cell_lines_period_list) | **GET** /api/v1/cell_lines | List cell lines



## cell_lines_period_get

> models::CellLinesRecord cell_lines_period_get(id)
One cell line

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::CellLinesRecord**](CellLinesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## cell_lines_period_list

> models::CellLinesResponse cell_lines_period_list(limit, sort)
List cell lines

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::CellLinesResponse**](CellLinesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

