import { defineStore } from "pinia";
import router from "../../router";
// @ts-ignore
import { LocalStorage } from "/src/utils/util.storage";
// @ts-ignore
import * as UserApi from "/src/api/modules/api.user";
import { RegisterReq, SmsLoginReq } from "/src/api/modules/api.user";
// @ts-ignore
import { LoginReq, UserInfoRes } from "/@/api/modules/api.user";
import { message, Modal, notification } from "ant-design-vue";
import { useI18n } from "vue-i18n";

import { mitter } from "/src/utils/util.mitt";
import {utils} from "/@/utils";

interface UserState {
  userInfo: Nullable<UserInfoRes>;
  token?: string;
}

const USER_INFO_KEY = "USER_INFO";
const TOKEN_KEY = "TOKEN";
export const useUserStore = defineStore({
  id: "app.user",
  state: (): UserState => ({
    // user info
    userInfo: null,
    // token
    token: undefined
  }),
  getters: {
    getUserInfo(): UserInfoRes {
      return this.userInfo || LocalStorage.get(USER_INFO_KEY) || {};
    },
    getToken(): string {
      return this.token || LocalStorage.get(TOKEN_KEY);
    },
    isAdmin(): boolean {
      return this.getUserInfo.roleIds?.includes(1) || this.getUserInfo.id === 1;
    }
  },
  actions: {
    setToken(info: string, expire: number) {
      this.token = info;
      LocalStorage.set(TOKEN_KEY, this.token, expire);
    },
    setUserInfo(info: UserInfoRes) {
      this.userInfo = info;
      LocalStorage.set(USER_INFO_KEY, info);
    },
    resetState() {
      this.userInfo = null;
      this.token = "";
      LocalStorage.remove(TOKEN_KEY);
      LocalStorage.remove(USER_INFO_KEY);
    },
    async register(user: RegisterReq) {
      await UserApi.register(user);
      notification.success({
        message: "注册成功，请登录"
      });
      await router.replace("/login");
    },
    /**
     * @description: login
     */
    async login(loginType: string, params: LoginReq | SmsLoginReq): Promise<any> {
      try {
        let loginRes: any = null;
        if (loginType === "sms") {
          loginRes = await UserApi.loginBySms(params as SmsLoginReq);
        } else {
          // 密码先进行hash
          const password =  await utils.hash.sha256(params.password)
          loginRes = await UserApi.login({
            ...params,
            password
          }as LoginReq);
        }

        const { accessToken, expires } = loginRes.accessInfo;
        // save token
        this.setToken(accessToken, expires);
        // get user info
        return await this.onLoginSuccess(loginRes);
      } catch (error) {
        return null;
      }
    },
    async getUserInfoAction(): Promise<UserInfoRes> {
      const userInfoRes = await UserApi.mine();
      this.setUserInfo(userInfoRes.userInfo);
      return userInfoRes.userInfo;
    },

    async loadUserInfo() {
      await this.getUserInfoAction();
    },

    async onLoginSuccess(loginRes: any) {
      // await this.getUserInfoAction();
      // const userInfo = await this.getUserInfoAction();
      let userInfo = loginRes.userInfo;
      this.setUserInfo(userInfo);
      mitter.emit("app.login", { userInfo, accessInfo: loginRes.accessInfo });
      await router.replace("/");
      return userInfo;
    },

    /**
     * @description: logout
     */
    logout(goLogin = true) {
      this.resetState();
      goLogin && router.push("/login");
      mitter.emit("app.logout");
    },

    /**
     * @description: Confirm before logging out
     */
    confirmLoginOut() {
      const { t } = useI18n();
      Modal.confirm({
        iconType: "warning",
        title: t("app.login.logoutTip"),
        content: t("app.login.logoutMessage"),
        onOk: async () => {
          await this.logout(true);
        }
      });
    }
  }
});
