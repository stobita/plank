import axios from "axios";

const baseDomain = "http://localhost:8080";
const baseURL = `${baseDomain}/api/v1`;

export type Collection<T> = {
  items: T[];
};

export default axios.create({
  baseURL
});
