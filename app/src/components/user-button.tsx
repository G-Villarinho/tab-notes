import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { LogOut, User } from "lucide-react";

import { useMutation } from "@tanstack/react-query";
import { isAxiosError } from "axios";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { logout } from "@/api/logout";
import { UserAvatar } from "./user-avatar";
import { useAuth } from "@/hooks/use-auth";

interface UserMenuProps {
  user: {
    name: string;
    email: string;
    username: string;
  };
}

export function UserButton({ user }: UserMenuProps) {
  const navigate = useNavigate();
  const { setUser } = useAuth();

  const { mutateAsync: logoutFn } = useMutation({ mutationFn: logout });

  async function handleLogout() {
    try {
      await logoutFn();
      setUser(null);
      navigate("/login");
    } catch (error) {
      if (isAxiosError(error)) {
        toast.error("Erro ao sair da conta");
      }
    }
  }

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <button className="flex items-center gap-2 rounded-md px-2 py-1 hover:bg-muted transition focus:outline-none">
          <UserAvatar
            name={user.name}
            username={user.username}
            className="h-9 w-9"
          />
        </button>
      </DropdownMenuTrigger>

      <DropdownMenuContent align="end" className="w-64">
        <DropdownMenuLabel className="flex flex-col gap-0.5">
          <span className="text-sm font-medium text-foreground truncate">
            {user.name}
          </span>
          <span className="text-xs text-muted-foreground truncate">
            @{user.username}
          </span>
          <span className="text-[11px] text-muted-foreground/70 font-mono truncate">
            {user.email}
          </span>
        </DropdownMenuLabel>

        <DropdownMenuSeparator />

        <DropdownMenuGroup>
          <DropdownMenuItem className="cursor-pointer">
            <User className="mr-2 h-4 w-4 text-muted-foreground" />
            <span>Meu perfil</span>
          </DropdownMenuItem>
        </DropdownMenuGroup>

        <DropdownMenuSeparator />

        <DropdownMenuItem
          onClick={handleLogout}
          className="text-red-500 cursor-pointer focus:text-red-600"
        >
          <LogOut className="mr-2 h-4 w-4" />
          <span>Sair</span>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
