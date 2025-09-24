import type { UpgradeDashboardInfo } from "./model/upgradeDashboardModel";

import type { BaseDataResp } from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  GetUpgradeDashboard = "/upgrade/upgrade_dashboard",
}

/**
 *  @description: Get upgrade dashboard By ID
 */
export const getUpgradeDashboard = () => {
  return requestClient.post<BaseDataResp<UpgradeDashboardInfo>>(
    Api.GetUpgradeDashboard,
  );
};
