import { createContext, Dispatch, SetStateAction } from "react";
import { ThemeType, themeNames } from "../theme";

type ViewContextProps = {
  themeName: ThemeType;
  setThemeName: Dispatch<SetStateAction<ThemeType>>;

  createBoardActive: boolean;
  setCreateBoardActive: Dispatch<SetStateAction<boolean>>;
  createSectionActive: boolean;
  setCreateSectionActive: Dispatch<SetStateAction<boolean>>;

  currentBoardId: number;
  setCurrentBoardId: Dispatch<SetStateAction<number>>;
};

const defaultProps = {
  themeName: themeNames[0],
  setThemeName: () => {},

  createBoardActive: false,
  setCreateBoardActive: () => {},
  createSectionActive: false,
  setCreateSectionActive: () => {},

  currentBoardId: 0,
  setCurrentBoardId: () => {}
};

export const ViewContext = createContext<ViewContextProps>(defaultProps);
