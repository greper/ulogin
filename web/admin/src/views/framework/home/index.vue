<template>
  <fs-page class="home—index">
    <!--    <page-content />-->
    <dashboard-user />
    <change-password-button ref="changePasswordButtonRef" :show-button="false"></change-password-button>
  </fs-page>
</template>

<script lang="ts" setup>
import DashboardUser from "./dashboard/index.vue";
import { useUserStore } from "/@/store/modules/user";
import ChangePasswordButton from "/@/views/certd/mine/change-password-button.vue";
import { onMounted, ref } from "vue";
import { Modal } from "ant-design-vue";

const userStore = useUserStore();
const changePasswordButtonRef = ref();
onMounted(() => {
  if (userStore.getUserInfo.isWeak === true) {
    Modal.info({
      title: "修改密码",
      content: "为了您的账户安全，请立即修改密码",
      onOk: () => {
        changePasswordButtonRef.value.open({
          password: "123456"
        });
      },
      okText: "立即修改"
    });
  }
});
</script>
<style lang="less">
.home—index {
  background-color: #eee;
}
</style>
