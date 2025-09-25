package examples

type Example interface {
	Init()
	Run()
}

func InitExamples(examples []Example) {
	for _, e := range examples {
		e.Init()
	}
}

func RunExamples(examples []Example) {
	for _, e := range examples {
		e.Run()
	}
}
