# PharmaObservatory::CompanyFundingEventsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **funding_event_id** | **String** | The funding event key. | [optional] |
| **issuer_name** | **String** | The issuer&#39;s name as filed. | [optional] |
| **issuer_org_id** | **String** | Legal-entity level: the issuer&#39;s organization record. | [optional] |
| **issuer_company_id** | **String** | Company level: the issuer&#39;s company (null when the issuer did not resolve to a company). | [optional] |
| **issuer_company_name** | **String** | Display name of the issuer&#39;s company. | [optional] |
| **issuer_group_id** | **String** | Group level: the issuer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **issuer_group_name** | **String** | Display name of the issuer&#39;s group. | [optional] |
| **event_date** | **Date** | The date the source gives for the event; event_stage says what the date means. | [optional] |
| **fiscal_year** | **Integer** | The fiscal year of the US National Institutes of Health (NIH), for a grant; null for the other sources. | [optional] |
| **source** | **String** | The source lane: reporter (NIH RePORTER grants), form_d (SEC Form D Regulation D offerings), edgar_offerings (SEC registered offerings) or sec_extracted (financings extracted from SEC filings). | [optional] |
| **round_type** | **String** | The kind of round (for example grant_nih, sbir, reg_d_private, ipo, follow_on, public_offering, atm, registered_direct, debt). | [optional] |
| **round_subtype** | **String** | A finer type, such as the grant activity code, the Regulation D exemption or the SEC form. | [optional] |
| **is_dilutive** | **Boolean** | Whether the round dilutes existing equity; null where the instrument is not classified. | [optional] |
| **amount_usd** | **Float** | The disclosed amount in whole US dollars; null where the amount is not disclosed. | [optional] |
| **currency** | **String** | The currency the source reports the amount in. | [optional] |
| **amount_is_disclosed** | **Boolean** | Whether the source states an amount for the event. | [optional] |
| **capital_treatment** | **String** | Whether the event&#39;s dollars count as capital raised, and if not why not (for example structured_source, actual_inflow, capacity, committed). | [optional] |
| **is_actual_issuer_inflow** | **Boolean** | True where the event counts as capital the issuer raised. | [optional] |
| **event_stage** | **String** | The lifecycle stage the event&#39;s date refers to (for example sold_to_date, funded, closed, registered, announced). | [optional] |
| **administering_ic** | **String** | The National Institutes of Health institute or center that administers the grant (for example CA for the National Cancer Institute); null for the other sources. | [optional] |
| **n_awards** | **Integer** | Grant applications folded into the event; 1 for the other sources. | [optional] |
| **cumulative_raised_usd** | **Float** | The running total of disclosed amounts over the company&#39;s own events up to and including this one, in date order; null where the event has no company. | [optional] |
| **source_url** | **String** | Link to the filing the event was read from, where the source has one. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyFundingEventsRow.new(
  funding_event_id: null,
  issuer_name: null,
  issuer_org_id: null,
  issuer_company_id: null,
  issuer_company_name: null,
  issuer_group_id: null,
  issuer_group_name: null,
  event_date: null,
  fiscal_year: null,
  source: null,
  round_type: null,
  round_subtype: null,
  is_dilutive: null,
  amount_usd: null,
  currency: null,
  amount_is_disclosed: null,
  capital_treatment: null,
  is_actual_issuer_inflow: null,
  event_stage: null,
  administering_ic: null,
  n_awards: null,
  cumulative_raised_usd: null,
  source_url: null
)
```

