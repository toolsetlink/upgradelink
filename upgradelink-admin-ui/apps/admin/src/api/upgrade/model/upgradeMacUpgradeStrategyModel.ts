import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: UpgradeMacUpgradeStrategy info response
 */
export interface UpgradeMacUpgradeStrategyInfo {
  id?: number;
  enable?: number;
  name?: string;
  description?: string;
  macId?: number;
  macVersionId?: number;
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
  beginDatetime?: string;
  endDatetime?: string;
  isGray?: number;
  grayDataList?: []*MacGrayDataInfo;
  isFlowLimit?: number;
  flowLimitDataList?: []*MacFlowLimitDataInfo;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeMacUpgradeStrategy list response
 */

export type UpgradeMacUpgradeStrategyListResp = BaseListResp<UpgradeMacUpgradeStrategyInfo>;
