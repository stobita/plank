import { createContext } from "react";
import { ThemeType, themeNames } from "../theme";

type ViewContextProps = {
  themeName: ThemeType;
  setThemeName: React.Dispatch<React.SetStateAction<ThemeType>>;

  createBoardActive: boolean;
  setCreateBoardActive: (b: boolean) => void;
};

const defaultProps = {
  themeName: themeNames[0],
  setThemeName: () => {},

  createBoardActive: false,
  setCreateBoardActive: () => {}
};

export const ViewContext = createContext<ViewContextProps>(defaultProps);
