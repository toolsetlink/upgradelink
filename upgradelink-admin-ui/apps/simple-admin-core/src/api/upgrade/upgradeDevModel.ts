import type {
  UpgradeDevModelInfo,
  UpgradeDevModelListResp,
} from "./model/upgradeDevModelModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeDevModel = "/upgrade/upgrade_dev_model/create",
  DeleteUpgradeDevModel = "/upgrade/upgrade_dev_model/delete",
  GetUpgradeDevModelById = "/upgrade/upgrade_dev_model",
  GetUpgradeDevModelList = "/upgrade/upgrade_dev_model/list",
  UpdateUpgradeDevModel = "/upgrade/upgrade_dev_model/update",
}

/**
 * @description: Get upgrade dev model list
 */

export const getUpgradeDevModelList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeDevModelListResp>>(
    Api.GetUpgradeDevModelList,
    params,
  );
};

/**
 *  @description: Create a new upgrade dev model
 */
export const createUpgradeDevModel = (params: UpgradeDevModelInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeDevModel, params);
};

/**
 *  @description: Update the upgrade dev model
 */
export const updateUpgradeDevModel = (params: UpgradeDevModelInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeDevModel, params);
};

/**
 *  @description: Delete upgrade dev models
 */
export const deleteUpgradeDevModel = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeDevModel, params);
};

/**
 *  @description: Get upgrade dev model By ID
 */
export const getUpgradeDevModelById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeDevModelInfo>>(
    Api.GetUpgradeDevModelById,
    params,
  );
};
