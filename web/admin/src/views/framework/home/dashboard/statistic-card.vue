<template>
  <div class="statistic-card">
    <a-card>
      <div class="data-item">
        <div class="header">
          <div class="title">{{ title }}</div>
          <div class="more"></div>
        </div>
        <div class="content">
          <div v-if="!slots.default" class="statistic">
            <div v-if="count !== 0" class="value">{{ count }}</div>
            <a-empty v-else></a-empty>
          </div>
          <slot></slot>
        </div>
        <div v-if="slots.footer" class="footer">
          <slot name="footer"></slot>
        </div>
      </div>
    </a-card>
  </div>
</template>
<script setup lang="ts">
import { FsIcon } from "@fast-crud/fast-crud";
const props = defineProps<{
  title: string;
  count?: number;
}>();
const slots = defineSlots();
</script>
<style lang="less">
.statistic-card {
  margin-bottom: 10px;
  .icon-text {
    display: inline-flex;
    justify-content: left;
    align-items: center;
    .fs-icon {
      margin-right: 5px;
      font-size: 14px;
    }
  }
  .data-item {
    display: flex;
    flex-direction: column;
    height: 180px;
    .header {
      display: flex;
      justify-content: space-between;
      //padding-bottom: 10px;
      color: #8077a4;
    }
    .content {
      display: flex;
      flex-direction: column;
      flex: 1;
      .statistic {
        height: 100%;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        .value {
          font-size: 50px;
          font-weight: 700;
          color: #2c254e;
        }
      }

      x-vue-echarts {
      }
    }
    .footer {
      color: #8077a4;
      font-size: 12px;
      text-align: right;
      display: flex;
      justify-content: right;
      align-items: center;
      border-color: #e8e8e8;
      border-style: dashed;
      border-width: 1px 0 0;
      padding-top: 15px;
      > * {
        cursor: pointer;
      }
      margin-bottom: -10px;
    }
  }
}
</style>
