import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { z } from "@vben/common-ui";
import { $t } from "@vben/locales";

import { getUpgradeConfigurationList } from "#/api/upgrade/upgradeConfiguration";
import { getUpgradeElectronList } from "#/api/upgrade/upgradeElectron";
import { getUpgradeFileList } from "#/api/upgrade/upgradeFile";
import { getUpgradeApkList } from "#/api/upgrade/upgradeApk";
import { getUpgradeLnxList } from "#/api/upgrade/upgradeLnx";
import { getUpgradeMacList } from "#/api/upgrade/upgradeMac";
import { getUpgradeTauriList } from "#/api/upgrade/upgradeTauri";
import { getUpgradeUrlList } from "#/api/upgrade/upgradeUrl";
import { getUpgradeWinList } from "#/api/upgrade/upgradeWin";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("sys.companySecret.accessKey"),
      field: "accessKey",
      width: 200,
    },
    {
      title: $t("sys.companySecret.secretKey"),
      field: "secretKey",
      width: 400,
    },
    {
      title: $t("sys.companySecret.enable"),
      field: "enable",
      slots: {
        default: (e) => {
          switch (e.row.enable) {
            case 1: {
              return $t("common.yes");
            }
            default: {
              return $t("common.no");
            }
          }
        },
      },
    },
    {
      field: "validityDatetime",
      title: $t("sys.companySecret.validityDatetimeTableTitle"),
      slots: {
        default: (e) => {
          return e.row.validityDatetime ?? $t("sys.companySecret.unlimited");
        },
      },
    },
    {
      field: "ruleDataUrl",
      title: $t("sys.companySecret.ruleDataTableTitle"),
      slots: {
        default: (e) => {
          return e.row.ruleDataUrl.length === 0 &&
          e.row.ruleDataFile.length === 0 &&
          e.row.ruleDataConfiguration.length === 0 &&
          e.row.ruleDataTauri.length === 0 &&
          e.row.ruleDataElectron.length === 0 &&
          e.row.ruleDataApk.length === 0 &&
          e.row.ruleDataMac.length === 0 &&
          e.row.ruleDataWin.length === 0 &&
          e.row.ruleDataLnx.length === 0
              ? $t("sys.companySecret.unlimited")
              : $t("common.yes");
        },
      },
    },
    {
      title: $t("sys.companySecret.description"),
      field: "description",
    },

    {
      title: $t("common.createTime"),
      field: "createdAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "name",
      label: $t("sys.companySecret.accessKey"),
      component: "Input",
    },
    {
      fieldName: "name",
      label: $t("sys.companySecret.secretKey"),
      component: "Input",
    },
    {
      fieldName: "enable",
      label: $t("sys.companySecret.enable"),
      component: "Select",
      componentProps: {
        options: [
          { label: $t("common.no"), value: 0 },
          { label: $t("common.yes"), value: 1 },
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
      label: $t("sys.companySecret.enable"),
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
      fieldName: "validityDatetime",
      label: $t("sys.companySecret.validityDatetimeTableTitle"),
      component: "DatePicker",
      // rules: "selectRequired",
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
      fieldName: "ruleDataUrl",
      label: $t("sys.companySecret.ruleDataUrl"),
      component: "CheckboxGroup",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeVersionType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeUrlList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.name}`,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.ruleDataUrl?.includes(item.id),
              };
            }),
          };
        },
        triggerFields: ["ruleDataUrl"],
      },
    },

    {
      fieldName: "ruleDataFile",
      label: $t("sys.companySecret.ruleDataFile"),
      component: "CheckboxGroup",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeVersionType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeFileList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.name}`,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.ruleDataFile?.includes(item.id),
              };
            }),
          };
        },
        triggerFields: ["ruleDataFile"],
      },
    },
    {
      fieldName: "ruleDataConfiguration",
      label: $t("sys.companySecret.ruleDataConfiguration"),
      component: "CheckboxGroup",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeVersionType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeConfigurationList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.name}`,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.ruleDataConfiguration?.includes(item.id),
              };
            }),
          };
        },
        triggerFields: ["ruleDataConfiguration"],
      },
    },
    {
      fieldName: "ruleDataTauri",
      label: $t("sys.companySecret.ruleDataTauri"),
      component: "CheckboxGroup",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeVersionType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeTauriList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.name}`,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.ruleDataTauri?.includes(item.id),
              };
            }),
          };
        },
        triggerFields: ["ruleDataTauri"],
      },
    },
    {
      fieldName: "ruleDataElectron",
      label: $t("sys.companySecret.ruleDataElectron"),
      component: "CheckboxGroup",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeVersionType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeElectronList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.name}`,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.ruleDataElectron?.includes(item.id),
              };
            }),
          };
        },
        triggerFields: ["ruleDataElectron"],
      },
    },
    {
      fieldName: "ruleDataApk",
      label: $t("sys.companySecret.ruleDataApk"),
      component: "CheckboxGroup",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeVersionType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeApkList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.name}`,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.ruleDataApk?.includes(item.id),
              };
            }),
          };
        },
        triggerFields: ["ruleDataApk"],
      },
    },
    {
      fieldName: "ruleDataMac",
      label: $t("sys.companySecret.ruleDataMac"),
      component: "CheckboxGroup",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeVersionType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeMacList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.name}`,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.ruleDataMac?.includes(item.id),
              };
            }),
          };
        },
        triggerFields: ["ruleDataMac"],
      },
    },
    {
      fieldName: "ruleDataWin",
      label: $t("sys.companySecret.ruleDataWin"),
      component: "CheckboxGroup",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeVersionType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeWinList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.name}`,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.ruleDataWin?.includes(item.id),
              };
            }),
          };
        },
        triggerFields: ["ruleDataWin"],
      },
    },
    {
      fieldName: "ruleDataLnx",
      label: $t("sys.companySecret.ruleDataLnx"),
      component: "CheckboxGroup",
      dependencies: {
        async componentProps(values) {
          if (values.upgradeVersionType === 0) {
            return {
              options: [],
            };
          }
          const res = await getUpgradeLnxList({
            page: 1,
            pageSize: 1000,
          });
          return {
            options: res.data.data.map((item: any) => {
              return {
                label: `${item.name}`,
                value: item.id,
                // 设置默认选中逻辑
                checked: values.ruleDataLnx?.includes(item.id),
              };
            }),
          };
        },
        triggerFields: ["ruleDataLnx"],
      },
    },
    {
      fieldName: "description",
      label: $t("sys.companySecret.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
      },
      rules: z.string().default("").optional(),
    },
  ],
};
