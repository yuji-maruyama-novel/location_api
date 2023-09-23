package main

import (
	"db"
	"fmt"
	"gqlgen/graph"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "3004"

func main() {
	loadEnv()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// ログファイル名に日付を含める
	logFileName := "./logs/" + time.Now().Format("2006-01-02") + ".log"

	// ログファイルを作成または開く
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}

	// ログ出力をファイルに設定
	log.SetOutput(logFile)

	// データベースへの接続処理
	db := db.ConnectGORM(logFile)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// .envファイル全体の読み込み
func loadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}
