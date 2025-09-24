import type {
  UpgradeApkVersionInfo,
  UpgradeApkVersionListResp,
} from "./model/upgradeApkVersionModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeApkVersion = "/upgrade/upgrade_apk_version/create",
  DeleteUpgradeApkVersion = "/upgrade/upgrade_apk_version/delete",
  GetUpgradeApkVersionById = "/upgrade/upgrade_apk_version",
  GetUpgradeApkVersionList = "/upgrade/upgrade_apk_version/list",
  UpdateUpgradeApkVersion = "/upgrade/upgrade_apk_version/update",
}

/**
 * @description: Get upgrade apk version list
 */

export const getUpgradeApkVersionList = (params: {
  apkId: any | undefined;
  page: number;
  pageSize: number;
}) => {
  return requestClient.post<BaseDataResp<UpgradeApkVersionListResp>>(
    Api.GetUpgradeApkVersionList,
    params,
  );
};

/**
 *  @description: Create a new upgrade apk version
 */
export const createUpgradeApkVersion = (params: UpgradeApkVersionInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeApkVersion, params);
};

/**
 *  @description: Update the upgrade apk version
 */
export const updateUpgradeApkVersion = (params: UpgradeApkVersionInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeApkVersion, params);
};

/**
 *  @description: Delete upgrade apk versions
 */
export const deleteUpgradeApkVersion = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeApkVersion, params);
};

/**
 *  @description: Get upgrade apk version By ID
 */
export const getUpgradeApkVersionById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeApkVersionInfo>>(
    Api.GetUpgradeApkVersionById,
    params,
  );
};
