import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { $t } from "@vben/locales";

import { getUpgradeElectronList } from "#/api/upgrade/upgradeElectron";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeElectron.name"),
      field: "electronName",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.versionName"),
      field: "versionName",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.versionCode"),
      field: "versionCode",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.platform"),
      field: "platform",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.arch"),
      field: "arch",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.installVersionFileSize"),
      field: "installVersionFileSize",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.upgradeVersionFileSize"),
      field: "upgradeVersionFileSize",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.description"),
      field: "description",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeElectronVersion.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "tauriId",
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
      fieldName: "versionName",
      label: $t("upgrade.upgradeElectronVersion.versionName"),
      component: "Input",
    },
    {
      fieldName: "versionCode",
      label: $t("upgrade.upgradeElectronVersion.versionCode"),
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
      fieldName: "electronId",
      label: $t("upgrade.upgradeElectron.name"),
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
      fieldName: "versionName",
      label: $t("upgrade.upgradeElectronVersion.versionName"),
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
      label: $t("upgrade.upgradeElectronVersion.versionCode"),
      component: "InputNumber",
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        show: (values) => values.id,
        triggerFields: ["id"],
      },
    },
    {
      fieldName: "platform",
      label: $t("upgrade.upgradeElectronVersion.platform"),
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
      label: $t("upgrade.upgradeElectronVersion.arch"),
      rules: "required",
      component: "Select",
      componentProps: {
        allowClear: true,
        filterOption: true,
        options: [
          {
            label: "x64",
            value: "x64",
          },
          {
            label: "arm64",
            value: "arm64",
          },
        ],
        placeholder: "请选择",
        showSearch: true,
      },
    },
    {
      fieldName: "InstallCloudFileName",
      label: $t("upgrade.upgradeElectronVersion.installCloudFileName"),
      component: "Input",
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        show: (values) => values.id && ["darwin"].includes(values.platform),
        triggerFields: ["id", "platform"],
      },
    },
    {
      fieldName: "installCloudFileId",
      label: $t("upgrade.upgradeElectronVersion.installFileName"),
      component: "UploadDraggerOne",
      // rules: "required",
      componentProps: {
        multiple: false,
        provider: "cloud-default",
      },
      dependencies: {
        show: (values) => !values.id && ["darwin"].includes(values.platform),
        triggerFields: ["id", "platform"],
      },
    },
    {
      fieldName: "installSha512",
      label: $t("upgrade.upgradeElectronVersion.installSha512"),
      component: "Textarea",
      // rules: "required",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
        // showCount: true, // 显示字数统计（可选）
      },
      dependencies: {
        disabled(values) {
          return !!values.id;
        },
        show: (values) => !values.id && ["darwin"].includes(values.platform),
        triggerFields: ["id", "platform"],
      },
    },
    {
      fieldName: "cloudFileName",
      label: $t("upgrade.upgradeElectronVersion.cloudFileName"),
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
      label: $t("upgrade.upgradeElectronVersion.editFileName"),
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
      fieldName: "sha512",
      label: $t("upgrade.upgradeElectronVersion.sha512"),
      component: "Textarea",
      rules: "required",
      componentProps: {
        autoSize: { minRows: 5 }, // 自动调整高度（可选）
        // showCount: true, // 显示字数统计（可选）
      },
    },

    {
      fieldName: "description",
      label: $t("upgrade.upgradeElectronVersion.description"),
      component: "Textarea",
      componentProps: {
        autoSize: { minRows: 10 }, // 自动调整高度（可选）
        // showCount: true, // 显示字数统计（可选）
      },
    },
  ],
};
