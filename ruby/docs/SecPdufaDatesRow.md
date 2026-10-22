# PharmaObservatory::SecPdufaDatesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **pdufa_observation_id** | **String** | The extracted disclosure id (filing document id and disclosure index). | [optional] |
| **event_category** | **String** | pdufa, advisory_committee, priority_review or other. | [optional] |
| **status** | **String** | The disclosure&#39;s status: pending, approved, crl (complete response letter), or withdrawn. | [optional] |
| **is_decided** | **Boolean** | True when the status is approved, crl or withdrawn. | [optional] |
| **is_forward** | **Boolean** | True when the disclosure is a pending, forward-looking date. | [optional] |
| **application_type** | **String** | The application type the filing names (NDA, BLA and so on). | [optional] |
| **designation** | **String** | The review designation the filing names (priority review, fast track and so on). | [optional] |
| **indication** | **String** | The indication the filing names. | [optional] |
| **date_phrase** | **String** | The date as the filing words it. | [optional] |
| **date_start** | **Date** | Start of the date interval the phrase means. | [optional] |
| **date_end** | **Date** | End of the date interval the phrase means. | [optional] |
| **date_precision** | **String** | How exact the phrase is (day, month, quarter and so on). | [optional] |
| **is_in_period_now** | **Boolean** | True when today falls in the interval and the disclosure is not decided. | [optional] |
| **drug_name** | **String** | The drug name as filed. | [optional] |
| **sponsor_name** | **String** | The filer&#39;s name. | [optional] |
| **sponsor_org_id** | **String** | Legal-entity level: the SEC registrant&#39;s organization record (null when it is not in the spine). | [optional] |
| **sponsor_company_id** | **String** | Company level: the filer&#39;s company (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the filer&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the filer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the filer&#39;s group. | [optional] |
| **ticker** | **String** | The stock ticker the filing names. | [optional] |
| **source_url** | **String** | The filing&#39;s address on the SEC site. | [optional] |
| **confidence** | **Float** | The extraction&#39;s own confidence, from 0 to 1. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SecPdufaDatesRow.new(
  molecule_id: null,
  pdufa_observation_id: null,
  event_category: null,
  status: null,
  is_decided: null,
  is_forward: null,
  application_type: null,
  designation: null,
  indication: null,
  date_phrase: null,
  date_start: null,
  date_end: null,
  date_precision: null,
  is_in_period_now: null,
  drug_name: null,
  sponsor_name: null,
  sponsor_org_id: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  ticker: null,
  source_url: null,
  confidence: null
)
```

