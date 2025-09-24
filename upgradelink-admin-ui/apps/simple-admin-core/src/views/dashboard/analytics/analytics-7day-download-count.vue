<script lang="ts" setup>
import type { EchartsUIType } from "@vben/plugins/echarts";

import { ref, watch } from "vue";

import { EchartsUI, useEcharts } from "@vben/plugins/echarts";

// 定义props
const props = defineProps<{
  chartData?: {
    series: Array<{
      appName: string;
      data: number[];
    }>;
    timeData: string[];
  };
}>();

const chartRef = ref<EchartsUIType>();
const { renderEcharts } = useEcharts(chartRef);

watch(
  () => props.chartData,
  (newVal) => {
    if (newVal) {
      renderEcharts({
        grid: {
          bottom: 0,
          containLabel: true,
          left: "1%",
          right: "1%",
          top: "2 %",
        },
        series: newVal.series.map((item) => ({
          name: item.appName,
          areaStyle: {},
          data: item.data,
          itemStyle: {},
          smooth: true,
          type: "line",
        })),
        tooltip: {
          axisPointer: {
            lineStyle: {
              width: 1,
            },
          },
          trigger: "axis",
        },
        xAxis: {
          axisTick: {
            show: false,
          },
          boundaryGap: false,
          data: newVal.timeData,
          splitLine: {
            lineStyle: {
              type: "solid",
              width: 1,
            },
            show: true,
          },
          type: "category",
        },
        yAxis: [
          {
            axisTick: {
              show: false,
            },
            splitArea: {
              show: true,
            },
            splitNumber: 4,
            type: "value",
          },
        ],
      });
    }
  },
  { immediate: true },
);
</script>

<template>
  <EchartsUI ref="chartRef" />
</template>
