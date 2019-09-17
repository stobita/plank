import { createContext, Dispatch, SetStateAction } from "react";

type EventContextProps = {
  updatedBoardIds: number[];
  setUpdatedBoardIds: Dispatch<SetStateAction<number[]>>;
};

const defaultProps = {
  updatedBoardIds: [],
  setUpdatedBoardIds: () => {}
};

export const EventContext = createContext<EventContextProps>(defaultProps);
