import { ref } from "vue";
import { getCommonColumnDefine } from "./common";
import { AddReq, CreateCrudOptionsProps, CreateCrudOptionsRet, DelReq, EditReq, UserPageQuery, UserPageRes } from "@fast-crud/fast-crud";
import { createNotificationApi } from "/@/views/certd/notification/api";
const api = createNotificationApi();
export default function ({ crudExpose, context }: CreateCrudOptionsProps): CreateCrudOptionsRet {
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

  const typeRef = ref();
  const commonColumnsDefine = getCommonColumnDefine(crudExpose, typeRef, api);
  return {
    crudOptions: {
      request: {
        pageRequest,
        addRequest,
        editRequest,
        delRequest
      },
      form: {
        labelCol: {
          //固定label宽度
          span: null,
          style: {
            width: "145px"
          }
        }
      },
      rowHandle: {
        width: 200
      },
      columns: {
        ...commonColumnsDefine
      }
    }
  };
}
