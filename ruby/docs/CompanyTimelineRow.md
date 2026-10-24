# PharmaObservatory::CompanyTimelineRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **timeline_event_id** | **String** | md5 of the lane, event type, source record and company: a stable surrogate id for the row. | [optional] |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **reporting_org_id** | **String** | Legal-entity level: the organization record the lane reported the event on (the lowest id when several records of the company report it). | [optional] |
| **source_key** | **String** | The lane that supplied the event. | [optional] |
| **source_entity_key** | **String** | The lane&#39;s own key for the thing that happened (trial id, funding event id, patent id, enforcement event id, statement id). | [optional] |
| **event_type** | **String** | The fine event vocabulary; funding types are the funding round types verbatim. | [optional] |
| **event_family** | **String** | The family of the event type, assigned by name from the lane registry: clinical, capital, intellectual_property, regulatory, compliance or other. | [optional] |
| **event_date** | **Date** | The first day consistent with the evidence, never the raw value when the evidence is coarser than a day. | [optional] |
| **date_precision** | **String** | How precise the evidence is: day, month, quarter, half or year. | [optional] |
| **precision_window_end** | **Date** | The last day consistent with the evidence; equals event_date at day precision. | [optional] |
| **date_precision_basis** | **String** | Where the precision claim comes from. Provenance for an audit, not a vocabulary to branch on. | [optional] |
| **is_period** | **Boolean** | True when the company did something across an interval (a trial). Independent of date_precision. | [optional] |
| **span_end_date** | **Date** | The end of a real period; null for a non-period row and for an open-ended period. | [optional] |
| **span_end_is_projected** | **Boolean** | True when the end (or the event itself, for a decision still ahead) is planned rather than realized. | [optional] |
| **display_label** | **String** | Presentation label for the event, not a stable contract. | [optional] |
| **detail_text** | **String** | Presentation detail, including any amount as formatted text. | [optional] |
| **link_kind** | **String** | What link_ref points at: trial, funding_event, patent, external or none. | [optional] |
| **link_ref** | **String** | The natural id of the linked thing, or an absolute URL when link_kind is external; null exactly when link_kind is none. | [optional] |
| **restriction_tier** | **String** | The tier vocabulary is none or licensable. Every row here is none: every lane is an open source, and a data test fails if a restricted lane is ever admitted. | [optional] |
| **source_license** | **String** | The licence of the lane&#39;s source. | [optional] |
| **commercial_use_allowed** | **Boolean** | Whether the lane&#39;s licence permits commercial use. | [optional] |
| **attribution** | **String** | The credit line the lane&#39;s source requires; show it wherever the event is shown. Null when none is owed. | [optional] |
| **source_label** | **String** | The lane&#39;s display name. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyTimelineRow.new(
  timeline_event_id: null,
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  reporting_org_id: null,
  source_key: null,
  source_entity_key: null,
  event_type: null,
  event_family: null,
  event_date: null,
  date_precision: null,
  precision_window_end: null,
  date_precision_basis: null,
  is_period: null,
  span_end_date: null,
  span_end_is_projected: null,
  display_label: null,
  detail_text: null,
  link_kind: null,
  link_ref: null,
  restriction_tier: null,
  source_license: null,
  commercial_use_allowed: null,
  attribution: null,
  source_label: null
)
```

