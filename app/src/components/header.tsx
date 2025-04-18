import Icon from "@/assets/icon.svg";
import { ThemeToggle } from "./theme-toggle";
import { UserButton } from "./user-button";
import { useAuth } from "@/hooks/use-auth";
import { Link } from "react-router-dom";
import { UserSearchInput } from "./user-search-input";

export function Header() {
  const { user } = useAuth();

  return (
    <header className="w-full border-b border-border px-72 py-2 flex items-center justify-between gap-4 bg-background relative">
      <div className="flex items-center gap-2 flex-1 max-w-xl relative">
        <Link to="/home">
          <img src={Icon} alt="Tab Notes" className="w-10 h-10" />
        </Link>

        <UserSearchInput />
      </div>

      <div className="gap-1 flex items-center">
        <ThemeToggle />
        {user && <UserButton user={user} />}
      </div>
    </header>
  );
}
