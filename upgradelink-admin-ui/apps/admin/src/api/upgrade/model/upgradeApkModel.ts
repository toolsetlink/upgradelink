import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeApk info response
 */
export interface UpgradeApkInfo {
  id?: number;
  key?: string;
  name?: string;
  packageName?: string;
  description?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeApk list response
 */

export type UpgradeApkListResp = BaseListResp<UpgradeApkInfo>;
