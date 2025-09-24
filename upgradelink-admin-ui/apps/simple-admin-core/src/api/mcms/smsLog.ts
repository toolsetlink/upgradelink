import type { SmsLogInfo, SmsLogListResp } from "./model/smsLogModel";

import type {
  BaseDataResp,
  BaseListReq,
  BaseResp,
  BaseUUIDReq,
  BaseUUIDsReq,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateSmsLog = "/sys-api/sms_log/create",
  DeleteSmsLog = "/sys-api/sms_log/delete",
  GetSmsLogById = "/sys-api/sms_log",
  GetSmsLogList = "/sys-api/sms_log/list",
  UpdateSmsLog = "/sys-api/sms_log/update",
}

/**
 * @description: Get sms log list
 */

export const getSmsLogList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<SmsLogListResp>>(
    Api.GetSmsLogList,
    params,
  );
};

/**
 *  @description: Create a new sms log
 */
export const createSmsLog = (params: SmsLogInfo) => {
  return requestClient.post<BaseResp>(Api.CreateSmsLog, params);
};

/**
 *  @description: Update the sms log
 */
export const updateSmsLog = (params: SmsLogInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateSmsLog, params);
};

/**
 *  @description: Delete sms logs
 */
export const deleteSmsLog = (params: BaseUUIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteSmsLog, params);
};

/**
 *  @description: Get sms log By ID
 */
export const getSmsLogById = (params: BaseUUIDReq) => {
  return requestClient.post<BaseDataResp<SmsLogInfo>>(
    Api.GetSmsLogById,
    params,
  );
};
