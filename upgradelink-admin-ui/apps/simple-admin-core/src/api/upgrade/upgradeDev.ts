import type {
  UpgradeDevInfo,
  UpgradeDevListResp,
} from "./model/upgradeDevModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeDev = "/upgrade/upgrade_dev/create",
  DeleteUpgradeDev = "/upgrade/upgrade_dev/delete",
  GetUpgradeDevById = "/upgrade/upgrade_dev",
  GetUpgradeDevList = "/upgrade/upgrade_dev/list",
  UpdateUpgradeDev = "/upgrade/upgrade_dev/update",
}

/**
 * @description: Get upgrade dev list
 */

export const getUpgradeDevList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeDevListResp>>(
    Api.GetUpgradeDevList,
    params,
  );
};

/**
 *  @description: Create a new upgrade dev
 */
export const createUpgradeDev = (params: UpgradeDevInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUpgradeDev, params);
};

/**
 *  @description: Update the upgrade dev
 */
export const updateUpgradeDev = (params: UpgradeDevInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUpgradeDev, params);
};

/**
 *  @description: Delete upgrade devs
 */
export const deleteUpgradeDev = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUpgradeDev, params);
};

/**
 *  @description: Get upgrade dev By ID
 */
export const getUpgradeDevById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeDevInfo>>(
    Api.GetUpgradeDevById,
    params,
  );
};
