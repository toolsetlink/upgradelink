import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { z } from "@vben/common-ui";
import { $t } from "@vben/locales";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      title: $t("upgrade.upgradeCompanyIncome.incomeType"),
      field: "incomeType",
      slots: {
        default: (e) => {
          switch (e.row.incomeType) {
            case 0: {
              return $t("upgrade.upgradeCompanyIncome.incomeTypeZero");
            }
            default: {
              return $t("upgrade.upgradeCompanyIncome.incomeTypeOne");
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeCompanyIncome.incomeAmount"),
      field: "incomeAmount",
      // 添加 formatter 函数进行单位转换
      formatter: (value) => {
        let num = 0;
        // 处理对象类型的值
        // eslint-disable-next-line unicorn/prefer-ternary
        if (typeof value === "object" && value !== null) {
          // 假设对象中存储数值的属性是 value 或其他字段
          // 根据实际结构修改这里的属性名
          num = Number(value.cellValue || value.cellValue || value);
        } else {
          // 非对象类型直接转换
          num = Number(value);
        }

        // 检查是否为有效数字
        if (Number.isNaN(num)) {
          return "0.00";
        }

        // 分转元计算
        const result = num / 100;
        return result.toFixed(2);
      },
    },
    {
      title: $t("upgrade.upgradeCompanyIncome.incomeTime"),
      field: "incomeTime",
      formatter: "formatDate",
    },
    {
      title: $t("upgrade.upgradeCompanyIncome.remark"),
      field: "remark",
    },
    {
      title: $t("upgrade.upgradeCompanyIncome.status"),
      field: "status",
      slots: {
        default: (e) => {
          switch (e.row.status) {
            case 0: {
              return $t("upgrade.upgradeCompanyIncome.statusZero");
            }
            default: {
              return $t("upgrade.upgradeCompanyIncome.statusOne");
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeCompanyIncome.createAt"),
      field: "createAt",
      formatter: "formatDate",
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
          { label: $t("upgrade.upgradeCompanyIncome.statusZero"), value: 0 },
          { label: $t("upgrade.upgradeCompanyIncome.statusOne"), value: 1 },
        ],
      },
    },
  ],
};

export const dataFormSchemas: VbenFormProps = {
  schema: [],
};
