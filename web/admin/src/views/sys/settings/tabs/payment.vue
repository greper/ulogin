<template>
  <div class="sys-settings-form sys-settings-payment">
    <a-form ref="formRef" :model="formState" :label-col="{ span: 8 }" :wrapper-col="{ span: 16 }" autocomplete="off">
      <div>支付方式</div>
      <a-form-item label="彩虹易支付" :name="['yizhifu', 'enabled']" :required="true">
        <a-switch v-model:checked="formState.yizhifu.enabled" />
      </a-form-item>
      <a-form-item v-if="formState.yizhifu.enabled" label="易支付配置" :name="['yizhifu', 'accessId']" :required="true">
        <access-selector v-model="formState.yizhifu.accessId" type="yizhifu" from="sys" />
        <div class="helper">
          <a href="https://certd.docmirror.cn/comm/payments/yizhifu.html">彩虹易支付配置帮助文档</a>
        </div>
      </a-form-item>

      <a-form-item label="支付宝" :name="['alipay', 'enabled']" :required="true">
        <a-switch v-model:checked="formState.alipay.enabled" />
      </a-form-item>
      <a-form-item v-if="formState.alipay.enabled" label="支付宝配置" :name="['alipay', 'accessId']" :required="true">
        <access-selector v-model="formState.alipay.accessId" type="alipay" from="sys" />
        <div class="helper">需要开通电脑网站支付， <a href="https://certd.docmirror.cn/comm/payments/alipay.html">支付宝配置帮助文档</a></div>
      </a-form-item>

      <a-form-item label="微信支付" :name="['wxpay', 'enabled']" :required="true">
        <a-switch v-model:checked="formState.wxpay.enabled" />
      </a-form-item>
      <a-form-item v-if="formState.wxpay.enabled" label="微信支付配置" :name="['wxpay', 'accessId']" :required="true">
        <access-selector v-model="formState.wxpay.accessId" type="wxpay" from="sys" />
        <div class="helper">需要开通Native支付， <a href="https://certd.docmirror.cn/comm/payments/wxpay.html">微信配置帮助文档</a></div>
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
        <loading-button type="primary" html-type="button" :click="onClick">保存</loading-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="tsx">
import { reactive, ref } from "vue";
import { merge } from "lodash-es";
import { notification } from "ant-design-vue";
import { request } from "/@/api/service";

defineOptions({
  name: "SettingPayment"
});

const api = {
  async SettingGet() {
    return await request({
      url: "/sys/settings/payment/get",
      method: "post"
    });
  },
  async SettingSave(data: any) {
    return await request({
      url: "/sys/settings/payment/save",
      method: "post",
      data
    });
  }
};

const formRef = ref<any>(null);
type PaymentItem = {
  enabled: boolean;
  accessId?: number;
};

const formState = reactive<
  Partial<{
    yizhifu: PaymentItem;
    alipay: PaymentItem;
    wxpay: PaymentItem;
  }>
>({
  yizhifu: { enabled: false },
  alipay: { enabled: false },
  wxpay: { enabled: false }
});

async function loadSettings() {
  const data: any = await api.SettingGet();
  merge(formState, data);
}

loadSettings();
const onClick = async () => {
  const form = await formRef.value.validateFields();
  await api.SettingSave(form);
  await loadSettings();
  notification.success({
    message: "保存成功"
  });
};
</script>
<style lang="less">
.sys-settings-base {
}
</style>
