import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeLnxVersion info response
 */
export interface UpgradeLnxVersionInfo {
  id?: number;
  lnxId?: number;
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
 *  @description: UpgradeLnxVersion list response
 */

export type UpgradeLnxVersionListResp = BaseListResp<UpgradeLnxVersionInfo>;
