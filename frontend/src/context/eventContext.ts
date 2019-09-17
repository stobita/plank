import { createContext } from "react";

type EventContextProps = {};

const defaultProps = {};

export const EventContext = createContext<EventContextProps>(defaultProps);
