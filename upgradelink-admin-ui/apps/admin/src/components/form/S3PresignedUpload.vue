<script lang="ts" setup>
import type { UploadProps } from "ant-design-vue";

import { ref, watch } from "vue";

import { $t } from "@vben/locales";

import { InboxOutlined, LoadingOutlined } from "@ant-design/icons-vue";
import { useVModel } from "@vueuse/core";
import { message, UploadDragger } from "ant-design-vue";

import CryptoJS from "crypto-js";

import { getPresignedUrl, createCloudFile } from "#/api/sys/cloudFile";

defineOptions({
  name: "S3PresignedUpload",
  inheritAttrs: false,
});

const props = defineProps({
  value: {
    type: [Array, Object, String, Number],
    default: undefined,
  },
  accept: {
    type: [String, Array],
    default: null,
    description: "接受的文件类型，例如 '.jpg,.png' 或 ['.jpg', '.png']",
  },
});

const emits = defineEmits(["update:value"]);

const fileId = ref<string>();
const fileList = ref<UploadProps["fileList"]>();
const isUploading = ref<boolean>(false);

const state = useVModel(props, "value", emits, {
  defaultValue: props.value,
  passive: true,
});

function checkFileType(file: File): boolean {
  if (!props.accept) return true;

  const fileName = file.name.toLowerCase();
  let allowedTypes: string[] = [];

  if (typeof props.accept === "string") {
    allowedTypes = props.accept
      .split(",")
      .map((type) => type.trim().toLowerCase());
  } else if (Array.isArray(props.accept)) {
    allowedTypes = props.accept.map((type) => type.trim().toLowerCase());
  }

  return allowedTypes.some((type) => {
    if (type.startsWith(".")) {
      return fileName.endsWith(type);
    }
    const fileExt = fileName.slice(fileName.lastIndexOf("."));
    return (file as any).type?.includes(type) || allowedTypes.includes(fileExt);
  });
}

function getAllowedTypesText(): string {
  if (!props.accept) return $t("component.upload.supportAnyFormatOne");

  let types: string[] = [];
  if (typeof props.accept === "string") {
    types = props.accept.split(",").map((type) => type.trim());
  } else if (Array.isArray(props.accept)) {
    types = props.accept.map((type) => type.trim());
  }

  return types.join(", ");
}

// 计算文件的 MD5 值
async function calculateFileMD5(file: File): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    
    reader.onload = (event) => {
      try {
        const arrayBuffer = event.target?.result as ArrayBuffer;
        const uint8Array = new Uint8Array(arrayBuffer);
        // 将 Uint8Array 转换为 WordArray，然后计算 MD5
        const wordArray = CryptoJS.lib.WordArray.create(uint8Array);
        const md5Hash = CryptoJS.MD5(wordArray).toString();
        resolve(md5Hash);
      } catch (error) {
        reject(error);
      }
    };
    
    reader.onerror = (error) => {
      reject(error);
    };
    
    reader.readAsArrayBuffer(file);
  });
}

async function handleUpload(file: any) {
  if (!checkFileType(file.file)) {
    message.error(
      `${$t("component.upload.fileTypeNotAllowed")}: ${getAllowedTypesText()}`,
    );

    if (fileList.value !== undefined) {
      fileList.value.forEach((item) => {
        if (item.uid === file.file.uid) {
          item.status = "error";
        }
      });
    }
    return;
  }

  isUploading.value = true;

  try {
      // 1. 计算文件 MD5 值
      const md5Hash = await calculateFileMD5(file.file);
      
      // 2. 获取预签名URL
      const presignedUrlResponse = await getPresignedUrl(file.file.name, file.file.type);
      const presignedUrlData = presignedUrlResponse.data;
      
      // 3. 直接上传到S3
      const uploadResponse = await fetch(presignedUrlData.presignedUrl, {
        method: 'PUT',
        body: file.file,
        headers: {
          'Content-Type': file.file.type
        }
      });
      
      if (!uploadResponse.ok) {
        throw new Error('Failed to upload file to S3');
      }
      
      // 4. 调用 createCloudFile 接口落库
      const createResponse = await createCloudFile({
        name: presignedUrlData.fileName,
        url: presignedUrlData.fileUrl,
        size: file.file.size,
        md5: md5Hash,
        fileType: presignedUrlData.fileType
      });
      
      // 5. 获取创建成功后返回的文件 ID
      const createdFileId = createResponse.data.id;
    
    if (fileList.value !== undefined) {
      fileList.value.forEach((item) => {
        if (item.uid === file.file.uid) {
          item.status = "done";
          fileId.value = createdFileId;
          message.success($t("component.upload.uploadSuccess"));
        }
      });
    }
  } catch (error) {
    console.error("上传失败:", error);
    message.error($t("component.upload.uploadFailed"));
  } finally {
    isUploading.value = false;
  }
}

watch(
  () => fileId.value,
  (newId) => {
    state.value = newId;
  },
);
</script>

<template>
  <div>
    <UploadDragger
      v-if="!fileList?.length"
      v-model:file-list="fileList"
      :custom-request="handleUpload"
      :max-count="1"
      :show-upload-list="false"
      :accept="
        props.accept && Array.isArray(props.accept)
          ? props.accept.join(',')
          : props.accept
      "
      v-bind="$attrs"
    >
      <p class="ant-upload-drag-icon">
        <InboxOutlined />
      </p>
      <p class="ant-upload-text">
        {{ $t("component.upload.uploadHelpMessage") }}
      </p>
      <p class="ant-upload-hint">
        {{ $t("component.upload.supportFormat") }}{{ getAllowedTypesText() }}
      </p>
    </UploadDragger>

    <div v-if="fileList?.length" class="uploaded-file">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <span v-if="isUploading" class="mr-2 text-primary">
            <LoadingOutlined spin />
          </span>
          {{ fileList[0]?.name || $t("common.unnamedFile") }}
        </div>
        <span v-if="isUploading" class="text-sm text-primary">
          {{ $t("component.upload.uploading") }}
        </span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.ant-upload-drag-icon {
  font-size: 48px;
  color: #1677ff;
}

.ant-upload-text {
  margin-top: 16px;
  font-size: 16px;
  font-weight: 500;
}

.ant-upload-hint {
  margin-top: 8px;
  color: #6b7280;
  font-size: 14px;
}

.uploaded-file {
  padding: 8px;
  color: rgba(0, 0, 0, 0.88);
  font-size: 14px;
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
}
</style>
