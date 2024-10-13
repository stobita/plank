import React, { ReactNode } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import { ViewContextProvider } from "./ViewContextProvider";
import { DataContextProvider } from "./DataContextProvider";
import { EventContextProvider } from "./EventContextProvider";

interface Props {
  children: ReactNode;
}

export const Provider = (props: Props) => {
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
