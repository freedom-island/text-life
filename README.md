# Text Life Server

开源参与者，欢迎你！

如果你对任意的规范感到不满，请参与 Issue 中的相关讨论。

如果你对游戏的内容有所建议，也请搜索并参与已存在的 Issue，或发起尚未创建的 Issue。

## 客户端请读

### 接口规范

服务接口采用 Protobuf 作为结构，但我对其了解不多（熟悉这些工具，是这个项目的目标之一），希望想这些教程能提供一些帮助：

- [Websocket and Protobuf integration in Javascript -
  Andrea Olivato](https://dev.to/andreaolivato/websocket-and-protobuf-integration-in-javascript-3m5p)

> 下文中，[ 和 ] 标识一个部分，并不是需要你输入；被方括号包裹的内容，才是需要你定义/使用的。

接口发起，请使用 `[路由][|][JSON?]`

## 服务端请读

创建路由：`protoc -I=proto --go_out=. proto/api.proto`
