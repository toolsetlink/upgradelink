/**
 * @zh_CN 登录页面 url 地址
 */
export const LOGIN_PATH = '/auth/login';

export interface LanguageOption {
  label: string;
  value: 'en-US' | 'zh-CN' | 'ja-JP' | 'pt-BR' | 'ru-RU' | 'es-ES' | 'ko-KR' | 'fa-IR';
}

/**
 * Supported languages
 */
export const SUPPORT_LANGUAGES: LanguageOption[] = [
  {
    label: '简体中文',
    value: 'zh-CN',
  },
  {
    label: 'English',
    value: 'en-US',
  },
  {
    label: '日本語',
    value: 'ja-JP',
  },
  {
    label: 'Português (Brasil)',
    value: 'pt-BR',
  },
  {
    label: 'Русский',
    value: 'ru-RU',
  },
  {
    label: 'Español',
    value: 'es-ES',
  },
  {
    label: '한국어',
    value: 'ko-KR',
  },
  {
    label: 'فارسی',
    value: 'fa-IR',
  },
];
