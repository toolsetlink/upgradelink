import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeTauriVersion info response
 */
export interface UpgradeTauriVersionInfo {
  id?: number;
  companyId?: number;
  tauriId?: number;
  cloudFileId?: string;
  versionName?: string;
  versionCode?: number;
  target?: string;
  arch?: string;
  description?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeTauriVersion list response
 */

export type UpgradeTauriVersionListResp = BaseListResp<UpgradeTauriVersionInfo>;
