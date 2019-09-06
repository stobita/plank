import colors from "./colors";

export type ThemeType = "light" | "dark";

const lightTheme = {
  bg: colors.mainWhite,
  text: colors.mainBlack,
  border: colors.borderGray
};

const darkTheme = {
  bg: colors.mainBlack,
  text: colors.mainWhite,
  border: colors.borderGray
};

const theme = {
  light: lightTheme,
  dark: darkTheme
};

export default theme;
