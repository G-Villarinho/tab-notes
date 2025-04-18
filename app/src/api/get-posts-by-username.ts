import { api } from "@/lib/axios";
import { PostResponse } from "./responses/post";

export interface GetPostsByUsernameParams {
  username?: string;
}

export async function getPostsByUsername({
  username,
}: GetPostsByUsernameParams) {
  const response = await api.get<PostResponse[]>(`/users/${username}/posts`);

  return response.data;
}
