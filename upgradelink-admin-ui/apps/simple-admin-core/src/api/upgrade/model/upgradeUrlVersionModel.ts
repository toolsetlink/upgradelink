import type { BaseListResp } from "#/api/model/baseModel";

/**
 *  @description: UpgradeUrlVersion info response
 */
export interface UpgradeUrlVersionInfo {
  id?: number;
  companyId?: number;
  urlId?: number;
  urlName?: string;
  urlPath?: string;
  versionName?: string;
  versionCode?: number;
  description?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeUrlVersion list response
 */

export type UpgradeUrlVersionListResp = BaseListResp<UpgradeUrlVersionInfo>;
