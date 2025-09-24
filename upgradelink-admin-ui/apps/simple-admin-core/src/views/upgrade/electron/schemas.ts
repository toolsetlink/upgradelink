import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { h } from "vue";

import { $t } from "@vben/locales";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeElectron.key"),
      field: "key",
    },
    {
      title: $t("upgrade.upgradeElectron.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeElectron.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeElectron.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeElectron.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "key",
      label: $t("upgrade.upgradeElectron.key"),
      component: "Input",
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeElectron.name"),
      component: "Input",
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
      fieldName: "key",
      label: $t("upgrade.upgradeElectron.key"),
      component: "Input",
      rules: "required",
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        show: (values) => values.id,
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeElectron.name"),
      component: "Input",
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeElectron.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
    },
    {
      component: "Divider",
      fieldName: "_divider",
      formItemClass: "col-span-1",
      hideLabel: true,
      renderComponentContent: () => {
        return {
          default: () => h("div", "开源项目信息配置"),
        };
      },
    },
    {
      fieldName: "githubUrl",
      label: $t("upgrade.upgradeElectron.githubUrl"),
      component: "Input",
    },
  ],
};
