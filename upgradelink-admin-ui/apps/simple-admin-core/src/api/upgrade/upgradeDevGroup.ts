import type {
  UpgradeDevGroupInfo,
  UpgradeDevGroupListResp,
} from "./model/upgradeDevGroupModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeDevGroup = "/upgrade/upgrade_dev_group/create",
  DeleteUpgradeDevGroup = "/upgrade/upgrade_dev_group/delete",
  GetUpgradeDevGroupById = "/upgrade/upgrade_dev_group",
  GetUpgradeDevGroupList = "/upgrade/upgrade_dev_group/list",
  UpdateUpgradeDevGroup = "/upgrade/upgrade_dev_group/update",
}

/**
 * @description: Get upgrade dev group list
 */

export const getUpgradeDevGroupList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeDevGroupListResp>>(
    Api.GetUpgradeDevGroupList,
    params,
  );
};

/**
 *  @description: Create a new upgrade dev group
 */
export const createUpgradeDevGroup = (params: UpgradeDevGroupInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeDevGroup, params);
};

/**
 *  @description: Update the upgrade dev group
 */
export const updateUpgradeDevGroup = (params: UpgradeDevGroupInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeDevGroup, params);
};

/**
 *  @description: Delete upgrade dev groups
 */
export const deleteUpgradeDevGroup = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeDevGroup, params);
};

/**
 *  @description: Get upgrade dev group By ID
 */
export const getUpgradeDevGroupById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeDevGroupInfo>>(
    Api.GetUpgradeDevGroupById,
    params,
  );
};
