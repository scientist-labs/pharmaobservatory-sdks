# BiomarkerAlleleFrequenciesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**route_key** | Option<**String**> | The route key of the gene the variant belongs to, the entry key for the biomarkers object type. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol. | [optional]
**variant_route_key** | Option<**String**> | The variant's own route key (the id its web page uses). | [optional]
**variant_biomarker_id** | Option<**String**> | The variant's own biomarker key. | [optional]
**amino_acid_change** | Option<**String**> | The protein change, for example p.V600E. | [optional]
**biomarker_variant_coordinate_id** | Option<**String**> | The key of the DNA change (the coordinate path) the row describes. | [optional]
**snv_ordinal** | Option<**i32**> | Which single-base DNA change of the protein change this is, from one; most changes have one. | [optional]
**is_ambiguous_path** | Option<**bool**> | Whether the protein change can be produced by more than one single-base DNA change, so this row is one candidate rather than the only one. | [optional]
**snv_path_count** | Option<**i32**> | How many single-base DNA changes produce the protein change. | [optional]
**refseq_transcript_accession** | Option<**String**> | The reference transcript the protein change was mapped on (a RefSeq accession). | [optional]
**strand** | Option<**String**> | The strand of the gene on the genome, plus or minus. | [optional]
**queried_contig** | Option<**String**> | The chromosome as gnomAD spells it. | [optional]
**queried_position** | Option<**i64**> | The genomic position looked up in gnomAD (GRCh38). | [optional]
**reference_allele** | Option<**String**> | The reference base on the genomic plus strand. | [optional]
**alternate_allele** | Option<**String**> | The alternate base on the genomic plus strand. | [optional]
**gnomad_release** | Option<**String**> | The gnomAD release the lookup used. | [optional]
**coverage_status** | Option<**String**> | What the lookup found. observed: the DNA change is in gnomAD. absent_from_gnomad: the position was queried and the change is not there. no_gnomad_release_for_contig: gnomAD has no file for the chromosome. query_failed: the lookup failed. | [optional]
**filter_status** | Option<**String**> | gnomAD's own quality filter for the site, verbatim (PASS or a semicolon list). Null unless observed. | [optional]
**population_code** | Option<**String**> | The population. overall is the unstratified total; the others are gnomAD genetic ancestry groups: afr African/African American, amr Admixed American, asj Ashkenazi Jewish, eas East Asian, fin Finnish, mid Middle Eastern, nfe Non-Finnish European, remaining the remaining individuals, sas South Asian. | [optional]
**allele_count** | Option<**i64**> | Alternate alleles observed in the population. Null unless observed. | [optional]
**allele_number** | Option<**i64**> | Alleles called in the population, the denominator of allele_frequency. Null unless observed. | [optional]
**allele_frequency** | Option<**f64**> | The share of called alleles that are the alternate allele (allele_count over allele_number), as a fraction from zero to one. Null unless observed. | [optional]
**homozygote_count** | Option<**i64**> | People homozygous for the alternate allele. Null unless observed. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


