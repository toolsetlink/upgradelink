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
      title: $t("upgrade.upgradeFile.key"),
      field: "key",
    },
    {
      title: $t("upgrade.upgradeFile.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeFile.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeFile.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeFile.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "key",
      label: $t("upgrade.upgradeFile.key"),
      component: "Input",
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeFile.name"),
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
      label: $t("upgrade.upgradeFile.key"),
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
      label: $t("upgrade.upgradeFile.name"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeFile.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
    },
  ],
};
