# influxdb-example
添加golang使用influxdb例子
## influxdb install
[下载地址](https://portal.influxdata.com/downloads)
选择InfluxDB,根据不同平台选择下载安装

## 关键概念
- **database** : InfluxDB 的数据库类似与传统关系型数据库，再InfluxDB中database作为用户
、保留策略（retention policies）、 连续查询（continuous queries）、要保存的时间序列数据等的逻辑容器
- **measurement** : measurement类似其他关系型数据库中的table，负责保存tags，fields，time
- **retention policy** : 描述了数据保留的时间策略以及如何拷贝存储到cluster上数据。一个measurement可以属于多个retention policy
- **fields** : fields包括fields key和fields value;fields key为strings类型且作为元数据存储，fields value是strings, floats, integers, or Booleans类型,由于influxDB是时间序列数据库，所以fields value总是与时间戳相关联。另外在存储时fields是必须包含的并且fields是无法被索引到的。
- **field set** : field set 由一对field-key和field-value组成
- **point** : 
- **tags** : tags包括tags key和tags value且都为strings类型并作为元数据存储。tags在存储时是可选的，但是tags 是可以被索引的，所以，如果需要提高查询效率，把需要查询的字段作为tags存储是一个好方法
- **series** : series 是retention policy, measurement,tag 的集合
- **timestamp** : 展示日期和时间，与特定数据相关联。每条数据都会有time字段，如果插入时不指定时间会自动使用当前时间。
- **point** ： point是在相同的series里面拥有相同timestamp的field set。
- [官方文档](https://docs.influxdata.com/influxdb/v1.6/concepts/key_concepts/)
