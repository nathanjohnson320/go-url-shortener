import { mande } from "mande";

export type Url = {
  id: number;
  long_url: string;
  short_url: string;
};

const urls = mande("http://localhost:8080/urls");

export function createUrl(long_url: string) {
  return urls.post<Url>({ long_url });
}

export function deleteUrl(id: number) {
  return urls.delete(id);
}

export function listUrls() {
  return urls.get<Url[]>();
}
