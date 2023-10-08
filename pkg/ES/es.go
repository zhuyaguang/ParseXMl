package ES

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"patentExtr/pkg"
)

func InitClient() (*elastic.Client, context.Context) {
	// 创建ES client用于后续操作ES
	client, err := elastic.NewClient(
		// 设置ES服务地址，支持多个地址
		elastic.SetURL("http://10.101.32.33:9200", "http://10.101.32.33:9201"),
		// 设置基于http base auth验证的账号和密码
		elastic.SetBasicAuth("user", "secret"))
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()
	return client, ctx
}

func InsertToES(es *elastic.Client, doc pkg.Document, ctx context.Context) error {

	// 使用client创建一个新的文档
	put1, err := es.Index().
		Index("patent566"). // 设置索引名称
		Id("1").            // 设置文档id
		BodyJson(doc).      // 指定前面声明的微博内容
		Do(ctx)             // 执行请求，需要传入一个上下文对象
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)
	return nil
}
