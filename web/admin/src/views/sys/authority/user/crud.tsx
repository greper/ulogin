import * as api from "./api";
import { AddReq, CreateCrudOptionsProps, CreateCrudOptionsRet, DelReq, dict, EditReq, UserPageQuery, UserPageRes } from "@fast-crud/fast-crud";
import { useUserStore } from "/@/store/modules/user";

export default function ({ crudExpose }: CreateCrudOptionsProps): CreateCrudOptionsRet {
  const pageRequest = async (query: UserPageQuery): Promise<UserPageRes> => {
    return await api.GetList(query);
  };
  const editRequest = async ({ form, row }: EditReq) => {
    form.id = row.id;
    return await api.UpdateObj(form);
  };
  const delRequest = async ({ row }: DelReq) => {
    return await api.DelObj(row.id);
  };

  const addRequest = async ({ form }: AddReq) => {
    return await api.AddObj(form);
  };

  const userStore = useUserStore();

  return {
    crudOptions: {
      request: {
        pageRequest,
        addRequest,
        editRequest,
        delRequest
      },
      rowHandle: {
        fixed: "right"
      },
      table: {
        scroll: {
          //使用固定列时需要设置此值，并且大于等于列宽度之和的值
          x: 1400
        }
      },
      columns: {
        id: {
          title: "id",
          type: "text",
          form: { show: false }, // 表单配置
          column: {
            width: 100,
            sorter: true
          }
        },
        createTime: {
          title: "创建时间",
          type: "datetime",
          form: { show: false }, // 表单配置
          column: {
            width: 180,
            sorter: true
          }
        },
        // updateTime: {
        //   title: "修改时间",
        //   type: "datetime",
        //   form: { show: false }, // 表单配置
        //   column: {
        //     sortable: "update_time",
        //     width: 180
        //   }
        // },
        username: {
          title: "用户名",
          type: "text",
          search: { show: true }, // 开启查询
          form: {
            rules: [
              { required: true, message: "请输入用户名" },
              { max: 50, message: "最大50个字符" }
            ]
          },
          editForm: { component: { disabled: false } },
          column: {
            sorter: true,
            width: 200
          }
        },
        password: {
          title: "密码",
          type: "text",
          key: "password",
          column: {
            show: false
          },
          form: {
            rules: [{ max: 50, message: "最大50个字符" }],
            component: {
              showPassword: true
            },
            helper: "填写则修改密码"
          }
        },
        nickName: {
          title: "昵称",
          type: "text",
          search: { show: true }, // 开启查询
          form: {
            rules: [{ max: 50, message: "最大50个字符" }]
          },
          column: {
            sorter: true
          }
        },
        email: {
          title: "邮箱",
          type: "text",
          search: { show: true }, // 开启查询
          form: {
            rules: [{ max: 50, message: "最大50个字符" }]
          },
          column: {
            sorter: true,
            width: 160
          }
        },
        mobile: {
          title: "手机号",
          type: "text",
          search: { show: true }, // 开启查询
          form: {
            rules: [{ max: 50, message: "最大50个字符" }]
          },
          column: {
            sorter: true,
            width: 130
          }
        },
        avatar: {
          title: "头像",
          type: "cropper-uploader",
          column: {
            width: 70,
            component: {
              //设置高度，修复操作列错位的问题
              style: {
                height: "30px",
                width: "auto"
              },
              buildUrl(key: string) {
                return `/api/basic/file/download?&key=` + key;
              }
            }
          },
          form: {
            component: {
              vModel: "modelValue",
              valueType: "key",
              cropper: {
                aspectRatio: 1,
                autoCropArea: 1,
                viewMode: 0
              },
              onReady: null,
              uploader: {
                type: "form",
                action: "/basic/file/upload",
                name: "file",
                headers: {
                  Authorization: "Bearer " + userStore.getToken
                },
                successHandle(res: any) {
                  return res;
                }
              },
              buildUrl(key: string) {
                return `/api/basic/file/download?&key=` + key;
              }
            }
          }
        },
        status: {
          title: "状态",
          type: "dict-switch",
          dict: dict({
            data: [
              { label: "启用", value: 1, color: "green" },
              { label: "禁用", value: 0, color: "red" }
            ]
          }),
          column: {
            align: "center",
            sorter: true,
            width: 100
          }
        },
        remark: {
          title: "备注",
          type: "text",
          column: {
            sorter: true
          },
          form: {
            rules: [{ max: 100, message: "最大100个字符" }]
          }
        },
        roles: {
          title: "角色",
          type: "dict-select",
          dict: dict({
            url: "/sys/authority/role/list",
            value: "id",
            label: "name"
          }), // 数据字典
          form: {
            component: { mode: "multiple" }
          },
          column: {
            width: 250,
            sortable: true
          }
        }
      }
    }
  };
}
