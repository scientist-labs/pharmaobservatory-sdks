# \BiomarkersApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**biomarkers_period_biomarker_bqp_assertions**](BiomarkersApi.md#biomarkers_period_biomarker_bqp_assertions) | **GET** /api/v1/biomarkers/{id}/biomarker_bqp_assertions | FDA biomarker qualification projects
[**biomarkers_period_biomarker_cdx_assertions**](BiomarkersApi.md#biomarkers_period_biomarker_cdx_assertions) | **GET** /api/v1/biomarkers/{id}/biomarker_cdx_assertions | FDA companion diagnostic assertions
[**biomarkers_period_biomarker_evidence**](BiomarkersApi.md#biomarkers_period_biomarker_evidence) | **GET** /api/v1/biomarkers/{id}/biomarker_evidence | Biomarker evidence
[**biomarkers_period_biomarker_literature**](BiomarkersApi.md#biomarkers_period_biomarker_literature) | **GET** /api/v1/biomarkers/{id}/biomarker_literature | Trial-cited literature mentions
[**biomarkers_period_biomarker_publishing_companies**](BiomarkersApi.md#biomarkers_period_biomarker_publishing_companies) | **GET** /api/v1/biomarkers/{id}/biomarker_publishing_companies | Companies publishing on a biomarker
[**biomarkers_period_biomarker_variant_mentions**](BiomarkersApi.md#biomarkers_period_biomarker_variant_mentions) | **GET** /api/v1/biomarkers/{id}/biomarker_variant_mentions | Protein variant mentions
[**biomarkers_period_get**](BiomarkersApi.md#biomarkers_period_get) | **GET** /api/v1/biomarkers/{id} | One biomarker
[**biomarkers_period_list**](BiomarkersApi.md#biomarkers_period_list) | **GET** /api/v1/biomarkers | List biomarkers



## biomarkers_period_biomarker_bqp_assertions

> models::BiomarkerBqpAssertionsResponse biomarkers_period_biomarker_bqp_assertions(id, limit, sort)
FDA biomarker qualification projects

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::BiomarkerBqpAssertionsResponse**](BiomarkerBqpAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_biomarker_cdx_assertions

> models::BiomarkerCdxAssertionsResponse biomarkers_period_biomarker_cdx_assertions(id, limit, sort)
FDA companion diagnostic assertions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::BiomarkerCdxAssertionsResponse**](BiomarkerCdxAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_biomarker_evidence

> models::BiomarkerEvidenceResponse biomarkers_period_biomarker_evidence(id, limit, sort)
Biomarker evidence

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::BiomarkerEvidenceResponse**](BiomarkerEvidenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_biomarker_literature

> models::BiomarkerLiteratureResponse biomarkers_period_biomarker_literature(id, limit, sort)
Trial-cited literature mentions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::BiomarkerLiteratureResponse**](BiomarkerLiteratureResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_biomarker_publishing_companies

> models::BiomarkerPublishingCompaniesResponse biomarkers_period_biomarker_publishing_companies(id, limit, sort)
Companies publishing on a biomarker

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::BiomarkerPublishingCompaniesResponse**](BiomarkerPublishingCompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_biomarker_variant_mentions

> models::BiomarkerVariantMentionsResponse biomarkers_period_biomarker_variant_mentions(id, limit, sort)
Protein variant mentions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::BiomarkerVariantMentionsResponse**](BiomarkerVariantMentionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_get

> models::BiomarkersRecord biomarkers_period_get(id)
One biomarker

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersRecord**](BiomarkersRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_list

> models::BiomarkersResponse biomarkers_period_list(limit, sort)
List biomarkers

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**limit** | Option<**i32**> |  |  |
**sort** | Option<**String**> |  |  |

### Return type

[**models::BiomarkersResponse**](BiomarkersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

