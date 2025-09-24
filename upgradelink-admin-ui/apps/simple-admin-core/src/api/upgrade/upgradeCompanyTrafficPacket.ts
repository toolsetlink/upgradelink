import type {
  UpgradeCompanyTrafficPacketInfo,
  UpgradeCompanyTrafficPacketListResp,
} from "./model/upgradeCompanyTrafficPacketModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateUpgradeCompanyTrafficPacket = "/upgrade/upgrade_company_traffic_packet/create",
  DeleteUpgradeCompanyTrafficPacket = "/upgrade/upgrade_company_traffic_packet/delete",
  GetUpgradeCompanyTrafficPacketById = "/upgrade/upgrade_company_traffic_packet",
  GetUpgradeCompanyTrafficPacketList = "/upgrade/upgrade_company_traffic_packet/list",
  UpdateUpgradeCompanyTrafficPacket = "/upgrade/upgrade_company_traffic_packet/update",
}

/**
 * @description: Get upgrade company traffic packet list
 */

export const getUpgradeCompanyTrafficPacketList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UpgradeCompanyTrafficPacketListResp>>(
    Api.GetUpgradeCompanyTrafficPacketList,
    params,
  );
};

/**
 *  @description: Create a new upgrade company traffic packet
 */
export const createUpgradeCompanyTrafficPacket = (
  params: UpgradeCompanyTrafficPacketInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.CreateUpgradeCompanyTrafficPacket,
    params,
  );
};

/**
 *  @description: Update the upgrade company traffic packet
 */
export const updateUpgradeCompanyTrafficPacket = (
  params: UpgradeCompanyTrafficPacketInfo,
) => {
  return requestClient.post<BaseResp>(
    Api.UpdateUpgradeCompanyTrafficPacket,
    params,
  );
};

/**
 *  @description: Delete upgrade company traffic packets
 */
export const deleteUpgradeCompanyTrafficPacket = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(
    Api.DeleteUpgradeCompanyTrafficPacket,
    params,
  );
};

/**
 *  @description: Get upgrade company traffic packet By ID
 */
export const getUpgradeCompanyTrafficPacketById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<UpgradeCompanyTrafficPacketInfo>>(
    Api.GetUpgradeCompanyTrafficPacketById,
    params,
  );
};
