import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { $t } from "@vben/locales";

import { getUpgradeConfigurationList } from "#/api/upgrade/upgradeConfiguration";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeConfiguration.name"),
      field: "configurationName",
    },
    {
      title: $t("upgrade.upgradeConfigurationVersion.versionName"),
      field: "versionName",
    },
    {
      title: $t("upgrade.upgradeConfigurationVersion.versionCode"),
      field: "versionCode",
    },
    {
      title: $t("upgrade.upgradeConfigurationVersion.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeConfigurationVersion.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeConfigurationVersion.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "configurationId",
      label: $t("upgrade.upgradeConfiguration.name"),
      component: "ApiSelect",
      componentProps: {
        api: getUpgradeConfigurationList,
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
      label: $t("upgrade.upgradeConfigurationVersion.versionName"),
      component: "Input",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeConfigurationVersion.versionCode"),
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
      fieldName: "configurationId",
      label: $t("upgrade.upgradeConfiguration.name"),
      component: "ApiSelect",
      rules: "required",
      componentProps: {
        api: getUpgradeConfigurationList,
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
      fieldName: "content",
      label: $t("upgrade.upgradeConfigurationVersion.content"),
      rules: "required",
      component: "CodeEditor",
    },
    {
      fieldName: "versionName",
      label: $t("upgrade.upgradeConfigurationVersion.versionName"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeConfigurationVersion.versionCode"),
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
      label: $t("upgrade.upgradeConfigurationVersion.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 3 }, // 自动调整高度（可选）
        // showCount: true, // 显示字数统计（可选）
      },
    },
  ],
};
