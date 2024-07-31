package models

import "time"

type Article struct{
    Title string `json:"title"`
    Author string `json:"author"`
    Content string `json:"content"`
    ID int `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
