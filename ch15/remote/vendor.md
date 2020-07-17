### Go 未解决的依赖问题
1. 同一环境下，不同项目使用同一包的不同版本
2. 无法管理对包的特定版本的依赖

### vendor 路径
查找依赖包路径的解决方案如下：
1. 当前包下的 vendor 目录
2. 向上级目录查找，直到找到 src 下的 vendor 目录
3. 在 GOPATH 下面查找依赖包
4. 在 GOROOT 目录下查找


### 常用的依赖管理工具
godep https://github.com/tools/godep
glide https://github.com/Masterminds/glide
dep https://github.com/golang/dep

### windows平台使用glide

1. 下载安装glide

go get github.com/Masterminds/glide

glide的源码以及exe文件在第一个gopath目录，如果不知道哪个是第一个gopath，echo一下

echo %GOPATH%

2. 把glide.exe加入系统环境变量path目录

建议直接把glide.exe拷贝到GO安装目录/bin

3. 进入自己的go work工程

glide init//初始化包依赖配置 # Y->Y->N->S->N->Y 生成glide.yaml文件

glide get github.com/streadway/amqp //我这里是安装rabbitmq,自行按需修改

4. Unable to export dependencies to vendor directory

在windows平台glide get的时候，总是提示如下错误：

```
D:\goopen\src>glide get github.com/streadway/amqp
···
[ERROR] Unable to export dependencies to vendor directory: Error moving files: e

xit status 1. output:

0
```
这个bug只在windows平台有
[github pr](https://github.com/Masterminds/glide/pull/889/commits/cc37dc711a3191c2b91b01b9593c685660eeb9af)

5.rebuild glide
按照这个pr修改glide源码，进入glide目录，go build，然后把重新生成的glide拷贝到go安装目录/bin

### mac下使用 glide
```
brew install glide
glide init  # Y->Y->N->S->N->Y 生成glide.yaml文件
glide install
```

### glide 简单使用

1. 导一个包测试
```
glide get github.com/mattn/go-adodb
```
2. 导入单个包源
```
glide get --all-dependencies -s -v github.com/mattn/go-adodb
```
3. 根据指定版本号导入项目，如
```
glide get github.com/go-sql-driver/mysql#v1.2
```