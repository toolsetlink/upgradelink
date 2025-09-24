import type {
  UpgradeApkInfo,
  UpgradeApkListResp,
} from "./model/upgradeApkModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeApk = "/upgrade/upgrade_apk/create",
  DeleteUpgradeApk = "/upgrade/upgrade_apk/delete",
  GetUpgradeApkById = "/upgrade/upgrade_apk",
  GetUpgradeApkList = "/upgrade/upgrade_apk/list",
  UpdateUpgradeApk = "/upgrade/upgrade_apk/update",
}

/**
 * @description: Get upgrade apk list
 */

export const getUpgradeApkList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeApkListResp>>(
    Api.GetUpgradeApkList,
    params,
  );
};

/**
 *  @description: Create a new upgrade apk
 */
export const createUpgradeApk = (params: UpgradeApkInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeApk, params);
};

/**
 *  @description: Update the upgrade apk
 */
export const updateUpgradeApk = (params: UpgradeApkInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeApk, params);
};

/**
 *  @description: Delete upgrade apks
 */
export const deleteUpgradeApk = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeApk, params);
};

/**
 *  @description: Get upgrade apk By ID
 */
export const getUpgradeApkById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeApkInfo>>(
    Api.GetUpgradeApkById,
    params,
  );
};
