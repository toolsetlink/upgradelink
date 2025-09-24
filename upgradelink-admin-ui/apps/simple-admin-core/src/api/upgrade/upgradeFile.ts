import type {
  UpgradeFileInfo,
  UpgradeFileListResp,
} from "./model/upgradeFileModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeFile = "/upgrade/upgrade_file/create",
  DeleteUpgradeFile = "/upgrade/upgrade_file/delete",
  GetUpgradeFileById = "/upgrade/upgrade_file",
  GetUpgradeFileList = "/upgrade/upgrade_file/list",
  UpdateUpgradeFile = "/upgrade/upgrade_file/update",
}

/**
 * @description: Get upgrade file list
 */

export const getUpgradeFileList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeFileListResp>>(
    Api.GetUpgradeFileList,
    params,
  );
};

/**
 *  @description: Create a new upgrade file
 */
export const createUpgradeFile = (params: UpgradeFileInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeFile, params);
};

/**
 *  @description: Update the upgrade file
 */
export const updateUpgradeFile = (params: UpgradeFileInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeFile, params);
};

/**
 *  @description: Delete upgrade files
 */
export const deleteUpgradeFile = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeFile, params);
};

/**
 *  @description: Get upgrade file By ID
 */
export const getUpgradeFileById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeFileInfo>>(
    Api.GetUpgradeFileById,
    params,
  );
};
