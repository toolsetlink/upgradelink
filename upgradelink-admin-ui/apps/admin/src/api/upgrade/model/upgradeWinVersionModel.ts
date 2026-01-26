import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeWinVersion info response
 */
export interface UpgradeWinVersionInfo {
  id?: number;
  winId?: number;
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
 *  @description: UpgradeWinVersion list response
 */

export type UpgradeWinVersionListResp = BaseListResp<UpgradeWinVersionInfo>;
