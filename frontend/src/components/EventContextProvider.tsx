import React, { ReactNode, useState } from "react";
import { EventContext } from "../context/eventContext";

interface Props {
  children: ReactNode;
}

type EventMessageData = {
  boardId: number;
};

export const EventContextProvider = (props: Props) => {
  const [updatedBoardIds, setUpdatedBoardIds] = useState<number[]>([]);
  const ev = new EventSource("http://localhost:8080/api/v1/sse");
  ev.addEventListener("open", e => {
    console.log("open:", e);
  });
  ev.addEventListener("error", e => {
    console.log("error:", e);
  });
  ev.addEventListener("message", e => {
    const data: EventMessageData = JSON.parse(e.data);
    setUpdatedBoardIds(prev => [...prev, data.boardId]);
  });
  return (
    <EventContext.Provider
      value={{
        updatedBoardIds,
        setUpdatedBoardIds
      }}
    >
      {props.children}
    </EventContext.Provider>
  );
};
