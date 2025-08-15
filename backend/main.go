package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "backend/config"
    "backend/routes"
)

func main() {
    // 環境変数から設定を読み込み
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // データベース接続
    db, err := config.ConnectDatabase()
    if err != nil {
        log.Fatal("データベース接続エラー:", err)
    }
    defer db.Close()

    // Ginルーターを作成
    router := gin.Default()
    
    // CORS設定
    router.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

    // ルート設定
    routes.SetupRoutes(router, db)

    // サーバー起動
    log.Printf("MCR部ホームページバックエンドサーバーを起動中... ポート: %s", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}