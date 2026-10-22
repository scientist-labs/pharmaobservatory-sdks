# \RwdSourcesApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**rwd_sources_period_get**](RwdSourcesApi.md#rwd_sources_period_get) | **GET** /api/v1/rwd_sources/{id} | One real-world data source
[**rwd_sources_period_list**](RwdSourcesApi.md#rwd_sources_period_list) | **GET** /api/v1/rwd_sources | List real-world data sources
[**rwd_sources_period_rwd_source_age_bands**](RwdSourcesApi.md#rwd_sources_period_rwd_source_age_bands) | **GET** /api/v1/rwd_sources/{id}/rwd_source_age_bands | Real-world data source age bands
[**rwd_sources_period_rwd_source_conditions**](RwdSourcesApi.md#rwd_sources_period_rwd_source_conditions) | **GET** /api/v1/rwd_sources/{id}/rwd_source_conditions | Real-world data sources named by registered studies of a condition
[**rwd_sources_period_rwd_source_etl_mappings**](RwdSourcesApi.md#rwd_sources_period_rwd_source_etl_mappings) | **GET** /api/v1/rwd_sources/{id}/rwd_source_etl_mappings | Real-world data source common data model mappings
[**rwd_sources_period_rwd_source_linkages**](RwdSourcesApi.md#rwd_sources_period_rwd_source_linkages) | **GET** /api/v1/rwd_sources/{id}/rwd_source_linkages | Real-world data source linkages
[**rwd_sources_period_rwd_source_types**](RwdSourcesApi.md#rwd_sources_period_rwd_source_types) | **GET** /api/v1/rwd_sources/{id}/rwd_source_types | Real-world data source types
[**rwd_sources_period_rwd_source_vocabularies**](RwdSourcesApi.md#rwd_sources_period_rwd_source_vocabularies) | **GET** /api/v1/rwd_sources/{id}/rwd_source_vocabularies | Real-world data source coding schemes
[**rwd_sources_period_rwd_study_data_sources**](RwdSourcesApi.md#rwd_sources_period_rwd_study_data_sources) | **GET** /api/v1/rwd_sources/{id}/rwd_study_data_sources | Registered real-world data study data sources



## rwd_sources_period_get

> models::RwdSourcesRecord rwd_sources_period_get(id)
One real-world data source

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |

### Return type

[**models::RwdSourcesRecord**](RwdSourcesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_sources_period_list

> models::RwdSourcesResponse rwd_sources_period_list(cursor, limit, sort, group_by, measure)
List real-world data sources

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RwdSourcesResponse**](RwdSourcesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_sources_period_rwd_source_age_bands

> models::RwdSourceAgeBandsResponse rwd_sources_period_rwd_source_age_bands(id, cursor, limit, sort, group_by, measure)
Real-world data source age bands

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

[**models::RwdSourceAgeBandsResponse**](RwdSourceAgeBandsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_sources_period_rwd_source_conditions

> models::RwdSourceConditionsResponse rwd_sources_period_rwd_source_conditions(id, cursor, limit, sort, group_by, measure)
Real-world data sources named by registered studies of a condition

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

[**models::RwdSourceConditionsResponse**](RwdSourceConditionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_sources_period_rwd_source_etl_mappings

> models::RwdSourceEtlMappingsResponse rwd_sources_period_rwd_source_etl_mappings(id, cursor, limit, sort, group_by, measure)
Real-world data source common data model mappings

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

[**models::RwdSourceEtlMappingsResponse**](RwdSourceEtlMappingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_sources_period_rwd_source_linkages

> models::RwdSourceLinkagesResponse rwd_sources_period_rwd_source_linkages(id, cursor, limit, sort, group_by, measure)
Real-world data source linkages

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

[**models::RwdSourceLinkagesResponse**](RwdSourceLinkagesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_sources_period_rwd_source_types

> models::RwdSourceTypesResponse rwd_sources_period_rwd_source_types(id, cursor, limit, sort, group_by, measure)
Real-world data source types

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

[**models::RwdSourceTypesResponse**](RwdSourceTypesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_sources_period_rwd_source_vocabularies

> models::RwdSourceVocabulariesResponse rwd_sources_period_rwd_source_vocabularies(id, cursor, limit, sort, group_by, measure)
Real-world data source coding schemes

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

[**models::RwdSourceVocabulariesResponse**](RwdSourceVocabulariesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rwd_sources_period_rwd_study_data_sources

> models::RwdStudyDataSourcesResponse rwd_sources_period_rwd_study_data_sources(id, cursor, limit, sort, group_by, measure)
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

