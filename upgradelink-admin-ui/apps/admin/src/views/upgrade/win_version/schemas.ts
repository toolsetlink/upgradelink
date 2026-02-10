import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { $t } from "@vben/locales";

import { getUpgradeWinList } from "#/api/upgrade/upgradeWin";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeWin.name"),
      field: "winName",
      width: 120,
    },
    {
      title: $t("upgrade.upgradeWinVersion.cloudFileName"),
      field: "cloudFileName",
    },
    {
      title: $t("upgrade.upgradeWinVersion.versionName"),
      field: "versionName",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeWinVersion.versionCode"),
      field: "versionCode",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeWinVersion.arch"),
      field: "arch",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeWinVersion.versionFileSize"),
      field: "versionFileSize",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeWinVersion.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeWinVersion.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "winId",
      label: $t("upgrade.upgradeWin.name"),
      component: "ApiSelect",
      componentProps: {
        api: getUpgradeWinList,
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
      label: $t("upgrade.upgradeWinVersion.versionName"),
      component: "Input",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeWinVersion.versionCode"),
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
      fieldName: "winId",
      label: $t("upgrade.upgradeWin.name"),
      component: "ApiSelect",
      rules: "required",
      componentProps: {
        api: getUpgradeWinList,
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
      label: $t("upgrade.upgradeWinVersion.cloudFileName"),
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
      label: $t("upgrade.upgradeWinVersion.versionFile"),
      component: "S3PresignedUpload",
      rules: "required",
      componentProps: {
        accept: [".exe"],
      },
      dependencies: {
        show: (values) => !values.id,
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "versionName",
      label: $t("upgrade.upgradeWinVersion.versionName"),
      help: $t("upgrade.upgradeWinVersion.versionNameHelp"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeWinVersion.versionCode"),
      help: $t("upgrade.upgradeWinVersion.versionCodeHelp"),
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
      fieldName: "arch",
      label: $t("upgrade.upgradeMacVersion.arch"),
      rules: "required",
      component: "Select",
      componentProps: {
        allowClear: true,
        filterOption: true,
        options: [
          {
            label: "x64",
            value: "x64",
          },
          {
            label: "arm64",
            value: "arm64",
          },
        ],
        placeholder: "请选择",
        showSearch: true,
      },
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeWinVersion.description"),
      component: "TextMarkdownEditor",
      componentProps: {
        placeholder: "请输入描述信息",
        autoSize: { minRows: 3 },
      },
    },
  ],
};
