import { defineOverridesPreferences } from '@vben/preferences';

/**
 * @description 项目配置文件
 * 只需要覆盖项目中的一部分配置，不需要的配置不用覆盖，会自动使用默认配置
 * !!! 更改配置后请清空缓存，否则可能不生效
 */
export const overridesPreferences = defineOverridesPreferences({
  // overrides
  app: {
    name: import.meta.env.VITE_APP_TITLE,
    defaultAvatar: "",
    accessMode: "backend",
    defaultHomePath: "/dashboard", // 默认首页路径
    enableCheckUpdates: true, // 是否开启检查更新
    checkUpdatesInterval: 60,  // 检查更新的时间间隔，单位为分钟
  },
  logo: {
    enable: true,
    fit: 'contain',
    source: 'https://www.toolsetlink.com/dl-logo-48x.svg',
  },
  copyright: {
    companyName: "UpgradeLink",
    companySiteLink: "https://www.toolsetlink.com/",
    date: "2026",
    enable: true,
    icp: "",
    icpLink: "",
    settingShow: true,
  },
  widget: {
    notification: false,
    timezone: false,
  },
});
