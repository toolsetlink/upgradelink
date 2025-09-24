import type { SendEmailReq, SendSmsReq } from "./model/messageModel";

import type { BaseResp } from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  SendEmail = "/sys-api/email/send",
  SendSms = "/sys-api/sms/send",
}

/**
 * @description: Send Email
 */

export const sendEmail = (params: SendEmailReq) => {
  return requestClient.post<BaseResp>(Api.SendEmail, params);
};

/**
 * @description: Send Sms
 */

export const sendSms = (params: SendSmsReq) => {
  return requestClient.post<BaseResp>(Api.SendSms, params);
};
