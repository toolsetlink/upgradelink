import type {
  UpgradeTauriVersionInfo,
  UpgradeTauriVersionListResp,
} from "./model/upgradeTauriVersionModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeTauriVersion = "/upgrade/upgrade_tauri_version/create",
  DeleteUpgradeTauriVersion = "/upgrade/upgrade_tauri_version/delete",
  GetUpgradeTauriVersionById = "/upgrade/upgrade_tauri_version",
  GetUpgradeTauriVersionList = "/upgrade/upgrade_tauri_version/list",
  UpdateUpgradeTauriVersion = "/upgrade/upgrade_tauri_version/update",
}

/**
 * @description: Get upgrade tauri version list
 */

export const getUpgradeTauriVersionList = (params: {
  page: number;
  pageSize: number;
  tauriId: any | undefined;
}) => {
  return requestClient.post<BaseDataResp<UpgradeTauriVersionListResp>>(
    Api.GetUpgradeTauriVersionList,
    params,
  );
};

/**
 *  @description: Create a new upgrade tauri version
 */
export const createUpgradeTauriVersion = (params: UpgradeTauriVersionInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeTauriVersion, params);
};

/**
 *  @description: Update the upgrade tauri version
 */
export const updateUpgradeTauriVersion = (params: UpgradeTauriVersionInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeTauriVersion, params);
};

/**
 *  @description: Delete upgrade tauri versions
 */
export const deleteUpgradeTauriVersion = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeTauriVersion, params);
};

/**
 *  @description: Get upgrade tauri version By ID
 */
export const getUpgradeTauriVersionById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeTauriVersionInfo>>(
    Api.GetUpgradeTauriVersionById,
    params,
  );
};
