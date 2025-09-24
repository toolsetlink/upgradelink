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
      title: $t("upgrade.upgradeApk.key"),
      field: "key",
    },
    {
      title: $t("upgrade.upgradeApk.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeApk.packageName"),
      field: "packageName",
    },
    {
      title: $t("upgrade.upgradeApk.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeApk.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeApk.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "key",
      label: $t("upgrade.upgradeApk.key"),
      component: "Input",
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeApk.name"),
      component: "Input",
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeApk.packageName"),
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
      label: $t("upgrade.upgradeApk.key"),
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
      label: $t("upgrade.upgradeApk.name"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "packageName",
      label: $t("upgrade.upgradeApk.packageName"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeApk.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
    },
  ],
};
