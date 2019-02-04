package lesson01

import (
	"io"

	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	opentracing "github.com/opentracing/opentracing-go"
	spanLog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
)

// Client 構造体
type Client struct {
	tracer opentracing.Tracer
	logger log.Factory
	closer io.Closer
}

// ConfigOptions オプション
type ConfigOptions struct {
	Message string
}

var options = &ConfigOptions{}

// NewClient Client構造体を作成する
func NewClient(o *ConfigOptions, tracer opentracing.Tracer, logger log.Factory, closer io.Closer) *Client {
	options = o
	return &Client{
		tracer: tracer,
		logger: logger,
		closer: closer,
	}
}

// RunE プログラム開始。エラーを返す
func (c *Client) RunE() error {
	c.logger.Bg().Info("Sandbox Start")
	c.logger.Bg().Info("message", zap.String("message", options.Message))

	span := c.tracer.StartSpan("say-hello")
	span.SetTag("hello-to", options.Message)

	span.LogFields(
		spanLog.String("event", "hoge"),
	)

	span.LogKV("event", "println")

	span.Finish()
	return nil
}
