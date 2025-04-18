import { api } from "@/lib/axios";

export interface UpdateProfileRequest {
  name: string;
  username: string;
}

export async function updateProfile({ name, username }: UpdateProfileRequest) {
  await api.put("/users", { name, username });
}
