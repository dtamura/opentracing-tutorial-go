package lesson02

import (
	"context"
	"fmt"

	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	opentracing "github.com/opentracing/opentracing-go"
	spanLog "github.com/opentracing/opentracing-go/log"
)

// Client 構造体
type Client struct {
	tracer opentracing.Tracer
	logger log.Factory
}

// ConfigOptions オプション
type ConfigOptions struct {
	Message string
}

var options ConfigOptions

// NewClient Client構造体を作成する
func NewClient(o ConfigOptions, tracer opentracing.Tracer, logger log.Factory) *Client {
	options = o
	return &Client{
		tracer: tracer,
		logger: logger,
	}
}

// Run プログラム開始
func (c *Client) Run() {
	c.logger.Bg().Info("Leson02 Start")

	span := c.tracer.StartSpan("say-hello") // "say-hello" という名称のSpanを生成
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	helloStr := options.Message
	span.SetTag("hello-to", helloStr) // Tagに"hello-to"をセット
	defer span.Finish()

	str := formatString(ctx, helloStr)
	printHello(ctx, str)
}

func formatString(ctx context.Context, helloTo string) string {
	// span := rootSpan.Tracer().StartSpan("formatString")
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		spanLog.String("event", "string-format"),
		spanLog.String("value", helloStr),
	)

	return helloStr
}

// printHello は文字列を出力する
func printHello(ctx context.Context, helloStr string) {
	// span := rootSpan.Tracer().StartSpan("printHello")
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	println(helloStr)
	span.LogKV("event", "println")
}
