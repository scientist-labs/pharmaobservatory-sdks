# \SuppliersApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**suppliers_period_get**](SuppliersApi.md#suppliers_period_get) | **GET** /api/v1/suppliers/{id} | One supplier
[**suppliers_period_list**](SuppliersApi.md#suppliers_period_list) | **GET** /api/v1/suppliers | List suppliers
[**suppliers_period_provider_biomarker_capability**](SuppliersApi.md#suppliers_period_provider_biomarker_capability) | **GET** /api/v1/suppliers/{id}/provider_biomarker_capability | Supplier biomarker capability
[**suppliers_period_provider_biomarker_capability_summary**](SuppliersApi.md#suppliers_period_provider_biomarker_capability_summary) | **GET** /api/v1/suppliers/{id}/provider_biomarker_capability_summary | Supplier biomarker capability, summarised
[**suppliers_period_provider_clia_licensure**](SuppliersApi.md#suppliers_period_provider_clia_licensure) | **GET** /api/v1/suppliers/{id}/provider_clia_licensure | Supplier CLIA laboratory licensure
[**suppliers_period_provider_rwd_delivery**](SuppliersApi.md#suppliers_period_provider_rwd_delivery) | **GET** /api/v1/suppliers/{id}/provider_rwd_delivery | RX provider real-world data study delivery track record (HMA-EMA catalogue)
[**suppliers_period_provider_rwd_delivery_study**](SuppliersApi.md#suppliers_period_provider_rwd_delivery_study) | **GET** /api/v1/suppliers/{id}/provider_rwd_delivery_study | RX provider real-world data study final-report dates below the floor (HMA-EMA catalogue)
[**suppliers_period_provider_rwd_footprint**](SuppliersApi.md#suppliers_period_provider_rwd_footprint) | **GET** /api/v1/suppliers/{id}/provider_rwd_footprint | Supplier real-world data footprint (HMA-EMA catalogue)
[**suppliers_period_rwd_us_network_memberships**](SuppliersApi.md#suppliers_period_rwd_us_network_memberships) | **GET** /api/v1/suppliers/{id}/rwd_us_network_memberships | US real-world data network membership (N3C)
[**suppliers_period_supplier_rwd_declarations**](SuppliersApi.md#suppliers_period_supplier_rwd_declarations) | **GET** /api/v1/suppliers/{id}/supplier_rwd_declarations | Supplier declared real-world data facts
[**suppliers_period_supplier_services**](SuppliersApi.md#suppliers_period_supplier_services) | **GET** /api/v1/suppliers/{id}/supplier_services | Supplier services
[**suppliers_period_supplier_sites**](SuppliersApi.md#suppliers_period_supplier_sites) | **GET** /api/v1/suppliers/{id}/supplier_sites | Supplier sites



## suppliers_period_get

> models::SuppliersRecord suppliers_period_get(id)
One supplier

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |

### Return type

[**models::SuppliersRecord**](SuppliersRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_list

> models::SuppliersResponse suppliers_period_list(cursor, limit, sort, group_by, measure)
List suppliers

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::SuppliersResponse**](SuppliersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_provider_biomarker_capability

> models::ProviderBiomarkerCapabilityResponse suppliers_period_provider_biomarker_capability(id, cursor, limit, sort, group_by, measure)
Supplier biomarker capability

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

[**models::ProviderBiomarkerCapabilityResponse**](ProviderBiomarkerCapabilityResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_provider_biomarker_capability_summary

> models::ProviderBiomarkerCapabilitySummaryResponse suppliers_period_provider_biomarker_capability_summary(id, cursor, limit, sort, group_by, measure)
Supplier biomarker capability, summarised

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

[**models::ProviderBiomarkerCapabilitySummaryResponse**](ProviderBiomarkerCapabilitySummaryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_provider_clia_licensure

> models::ProviderCliaLicensureResponse suppliers_period_provider_clia_licensure(id, cursor, limit, sort, group_by, measure)
Supplier CLIA laboratory licensure

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

[**models::ProviderCliaLicensureResponse**](ProviderCliaLicensureResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_provider_rwd_delivery

> models::ProviderRwdDeliveryResponse suppliers_period_provider_rwd_delivery(id, cursor, limit, sort, group_by, measure)
RX provider real-world data study delivery track record (HMA-EMA catalogue)

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

[**models::ProviderRwdDeliveryResponse**](ProviderRwdDeliveryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_provider_rwd_delivery_study

> models::ProviderRwdDeliveryStudyResponse suppliers_period_provider_rwd_delivery_study(id, cursor, limit, sort, group_by, measure)
RX provider real-world data study final-report dates below the floor (HMA-EMA catalogue)

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

[**models::ProviderRwdDeliveryStudyResponse**](ProviderRwdDeliveryStudyResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_provider_rwd_footprint

> models::ProviderRwdFootprintResponse suppliers_period_provider_rwd_footprint(id, cursor, limit, sort, group_by, measure)
Supplier real-world data footprint (HMA-EMA catalogue)

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

[**models::ProviderRwdFootprintResponse**](ProviderRwdFootprintResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_rwd_us_network_memberships

> models::RwdUsNetworkMembershipsResponse suppliers_period_rwd_us_network_memberships(id, cursor, limit, sort, group_by, measure)
US real-world data network membership (N3C)

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

[**models::RwdUsNetworkMembershipsResponse**](RwdUsNetworkMembershipsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_supplier_rwd_declarations

> models::SupplierRwdDeclarationsResponse suppliers_period_supplier_rwd_declarations(id, cursor, limit, sort, group_by, measure)
Supplier declared real-world data facts

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

[**models::SupplierRwdDeclarationsResponse**](SupplierRwdDeclarationsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_supplier_services

> models::SupplierServicesResponse suppliers_period_supplier_services(id, cursor, limit, sort, group_by, measure)
Supplier services

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

[**models::SupplierServicesResponse**](SupplierServicesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## suppliers_period_supplier_sites

> models::SupplierSitesResponse suppliers_period_supplier_sites(id, cursor, limit, sort, group_by, measure)
Supplier sites

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

[**models::SupplierSitesResponse**](SupplierSitesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

