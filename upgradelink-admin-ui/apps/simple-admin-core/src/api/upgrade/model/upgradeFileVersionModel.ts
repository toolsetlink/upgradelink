import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeFileVersion info response
 */
export interface UpgradeFileVersionInfo {
  id?: number;
  fileId?: number;
  cloudFileId?: string;
  versionName?: string;
  versionCode?: number;
  description?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeFileVersion list response
 */

export type UpgradeFileVersionListResp = BaseListResp<UpgradeFileVersionInfo>;
