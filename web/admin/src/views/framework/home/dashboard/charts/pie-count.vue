<template>
  <v-chart class="chart" :option="option" autoresize />
</template>

<script setup lang="ts">
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { PieChart } from "echarts/charts";
import { TitleComponent, TooltipComponent, LegendComponent, GridComponent } from "echarts/components";
import VChart, { THEME_KEY } from "vue-echarts";
import { ref, provide, defineProps } from "vue";
import { ChartItem } from "./d";

use([CanvasRenderer, PieChart, TitleComponent, TooltipComponent, GridComponent, LegendComponent]);

provide(THEME_KEY, "");

const props = defineProps<{
  data: ChartItem[];
}>();

const option = ref({
  color: ["#91cc75", "#73c0de", "#ee6666", "#fac858", "#5470c6", "#3ba272", "#fc8452", "#9a60b4", "#ea7ccc", "#5470c6"],
  tooltip: {
    trigger: "item"
  },
  legend: {
    orient: "vertical",
    bottom: "5%",
    left: "left"
  },
  grid: {
    top: "20px",
    left: "20px",
    right: "20px",
    bottom: "10px",
    containLabel: true
  },
  series: [
    {
      center: ["60%", "50%"],
      name: "状态",
      type: "pie",
      radius: "80%",
      avoidLabelOverlap: false,
      itemStyle: {
        borderRadius: 0,
        borderColor: "#fff",
        borderWidth: 1
      },
      label: {
        show: false,
        position: "center"
      },
      emphasis: {
        label: {
          show: false,
          fontSize: 18,
          fontWeight: "bold"
        }
      },
      labelLine: {
        show: false
      },
      data: props.data
    }
  ]
});
</script>

<style lang="less">
.chart {
}
</style>
