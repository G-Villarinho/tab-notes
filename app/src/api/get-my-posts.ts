import { api } from "@/lib/axios";
import { PostResponse } from "./responses/post";

export async function getMyPosts() {
  const response = await api.get<PostResponse[]>("/me/posts");

  return response.data;
}
