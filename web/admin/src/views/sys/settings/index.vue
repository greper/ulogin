<template>
  <fs-page class="page-sys-settings">
    <!--    <template #header>-->
    <!--      <div class="title">系统设置</div>-->
    <!--    </template>-->
    <div class="sys-settings-body">
      <a-tabs :active-key="activeKey" type="card" class="sys-settings-tabs" @update:active-key="onChange">
        <a-tab-pane key="" tab="基本设置">
          <SettingBase v-if="activeKey === ''" />
        </a-tab-pane>
        <a-tab-pane key="register" tab="注册设置">
          <SettingRegister v-if="activeKey === 'register'" />
        </a-tab-pane>
        <a-tab-pane v-if="settingsStore.isComm" key="payment" tab="支付设置">
          <SettingPayment v-if="activeKey === 'payment'" />
        </a-tab-pane>
      </a-tabs>
    </div>
  </fs-page>
</template>

<script setup lang="tsx">
import SettingBase from "/@/views/sys/settings/tabs/base.vue";
import SettingRegister from "/@/views/sys/settings/tabs/register.vue";
import SettingPayment from "/@/views/sys/settings/tabs/payment.vue";
import { useRoute, useRouter } from "vue-router";
import { ref } from "vue";
import { useSettingStore } from "/@/store/modules/settings";
defineOptions({
  name: "SysSettings"
});
const settingsStore = useSettingStore();
const activeKey = ref("");
const route = useRoute();
const router = useRouter();
if (route.query.tab) {
  activeKey.value = (route.query.tab as string) || "";
}

function onChange(value: string) {
  // activeKey.value = value;
  // 创建一个新的查询参数对象
  const query: any = {};
  if (value !== "") {
    query.tab = value;
  }
  // 使用`push`方法更新路由，保留其他查询参数不变
  router.push({ path: route.path, query });
}
</script>

<style lang="less">
.page-sys-settings {
  .sys-settings-form {
    width: 500px;
    margin: 20px;
  }

  .sys-settings-body {
    height: 100%;
    padding-top: 20px;
    padding-left: 20px;
    .sys-settings-tabs {
      height: 100%;
      display: flex;
      flex-direction: column;
      .ant-tabs-content-holder {
        flex: 1;
        overflow: auto;
      }
    }
  }
}
</style>
