package proteintranslation

import (
	"errors"
)

var (
	ErrStop        = errors.New("STOP codon")
	ErrInvalidBase = errors.New("invalid base")
)

func FromRNA(rna string) ([]string, error) {
	var codons []string

	for i := 0; i < len(rna); i += 3 {
		codon, err := FromCodon(rna[i:min(i+3, len(rna))])

		switch {
		case errors.Is(err, ErrStop):
			return codons, nil
		case errors.Is(err, ErrInvalidBase):
			return []string{}, ErrInvalidBase
		default:
			codons = append(codons, codon)
		}
	}

	return codons, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
