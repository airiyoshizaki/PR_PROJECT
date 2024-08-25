package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	handler "github.com/route-test/interface/handler"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/route-test/docs" // これ注視する
)

// @title routetest.
// @version 1.0
// @description  routetest.
// @host localhost:6061
func main() {

	// 新しいServeMuxを作成し、ルートパスを登録
	r := chi.NewRouter()

	r.Get("/", handler.HomeHandler)
	r.Get("/about", handler.AboutHandler)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:6061/swagger/doc.json"),
	))

	// サーバー設定を含むhttp.Server構造体を作成
	server := &http.Server{
		Addr:         ":8080",          // サーバーがリッスンするアドレスとポート
		Handler:      r,                // リクエストを処理するハンドラー
		ReadTimeout:  15 * time.Second, // リクエストヘッダーの読み取りタイムアウト
		WriteTimeout: 15 * time.Second, // レスポンス書き込みのタイムアウト
	}

	// サーバーを起動し、リクエストを待機
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
