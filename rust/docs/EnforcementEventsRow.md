# EnforcementEventsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**event_id** | Option<**String**> | The event id: a hash of the lane and the lane's own key, unique across lanes. | [optional]
**source** | Option<**String**> | The lane the event came from: inspection_483 (an FDA inspection), import_refusal, warning_letter, untitled_letter, debarment or recall. | [optional]
**event_date** | Option<[**String**](string.md)> | The event date: the inspection date, refusal date, letter issue date (posted date when there is none), debarment effective date or recall initiation date. | [optional]
**domain** | Option<**String**> | The regulatory domain of the event, derived from the lane's own classifier: drug, biologic, device, veterinary, food, cosmetic, tobacco or other. | [optional]
**is_biomedical** | Option<**bool**> | True when domain is drug, biologic or device. | [optional]
**headline** | Option<**String**> | A one-line description: the inspection product type, the refused product, the letter subject, the debarment type or the recalled product. | [optional]
**detail** | Option<**String**> | The lane's supporting text: the cited regulations of an inspection, the refusal charges, the issuing office of a warning letter, the debarment term and citation, or the recall reason. | [optional]
**classification** | Option<**String**> | The inspection outcome (NAI no action indicated, VAI voluntary action indicated, OAI official action indicated); null for every other lane. | [optional]
**recall_class** | Option<**String**> | The recall severity class (Class I, II or III); null for every other lane. | [optional]
**had_483** | Option<**bool**> | True when the inspection ended with a Form 483 of observations; null for every other lane. | [optional]
**source_url** | Option<**String**> | The regulator's page for the event, where the lane has one. | [optional]
**source_ref** | Option<**String**> | The lane's own reference, where it has one: the letter id, debarment id or recall number. | [optional]
**device_class** | Option<**String**> | The FDA device class of an import-refusal product, from the device classification dimension; null for every other lane. | [optional]
**device_specialty** | Option<**String**> | The medical specialty of an import-refusal device product; null for every other lane. | [optional]
**is_low_signal_device** | Option<**bool**> | True for a device event the biomedical view hides by default (a low-signal commodity class). | [optional]
**firm_name** | Option<**String**> | The firm's name as filed with the regulator (the party named on the record, an individual for an individual debarment). | [optional]
**firm_org_id** | Option<**String**> | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional]
**firm_company_id** | Option<**String**> | Company level: the firm's company (null when unresolved); the entry key for the companies object type. | [optional]
**firm_company_name** | Option<**String**> | Display name of the firm's company. | [optional]
**firm_group_id** | Option<**String**> | Group level: the firm's current ultimate parent; a root's group is itself. | [optional]
**firm_group_name** | Option<**String**> | Display name of the firm's group. | [optional]
**match_tier** | Option<**String**> | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous or unresolved. | [optional]
**org_resolution_method** | Option<**String**> | The route that attached the firm: fei_bridge (the FDA establishment identifier bridge), fei_firm_mint or fei_firm_fold (the operating firm's registry record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional]
**is_org_resolved** | Option<**bool**> | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


