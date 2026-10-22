# RareDiseasePrevalenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**prevalence_assertion_id** | Option<**String**> | The identifier of the prevalence observation. | [optional]
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**prevalence_type** | Option<**String**> | The kind of observation: point prevalence, prevalence at birth, annual incidence, lifetime prevalence or cases and families. | [optional]
**prevalence_class** | Option<**String**> | The Orphanet prevalence class of the observation, for example 1-9 / 100 000; null when the source gives none. | [optional]
**prevalence_value_source** | Option<**String**> | The prevalence value as the source writes it. | [optional]
**normalized_prevalence_ratio** | Option<**f64**> | The observation as a ratio (cases per person) where one can be derived. | [optional]
**prevalence_lower_bound** | Option<**f64**> | The lower bound of the prevalence class as a ratio. | [optional]
**prevalence_upper_bound** | Option<**f64**> | The upper bound of the prevalence class as a ratio. | [optional]
**population_denominator** | Option<**f64**> | The denominator of the figure (for example 100000). | [optional]
**population_unit** | Option<**String**> | The unit of the denominator. | [optional]
**population_scope** | Option<**String**> | The population the observation covers. | [optional]
**age_scope** | Option<**String**> | The age group the observation covers. | [optional]
**qualification** | Option<**String**> | How the figure is qualified at the source (value and class, class only, cases, families). | [optional]
**validation_status** | Option<**String**> | validated when Orphanet marks the observation validated, not_yet_validated otherwise. | [optional]
**geography** | Option<**String**> | The geography the observation covers (Worldwide, a country or a region). | [optional]
**observation_date** | Option<[**String**](string.md)> | The date of the observation where the source gives one. | [optional]
**uncertainty_lower** | Option<**f64**> | The lower end of the stated uncertainty range. | [optional]
**uncertainty_upper** | Option<**f64**> | The upper end of the stated uncertainty range. | [optional]
**source** | Option<**String**> | The evidence source Orphanet cites for the observation. | [optional]
**source_snapshot_date** | Option<[**String**](string.md)> | The date of the Orphadata Science snapshot the row was read from. | [optional]
**source_url** | Option<**String**> | The URL of the source file the row was read from. | [optional]
**license_name** | Option<**String**> | The licence of the Orphadata Science file. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


