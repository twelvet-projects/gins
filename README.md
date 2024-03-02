# gins

基于Gin 进行常用第三方插件封装，保持它的轻量、性能
拥有日志、配置、数据库(Mysql、Mssql、Oracle、Pgsql、Sqlite)、Redis、Swagger基础功能

直接上手业务，无需过多关注基础建设

## 在线体验

- admin/123456

演示地址：[https://cloud.twelvet.cn](https://cloud.twelvet.cn)

## 支持Linux一件Docker启动
需要自行安装go、docker、docker-compose
```shell
# build
go build -v
# 进入脚本目录
cd ./docker
# 可执行权限
chmod 751 deploy.sh
# 执行启动（按需执行参数，[init|port|base|server|stop|rm]）
# 初始化
./deploy.sh init
# 基础服务
./deploy.sh base
# 启动gins
./deploy.sh server
```
## 开源共建

### 开源协议

twelvet 开源软件遵循 [Apache 2.0 协议](https://www.apache.org/licenses/LICENSE-2.0.html)。
允许商业使用，但务必保留类作者、Copyright 信息。

### 其他说明

1. 欢迎提交 [PR](https://github.com/twelvet-projects/twelvet/pulls)，注意对应提交对应分支
   代码规范 [spring-javaformat](https://github.com/spring-io/spring-javaformat)

   <details>
    <summary>代码规范说明</summary>

    1. 由于 <a href="https://github.com/spring-io/spring-javaformat" target="_blank">spring-javaformat</a>
       强制所有代码按照指定格式排版，未按此要求提交的代码将不能通过合并（打包）
    2. 如果使用 IntelliJ IDEA
       开发，请安装自动格式化软件 <a href="https://repo1.maven.org/maven2/io/spring/javaformat/spring-javaformat-intellij-idea-plugin/" target="_blank">
       spring-javaformat-intellij-idea-plugin</a>
    3. 其他开发工具，请参考 <a href="https://github.com/spring-io/spring-javaformat" target="_blank">
       spring-javaformat</a>
       说明，或`提交代码前`在项目根目录运行下列命令（需要开发者电脑支持`mvn`命令）进行代码格式化
       ```
       mvn spring-javaformat:apply
       ```
   </details>

2. 欢迎提交 [issue](https://github.com/twelvet-projects/twelvet/issues)，请写清楚遇到问题的原因、开发环境、复显步骤。

## 🤝鸣谢

感谢jetbrains提供的许可证[![jetbrains](https://cloud.twelvet.cn/jetbrains.png)](https://www.jetbrains.com?from=https://github.com/twelvet-projects/twelvet)
