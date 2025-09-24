<script lang="ts" setup>
import { ref, watch } from "vue";

import { InputNumber, Select, TimePicker } from "ant-design-vue";
import dayjs from "dayjs";

const props = defineProps({
  timeMode: {
    type: String,
    default: "datetime",
  },
  valueFormat: {
    type: String,
    default: "unixmilli",
  },
  value: {
    type: Array<{
      beginTime: string;
      dimension: number;
      enable: number;
      endTime: string;
      limit: number;
    }>,
    default: () => [],
  },
});

const emit = defineEmits(["blur", "change"]);

const showTimePicker = ref<boolean>();
showTimePicker.value = props.timeMode === "datetime";

const modelValues = ref<
  Array<{
    beginTime: string;
    dimension: number;
    enable: number;
    endTime: string;
    limit: number;
  }>
>([]);

watch(
  () => props.value,
  () => {
    if (props.value && props.value.length > 0) {
      if (props.value.length > 0) {
        modelValues.value = props.value.map((item) => ({
          ...item,
          beginTime: dayjs(item.beginTime, "HH:mm:ss").format("HH:mm:ss"),
          endTime: dayjs(item.endTime, "HH:mm:ss").format("HH:mm:ss"),
        }));
      }
    } else {
      modelValues.value = [
        {
          enable: 1,
          beginTime: dayjs("00:00:00", "HH:mm:ss").format("HH:mm:ss"),
          endTime: dayjs("00:00:00", "HH:mm:ss").format("HH:mm:ss"),
          dimension: 1,
          limit: 10,
        },
      ];

      onChange();
    }
  },
  {
    immediate: true,
  },
);

function addInputGroup() {
  modelValues.value.push({
    enable: 1,
    beginTime: dayjs("00:00:00", "HH:mm:ss").format("HH:mm:ss"),
    endTime: dayjs("00:00:00", "HH:mm:ss").format("HH:mm:ss"),
    dimension: 1,
    limit: 10,
  });

  onChange();
}

function removeInputGroup(index: number) {
  modelValues.value.splice(index, 1);
  onChange();
}

function onChange() {
  emit("change", modelValues.value);
}

function handleTimePickerChange(
  date: dayjs.Dayjs | null,
  field: "beginTime" | "endTime",
  index: number,
) {
  if (date) {
    modelValues.value[index][field] = date.format("HH:mm:ss");
    onChange();
  }
}
</script>

<template>
  <div>
    <div
      v-for="(modelValue, index) in modelValues"
      :key="index"
      class="flex w-full gap-1 mb-2"
    >
      <Select
        v-model:value="modelValue.enable"
        class="w-[80px]"
        placeholder="启用状态"
        allow-clear
        :class="{ 'valid-success': !!modelValue.enable }"
        :options="[
          { label: '启用', value: 1 },
          { label: '禁用', value: 0 },
        ]"
        @blur="emit('blur')"
        @change="onChange"
      />
      <TimePicker
        v-bind="$attrs"
        :value="
          modelValue.beginTime ? dayjs(modelValue.beginTime, 'HH:mm:ss') : ''
        "
        format="HH:mm:ss"
        allow-clear
        @change="(date) => handleTimePickerChange(date, 'beginTime', index)"
      />
      <TimePicker
        v-bind="$attrs"
        :value="modelValue.endTime ? dayjs(modelValue.endTime, 'HH:mm:ss') : ''"
        format="HH:mm:ss"
        allow-clear
        @change="(date) => handleTimePickerChange(date, 'endTime', index)"
      />
      <Select
        v-model:value="modelValue.dimension"
        class="w-[80px]"
        placeholder="流控维护"
        allow-clear
        :class="{ 'valid-success': !!modelValue.dimension }"
        :options="[
          { label: '秒', value: 1 },
          { label: '分钟', value: 2 },
          { label: '小时', value: 3 },
          { label: '填', value: 4 },
        ]"
        @blur="emit('blur')"
        @change="onChange"
      />
      <InputNumber
        placeholder="请输入限制数量"
        class="flex-1"
        allow-clear
        :class="{ 'valid-success': modelValue.limit }"
        v-model:value="modelValue.limit"
        type="number"
        @blur="emit('blur')"
        @change="onChange"
      />
      <button
        v-if="index > 0"
        @click="removeInputGroup(index)"
        class="px-2 py-1 bg-red-500 text-white rounded"
      >
        删除
      </button>
      <button
        v-if="index === 0"
        @click="addInputGroup"
        class="px-2 py-1 bg-blue-500 text-white rounded"
      >
        添加
      </button>
    </div>
  </div>
</template>
