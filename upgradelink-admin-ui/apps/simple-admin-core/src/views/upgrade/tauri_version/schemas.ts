import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { $t } from "@vben/locales";

import { getUpgradeTauriList } from "#/api/upgrade/upgradeTauri";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeTauri.name"),
      field: "tauriName",
    },
    {
      title: $t("upgrade.upgradeTauriVersion.target"),
      field: "target",
    },
    {
      title: $t("upgrade.upgradeTauriVersion.arch"),
      field: "arch",
    },
    {
      title: $t("upgrade.upgradeTauriVersion.versionName"),
      field: "versionName",
    },
    {
      title: $t("upgrade.upgradeTauriVersion.versionCode"),
      field: "versionCode",
    },
    {
      title: $t("upgrade.upgradeTauriVersion.installVersionFileSize"),
      field: "installVersionFileSize",
    },
    {
      title: $t("upgrade.upgradeTauriVersion.upgradeVersionFileSize"),
      field: "upgradeVersionFileSize",
    },
    {
      title: $t("upgrade.upgradeTauriVersion.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeTauriVersion.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeTauriVersion.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "tauriId",
      label: $t("upgrade.upgradeTauri.name"),
      component: "ApiSelect",
      componentProps: {
        api: getUpgradeTauriList,
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
      fieldName: "versionName",
      label: $t("upgrade.upgradeTauriVersion.versionName"),
      component: "Input",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeTauriVersion.versionCode"),
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
      fieldName: "tauriId",
      label: $t("upgrade.upgradeTauri.name"),
      component: "ApiSelect",
      rules: "required",
      componentProps: {
        api: getUpgradeTauriList,
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
      fieldName: "versionName",
      label: $t("upgrade.upgradeTauriVersion.versionName"),
      component: "Input",
      rules: "required",
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeTauriVersion.versionCode"),
      component: "InputNumber",
      // rules: "required",
      // componentProps: {
      //   disabled: true,
      // },
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        show: (values) => values.id,
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "target",
      label: $t("upgrade.upgradeTauriVersion.target"),
      rules: "required",
      component: "Select",
      componentProps: {
        allowClear: true,
        filterOption: true,
        options: [
          {
            label: "linux",
            value: "linux",
          },
          {
            label: "darwin",
            value: "darwin",
          },
          {
            label: "windows",
            value: "windows",
          },
        ],
        placeholder: "请选择",
        showSearch: true,
      },
    },
    {
      fieldName: "arch",
      label: $t("upgrade.upgradeTauriVersion.arch"),
      rules: "required",
      component: "Select",
      componentProps: {
        allowClear: true,
        filterOption: true,
        options: [
          {
            label: "x86_64",
            value: "x86_64",
          },
          {
            label: "i686",
            value: "i686",
          },
          {
            label: "aarch64",
            value: "aarch64",
          },
          {
            label: "armv7",
            value: "armv7",
          },
        ],
        placeholder: "请选择",
        showSearch: true,
      },
    },
    {
      fieldName: "InstallCloudFileName",
      label: $t("upgrade.upgradeTauriVersion.installCloudFileName"),
      component: "Input",
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        show: (values) => values.id && ["darwin"].includes(values.target),
        triggerFields: ["id", "target"],
      },
    },
    {
      fieldName: "installCloudFileId",
      label: $t("upgrade.upgradeTauriVersion.installFileName"),
      component: "UploadDraggerOne",
      // rules: "required",
      componentProps: {
        multiple: false,
        provider: "cloud-default",
      },
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        show: (values) => !values.id && ["darwin"].includes(values.target),
        triggerFields: ["id", "target"],
      },
    },
    {
      fieldName: "cloudFileName",
      label: $t("upgrade.upgradeTauriVersion.cloudFileName"),
      component: "Input",
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        show: (values) => values.id,
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "cloudFileId",
      label: $t("upgrade.upgradeTauriVersion.editFileName"),
      component: "UploadDraggerOne",
      rules: "required",
      componentProps: {
        multiple: false,
        provider: "cloud-default",
      },
      dependencies: {
        show: (values) => !values.id,
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "signature",
      label: $t("upgrade.upgradeTauriVersion.signature"),
      component: "Textarea",
      rules: "required",
      componentProps: {
        autoSize: { minRows: 10 }, // 自动调整高度（可选）
        // showCount: true, // 显示字数统计（可选）
      },
    },
    {
      fieldName: "description",
      label: $t("upgrade.upgradeTauriVersion.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 10 }, // 自动调整高度（可选）
        // showCount: true, // 显示字数统计（可选）
      },
    },
  ],
};
