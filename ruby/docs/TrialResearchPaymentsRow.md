# PharmaObservatory::TrialResearchPaymentsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **payment_key** | **String** | A stable key for the trial, payer and program year row. | [optional] |
| **payer_name** | **String** | The paying manufacturer&#39;s name as reported. | [optional] |
| **payer_org_id** | **String** | Legal-entity level: the payer rolled up to its ultimate parent organization record (null when unresolved). Payments from subsidiaries of one parent share one row per trial and program year. | [optional] |
| **payer_company_id** | **String** | Company level: the company the payer belongs to (null when unresolved). | [optional] |
| **payer_company_name** | **String** | The payer&#39;s company name. | [optional] |
| **payer_group_id** | **String** | Group level: the current ultimate parent of the payer&#39;s company. | [optional] |
| **payer_group_name** | **String** | The ultimate parent company&#39;s name. | [optional] |
| **program_year** | **Integer** | The Open Payments program year of the payments. | [optional] |
| **research_payments** | **Integer** | Payment records behind the row. | [optional] |
| **total_amount_usd** | **Float** | Total reported dollars for the row, counted once. | [optional] |
| **pi_count** | **Integer** | Distinct resolved principal investigators named on the payment records. | [optional] |
| **any_disputed** | **Boolean** | True when any payment record in the row was disputed by the recipient. | [optional] |
| **any_preclinical** | **Boolean** | True when any payment record in the row was marked as preclinical research. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialResearchPaymentsRow.new(
  trial_id: null,
  payment_key: null,
  payer_name: null,
  payer_org_id: null,
  payer_company_id: null,
  payer_company_name: null,
  payer_group_id: null,
  payer_group_name: null,
  program_year: null,
  research_payments: null,
  total_amount_usd: null,
  pi_count: null,
  any_disputed: null,
  any_preclinical: null
)
```

