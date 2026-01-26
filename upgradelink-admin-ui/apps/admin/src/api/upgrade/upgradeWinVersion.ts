import type {
  UpgradeWinVersionInfo,
  UpgradeWinVersionListResp,
} from "./model/upgradeWinVersionModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeWinVersion = "/upgrade_win_version/create",
  DeleteUpgradeWinVersion = "/upgrade_win_version/delete",
  GetUpgradeWinVersionById = "/upgrade_win_version",
  GetUpgradeWinVersionList = "/upgrade_win_version/list",
  UpdateUpgradeWinVersion = "/upgrade_win_version/update",
}

/**
 * @description: Get upgrade win version list
 */

export const getUpgradeWinVersionList = (params: {
  page: number;
  pageSize: number;
  winId: any | undefined;
}) => {
  return requestClient.post<BaseDataResp<UpgradeWinVersionListResp>>(
    Api.GetUpgradeWinVersionList,
    params,
  );
};
/**
 *  @description: Create a new upgrade win version
 */
export const createUpgradeWinVersion = (params: UpgradeWinVersionInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeWinVersion, params);
};

/**
 *  @description: Update the upgrade win version
 */
export const updateUpgradeWinVersion = (params: UpgradeWinVersionInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeWinVersion, params);
};

/**
 *  @description: Delete upgrade win versions
 */
export const deleteUpgradeWinVersion = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeWinVersion, params);
};

/**
 *  @description: Get upgrade win version By ID
 */
export const getUpgradeWinVersionById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeWinVersionInfo>>(
    Api.GetUpgradeWinVersionById,
    params,
  );
};
