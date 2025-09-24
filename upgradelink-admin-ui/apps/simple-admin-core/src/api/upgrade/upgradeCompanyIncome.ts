import type {
  UpgradeCompanyIncomeInfo,
  UpgradeCompanyIncomeListResp,
} from "./model/upgradeCompanyIncomeModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeCompanyIncome = "/upgrade/upgrade_company_income/create",
  DeleteUpgradeCompanyIncome = "/upgrade/upgrade_company_income/delete",
  GetUpgradeCompanyIncomeById = "/upgrade/upgrade_company_income",
  GetUpgradeCompanyIncomeList = "/upgrade/upgrade_company_income/list",
  UpdateUpgradeCompanyIncome = "/upgrade/upgrade_company_income/update",
}

/**
 * @description: Get upgrade company income list
 */

export const getUpgradeCompanyIncomeList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeCompanyIncomeListResp>>(
    Api.GetUpgradeCompanyIncomeList,
    params,
  );
};

/**
 *  @description: Create a new upgrade company income
 */
export const createUpgradeCompanyIncome = (
  params: UpgradeCompanyIncomeInfo,
) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeCompanyIncome, params);
};

/**
 *  @description: Update the upgrade company income
 */
export const updateUpgradeCompanyIncome = (
  params: UpgradeCompanyIncomeInfo,
) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeCompanyIncome, params);
};

/**
 *  @description: Delete upgrade company incomes
 */
export const deleteUpgradeCompanyIncome = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeCompanyIncome, params);
};

/**
 *  @description: Get upgrade company income By ID
 */
export const getUpgradeCompanyIncomeById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeCompanyIncomeInfo>>(
    Api.GetUpgradeCompanyIncomeById,
    params,
  );
};
