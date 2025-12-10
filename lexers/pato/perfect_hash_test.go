package pato

import (
	"testing"
)

// TestFindPerfectHash searches for a perfect hash function for the keywords.
// This test is normally skipped unless -run=TestFindPerfectHash is specified.
// It found the current perfect hash: c0=4 c1=26 c2=17 c3=11 cLast=28 cLen=29
func TestFindPerfectHash(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping hash search in short mode")
	}

	// Collect all keywords
	var keywords []string
	var kwToks []Token
	for tok := keywordBeg + 1; tok < keywordEnd; tok++ {
		keywords = append(keywords, tok.String())
		kwToks = append(kwToks, tok)
	}
	t.Logf("Searching perfect hash for %d keywords", len(keywords))
	attempts := 0
	tableSizes := []uint{1 << 4, 1 << 5, 1 << 6, 1 << 7, 1 << 8, 1 << 9, 1 << 10}
	const maxCoef = 32

	// Coefficients for perfect hash function.
	// First coeficcient is length multiplication.
	// This method of searching for a perfect hash is quite robust
	// and works even with Fortran keywords (70+ keywords, ~20 sharing identical END start)
	var coefs [3]uint
	for _, tableSize := range tableSizes {
		kwMap := make([]Token, tableSize)
		t.Logf("Trying table size %d...", tableSize)
		for i := range coefs {
			coefs[i] = 1 // start all coefficients at 1.
		}
		attempt := 0
		for {
			attempt++
			mask := tableSize - 1
			collision := false
			clear(kwMap)
			for i, kw := range keywords {
				h := uint(len(kw)) * coefs[0]
				for i := 1; i < len(coefs) && i < len(kw); i++ {
					h += uint(kw[i-1]) * coefs[i]
				}
				h &= mask
				tok := kwMap[h]
				if tok != 0 {
					collision = true
					break
				}
				kwMap[h] = kwToks[i]
			}
			attempts++
			if !collision {
				t.Logf("FOUND with tableSize=%d after %d attempts! (%d total attempts)", tableSize, attempt+1, attempts)
				t.Logf("%v", coefs)
				return
			}
			for i := 0; coefs[i] == maxCoef; i++ {

			}
			coefs[0]++
			for i := 0; coefs[i] == maxCoef && i < len(coefs)-1; i++ {
				coefs[i] = 1
				coefs[i+1]++
			}
			if coefs[len(coefs)-1] == maxCoef+1 {
				break
			}
		}
	}
	t.Error("No perfect hash found after", attempts, "attempts")
}
