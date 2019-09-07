import colors from "./colors";

export const themeNames = ["light", "dark"] as const;
export type ThemeType = typeof themeNames[number];

const lightTheme = {
  primary: colors.primary,
  bg: colors.mainWhite,
  text: colors.mainBlack,
  border: colors.borderGray
};

const darkTheme = {
  primary: colors.primary,
  bg: colors.mainBlack,
  text: colors.mainWhite,
  border: colors.borderGray
};

const theme = {
  light: lightTheme,
  dark: darkTheme
};

export default theme;
