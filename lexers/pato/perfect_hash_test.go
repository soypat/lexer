package pato

import (
	"errors"
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
	// Coefficients for perfect hash function.
	// First coeficcient is length multiplication.
	// This method of searching for a perfect hash is quite robust
	// and works even with Fortran keywords (70+ keywords, ~20 sharing identical END start)
	const maxCoef = 32
	const onlyPow2Coefs = true
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
			if onlyPow2Coefs {
				coefs[0] *= 2
			} else {
				coefs[0]++
			}

			for i := 0; coefs[i] == maxCoef && i < len(coefs)-1; i++ {
				coefs[i] = 1
				if onlyPow2Coefs {
					coefs[i+1] *= 2
				} else {
					coefs[i+1]++
				}
			}
			if coefs[len(coefs)-1] == maxCoef+1 {
				break
			}
		}
	}
	t.Error("No perfect hash found after", attempts, "attempts")
}

type PerfectHashFinder struct {
	TableSize      int
	DefaultMaxCoef uint
	// HashLastIndices
	hashmap []uint
	mask    uint
}
type Coef struct {
	IndexApplied int  // Index at which hash consumes byte.
	Value        uint // Coefficient value to multiply byte at index.
	MaxValue     uint
	StartValue   uint
	OnlyPow2     bool
}

func (c *Coef) init() {
	if c.StartValue == 0 {
		c.Value = 1
	} else {
		c.Value = c.StartValue
	}
}
func (c *Coef) increment() {
	if c.OnlyPow2 {
		c.Value *= 2
	} else {
		c.Value++
	}
}
func (c *Coef) saturated() bool { return c.Value >= c.MaxValue }

func (phf *PerfectHashFinder) Search(coefs []Coef, inputs []string) error {
	tblsz := phf.TableSize
	if tblsz == 0 {
		return errors.New("zero table size")
	}
	hashmap := make([]uint, tblsz)
	for i := range coefs {
		coefs[i].init()
		if coefs[i].MaxValue == 0 {
			coefs[i].MaxValue = phf.DefaultMaxCoef
		}
	}
	mask := uint(tblsz) - 1
	currentAttempt := 0
	for {
		currentAttempt++
		attemptSuccess := true
		clear(hashmap)
		allSat := true
		for i := range coefs {
			allSat = allSat && coefs[i].saturated()
		}
		if allSat {
			break // We are done.
		}
		for _, kw := range inputs {
			h := phf.apply(mask, coefs, kw)
			tok := hashmap[h]
			if tok != 0 {
				attemptSuccess = false
				break
			}
			hashmap[h] = 1
		}
		if attemptSuccess {
			return nil
		}
		coefs[0].increment()
		for i := 0; coefs[i].saturated() && i < len(coefs)-1; i++ {
			coefs[i].increment()
		}
	}
	return errors.New("no coefficients found")
}

func (phf *PerfectHashFinder) Apply(coefs []Coef, kw string) uint {
	h := phf.apply(uint(phf.TableSize)-1, coefs, kw)
	return h
}

func (phf *PerfectHashFinder) apply(mask uint, coefs []Coef, kw string) uint {
	h := uint(len(kw)) * coefs[len(coefs)-1].Value
	for i := 1; i < len(coefs)-1; i++ {
		idx := coefs[i].IndexApplied
		if idx < 0 && -idx < len(kw) {
			h += uint(kw[len(kw)+idx]) * coefs[i].Value
		} else if idx >= 0 && idx < len(kw) {
			h += uint(kw[idx]) * coefs[i].Value
		}
	}
	return h & mask
}
