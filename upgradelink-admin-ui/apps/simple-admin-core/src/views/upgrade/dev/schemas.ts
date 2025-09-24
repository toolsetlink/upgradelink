import type { VbenFormProps } from "@vben/common-ui";

import type { VxeGridProps } from "#/adapter/vxe-table";

import { $t } from "@vben/locales";

import { getUpgradeDevGroupList } from "#/api/upgrade/upgradeDevGroup";

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: "checkbox",
      width: 60,
    },
    {
      title: $t("upgrade.upgradeDev.key"),
      field: "key",
    },
    // {
    //   title: $t("upgrade.upgradeDev.devGroupIds"),
    //   field: "devGroupIds",
    // },
    {
      title: $t("upgrade.upgradeDev.createAt"),
      field: "createAt",
      formatter: "formatDateTime",
    },
    {
      title: $t("upgrade.upgradeDev.updateAt"),
      field: "updateAt",
      formatter: "formatDateTime",
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: "key",
      label: $t("upgrade.upgradeDev.key"),
      component: "Input",
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
      fieldName: "key",
      label: $t("upgrade.upgradeDev.key"),
      component: "Input",
    },
    // {
    //   fieldName: "devGroupIds",
    //   label: $t("upgrade.upgradeDev.devGroupName"),
    //   component: "ApiSelect",
    //   componentProps: {
    //     api: getUpgradeDevGroupList,
    //     params: {
    //       page: 1,
    //       pageSize: 1000,
    //       name: "",
    //     },
    //     resultField: "data.data",
    //     labelField: "name",
    //     valueField: "id",
    //     multiple: false,
    //   },
    // },
    {
      label: $t("upgrade.upgradeDev.devGroupName"),
      fieldName: "devGroupIds",
      component: "CheckboxGroup",
      rules: "selectRequired",
      dependencies: {
        async componentProps(values) {
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
                checked: values.devGroupIds?.includes(item.id),
              };
            }),
          };
        },
        triggerFields: ["devGroupIds"],
      },
    },
  ],
};
