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
      title: $t("upgrade.upgradeDevGroup.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeDevGroup.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeDevGroup.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "name",
      label: $t("upgrade.upgradeDevGroup.name"),
      component: "Input",
    },
    {
      fieldName: "createAt",
      label: $t("upgrade.upgradeDevGroup.createAt"),
      component: "InputNumber",
    },
    {
      fieldName: "updateAt",
      label: $t("upgrade.upgradeDevGroup.updateAt"),
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
      fieldName: "name",
      label: $t("upgrade.upgradeDevGroup.name"),
      component: "Input",
    },
  ],
};
