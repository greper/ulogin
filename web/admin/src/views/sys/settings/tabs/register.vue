<template>
  <div class="sys-settings-form sys-settings-register">
    <a-form :model="formState" name="register" :label-col="{ span: 8 }" :wrapper-col="{ span: 16 }" autocomplete="off" @finish="onFinish">
      <a-form-item label="开启自助注册" :name="['public', 'registerEnabled']">
        <a-switch v-model:checked="formState.public.registerEnabled" />
      </a-form-item>
      <template v-if="formState.public.registerEnabled">
        <a-form-item label="限制用户流水线数量" :name="['public', 'limitUserPipelineCount']">
          <a-input-number v-model:value="formState.public.limitUserPipelineCount" />
          <div class="helper">0为不限制</div>
        </a-form-item>
        <a-form-item label="管理其他用户流水线" :name="['public', 'managerOtherUserPipeline']">
          <a-switch v-model:checked="formState.public.managerOtherUserPipeline" />
        </a-form-item>

        <a-form-item label="开启用户名注册" :name="['public', 'usernameRegisterEnabled']">
          <a-switch v-model:checked="formState.public.usernameRegisterEnabled" />
        </a-form-item>

        <a-form-item label="开启邮箱注册" :name="['public', 'emailRegisterEnabled']">
          <div class="flex-o">
            <a-switch v-model:checked="formState.public.emailRegisterEnabled" :disabled="!settingsStore.isPlus" title="专业版功能" />
            <vip-button class="ml-5" mode="button"></vip-button>
          </div>
          <div class="helper">需要<router-link to="/sys/settings/email">设置邮箱服务器</router-link></div>
        </a-form-item>
        <a-form-item label="开启手机号登录、注册" :name="['public', 'smsLoginEnabled']">
          <div class="flex-o">
            <a-switch v-model:checked="formState.public.smsLoginEnabled" :disabled="!settingsStore.isComm" title="商业版功能" />
            <vip-button class="ml-5" mode="comm"></vip-button>
          </div>
        </a-form-item>
        <template v-if="formState.public.smsLoginEnabled">
          <a-form-item label="短信提供商" :name="['private', 'sms', 'type']">
            <a-select v-model:value="formState.private.sms.type" @change="smsTypeChange">
              <a-select-option value="aliyun">阿里云短信</a-select-option>
              <a-select-option value="yfysms">易发云短信</a-select-option>
            </a-select>
          </a-form-item>
          <template v-for="item of smsTypeDefineInputs" :key="item.simpleKey">
            <fs-form-item v-model="formState.private.sms.config[item.simpleKey]" :path="'private.sms.config' + item.key" :item="item" />
          </template>

          <a-form-item label="短信测试">
            <div class="flex">
              <a-input v-model:value="testMobile" placeholder="输入测试手机号" />
              <loading-button class="ml-5" title="保存后再点击测试" type="primary" :click="testSendSms">测试</loading-button>
            </div>
            <div class="helper">保存后再点击测试</div>
          </a-form-item>
        </template>
      </template>

      <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
        <a-button :loading="saveLoading" type="primary" html-type="submit">保存</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="tsx">
import { reactive, ref, Ref } from "vue";
import { GetSmsTypeDefine, SysSettings } from "/@/views/sys/settings/api";
import * as api from "/@/views/sys/settings/api";
import { merge } from "lodash-es";
import { useSettingStore } from "/@/store/modules/settings";
import { notification } from "ant-design-vue";

defineOptions({
  name: "SettingRegister"
});

const testMobile = ref("");
async function testSendSms() {
  if (!testMobile.value) {
    notification.error({
      message: "请输入测试手机号"
    });
    return;
  }
  await api.TestSms({
    mobile: testMobile.value
  });
  notification.success({
    message: "发送成功"
  });
}
const formState = reactive<Partial<SysSettings>>({
  public: {
    registerEnabled: false
  },
  private: {
    sms: {
      type: "aliyun",
      config: {}
    }
  }
});

const rules = {
  leastOneLogin: {
    validator: (rule: any, value: any) => {
      if (!formState.public.passwordLoginEnabled && !formState.public.smsLoginEnabled) {
        return Promise.reject("密码登录和手机号登录至少开启一个");
      }
      return Promise.resolve();
    }
  },
  required: {
    required: true,
    trigger: "change",
    message: "此项必填"
  }
};

async function smsTypeChange(value: string) {
  if (formState.private?.sms?.config) {
    formState.private.sms.config = {};
  }

  await loadTypeDefine(value);
}
const smsTypeDefineInputs: Ref = ref({});
async function loadTypeDefine(type: string) {
  const define: any = await api.GetSmsTypeDefine(type);
  const keys = Object.keys(define.input);
  const inputs: any = {};
  keys.forEach((key) => {
    const value = define.input[key];
    value.simpleKey = key;
    value.key = "private.sms.config." + key;
    if (!value.component) {
      value.component = {
        name: "a-input"
      };
    }
    if (!value.component.name) {
      value.component.vModel = "value";
    }
    if (!value.rules) {
      value.rules = [];
    }
    if (value.required) {
      value.rules.push(rules.required);
    }

    inputs[key] = define.input[key];
  });
  smsTypeDefineInputs.value = inputs;
}

async function loadSysSettings() {
  const data: any = await api.SysSettingsGet();
  merge(formState, data);
  if (data?.private.sms?.type) {
    await loadTypeDefine(data.private.sms.type);
  }
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
</script>
<style lang="less">
.sys-settings-site {
}
</style>
