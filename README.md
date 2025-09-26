# Upgradelink

## 它是做什么的
UpgradeLink 是**全端支持的应用升级系统与应用升级分发平台**，为应用提供一站式的应用升级及分发解决方案，核心价值包括：
- 有效降低技术门槛
- 减少研发成本
- 助力业务快速搭建稳定高质量的应用

# 介绍

## 1、全端应用升级支持
覆盖主流应用开发框架/系统，提供针对性升级能力：

| 支持类型        | 核心功能描述                                |
|-------------|---------------------------------------|
| 安卓应用升级      | 支持 APK 文件上传与管理，提供专属升级策略，用于应用更新        |
| Tauri 应用升级  | 完全兼容 Tauri 官方升级组件接口，提供标准化升级策略与流程管理    |
| Electron 升级 | 完全兼容 Electron 官方升级组件接口，提供标准化升级策略与流程管理 |


## 2、核心功能模块

### 企业级安全防护（🔐）
- API 服务内置多重安全机制，采用**签名验证、防重放攻击、请求频率限制**等技术
- 保障应用升级过程零风险
### 开箱即用 API-SDK（📦）
| [Go](https://github.com/toolsetlink/upgradelink-api-go) | [Java](https://github.com/toolsetlink/upgradelink-api-java)
| [Dart](https://github.com/toolsetlink/upgradelink-api-dart)
| [Android](https://github.com/toolsetlink/upgradelink-api-android)
| [TypeScript](https://github.com/toolsetlink/upgradelink-api-ts) |

- 支持主流开发语言：Go、Java、Dart、Android、TypeScript
- 优势：提供便捷的 SDK 接入方式，帮助快速实现应用升级功能
### 灵活升级策略（📁）
- 支持维度：设备、机型等
- 管理方式：可通过可视化控制台精准管理升级包分发

## 服务成就数据
- 托管应用：100+
- 托管应用版本：900+
- 加速下载次数：80,000+


# 托管的开源项目
企业版 开源计划-免费提供服务 [计划链接](https://www.toolsetlink.com/upgrade/open-source/plan.html)

| [note-gen (AI笔记软件)](https://notegen.top/en)         | [BongoCat（桌宠）](https://github.com/ayangweb/BongoCat)   | [MarkFlowy (MD 编辑器)](https://github.com/drl990114/MarkFlowy)    | [lazyeat (手势控制)](https://github.com/lanxiuyun/lazyeat)  |


#  快速开始只需4步。

## 1. 环境准备
需要安装[Docker](https://www.docker.com/)。
## 2. 下载项目
### 2.1. 从GitHub下载项目
```shell
git clone https://github.com/toolsetlink/upgradelink
```
## 3.进入项目 启动development目录下的 mysql 与 redis
> 注意：如果有独立的mysql 与 redis 环境 参考 自行build 文档。

### 3.1 启动 mysql

```shell
cd upgradelink/development/mysql-8.4.3
```

```shell
docker-compose up -d
```


### 3.1 启动 redis

```shell
cd upgradelink/development/redis-6.0.20
```

```shell
docker-compose up -d
```

## 4. 启动UpgradeLink

首次执行命令时，会自动下载所需的相关Docker镜像，需要等待的时长取决于网络速度。您也可以提前下载好相关镜像，以缩短执行部署命令的等待时间。

```shell
docker run -d --add-host=host.docker.internal:host-gateway -p 8081:8080 -p 8888:8888 toolsetlink/upgradelink-standalone:v2.0.6
```


# 其他开源项目的快速入门：
- [自行 build 文档](https://www.toolsetlink.com/upgrade/deploy/quick-start-docker2.html)


# 文档
您可以从 [UpgradeLink](https://www.toolsetlink.com/upgrade/) 网站查看完整文档。

所有最新和长期通知也可以在此处找到 [UpgradeLink 通知问题](https://github.com/toolsetlink/upgradelink/issues)。


# 贡献
欢迎贡献者加入 UpgradeLink 项目。请 [进群](https://www.toolsetlink.com/upgrade/communication-group.html) 了解如何为这个项目做出贡献。


# 其他相关项目存储库

## Sdk
- [Go](https://github.com/toolsetlink/upgradelink-api-go)   GO sdk
- [Java](https://github.com/toolsetlink/upgradelink-api-java)   Java sdk
- [Dart](https://github.com/toolsetlink/upgradelink-api-dart)     Dart sdk
- [Android](https://github.com/toolsetlink/upgradelink-api-android)  Android sdk
- [TypeScript](https://github.com/toolsetlink/upgradelink-api-ts) TypeScript sdk




# Contributors

<a href="https://github.com/toolsetlink/upgradelink/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=toolsetlink/upgradelink" />
</a>

# Star History

[![Star History Chart](https://api.star-history.com/svg?repos=toolsetlink/upgradelink&type=Date)](https://www.star-history.com/#toolsetlink/upgradelink&Date)