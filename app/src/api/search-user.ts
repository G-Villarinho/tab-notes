import { api } from "@/lib/axios";

export interface SearchUserQueryParams {
  q: string;
}

export interface SearchUserResponse {
  name: string;
  username: string;
}

export async function searchUser({
  q,
}: SearchUserQueryParams): Promise<SearchUserResponse[]> {
  const response = await api.get<SearchUserResponse[]>("/users", {
    params: { q },
  });

  return response.data;
}
