import { TinyColor } from '@ctrl/tinycolor';
export { TinyColor } from '@ctrl/tinycolor';
import { getColors } from 'theme-colors';

function isDarkColor(color) {
  return new TinyColor(color).isDark();
}
function isLightColor(color) {
  return new TinyColor(color).isLight();
}

function convertToHsl(color) {
  const { a, h, l, s } = new TinyColor(color).toHsl();
  const hsl = `hsl(${Math.round(h)} ${Math.round(s * 100)}% ${Math.round(l * 100)}%)`;
  return a < 1 ? `${hsl} ${a}` : hsl;
}
function convertToHslCssVar(color) {
  const { a, h, l, s } = new TinyColor(color).toHsl();
  const hsl = `${Math.round(h)} ${Math.round(s * 100)}% ${Math.round(l * 100)}%`;
  return a < 1 ? `${hsl} / ${a}` : hsl;
}
function convertToRgb(str) {
  return new TinyColor(str.replaceAll(/deg|grad|rad|turn/g, "")).toRgbString();
}
function isValidColor(color) {
  if (!color) {
    return false;
  }
  return new TinyColor(color).isValid;
}

function generatorColorVariables(colorItems) {
  const colorVariables = {};
  colorItems.forEach(({ alias, color, name }) => {
    if (color) {
      const colorsMap = getColors(new TinyColor(color).toHexString());
      let mainColor = colorsMap["500"];
      const colorKeys = Object.keys(colorsMap);
      colorKeys.forEach((key) => {
        const colorValue = colorsMap[key];
        if (colorValue) {
          const hslColor = convertToHslCssVar(colorValue);
          colorVariables[`--${name}-${key}`] = hslColor;
          if (alias) {
            colorVariables[`--${alias}-${key}`] = hslColor;
          }
          if (key === "500") {
            mainColor = hslColor;
          }
        }
      });
      if (alias && mainColor) {
        colorVariables[`--${alias}`] = mainColor;
      }
    }
  });
  return colorVariables;
}

export { convertToHsl, convertToHslCssVar, convertToRgb, generatorColorVariables, isDarkColor, isLightColor, isValidColor };
