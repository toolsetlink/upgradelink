import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: UpgradeCompanyTrafficPacket info response
 */
export interface UpgradeCompanyTrafficPacketInfo {
  id?: number;
  companyId?: number;
  packetId?: number;
  startTime?: number;
  endTime?: number;
  initialSize?: number;
  remainingSize?: number;
  status?: number;
  exchangeTime?: number;
  createAt?: number;
  updateAt?: number;
}

/**
 *  @description: UpgradeCompanyTrafficPacket list response
 */

export type UpgradeCompanyTrafficPacketListResp =
  BaseListResp<UpgradeCompanyTrafficPacketInfo>;
