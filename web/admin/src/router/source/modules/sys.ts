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
      permission: "sys:settings:view"
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
          permission: "sys:auth:user:view"
        }
      },

      {
        title: "系统设置",
        name: "SysSettings",
        path: "/sys/settings",
        component: "/sys/settings/index.vue",
        meta: {
          icon: "ion:settings-outline",
          permission: "sys:settings:view"
        }
      },

      {
        title: "邮件服务器设置",
        name: "EmailSetting",
        path: "/sys/settings/email",
        component: "/sys/settings/email/index.vue",
        meta: {
          permission: "sys:settings:view",
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
          show: () => {
            const settingStore = useSettingStore();
            return settingStore.isComm;
          },
          icon: "ion:document-text-outline",
          permission: "sys:settings:view"
        }
      },
      {
        title: "顶部菜单设置",
        name: "HeaderMenus",
        path: "/sys/settings/header-menus",
        component: "/sys/settings/header-menus/index.vue",
        meta: {
          show: () => {
            const settingStore = useSettingStore();
            return settingStore.isComm;
          },
          icon: "ion:menu",
          permission: "sys:settings:view"
        }
      },
      {
        title: "账号绑定",
        name: "AccountBind",
        path: "/sys/account",
        component: "/sys/account/index.vue",
        meta: {
          icon: "ion:golf-outline",
          permission: "sys:settings:view"
        }
      },
      {
        title: "权限管理",
        name: "PermissionManager",
        path: "/sys/authority/permission",
        component: "/sys/authority/permission/index.vue",
        meta: {
          icon: "ion:list-outline",
          //需要校验权限
          permission: "sys:auth:per:view"
        }
      },
      {
        title: "角色管理",
        name: "RoleManager",
        path: "/sys/authority/role",
        component: "/sys/authority/role/index.vue",
        meta: {
          icon: "ion:people-outline",
          permission: "sys:auth:role:view"
        }
      },
      {
        title: "用户管理",
        name: "UserManager",
        path: "/sys/authority/user",
        component: "/sys/authority/user/index.vue",
        meta: {
          icon: "ion:person-outline",
          permission: "sys:auth:user:view"
        }
      },
    ]
  }
];
