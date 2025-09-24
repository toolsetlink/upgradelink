import type {
  UpgradeElectronVersionInfo,
  UpgradeElectronVersionListResp,
} from "./model/upgradeElectronVersionModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeElectronVersion = "/upgrade/upgrade_electron_version/create",
  DeleteUpgradeElectronVersion = "/upgrade/upgrade_electron_version/delete",
  GetUpgradeElectronVersionById = "/upgrade/upgrade_electron_version",
  GetUpgradeElectronVersionList = "/upgrade/upgrade_electron_version/list",
  UpdateUpgradeElectronVersion = "/upgrade/upgrade_electron_version/update",
}

/**
 * @description: Get upgrade electron version list
 */

export const getUpgradeElectronVersionList = (params: {
  electronId: any | undefined;
  page: number;
  pageSize: number;
}) => {
  return requestClient.post<BaseDataResp<UpgradeElectronVersionListResp>>(
    Api.GetUpgradeElectronVersionList,
    params,
  );
};

/**
 *  @description: Create a new upgrade electron version
 */
export const createUpgradeElectronVersion = (
  params: UpgradeElectronVersionInfo,
) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeElectronVersion, params);
};

/**
 *  @description: Update the upgrade electron version
 */
export const updateUpgradeElectronVersion = (
  params: UpgradeElectronVersionInfo,
) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeElectronVersion, params);
};

/**
 *  @description: Delete upgrade electron versions
 */
export const deleteUpgradeElectronVersion = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeElectronVersion, params);
};

/**
 *  @description: Get upgrade electron version By ID
 */
export const getUpgradeElectronVersionById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeElectronVersionInfo>>(
    Api.GetUpgradeElectronVersionById,
    params,
  );
};
