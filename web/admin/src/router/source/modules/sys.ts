import LayoutPass from "/@/layout/layout-pass.vue";
import { useSettingStore } from "/@/store/modules/settings";

export const sysResources = [
  {
    title: "系统管理",
    name: "SysRoot",
    path: "/sys",
    redirect: "/sys/settings",
    component: LayoutPass,
    meta: {
      icon: "ion:settings-outline",
    },
    children: [
      {
        title: "控制台",
        name: "SysConsole",
        path: "/sys/console",
        component: "/sys/console/index.vue",
        meta: {
          show: () => {
            const settingStore = useSettingStore();
            return settingStore.isComm;
          },
          icon: "ion:speedometer-outline",
        }
      },
      {
        title: "系统设置",
        name: "SysSettings",
        path: "/sys/settings",
        component: "/sys/settings/index.vue",
        meta: {
          icon: "ion:settings-outline",
        }
      },
      {
        title: "邮件服务器设置",
        name: "EmailSetting",
        path: "/sys/settings/email",
        component: "/sys/settings/email/index.vue",
        meta: {
          icon: "ion:mail-outline",
          auth: true
        }
      },
      {
        title: "站点个性化",
        name: "SiteSetting",
        path: "/sys/site",
        component: "/sys/site/index.vue",
        meta: {
          icon: "ion:document-text-outline",
        }
      },
      {
        title: "用户管理",
        name: "UserManager",
        path: "/sys/authority/user",
        component: "/sys/authority/user/index.vue",
        meta: {
          icon: "ion:person-outline",
        }
      },
    ]
  }
];
