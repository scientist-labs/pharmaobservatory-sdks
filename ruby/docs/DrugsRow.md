# PharmaObservatory::DrugsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **molecule_name** | **String** | The drug&#39;s display name, from its lead record. | [optional] |
| **substance_name** | **String** | The preferred substance name in the FDA substance registry. | [optional] |
| **inn_name** | **String** | The International Nonproprietary Name (INN) where one is recorded. | [optional] |
| **unii** | **String** | The FDA Unique Ingredient Identifier (UNII) of the lead record. | [optional] |
| **cas** | **String** | The Chemical Abstracts Service (CAS) registry number of the lead record. | [optional] |
| **chembl_id** | **String** | The ChEMBL compound id of the lead record. | [optional] |
| **substance_type** | **String** | The registry&#39;s substance type (chemical, protein, mixture and so on). | [optional] |
| **atc_code** | **String** | The Anatomical Therapeutic Chemical (ATC) code. | [optional] |
| **atc_class** | **String** | The top-level ATC class. | [optional] |
| **pharm_class** | **String** | The pharmacologic class the FDA registry gives the substance. | [optional] |
| **is_biologic** | **Boolean** | True when the substance is a biologic, by the registry, Drugs@FDA or the Purple Book. | [optional] |
| **is_marketed** | **Boolean** | True when the drug has a package in the US National Drug Code (NDC) directory. | [optional] |
| **withdrawn** | **Boolean** | True when ChEMBL records the drug as withdrawn somewhere. | [optional] |
| **approved_anywhere** | **Boolean** | True when the drug is approved in at least one jurisdiction. | [optional] |
| **approved_in** | **Array&lt;String&gt;** | The regulators whose approval is recorded (FDA, EMA, Health Canada). | [optional] |
| **first_approval_year** | **Integer** | The earliest approval year ChEMBL records. | [optional] |
| **max_phase** | **Integer** | The highest development phase reached across the folded records (4 is approved). | [optional] |
| **n_trials** | **Integer** | Registered trials that study the drug or a record folded under it. | [optional] |
| **n_members** | **Integer** | Drug records folded into this drug, the lead included. | [optional] |
| **member_ids** | **Array&lt;String&gt;** | The drug record ids folded into this drug, the lead included. The ids other datasets name drug_id. | [optional] |
| **variant_count** | **Integer** | Variant records folded under the molecule; null when the drug is not a molecule. | [optional] |
| **brands** | **Array&lt;String&gt;** | The brand keys that resolve to the drug. | [optional] |
| **brand_count** | **Integer** | The number of brand keys. | [optional] |
| **is_molecule** | **Boolean** | True when the drug has a chemical-entities page. | [optional] |
| **is_combination** | **Boolean** | True when the drug is a combination product of two or more components. | [optional] |
| **has_boxed_warning** | **Boolean** | True when the drug&#39;s current label carries a boxed warning; null when no label is on record. | [optional] |
| **smiles** | **String** | The canonical SMILES string of the structure (ChEMBL). | [optional] |
| **inchikey** | **String** | The InChIKey of the structure (ChEMBL). | [optional] |
| **n_patent_mentions** | **Integer** | Patents that mention the drug in the SureChEMBL corpus. | [optional] |
| **n_patent_families** | **Integer** | Distinct patent families among those mentions. | [optional] |
| **first_patent_mention** | **Date** | Publication date of the earliest mentioning patent. | [optional] |
| **last_patent_mention** | **Date** | Publication date of the latest mentioning patent. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugsRow.new(
  molecule_id: null,
  molecule_name: null,
  substance_name: null,
  inn_name: null,
  unii: null,
  cas: null,
  chembl_id: null,
  substance_type: null,
  atc_code: null,
  atc_class: null,
  pharm_class: null,
  is_biologic: null,
  is_marketed: null,
  withdrawn: null,
  approved_anywhere: null,
  approved_in: null,
  first_approval_year: null,
  max_phase: null,
  n_trials: null,
  n_members: null,
  member_ids: null,
  variant_count: null,
  brands: null,
  brand_count: null,
  is_molecule: null,
  is_combination: null,
  has_boxed_warning: null,
  smiles: null,
  inchikey: null,
  n_patent_mentions: null,
  n_patent_families: null,
  first_patent_mention: null,
  last_patent_mention: null
)
```

