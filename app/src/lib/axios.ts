import axios from "axios";
import camelcaseKeys from "camelcase-keys";
import { env } from "@/env";

export const api = axios.create({
  baseURL: env.VITE_API_URL,
  withCredentials: true,
});

if (env.VITE_ENABLE_API_DELAY) {
  api.interceptors.request.use(async (config) => {
    await new Promise((resolve) =>
      setTimeout(resolve, Math.round(Math.random() * 1000))
    );
    return config;
  });
}

api.interceptors.response.use((response) => {
  const contentType = response.headers["content-type"];

  if (
    contentType?.includes("application/json") &&
    response.data &&
    typeof response.data === "object"
  ) {
    response.data = camelcaseKeys(response.data, { deep: true });
  }

  return response;
});
