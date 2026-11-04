# \SubscriptionsApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**changes_period_list**](SubscriptionsApi.md#changes_period_list) | **GET** /api/v1/changes | Changes since a cursor
[**subscriptions_period_create**](SubscriptionsApi.md#subscriptions_period_create) | **POST** /api/v1/subscriptions | Follow an entity or a dataset query
[**subscriptions_period_delete**](SubscriptionsApi.md#subscriptions_period_delete) | **DELETE** /api/v1/subscriptions/{id} | Stop following
[**subscriptions_period_list**](SubscriptionsApi.md#subscriptions_period_list) | **GET** /api/v1/subscriptions | List your subscriptions



## changes_period_list

> models::ChangesList200Response changes_period_list(cursor, limit)
Changes since a cursor

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> |  |  |
**limit** | Option<**i32**> |  |  |

### Return type

[**models::ChangesList200Response**](changes_list_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## subscriptions_period_create

> models::SubscriptionsCreate201Response subscriptions_period_create(subscriptions_create_request)
Follow an entity or a dataset query

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**subscriptions_create_request** | Option<[**SubscriptionsCreateRequest**](SubscriptionsCreateRequest.md)> |  |  |

### Return type

[**models::SubscriptionsCreate201Response**](subscriptions_create_201_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## subscriptions_period_delete

> models::SubscriptionsDelete200Response subscriptions_period_delete(id)
Stop following

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |

### Return type

[**models::SubscriptionsDelete200Response**](subscriptions_delete_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## subscriptions_period_list

> models::SubscriptionsList200Response subscriptions_period_list()
List your subscriptions

### Parameters

This endpoint does not need any parameter.

### Return type

[**models::SubscriptionsList200Response**](subscriptions_list_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

