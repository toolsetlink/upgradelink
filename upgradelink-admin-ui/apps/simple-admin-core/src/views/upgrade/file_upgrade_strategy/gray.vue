<script lang="ts" setup>
import { ref, watch } from "vue";

import { InputNumber, RangePicker, Select } from "ant-design-vue";
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
      beginDatetime: string;
      enable: number;
      endDatetime: string;
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
    beginDatetime: string;
    enable: number;
    endDatetime: string;
    limit: number;
  }>
>([]);

watch(
  () => props.value,
  () => {
    // 生成当前日期的格式化字符串 (YYYY-MM-DD 00:00:00)
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, "0");
    const day = String(now.getDate()).padStart(2, "0");
    const currentDateStr = `${year}-${month}-${day} 00:00:00`;

    if (props.value && props.value.length > 0) {
      if (props.value.length > 0) {
        modelValues.value = props.value.map((item) => ({
          ...item,
          beginDatetime: item.beginDatetime || currentDateStr,
          endDatetime: item.endDatetime || currentDateStr,
        }));
      }
    } else {
      modelValues.value = [
        {
          enable: 1,
          beginDatetime: currentDateStr,
          endDatetime: currentDateStr,
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
  // 生成当前日期的格式化字符串 (YYYY-MM-DD 00:00:00)
  const now = new Date();
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, "0");
  const day = String(now.getDate()).padStart(2, "0");
  const currentDateStr = `${year}-${month}-${day} 00:00:00`;

  modelValues.value.push({
    enable: 1,
    beginDatetime: currentDateStr,
    endDatetime: currentDateStr,
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

function handleRangePickerChange(
  date: [dayjs.Dayjs, dayjs.Dayjs] | [string, string],
  index: number,
) {
  let dayjsDates: [dayjs.Dayjs, dayjs.Dayjs] | [string, string];
  // eslint-disable-next-line prefer-const
  dayjsDates =
    typeof date[0] === "string" ? [dayjs(date[0]), dayjs(date[1])] : date;

  if (dayjsDates && dayjsDates.length === 2) {
    modelValues.value[index].beginDatetime = dayjsDates[0].format(
      "YYYY-MM-DD HH:mm:ss",
    );
    modelValues.value[index].endDatetime = dayjsDates[1].format(
      "YYYY-MM-DD HH:mm:ss",
    );
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
      <RangePicker
        v-bind="$attrs"
        :value="[
          modelValue.beginDatetime ? dayjs(modelValue.beginDatetime) : '',
          modelValue.endDatetime ? dayjs(modelValue.endDatetime) : '',
        ]"
        :show-time="showTimePicker"
        allow-clear
        @change="(date) => handleRangePickerChange(date, index)"
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
