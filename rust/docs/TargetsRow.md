# TargetsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional]
**symbol** | Option<**String**> | The target's HGNC gene symbol; the target page address accepts it. | [optional]
**name** | Option<**String**> | The gene's name. | [optional]
**biotype** | Option<**String**> | Open Targets biotype of the gene (protein_coding and so on); null where the source gives none. | [optional]
**uniprot_accession** | Option<**String**> | The UniProt accession of the gene's canonical protein; null where there is none. | [optional]
**protein_name** | Option<**String**> | The protein's recommended name from UniProt; null where the gene has no protein record. | [optional]
**hgnc_id** | Option<**String**> | The HGNC gene id (HGNC:nnnn) from UniProt; null where the gene has no protein record. | [optional]
**ncbi_gene_id** | Option<**String**> | The NCBI Gene id from UniProt; null where the gene has no protein record. | [optional]
**has_biomarker** | Option<**bool**> | Whether the gene has a biomarker record; the biomarkers datasets enter on the gene symbol. | [optional]
**n_drugs** | Option<**i64**> | Distinct drugs that act on the target, any evidence tier. | [optional]
**n_drugs_curated** | Option<**i64**> | Of n_drugs, those with a curated Open Targets mechanism edge (the floor). | [optional]
**n_drugs_investigational** | Option<**i64**> | Of n_drugs, those found only through ChEMBL or the curated investigational assets. | [optional]
**moa_floor_only** | Option<**bool**> | True when no drug beyond the curated floor is known for the target. | [optional]
**n_trials** | Option<**i64**> | Distinct trials whose drug interventions act on the target. | [optional]
**n_trials_curated** | Option<**i64**> | Of n_trials, those reached through a curated mechanism edge. | [optional]
**n_trials_investigational** | Option<**i64**> | Of n_trials, those reached only through ChEMBL or curated investigational assets. | [optional]
**n_programs** | Option<**f64**> | Distinct drug and indication programs against the target (a specific indication with at least two trials); the crowding headline. | [optional]
**n_active_programs** | Option<**f64**> | Of n_programs, those with no terminated, withdrawn or suspended trial. | [optional]
**lead_phase** | Option<**i32**> | The furthest phase band (0 to 4) that carries any program; 0 when the target has none. | [optional]
**programs_rank** | Option<**i64**> | Rank of n_programs among targets that have programs (1 is the most crowded); null when the target has none. | [optional]
**programs_pctile** | Option<**f64**> | Share of programmed targets with fewer programs than this one, 0 to 1; null when the target has none. | [optional]
**n_approved_drugs** | Option<**i64**> | Distinct approved drugs with a curated mechanism edge to the target. | [optional]
**n_publication_mention_papers** | Option<**i64**> | Distinct papers with a gene mention in the trial-cited PubMed corpus. | [optional]
**has_ot_association** | Option<**bool**> | True when any Open Targets datasource has a target-disease association for the gene. | [optional]
**n_pdb** | Option<**i32**> | Protein Data Bank structures mapped to the gene's protein. | [optional]
**has_pdb** | Option<**bool**> | True when at least one Protein Data Bank structure maps to the gene's protein. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


