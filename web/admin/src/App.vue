<template>
  <a-config-provider :locale="locale" :theme="settingStore.themeToken">
    <contextHolder />
    <fs-form-provider>
      <router-view v-if="routerEnabled" />
    </fs-form-provider>
  </a-config-provider>
</template>

<script lang="ts" setup>
import zhCN from "ant-design-vue/es/locale/zh_CN";
import enUS from "ant-design-vue/es/locale/en_US";
import { provide, ref } from "vue";
import { usePageStore } from "/src/store/modules/page";
import { useSettingStore } from "/@/store/modules/settings";
import "dayjs/locale/zh-cn";
import "dayjs/locale/en";
import dayjs from "dayjs";
import { Modal } from "ant-design-vue";

defineOptions({
  name: "App"
});
const [modal, contextHolder] = Modal.useModal();
provide("modal", modal);
//刷新页面方法
const routerEnabled = ref(true);
const locale = ref(zhCN);
async function reload() {
  // routerEnabled.value = false;
  // await nextTick();
  // routerEnabled.value = true;
}
function localeChanged(value: any) {
  console.log("locale changed:", value);
  if (value === "zh-cn") {
    locale.value = zhCN;
    dayjs.locale("zh-cn");
  } else if (value === "en") {
    locale.value = enUS;
    dayjs.locale("en");
  }
}
localeChanged("zh-cn");
provide("fn:router.reload", reload);
provide("fn:locale.changed", localeChanged);

//其他初始化
// const resourceStore = useResourceStore();
// resourceStore.init();
const pageStore = usePageStore();
pageStore.init();
const settingStore = useSettingStore();
</script>
