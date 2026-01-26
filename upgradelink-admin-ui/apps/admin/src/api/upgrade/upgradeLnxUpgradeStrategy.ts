import type {
  UpgradeLnxUpgradeStrategyInfo,
  UpgradeLnxUpgradeStrategyListResp,
} from "./model/upgradeLnxUpgradeStrategyModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeLnxUpgradeStrategy = "/upgrade_lnx_upgrade_strategy/create",
  DeleteUpgradeLnxUpgradeStrategy = "/upgrade_lnx_upgrade_strategy/delete",
  GetUpgradeLnxUpgradeStrategyById = "/upgrade_lnx_upgrade_strategy",
  GetUpgradeLnxUpgradeStrategyList = "/upgrade_lnx_upgrade_strategy/list",
  UpdateUpgradeLnxUpgradeStrategy = "/upgrade_lnx_upgrade_strategy/update",
}

/**
 * @description: Get upgrade lnx upgrade strategy list
 */

export const getUpgradeLnxUpgradeStrategyList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeLnxUpgradeStrategyListResp>>(
    Api.GetUpgradeLnxUpgradeStrategyList,
    params,
  );
};

/**
 *  @description: Create a new upgrade lnx upgrade strategy
 */
export const createUpgradeLnxUpgradeStrategy = (
  params: UpgradeLnxUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeLnxUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Update the upgrade lnx upgrade strategy
 */
export const updateUpgradeLnxUpgradeStrategy = (
  params: UpgradeLnxUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeLnxUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Delete upgrade lnx upgrade strategys
 */
export const deleteUpgradeLnxUpgradeStrategy = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeLnxUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Get upgrade lnx upgrade strategy By ID
 */
export const getUpgradeLnxUpgradeStrategyById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeLnxUpgradeStrategyInfo>>(
    Api.GetUpgradeLnxUpgradeStrategyById,
    params,
  );
};
