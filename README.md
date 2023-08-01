# ParseXMl

专利总共分四种 
```shell
const PATENT1 = "中国外观设计专利授权公告标准化著录项目及切图数据" //30-S  industrial design
const PATENT2 = "中国发明专利申请公布标准化全文文本数据" //10-A  invention publication 
const PATENT3 = "中国发明专利授权公告标准化全文文本数据" //10-B  invention grant 
const PATENT4 = "中国实用新型专利授权公告标准化全文文本数据" //20-U utility model 
```
分别处理后保存在`/output`目录下面的 30-S 10-A 10-B 20-U 目录下

## 第一步：将专利的压缩包解压

 `/data/output/10-A/20220104/1/CN102020000545134CN00001138834480AFULZH20220104CN00A`


解压目录分别是 10-A 专利类型 保存日期 存储文件夹序号 专利目录

## 第二步：解析 xml 文件

保存成 CN102020000545134CN00001138834480AFULZH20220104CN00A.json 文件

启动命令：
nohup ./patentExtr --data="/data/sipo" --output="/data/output" --s-start="20220101" --a-start="20220101" --b-start="20220101" --u-start="20220101" > log.log 2>&1 &


## 未插入数据库
// /data/output/10-B/20220816/9/CN102022000190631CN00001142615810BFULZH20220513CN00R/CN102022000190631CN00001142615810BFULZH20220513CN00R.XML

## 下次插入的起点
// 最开始起点 [20220101 20220101  20220101 20220101]

## todo
* 格式化日志，显示时间
* 错误持久化，对漏掉对专利进行第二次插入