import repository, { Collection } from "./repository";
import { Board } from "../model/model";

const resource = "/boards";

export type CreateBoardPayload = {
  title: string;
};

export default {
  async getBoards(): Promise<Board[]> {
    const res = await repository
      .get<Collection<Board>>(`${resource}`)
      .catch(e => {
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
  }
};
