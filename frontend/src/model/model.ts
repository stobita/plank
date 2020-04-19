export type Board = {
  id: number;
  name: string;
};

export const boardInit = {
  id: 0,
  name: "",
};

export type Section = {
  id: number;
  name: string;
  cards: Card[];
  board: Board;
  position: number;
};

export const sectionInit: Section = {
  id: 0,
  name: "",
  board: boardInit,
  cards: [],
  position: 0,
};

export type Card = {
  id: number;
  name: string;
  description: string;
  section: Section;
  position?: number;
  labels: Label[];
};

export type Label = {
  id: number;
  name: string;
};

export const cardInit: Card = {
  id: 0,
  name: "",
  description: "",
  section: sectionInit,
  labels: [],
};
