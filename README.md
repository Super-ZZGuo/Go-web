# Gee-web

> 一个用 Go 语言实现一个简单的 Web 框架，设计思路基于[Gin](https://github.com/gin-gonic/gin)

功能实现较为简单，实现了一个Web框架最为简单基础的一些功能：

- 构建Context(上下文)来**解析请求**、**快速构建常用respone(响应)**
- 使用Trie(前缀树)来进行**路由匹配**
- 实现同一实例下不同路由的**分组控制**，便于**中间件**对于对应的分组路由的**功能加强**

> 受限于一些原因，无法使用Go完成一个完整的项目
> 
> 故希望能通过该项目来提高自己对于GO的认识


