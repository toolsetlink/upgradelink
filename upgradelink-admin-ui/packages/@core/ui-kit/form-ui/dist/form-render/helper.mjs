import { isObject, isString } from "@vben-core/shared/utils";
export function getBaseRules(schema) {
  if (!schema || isString(schema)) return null;
  if ("innerType" in schema._def)
    return getBaseRules(schema._def.innerType);
  if ("schema" in schema._def)
    return getBaseRules(schema._def.schema);
  return schema;
}
export function getDefaultValueInZodStack(schema) {
  if (!schema || isString(schema)) {
    return;
  }
  const typedSchema = schema;
  if (typedSchema._def.typeName === "ZodDefault")
    return typedSchema._def.defaultValue();
  if ("innerType" in typedSchema._def) {
    return getDefaultValueInZodStack(
      typedSchema._def.innerType
    );
  }
  if ("schema" in typedSchema._def) {
    return getDefaultValueInZodStack(
      typedSchema._def.schema
    );
  }
  return void 0;
}
export function isEventObjectLike(obj) {
  if (!obj || !isObject(obj)) {
    return false;
  }
  return Reflect.has(obj, "target") && Reflect.has(obj, "stopPropagation");
}
