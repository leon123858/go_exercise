# data-api

## test

note: should turn on local testing DB

```
make test
```

## deploy

0. `make deploy` 建置和發佈
1. deploy cloud run by registry in cloud console
2. should set env in cloud run named "postgresPassword" which have db password
3. click 持續部署 in cloud run console, target is `/data-api/Dockerfile` (many permission should be set)
4. go to cloud build and set "手動觸發" (之後可以到 cloud build 手動 CD)

## swagger 語法

```
@Summary 是对该接口的一个描述
@Id 是一个全局标识符，所有的接口文档中 Id 不能标注
@Tags 是对接口的标注，同一个 tag 为一组，这样方便我们整理接口
@Security ApiKeyAuth 表示这是一个需要认证才可以调用的接口，对应// @securityDefinitions.apikey ApiKeyAuth
@Version 表明该接口的版本
@Accept 表示该该请求的请求类型
@Param 表示参数 分别有以下参数 参数名词 参数类型 数据类型 是否必须 注释 属性(可选参数),参数之间用空格隔开。
@Success 表示请求成功后返回，它有以下参数 请求返回状态码，参数类型，数据类型，注释
@Failure 请求失败后返回，参数同上
@Router 该函数定义了请求路由并且包含路由的请求方式。
```
