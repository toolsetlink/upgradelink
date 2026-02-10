import type {
  CloudFileDeleteReq,
  CloudFileInfo,
  CloudFileListResp,
} from "./model/cloudFileModel";

import type {
  BaseDataResp,
  BaseListReq,
  BaseResp,
  BaseUUIDReq,
  BaseUUIDsReq,
} from "#/api/model/baseModel";

import { requestClient } from "#/api/request";

enum Api {
  CreateCloudFile = "/cloud_file/create",
  DeleteCloudFile = "/cloud_file/delete",
  DeleteCloudFileByUrl = "/cloud_file/delete_by_url",
  GetCloudFileById = "/cloud_file",
  GetCloudFileList = "/cloud_file/list",
  UpdateCloudFile = "/cloud_file/update",
  uploadFile = "/cloud_file/upload",
  GetPresignedUrl = "/cloud_file/presigned_url",
}

/**
 * @description: Get cloud file list
 */

export const getCloudFileList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<CloudFileListResp>>(
      Api.GetCloudFileList,
      params,
  );
};

/**
 *  @description: Create a new cloud file
 */
export const createCloudFile = (params: CloudFileInfo) => {
  return requestClient.post<BaseDataResp<CloudFileInfo>>(Api.CreateCloudFile, params);
};

/**
 *  @description: Update the cloud file
 */
export const updateCloudFile = (params: CloudFileInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateCloudFile, params);
};

/**
 *  @description: Delete cloud files
 */
export const deleteCloudFile = (params: BaseUUIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteCloudFile, params);
};

/**
 *  @description: Get cloud file By ID
 */
export const getCloudFileById = (params: BaseUUIDReq) => {
  return requestClient.post<BaseDataResp<CloudFileInfo>>(
      Api.GetCloudFileById,
      params,
  );
};

/**
 * @description: Get presigned URL for direct upload
 */
export function getPresignedUrl(filename: string, contentType: string) {
  return requestClient.post<BaseDataResp<{
    presignedUrl: string;
    fileUrl: string;
    fileName: string;
    fileSuffix: string;
    fileType: string;
  }>>(Api.GetPresignedUrl, { filename, contentType });
}

/**
 * @description: Upload file directly to S3 using presigned URL
 */
export async function uploadCloudFile(file: File, provider: string = "") {
  // First get presigned URL
  const presignedUrlResponse = await getPresignedUrl(file.name, file.type);
  const presignedUrlData = presignedUrlResponse.data;

  // Upload file directly to S3 using presigned URL
  const uploadResponse = await fetch(presignedUrlData.presignedUrl, {
    method: 'PUT',
    body: file,
    headers: {
      'Content-Type': file.type
    }
  });

  if (!uploadResponse.ok) {
    throw new Error('Failed to upload file to S3');
  }

  // Return the same structure as before for backward compatibility
  return {
    code: 0,
    msg: 'success',
    data: {
      name: presignedUrlData.fileName,
      url: presignedUrlData.fileUrl,
      size: file.size,
      fileType: presignedUrlData.fileType,
    }
  };
}

/**
 *  @description: Delete cloud file by url
 */
export const deleteCloudFileByUrl = (params: CloudFileDeleteReq) => {
  return requestClient.post<BaseResp>(Api.DeleteCloudFileByUrl, params);
};
