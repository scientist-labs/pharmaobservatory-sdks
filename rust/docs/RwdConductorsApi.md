# \RwdConductorsApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**rwd_conductors_period_get**](RwdConductorsApi.md#rwd_conductors_period_get) | **GET** /api/v1/rwd_conductors/{id} | One real-world data study conductor
[**rwd_conductors_period_list**](RwdConductorsApi.md#rwd_conductors_period_list) | **GET** /api/v1/rwd_conductors | List real-world data study conductors
[**rwd_conductors_period_rwd_conductor_delivery**](RwdConductorsApi.md#rwd_conductors_period_rwd_conductor_delivery) | **GET** /api/v1/rwd_conductors/{id}/rwd_conductor_delivery | Real-world data study conductor delivery by mandate (HMA-EMA catalogue)
[**rwd_conductors_period_rwd_conductor_footprint**](RwdConductorsApi.md#rwd_conductors_period_rwd_conductor_footprint) | **GET** /api/v1/rwd_conductors/{id}/rwd_conductor_footprint | Real-world data study conductor footprint (HMA-EMA catalogue)
[**rwd_conductors_period_rwd_conductor_studies**](RwdConductorsApi.md#rwd_conductors_period_rwd_conductor_studies) | **GET** /api/v1/rwd_conductors/{id}/rwd_conductor_studies | Real-world data study conductor studies (HMA-EMA catalogue)



## rwd_conductors_period_get

> models::RwdConductorsRecord rwd_conductors_period_get(id)
One real-world data study conductor

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::RwdConductorsRecord**](RwdConductorsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_conductors_period_list

> models::RwdConductorsResponse rwd_conductors_period_list(cursor, limit, sort, group_by, measure)
List real-world data study conductors

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdConductorsResponse**](RwdConductorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_conductors_period_rwd_conductor_delivery

> models::RwdConductorDeliveryResponse rwd_conductors_period_rwd_conductor_delivery(id, cursor, limit, sort, group_by, measure)
Real-world data study conductor delivery by mandate (HMA-EMA catalogue)

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

[**models::RwdConductorDeliveryResponse**](RwdConductorDeliveryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_conductors_period_rwd_conductor_footprint

> models::RwdConductorFootprintResponse rwd_conductors_period_rwd_conductor_footprint(id, cursor, limit, sort, group_by, measure)
Real-world data study conductor footprint (HMA-EMA catalogue)

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

[**models::RwdConductorFootprintResponse**](RwdConductorFootprintResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_conductors_period_rwd_conductor_studies

> models::RwdConductorStudiesResponse rwd_conductors_period_rwd_conductor_studies(id, cursor, limit, sort, group_by, measure)
Real-world data study conductor studies (HMA-EMA catalogue)

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

[**models::RwdConductorStudiesResponse**](RwdConductorStudiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

