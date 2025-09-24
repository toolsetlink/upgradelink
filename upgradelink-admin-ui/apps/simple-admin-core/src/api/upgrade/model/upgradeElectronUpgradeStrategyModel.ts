import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeElectronUpgradeStrategy info response
 */
export interface UpgradeElectronUpgradeStrategyInfo {
  id?: number;
  enable?: number;
  name?: string;
  description?: string;
  electronId?: number;
  electronVersionId?: number;
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
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeElectronUpgradeStrategy list response
 */

export type UpgradeElectronUpgradeStrategyListResp =
  BaseListResp<UpgradeElectronUpgradeStrategyInfo>;
