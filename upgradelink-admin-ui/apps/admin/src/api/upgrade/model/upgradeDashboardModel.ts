import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: 7 days device count data
 */
export interface DevCount7DaySeriesItem {
  data: number[];
  appName: string;
}

/**
 *  @description: 7 days device count info
 */
export interface DevCount7Day {
  timeData: string[];
  series: DevCount7DaySeriesItem[];
}

/**
 *  @description: 7 days request count data
 */
export interface ReqCount7DaySeriesItem {
  data: number[];
  appName: string;
}

/**
 *  @description: 7 days request count info
 */
export interface ReqCount7Day {
  timeData: string[];
  series: ReqCount7DaySeriesItem[];
}

/**
 *  @description: 7 days request count data
 */
export interface DownloadCount7DaySeriesItem {
  data: number[];
  appName: string;
}

/**
 *  @description: 7 days request count info
 */
export interface DownloadCount7Day {
  timeData: string[];
  series: DownloadCount7DaySeriesItem[];
}

/**
 *  @description: 7 days request count data
 */
export interface TrafficUsageCount7DaySeriesItem {
  data: number[];
  appName: string;
}

/**
 *  @description: 7 days request count info
 */
export interface TrafficUsageCount7Day {
  timeData: string[];
  series: TrafficUsageCount7DaySeriesItem[];
}

/**
 *  @description: 7 days request count data
 */
export interface AppStartCount7DaySeriesItem {
  data: number[];
  appName: string;
}

/**
 *  @description: 7 days request count info
 */
export interface AppStartCount7Day {
  timeData: string[];
  series: AppStartCount7DaySeriesItem[];
}

/**
 *  @description: 7 days request count data
 */
export interface AppGetStrategyCount7DaySeriesItem {
  data: number[];
  appName: string;
}

/**
 *  @description: 7 days request count info
 */
export interface AppGetStrategyCount7Day {
  timeData: string[];
  series: AppGetStrategyCount7DaySeriesItem[];
}

/**
 *  @description: 7 days request count data
 */
export interface AppUpgradeCount7DaySeriesItem {
  data: number[];
  appName: string;
}

/**
 *  @description: 7 days request count info
 */
export interface AppUpgradeCount7Day {
  timeData: string[];
  series: AppUpgradeCount7DaySeriesItem[];
}

/**
 *  @description: UpgradeDashboard info response
 */
export interface UpgradeDashboardInfo {
  yesterdayNewDevCount?: number;
  totalDevCount?: number;
  yesterdayReqCount?: number;
  totalReqCount?: number;
  yesterdayDownloadCount?: number;
  totalDownloadCount?: number;
  yesterdayAppGetStrategyCount?: number;
  totalAppGetStrategyCount?: number;
  yesterdayAppStartCount?: number;
  totalAppStartCount?: number;
  yesterdayAppUpgradeCount?: number;
  totalAppUpgradeCount?: number;
  devCount7Day?: DevCount7Day;
  reqCount7Day?: ReqCount7Day;
  trafficUsageCount7Day?: TrafficUsageCount7Day;
  downloadCount7Day?: DownloadCount7Day;
  appGetStrategyCount7Day?: AppGetStrategyCount7Day;
  appStartCount7Day?: AppStartCount7Day;
  appUpgradeCount7Day?: AppUpgradeCount7Day;
}

/**
 *  @description: UpgradeDashboard list response
 */

export type UpgradeDashboardListResp = BaseListResp<UpgradeDashboardInfo>;
