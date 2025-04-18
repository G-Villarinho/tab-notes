import { BadgeCheck } from "lucide-react";
import { UserAvatar } from "@/components/user-avatar";
import { Link } from "react-router-dom";

interface ProfileSummaryCardProps {
  user: {
    name: string;
    username: string;
    followers: number;
    following: number;
  };
}

export function ProfileSummaryCard({ user }: ProfileSummaryCardProps) {
  return (
    <div className="rounded-xl border shadow-sm bg-card overflow-hidden text-center min-h-[192px]">
      <div className="h-16 w-full bg-muted" />

      <div className="-mt-8 flex justify-center">
        <Link to={`/${user.username}`}>
          <UserAvatar
            name={user.name}
            username={user.username}
            className="h-12 w-12"
          />
        </Link>
      </div>

      <div className="px-4 pt-2 pb-4 space-y-1">
        <div className="flex justify-center items-center gap-1">
          <Link to={`/${user.username}`}>
            <h3 className="font-semibold text-base hover:underline transition-colors">
              {user.name}
            </h3>
          </Link>
          <BadgeCheck className="w-4 h-4 text-primary" />
        </div>

        <Link to={`/${user.username}`}>
          <p className="text-sm text-muted-foreground hover:underline transition-colors">
            @{user.username}
          </p>
        </Link>

        <div className="pt-2 text-xs text-muted-foreground flex justify-center gap-4">
          <span>
            <strong className="text-foreground">{user.followers}</strong>{" "}
            seguidores
          </span>
          <span>
            <strong className="text-foreground">{user.following}</strong>{" "}
            seguindo
          </span>
        </div>
      </div>
    </div>
  );
}
