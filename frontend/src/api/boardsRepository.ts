import repository, { Collection } from "./repository";
import { Board, Section, Label } from "../model/model";

const resource = "/boards";
const childResouce = {
  sections: "sections",
  rabels: "labels",
};

export type CreateBoardPayload = {
  name: string;
};

export type UpdateBoardPayload = {
  name: string;
};

export type CreateSectionPayload = {
  name: string;
};

export type UpdateSectionPayload = {
  name: string;
};

export default {
  async getBoards(): Promise<Board[]> {
    const res = await repository
      .get<Collection<Board>>(`${resource}`)
      .catch((e) => {
        console.error(e);
        return null;
      });
    if (res === null || !res.data.items) {
      return [];
    }
    return res.data.items;
  },
  async createBoard(payload: CreateBoardPayload): Promise<Board> {
    const res = await repository.post<Board>(`${resource}`, payload);
    return res.data;
  },
  async updateBoard(
    boardId: number,
    payload: UpdateBoardPayload
  ): Promise<Board> {
    const res = await repository.put(`${resource}/${boardId}`, payload);
    return res.data;
  },
  async deleteBoard(boardId: number): Promise<void> {
    await repository.delete(`${resource}/${boardId}`);
    return;
  },
  async getBoardSections(boardId: number): Promise<Section[]> {
    const res = await repository
      .get<Collection<Section>>(
        `${resource}/${boardId}/${childResouce.sections}`
      )
      .catch((e) => {
        console.error(e);
        return null;
      });
    if (res === null || !res.data.items) {
      return [];
    }
    return res.data.items;
  },
  async getLabelSections(boardId: number, labelId: number): Promise<Section[]> {
    const res = await repository
      .get<Collection<Section>>(
        `${resource}/${boardId}/${childResouce.rabels}/${labelId}/${childResouce.sections}`
      )
      .catch((e) => {
        console.error(e);
        return null;
      });
    if (res === null || !res.data.items) {
      return [];
    }
    return res.data.items;
  },
  async createSection(
    boardId: number,
    payload: CreateSectionPayload
  ): Promise<Section> {
    const res = await repository.post<Section>(
      `${resource}/${boardId}/${childResouce.sections}`,
      payload
    );
    return res.data;
  },
  async updateSection(
    boardId: number,
    sectionId: number,
    payload: UpdateSectionPayload
  ): Promise<Section> {
    const res = await repository.put(
      `${resource}/${boardId}/${childResouce.sections}/${sectionId}`,
      payload
    );
    return res.data;
  },
  async deleteSection(boardId: number, sectionId: number): Promise<void> {
    await repository.delete(
      `${resource}/${boardId}/${childResouce.sections}/${sectionId}`
    );
    return;
  },
  async reorderSection(
    boardId: number,
    sectionId: number,
    position: number
  ): Promise<void> {
    await repository.put(
      `${resource}/${boardId}/${childResouce.sections}/${sectionId}/reorder`,
      {
        position,
      }
    );
  },
  async getLabels(boardId: number): Promise<Label[]> {
    const res = await repository
      .get<Collection<Label>>(`${resource}/${boardId}/${childResouce.rabels}`)
      .catch((e) => {
        console.log(e);
        return null;
      });
    if (res == null || !res.data.items) {
      return [];
    }
    return res.data.items;
  },
};
