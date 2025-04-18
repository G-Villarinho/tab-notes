import { api } from "@/lib/axios";

export async function logout() {
    await api.post("/logout")
}