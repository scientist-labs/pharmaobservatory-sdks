# DrugsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**molecule_name** | Option<**String**> | The drug's display name, from its lead record. | [optional]
**substance_name** | Option<**String**> | The preferred substance name in the FDA substance registry. | [optional]
**inn_name** | Option<**String**> | The International Nonproprietary Name (INN) where one is recorded. | [optional]
**unii** | Option<**String**> | The FDA Unique Ingredient Identifier (UNII) of the lead record. | [optional]
**cas** | Option<**String**> | The Chemical Abstracts Service (CAS) registry number of the lead record. | [optional]
**chembl_id** | Option<**String**> | The ChEMBL compound id of the lead record. | [optional]
**substance_type** | Option<**String**> | The registry's substance type (chemical, protein, mixture and so on). | [optional]
**atc_code** | Option<**String**> | The Anatomical Therapeutic Chemical (ATC) code. | [optional]
**atc_class** | Option<**String**> | The top-level ATC class. | [optional]
**pharm_class** | Option<**String**> | The pharmacologic class the FDA registry gives the substance. | [optional]
**is_biologic** | Option<**bool**> | True when the substance is a biologic, by the registry, Drugs@FDA or the Purple Book. | [optional]
**is_marketed** | Option<**bool**> | True when the drug has a package in the US National Drug Code (NDC) directory. | [optional]
**withdrawn** | Option<**bool**> | True when ChEMBL records the drug as withdrawn somewhere. | [optional]
**approved_anywhere** | Option<**bool**> | True when the drug is approved in at least one jurisdiction. | [optional]
**approved_in** | Option<**Vec<String>**> | The regulators whose approval is recorded (FDA, EMA, Health Canada). | [optional]
**first_approval_year** | Option<**i32**> | The earliest approval year ChEMBL records. | [optional]
**max_phase** | Option<**i32**> | The highest development phase reached across the folded records (4 is approved). | [optional]
**n_trials** | Option<**i32**> | Registered trials that study the drug or a record folded under it. | [optional]
**n_members** | Option<**i32**> | Drug records folded into this drug, the lead included. | [optional]
**member_ids** | Option<**Vec<String>**> | The drug record ids folded into this drug, the lead included. The ids other datasets name drug_id. | [optional]
**variant_count** | Option<**i32**> | Variant records folded under the molecule; null when the drug is not a molecule. | [optional]
**brands** | Option<**Vec<String>**> | The brand keys that resolve to the drug. | [optional]
**brand_count** | Option<**i32**> | The number of brand keys. | [optional]
**is_molecule** | Option<**bool**> | True when the drug has a chemical-entities page. | [optional]
**is_combination** | Option<**bool**> | True when the drug is a combination product of two or more components. | [optional]
**has_boxed_warning** | Option<**bool**> | True when the drug's current label carries a boxed warning; null when no label is on record. | [optional]
**smiles** | Option<**String**> | The canonical SMILES string of the structure (ChEMBL). | [optional]
**inchikey** | Option<**String**> | The InChIKey of the structure (ChEMBL). | [optional]
**n_patent_mentions** | Option<**i32**> | Patents that mention the drug in the SureChEMBL corpus. | [optional]
**n_patent_families** | Option<**i32**> | Distinct patent families among those mentions. | [optional]
**first_patent_mention** | Option<[**String**](string.md)> | Publication date of the earliest mentioning patent. | [optional]
**last_patent_mention** | Option<[**String**](string.md)> | Publication date of the latest mentioning patent. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


