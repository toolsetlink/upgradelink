import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { h } from "vue";

import { z } from "@vben/common-ui";
import { $t } from "@vben/locales";

import { getUpgradeConfigurationList } from "#/api/upgrade/upgradeConfiguration";
import { getUpgradeConfigurationVersionList } from "#/api/upgrade/upgradeConfigurationVersion";
import { getUpgradeDevGroupList } from "#/api/upgrade/upgradeDevGroup";
import { getUpgradeDevModelList } from "#/api/upgrade/upgradeDevModel";

import FlowLimitFields from "./flowlimit.vue";
import GrayFields from "./gray.vue";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeConfigurationUpgradeStrategy.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeConfiguration.name"),
      field: "configurationName",
    },
    {
      title: $t("upgrade.upgradeConfigurationVersion.versionName"),
      field: "configurationVersionName",
    },
    {
      title: $t("upgrade.upgradeConfigurationVersion.versionCode"),
      field: "configurationVersionCode",
    },
    {
      title: $t("upgrade.upgradeConfigurationUpgradeStrategy.upgradeType"),
      field: "upgradeType",
      slots: {
        default: (e) => {
          switch (e.row.upgradeType) {
            case 1: {
              return $t(
                "upgrade.upgradeConfigurationUpgradeStrategy.upgradeTypeOne",
              );
            }
            case 2: {
              return $t(
                "upgrade.upgradeConfigurationUpgradeStrategy.upgradeTypeTwo",
              );
            }
            case 3: {
              return $t(
                "upgrade.upgradeConfigurationUpgradeStrategy.upgradeTypeThree",
              );
            }
            default: {
              return $t(
                "upgrade.upgradeConfigurationUpgradeStrategy.upgradeTypeZero",
              );
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeConfigurationUpgradeStrategy.upgradeDevType"),
      field: "upgradeDevType",
      slots: {
        default: (e) => {
          switch (e.row.upgradeDevType) {
            case 1: {
              return $t(
                "upgrade.upgradeConfigurationUpgradeStrategy.upgradeDevTypeOne",
              );
            }
            case 2: {
              return $t(
                "upgrade.upgradeConfigurationUpgradeStrategy.upgradeDevTypeTwo",
              );
            }
            default: {
              return $t(
                "upgrade.upgradeConfigurationUpgradeStrategy.upgradeDevTypeZero",
              );
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeConfigurationUpgradeStrategy.beginDatetime"),
      field: "beginDatetime",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeConfigurationUpgradeStrategy.endDatetime"),
      field: "endDatetime",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeConfigurationUpgradeStrategy.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "name",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.name"),
      component: "Input",
    },
    {
      fieldName: "configurationId",
      label: $t("upgrade.upgradeConfiguration.name"),
      component: "ApiSelect",
      componentProps: {
        api: getUpgradeConfigurationList,
        params: {
          page: 1,
          pageSize: 1000,
          name: "",
        },
        resultField: "data.data",
        labelField: "name",
        valueField: "id",
        multiple: false,
      },
    },
    {
      fieldName: "enable",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.enable"),
      component: "Select",
      componentProps: {
        options: [
          { label: "否", value: 0 },
          { label: "是", value: 1 },
        ],
      },
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
      fieldName: "enable",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.enable"),
      component: "RadioButtonGroup",
      rules: "required",
      defaultValue: 1,
      componentProps: {
        options: [
          { label: $t("common.yes"), value: 1 },
          { label: $t("common.no"), value: 0 },
        ],
      },
    },
    {
      fieldName: "name",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.name"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
      rules: z.string().default("").optional(),
    },
    {
      fieldName: "configurationId",
      label: $t("upgrade.upgradeConfiguration.name"),
      component: "ApiSelect",
      rules: "required",
      componentProps: {
        api: getUpgradeConfigurationList,
        params: {
          page: 1,
          pageSize: 1000,
          name: "",
        },
        resultField: "data.data",
        labelField: "name",
        valueField: "id",
        multiple: false,
      },
    },
    {
      fieldName: "configurationVersionId",
      label: $t("upgrade.upgradeConfigurationVersion.versionName"),
      component: "ApiSelect",
      rules: "required",
      dependencies: {
        async componentProps(values) {
          if (!values.configurationId) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeConfigurationVersionList({
            page: 1,
            pageSize: 1000,
            configurationId: values.configurationId,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: item.versionName,
                value: item.id,
              };
            }),
          };
        },
        triggerFields: ["configurationId"],
      },
    },
    {
      fieldName: "beginDatetime",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.beginDatetime"),
      component: "DatePicker",
      rules: "selectRequired",
      componentProps: {
        // 传给后端的时间格式--
        valueFormat: "YYYY-MM-DD HH:mm:ss",
        // 显示的时间格式
        showTime: {
          format: "YYYY-MM-DD HH:mm:ss",
        },
        style: {
          width: "100%",
        },
      },
    },
    {
      fieldName: "endDatetime",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.endDatetime"),
      component: "DatePicker",
      rules: "selectRequired",
      componentProps: {
        // 传给后端的时间格式--
        valueFormat: "YYYY-MM-DD HH:mm:ss",
        // 显示的时间格式
        showTime: {
          format: "YYYY-MM-DD HH:mm:ss",
        },
        style: {
          width: "100%",
        },
      },
    },
    {
      fieldName: "upgradeType",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.upgradeType"),
      component: "RadioButtonGroup",
      defaultValue: 1,
      componentProps: {
        options: [
          {
            label: $t(
              "upgrade.upgradeConfigurationUpgradeStrategy.upgradeTypeOne",
            ),
            value: 1,
          },
          {
            label: $t(
              "upgrade.upgradeConfigurationUpgradeStrategy.upgradeTypeTwo",
            ),
            value: 2,
          },
          {
            label: $t(
              "upgrade.upgradeConfigurationUpgradeStrategy.upgradeTypeThree",
            ),
            value: 3,
          },
        ],
      },
    },
    {
      fieldName: "promptUpgradeContent",
      label: $t(
        "upgrade.upgradeConfigurationUpgradeStrategy.promptUpgradeContent",
      ),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
      dependencies: {
        show: (values) => [1].includes(values.upgradeType),
        triggerFields: ["upgradeType"],
      },
      rules: z.string().default("").optional(),
    },
    {
      component: "Divider",
      fieldName: "_divider",
      formItemClass: "col-span-1",
      hideLabel: true,
      renderComponentContent: () => {
        return {
          default: () => h("div", "升级条件"),
        };
      },
    },
    {
      fieldName: "upgradeDevType",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.upgradeDevType"),
      component: "RadioButtonGroup",
      defaultValue: 0,
      componentProps: {
        options: [
          {
            label: $t(
              "upgrade.upgradeConfigurationUpgradeStrategy.upgradeDevTypeZero",
            ),
            value: 0,
          },
          {
            label: $t(
              "upgrade.upgradeConfigurationUpgradeStrategy.upgradeDevTypeOne",
            ),
            value: 1,
          },
          {
            label: $t(
              "upgrade.upgradeConfigurationUpgradeStrategy.upgradeDevTypeTwo",
            ),
            value: 2,
          },
        ],
      },
    },
    {
      fieldName: "upgradeDevTypeZeroData",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.upgradeDevData"),
      component: "Input",
      dependencies: {
        // 当 upgradeDevType 为 1 或 2 时显示
        // show: (values) => [1, 2].includes(values.upgradeDevType),
        show: false,
        triggerFields: ["upgradeDevType"],
      },
      rules: z.string().default("").optional(),
    },
    {
      fieldName: "upgradeDevTypeOneData",
      component: "CheckboxGroup",
      rules: "selectRequired",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeDevType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeDevGroupList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: item.name,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.upgradeDevTypeOneData?.includes(item.id),
              };
            }),
          };
        },
        // 当 upgradeDevType 为 2 时显示
        show: (values) => [1].includes(values.upgradeDevType),
        triggerFields: ["upgradeDevType"],
      },
    },
    {
      fieldName: "upgradeDevTypeTwoData",
      component: "CheckboxGroup",
      rules: "selectRequired",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeDevType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeDevModelList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: item.name,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.upgradeDevTypeTwoData?.includes(item.id),
              };
            }),
          };
        },
        // 当 upgradeDevType 为 2 时显示
        show: (values) => [2].includes(values.upgradeDevType),
        triggerFields: ["upgradeDevType"],
      },
    },
    {
      fieldName: "upgradeVersionType",
      label: $t(
        "upgrade.upgradeConfigurationUpgradeStrategy.upgradeVersionType",
      ),
      component: "RadioButtonGroup",
      defaultValue: 0,
      componentProps: {
        options: [
          {
            label: $t(
              "upgrade.upgradeConfigurationUpgradeStrategy.upgradeVersionTypeZero",
            ),
            value: 0,
          },
          {
            label: $t(
              "upgrade.upgradeConfigurationUpgradeStrategy.upgradeVersionTypeOne",
            ),
            value: 1,
          },
        ],
      },
    },
    {
      fieldName: "upgradeVersionTypeZeroData",
      label: $t(
        "upgrade.upgradeConfigurationUpgradeStrategy.upgradeVersionData",
      ),
      component: "Input",
      dependencies: {
        show: false,
        triggerFields: ["upgradeVersionType"],
      },
      rules: z.string().default("").optional(),
    },
    {
      fieldName: "upgradeVersionTypeOneData",
      component: "CheckboxGroup",
      rules: "selectRequired",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeVersionType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeConfigurationVersionList({
            page: 1,
            pageSize: 1000,
            configurationId: values.configurationId,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.configurationName}-${item.versionName}-${item.versionCode}`,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.upgradeVersionTypeOneData?.includes(item.id),
              };
            }),
          };
        },
        // 当 upgradeVersionType 为 1 时显示
        show: (values) => [1].includes(values.upgradeVersionType),
        triggerFields: ["upgradeVersionType"],
      },
    },

    {
      component: "Divider",
      fieldName: "_divider",
      formItemClass: "col-span-1",
      hideLabel: true,
      renderComponentContent: () => {
        return {
          default: () => h("div", "流量策略"),
        };
      },
    },
    {
      fieldName: "isGray",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.isGray"),
      component: "RadioButtonGroup",
      rules: "required",
      defaultValue: 0,
      componentProps: {
        options: [
          { label: $t("common.yes"), value: 1 },
          { label: $t("common.no"), value: 0 },
        ],
      },
    },
    {
      fieldName: "grayDataList",
      label: "灰度策略",
      component: (e) => {
        return h(GrayFields, {
          ...e,
          value: e.modelValue,
        });
      },
      disabledOnChangeListener: false,
      dependencies: {
        show: (values) => [1].includes(values.isGray),
        triggerFields: ["isGray"],
      },
    },
    {
      fieldName: "isFlowLimit",
      label: $t("upgrade.upgradeConfigurationUpgradeStrategy.isFlowLimit"),
      component: "RadioButtonGroup",
      rules: "required",
      defaultValue: 0,
      componentProps: {
        options: [
          { label: $t("common.yes"), value: 1 },
          { label: $t("common.no"), value: 0 },
        ],
      },
    },
    {
      fieldName: "flowLimitDataList",
      label: "流控策略",
      component: (e) => {
        return h(FlowLimitFields, {
          ...e,
          value: e.modelValue,
        });
      },
      disabledOnChangeListener: false,
      dependencies: {
        show: (values) => [1].includes(values.isFlowLimit),
        triggerFields: ["isFlowLimit"],
      },
    },
  ],
};
