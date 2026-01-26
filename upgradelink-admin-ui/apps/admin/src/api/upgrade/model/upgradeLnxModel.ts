import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeLnx info response
 */
export interface UpgradeLnxInfo {
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
 *  @description: UpgradeLnx list response
 */

export type UpgradeLnxListResp = BaseListResp<UpgradeLnxInfo>;
