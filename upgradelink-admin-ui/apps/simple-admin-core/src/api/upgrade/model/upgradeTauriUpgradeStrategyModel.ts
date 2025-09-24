import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeTauriUpgradeStrategy info response
 */
export interface UpgradeTauriUpgradeStrategyInfo {
  id?: number;
  enable?: number;
  name?: string;
  description?: string;
  tauriId?: number;
  tauriVersionId?: number;
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
 *  @description: UpgradeTauriUpgradeStrategy list response
 */

export type UpgradeTauriUpgradeStrategyListResp =
  BaseListResp<UpgradeTauriUpgradeStrategyInfo>;
