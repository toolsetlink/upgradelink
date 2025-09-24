import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeElectron info response
 */
export interface UpgradeElectronInfo {
  id?: number;
  key?: string;
  name?: string;
  description?: string;
  githubUrl?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeElectron list response
 */

export type UpgradeElectronListResp = BaseListResp<UpgradeElectronInfo>;
