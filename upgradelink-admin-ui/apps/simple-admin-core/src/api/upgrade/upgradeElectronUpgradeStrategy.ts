import type {
  UpgradeElectronUpgradeStrategyInfo,
  UpgradeElectronUpgradeStrategyListResp,
} from "./model/upgradeElectronUpgradeStrategyModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeElectronUpgradeStrategy = "/upgrade/upgrade_electron_upgrade_strategy/create",
  DeleteUpgradeElectronUpgradeStrategy = "/upgrade/upgrade_electron_upgrade_strategy/delete",
  GetUpgradeElectronUpgradeStrategyById = "/upgrade/upgrade_electron_upgrade_strategy",
  GetUpgradeElectronUpgradeStrategyList = "/upgrade/upgrade_electron_upgrade_strategy/list",
  UpdateUpgradeElectronUpgradeStrategy = "/upgrade/upgrade_electron_upgrade_strategy/update",
}

/**
 * @description: Get upgrade electron upgrade strategy list
 */

export const getUpgradeElectronUpgradeStrategyList = (params: BaseListReq) => {
  return requestClient.post<
    BaseDataResp<UpgradeElectronUpgradeStrategyListResp>
  >(Api.GetUpgradeElectronUpgradeStrategyList, params);
};

/**
 *  @description: Create a new upgrade electron upgrade strategy
 */
export const createUpgradeElectronUpgradeStrategy = (
  params: UpgradeElectronUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeElectronUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Update the upgrade electron upgrade strategy
 */
export const updateUpgradeElectronUpgradeStrategy = (
  params: UpgradeElectronUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeElectronUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Delete upgrade electron upgrade strategys
 */
export const deleteUpgradeElectronUpgradeStrategy = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeElectronUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Get upgrade electron upgrade strategy By ID
 */
export const getUpgradeElectronUpgradeStrategyById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeElectronUpgradeStrategyInfo>>(
    Api.GetUpgradeElectronUpgradeStrategyById,
    params,
  );
};
