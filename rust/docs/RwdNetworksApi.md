# \RwdNetworksApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**rwd_networks_period_get**](RwdNetworksApi.md#rwd_networks_period_get) | **GET** /api/v1/rwd_networks/{id} | One real-world data register network
[**rwd_networks_period_list**](RwdNetworksApi.md#rwd_networks_period_list) | **GET** /api/v1/rwd_networks | List real-world data register networks



## rwd_networks_period_get

> models::RwdNetworksRecord rwd_networks_period_get(id)
One real-world data register network

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |

### Return type

[**models::RwdNetworksRecord**](RwdNetworksRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_networks_period_list

> models::RwdNetworksResponse rwd_networks_period_list(cursor, limit, sort, group_by, measure)
List real-world data register networks

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdNetworksResponse**](RwdNetworksResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

