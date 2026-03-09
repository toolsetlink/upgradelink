# UpgradeLink 🚀

Cross-Platform App Upgrade & Distribution Platform | Save 90% of the cost in building upgrade services for developers

[![GitHub Repo stars](https://img.shields.io/github/stars/toolsetlink/upgradelink)](https://github.com/toolsetlink/upgradelink)
[![stars](https://gitcode.com/toolsetlink/upgradelink/star/badge.svg)](https://gitcode.com/toolsetlink/upgradelink)
[![stars](https://gitee.com/toolsetlink/upgradelink/badge/star.svg)](https://gitee.com/toolsetlink/upgradelink)

[English](README.md) | [中文](README_zh.md) | [日本語](README_ja.md) | [Português](README_pt.md) | [Русский](README_ru.md) | [Español](README_es.md) | [한국어](README_ko.md) | [فارسی](README_fa.md)


## Why Choose UpgradeLink?
Whether it's a small tool by individual developers or a complex enterprise-level application, you can instantly integrate full-stack upgrade capabilities without building services from scratch:
- 🚀 **Out-of-the-Box**: No need to dive deep into upgrade logic—SDK integration takes only 10 minutes
- 🌍 **Cross-Platform Coverage**: Supports Windows/Linux/Mac/Android/Tauri/Electron and other full scenarios
- 💰 **Cost Reduction**: Eliminate R&D investment in server setup and cross-platform adaptation
- 🔒 **Enterprise-Grade Reliability**: Security protection + flexible strategies, supporting canary release/targeted upgrade
- 📦 **Flexible Storage Support**: Supports S3 protocol object storage, customizable CDN download address


## Core Features: Full-Scenario Upgrade Solutions
### 1. Platform/Framework-Specific Upgrades
| Supported Type       | Core Capabilities                          | Use Cases                  |
|----------------------|---------------------------------------------|----------------------------|
| Windows Apps         | Dedicated upgrade strategy management       | PC desktop applications    |
| Linux Apps           | Adapt to server/desktop upgrade requirements| Server tools, Linux desktop apps |
| Mac Apps             | Apple ecosystem-compliant upgrade workflow  | Mac desktop applications   |
| Android Apps         | Supports delta updates (incremental updates)| Mobile apps                |
| Tauri Apps           | Compatible with official component APIs for seamless upgrades | Tauri cross-platform apps |
| Electron Apps        | Compatible with official upgrade APIs, zero-modification integration | Electron cross-platform apps |

### 2. Flexible Resource Upgrades
| Upgrade Type         | Feature Description                          | Use Cases                  |
|----------------------|---------------------------------------------|----------------------------|
| Configuration Upgrade| Custom JSON configuration with online dynamic updates | Apps requiring dynamic parameter adjustments |
| File Upgrade         | Upload any files to support resource supplement/replacement | Apps needing to update materials or plugins |
| URL Upgrade          | Custom download links with self-managed file storage | Apps with existing cloud storage services |

### 3. Flexible Storage and Distribution
- 📦 **S3 Protocol Support**: All application version files can be stored in S3-compatible object storage services
- 🌐 **Customizable CDN**: Supports configuring custom CDN download addresses to optimize global user download experience


## 10-Minute Integration Guide 🔧
### Step 1: Select the Corresponding SDK
Supports 6 mainstream languages with no complex configuration:
- 🐹 [Go SDK](https://github.com/toolsetlink/upgradelink-api-go)
- ☕ [Java SDK](https://github.com/toolsetlink/upgradelink-api-java)
- 🐍 [Python SDK](https://github.com/toolsetlink/upgradelink-api-python)
- 🦋 [Dart SDK](https://github.com/toolsetlink/upgradelink-api-dart)
- 🤖 [Android SDK](https://github.com/toolsetlink/upgradelink-api-android)
- 🟦 [TypeScript SDK](https://github.com/toolsetlink/upgradelink-api-typescript)
- 🟣 [C# SDK](https://github.com/toolsetlink/upgradelink-api-csharp)

### Step 2: Refer to Documentation
- [Complete User Guide](https://www.toolsetlink.com/en/upgrade/what-is-upgrade.html)
- [Quick Start Tutorial](https://www.toolsetlink.com/en/upgrade/deploy/quick-start-docker.html)


## Deployment Methods
Supports multiple deployment options to meet your needs:
- [Standalone Quick Start (Docker)](https://www.toolsetlink.com/en/upgrade/deploy/quick-start-docker2.html)
- [Docker Compose Deployment](https://www.toolsetlink.com/en/upgrade/deploy/docker-compose.html)
  - [Clustered Deployment](https://www.toolsetlink.com/en/upgrade/deploy/cluster-docker.html)


## Who's Using?

Multiple open-source projects have been integrated, categorized by technical framework as follows:

### 🚀 Tauri Framework Projects
| Project Name | Use Case | Stars |
| :--- | :--- |:---------------------------------------------------------------------------------------------------------------------------:|
| note-gen | Cross-platform note-taking tool | [![GitHub Repo stars](https://img.shields.io/github/stars/codexu/note-gen)](https://github.com/codexu/note-gen) |
| BongoCat | Fun desktop pet | [![GitHub Repo stars](https://img.shields.io/github/stars/ayangweb/BongoCat)](https://github.com/ayangweb/BongoCat) |
| Hula | Open-source instant messaging system | [![GitHub Repo stars](https://img.shields.io/github/stars/HuLaSpark/HuLa)](https://github.com/HuLaSpark/HuLa) |
| MarkFlowy | Lightweight Markdown tool | [![GitHub Repo stars](https://img.shields.io/github/stars/drl990114/MarkFlowy)](https://github.com/drl990114/MarkFlowy) |
| lazyeat | Gesture control tool | [![GitHub Repo stars](https://img.shields.io/github/stars/lanxiuyun/lazyeat)](https://github.com/lanxiuyun/lazyeat) |
| Prompt-Tools | AI prompt tool | [![GitHub Repo stars](https://img.shields.io/github/stars/jwangkun/Prompt-Tools)](https://github.com/jwangkun/Prompt-Tools) |
| MarkFly | Markdown editor | [![GitHub Repo stars](https://img.shields.io/github/stars/jwangkun/MarkFly)](https://github.com/jwangkun/MarkFly) |
| rank-analysis | LoL ranking analysis tool | [![GitHub Repo stars](https://img.shields.io/github/stars/wnzzer/rank-analysis)](https://github.com/wnzzer/rank-analysis) |
| Welight | WeChat Official Account intelligent layout application | [Official Website](https://waer.ltd) |

---

### 📱 Android Projects
| Project Name | Use Case |                                                             Stars                                                             |
| :--- | :--- |:-----------------------------------------------------------------------------------------------------------------------------:|
| 聚在工大 | Campus service | [![GitHub Repo stars](https://img.shields.io/github/stars/Chiu-xaH/HFUT-Schedule)](https://github.com/Chiu-xaH/HFUT-Schedule) |
| 脉扑 | Health and wellness |                                                               [Official Website](https://www.maipusc.com)                                                                |

---

### 💻 Electron Framework Projects
| Project Name | Use Case | Stars |
| :--- | :--- |:-----:|
| es-client | ElasticSearch client | [Official Website](https://es-client.esion.xyz) |


## Contribution Guide
We welcome contributions! 
> Tech Stack: Backend built with [Go Zero](https://go-zero.dev/), frontend built with [Vben](https://doc.vben.pro/) (Vue3). Admin backend references permission-related implementations from [simple-admin](https://doc.ryansu.tech/) (microservice architecture not adopted).


## Notifications
- Latest Updates: [UpgradeLink Notifications](https://github.com/toolsetlink/upgradelink/issues)


## Milestones
### v3.0.0 (2026-01-26)
- ⚡ **Second Development Support**: Core new second development support capability, with supporting documentation continuously updated
- 🔧 **Backend Reconstruction**: Comprehensive reconstruction of admin backend system, merging multiple services into 1 service, reducing 7 project modules to 3
- 📈 **Reduce Maintenance Cost**: Clean up useless backend modules, significantly reduce project maintenance costs

### v2.0.0 (2025-09-24)
- 🚀 **Official Open Source**: Project officially open sourced
- ⚡ **High-Speed Traffic Package**: Support high-speed traffic package to help enterprise users upgrade applications quickly
- 📦 **Multi-Platform Support**: Support Windows/Linux/Mac application upgrade modules
- 🔧 **Automated Deployment**: Add Tauri/Electron Github Action modules to support automated deployment and upgrade

### v1.x Series
- 🌱 **Basic Capability Building**: From v1.0.0 to v1.9.0, gradually implement core functions
- 📱 **Platform Expansion**: Support Android application upgrade, configuration upgrade and other functions
- 📊 **Data Statistics**: Add application traffic statistics, revenue module, etc.
- 🌐 **Global Access**: Application download and distribution pages support global access

## Future Plans
- 🔄 **Continue to Improve Second Development Documentation**: Based on v3.0.0 second development support capability, iteratively improve development documentation
- 📱 **More Platform Support**: Expand upgrade capabilities for more application platforms
- 🎯 **Feature Enhancement**: Implement more advanced functions based on the reconstructed admin system
- 🌍 **International Promotion**: Improve English translation, support more language versions

## Project Update Records
View the complete project update records: [UpgradeLink Update Records](https://www.toolsetlink.com/en/upgrade/update-record.html)

## Contributors
<a href="https://github.com/toolsetlink/upgradelink/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=toolsetlink/upgradelink" />
</a>
