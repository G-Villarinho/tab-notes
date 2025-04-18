import { api } from "@/lib/axios";

export interface GetFeedResponse {
  postId: string;
  title: string;
  content: string;
  likes: number;
  createdAt: string;
  authorName: string;
  authorUsername: string;
  likedByUser: boolean;
}

export interface GetFeedQueryParams {
  offset: number;
  limit: number;
}

export async function getFeed(params: GetFeedQueryParams) {
  const response = await api.get<GetFeedResponse[]>("/feed", {
    params,
  });

  return response.data;
}
