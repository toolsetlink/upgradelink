import type {
  UpgradeUrlVersionInfo,
  UpgradeUrlVersionListResp,
} from "./model/upgradeUrlVersionModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeUrlVersion = "/upgrade/upgrade_url_version/create",
  DeleteUpgradeUrlVersion = "/upgrade/upgrade_url_version/delete",
  GetUpgradeUrlVersionById = "/upgrade/upgrade_url_version",
  GetUpgradeUrlVersionList = "/upgrade/upgrade_url_version/list",
  UpdateUpgradeUrlVersion = "/upgrade/upgrade_url_version/update",
}

/**
 * @description: Get upgrade url version list
 */

export const getUpgradeUrlVersionList = (params: {
  page: number;
  pageSize: number;
  urlId: any | undefined;
}) => {
  return requestClient.post<BaseDataResp<UpgradeUrlVersionListResp>>(
    Api.GetUpgradeUrlVersionList,
    params,
  );
};

/**
 *  @description: Create a new upgrade url version
 */
export const createUpgradeUrlVersion = (params: UpgradeUrlVersionInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeUrlVersion, params);
};

/**
 *  @description: Update the upgrade url version
 */
export const updateUpgradeUrlVersion = (params: UpgradeUrlVersionInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeUrlVersion, params);
};

/**
 *  @description: Delete upgrade url versions
 */
export const deleteUpgradeUrlVersion = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeUrlVersion, params);
};

/**
 *  @description: Get upgrade url version By ID
 */
export const getUpgradeUrlVersionById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeUrlVersionInfo>>(
    Api.GetUpgradeUrlVersionById,
    params,
  );
};
