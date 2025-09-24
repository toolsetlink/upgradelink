import type {
  UpgradeElectronInfo,
  UpgradeElectronListResp,
} from "./model/upgradeElectronModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeElectron = "/upgrade/upgrade_electron/create",
  DeleteUpgradeElectron = "/upgrade/upgrade_electron/delete",
  GetUpgradeElectronById = "/upgrade/upgrade_electron",
  GetUpgradeElectronList = "/upgrade/upgrade_electron/list",
  UpdateUpgradeElectron = "/upgrade/upgrade_electron/update",
}

/**
 * @description: Get upgrade electron list
 */

export const getUpgradeElectronList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeElectronListResp>>(
    Api.GetUpgradeElectronList,
    params,
  );
};

/**
 *  @description: Create a new upgrade electron
 */
export const createUpgradeElectron = (params: UpgradeElectronInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeElectron, params);
};

/**
 *  @description: Update the upgrade electron
 */
export const updateUpgradeElectron = (params: UpgradeElectronInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeElectron, params);
};

/**
 *  @description: Delete upgrade electrons
 */
export const deleteUpgradeElectron = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeElectron, params);
};

/**
 *  @description: Get upgrade electron By ID
 */
export const getUpgradeElectronById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeElectronInfo>>(
    Api.GetUpgradeElectronById,
    params,
  );
};
