<script lang="ts" setup>
import type { UpgradeConfigurationVersionInfo } from "#/api/upgrade/model/upgradeConfigurationVersionModel";

import { ref } from "vue";

import { useVbenModal } from "@vben/common-ui";
import { $t } from "@vben/locales";

import { message } from "ant-design-vue";

import { useVbenForm } from "#/adapter/form";
import {
  createUpgradeConfigurationVersion,
  updateUpgradeConfigurationVersion,
} from "#/api/upgrade/upgradeConfigurationVersion";

import { dataFormSchemas } from "./schemas";

defineOptions({
  name: "UpgradeConfigurationVersionForm",
});

const record = ref();
const isUpdate = ref(false);
const gridApi = ref();

async function onSubmit(values: Record<string, any>) {
  const result = isUpdate.value
    ? await updateUpgradeConfigurationVersion(
        values as UpgradeConfigurationVersionInfo,
      )
    : await createUpgradeConfigurationVersion(
        values as UpgradeConfigurationVersionInfo,
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

const [Modal, modalApi] = useVbenModal({
  fullscreenButton: false,
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
        ? $t(
            "upgrade.upgradeConfigurationVersion.editUpgradeConfigurationVersion",
          )
        : $t(
            "upgrade.upgradeConfigurationVersion.addUpgradeConfigurationVersion",
          ),
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
