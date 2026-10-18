# \CellLinesApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cell_lines_period_get**](CellLinesApi.md#cell_lines_period_get) | **GET** /api/v1/cell-lines/{id} | One cell-lines record
[**cell_lines_period_identity**](CellLinesApi.md#cell_lines_period_identity) | **GET** /api/v1/cell-lines/{id}/identity | Identity and availability
[**cell_lines_period_list**](CellLinesApi.md#cell_lines_period_list) | **GET** /api/v1/cell-lines | Cell line identity safety
[**cell_lines_period_safety_resolution**](CellLinesApi.md#cell_lines_period_safety_resolution) | **GET** /api/v1/cell-lines/{id}/safety_resolution | Safety resolution
[**cell_lines_period_status_counts**](CellLinesApi.md#cell_lines_period_status_counts) | **GET** /api/v1/cell-lines/status_counts | Cell line safety summary



## cell_lines_period_get

> models::CellLinesGet200Response cell_lines_period_get(id)
One cell-lines record

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::CellLinesGet200Response**](cell_lines_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## cell_lines_period_identity

> models::CellLinesIdentity200Response cell_lines_period_identity(id)
Identity and availability

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::CellLinesIdentity200Response**](cell_lines_identity_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## cell_lines_period_list

> models::CellLinesList200Response cell_lines_period_list(cursor, facets, limit, safety_status)
Cell line identity safety

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> |  |  |
**facets** | Option<**bool**> |  |  |
**limit** | Option<**i32**> |  |  |
**safety_status** | Option<**String**> |  |  |

### Return type

[**models::CellLinesList200Response**](cell_lines_list_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## cell_lines_period_safety_resolution

> models::CellLinesSafetyResolution200Response cell_lines_period_safety_resolution(id)
Safety resolution

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::CellLinesSafetyResolution200Response**](cell_lines_safety_resolution_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## cell_lines_period_status_counts

> models::CellLinesStatusCounts200Response cell_lines_period_status_counts(species)
Cell line safety summary

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**species** | Option<**String**> |  |  |

### Return type

[**models::CellLinesStatusCounts200Response**](cell_lines_status_counts_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

