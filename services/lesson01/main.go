package lesson01

import (
	"io"

	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	opentracing "github.com/opentracing/opentracing-go"
	spanLog "github.com/opentracing/opentracing-go/log"
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
	c.logger.Bg().Info("Lesson01 Start")

	span := c.tracer.StartSpan("say-hello") // "say-hello" という名称のSpanを生成
	helloStr := options.Message
	span.SetTag("hello-to", helloStr) // Tagに"hello-to"をセット

	span.LogFields(
		spanLog.String("event", "string-format"),
		spanLog.String("value", helloStr),
	)

	println(helloStr)
	span.LogKV("event", "println")

	span.Finish()
	return nil
}
