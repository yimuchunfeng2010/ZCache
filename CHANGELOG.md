# Change Log

A light Cache writed by golang

# [version:0.0.1]
### Features
初始框架

# [version:0.0.2]
### Features
实现基本Get/Set接口


# [version:0.0.3]
### Features
实现基本Update/Delete接口

# [version:0.0.4]
### Features
更改核心数据结构 由 二叉树 改为 hash表+二叉树

# [version:0.0.5]
### Features
增加读写锁

# [version:0.0.6]
### Features
增加测试接口

# [version:0.0.7]
### Features
增加查询所有数据的接口

# [version:0.0.8]
### Features
增加数据存库接口

# [version:0.0.9]
### Features
增加数据文本导入接口

# [version:0.1.0]
### Features
增加发送邮件相关代码

# [version:0.1.1]
### Features
增加对集群健康度的校验

# [version:0.1.2]
### Features
增加系统退出信号处理，保存数据(功能待放开)

# [version:0.1.3]
### Features
增加全局key个数计数功能

# [version:0.1.4]
### Features
增加值增加/减少功能

# [version:0.1.5]
### Features
增加删除全部缓存功能(删除前先存库)

# [version:0.1.6]
### Fix
修复无法获取全部key/value问题

# [version:0.1.7]
### Features
增加扩容功能

# [version:0.1.8]
### Features
增加redis导入导出功能

# [version:0.1.9]
### Features
编写dockerfile文件

# [version:0.2.0]
### Features
增加系统定时任务框架

# [version:0.2.1]
### Features
增加日志系统缓存功能

# [version:0.2.2]
### Features
增加基于zookeeper的分布式读写锁功能

# [version:0.2.3]
### Features
增加基于zookeeper的集群监控功能

# [version:0.2.4]
### Features
增加集群亚健康/不健康时的数据备份功能

# [version:0.2.5]
### Features
增加集群数据Hash值的计算

# [version:0.2.6]
### Features
实现部分客户端接口功能

# [version:0.2.7]
### Fix
暂时屏蔽Zookeeper功能及修改客户端接口实现

# [version:0.2.8]
### Features
增加内部增删改接口

# [version:0.2.9]
### Fix
修改增删改接口以实现内部集群修改操作

# [version:0.3.0]
### Features
增加提交和撤销任务接口

# [version:0.3.1]
### Features
增加定时清理超时Job

# [version:0.3.2]
### Features
添加.gitignore文件

# [version:0.3.3]
### Features
引入grpc框架，实现部分rpc接口

# [version:0.3.4]
### Features
实现grpc服务端接口

# [version:0.3.5]
### Features
业务接入grpc