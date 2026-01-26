import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { $t } from "@vben/locales";

import { getUpgradeApkList } from "#/api/upgrade/upgradeApk";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeApk.name"),
      field: "apkName",
      width: 120,
    },
    {
      title: $t("upgrade.upgradeApkVersion.cloudFileName"),
      field: "cloudFileName",
    },
    {
      title: $t("upgrade.upgradeApkVersion.versionName"),
      field: "versionName",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeApkVersion.versionCode"),
      field: "versionCode",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeApkVersion.versionFileSize"),
      field: "versionFileSize",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeApkVersion.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeApkVersion.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "apkId",
      label: $t("upgrade.upgradeApk.name"),
      component: "ApiSelect",
      componentProps: {
        api: getUpgradeApkList,
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
      label: $t("upgrade.upgradeApkVersion.versionName"),
      component: "Input",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeApkVersion.versionCode"),
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
      fieldName: "apkId",
      label: $t("upgrade.upgradeApk.name"),
      component: "ApiSelect",
      rules: "required",
      componentProps: {
        api: getUpgradeApkList,
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
      label: $t("upgrade.upgradeApkVersion.cloudFileName"),
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
      label: $t("upgrade.upgradeApkVersion.versionFile"),
      component: "UploadDraggerOne",
      rules: "required",
      componentProps: {
        multiple: false,
        provider: "cloud-default",
        accept: ".apk",
      },
      dependencies: {
        show: (values) => !values.id,
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "versionName",
      label: $t("upgrade.upgradeApkVersion.versionName"),
      help: $t("upgrade.upgradeApkVersion.versionNameHelp"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeApkVersion.versionCode"),
      help: $t("upgrade.upgradeApkVersion.versionCodeHelp"),
      component: "InputNumber",
      componentProps: {
        style: { width: "200px" },
      },
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
      label: $t("upgrade.upgradeApkVersion.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 10 }, // 自动调整高度（可选）
        // showCount: true, // 显示字数统计（可选）
      },
    },
  ],
};
