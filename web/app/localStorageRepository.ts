'use client'

import { ThemeType, themeNames } from "../theme";
const themeKey = "plank.theme";

export const localStorageRepository = {
  getThemeName: (): ThemeType => {
    try {
      const value = localStorage.getItem(themeKey);
      const guard = (v: string): v is ThemeType => {
        return themeNames.some(name => name === v);
      };
      if (value !== null && guard(value)) {
        return value;
      } else {
        return "light";
      }
    } catch (e) {
      console.error(e);
      return "light";
    }
  },
  setThemeName: (name: ThemeType) => {
    localStorage.setItem(themeKey, name);
  }
};
