import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { z } from "@vben/common-ui";
import { $t } from "@vben/locales";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      title: $t("upgrade.upgradeCompanyTrafficPacket.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeCompanyTrafficPacket.key"),
      field: "key",
    },
    {
      title: $t("upgrade.upgradeCompanyTrafficPacket.startTime"),
      field: "startTime",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeCompanyTrafficPacket.endTime"),
      field: "endTime",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeCompanyTrafficPacket.initialSize"),
      field: "initialSize",
    },
    {
      title: $t("upgrade.upgradeCompanyTrafficPacket.remainingSize"),
      field: "remainingSize",
    },
    {
      title: $t("upgrade.upgradeCompanyTrafficPacket.status"),
      field: "status",
      slots: {
        default: (e) => {
          switch (e.row.status) {
            case 0: {
              return $t("upgrade.upgradeCompanyTrafficPacket.statusZero");
            }
            case 1: {
              return $t("upgrade.upgradeCompanyTrafficPacket.statusOne");
            }
            default: {
              return $t("upgrade.upgradeCompanyTrafficPacket.statusTwo");
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeCompanyTrafficPacket.exchangeTime"),
      field: "exchangeTime",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "status",
      label: $t("upgrade.upgradeCompanyIncome.status"),
      rules: z.number().max(10).optional(),
      component: "Select",
      componentProps: {
        options: [
          {
            label: $t("upgrade.upgradeCompanyTrafficPacket.statusZero"),
            value: 0,
          },
          {
            label: $t("upgrade.upgradeCompanyTrafficPacket.statusOne"),
            value: 1,
          },
          {
            label: $t("upgrade.upgradeCompanyTrafficPacket.statusTwo"),
            value: 2,
          },
        ],
      },
    },
  ],
};

export const dataFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "key",
      label: $t("upgrade.upgradeCompanyTrafficPacket.key"),
      component: "Input",
      rules: "required",
    },
  ],
};
