import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { $t } from "@vben/locales";

export const tableColumns: VxeGridProps = {
  columns: [
    // {
    //   type: "checkbox",
    //   width: 60,
    // },
    {
      title: $t("sys.companySecret.accessKey"),
      field: "access_key",
    },
    {
      title: $t("sys.companySecret.secretKey"),
      field: "secret_key",
    },
    // {
    //   title: $t("common.createTime"),
    //   field: "createdAt",
    //   formatter: "formatDateTime",
    // },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [],
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
      fieldName: "access_key",
      label: $t("sys.companySecret.accessKey"),
      component: "Input",
    },
    {
      fieldName: "secret_key",
      label: $t("sys.companySecret.secretKey"),
      component: "Input",
    },
  ],
};
