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
import { ModalApi } from "./modal-api.mjs";
import VbenModal from "./modal.vue";
const USER_MODAL_INJECT_KEY = Symbol("VBEN_MODAL_INJECT");
const DEFAULT_MODAL_PROPS = {};
export function setDefaultModalProps(props) {
  Object.assign(DEFAULT_MODAL_PROPS, props);
}
export function useVbenModal(options = {}) {
  const { connectedComponent } = options;
  if (connectedComponent) {
    const extendedApi2 = reactive({});
    const isModalReady = ref(true);
    const Modal2 = defineComponent(
      (props, { attrs, slots }) => {
        provide(USER_MODAL_INJECT_KEY, {
          extendApi(api2) {
            Object.setPrototypeOf(extendedApi2, api2);
          },
          options,
          async reCreateModal() {
            isModalReady.value = false;
            await nextTick();
            isModalReady.value = true;
          }
        });
        checkProps(extendedApi2, {
          ...props,
          ...attrs,
          ...slots
        });
        return () => h(
          isModalReady.value ? connectedComponent : "div",
          {
            ...props,
            ...attrs
          },
          slots
        );
      },
      // eslint-disable-next-line vue/one-component-per-file
      {
        name: "VbenParentModal",
        inheritAttrs: false
      }
    );
    return [Modal2, extendedApi2];
  }
  const injectData = inject(USER_MODAL_INJECT_KEY, {});
  const mergedOptions = {
    ...DEFAULT_MODAL_PROPS,
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
      injectData.reCreateModal?.();
    }
  };
  const api = new ModalApi(mergedOptions);
  const extendedApi = api;
  extendedApi.useStore = (selector) => {
    return useStore(api.store, selector);
  };
  const Modal = defineComponent(
    (props, { attrs, slots }) => {
      return () => h(
        VbenModal,
        {
          ...props,
          ...attrs,
          modalApi: extendedApi
        },
        slots
      );
    },
    // eslint-disable-next-line vue/one-component-per-file
    {
      name: "VbenModal",
      inheritAttrs: false
    }
  );
  injectData.extendApi?.(extendedApi);
  return [Modal, extendedApi];
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
        `[Vben Modal]: When 'connectedComponent' exists, do not set props or slots '${attr}', which will increase complexity. If you need to modify the props of Modal, please use useVbenModal or api.`
      );
    }
  }
}
