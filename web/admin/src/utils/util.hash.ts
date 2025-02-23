export const hashUtils = {
  md5(data: string) {
    throw new Error("Not implemented");
  },
  async sha256(data:string){
      const msgBuffer = new TextEncoder().encode(data); // 将字符串转换为ArrayBuffer
      const hashBuffer = await crypto.subtle.digest('SHA-256', msgBuffer); // 使用SHA-256算法进行哈希计算
      const hashArray = Array.from(new Uint8Array(hashBuffer)); // 将ArrayBuffer转换为Uint8Array，然后转换为数组
       // 将每个字节转换为十六进制并拼接成字符串
    return hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
  }
};
