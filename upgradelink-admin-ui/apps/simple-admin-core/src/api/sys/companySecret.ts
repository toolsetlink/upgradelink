import type {
  CompanySecretInfo,
  CompanySecretListResp,
} from "./model/companySecretModel";

import type {
  BaseDataResp,
  BaseIDReq,
  BaseIDsReq,
  BaseListReq,
  BaseResp,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateCompanySecret = "/sys-api/company_secret/create",
  DeleteCompanySecret = "/sys-api/company_secret/delete",
  GetCompanySecretById = "/sys-api/company_secret",
  GetCompanySecretList = "/sys-api/company_secret/list",
  UpdateCompanySecret = "/sys-api/company_secret/update",
}

/**
 * @description: Get company secret list
 */

export const getCompanySecretList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<CompanySecretListResp>>(
    Api.GetCompanySecretList,
    params,
  );
};

/**
 *  @description: Create a new company secret
 */
export const createCompanySecret = (params: CompanySecretInfo) => {
  return requestClient.post<BaseResp>(Api.CreateCompanySecret, params);
};

/**
 *  @description: Update the company secret
 */
export const updateCompanySecret = (params: CompanySecretInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateCompanySecret, params);
};

/**
 *  @description: Delete company secrets
 */
export const deleteCompanySecret = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteCompanySecret, params);
};

/**
 *  @description: Get company secret By ID
 */
export const getCompanySecretById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<CompanySecretInfo>>(
    Api.GetCompanySecretById,
    params,
  );
};
