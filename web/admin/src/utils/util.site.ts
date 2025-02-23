import { env } from "./util.env";
export const site = {
  /**
   * @description 更新标题
   * @param titleText
   */
  title: function (titleText: string, baseTitle?: string) {
    const processTitle = baseTitle || env.TITLE || "Certd";
    window.document.title = `${processTitle}${titleText ? ` | ${titleText}` : ""}`;
  }
};
