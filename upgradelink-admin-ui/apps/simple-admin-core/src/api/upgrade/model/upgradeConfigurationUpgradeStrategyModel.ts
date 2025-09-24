import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeConfigurationUpgradeStrategy info response
 */
export interface UpgradeConfigurationUpgradeStrategyInfo {
  id?: number;
  enable?: number;
  name?: string;
  description?: string;
  configurationId?: number;
  configurationVersionId?: number;
  beginDatetime?: string;
  endDatetime?: string;
  upgradeType?: number;
  promptUpgradeContent?: string;
  upgradeDevType?: number;
  upgradeDevData?: string;
  upgradeDevTypeZeroData?: string;
  upgradeDevTypeOneData?: number[];
  upgradeDevTypeTwoData?: number[];
  upgradeVersionType?: number;
  upgradeVersionData?: string;
  upgradeVersionTypeZeroData?: string;
  upgradeVersionTypeOneData?: number[];
  isGray?: number;
}

/**
 *  @description: UpgradeConfigurationUpgradeStrategy list response
 */

export type UpgradeConfigurationUpgradeStrategyListResp =
  BaseListResp<UpgradeConfigurationUpgradeStrategyInfo>;
