import { api } from "@/lib/axios";

export interface GetUserSessionsResponse {
  id: string;
  verifiedAt: string;
  currentSessionId: string;
  RevokedAt?: string | null;
  createdAt: string;
}

export async function getUserSessions() {
  const response = await api.get<GetUserSessionsResponse[]>("/me/sessions");

  return response.data;
}
