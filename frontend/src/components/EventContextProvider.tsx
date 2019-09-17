import React, { ReactNode } from "react";
import { EventContext } from "../context/eventContext";

interface Props {
  children: ReactNode;
}

export const EventContextProvider = (props: Props) => {
  const ev = new EventSource("http://localhost:8080/api/v1/sse");
  ev.addEventListener("open", e => {
    console.log("open:", e);
  });
  ev.addEventListener("error", e => {
    console.log("error:", e);
  });
  ev.addEventListener("message", e => {
    console.log("message");
    console.log(e);
  });
  return (
    <EventContext.Provider value={{}}>{props.children}</EventContext.Provider>
  );
};
