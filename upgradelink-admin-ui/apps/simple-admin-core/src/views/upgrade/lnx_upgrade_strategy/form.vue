<script lang="ts" setup>
import type { UpgradeLnxUpgradeStrategyInfo } from "#/api/upgrade/model/upgradeLnxUpgradeStrategyModel";

import { ref } from "vue";

import { useVbenDrawer } from "@vben/common-ui";
import { $t } from "@vben/locales";

import { message } from "ant-design-vue";

import { useVbenForm } from "#/adapter/form";
import {
  createUpgradeLnxUpgradeStrategy,
  updateUpgradeLnxUpgradeStrategy,
} from "#/api/upgrade/upgradeLnxUpgradeStrategy";

import { dataFormSchemas } from "./schemas";

defineOptions({
  name: "UpgradeLnxUpgradeStrategyForm",
});

const record = ref();
const isUpdate = ref(false);
const gridApi = ref();

async function onSubmit(values: Record<string, any>) {
  const result = isUpdate.value
    ? await updateUpgradeLnxUpgradeStrategy(
        values as UpgradeLnxUpgradeStrategyInfo,
      )
    : await createUpgradeLnxUpgradeStrategy(
        values as UpgradeLnxUpgradeStrategyInfo,
      );
  if (result.code === 0) {
    message.success(result.msg);
    gridApi.value.reload();
  }
}

const [Form, formApi] = useVbenForm({
  handleSubmit: onSubmit,
  schema: [...(dataFormSchemas.schema as any)],
  showDefaultActions: false,
  layout: "vertical",
});

const [Modal, modalApi] = useVbenDrawer({
  // fullscreenButton: false,
  onCancel() {
    modalApi.close();
  },
  onConfirm: async () => {
    const validationResult = await formApi.validate();
    if (validationResult.valid) {
      await formApi.submitForm();
      modalApi.close();
    }
  },
  onOpenChange(isOpen: boolean) {
    isUpdate.value = modalApi.getData()?.isUpdate;
    record.value = isOpen ? modalApi.getData()?.record || {} : {};
    gridApi.value = isOpen ? modalApi.getData()?.gridApi : null;
    if (isOpen) {
      formApi.setValues(record.value);
    }
    modalApi.setState({
      title: isUpdate.value
        ? $t("upgrade.upgradeLnxUpgradeStrategy.editUpgradeLnxUpgradeStrategy")
        : $t("upgrade.upgradeLnxUpgradeStrategy.addUpgradeLnxUpgradeStrategy"),
    });
  },
});

defineExpose(modalApi);
</script>
<template>
  <Modal>
    <Form />
  </Modal>
</template>
