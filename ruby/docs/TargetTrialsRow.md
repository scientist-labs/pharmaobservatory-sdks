# PharmaObservatory::TargetTrialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] |
| **target_symbol** | **String** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] |
| **trial_id** | **String** | The ClinicalTrials.gov NCT id, the id the trials datasets use. | [optional] |
| **evidence_source** | **String** | The best evidence tier among the trial&#39;s drugs that act on the target. | [optional] |
| **brief_title** | **String** | The trial&#39;s brief title. | [optional] |
| **phase** | **String** | The registry&#39;s phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional] |
| **phase_label** | **String** | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial&#39;s phase column holds the registry code. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall recruitment status. | [optional] |
| **start_year** | **Integer** | The year the trial started. | [optional] |
| **lead_sponsor_name** | **String** | The lead sponsor&#39;s name as filed. | [optional] |
| **lead_sponsor_org_id** | **String** | Legal-entity level: the lead sponsor organization; null when it does not resolve. | [optional] |
| **lead_sponsor_company_id** | **String** | Company level: the lead sponsor&#39;s company; null when none. | [optional] |
| **lead_sponsor_company_name** | **String** | Display name of the lead sponsor&#39;s company. | [optional] |
| **lead_sponsor_group_id** | **String** | Group level: the lead sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **lead_sponsor_group_name** | **String** | Display name of the lead sponsor&#39;s group. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetTrialsRow.new(
  target_id: null,
  target_symbol: null,
  trial_id: null,
  evidence_source: null,
  brief_title: null,
  phase: null,
  phase_label: null,
  overall_status: null,
  start_year: null,
  lead_sponsor_name: null,
  lead_sponsor_org_id: null,
  lead_sponsor_company_id: null,
  lead_sponsor_company_name: null,
  lead_sponsor_group_id: null,
  lead_sponsor_group_name: null
)
```

