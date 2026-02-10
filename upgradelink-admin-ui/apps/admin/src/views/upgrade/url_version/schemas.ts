import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { $t } from "@vben/locales";

import { getUpgradeUrlList } from "#/api/upgrade/upgradeUrl";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeUrl.name"),
      field: "urlName",
      width: 120,
    },
    {
      title: $t("upgrade.upgradeUrlVersion.versionName"),
      field: "versionName",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeUrlVersion.versionCode"),
      field: "versionCode",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeUrlVersion.urlPath"),
      field: "urlPath",
    },
    {
      title: $t("upgrade.upgradeUrlVersion.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeUrlVersion.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "urlId",
      label: $t("upgrade.upgradeUrl.name"),
      component: "ApiSelect",
      componentProps: {
        api: getUpgradeUrlList,
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
      label: $t("upgrade.upgradeUrlVersion.versionName"),
      component: "Input",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeUrlVersion.versionCode"),
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
      fieldName: "urlId",
      label: $t("upgrade.upgradeUrl.name"),
      component: "ApiSelect",
      rules: "required",
      componentProps: {
        api: getUpgradeUrlList,
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
      fieldName: "urlPath",
      label: $t("upgrade.upgradeUrlVersion.urlPath"),
      help: $t("upgrade.upgradeUrlVersion.urlPathHelp"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "versionName",
      label: $t("upgrade.upgradeUrlVersion.versionName"),
      help: $t("upgrade.upgradeUrlVersion.versionNameHelp"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeUrlVersion.versionCode"),
      help: $t("upgrade.upgradeUrlVersion.versionCodeHelp"),
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
      label: $t("upgrade.upgradeUrlVersion.description"),
      component: "TextMarkdownEditor",
      componentProps: {
        placeholder: "请输入描述信息",
        autoSize: { minRows: 3 },
      },
    },
  ],
};
