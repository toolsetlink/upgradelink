import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeFileUpgradeStrategy info response
 */
export interface UpgradeFileUpgradeStrategyInfo {
  id?: number;
  enable?: number;
  name?: string;
  description?: string;
  fileId?: number;
  fileVersionId?: number;
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
 *  @description: UpgradeFileUpgradeStrategy list response
 */

export type UpgradeFileUpgradeStrategyListResp =
  BaseListResp<UpgradeFileUpgradeStrategyInfo>;
