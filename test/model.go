package main

import (
	"fmt"
	"go_blog/internal/dao"
	"go_blog/internal/model"
)

// var env, _ = os.ReadFile(".env")
// var configEnv = flag.String("c", "dev", "the config file")
// 通用调用方法
func main() {
	//svc.GetSvc()
	//var articles []model.Article
	qb := dao.NewQueryBuilder("blog1", &model.Article{}) // 注意这里改为 &[]model.Article{}
	conditions := dao.Conditions{
		{Field: "id", Operator: "=", Value: []any{1}},
	}
	result := qb.FindOne(dao.QueryOptions{
		Conditions: conditions,
		OrderBy:    "id DESC",
	})
	fmt.Printf("%+v\n", result)

	if result.Error != nil {
		fmt.Printf("查询错误: %v\n", result.Error)
		return
	}

	// 类型断言，将result.Data转换为[]model.Article的指针
	articleList, ok := result.Data.(*model.Article)
	if !ok {
		fmt.Println("数据类型转换失败")
		return
	}
	fmt.Printf("Article: %+v\n", articleList)
	fmt.Printf("第1条记录: ID=%d, Title=%s, Desc=%s\n, 创建时间：%s", articleList.ID, articleList.Title, articleList.Desc, articleList.CreatedAt)
	// 现在可以安全使用 articleList
	//for _, article := range *articleList {
	//	fmt.Printf("Article: %+v\n", article)
	//}
	//
	//// 打印每条数据
	//fmt.Printf("共查询到 %d 条记录\n", len(*articleList))
	//for i, article := range *articleList {
	//	fmt.Printf("第%d条记录: ID=%d, Title=%s, Desc=%s\n", i+1, article.ID, article.Title, article.Desc)
	//}
}
