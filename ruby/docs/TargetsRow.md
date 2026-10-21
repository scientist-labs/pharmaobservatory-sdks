# PharmaObservatory::TargetsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] |
| **symbol** | **String** | The target&#39;s HGNC gene symbol; the target page address accepts it. | [optional] |
| **name** | **String** | The gene&#39;s name. | [optional] |
| **biotype** | **String** | Open Targets biotype of the gene (protein_coding and so on); null where the source gives none. | [optional] |
| **uniprot_accession** | **String** | The UniProt accession of the gene&#39;s canonical protein; null where there is none. | [optional] |
| **protein_name** | **String** | The protein&#39;s recommended name from UniProt; null where the gene has no protein record. | [optional] |
| **hgnc_id** | **String** | The HGNC gene id (HGNC:nnnn) from UniProt; null where the gene has no protein record. | [optional] |
| **ncbi_gene_id** | **String** | The NCBI Gene id from UniProt; null where the gene has no protein record. | [optional] |
| **has_biomarker** | **Boolean** | Whether the gene has a biomarker record; the biomarkers datasets enter on the gene symbol. | [optional] |
| **n_drugs** | **Integer** | Distinct drugs that act on the target, any evidence tier. | [optional] |
| **n_drugs_curated** | **Integer** | Of n_drugs, those with a curated Open Targets mechanism edge (the floor). | [optional] |
| **n_drugs_investigational** | **Integer** | Of n_drugs, those found only through ChEMBL or the curated investigational assets. | [optional] |
| **moa_floor_only** | **Boolean** | True when no drug beyond the curated floor is known for the target. | [optional] |
| **n_trials** | **Integer** | Distinct trials whose drug interventions act on the target. | [optional] |
| **n_trials_curated** | **Integer** | Of n_trials, those reached through a curated mechanism edge. | [optional] |
| **n_trials_investigational** | **Integer** | Of n_trials, those reached only through ChEMBL or curated investigational assets. | [optional] |
| **n_programs** | **Float** | Distinct drug and indication programs against the target (a specific indication with at least two trials); the crowding headline. | [optional] |
| **n_active_programs** | **Float** | Of n_programs, those with no terminated, withdrawn or suspended trial. | [optional] |
| **lead_phase** | **Integer** | The furthest phase band (0 to 4) that carries any program; 0 when the target has none. | [optional] |
| **programs_rank** | **Integer** | Rank of n_programs among targets that have programs (1 is the most crowded); null when the target has none. | [optional] |
| **programs_pctile** | **Float** | Share of programmed targets with fewer programs than this one, 0 to 1; null when the target has none. | [optional] |
| **n_approved_drugs** | **Integer** | Distinct approved drugs with a curated mechanism edge to the target. | [optional] |
| **n_publication_mention_papers** | **Integer** | Distinct papers with a gene mention in the trial-cited PubMed corpus. | [optional] |
| **has_ot_association** | **Boolean** | True when any Open Targets datasource has a target-disease association for the gene. | [optional] |
| **n_pdb** | **Integer** | Protein Data Bank structures mapped to the gene&#39;s protein. | [optional] |
| **has_pdb** | **Boolean** | True when at least one Protein Data Bank structure maps to the gene&#39;s protein. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetsRow.new(
  target_id: null,
  symbol: null,
  name: null,
  biotype: null,
  uniprot_accession: null,
  protein_name: null,
  hgnc_id: null,
  ncbi_gene_id: null,
  has_biomarker: null,
  n_drugs: null,
  n_drugs_curated: null,
  n_drugs_investigational: null,
  moa_floor_only: null,
  n_trials: null,
  n_trials_curated: null,
  n_trials_investigational: null,
  n_programs: null,
  n_active_programs: null,
  lead_phase: null,
  programs_rank: null,
  programs_pctile: null,
  n_approved_drugs: null,
  n_publication_mention_papers: null,
  has_ot_association: null,
  n_pdb: null,
  has_pdb: null
)
```

