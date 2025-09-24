import type {
  UpgradeFileVersionInfo,
  UpgradeFileVersionListResp,
} from "./model/upgradeFileVersionModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeFileVersion = "/upgrade/upgrade_file_version/create",
  DeleteUpgradeFileVersion = "/upgrade/upgrade_file_version/delete",
  GetUpgradeFileVersionById = "/upgrade/upgrade_file_version",
  GetUpgradeFileVersionList = "/upgrade/upgrade_file_version/list",
  UpdateUpgradeFileVersion = "/upgrade/upgrade_file_version/update",
}

/**
 * @description: Get upgrade file version list
 */

export const getUpgradeFileVersionList = (params: {
  fileId: any | undefined;
  page: number;
  pageSize: number;
}) => {
  return requestClient.post<BaseDataResp<UpgradeFileVersionListResp>>(
    Api.GetUpgradeFileVersionList,
    params,
  );
};

/**
 *  @description: Create a new upgrade file version
 */
export const createUpgradeFileVersion = (params: UpgradeFileVersionInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeFileVersion, params);
};

/**
 *  @description: Update the upgrade file version
 */
export const updateUpgradeFileVersion = (params: UpgradeFileVersionInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeFileVersion, params);
};

/**
 *  @description: Delete upgrade file versions
 */
export const deleteUpgradeFileVersion = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeFileVersion, params);
};

/**
 *  @description: Get upgrade file version By ID
 */
export const getUpgradeFileVersionById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeFileVersionInfo>>(
    Api.GetUpgradeFileVersionById,
    params,
  );
};
