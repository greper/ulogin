<template>
  <fs-page class="page-user-profile">
    <template #header>
      <div class="title">我的信息</div>
    </template>
    <div class="p-10">
      <a-descriptions title="" bordered>
        <a-descriptions-item label="用户名">{{ userInfo.username }}</a-descriptions-item>
        <a-descriptions-item label="头像">
          <a-avatar v-if="userInfo.avatar" size="large" :src="'/api/basic/file/download?&key=' + userInfo.avatar" style="background-color: #eee"> </a-avatar>
          <a-avatar v-else size="large" style="background-color: #00b4f5">
            {{ userInfo.username }}
          </a-avatar>
        </a-descriptions-item>
        <a-descriptions-item label="昵称">{{ userInfo.nickName }}</a-descriptions-item>
        <a-descriptions-item label="邮箱">{{ userInfo.email }}</a-descriptions-item>
        <a-descriptions-item label="手机号">{{ userInfo.phoneCode }}{{ userInfo.mobile }}</a-descriptions-item>
        <a-descriptions-item label="修改密码">
          <change-password-button :show-button="true"> </change-password-button>
        </a-descriptions-item>
      </a-descriptions>
    </div>
  </fs-page>
</template>

<script lang="ts" setup>
import * as api from "./api";
import { Ref, ref } from "vue";
import ChangePasswordButton from "/@/views/certd/mine/change-password-button.vue";

defineOptions({
  name: "UserProfile"
});

const userInfo: Ref = ref({});

const getUserInfo = async () => {
  userInfo.value = await api.getMineInfo();
};
getUserInfo();
</script>
