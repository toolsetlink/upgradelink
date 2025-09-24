<script lang="ts" setup>
import type { UpgradeCompanyTrafficPacketInfo } from "#/api/upgrade/model/upgradeCompanyTrafficPacketModel";

import { ref } from "vue";

import { useVbenModal } from "@vben/common-ui";
import { $t } from "@vben/locales";

import { message } from "ant-design-vue";

import { useVbenForm } from "#/adapter/form";
import {
  createUpgradeCompanyTrafficPacket,
  updateUpgradeCompanyTrafficPacket,
} from "#/api/upgrade/upgradeCompanyTrafficPacket";

import { dataFormSchemas } from "./schemas";

defineOptions({
  name: "UpgradeCompanyTrafficPacketForm",
});

const record = ref();
const isUpdate = ref(false);
const gridApi = ref();

async function onSubmit(values: Record<string, any>) {
  const result = isUpdate.value
    ? await updateUpgradeCompanyTrafficPacket(
        values as UpgradeCompanyTrafficPacketInfo,
      )
    : await createUpgradeCompanyTrafficPacket(
        values as UpgradeCompanyTrafficPacketInfo,
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
            "upgrade.upgradeCompanyTrafficPacket.editUpgradeCompanyTrafficPacket",
          )
        : $t(
            "upgrade.upgradeCompanyTrafficPacket.addUpgradeCompanyTrafficPacket",
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
