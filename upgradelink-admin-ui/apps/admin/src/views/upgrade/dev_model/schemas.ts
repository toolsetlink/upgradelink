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
      title: $t("upgrade.upgradeDevModel.key"),
      field: "key",
    },
    {
      title: $t("upgrade.upgradeDevModel.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeDevModel.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeDevModel.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "key",
      label: $t("upgrade.upgradeDevModel.key"),
      component: "Input",
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeDevModel.name"),
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
      label: $t("upgrade.upgradeDevModel.key"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeDevModel.name"),
      component: "Input",
      rules: "required",
    },
  ],
};
