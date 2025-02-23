// @ts-ignore
import { useI18n } from "vue-i18n";
import { AddReq, CreateCrudOptionsProps, CreateCrudOptionsRet, DelReq, dict, EditReq, UserPageQuery, UserPageRes } from "@fast-crud/fast-crud";
import { OPEN_API_DOC, openkeyApi } from "./api";
import { useModal } from "/@/use/use-modal";

export default function ({ crudExpose, context }: CreateCrudOptionsProps): CreateCrudOptionsRet {
  const { t } = useI18n();
  const api = openkeyApi;
  const pageRequest = async (query: UserPageQuery): Promise<UserPageRes> => {
    return await api.GetList(query);
  };
  const editRequest = async (req: EditReq) => {
    const { form, row } = req;
    form.id = row.id;
    const res = await api.UpdateObj(form);
    return res;
  };
  const delRequest = async (req: DelReq) => {
    const { row } = req;
    return await api.DelObj(row.id);
  };

  const addRequest = async (req: AddReq) => {
    const { form } = req;
    const res = await api.AddObj(form);
    return res;
  };
  const model = useModal();
  return {
    crudOptions: {
      request: {
        pageRequest,
        addRequest,
        editRequest,
        delRequest
      },
      search: {
        show: false
      },
      form: {
        labelCol: {
          //固定label宽度
          span: null,
          style: {
            width: "100px"
          }
        },
        col: {
          span: 22
        },
        wrapper: {
          width: 600
        }
      },
      actionbar: {
        buttons: {
          add: {
            text: "生成新的Key"
          }
        }
      },
      rowHandle: {
        width: 300,
        fixed: "right",
        buttons: {
          view: { show: true },
          copy: { show: false },
          edit: { show: false },
          remove: { show: true },
          gen: {
            text: "接口测试",
            size: "mini",
            icon: "devicon-plain:vitest",
            type: "primary",
            async click({ row }) {
              const apiToken = await api.GetApiToken(row.id);

              model.success({
                title: "x-certd-token",
                maskClosable: true,
                okText: "确定",
                width: 600,
                content: () => {
                  return (
                    <div>
                      <div class={"m-10 p-10"}>
                        测试x-certd-token如下，您可以在3分钟内使用它进行
                        <a href={OPEN_API_DOC} target={"_blank"}>
                          开放接口
                        </a>
                        请求测试
                      </div>
                      <div class={"m-10 p-10"} style={{ border: "1px solid #333" }}>
                        <fs-copyable model-value={apiToken}></fs-copyable>
                      </div>
                    </div>
                  );
                }
              });
            }
          }
        }
      },
      columns: {
        id: {
          title: "ID",
          key: "id",
          type: "number",
          search: {
            show: false
          },
          column: {
            width: 100,
            editable: {
              disabled: true
            }
          },
          form: {
            show: false
          }
        },
        keyId: {
          title: "KeyId",
          type: ["text", "copyable"],
          search: {
            show: true
          },
          form: {
            show: false
          },
          column: {
            width: 250,
            sorter: true
          }
        },
        keySecret: {
          title: "KeySecret",
          type: ["text", "copyable"],
          form: {
            show: false
          },
          column: {
            width: 580,
            sorter: true
          }
        },
        scope: {
          title: "权限范围",
          type: "dict-radio",
          dict: dict({
            data: [
              { label: "仅开放接口", value: "open", color: "blue" },
              { label: "账户所有权限", value: "user", color: "red" }
            ]
          }),
          form: {
            value: "open",
            show: true,
            rules: [{ required: true, message: "此项必填" }],
            helper: "仅开放接口只可以访问开放接口，账户所有权限可以访问所有接口",
            component: {
              vModel: "value"
            }
          },
          column: {
            width: 120,
            align: "center",
            sorter: true
          }
        },
        createTime: {
          title: "创建时间",
          type: "datetime",
          search: {
            show: false
          },
          form: {
            show: false
          }
        }
      }
    }
  };
}
