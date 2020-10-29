# 简介

本项目是利用钉钉的企业内部机器人且基于gin框架编写的简单骰娘机器人~~（只会投骰子而已）~~

如果你的技术栈恰好是Golang，那么你可以很容易的就能启动起此项目，由于`Golang`已有的包没有实现一个`d`运算符，因此自己按照网上的代码写了个`AST`语法解析器，能够满足基本的投骰子需求

# 配置

## 代码配置

本项目不需要配置太多，代码层面只需要改这三个配置即可:

`Utils/BotStructUtils.go`

```go
var Userconfig = map[string]string{
	"AppKey": "xxxxx", //这里填钉钉开发者后台给出的APPsecret
	"secretkey": "xxx", // 这里填机器人界面下的密钥SE开头的
	"webhook": "xxx", //这里填机器人的webhook地址
}
```

只需要将这三个值改好就行，默认启动的端口为`45678`，如果需要更改请在`build`前改`main.go`中的内容即可:

`main.go`

```go
func main() {
	r := router.GetRouters()
	r.Run("0.0.0.0:45678")
	//如果需要改到其他端口改成0.0.0.0:Port即可
}
```

## 钉钉配置

1. 首先创建一个企业~~(亲友群)~~，然后按照下图创建一个机器人

![](https://qiniu.ebounce.cn/blog/20201029194809.png?imageView2/0/q/75%7Cwatermark/2/text/QEVib3VuY2U=/font/5b6u6L2v6ZuF6buR/fontsize/460/fill/IzlFOTc5Nw==/dissolve/50/gravity/SouthEast/dx/10/dy/10%7Cimageslim)

2. 钉钉中更改安全设置：

   ![](https://qiniu.ebounce.cn/blog/20201029195011.png?imageView2/0/q/75%7Cwatermark/2/text/QEVib3VuY2U=/font/5b6u6L2v6ZuF6buR/fontsize/460/fill/IzlFOTc5Nw==/dissolve/50/gravity/SouthEast/dx/10/dy/10%7Cimageslim)

   这里加签中的内容填到代码里的`secretkey`中

3. 找到`webhook`地址：



![](https://qiniu.ebounce.cn/blog/20201029195133.png?imageView2/0/q/75%7Cwatermark/2/text/QEVib3VuY2U=/font/5b6u6L2v6ZuF6buR/fontsize/460/fill/IzlFOTc5Nw==/dissolve/50/gravity/SouthEast/dx/10/dy/10%7Cimageslim)

这里`webhook`的值填到代码中的`webhook`即可

4.填写接收信息的服务器地址：

![](https://qiniu.ebounce.cn/blog/20201029195450.png?imageView2/0/q/75%7Cwatermark/2/text/QEVib3VuY2U=/font/5b6u6L2v6ZuF6buR/fontsize/460/fill/IzlFOTc5Nw==/dissolve/50/gravity/SouthEast/dx/10/dy/10%7Cimageslim)

这样钉钉就算配置好了

## 构建项目并启动

1. 请保证本地有go语言环境，具体可以参考 [Go安装](https://studygolang.com/dl)

2. 本项目使用go mod搭建，所以请使用go mod 参考 [goproxy.io](https://goproxy.io/zh/)

3. 最后到你服务器的对应目录`main.go`的对应目录，运行

```shell
# linux or mac
go build main.go
./main & #保证在root权限下运行
# window
go build main.go
main.exe
```

最后效果图大概是这样：

![](https://qiniu.ebounce.cn/blog/20201029200159.png?imageView2/0/q/75%7Cwatermark/2/text/QEVib3VuY2U=/font/5b6u6L2v6ZuF6buR/fontsize/460/fill/IzlFOTc5Nw==/dissolve/50/gravity/SouthEast/dx/10/dy/10%7Cimageslim)

