<template>
  <fs-page class="page-cert-notification-modal">
    <fs-crud ref="crudRef" v-bind="crudBinding"> </fs-crud>
  </fs-page>
</template>

<script lang="ts">
import { defineComponent, onMounted, watch } from "vue";
import { useFs } from "@fast-crud/fast-crud";
import createCrudOptions from "./crud";
import { createNotificationApi } from "../../api";

export default defineComponent({
  name: "CertNotificationModal",
  props: {
    modelValue: {}
  },
  emits: ["update:modelValue"],
  setup(props, ctx) {
    const api = createNotificationApi();
    const context: any = { props, ctx, api };
    const { crudBinding, crudRef, crudExpose } = useFs({ createCrudOptions, context });
    onMounted(() => {
      crudExpose.doRefresh();
    });

    return {
      crudBinding,
      crudRef
    };
  }
});
</script>
<style lang="less">
.page-cert-notification {
}
</style>
