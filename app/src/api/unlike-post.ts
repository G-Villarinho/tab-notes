import { api } from "@/lib/axios";

export interface UnlikePostParams {
  postId: string;
}

export async function unlikePost({ postId }: UnlikePostParams) {
  await api.post(`/posts/${postId}/unlike`);
}