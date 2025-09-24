import type {
  UpgradeConfigurationInfo,
  UpgradeConfigurationListResp,
} from "./model/upgradeConfigurationModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeConfiguration = "/upgrade/upgrade_configuration/create",
  DeleteUpgradeConfiguration = "/upgrade/upgrade_configuration/delete",
  GetUpgradeConfigurationById = "/upgrade/upgrade_configuration",
  GetUpgradeConfigurationList = "/upgrade/upgrade_configuration/list",
  UpdateUpgradeConfiguration = "/upgrade/upgrade_configuration/update",
}

/**
 * @description: Get upgrade configuration list
 */

export const getUpgradeConfigurationList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeConfigurationListResp>>(
    Api.GetUpgradeConfigurationList,
    params,
  );
};

/**
 *  @description: Create a new upgrade configuration
 */
export const createUpgradeConfiguration = (
  params: UpgradeConfigurationInfo,
) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeConfiguration, params);
};

/**
 *  @description: Update the upgrade configuration
 */
export const updateUpgradeConfiguration = (
  params: UpgradeConfigurationInfo,
) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeConfiguration, params);
};

/**
 *  @description: Delete upgrade configurations
 */
export const deleteUpgradeConfiguration = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeConfiguration, params);
};

/**
 *  @description: Get upgrade configuration By ID
 */
export const getUpgradeConfigurationById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeConfigurationInfo>>(
    Api.GetUpgradeConfigurationById,
    params,
  );
};
