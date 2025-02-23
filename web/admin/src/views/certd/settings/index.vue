<template>
  <fs-page class="page-user-settings">
    <template #header>
      <div class="title">设置</div>
    </template>
    <div class="user-settings-form settings-form">
      <a-form
        :model="formState"
        name="basic"
        :label-col="{ span: 8 }"
        :wrapper-col="{ span: 16 }"
        autocomplete="off"
        @finish="onFinish"
        @finish-failed="onFinishFailed"
      >
        <a-form-item label="默认定时设置" name="defaultCron">
          <notification-selector v-model="formState.defaultCron" />
          <div class="helper">创建流水线时默认使用此定时时间</div>
        </a-form-item>

        <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
          <a-button :loading="saveLoading" type="primary" html-type="submit">保存</a-button>
        </a-form-item>
      </a-form>
    </div>
  </fs-page>
</template>

<script setup lang="tsx">
import { reactive, ref } from "vue";
import * as api from "./api";
import { UserSettings } from "./api";
import { notification } from "ant-design-vue";
import { merge } from "lodash-es";
import NotificationSelector from "/@/views/certd/notification/notification-selector/index.vue";

defineOptions({
  name: "UserSettings"
});

const formState = reactive<Partial<UserSettings>>({});

async function loadUserSettings() {
  const data: any = await api.UserSettingsGet();
  merge(formState, data);
}

const saveLoading = ref(false);
loadUserSettings();
const onFinish = async (form: any) => {
  try {
    saveLoading.value = true;
    await api.UserSettingsSave(form);
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
</script>

<style lang="less">
.page-user-settings {
  .user-settings-form {
    width: 500px;
    margin: 20px;
  }
}
</style>
