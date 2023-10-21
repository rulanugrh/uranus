package monitoring

import (
	"context"
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/rulanugrh/uranus/configs"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitConfig() {
	conf := configs.GetConfig()
	config := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 10,
		},

		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: conf.Jaeger.Host + conf.Jaeger.Port,
		},
	}

	closer, err := config.InitGlobalTracer("uranus-service")
	if err != nil {
		log.Println("cant init global tracer jaeger")
	}
	defer closer.Close()

}

func StartTracing(ctx context.Context, name string) opentracing.Span {
	trace, err := opentracing.StartSpanFromContext(ctx, name)
	if err.Err() != nil {
		return nil
	}

	return trace
}
