import colors from "./colors";

export const themeNames = ["light", "dark"] as const;
export type ThemeType = typeof themeNames[number];

const lightTheme = {
  primary: colors.primary,
  main: colors.mainWhite,
  bg: colors.thinGray,
  text: colors.mainBlack,
  border: colors.borderGray
};

const darkTheme = {
  primary: colors.primary,
  main: colors.mainBlack,
  bg: colors.mainGray,
  text: colors.mainWhite,
  border: colors.borderGray
};

const theme = {
  light: lightTheme,
  dark: darkTheme
};

export default theme;
