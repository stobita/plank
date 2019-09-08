import React, { ReactNode } from "react";
import { ViewContextProvider } from "./components/ViewContextProvider";
import { DataContextProvider } from "./components/DataContextProvider";

interface Props {
  children: ReactNode;
}

export default (props: Props) => {
  return (
    <ViewContextProvider>
      <DataContextProvider>{props.children}</DataContextProvider>
    </ViewContextProvider>
  );
};
