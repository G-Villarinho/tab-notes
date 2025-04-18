import { api } from "@/lib/axios";

export interface GetProfileByUsernameParams {
  username?: string;
}

export interface GetProfileByUsername {
  name: string;
  username: string;
  followers: number;
  following: number;
  followedByMe: boolean;
  followingMe: boolean;
}

export async function getProfileByUsername({
  username,
}: GetProfileByUsernameParams) {
  const response = await api.get<GetProfileByUsername>(`/users/${username}`);

  return response.data;
}
