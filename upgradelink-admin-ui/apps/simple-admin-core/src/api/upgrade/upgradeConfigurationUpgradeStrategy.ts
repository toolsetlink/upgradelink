import type {
  UpgradeConfigurationUpgradeStrategyInfo,
  UpgradeConfigurationUpgradeStrategyListResp,
} from "./model/upgradeConfigurationUpgradeStrategyModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeConfigurationUpgradeStrategy = "/upgrade/upgrade_configuration_upgrade_strategy/create",
  DeleteUpgradeConfigurationUpgradeStrategy = "/upgrade/upgrade_configuration_upgrade_strategy/delete",
  GetUpgradeConfigurationUpgradeStrategyById = "/upgrade/upgrade_configuration_upgrade_strategy",
  GetUpgradeConfigurationUpgradeStrategyList = "/upgrade/upgrade_configuration_upgrade_strategy/list",
  UpdateUpgradeConfigurationUpgradeStrategy = "/upgrade/upgrade_configuration_upgrade_strategy/update",
}

/**
 * @description: Get upgrade configuration upgrade strategy list
 */

export const getUpgradeConfigurationUpgradeStrategyList = (
  params: BaseListReq,
) => {
  return requestClient.post<
    BaseDataResp<UpgradeConfigurationUpgradeStrategyListResp>
  >(Api.GetUpgradeConfigurationUpgradeStrategyList, params);
};

/**
 *  @description: Create a new upgrade configuration upgrade strategy
 */
export const createUpgradeConfigurationUpgradeStrategy = (
  params: UpgradeConfigurationUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeConfigurationUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Update the upgrade configuration upgrade strategy
 */
export const updateUpgradeConfigurationUpgradeStrategy = (
  params: UpgradeConfigurationUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeConfigurationUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Delete upgrade configuration upgrade strategys
 */
export const deleteUpgradeConfigurationUpgradeStrategy = (
  params: BaseIDsReq,
) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeConfigurationUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Get upgrade configuration upgrade strategy By ID
 */
export const getUpgradeConfigurationUpgradeStrategyById = (
  params: BaseIDReq,
) => {
  return requestClient.post<
    BaseDataResp<UpgradeConfigurationUpgradeStrategyInfo>
  >(Api.GetUpgradeConfigurationUpgradeStrategyById, params);
};
