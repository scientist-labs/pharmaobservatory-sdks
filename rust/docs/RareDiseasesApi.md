# \RareDiseasesApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**rare_diseases_period_get**](RareDiseasesApi.md#rare_diseases_period_get) | **GET** /api/v1/rare_diseases/{id} | One rare disease
[**rare_diseases_period_list**](RareDiseasesApi.md#rare_diseases_period_list) | **GET** /api/v1/rare_diseases | List rare diseases
[**rare_diseases_period_rare_disease_genes**](RareDiseasesApi.md#rare_diseases_period_rare_disease_genes) | **GET** /api/v1/rare_diseases/{id}/rare_disease_genes | Rare disease genes
[**rare_diseases_period_rare_disease_hierarchy**](RareDiseasesApi.md#rare_diseases_period_rare_disease_hierarchy) | **GET** /api/v1/rare_diseases/{id}/rare_disease_hierarchy | Rare disease hierarchy
[**rare_diseases_period_rare_disease_investigator_coverage**](RareDiseasesApi.md#rare_diseases_period_rare_disease_investigator_coverage) | **GET** /api/v1/rare_diseases/{id}/rare_disease_investigator_coverage | Rare disease investigator coverage
[**rare_diseases_period_rare_disease_investigators**](RareDiseasesApi.md#rare_diseases_period_rare_disease_investigators) | **GET** /api/v1/rare_diseases/{id}/rare_disease_investigators | Rare disease investigators
[**rare_diseases_period_rare_disease_mappings**](RareDiseasesApi.md#rare_diseases_period_rare_disease_mappings) | **GET** /api/v1/rare_diseases/{id}/rare_disease_mappings | Rare disease mappings
[**rare_diseases_period_rare_disease_organizations**](RareDiseasesApi.md#rare_diseases_period_rare_disease_organizations) | **GET** /api/v1/rare_diseases/{id}/rare_disease_organizations | Rare disease organizations
[**rare_diseases_period_rare_disease_prevalence**](RareDiseasesApi.md#rare_diseases_period_rare_disease_prevalence) | **GET** /api/v1/rare_diseases/{id}/rare_disease_prevalence | Rare disease prevalence
[**rare_diseases_period_rare_disease_rarity**](RareDiseasesApi.md#rare_diseases_period_rare_disease_rarity) | **GET** /api/v1/rare_diseases/{id}/rare_disease_rarity | Rare disease rarity assessments
[**rare_diseases_period_rare_disease_rarity_evidence**](RareDiseasesApi.md#rare_diseases_period_rare_disease_rarity_evidence) | **GET** /api/v1/rare_diseases/{id}/rare_disease_rarity_evidence | Rare disease rarity evidence
[**rare_diseases_period_rare_disease_replacements**](RareDiseasesApi.md#rare_diseases_period_rare_disease_replacements) | **GET** /api/v1/rare_diseases/{id}/rare_disease_replacements | Rare disease replacements
[**rare_diseases_period_rare_disease_site_coverage**](RareDiseasesApi.md#rare_diseases_period_rare_disease_site_coverage) | **GET** /api/v1/rare_diseases/{id}/rare_disease_site_coverage | Rare disease site coverage
[**rare_diseases_period_rare_disease_sponsors**](RareDiseasesApi.md#rare_diseases_period_rare_disease_sponsors) | **GET** /api/v1/rare_diseases/{id}/rare_disease_sponsors | Rare disease sponsors
[**rare_diseases_period_rare_disease_trial_coverage**](RareDiseasesApi.md#rare_diseases_period_rare_disease_trial_coverage) | **GET** /api/v1/rare_diseases/{id}/rare_disease_trial_coverage | Rare disease trial coverage
[**rare_diseases_period_rare_disease_trial_sites**](RareDiseasesApi.md#rare_diseases_period_rare_disease_trial_sites) | **GET** /api/v1/rare_diseases/{id}/rare_disease_trial_sites | Rare disease trial sites
[**rare_diseases_period_rare_disease_trials**](RareDiseasesApi.md#rare_diseases_period_rare_disease_trials) | **GET** /api/v1/rare_diseases/{id}/rare_disease_trials | Rare disease trials



## rare_diseases_period_get

> models::RareDiseasesRecord rare_diseases_period_get(id)
One rare disease

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |

### Return type

[**models::RareDiseasesRecord**](RareDiseasesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_list

> models::RareDiseasesResponse rare_diseases_period_list(cursor, limit, sort, group_by, measure)
List rare diseases

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cursor** | Option<**String**> | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. |  |
**limit** | Option<**i32**> | Rows per page, up to 500 (default 50). |  |
**sort** | Option<**String**> | Comma separated columns, a leading minus for descending: sort=a,-b. |  |
**group_by** | Option<**String**> | A dimension to group by; a date takes :year, :quarter or :month. |  |
**measure** | Option<**String**> | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). |  |

### Return type

[**models::RareDiseasesResponse**](RareDiseasesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_genes

> models::RareDiseaseGenesResponse rare_diseases_period_rare_disease_genes(id, cursor, limit, sort, group_by, measure)
Rare disease genes

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

[**models::RareDiseaseGenesResponse**](RareDiseaseGenesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_hierarchy

> models::RareDiseaseHierarchyResponse rare_diseases_period_rare_disease_hierarchy(id, cursor, limit, sort, group_by, measure)
Rare disease hierarchy

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

[**models::RareDiseaseHierarchyResponse**](RareDiseaseHierarchyResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_investigator_coverage

> models::RareDiseaseInvestigatorCoverageResponse rare_diseases_period_rare_disease_investigator_coverage(id, cursor, limit, sort, group_by, measure)
Rare disease investigator coverage

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

[**models::RareDiseaseInvestigatorCoverageResponse**](RareDiseaseInvestigatorCoverageResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_investigators

> models::RareDiseaseInvestigatorsResponse rare_diseases_period_rare_disease_investigators(id, cursor, limit, sort, group_by, measure)
Rare disease investigators

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

[**models::RareDiseaseInvestigatorsResponse**](RareDiseaseInvestigatorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_mappings

> models::RareDiseaseMappingsResponse rare_diseases_period_rare_disease_mappings(id, cursor, limit, sort, group_by, measure)
Rare disease mappings

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

[**models::RareDiseaseMappingsResponse**](RareDiseaseMappingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_organizations

> models::RareDiseaseOrganizationsResponse rare_diseases_period_rare_disease_organizations(id, cursor, limit, sort, group_by, measure)
Rare disease organizations

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

[**models::RareDiseaseOrganizationsResponse**](RareDiseaseOrganizationsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_prevalence

> models::RareDiseasePrevalenceResponse rare_diseases_period_rare_disease_prevalence(id, cursor, limit, sort, group_by, measure)
Rare disease prevalence

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

[**models::RareDiseasePrevalenceResponse**](RareDiseasePrevalenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_rarity

> models::RareDiseaseRarityResponse rare_diseases_period_rare_disease_rarity(id, cursor, limit, sort, group_by, measure)
Rare disease rarity assessments

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

[**models::RareDiseaseRarityResponse**](RareDiseaseRarityResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_rarity_evidence

> models::RareDiseaseRarityEvidenceResponse rare_diseases_period_rare_disease_rarity_evidence(id, cursor, limit, sort, group_by, measure)
Rare disease rarity evidence

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

[**models::RareDiseaseRarityEvidenceResponse**](RareDiseaseRarityEvidenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_replacements

> models::RareDiseaseReplacementsResponse rare_diseases_period_rare_disease_replacements(id, cursor, limit, sort, group_by, measure)
Rare disease replacements

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

[**models::RareDiseaseReplacementsResponse**](RareDiseaseReplacementsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_site_coverage

> models::RareDiseaseSiteCoverageResponse rare_diseases_period_rare_disease_site_coverage(id, cursor, limit, sort, group_by, measure)
Rare disease site coverage

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

[**models::RareDiseaseSiteCoverageResponse**](RareDiseaseSiteCoverageResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_sponsors

> models::RareDiseaseSponsorsResponse rare_diseases_period_rare_disease_sponsors(id, cursor, limit, sort, group_by, measure)
Rare disease sponsors

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

[**models::RareDiseaseSponsorsResponse**](RareDiseaseSponsorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_trial_coverage

> models::RareDiseaseTrialCoverageResponse rare_diseases_period_rare_disease_trial_coverage(id, cursor, limit, sort, group_by, measure)
Rare disease trial coverage

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

[**models::RareDiseaseTrialCoverageResponse**](RareDiseaseTrialCoverageResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_trial_sites

> models::RareDiseaseTrialSitesResponse rare_diseases_period_rare_disease_trial_sites(id, cursor, limit, sort, group_by, measure)
Rare disease trial sites

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

[**models::RareDiseaseTrialSitesResponse**](RareDiseaseTrialSitesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rare_diseases_period_rare_disease_trials

> models::RareDiseaseTrialsResponse rare_diseases_period_rare_disease_trials(id, cursor, limit, sort, group_by, measure)
Rare disease trials

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

[**models::RareDiseaseTrialsResponse**](RareDiseaseTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

