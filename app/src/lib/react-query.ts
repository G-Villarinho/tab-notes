import { QueryClient } from "@tanstack/react-query";
import toast from "react-hot-toast";

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry(failureCount) {
        if (failureCount >= 3) {
          toast.error(
            "O aplicativo está demorando mais do que o esperado para carregar. Por favor, tente novamente em alguns minutos.",
            {
              id: "network-error",
            }
          );
          return false;
        }
        return true;
      },
    },
  },
});