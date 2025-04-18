import { AuthContext } from "@/contexts/auth/context";
import { useContext } from "react";

export function useAuth() {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error("useAuth must be used within a UserAuthProvider");
  }

  return context;
}
