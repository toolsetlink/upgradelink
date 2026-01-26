import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: UpgradeWinUpgradeStrategy info response
 */
export interface UpgradeWinUpgradeStrategyInfo {
  id?: number;
  enable?: number;
  name?: string;
  description?: string;
  winId?: number;
  winVersionId?: number;
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
  grayDataList?: []*WinGrayDataInfo;
  isFlowLimit?: number;
  flowLimitDataList?: []*WinFlowLimitDataInfo;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeWinUpgradeStrategy list response
 */

export type UpgradeWinUpgradeStrategyListResp = BaseListResp<UpgradeWinUpgradeStrategyInfo>;
