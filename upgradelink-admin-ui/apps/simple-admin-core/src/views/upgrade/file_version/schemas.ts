import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { $t } from "@vben/locales";

import { getUpgradeFileList } from "#/api/upgrade/upgradeFile";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeFile.name"),
      field: "fileName",
    },
    {
      title: $t("upgrade.upgradeFileVersion.cloudFileName"),
      field: "cloudFileName",
    },
    {
      title: $t("upgrade.upgradeFileVersion.versionName"),
      field: "versionName",
    },
    {
      title: $t("upgrade.upgradeFileVersion.versionCode"),
      field: "versionCode",
    },
    {
      title: $t("upgrade.upgradeFileVersion.versionFileSize"),
      field: "versionFileSize",
    },
    {
      title: $t("upgrade.upgradeFileVersion.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeFileVersion.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeFileVersion.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "fileId",
      label: $t("upgrade.upgradeFile.name"),
      component: "ApiSelect",
      componentProps: {
        api: getUpgradeFileList,
        params: {
          page: 1,
          pageSize: 1000,
          name: "",
        },
        resultField: "data.data",
        labelField: "name",
        valueField: "id",
        multiple: false,
      },
    },
    {
      fieldName: "versionName",
      label: $t("upgrade.upgradeFileVersion.versionName"),
      component: "Input",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeFileVersion.versionCode"),
      component: "InputNumber",
    },
  ],
};

export const dataFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "id",
      label: "ID",
      component: "Input",
      dependencies: {
        show: false,
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "fileId",
      label: $t("upgrade.upgradeFile.name"),
      component: "ApiSelect",
      rules: "required",
      componentProps: {
        api: getUpgradeFileList,
        params: {
          page: 1,
          pageSize: 1000,
          name: "",
        },
        resultField: "data.data",
        labelField: "name",
        valueField: "id",
        multiple: false,
      },
    },
    {
      fieldName: "cloudFileName",
      label: $t("upgrade.upgradeFileVersion.cloudFileName"),
      component: "Input",
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        show: (values) => values.id,
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "cloudFileId",
      label: $t("upgrade.upgradeFileVersion.versionFile"),
      component: "UploadDraggerOne",
      rules: "required",
      componentProps: {
        multiple: false,
        provider: "cloud-default",
      },
      dependencies: {
        show: (values) => !values.id,
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "versionName",
      label: $t("upgrade.upgradeFileVersion.versionName"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeFileVersion.versionCode"),
      component: "InputNumber",
      rules: "required",
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeFileVersion.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 10 }, // 自动调整高度（可选）
        // showCount: true, // 显示字数统计（可选）
      },
    },
  ],
};
