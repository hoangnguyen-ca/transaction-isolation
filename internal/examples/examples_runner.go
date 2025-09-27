package examples

type Example interface {
	Init()
	Run(iterations int)
}

func InitExamples(examples []Example) {
	for _, e := range examples {
		e.Init()
	}
}

func RunExamples(examples []Example, iterations int) {
	for _, e := range examples {
		e.Run(iterations)
	}
}
