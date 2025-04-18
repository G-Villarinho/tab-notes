import { api } from "@/lib/axios"

export interface GetProfileResponse {
    name: string
    email: string
    username: string
    followers: number
    following: number
}

export async function getProfile() {
    const response = await api.get<GetProfileResponse>("/me");

    return response.data;
}