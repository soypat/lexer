package pato

import (
	"errors"
	"slices"
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
	tableSizes := []int{4, 5, 6, 7, 8, 9, 10}
	// Coefficients for perfect hash function.
	// First coeficcient is length multiplication.
	// This method of searching for a perfect hash is quite robust
	// and works even with Fortran keywords (70+ keywords, ~20 sharing identical END start)
	const maxCoef = 32
	coefs := make([]Coef, 3)
	for i := range len(coefs) {
		coefs[i].IndexApplied = i
		coefs[i].Op = TokPlus
		coefs[i].OnlyPow2 = true // Beware: Is quite limiting when one has lots of keywords, but has performance benefits.
	}
	phf := PerfectHashFinder{
		DefaultMaxCoef: maxCoef,
	}
	for _, tableSize := range tableSizes {
		phf.TableSizeBits = tableSize
		t.Logf("Trying table size %d...", 1<<tableSize)
		currentAttempt, err := phf.Search(coefs, keywords)
		attempts += currentAttempt
		if err != nil && err != ErrNoCoefficientsFound {
			t.Fatal(err)
		} else if err == nil {
			t.Logf("FOUND with tableSize=%d after %d attempts! (%d total attempts)", tableSize, currentAttempt, attempts)
			for i := range coefs {
				t.Logf("coef%d=%d op=%s", i, coefs[i].Value, coefs[i].Op.String())
			}
			return
		}

	}
	t.Error("No perfect hash found after", attempts, "attempts")
}

type PerfectHashFinder struct {
	TableSizeBits  int
	DefaultMaxCoef uint
	// HashLastIndices
	hashmap []uint
}
type Coef struct {
	IndexApplied int  // Index at which hash consumes byte. Negative value indexes from the end.
	Value        uint // Coefficient value to multiply byte at index.
	MaxValue     uint
	StartValue   uint
	OnlyPow2     bool
	Op           Token
}

var ErrNoCoefficientsFound = errors.New("no coefficients found")

func (c *Coef) init() {
	if c.StartValue == 0 {
		c.Value = 1
	} else {
		c.Value = c.StartValue
	}
	if c.Op == 0 {
		c.Op = TokPlus
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

func (phf *PerfectHashFinder) Search(coefs []Coef, inputs []string) (int, error) {
	if phf.TableSizeBits <= 0 || phf.TableSizeBits > 32 {
		return 0, errors.New("zero/negative bits for table size or too large")
	} else if len(coefs) == 0 {
		return 0, errors.New("require at least one coefficient to find perfect hash")
	} else if len(inputs) == 0 {
		return 0, errors.New("zero inputs")
	}
	err := phf.ConfigureCoefsWithDefaults(coefs)
	if err != nil {
		return 0, err
	}
	tblsz := 1 << phf.TableSizeBits
	phf.hashmap = slices.Grow(phf.hashmap[:0], tblsz)[:tblsz]
	hashmap := phf.hashmap
	mask := uint(tblsz) - 1
	currentAttempt := 0
	for {
		currentAttempt++
		attemptSuccess := true
		clear(hashmap)
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
			return currentAttempt, nil
		}
		coefs[0].increment()
		for i := 0; coefs[i].saturated() && i < len(coefs)-1; i++ {
			coefs[i].init()
			coefs[i+1].increment()
		}
		// Check for super-saturation.
		if coefs[len(coefs)-1].Value > coefs[len(coefs)-1].MaxValue {
			break
		}
	}
	return currentAttempt, ErrNoCoefficientsFound
}

func (phf *PerfectHashFinder) ConfigureCoefsWithDefaults(coefs []Coef) error {
	for i := range coefs {
		coefs[i].init()
		if coefs[i].MaxValue == 0 {
			if phf.DefaultMaxCoef <= 0 {
				return errors.New("default max coefficient need be set and positive for input")
			}
			coefs[i].MaxValue = phf.DefaultMaxCoef
		}
	}
	return nil
}

func (phf *PerfectHashFinder) Apply(coefs []Coef, kw string) uint {
	h := phf.apply((1<<phf.TableSizeBits)-1, coefs, kw)
	return h
}

func (phf *PerfectHashFinder) apply(mask uint, coefs []Coef, kw string) uint {
	h := uint(len(kw)) * coefs[len(coefs)-1].Value
	for i := 0; i < len(coefs)-1; i++ {
		idx := coefs[i].IndexApplied
		var a uint
		if idx < 0 && -idx <= len(kw) {
			a = uint(kw[len(kw)+idx]) * coefs[i].Value
		} else if idx >= 0 && idx < len(kw) {
			a = uint(kw[idx]) * coefs[i].Value
		}
		switch coefs[i].Op {
		case TokPlus:
			h += a
		case TokHat:
			h ^= a
		case TokAsterisk:
			h *= a
		default:
			panic("unsupported operation")
		}

	}
	return h & mask
}
