package entity

import "time"

type Video struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FileName  string    `json:"file_name" gorm:"size:255;not null"`
	FilePath  string    `json:"file_path" gorm:"size:255;not null"`
	Duration  int       `json:"duration" gorm:"not null"`
	Size      int64     `json:"size" gorm:"not null"`
	CreatedAt time.Time `json:"upload_time" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

type TrimVideoRequest struct {
	VideoID   int    `json:"video_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type MergeVideoRequest struct {
	VideoIDs []int64 `json:"video_ids"`
}

type ShareVideoRequest struct {
	VideoID    int    `json:"video_id"`
	ExpiryTime string `json:"expiry_time"`
}
