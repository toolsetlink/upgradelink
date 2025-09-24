import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { h } from "vue";

import { z } from "@vben/common-ui";
import { $t } from "@vben/locales";

import { getUpgradeDevGroupList } from "#/api/upgrade/upgradeDevGroup";
import { getUpgradeDevModelList } from "#/api/upgrade/upgradeDevModel";
import { getUpgradeUrlList } from "#/api/upgrade/upgradeUrl";
import { getUpgradeUrlVersionList } from "#/api/upgrade/upgradeUrlVersion";

import FlowLimitFields from "./flowlimit.vue";
import GrayFields from "./gray.vue";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeUrlUpgradeStrategy.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeUrl.name"),
      field: "urlName",
    },
    {
      title: $t("upgrade.upgradeUrlVersion.versionName"),
      field: "urlVersionName",
    },
    {
      title: $t("upgrade.upgradeUrlVersion.versionCode"),
      field: "urlVersionCode",
    },
    {
      title: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeType"),
      field: "upgradeType",
      slots: {
        default: (e) => {
          switch (e.row.upgradeType) {
            case 1: {
              return $t("upgrade.upgradeUrlUpgradeStrategy.upgradeTypeOne");
            }
            case 2: {
              return $t("upgrade.upgradeUrlUpgradeStrategy.upgradeTypeTwo");
            }
            case 3: {
              return $t("upgrade.upgradeUrlUpgradeStrategy.upgradeTypeThree");
            }
            default: {
              return $t("upgrade.upgradeUrlUpgradeStrategy.upgradeTypeZero");
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeDevType"),
      field: "upgradeDevType",
      slots: {
        default: (e) => {
          switch (e.row.upgradeDevType) {
            case 1: {
              return $t("upgrade.upgradeUrlUpgradeStrategy.upgradeDevTypeOne");
            }
            case 2: {
              return $t("upgrade.upgradeUrlUpgradeStrategy.upgradeDevTypeTwo");
            }
            default: {
              return $t("upgrade.upgradeUrlUpgradeStrategy.upgradeDevTypeZero");
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeUrlUpgradeStrategy.beginDatetime"),
      field: "beginDatetime",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeUrlUpgradeStrategy.endDatetime"),
      field: "endDatetime",
      formatter: "formatDateTime",
    },
    // {
    //   title: $t("upgrade.upgradeUrlUpgradeStrategy.enable"),
    //   field: "enable",
    //   slots: {
    //     default: (e) =>
    //       h(Switch, {
    //         checked: e.row.enable === 1,
    //         onClick: () => {
    //           const newEnable = e.row.enable === 1 ? 2 : 1;
    //           updateUpgradeUrlUpgradeStrategy({
    //             id: e.row.id,
    //             enable: newEnable,
    //           }).then(() => {
    //             e.row.enable = newEnable;
    //           });
    //         },
    //       }),
    //   },
    // },
    // {
    //   title: $t('upgrade.upgradeUrlUpgradeStrategy.createAt'),
    //   field: 'createAt',
    //   formatter: 'formatDateTime',
    // },
    {
      title: $t("upgrade.upgradeUrlUpgradeStrategy.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "name",
      label: $t("upgrade.upgradeUrlUpgradeStrategy.name"),
      component: "Input",
    },
    {
      fieldName: "urlId",
      label: $t("upgrade.upgradeUrl.name"),
      component: "ApiSelect",
      componentProps: {
        api: getUpgradeUrlList,
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
      label: $t("upgrade.upgradeUrlUpgradeStrategy.enable"),
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
      label: $t("upgrade.upgradeUrlUpgradeStrategy.enable"),
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
      label: $t("upgrade.upgradeUrlUpgradeStrategy.name"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeUrlUpgradeStrategy.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
      rules: z.string().default("").optional(),
    },
    {
      fieldName: "urlId",
      label: $t("upgrade.upgradeUrl.name"),
      component: "ApiSelect",
      rules: "required",
      componentProps: {
        api: getUpgradeUrlList,
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
      fieldName: "urlVersionId",
      label: $t("upgrade.upgradeUrlVersion.versionName"),
      component: "ApiSelect",
      rules: "required",
      dependencies: {
        async componentProps(values) {
          if (!values.urlId) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeUrlVersionList({
            page: 1,
            pageSize: 1000,
            urlId: values.urlId,
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
        triggerFields: ["urlId"],
      },
    },
    {
      fieldName: "beginDatetime",
      label: $t("upgrade.upgradeUrlUpgradeStrategy.beginDatetime"),
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
      label: $t("upgrade.upgradeUrlUpgradeStrategy.endDatetime"),
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
      label: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeType"),
      component: "RadioButtonGroup",
      defaultValue: 1,
      componentProps: {
        options: [
          {
            label: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeTypeOne"),
            value: 1,
          },
          {
            label: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeTypeTwo"),
            value: 2,
          },
          {
            label: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeTypeThree"),
            value: 3,
          },
        ],
      },
    },
    {
      fieldName: "promptUpgradeContent",
      label: $t("upgrade.upgradeUrlUpgradeStrategy.promptUpgradeContent"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
      // dependencies: {
      //   show: (values) => [1].includes(values.upgradeType),
      //   triggerFields: ["upgradeType"],
      // },
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
      label: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeDevType"),
      component: "RadioButtonGroup",
      defaultValue: 0,
      componentProps: {
        options: [
          {
            label: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeDevTypeZero"),
            value: 0,
          },
          {
            label: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeDevTypeOne"),
            value: 1,
          },
          {
            label: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeDevTypeTwo"),
            value: 2,
          },
        ],
      },
    },
    {
      fieldName: "upgradeDevTypeZeroData",
      label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeDevData"),
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
      label: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeVersionType"),
      component: "RadioButtonGroup",
      defaultValue: 0,
      componentProps: {
        options: [
          {
            label: $t(
              "upgrade.upgradeUrlUpgradeStrategy.upgradeVersionTypeZero",
            ),
            value: 0,
          },
          {
            label: $t(
              "upgrade.upgradeUrlUpgradeStrategy.upgradeVersionTypeOne",
            ),
            value: 1,
          },
        ],
      },
    },
    {
      fieldName: "upgradeVersionTypeZeroData",
      label: $t("upgrade.upgradeUrlUpgradeStrategy.upgradeVersionData"),
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
          const res = await getUpgradeUrlVersionList({
            page: 1,
            pageSize: 1000,
            urlId: values.urlId,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.urlName}-${item.versionName}-${item.versionCode}`,
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
      label: $t("upgrade.upgradeUrlUpgradeStrategy.isGray"),
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
      label: $t("upgrade.upgradeUrlUpgradeStrategy.isFlowLimit"),
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
