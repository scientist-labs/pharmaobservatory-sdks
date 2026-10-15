# \ApiV1BiomarkerAssayScreenApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_biomarker_assay_screen_get**](ApiV1BiomarkerAssayScreenApi.md#api_v1_biomarker_assay_screen_get) | **GET** /api/v1/biomarker_assay_screen | index



## api_v1_biomarker_assay_screen_get

> models::ApiV1BiomarkerAssayScreenGet200Response api_v1_biomarker_assay_screen_get(assay_provider_company, document_id, gene, gene_resolution_status, pmid, validation_type_mentioned)
index

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**assay_provider_company** | Option<**String**> |  |  |
**document_id** | Option<**String**> |  |  |
**gene** | Option<**String**> |  |  |
**gene_resolution_status** | Option<**String**> |  |  |
**pmid** | Option<**i32**> |  |  |
**validation_type_mentioned** | Option<**String**> |  |  |

### Return type

[**models::ApiV1BiomarkerAssayScreenGet200Response**](_api_v1_biomarker_assay_screen_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

