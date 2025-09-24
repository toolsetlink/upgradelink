import type { BaseListResp } from "#/api/model/baseModel";

/**
 *  @description: UpgradeDevModel info response
 */
export interface UpgradeDevModelInfo {
  id?: number;
  companyId?: number;
  key?: string;
  name?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeDevModel list response
 */

export type UpgradeDevModelListResp = BaseListResp<UpgradeDevModelInfo>;
