package protein

const testVersion = 1

var codonProteinMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UCC": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGC": "Cysteine",
	"UGU": "Cysteine",
	"UGG": "Tryptophan",
	"UAG": "STOP",
	"UAA": "STOP",
	"UGA": "STOP",
}

func FromCodon(rna string) string {
	return codonProteinMap[rna]
}

func FromRNA(rna string) []string {

	var codon = []rune{}
	var proteinList = []string{}

	for i, x := range rna {
		codon = append(codon, x)
		if (i+1)%3 == 0 {
			protein := codonProteinMap[string(codon)]
			if protein == "STOP" {
				break
			}

			proteinList = append(proteinList, protein)
			codon = []rune{}
		}
	}

	return proteinList
}
