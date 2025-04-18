import { useEffect, useState, useCallback } from "react";
import { getProfile, GetProfileResponse } from "@/api/get-profile";
import { AuthContext } from "./context";

interface AuthProviderProps {
  children: React.ReactNode;
}

export function AuthProvider({ children }: AuthProviderProps) {
  const [user, setUser] = useState<GetProfileResponse | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  const loadUserProfile = useCallback(async () => {
    setIsLoading(true);
    try {
      const profile = await getProfile();
      setUser(profile);
    } catch {
      setUser(null);
    } finally {
      setIsLoading(false);
    }
  }, []);

  useEffect(() => {
    loadUserProfile();
  }, [loadUserProfile]);

  const isAuthenticated = Boolean(user);

  return (
    <AuthContext.Provider
      value={{
        user: user ?? null,
        isLoading,
        isAuthenticated,
        setUser: (u) => setUser(u ?? null),
        refetch: loadUserProfile,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
