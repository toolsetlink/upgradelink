import type {
  UpgradeTauriUpgradeStrategyInfo,
  UpgradeTauriUpgradeStrategyListResp,
} from "./model/upgradeTauriUpgradeStrategyModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeTauriUpgradeStrategy = "/upgrade/upgrade_tauri_upgrade_strategy/create",
  DeleteUpgradeTauriUpgradeStrategy = "/upgrade/upgrade_tauri_upgrade_strategy/delete",
  GetUpgradeTauriUpgradeStrategyById = "/upgrade/upgrade_tauri_upgrade_strategy",
  GetUpgradeTauriUpgradeStrategyList = "/upgrade/upgrade_tauri_upgrade_strategy/list",
  UpdateUpgradeTauriUpgradeStrategy = "/upgrade/upgrade_tauri_upgrade_strategy/update",
}

/**
 * @description: Get upgrade tauri upgrade strategy list
 */

export const getUpgradeTauriUpgradeStrategyList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeTauriUpgradeStrategyListResp>>(
    Api.GetUpgradeTauriUpgradeStrategyList,
    params,
  );
};

/**
 *  @description: Create a new upgrade tauri upgrade strategy
 */
export const createUpgradeTauriUpgradeStrategy = (
  params: UpgradeTauriUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeTauriUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Update the upgrade tauri upgrade strategy
 */
export const updateUpgradeTauriUpgradeStrategy = (
  params: UpgradeTauriUpgradeStrategyInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeTauriUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Delete upgrade tauri upgrade strategys
 */
export const deleteUpgradeTauriUpgradeStrategy = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeTauriUpgradeStrategy,
    params,
  );
};

/**
 *  @description: Get upgrade tauri upgrade strategy By ID
 */
export const getUpgradeTauriUpgradeStrategyById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeTauriUpgradeStrategyInfo>>(
    Api.GetUpgradeTauriUpgradeStrategyById,
    params,
  );
};
