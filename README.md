# fuwu
简单服务管理（fuwu five 555）

## 功能

- start serverName
- stop serverName (same close)
- status serverName
- restart serverName (close and start)
- list (show all serverName and status is running)
- refresh (read fuwu.yml and refresh server list)

- server (-log) start daemon server

## 核心服务

- http 守护进程服务 (核心服务失败重启)
- 开放端口 555 提供 api 操作
- 服务跟随启动
- 服务失败自动重启
- 服务日志管理
- 服务配置热更新
- （服务日志定期删除）
- （简单管理页面）

> (xxx) 待定功能

