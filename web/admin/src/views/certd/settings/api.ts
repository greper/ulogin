// @ts-ignore
import { request } from "/@/api/service";
import { SysPrivateSetting, SysPublicSetting } from "/@/api/modules/api.basic";
const apiPrefix = "/user/settings";
export type UserSettings = {
  defaultNotification?: number;
  defaultCron?: string;
};

export async function UserSettingsGet() {
  const res = await request({
    url: apiPrefix + "/getDefault",
    method: "post"
  });
  if (!res) {
    return {};
  }
  return res;
}

export async function UserSettingsSave(setting: any) {
  return await request({
    url: apiPrefix + "/saveDefault",
    method: "post",
    data: setting
  });
}
