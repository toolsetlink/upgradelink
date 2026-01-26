import { isRef, toRaw } from "vue";
import { Store } from "@vben-core/shared/store";
import {
  bindMethods,
  createMerge,
  formatDate,
  isDate,
  isDayjsObject,
  isFunction,
  isObject,
  mergeWithArrayOverride,
  StateHandler
} from "@vben-core/shared/utils";
function getDefaultState() {
  return {
    actionWrapperClass: "",
    collapsed: false,
    collapsedRows: 1,
    collapseTriggerResize: false,
    commonConfig: {},
    handleReset: void 0,
    handleSubmit: void 0,
    handleValuesChange: void 0,
    handleCollapsedChange: void 0,
    layout: "horizontal",
    resetButtonOptions: {},
    schema: [],
    scrollToFirstError: false,
    showCollapseButton: false,
    showDefaultActions: true,
    submitButtonOptions: {},
    submitOnChange: false,
    submitOnEnter: false,
    wrapperClass: "grid-cols-1"
  };
}
export class FormApi {
  // private api: Pick<VbenFormProps, 'handleReset' | 'handleSubmit'>;
  form = {};
  isMounted = false;
  state = null;
  stateHandler;
  store;
  /**
   * 组件实例映射
   */
  componentRefMap = /* @__PURE__ */ new Map();
  // 最后一次点击提交时的表单值
  latestSubmissionValues = null;
  prevState = null;
  constructor(options = {}) {
    const { ...storeState } = options;
    const defaultState = getDefaultState();
    this.store = new Store(
      {
        ...defaultState,
        ...storeState
      },
      {
        onUpdate: () => {
          this.prevState = this.state;
          this.state = this.store.state;
          this.updateState();
        }
      }
    );
    this.state = this.store.state;
    this.stateHandler = new StateHandler();
    bindMethods(this);
  }
  /**
   * 获取字段组件实例
   * @param fieldName 字段名
   * @returns 组件实例
   */
  getFieldComponentRef(fieldName) {
    let target = this.componentRefMap.has(fieldName) ? this.componentRefMap.get(fieldName) : void 0;
    if (target && target.$.type.name === "AsyncComponentWrapper" && target.$.subTree.ref) {
      if (Array.isArray(target.$.subTree.ref)) {
        if (target.$.subTree.ref.length > 0 && isRef(target.$.subTree.ref[0]?.r)) {
          target = target.$.subTree.ref[0]?.r.value;
        }
      } else if (isRef(target.$.subTree.ref.r)) {
        target = target.$.subTree.ref.r.value;
      }
    }
    return target;
  }
  /**
   * 获取当前聚焦的字段，如果没有聚焦的字段则返回undefined
   */
  getFocusedField() {
    for (const fieldName of this.componentRefMap.keys()) {
      const ref = this.getFieldComponentRef(fieldName);
      if (ref) {
        let el = null;
        if (ref instanceof HTMLElement) {
          el = ref;
        } else if (ref.$el instanceof HTMLElement) {
          el = ref.$el;
        }
        if (!el) {
          continue;
        }
        if (el === document.activeElement || el.contains(document.activeElement)) {
          return fieldName;
        }
      }
    }
    return void 0;
  }
  getLatestSubmissionValues() {
    return this.latestSubmissionValues || {};
  }
  getState() {
    return this.state;
  }
  async getValues() {
    const form = await this.getForm();
    return form.values ? this.handleRangeTimeValue(form.values) : {};
  }
  async isFieldValid(fieldName) {
    const form = await this.getForm();
    return form.isFieldValid(fieldName);
  }
  merge(formApi) {
    const chain = [this, formApi];
    const proxy = new Proxy(formApi, {
      get(target, prop) {
        if (prop === "merge") {
          return (nextFormApi) => {
            chain.push(nextFormApi);
            return proxy;
          };
        }
        if (prop === "submitAllForm") {
          return async (needMerge = true) => {
            try {
              const results = await Promise.all(
                chain.map(async (api) => {
                  const validateResult = await api.validate();
                  if (!validateResult.valid) {
                    return;
                  }
                  const rawValues = toRaw(await api.getValues() || {});
                  return rawValues;
                })
              );
              if (needMerge) {
                const mergedResults = Object.assign({}, ...results);
                return mergedResults;
              }
              return results;
            } catch (error) {
              console.error("Validation error:", error);
            }
          };
        }
        return target[prop];
      }
    });
    return proxy;
  }
  mount(formActions, componentRefMap) {
    if (!this.isMounted) {
      Object.assign(this.form, formActions);
      this.stateHandler.setConditionTrue();
      this.setLatestSubmissionValues({
        ...toRaw(this.handleRangeTimeValue(this.form.values))
      });
      this.componentRefMap = componentRefMap;
      this.isMounted = true;
    }
  }
  /**
   * 根据字段名移除表单项
   * @param fields
   */
  async removeSchemaByFields(fields) {
    const fieldSet = new Set(fields);
    const schema = this.state?.schema ?? [];
    const filterSchema = schema.filter((item) => !fieldSet.has(item.fieldName));
    this.setState({
      schema: filterSchema
    });
  }
  /**
   * 重置表单
   */
  async resetForm(state, opts) {
    const form = await this.getForm();
    return form.resetForm(state, opts);
  }
  async resetValidate() {
    const form = await this.getForm();
    const fields = Object.keys(form.errors.value);
    fields.forEach((field) => {
      form.setFieldError(field, void 0);
    });
  }
  /**
   * 滚动到第一个错误字段
   * @param errors 验证错误对象
   */
  scrollToFirstError(errors) {
    const firstErrorFieldName = typeof errors === "string" ? errors : Object.keys(errors)[0];
    if (!firstErrorFieldName) {
      return;
    }
    let el = document.querySelector(
      `[name="${firstErrorFieldName}"]`
    );
    if (!el) {
      const componentRef = this.getFieldComponentRef(firstErrorFieldName);
      if (componentRef && componentRef.$el instanceof HTMLElement) {
        el = componentRef.$el;
      }
    }
    if (el) {
      el.scrollIntoView({
        behavior: "smooth",
        block: "center",
        inline: "nearest"
      });
    }
  }
  async setFieldValue(field, value, shouldValidate) {
    const form = await this.getForm();
    form.setFieldValue(field, value, shouldValidate);
  }
  setLatestSubmissionValues(values) {
    this.latestSubmissionValues = { ...toRaw(values) };
  }
  setState(stateOrFn) {
    if (isFunction(stateOrFn)) {
      this.store.setState((prev) => {
        return mergeWithArrayOverride(stateOrFn(prev), prev);
      });
    } else {
      this.store.setState((prev) => mergeWithArrayOverride(stateOrFn, prev));
    }
  }
  /**
   * 设置表单值
   * @param fields record
   * @param filterFields 过滤不在schema中定义的字段 默认为true
   * @param shouldValidate
   */
  async setValues(fields, filterFields = true, shouldValidate = false) {
    const form = await this.getForm();
    if (!filterFields) {
      form.setValues(fields, shouldValidate);
      return;
    }
    const fieldMergeFn = createMerge((obj, key, value) => {
      if (key in obj) {
        obj[key] = !Array.isArray(obj[key]) && isObject(obj[key]) && !isDayjsObject(obj[key]) && !isDate(obj[key]) ? fieldMergeFn(value, obj[key]) : value;
      }
      return true;
    });
    const filteredFields = fieldMergeFn(fields, form.values);
    form.setValues(filteredFields, shouldValidate);
  }
  async submitForm(e) {
    e?.preventDefault();
    e?.stopPropagation();
    const form = await this.getForm();
    await form.submitForm();
    const rawValues = toRaw(await this.getValues());
    await this.state?.handleSubmit?.(rawValues);
    return rawValues;
  }
  unmount() {
    this.form?.resetForm?.();
    this.latestSubmissionValues = null;
    this.isMounted = false;
    this.stateHandler.reset();
  }
  updateSchema(schema) {
    const updated = [...schema];
    const hasField = updated.every(
      (item) => Reflect.has(item, "fieldName") && item.fieldName
    );
    if (!hasField) {
      console.error(
        "All items in the schema array must have a valid `fieldName` property to be updated"
      );
      return;
    }
    const currentSchema = [...this.state?.schema ?? []];
    const updatedMap = {};
    updated.forEach((item) => {
      if (item.fieldName) {
        updatedMap[item.fieldName] = item;
      }
    });
    currentSchema.forEach((schema2, index) => {
      const updatedData = updatedMap[schema2.fieldName];
      if (updatedData) {
        currentSchema[index] = mergeWithArrayOverride(
          updatedData,
          schema2
        );
      }
    });
    this.setState({ schema: currentSchema });
  }
  async validate(opts) {
    const form = await this.getForm();
    const validateResult = await form.validate(opts);
    if (Object.keys(validateResult?.errors ?? {}).length > 0) {
      console.error("validate error", validateResult?.errors);
      if (this.state?.scrollToFirstError) {
        this.scrollToFirstError(validateResult.errors);
      }
    }
    return validateResult;
  }
  async validateAndSubmitForm() {
    const form = await this.getForm();
    const { valid, errors } = await form.validate();
    if (!valid) {
      if (this.state?.scrollToFirstError) {
        this.scrollToFirstError(errors);
      }
      return;
    }
    return await this.submitForm();
  }
  async validateField(fieldName, opts) {
    const form = await this.getForm();
    const validateResult = await form.validateField(fieldName, opts);
    if (Object.keys(validateResult?.errors ?? {}).length > 0) {
      console.error("validate error", validateResult?.errors);
      if (this.state?.scrollToFirstError) {
        this.scrollToFirstError(fieldName);
      }
    }
    return validateResult;
  }
  async getForm() {
    if (!this.isMounted) {
      await this.stateHandler.waitForCondition();
    }
    if (!this.form?.meta) {
      throw new Error("<VbenForm /> is not mounted");
    }
    return this.form;
  }
  handleMultiFields = (originValues) => {
    const arrayToStringFields = this.state?.arrayToStringFields;
    if (!arrayToStringFields || !Array.isArray(arrayToStringFields)) {
      return;
    }
    const processFields = (fields, separator = ",") => {
      this.processFields(fields, separator, originValues, (value, sep) => {
        if (Array.isArray(value)) {
          return value.join(sep);
        } else if (typeof value === "string") {
          if (value === "") {
            return [];
          }
          const escapedSeparator = sep.replaceAll(
            /[.*+?^${}()|[\]\\]/g,
            String.raw`\$&`
          );
          return value.split(new RegExp(escapedSeparator));
        } else {
          return value;
        }
      });
    };
    if (arrayToStringFields.every((item) => typeof item === "string")) {
      const lastItem = arrayToStringFields[arrayToStringFields.length - 1] || "";
      const fields = lastItem.length === 1 ? arrayToStringFields.slice(0, -1) : arrayToStringFields;
      const separator = lastItem.length === 1 ? lastItem : ",";
      processFields(fields, separator);
      return;
    }
    arrayToStringFields.forEach((fieldConfig) => {
      if (Array.isArray(fieldConfig)) {
        const [fields, separator = ","] = fieldConfig;
        if (!Array.isArray(fields)) {
          console.warn(
            `Invalid field configuration: fields should be an array of strings, got ${typeof fields}`
          );
          return;
        }
        processFields(fields, separator);
      }
    });
  };
  handleRangeTimeValue = (originValues) => {
    const values = { ...originValues };
    const fieldMappingTime = this.state?.fieldMappingTime;
    this.handleMultiFields(values);
    if (!fieldMappingTime || !Array.isArray(fieldMappingTime)) {
      return values;
    }
    fieldMappingTime.forEach(
      ([field, [startTimeKey, endTimeKey], format = "YYYY-MM-DD"]) => {
        if (startTimeKey && endTimeKey && values[field] === null) {
          Reflect.deleteProperty(values, startTimeKey);
          Reflect.deleteProperty(values, endTimeKey);
        }
        if (!values[field]) {
          Reflect.deleteProperty(values, field);
          return;
        }
        const [startTime, endTime] = values[field];
        if (format === null) {
          values[startTimeKey] = startTime;
          values[endTimeKey] = endTime;
        } else if (isFunction(format)) {
          values[startTimeKey] = format(startTime, startTimeKey);
          values[endTimeKey] = format(endTime, endTimeKey);
        } else {
          const [startTimeFormat, endTimeFormat] = Array.isArray(format) ? format : [format, format];
          values[startTimeKey] = startTime ? formatDate(startTime, startTimeFormat) : void 0;
          values[endTimeKey] = endTime ? formatDate(endTime, endTimeFormat) : void 0;
        }
        Reflect.deleteProperty(values, field);
      }
    );
    return values;
  };
  processFields = (fields, separator, originValues, transformFn) => {
    fields.forEach((field) => {
      const value = originValues[field];
      if (value === void 0 || value === null) {
        return;
      }
      originValues[field] = transformFn(value, separator);
    });
  };
  updateState() {
    const currentSchema = this.state?.schema ?? [];
    const prevSchema = this.prevState?.schema ?? [];
    if (currentSchema.length < prevSchema.length) {
      const currentFields = new Set(
        currentSchema.map((item) => item.fieldName)
      );
      const deletedSchema = prevSchema.filter(
        (item) => !currentFields.has(item.fieldName)
      );
      for (const schema of deletedSchema) {
        this.form?.setFieldValue?.(schema.fieldName, void 0);
      }
    }
  }
}
