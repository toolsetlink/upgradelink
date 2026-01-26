import type {
  UpgradeLnxVersionInfo,
  UpgradeLnxVersionListResp,
} from "./model/upgradeLnxVersionModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeLnxVersion = "/upgrade_lnx_version/create",
  DeleteUpgradeLnxVersion = "/upgrade_lnx_version/delete",
  GetUpgradeLnxVersionById = "/upgrade_lnx_version",
  GetUpgradeLnxVersionList = "/upgrade_lnx_version/list",
  UpdateUpgradeLnxVersion = "/upgrade_lnx_version/update",
}

/**
 * @description: Get upgrade lnx version list
 */

export const getUpgradeLnxVersionList = (params: {
  lnxId: any | undefined;
  page: number;
  pageSize: number;
}) => {
  return requestClient.post<BaseDataResp<UpgradeLnxVersionListResp>>(
    Api.GetUpgradeLnxVersionList,
    params,
  );
};

/**
 *  @description: Create a new upgrade lnx version
 */
export const createUpgradeLnxVersion = (params: UpgradeLnxVersionInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeLnxVersion, params);
};

/**
 *  @description: Update the upgrade lnx version
 */
export const updateUpgradeLnxVersion = (params: UpgradeLnxVersionInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeLnxVersion, params);
};

/**
 *  @description: Delete upgrade lnx versions
 */
export const deleteUpgradeLnxVersion = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeLnxVersion, params);
};

/**
 *  @description: Get upgrade lnx version By ID
 */
export const getUpgradeLnxVersionById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeLnxVersionInfo>>(
    Api.GetUpgradeLnxVersionById,
    params,
  );
};
