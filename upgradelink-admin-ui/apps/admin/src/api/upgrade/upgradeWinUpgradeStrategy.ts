import type {
  UpgradeWinUpgradeStrategyInfo,
  UpgradeWinUpgradeStrategyListResp,
} from "./model/upgradeWinUpgradeStrategyModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeWinUpgradeStrategy = "/upgrade_win_upgrade_strategy/create",
  DeleteUpgradeWinUpgradeStrategy = "/upgrade_win_upgrade_strategy/delete",
  GetUpgradeWinUpgradeStrategyById = "/upgrade_win_upgrade_strategy",
  GetUpgradeWinUpgradeStrategyList = "/upgrade_win_upgrade_strategy/list",
  UpdateUpgradeWinUpgradeStrategy = "/upgrade_win_upgrade_strategy/update",
}

/**
 * @description: Get upgrade win upgrade strategy list
 */

export const getUpgradeWinUpgradeStrategyList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeWinUpgradeStrategyListResp>>(
    Api.GetUpgradeWinUpgradeStrategyList,
    params,
  );
};

/**
 *  @description: Create a new upgrade win upgrade strategy
 */
export const createUpgradeWinUpgradeStrategy = (
  params: UpgradeWinUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeWinUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Update the upgrade win upgrade strategy
 */
export const updateUpgradeWinUpgradeStrategy = (
  params: UpgradeWinUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeWinUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Delete upgrade win upgrade strategys
 */
export const deleteUpgradeWinUpgradeStrategy = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeWinUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Get upgrade win upgrade strategy By ID
 */
export const getUpgradeWinUpgradeStrategyById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeWinUpgradeStrategyInfo>>(
    Api.GetUpgradeWinUpgradeStrategyById,
    params,
  );
};
