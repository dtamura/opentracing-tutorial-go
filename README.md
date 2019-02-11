Sample
=============================

Lesson01
------------------------

```cmd
make
export JAEGER_AGENT_HOST=<HOST>
bin/opentracing-tutorial lesson01 -m hoge
```


Lesson02 Context and Tracing Functions
-------------------------

- 複数の関数をトレースする
- 複数のスパンを1つのトレースに結合する
- Contextによって伝播させる


- 入力: 文字列
- formatString: 入力文字列を整形
- printHello: 文字列を表示

```cmd
make
export JAEGER_AGENT_HOST=<HOST>
bin/opentracing-tutorial lesson02 -m hoge
```

Lesson03 Tracing RPC Requests
-------------------

- 複数のマイクロサービス間のトランザクションをトレースする
- `Inject` , `Extract` を用いてプロセス間のContextを伝播する
- OpenTracing推奨のTagを適用する


lesson02の各関数をHTTPリクエストに変更した

Injectを使用することにより、RequestHeaderに Uber-Trace-Idが付与される


Lesson04 Baggage
-----------------------------

- 分散したコンテキストの伝播を理解する
- baggage