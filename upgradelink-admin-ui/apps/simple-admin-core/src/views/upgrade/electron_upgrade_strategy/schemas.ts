import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { h } from "vue";

import { z } from "@vben/common-ui";
import { $t } from "@vben/locales";

import { getUpgradeDevGroupList } from "#/api/upgrade/upgradeDevGroup";
import { getUpgradeDevModelList } from "#/api/upgrade/upgradeDevModel";
import { getUpgradeElectronList } from "#/api/upgrade/upgradeElectron";
import { getUpgradeElectronVersionList } from "#/api/upgrade/upgradeElectronVersion";
import FlowLimitFields from "#/views/upgrade/electron_upgrade_strategy/flowlimit.vue";
import GrayFields from "#/views/upgrade/electron_upgrade_strategy/gray.vue";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeElectronUpgradeStrategy.name"),
      field: "name",
    },
    {
      title: $t("upgrade.upgradeElectron.name"),
      field: "electronName",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.platform"),
      field: "electronPlatform",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.arch"),
      field: "electronArch",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.versionName"),
      field: "electronVersionName",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.versionCode"),
      field: "electronVersionCode",
    },
    {
      title: $t("upgrade.upgradeElectronUpgradeStrategy.upgradeType"),
      field: "upgradeType",
      slots: {
        default: (e) => {
          switch (e.row.upgradeType) {
            case 1: {
              return $t(
                "upgrade.upgradeElectronUpgradeStrategy.upgradeTypeOne",
              );
            }
            case 2: {
              return $t(
                "upgrade.upgradeElectronUpgradeStrategy.upgradeTypeTwo",
              );
            }
            case 3: {
              return $t(
                "upgrade.upgradeElectronUpgradeStrategy.upgradeTypeThree",
              );
            }
            default: {
              return $t(
                "upgrade.upgradeElectronUpgradeStrategy.upgradeTypeZero",
              );
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeElectronUpgradeStrategy.upgradeDevType"),
      field: "upgradeDevType",
      slots: {
        default: (e) => {
          switch (e.row.upgradeDevType) {
            case 1: {
              return $t(
                "upgrade.upgradeElectronUpgradeStrategy.upgradeDevTypeOne",
              );
            }
            case 2: {
              return $t(
                "upgrade.upgradeElectronUpgradeStrategy.upgradeDevTypeTwo",
              );
            }
            default: {
              return $t(
                "upgrade.upgradeElectronUpgradeStrategy.upgradeDevTypeZero",
              );
            }
          }
        },
      },
    },
    {
      title: $t("upgrade.upgradeElectronUpgradeStrategy.beginDatetime"),
      field: "beginDatetime",
    },
    {
      title: $t("upgrade.upgradeElectronUpgradeStrategy.endDatetime"),
      field: "endDatetime",
    },
    {
      title: $t("upgrade.upgradeElectronUpgradeStrategy.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "name",
      label: $t("upgrade.upgradeElectronUpgradeStrategy.name"),
      component: "Input",
    },
    {
      fieldName: "electronId",
      label: $t("upgrade.upgradeElectron.name"),
      component: "ApiSelect",
      componentProps: {
        api: getUpgradeElectronList,
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
      label: $t("upgrade.upgradeElectronUpgradeStrategy.enable"),
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
      label: $t("upgrade.upgradeElectronUpgradeStrategy.enable"),
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
      label: $t("upgrade.upgradeElectronUpgradeStrategy.name"),
      component: "Input",
      rules: "required",
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeElectronUpgradeStrategy.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
      rules: z.string().default("").optional(),
    },
    {
      fieldName: "electronId",
      label: $t("upgrade.upgradeElectronUpgradeStrategy.electronId"),
      component: "ApiSelect",
      rules: "required",
      componentProps: {
        api: getUpgradeElectronList,
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
      fieldName: "electronVersionId",
      label: $t("upgrade.upgradeElectronVersion.versionName"),
      component: "ApiSelect",
      rules: "required",
      dependencies: {
        async componentProps(values) {
          if (!values.electronId) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeElectronVersionList({
            page: 1,
            pageSize: 1000,
            electronId: values.electronId,
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
        triggerFields: ["electronId"],
      },
    },
    {
      fieldName: "beginDatetime",
      label: $t("upgrade.upgradeElectronUpgradeStrategy.beginDatetime"),
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
      label: $t("upgrade.upgradeElectronUpgradeStrategy.endDatetime"),
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
      label: $t("upgrade.upgradeElectronUpgradeStrategy.upgradeType"),
      component: "RadioButtonGroup",
      defaultValue: 1,
      componentProps: {
        options: [
          {
            label: $t("upgrade.upgradeElectronUpgradeStrategy.upgradeTypeOne"),
            value: 1,
          },
          {
            label: $t("upgrade.upgradeElectronUpgradeStrategy.upgradeTypeTwo"),
            value: 2,
          },
          {
            label: $t(
              "upgrade.upgradeElectronUpgradeStrategy.upgradeTypeThree",
            ),
            value: 3,
          },
        ],
      },
    },
    {
      fieldName: "promptUpgradeContent",
      label: $t("upgrade.upgradeElectronUpgradeStrategy.promptUpgradeContent"),
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
      label: $t("upgrade.upgradeElectronUpgradeStrategy.upgradeDevType"),
      component: "RadioButtonGroup",
      defaultValue: 0,
      componentProps: {
        options: [
          {
            label: $t(
              "upgrade.upgradeElectronUpgradeStrategy.upgradeDevTypeZero",
            ),
            value: 0,
          },
          {
            label: $t(
              "upgrade.upgradeElectronUpgradeStrategy.upgradeDevTypeOne",
            ),
            value: 1,
          },
          {
            label: $t(
              "upgrade.upgradeElectronUpgradeStrategy.upgradeDevTypeTwo",
            ),
            value: 2,
          },
        ],
      },
    },
    {
      fieldName: "upgradeDevTypeZeroData",
      label: $t("upgrade.upgradeElectronUpgradeStrategy.upgradeDevData"),
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
      label: $t("upgrade.upgradeElectronUpgradeStrategy.upgradeVersionType"),
      component: "RadioButtonGroup",
      defaultValue: 0,
      componentProps: {
        options: [
          {
            label: $t(
              "upgrade.upgradeElectronUpgradeStrategy.upgradeVersionTypeZero",
            ),
            value: 0,
          },
          {
            label: $t(
              "upgrade.upgradeElectronUpgradeStrategy.upgradeVersionTypeOne",
            ),
            value: 1,
          },
        ],
      },
    },
    {
      fieldName: "upgradeVersionTypeZeroData",
      label: $t("upgrade.upgradeElectronUpgradeStrategy.upgradeVersionData"),
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
          const res = await getUpgradeElectronVersionList({
            page: 1,
            pageSize: 1000,
            electronId: values.electronId,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.electronName}-${item.target}-${item.arch}-${item.versionName}-${item.versionCode}`,
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
      label: $t("upgrade.upgradeElectronUpgradeStrategy.isGray"),
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
      label: $t("upgrade.upgradeElectronUpgradeStrategy.isFlowLimit"),
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
