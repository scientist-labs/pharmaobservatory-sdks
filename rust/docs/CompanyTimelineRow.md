# CompanyTimelineRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**timeline_event_id** | Option<**String**> | md5 of the lane, event type, source record and company: a stable surrogate id for the row. | [optional]
**company_id** | Option<**String**> | Company level: the company the row is about; the entry key for the companies object type. | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**reporting_org_id** | Option<**String**> | Legal-entity level: the organization record the lane reported the event on (the lowest id when several records of the company report it). | [optional]
**source_key** | Option<**String**> | The lane that supplied the event. | [optional]
**source_entity_key** | Option<**String**> | The lane's own key for the thing that happened (trial id, funding event id, patent id, enforcement event id, statement id). | [optional]
**event_type** | Option<**String**> | The fine event vocabulary; funding types are the funding round types verbatim. | [optional]
**event_family** | Option<**String**> | The family of the event type, assigned by name from the lane registry: clinical, capital, intellectual_property, regulatory, compliance or other. | [optional]
**event_date** | Option<[**String**](string.md)> | The first day consistent with the evidence, never the raw value when the evidence is coarser than a day. | [optional]
**date_precision** | Option<**String**> | How precise the evidence is: day, month, quarter, half or year. | [optional]
**precision_window_end** | Option<[**String**](string.md)> | The last day consistent with the evidence; equals event_date at day precision. | [optional]
**date_precision_basis** | Option<**String**> | Where the precision claim comes from. Provenance for an audit, not a vocabulary to branch on. | [optional]
**is_period** | Option<**bool**> | True when the company did something across an interval (a trial). Independent of date_precision. | [optional]
**span_end_date** | Option<[**String**](string.md)> | The end of a real period; null for a non-period row and for an open-ended period. | [optional]
**span_end_is_projected** | Option<**bool**> | True when the end (or the event itself, for a decision still ahead) is planned rather than realized. | [optional]
**display_label** | Option<**String**> | Presentation label for the event, not a stable contract. | [optional]
**detail_text** | Option<**String**> | Presentation detail, including any amount as formatted text. | [optional]
**link_kind** | Option<**String**> | What link_ref points at: trial, funding_event, patent, external or none. | [optional]
**link_ref** | Option<**String**> | The natural id of the linked thing, or an absolute URL when link_kind is external; null exactly when link_kind is none. | [optional]
**restriction_tier** | Option<**String**> | The tier vocabulary is none or licensable. Every row here is none: every lane is an open source, and a data test fails if a restricted lane is ever admitted. | [optional]
**source_license** | Option<**String**> | The licence of the lane's source. | [optional]
**commercial_use_allowed** | Option<**bool**> | Whether the lane's licence permits commercial use. | [optional]
**attribution** | Option<**String**> | The credit line the lane's source requires; show it wherever the event is shown. Null when none is owed. | [optional]
**source_label** | Option<**String**> | The lane's display name. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


