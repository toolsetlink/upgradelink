import {
  defineComponent,
  h,
  inject,
  nextTick,
  provide,
  reactive,
  ref
} from "vue";
import { useStore } from "@vben-core/shared/store";
import { DrawerApi } from "./drawer-api.mjs";
import VbenDrawer from "./drawer.vue";
const USER_DRAWER_INJECT_KEY = Symbol("VBEN_DRAWER_INJECT");
const DEFAULT_DRAWER_PROPS = {};
export function setDefaultDrawerProps(props) {
  Object.assign(DEFAULT_DRAWER_PROPS, props);
}
export function useVbenDrawer(options = {}) {
  const { connectedComponent } = options;
  if (connectedComponent) {
    const extendedApi2 = reactive({});
    const isDrawerReady = ref(true);
    const Drawer2 = defineComponent(
      (props, { attrs, slots }) => {
        provide(USER_DRAWER_INJECT_KEY, {
          extendApi(api2) {
            Object.setPrototypeOf(extendedApi2, api2);
          },
          options,
          async reCreateDrawer() {
            isDrawerReady.value = false;
            await nextTick();
            isDrawerReady.value = true;
          }
        });
        checkProps(extendedApi2, {
          ...props,
          ...attrs,
          ...slots
        });
        return () => h(
          isDrawerReady.value ? connectedComponent : "div",
          { ...props, ...attrs },
          slots
        );
      },
      // eslint-disable-next-line vue/one-component-per-file
      {
        name: "VbenParentDrawer",
        inheritAttrs: false
      }
    );
    return [Drawer2, extendedApi2];
  }
  const injectData = inject(USER_DRAWER_INJECT_KEY, {});
  const mergedOptions = {
    ...DEFAULT_DRAWER_PROPS,
    ...injectData.options,
    ...options
  };
  mergedOptions.onOpenChange = (isOpen) => {
    options.onOpenChange?.(isOpen);
    injectData.options?.onOpenChange?.(isOpen);
  };
  const onClosed = mergedOptions.onClosed;
  mergedOptions.onClosed = () => {
    onClosed?.();
    if (mergedOptions.destroyOnClose) {
      injectData.reCreateDrawer?.();
    }
  };
  const api = new DrawerApi(mergedOptions);
  const extendedApi = api;
  extendedApi.useStore = (selector) => {
    return useStore(api.store, selector);
  };
  const Drawer = defineComponent(
    (props, { attrs, slots }) => {
      return () => h(VbenDrawer, { ...props, ...attrs, drawerApi: extendedApi }, slots);
    },
    // eslint-disable-next-line vue/one-component-per-file
    {
      name: "VbenDrawer",
      inheritAttrs: false
    }
  );
  injectData.extendApi?.(extendedApi);
  return [Drawer, extendedApi];
}
async function checkProps(api, attrs) {
  if (!attrs || Object.keys(attrs).length === 0) {
    return;
  }
  await nextTick();
  const state = api?.store?.state;
  if (!state) {
    return;
  }
  const stateKeys = new Set(Object.keys(state));
  for (const attr of Object.keys(attrs)) {
    if (stateKeys.has(attr) && !["class"].includes(attr)) {
      console.warn(
        `[Vben Drawer]: When 'connectedComponent' exists, do not set props or slots '${attr}', which will increase complexity. If you need to modify the props of Drawer, please use useVbenDrawer or api.`
      );
    }
  }
}
