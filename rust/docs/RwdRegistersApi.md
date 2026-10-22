# \RwdRegistersApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**rwd_registers_period_get**](RwdRegistersApi.md#rwd_registers_period_get) | **GET** /api/v1/rwd_registers/{id} | One real-world data register
[**rwd_registers_period_list**](RwdRegistersApi.md#rwd_registers_period_list) | **GET** /api/v1/rwd_registers | List real-world data registers
[**rwd_registers_period_rwd_coverage**](RwdRegistersApi.md#rwd_registers_period_rwd_coverage) | **GET** /api/v1/rwd_registers/{id}/rwd_coverage | HMA-EMA catalogue denominator chains
[**rwd_registers_period_rwd_withheld_fields**](RwdRegistersApi.md#rwd_registers_period_rwd_withheld_fields) | **GET** /api/v1/rwd_registers/{id}/rwd_withheld_fields | HMA-EMA catalogue columns withheld from publication



## rwd_registers_period_get

> models::RwdRegistersRecord rwd_registers_period_get(id)
One real-world data register

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::RwdRegistersRecord**](RwdRegistersRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_registers_period_list

> models::RwdRegisterProvenanceResponse rwd_registers_period_list(cursor, limit, sort, group_by, measure)
List real-world data registers

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdRegisterProvenanceResponse**](RwdRegisterProvenanceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_registers_period_rwd_coverage

> models::RwdCoverageResponse rwd_registers_period_rwd_coverage(id, cursor, limit, sort, group_by, measure)
HMA-EMA catalogue denominator chains

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdCoverageResponse**](RwdCoverageResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_registers_period_rwd_withheld_fields

> models::RwdWithheldFieldsResponse rwd_registers_period_rwd_withheld_fields(id, cursor, limit, sort, group_by, measure)
HMA-EMA catalogue columns withheld from publication

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdWithheldFieldsResponse**](RwdWithheldFieldsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

