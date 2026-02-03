# UpgradeLink 🚀

پلتفرم ارتقا و توزیع اپلیکیشن‌های چند‌پلتفرمی | ۹۰ درصد هزینه ساخت سرویس‌های ارتقا را برای توسعه‌دهندگان کاهش دهد

[![GitHub Repo stars](https://img.shields.io/github/stars/toolsetlink/upgradelink)](https://github.com/toolsetlink/upgradelink)
[![stars](https://gitcode.com/toolsetlink/upgradelink/star/badge.svg)](https://gitcode.com/toolsetlink/upgradelink)
[![stars](https://gitee.com/toolsetlink/upgradelink/badge/star.svg)](https://gitee.com/toolsetlink/upgradelink)

[English](README.md) | [中文](README_zh.md) | [日本語](README_ja.md) | [Português](README_pt.md) | [Русский](README_ru.md) | [Español](README_es.md) | [한국어](README_ko.md) | [فارسی](README_fa.md)


## چرا UpgradeLink را انتخاب کنم؟
آیا ابزاری کوچک توسط توسعه‌دهندگان فردی یا یک اپلیکیشن پیچیده شرکتی باشد، می‌توانید به صورت آنی قابلیت‌های ارتقا کلاس را بدون ساخت سرویس‌ها از ابتدا ادغام کنید:
- 🚀 **از جایی که هست آماده**: نیازی به غوطه‌ور شدن در منطق ارتقا نیست — ادغام SDK فقط ۱۰ دقیقه طول می‌کشد
- 🌍 **پوشش چند‌پلتفرمی**: از Windows/Linux/Mac/Android/Tauri/Electron و سایر سناریوهای کامل پشتیبانی می‌کند
- 💰 **کاهش هزینه**: سرمایه‌گذاری R&D در راه‌اندازی سرور و سازگاری چند‌پلتفرمی را حذف می‌کند
- 🔒 **امنیتی شرکتی**: محافظت امنیتی + استراتژی‌های انعطاف‌پذیر، پشتیبانی از انتشار بزیرگانه/ارتقا هدفمند
- 📦 **پشتیبانی از ذخیره‌سازی انعطاف‌پذیر**: از ذخیره‌سازی اشیاء با پروتکل S3 پشتیبانی می‌کند، آدرس دانلود CDN قابل سفارش است


## ویژگی‌های اصلی: راه‌حل‌های ارتقا برای تمام سناریوها
### ۱. ارتقاهاي مخصوص پلتفرم/چارچوب
| نوع پشتیبانی شده       | قابلیت‌های اصلی                          | موارد استفاده                  |
|----------------------|-------------------------------------------|-------------------------------|
| اپلیکیشن‌های Windows  | مدیریت استراتژی ارتقا اختصاصی           | اپلیکیشن‌های دسکتاپ PC        |
| اپلیکیشن‌های Linux    | سازگاری با نیازهای ارتقا سرور/دسکتاپ      | ابزارهای سرور، اپلیکیشن‌های دسکتاپ Linux |
| اپلیکیشن‌های Mac      | روند ارتقا سازگار با экوسیستم Apple       | اپلیکیشن‌های دسکتاپ Mac       |
| اپلیکیشن‌های Android  | پشتیبانی از بروزرسانی‌های دلتا (بروزرسانی‌های افزایشی) | اپلیکیشن‌های موبایل            |
| اپلیکیشن‌های Tauri    | سازگاری با APIهای قطعه‌های رسمی برای ارتقاها بدون وقفه | اپلیکیشن‌های چند‌پلتفرمی Tauri |
| اپلیکیشن‌های Electron | سازگاری با APIهای ارتقا رسمی، ادغام بدون تغییر | اپلیکیشن‌های چند‌پلتفرمی Electron |

### ۲. ارتقاهاي منابع انعطاف‌پذیر
| نوع ارتقا             | توضیح ویژگی                               | موارد استفاده                  |
|----------------------|-------------------------------------------|-------------------------------|
| ارتقا پیکربندی       | پیکربندی JSON سفارشی با بروزرسانی‌های پویا آنلاین | اپلیکیشن‌هایی که نیاز به تنظیمات پویا پارامترها دارند |
| ارتقا فایل            | بارگذاری هر فایلی برای پشتیبانی از تکمیل/جایگزینی منابع | اپلیکیشن‌هایی که نیاز به بروزرسانی مواد یا پلاگین‌ها دارند |
| ارتقا URL             | لینک‌های دانلود سفارشی با ذخیره‌سازی فایل خودمدار | اپلیکیشن‌هایی با سرویس‌های ذخیره‌سازی ابری موجود |

### ۳. ذخیره‌سازی و توزیع انعطاف‌پذیر
- 📦 **پشتیبانی از پروتکل S3**: تمام فایل‌های نسخه اپلیکیشن را می‌توان در سرویس‌های ذخیره‌سازی اشیاء سازگار با S3 ذخیره کرد
- 🌐 **CDN قابل سفارش**: از پیکربندی آدرس‌های دانلود CDN سفارشی برای بهینه‌سازی تجربه دانلود کاربران جهانی پشتیبانی می‌کند


## راهنمای ادغام ۱۰ دقیقه ای 🔧
### مرحله ۱: SDK مربوطه را انتخاب کنید
۶ زبان اصلی را بدون پیکربندی پیچیده پشتیبانی می‌کند:
- 🐹 [Go SDK](https://github.com/toolsetlink/upgradelink-api-go)
- ☕ [Java SDK](https://github.com/toolsetlink/upgradelink-api-java)
- 🐍 [Python SDK](https://github.com/toolsetlink/upgradelink-api-python)
- 🦋 [Dart SDK](https://github.com/toolsetlink/upgradelink-api-dart)
- 🤖 [Android SDK](https://github.com/toolsetlink/upgradelink-api-android)
- 🟦 [TypeScript SDK](https://github.com/toolsetlink/upgradelink-api-typescript)
- 🟣 [C# SDK](https://github.com/toolsetlink/upgradelink-api-csharp)

### مرحله ۲: به مستندات مراجعه کنید
- [راهنمای استفاده کامل](https://www.toolsetlink.com/en/upgrade/what-is-upgrade.html)
- [آموزش شروع سریع](https://www.toolsetlink.com/en/upgrade/deploy/quick-start-docker.html)


## روش‌های استقرار
چندین گزینه استقرار را برای برآورده کردن نیازهای شما پشتیبانی می‌کند:
- [شروع سریع مستقل (Docker)](https://www.toolsetlink.com/en/upgrade/deploy/quick-start-docker2.html)
- [استقرار Docker Compose](https://www.toolsetlink.com/en/upgrade/deploy/docker-compose.html)
  - [استقرار خوشه‌ای](https://www.toolsetlink.com/en/upgrade/deploy/cluster-docker.html)


## چه کسی در حال استفاده است؟

چندین پروژه متن‌باز قبلاً ادغام شده‌اند که بر اساس چارچوب فنی دسته‌بندی شده‌اند:

### 🚀 پروژه‌های چارچوب Tauri
| نام پروژه | مورد استفاده | Stars |
| :--- | :--- |:---------------------------------------------------------------------------------------------------------------------------:|
| note-gen | ابزار یادداشت چند‌پلتفرمی | [![GitHub Repo stars](https://img.shields.io/github/stars/codexu/note-gen)](https://github.com/codexu/note-gen) |
| BongoCat | حیوان خانگی دسکتاپ سرگرم‌کننده | [![GitHub Repo stars](https://img.shields.io/github/stars/ayangweb/BongoCat)](https://github.com/ayangweb/BongoCat) |
| Hula | سیستم پیام‌رسانی آنی متن‌باز | [![GitHub Repo stars](https://img.shields.io/github/stars/HuLaSpark/HuLa)](https://github.com/HuLaSpark/HuLa) |
| MarkFlowy | ابزار Markdown سبک وزن | [![GitHub Repo stars](https://img.shields.io/github/stars/drl990114/MarkFlowy)](https://github.com/drl990114/MarkFlowy) |
| lazyeat | ابزار کنترل با ژست | [![GitHub Repo stars](https://img.shields.io/github/stars/lanxiuyun/lazyeat)](https://github.com/lanxiuyun/lazyeat) |
| Prompt-Tools | ابزار提示词 AI | [![GitHub Repo stars](https://img.shields.io/github/stars/jwangkun/Prompt-Tools)](https://github.com/jwangkun/Prompt-Tools) |
| MarkFly | ویرایشگر Markdown | [![GitHub Repo stars](https://img.shields.io/github/stars/jwangkun/MarkFly)](https://github.com/jwangkun/MarkFly) |
| rank-analysis | ابزار تحلیل رتبه‌بندی LoL | [![GitHub Repo stars](https://img.shields.io/github/stars/wnzzer/rank-analysis)](https://github.com/wnzzer/rank-analysis) |
| Welight | اپلیکیشن چیدمان هوشمند برای حساب رسمی WeChat | [وب‌سایت رسمی](https://waer.ltd) |

---

### 📱 پروژه‌های Android
| نام پروژه | مورد استفاده | Stars |
| :--- | :--- |:-----------------------------------------------------------------------------------------------------------------------------:|
| 聚在工大 | سرویس دانشگاهی | [![GitHub Repo stars](https://img.shields.io/github/stars/Chiu-xaH/HFUT-Schedule)](https://github.com/Chiu-xaH/HFUT-Schedule) |
| 脉扑 | سلامت و رفاه | [وب‌سایت رسمی](https://www.maipusc.com) |

---

### 💻 پروژه‌های چارچوب Electron
| نام پروژه | مورد استفاده | Stars |
| :--- | :--- |:-----:|
| es-client | کلاینت ElasticSearch | [وب‌سایت رسمی](https://es-client.esion.xyz) |


## راهنمای مشارکت
ما مشارکت‌ها را خوشحال می‌بینیم! 
> استک تکنولوژی: بک‌اند بر اساس [Go Zero](https://go-zero.dev/) توسعه یافته، فرانت‌اند بر اساس [Vben](https://doc.vben.pro/) (Vue3) توسعه یافته. بک‌اند ادمین از پیاده‌سازی‌های مربوط به مجوز [simple-admin](https://doc.ryansu.tech/) (архیتکچر میکروسرویس آن اتخاذ نشده) مرجع می‌کند.


## اطلاعیه‌ها و بازخورد
- آخرین بروزرسانی‌ها: [اطلاعیه‌های UpgradeLink](https://github.com/toolsetlink/upgradelink/issues)
- بازخورد: [با ما تماس بگیرید](https://www.toolsetlink.com/en/upgrade/contact-us.html)


## نقطه‌های عطف
### v3.0.0 (2026-01-26)
- ⚡ **پشتیبانی از توسعه ثانویه**: قابلیت جدید پشتیبانی از توسعه ثانویه، با مستندات پشتیبانی که به طور مداوم به‌روزرسانی می‌شوند
- 🔧 **تغییر ساختار بک‌اند**: بازسازی کامل سیستم بک‌اند ادمین، ادغام چندین سرویس در ۱ سرویس، کاهش ۷ ماژول پروژه به ۳
- 📈 **کاهش هزینه نگهداری**: پاکسازی ماژول‌های بیکار بک‌اند، کاهش قابل توجه هزینه نگهداری پروژه

### v2.0.0 (2025-09-24)
- 🚀 **انعطاف‌پذیری رسمی**: پروژه به صورت رسمی به عنوان متن‌باز منتشر شد
- ⚡ **بسته ترافیک با سرعت بالا**: پشتیبانی از بسته‌های ترافیک با سرعت بالا برای کمک به کاربران شرکت‌ها به‌منظور ارتقا سریع اپلیکیشن‌ها
- 📦 **پشتیبانی چند‌پلتفرمی**: پشتیبانی از ماژول‌های ارتقا اپلیکیشن‌های Windows/Linux/Mac
- 🔧 **استقرار اتوماتیک**: اضافه کردن ماژول GitHub Action برای Tauri/Electron، پشتیبانی از استقرار و ارتقا اتوماتیک

### سری v1.x
- 🌱 **ساخت قابلیت‌های پایه**: از v1.0.0 تا v1.9.0، پیاده‌سازی تدریجی قابلیت‌های اصلی
- 📱 **گسترش پلتفرم**: پشتیبانی از ارتقا اپلیکیشن‌های Android، ارتقا پیکربندی و سایر قابلیت‌ها
- 📊 **آمار داده**: اضافه کردن آمار ترافیک اپلیکیشن، ماژول درآمد و غیره
- 🌐 **دسترسی جهانی**: صفحات دانلود و توزیع اپلیکیشن از دسترسی جهانی پشتیبانی می‌کنند

## برنامه‌های آتی
- 🔄 **ادامه بهبود مستندات توسعه ثانویه**: بر اساس قابلیت پشتیبانی از توسعه ثانویه v3.0.0، بهبود تدریجی مستندات توسعه
- 📱 **پشتیبانی از پلتفرم‌های بیشتر**: گسترش قابلیت‌های ارتقا برای پلتفرم‌های اپلیکیشن بیشتر
- 🎯 **افزایش قابلیت‌ها**: پیاده‌سازی قابلیت‌های پیشرفته‌تر بر اساس سیستم ادمین بازسازی شده
- 🌍 **ترویج بین‌المللی**: بهبود ترجمه انگلیسی، پشتیبانی از نسخه‌های بیشتر در زبان‌های دیگر

## رکوردهای بروزرسانی پروژه
برای مشاهده رکوردهای کامل بروزرسانی پروژه: [رکوردهای بروزرسانی UpgradeLink](https://www.toolsetlink.com/upgrade/update-record.html)

## مشارکت‌کنندگان
<a href="https://github.com/toolsetlink/upgradelink/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=toolsetlink/upgradelink" />
</a>