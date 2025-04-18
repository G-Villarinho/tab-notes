import { useQuery } from "@tanstack/react-query";
import { api } from "@/lib/axios";
import { useDebounce } from "./use-debounce";

interface SearchUser {
  name: string;
  username: string;
}

export function useUsersSearch(query: string) {
  const debounced = useDebounce(query);

  const { data, isLoading } = useQuery<SearchUser[]>({
    queryKey: ["search-users", debounced],
    queryFn: async () => {
      const response = await api.get("/users?q=" + debounced);
      return response.data;
    },
    enabled: debounced.length > 1,
  });

  return { users: data ?? [], isLoading };
}
