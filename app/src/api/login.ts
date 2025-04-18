import { api } from "@/lib/axios"

export interface LoginRequest {
    email: string
}

export async function login({ email }: LoginRequest) {
    await api.post("/authenticate", { email })
}