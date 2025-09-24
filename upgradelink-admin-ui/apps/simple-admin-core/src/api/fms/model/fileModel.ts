import type { BaseListResp } from "../../model/baseModel";

/**
 *  author: toolsetlink
 *  @description: file info response
 */
export interface fileInfo {
  id: string;
  createdAt?: number;
  name: string;
  fileType: string;
  size: number;
  path: string;
  publicPath: string;
  tagIds: number[];
}

/**
 *  author: toolsetlink
 *  @description: file list response
 */

export type FileListResp = BaseListResp<fileInfo>;

/**
 *  author: toolsetlink
 *  @description: change status request
 */
export interface changeStatusReq {
  id: string;
  status: boolean;
}

/**
 *  author: toolsetlink
 *  @description: update file info request
 */
export interface updateFileInfoReq {
  id: string;
  name: string;
  tagIds: number[];
}

/**
 *  @description: file deletion request
 */

export interface FileDeleteReq {
  url: string;
}
