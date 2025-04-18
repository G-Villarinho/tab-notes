import { api } from "@/lib/axios";

export interface FollowUserParams {
  username: string;
}

export async function followUser({ username }: FollowUserParams) {
  const response = await api.post(`/users/${username}/follow`);

  return response.data;
}
