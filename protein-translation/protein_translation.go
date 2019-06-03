package protein

import "errors"

// ErrStop a stop codon
var ErrStop = errors.New("Stop")

// ErrInvalidBase Is an eror if the base is invalid
var ErrInvalidBase = errors.New("InvalidBase")

// FromCodon stuff
func FromCodon(in string) (string, error) {
	switch in {
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

// FromRNA stuff
func FromRNA(in string) ([]string, error) {
	out := make([]string, 0)
	for i := 0; i < len(in); i += 3 {
		codon, err := FromCodon(in[i : i+3])
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			return out, ErrInvalidBase
		}
		out = append(out, codon)
	}
	return out, nil
}
