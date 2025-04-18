import { api } from "@/lib/axios";

export interface UnfollowUserParams {
  username: string;
}

export async function unfollowUser({ username }: UnfollowUserParams) {
  const response = await api.post(`/users/${username}/unfollow`);

  return response.data;
}
