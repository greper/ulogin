import { request } from "/@/api/service";

export async function GetStatisticCount() {
  return await request({
    url: "/sys/statistic/count",
    method: "POST"
  });
}
