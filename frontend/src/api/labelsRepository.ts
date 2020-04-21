import repository, { Collection } from "./repository";
import { Label, Section } from "../model/model";

const resource = "/labels";

export default {
  async getLabels(): Promise<Label[]> {
    const res = await repository
      .get<Collection<Label>>(`${resource}`)
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
