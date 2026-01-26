import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { $t } from "@vben/locales";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeLnx.key"),
      field: "key",
    },
    {
      title: $t("upgrade.upgradeLnx.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeLnx.packageName"),
      field: "packageName",
    },
    {
      title: $t("upgrade.upgradeLnx.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeLnx.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeLnx.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "key",
      label: $t("upgrade.upgradeLnx.key"),
      component: "Input",
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeLnx.name"),
      component: "Input",
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeLnx.packageName"),
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
      label: $t("upgrade.upgradeLnx.key"),
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
      label: $t("upgrade.upgradeLnx.name"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "packageName",
      label: $t("upgrade.upgradeLnx.packageName"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeLnx.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
    },
  ],
};
