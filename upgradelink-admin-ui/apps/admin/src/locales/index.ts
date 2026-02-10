import type { Locale } from 'ant-design-vue/es/locale';

import type { App } from 'vue';

import type { LocaleSetupOptions, SupportedLanguagesType } from '@vben/locales';

import { ref } from 'vue';

import {
  $t,
  setupI18n as coreSetup,
  loadLocalesMapFromDir,
} from '@vben/locales';
import { preferences } from '@vben/preferences';

import antdEnLocale from 'ant-design-vue/es/locale/en_US';
import antdDefaultLocale from 'ant-design-vue/es/locale/zh_CN';
import antdJaLocale from 'ant-design-vue/es/locale/ja_JP';
import antdPtLocale from 'ant-design-vue/es/locale/pt_BR';
import antdRuLocale from 'ant-design-vue/es/locale/ru_RU';
import antdEsLocale from 'ant-design-vue/es/locale/es_ES';
import antdKoLocale from 'ant-design-vue/es/locale/ko_KR';
import dayjs from 'dayjs';

const antdLocale = ref<Locale>(antdDefaultLocale);

const modules = import.meta.glob('./langs/**/*.json');

const localesMap = loadLocalesMapFromDir(
  /\.\/langs\/([^/]+)\/(.*)\.json$/,
  modules,
);
/**
 * 加载应用特有的语言包
 * 这里也可以改造为从服务端获取翻译数据
 * @param lang
 */
async function loadMessages(lang: SupportedLanguagesType) {
  const [appLocaleMessages] = await Promise.all([
    localesMap[lang]?.(),
    loadThirdPartyMessage(lang),
  ]);
  return appLocaleMessages?.default || {};
}

/**
 * 加载第三方组件库的语言包
 * @param lang
 */
async function loadThirdPartyMessage(lang: SupportedLanguagesType) {
  await Promise.all([loadAntdLocale(lang), loadDayjsLocale(lang)]);
}

/**
 * 加载dayjs的语言包
 * @param lang
 */
async function loadDayjsLocale(lang: SupportedLanguagesType) {
  let locale;
  switch (lang) {
    case 'en-US': {
      locale = await import('dayjs/locale/en');
      break;
    }
    case 'zh-CN': {
      locale = await import('dayjs/locale/zh-cn');
      break;
    }
    case 'ja-JP': {
      locale = await import('dayjs/locale/ja');
      break;
    }
    case 'pt-BR': {
      locale = await import('dayjs/locale/pt-br');
      break;
    }
    case 'ru-RU': {
      locale = await import('dayjs/locale/ru');
      break;
    }
    case 'es-ES': {
      locale = await import('dayjs/locale/es');
      break;
    }
    case 'ko-KR': {
      locale = await import('dayjs/locale/ko');
      break;
    }
    case 'fa-IR': {
      locale = await import('dayjs/locale/fa');
      break;
    }
    // 默认使用英语
    default: {
      locale = await import('dayjs/locale/en');
    }
  }
  if (locale) {
    dayjs.locale(locale);
  } else {
    console.error(`Failed to load dayjs locale for ${lang}`);
  }
}

/**
 * 加载antd的语言包
 * @param lang
 */
async function loadAntdLocale(lang: SupportedLanguagesType) {
  switch (lang) {
    case 'en-US': {
      antdLocale.value = antdEnLocale;
      break;
    }
    case 'zh-CN': {
      antdLocale.value = antdDefaultLocale;
      break;
    }
    case 'ja-JP': {
      antdLocale.value = antdJaLocale;
      break;
    }
    case 'pt-BR': {
      antdLocale.value = antdPtLocale;
      break;
    }
    case 'ru-RU': {
      antdLocale.value = antdRuLocale;
      break;
    }
    case 'es-ES': {
      antdLocale.value = antdEsLocale;
      break;
    }
    case 'ko-KR': {
      antdLocale.value = antdKoLocale;
      break;
    }
    // Persian (fa-IR) is not supported by antd, default to English
    case 'fa-IR': {
      antdLocale.value = antdEnLocale;
      break;
    }
  }
}

async function setupI18n(app: App, options: LocaleSetupOptions = {}) {
  await coreSetup(app, {
    defaultLocale: preferences.app.locale,
    loadMessages,
    missingWarn: !import.meta.env.PROD,
    ...options,
  });
}

export { $t, antdLocale, setupI18n };
