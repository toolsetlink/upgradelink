import type {
  UpgradeLnxInfo,
  UpgradeLnxListResp,
} from "./model/upgradeLnxModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeLnx = "/upgrade_lnx/create",
  DeleteUpgradeLnx = "/upgrade_lnx/delete",
  GetUpgradeLnxById = "/upgrade_lnx",
  GetUpgradeLnxList = "/upgrade_lnx/list",
  UpdateUpgradeLnx = "/upgrade_lnx/update",
}

/**
 * @description: Get upgrade lnx list
 */

export const getUpgradeLnxList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeLnxListResp>>(
    Api.GetUpgradeLnxList,
    params,
  );
};

/**
 *  @description: Create a new upgrade lnx
 */
export const createUpgradeLnx = (params: UpgradeLnxInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeLnx, params);
};

/**
 *  @description: Update the upgrade lnx
 */
export const updateUpgradeLnx = (params: UpgradeLnxInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeLnx, params);
};

/**
 *  @description: Delete upgrade lnxs
 */
export const deleteUpgradeLnx = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeLnx, params);
};

/**
 *  @description: Get upgrade lnx By ID
 */
export const getUpgradeLnxById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeLnxInfo>>(
    Api.GetUpgradeLnxById,
    params,
  );
};
