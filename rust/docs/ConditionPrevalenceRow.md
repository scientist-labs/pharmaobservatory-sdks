# ConditionPrevalenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**disease_prevalence_id** | Option<**String**> | The figure's key. | [optional]
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**condition_name** | Option<**String**> | The condition's MeSH heading. | [optional]
**condition_mapping_relation** | Option<**String**> | How the figure's disease maps to the condition: exact, narrower, broader, undetermined or unspecified. | [optional]
**disease_id** | Option<**String**> | The rare disease key of the figure's disease (orpha:<code>). | [optional]
**disease_name** | Option<**String**> | The disease's name in the source. | [optional]
**orpha_code** | Option<**String**> | The Orphanet disease number. | [optional]
**prevalence_type** | Option<**String**> | What the figure measures: point_prevalence, lifetime_prevalence, annual_incidence, prevalence_at_birth or cases_families. | [optional]
**prevalence_class** | Option<**String**> | The source's prevalence band, as filed. | [optional]
**prevalence_lower_bound** | Option<**f64**> | Lower bound of the prevalence band as a ratio of the population (0.0001 is 1 in 10,000). | [optional]
**prevalence_upper_bound** | Option<**f64**> | Upper bound of the prevalence band as a ratio of the population (0.0005 is 5 in 10,000). | [optional]
**population_denominator** | Option<**f64**> | The population size the source states the band against (for example 10000 for a band per 10,000 persons). | [optional]
**population_unit** | Option<**String**> | The unit of the population denominator. | [optional]
**qualification** | Option<**String**> | The source's qualification of the figure (class_only, value_and_class and so on). | [optional]
**validation_status** | Option<**String**> | The source's validation status for the figure (validated or not_yet_validated). | [optional]
**geography** | Option<**String**> | The geography the figure describes, as the source states it. | [optional]
**jurisdiction_code** | Option<**String**> | US, EU or GB where the geography is that jurisdiction; null otherwise. | [optional]
**observation_date** | Option<[**String**](string.md)> | The date the source assigns to the observation. | [optional]
**source_snapshot_date** | Option<[**String**](string.md)> | The date of the source release the figure was read from. | [optional]
**source** | Option<**String**> | The publication reference the source cites for the figure. | [optional]
**source_url** | Option<**String**> | The source release URL. | [optional]
**license_name** | Option<**String**> | The figure's licence. | [optional]
**attribution** | Option<**String**> | The credit line the licence requires. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


