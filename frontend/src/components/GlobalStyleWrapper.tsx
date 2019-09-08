import React, { useContext } from "react";
import { createGlobalStyle } from "styled-components";
import { ViewContext } from "../context/viewContext";

export const GlobalStyleWrapper = () => {
  const { createBoardActive } = useContext(ViewContext);
  const GlobalStyle = createGlobalStyle<{ noScroll: boolean }>`
    body {
      background: ${props => props.theme.bg};
      overflow: ${props => (props.noScroll ? "hidden" : "auto")}
    }
    ::-webkit-scrollbar {
      width: 10px;
      height: 10px;
    }
    ::-webkit-scrollbar-track{
      background: ${props => props.theme.bg};
    }
    ::-webkit-scrollbar-thumb{
      background: ${props => props.theme.main};
      box-shadow: none;
    }
    ::-webkit-scrollbar-corner{
      display:none;
    }
  `;
  return <GlobalStyle noScroll={createBoardActive} />;
};
