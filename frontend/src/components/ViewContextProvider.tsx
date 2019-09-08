import React, { ReactNode, useEffect, useState } from "react";
import { ViewContext } from "../context/viewContext";
import { localStorageRepository } from "../localStorageRepository";
import theme, { ThemeType } from "../theme";
import { ThemeProvider } from "styled-components";
import { GlobalStyleWrapper } from "./GlobalStyleWrapper";

interface Props {
  children: ReactNode;
}

export const ViewContextProvider = (props: Props) => {
  const currentThemeName = localStorageRepository.getThemeName();
  const [themeName, setThemeName] = useState<ThemeType>(currentThemeName);
  const [currentBoardId, setCurrentBoardId] = useState(0);
  const [createBoardActive, setCreateBoardActive] = useState(false);
  const [createSectionActive, setCreateSectionActive] = useState(false);

  useEffect(() => {
    localStorageRepository.setThemeName(themeName);
  }, [themeName]);

  return (
    <ThemeProvider theme={theme[themeName]}>
      <ViewContext.Provider
        value={{
          themeName,
          setThemeName,
          currentBoardId,
          setCurrentBoardId,
          createBoardActive,
          setCreateBoardActive,
          createSectionActive,
          setCreateSectionActive
        }}
      >
        <GlobalStyleWrapper />
        {props.children}
      </ViewContext.Provider>
    </ThemeProvider>
  );
};
