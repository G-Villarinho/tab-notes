import { api } from "@/lib/axios";

export interface CreatePostRequest {
    title: string;
    content: string;
}

export interface CreatePostResponse {
    id: string
    title: string;
    content: string;
    Likes: number;
    createdAt: string;
}

export async function createPost(data: CreatePostRequest) {
    const response = await api.post<CreatePostResponse>("/posts", data);

    return response.data;
}