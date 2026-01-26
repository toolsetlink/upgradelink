<script lang="ts" setup>
import type { UploadProps } from "ant-design-vue";

import { ref, watch } from "vue";

import { $t } from "@vben/locales";

import { InboxOutlined, LoadingOutlined } from "@ant-design/icons-vue";
import { useVModel } from "@vueuse/core";
import { message, UploadDragger } from "ant-design-vue";

import { uploadCloudFile } from "#/api/sys/cloudFile";

defineOptions({
  name: "UploadDraggerOne",
  inheritAttrs: false,
});

const props = defineProps({
  multiple: {
    type: Boolean,
    default: false,
  },
  provider: {
    type: String,
    default: "local",
  },
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
    const result = (await uploadCloudFile(
      file.file,
      "",
    )) as any;
    if (fileList.value !== undefined) {
      fileList.value.forEach((item) => {
        if (item.uid === file.file.uid) {
          const resultStatus = result.code !== undefined && result.code === 0;
          item.status = resultStatus ? "done" : "error";
          if (resultStatus) {
            fileId.value = result.data.id;
            message.success($t("component.upload.uploadSuccess"));
          }
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
      :multiple="props.multiple"
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
.uploaded-file {
  padding: 8px;
  color: rgba(0, 0, 0, 0.88);
  font-size: 14px;
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
}
</style>
