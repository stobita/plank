import React, { ReactNode } from "react";
import { ViewContextProvider } from "./components/ViewContextProvider";
import { DataContextProvider } from "./components/DataContextProvider";
import { EventContextProvider } from "./components/EventContextProvider";

interface Props {
  children: ReactNode;
}

export default (props: Props) => {
  return (
    <ViewContextProvider>
      <EventContextProvider>
        <DataContextProvider>{props.children}</DataContextProvider>
      </EventContextProvider>
    </ViewContextProvider>
  );
};
