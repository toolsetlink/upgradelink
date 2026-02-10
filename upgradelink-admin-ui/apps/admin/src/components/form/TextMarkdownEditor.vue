<script lang="ts" setup>
import { ref, watch, computed } from "vue";
import { Codemirror } from "vue-codemirror";
import { useVModel } from "@vueuse/core";
import { Button, Tooltip, Space, Tag } from "ant-design-vue";
import { FileTextOutlined, CodeOutlined } from "@ant-design/icons-vue";

import { usePreferences } from "@vben/preferences";
import { $t } from "@vben/locales";

// 导入 CodeMirror
import { markdown } from "@codemirror/lang-markdown";
import { githubDark, githubLight } from "@uiw/codemirror-theme-github";
import { indentOnInput } from "@codemirror/language";

const props = defineProps({
  value: { type: String, default: undefined },
  placeholder: { type: String, default: "" },
  autoSize: { type: Object, default: () => ({ minRows: 3 }) },
});

const emits = defineEmits(["update:value"]);
const state = useVModel(props, "value", emits, {
  defaultValue: props.value,
  passive: true,
});

const inputValue = ref<string>(props.value || "");
const editorMode = ref<'text' | 'markdown'>('text');

const preference = usePreferences();
const isDark = computed(() => preference.isDark.value);

// 监听外部值变化
watch(
  () => props.value,
  (newValue) => {
    inputValue.value = newValue || "";
  }
);

// 切换编辑模式
function toggleMode() {
  editorMode.value = editorMode.value === 'text' ? 'markdown' : 'text';
}

// 获取当前模式的显示信息
const currentModeInfo = computed(() => {
  if (editorMode.value === 'markdown') {
    return {
      label: $t('component.markdownEditor.markdownEditor'),
      color: 'blue',
      description: $t('component.markdownEditor.markdownDescription')
    };
  } else {
    return {
      label: $t('component.markdownEditor.textEditor'),
      color: 'green',
      description: $t('component.markdownEditor.textDescription')
    };
  }
});

// CodeMirror 配置
const extensions = computed(() => {
  const exts = [];
  if (isDark.value) {
    exts.push(githubDark);
  } else {
    exts.push(githubLight);
  }
  if (editorMode.value === 'markdown') {
    exts.push(markdown());
  }
  exts.push(indentOnInput());
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
</script>
<template>
  <div 
    v-bind="$attrs"
    style="width: 100%; border-radius: 8px; box-shadow: 0 1px 2px rgba(0,0,0,0.08); overflow: hidden; border: 1px solid #f0f0f0;"
  >
    <!-- 工具栏 -->
    <div style="padding: 12px; background: rgba(0,0,0,0.02); border-bottom: 1px solid #f0f0f0;">
      <Space style="width: 100%; justify-content: space-between; align-items: center;">
        <!-- 当前模式信息 -->
        <Space align="center">
          <Tag :color="currentModeInfo.color" style="font-size: 13px; font-weight: 500;">
            {{ currentModeInfo.label }}
          </Tag>
          <span style="font-size: 12px; color: #666;">{{ currentModeInfo.description }}</span>
        </Space>
        
        <!-- 操作按钮 -->
        <Tooltip :title="editorMode === 'text' ? $t('component.markdownEditor.switchToMarkdown') : $t('component.markdownEditor.switchToText')">
          <Button 
            type="primary" 
            shape="circle" 
            size="small"
            @click="toggleMode"
          >
            <template #icon>
              <CodeOutlined v-if="editorMode === 'text'" /><FileTextOutlined v-else />
            </template>
          </Button>
        </Tooltip>
      </Space>
    </div>
    
    <!-- 编辑器 -->
    <div style="padding: 12px;">
      <Codemirror
        v-model="state"
        :autofocus="true"
        :extensions="extensions"
        :indent-with-tab="true"
        :style="editorStyle"
        :tab-size="2"
        :placeholder="props.placeholder"
      />
    </div>
  </div>
</template>