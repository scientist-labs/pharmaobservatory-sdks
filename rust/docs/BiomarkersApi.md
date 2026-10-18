# \BiomarkersApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**biomarkers_period_assay_provider_ledger**](BiomarkersApi.md#biomarkers_period_assay_provider_ledger) | **GET** /api/v1/biomarkers/{id}/assay_provider_ledger | Assay-provider ledger
[**biomarkers_period_cgi**](BiomarkersApi.md#biomarkers_period_cgi) | **GET** /api/v1/biomarkers/{id}/cgi | Actionable biomarkers
[**biomarkers_period_coverage_by_traversal**](BiomarkersApi.md#biomarkers_period_coverage_by_traversal) | **GET** /api/v1/biomarkers/{id}/coverage_by_traversal | Coverage by traversal
[**biomarkers_period_evidence_over_time**](BiomarkersApi.md#biomarkers_period_evidence_over_time) | **GET** /api/v1/biomarkers/{id}/evidence_over_time | Evidence over time
[**biomarkers_period_fda_bqp**](BiomarkersApi.md#biomarkers_period_fda_bqp) | **GET** /api/v1/biomarkers/{id}/fda_bqp | FDA biomarker qualification projects
[**biomarkers_period_fda_cdx**](BiomarkersApi.md#biomarkers_period_fda_cdx) | **GET** /api/v1/biomarkers/{id}/fda_cdx | FDA companion diagnostic assertions
[**biomarkers_period_get**](BiomarkersApi.md#biomarkers_period_get) | **GET** /api/v1/biomarkers/{id} | One biomarkers record
[**biomarkers_period_hub_tier_tables**](BiomarkersApi.md#biomarkers_period_hub_tier_tables) | **GET** /api/v1/biomarkers/hub_tier_tables | Biomarker evidence, labeled by type
[**biomarkers_period_identity**](BiomarkersApi.md#biomarkers_period_identity) | **GET** /api/v1/biomarkers/{id}/identity | Biomarker identity
[**biomarkers_period_list**](BiomarkersApi.md#biomarkers_period_list) | **GET** /api/v1/biomarkers | Biomarkers
[**biomarkers_period_literature**](BiomarkersApi.md#biomarkers_period_literature) | **GET** /api/v1/biomarkers/{id}/literature | Trial-cited literature mentions
[**biomarkers_period_ot_association**](BiomarkersApi.md#biomarkers_period_ot_association) | **GET** /api/v1/biomarkers/{id}/ot_association | Target-disease associations
[**biomarkers_period_population_allele_frequency**](BiomarkersApi.md#biomarkers_period_population_allele_frequency) | **GET** /api/v1/biomarkers/{id}/population_allele_frequency | Population allele frequency
[**biomarkers_period_publishing_org**](BiomarkersApi.md#biomarkers_period_publishing_org) | **GET** /api/v1/biomarkers/{id}/publishing_org | Organizations publishing
[**biomarkers_period_related_biomarkers**](BiomarkersApi.md#biomarkers_period_related_biomarkers) | **GET** /api/v1/biomarkers/{id}/related_biomarkers | Related biomarkers
[**biomarkers_period_related_conditions**](BiomarkersApi.md#biomarkers_period_related_conditions) | **GET** /api/v1/biomarkers/{id}/related_conditions | Related conditions
[**biomarkers_period_related_organizations**](BiomarkersApi.md#biomarkers_period_related_organizations) | **GET** /api/v1/biomarkers/{id}/related_organizations | Related organisations
[**biomarkers_period_relations**](BiomarkersApi.md#biomarkers_period_relations) | **GET** /api/v1/biomarkers/{id}/relations | Literature-extracted relations
[**biomarkers_period_sample_screening**](BiomarkersApi.md#biomarkers_period_sample_screening) | **GET** /api/v1/biomarkers/{id}/sample_screening | Human sample screening reference
[**biomarkers_period_trial_mention**](BiomarkersApi.md#biomarkers_period_trial_mention) | **GET** /api/v1/biomarkers/{id}/trial_mention | Trial mentions
[**biomarkers_period_variant_mentions**](BiomarkersApi.md#biomarkers_period_variant_mentions) | **GET** /api/v1/biomarkers/{id}/variant_mentions | Protein variant mentions
[**biomarkers_period_vendor_sourcing**](BiomarkersApi.md#biomarkers_period_vendor_sourcing) | **GET** /api/v1/biomarkers/{id}/vendor_sourcing | Vendor sourcing
[**biomarkers_period_where_it_appears**](BiomarkersApi.md#biomarkers_period_where_it_appears) | **GET** /api/v1/biomarkers/{id}/where_it_appears | Where it appears



## biomarkers_period_assay_provider_ledger

> models::BiomarkersAssayProviderLedger200Response biomarkers_period_assay_provider_ledger(id)
Assay-provider ledger

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersAssayProviderLedger200Response**](biomarkers_assay_provider_ledger_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_cgi

> models::BiomarkersCgi200Response biomarkers_period_cgi(id)
Actionable biomarkers

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersCgi200Response**](biomarkers_cgi_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_coverage_by_traversal

> models::BiomarkersCoverageByTraversal200Response biomarkers_period_coverage_by_traversal(id)
Coverage by traversal

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersCoverageByTraversal200Response**](biomarkers_coverage_by_traversal_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_evidence_over_time

> models::BiomarkersEvidenceOverTime200Response biomarkers_period_evidence_over_time(id)
Evidence over time

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersEvidenceOverTime200Response**](biomarkers_evidence_over_time_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_fda_bqp

> models::BiomarkersFdaBqp200Response biomarkers_period_fda_bqp(id)
FDA biomarker qualification projects

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersFdaBqp200Response**](biomarkers_fda_bqp_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_fda_cdx

> models::BiomarkersFdaCdx200Response biomarkers_period_fda_cdx(id)
FDA companion diagnostic assertions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersFdaCdx200Response**](biomarkers_fda_cdx_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_get

> models::BiomarkersGet200Response biomarkers_period_get(id)
One biomarkers record

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersGet200Response**](biomarkers_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_hub_tier_tables

> models::BiomarkersHubTierTables200Response biomarkers_period_hub_tier_tables()
Biomarker evidence, labeled by type

### Parameters

This endpoint does not need any parameter.

### Return type

[**models::BiomarkersHubTierTables200Response**](biomarkers_hub_tier_tables_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_identity

> models::BiomarkersIdentity200Response biomarkers_period_identity(id)
Biomarker identity

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersIdentity200Response**](biomarkers_identity_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_list

> models::BiomarkersList200Response biomarkers_period_list(cursor, dir, evidence, facets, identity, kind, limit, q, sort, validated)
Biomarkers

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> |  |  |
**dir** | Option<**String**> |  |  |
**evidence** | Option<**i32**> |  |  |
**facets** | Option<**bool**> |  |  |
**identity** | Option<**String**> |  |  |
**kind** | Option<**String**> |  |  |
**limit** | Option<**i32**> |  |  |
**q** | Option<**String**> |  |  |
**sort** | Option<**String**> |  |  |
**validated** | Option<**i32**> |  |  |

### Return type

[**models::BiomarkersList200Response**](biomarkers_list_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_literature

> models::BiomarkersLiterature200Response biomarkers_period_literature(id)
Trial-cited literature mentions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersLiterature200Response**](biomarkers_literature_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_ot_association

> models::BiomarkersOtAssociation200Response biomarkers_period_ot_association(id)
Target-disease associations

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersOtAssociation200Response**](biomarkers_ot_association_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_population_allele_frequency

> models::BiomarkersPopulationAlleleFrequency200Response biomarkers_period_population_allele_frequency(id)
Population allele frequency

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersPopulationAlleleFrequency200Response**](biomarkers_population_allele_frequency_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_publishing_org

> models::BiomarkersPublishingOrg200Response biomarkers_period_publishing_org(id)
Organizations publishing

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersPublishingOrg200Response**](biomarkers_publishing_org_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_related_biomarkers

> models::BiomarkersRelatedBiomarkers200Response biomarkers_period_related_biomarkers(id)
Related biomarkers

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersRelatedBiomarkers200Response**](biomarkers_related_biomarkers_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_related_conditions

> models::BiomarkersRelatedConditions200Response biomarkers_period_related_conditions(id)
Related conditions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersRelatedConditions200Response**](biomarkers_related_conditions_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_related_organizations

> models::BiomarkersRelatedOrganizations200Response biomarkers_period_related_organizations(id)
Related organisations

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersRelatedOrganizations200Response**](biomarkers_related_organizations_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_relations

> models::BiomarkersRelations200Response biomarkers_period_relations(id, chemicals)
Literature-extracted relations

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**chemicals** | Option<**String**> |  |  |

### Return type

[**models::BiomarkersRelations200Response**](biomarkers_relations_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_sample_screening

> models::BiomarkersSampleScreening200Response biomarkers_period_sample_screening(id, positive_samples)
Human sample screening reference

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |
**positive_samples** | Option<**i32**> |  |  |

### Return type

[**models::BiomarkersSampleScreening200Response**](biomarkers_sample_screening_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_trial_mention

> models::BiomarkersTrialMention200Response biomarkers_period_trial_mention(id)
Trial mentions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersTrialMention200Response**](biomarkers_trial_mention_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_variant_mentions

> models::BiomarkersVariantMentions200Response biomarkers_period_variant_mentions(id)
Protein variant mentions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersVariantMentions200Response**](biomarkers_variant_mentions_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_vendor_sourcing

> models::BiomarkersVendorSourcing200Response biomarkers_period_vendor_sourcing(id)
Vendor sourcing

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersVendorSourcing200Response**](biomarkers_vendor_sourcing_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## biomarkers_period_where_it_appears

> models::BiomarkersWhereItAppears200Response biomarkers_period_where_it_appears(id)
Where it appears

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** |  | [required] |

### Return type

[**models::BiomarkersWhereItAppears200Response**](biomarkers_where_it_appears_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

