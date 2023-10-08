package ES

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"patentExtr/pkg"
	"testing"
)

func TestInsertToES(t *testing.T) {
	es, ctx := InitClient()
	// 准备要插入的文档
	doc := pkg.Document{
		Timestamp: "2022-03-01T12:00:00Z",
		Version:   "1.0",
		Abstract:  "这是一个示例文本33333。",
	}
	InsertToES(es, doc, ctx)

}

func TestConnect(t *testing.T) {

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
	ctx := context.Background()

	println(client)

	// 根据id查询文档
	get1, err := client.Get().
		Index("patent1").          // 指定索引名
		Id("1456232735810129999"). // 设置文档id
		Do(ctx)                    // 执行请求
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("文档id=%s 版本号=%d 索引名=%s\n", get1.Id, get1.Version, get1.Index)
	}

	// 手动将文档内容转换成go struct对象
	msg2 := pkg.Document{}
	// 提取文档内容，原始类型是json数据
	data, _ := get1.Source.MarshalJSON()
	// 将json转成struct结果
	json.Unmarshal(data, &msg2)
	// 打印结果
	fmt.Println(msg2.Name)

}
