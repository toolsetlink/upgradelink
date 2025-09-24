<script lang="ts" setup>
import type { AnalysisOverviewItem } from "@vben/common-ui";
import type { TabOption } from "@vben/types";

import { onMounted, ref } from "vue";

import { AnalysisChartsTabs, AnalysisOverview } from "@vben/common-ui";

import { getUpgradeDashboard } from "#/api/upgrade/upgradeDashboard";

// import Analytics7DayDevCount from "./analytics-7day-dev-count.vue";
// import Analytics7dayReqCount from "./analytics-7day-req-count.vue";
import Analytics7dayAppGetStrategyCount from "./analytics-7day-appgetstrategy-count.vue";
import Analytics7dayAppStartCount from "./analytics-7day-appstart-count.vue";
import Analytics7dayAppUpgradeCount from "./analytics-7day-appupgrade-count.vue";
import Analytics7DayDownloadCount from "./analytics-7day-download-count.vue";
import Analytics7DayTrafficUsageCount from "./analytics-7day-trafficusage-count.vue";

// 新增状态变量
const dashboardData = ref<DashboardData | null>(null);

// 正确使用 ref 创建响应式对象
const overviewItems = ref<AnalysisOverviewItem[]>([
  // {
  //   icon: "i-fluent:device-phone-24-filled",
  //   title: "昨日新增设备量",
  //   totalTitle: "总用户量",
  //   totalValue: 0,
  //   value: 0,
  // },
  // {
  //   icon: "i-fluent:device-phone-24-filled",
  //   title: "昨日请求次数",
  //   totalTitle: "总请求次数",
  //   totalValue: 0,
  //   value: 0,
  // },
  {
    icon: "i-fluent:device-phone-24-filled",
    title: "昨日下载量",
    totalTitle: "总下载量",
    totalValue: 0,
    value: 0,
  },
  {
    icon: "i-fluent:device-phone-24-filled",
    title: "昨日请求升级量",
    totalTitle: "总应用请求升级量",
    totalValue: 0,
    value: 0,
  },
  {
    icon: "i-fluent:device-phone-24-filled",
    title: "昨日升级量",
    totalTitle: "总应用升级量",
    totalValue: 0,
    value: 0,
  },
  {
    icon: "i-fluent:device-phone-24-filled",
    title: "昨日启动量",
    totalTitle: "总应用启动量",
    totalValue: 0,
    value: 0,
  },
]);

onMounted(async () => {
  try {
    const response = await getUpgradeDashboard();
    // 统一存储接口数据
    dashboardData.value = response.data;

    // 更新overview数据
    overviewItems.value = [
      {
        ...overviewItems.value[0],
        value: response.data.yesterdayDownloadCount,
        totalValue: response.data.totalDownloadCount,
      },
      {
        ...overviewItems.value[1],
        value: response.data.yesterdayAppGetStrategyCount,
        totalValue: response.data.totalAppGetStrategyCount,
      },
      {
        ...overviewItems.value[2],
        value: response.data.yesterdayAppUpgradeCount,
        totalValue: response.data.totalAppUpgradeCount,
      },
      {
        ...overviewItems.value[3],
        value: response.data.yesterdayAppStartCount,
        totalValue: response.data.totalAppStartCount,
      },
    ];
  } catch (error) {
    console.error("获取数据失败:", error);
  }
});

// const chartTabs1: TabOption[] = [
//   {
//     label: "近7天设备新增量",
//     value: "7day-dev-count",
//   },
// ];
//
// const chartTabs2: TabOption[] = [
//   {
//     label: "近7天接口请求次数",
//     value: "7day-req-count",
//   },
// ];

const chartTabs5: TabOption[] = [
  {
    label: "近7天流量使用量(G)",
    value: "7day-trafficusage-count",
  },
];

const chartTabs1: TabOption[] = [
  {
    label: "近7天下载次数",
    value: "7day-download-count",
  },
];

const chartTabs4: TabOption[] = [
  {
    label: "近7天应用请求升级次数",
    value: "7day-appgetstrategy-count",
  },
];

