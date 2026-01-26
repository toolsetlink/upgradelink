import { clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';
import dayjs from 'dayjs';
import timezone from 'dayjs/plugin/timezone';
import utc from 'dayjs/plugin/utc';
import { isString, isObject } from '@vue/shared';
export { isFunction, isObject, isString } from '@vue/shared';
import { createDefu } from 'defu';
export { createDefu as createMerge, defu as merge } from 'defu';
export { get, isEqual, set } from 'es-toolkit/compat';
export { default as cloneDeep } from 'lodash.clonedeep';

function cn(...inputs) {
  return twMerge(clsx(inputs));
}

dayjs.extend(utc);
dayjs.extend(timezone);
function formatDate(time, format = "YYYY-MM-DD") {
  try {
    const date = dayjs.isDayjs(time) ? time : dayjs(time);
    if (!date.isValid()) {
      throw new Error("Invalid date");
    }
    return date.tz().format(format);
  } catch (error) {
    console.error(`Error formatting date: ${error}`);
    return String(time ?? "");
  }
}
function formatDateTime(time) {
  return formatDate(time, "YYYY-MM-DD HH:mm:ss");
}
function isDate(value) {
  return value instanceof Date;
}
function isDayjsObject(value) {
  return dayjs.isDayjs(value);
}
const getSystemTimezone = () => {
  return dayjs.tz.guess();
};
let currentTimezone = getSystemTimezone();
const setCurrentTimezone = (timezone2) => {
  currentTimezone = timezone2 || getSystemTimezone();
  dayjs.tz.setDefault(currentTimezone);
};
const getCurrentTimezone = () => {
  return currentTimezone;
};

function arraysEqual(a, b) {
  if (a.length !== b.length) return false;
  const counter = /* @__PURE__ */ new Map();
  for (const value of a) {
    counter.set(value, (counter.get(value) || 0) + 1);
  }
  for (const value of b) {
    const count = counter.get(value);
    if (count === void 0 || count === 0) {
      return false;
    }
    counter.set(value, count - 1);
  }
  return true;
}
function diff(obj1, obj2) {
  function findDifferences(o1, o2) {
    if (Array.isArray(o1) && Array.isArray(o2)) {
      if (!arraysEqual(o1, o2)) {
        return o2;
      }
      return void 0;
    }
    if (typeof o1 === "object" && typeof o2 === "object" && o1 !== null && o2 !== null) {
      const diffResult = {};
      const keys = /* @__PURE__ */ new Set([...Object.keys(o1), ...Object.keys(o2)]);
      keys.forEach((key) => {
        const valueDiff = findDifferences(o1[key], o2[key]);
        if (valueDiff !== void 0) {
          diffResult[key] = valueDiff;
        }
      });
      return Object.keys(diffResult).length > 0 ? diffResult : void 0;
    }
    return o1 === o2 ? void 0 : o2;
  }
  return findDifferences(obj1, obj2);
}

function getElementVisibleRect(element) {
  if (!element) {
    return {
      bottom: 0,
      height: 0,
      left: 0,
      right: 0,
      top: 0,
      width: 0
    };
  }
  const rect = element.getBoundingClientRect();
  const viewHeight = Math.max(
    document.documentElement.clientHeight,
    window.innerHeight
  );
  const top = Math.max(rect.top, 0);
  const bottom = Math.min(rect.bottom, viewHeight);
  const viewWidth = Math.max(
    document.documentElement.clientWidth,
    window.innerWidth
  );
  const left = Math.max(rect.left, 0);
  const right = Math.min(rect.right, viewWidth);
  return {
    bottom,
    height: Math.max(0, bottom - top),
    left,
    right,
    top,
    width: Math.max(0, right - left)
  };
}
function getScrollbarWidth() {
  const scrollDiv = document.createElement("div");
  scrollDiv.style.visibility = "hidden";
  scrollDiv.style.overflow = "scroll";
  scrollDiv.style.position = "absolute";
  scrollDiv.style.top = "-9999px";
  document.body.append(scrollDiv);
  const innerDiv = document.createElement("div");
  scrollDiv.append(innerDiv);
  const scrollbarWidth = scrollDiv.offsetWidth - innerDiv.offsetWidth;
  scrollDiv.remove();
  return scrollbarWidth;
}
function needsScrollbar() {
  const doc = document.documentElement;
  const body = document.body;
  const overflowY = window.getComputedStyle(body).overflowY;
  if (overflowY === "scroll" || overflowY === "auto") {
    return doc.scrollHeight > window.innerHeight;
  }
  return doc.scrollHeight > window.innerHeight;
}
function triggerWindowResize() {
  const resizeEvent = new Event("resize");
  window.dispatchEvent(resizeEvent);
}

function openWindow(url, options = {}) {
  const { noopener = true, noreferrer = true, target = "_blank" } = options;
  const features = [noopener && "noopener=yes", noreferrer && "noreferrer=yes"].filter(Boolean).join(",");
  window.open(url, target, features);
}
function openRouteInNewWindow(path) {
  const { hash, origin } = location;
  const fullPath = path.startsWith("/") ? path : `/${path}`;
  const url = `${origin}${hash && !fullPath.startsWith("/#") ? "/#" : ""}${fullPath}`;
  openWindow(url, { target: "_blank" });
}

const DEFAULT_FILENAME = "downloaded_file";
async function downloadFileFromUrl({
  fileName,
  source,
  target = "_blank"
}) {
  if (!source || typeof source !== "string") {
    throw new Error("Invalid URL.");
  }
  const isChrome = window.navigator.userAgent.toLowerCase().includes("chrome");
  const isSafari = window.navigator.userAgent.toLowerCase().includes("safari");
  if (/iP/.test(window.navigator.userAgent)) {
    console.error("Your browser does not support download!");
    return;
  }
  if (isChrome || isSafari) {
    triggerDownload(source, resolveFileName(source, fileName));
    return;
  }
  if (!source.includes("?")) {
    source += "?download";
  }
  openWindow(source, { target });
}
function downloadFileFromBase64({ fileName, source }) {
  if (!source || typeof source !== "string") {
    throw new Error("Invalid Base64 data.");
  }
  const resolvedFileName = fileName || DEFAULT_FILENAME;
  triggerDownload(source, resolvedFileName);
}
async function downloadFileFromImageUrl({
  fileName,
  source
}) {
  const base64 = await urlToBase64(source);
  downloadFileFromBase64({ fileName, source: base64 });
}
function downloadFileFromBlob({
  fileName = DEFAULT_FILENAME,
  source
}) {
  if (!(source instanceof Blob)) {
    throw new TypeError("Invalid Blob data.");
  }
  const url = URL.createObjectURL(source);
  triggerDownload(url, fileName);
}
function downloadFileFromBlobPart({
  fileName = DEFAULT_FILENAME,
  source
}) {
  const blob = source instanceof Blob ? source : new Blob([source], { type: "application/octet-stream" });
  const url = URL.createObjectURL(blob);
  triggerDownload(url, fileName);
}
function urlToBase64(url, mineType) {
  return new Promise((resolve, reject) => {
    let canvas = document.createElement("CANVAS");
    const ctx = canvas?.getContext("2d");
    const img = new Image();
    img.crossOrigin = "";
    img.addEventListener("load", () => {
      if (!canvas || !ctx) {
        return reject(new Error("Failed to create canvas."));
      }
      canvas.height = img.height;
      canvas.width = img.width;
      ctx.drawImage(img, 0, 0);
      const dataURL = canvas.toDataURL(mineType || "image/png");
      canvas = null;
      resolve(dataURL);
    });
    img.src = url;
  });
}
function triggerDownload(href, fileName, revokeDelay = 100) {
  const defaultFileName = "downloaded_file";
  const finalFileName = fileName || defaultFileName;
  const link = document.createElement("a");
  link.href = href;
  link.download = finalFileName;
  link.style.display = "none";
  if (link.download === void 0) {
    link.setAttribute("target", "_blank");
  }
  document.body.append(link);
  link.click();
  link.remove();
  setTimeout(() => URL.revokeObjectURL(href), revokeDelay);
}
function resolveFileName(url, fileName) {
  return fileName || url.slice(url.lastIndexOf("/") + 1) || DEFAULT_FILENAME;
}

function isUndefined(value) {
  return value === void 0;
}
function isBoolean(value) {
  return typeof value === "boolean";
}
function isEmpty(value) {
  if (value === null || value === void 0) {
    return true;
  }
  if (Array.isArray(value) || isString(value)) {
    return value.length === 0;
  }
  if (value instanceof Map || value instanceof Set) {
    return value.size === 0;
  }
  if (isObject(value)) {
    return Object.keys(value).length === 0;
  }
  return false;
}
function isHttpUrl(url) {
  if (!url) {
    return false;
  }
  const httpRegex = /^https?:\/\/.*$/;
  return httpRegex.test(url);
}
function isWindow(value) {
  return typeof window !== "undefined" && value !== null && value === value.window;
}
function isMacOs() {
  const macRegex = /macintosh|mac os x/i;
  return macRegex.test(navigator.userAgent);
}
function isWindowsOs() {
  const windowsRegex = /windows|win32/i;
  return windowsRegex.test(navigator.userAgent);
}
function isNumber(value) {
  return typeof value === "number" && Number.isFinite(value);
}
function getFirstNonNullOrUndefined(...values) {
  for (const value of values) {
    if (value !== void 0 && value !== null) {
      return value;
    }
  }
  return void 0;
}

function capitalizeFirstLetter(string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}
function toLowerCaseFirstLetter(str) {
  if (!str) return str;
  return str.charAt(0).toLowerCase() + str.slice(1);
}
function toCamelCase(key, parentKey) {
  if (!parentKey) {
    return key;
  }
  return parentKey + key.charAt(0).toUpperCase() + key.slice(1);
}
function kebabToCamelCase(str) {
  return str.split("-").filter(Boolean).map(
    (word, index) => index === 0 ? word : word.charAt(0).toUpperCase() + word.slice(1)
  ).join("");
}

const mergeWithArrayOverride = createDefu((originObj, key, updates) => {
  if (Array.isArray(originObj[key]) && Array.isArray(updates)) {
    originObj[key] = updates;
    return true;
  }
});

let nProgressInstance = null;
async function loadNprogress() {
  if (nProgressInstance) {
    return nProgressInstance;
  }
  nProgressInstance = await import('nprogress');
  nProgressInstance.configure({
    showSpinner: true,
    speed: 300
  });
  return nProgressInstance;
}
async function startProgress() {
  const nprogress = await loadNprogress();
  nprogress?.start();
}
async function stopProgress() {
  const nprogress = await loadNprogress();
  nprogress?.done();
}

function loadScript(src) {
  return new Promise((resolve, reject) => {
    if (document.querySelector(`script[src="${src}"]`)) {
      return resolve();
    }
    const script = document.createElement("script");
    script.src = src;
    script.addEventListener("load", () => resolve());
    script.addEventListener(
      "error",
      () => reject(new Error(`Failed to load script: ${src}`))
    );
    document.head.append(script);
  });
}

class StateHandler {
  condition = false;
  rejectCondition = null;
  resolveCondition = null;
  isConditionTrue() {
    return this.condition;
  }
  reset() {
    this.condition = false;
    this.clearPromises();
  }
  // 触发状态为 false 时，reject
  setConditionFalse() {
    this.condition = false;
    if (this.rejectCondition) {
      this.rejectCondition();
      this.clearPromises();
    }
  }
  // 触发状态为 true 时，resolve
  setConditionTrue() {
    this.condition = true;
    if (this.resolveCondition) {
      this.resolveCondition();
      this.clearPromises();
    }
  }
  // 返回一个 Promise，等待 condition 变为 true
  waitForCondition() {
    return new Promise((resolve, reject) => {
      if (this.condition) {
        resolve();
      } else {
        this.resolveCondition = resolve;
        this.rejectCondition = reject;
      }
    });
  }
  // 清理 resolve/reject 函数
  clearPromises() {
    this.resolveCondition = null;
    this.rejectCondition = null;
  }
}

async function to(promise, errorExt) {
  try {
    const data = await promise;
    const result = [null, data];
    return result;
  } catch (error) {
    if (errorExt) {
      const parsedError = Object.assign({}, error, errorExt);
      return [parsedError, void 0];
    }
    return [error, void 0];
  }
}

function traverseTreeValues(tree, getValue, options) {
  const result = [];
  const { childProps } = options || {
    childProps: "children"
  };
  const dfs = (treeNode) => {
    const value = getValue(treeNode);
    result.push(value);
    const children = treeNode?.[childProps];
    if (!children) {
      return;
    }
    if (children.length > 0) {
      for (const child of children) {
        dfs(child);
      }
    }
  };
  for (const treeNode of tree) {
    dfs(treeNode);
  }
  return result.filter(Boolean);
}
function filterTree(tree, filter, options) {
  const { childProps } = options || {
    childProps: "children"
  };
  const _filterTree = (nodes) => {
    return nodes.filter((node) => {
      if (filter(node)) {
        if (node[childProps]) {
          node[childProps] = _filterTree(node[childProps]);
        }
        return true;
      }
      return false;
    });
  };
  return _filterTree(tree);
}
function mapTree(tree, mapper, options) {
  const { childProps } = options || {
    childProps: "children"
  };
  return tree.map((node) => {
    const mapperNode = mapper(node);
    if (mapperNode[childProps]) {
      mapperNode[childProps] = mapTree(mapperNode[childProps], mapper, options);
    }
    return mapperNode;
  });
}
function sortTree(treeData, sortFunction, options) {
  const { childProps } = options || {
    childProps: "children"
  };
  return treeData.toSorted(sortFunction).map((item) => {
    const children = item[childProps];
    if (children && Array.isArray(children) && children.length > 0) {
      return {
        ...item,
        [childProps]: sortTree(children, sortFunction, options)
      };
    }
    return item;
  });
}

function uniqueByField(arr, key) {
  const seen = /* @__PURE__ */ new Map();
  return arr.filter((item) => {
    const value = item[key];
    return seen.has(value) ? false : (seen.set(value, item), true);
  });
}

function updateCSSVariables(variables, id = "__vben-styles__") {
  const styleElement = document.querySelector(`#${id}`) || document.createElement("style");
  styleElement.id = id;
  let cssText = ":root {";
  for (const key in variables) {
    if (Object.prototype.hasOwnProperty.call(variables, key)) {
      cssText += `${key}: ${variables[key]};`;
    }
  }
  cssText += "}";
  styleElement.textContent = cssText;
  if (!document.querySelector(`#${id}`)) {
    setTimeout(() => {
      document.head.append(styleElement);
    });
  }
}

function bindMethods(instance) {
  const prototype = Object.getPrototypeOf(instance);
  const propertyNames = Object.getOwnPropertyNames(prototype);
  propertyNames.forEach((propertyName) => {
    const descriptor = Object.getOwnPropertyDescriptor(prototype, propertyName);
    const propertyValue = instance[propertyName];
    if (typeof propertyValue === "function" && propertyName !== "constructor" && descriptor && !descriptor.get && !descriptor.set) {
      instance[propertyName] = propertyValue.bind(instance);
    }
  });
}
function getNestedValue(obj, path) {
  if (typeof path !== "string" || path.length === 0) {
    throw new Error("Path must be a non-empty string");
  }
  const keys = path.split(".");
  let current = obj;
  for (const key of keys) {
    if (current === null || current === void 0) {
      return void 0;
    }
    current = current[key];
  }
  return current;
}

export { StateHandler, arraysEqual, bindMethods, capitalizeFirstLetter, cn, diff, downloadFileFromBase64, downloadFileFromBlob, downloadFileFromBlobPart, downloadFileFromImageUrl, downloadFileFromUrl, filterTree, formatDate, formatDateTime, getCurrentTimezone, getElementVisibleRect, getFirstNonNullOrUndefined, getNestedValue, getScrollbarWidth, getSystemTimezone, isBoolean, isDate, isDayjsObject, isEmpty, isHttpUrl, isMacOs, isNumber, isUndefined, isWindow, isWindowsOs, kebabToCamelCase, loadScript, mapTree, mergeWithArrayOverride, needsScrollbar, openRouteInNewWindow, openWindow, setCurrentTimezone, sortTree, startProgress, stopProgress, to, toCamelCase, toLowerCaseFirstLetter, traverseTreeValues, triggerDownload, triggerWindowResize, uniqueByField, updateCSSVariables, urlToBase64 };
