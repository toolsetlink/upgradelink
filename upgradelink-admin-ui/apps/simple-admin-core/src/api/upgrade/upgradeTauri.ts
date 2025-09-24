import type {
  UpgradeTauriInfo,
  UpgradeTauriListResp,
} from "./model/upgradeTauriModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeTauri = "/upgrade/upgrade_tauri/create",
  DeleteUpgradeTauri = "/upgrade/upgrade_tauri/delete",
  GetUpgradeTauriById = "/upgrade/upgrade_tauri",
  GetUpgradeTauriList = "/upgrade/upgrade_tauri/list",
  UpdateUpgradeTauri = "/upgrade/upgrade_tauri/update",
}

/**
 * @description: Get upgrade tauri list
 */

export const getUpgradeTauriList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeTauriListResp>>(
    Api.GetUpgradeTauriList,
    params,
  );
};

/**
 *  @description: Create a new upgrade tauri
 */
export const createUpgradeTauri = (params: UpgradeTauriInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeTauri, params);
};

/**
 *  @description: Update the upgrade tauri
 */
export const updateUpgradeTauri = (params: UpgradeTauriInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeTauri, params);
};

/**
 *  @description: Delete upgrade tauris
 */
export const deleteUpgradeTauri = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeTauri, params);
};

/**
 *  @description: Get upgrade tauri By ID
 */
export const getUpgradeTauriById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeTauriInfo>>(
    Api.GetUpgradeTauriById,
    params,
  );
};
