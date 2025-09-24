import type { RoleInfo, RoleListResp } from "./model/roleModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateRole = "/sys-api/role/create",
  DeleteRole = "/sys-api/role/delete",
  GetRoleById = "/sys-api/role",
  GetRoleList = "/sys-api/role/list",
  UpdateRole = "/sys-api/role/update",
}

/**
 * @description: Get role list
 */

export const getRoleList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<RoleListResp>>(
    Api.GetRoleList,
    params,
  );
};

/**
 *  @description: Create a new role
 */
export const createRole = (params: RoleInfo) => {
  return requestClient.post<BaseResp>(Api.CreateRole, params);
};

/**
 *  @description: Update the role
 */
export const updateRole = (params: RoleInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateRole, params);
};

/**
 *  @description: Delete roles
 */
export const deleteRole = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteRole, params);
};

/**
 *  @description: Get role By ID
 */
export const getRoleById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<RoleInfo>>(Api.GetRoleById, params);
};
