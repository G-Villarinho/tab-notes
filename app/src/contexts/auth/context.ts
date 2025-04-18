import { createContext } from "react";
import { GetProfileResponse } from "@/api/get-profile";

interface AuthContextProps {
  user: GetProfileResponse | null;
  isLoading: boolean;
  isAuthenticated: boolean;
  setUser: (user: GetProfileResponse | null) => void;
  refetch: () => Promise<void>;
}

export const AuthContext = createContext<AuthContextProps | undefined>(
  undefined
);
