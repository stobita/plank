import React, { ReactNode, useEffect, useState } from "react";
import { ViewContext } from "../context/viewContext";
import { localStorageRepository } from "../localStorageRepository";
import theme, { ThemeType } from "../theme";
import { ThemeProvider } from "styled-components";
import { GlobalStyleWrapper } from "./GlobalStyleWrapper";
import { Board } from "../model/model";

interface Props {
  children: ReactNode;
}

export const ViewContextProvider = (props: Props) => {
  const currentThemeName = localStorageRepository.getThemeName();
  const [themeName, setThemeName] = useState<ThemeType>(currentThemeName);
  const [currentBoard, setCurrentBoard] = useState<Board>({ id: 0, name: "" });
  const [createBoardActive, setCreateBoardActive] = useState(false);
  const [boardSettingActive, setBoardSettingActive] = useState(false);
  const [createSectionActive, setCreateSectionActive] = useState(false);
  const [sectionFilterActive, setSectionFilterActive] = useState(false);

  useEffect(() => {
    localStorageRepository.setThemeName(themeName);
  }, [themeName]);

  useEffect(() => {
    setCreateBoardActive(false);
    setBoardSettingActive(false);
  }, [currentBoard]);

  useEffect(() => {
    if (boardSettingActive && createBoardActive) {
      setBoardSettingActive(false);
    }
  }, [createBoardActive, boardSettingActive]);

  useEffect(() => {
    if (createBoardActive && boardSettingActive) {
      setCreateBoardActive(false);
    }
  }, [boardSettingActive, createBoardActive]);

  return (
    <ThemeProvider theme={theme[themeName]}>
      <ViewContext.Provider
        value={{
          themeName,
          setThemeName,
          currentBoard,
          setCurrentBoard,
          createBoardActive,
          setCreateBoardActive,
          boardSettingActive,
          setBoardSettingActive,
          createSectionActive,
          setCreateSectionActive,
          sectionFilterActive,
          setSectionFilterActive,
        }}
      >
        <GlobalStyleWrapper />
        {props.children}
      </ViewContext.Provider>
    </ThemeProvider>
  );
};
