import React, { ReactNode, useEffect, useState } from "react";
import { ThemeType } from "./theme";
import { ThemeProvider } from "styled-components";
import theme from "./theme";
import { localStorageRepository } from "./localStorageRepository";
import { Board } from "./model/model";
import { DataContext } from "./context/dataContext";
import { ViewContext } from "./context/viewContext";

interface Props {
  children: ReactNode;
}

export default (props: Props) => {
  const currentThemeName = localStorageRepository.getThemeName();
  const [themeName, setThemeName] = React.useState<ThemeType>(currentThemeName);
  const [createBoardActive, setCreateBoardActive] = React.useState(false);
  const [boards, setBoards] = useState<Board[]>([]);

  useEffect(() => {
    localStorageRepository.setThemeName(themeName);
  }, [themeName]);

  return (
    <ViewContext.Provider
      value={{
        themeName,
        setThemeName,
        createBoardActive,
        setCreateBoardActive
      }}
    >
      <ThemeProvider theme={theme[themeName]}>
        <DataContext.Provider value={{ boards, setBoards }}>
          <div>{props.children}</div>
        </DataContext.Provider>
      </ThemeProvider>
    </ViewContext.Provider>
  );
};
