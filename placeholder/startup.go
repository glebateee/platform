package placeholder

import (
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"sync"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		//&SimpleMessageComponent{},
		handling.NewRouter(
			handling.HandlerEntry{Prefix: "", Handler: NameHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: DayHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: WeatherHandler{}},
		).AddMethodAlias("/", NameHandler.GetNames),
	)
}

func Start() {
	results, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
