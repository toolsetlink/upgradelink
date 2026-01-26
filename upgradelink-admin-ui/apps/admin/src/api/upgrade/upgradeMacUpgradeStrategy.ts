import type {
  UpgradeMacUpgradeStrategyInfo,
  UpgradeMacUpgradeStrategyListResp,
} from "./model/upgradeMacUpgradeStrategyModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeMacUpgradeStrategy = "/upgrade_mac_upgrade_strategy/create",
  DeleteUpgradeMacUpgradeStrategy = "/upgrade_mac_upgrade_strategy/delete",
  GetUpgradeMacUpgradeStrategyById = "/upgrade_mac_upgrade_strategy",
  GetUpgradeMacUpgradeStrategyList = "/upgrade_mac_upgrade_strategy/list",
  UpdateUpgradeMacUpgradeStrategy = "/upgrade_mac_upgrade_strategy/update",
}

/**
 * @description: Get upgrade mac upgrade strategy list
 */

export const getUpgradeMacUpgradeStrategyList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeMacUpgradeStrategyListResp>>(
    Api.GetUpgradeMacUpgradeStrategyList,
    params,
  );
};

/**
 *  @description: Create a new upgrade mac upgrade strategy
 */
export const createUpgradeMacUpgradeStrategy = (
  params: UpgradeMacUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeMacUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Update the upgrade mac upgrade strategy
 */
export const updateUpgradeMacUpgradeStrategy = (
  params: UpgradeMacUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeMacUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Delete upgrade mac upgrade strategys
 */
export const deleteUpgradeMacUpgradeStrategy = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeMacUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Get upgrade mac upgrade strategy By ID
 */
export const getUpgradeMacUpgradeStrategyById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeMacUpgradeStrategyInfo>>(
    Api.GetUpgradeMacUpgradeStrategyById,
    params,
  );
};
