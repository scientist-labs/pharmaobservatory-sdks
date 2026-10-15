# \ApiV1BiomarkerAssayWhitespaceApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_biomarker_assay_whitespace_get**](ApiV1BiomarkerAssayWhitespaceApi.md#api_v1_biomarker_assay_whitespace_get) | **GET** /api/v1/biomarker_assay_whitespace | index



## api_v1_biomarker_assay_whitespace_get

> models::ApiV1BiomarkerAssayWhitespaceGet200Response api_v1_biomarker_assay_whitespace_get(bucket, gene, gene_evidence_tier, gene_supply_status, limit, q)
index

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**bucket** | Option<**String**> |  |  |
**gene** | Option<**String**> |  |  |
**gene_evidence_tier** | Option<**String**> |  |  |
**gene_supply_status** | Option<**String**> |  |  |
**limit** | Option<**i32**> |  |  |
**q** | Option<**String**> |  |  |

### Return type

[**models::ApiV1BiomarkerAssayWhitespaceGet200Response**](_api_v1_biomarker_assay_whitespace_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

