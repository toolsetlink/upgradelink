import { createJiti } from "../../../../../../node_modules/.pnpm/jiti@2.6.1/node_modules/jiti/lib/jiti.mjs";

const jiti = createJiti(import.meta.url, {
  "interopDefault": true,
  "alias": {
    "@vben-core/shared": "/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/packages/@core/base/shared"
  },
  "transformOptions": {
    "babel": {
      "plugins": []
    }
  }
})

/** @type {import("/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/packages/@core/base/shared/src/utils/index.js")} */
const _module = await jiti.import("/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/packages/@core/base/shared/src/utils/index.ts");

export const get = _module.get;
export const isEqual = _module.isEqual;
export const set = _module.set;
export const cloneDeep = _module.cloneDeep;
export const cn = _module.cn;
export const formatDate = _module.formatDate;
export const formatDateTime = _module.formatDateTime;
export const isDate = _module.isDate;
export const isDayjsObject = _module.isDayjsObject;
export const getSystemTimezone = _module.getSystemTimezone;
export const setCurrentTimezone = _module.setCurrentTimezone;
export const getCurrentTimezone = _module.getCurrentTimezone;
export const arraysEqual = _module.arraysEqual;
export const diff = _module.diff;
export const getElementVisibleRect = _module.getElementVisibleRect;
export const getScrollbarWidth = _module.getScrollbarWidth;
export const needsScrollbar = _module.needsScrollbar;
export const triggerWindowResize = _module.triggerWindowResize;
export const downloadFileFromUrl = _module.downloadFileFromUrl;
export const downloadFileFromBase64 = _module.downloadFileFromBase64;
export const downloadFileFromImageUrl = _module.downloadFileFromImageUrl;
export const downloadFileFromBlob = _module.downloadFileFromBlob;
export const downloadFileFromBlobPart = _module.downloadFileFromBlobPart;
export const urlToBase64 = _module.urlToBase64;
export const triggerDownload = _module.triggerDownload;
export const getFirstNonNullOrUndefined = _module.getFirstNonNullOrUndefined;
export const isBoolean = _module.isBoolean;
export const isEmpty = _module.isEmpty;
export const isFunction = _module.isFunction;
export const isHttpUrl = _module.isHttpUrl;
export const isMacOs = _module.isMacOs;
export const isNumber = _module.isNumber;
export const isObject = _module.isObject;
export const isString = _module.isString;
export const isUndefined = _module.isUndefined;
export const isWindow = _module.isWindow;
export const isWindowsOs = _module.isWindowsOs;
export const capitalizeFirstLetter = _module.capitalizeFirstLetter;
export const kebabToCamelCase = _module.kebabToCamelCase;
export const toCamelCase = _module.toCamelCase;
export const toLowerCaseFirstLetter = _module.toLowerCaseFirstLetter;
export const mergeWithArrayOverride = _module.mergeWithArrayOverride;
export const key = _module.key;
export const updates = _module.updates;
export const createMerge = _module.createMerge;
export const merge = _module.merge;
export const startProgress = _module.startProgress;
export const stopProgress = _module.stopProgress;
export const loadScript = _module.loadScript;
export const StateHandler = _module.StateHandler;
export const to = _module.to;
export const filterTree = _module.filterTree;
export const mapTree = _module.mapTree;
export const sortTree = _module.sortTree;
export const traverseTreeValues = _module.traverseTreeValues;
export const uniqueByField = _module.uniqueByField;
export const updateCSSVariables = _module.updateCSSVariables;
export const bindMethods = _module.bindMethods;
export const getNestedValue = _module.getNestedValue;
export const openRouteInNewWindow = _module.openRouteInNewWindow;
export const openWindow = _module.openWindow;