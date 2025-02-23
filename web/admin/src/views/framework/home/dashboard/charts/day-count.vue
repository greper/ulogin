<template>
  <v-chart v-if="props.data" class="chart" :option="option" autoresize />
</template>

<script setup lang="ts">
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { PieChart, LineChart } from "echarts/charts";
import { TitleComponent, TooltipComponent, LegendComponent, GridComponent } from "echarts/components";
import VChart, { THEME_KEY } from "vue-echarts";
import { ref, provide, defineProps, computed } from "vue";
import { ChartItem } from "/@/views/framework/home/dashboard/charts/d";

use([CanvasRenderer, PieChart, LineChart, TitleComponent, TooltipComponent, GridComponent, LegendComponent]);

provide(THEME_KEY, "");

const props = withDefaults(
  defineProps<{
    data: ChartItem[];
    title: string;
  }>(),
  {
    data: () => {
      return [];
    }
  }
);

const dates = computed(() => {
  return props.data.map((item) => {
    return item.name;
  });
});
const counts = computed(() => {
  return props.data.map((item) => {
    return item.value;
  });
});
const option = ref({
  color: ["#91cc75", "#73c0de", "#ee6666", "#fac858", "#3ba272", "#fc8452", "#9a60b4", "#ea7ccc", "#5470c6"],
  title: {
    show: props.data.length === 0, // 没数据才显示
    extStyle: {
      color: "grey",
      fontSize: 20
    },
    text: "暂无数据",
    left: "center",
    top: "center"
  },

  tooltip: {
    trigger: "item"
  },
  // tooltip: {
  //   trigger: "axis",
  //   axisPointer: {
  //     type: "cross",
  //     label: {
  //       backgroundColor: "#6a7985"
  //     }
  //   }
  // },
  // legend: {
  //   data: ["Email",]
  // },
  grid: {
    top: "20px",
    left: "20px",
    right: "20px",
    bottom: "10px",
    containLabel: true
  },
  xAxis: [
    {
      type: "category",
      boundaryGap: false,
      data: dates
    }
  ],
  yAxis: [
    {
      type: "value"
    }
  ],
  series: [
    {
      name: props.title,
      type: "line",
      stack: "Total",
      label: {
        show: true,
        position: "top"
      },
      smooth: true,
      areaStyle: {},
      emphasis: {
        focus: "series"
      },
      data: counts
    }
  ]
});
</script>

<style lang="less"></style>
