# PharmaObservatory::DrugPatentCliffRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **drug_id** | **String** | The drug key. | [optional] |
| **drug_name** | **String** | The drug&#39;s name. | [optional] |
| **pharm_class** | **String** | The drug&#39;s pharmacologic class. | [optional] |
| **atc_class** | **String** | The drug&#39;s Anatomical Therapeutic Chemical (ATC) class. | [optional] |
| **brand** | **String** | The brand name. | [optional] |
| **candidate_brands** | **String** | Other candidate brand names, comma separated. | [optional] |
| **cliff_date** | **Date** | The date exclusivity is expected to end (loss of exclusivity); always in the future when the dataset is built. | [optional] |
| **cliff_year** | **Integer** | The year of cliff_date. | [optional] |
| **cliff_basis** | **String** | Which regime set the date: curated_loe, orange_book_patent, pte_expiry or bpcia_exclusivity. | [optional] |
| **cliff_source** | **String** | Source of a curated date (null on the other bases). | [optional] |
| **cliff_confidence** | **String** | Confidence of a curated date (null on the other bases). | [optional] |
| **has_generic** | **Boolean** | True when a generic is already marketed. | [optional] |
| **biosimilar_available** | **Boolean** | True when a biosimilar is already marketed. | [optional] |
| **applicant_name** | **String** | The brand or originator applicant as the Orange Book or Purple Book names it. | [optional] |
| **applicant_org_id** | **String** | Legal-entity level: the organization record the applicant resolved to (null when unresolved). | [optional] |
| **applicant_company_id** | **String** | Company level: the company the applicant belongs to (null when unresolved). | [optional] |
| **applicant_company_name** | **String** | Display name of the applicant&#39;s company. | [optional] |
| **applicant_group_id** | **String** | Group level: the applicant&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **applicant_group_name** | **String** | Display name of the applicant&#39;s group. | [optional] |
| **applicant_match_scope** | **String** | How the applicant name matched an organization record. | [optional] |
| **medicare_spend** | **Float** | Latest-year Medicare Part B and D spend for the drug, in US dollars. | [optional] |
| **medicare_spend_year** | **Integer** | The year of medicare_spend. | [optional] |
| **disclosed_revenue_usd_millions** | **Float** | Latest full-year disclosed worldwide revenue, in US millions. | [optional] |
| **disclosed_revenue_year** | **Integer** | The fiscal year of the disclosed revenue. | [optional] |
| **size_usd** | **Float** | The single size axis in US dollars: disclosed revenue when present, else Medicare spend. | [optional] |
| **size_basis** | **String** | Which source size_usd used: disclosed_worldwide_revenue, medicare_spend or none. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugPatentCliffRow.new(
  drug_id: null,
  drug_name: null,
  pharm_class: null,
  atc_class: null,
  brand: null,
  candidate_brands: null,
  cliff_date: null,
  cliff_year: null,
  cliff_basis: null,
  cliff_source: null,
  cliff_confidence: null,
  has_generic: null,
  biosimilar_available: null,
  applicant_name: null,
  applicant_org_id: null,
  applicant_company_id: null,
  applicant_company_name: null,
  applicant_group_id: null,
  applicant_group_name: null,
  applicant_match_scope: null,
  medicare_spend: null,
  medicare_spend_year: null,
  disclosed_revenue_usd_millions: null,
  disclosed_revenue_year: null,
  size_usd: null,
  size_basis: null
)
```

