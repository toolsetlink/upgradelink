<script lang="ts" setup>
import { onBeforeUpdate, ref, computed } from "vue";
import { Codemirror } from "vue-codemirror";

import { usePreferences } from "@vben/preferences";
import { $t } from "@vben/locales";

import { json } from "@codemirror/lang-json";
import { githubDark, githubLight } from "@uiw/codemirror-theme-github";
import { useVModel } from "@vueuse/core";
import JSONBigInt from "json-bigint";

import { message } from "ant-design-vue";
import { Button, Tooltip, Space } from "ant-design-vue";
import { 
  CheckCircleOutlined, 
  CloseCircleOutlined,
  FormatPainterOutlined,
  CopyOutlined,
  ReloadOutlined,
  CheckOutlined
} from "@ant-design/icons-vue";

const props = defineProps({
  value: { type: String, default: undefined },
  autoValidate: { type: Boolean, default: true },
  autoFormat: { type: Boolean, default: true },
});

const emits = defineEmits(["update:value", "validate"]);
const state = useVModel(props, "value", emits, {
  defaultValue: props.value,
  passive: true,
});

const inputValue = ref<string>("");
const validationError = ref<string>("");
const isValid = ref<boolean>(true);

onBeforeUpdate(() => {
  inputValue.value = props.value as string;
  if (props.autoValidate) {
    validateJSON(props.value as string);
  }
});

const preference = usePreferences();
const isDark = computed(() => preference.isDark.value);

const extensions = computed(() => {
  const exts = [];
  if (isDark.value) {
    exts.push(githubDark);
  } else {
    exts.push(githubLight);
  }
  exts.push(json());
  return exts;
});

// 编辑器样式
const editorStyle = computed(() => {
  return {
    width: '100%',
    minHeight: '400px',
    border: `1px solid ${isDark.value ? '#303030' : '#d9d9d9'}`,
    borderRadius: '4px',
    backgroundColor: isDark.value ? '#1f1f1f' : '#ffffff'
  };
});

function validateJSON(value: string) {
  if (!value) {
    isValid.value = true;
    validationError.value = "";
    emits("validate", true);
    return true;
  }

  try {
    JSONBigInt({ storeAsString: true, strict: true }).parse(value);
    isValid.value = true;
    validationError.value = "";
    emits("validate", true);
    return true;
  } catch (error: any) {
    isValid.value = false;
    validationError.value = error.message;
    emits("validate", false);
    return false;
  }
}

function beautifyJSON() {
  if (!inputValue.value) {
    message.warning($t("component.jsonEditor.pleaseInputJson"));
    return;
  }

  try {
    const parsed = JSONBigInt({ storeAsString: true, strict: true }).parse(inputValue.value);
    const beautified = JSON.stringify(parsed, null, 2);
    state.value = beautified;
    validateJSON(beautified);
    message.success($t("component.jsonEditor.beautifySuccess"));
  } catch (error) {
    message.error($t("component.jsonEditor.beautifyError"));
  }
}

function handleValueChange(v: string) {
  emits("update:value", v);
  if (props.autoValidate) {
    validateJSON(v);
  }
}

function copyToClipboard() {
  if (!inputValue.value) {
    message.warning($t("component.jsonEditor.noContentToCopy"));
    return;
  }
  navigator.clipboard.writeText(inputValue.value).then(() => {
    message.success($t("component.jsonEditor.copySuccess"));
  }).catch(() => {
    message.error($t("component.jsonEditor.copyError"));
  });
}

function resetEditor() {
  state.value = "";
  validationError.value = "";
  isValid.value = true;
  message.info($t("component.jsonEditor.editorReset"));
}

function testJSON() {
  if (!inputValue.value) {
    message.warning($t("component.jsonEditor.pleaseInputJson"));
    return;
  }
  const result = validateJSON(inputValue.value);
  if (result) {
    message.success($t("component.jsonEditor.jsonValid"));
  } else {
    message.error($t("component.jsonEditor.jsonInvalid"));
  }
}
</script>
<template>
  <div style="width: 100%; border-radius: 8px; box-shadow: 0 1px 2px rgba(0,0,0,0.08); overflow: hidden; border: 1px solid #f0f0f0;">
    <!-- 工具栏 -->
    <div style="padding: 12px; background: rgba(0,0,0,0.02); border-bottom: 1px solid #f0f0f0;">
      <Space style="width: 100%; justify-content: space-between; align-items: center;">
        <!-- 验证状态 -->
        <Space align="center">
          <Tooltip v-if="validationError" :title="$t('component.jsonEditor.jsonFormatError')">
            <CloseCircleOutlined style="color: #ff4d4f; font-size: 16px;" />
          </Tooltip>
          <Tooltip v-else :title="$t('component.jsonEditor.jsonFormatValid')">
            <CheckCircleOutlined style="color: #52c41a; font-size: 16px;" />
          </Tooltip>
          <span v-if="validationError" style="font-size: 12px; color: #666;">{{ validationError }}</span>
          <span v-else style="font-size: 12px; color: #666;">{{ $t('component.jsonEditor.jsonFormatValid') }}</span>
        </Space>
        
        <!-- 操作按钮 -->
        <Space size="small">
          <Tooltip :title="$t('component.jsonEditor.beautifyJsonFormat')">
            <Button 
              type="primary" 
              shape="circle" 
              size="small"
              @click="beautifyJSON"
            >
              <FormatPainterOutlined />
            </Button>
          </Tooltip>
          
          <Tooltip :title="$t('component.jsonEditor.validateJsonFormat')">
            <Button 
              shape="circle" 
              size="small"
              @click="testJSON"
            >
              <CheckOutlined />
            </Button>
          </Tooltip>
          
          <Tooltip :title="$t('component.jsonEditor.copyToClipboard')">
            <Button 
              shape="circle" 
              size="small"
              @click="copyToClipboard"
            >
              <CopyOutlined />
            </Button>
          </Tooltip>
          
          <Tooltip :title="$t('component.jsonEditor.resetEditor')">
            <Button 
              danger 
              shape="circle" 
              size="small"
              @click="resetEditor"
            >
              <ReloadOutlined />
            </Button>
          </Tooltip>
        </Space>
      </Space>
    </div>
    
    <!-- 编辑器 -->
    <div style="padding: 12px;">
      <Codemirror
        v-bind="$attrs"
        v-model="state"
        :autofocus="true"
        :extensions="extensions"
        :indent-with-tab="true"
        :style="editorStyle"
        :tab-size="2"
        @change="handleValueChange"
      />
    </div>
  </div>
</template>