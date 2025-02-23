import { useRoute, useRouter } from "vue-router";
import { ref, watch, onMounted, onUnmounted, resolveComponent, nextTick, defineComponent } from "vue";
import * as _ from "lodash-es";
import BScroll from "better-scroll";
import "./index.less";
import { utils } from "@fast-crud/fast-crud";
import { routerUtils } from "/@/utils/util.router";

defineOptions()

export default defineComponent({
  name: "FsMenu",
  inheritAttrs: true,
  props: {
    menus: {},
    expandSelected: {
      default: false
    },
    scroll: {}
  },
  setup(props, ctx) {
    async function onSelect(item: any) {
      await routerUtils.open(item.key);
    }

    const items = ref([]);

    function buildItemMenus(menus: any) {
      if (menus == null) {
        return;
      }
      const list: any = [];
      for (const sub of menus) {
        const item: any = {
          key: sub.path,
          label: sub.title,
          title: sub.title,
          icon: () => {
            return <fsIcon icon={sub.icon ?? sub.meta?.icon} />;
          }
        };
        list.push(item);
        if (sub.children && sub.children.length > 0) {
          item.children = buildItemMenus(sub.children);
        }
      }
      return list;
    }
    items.value = buildItemMenus(props.menus);
    console.log("items", items.value);
    const fsIcon = resolveComponent("FsIcon");

    const buildMenus = (children: any) => {
      const slots: any = [];
      if (children == null) {
        return slots;
      }
      for (const sub of children) {
        if (sub.meta?.show != null) {
          if (sub.meta.show === false || (typeof sub.meta.show === "function" && !sub.meta.show())) {
            continue;
          }
        }
        const title: any = () => {
          const icon = sub.icon || sub?.meta?.icon;
          if (icon) {
            // @ts-ignore , anticon必须要有，不然不能折叠
            return (
              <div class={"menu-item-title"}>
                <fsIcon class={"anticon"} icon={icon} />
                <span>{sub.title}</span>
              </div>
            );
          }
          return sub.title;
        };
        if (sub.children && sub.children.length > 0) {
          // eslint-disable-next-line @typescript-eslint/no-unused-vars
          const subSlots = {
            default: () => {
              return buildMenus(sub.children);
            },
            title
          };
          function onTitleClick() {
            if (sub.path && ctx.attrs.mode === "horizontal") {
              open(sub.path);
            }
          }
          slots.push(<a-sub-menu key={sub.path} v-slots={subSlots} />);
        } else {
          slots.push(
            <a-menu-item key={sub.path} title={sub.title}>
              {title}
            </a-menu-item>
          );
        }
      }
      return slots;
    };
    const slots = {
      default() {
        return buildMenus(props.menus);
      }
    };
    const selectedKeys = ref([]);
    const openKeys = ref([]);

    const route = useRoute();
    const router = useRouter();

    function openSelectedParents(fullPath: any) {
      if (!props.expandSelected) {
        return;
      }
      if (props.menus == null) {
        return;
      }
      const keys: any = [];
      let changed = false;
      utils.deepdash.forEachDeep(props.menus, (value: any, key: any, parent: any, context: any) => {
        if (value == null) {
          return;
        }
        if (value.path === fullPath) {
          _.forEach(context.parents, (item) => {
            if (item.value instanceof Array) {
              return;
            }
            keys.push(item.value.path);
          });
        }
      });
      if (keys.length > 0) {
        for (const key of keys) {
          if (openKeys.value.indexOf(key) === -1) {
            openKeys.value.push(key);
            changed = true;
          }
        }
      }
      return changed;
    }

    // const { asideMenuRef, onOpenChange } = useBetterScroll(props.scroll as any);

    watch(
      () => {
        return route.fullPath;
      },
      (path) => {
        // path = route.fullPath;
        selectedKeys.value = [path];
        const changed = openSelectedParents(path);
        if (changed) {
          // onOpenChange();
        }
      },
      {
        immediate: true
      }
    );
    return () => {
      const menu = (
        <a-menu
          mode={"inline"}
          theme={"light"}
          // v-slots={slots}
          // onClick={onSelect}
          // onOpenChange={onOpenChange}
          v-models={[
            [openKeys.value, "openKeys"],
            [selectedKeys.value, "selectedKeys"]
          ]}
          items={items.value}
          inlineCollapsed={!props.expandSelected}
        />
      );
      const classNames = { "fs-menu-wrapper": true, "fs-menu-better-scroll": props.scroll };
      return <div>{menu}</div>;
    };
  }
});
