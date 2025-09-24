import type {
  UpgradeUrlUpgradeStrategyInfo,
  UpgradeUrlUpgradeStrategyListResp,
} from "./model/upgradeUrlUpgradeStrategyModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeUrlUpgradeStrategy = "/upgrade/upgrade_url_upgrade_strategy/create",
  DeleteUpgradeUrlUpgradeStrategy = "/upgrade/upgrade_url_upgrade_strategy/delete",
  GetUpgradeUrlUpgradeStrategyById = "/upgrade/upgrade_url_upgrade_strategy",
  GetUpgradeUrlUpgradeStrategyList = "/upgrade/upgrade_url_upgrade_strategy/list",
  UpdateUpgradeUrlUpgradeStrategy = "/upgrade/upgrade_url_upgrade_strategy/update",
}

/**
 * @description: Get upgrade url upgrade strategy list
 */

export const getUpgradeUrlUpgradeStrategyList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeUrlUpgradeStrategyListResp>>(
    Api.GetUpgradeUrlUpgradeStrategyList,
    params,
  );
};

/**
 *  @description: Create a new upgrade url upgrade strategy
 */
export const createUpgradeUrlUpgradeStrategy = (
  params: UpgradeUrlUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeUrlUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Update the upgrade url upgrade strategy
 */
export const updateUpgradeUrlUpgradeStrategy = (
  params: UpgradeUrlUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeUrlUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Delete upgrade url upgrade strategys
 */
export const deleteUpgradeUrlUpgradeStrategy = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeUrlUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Get upgrade url upgrade strategy By ID
 */
export const getUpgradeUrlUpgradeStrategyById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeUrlUpgradeStrategyInfo>>(
    Api.GetUpgradeUrlUpgradeStrategyById,
    params,
  );
};
