import { request } from "/src/api/service";
const apiPrefix = "/mine/email";

export async function TestSend(receiver: string) {
  await request({
    url: apiPrefix + "/test",
    method: "post",
    data: {
      receiver
    }
  });
}
