import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: UpgradeLnxUpgradeStrategy info response
 */
export interface UpgradeLnxUpgradeStrategyInfo {
  id?: number;
  enable?: number;
  name?: string;
  description?: string;
  lnxId?: number;
  lnxVersionId?: number;
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
  grayDataList?: []*LnxGrayDataInfo;
  isFlowLimit?: number;
  flowLimitDataList?: []*LnxFlowLimitDataInfo;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeLnxUpgradeStrategy list response
 */

export type UpgradeLnxUpgradeStrategyListResp = BaseListResp<UpgradeLnxUpgradeStrategyInfo>;
