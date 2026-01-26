import type { BaseListResp } from "#/api/model/baseModel";

/**
 *  @description: UpgradeUrlUpgradeStrategy info response
 */
export interface UpgradeUrlUpgradeStrategyInfo {
  id?: number;
  companyId?: number;
  enable?: number;
  name?: string;
  description?: string;
  urlId?: number;
  urlVersionId?: number;
  urlVersionCode?: number;
  upgradeType?: number;
  promptUpgradeContent?: string;
  upgradeDevType?: number;
  upgradeDevData?: string;
  beginDatetime?: string;
  endDatetime?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeUrlUpgradeStrategy list response
 */

export type UpgradeUrlUpgradeStrategyListResp =
  BaseListResp<UpgradeUrlUpgradeStrategyInfo>;
