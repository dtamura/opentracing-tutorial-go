package lesson04

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	spanLog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
)

// Client 構造体
type Client struct {
	tracer opentracing.Tracer
	logger log.Factory
}

// ConfigOptions オプション
type ConfigOptions struct {
	Message  string
	Greeting string
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
	c.logger.Bg().Info("lesson04 Start")

	span := c.tracer.StartSpan("say-hello")
	span.SetBaggageItem("greeting", options.Greeting) // Baggageで文字列を送信
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	helloStr := options.Message
	span.SetTag("hello-to", helloStr)
	defer span.Finish()

	str, err := formatString(ctx, helloStr)
	if err != nil {
		c.logger.Bg().Error("Error while formatString", zap.Error(err))
		return
	}

	err = printHello(ctx, str)
	if err != nil {
		c.logger.Bg().Error("Error while PrintString", zap.Error(err))
	}
}

func formatString(ctx context.Context, helloTo string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloTo", helloTo)
	url := "http://localhost:8081/format?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	// Inject
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	client := &http.Client{}
	resp, err := client.Do(req) // リクエスト送信
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}
	helloStr := string(body)

	span.LogFields(
		spanLog.String("event", "string-format"),
		spanLog.String("value", helloStr),
	)

	return helloStr, nil
}

// printHello は文字列を出力する
func printHello(ctx context.Context, helloStr string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloStr", helloStr)
	url := "http://localhost:8082/publish?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)

	// Inject
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	client := &http.Client{}
	resp, err := client.Do(req) // リクエスト送信
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}

	span.LogKV("event", "println")
	return nil
}
