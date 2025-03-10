export default [
  {
    path: "/login",
    method: "post",
    handle() {
      return {
        code: 0,
        message: "success",
        data: {
          token: "faker token",
          expire: 10000
        }
      };
    }
  },
  {
    path: "/sys/authority/user/mine",
    method: "get",
    handle() {
      return {
        code: 0,
        message: "success",
        data: {
          id: 1,
          username: "username",
          nickName: "admin"
        }
      };
    }
  }
];
