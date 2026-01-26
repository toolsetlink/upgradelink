import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeMac info response
 */
export interface UpgradeMacInfo {
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
 *  @description: UpgradeMac list response
 */

export type UpgradeMacListResp = BaseListResp<UpgradeMacInfo>;
