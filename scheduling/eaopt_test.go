package scheduling

import (
	"bytes"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/MaxHalford/eaopt"
)

func TestGeneticLib(t *testing.T) {
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	ga.NGenerations = 10e9
	timeout := time.NewTimer(30 * time.Second)
	ga.EarlyStop = func(ga *eaopt.GA) bool {
		select {
		case <-timeout.C:
			return true
		default:
			return ga.HallOfFame[0].Fitness == 0
		}
	}
	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		// Concatenate the elements from the best individual and display the result
		var buffer bytes.Buffer
		for _, letter := range ga.HallOfFame[0].Genome.(GeneticString) {
			buffer.WriteString(letter)
		}
		t.Logf("%d) Result -> %s (%.0f mismatches)\n", ga.Generations, buffer.String(), ga.HallOfFame[0].Fitness)
	}

	// Run the GA
	err = ga.Minimize(MakeGeneticStrings)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	result := ga.HallOfFame[0].Genome.(GeneticString)
	if strings.Join(result, "") != strings.Join(testTarget, "") {
		t.Fail()
	}
}

var (
	testCorpus = strings.Split("abcdefghijklmnopqrstuvwxyz ", "")
	testTarget = strings.Split("zhenghaoren shabi", "")
)

type GeneticString []string

// Evaluate a Strings slice by counting the number of mismatches between itself
// and the target string.
func (X GeneticString) Evaluate() (mismatches float64, err error) {
	for i, s := range X {
		if s != testTarget[i] {
			mismatches++
		}
	}
	return
}

// Mutate a Strings slice by replacing it's elements by random characters
// contained in  a corpus.
func (X GeneticString) Mutate(rng *rand.Rand) {
	eaopt.MutUniformString(X, testCorpus, 3, rng)
}

// Crossover a Strings slice with another by applying 2-point crossover.
func (X GeneticString) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXString(X, Y.(GeneticString), 2, rng)
}

// Clone a Strings slice..
func (X GeneticString) Clone() eaopt.Genome {
	var XX = make(GeneticString, len(X))
	copy(XX, X)
	return XX
}

// MakeStrings creates random Strings slices by picking random characters from a
// corpus.
func MakeGeneticStrings(rng *rand.Rand) eaopt.Genome {
	return GeneticString(eaopt.InitUnifString(uint(len(testTarget)), testCorpus, rng))
}
