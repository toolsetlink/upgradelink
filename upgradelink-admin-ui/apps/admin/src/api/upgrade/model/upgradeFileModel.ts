import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeFile info response
 */
export interface UpgradeFileInfo {
  id?: number;
  key?: string;
  name?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeFile list response
 */

export type UpgradeFileListResp = BaseListResp<UpgradeFileInfo>;
