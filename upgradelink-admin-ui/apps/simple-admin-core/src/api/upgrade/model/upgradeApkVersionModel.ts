import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeApkVersion info response
 */
export interface UpgradeApkVersionInfo {
  id?: number;
  apkId?: number;
  cloudFileId?: string;
  versionName?: string;
  versionCode?: number;
  description?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeApkVersion list response
 */

export type UpgradeApkVersionListResp = BaseListResp<UpgradeApkVersionInfo>;
