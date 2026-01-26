import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeConfigurationVersion info response
 */
export interface UpgradeConfigurationVersionInfo {
  id?: number;
  companyId?: number;
  configurationId?: number;
  content?: string;
  versionName?: string;
  versionCode?: number;
  description?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeConfigurationVersion list response
 */

export type UpgradeConfigurationVersionListResp =
  BaseListResp<UpgradeConfigurationVersionInfo>;
