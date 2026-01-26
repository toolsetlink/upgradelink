import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: UpgradeApkUpgradeStrategy info response
 */
export interface UpgradeApkUpgradeStrategyInfo {
  id?: number;
  enable?: number;
  name?: string;
  description?: string;
  apkId?: number;
  apkVersionId?: number;
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
  grayDataList?: []*ApkGrayDataInfo;
  isFlowLimit?: number;
  flowLimitDataList?: []*ApkFlowLimitDataInfo;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeApkUpgradeStrategy list response
 */

export type UpgradeApkUpgradeStrategyListResp = BaseListResp<UpgradeApkUpgradeStrategyInfo>;
