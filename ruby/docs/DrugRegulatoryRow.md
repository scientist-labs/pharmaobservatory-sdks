# PharmaObservatory::DrugRegulatoryRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **fda_approved** | **Boolean** | True when the drug has a record in Drugs@FDA. | [optional] |
| **in_ndc** | **Boolean** | True when the drug has a package in the National Drug Code (NDC) directory. | [optional] |
| **has_generic** | **Boolean** | True when a generic is on record. | [optional] |
| **biosimilar_available** | **Boolean** | True when the Purple Book lists a biosimilar. | [optional] |
| **interchangeable_available** | **Boolean** | True when the Purple Book lists an interchangeable biosimilar. | [optional] |
| **first_licensure_date** | **Date** | The earliest biologic licensure date in the Purple Book. | [optional] |
| **pte_expiry** | **Date** | The extended expiry of a granted patent term extension. | [optional] |
| **patent_protected** | **Boolean** | True when an Orange Book patent is still active. | [optional] |
| **latest_patent_expiry** | **Date** | The latest expiry among the Orange Book patents. | [optional] |
| **applicant_name** | **String** | The applicant (company) as filed in the Orange or Purple Book. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugRegulatoryRow.new(
  molecule_id: null,
  fda_approved: null,
  in_ndc: null,
  has_generic: null,
  biosimilar_available: null,
  interchangeable_available: null,
  first_licensure_date: null,
  pte_expiry: null,
  patent_protected: null,
  latest_patent_expiry: null,
  applicant_name: null
)
```

