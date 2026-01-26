import { h } from "vue";
import {
  VbenButton,
  VbenCheckbox,
  Input as VbenInput,
  VbenInputPassword,
  VbenPinInput,
  VbenSelect
} from "@vben-core/shadcn-ui";
import { globalShareState } from "@vben-core/shared/global-state";
import { defineRule } from "vee-validate";
const DEFAULT_MODEL_PROP_NAME = "modelValue";
export const DEFAULT_FORM_COMMON_CONFIG = {};
export const COMPONENT_MAP = {
  DefaultButton: h(VbenButton, { size: "sm", variant: "outline" }),
  PrimaryButton: h(VbenButton, { size: "sm", variant: "default" }),
  VbenCheckbox,
  VbenInput,
  VbenInputPassword,
  VbenPinInput,
  VbenSelect
};
export const COMPONENT_BIND_EVENT_MAP = {
  VbenCheckbox: "checked"
};
export function setupVbenForm(options) {
  const { config, defineRules } = options;
  const {
    disabledOnChangeListener = true,
    disabledOnInputListener = true,
    emptyStateValue = void 0
  } = config || {};
  Object.assign(DEFAULT_FORM_COMMON_CONFIG, {
    disabledOnChangeListener,
    disabledOnInputListener,
    emptyStateValue
  });
  if (defineRules) {
    for (const key of Object.keys(defineRules)) {
      defineRule(key, defineRules[key]);
    }
  }
  const baseModelPropName = config?.baseModelPropName ?? DEFAULT_MODEL_PROP_NAME;
  const modelPropNameMap = config?.modelPropNameMap;
  const components = globalShareState.getComponents();
  for (const component of Object.keys(components)) {
    const key = component;
    COMPONENT_MAP[key] = components[component];
    if (baseModelPropName !== DEFAULT_MODEL_PROP_NAME) {
      COMPONENT_BIND_EVENT_MAP[key] = baseModelPropName;
    }
    if (modelPropNameMap && modelPropNameMap[key]) {
      COMPONENT_BIND_EVENT_MAP[key] = modelPropNameMap[key];
    }
  }
}
