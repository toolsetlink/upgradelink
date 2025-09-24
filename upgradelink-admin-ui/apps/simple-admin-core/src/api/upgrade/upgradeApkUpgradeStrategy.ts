import type {
  UpgradeApkUpgradeStrategyInfo,
  UpgradeApkUpgradeStrategyListResp,
} from "./model/upgradeApkUpgradeStrategyModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeApkUpgradeStrategy = "/upgrade/upgrade_apk_upgrade_strategy/create",
  DeleteUpgradeApkUpgradeStrategy = "/upgrade/upgrade_apk_upgrade_strategy/delete",
  GetUpgradeApkUpgradeStrategyById = "/upgrade/upgrade_apk_upgrade_strategy",
  GetUpgradeApkUpgradeStrategyList = "/upgrade/upgrade_apk_upgrade_strategy/list",
  UpdateUpgradeApkUpgradeStrategy = "/upgrade/upgrade_apk_upgrade_strategy/update",
}

/**
 * @description: Get upgrade apk upgrade strategy list
 */

export const getUpgradeApkUpgradeStrategyList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeApkUpgradeStrategyListResp>>(
    Api.GetUpgradeApkUpgradeStrategyList,
    params,
  );
};

/**
 *  @description: Create a new upgrade apk upgrade strategy
 */
export const createUpgradeApkUpgradeStrategy = (
  params: UpgradeApkUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeApkUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Update the upgrade apk upgrade strategy
 */
export const updateUpgradeApkUpgradeStrategy = (
  params: UpgradeApkUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeApkUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Delete upgrade apk upgrade strategys
 */
export const deleteUpgradeApkUpgradeStrategy = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeApkUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Get upgrade apk upgrade strategy By ID
 */
export const getUpgradeApkUpgradeStrategyById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeApkUpgradeStrategyInfo>>(
    Api.GetUpgradeApkUpgradeStrategyById,
    params,
  );
};
