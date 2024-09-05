import { createContext, Dispatch, SetStateAction } from "react";
import { ThemeType, themeNames } from "../theme";
import { Board } from "../model/model";

type ViewContextProps = {
  themeName: ThemeType;
  setThemeName: Dispatch<SetStateAction<ThemeType>>;

  createBoardActive: boolean;
  setCreateBoardActive: Dispatch<SetStateAction<boolean>>;

  boardSettingActive: boolean;
  setBoardSettingActive: Dispatch<SetStateAction<boolean>>;

  createSectionActive: boolean;
  setCreateSectionActive: Dispatch<SetStateAction<boolean>>;

  currentBoard: Board;
  setCurrentBoard: Dispatch<SetStateAction<Board>>;

  sectionFilterActive: boolean;
  setSectionFilterActive: Dispatch<SetStateAction<boolean>>;
};

const defaultProps = {
  themeName: themeNames[0],
  setThemeName: () => {},

  createBoardActive: false,
  setCreateBoardActive: () => {},

  boardSettingActive: false,
  setBoardSettingActive: () => {},

  createSectionActive: false,
  setCreateSectionActive: () => {},

  currentBoard: { id: 0, name: "" },
  setCurrentBoard: () => {},

  sectionFilterActive: false,
  setSectionFilterActive: () => {},
};

export const ViewContext = createContext<ViewContextProps>(defaultProps);
