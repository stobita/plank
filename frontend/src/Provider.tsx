import React, { ReactNode } from "react";
import { DndProvider } from "react-dnd";
import HTML5Backend from "react-dnd-html5-backend";
import { ViewContextProvider } from "./components/ViewContextProvider";
import { DataContextProvider } from "./components/DataContextProvider";
import { EventContextProvider } from "./components/EventContextProvider";
import { MoveContextProvider } from "./components/MoveContextProvider";

interface Props {
  children: ReactNode;
}

export default (props: Props) => {
  return (
    <DataContextProvider>
      <ViewContextProvider>
        <MoveContextProvider>
          <EventContextProvider>
            <DndProvider backend={HTML5Backend}>{props.children}</DndProvider>
          </EventContextProvider>
        </MoveContextProvider>
      </ViewContextProvider>
    </DataContextProvider>
  );
};
