<template>
  <a-button v-if="showButton" type="primary" @click="open">修改密码</a-button>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { CrudOptions, useColumns, useFormWrapper } from "@fast-crud/fast-crud";
import * as api from "/@/views/certd/mine/api";
import { notification } from "ant-design-vue";
import { useUserStore } from "/@/store/modules/user";

defineProps<{
  showButton: boolean;
}>();

let passwordFormRef = ref();

const validatePass1 = async (rule: any, value: any) => {
  if (value === "") {
    throw new Error("请输入密码");
  }
  const formData = passwordFormRef.value.getFormData();
  if (formData.confirmNewPassword !== "") {
    passwordFormRef.value.formRef.formRef.validateFields(["confirmNewPassword"]);
  }
  if (formData.password === formData.newPassword) {
    throw new Error("新密码不能和旧密码相同");
  }
};
const validatePass2 = async (rule: any, value: any) => {
  if (value === "") {
    throw new Error("请再次输入密码");
  } else if (value !== passwordFormRef.value.getFormData().newPassword) {
    throw new Error("两次输入密码不一致!");
  }
};
const userStore = useUserStore();
const { openDialog } = useFormWrapper();
const { buildFormOptions } = useColumns();
const passwordFormOptions: CrudOptions = {
  form: {
    col: {
      span: 24
    },
    wrapper: {
      title: "修改密码",
      width: "500px"
    },
    async doSubmit({ form }) {
      await api.changePassword(form);
      //重新加载用户信息
      await userStore.loadUserInfo();
    },
    async afterSubmit() {
      notification.success({ message: "修改成功" });
    }
  },
  columns: {
    password: {
      title: "旧密码",
      type: "password",
      form: {
        rules: [{ required: true, message: "请输入旧密码" }]
      }
    },
    newPassword: {
      title: "新密码",
      type: "password",
      form: {
        rules: [
          { required: true, message: "请输入确认密码" },
          //@ts-ignore
          { validator: validatePass1, trigger: "blur" }
        ]
      }
    },
    confirmNewPassword: {
      title: "确认新密码",
      type: "password",
      form: {
        rules: [
          { required: true, message: "请输入确认密码" },
          //@ts-ignore
          { validator: validatePass2, trigger: "blur" }
        ]
      }
    }
  }
};

async function open(opts: { password: "" }) {
  const formOptions = buildFormOptions(passwordFormOptions);
  formOptions.newInstance = true; //新实例打开
  passwordFormRef.value = await openDialog(formOptions);
  passwordFormRef.value.setFormData({
    password: opts.password
  });
  console.log(passwordFormRef.value);
}

const scope = ref({
  open: open
});

defineExpose(scope.value);
</script>
