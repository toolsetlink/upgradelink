import type {
  UpgradeConfigurationVersionInfo,
  UpgradeConfigurationVersionListResp,
} from "./model/upgradeConfigurationVersionModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeConfigurationVersion = "/upgrade/upgrade_configuration_version/create",
  DeleteUpgradeConfigurationVersion = "/upgrade/upgrade_configuration_version/delete",
  GetUpgradeConfigurationVersionById = "/upgrade/upgrade_configuration_version",
  GetUpgradeConfigurationVersionList = "/upgrade/upgrade_configuration_version/list",
  UpdateUpgradeConfigurationVersion = "/upgrade/upgrade_configuration_version/update",
}

/**
 * @description: Get upgrade configuration version list
 */

export const getUpgradeConfigurationVersionList = (params: {
  configurationId: any | undefined;
  page: number;
  pageSize: number;
}) => {
  return requestClient.post<BaseDataResp<UpgradeConfigurationVersionListResp>>(
    Api.GetUpgradeConfigurationVersionList,
    params,
  );
};

/**
 *  @description: Create a new upgrade configuration version
 */
export const createUpgradeConfigurationVersion = (
  params: UpgradeConfigurationVersionInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeConfigurationVersion,
    params,
  );
};

/**
 *  @description: Update the upgrade configuration version
 */
export const updateUpgradeConfigurationVersion = (
  params: UpgradeConfigurationVersionInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeConfigurationVersion,
    params,
  );
};

/**
 *  @description: Delete upgrade configuration versions
 */
export const deleteUpgradeConfigurationVersion = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeConfigurationVersion,
    params,
  );
};

/**
 *  @description: Get upgrade configuration version By ID
 */
export const getUpgradeConfigurationVersionById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeConfigurationVersionInfo>>(
    Api.GetUpgradeConfigurationVersionById,
    params,
  );
};
