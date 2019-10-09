---
title: day_10
---

### 状态码
- 100 请求确认
- 200 请求成功
- 302 重定向
- 400 异常请求
- 401 没有权限
- 403 禁止访问
- 404 寻找不到资源
- 500 内部服务器错误

### panic处理
- 封装一个函数先做 异常处理，然后执行业务函数


### 通知
- 主线程通知子线程退出 使用context
- 子线程通知主线程退出 使用sync.WaitGroup


### 事务 ACID
- A 原子性
- C 一致性
- I 隔离性
- D 持久性

### MySQL事务
- import sqlx
- Db.Begin() 开始事务
- Db.Submit() 提交事务
- Db.Rollback() 回滚事务
