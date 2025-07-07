package model

import (
	"Project1/message"
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	ArticleTitle string `gorm:"type:varchar(255);not null" json:"article_title"`
	Author       string `gorm:"type:varchar(255);not null" json:"author"`
	Content      string `gorm:"type:varchar(255);not null" json:"content"`
}

func CreateArticle(article *Article) int {
	if err := db.Create(&article).Error; err != nil {
		return message.ERROR
	}
	return message.SUCCESS

}

func GetArticleByTitle(title string) (*Article, int) {
	//根据书名查询信息
	// 1. 尝试从Redis获取
	cacheKey := fmt.Sprintf("article_%s", title)

	var val, err = rdb.Get(context.Background(), cacheKey).Result()
	if err == nil {
		article := &Article{}
		if json.Unmarshal([]byte(val), article) == nil {
			return article, message.SUCCESS // 缓存命中
		}
	}

	//2.缓存没有命中,从数据库中取
	var article Article
	if err := db.First(&article, "article_title = ?", title).Error; err != nil {
		return nil, message.ERROR
	}
	//3.存入redis
	JsonData, _ := json.Marshal(article)
	rdb.Set(context.Background(), cacheKey, JsonData, 3600*time.Second)
	return &article, message.SUCCESS
}
