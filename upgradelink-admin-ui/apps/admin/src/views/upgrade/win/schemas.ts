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
      title: $t("upgrade.upgradeWin.key"),
      field: "key",
    },
    {
      title: $t("upgrade.upgradeWin.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeWin.packageName"),
      field: "packageName",
    },
    {
      title: $t("upgrade.upgradeWin.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeWin.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeWin.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "key",
      label: $t("upgrade.upgradeWin.key"),
      component: "Input",
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeWin.name"),
      component: "Input",
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeWin.packageName"),
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
      label: $t("upgrade.upgradeWin.key"),
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
      label: $t("upgrade.upgradeWin.name"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "packageName",
      label: $t("upgrade.upgradeWin.packageName"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeWin.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
    },
  ],
};
