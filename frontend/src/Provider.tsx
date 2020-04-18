import React, { ReactNode } from "react";
import { DndProvider } from "react-dnd";
import HTML5Backend from "react-dnd-html5-backend";
import { ViewContextProvider } from "./components/ViewContextProvider";
import { DataContextProvider } from "./components/DataContextProvider";
import { EventContextProvider } from "./components/EventContextProvider";

interface Props {
  children: ReactNode;
}

export default (props: Props) => {
  return (
    <DataContextProvider>
      <ViewContextProvider>
        <EventContextProvider>
          <DndProvider backend={HTML5Backend}>{props.children}</DndProvider>
        </EventContextProvider>
      </ViewContextProvider>
    </DataContextProvider>
  );
};
