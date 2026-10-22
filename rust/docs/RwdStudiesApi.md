# \RwdStudiesApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**rwd_studies_period_get**](RwdStudiesApi.md#rwd_studies_period_get) | **GET** /api/v1/rwd_studies/{id} | One real-world data study
[**rwd_studies_period_list**](RwdStudiesApi.md#rwd_studies_period_list) | **GET** /api/v1/rwd_studies | List real-world data studies
[**rwd_studies_period_rwd_conductor_studies**](RwdStudiesApi.md#rwd_studies_period_rwd_conductor_studies) | **GET** /api/v1/rwd_studies/{id}/rwd_conductor_studies | Real-world data study conductor studies (HMA-EMA catalogue)
[**rwd_studies_period_rwd_study_conditions**](RwdStudiesApi.md#rwd_studies_period_rwd_study_conditions) | **GET** /api/v1/rwd_studies/{id}/rwd_study_conditions | Registered real-world data study conditions
[**rwd_studies_period_rwd_study_data_sources**](RwdStudiesApi.md#rwd_studies_period_rwd_study_data_sources) | **GET** /api/v1/rwd_studies/{id}/rwd_study_data_sources | Registered real-world data study data sources
[**rwd_studies_period_rwd_study_drugs**](RwdStudiesApi.md#rwd_studies_period_rwd_study_drugs) | **GET** /api/v1/rwd_studies/{id}/rwd_study_drugs | Registered real-world data study drugs
[**rwd_studies_period_rwd_study_milestones**](RwdStudiesApi.md#rwd_studies_period_rwd_study_milestones) | **GET** /api/v1/rwd_studies/{id}/rwd_study_milestones | Registered real-world data study milestones



## rwd_studies_period_get

> models::RwdStudiesRecord rwd_studies_period_get(id)
One real-world data study

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |

### Return type

[**models::RwdStudiesRecord**](RwdStudiesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_studies_period_list

> models::RwdStudiesResponse rwd_studies_period_list(cursor, limit, sort, group_by, measure)
List real-world data studies

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdStudiesResponse**](RwdStudiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_studies_period_rwd_conductor_studies

> models::RwdConductorStudiesResponse rwd_studies_period_rwd_conductor_studies(id, cursor, limit, sort, group_by, measure)
Real-world data study conductor studies (HMA-EMA catalogue)

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |
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


## rwd_studies_period_rwd_study_conditions

> models::RwdStudyConditionsResponse rwd_studies_period_rwd_study_conditions(id, cursor, limit, sort, group_by, measure)
Registered real-world data study conditions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdStudyConditionsResponse**](RwdStudyConditionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_studies_period_rwd_study_data_sources

> models::RwdStudyDataSourcesResponse rwd_studies_period_rwd_study_data_sources(id, cursor, limit, sort, group_by, measure)
Registered real-world data study data sources

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdStudyDataSourcesResponse**](RwdStudyDataSourcesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_studies_period_rwd_study_drugs

> models::RwdStudyDrugsResponse rwd_studies_period_rwd_study_drugs(id, cursor, limit, sort, group_by, measure)
Registered real-world data study drugs

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdStudyDrugsResponse**](RwdStudyDrugsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_studies_period_rwd_study_milestones

> models::RwdStudyMilestonesResponse rwd_studies_period_rwd_study_milestones(id, cursor, limit, sort, group_by, measure)
Registered real-world data study milestones

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdStudyMilestonesResponse**](RwdStudyMilestonesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

