package scheduling

import (
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	"courseScheduling/models"
	"github.com/MaxHalford/eaopt"
)

type Generator struct {
	params *Params

	cntLessons   int
	allPlacement []placement
	timespanMap  map[int]*models.Timespan

	p *sync.Pool
}

func NewGenerator(params *Params) *Generator {
	g := &Generator{params: params}
	g.cntLessons = 0
	for i := range params.AllInstructedClazz {
		g.cntLessons += params.AllInstructedClazz[i].Instruct.Course.LessonsPerWeek
	}
	g.allPlacement = product2Placement(g.params.AllClazzroom, g.params.AllTimespan, availableWeekday)
	g.timespanMap = make(map[int]*models.Timespan, len(g.params.AllTimespan))
	for i := range g.params.AllTimespan {
		g.timespanMap[g.params.AllTimespan[i].Id] = g.params.AllTimespan[i]
	}
	// g.p = NewSyncPool(1, g.cntLessons, 100)
	g.p = &sync.Pool{
		New: func() interface{} {
			return make([]interface{}, 0, g.cntLessons)
		},
	}
	return g
}

var availableWeekday = []int{1, 2, 3, 4, 5}

func (g *Generator) GenerateSchedule() (result *GeneticSchedule) {
	g.printParams()
	config := eaopt.NewDefaultGAConfig()
	config.NGenerations = Config.MaxGenerations
	config.ParallelEval = true
	config.Model = eaopt.ModMutationOnly{Strict: true}
	timeout := time.NewTimer(Config.Timeout)
	var (
		LastFitness     = math.NaN()
		LastFitnessKeep = 0
	)
	// Stop when fitness is unchanged (precision is as FitnessJudgePrecision) in 500 generations
	config.EarlyStop = func(ga *eaopt.GA) bool {
		bestCandidate := ga.HallOfFame[0].Genome.(*GeneticSchedule)
		invalid := bestCandidate.Invalidity()
		if invalid == 0.0 {
			ga.Model = &eaopt.ModDownToSize{
				NOffsprings: 101,
				SelectorA:   eaopt.SelTournament{5},
				SelectorB:   eaopt.SelElitism{},
				MutRate:     0.9,
				CrossRate:   0.01,
			}
		} else {
			ga.Model = &eaopt.ModMutationOnly{Strict: true}
		}
		select {
		case <-timeout.C:
			return true
		default:
			return LastFitnessKeep > Config.StopWhenFitnessKeep
		}
	}
	// Add a custom print function to track progress
	config.Callback = func(ga *eaopt.GA) {
		//if ga.HallOfFame[0].Genome == nil {
		//	return
		//}
		bestCandidate := ga.HallOfFame[0].Genome.(*GeneticSchedule)
		fitness := ga.HallOfFame[0].Fitness
		invalid := bestCandidate.Invalidity()
		if invalid == 0.0 && math.Abs(fitness-LastFitness) < Config.FitnessJudgePrecision {
			LastFitnessKeep++
		} else {
			LastFitness = fitness
			LastFitnessKeep = 0
		}
		if LastFitnessKeep == 0 {
			_, err := bestCandidate.Evaluate()
			if err != nil {
				log.Println(err)
				return
			}
			invalidity := bestCandidate.Invalidity()

			log.Printf("%d) Result -> \n"+
				"h:%+v s:%+v\n "+
				"Invalidity:%.0f HowBad:%f\n",
				ga.Generations,
				// bestCandidate,
				bestCandidate.scores.h,
				bestCandidate.scores.s,
				invalidity, ga.HallOfFame[0].Fitness)
		}
	}
	var ga, err = config.NewGA()
	if err != nil {
		log.Println(err)
		return nil
	}
	// Run the GA
	err = ga.Minimize(func(rng *rand.Rand) eaopt.Genome {
		return MakeGeneticSchedule(g, rng)
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	// Final result
	if len(ga.HallOfFame) < 1 {
		log.Printf("We fail to find a solution after %d generations in %s\n", ga.Generations, ga.Age)
		return nil
	}
	result = ga.HallOfFame[0].Genome.(*GeneticSchedule)
	if result.Invalidity() != 0.0 {
		log.Printf("We find an invalid solution after %d generations in %s\n", ga.Generations, ga.Age)
		return nil
	}
	log.Printf("FINAL %d) Result -> \n"+
		"%v \n"+
		"h:%+v s:%+v\n "+
		"Invalidity:%.0f HowBad:%f\n",
		ga.Generations,
		result,
		result.scores.h,
		result.scores.s,
		result.Invalidity(), ga.HallOfFame[0].Fitness)
	log.Printf("Optimal solution obtained after %d generations in %s\n", ga.Generations, ga.Age)
	return result
}

func (g *Generator) printParams() {
	params := g.params
	log.Println("AllInstructs:")
	for _, item := range params.AllInstructedClazz {
		log.Printf("%+v", item)
	}
	log.Println()

	log.Println("AllClazzroom:")
	for _, item := range params.AllClazzroom {
		log.Printf("%+v", item)
	}
	log.Println()

	log.Println("AllTimespan:")
	for _, item := range params.AllTimespan {
		log.Printf("%+v", item)
	}
	log.Println()
}
