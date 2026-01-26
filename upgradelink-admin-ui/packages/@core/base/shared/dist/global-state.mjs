class GlobalShareState {
  #components = {};
  #message = {};
  /**
   * 定义框架内部各个场景的消息提示
   */
  defineMessage({ copyPreferencesSuccess }) {
    this.#message = {
      copyPreferencesSuccess
    };
  }
  getComponents() {
    return this.#components;
  }
  getMessage() {
    return this.#message;
  }
  setComponents(value) {
    this.#components = value;
  }
}
const globalShareState = new GlobalShareState();

export { globalShareState };
