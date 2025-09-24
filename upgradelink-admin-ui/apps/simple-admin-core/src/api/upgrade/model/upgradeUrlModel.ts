import type { BaseListResp } from "#/api/model/baseModel";

/**
 *  @description: UpgradeUrl info response
 */
export interface UpgradeUrlInfo {
  id?: number;
  companyId?: number;
  key?: string;
  name?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeUrl list response
 */

export type UpgradeUrlListResp = BaseListResp<UpgradeUrlInfo>;