const chartTabs3: TabOption[] = [
  {
    label: "近7天应用升级次数",
    value: "7day-appupgrade-count",
  },
];

const chartTabs2: TabOption[] = [
  {
    label: "近7天应用启动次数",
    value: "7day-appstart-count",
  },
];

// const chartTabs: TabOption[] = [
//   {
//     label: "设备版本分布",
//     value: "trends",
//   },
//   {
//     label: "设备机型分布",
//     value: "visits",
//   },
// ];
</script>

<template>
  <div class="p-5">
    <AnalysisOverview :items="overviewItems" />

    <!--    <AnalysisChartsTabs :tabs="chartTabs1" class="mt-5">-->
    <!--      <template #7day-dev-count>-->
    <!--        &lt;!&ndash; 传递数据给子组件 &ndash;&gt;-->
    <!--        <Analytics7DayDevCount :chart-data="dashboardData?.devCount7Day" />-->
    <!--      </template>-->
    <!--    </AnalysisChartsTabs>-->

    <!--    <AnalysisChartsTabs :tabs="chartTabs2" class="mt-5">-->
    <!--      <template #7day-req-count>-->
    <!--        <Analytics7dayReqCount :chart-data="dashboardData?.reqCount7Day" />-->
    <!--      </template>-->
    <!--    </AnalysisChartsTabs>-->

    <AnalysisChartsTabs :tabs="chartTabs5" class="mt-5">
      <template #7day-trafficusage-count>
        <Analytics7DayTrafficUsageCount
          :chart-data="dashboardData?.trafficUsageCount7Day"
        />
      </template>
    </AnalysisChartsTabs>

    <AnalysisChartsTabs :tabs="chartTabs1" class="mt-5">
      <template #7day-download-count>
        <Analytics7DayDownloadCount
          :chart-data="dashboardData?.downloadCount7Day"
        />
      </template>
    </AnalysisChartsTabs>

    <AnalysisChartsTabs :tabs="chartTabs4" class="mt-5">
      <template #7day-appgetstrategy-count>
        <Analytics7dayAppGetStrategyCount
          :chart-data="dashboardData?.appGetStrategyCount7Day"
        />
      </template>
    </AnalysisChartsTabs>

    <AnalysisChartsTabs :tabs="chartTabs3" class="mt-5">
      <template #7day-appupgrade-count>
        <Analytics7dayAppUpgradeCount
          :chart-data="dashboardData?.appUpgradeCount7Day"
        />
      </template>
    </AnalysisChartsTabs>

    <AnalysisChartsTabs :tabs="chartTabs2" class="mt-5">
      <template #7day-appstart-count>
        <Analytics7dayAppStartCount
          :chart-data="dashboardData?.appStartCount7Day"
        />
      </template>
    </AnalysisChartsTabs>

    <!--    <AnalysisChartsTabs :tabs="chartTabs" class="mt-5">-->
    <!--      <template #trends>-->
    <!--        <AnalyticsTrends />-->
    <!--      </template>-->
    <!--      <template #visits>-->
    <!--        <AnalyticsVisits />-->
    <!--      </template>-->
    <!--    </AnalysisChartsTabs>-->

    <!--    <div class="mt-5 w-full md:flex">-->
    <!--      <AnalysisChartCard class="mt-5 md:mr-4 md:mt-0 md:w-1/3" title="访问数量">-->
    <!--        <AnalyticsVisitsData />-->
    <!--      </AnalysisChartCard>-->
    <!--      <AnalysisChartCard class="mt-5 md:mr-4 md:mt-0 md:w-1/3" title="访问来源">-->
    <!--        <AnalyticsVisitsSource />-->
    <!--      </AnalysisChartCard>-->
    <!--      <AnalysisChartCard class="mt-5 md:mt-0 md:w-1/3" title="访问来源">-->
    <!--        <AnalyticsVisitsSales />-->
    <!--      </AnalysisChartCard>-->
    <!--    </div>-->
  </div>
</template>
