package rnatranscription

func ToRNA(dna string) string {
	rna := make([]byte, len(dna))

	for i, nucleotide := range dna {
		switch nucleotide {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		default:
			panic("invalid nucleotide in dna string")
		}
	}

	return string(rna)
}
