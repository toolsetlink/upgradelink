import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeElectronVersion info response
 */
export interface UpgradeElectronVersionInfo {
  id?: number;
  companyId?: number;
  electronId?: number;
  cloudFileId?: string;
  versionName?: string;
  versionCode?: number;
  platform?: string;
  target?: string;
  arch?: string;
  sha512?: string;
  description?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeElectronVersion list response
 */

export type UpgradeElectronVersionListResp =
  BaseListResp<UpgradeElectronVersionInfo>;
