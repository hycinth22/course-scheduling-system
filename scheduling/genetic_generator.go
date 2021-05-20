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

type Params struct {
	AllInstructedClazz []*models.InstructedClazz
	AllClazzroom       []*models.Clazzroom
	AllTimespan        []*models.Timespan
}

type Generator struct {
	params *Params

	cntLessons   int
	allPlacement []placement
	timespanMap  map[int]*models.Timespan

	p *sync.Pool

	config ConfigType
}

func NewGenerator(params *Params, config ConfigType) *Generator {
	g := &Generator{params: params, config: config}
	g.cntLessons = 0
	for i := range params.AllInstructedClazz {
		g.cntLessons += params.AllInstructedClazz[i].Instruct.Course.LessonsPerWeek
	}
	g.allPlacement = product2Placement(g.params.AllClazzroom, g.params.AllTimespan, availableWeekday)
	g.timespanMap = make(map[int]*models.Timespan, len(g.params.AllTimespan))
	for i := range g.params.AllTimespan {
		g.timespanMap[g.params.AllTimespan[i].Id] = g.params.AllTimespan[i]
	}
	g.p = &sync.Pool{
		New: func() interface{} {
			return make([]IEqual, 0, g.cntLessons)
		},
	}
	return g
}

var availableWeekday = []int{1, 2, 3, 4, 5}

func (g *Generator) GenerateSchedule() (result *GeneticSchedule, float64 float64) {
	var stage = 1
	var everStage2 = false
	g.printParams()
	config := eaopt.GAConfig{
		NPops:        g.config.NumOfPopulations,
		PopSize:      g.config.SizeOfPopulation,
		NGenerations: g.config.MaxGenerations,
		HofSize:      1,
		ParallelEval: true,
		Migrator:     eaopt.MigRing{NMigrants: 20},
		MigFrequency: 5,
		Speciator:    nil,
		Logger:       nil,
	}
	config.Model = &eaopt.ModGenerational{
		Selector:  eaopt.SelElitism{},
		MutRate:   g.config.MutationRate,
		CrossRate: g.config.CrossoverRate,
	}
	var timeout *time.Timer
	var (
		LastFitness     = math.NaN()
		LastFitnessKeep = 0
	)
	// Stop when fitness is unchanged (precision is as FitnessJudgePrecision) in g.config.StopWhenFitnessKeep generations
	config.EarlyStop = func(ga *eaopt.GA) bool {
		bestCandidate := ga.HallOfFame[0].Genome.(*GeneticSchedule)
		invalid := bestCandidate.Invalidity()
		if invalid == 0 {
			//ga.Model = &eaopt.ModDownToSize{
			//	NOffsprings: g.config.NumOfOffsprings,
			//	SelectorA:   eaopt.SelElitism{},
			//	SelectorB:   eaopt.SelElitism{},
			//	MutRate:     g.config.MutationRate,
			//	CrossRate:   g.config.CrossoverRate,
			//}
			//ga.Model = &eaopt.ModSteadyState{
			//	Selector:  eaopt.SelElitism{},
			//	MutRate:   g.config.MutationRate,
			//	CrossRate:   g.config.CrossoverRate,
			//	KeepBest: true,
			//}
			ga.Model = &eaopt.ModGenerational{
				Selector:  eaopt.SelElitism{},
				MutRate:   g.config.MutationRate,
				CrossRate: g.config.CrossoverRate,
			}
			stage = 2
			everStage2 = true
		} else {
			stage = 1
			/*			ga.Model = &eaopt.ModGenerational{
						Selector:  eaopt.SelElitism{},
						MutRate:   g.config.MutationRate,
						CrossRate: g.config.CrossoverRate,
					}*/
			//ga.Model = &eaopt.ModMutationOnly{Strict: true}
			ga.Model = &eaopt.ModDownToSize{
				NOffsprings: g.config.NumOfOffsprings,
				SelectorA:   eaopt.SelElitism{},
				SelectorB:   eaopt.SelElitism{},
				MutRate:     g.config.MutationRate,
				CrossRate:   g.config.CrossoverRate,
			}
		}
		select {
		case <-timeout.C:
			return true
		default:
			return LastFitnessKeep > g.config.StopWhenFitnessKeep
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
		if invalid == 0.0 && math.Abs(fitness-LastFitness) < g.config.FitnessJudgePrecision {
			LastFitnessKeep++
		} else {
			LastFitness = fitness
			LastFitnessKeep = 0
		}

		//if LastFitnessKeep == 0 {
		_, err := bestCandidate.Evaluate()
		if err != nil {
			log.Println(err)
			return
		}
		invalidity := bestCandidate.Invalidity()
		if invalidity > 0 {
			ga.Model = &eaopt.ModDownToSize{
				NOffsprings: g.config.NumOfOffsprings,
				SelectorA:   eaopt.SelElitism{},
				SelectorB:   eaopt.SelElitism{},
				MutRate:     g.config.MutationRate,
				CrossRate:   g.config.CrossoverRate,
			}
		}
		log.Printf("%d) Result -> \n"+
			"h:%+v s:%+v\n "+
			"Invalidity:%d HowBad:%f\n %d %v",
			ga.Generations,
			// bestCandidate,
			bestCandidate.scores.h,
			bestCandidate.scores.s,
			invalidity, ga.HallOfFame[0].Fitness, stage, everStage2)
		//}
	}
	var ga, err = config.NewGA()
	if err != nil {
		log.Println(err)
		return nil, math.Inf(1)
	}
	// Run the GA
	timeout = time.NewTimer(g.config.Timeout)
	err = ga.Minimize(func(rng *rand.Rand) eaopt.Genome {
		return MakeGeneticSchedule(g, rng)
	})
	if err != nil {
		log.Println(err)
		return nil, math.Inf(1)
	}

	// Final result
	if len(ga.HallOfFame) < 1 {
		log.Printf("We fail to find a solution after %d generations in %s\n", ga.Generations, ga.Age)
		return nil, math.Inf(1)
	}
	if !ga.HallOfFame.IsSortedByFitness() {
		ga.HallOfFame.SortByFitness()
	}
	result = ga.HallOfFame[0].Genome.(*GeneticSchedule)
	if result.Invalidity() != 0.0 {
		log.Printf("We find an invalid solution after %d generations in %s\n", ga.Generations, ga.Age)
		return nil, math.Inf(1)
	}
	log.Printf("FINAL %d) Result -> \n"+
		"%v \n"+
		"h:%+v s:%+v\n "+
		"Invalidity:%d HowBad:%f\n",
		ga.Generations,
		result,
		result.scores.h,
		result.scores.s,
		result.Invalidity(), ga.HallOfFame[0].Fitness)
	log.Printf("Optimal solution obtained after %d generations in %s\n", ga.Generations, ga.Age)
	return result, ga.HallOfFame[0].Fitness
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
