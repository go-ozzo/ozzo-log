# ozzo-log

[![GoDoc](https://godoc.org/github.com/go-ozzo/ozzo-log?status.png)](http://godoc.org/github.com/go-ozzo/ozzo-log)
[![Build Status](https://travis-ci.org/go-ozzo/ozzo-log.svg?branch=master)](https://travis-ci.org/go-ozzo/ozzo-log)
[![Coverage](http://gocover.io/_badge/github.com/go-ozzo/ozzo-log)](http://gocover.io/github.com/go-ozzo/ozzo-log)

## 其他语言

[English](../README.md) [Русский](/docs/README-ru.md)

## 说明

ozzo-log is a Go package providing enhanced logging support for Go programs. It has the following features:<br>
ozzo-log 是给 Go 程序提供加强型日志功能的 go 包。它支持以下功能：

* High performance through asynchronous logging;
* 通过异步记录实现的高性能
* Recording message severity levels;
* 记录消息的严重级别(severity level)
* Recording message categories;
* 记录消息的分类
* Recording message call stacks;
* 记录错误信息的调用栈(call stack)
* Filtering via severity levels and categories;
* 根据严重级别和消息分类进行过滤分流
* Customizable message format;
* 可定制的信息格式
* Configurable and pluggable message handling through log targets;
* 通过日志目的地(log target)对象，实现可配置、可接入的信息处理过程。
* Included console, file, network, and email log targets.
* 包括 console、file、network 以及 email 多种日志目的地

## 需求

Go 1.2 或以上。

## 安装

Run the following command to install the package:<br>
执行以下指令安装此包：

```
go get github.com/go-ozzo/ozzo-log
```

## 准备开始

The following code snippet shows how you can use this package.<br>
以下代码片段展示了如何使用此包：

```go
package main

import (
	"github.com/go-ozzo/ozzo-log"
)

func main() {
    // creates the root logger
    // 创建根记录器(root logger)
	logger := log.NewLogger()

	// adds a console target and a file target
	// 添加一个控制台目的地和一个文件目的地
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = "app.log"
	t2.MaxLevel = log.LevelError
	logger.Targets = append(logger.Targets, t1, t2)

	logger.Open()
	defer logger.Close()

	// calls log methods to log various log messages
	// 调用不同的记录方法，随便记录一些日志信息。
	logger.Error("plain text error")
	logger.Error("error with format: %v", true)
	logger.Debug("some debug info")

	// customizes log category
	// 自定义日志类别
	l := logger.GetLogger("app.services")
	l.Info("some info")
	l.Warning("some warning")

	...
}
```

## 记录器与目的地 (Logger and Targets)

A logger provides various log methods that can be called by application code
to record messages of various severity levels.<br>
记录器提供了多种不同的日志方法。应用代码可以用它们来记录不同严重级别的信息。

A target filters log messages by their severity levels and message categories
and processes the filtered messages in various ways, such as saving them in files,
sending them in emails, etc.<br>
一个目的地对象会通过这些消息的严重级别和分类对他们进行过滤，以不同的方式对过滤后的数据进行处理，比如写入文件或者以邮件形式发送等等。

A logger can be equipped with multiple targets each with different filtering conditions.<br>
记录器可以装配不同的目的地对象，并给他们应用不同的过滤规则。

The following targets are included in the ozzo-log package.<br>
以下过滤器乃 ozzo-log 原生自带：

* `ConsoleTarget`: displays filtered messages to console window
* `ConsoleTarget`：于控制台窗口内显示过滤后的信息
* `FileTarget`: saves filtered messages in a file (supporting file rotating)
* `FileTarget`：将过滤后的消息写入文件（支持文件分页）
* `NetworkTarget`: sends filtered messages to an address on a network
* `NetworkTarget`：将过滤后的信息发送与网络内的某一地址
* `MailTarget`: sends filtered messages in emails
* `MailTarget`：把过滤后的信息以邮件形式发送

You can create a logger, configure its targets, and start to use logger with the following code:<br>
你可以创建一个记录器，配置好它的消息目的地，然后就可以使用如下方法进行日志记录了：

```go
// creates the root logger
// 创建根记录器
logger := log.NewLogger()
logger.Targets = append(logger.Targets, target1, target2, ...)
logger.Open()
...于此处调用日志方法...
logger.Close()
```

## 严重级别 (Severity Levels)

You can log a message of a particular severity level (following the RFC5424 standard)
by calling one of the following methods of the `Logger` struct:<br>
记录消息时，可以通过调用  `Logger` 结构的特定方法标明特定的严重级别（级别的规定符合 RFC5424 标准）

* `Emergency()`: the system is unusable.
* `Emergency()`：系统已无法使用
* `Alert()`: action must be taken immediately.
* `Alert()`：必须立刻采取措施
* `Critical()`: critical conditions.
* `Critical()`：危笃状态
* `Error()`: error conditions.
* `Error()`：错误情形
* `Warning()`: warning conditions.
* `Warning()`：警告情形
* `Notice()`: normal but significant conditions.
* `Notice()`：正常但重要的情形
* `Info()`: informational purpose.
* `Info()`：用于记录信息
* `Debug()`: debugging purpose.
* `Debug()`：用于调试

## 信息分类 (Message Categories)

Each log message is associated with a category which can be used to group messages.
For example, you may use the same category for messages logged by the same Go package.
This will allow you to selectively send messages to different targets.<br>
每一条日志消息都会关联有一个用于信息分组的类别。比如，你可以给来自同一个 Go 包的日志记录相同的类别标识。之后就可以有选择地吧信息发送到不同的目的地。

When you call `log.NewLogger()`, a root logger is returned which logs messages using
the category named as `app`. To log messages with a different category, call the `GetLogger()`
method of the root logger or a parent logger to get a child logger and then call its
log methods:<br>
调用 `log.NewLogger()` 方法会返回一个设定为 `app`（默认值）类型的根记录器。要使用不同的类别记录，则可以调用某根记录器或父记录器的 `GetLogger()` 方法，从而获得一个不同类别的子记录器。

```go
logger := log.NewLogger()
// 消息归类于 "app"
logger.Error("...")

l1 := logger.GetLogger("system")
// 消息归类于 "system"
l1.Error("...")

l2 := l1.GetLogger("app.models")
// 消息归类于 "app.models"
l2.Error("...")
```

## 信息格式 (Message Formatting)

By default, each log message takes this format when being sent to different targets:<br>
默认情况下，发送给不同目的地的消息均会使用以下缺省格式：

```
2015-10-22T08:39:28-04:00 [Error][app.models] something is wrong
...调用栈（如果启用）...
```

You may customize the message format by specifying your own message formatter when calling
`Logger.GetLogger()`. For example,<br>
在调用 `Logger.GetLogger()` 时，可以通过指定你自己的格式化器来自定义信息格式。比如：

```go
logger := log.NewLogger()
logger = logger.GetLogger("app", func (l *Logger, e *Entry) string {
    return fmt.Sprintf("%v [%v][%v] %v%v", e.Time.Format(time.RFC822Z), e.Level, e.Category, e.Message, e.CallStack)
})
```


## 记录调用栈

By setting `Logger.CallStackDepth` as a positive number, it is possible to record call stack information for
each log method call. You may further configure `Logger.CallStackFilter` so that only call stack frames containing
the specified substring will be recorded. For example,<br>
通过给 `Logger.CallStackDepth` 设置一个正整数，为记录下每次日志方法调用的调用栈信息是可能的。也可以再进一步，通过配置 `Logger.CallStackFilter`，让其只记录包含特定子字符串的调用栈栈帧。以下是示例：

```go
logger := log.NewLogger()
// record call stacks containing "myapp/src" up to 5 frames per log message
// 记录 "myapp/src" 的调用栈，最深可达每消息 5 栈帧
logger.CallStackDepth = 5
logger.CallStackFilter = "myapp/src"
```

## 信息过滤 (Message Filtering)

By default, messages of *all* severity levels will be recorded. You may customize
`Logger.MaxLevel` to change this behavior. For example,<br>
默认，**所有**的严重级别都会被记录，但是你可以修改 `Logger.MaxLevel` 从而改变默认的行为。如：

```go
logger := log.NewLogger()
// only record messages between Emergency and Warning levels
// 只记录 Emergency（紧急） 和 Warning（警告） 级别的消息
logger.MaxLevel = log.LevelWarning
```

Besides filtering messages at the logger level, a finer grained message filtering can be done
at target level. For each target, you can specify its `MaxLevel` similar to that with the logger;
you can also specify which categories of the messages the target should handle. For example,<br>
除了在记录器层级进行过滤之外，也可以通过日志目的地层级进行更加细粒度地过滤。对于每一个目的地，可以单独指定它的 `MaxLevel`，和在记录器层级用法类似；你也能指定各个目的地应该处理哪些消息类别。举个栗子解释一下：

```go
target := log.NewConsoleTarget()
// handle messages between Emergency and Info levels
// 此目的地会处理 Emergency 级别和 Info 级别之间的消息
target.MaxLevel = log.LevelInfo
// handle messages of categories which start with "system.db." or "app."
// 处理所有消息分类以 "system.db." 或 "app." 开头的消息
target.Categories = []string{"system.db.*", "app.*"}
```

## 配置记录器

When an application is deployed for production, a common need is to allow changing
the logging configuration of the application without recompiling its source code.
ozzo-log is designed with this in mind.<br>
当应用被部署于生产环境时，人们会有一个非常常见的需求，也就是对记录器进行配置，而无需重新编译源代码。ozzo-log 在设计时就包含了这个考量。

For example, you can use a JSON file to specify how the application and its
logger should be configured:<br>
具体来说，你可以用一个 JSON 文件来配置日志器：

```
{
    "Logger": {
        "Targets": [
            {
                "type": "ConsoleTarget",
            },
            {
                "type": "FileTarget",
                "FileName": "app.log",
                "MaxLevel": 4   // Warning or above
            }
        ]
    }
}
```

Assuming the JSON file is `app.json`, in your application code you can use the `ozzo-config` package
to load the JSON file and configure the logger used by the application:<br>
打个比方，若你的 JSON 文件名叫 `app.json`，我们便可通过 `ozzo-config` 包实现对 JSON文件的导入和对日志记录器的配置：

```go
package main

import (
	"github.com/go-ozzo/ozzo-config"
    "github.com/go-ozzo/ozzo-log"
)

func main() {
    c := config.New()
    c.Load("app.json")
    // register the target types to allow configuring Logger.Targets.
    // 注册供 Logger.Targets 使用的日志目的地类型
    c.Register("ConsoleTarget", log.NewConsoleTarget)
    c.Register("FileTarget", log.NewFileTarget)

    logger := log.NewLogger()
    if err := c.Configure(logger, "Logger"); err != nil {
        panic(err)
    }
}
```

To change the logger configuration, simply modify the JSON file without
recompiling the Go source files.<br>
要修改 logger 的配置，只需修改 JSON 文件就好了，无需重新编译 Go 的源文件。