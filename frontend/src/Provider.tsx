import React, { ReactNode, useEffect, useState } from "react";
import { ThemeType } from "./theme";
import { ThemeProvider } from "styled-components";
import theme from "./theme";
import { localStorageRepository } from "./localStorageRepository";
import { Board } from "./model/model";

interface Props {
  children: ReactNode;
}

type MainContextType = {
  themeName: ThemeType;
  setThemeName: React.Dispatch<React.SetStateAction<ThemeType>>;
};

export const MainContext = React.createContext<MainContextType>({
  themeName: "light",
  setThemeName: () => {}
});

type DataContextType = {
  boards: Board[];
  setBoards: React.Dispatch<React.SetStateAction<Board[]>>;
};
export const DataContext = React.createContext<DataContextType>({
  boards: [],
  setBoards: () => {}
});

type ModalContextType = {
  createBoardActive: boolean;
  setCreateBoardActive: (b: boolean) => void;
};

export const ModalContext = React.createContext<ModalContextType>({
  createBoardActive: false,
  setCreateBoardActive: () => {}
});

export default (props: Props) => {
  const currentThemeName = localStorageRepository.getThemeName();
  const [themeName, setThemeName] = React.useState<ThemeType>(currentThemeName);
  const [createBoardActive, setCreateBoardActive] = React.useState(false);
  const [boards, setBoards] = useState<Board[]>([]);

  useEffect(() => {
    localStorageRepository.setThemeName(themeName);
  }, [themeName]);

  return (
    <MainContext.Provider value={{ themeName, setThemeName }}>
      <ModalContext.Provider
        value={{
          createBoardActive,
          setCreateBoardActive
        }}
      >
        <ThemeProvider theme={theme[themeName]}>
          <DataContext.Provider value={{ boards, setBoards }}>
            <div>{props.children}</div>
          </DataContext.Provider>
        </ThemeProvider>
      </ModalContext.Provider>
    </MainContext.Provider>
  );
};
