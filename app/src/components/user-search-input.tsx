// src/components/user-search-input.tsx
import { useState } from "react";
import { Link } from "react-router-dom";
import { useUsersSearch } from "@/hooks/use-users-search";
import { Input } from "./ui/input";
import { UserAvatar } from "./user-avatar";

interface UserSearchInputProps {
  placeholder?: string;
  onSelect?: () => void;
}

export function UserSearchInput({
  placeholder = "Buscar usuários...",
  onSelect,
}: UserSearchInputProps) {
  const [query, setQuery] = useState("");
  const { users, isLoading } = useUsersSearch(query);
  const showDropdown = query.length > 1 && (isLoading || users.length > 0);

  return (
    <div className="w-full relative">
      <Input
        type="search"
        placeholder={placeholder}
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        className="w-full"
      />

      {showDropdown && (
        <div className="absolute z-50 mt-1 w-full rounded-md border bg-popover shadow-md max-h-60 overflow-y-auto">
          {isLoading ? (
            <p className="text-sm p-2 text-muted-foreground">Carregando...</p>
          ) : users.length === 0 ? (
            <p className="text-sm p-2 text-muted-foreground">
              Nenhum usuário encontrado
            </p>
          ) : (
            users.map((user) => (
              <Link
                to={`/${user.username}`}
                key={user.username}
                className="flex items-center gap-3 px-3 py-2 text-sm hover:bg-muted transition"
                onClick={() => {
                  setQuery("");
                  onSelect?.();
                }}
              >
                <UserAvatar
                  name={user.name}
                  username={user.username}
                  className="h-8 w-8"
                />
                <div>
                  <p className="font-medium">{user.name}</p>
                  <span className="text-muted-foreground text-xs">
                    @{user.username}
                  </span>
                </div>
              </Link>
            ))
          )}
        </div>
      )}
    </div>
  );
}
