# PharmaObservatory::CatalystsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **event_id** | **String** | The event key: event_type, a colon and source_record_id. Unique. | [optional] |
| **event_type** | **String** | What the event is: readout (a trial&#39;s primary completion), patent_cliff (loss of exclusivity), pdufa (an FDA decision date or advisory committee) or press_release (a Benzinga press-release event). | [optional] |
| **event_category** | **String** | The kind of event within its type. readout: primary_completion. patent_cliff: the cliff basis (curated_loe, orange_book_patent, pte_expiry, bpcia_exclusivity). pdufa: the calendar&#39;s own category (pdufa, advisory_committee and others). press_release: the Benzinga event token, lowercased (clinical_trial_result, pdufa_assignment, data_presentation and others). Open vocabulary on the last two. | [optional] |
| **source** | **String** | The lane the row came from: ctgov (ClinicalTrials.gov), drug_exclusivity (the exclusivity registers), benzinga, both or sec_extraction (the PDUFA calendar, the same values as pdufa_dates) or benzinga_press_release. | [optional] |
| **source_record_id** | **String** | The key of the row in its source, to join back to trials.trial_id, drug_patent_cliff.drug_id, pdufa_dates.pdufa_date_id or the press-release event id. | [optional] |
| **restriction_tier** | **String** | none for an open row (ClinicalTrials.gov, the exclusivity registers, the SEC lane of the PDUFA calendar); licensable for a Benzinga row (the calendar or a press release). The api view gates rows on it. | [optional] |
| **event_date** | **Date** | The date the calendar sorts on: a readout&#39;s primary completion date, the cliff date, the decision date, or the end of the press-release event&#39;s date window. | [optional] |
| **event_year** | **Integer** | The calendar year of event_date; the x axis of a by-year count. | [optional] |
| **date_start** | **Date** | Start of the date interval the source states (the first day of the month for a readout the registry filed by month). | [optional] |
| **date_end** | **Date** | End of the date interval the source states. | [optional] |
| **date_precision** | **String** | Precision of the date (day, month, quarter, half, year, fuzzy, relative). A readout is month where the registry filed only a month, in which case event_date is the last day of that month. | [optional] |
| **date_phrase** | **String** | The date as the source words it (an ISO date for a readout and a cliff; the press release&#39;s own phrase, such as H2 2026). | [optional] |
| **date_status** | **String** | For a readout, whether the registry&#39;s completion date is estimated or actual (an anticipated date counts as estimated); null on the other legs. | [optional] |
| **trial_id** | **String** | The ClinicalTrials.gov registry id (readout rows). | [optional] |
| **trial_title** | **String** | The trial&#39;s brief title (readout rows). | [optional] |
| **phase** | **String** | The registry&#39;s phase code (readout rows). | [optional] |
| **phase_label** | **String** | The trial phase as a display label (readout rows). | [optional] |
| **overall_status** | **String** | The registry&#39;s overall status code (readout rows). | [optional] |
| **related_trial_ids** | **Array&lt;String&gt;** | The registry ids a press-release event names; null on the other legs. | [optional] |
| **outcome** | **String** | The outcome a press-release event reports (a decided event stays in the forward book); null on the other legs. | [optional] |
| **indication_text** | **String** | The indication as the source words it (PDUFA and press-release rows). | [optional] |
| **condition_id** | **String** | The condition key (PDUFA rows). | [optional] |
| **condition_name** | **String** | The condition&#39;s name (PDUFA rows). | [optional] |
| **source_type** | **String** | Kind of source document (SEC filing for the SEC lane, Press release for a press-release event); null on the open legs. | [optional] |
| **source_url** | **String** | Link to the source document (PDUFA and press-release rows). | [optional] |
| **drug_name** | **String** | The drug name as filed or its display name; null on a readout (its drugs are the list drug_ids). | [optional] |
| **drug_id** | **String** | The drug key of a patent-cliff, PDUFA or press-release row; null on a readout and where the drug did not resolve. | [optional] |
| **molecule_id** | **String** | The molecule key (PDUFA and press-release rows). | [optional] |
| **drug_ids** | **Array&lt;String&gt;** | Every drug key of the event: a readout carries each drug of its trial, sorted; the other legs carry their one drug, or an empty array. The drug predicate of a watchlist is &#39;&lt;drug&gt;&#39; &#x3D; ANY(drug_ids). | [optional] |
| **sponsor_role** | **String** | What the organization is on the row: lead_sponsor (readout), applicant (patent cliff, the FDA applicant of record), sponsor (PDUFA) or issuer (press release, the company the ticker names). | [optional] |
| **sponsor_name** | **String** | The organization&#39;s name as filed; null on a press-release row (the source names a ticker, not an organization). | [optional] |
| **sponsor_org_id** | **String** | Legal-entity level: the organization record (a readout&#39;s lead sponsor, a cliff&#39;s applicant, a PDUFA row&#39;s filing organization; for a press release the identity hub&#39;s resolved entity for the ticker). Null when unresolved. | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the organization belongs to (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the organization&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the organization&#39;s group. | [optional] |
| **ticker** | **String** | Exchange ticker (PDUFA and press-release rows). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CatalystsRow.new(
  event_id: null,
  event_type: null,
  event_category: null,
  source: null,
  source_record_id: null,
  restriction_tier: null,
  event_date: null,
  event_year: null,
  date_start: null,
  date_end: null,
  date_precision: null,
  date_phrase: null,
  date_status: null,
  trial_id: null,
  trial_title: null,
  phase: null,
  phase_label: null,
  overall_status: null,
  related_trial_ids: null,
  outcome: null,
  indication_text: null,
  condition_id: null,
  condition_name: null,
  source_type: null,
  source_url: null,
  drug_name: null,
  drug_id: null,
  molecule_id: null,
  drug_ids: null,
  sponsor_role: null,
  sponsor_name: null,
  sponsor_org_id: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  ticker: null
)
```

