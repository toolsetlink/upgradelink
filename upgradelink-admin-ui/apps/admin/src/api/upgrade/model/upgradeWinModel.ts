import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeWin info response
 */
export interface UpgradeWinInfo {
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
 *  @description: UpgradeWin list response
 */

export type UpgradeWinListResp = BaseListResp<UpgradeWinInfo>;
