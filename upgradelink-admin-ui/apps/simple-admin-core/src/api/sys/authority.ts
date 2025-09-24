import type {
  ApiAuthorityReq,
  ApiAuthorityResp,
  ApiListReq,
  MenuAuthorityInfo,
} from "./model/authorityModel";

import type { BaseDataResp, BaseIDReq, BaseResp } from "#/api/model/baseModel";
import type { ApiListResp } from "#/api/sys/model/apiModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateOrUpdateApiAuthority = "/sys-api/authority/api/create_or_update",
  CreateOrUpdateMenuAuthority = "/sys-api/authority/menu/create_or_update",
  GetApiList = "/sys-api/api/list",
  GetRoleApiList = "/sys-api/authority/api/role",
  GetRoleMenuList = "/sys-api/authority/menu/role",
}

/**
 *  author: toolsetlink
 *  @description: this function is used to get api list for authorization
 */

export const getApiList = (params: ApiListReq) => {
  return requestClient.post<BaseDataResp<ApiListResp>>(Api.GetApiList, params);
};

/**
 * @description: Get api authorization list
 */

export const getApiAuthority = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<ApiAuthorityResp>>(
    Api.GetRoleApiList,
    params,
  );
};

/**
 *  author: ryan
 *  @description: create or update api authorization
 */
export const createOrUpdateApiAuthority = (params: ApiAuthorityReq) => {
  return requestClient.post<BaseResp>(Api.CreateOrUpdateApiAuthority, params);
};

/**
 *  author: toolsetlink
 *  @description:
 */

export const createOrUpdateMenuAuthority = (params: MenuAuthorityInfo) => {
  return requestClient.post<BaseResp>(Api.CreateOrUpdateMenuAuthority, params);
};

/**
 *  author: toolsetlink
 *  @description: get role's menu authorization ids
 */

export const getMenuAuthority = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<MenuAuthorityInfo>>(
    Api.GetRoleMenuList,
    params,
  );
};
