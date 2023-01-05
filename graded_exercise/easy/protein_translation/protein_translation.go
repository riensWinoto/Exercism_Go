package main

import (
	"errors"
)

var (
	ErrStop        error = errors.New("ErrStop")
	ErrInvalidBase error = errors.New("ErrInvalidBase")
)

func FromRNA(rna string) ([]string, error) {
	var translateProtein = []string{}
	for i := 0; i < len(rna); {
		codonx, err := FromCodon(rna[i : i+3])
		switch err {
		case ErrStop:
			return translateProtein, nil
		case ErrInvalidBase:
			return translateProtein, err
		}
		translateProtein = append(translateProtein, codonx)
		i += 3
	}
	return translateProtein, nil
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
