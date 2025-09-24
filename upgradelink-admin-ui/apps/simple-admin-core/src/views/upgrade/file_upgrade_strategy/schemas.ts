import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { h } from "vue";

import { z } from "@vben/common-ui";
import { $t } from "@vben/locales";

import { getUpgradeDevGroupList } from "#/api/upgrade/upgradeDevGroup";
import { getUpgradeDevModelList } from "#/api/upgrade/upgradeDevModel";
import { getUpgradeFileList } from "#/api/upgrade/upgradeFile";
import { getUpgradeFileVersionList } from "#/api/upgrade/upgradeFileVersion";

import FlowLimitFields from "./flowlimit.vue";
import GrayFields from "./gray.vue";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeFileUpgradeStrategy.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeFile.name"),
      field: "fileName",
    },
    {
      title: $t("upgrade.upgradeFileVersion.versionName"),
      field: "fileVersionName",
    },
    {
      title: $t("upgrade.upgradeFileVersion.versionCode"),
      field: "fileVersionCode",
    },
    {
      title: $t("upgrade.upgradeFileUpgradeStrategy.upgradeType"),
      field: "upgradeType",
      slots: {
        default: (e) => {
          switch (e.row.upgradeType) {
            case 1: {
              return $t("upgrade.upgradeFileUpgradeStrategy.upgradeTypeOne");
            }
            case 2: {
              return $t("upgrade.upgradeFileUpgradeStrategy.upgradeTypeTwo");
            }
            case 3: {
              return $t("upgrade.upgradeFileUpgradeStrategy.upgradeTypeThree");
            }
            default: {
              return $t("upgrade.upgradeFileUpgradeStrategy.upgradeTypeZero");
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeFileUpgradeStrategy.upgradeDevType"),
      field: "upgradeDevType",
      slots: {
        default: (e) => {
          switch (e.row.upgradeDevType) {
            case 1: {
              return $t("upgrade.upgradeFileUpgradeStrategy.upgradeDevTypeOne");
            }
            case 2: {
              return $t("upgrade.upgradeFileUpgradeStrategy.upgradeDevTypeTwo");
            }
            default: {
              return $t(
                "upgrade.upgradeFileUpgradeStrategy.upgradeDevTypeZero",
              );
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeFileUpgradeStrategy.beginDatetime"),
      field: "beginDatetime",
    },
    {
      title: $t("upgrade.upgradeFileUpgradeStrategy.endDatetime"),
      field: "endDatetime",
    },
    // {
    //   title: $t("upgrade.upgradeFileUpgradeStrategy.enable"),
    //   field: "enable",
    //   slots: {
    //     default: (e) =>
    //       h(Switch, {
    //         checked: e.row.enable === 1,
    //         onClick: () => {
    //           const newEnable = e.row.enable === 1 ? 2 : 1;
    //           updateUpgradeFileUpgradeStrategy({
    //             id: e.row.id,
    //             enable: newEnable,
    //           }).then(() => {
    //             e.row.enable = newEnable;
    //           });
    //         },
    //       }),
    //   },
    // },
    {
      title: $t("upgrade.upgradeFileUpgradeStrategy.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "name",
      label: $t("upgrade.upgradeFileUpgradeStrategy.name"),
      component: "Input",
    },
    {
      fieldName: "fileId",
      label: $t("upgrade.upgradeFile.name"),
      component: "ApiSelect",
      componentProps: {
        api: getUpgradeFileList,
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
      label: $t("upgrade.upgradeFileUpgradeStrategy.enable"),
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
      label: $t("upgrade.upgradeFileUpgradeStrategy.enable"),
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
      label: $t("upgrade.upgradeFileUpgradeStrategy.name"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeFileUpgradeStrategy.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
      rules: z.string().default("").optional(),
    },
    {
      fieldName: "fileId",
      label: $t("upgrade.upgradeFileUpgradeStrategy.fileId"),
      component: "ApiSelect",
      rules: "required",
      componentProps: {
        api: getUpgradeFileList,
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
      fieldName: "fileVersionId",
      label: $t("upgrade.upgradeFileVersion.versionName"),
      component: "ApiSelect",
      rules: "required",
      dependencies: {
        async componentProps(values) {
          if (!values.fileId) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeFileVersionList({
            page: 1,
            pageSize: 1000,
            fileId: values.fileId,
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
        triggerFields: ["fileId"],
      },
    },
    {
      fieldName: "beginDatetime",
      label: $t("upgrade.upgradeFileUpgradeStrategy.beginDatetime"),
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
      label: $t("upgrade.upgradeFileUpgradeStrategy.endDatetime"),
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
      label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeType"),
      component: "RadioButtonGroup",
      defaultValue: 1,
      componentProps: {
        options: [
          {
            label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeTypeOne"),
            value: 1,
          },
          {
            label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeTypeTwo"),
            value: 2,
          },
          {
            label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeTypeThree"),
            value: 3,
          },
        ],
      },
    },
    {
      fieldName: "promptUpgradeContent",
      label: $t("upgrade.upgradeFileUpgradeStrategy.promptUpgradeContent"),
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
      label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeDevType"),
      component: "RadioButtonGroup",
      defaultValue: 0,
      componentProps: {
        options: [
          {
            label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeDevTypeZero"),
            value: 0,
          },
          {
            label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeDevTypeOne"),
            value: 1,
          },
          {
            label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeDevTypeTwo"),
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
      label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeVersionType"),
      component: "RadioButtonGroup",
      defaultValue: 0,
      componentProps: {
        options: [
          {
            label: $t(
              "upgrade.upgradeFileUpgradeStrategy.upgradeVersionTypeZero",
            ),
            value: 0,
          },
          {
            label: $t(
              "upgrade.upgradeFileUpgradeStrategy.upgradeVersionTypeOne",
            ),
            value: 1,
          },
        ],
      },
    },
    {
      fieldName: "upgradeVersionTypeZeroData",
      label: $t("upgrade.upgradeFileUpgradeStrategy.upgradeVersionData"),
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
          const res = await getUpgradeFileVersionList({
            page: 1,
            pageSize: 1000,
            fileId: values.fileId,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.fileName}-${item.versionName}-${item.versionCode}`,
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
      label: $t("upgrade.upgradeFileUpgradeStrategy.isGray"),
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
      label: $t("upgrade.upgradeFileUpgradeStrategy.isFlowLimit"),
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
