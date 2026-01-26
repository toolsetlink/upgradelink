export { TinyColor } from '@ctrl/tinycolor';

declare function isDarkColor(color: string): boolean;
declare function isLightColor(color: string): boolean;

declare function convertToHsl(color: string): string;
declare function convertToHslCssVar(color: string): string;
declare function convertToRgb(str: string): string;
declare function isValidColor(color?: string): boolean;

interface ColorItem {
    alias?: string;
    color: string;
    name: string;
}
declare function generatorColorVariables(colorItems: ColorItem[]): Record<string, string>;

export { convertToHsl, convertToHslCssVar, convertToRgb, generatorColorVariables, isDarkColor, isLightColor, isValidColor };
