import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeTauri info response
 */
export interface UpgradeTauriInfo {
  id?: number;
  companyId?: number;
  key?: string;
  name?: string;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeTauri list response
 */

export type UpgradeTauriListResp = BaseListResp<UpgradeTauriInfo>;
