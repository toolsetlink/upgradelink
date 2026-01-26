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
      width: 120,
    },
    {
      title: $t("upgrade.upgradeConfigurationVersion.versionName"),
      field: "versionName",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeConfigurationVersion.versionCode"),
      field: "versionCode",
      width: 100,
    },
    {
      title: $t("upgrade.upgradeConfigurationVersion.description"),
      field: "description",
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
      component: "CodeEditor",
      rules: "required",
    },
    {
      fieldName: "versionName",
      label: $t("upgrade.upgradeConfigurationVersion.versionName"),
      help: $t("upgrade.upgradeConfigurationVersion.versionNameHelp"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeConfigurationVersion.versionCode"),
      help: $t("upgrade.upgradeConfigurationVersion.versionCodeHelp"),
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
      label: $t("upgrade.upgradeConfigurationVersion.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 3 }, // 自动调整高度（可选）
        // showCount: true, // 显示字数统计（可选）
      },
    },
  ],
};
