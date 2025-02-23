import { request } from "/@/api/service";
import { LocalStorage } from "/@/utils/util.storage";

export async function GetStatisticCount() {
  return await request({
    url: "/statistic/count",
    method: "POST"
  });
}

export async function GetLatestVersion() {
  const latest = LocalStorage.get("latestVersion");
  if (latest) {
    return latest;
  }
  const res = await request({
    url: "/app/latest",
    method: "GET",
    unpack: true
  });
  try {
    const latest = res;
    LocalStorage.set("latestVersion", latest, 60 * 60 * 24);
    return latest;
  } catch (e: any) {
    console.error(e);
    return null;
  }
}
