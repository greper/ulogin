export const amountUtils = {
  toCent(amount: number): number {
    return parseInt((amount * 100).toFixed(0));
  },

  toYuan(amount: number): number {
    return parseFloat((amount / 100).toFixed(2));
  }
};
