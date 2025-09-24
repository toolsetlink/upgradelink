import type {
  UpgradeUrlInfo,
  UpgradeUrlListResp,
} from "./model/upgradeUrlModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeUrl = "/upgrade/upgrade_url/create",
  DeleteUpgradeUrl = "/upgrade/upgrade_url/delete",
  GetUpgradeUrlById = "/upgrade/upgrade_url",
  GetUpgradeUrlList = "/upgrade/upgrade_url/list",
  UpdateUpgradeUrl = "/upgrade/upgrade_url/update",
}

/**
 * @description: Get upgrade url list
 */

export const getUpgradeUrlList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeUrlListResp>>(
    Api.GetUpgradeUrlList,
    params,
  );
};

/**
 *  @description: Create a new upgrade url
 */
export const createUpgradeUrl = (params: UpgradeUrlInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeUrl, params);
};

/**
 *  @description: Update the upgrade url
 */
export const updateUpgradeUrl = (params: UpgradeUrlInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeUrl, params);
};

/**
 *  @description: Delete upgrade urls
 */
export const deleteUpgradeUrl = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeUrl, params);
};

/**
 *  @description: Get upgrade url By ID
 */
export const getUpgradeUrlById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeUrlInfo>>(
    Api.GetUpgradeUrlById,
    params,
  );
};
