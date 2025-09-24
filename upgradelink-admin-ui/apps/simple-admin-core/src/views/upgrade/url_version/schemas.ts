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
    },
    {
      title: $t("upgrade.upgradeUrlVersion.versionName"),
      field: "versionName",
    },
    {
      title: $t("upgrade.upgradeUrlVersion.versionCode"),
      field: "versionCode",
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
      title: $t("upgrade.upgradeUrlVersion.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
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
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "versionName",
      label: $t("upgrade.upgradeUrlVersion.versionName"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeUrlVersion.versionCode"),
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
      label: $t("upgrade.upgradeUrlVersion.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 10 }, // 自动调整高度（可选）
        // showCount: true, // 显示字数统计（可选）
      },
    },
  ],
};
