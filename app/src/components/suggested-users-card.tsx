import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { UserAvatar } from "@/components/user-avatar";

const suggestedUsers = [
  {
    name: "Caio Gabriel",
    username: "caio.dev",
  },
  {
    name: "Ana Monteiro",
    username: "ana.codes",
  },
  {
    name: "Pedro Lima",
    username: "pedrolima",
  },
];

export function SuggestedUsersCard() {
  return (
    <div className="rounded-xl border p-4 shadow-sm bg-card space-y-4">
      <h2 className="text-sm font-semibold text-muted-foreground">
        Você pode gostar...
      </h2>

      <ul className="space-y-3">
        {suggestedUsers.map((user) => {
          return (
            <li key={user.username} className="flex items-center gap-3">
              <UserAvatar
                name={user.name}
                username={user.username}
                className="h-10 w-10"
              />
              <div className="flex-1 text-sm leading-tight">
                <p className="font-medium text-foreground">{user.name}</p>
                <p className="text-xs text-muted-foreground">
                  @{user.username}
                </p>
              </div>
              <Button variant="outline" size="xs" className="shrink-0">
                <Plus className="w-4 h-4" />
              </Button>
            </li>
          );
        })}
      </ul>
    </div>
  );
}
