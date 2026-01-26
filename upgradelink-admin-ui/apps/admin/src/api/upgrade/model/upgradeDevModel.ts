import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeDev info response
 */
export interface UpgradeDevInfo {
  id?: number;
  key?: string;
  isDel?: number;
  devGroupIds?: number[];
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeDev list response
 */

export type UpgradeDevListResp = BaseListResp<UpgradeDevInfo>;
