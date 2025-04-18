import { api } from "@/lib/axios";

export interface RegisterUserRequest {
    username: string;
    name: string;
    email: string;
}

export async function registerUser({ username, name, email }: RegisterUserRequest) {
    await api.post("/register", {
        username,
        name,
        email,
    });
}