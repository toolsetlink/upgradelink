import type { BaseListResp } from "../../model/baseModel";

/**
 *  @description: CompanySecret info response
 */
export interface CompanySecretInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  access_key?: string;
  secret_key?: string;
}

/**
 *  @description: CompanySecret list response
 */

export type CompanySecretListResp = BaseListResp<CompanySecretInfo>;
