// src/components/user-avatar.tsx

import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { cn } from "@/lib/utils"; // caso esteja usando essa função utilitária

function getPastelColorFromString(input: string): string {
  input = input.toLowerCase();
  let hash = 0;
  for (let i = 0; i < input.length; i++) {
    hash = input.charCodeAt(i) + ((hash << 5) - hash);
  }
  const hue = hash % 360;
  return `hsl(${hue}, 40%, 80%)`;
}

interface UserAvatarProps {
  name: string;
  username: string;
  className?: string;
}

export function UserAvatar({ name, username, className }: UserAvatarProps) {
  const initials = name
    .split(" ")
    .map((n) => n[0])
    .join("")
    .slice(0, 2)
    .toUpperCase();

  const color = getPastelColorFromString(username);

  return (
    <Avatar className={cn("h-10 w-10", className)}>
      <AvatarFallback
        className="text-sm font-semibold text-black"
        style={{ backgroundColor: color }}
      >
        {initials}
      </AvatarFallback>
    </Avatar>
  );
}
