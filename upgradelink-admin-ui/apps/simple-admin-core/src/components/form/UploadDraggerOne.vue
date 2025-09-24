<script lang="ts" setup>
import type { UploadProps } from "ant-design-vue";

import { ref, watch } from "vue";

import { $t } from "@vben/locales";

import { InboxOutlined, LoadingOutlined } from "@ant-design/icons-vue";
import { useClipboard, useVModel } from "@vueuse/core";
import { message, UploadDragger } from "ant-design-vue";

import { uploadCloudFile } from "#/api/fms/cloudFile";
import { uploadFile } from "#/api/fms/file";

defineOptions({
  name: "UploadDraggerOne",
  inheritAttrs: false, // 禁止自动继承属性
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
});

const emits = defineEmits(["update:value"]);

const { copy } = useClipboard();

const fileUrlDict: { [key: string]: string } = {};

const fileId = ref<string>();
const fileList = ref<UploadProps["fileList"]>();
const isUploading = ref<boolean>(false); // 新增：上传状态标识

const state = useVModel(props, "value", emits, {
  defaultValue: props.value,
  passive: true,
});

async function handleUpload(file: any) {
  isUploading.value = true; // 开始上传，设置状态为 true

  try {
    if (props.provider === "local") {
      const result = (await uploadFile(file.file)) as any;
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
    } else {
      const result = (await uploadCloudFile(
        file.file,
        props.provider === "cloud-default" ? "" : props.provider,
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
    }
  } catch (error) {
    // 处理错误
    console.error("上传失败:", error);
    message.error($t("component.upload.uploadFailed"));
  } finally {
    isUploading.value = false; // 上传完成（无论成功或失败），设置状态为 false
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
  <!-- 上传区域 -->
  <div>
    <UploadDragger
      v-if="!fileList?.length"
      v-model:file-list="fileList"
      :custom-request="handleUpload"
      :multiple="props.multiple"
      :max-count="1"
      :show-upload-list="false"
      v-bind="$attrs"
    >
      <p class="ant-upload-drag-icon">
        <InboxOutlined />
      </p>
      <p class="ant-upload-text">
        {{ $t("component.upload.uploadHelpMessage") }}
      </p>
      <p class="ant-upload-hint">
        {{ $t("component.upload.supportAnyFormatOne") }}
      </p>
    </UploadDragger>

    <!-- 文件列表展示 -->
    <div v-if="fileList?.length" class="uploaded-file">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <!-- 显示加载状态图标 -->
          <span v-if="isUploading" class="mr-2 text-primary">
            <LoadingOutlined spin />
          </span>
          {{ fileList[0]?.name || $t("common.unnamedFile") }}
        </div>
        <!-- 显示上传中文字 -->
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
