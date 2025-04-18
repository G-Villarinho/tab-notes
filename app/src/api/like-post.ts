import { api } from "@/lib/axios";

export interface LikePostParams {
    postId: string;
}

export async function likePost({ postId }: LikePostParams) {
    await api.post(`/posts/${postId}/like`);
}