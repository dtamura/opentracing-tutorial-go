package sandbox

import (
	"io"

	"github.com/dtamura/hello-cobra/lib/log"
	opentracing "github.com/opentracing/opentracing-go"
	spanLog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
)

// Server
type Server struct {
	tracer opentracing.Tracer
	logger log.Factory
	closer io.Closer
}

// ConfigOptions オプション
type ConfigOptions struct {
	Message string
}

var options = &ConfigOptions{}

// NewServer creates a new frontend.Server
func NewServer(o *ConfigOptions, tracer opentracing.Tracer, logger log.Factory, closer io.Closer) *Server {
	options = o
	return &Server{
		tracer: tracer,
		logger: logger,
		closer: closer,
	}
}

// RunE start service
func (s *Server) RunE() error {
	s.logger.Bg().Info("Sandbox Start")
	s.logger.Bg().Info("message", zap.String("message", options.Message))

	span := s.tracer.StartSpan("say-hello")
	span.SetTag("hello-to", options.Message)

	span.LogFields(
		spanLog.String("event", "hoge"),
	)

	span.LogKV("event", "println")

	span.Finish()
	return nil
}
