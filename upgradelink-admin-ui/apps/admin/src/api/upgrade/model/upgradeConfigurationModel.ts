import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeConfiguration info response
 */
export interface UpgradeConfigurationInfo {
  id?: number;
  companyId?: number;
  key?: string;
  name?: string;
  description?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeConfiguration list response
 */

export type UpgradeConfigurationListResp =
  BaseListResp<UpgradeConfigurationInfo>;
