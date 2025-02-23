<template>
  <a-button class="cd-secret-plain-getter ml-5" @click="showPlain">
    <fs-icon class="pointer" :icon="computedIcon" />
  </a-button>
</template>
<script lang="ts" setup>
import { computed, inject, ref } from "vue";
defineOptions({
  name: "SecretPlainGetter"
});

const props = defineProps<{
  modelValue?: string;
  accessId?: number;
  inputKey: string;
}>();

const emit = defineEmits(["update:modelValue"]);
const showRef = ref(false);
const computedIcon = computed(() => {
  return showRef.value ? "ion:eye-outline" : "ion:eye-off-outline";
});
const accessApi: any = inject("accessApi");
async function showPlain() {
  showRef.value = true;
  if (props.accessId) {
    const plain = await accessApi.GetSecretPlain(props.accessId, props.inputKey);
    emit("update:modelValue", plain);
  }
}
</script>
<style lang="less"></style>
