# PharmaObservatory::BiomarkerAlleleFrequenciesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The route key of the gene the variant belongs to, the entry key for the biomarkers object type. | [optional] |
| **gene_symbol** | **String** | The gene symbol. | [optional] |
| **variant_route_key** | **String** | The variant&#39;s own route key (the id its web page uses). | [optional] |
| **variant_biomarker_id** | **String** | The variant&#39;s own biomarker key. | [optional] |
| **amino_acid_change** | **String** | The protein change, for example p.V600E. | [optional] |
| **biomarker_variant_coordinate_id** | **String** | The key of the DNA change (the coordinate path) the row describes. | [optional] |
| **snv_ordinal** | **Integer** | Which single-base DNA change of the protein change this is, from one; most changes have one. | [optional] |
| **is_ambiguous_path** | **Boolean** | Whether the protein change can be produced by more than one single-base DNA change, so this row is one candidate rather than the only one. | [optional] |
| **snv_path_count** | **Integer** | How many single-base DNA changes produce the protein change. | [optional] |
| **refseq_transcript_accession** | **String** | The reference transcript the protein change was mapped on (a RefSeq accession). | [optional] |
| **strand** | **String** | The strand of the gene on the genome, plus or minus. | [optional] |
| **queried_contig** | **String** | The chromosome as gnomAD spells it. | [optional] |
| **queried_position** | **Integer** | The genomic position looked up in gnomAD (GRCh38). | [optional] |
| **reference_allele** | **String** | The reference base on the genomic plus strand. | [optional] |
| **alternate_allele** | **String** | The alternate base on the genomic plus strand. | [optional] |
| **gnomad_release** | **String** | The gnomAD release the lookup used. | [optional] |
| **coverage_status** | **String** | What the lookup found. observed: the DNA change is in gnomAD. absent_from_gnomad: the position was queried and the change is not there. no_gnomad_release_for_contig: gnomAD has no file for the chromosome. query_failed: the lookup failed. | [optional] |
| **filter_status** | **String** | gnomAD&#39;s own quality filter for the site, verbatim (PASS or a semicolon list). Null unless observed. | [optional] |
| **population_code** | **String** | The population. overall is the unstratified total; the others are gnomAD genetic ancestry groups: afr African/African American, amr Admixed American, asj Ashkenazi Jewish, eas East Asian, fin Finnish, mid Middle Eastern, nfe Non-Finnish European, remaining the remaining individuals, sas South Asian. | [optional] |
| **allele_count** | **Integer** | Alternate alleles observed in the population. Null unless observed. | [optional] |
| **allele_number** | **Integer** | Alleles called in the population, the denominator of allele_frequency. Null unless observed. | [optional] |
| **allele_frequency** | **Float** | The share of called alleles that are the alternate allele (allele_count over allele_number), as a fraction from zero to one. Null unless observed. | [optional] |
| **homozygote_count** | **Integer** | People homozygous for the alternate allele. Null unless observed. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerAlleleFrequenciesRow.new(
  route_key: null,
  gene_symbol: null,
  variant_route_key: null,
  variant_biomarker_id: null,
  amino_acid_change: null,
  biomarker_variant_coordinate_id: null,
  snv_ordinal: null,
  is_ambiguous_path: null,
  snv_path_count: null,
  refseq_transcript_accession: null,
  strand: null,
  queried_contig: null,
  queried_position: null,
  reference_allele: null,
  alternate_allele: null,
  gnomad_release: null,
  coverage_status: null,
  filter_status: null,
  population_code: null,
  allele_count: null,
  allele_number: null,
  allele_frequency: null,
  homozygote_count: null
)
```

