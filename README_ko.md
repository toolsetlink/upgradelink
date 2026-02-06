# UpgradeLink 🚀

크로스 플랫폼 앱 업그레이드 및 배포 플랫폼 | 개발자를 위한 업그레이드 서비스 구축 비용의 90% 절감

[![GitHub Repo stars](https://img.shields.io/github/stars/toolsetlink/upgradelink)](https://github.com/toolsetlink/upgradelink)
[![stars](https://gitcode.com/toolsetlink/upgradelink/star/badge.svg)](https://gitcode.com/toolsetlink/upgradelink)
[![stars](https://gitee.com/toolsetlink/upgradelink/badge/star.svg)](https://gitee.com/toolsetlink/upgradelink)

[English](README.md) | [中文](README_zh.md) | [日本語](README_ja.md) | [Português](README_pt.md) | [Русский](README_ru.md) | [Español](README_es.md) | [한국어](README_ko.md) | [فارسی](README_fa.md)


## UpgradeLink를 선택해야 하는 이유?
개인 개발자의 소규모 도구이든 복잡한 기업급 애플리케이션이든, 서버를 처음부터 구축할 필요 없이 즉시 풀스택 업그레이드 기능을 통합할 수 있습니다:
- 🚀 **바로 사용 가능**: 업그레이드 논리에 깊이 들어갈 필요 없음 — SDK 통합에 단 10분 소요
- 🌍 **전 플랫폼 지원**: Windows/Linux/Mac/Android/Tauri/Electron 등 모든 시나리오 지원
- 💰 **비용 크게 절감**: 서버 설정 및 크로스 플랫폼 적응에 대한 R&D 투자 제거
- 🔒 **기업급 신뢰성**: 보안 보호 + 유연한 전략, 카나리아 릴리스/타겟 업그레이드 지원
- 📦 **유연한 저장소 지원**: S3 프로토콜 객체 저장소 지원, CDN 다운로드 주소 사용자 정의 가능


## 핵심 기능: 전체 시나리오 업그레이드 솔루션
### 1. 플랫폼/프레임워크별 업그레이드
| 지원 유형         | 핵심 기능                          | 사용 사례                  |
|------------------|-----------------------------------|---------------------------|
| Windows 앱       | 전용 업그레이드 전략 관리         | PC 데스크톱 애플리케이션   |
| Linux 앱         | 서버/데스크톱 업그레이드 요구사항에 적응 | 서버 도구, Linux 데스크톱 앱 |
| Mac 앱           | Apple 생태계와 호환되는 업그레이드 흐름 | Mac 데스크톱 애플리케이션 |
| Android 앱       | 델타 업데이트 (증분 업데이트) 지원 | 모바일 앱                  |
| Tauri 앱         | 공식 구성요소 API와 호환, 원활한 업그레이드 | Tauri 크로스 플랫폼 앱    |
| Electron 앱      | 공식 업그레이드 API와 호환, 제로 수정 통합 | Electron 크로스 플랫폼 앱 |

### 2. 유연한 리소스 업그레이드
| 업그레이드 유형   | 기능 설명                          | 사용 사례                  |
|------------------|-----------------------------------|---------------------------|
| 구성 업그레이드   | 사용자 정의 JSON 구성, 온라인 동적 업데이트 지원 | 동적 매개변수 조정이 필요한 앱 |
| 파일 업그레이드   | 리소스 보충/교체를 지원하기 위해 모든 파일 업로드 | 자료 또는 플러그인 업데이트가 필요한 앱 |
| URL 업그레이드    | 파일 저장소를 자체 관리하는 사용자 정의 다운로드 링크 | 기존 클라우드 저장소 서비스가 있는 앱 |

### 3. 유연한 저장 및 배포
- 📦 **S3 프로토콜 지원**: 모든 애플리케이션 버전 파일을 S3 호환 객체 저장 서비스에 저장할 수 있음
- 🌐 **사용자 정의 CDN**: 글로벌 사용자 다운로드 경험을 최적화하기 위해 사용자 정의 CDN 다운로드 주소 구성 지원


## 10분 통합 가이드 🔧
### 단계 1: 해당 SDK 선택
복잡한 구성 없이 6가지 주요 언어를 지원:
- 🐹 [Go SDK](https://github.com/toolsetlink/upgradelink-api-go)
- ☕ [Java SDK](https://github.com/toolsetlink/upgradelink-api-java)
- 🐍 [Python SDK](https://github.com/toolsetlink/upgradelink-api-python)
- 🦋 [Dart SDK](https://github.com/toolsetlink/upgradelink-api-dart)
- 🤖 [Android SDK](https://github.com/toolsetlink/upgradelink-api-android)
- 🟦 [TypeScript SDK](https://github.com/toolsetlink/upgradelink-api-typescript)
- 🟣 [C# SDK](https://github.com/toolsetlink/upgradelink-api-csharp)

### 단계 2: 문서 참조
- [전체 사용 가이드](https://www.toolsetlink.com/ko/upgrade/what-is-upgrade.html)
- [빠른 시작 튜토리얼](https://www.toolsetlink.com/ko/upgrade/deploy/quick-start-docker.html)


## 배포 방법
여러 배포 옵션을 지원하여 필요에 따라 선택하세요:
- [독립 실행형 빠른 시작 (Docker)](https://www.toolsetlink.com/ko/upgrade/deploy/quick-start-docker2.html)
- [Docker Compose 배포](https://www.toolsetlink.com/ko/upgrade/deploy/docker-compose.html)
  - [클러스터 배포](https://www.toolsetlink.com/ko/upgrade/deploy/cluster-docker.html)


## 누가 사용하고 있나요?

기술 프레임워크별로 분류된 여러 오픈 소스 프로젝트가 이미 통합되었습니다:

### 🚀 Tauri 프레임워크 프로젝트
| 프로젝트 이름 | 사용 사례 |                                                            Stars                                                            | 
| :--- | :--- |:---------------------------------------------------------------------------------------------------------------------------:| 
| note-gen | 크로스 플랫폼 노트 도구 |       [![GitHub Repo stars](https://img.shields.io/github/stars/codexu/note-gen)](https://github.com/codexu/note-gen)       | 
| BongoCat | 재미있는 데스크톱 펫 |     [![GitHub Repo stars](https://img.shields.io/github/stars/ayangweb/BongoCat)](https://github.com/ayangweb/BongoCat)     | 
| Hula | 오픈 소스 인스턴트 메시징 시스템 |        [![GitHub Repo stars](https://img.shields.io/github/stars/HuLaSpark/HuLa)](https://github.com/HuLaSpark/HuLa)        | 
| MarkFlowy | 경량 Markdown 도구 |   [![GitHub Repo stars](https://img.shields.io/github/stars/drl990114/MarkFlowy)](https://github.com/drl990114/MarkFlowy)   |
| lazyeat | 제스처 제어 도구 |     [![GitHub Repo stars](https://img.shields.io/github/stars/lanxiuyun/lazyeat)](https://github.com/lanxiuyun/lazyeat)     | 
| Prompt-Tools | AI 프롬프트 도구 | [![GitHub Repo stars](https://img.shields.io/github/stars/jwangkun/Prompt-Tools)](https://github.com/jwangkun/Prompt-Tools) | 
| MarkFly | Markdown 편집기 |      [![GitHub Repo stars](https://img.shields.io/github/stars/jwangkun/MarkFly)](https://github.com/jwangkun/MarkFly)      | 
| rank-analysis | LoL 랭킹 분석 도구 |  [![GitHub Repo stars](https://img.shields.io/github/stars/wnzzer/rank-analysis)](https://github.com/wnzzer/rank-analysis)  |
| Welight | 위챗 공식 계정 지능형 레이아웃 애플리케이션 |                                                   [공식 웹사이트](https://waer.ltd)                                                    |  

---

### 📱 Android 프로젝트
| 프로젝트 이름 | 사용 사례 |                                                             Stars                                                             |
| :--- | :--- |:-----------------------------------------------------------------------------------------------------------------------------:|
| 聚在工大 | 캠퍼스 서비스 | [![GitHub Repo stars](https://img.shields.io/github/stars/Chiu-xaH/HFUT-Schedule)](https://github.com/Chiu-xaH/HFUT-Schedule) |
| 脉扑 | 건강과 웰니스 |                                                                [공식 웹사이트](https://www.maipusc.com)                                                                  |

---

### 💻 Electron 프레임워크 프로젝트
| 프로젝트 이름 | 사용 사례 | Stars |
| :--- | :--- |:-----:|
| es-client | ElasticSearch 클라이언트 |   [공식 웹사이트](https://es-client.esion.xyz)    |


## 기여 가이드
기여를 환영합니다! 
> 기술 스택: 백엔드는 [Go Zero](https://go-zero.dev/) 기반으로 개발, 프론트엔드는 [Vben](https://doc.vben.pro/) (Vue3) 기반으로 개발. Admin 백엔드는 [simple-admin](https://doc.ryansu.tech/)의 권한 관련 구현을 참조 (마이크로서비스 아키텍처는 채택하지 않음).


## 알림 
- 최신 업데이트: [UpgradeLink 알림](https://github.com/toolsetlink/upgradelink/issues)


## 마일스톤
### v3.0.0 (2026-01-26)
- ⚡ **이차 개발 지원**: 핵심 신규 이차 개발 지원 기능, 지원 문서 지속 업데이트
- 🔧 **백엔드 재구축**: 관리 백엔드 시스템 전면 재구축, 다중 서비스를 1개 서비스로 병합, 7개 프로젝트 모듈을 3개로 축소
- 📈 **유지 관리 비용 절감**: 백엔드 불필요한 모듈 정리, 프로젝트 유지 관리 비용 크게 절감

### v2.0.0 (2025-09-24)
- 🚀 **공식 오픈 소스**: 프로젝트 공식 오픈 소스 출시
- ⚡ **고속 트래픽 패키지**: 기업 사용자가 애플리케이션을 빠르게 업그레이드할 수 있도록 고속 트래픽 패키지 지원
- 📦 **다중 플랫폼 지원**: Windows/Linux/Mac 애플리케이션 업그레이드 모듈 지원
- 🔧 **자동화 배포**: Tauri/Electron GitHub Action 모듈 추가, 자동화 배포 및 업그레이드 지원

### v1.x 시리즈
- 🌱 **기본 기능 구축**: v1.0.0부터 v1.9.0까지 핵심 기능 단계적 구현
- 📱 **플랫폼 확장**: Android 애플리케이션 업그레이드, 구성 업그레이드 등 기능 지원
- 📊 **데이터 통계**: 애플리케이션 트래픽 통계, 수입 모듈 등 추가
- 🌐 **글로벌 액세스**: 애플리케이션 다운로드 및 배포 페이지 글로벌 액세스 지원

## 향후 계획
- 🔄 **이차 개발 문서 지속 개선**: v3.0.0 이차 개발 지원 기능을 기반으로 개발 문서를 반복적으로 개선
- 📱 **더 많은 플랫폼 지원**: 더 많은 애플리케이션 플랫폼에 대한 업그레이드 기능 확장
- 🎯 **기능 강화**: 재구축된 관리 시스템을 기반으로 더 많은 고급 기능 구현
- 🌍 **국제화 홍보**: 영어 번역 개선, 더 많은 언어 버전 지원

## 프로젝트 업데이트 기록
프로젝트 업데이트 기록 전체 보기: [UpgradeLink 업데이트 기록](https://www.toolsetlink.com/ko/upgrade/update-record.html)

## 기여자
<a href="https://github.com/toolsetlink/upgradelink/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=toolsetlink/upgradelink" />
</a>