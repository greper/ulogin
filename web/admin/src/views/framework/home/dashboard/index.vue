<template>
  <div class="dashboard-user">
    <div class="header-profile">
      <div class="avatar">
        <a-avatar v-if="userInfo.avatar" size="large" :src="'/api/basic/file/download?&key=' + userInfo.avatar" style="background-color: #eee"> </a-avatar>
        <a-avatar v-else size="large" style="background-color: #00b4f5">
          {{ userInfo.username }}
        </a-avatar>
      </div>
      <div class="text">
        <div class="left">
          <div>
            <span>您好，{{ userInfo.nickName || userInfo.username }}， 欢迎使用 【{{ siteInfo.title }}】</span>
          </div>
          <div class="flex-o">
            <a-tag color="green" class="flex-inline pointer m-0"> <fs-icon icon="ion:time-outline"></fs-icon> {{ now }}</a-tag>
            <template v-if="userStore.isAdmin">
              <a-divider type="vertical" />
              <a-badge :dot="hasNewVersion">
                <a-tag color="blue" class="flex-inline pointer m-0" :title="'最新版本:' + latestVersion" @click="openUpgradeUrl()">
                  <fs-icon icon="ion:rocket-outline" class="mr-5"></fs-icon>
                  v{{ version }}
                </a-tag>
              </a-badge>
            </template>
            <template v-if="settingsStore.isComm">
              <a-divider type="vertical" />
            </template>
          </div>
        </div>
      </div>
      <div class="suggest">
      </div>
    </div>
    <div v-if="!settingStore.isComm" class="warning">
      <a-alert type="warning" show-icon>
        <template #message>
          证书和授权为敏感信息，不要使用来历不明的在线Certd服务和镜像，以免泄露；请务必私有化部署使用，认准官方版本发布渠道：
          <a class="ml-5 flex-inline" href="https://gitee.com/certd/certd" target="_blank">gitee</a>、
          <a class="ml-5 flex-inline" href="https://github.com/certd/certd" target="_blank">github</a>、
          <a class="ml-5 flex-inline" href="https://certd.docmirror.cn" target="_blank">帮助文档</a>
        </template>
      </a-alert>
    </div>

    <div class="statistic-data m-20">
      <a-row :gutter="20">
        <a-col :span="6">
          <statistic-card title="证书流水线数量" :count="count.pipelineCount">
            <template v-if="count.pipelineCount === 0" #default>
              <div class="flex-center flex-1 flex-col">
                <div style="font-size: 18px; font-weight: 700">您还没有证书流水线</div>
                <fs-button type="primary" class="mt-10" icon="ion:add-circle-outline" @click="goPipeline">立即创建</fs-button>
              </div>
            </template>
            <template #footer>
              <router-link to="/certd/pipeline" class="flex"><fs-icon icon="ion:settings-outline" class="mr-5 fs-16" /> 管理流水线</router-link>
            </template>
          </statistic-card>
        </a-col>
        <a-col :span="6">
          <statistic-card title="流水线状态" :footer="false">
            <pie-count v-if="count.pipelineStatusCount" :data="count.pipelineStatusCount"></pie-count>
          </statistic-card>
        </a-col>
        <a-col :span="6">
          <statistic-card title="最近运行统计" :footer="false">
            <day-count v-if="count.historyCountPerDay" :data="count.historyCountPerDay" title="运行次数"></day-count>
          </statistic-card>
        </a-col>
        <a-col :span="6">
          <statistic-card title="最快到期证书">
          </statistic-card>
        </a-col>
      </a-row>
    </div>

    <div v-if="pluginGroups" class="plugin-list">
      <a-card>
        <template #title>
          已支持的部署任务总览 <a-tag color="green">{{ pluginGroups.groups.all.plugins.length }}</a-tag>
        </template>
        <a-row :gutter="10">
          <a-col v-for="item of pluginGroups.groups.all.plugins" :key="item.name" class="plugin-item-col" :span="4">
            <a-card>
              <a-tooltip :title="item.desc" class="flex-between">
                <div class="plugin-item pointer">
                  <div class="icon">
                    <fs-icon :icon="item.icon" class="font-size-16 color-blue" />
                  </div>
                  <div class="text">
                    <div class="title">{{ item.title }}</div>
                  </div>
                </div>
              </a-tooltip>
            </a-card>
          </a-col>
        </a-row>
      </a-card>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { FsIcon } from "@fast-crud/fast-crud";
