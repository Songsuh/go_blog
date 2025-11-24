package model

const TableNameArticle = "article"

type Article struct {
	BaseModel
	ID         int32  `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:id" json:"Id"`
	Title      string `gorm:"column:title;type:varchar(100);not null;comment:标题" json:"Title"`
	Desc       string `gorm:"column:desc;type:longtext;not null;comment:描述" json:"Desc"`
	Content    string `gorm:"column:content;type:longtext;not null;comment:内容" json:"Content"`
	Img        string `gorm:"column:img;type:longtext;not null;comment:图片地址" json:"Img"`
	Type       string `gorm:"column:type;type:tinyint(3);not null;comment:类型(1-原创 2-转载 3-翻译)" json:"Type"`
	Status     string `gorm:"column:status;type:tinyint(1);not null;comment:状态(1-公开 2-私密)" json:"Status"`
	IsStop     string `gorm:"column:author;type:tinyint(1);not null;comment:是否停止" json:"IsStop"`
	IsDelete   string `gorm:"column:isDelete;type:tinyint(1);not null;comment:是否删除" json:"IsDelete"`
	OriginUrl  string `gorm:"column:originUrl;type:longtext;not null;comment:原文链接" json:"OriginUrl"`
	CategoryId string `gorm:"column:categoryId;type:bigint(20);not null;comment:分类Id" json:"CategoryId"`
	UserId     string `gorm:"column:userId;type:bigint(20);not null;comment:用户Id" json:"UserId"`
}

func (article *Article) TableName() string {
	return TableNameArticle
}
