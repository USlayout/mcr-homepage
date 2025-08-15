package routes

import (
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
    // ヘルスチェック
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "OK",
            "message": "MCR部ホームページバックエンドが正常に動作中",
            "version": "1.0.0",
        })
    })

    // API v1 グループ
    api := router.Group("/api/v1")
    {
        // 公開API（認証不要）
        public := api.Group("/public")
        {
            public.GET("/posts", getPublicPosts)         // 公開投稿一覧
            public.GET("/posts/:id", getPostByID)        // 投稿詳細
            public.GET("/events", getPublicEvents)       // 公開イベント一覧
            public.GET("/gallery", getPublicGallery)     // 公開ギャラリー
            public.GET("/downloads", getPublicDownloads) // 公開ダウンロード
        }

        // 認証が必要なAPI
        auth := api.Group("/auth")
        {
            auth.POST("/login", login)          // ログイン
            auth.POST("/logout", logout)        // ログアウト
            auth.GET("/me", getCurrentUser)     // 現在のユーザー情報
        }

        // 管理者・部員専用API
        admin := api.Group("/admin")
        // admin.Use(authMiddleware) // 認証ミドルウェア（後で実装）
        {
            admin.POST("/posts", createPost)           // 投稿作成
            admin.PUT("/posts/:id", updatePost)        // 投稿更新
            admin.DELETE("/posts/:id", deletePost)     // 投稿削除
            admin.POST("/events", createEvent)         // イベント作成
            admin.PUT("/events/:id", updateEvent)      // イベント更新
            admin.DELETE("/events/:id", deleteEvent)   // イベント削除
        }

        // ファイルアップロード
        upload := api.Group("/upload")
        // upload.Use(authMiddleware)
        {
            upload.POST("/image", uploadImage)         // 画像アップロード
            upload.POST("/video", uploadVideo)         // 動画アップロード
            upload.POST("/file", uploadFile)           // ファイルアップロード
        }
    }

    // 静的ファイル配信
    router.Static("/uploads", "./uploads")
}

// 仮の実装（後で各handlerファイルに移動）
func getPublicPosts(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "公開投稿一覧",
        "posts": []gin.H{
            {"id": 1, "title": "MCR部の活動が始まりました！", "type": "news"},
            {"id": 2, "title": "3Dモデル制作に挑戦", "type": "activity"},
        },
    })
}

func getPostByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": "投稿詳細",
        "id": id,
    })
}

func getPublicEvents(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "公開イベント一覧",
        "events": []gin.H{
            {"id": 1, "title": "MCR部体験会", "date": "2025-09-01"},
            {"id": 2, "title": "プログラミング勉強会", "date": "2025-09-15"},
        },
    })
}

func getPublicGallery(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "公開ギャラリー",
        "gallery": []gin.H{
            {"id": 1, "title": "初回Blender作品", "type": "3d_model"},
            {"id": 2, "title": "部活動写真", "type": "image"},
        },
    })
}

func getPublicDownloads(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "公開ダウンロード",
        "downloads": []gin.H{
            {"id": 1, "title": "MCRツール v1.0", "type": "zip"},
            {"id": 2, "title": "部員制作ゲーム", "type": "unity"},
        },
    })
}

// 認証系（後で実装）
func login(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "ログイン（未実装）"})
}

func logout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "ログアウト（未実装）"})
}

func getCurrentUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "現在のユーザー（未実装）"})
}

// 管理系（後で実装）
func createPost(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "投稿作成（未実装）"})
}

func updatePost(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "投稿更新（未実装）"})
}

func deletePost(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "投稿削除（未実装）"})
}

func createEvent(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "イベント作成（未実装）"})
}

func updateEvent(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "イベント更新（未実装）"})
}

func deleteEvent(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "イベント削除（未実装）"})
}

// アップロード系（後で実装）
func uploadImage(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "画像アップロード（未実装）"})
}

func uploadVideo(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "動画アップロード（未実装）"})
}

func uploadFile(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "ファイルアップロード（未実装）"})
}
