# gins

基于Gin 进行常用第三方插件封装的框架(Mysql, Redis, Viper, Zap等)，保持它的轻量、性能

直接上手业务，无需过多关注基础建设

## 支持Linux一件Docker启动
内存 > 16
需要自行安装docker、docker-compose
```shell
# mvn
mvn clean && mvn install
# 进入脚本目录
cd ./docker
# 可执行权限
chmod 751 deploy.sh
# 执行启动（按需执行参数，[init|port|base|server|stop|rm]）
# 基础服务
./deploy.sh base
# 启动gins
./deploy.sh server
```

## TwelveT交流群

QQ群： [![加入QQ群](https://img.shields.io/badge/985830229-blue.svg)](https://jq.qq.com/?_wv=1027&k=cznM6Q00) 点击按钮入群。
