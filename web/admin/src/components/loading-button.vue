<template>
  <a-button :loading="loading" @click="onClick">
    <slot></slot>
  </a-button>
</template>

<script setup lang="ts">
import { ref } from "vue";
const props = defineProps<{
  click?: () => Promise<void>;
}>();

const loading = ref(false);
async function onClick() {
  loading.value = true;
  try {
    if (props.click) {
      await props.click();
    }
  } finally {
    loading.value = false;
  }
}
</script>
