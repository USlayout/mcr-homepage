package models

import (
    "time"
)

// 投稿（活動報告・お知らせ）
type Post struct {
    ID          int       `json:"id" db:"id"`
    Title       string    `json:"title" db:"title"`
    Content     string    `json:"content" db:"content"`
    Type        string    `json:"type" db:"type"`           // "news", "activity", "result"
    AuthorID    int       `json:"author_id" db:"author_id"`
    AuthorName  string    `json:"author_name" db:"author_name"`
    ImageURL    string    `json:"image_url" db:"image_url"`
    VideoURL    string    `json:"video_url" db:"video_url"`
    Tags        string    `json:"tags" db:"tags"`           // JSON形式で複数タグ
    IsPublished bool      `json:"is_published" db:"is_published"`
    ViewCount   int       `json:"view_count" db:"view_count"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// イベント（大会・活動予定）
type Event struct {
    ID          int       `json:"id" db:"id"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    EventDate   time.Time `json:"event_date" db:"event_date"`
    Location    string    `json:"location" db:"location"`
    Type        string    `json:"type" db:"type"`           // "competition", "meeting", "exhibition"
    IsPublic    bool      `json:"is_public" db:"is_public"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// ユーザー（部員）
type User struct {
    ID        int       `json:"id" db:"id"`
    Username  string    `json:"username" db:"username"`
    Email     string    `json:"email" db:"email"`
    Password  string    `json:"-" db:"password"`          // JSONでは出力しない
    Role      string    `json:"role" db:"role"`           // "admin", "member", "graduate"
    Grade     int       `json:"grade" db:"grade"`         // 学年
    RealName  string    `json:"real_name" db:"real_name"`
    Bio       string    `json:"bio" db:"bio"`             // 自己紹介
    Skills    string    `json:"skills" db:"skills"`       // 得意分野（JSON）
    IsActive  bool      `json:"is_active" db:"is_active"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// ギャラリー（作品・写真）
type Gallery struct {
    ID          int       `json:"id" db:"id"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    Type        string    `json:"type" db:"type"`           // "image", "video", "3d_model", "unity"
    FileURL     string    `json:"file_url" db:"file_url"`
    ThumbnailURL string   `json:"thumbnail_url" db:"thumbnail_url"`
    AuthorID    int       `json:"author_id" db:"author_id"`
    AuthorName  string    `json:"author_name" db:"author_name"`
    Category    string    `json:"category" db:"category"`   // "blender", "programming", "photo"
    Tags        string    `json:"tags" db:"tags"`
    ViewCount   int       `json:"view_count" db:"view_count"`
    IsPublic    bool      `json:"is_public" db:"is_public"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// ダウンロード（制作物）
type Download struct {
    ID           int       `json:"id" db:"id"`
    Title        string    `json:"title" db:"title"`
    Description  string    `json:"description" db:"description"`
    FileURL      string    `json:"file_url" db:"file_url"`
    FileName     string    `json:"file_name" db:"file_name"`
    FileSize     int64     `json:"file_size" db:"file_size"`
    FileType     string    `json:"file_type" db:"file_type"`   // "zip", "exe", "blend", "unity"
    Category     string    `json:"category" db:"category"`
    AuthorID     int       `json:"author_id" db:"author_id"`
    AuthorName   string    `json:"author_name" db:"author_name"`
    Version      string    `json:"version" db:"version"`
    Requirements string    `json:"requirements" db:"requirements"` // 動作環境
    DownloadCount int      `json:"download_count" db:"download_count"`
    IsPublic     bool      `json:"is_public" db:"is_public"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
