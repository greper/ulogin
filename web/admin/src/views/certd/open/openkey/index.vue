<template>
  <fs-page>
    <template #header>
      <div class="title">开放接口密钥管理</div>
      <div class="more">
        <a :href="OPEN_API_DOC" target="_blank">开放接口文档</a>
      </div>
    </template>
    <fs-crud ref="crudRef" v-bind="crudBinding"> </fs-crud>
  </fs-page>
</template>

<script lang="ts" setup>
import { onActivated, onMounted } from "vue";
import { useFs } from "@fast-crud/fast-crud";
import createCrudOptions from "./crud";
import { OPEN_API_DOC } from "/@/views/certd/open/openkey/api";

defineOptions({
  name: "OpenKey"
});
const { crudBinding, crudRef, crudExpose } = useFs({ createCrudOptions, context: {} });

// 页面打开后获取列表数据
onMounted(() => {
  crudExpose.doRefresh();
});
onActivated(() => {
  crudExpose.doRefresh();
});
</script>
