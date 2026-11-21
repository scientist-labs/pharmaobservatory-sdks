# \RwdUsNetworksApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**rwd_us_networks_period_get**](RwdUsNetworksApi.md#rwd_us_networks_period_get) | **GET** /api/v1/rwd_us_networks/{id} | One us real-world data network
[**rwd_us_networks_period_list**](RwdUsNetworksApi.md#rwd_us_networks_period_list) | **GET** /api/v1/rwd_us_networks | List us real-world data networks
[**rwd_us_networks_period_rwd_us_network_memberships**](RwdUsNetworksApi.md#rwd_us_networks_period_rwd_us_network_memberships) | **GET** /api/v1/rwd_us_networks/{id}/rwd_us_network_memberships | US real-world data network membership (N3C)
[**rwd_us_networks_period_rwd_us_network_projects**](RwdUsNetworksApi.md#rwd_us_networks_period_rwd_us_network_projects) | **GET** /api/v1/rwd_us_networks/{id}/rwd_us_network_projects | US real-world data network projects (N3C)



## rwd_us_networks_period_get

> models::RwdUsNetworksRecord rwd_us_networks_period_get(id)
One us real-world data network

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::RwdUsNetworksRecord**](RwdUsNetworksRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_us_networks_period_list

> models::RwdUsNetworksResponse rwd_us_networks_period_list(cursor, limit, sort, group_by, measure)
List us real-world data networks

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdUsNetworksResponse**](RwdUsNetworksResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_us_networks_period_rwd_us_network_memberships

> models::RwdUsNetworkMembershipsResponse rwd_us_networks_period_rwd_us_network_memberships(id, cursor, limit, sort, group_by, measure)
US real-world data network membership (N3C)

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

[**models::RwdUsNetworkMembershipsResponse**](RwdUsNetworkMembershipsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_us_networks_period_rwd_us_network_projects

> models::RwdUsNetworkProjectsResponse rwd_us_networks_period_rwd_us_network_projects(id, cursor, limit, sort, group_by, measure)
US real-world data network projects (N3C)

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

[**models::RwdUsNetworkProjectsResponse**](RwdUsNetworkProjectsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

