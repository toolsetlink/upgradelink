import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeCompanyIncome info response
 */
export interface UpgradeCompanyIncomeInfo {
  id?: number;
  companyId?: number;
  incomeType?: number;
  incomeAmount?: number;
  remark?: string;
  status?: number;
  isDel?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeCompanyIncome list response
 */

export type UpgradeCompanyIncomeListResp =
  BaseListResp<UpgradeCompanyIncomeInfo>;
