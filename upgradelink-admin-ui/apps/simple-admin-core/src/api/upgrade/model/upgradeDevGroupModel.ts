import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeDevGroup info response
 */
export interface UpgradeDevGroupInfo {
  id?: number;
  name?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeDevGroup list response
 */

export type UpgradeDevGroupListResp = BaseListResp<UpgradeDevGroupInfo>;
