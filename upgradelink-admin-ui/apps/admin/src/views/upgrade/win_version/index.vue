<script lang="ts" setup>
import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridListeners, VxeGridProps } from "#/adapter/vxe-table";
import type { UpgradeWinVersionInfo } from "#/api/upgrade/model/upgradeWinVersionModel";
import type { ActionItem } from "#/components/table/table-action";

import { h, ref } from "vue";

import { Page, useVbenModal } from "@vben/common-ui";
import { $t } from "@vben/locales";

import { Button, Modal, QRCode } from "ant-design-vue";
import { isPlainObject } from "remeda";

import { useVbenVxeGrid } from "#/adapter/vxe-table";
import {
  deleteUpgradeWinVersion,
  getUpgradeWinVersionList,
} from "#/api/upgrade/upgradeWinVersion";
import { TableAction } from "#/components/table/table-action";

import UpgradeWinVersionForm from "./form.vue";
import { searchFormSchemas, tableColumns } from "./schemas";

defineOptions({
  name: "UpgradeWinVersionManagement",
});

// ---------------- form -----------------

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: UpgradeWinVersionForm,
});

const showDeleteButton = ref<boolean>(false);
const qrCodeVisible = ref<boolean>(false);
const currentQrCodeUrl = ref<string>("");
const currentVersionName = ref<string>("");

const gridEvents: VxeGridListeners<any> = {
  checkboxChange(e) {
    showDeleteButton.value = e.$table.getCheckboxRecords().length > 0;
  },
  checkboxAll(e) {
    showDeleteButton.value = e.$table.getCheckboxRecords().length > 0;
  },
};

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  schema: [...(searchFormSchemas.schema as any)],
  // 控制表单是否显示折叠按钮
  showCollapseButton: true,
  // 按下回车时是否提交表单
  submitOnEnter: false,
};

// ------------- table --------------------

const gridOptions: VxeGridProps<UpgradeWinVersionInfo> = {
  checkboxConfig: {
    highlight: true,
  },
  toolbarConfig: {
    slots: {
      buttons: "toolbar-buttons",
    },
  },
  columns: [
    ...(tableColumns.columns as any),
    {
      title: $t("common.action"),
      fixed: "right",
      field: "action",
      width: "150",
      slots: {
        default: ({ row }) =>
          h(TableAction, {
            actions: [
              {
                type: "link",
                icon: "ant-design:qrcode-outlined",
                tooltip: $t("common.showQrcode"),
                onClick: showQrCode.bind(null, row),
              },
              {
                type: "link",
                icon: "ant-design:cloud-download-outlined",
                tooltip: $t("common.download"),
                onClick: handleDownload.bind(null, row),
              },
              {
                type: "link",
                icon: "clarity:note-edit-line",
                tooltip: $t("common.edit"),
                onClick: openFormModal.bind(null, row),
              },
              {
                icon: "ant-design:delete-outlined",
                type: "link",
                color: "error",
                tooltip: $t("common.delete"),
                popConfirm: {
                  title: $t("common.deleteConfirm"),
                  placement: "left",
                  confirm: batchDelete.bind(null, [row.id]),
                },
              },
            ] as ActionItem[],
          }),
      },
    },
  ],
  height: "auto",
  keepSource: true,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getUpgradeWinVersionList({
          page: page.currentPage,
          pageSize: page.pageSize,
          ...formValues,
        });
        return {
          items: res.data.data,
          total: res.data.total,
        };
      },
    },
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions,
  gridOptions,
  gridEvents,
});

function openFormModal(record: any) {
  if (isPlainObject(record)) {
    formModalApi.setData({
      record,
      isUpdate: true,
      gridApi,
    });
  } else {
    formModalApi.setData({
      record: null,
      isUpdate: false,
      gridApi,
    });
  }
  formModalApi.open();
}

function handleBatchDelete() {
  Modal.confirm({
    title: $t("common.deleteConfirm"),
    async onOk() {
      const ids = gridApi.grid.getCheckboxRecords().map((item: any) => item.id);

      batchDelete(ids);
    },
  });
}

async function batchDelete(ids: any[]) {
  const result = await deleteUpgradeWinVersion({
    ids,
  });
  if (result.code === 0) {
    await gridApi.reload();
    showDeleteButton.value = false;
  }
}

async function handleDownload(record: any) {
  const link = document.createElement("a");
  link.href = record.cloudFilePath;
  link.download = record.versionName;
  link.click();
  link.remove();
  URL.revokeObjectURL(link.href);
}

function showQrCode(record: any) {
  currentQrCodeUrl.value = record.cloudFilePath;
  currentVersionName.value = record.versionName;
  qrCodeVisible.value = true;
}

const downloadQrCode = () => {
  // 获取二维码Canvas元素
  const qrCodeCanvas = document.querySelector(
    ".ant-qrcode canvas",
  ) as HTMLCanvasElement;
  if (qrCodeCanvas) {
    // 创建下载链接
    const link = document.createElement("a");
    link.download = `qrcode-${currentVersionName.value || "download"}.png`;
    link.href = qrCodeCanvas.toDataURL("image/png");
    link.click();
    link.remove();
  }
};
</script>

<template>
  <Page auto-content-height>
    <FormModal />
    <Modal v-model:visible="qrCodeVisible" :footer="null" width="300px">
      <div style="text-align: center; padding: 20px">
        <QRCode
          :value="currentQrCodeUrl"
          :size="200"
          color="#000"
          bg-color="#fff"
        />
        <div style="margin-top: 20px">
          <Button type="primary" @click="downloadQrCode">
            {{ $t("common.downloadQrcode") }}
          </Button>
        </div>
      </div>
    </Modal>
    <Grid>
      <template #toolbar-buttons>
        <Button
          v-show="showDeleteButton"
          danger
          type="primary"
          @click="handleBatchDelete"
        >
          {{ $t("common.delete") }}
        </Button>
      </template>

      <template #toolbar-tools>
        <Button type="primary" @click="openFormModal">
          {{ $t("upgrade.upgradeWinVersion.addUpgradeWinVersion") }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
