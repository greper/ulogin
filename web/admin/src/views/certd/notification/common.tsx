import { ColumnCompositionProps, compute, dict } from "@fast-crud/fast-crud";
import { computed, provide, ref, toRef } from "vue";
import { useReference } from "/@/use/use-refrence";
import { forEach, get, merge, set } from "lodash-es";
import { Modal } from "ant-design-vue";
import * as api from "/@/views/sys/cname/provider/api";
import { mitter } from "/@/utils/util.mitt";

export function notificationProvide(api: any) {
  provide("notificationApi", api);
  provide("get:plugin:type", () => {
    return "notification";
  });
}

export function getCommonColumnDefine(crudExpose: any, typeRef: any, api: any) {
  const notificationTypeDictRef = dict({
    url: "/pi/notification/getTypeDict"
  });
  const defaultPluginConfig = {
    component: {
      name: "a-input",
      vModel: "value"
    }
  };

  function buildDefineFields(define: any, form: any, mode: string) {
    const formWrapperRef = crudExpose.getFormWrapperRef();
    const columnsRef = toRef(formWrapperRef.formOptions, "columns");

    for (const key in columnsRef.value) {
      if (key.indexOf(".") >= 0) {
        delete columnsRef.value[key];
      }
    }
    console.log('crudBinding.value[mode + "Form"].columns', columnsRef.value);
    forEach(define.input, (value: any, mapKey: any) => {
      const key = "body." + mapKey;
      const field = {
        ...value,
        key
      };
      const column = merge({ title: key }, defaultPluginConfig, field);
      //eval
      useReference(column);

      if (column.required) {
        if (!column.rules) {
          column.rules = [];
        }
        column.rules.push({ required: true, message: "此项必填" });
      }

      //设置默认值
      if (column.value != null && get(form, key) == null) {
        set(form, key, column.value);
      }
      //字段配置赋值
      columnsRef.value[key] = column;
      console.log("form", columnsRef.value, form);
    });
  }

  const currentDefine = ref();

  return {
    id: {
      title: "ID",
      key: "id",
      type: "number",
      column: {
        width: 100
      },
      form: {
        show: false
      }
    },
    type: {
      title: "通知类型",
      type: "dict-select",
      dict: notificationTypeDictRef,
      search: {
        show: false
      },
      column: {
        width: 200,
        component: {
          color: "auto"
        }
      },
      editForm: {
        component: {
          disabled: false
        }
      },
      form: {
        component: {
          disabled: false,
          showSearch: true,
          filterOption: (input: string, option: any) => {
            input = input?.toLowerCase();
            return option.value.toLowerCase().indexOf(input) >= 0 || option.label.toLowerCase().indexOf(input) >= 0;
          },
          renderLabel(item: any) {
            return (
              <span class={"flex-o flex-between"}>
                {item.label}
                {item.needPlus && <fs-icon icon={"mingcute:vip-1-line"} className={"color-plus"}></fs-icon>}
              </span>
            );
          }
        },
        rules: [{ required: true, message: "请选择通知类型" }],
        valueChange: {
          immediate: true,
          async handle({ value, mode, form, immediate }) {
            if (value == null) {
              return;
            }
            const lastTitle = currentDefine.value?.title;
            const define = await api.GetProviderDefine(value);
            currentDefine.value = define;
            console.log("define", define);

            if (!immediate) {
              form.body = {};
              if (define.needPlus) {
                mitter.emit("openVipModal");
              }
            }

            if (!form.name || form.name === lastTitle) {
              form.name = define.title;
            }
            buildDefineFields(define, form, mode);
          }
        },
        helper: computed(() => {
          const define = currentDefine.value;
          if (define == null) {
            return "";
          }
          return define.desc;
        })
      }
    } as ColumnCompositionProps,
    name: {
      title: "通知名称",
      search: {
        show: true
      },
      type: ["text"],
      form: {
        rules: [{ required: true, message: "请填写名称" }],
        helper: "随便填，当多个相同类型的通知时，便于区分"
      },
      column: {
        width: 200
      }
    },
    isDefault: {
      title: "是否默认",
      type: "dict-switch",
      dict: dict({
        data: [
          { label: "是", value: true, color: "success" },
          { label: "否", value: false, color: "default" }
        ]
      }),
      form: {
        value: false,
        rules: [{ required: true, message: "请选择是否默认" }],
        order: 999
      },
      column: {
        align: "center",
        width: 100,
        component: {
          name: "a-switch",
          vModel: "checked",
          disabled: compute(({ value }) => {
            return value === true;
          }),
          on: {
            change({ row }) {
              Modal.confirm({
                title: "提示",
                content: "确定设置为默认通知？",
                onOk: async () => {
                  await api.SetDefault(row.id);
                  await crudExpose.doRefresh();
                },
                onCancel: async () => {
                  await crudExpose.doRefresh();
                }
              });
            }
          }
        }
      }
    } as ColumnCompositionProps,
    test: {
      title: "测试",
      form: {
        show: compute(({ form }) => {
          return !!form.type;
        }),
        component: {
          name: "api-test",
          action: "TestRequest"
        },
        order: 990,
        col: {
          span: 24
        }
      },
      column: {
        show: false
      }
    },
    setting: {
      column: { show: false },
      form: {
        show: false,
        valueBuilder({ value, form }) {
          form.body = {};
          if (!value) {
            return;
          }
          const setting = JSON.parse(value);
          for (const key in setting) {
            form.body[key] = setting[key];
          }
        },
        valueResolve({ form }) {
          const setting = form.body;
          form.setting = JSON.stringify(setting);
        }
      }
    } as ColumnCompositionProps
  };
}
