# PharmaObservatory::PdufaDatesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pdufa_date_id** | **String** | The calendar entry id. | [optional] |
| **source** | **String** | Which lane the row came from: benzinga, sec_extraction or both. | [optional] |
| **restriction_tier** | **String** | none for a row from SEC filings; licensable for a row from the Benzinga calendar. | [optional] |
| **decision_date** | **Date** | The FDA decision date (PDUFA is the date under the Prescription Drug User Fee Act). | [optional] |
| **date_start** | **Date** | Start of the date interval the source states. | [optional] |
| **date_end** | **Date** | End of the date interval the source states. | [optional] |
| **date_precision** | **String** | Precision of the date (day, month, quarter). | [optional] |
| **date_phrase** | **String** | The date as the source words it. | [optional] |
| **event_type** | **String** | Event type. | [optional] |
| **event_category** | **String** | Event category (pdufa, adcomm and so on). | [optional] |
| **drug_name** | **String** | The drug name as filed. | [optional] |
| **drug_id** | **String** | The drug key. | [optional] |
| **molecule_id** | **String** | The molecule key. | [optional] |
| **sponsor_name** | **String** | The sponsor name as filed. | [optional] |
| **sponsor_org_id** | **String** | Legal-entity level: the organization record the sponsor resolved to (null when unresolved). | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the sponsor belongs to (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **ticker** | **String** | Exchange ticker. | [optional] |
| **condition_id** | **String** | The condition key. | [optional] |
| **condition_name** | **String** | The condition&#39;s name. | [optional] |
| **indication_text** | **String** | The indication as the source words it. | [optional] |
| **source_type** | **String** | Kind of source document. | [optional] |
| **source_link** | **String** | Link to the source document. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::PdufaDatesRow.new(
  pdufa_date_id: null,
  source: null,
  restriction_tier: null,
  decision_date: null,
  date_start: null,
  date_end: null,
  date_precision: null,
  date_phrase: null,
  event_type: null,
  event_category: null,
  drug_name: null,
  drug_id: null,
  molecule_id: null,
  sponsor_name: null,
  sponsor_org_id: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  ticker: null,
  condition_id: null,
  condition_name: null,
  indication_text: null,
  source_type: null,
  source_link: null
)
```

