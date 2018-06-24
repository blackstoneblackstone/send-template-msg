# 微信发送模板消息

> 微信发送模板消息，由于每次发送都要跟微信服务器做交互，所以高并发的处理这个异步任务非常重要。
> 毫无疑问，go语言是最理想的实现语言。


## init

- 修改config.toml.bak文件为config.toml
- 填写文件内容，redis和mysql的配置。
- 建立数据库wpx-js, 运行mysql.sql文件，建立数据表