import { useUserStore } from "/@/store/modules/user";
import { computed, ComputedRef, onMounted, Ref, ref } from "vue";
import dayjs from "dayjs";
import StatisticCard from "/@/views/framework/home/dashboard/statistic-card.vue";
import DayCount from "./charts/day-count.vue";
import PieCount from "./charts/pie-count.vue";
import ExpiringList from "./charts/expiring-list.vue";
import { useSettingStore } from "/@/store/modules/settings";
import { SiteInfo } from "/@/api/modules/api.basic";
import { UserInfoRes } from "/@/api/modules/api.user";
import { GetStatisticCount } from "/@/views/framework/home/dashboard/api";
import { useRouter } from "vue-router";
import * as api from "./api";
defineOptions({
  name: "DashboardUser"
});

const version = ref(import.meta.env.VITE_APP_VERSION);
const latestVersion = ref();
const hasNewVersion = computed(() => {
  if (!latestVersion.value) {
    return false;
  }
  if (latestVersion.value === version.value) {
    return false;
  }
  //分段比较
  const current = version.value.split(".");
  const latest = latestVersion.value.split(".");
  for (let i = 0; i < current.length; i++) {
    if (parseInt(latest[i]) < parseInt(current[i])) {
      return false;
    }
  }
  return true;
});
async function loadLatestVersion() {
  latestVersion.value = await api.GetLatestVersion();
  console.log("latestVersion", latestVersion.value);
}
const settingStore = useSettingStore();
const siteInfo: Ref<SiteInfo> = computed(() => {
  return settingStore.siteInfo;
});
const settingsStore = useSettingStore();
const userStore = useUserStore();
const userInfo: ComputedRef<UserInfoRes> = computed(() => {
  return userStore.getUserInfo;
});
const now = computed(() => {
  return dayjs().format("YYYY-MM-DD HH:mm:ss");
});
const router = useRouter();
function goPipeline() {
  router.push({ path: "/certd/pipeline" });
}

const count: any = ref({});
function transformStatusCount() {
  const data = count.value.pipelineStatusCount;
  const sorted = [
    { name: "success", label: "成功" },
    { name: "start", label: "运行中" },
    { name: "error", label: "失败" },
    { name: "canceled", label: "已取消" },
    { name: null, label: "未执行" }
  ];
  const result = [];
  for (const item of sorted) {
    const find = data.find((v: any) => v.status === item.name);
    if (find) {
      result.push({ name: item.label, value: find.count });
    } else {
      result.push({ name: item.label, value: 0 });
    }
  }
  count.value.pipelineStatusCount = result;
}
async function loadCount() {
  count.value = await GetStatisticCount();
  transformStatusCount();
  count.value.historyCountPerDay = count.value.historyCountPerDay.map((item: any) => {
    return {
      name: item.date,
      value: item.count
    };
  });
}

async function loadPluginGroups() {
  const groups = await pluginApi.GetGroups({});
  pluginGroups.value = new PluginGroups(groups);
}

const pluginGroups = ref();
onMounted(async () => {
  await userStore.loadUserInfo();
  loadLatestVersion();
  loadCount();
  loadPluginGroups();
});

function openUpgradeUrl() {
  window.open("https://certd.docmirror.cn/guide/install/upgrade.html");
}
</script>

<style lang="less">
.dashboard-user {
  .warning {
    .ant-alert {
      border-left: 0;
      border-right: 0;
      border-radius: 0;
    }
  }
  .header-profile {
    display: flex;
    align-items: center;
    padding: 20px;
    background-color: #fff;

    .avatar {
      margin-right: 10px;
    }
    .text {
      flex: 1;
      display: flex;
      flex-direction: row;
      .left {
        display: flex;
        flex-direction: column;
        justify-content: center;
        > div {
          margin: 4px;
        }
      }
    }
  }
  .notice {
    padding: 20px;
  }
  .plugin-list {
    margin: 0 20px;

    .plugin-item-col {
      margin-bottom: 10px;
      .plugin-item {
        display: flex;
        justify-items: center;
        line-height: 20px;
        .icon {
          display: flex;
          justify-items: center;
          font-size: 20px;
          margin-right: 8px;
        }
        .text {
          overflow: hidden;
          text-overflow: ellipsis;
          word-break: keep-all;
          white-space: nowrap;
        }
      }
    }
  }
}
</style>
