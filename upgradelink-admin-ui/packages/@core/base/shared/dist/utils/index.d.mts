import { ClassValue } from 'clsx';
import dayjs from 'dayjs';
import * as defu from 'defu';
export { createDefu as createMerge, defu as merge } from 'defu';
export { get, isEqual, set } from 'es-toolkit/compat';
export { default as cloneDeep } from 'lodash.clonedeep';
export { isFunction, isObject, isString } from '@vue/shared';

declare function cn(...inputs: ClassValue[]): string;

type FormatDate = Date | dayjs.Dayjs | number | string;
type Format = 'HH' | 'HH:mm' | 'HH:mm:ss' | 'YYYY' | 'YYYY-MM' | 'YYYY-MM-DD' | 'YYYY-MM-DD HH' | 'YYYY-MM-DD HH:mm' | 'YYYY-MM-DD HH:mm:ss' | (string & {});
declare function formatDate(time?: FormatDate, format?: Format): string;
declare function formatDateTime(time?: FormatDate): string;
declare function isDate(value: any): value is Date;
declare function isDayjsObject(value: any): value is dayjs.Dayjs;
declare const getSystemTimezone: () => string;
declare const setCurrentTimezone: (timezone?: string) => void;
declare const getCurrentTimezone: () => string;

declare function arraysEqual<T>(a: T[], b: T[]): boolean;
type DiffResult<T> = Partial<{
    [K in keyof T]: T[K] extends object ? DiffResult<T[K]> : T[K];
}>;
declare function diff<T extends Record<string, any>>(obj1: T, obj2: T): DiffResult<T>;

interface VisibleDomRect {
    bottom: number;
    height: number;
    left: number;
    right: number;
    top: number;
    width: number;
}
declare function getElementVisibleRect(element?: HTMLElement | null | undefined): VisibleDomRect;
declare function getScrollbarWidth(): number;
declare function needsScrollbar(): boolean;
declare function triggerWindowResize(): void;

interface DownloadOptions<T = string> {
    fileName?: string;
    source: T;
    target?: string;
}
declare function downloadFileFromUrl({ fileName, source, target, }: DownloadOptions): Promise<void>;
declare function downloadFileFromBase64({ fileName, source }: DownloadOptions): void;
declare function downloadFileFromImageUrl({ fileName, source, }: DownloadOptions): Promise<void>;
declare function downloadFileFromBlob({ fileName, source, }: DownloadOptions<Blob>): void;
declare function downloadFileFromBlobPart({ fileName, source, }: DownloadOptions<BlobPart>): void;
declare function urlToBase64(url: string, mineType?: string): Promise<string>;
declare function triggerDownload(href: string, fileName: string | undefined, revokeDelay?: number): void;

declare function isUndefined(value?: unknown): value is undefined;
declare function isBoolean(value: unknown): value is boolean;
declare function isEmpty<T = unknown>(value?: T): value is T;
declare function isHttpUrl(url?: string): boolean;
declare function isWindow(value: any): value is Window;
declare function isMacOs(): boolean;
declare function isWindowsOs(): boolean;
declare function isNumber(value: any): value is number;
declare function getFirstNonNullOrUndefined<T>(...values: (null | T | undefined)[]): T | undefined;

declare function capitalizeFirstLetter(string: string): string;
declare function toLowerCaseFirstLetter(str: string): string;
declare function toCamelCase(key: string, parentKey: string): string;
declare function kebabToCamelCase(str: string): string;

declare const mergeWithArrayOverride: <Source extends {
    [x: string]: any;
    [x: number]: any;
    [x: symbol]: any;
}, Defaults extends Array<{
    [x: string]: any;
    [x: number]: any;
    [x: symbol]: any;
} | (number | boolean | any[] | Record<never, any> | null | undefined)>>(source: Source, ...defaults: Defaults) => defu.Defu<Source, Defaults>;

declare function startProgress(): Promise<void>;
declare function stopProgress(): Promise<void>;

declare function loadScript(src: string): Promise<void>;

declare class StateHandler {
    private condition;
    private rejectCondition;
    private resolveCondition;
    isConditionTrue(): boolean;
    reset(): void;
    setConditionFalse(): void;
    setConditionTrue(): void;
    waitForCondition(): Promise<void>;
    private clearPromises;
}

declare function to<T, U = Error>(promise: Readonly<Promise<T>>, errorExt?: object): Promise<[null, T] | [U, undefined]>;

interface TreeConfigOptions {
    childProps: string;
}
declare function traverseTreeValues<T, V>(tree: T[], getValue: (node: T) => V, options?: TreeConfigOptions): V[];
declare function filterTree<T extends Record<string, any>>(tree: T[], filter: (node: T) => boolean, options?: TreeConfigOptions): T[];
declare function mapTree<T, V extends Record<string, any>>(tree: T[], mapper: (node: T) => V, options?: TreeConfigOptions): V[];
declare function sortTree<T extends Record<string, any>>(treeData: T[], sortFunction: (a: T, b: T) => number, options?: TreeConfigOptions): T[];

declare function uniqueByField<T>(arr: T[], key: keyof T): T[];

declare function updateCSSVariables(variables: {
    [key: string]: string;
}, id?: string): void;

declare function bindMethods<T extends object>(instance: T): void;
declare function getNestedValue<T>(obj: T, path: string): any;

interface OpenWindowOptions {
    noopener?: boolean;
    noreferrer?: boolean;
    target?: '_blank' | '_parent' | '_self' | '_top' | string;
}
declare function openWindow(url: string, options?: OpenWindowOptions): void;
declare function openRouteInNewWindow(path: string): void;

export { StateHandler, arraysEqual, bindMethods, capitalizeFirstLetter, cn, diff, downloadFileFromBase64, downloadFileFromBlob, downloadFileFromBlobPart, downloadFileFromImageUrl, downloadFileFromUrl, filterTree, formatDate, formatDateTime, getCurrentTimezone, getElementVisibleRect, getFirstNonNullOrUndefined, getNestedValue, getScrollbarWidth, getSystemTimezone, isBoolean, isDate, isDayjsObject, isEmpty, isHttpUrl, isMacOs, isNumber, isUndefined, isWindow, isWindowsOs, kebabToCamelCase, loadScript, mapTree, mergeWithArrayOverride, needsScrollbar, openRouteInNewWindow, openWindow, setCurrentTimezone, sortTree, startProgress, stopProgress, to, toCamelCase, toLowerCaseFirstLetter, traverseTreeValues, triggerDownload, triggerWindowResize, uniqueByField, updateCSSVariables, urlToBase64 };
export type { VisibleDomRect };
