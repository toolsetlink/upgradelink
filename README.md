# UpgradeLink 🚀

全端应用升级与分发平台 | 为开发者省去90%升级服务搭建成本

[![GitHub Repo stars](https://img.shields.io/github/stars/toolsetlink/upgradelink)](https://github.com/toolsetlink/upgradelink)
[![stars](https://gitcode.com/toolsetlink/upgradelink/star/badge.svg)](https://gitcode.com/toolsetlink/upgradelink)
[![stars](https://gitee.com/toolsetlink/upgradelink/badge/star.svg)](https://gitee.com/toolsetlink/upgradelink)

[English](README_en.md) | [中文](README.md) | [日本語](README_ja.md)


## 为什么选择 UpgradeLink？
无论是个人开发者的小工具，还是企业级复杂应用，都能一键接入全端升级能力，无需从零搭建服务：
- 🚀 **开箱即用**：无需深耕升级逻辑，SDK接入10分钟完成
- 🌍 **全端覆盖**：Windows/Linux/Mac/安卓/Tauri/Electron等全场景支持
- 💰 **成本锐减**：省去服务器搭建、跨端适配的研发投入
- 🔒 **企业级可靠**：安全防护+灵活策略，支持灰度发布/定向升级
- 📦 **灵活存储支持**：支持S3协议对象存储，可自定义CDN下载地址


## 核心功能：全场景升级解决方案
### 1. 多平台/框架专属升级
| 支持类型         | 核心能力                          | 适用场景                  |
|------------------|-----------------------------------|---------------------------|
| Windows 应用     | 专属升级策略管理                  | PC桌面应用                |
| Linux 应用       | 适配服务器/桌面端升级需求         | 服务器工具、Linux桌面应用 |
| Mac 应用         | 符合苹果生态的升级流程            | Mac桌面应用               |
| 安卓应用         | 支持差分升级（增量更新）          | 手机APP                   |
| Tauri 应用       | 兼容官方组件接口，无缝升级        | Tauri跨端应用             |
| Electron 应用    | 兼容官方升级接口，零改造接入      | Electron跨端应用          |

### 2. 灵活资源升级
| 升级类型   | 功能描述                          | 适用场景                  |
|------------|-----------------------------------|---------------------------|
| 配置升级   | 自定义JSON配置，支持在线动态更新  | 需要动态调整参数的应用    |
| 文件升级   | 上传任意文件，支持资源补充/替换   | 需更新素材、插件的应用      |
| URL升级    | 自定义下载链接，自行维护文件存储  | 已有云存储服务的应用      |

### 3. 灵活存储与分发
- 📦 **S3协议支持**：所有应用版本文件均可存储在支持S3协议的对象存储服务中
- 🌐 **自定义CDN**：支持配置自定义CDN下载地址，优化全球用户下载体验


## 10分钟接入指南 🔧
### 第一步：选对应SDK
支持6种主流语言，无需复杂配置：
- 🐹 [Go SDK](https://github.com/toolsetlink/upgradelink-api-go)
- ☕ [Java SDK](https://github.com/toolsetlink/upgradelink-api-java)
- 🐍 [Python SDK](https://github.com/toolsetlink/upgradelink-api-python)
- 🦋 [Dart SDK](https://github.com/toolsetlink/upgradelink-api-dart)
- 🤖 [Android SDK](https://github.com/toolsetlink/upgradelink-api-android)
- 🟦 [TypeScript SDK](https://github.com/toolsetlink/upgradelink-api-typescript)
- 🟣 [C# SDK](https://github.com/toolsetlink/upgradelink-api-csharp)

### 第二步：参考文档
- [完整使用文档](https://www.toolsetlink.com/upgrade/what-is-upgrade.html)
- [快速接入教程](https://www.toolsetlink.com/upgrade/deploy/quick-start-docker.html)


## 部署方式
支持多种部署方案，按需选择：
- [单机快速启动（Docker）](https://www.toolsetlink.com/upgrade/deploy/quick-start-docker2.html)
- [Docker Compose 部署](https://www.toolsetlink.com/upgrade/deploy/docker-compose.html)
- [集群化部署](https://www.toolsetlink.com/upgrade/deploy/cluster-docker.html)


## 谁在使用？

已有多个开源项目接入，按技术框架分类如下：

### 🚀 Tauri 框架项目
| 项目名称 | 应用场景 |                                                            Stars                                                            | 
| :--- | :--- |:---------------------------------------------------------------------------------------------------------------------------:| 
| note-gen | 跨平台笔记工具 |       [![GitHub Repo stars](https://img.shields.io/github/stars/codexu/note-gen)](https://github.com/codexu/note-gen)       | 
| BongoCat | 趣味桌面宠物 |     [![GitHub Repo stars](https://img.shields.io/github/stars/ayangweb/BongoCat)](https://github.com/ayangweb/BongoCat)     | 
| Hula | 开源即时通讯系统 |        [![GitHub Repo stars](https://img.shields.io/github/stars/HuLaSpark/HuLa)](https://github.com/HuLaSpark/HuLa)        | 
| MarkFlowy | 轻量 Markdown 工具 |   [![GitHub Repo stars](https://img.shields.io/github/stars/drl990114/MarkFlowy)](https://github.com/drl990114/MarkFlowy)   |
| lazyeat | 手势控制工具 |     [![GitHub Repo stars](https://img.shields.io/github/stars/lanxiuyun/lazyeat)](https://github.com/lanxiuyun/lazyeat)     | 
| Prompt-Tools | AI 提示词工具 | [![GitHub Repo stars](https://img.shields.io/github/stars/jwangkun/Prompt-Tools)](https://github.com/jwangkun/Prompt-Tools) | 
| MarkFly | Markdown 编辑器 |      [![GitHub Repo stars](https://img.shields.io/github/stars/jwangkun/MarkFly)](https://github.com/jwangkun/MarkFly)      | 
| rank-analysis | LoL 排位分析工具 |  [![GitHub Repo stars](https://img.shields.io/github/stars/wnzzer/rank-analysis)](https://github.com/wnzzer/rank-analysis)  |
| Welight | 公众号智能排版应用 |                                                   [官网](https://waer.ltd)                                                    |  

---

### 📱 Android 项目
| 项目名称 | 应用场景 |                                                             Stars                                                             |
| :--- | :--- |:-----------------------------------------------------------------------------------------------------------------------------:|
| 聚在工大 | 校园服务 | [![GitHub Repo stars](https://img.shields.io/github/stars/Chiu-xaH/HFUT-Schedule)](https://github.com/Chiu-xaH/HFUT-Schedule) |
| 脉扑 | 健康生活 |                                                                [官网](https://www.maipusc.com)                                                                  |

---

### 💻 Electron 框架项目
| 项目名称 | 应用场景 | Stars |
| :--- | :--- |:-----:|
| es-client | ElasticSearch 客户端 |   [官网](https://es-client.esion.xyz)    |



## 企业版
免费提供企业级服务，免部署，详情见 [企业版免费使用](https://www.toolsetlink.com/upgrade/company/company.html)


## 贡献指南
欢迎加入开发！通过 [官方交流群](https://www.toolsetlink.com/upgrade/communication-group.html) 了解贡献流程。

> 技术栈：基于 [Go Zero](https://go-zero.dev/) 与 [simple-admin](https://doc.ryansu.tech/) 开发


## 通知与反馈
- 最新动态：[UpgradeLink 通知](https://github.com/toolsetlink/upgradelink/issues)
- 问题反馈：[联系我们](https://www.toolsetlink.com/upgrade/contact-us.html)


## 贡献者
<a href="https://github.com/toolsetlink/upgradelink/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=toolsetlink/upgradelink" />
</a>

# Star History

[![Star History Chart](https://api.star-history.com/svg?repos=toolsetlink/upgradelink&type=Date)](https://www.star-history.com/#toolsetlink/upgradelink&Date)