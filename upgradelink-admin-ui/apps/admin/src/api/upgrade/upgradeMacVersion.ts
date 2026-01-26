import type {
  UpgradeMacVersionInfo,
  UpgradeMacVersionListResp,
} from "./model/upgradeMacVersionModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeMacVersion = "/upgrade_mac_version/create",
  DeleteUpgradeMacVersion = "/upgrade_mac_version/delete",
  GetUpgradeMacVersionById = "/upgrade_mac_version",
  GetUpgradeMacVersionList = "/upgrade_mac_version/list",
  UpdateUpgradeMacVersion = "/upgrade_mac_version/update",
}

/**
 * @description: Get upgrade mac version list
 */

export const getUpgradeMacVersionList = (params: {
  macId: any | undefined;
  page: number;
  pageSize: number;
}) => {
  return requestClient.post<BaseDataResp<UpgradeMacVersionListResp>>(
    Api.GetUpgradeMacVersionList,
    params,
  );
};

/**
 *  @description: Create a new upgrade mac version
 */
export const createUpgradeMacVersion = (params: UpgradeMacVersionInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeMacVersion, params);
};

/**
 *  @description: Update the upgrade mac version
 */
export const updateUpgradeMacVersion = (params: UpgradeMacVersionInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeMacVersion, params);
};

/**
 *  @description: Delete upgrade mac versions
 */
export const deleteUpgradeMacVersion = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeMacVersion, params);
};

/**
 *  @description: Get upgrade mac version By ID
 */
export const getUpgradeMacVersionById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeMacVersionInfo>>(
    Api.GetUpgradeMacVersionById,
    params,
  );
};
