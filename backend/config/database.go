package config

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

type Database struct {
    DB *sql.DB
}

func ConnectDatabase() (*sql.DB, error) {
    // 環境変数から接続情報を取得
    dbUser := getEnv("DB_USER", "mcr_user")
    dbPassword := getEnv("DB_PASSWORD", "mcr_password")
    dbHost := getEnv("DB_HOST", "database")
    dbPort := getEnv("DB_PORT", "3306")
    dbName := getEnv("DB_NAME", "mcr_homepage")

    // データソース名を構築
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    // データベースに接続
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("データベース接続エラー: %v", err)
    }

    // 接続テスト
    if err = db.Ping(); err != nil {
        return nil, fmt.Errorf("データベース疎通エラー: %v", err)
    }

    return db, nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
