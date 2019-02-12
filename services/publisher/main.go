package publisher

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

// Server サービス構造体
type Server struct {
	addr   string
	tracer opentracing.Tracer
	logger log.Factory
}

// ConfigOptions オプション
type ConfigOptions struct {
	Port int // ポート番号
}

// NewServer publisher.Server を作成
func NewServer(options ConfigOptions, tracer opentracing.Tracer, logger log.Factory) *Server {
	return &Server{
		addr:   net.JoinHostPort("0.0.0.0", strconv.Itoa(options.Port)),
		tracer: tracer,
		logger: logger,
	}
}

// RunE サーバーを開始する
func (s *Server) RunE() error {
	mux := s.createServerMux()
	s.logger.Bg().Info("Starting Publisher Server", zap.String("address", s.addr))
	return http.ListenAndServe(s.addr, mux)
}

func (s *Server) createServerMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/publish", s.publish)
	return mux
}

func (s *Server) publish(w http.ResponseWriter, r *http.Request) {
	// Lesson02ではこのようにContextからSpanを生成していた
	// ctx := r.Context()
	// span, ctx := opentracing.StartSpanFromContext(ctx, "publisher")

	// Lesson03 Extract
	// 伝播されてきたSpanから子のSpanを作成し、作成したSpanに span.kind=server というタグをセットする
	spanCtx, _ := s.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := s.tracer.StartSpan("publish", ext.RPCServerOption(spanCtx))
	defer span.Finish()
	helloStr := r.FormValue("helloStr")
	fmt.Println(helloStr)
}
