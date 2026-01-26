import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeMacVersion info response
 */
export interface UpgradeMacVersionInfo {
  id?: number;
  macId?: number;
  cloudFileId?: string;
  versionName?: string;
  versionCode?: number;
  versionFileSize?: string;
  description?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeMacVersion list response
 */

export type UpgradeMacVersionListResp = BaseListResp<UpgradeMacVersionInfo>;
