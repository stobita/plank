import { createContext } from "react";
import { Board } from "../model/model";

type DataContextProps = {
  boards: Board[];
  setBoards: React.Dispatch<React.SetStateAction<Board[]>>;
};

const defaultProps = {
  boards: [],
  setBoards: () => {}
};

export const DataContext = createContext<DataContextProps>(defaultProps);
