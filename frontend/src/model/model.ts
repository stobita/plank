export type Board = {
  id: number;
  name: string;
};

export const boardInit = {
  id: 0,
  name: ""
};

export type Section = {
  id: number;
  name: string;
  cards: Card[];
  board: Board;
};

export const sectionInit: Section = {
  id: 0,
  name: "",
  board: boardInit,
  cards: []
};

export type Card = {
  id: number;
  name: string;
  description: string;
  section: Section;
};

export const cardInit: Card = {
  id: 0,
  name: "",
  description: "",
  section: sectionInit
};
