package model

const TableNameArticleTag = "article_tag"

type ArticleTag struct {
	ArticleID int64 `gorm:"column:article_id" json:"article_id"`
	TagID     int64 `gorm:"column:tag_id" json:"tag_id"`
}

func (articleTag *ArticleTag) TableName() string {
	return TableNameArticleTag
}
