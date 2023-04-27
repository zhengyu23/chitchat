## 项目简介  
本项目是一个使用Go语言+Go原生包编写的论坛Demo，后端部分使用MVC架构+PostgreSQL数据库，前端部分使用Bootstrap框架，旨在为新手提供一个简洁明了的Go-Web项目实例。

项目拥有论坛、用户、私信三个主要模块，可以实现用户注册、登录、发帖、回帖、私信等功能。代码结构清晰，模块化设计，易于维护和扩展。  

这是我学习Go-Web开发的第一个项目，希望该项目达到适合新手学习+小而精的效果，但目前还存在一些架构、安全和代码冗余问题，后续有空我会进一步改进(找到最适合该项目的状态)，也欢迎大家一同开发。  


### 主要技术栈
- Go编程基础：字面量，常量，变量，类型，操作符，表达式，语句，错误处理
- 数据结构：字符串，数组
- SQL，PostGreSQL
- html，template,Bootstrap框架

## 学习指路
database/sql文档(英):https://pkg.go.dev/database/sql  
net/http文档(英):https://pkg.go.dev/net/http  
html/template文档(英):https://pkg.go.dev/html/template  
log文档(英):https://pkg.go.dev/log  
Bootstrap文档(中):https://v3.bootcss.com/  
Markdown文档(中)：https://markdown.com.cn/basic-syntax/line-breaks.html

书籍参考：《Go Web Programming》，作者郑兆雄(新加坡)，译者黄健宏，2017年12月第1版。  
书籍代码库：https://github.com/sausheong/gwp  
作者郑兆雄Github主页：https://github.com/sausheong  

## 代码结构  
项目的代码结构如下：

`data`：数据操作层(Model)，包括数据的增删改查等操作。  
`templates`：视图层(View)，负责展示数据并与用户进行交互。  
`route`：业务逻辑层(Controller)，负责处理具体的业务逻辑，例如用户注册、登录、商品展示等。  
`static`：静态资源，例如图片、CSS、JavaScript等(暂无)。
`main.go`：代码入口，负责初始化整个应用。  
`utils.go`：通用方法，例如字符串处理、日期格式化等。  
`config.json`：配置文件，存储应用程序的配置信息，例如数据库连接信息、端口号等。 

# 运行方法  
- 克隆本仓库到本地。    
- 安装Go环境。    
- 配置config.json文件。（默认地址为"localhost:8080"）
- 修改data文件中db.go相应配置（已写在注释）。
- 在PostgreSQL中运行data文件中的setup.sql文件，创建数据库表和属性。
- 运行命令go run main.go。  
- 打开浏览器，访问http://localhost:8080。

# 如何贡献
欢迎贡献代码，您可以通过以下方式参与本Demo的开发  
- Fork本仓库到自己的Github账号下。  
- 在自己的仓库中进行代码修改。  
- 提交Pull Request，等待合并。    
`Ps：你可以添加注释，让它更受新手青睐；或者，修改代码，让它更清晰简洁。`

### 开发流程
- 需求分析
- 概要设计
  - 模块设计
  - 数据设计
- 详细设计
  - 模块详细设计
  - 数据详细设计
  - 接口设计
- 编写代码
- 测试

### 主要使用包
- "html/template"
- "net/http"  
- 
#### 创建时间：2023年4月12日16:32:20
#### 完成时间：2023年4月12日01:45:22

### 下一步优化 
(now: 2023年4月12日)
- 异常处理 + 日志打印
- 响应超时
- 登录验证（通过书本中嵌套处理器的方式，但又想到这可能涉及“接口转接”的问题，不知道这个词用得对不对）  

(end：时间xxx)  
(now: 时间xxx)...  



`写在最后: 这个项目我构思了好几天，期间我总想着如何让代码结构更加抽象，让设计更加抽象，但总觉得还差点什么导致代码不够简洁（所以推翻了3版的代码架构设计，文档后续我会补上来），最终是参考了书籍代码库里的代码架构，从中得到启发并顺利完成。我发现自己构思的时候总会花去大量时间，但其实别人代码已足够优秀，往往可从阅读别人代码中获取灵感而不是一昧地独自思考，这是我做得不好的地方。同时，这个项目也让我对go基本知识更加掌握，学习了template包的用法，熟悉了SQL语法和PostGreSQL语法，也弥补了一点不会写、看前端和不会写Markdown的缺点。`