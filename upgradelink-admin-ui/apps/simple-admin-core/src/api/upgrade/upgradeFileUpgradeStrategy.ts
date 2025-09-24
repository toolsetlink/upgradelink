import type {
  UpgradeFileUpgradeStrategyInfo,
  UpgradeFileUpgradeStrategyListResp,
} from "./model/upgradeFileUpgradeStrategyModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeFileUpgradeStrategy = "/upgrade/upgrade_file_upgrade_strategy/create",
  DeleteUpgradeFileUpgradeStrategy = "/upgrade/upgrade_file_upgrade_strategy/delete",
  GetUpgradeFileUpgradeStrategyById = "/upgrade/upgrade_file_upgrade_strategy",
  GetUpgradeFileUpgradeStrategyList = "/upgrade/upgrade_file_upgrade_strategy/list",
  UpdateUpgradeFileUpgradeStrategy = "/upgrade/upgrade_file_upgrade_strategy/update",
}

/**
 * @description: Get upgrade file upgrade strategy list
 */

export const getUpgradeFileUpgradeStrategyList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeFileUpgradeStrategyListResp>>(
    Api.GetUpgradeFileUpgradeStrategyList,
    params,
  );
};

/**
 *  @description: Create a new upgrade file upgrade strategy
 */
export const createUpgradeFileUpgradeStrategy = (
  params: UpgradeFileUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeFileUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Update the upgrade file upgrade strategy
 */
export const updateUpgradeFileUpgradeStrategy = (
  params: UpgradeFileUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeFileUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Delete upgrade file upgrade strategys
 */
export const deleteUpgradeFileUpgradeStrategy = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeFileUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Get upgrade file upgrade strategy By ID
 */
export const getUpgradeFileUpgradeStrategyById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeFileUpgradeStrategyInfo>>(
    Api.GetUpgradeFileUpgradeStrategyById,
    params,
  );
};
