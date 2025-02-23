<template>
  <div class="sys-settings-form sys-settings-base">
    <a-form
      :model="formState"
      name="basic"
      :label-col="{ span: 8 }"
      :wrapper-col="{ span: 16 }"
      autocomplete="off"
      @finish="onFinish"
      @finish-failed="onFinishFailed"
    >
      <a-form-item label="ICP备案号" :name="['public', 'icpNo']">
        <a-input v-model:value="formState.public.icpNo" placeholder="粤ICP备xxxxxxx号" />
      </a-form-item>

      <a-form-item label="HTTP代理" :name="['private', 'httpProxy']" :rules="urlRules">
        <a-input v-model:value="formState.private.httpProxy" placeholder="http://192.168.1.2:18010/" />
        <div class="helper">当某些网站被墙时可以配置</div>
      </a-form-item>

      <a-form-item label="HTTPS代理" :name="['private', 'httpsProxy']" :rules="urlRules">
        <div class="flex">
          <a-input v-model:value="formState.private.httpsProxy" placeholder="http://192.168.1.2:18010/" />
          <a-button class="ml-5" type="primary" :loading="testProxyLoading" title="保存后，再点击测试" @click="testProxy">测试</a-button>
        </div>
        <div class="helper">一般这两个代理填一样的，保存后再测试</div>
      </a-form-item>

      <a-form-item label="双栈网络" :name="['private', 'dnsResultOrder']">
        <a-select v-model:value="formState.private.dnsResultOrder">
          <a-select-option value="verbatim">默认</a-select-option>
          <a-select-option value="ipv4first">IPV4优先</a-select-option>
          <a-select-option value="ipv6first">IPV6优先</a-select-option>
        </a-select>
        <div class="helper">如果选择IPv6优先，需要在docker-compose.yaml中启用ipv6</div>
      </a-form-item>

      <a-form-item label="启用公共CNAME服务" :name="['private', 'commonCnameEnabled']">
        <a-switch v-model:checked="formState.private.commonCnameEnabled" />
        <div class="helper">
          是否可以使用公共CNAME服务，如果禁用，且没有设置<router-link to="/sys/cname/provider">自定义CNAME服务</router-link>，则无法使用CNAME代理方式申请证书
        </div>
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
        <a-button :loading="saveLoading" type="primary" html-type="submit">保存</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="tsx">
import { reactive, ref } from "vue";
import { SysSettings } from "/@/views/sys/settings/api";
import * as api from "/@/views/sys/settings/api";
import { merge } from "lodash-es";
import { useSettingStore } from "/@/store/modules/settings";
import { notification } from "ant-design-vue";
import { util } from "/@/utils";

defineOptions({
  name: "SettingBase"
});

const formState = reactive<Partial<SysSettings>>({
  public: {
    icpNo: ""
  },
  private: {}
});

const urlRules = ref({
  type: "url",
  message: "请输入正确的URL"
});

async function loadSysSettings() {
  const data: any = await api.SysSettingsGet();
  merge(formState, data);
}

const saveLoading = ref(false);
loadSysSettings();
const settingsStore = useSettingStore();
const onFinish = async (form: any) => {
  try {
    saveLoading.value = true;
    await api.SysSettingsSave(form);
    await settingsStore.loadSysSettings();
    notification.success({
      message: "保存成功"
    });
  } finally {
    saveLoading.value = false;
  }
};

const onFinishFailed = (errorInfo: any) => {
  // console.log("Failed:", errorInfo);
};

async function stopOtherUserTimer() {
  await api.stopOtherUserTimer();
  notification.success({
    message: "停止成功"
  });
}

const testProxyLoading = ref(false);
async function testProxy() {
  testProxyLoading.value = true;
  try {
    const res = await api.TestProxy();
    let success = true;
    if (res.google !== true || res.baidu !== true) {
      success = false;
    }
    const content = () => {
      return (
        <div>
          <div>Google: {res.google === true ? "成功" : util.maxLength(res.google)}</div>
          <div>Baidu: {res.baidu === true ? "成功" : util.maxLength(res.google)}</div>
        </div>
      );
    };
    if (!success) {
      notification.error({
        message: "测试失败",
        description: content
      });
      return;
    }
    notification.success({
      message: "测试完成",
      description: content
    });
  } finally {
    testProxyLoading.value = false;
  }
}
</script>
<style lang="less">
.sys-settings-base {
}
</style>
